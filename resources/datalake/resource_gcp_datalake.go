// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datalake

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client/operations"
	datalakemodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ resource.ResourceWithConfigure   = &gcpDatalakeResource{}
	_ resource.ResourceWithImportState = &gcpDatalakeResource{}
	_ resource.ResourceWithModifyPlan  = &gcpDatalakeResource{}
)

type gcpDatalakeResource struct {
	client *cdp.Client
}

func (r *gcpDatalakeResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if !req.State.Raw.IsNull() {
		resp.Diagnostics.AddError(
			"Resource Update Considerations",
			"Due to provider limitations of this technical preview, modifications are not possible. "+
				"Use the web interface or the CLI to update this resource.")
	}
}

func (r *gcpDatalakeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func NewGcpDatalakeResource() resource.Resource {
	return &gcpDatalakeResource{}
}

func (r *gcpDatalakeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_datalake_gcp_datalake"
}

func (r *gcpDatalakeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *gcpDatalakeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state gcpDatalakeResourceModel
	diags := req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Got error while trying to set plan")
		return
	}

	client := r.client.Datalake

	params := operations.NewCreateGCPDatalakeParamsWithContext(ctx)
	params.WithInput(toGcpDatalakeRequest(ctx, &state))
	responseOk, err := client.Operations.CreateGCPDatalake(params)
	if err != nil {
		utils.AddDatalakeDiagnosticsError(err, &resp.Diagnostics, "create GCP Datalake")
		return
	}

	datalakeResp := responseOk.Payload
	toGcpDatalakeResourceModel(datalakeResp, &state)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !(state.PollingOptions != nil && state.PollingOptions.Async.ValueBool()) {
		stateSaver := func(dlDtl *datalakemodels.DatalakeDetails) {
			datalakeDetailsToGcpDatalakeResourceModel(ctx, dlDtl, &state, state.PollingOptions, &resp.Diagnostics)
			diags = resp.State.Set(ctx, state)
			resp.Diagnostics.Append(diags...)
		}
		if err := waitForDatalakeToBeRunning(ctx, state.DatalakeName.ValueString(), time.Hour, callFailureThreshold, r.client.Datalake, state.PollingOptions, stateSaver); err != nil {
			utils.AddDatalakeDiagnosticsError(err, &resp.Diagnostics, "create AWS Datalake")
			return
		}
	}

	descParams := operations.NewDescribeDatalakeParamsWithContext(ctx)
	descParams.WithInput(&datalakemodels.DescribeDatalakeRequest{DatalakeName: state.DatalakeName.ValueStringPointer()})
	descResponseOk, err := client.Operations.DescribeDatalake(descParams)
	if err != nil {
		utils.AddDatalakeDiagnosticsError(err, &resp.Diagnostics, "create GCP Datalake")
		return
	}

	descDlResp := descResponseOk.Payload
	datalakeDetailsToGcpDatalakeResourceModel(ctx, descDlResp.Datalake, &state, state.PollingOptions, &resp.Diagnostics)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *gcpDatalakeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state gcpDatalakeResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Datalake

	dlName := state.DatalakeName.ValueString()
	if len(dlName) == 0 {
		dlName = state.ID.ValueString()
	}
	params := operations.NewDescribeDatalakeParamsWithContext(ctx)
	params.WithInput(&datalakemodels.DescribeDatalakeRequest{DatalakeName: &dlName})
	responseOk, err := client.Operations.DescribeDatalake(params)
	if err != nil {
		if dlErr, ok := err.(*operations.DescribeDatalakeDefault); ok {
			if cdp.IsDatalakeError(dlErr.GetPayload(), "NOT_FOUND", "") {
				resp.Diagnostics.AddWarning("Resource not found on provider", "Data lake not found, removing from state.")
				tflog.Warn(ctx, "Data lake not found, removing from state", map[string]interface{}{
					"id": state.ID.ValueString(),
				})
				resp.State.RemoveResource(ctx)
				return
			}
		}
		utils.AddDatalakeDiagnosticsError(err, &resp.Diagnostics, "read GCP Datalake")
		return
	}

	datalakeResp := responseOk.Payload
	datalakeDetailsToGcpDatalakeResourceModel(ctx, datalakeResp.Datalake, &state, state.PollingOptions, &resp.Diagnostics)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *gcpDatalakeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *gcpDatalakeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state gcpDatalakeResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Datalake
	params := operations.NewDeleteDatalakeParamsWithContext(ctx)
	params.WithInput(&datalakemodels.DeleteDatalakeRequest{
		DatalakeName: state.DatalakeName.ValueStringPointer(),
		Force:        false,
	})
	_, err := client.Operations.DeleteDatalake(params)
	if err != nil {
		if dlErr, ok := err.(*operations.DescribeDatalakeDefault); ok {
			if cdp.IsDatalakeError(dlErr.GetPayload(), "NOT_FOUND", "") {
				tflog.Info(ctx, "Data lake already deleted", map[string]interface{}{
					"id": state.ID.ValueString(),
				})
				return
			}
		}
		utils.AddDatalakeDiagnosticsError(err, &resp.Diagnostics, "delete GCP Datalake")
		return
	}

	if err := waitForDatalakeToBeDeleted(ctx, state.DatalakeName.ValueString(), time.Hour, r.client.Datalake, state.PollingOptions); err != nil {
		utils.AddDatalakeDiagnosticsError(err, &resp.Diagnostics, "delete GCP Datalake")
		return
	}

}
