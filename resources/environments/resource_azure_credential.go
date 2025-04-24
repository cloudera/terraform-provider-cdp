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

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ resource.ResourceWithConfigure   = &azureCredentialResource{}
	_ resource.ResourceWithImportState = &azureCredentialResource{}
)

type azureCredentialResource struct {
	client *cdp.Client
}

func (r *azureCredentialResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func NewAzureCredentialResource() resource.Resource {
	return &azureCredentialResource{}
}

func (r *azureCredentialResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_azure_credential"
}

func (r *azureCredentialResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *azureCredentialResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data azureCredentialResourceModel
	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Environments

	params := operations.NewCreateAzureCredentialParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.CreateAzureCredentialRequest{
		CredentialName: data.CredentialName.ValueStringPointer(),
		Description:    data.Description.ValueString(),
		SubscriptionID: data.SubscriptionID.ValueString(),
		TenantID:       data.TenantID.ValueString(),
		AppBased: &environmentsmodels.CreateAzureCredentialRequestAppBased{
			ApplicationID: data.AppBased.ApplicationID.ValueString(),
			SecretKey:     data.AppBased.SecretKey.ValueString(),
		},
	})

	result, err := client.Operations.CreateAzureCredential(params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "create Azure Credential")
		return
	}

	data.Crn = types.StringPointerValue(result.Payload.Credential.Crn)
	data.ID = data.CredentialName

	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *azureCredentialResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state azureCredentialResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed value from CDP
	credentialName := state.CredentialName.ValueString()
	params := operations.NewListCredentialsParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.ListCredentialsRequest{CredentialName: credentialName})
	listCredentialsResp, err := r.client.Environments.Operations.ListCredentials(params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "read Azure Credential")
		return
	}

	// Overwrite items with refreshed state
	credentials := listCredentialsResp.GetPayload().Credentials
	if len(credentials) == 0 || *credentials[0].CredentialName != credentialName {
		resp.Diagnostics.AddError(
			"Error reading Azure Credential.",
			"Azure Credential not found, removing from state.")
		resp.State.RemoveResource(ctx) // deleted
		return
	}
	c := credentials[0]

	state.ID = types.StringPointerValue(c.Crn)
	state.CredentialName = types.StringPointerValue(c.CredentialName)
	state.Crn = types.StringPointerValue(c.Crn)
	if c.AzureCredentialProperties != nil {
		state.AppBased.ApplicationID = types.StringValue(c.AzureCredentialProperties.AppID)
		state.SubscriptionID = types.StringValue(c.AzureCredentialProperties.SubscriptionID)
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *azureCredentialResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *azureCredentialResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state azureCredentialResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewDeleteCredentialParamsWithContext(ctx).WithInput(&environmentsmodels.DeleteCredentialRequest{CredentialName: state.CredentialName.ValueStringPointer()})
	_, err := r.client.Environments.Operations.DeleteCredential(params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "delete Azure Credential")
		return
	}

}
