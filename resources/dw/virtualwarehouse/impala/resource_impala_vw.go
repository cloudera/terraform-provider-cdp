// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package impala

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

type impalaResource struct {
	client *cdp.Client
}

var (
	_ resource.Resource              = (*impalaResource)(nil)
	_ resource.ResourceWithConfigure = (*impalaResource)(nil)
)

func NewImpalaResource() resource.Resource {
	return &impalaResource{}
}

func (r *impalaResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *impalaResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dw_vw_impala"
}

func (r *impalaResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = impalaSchema
}

func (r *impalaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan resourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create the VW Request using a helper
	vwhCreateRequest, errDiag := r.createVwRequestFromPlan(&plan, ctx)
	tflog.Debug(ctx, fmt.Sprintf("CreateVw request: %+v", vwhCreateRequest))
	resp.Diagnostics.Append(errDiag...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Make API request to create VW
	response, err := r.client.Dw.Operations.CreateVw(
		operations.NewCreateVwParamsWithContext(ctx).WithInput(vwhCreateRequest),
	)
	if err != nil || response.GetPayload() == nil {
		resp.Diagnostics.AddError(
			"Error creating Impala virtual warehouse",
			fmt.Sprintf("Could not create Impala, unexpected error: %v", err),
		)
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("CreateVw response: %+v", response.GetPayload()))

	// Wait for the VW to reach Running state
	if err := r.waitForVwRunning(ctx, &plan, &response.GetPayload().VwID); err != nil {
		resp.Diagnostics.AddError(
			"Error waiting for Data Warehouse Impala virtual warehouse",
			fmt.Sprintf("Could not create Impala, unexpected error: %v", err),
		)
		return
	}

	// Fetch and map the response to plan using a helper
	if err := r.populatePlanFromDescribe(ctx, &plan, &response.GetPayload().VwID); err != nil {
		resp.Diagnostics.AddError(
			"Error creating Impala virtual warehouse",
			fmt.Sprintf("Could not describe Impala, unexpected error: %v", err),
		)
		return
	}

	// Save the updated plan into state
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *impalaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Warn(ctx, "Read operation is not implemented yet.")
}

func (r *impalaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Warn(ctx, "Update operation is not implemented yet.")
}

func (r *impalaResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	clusterID := state.ClusterID.ValueStringPointer()
	vwID := state.ID.ValueStringPointer()

	if clusterID == nil || vwID == nil {
		resp.Diagnostics.AddError(
			"Invalid State",
			"ClusterID or Virtual Warehouse ID is missing.",
		)
		return
	}

	err := r.deleteVirtualWarehouse(ctx, clusterID, vwID)
	if err != nil {
		if strings.Contains(err.Error(), "Virtual Warehouse not found") {
			return
		}
		resp.Diagnostics.AddError(
			"Error deleting Impala Virtual Warehouse",
			"Could not delete Impala Virtual Warehouse, unexpected error: "+err.Error(),
		)
		return
	}

	// Handle polling for not async call
	if opts := state.PollingOptions; opts == nil || !opts.Async.ValueBool() {
		err = r.pollForDeletion(ctx, state, clusterID, vwID)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error waiting for Impala Virtual Warehouse deletion",
				"Could not delete Impala Virtual Warehouse, unexpected error: "+err.Error(),
			)
		}
	}
}

func (r *impalaResource) stateRefresh(ctx context.Context, clusterID *string, vwID *string, callFailedCount *int, callFailureThreshold int) func() (any, string, error) {
	return func() (any, string, error) {
		tflog.Debug(ctx, "About to describe Impala")
		params := operations.NewDescribeVwParamsWithContext(ctx).
			WithInput(&models.DescribeVwRequest{ClusterID: clusterID, VwID: vwID})
		resp, err := r.client.Dw.Operations.DescribeVw(params)
		if err != nil {
			if strings.Contains(err.Error(), "Virtual Warehouse not found") {
				return &models.DescribeVwResponse{}, "Deleted", nil
			}
			*callFailedCount++
			if *callFailedCount <= callFailureThreshold {
				tflog.Warn(ctx, fmt.Sprintf("could not describe Data Warehouse Impala Virtual Warehouse "+
					"due to [%s] but threshold limit is not reached yet (%d out of %d).", err.Error(), callFailedCount, callFailureThreshold))
				return nil, "", nil
			}
			tflog.Error(ctx, fmt.Sprintf("error describing Data Warehouse Impala Virtual Warehouse due to [%s] "+
				"failure threshold limit exceeded.", err.Error()))
			return nil, "", err
		}
		*callFailedCount = 0
		vw := resp.GetPayload()
		tflog.Debug(ctx, fmt.Sprintf("Described Impala %s with status %s", vw.Vw.ID, vw.Vw.Status))
		return vw, vw.Vw.Status, nil
	}
}

