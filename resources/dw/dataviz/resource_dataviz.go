// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package dataviz

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

type datavizResource struct {
	client *cdp.Client
}

func (r *datavizResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

var (
	_ resource.Resource                = (*datavizResource)(nil)
	_ resource.ResourceWithConfigure   = (*datavizResource)(nil)
	_ resource.ResourceWithImportState = (*datavizResource)(nil)
)

func NewDataVizResource() resource.Resource {
	return &datavizResource{}
}

func (r *datavizResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *datavizResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dw_dataviz"
}

func (r *datavizResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = datavizSchema
}

func (r *datavizResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan resourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create new Data Visualization
	response, err := r.client.Dw.Operations.CreateDataVisualization(
		requestFromPlan(ctx, plan))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Data Visualization",
			fmt.Sprintf("Could not create Data Visualization, unexpected error: %v", err),
		)
		return
	}

	payload := response.GetPayload()
	clusterID := plan.ClusterID.ValueStringPointer()
	vizID := &payload.DataVisualizationID

	if opts := plan.PollingOptions; !(opts != nil && opts.Async.ValueBool()) {
		callFailedCount := 0
		stateConf := &retry.StateChangeConf{
			Pending:      []string{"Accepted", "Creating", "Created", "Starting"},
			Target:       []string{"Running"},
			Delay:        30 * time.Second,
			Timeout:      utils.GetPollingTimeout(&plan, 20*time.Minute),
			PollInterval: 30 * time.Second,
			Refresh:      r.stateRefresh(ctx, clusterID, vizID, &callFailedCount, utils.GetCallFailureThreshold(&plan, 3)),
		}
		if _, err = stateConf.WaitForStateContext(ctx); err != nil {
			resp.Diagnostics.AddError(
				"Error waiting for Data Warehouse hive virtual warehouse",
				"Could not create hive, unexpected error: "+err.Error(),
			)
			return
		}
	}
	desc := operations.NewDescribeDataVisualizationParamsWithContext(ctx).
		WithInput(&models.DescribeDataVisualizationRequest{DataVisualizationID: vizID, ClusterID: clusterID})
	describe, err := r.client.Dw.Operations.DescribeDataVisualization(desc)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Data Visualization",
			fmt.Sprintf("Could not describe Data Visualization, unexpected error: %v", err),
		)
		return
	}

	viz := describe.GetPayload()
	plan.ID = types.StringValue(viz.DataVisualization.ID)
	plan.Name = types.StringValue(viz.DataVisualization.Name)
	plan.Status = types.StringValue(viz.DataVisualization.Status)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *datavizResource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Warn(ctx, "Read operation is not implemented yet.")
}

func (r *datavizResource) Update(ctx context.Context, _ resource.UpdateRequest, _ *resource.UpdateResponse) {
	tflog.Warn(ctx, "Update operation is not implemented yet.")
}

func (r *datavizResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	//		var state resourceModel
	//
	//		resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	//		if resp.Diagnostics.HasError() {
	//			return
	//		}
	//
	//		clusterID := state.ClusterID.ValueStringPointer()
	//		vwID := state.ID.ValueStringPointer()
	//		op := operations.NewDeleteVwParamsWithContext(ctx).
	//			WithInput(&models.DeleteVwRequest{
	//				ClusterID: clusterID,
	//				VwID:      vwID,
	//			})
	//
	//		if _, err := r.client.Dw.Operations.DeleteVw(op); err != nil {
	//			if strings.Contains(err.Error(), "Virtual Warehouse not found") {
	//				return
	//			}
	//			resp.Diagnostics.AddError(
	//				"Error deleting Hive Virtual Warehouse",
	//				"Could not delete Hive Virtual Warehouse, unexpected error: "+err.Error(),
	//			)
	//			return
	//		}
	//
	//		if opts := state.PollingOptions; !(opts != nil && opts.Async.ValueBool()) {
	//			callFailedCount := 0
	//			stateConf := &retry.StateChangeConf{
	//				Pending:      []string{"Deleting", "Running", "Stopping", "Stopped", "Creating", "Created", "Starting", "Updating"},
	//				Target:       []string{"Deleted"}, // This is not an actual state, we added it to fake the state change
	//				Delay:        30 * time.Second,
	//				Timeout:      utils.GetPollingTimeout(&state, 20*time.Minute),
	//				PollInterval: 30 * time.Second,
	//				Refresh:      r.stateRefresh(ctx, clusterID, vwID, &callFailedCount, utils.GetCallFailureThreshold(&state, 3)),
	//			}
	//			if _, err := stateConf.WaitForStateContext(ctx); err != nil {
	//				resp.Diagnostics.AddError(
	//					"Error waiting for Data Warehouse Hive Virtual Warehouse",
	//					"Could not delete hive, unexpected error: "+err.Error(),
	//				)
	//				return
	//			}
	//		}
}

func (r *datavizResource) stateRefresh(context.Context, *string, *string, *int, int) func() (any, string, error) {
	//func (r *datavizResource) stateRefresh(ctx context.Context, clusterID *string, vwID *string, callFailedCount *int, callFailureThreshold int) func() (any, string, error) {
	return func() (any, string, error) {
		//tflog.Debug(ctx, "About to describe hive")
		//params := operations.NewDescribeVwParamsWithContext(ctx).
		//	WithInput(&models.DescribeVwRequest{ClusterID: clusterID, VwID: vwID})
		//resp, err := r.client.Dw.Operations.DescribeVw(params)
		//if err != nil {
		//	if strings.Contains(err.Error(), "Virtual Warehouse not found") {
		//		return &models.DescribeVwResponse{}, "Deleted", nil
		//	}
		//	*callFailedCount++
		//	if *callFailedCount <= callFailureThreshold {
		//		tflog.Warn(ctx, fmt.Sprintf("could not describe Data Warehouse Hive Virtual Warehouse "+
		//			"due to [%s] but threshold limit is not reached yet (%d out of %d).", err.Error(), callFailedCount, callFailureThreshold))
		//		return nil, "", nil
		//	}
		//	tflog.Error(ctx, fmt.Sprintf("error describing Data Warehouse Hive Virtual Warehouse due to [%s] "+
		//		"failure threshold limit exceeded.", err.Error()))
		//	return nil, "", err
		//}
		//*callFailedCount = 0
		//vw := resp.GetPayload()
		//tflog.Debug(ctx, fmt.Sprintf("Described Hive %s with status %s", vw.Vw.ID, vw.Vw.Status))
		//return vw, vw.Vw.Status, nil
		return nil, "", nil
	}
}

func requestFromPlan(ctx context.Context, plan resourceModel) *operations.CreateDataVisualizationParams {
	// Generate API request body from plan
	input := &models.CreateDataVisualizationRequest{
		ClusterID: plan.ClusterID.ValueStringPointer(),
		Name:      plan.Name.ValueStringPointer(),

		ImageVersion: plan.getImageVersion(),
		Config: &models.VizConfig{
			AdminGroups: utils.FromListValueToStringList(plan.AdminGroups),
			UserGroups:  utils.FromListValueToStringList(plan.UserGroups),
		},

		ResourceTemplate: plan.ResourceTemplate.ValueString(),
	}
	return operations.NewCreateDataVisualizationParamsWithContext(ctx).
		WithInput(input)
}
