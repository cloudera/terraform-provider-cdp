// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package environments

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

const describeLogPrefix = "Result of describe environment: "

var (
	_ resource.Resource = &gcpEnvironmentResource{}
)

type gcpEnvironmentResource struct {
	client *cdp.Client
}

func NewGcpEnvironmentResource() resource.Resource {
	return &gcpEnvironmentResource{}
}

func (r *gcpEnvironmentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_gcp_environment"
}

func (r *gcpEnvironmentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *gcpEnvironmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data gcpEnvironmentResourceModel
	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Got Error while trying to set plan")
		return
	}

	client := r.client.Environments

	params := operations.NewCreateGCPEnvironmentParamsWithContext(ctx)
	params.WithInput(toGcpEnvironmentRequest(ctx, &data))

	responseOk, err := client.Operations.CreateGCPEnvironment(params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "create GCP Environment")
		return
	}

	toGcpEnvironmentResource(ctx,
		utils.LogEnvironmentSilently(ctx, responseOk.Payload.Environment, describeLogPrefix),
		&data, &resp.Diagnostics)

	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	timeout := time.Hour * 1
	if err := waitForEnvironmentToBeAvailable(data.ID.ValueString(), timeout, client, ctx); err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "create GCP Environment")
		return
	}

	environmentName := data.EnvironmentName.ValueString()
	descParams := operations.NewDescribeEnvironmentParamsWithContext(ctx)
	descParams.WithInput(&environmentsmodels.DescribeEnvironmentRequest{
		EnvironmentName: &environmentName,
	})
	descEnvResp, err := r.client.Environments.Operations.DescribeEnvironment(descParams)
	if err != nil {
		if isEnvNotFoundError(err) {
			resp.Diagnostics.AddWarning("Resource not found on provider", "Environment not found, removing from state.")
			tflog.Warn(ctx, "Environment not found, removing from state", map[string]interface{}{
				"id": data.ID.ValueString(),
			})
			resp.State.RemoveResource(ctx)
			return
		}
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "create GCP Environment")
		return
	}

	toGcpEnvironmentResource(ctx, utils.LogEnvironmentSilently(ctx, descEnvResp.GetPayload().Environment, describeLogPrefix), &data, &resp.Diagnostics)
	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *gcpEnvironmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state gcpEnvironmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	environmentName := state.EnvironmentName.ValueString()
	params := operations.NewDescribeEnvironmentParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.DescribeEnvironmentRequest{
		EnvironmentName: &environmentName,
	})
	descEnvResp, err := r.client.Environments.Operations.DescribeEnvironment(params)
	if err != nil {
		if isEnvNotFoundError(err) {
			resp.Diagnostics.AddWarning("Resource not found on provider", "Environment not found, removing from state.")
			tflog.Warn(ctx, "Environment not found, removing from state", map[string]interface{}{
				"id": state.ID.ValueString(),
			})
			resp.State.RemoveResource(ctx)
			return
		}
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "read GCP Environment")
		return
	}

	toGcpEnvironmentResource(ctx, utils.LogEnvironmentSilently(ctx, descEnvResp.GetPayload().Environment, describeLogPrefix), &state, &resp.Diagnostics)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *gcpEnvironmentResource) Update(ctx context.Context, _ resource.UpdateRequest, _ *resource.UpdateResponse) {
	tflog.Warn(ctx, "Update operation is not implemented yet.")
}

func (r *gcpEnvironmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state gcpEnvironmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	environmentName := state.EnvironmentName.ValueString()
	params := operations.NewDeleteEnvironmentParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.DeleteEnvironmentRequest{EnvironmentName: &environmentName})
	_, err := r.client.Environments.Operations.DeleteEnvironment(params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "delete GCP Environment")
		return
	}

	timeout := time.Hour * 1
	err = waitForEnvironmentToBeDeleted(environmentName, timeout, r.client.Environments, ctx)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "delete GCP Environment")
		return
	}
}
