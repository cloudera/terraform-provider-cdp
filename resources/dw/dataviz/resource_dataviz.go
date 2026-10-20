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

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
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

var (
	_ resource.Resource                     = (*datavizResource)(nil)
	_ resource.ResourceWithConfigure        = (*datavizResource)(nil)
	_ resource.ResourceWithImportState      = (*datavizResource)(nil)
	_ resource.ResourceWithValidateConfig   = (*datavizResource)(nil)
	_ resource.ResourceWithConfigValidators = (*datavizResource)(nil)
)

const (
	resourceTemplateDefault string = "viz-default"
	resourceTemplateLow     string = "viz-low"
	resourceTemplateMedium  string = "viz-medium"
	resourceTemplateLarge   string = "viz-large"
)

func NewDataVizResource() resource.Resource {
	return &datavizResource{}
}

func (r *datavizResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *datavizResource) ConfigValidators(context.Context) []resource.ConfigValidator {
	return utils.PollingOptionsConflictValidators()
}

func (r *datavizResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data resourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if admGrp := data.AdminGroups; len(admGrp.Elements()) == 0 {
		resp.Diagnostics.AddAttributeError(
			path.Root("admin_groups"),
			"Invalid administrator groups",
			"The admin_groups must have at least one group.",
		)
	}

	// If resource template not defined then default will be applied -> viz-default, otherwise we validate
	if !data.ResourceTemplate.IsNull() && data.ResourceTemplate.ValueString() != "" {
		// oly allow certain values for the resource template field
		resourceTemplatePossibleValues := []string{
			resourceTemplateDefault,
			resourceTemplateLow,
			resourceTemplateMedium,
			resourceTemplateLarge,
		}
		if !isAnyOf(data.ResourceTemplate.ValueString(), resourceTemplatePossibleValues...) {
			resp.Diagnostics.AddAttributeError(
				path.Root("resource_template"),
				"Invalid resource template",
				fmt.Sprintf("The resource_template can be one of the following if defined: %s.", strings.Join(resourceTemplatePossibleValues, ", ")),
			)
		}
	}
}

func (r *datavizResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *datavizResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dw_data_visualization"
}

