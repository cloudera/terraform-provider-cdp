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
	environments "github.com/cloudera/terraform-provider-cdp/resources"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource = &azureImageTermsResource{}
)

type azureImageTermsResource struct {
	client *cdp.Client
}

<<<<<<< HEAD
=======
type azureImageTermsResourceModel struct {
	ID       types.String `tfsdk:"id"`
	Accepted types.Bool   `tfsdk:"accepted"`
}

>>>>>>> 6f63600 (CDPCP-10777 Added Azure Image Terms Policy resource)
func NewAzureImageTermsResource() resource.Resource {
	return &azureImageTermsResource{}
}

func (r *azureImageTermsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_azure_image_terms"
}

func (r *azureImageTermsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = environments.AzureImageTermsPolicySchema
}

func (r *azureImageTermsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *azureImageTermsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data azureImageTermsResourceModel
	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Environments

	params := operations.NewUpdateAzureImageTermsPolicyParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.UpdateAzureImageTermsPolicyRequest{
		Accepted: data.Accepted.ValueBoolPointer(),
	})

	_, err := client.Operations.UpdateAzureImageTermsPolicy(params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "create Azure Image Terms Policy")
		return
	}

	data.ID = types.StringValue(uuid.New().String())

	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

type getAzureImageTermsPolicyRequest struct{}

func (r *azureImageTermsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state azureImageTermsResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewGetAzureImageTermsPolicyParamsWithContext(ctx)
	params.WithInput(getAzureImageTermsPolicyRequest{})
	getPolicyResp, err := r.client.Environments.Operations.GetAzureImageTermsPolicy(params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "read Azure Image Terms Policy")
		return
	}

	state.Accepted = types.BoolPointerValue(getPolicyResp.Payload.Accepted)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *azureImageTermsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data azureImageTermsResourceModel
	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Environments

	params := operations.NewUpdateAzureImageTermsPolicyParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.UpdateAzureImageTermsPolicyRequest{
		Accepted: data.Accepted.ValueBoolPointer(),
	})

	_, err := client.Operations.UpdateAzureImageTermsPolicy(params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "update Azure Image Terms Policy")
		return
	}

	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *azureImageTermsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data azureImageTermsResourceModel
	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Environments

	defaultValue := false

	params := operations.NewUpdateAzureImageTermsPolicyParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.UpdateAzureImageTermsPolicyRequest{
		Accepted: &defaultValue,
	})

	_, err := client.Operations.UpdateAzureImageTermsPolicy(params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "delete Azure Image Terms Policy")
		return
	}
}
