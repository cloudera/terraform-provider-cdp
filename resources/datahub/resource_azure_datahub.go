// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datahub

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client/operations"
	datahubmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ resource.Resource = &azureDatahubResource{}
)

type azureDatahubResource struct {
	client *cdp.Client
}

func NewAzureDatahubResource() resource.Resource {
	return &azureDatahubResource{}
}

func (r *azureDatahubResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_datahub_azure_cluster"
}

func (r *azureDatahubResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *azureDatahubResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Azure cluster creation process requested.")
	var data datahubResourceModel
	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewCreateAzureClusterParamsWithContext(ctx)
	params.WithInput(fromModelToAzureRequest(data, ctx))

	res, err := r.client.Datahub.Operations.CreateAzureCluster(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Azure Data hub cluster.",
			"Got error while creating Azure Data hub cluster: "+err.Error(),
		)
		return
	}

	data.Crn = types.StringPointerValue(res.Payload.Cluster.Crn)
	data.ID = types.StringPointerValue(res.Payload.Cluster.Crn)
	data.Name = types.StringPointerValue(res.Payload.Cluster.ClusterName)
	data.Status = types.StringValue(res.Payload.Cluster.Status)

	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	status, err := waitForToBeAvailable(data.ID.ValueString(), r.client.Datahub, ctx)
	tflog.Debug(ctx, fmt.Sprintf("Cluster polling finished, setting status from '%s' to '%s'", data.Status.ValueString(), status))
	data.Status = types.StringValue(status)
	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err != nil {
		tflog.Debug(ctx, fmt.Sprintf("Cluster creation has ended up in error: %s", err.Error()))
		resp.Diagnostics.AddError(
			"Error creating Azure Data hub cluster",
			"Failure to poll of Azure Data hub cluster creation: "+err.Error(),
		)
		return
	}
}

func (r *azureDatahubResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state datahubResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewDescribeClusterParamsWithContext(ctx)
	params.WithInput(&datahubmodels.DescribeClusterRequest{
		ClusterName: state.Name.ValueStringPointer(),
	})

	result, err := r.client.Datahub.Operations.DescribeCluster(params)
	if err != nil {
		if isNotFoundError(err) {
			resp.Diagnostics.AddWarning("Resource not found on provider", "Azure Data hub cluster not found, removing from state.")
			tflog.Warn(ctx, "Azure Data hub cluster not found, removing from state", map[string]interface{}{"id": state.ID.ValueString()})
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error Reading Azure Data hub cluster",
			"Could not read Azure Data hub cluster: "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	cluster := result.Payload.Cluster

	state.ID = types.StringPointerValue(cluster.Crn)
	state.Crn = types.StringPointerValue(cluster.Crn)
	state.Status = types.StringValue(cluster.Status)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *azureDatahubResource) Update(ctx context.Context, _ resource.UpdateRequest, _ *resource.UpdateResponse) {
	tflog.Warn(ctx, "Update operation is not implemented yet.")
}

func (r *azureDatahubResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state datahubResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewDeleteClusterParamsWithContext(ctx).WithInput(&datahubmodels.DeleteClusterRequest{
		ClusterName: state.ID.ValueStringPointer(),
	})
	_, err := r.client.Datahub.Operations.DeleteCluster(params)
	if err != nil {
		if !isNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Error Deleting Azure Data hub cluster",
				"Could not delete Azure Data hub cluster due to: "+err.Error(),
			)
		}
		return
	}

	err = waitForToBeDeleted(state.Name.ValueString(), r.client.Datahub, ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting Azure Data hub cluster",
			"Failure to poll Azure Data hub deletion, unexpected error: "+err.Error(),
		)
		return
	}
}
