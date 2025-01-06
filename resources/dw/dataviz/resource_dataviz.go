// Copyright 2025 Cloudera. All Rights Reserved.
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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
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
	resp.TypeName = req.ProviderTypeName + "_dw_data_visualization"
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
	create, err := r.createDataViz(createRequestFromPlan(ctx, plan))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Data Visualization",
			fmt.Sprintf("Could not create Data Visualization, unexpected error: %v", err),
		)
		return
	}

	clusterID := plan.ClusterID.ValueStringPointer()
	vizID := &create.GetPayload().DataVisualizationID

	// Wait the desired state
	if opts := plan.PollingOptions; !(opts != nil && opts.Async.ValueBool()) {
		if _, err = r.retryStateConf(ctx, setupRetryCfg(clusterID, vizID), &plan).WaitForStateContext(ctx); err != nil {
			resp.Diagnostics.AddError(
				"Error waiting for Data Visualization",
				"Could not create Data Visualization, unexpected error: "+err.Error(),
			)
			return
		}
	}

	// Describe the fresh data
	describe, err := r.describeDataViz(describeRequest(ctx, clusterID, vizID))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Data Visualization",
			fmt.Sprintf("Could not describe Data Visualization, unexpected error: %v", err),
		)
		return
	}

	diags = resp.State.Set(
		ctx,
		stateFromDataViz(
			*clusterID,
			describe.GetPayload().DataVisualization,
			plan.ResourceTemplate,
			time.Now(),
			plan.PollingOptions,
		),
	)
	resp.Diagnostics.Append(diags...)
}

func (r *datavizResource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Warn(ctx, "Read operation is not implemented yet.")
}

func (r *datavizResource) Update(ctx context.Context, _ resource.UpdateRequest, _ *resource.UpdateResponse) {
	tflog.Warn(ctx, "Update operation is not implemented yet.")
}

func (r *datavizResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	clusterID := state.ClusterID.ValueStringPointer()
	vizID := state.ID.ValueStringPointer()

	if _, err := r.deleteDataViz(deleteRequest(ctx, clusterID, vizID)); err != nil {
		if strings.Contains(err.Error(), "unable to get viz-webapp") {
			return
		}
		resp.Diagnostics.AddError(
			"Error deleting Data Visualization",
			fmt.Sprintf("Could not delete Data Visualization, unexpected error: %v", err),
		)
		return
	}

	if opts := state.PollingOptions; !(opts != nil && opts.Async.ValueBool()) {
		if _, err := r.retryStateConf(ctx, teardownRetryCfg(clusterID, vizID), &state).WaitForStateContext(ctx); err != nil {
			resp.Diagnostics.AddError(
				"Error waiting for Data Visualization to delete",
				fmt.Sprintf("Could not delete Data Visualization, unexpected error: %v", err),
			)
			return
		}
	}
}

func (r *datavizResource) createDataViz(p *operations.CreateDataVisualizationParams) (*operations.CreateDataVisualizationOK, error) {
	return r.client.Dw.Operations.CreateDataVisualization(p)
}

func (r *datavizResource) describeDataViz(p *operations.DescribeDataVisualizationParams) (*operations.DescribeDataVisualizationOK, error) {
	return r.client.Dw.Operations.DescribeDataVisualization(p)
}

func (r *datavizResource) deleteDataViz(p *operations.DeleteDataVisualizationParams) (*operations.DeleteDataVisualizationOK, error) {
	return r.client.Dw.Operations.DeleteDataVisualization(p)
}

func createRequestFromPlan(ctx context.Context, plan resourceModel) *operations.CreateDataVisualizationParams {
	return operations.NewCreateDataVisualizationParamsWithContext(ctx).
		WithInput(&models.CreateDataVisualizationRequest{
			ClusterID: plan.ClusterID.ValueStringPointer(),
			Name:      plan.Name.ValueStringPointer(),

			ImageVersion: plan.ImageVersion.ValueString(),
			Config: &models.VizConfig{
				AdminGroups: utils.FromListValueToStringList(plan.AdminGroups),
				UserGroups:  utils.FromListValueToStringList(plan.UserGroups),
			},

			ResourceTemplate: plan.ResourceTemplate.ValueString(),
		})
}

func describeRequest(ctx context.Context, clusterID *string, vizID *string) *operations.DescribeDataVisualizationParams {
	return operations.NewDescribeDataVisualizationParamsWithContext(ctx).
		WithInput(&models.DescribeDataVisualizationRequest{
			ClusterID:           clusterID,
			DataVisualizationID: vizID,
		})
}

