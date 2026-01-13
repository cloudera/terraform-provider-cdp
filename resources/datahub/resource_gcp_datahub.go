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

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client/operations"
	datahubmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ resource.ResourceWithConfigure   = &gcpDatahubResource{}
	_ resource.ResourceWithImportState = &gcpDatahubResource{}
)

type gcpDatahubResource struct {
	client *cdp.Client
}

func (r *gcpDatahubResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func NewGcpDatahubResource() resource.Resource {
	return &gcpDatahubResource{}
}

func (r *gcpDatahubResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_datahub_gcp_cluster"
}

func (r *gcpDatahubResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *gcpDatahubResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "GCP cluster creation process requested.")
	var data gcpDatahubResourceModel
	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewCreateGCPClusterParamsWithContext(ctx)
	params.WithInput(fromModelToGcpRequest(data, ctx))

	res, err := r.client.Datahub.Operations.CreateGCPCluster(params)
	if err != nil {
		tflog.Error(ctx, err.Error())
		utils.AddDatahubDiagnosticsError(err, &resp.Diagnostics, "create GCP Datahub")
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

	status, err := waitForToBeAvailable(data.ID.ValueString(), r.client.Datahub, ctx, data.PollingOptions)
	tflog.Debug(ctx, fmt.Sprintf("Cluster polling finished, setting status from '%s' to '%s'", data.Status.ValueString(), status))
	//TODO: Should save to state fields filled by the backend from the response to make the resource more versatile for TF developers
	data.Status = types.StringValue(status)
	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err != nil {
		tflog.Debug(ctx, fmt.Sprintf("Cluster creation has ended up in error: %s", err.Error()))
		utils.AddDatahubDiagnosticsError(err, &resp.Diagnostics, "create GCP Datahub")
		return
	}
}

func (r *gcpDatahubResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state gcpDatahubResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	clusterName := state.Name.ValueString()
	if len(clusterName) == 0 {
		clusterName = state.ID.ValueString()
	}
	params := operations.NewDescribeClusterParamsWithContext(ctx)
	params.WithInput(&datahubmodels.DescribeClusterRequest{
		ClusterName: &clusterName,
	})

	result, err := r.client.Datahub.Operations.DescribeCluster(params)
	if err != nil {
		if isNotFoundError(err) {
			resp.Diagnostics.AddWarning("Resource not found on provider", "GCP Data hub cluster not found, removing from state.")
			tflog.Warn(ctx, "GCP Data hub cluster not found, removing from state", map[string]interface{}{"id": state.ID.ValueString()})
			resp.State.RemoveResource(ctx)
			return
		}
		utils.AddDatahubDiagnosticsError(err, &resp.Diagnostics, "read GCP Datahub")
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

func (r *gcpDatahubResource) Update(ctx context.Context, _ resource.UpdateRequest, _ *resource.UpdateResponse) {
	tflog.Warn(ctx, "Update operation is not implemented yet.")
}

func (r *gcpDatahubResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state gcpDatahubResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewDeleteClusterParamsWithContext(ctx).WithInput(&datahubmodels.DeleteClusterRequest{
		ClusterName: state.ID.ValueStringPointer(),
		Force:       state.forceDeleteRequested(),
	})
	if state.forceDeleteRequested() {
		tflog.Debug(ctx, fmt.Sprintf("Sending force delete request for cluster: %s", *params.Input.ClusterName))
	} else {
		tflog.Debug(ctx, fmt.Sprintf("Sending delete request for cluster: %s", *params.Input.ClusterName))
	}
	_, err := r.client.Datahub.Operations.DeleteCluster(params)
	if err != nil {
		if !isNotFoundError(err) {
			utils.AddDatahubDiagnosticsError(err, &resp.Diagnostics, "delete GCP Datahub")
		}
		return
	}

	err = waitForToBeDeleted(state.Name.ValueString(), r.client.Datahub, ctx, state.PollingOptions)
	if err != nil {
		utils.AddDatahubDiagnosticsError(err, &resp.Diagnostics, "delete GCP Datahub")
		return
	}
}