func (r *datavizResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	datavizSchema.Attributes["timeouts"] = timeouts.Attributes(ctx, timeouts.Opts{Create: true, Delete: true})
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
	create, err := r.createDataViz(ctx, createRequestFromPlan(ctx, plan))
	if err != nil {
		utils.AddDwDiagnosticsError(err, &resp.Diagnostics, "create Data Visualization")
		return
	}

	clusterID := plan.ClusterID.ValueStringPointer()
	vizID := &create.GetPayload().DataVisualizationID

	// Wait the desired state
	if !utils.ResolveAsync(plan.PollingOptions, plan.Async) {
		createTimeout, diagsTimeout := utils.ResolveTimeout(ctx, plan.PollingOptions, plan.Timeouts.Create, 20*time.Minute)
		resp.Diagnostics.Append(diagsTimeout...)
		if resp.Diagnostics.HasError() {
			return
		}
		if _, err = r.retryStateConf(ctx, setupRetryCfg(clusterID, vizID), createTimeout, utils.ResolveCallFailureThreshold(plan.PollingOptions, plan.CallFailureThreshold, 3)).WaitForStateContext(ctx); err != nil {
			resp.Diagnostics.AddError(
				"Error waiting for Data Visualization",
				"Could not create Data Visualization, unexpected error: "+err.Error(),
			)
			return
		}
	}

	// Describe the fresh data
	describe, err := r.describeDataViz(ctx, describeRequest(ctx, clusterID, vizID))
	if err != nil {
		utils.AddDwDiagnosticsError(err, &resp.Diagnostics, "describe Data Visualization")
		return
	}

	diags = resp.State.Set(
		ctx,
		stateFromDataViz(
			*clusterID,
			describe.GetPayload().DataVisualization,
			plan.ResourceTemplate.ValueStringPointer(),
			time.Now(),
			plan.PollingOptions,
			plan.Async,
			plan.CallFailureThreshold,
			plan.Timeouts,
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

	if _, err := r.deleteDataViz(ctx, deleteRequest(ctx, clusterID, vizID)); err != nil {
		if strings.Contains(err.Error(), "unable to get viz-webapp") {
			return
		}
		utils.AddDwDiagnosticsError(err, &resp.Diagnostics, "delete Data Visualization")
		return
	}

	if !utils.ResolveAsync(state.PollingOptions, state.Async) {
		deleteTimeout, diagsTimeout := utils.ResolveTimeout(ctx, state.PollingOptions, state.Timeouts.Delete, 20*time.Minute)
		resp.Diagnostics.Append(diagsTimeout...)
		if resp.Diagnostics.HasError() {
			return
		}
		if _, err := r.retryStateConf(ctx, teardownRetryCfg(clusterID, vizID), deleteTimeout, utils.ResolveCallFailureThreshold(state.PollingOptions, state.CallFailureThreshold, 3)).WaitForStateContext(ctx); err != nil {
			resp.Diagnostics.AddError(
				"Error waiting for Data Visualization to delete",
				fmt.Sprintf("Could not delete Data Visualization, unexpected error: %v", err),
			)
			return
		}
	}
}

func (r *datavizResource) createDataViz(ctx context.Context, p *operations.CreateDataVisualizationParams) (*operations.CreateDataVisualizationOK, error) {
	return r.client.Dw.Operations.CreateDataVisualizationContext(ctx, p)
}

func (r *datavizResource) describeDataViz(ctx context.Context, p *operations.DescribeDataVisualizationParams) (*operations.DescribeDataVisualizationOK, error) {
	return r.client.Dw.Operations.DescribeDataVisualizationContext(ctx, p)
}

func (r *datavizResource) deleteDataViz(ctx context.Context, p *operations.DeleteDataVisualizationParams) (*operations.DeleteDataVisualizationOK, error) {
	return r.client.Dw.Operations.DeleteDataVisualizationContext(ctx, p)
}

func createRequestFromPlan(ctx context.Context, plan resourceModel) *operations.CreateDataVisualizationParams {
	return operations.NewCreateDataVisualizationParams().
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
	return operations.NewDescribeDataVisualizationParams().
		WithInput(&models.DescribeDataVisualizationRequest{
			ClusterID:           clusterID,
			DataVisualizationID: vizID,
		})
}

func deleteRequest(ctx context.Context, clusterID *string, vizID *string) *operations.DeleteDataVisualizationParams {
	return operations.NewDeleteDataVisualizationParams().
		WithInput(&models.DeleteDataVisualizationRequest{
			ClusterID:           clusterID,
			DataVisualizationID: vizID,
		})
}

func (r *datavizResource) retryStateConf(
	ctx context.Context,
	cfg *retryStateCfg,
	timeout time.Duration,
	callFailureThreshold int,
) *retry.StateChangeConf {
	failedCnt := 0
	return &retry.StateChangeConf{
		Pending:      cfg.pending,
		Target:       cfg.target,
		Delay:        30 * time.Second,
		Timeout:      timeout,
		PollInterval: 10 * time.Second,
		Refresh:      r.stateRefresh(ctx, cfg.clusterID, cfg.vizID, &failedCnt, callFailureThreshold),
	}
}

func (r *datavizResource) stateRefresh(ctx context.Context, clusterID *string, vizID *string, failedCnt *int, failureThreshold int) func() (any, string, error) {
	return func() (any, string, error) {
		tflog.Debug(ctx, "Describing Data Visualisation")

		resp, err := r.describeDataViz(ctx, describeRequest(ctx, clusterID, vizID))
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Error describing Data Visualisation, error, %v", err))

			if strings.Contains(err.Error(), "unable to get viz-webapp") {
				// the &models.DescribeDataVisualizationResponse{} has to be a response otherwise it end up in an infinite loop
				return &models.DescribeDataVisualizationResponse{}, "Deleted", nil
			}

			// the "Data Visualization not found" will be the correct way of handling
			if strings.Contains(err.Error(), "Data Visualization not found") {
				// the &models.DescribeDataVisualizationResponse{} has to be a response otherwise it end up in an infinite loop
				return &models.DescribeDataVisualizationResponse{}, "Deleted", nil
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
	template *string,
	updated time.Time,
	pollingOptions *utils.PollingOptions,
	async types.Bool,
	callFailureThreshold types.Int64,
	timeoutsVal timeouts.Value,
) resourceModel {
	t := types.StringNull()
	if template != nil {
		t = types.StringValue(*template)
	}

	return resourceModel{
		ID:        types.StringValue(viz.ID),
		ClusterID: types.StringValue(clusterID),
		Name:      types.StringValue(viz.Name),

		ImageVersion:     types.StringValue(viz.ImageVersion),
		ResourceTemplate: t,

		UserGroups:  utils.FromStringListToListValue(viz.UserGroups),
		AdminGroups: utils.FromStringListToListValue(viz.AdminGroups),

		LastUpdated:          types.StringValue(updated.Format(time.RFC850)),
		Status:               types.StringValue(viz.Status),
		PollingOptions:       pollingOptions,
		Async:                async,
		CallFailureThreshold: callFailureThreshold,
		Timeouts:             timeoutsVal,
	}
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

func isAnyOf(v string, possibleValues ...string) bool {
	for _, pv := range possibleValues {
		if v == pv {
			return true
		}
	}
	return false
}
