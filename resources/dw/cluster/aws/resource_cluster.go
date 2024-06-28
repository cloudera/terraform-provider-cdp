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
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

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

	// Create new aws cluster
	response, err := r.client.Dw.Operations.CreateAwsCluster(clusterParams)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating data warehouse aws cluster",
			"Could not create cluster, unexpected error: "+err.Error(),
		)
		return
	}

	payload := response.GetPayload()
	desc := operations.NewDescribeClusterParamsWithContext(ctx).
		WithInput(&models.DescribeClusterRequest{ClusterID: &payload.ClusterID})
	describe, err := r.client.Dw.Operations.DescribeCluster(desc)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating data warehouse aws cluster",
			"Could not describe cluster, unexpected error: "+err.Error(),
		)
		return
	}

	cluster := describe.GetPayload()

	// Map response body to schema and populate Computed attribute values
	plan.ID = types.StringValue(cluster.Cluster.EnvironmentCrn)
	plan.Crn = types.StringValue(cluster.Cluster.EnvironmentCrn)
	plan.Name = types.StringValue(cluster.Cluster.Name)
	plan.ClusterID = types.StringValue(cluster.Cluster.ID)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

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

	op := operations.NewDeleteClusterParamsWithContext(ctx).
		WithInput(&models.DeleteClusterRequest{
			ClusterID: state.ClusterID.ValueStringPointer(),
			// Force:     true,
		})

	if _, err := r.client.Dw.Operations.DeleteCluster(op); err != nil {
		resp.Diagnostics.AddError(
			"Error deleting data warehouse aws cluster",
			"Could not delete cluster, unexpected error: "+err.Error(),
		)
		return
	}
}
