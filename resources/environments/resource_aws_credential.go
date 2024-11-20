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
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.ResourceWithConfigure   = &awsCredentialResource{}
	_ resource.ResourceWithImportState = &awsCredentialResource{}
)

var (
	credentialCreateRetryDuration = time.Duration(30) * time.Second
)

type awsCredentialResource struct {
	client *cdp.Client
}

func NewAwsCredentialResource() resource.Resource {
	return &awsCredentialResource{}
}

func (r *awsCredentialResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_aws_credential"
}

func (r *awsCredentialResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *awsCredentialResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan awsCredentialResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Environments

	params := operations.NewCreateAWSCredentialParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.CreateAWSCredentialRequest{
		CredentialName: plan.CredentialName.ValueStringPointer(),
		Description:    plan.Description.ValueString(),
		RoleArn:        plan.RoleArn.ValueStringPointer(),
	})

	err := retry.RetryContext(ctx, credentialCreateRetryDuration, func() *retry.RetryError {
		responseOk, err := client.Operations.CreateAWSCredential(params)
		if err != nil {
			if envErr, ok := err.(*operations.CreateAWSCredentialDefault); ok {
				if utils.IsRetryableError(envErr.Code()) {
					return retry.RetryableError(err)
				}
			}
			return retry.NonRetryableError(err)
		}
		plan.Crn = types.StringPointerValue(responseOk.Payload.Credential.Crn)
		plan.ID = types.StringPointerValue(params.Input.CredentialName)
		return nil
	})
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "create AWS Credential")
		return
	}

	// AWS credential creation does not support a couple of fields to be set during creation but can be
	// updated afterward thus if they are set in the plan, we need to update the credential right away
	if credentialNeedsToBeUpdatedAfterCreation(plan) {
		tflog.Debug(ctx, "Updating AWS credential required due to plan configuration")
		updateErr := r.updateCredential(ctx, client, plan)
		if updateErr != nil {
			utils.AddEnvironmentDiagnosticsError(updateErr, &resp.Diagnostics, "update AWS Credential")
			return
		}
	}

	// Save plan into Terraform state
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *awsCredentialResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state awsCredentialResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed value from CDP
	credentialName := state.ID.ValueString()
	c, err := FindCredentialByName(ctx, r.client, credentialName)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "read AWS Credential")
		return
	}
	if c == nil {
		resp.State.RemoveResource(ctx) // deleted
		return
	}

	// Overwrite items with refreshed state
	state.ID = types.StringPointerValue(c.CredentialName)
	state.CredentialName = types.StringPointerValue(c.CredentialName)
	state.Crn = types.StringPointerValue(c.Crn)
	if c.Description != "" {
		state.Description = types.StringValue(c.Description)
	} else {
		state.Description = types.StringNull()
	}
	if c.AwsCredentialProperties != nil {
		state.RoleArn = types.StringValue(c.AwsCredentialProperties.RoleArn)
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// FindCredentialByName reads and returns the credential from CDP if any.
func FindCredentialByName(ctx context.Context, cdpClient *cdp.Client, credentialName string) (*environmentsmodels.Credential, error) {
	params := operations.NewListCredentialsParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.ListCredentialsRequest{CredentialName: credentialName})
	listCredentialsResp, err := cdpClient.Environments.Operations.ListCredentials(params)
	if err != nil {
		return nil, err
	}
	credentials := listCredentialsResp.GetPayload().Credentials
	if len(credentials) == 0 || *credentials[0].CredentialName != credentialName {
		return nil, nil
	}
	return credentials[0], nil
}

func (r *awsCredentialResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan awsCredentialResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Environments

	err := r.updateCredential(ctx, client, plan)

	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "update AWS Credential")
		return
	}
	resp.State.Set(ctx, plan)
}

func (r *awsCredentialResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state awsCredentialResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	credentialName := state.ID.ValueString()
	params := operations.NewDeleteCredentialParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.DeleteCredentialRequest{CredentialName: &credentialName})
	_, err := r.client.Environments.Operations.DeleteCredential(params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "delete AWS Credential")
		return
	}
}

func (r *awsCredentialResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *awsCredentialResource) updateCredential(ctx context.Context, client *environmentsclient.Environments, plan awsCredentialResourceModel) error {
	params := operations.NewUpdateAwsCredentialParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.UpdateAwsCredentialRequest{
		RoleArn:                plan.RoleArn.ValueStringPointer(),
		Description:            plan.Description.ValueString(),
		CredentialName:         plan.CredentialName.ValueStringPointer(),
		VerifyPermissions:      plan.VerifyPermissions.ValueBoolPointer(),
		SkipOrgPolicyDecisions: plan.SkipOrgPolicyDecisions.ValueBoolPointer(),
	})
	return retry.RetryContext(ctx, credentialCreateRetryDuration, func() *retry.RetryError {
		tflog.Debug(ctx, "Updating AWS credential")
		_, err := client.Operations.UpdateAwsCredential(params)
		if err != nil {
			var envErr *operations.UpdateAwsCredentialDefault
			if errors.As(err, &envErr) {
				if utils.IsRetryableError(envErr.Code()) {
					return retry.RetryableError(err)
				}
			}
			return retry.NonRetryableError(err)
		}
		return nil
	})
}

func credentialNeedsToBeUpdatedAfterCreation(plan awsCredentialResourceModel) bool {
	return (!plan.SkipOrgPolicyDecisions.IsNull() && plan.SkipOrgPolicyDecisions.ValueBool()) || (!plan.VerifyPermissions.IsNull() || plan.VerifyPermissions.ValueBool())
}
