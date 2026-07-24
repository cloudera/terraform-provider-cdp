// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package gcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/resources/environments/envcommon"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ resource.ResourceWithConfigure   = &gcpEnvironmentResource{}
	_ resource.ResourceWithImportState = &gcpEnvironmentResource{}
)

type gcpEnvironmentResource struct {
	client *cdp.Client
}

func (r *gcpEnvironmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func NewGcpEnvironmentResource() resource.Resource {
	return &gcpEnvironmentResource{}
}

func (r *gcpEnvironmentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = GcpEnvironmentSchema
}

func (r *gcpEnvironmentResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_gcp_environment"
}

func (r *gcpEnvironmentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *gcpEnvironmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GcpEnvironmentResourceModel
	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Got Error while trying to set plan")
		return
	}

	client := r.client.Environments

	params := operations.NewCreateGCPEnvironmentParams()
	params.WithInput(toGcpEnvironmentRequest(ctx, &data))

	responseOk, err := client.Operations.CreateGCPEnvironmentContext(ctx, params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "create GCP Environment")
		return
	}

	ToGcpEnvironmentResource(ctx,
		utils.LogEnvironmentSilently(ctx, responseOk.Payload.Environment, envcommon.DescribeLogPrefix),
		&data, data.PollingOptions, &resp.Diagnostics)

	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	descEnvResp, err := envcommon.DescribeEnvironmentWithDiagnosticHandle(data.EnvironmentName.ValueString(), data.ID.ValueString(), ctx, r.client, &resp.Diagnostics, &resp.State)
	if err != nil {
		return
	}
	if data.PollingOptions == nil || !data.PollingOptions.Async.ValueBool() {
		stateSaver := func(env *environmentsmodels.Environment) {
			ToGcpEnvironmentResource(ctx, utils.LogEnvironmentSilently(ctx, env, envcommon.DescribeLogPrefix), &data, data.PollingOptions, &resp.Diagnostics)
			diags = resp.State.Set(ctx, data)
			resp.Diagnostics.Append(diags...)
		}
		descEnvResp, err = envcommon.WaitForCreateEnvironmentWithDiagnosticHandle(ctx, r.client, data.ID.ValueString(), data.EnvironmentName.ValueString(), resp, data.PollingOptions, stateSaver)
		if err != nil {
			return
		}
	}

	ToGcpEnvironmentResource(ctx, utils.LogEnvironmentSilently(ctx, descEnvResp, envcommon.DescribeLogPrefix), &data, data.PollingOptions, &resp.Diagnostics)
	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *gcpEnvironmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state GcpEnvironmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	envName := state.EnvironmentName.ValueString()
	if len(envName) == 0 {
		envName = state.ID.ValueString()
	}
	descEnvResp, err := envcommon.DescribeEnvironmentWithDiagnosticHandle(envName, state.ID.ValueString(), ctx, r.client, &resp.Diagnostics, &resp.State)
	if err != nil {
		return
	}
	ToGcpEnvironmentResource(ctx, descEnvResp, &state, state.PollingOptions, &resp.Diagnostics)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *gcpEnvironmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	envcommon.PerformEnvironmentUpdate(ctx, req, resp, r.client.Environments, updateGcpEnvironment)
}

func (r *gcpEnvironmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state GcpEnvironmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	cascading := state.cascadeDelete()
	forced := state.forceDelete()

	if err := envcommon.DeleteEnvironmentWithDiagnosticHandle(state.EnvironmentName.ValueString(), cascading, forced, ctx, r.client, resp, state.PollingOptions); err != nil {
		return
	}
}
