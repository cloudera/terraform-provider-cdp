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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ resource.ResourceWithConfigure   = &awsGovCloudAuditCredentialResource{}
	_ resource.ResourceWithImportState = &awsGovCloudAuditCredentialResource{}
)

type awsGovCloudAuditCredentialResource struct {
	client *cdp.Client
}

func NewAwsGovCloudAuditCredentialResource() resource.Resource {
	return &awsGovCloudAuditCredentialResource{}
}

func (r *awsGovCloudAuditCredentialResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_aws_gov_cloud_audit_credential"
}

func (r *awsGovCloudAuditCredentialResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *awsGovCloudAuditCredentialResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan awsAuditCredentialResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewSetAWSGovCloudAuditCredentialParams()
	params.WithInput(&environmentsmodels.SetAWSGovCloudAuditCredentialRequest{
		RoleArn: plan.RoleArn.ValueStringPointer(),
	})

	result, err := r.client.Environments.Operations.SetAWSGovCloudAuditCredentialContext(ctx, params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "create AWS GovCloud Audit Credential")
		return
	}
	if !credentialOrError(result.Payload.Credential, &resp.Diagnostics, "create AWS GovCloud Audit Credential") {
		return
	}

	mapAuditCredentialResponse(result.Payload.Credential, &plan.ID, &plan.CredentialName, &plan.Crn, &plan.Description)
	mapAwsCredentialProperties(result.Payload.Credential, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *awsGovCloudAuditCredentialResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state awsAuditCredentialResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	c, err := findAuditCredentialByName(ctx, r.client, state.ID.ValueString())
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "read AWS GovCloud Audit Credential")
		return
	}
	if c == nil {
		resp.State.RemoveResource(ctx)
		return
	}

	mapAuditCredentialResponse(c, &state.ID, &state.CredentialName, &state.Crn, &state.Description)
	mapAwsCredentialProperties(c, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *awsGovCloudAuditCredentialResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan awsAuditCredentialResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewSetAWSGovCloudAuditCredentialParams()
	params.WithInput(&environmentsmodels.SetAWSGovCloudAuditCredentialRequest{
		RoleArn: plan.RoleArn.ValueStringPointer(),
	})

	result, err := r.client.Environments.Operations.SetAWSGovCloudAuditCredentialContext(ctx, params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "update AWS GovCloud Audit Credential")
		return
	}
	if !credentialOrError(result.Payload.Credential, &resp.Diagnostics, "update AWS GovCloud Audit Credential") {
		return
	}

	mapAuditCredentialResponse(result.Payload.Credential, &plan.ID, &plan.CredentialName, &plan.Crn, &plan.Description)
	mapAwsCredentialProperties(result.Payload.Credential, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *awsGovCloudAuditCredentialResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state awsAuditCredentialResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	deleteAuditCredential(ctx, r.client, state.CredentialName.ValueString(), &resp.Diagnostics, "delete AWS GovCloud Audit Credential")
}

func (r *awsGovCloudAuditCredentialResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