func deleteRequest(ctx context.Context, clusterID *string, vizID *string) *operations.DeleteDataVisualizationParams {
	return operations.NewDeleteDataVisualizationParamsWithContext(ctx).
		WithInput(&models.DeleteDataVisualizationRequest{
			ClusterID:           clusterID,
			DataVisualizationID: vizID,
		})
}

func (r *datavizResource) retryStateConf(
	ctx context.Context,
	cfg *retryStateCfg,
	po utils.HasPollingOptions,
) *retry.StateChangeConf {
	failedCnt := 0
	return &retry.StateChangeConf{
		Pending:      cfg.pending,
		Target:       cfg.target,
		Delay:        30 * time.Second,
		Timeout:      utils.GetPollingTimeout(po, 20*time.Minute),
		PollInterval: 30 * time.Second,
		Refresh:      r.stateRefresh(ctx, cfg.clusterID, cfg.vizID, &failedCnt, utils.GetCallFailureThreshold(po, 3)),
	}
}

func (r *datavizResource) stateRefresh(ctx context.Context, clusterID *string, vizID *string, failedCnt *int, failureThreshold int) func() (any, string, error) {
	return func() (any, string, error) {
		tflog.Debug(ctx, "Describing Data Visualisation")

		resp, err := r.describeDataViz(describeRequest(ctx, clusterID, vizID))
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Error describing Data Visualisation, error, %v", err))

			if strings.Contains(err.Error(), "unable to get viz-webapp") {
				// the &models.DescribeDataVisualizationResponse{} has to be a response otherwise it end up in an infinite loop
				return &models.DescribeDataVisualizationResponse{}, "Deleted", nil
			}

			// the "Data Visualization not found" will be the correct way of handling
			if strings.Contains(err.Error(), "Data Visualization not found") {
				return nil, "Deleted", nil
			}

			*failedCnt++
			if *failedCnt <= failureThreshold {
				tflog.Warn(ctx, fmt.Sprintf("could not describe Data Visualization "+
					"due to [%s] but threshold limit is not reached yet (%d out of %d).", err.Error(), failedCnt, failureThreshold))
				return nil, "", nil
			}
			tflog.Error(ctx, fmt.Sprintf("error describing Data Visualization due to [%s] "+
				"failure threshold limit exceeded.", err.Error()))
			return nil, "", err
		}

		*failedCnt = 0
		dataViz := resp.GetPayload().DataVisualization
		tflog.Debug(ctx, fmt.Sprintf("Described Data Visualization %s with status %s", dataViz.ID, dataViz.Status))
		return dataViz, dataViz.Status, nil
	}
}

func stateFromDataViz(
	clusterID string,
	viz *models.DataVisualizationSummary,
	template types.String,
	updated time.Time,
	pollingOpts *utils.PollingOptions,
) resourceModel {
	return resourceModel{
		ID:        types.StringValue(viz.ID),
		ClusterID: types.StringValue(clusterID),
		Name:      types.StringValue(viz.Name),

		ImageVersion:     types.StringValue(viz.ImageVersion),
		ResourceTemplate: template,

		UserGroups:  stringList(viz.UserGroups),
		AdminGroups: stringList(viz.AdminGroups),

		LastUpdated: types.StringValue(updated.Format(time.RFC850)),
		Status:      types.StringValue(viz.Status),

		PollingOptions: pollingOpts,
	}
}

func stringList(s []string) types.List {
	var elems []attr.Value
	elems = make([]attr.Value, 0, len(s))
	for _, v := range s {
		elems = append(elems, types.StringValue(v))
	}
	var list types.List
	list, _ = types.ListValue(types.StringType, elems)
	return list
}

type retryStateCfg struct {
	clusterID *string
	vizID     *string
	pending   []string
	target    []string
}

func setupRetryCfg(clusterID *string, vizID *string) *retryStateCfg {
	return &retryStateCfg{
		clusterID: clusterID,
		vizID:     vizID,
		pending:   []string{"Accepted", "Creating", "Created", "Starting"},
		target:    []string{"Running"},
	}
}

func teardownRetryCfg(clusterID *string, vizID *string) *retryStateCfg {
	return &retryStateCfg{
		clusterID: clusterID,
		vizID:     vizID,
		pending:   []string{"Deleting", "Running", "Stopping", "Stopped", "Creating", "Created", "Starting", "Updating"},
		target:    []string{"Deleted"},
	}
}
