// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package iam

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/client/operations"
	iammodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ resource.Resource = &groupResource{}
)

type samlProvider struct {
	client *cdp.Client
}

func NewSamlProvider() resource.Resource {
	return &samlProvider{}
}

func (r *samlProvider) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_saml_provider"
}

func (r *samlProvider) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *samlProvider) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "SAML provider creation requested.")
	var plan samlProviderModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewCreateSamlProviderParamsWithContext(ctx)
	params.WithInput(&iammodels.CreateSamlProviderRequest{
		EnableScim:                      plan.EnableScim.ValueBool(),
		GenerateWorkloadUsernameByEmail: plan.GenerateWorkloadUsernameByEmail.ValueBool(),
		SamlMetadataDocument:            plan.SamlMetadataDocument.ValueString(),
		SamlProviderName:                plan.SamlProviderName.ValueStringPointer(),
		SyncGroupsOnLogin:               plan.SyncGroupsOnLogin.ValueBool(),
	})

	tflog.Debug(ctx, fmt.Sprintf("About to create SAML provider using the following request: %+v", *params.Input))

	response, err := r.client.Iam.Operations.CreateSamlProvider(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating SAML provider",
			"Got error while creating SAML provider: "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, "SAML provider creation finished successfully.")

	plan.GenerateWorkloadUsernameByEmail = types.BoolValue(response.Payload.SamlProvider.GenerateWorkloadUsernameByEmail)
	plan.SamlMetadataDocument = types.StringValue(response.Payload.SamlProvider.SamlMetadataDocument)
	plan.SyncGroupsOnLogin = types.BoolPointerValue(response.Payload.SamlProvider.SyncGroupsOnLogin)
	plan.SamlProviderName = types.StringPointerValue(response.Payload.SamlProvider.SamlProviderName)
	plan.SamlProviderId = types.StringPointerValue(response.Payload.SamlProvider.SamlProviderID)
	plan.CdpSpMetadata = types.StringValue(response.Payload.SamlProvider.CdpSpMetadata) // only for create & describe
	plan.CreationDate = types.StringValue(response.Payload.SamlProvider.CreationDate.String())
	plan.EnableScim = types.BoolValue(response.Payload.SamlProvider.EnableScim)
	plan.ScimURL = types.StringValue(response.Payload.SamlProvider.ScimURL)
	plan.Crn = types.StringPointerValue(response.Payload.SamlProvider.Crn)

	// Save data into Terraform state
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *samlProvider) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state samlProviderModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewDescribeSamlProviderParamsWithContext(ctx)
	params.WithInput(&iammodels.DescribeSamlProviderRequest{SamlProviderName: state.SamlProviderName.ValueString()})

	response, err := r.client.Iam.Operations.DescribeSamlProvider(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading SAML provider",
			"Got error while reading SAML provider: "+err.Error(),
		)
		return
	}

	state.GenerateWorkloadUsernameByEmail = types.BoolValue(response.Payload.SamlProvider.GenerateWorkloadUsernameByEmail)
	state.SamlMetadataDocument = types.StringValue(response.Payload.SamlProvider.SamlMetadataDocument)
	state.SyncGroupsOnLogin = types.BoolPointerValue(response.Payload.SamlProvider.SyncGroupsOnLogin)
	state.SamlProviderName = types.StringPointerValue(response.Payload.SamlProvider.SamlProviderName)
	state.SamlProviderId = types.StringPointerValue(response.Payload.SamlProvider.SamlProviderID)
	state.CdpSpMetadata = types.StringValue(response.Payload.SamlProvider.CdpSpMetadata) // only for create & describe
	state.CreationDate = types.StringValue(response.Payload.SamlProvider.CreationDate.String())
	state.EnableScim = types.BoolValue(response.Payload.SamlProvider.EnableScim)
	state.ScimURL = types.StringValue(response.Payload.SamlProvider.ScimURL)
	state.Crn = types.StringPointerValue(response.Payload.SamlProvider.Crn)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *samlProvider) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Get current state and the actual plan
	var plan, state samlProviderModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewUpdateSamlProviderParamsWithContext(ctx)
	params.WithInput(&iammodels.UpdateSamlProviderRequest{
		EnableScim:                      plan.EnableScim.ValueBool(),
		GenerateWorkloadUsernameByEmail: plan.GenerateWorkloadUsernameByEmail.ValueBool(),
		SamlMetadataDocument:            plan.SamlMetadataDocument.ValueString(),
		SamlProviderName:                plan.SamlProviderName.ValueStringPointer(),
		SyncGroupsOnLogin:               plan.SyncGroupsOnLogin.ValueBool(),
	})

	response, err := r.client.Iam.Operations.UpdateSamlProvider(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating SAML provider",
			"Got error while updating SAML provider: "+err.Error(),
		)
		return
	}

	state.GenerateWorkloadUsernameByEmail = types.BoolValue(response.Payload.SamlProvider.GenerateWorkloadUsernameByEmail)
	state.SamlMetadataDocument = types.StringValue(response.Payload.SamlProvider.SamlMetadataDocument)
	state.SyncGroupsOnLogin = types.BoolPointerValue(response.Payload.SamlProvider.SyncGroupsOnLogin)
	state.SamlProviderName = types.StringPointerValue(response.Payload.SamlProvider.SamlProviderName)
	state.SamlProviderId = types.StringPointerValue(response.Payload.SamlProvider.SamlProviderID)
	state.CreationDate = types.StringValue(response.Payload.SamlProvider.CreationDate.String())
	state.EnableScim = types.BoolValue(response.Payload.SamlProvider.EnableScim)
	state.ScimURL = types.StringValue(response.Payload.SamlProvider.ScimURL)
	state.Crn = types.StringPointerValue(response.Payload.SamlProvider.Crn)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *groupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state groupModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Iam

	groupName := state.ID.ValueString()
	params := operations.NewDeleteGroupParamsWithContext(ctx)
	params.WithInput(&iammodels.DeleteGroupRequest{GroupName: &groupName})
	_, err := client.Operations.DeleteGroup(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting Group",
			"Could not delete Group, unexpected error: "+err.Error(),
		)
		return
	}
}
