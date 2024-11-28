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

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

	// Generate API request body from plan
	vw := operations.NewCreateVwParamsWithContext(ctx).
		WithInput(&models.CreateVwRequest{
			Name:      plan.Name.ValueStringPointer(),
			ClusterID: plan.ClusterID.ValueStringPointer(),
			DbcID:     plan.DatabaseCatalogID.ValueStringPointer(),
			VwType:    models.VwTypeImpala.Pointer(),
		})

	// Create new virtual warehouse
	response, err := r.client.Dw.Operations.CreateVw(vw)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Impala virtual warehouse",
			"Could not create Impala, unexpected error: "+err.Error(),
		)
		return
	}

	payload := response.GetPayload()
	clusterID := plan.ClusterID.ValueStringPointer()
	vwID := &payload.VwID

	if opts := plan.PollingOptions; !(opts != nil && opts.Async.ValueBool()) {
		callFailedCount := 0
		stateConf := &retry.StateChangeConf{
			Pending:      []string{"Accepted", "Creating", "Created", "Starting"},
			Target:       []string{"Running"},
			Delay:        30 * time.Second,
			Timeout:      utils.GetPollingTimeout(&plan, 20*time.Minute),
			PollInterval: 30 * time.Second,
			Refresh:      r.stateRefresh(ctx, clusterID, vwID, &callFailedCount, utils.GetCallFailureThreshold(&plan, 3)),
		}
		if _, err = stateConf.WaitForStateContext(ctx); err != nil {
			resp.Diagnostics.AddError(
				"Error waiting for Data Warehouse Impala virtual warehouse",
				"Could not create Impala, unexpected error: "+err.Error(),
			)
			return
		}
	}
	desc := operations.NewDescribeVwParamsWithContext(ctx).
		WithInput(&models.DescribeVwRequest{VwID: vwID, ClusterID: clusterID})
	describe, err := r.client.Dw.Operations.DescribeVw(desc)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Impala virtual warehouse",
			"Could not describe Impala, unexpected error: "+err.Error(),
		)
		return
	}

	impala := describe.GetPayload()
	plan.ID = types.StringValue(impala.Vw.ID)
	plan.DatabaseCatalogID = types.StringValue(impala.Vw.DbcID)
	plan.Name = types.StringValue(impala.Vw.Name)
	plan.Status = types.StringValue(impala.Vw.Status)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
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
	op := operations.NewDeleteVwParamsWithContext(ctx).
		WithInput(&models.DeleteVwRequest{
			ClusterID: clusterID,
			VwID:      vwID,
		})

	if _, err := r.client.Dw.Operations.DeleteVw(op); err != nil {
		if strings.Contains(err.Error(), "Virtual Warehouse not found") {
			return
		}
		resp.Diagnostics.AddError(
			"Error deleting Impala Virtual Warehouse",
			"Could not delete Impala Virtual Warehouse, unexpected error: "+err.Error(),
		)
		return
	}

	if opts := state.PollingOptions; !(opts != nil && opts.Async.ValueBool()) {
		callFailedCount := 0
		stateConf := &retry.StateChangeConf{
			Pending:      []string{"Deleting", "Running", "Stopping", "Stopped", "Creating", "Created", "Starting", "Updating"},
			Target:       []string{"Deleted"},
			Delay:        30 * time.Second,
			Timeout:      utils.GetPollingTimeout(&state, 20*time.Minute),
			PollInterval: 30 * time.Second,
			Refresh:      r.stateRefresh(ctx, clusterID, vwID, &callFailedCount, utils.GetCallFailureThreshold(&state, 3)),
		}
		if _, err := stateConf.WaitForStateContext(ctx); err != nil {
			resp.Diagnostics.AddError(
				"Error waiting for Data Warehouse Impala Virtual Warehouse",
				"Could not delete Impala, unexpected error: "+err.Error(),
			)
			return
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
