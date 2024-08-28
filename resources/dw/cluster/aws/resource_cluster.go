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

var (
	_ resource.Resource              = (*dwClusterResource)(nil)
	_ resource.ResourceWithConfigure = (*dwClusterResource)(nil)
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

	// Generate API request body from plan
	clusterParams := operations.NewCreateAwsClusterParamsWithContext(ctx).
		WithInput(plan.convertToCreateAwsClusterRequest())

	// Create new AWS cluster
	response, err := r.client.Dw.Operations.CreateAwsCluster(clusterParams)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Data Warehouse AWS cluster",
			"Could not create cluster, unexpected error: "+err.Error(),
		)
		return
	}
	payload := response.GetPayload()
	clusterID := &payload.ClusterID
	plan.ClusterID = types.StringValue(*clusterID)

	desc := operations.NewDescribeClusterParamsWithContext(ctx).
		WithInput(&models.DescribeClusterRequest{ClusterID: clusterID})
	describe, err := r.client.Dw.Operations.DescribeCluster(desc)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Data Warehouse AWS cluster",
			"Could not describe cluster, unexpected error: "+err.Error(),
		)
		return
	}

	cluster := describe.GetPayload()

	// Map response body to schema and populate Computed attribute values
	plan.ID = types.StringValue(cluster.Cluster.EnvironmentCrn)
	plan.Crn = types.StringValue(cluster.Cluster.EnvironmentCrn)
	plan.Name = types.StringValue(cluster.Cluster.Name)
	plan.Status = types.StringValue(cluster.Cluster.Status)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if plan.PollingOptions.Async.ValueBool() {
		callFailedCount := 0
		stateConf := &retry.StateChangeConf{
			Pending:      []string{"Accepted", "Creating", "Created", "Starting"},
			Target:       []string{"Running"},
			Delay:        30 * time.Second,
			Timeout:      plan.getPollingTimeout(),
			PollInterval: 30 * time.Second,
			Refresh:      r.stateRefresh(ctx, clusterID, &callFailedCount, int(plan.PollingOptions.CallFailureThreshold.ValueInt64())),
		}
		if _, err = stateConf.WaitForStateContext(ctx); err != nil {
			resp.Diagnostics.AddError(
				"Error waiting for Data Warehouse AWS cluster",
				"Could not create cluster, unexpected error: "+err.Error(),
			)
			return
		}
		plan.Status = types.StringValue(cluster.Cluster.Status)
		plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *dwClusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Warn(ctx, "Read operation is not implemented yet.")
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
	op := operations.NewDeleteClusterParamsWithContext(ctx).
		WithInput(&models.DeleteClusterRequest{
			ClusterID: clusterID,
		})

	if _, err := r.client.Dw.Operations.DeleteCluster(op); err != nil {
		resp.Diagnostics.AddError(
			"Error deleting Data Warehouse AWS cluster",
			"Could not delete cluster, unexpected error: "+err.Error(),
		)
		return
	}

	if state.PollingOptions.Async.ValueBool() {
		callFailedCount := 0
		stateConf := &retry.StateChangeConf{
			Pending:      []string{"Deleting", "Running"},
			Target:       []string{"Deleted"}, // This is not an actual state, we added it to fake the state change
			Delay:        30 * time.Second,
			Timeout:      state.getPollingTimeout(),
			PollInterval: 30 * time.Second,
			Refresh:      r.stateRefresh(ctx, clusterID, &callFailedCount, int(state.PollingOptions.CallFailureThreshold.ValueInt64())),
		}
		if _, err := stateConf.WaitForStateContext(ctx); err != nil {
			resp.Diagnostics.AddError(
				"Esrror waiting for Data Warehouse AWS cluster",
				"Could not delete cluster, unexpected error: "+err.Error(),
			)
			return
		}
	}
}

func (r *dwClusterResource) stateRefresh(ctx context.Context, clusterID *string, callFailedCount *int, callFailureThreshold int) func() (any, string, error) {
	return func() (any, string, error) {
		tflog.Debug(ctx, "About to describe cluster")
		params := operations.NewDescribeClusterParamsWithContext(ctx).
			WithInput(&models.DescribeClusterRequest{ClusterID: clusterID})
		resp, err := r.client.Dw.Operations.DescribeCluster(params)
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
