// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package aws

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
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

type dwClusterResource struct {
	client *cdp.Client
}

func (r *dwClusterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

var (
	_ resource.Resource                = (*dwClusterResource)(nil)
	_ resource.ResourceWithConfigure   = (*dwClusterResource)(nil)
	_ resource.ResourceWithImportState = (*dwClusterResource)(nil)
)

func NewDwClusterResource() resource.Resource {
	return &dwClusterResource{}
}

func (r *dwClusterResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *dwClusterResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dw_aws_cluster"
}

func (r *dwClusterResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = dwClusterSchema
}

func (r *dwClusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan resourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, diags := r.createCluster(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := response.GetPayload()
	clusterID := &payload.ClusterID
	plan.ClusterID = types.StringValue(*clusterID)

	describe, err := r.describeCluster(ctx, clusterID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Data Warehouse AWS cluster",
			"Could not describe cluster, unexpected error: "+err.Error(),
		)
		return
	}
	plan.setResourceModel(ctx, describe.GetPayload())
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if opts := plan.PollingOptions; opts == nil || opts.Async.ValueBool() {
		cluster, diags := r.waitForClusterCreation(ctx, &plan, clusterID)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		plan.setResourceModel(ctx, cluster)
		diags = resp.State.Set(ctx, plan)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		catalog, diags := r.waitForDefaultDatabaseCatalogCreation(ctx, &plan, clusterID)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		diags = plan.setDefaultDatabaseCatalog(catalog)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *dwClusterResource) waitForClusterCreation(ctx context.Context, plan *resourceModel, clusterID *string) (*models.DescribeClusterResponse, diag.Diagnostics) {
	diags := diag.Diagnostics{}
	if _, err := retryStateConf(ctx, plan, setupRetryCfg(clusterID), r.clusterStateRefresh).WaitForStateContext(ctx); err != nil {
		diags.AddError(
			"Error waiting for Data Warehouse AWS cluster",
			"Could not create cluster, unexpected error: "+err.Error(),
		)
		return nil, diags
	}

	describe, err := r.describeCluster(ctx, clusterID)
	if err != nil {
		diags.AddError(
			"Error creating Data Warehouse AWS cluster",
			"Could not describe cluster, unexpected error: "+err.Error(),
		)
		return nil, diags
	}

	return describe.GetPayload(), diags
}

func (r *dwClusterResource) waitForDefaultDatabaseCatalogCreation(ctx context.Context, plan *resourceModel, clusterID *string) (*models.DbcSummary, diag.Diagnostics) {
	diags := diag.Diagnostics{}
	if _, err := retryStateConf(ctx, plan, setupRetryCfg(clusterID), r.databaseCatalogStateRefresh).WaitForStateContext(ctx); err != nil {
		diags.AddError(
			"Error waiting for Data Warehouse database catalog",
			fmt.Sprintf("Could not create database catalog, unexpected error: %v", err),
		)
		return nil, diags
	}
	catalog, err := r.getDatabaseCatalog(ctx, clusterID)
	if err != nil {
		diags.AddError(
			"Error finding Data Warehouse database catalog", fmt.Sprintf("unexpected error: %v", err),
		)
		return nil, diags
	}

	return catalog, diags
}

func (r *dwClusterResource) createCluster(ctx context.Context, plan resourceModel) (*operations.CreateAwsClusterOK, diag.Diagnostics) {
	// Generate API request body from plan
	req, diags := plan.convertToCreateAwsClusterRequest(ctx)
	if diags.HasError() {
		return nil, diags
	}
	clusterParams := operations.NewCreateAwsClusterParamsWithContext(ctx).
		WithInput(req)

	// Create new AWS cluster
	response, err := r.client.Dw.Operations.CreateAwsCluster(clusterParams)
	if err != nil {
		diags.AddError("Error creating Data Warehouse AWS cluster", "Could not create cluster, unexpected error: "+err.Error())
	}
	return response, diags
}

func (r *dwClusterResource) describeCluster(ctx context.Context, clusterID *string) (*operations.DescribeClusterOK, error) {
	desc := operations.NewDescribeClusterParamsWithContext(ctx).
		WithInput(&models.DescribeClusterRequest{ClusterID: clusterID})
	describe, err := r.client.Dw.Operations.DescribeCluster(desc)
	return describe, err
}

func (r *dwClusterResource) deleteCluster(ctx context.Context, clusterID *string) (*operations.DeleteClusterOK, error) {
	op := operations.NewDeleteClusterParamsWithContext(ctx).
		WithInput(&models.DeleteClusterRequest{
			ClusterID: clusterID,
		})
	resp, err := r.client.Dw.Operations.DeleteCluster(op)
	return resp, err
}

func (r *dwClusterResource) getDatabaseCatalog(ctx context.Context, clusterID *string) (*models.DbcSummary, error) {
	response, err := r.listDatabaseCatalogs(ctx, clusterID)
	if err != nil {
		err = fmt.Errorf("could not list database catalogs, unexpected error: %s", err.Error())
		return nil, err
	}
	resp := response.GetPayload()
	if len(resp.Dbcs) != 1 {
		err = fmt.Errorf("exactly one Data Warehouse database catalog should be deployed for cluster %s", *clusterID)
		return nil, err
	}
	return resp.Dbcs[0], nil
}

func (r *dwClusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	describe, err := r.describeCluster(ctx, state.ClusterID.ValueStringPointer())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Data Warehouse AWS cluster",
			"Could not describe cluster, unexpected error: "+err.Error(),
		)
		return
	}
	diags = state.setResourceModel(ctx, describe.GetPayload())
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *dwClusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Warn(ctx, "Update operation is not implemented yet.")
}