func (r *impalaResource) createVwRequestFromPlan(plan *resourceModel, ctx context.Context) (*models.CreateVwRequest, diag.Diagnostics) {
	var diagErrors diag.Diagnostics
	req := &models.CreateVwRequest{
		Name:      plan.Name.ValueStringPointer(),
		ClusterID: plan.ClusterID.ValueStringPointer(),
		DbcID:     plan.DatabaseCatalogID.ValueStringPointer(),
		VwType:    models.VwTypeImpala.Pointer(),
	}

	setIfNotEmpty := func(target *string, source string) {
		if source != "" {
			*target = source
		}
	}

	setIfPositive := func(target *int32, source int32) {
		if source > 0 {
			*target = source
		}
	}

	setIfTrue := func(target **bool, source bool) {
		if source {
			*target = &source
		}
	}

	setIfNotEmpty(&req.ImageVersion, plan.ImageVersion.ValueString())
	setIfNotEmpty(&req.InstanceType, plan.InstanceType.ValueString())
	setIfNotEmpty(&req.TShirtSize, plan.TShirtSize.ValueString())
	setIfNotEmpty(&req.AvailabilityZone, plan.AvailabilityZone.ValueString())

	setIfPositive(&req.NodeCount, plan.NodeCount.ValueInt32())

	req.EnableUnifiedAnalytics = plan.EnableUnifiedAnalytics.ValueBool()
	req.ImpalaQueryLog = plan.ImpalaQueryLog.ValueBool()
	setIfTrue(&req.PlatformJwtAuth, plan.PlatformJwtAuth.ValueBool())

	var err error
	if !plan.ImpalaOptions.IsNull() {
		req.ImpalaOptions, err = convertToAPIImpalaOptions(plan.ImpalaOptions, ctx)
		if err != nil {
			diagErrors.AddError("failed to convert Impala Options",
				err.Error())
		}
	}
	if !plan.Autoscaling.IsNull() {
		req.Autoscaling, err = convertToAPIAutoscaling(plan.Autoscaling, ctx)
		if err != nil {
			diagErrors.AddError("failed to convert API Autoscaling",
				err.Error())
		}
	}
	if !plan.ImpalaHASettings.IsNull() {
		req.ImpalaHaSettings, err = convertToAPIImpalaHASettings(plan.ImpalaHASettings, ctx)
		if err != nil {
			diagErrors.AddError("failed to convert Impala HA Settings",
				err.Error())
		}
	}
	if !plan.QueryIsolationOptions.IsNull() {
		req.QueryIsolationOptions = convertToAPIQueryIsolationOptions(plan.QueryIsolationOptions, ctx)
	}

	if !plan.EnableSSO.IsNull() && !plan.EnableSSO.IsUnknown() {
		boolValue, err1 := plan.EnableSSO.ToBoolValue(ctx)
		if err != nil {
			diagErrors.AddError("failed to convert Enable SSo",
				err.Error())
			return req, err1
		}
		tflog.Debug(ctx, fmt.Sprintf("Assigned value to EnableSSO: %v", boolValue))
		req.Config = &models.ServiceConfigReq{
			EnableSSO: boolValue.ValueBool(),
		}
	}

	if len(plan.Tags.Elements()) > 0 {
		req.Tags = convertToAPITagRequests(plan.Tags)
	}
	return req, diagErrors
}

func (r *impalaResource) waitForVwRunning(ctx context.Context, plan *resourceModel, vwID *string) error {
	clusterID := plan.ClusterID.ValueStringPointer()

	if opts := plan.PollingOptions; opts == nil || !opts.Async.ValueBool() {
		callFailedCount := 0
		stateConf := &retry.StateChangeConf{
			Pending:      []string{"Accepted", "Creating", "Created", "Starting"},
			Target:       []string{"Running"},
			Delay:        30 * time.Second,
			Timeout:      utils.GetPollingTimeout(plan, 20*time.Minute),
			PollInterval: 30 * time.Second,
			Refresh:      r.stateRefresh(ctx, clusterID, vwID, &callFailedCount, utils.GetCallFailureThreshold(plan, 3)),
		}
		_, err := stateConf.WaitForStateContext(ctx)
		return err
	}
	return nil
}

func (r *impalaResource) populatePlanFromDescribe(ctx context.Context, plan *resourceModel, vwID *string) error {
	desc := operations.NewDescribeVwParamsWithContext(ctx).
		WithInput(&models.DescribeVwRequest{VwID: vwID, ClusterID: plan.ClusterID.ValueStringPointer()})
	describe, err := r.client.Dw.Operations.DescribeVw(desc)
	if err != nil {
		return err
	}

	impala := describe.GetPayload()
	tflog.Info(context.Background(), fmt.Sprintf("API Response: %+v", impala))
	plan.setFromDescribeVwResponse(impala, ctx)

	return nil
}

func (r *impalaResource) deleteVirtualWarehouse(ctx context.Context, clusterID, vwID *string) error {
	op := operations.NewDeleteVwParamsWithContext(ctx).
		WithInput(&models.DeleteVwRequest{
			ClusterID: clusterID,
			VwID:      vwID,
		})

	_, err := r.client.Dw.Operations.DeleteVw(op)
	return err
}

func (r *impalaResource) pollForDeletion(ctx context.Context, state resourceModel, clusterID, vwID *string) error {
	callFailedCount := 0
	stateConf := &retry.StateChangeConf{
		Pending:      []string{"Deleting", "Running", "Stopping", "Stopped", "Creating", "Created", "Starting", "Updating"},
		Target:       []string{"Deleted"},
		Delay:        30 * time.Second,
		Timeout:      utils.GetPollingTimeout(&state, 20*time.Minute),
		PollInterval: 30 * time.Second,
		Refresh:      r.stateRefresh(ctx, clusterID, vwID, &callFailedCount, utils.GetCallFailureThreshold(&state, 3)),
	}

	_, err := stateConf.WaitForStateContext(ctx)
	return err
}
