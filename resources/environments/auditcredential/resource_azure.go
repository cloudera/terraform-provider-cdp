// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package auditcredential

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
	_ resource.ResourceWithConfigure   = &azureAuditCredentialResource{}
	_ resource.ResourceWithImportState = &azureAuditCredentialResource{}
)

type azureAuditCredentialResource struct {
	client *cdp.Client
}

func NewAzureAuditCredentialResource() resource.Resource {
	return &azureAuditCredentialResource{}
}

func (r *azureAuditCredentialResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_azure_audit_credential"
}

func (r *azureAuditCredentialResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *azureAuditCredentialResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan azureAuditCredentialResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewSetAzureAuditCredentialParams()
	params.WithInput(&environmentsmodels.SetAzureAuditCredentialRequest{
		SubscriptionID: plan.SubscriptionID.ValueStringPointer(),
		TenantID:       plan.TenantID.ValueStringPointer(),
		AppBased: &environmentsmodels.SetAzureAuditCredentialRequestAppBased{
			ApplicationID: plan.AppBased.ApplicationID.ValueStringPointer(),
			SecretKey:     plan.AppBased.SecretKey.ValueStringPointer(),
		},
	})

	result, err := r.client.Environments.Operations.SetAzureAuditCredentialContext(ctx, params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "create Azure Audit Credential")
		return
	}
	if !credentialOrError(result.Payload.Credential, &resp.Diagnostics, "create Azure Audit Credential") {
		return
	}

	mapAuditCredentialResponse(result.Payload.Credential, &plan.ID, &plan.CredentialName, &plan.Crn, &plan.Description)
	mapAzureCredentialProperties(result.Payload.Credential, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *azureAuditCredentialResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state azureAuditCredentialResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	c, err := findAuditCredentialByName(ctx, r.client, state.ID.ValueString())
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "read Azure Audit Credential")
		return
	}
	if c == nil {
		resp.State.RemoveResource(ctx)
		return
	}

	mapAuditCredentialResponse(c, &state.ID, &state.CredentialName, &state.Crn, &state.Description)
	mapAzureCredentialProperties(c, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *azureAuditCredentialResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan azureAuditCredentialResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewSetAzureAuditCredentialParams()
	params.WithInput(&environmentsmodels.SetAzureAuditCredentialRequest{
		SubscriptionID: plan.SubscriptionID.ValueStringPointer(),
		TenantID:       plan.TenantID.ValueStringPointer(),
		AppBased: &environmentsmodels.SetAzureAuditCredentialRequestAppBased{
			ApplicationID: plan.AppBased.ApplicationID.ValueStringPointer(),
			SecretKey:     plan.AppBased.SecretKey.ValueStringPointer(),
		},
	})

	result, err := r.client.Environments.Operations.SetAzureAuditCredentialContext(ctx, params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "update Azure Audit Credential")
		return
	}
	if !credentialOrError(result.Payload.Credential, &resp.Diagnostics, "update Azure Audit Credential") {
		return
	}

	mapAuditCredentialResponse(result.Payload.Credential, &plan.ID, &plan.CredentialName, &plan.Crn, &plan.Description)
	mapAzureCredentialProperties(result.Payload.Credential, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *azureAuditCredentialResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state azureAuditCredentialResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	deleteAuditCredential(ctx, r.client, state.CredentialName.ValueString(), &resp.Diagnostics, "delete Azure Audit Credential")
}

func (r *azureAuditCredentialResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func mapAzureCredentialProperties(c *environmentsmodels.Credential, state *azureAuditCredentialResourceModel) {
	if c.AzureCredentialProperties == nil {
		return
	}
	state.SubscriptionID = types.StringValue(c.AzureCredentialProperties.SubscriptionID)
	state.TenantID = types.StringValue(c.AzureCredentialProperties.TenantID)
	if state.AppBased == nil {
		state.AppBased = &AuditCredAppBased{}
	}
	if c.AzureCredentialProperties.AppID != "" {
		state.AppBased.ApplicationID = types.StringValue(c.AzureCredentialProperties.AppID)
	}
}
