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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource = &azureCredentialResource{}
)

type azureCredentialResource struct {
	client *cdp.Client
}

type AppBased struct {
	ApplicationID types.String `tfsdk:"application_id"`
	SecretKey     types.String `tfsdk:"secret_key"`
}

type azureCredentialResourceModel struct {
	ID             types.String `tfsdk:"id"`
	CredentialName types.String `tfsdk:"credential_name"`
	SubscriptionID types.String `tfsdk:"subscription_id"`
	TenantID       types.String `tfsdk:"tenant_id"`
	AppBased       *AppBased    `tfsdk:"app_based"`
	Crn            types.String `tfsdk:"crn"`
	Description    types.String `tfsdk:"description"`
}

func NewAzureCredentialResource() resource.Resource {
	return &azureCredentialResource{}
}

func (r *azureCredentialResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_azure_credential"
}

func (r *azureCredentialResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The Azure credential is used for authorization  to provision resources such as compute instances within your cloud provider account.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"credential_name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the CDP credential.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"subscription_id": schema.StringAttribute{
				Description: "The Azure subscription ID. Required for secret based credentials and should look like the following example: a8d4457d-310v-41p6-sc53-14g8d733e514",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required: true,
			},
			"tenant_id": schema.StringAttribute{
				Description: "The Azure AD tenant ID for the Azure subscription. Required for secret based credentials and should look like the following example: b10u3481-2451-10ba-7sfd-9o2d1v60185d",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"app_based": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"application_id": schema.StringAttribute{
						Description: "The ID of the application registered in Azure.",
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Required:  true,
						Sensitive: false,
					},
					"secret_key": schema.StringAttribute{
						Description: "The client secret key (also referred to as application password) for the registered application.",
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Required:  true,
						Sensitive: true,
					},
				},
			},
			"description": schema.StringAttribute{
				Description: "A description for the credential.",
				Optional:    true,
			},
			"crn": schema.StringAttribute{
				Description: "The CRN of the credential.",
				Computed:    true,
			},
		},
	}
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
		resp.Diagnostics.AddError(
			"Error creating Azure credential",
			"Got error while creating Azure credential: "+err.Error(),
		)
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
	credentialName := state.ID.ValueString()
	params := operations.NewListCredentialsParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.ListCredentialsRequest{CredentialName: credentialName})
	listCredentialsResp, err := r.client.Environments.Operations.ListCredentials(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Azure Credentials",
			"Could not read Azure Credentials: "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	// Overwrite items with refreshed state
	credentials := listCredentialsResp.GetPayload().Credentials
	if len(credentials) == 0 || *credentials[0].CredentialName != credentialName {
		resp.Diagnostics.AddError(
			"Summary: Credential could not found.",
			"Detailed: Credential not found, removing from state.")
		resp.State.RemoveResource(ctx) // deleted
		return
	}
	c := credentials[0]

	state.ID = types.StringPointerValue(c.CredentialName)
	state.CredentialName = types.StringPointerValue(c.CredentialName)
	state.Crn = types.StringPointerValue(c.Crn)
	state.AppBased.ApplicationID = types.StringValue(c.AzureCredentialProperties.AppID)
	state.SubscriptionID = types.StringValue(c.AzureCredentialProperties.SubscriptionID)

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

	params := operations.NewDeleteCredentialParams().WithInput(&environmentsmodels.DeleteCredentialRequest{CredentialName: state.ID.ValueStringPointer()})
	_, err := r.client.Environments.Operations.DeleteCredential(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Azure credential",
			"Could not delete Azure credential due to: "+err.Error(),
		)
		return
	}

}
