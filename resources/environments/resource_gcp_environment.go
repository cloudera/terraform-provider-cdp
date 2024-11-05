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

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

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
		&data, data.PollingOptions, &resp.Diagnostics)

	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	descEnvResp, err := describeEnvironmentWithDiagnosticHandle(data.EnvironmentName.ValueString(), data.ID.ValueString(), ctx, r.client, &resp.Diagnostics, &resp.State)
	if err != nil {
		return
	}
	if !(data.PollingOptions != nil && data.PollingOptions.Async.ValueBool()) {
		stateSaver := func(env *environmentsmodels.Environment) {
			toGcpEnvironmentResource(ctx, utils.LogEnvironmentSilently(ctx, env, describeLogPrefix), &data, data.PollingOptions, &resp.Diagnostics)
			diags = resp.State.Set(ctx, data)
			resp.Diagnostics.Append(diags...)
		}
		descEnvResp, err = waitForCreateEnvironmentWithDiagnosticHandle(ctx, r.client, data.ID.ValueString(), data.EnvironmentName.ValueString(), resp, data.PollingOptions, stateSaver)
		if err != nil {
			return
		}
	}

	toGcpEnvironmentResource(ctx, utils.LogEnvironmentSilently(ctx, descEnvResp, describeLogPrefix), &data, data.PollingOptions, &resp.Diagnostics)
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

	descEnvResp, err := describeEnvironmentWithDiagnosticHandle(state.EnvironmentName.ValueString(), state.ID.ValueString(), ctx, r.client, &resp.Diagnostics, &resp.State)
	if err != nil {
		return
	}
	toGcpEnvironmentResource(ctx, descEnvResp, &state, state.PollingOptions, &resp.Diagnostics)

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

	cascading := state.Cascading.ValueBool()
	if state.Cascading.IsNull() {
		cascading = true
	}
	if err := deleteEnvironmentWithDiagnosticHandle(state.EnvironmentName.ValueString(), cascading, ctx, r.client, resp, state.PollingOptions); err != nil {
		return
	}
}