func (r *dwClusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resourceModel

	// Read Terraform prior state into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	clusterID := state.ClusterID.ValueStringPointer()
	if _, err := r.deleteCluster(ctx, clusterID); err != nil {
		resp.Diagnostics.AddError(
			"Error deleting Data Warehouse AWS cluster",
			"Could not delete cluster, unexpected error: "+err.Error(),
		)
		return
	}

	if opts := state.PollingOptions; opts == nil || opts.Async.ValueBool() {
		if _, err := retryStateConf(ctx, &state, teardownRetryCfg(clusterID), r.clusterStateRefresh).WaitForStateContext(ctx); err != nil {
			resp.Diagnostics.AddError(
				"Error waiting for Data Warehouse AWS cluster",
				"Could not delete cluster, unexpected error: "+err.Error(),
			)
			return
		}
	}
}

type retryStateCfg struct {
	clusterID *string
	pending   []string
	target    []string
}

func setupRetryCfg(clusterID *string) *retryStateCfg {
	return &retryStateCfg{
		clusterID: clusterID,
		pending:   []string{"Accepted", "Creating", "Created", "Loading", "Starting"},
		target:    []string{"Running"},
	}
}

func teardownRetryCfg(clusterID *string) *retryStateCfg {
	return &retryStateCfg{
		clusterID: clusterID,
		pending:   []string{"Deleting", "Running", "Stopping", "Stopped", "Creating", "Created", "Starting", "Updating"},
		target:    []string{"Deleted"},
	}
}

type stateRefresherFunc func(
	ctx context.Context,
	clusterID *string,
	callFailedCount *int,
	callFailureThreshold int) func() (any, string, error)

func retryStateConf(
	ctx context.Context,
	po utils.HasPollingOptions,
	status *retryStateCfg,
	stateRefresher stateRefresherFunc) *retry.StateChangeConf {
	callFailedCount := 0
	return &retry.StateChangeConf{
		Pending:      status.pending,
		Target:       status.target, // Deleted is not an actual state, we added it to fake the state change
		Delay:        30 * time.Second,
		Timeout:      utils.GetPollingTimeout(po, 40*time.Minute),
		PollInterval: 30 * time.Second,
		Refresh:      stateRefresher(ctx, status.clusterID, &callFailedCount, utils.GetCallFailureThreshold(po, 3)),
	}
}

func (r *dwClusterResource) clusterStateRefresh(ctx context.Context, clusterID *string, callFailedCount *int, callFailureThreshold int) func() (any, string, error) {
	return func() (any, string, error) {
		tflog.Debug(ctx, "About to describe cluster")
		resp, err := r.describeCluster(ctx, clusterID)
		if err != nil {
			if strings.Contains(err.Error(), "NOT_FOUND") {
				return &models.DescribeClusterResponse{}, "Deleted", nil
			}
			*callFailedCount++
			if *callFailedCount <= callFailureThreshold {
				tflog.Warn(ctx, fmt.Sprintf("could not describe Data Warehouse AWS cluster "+
					"due to [%s] but threshold limit is not reached yet (%d out of %d).", err.Error(), callFailedCount, callFailureThreshold))
				return nil, "", nil
			}
			tflog.Error(ctx, fmt.Sprintf("error describing Data Warehouse AWS cluster due to [%s] "+
				"failure threshold limit exceeded.", err.Error()))
			return nil, "", err
		}
		*callFailedCount = 0
		cluster := resp.GetPayload()
		tflog.Debug(ctx, fmt.Sprintf("Described cluster %s with status %s", *clusterID, cluster.Cluster.Status))
		return cluster, cluster.Cluster.Status, nil
	}
}

func (r *dwClusterResource) databaseCatalogStateRefresh(ctx context.Context, clusterID *string, callFailedCount *int, callFailureThreshold int) func() (any, string, error) {
	return func() (any, string, error) {
		tflog.Debug(ctx, "About to get DBCs")
		response, err := r.listDatabaseCatalogs(ctx, clusterID)
		if err != nil {
			tflog.Error(ctx,
				fmt.Sprintf("could not list database catalogs, unexpected error: %s", err.Error()),
			)
			return nil, "", err
		}
		resp := response.GetPayload()
		if len(resp.Dbcs) == 0 {
			*callFailedCount++
			if *callFailedCount <= callFailureThreshold {
				tflog.Warn(ctx, fmt.Sprintf("could not find Data Warehouse database catalog "+
					"due to [%s] but threshold limit is not reached yet (%d out of %d).", err.Error(), callFailedCount, callFailureThreshold))
				return nil, "", nil
			}
			tflog.Error(ctx, fmt.Sprintf("error describing Data Warehouse database catalog due to [%s] "+
				"failure threshold limit exceeded.", err.Error()))
			return nil, "", err
		}
		if len(resp.Dbcs) > 1 {
			err = fmt.Errorf("found more than one Data Warehouse database catalog for cluster %s", *clusterID)
			tflog.Error(ctx, fmt.Sprintf("error describing Data Warehouse database catalog due to [%s] ", err.Error()))
			return nil, "", err
		}
		*callFailedCount = 0

		tflog.Debug(ctx, fmt.Sprintf("Found database catalog %s with status %s", resp.Dbcs[0].ID, resp.Dbcs[0].Status))
		return resp.Dbcs[0], resp.Dbcs[0].Status, nil
	}
}

func (r *dwClusterResource) listDatabaseCatalogs(ctx context.Context, clusterID *string) (*operations.ListDbcsOK, error) {
	catalogParams := operations.NewListDbcsParamsWithContext(ctx).WithInput(&models.ListDbcsRequest{ClusterID: clusterID})
	response, err := r.client.Dw.Operations.ListDbcs(catalogParams)
	return response, err
}
