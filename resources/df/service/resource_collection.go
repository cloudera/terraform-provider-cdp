// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package service

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/df/client/operations"
	dfmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/df/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ resource.Resource                = (*dfCollectionResource)(nil)
	_ resource.ResourceWithConfigure   = (*dfCollectionResource)(nil)
	_ resource.ResourceWithImportState = (*dfCollectionResource)(nil)
)

type dfCollectionResource struct {
	client *cdp.Client
}

type collectionModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Crn         types.String `tfsdk:"crn"`
}

func NewDfCollectionResource() resource.Resource {
	return &dfCollectionResource{}
}

func (r *dfCollectionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_df_collection"
}

func (r *dfCollectionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *dfCollectionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Creates a DataFlow catalog collection for organizing flow definitions.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Required: true, MarkdownDescription: "The collection name.",
			},
			"description": schema.StringAttribute{
				Optional: true, MarkdownDescription: "The collection description.",
			},
			"crn": schema.StringAttribute{
				Computed: true, MarkdownDescription: "The CRN of the collection.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *dfCollectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan collectionModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := &dfmodels.CreateCollectionRequest{
		Name: plan.Name.ValueStringPointer(),
	}
	if !plan.Description.IsNull() {
		input.Description = plan.Description.ValueString()
	}

	params := operations.NewCreateCollectionParamsWithContext(ctx).WithInput(input)
	result, err := r.client.Df.Operations.CreateCollection(params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating DataFlow collection", err.Error())
		return
	}

	col := result.GetPayload().CatalogCollection
	plan.Crn = types.StringValue(col.Crn)
	plan.ID = types.StringValue(col.Crn)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *dfCollectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state collectionModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewDescribeCollectionParamsWithContext(ctx).WithInput(&dfmodels.DescribeCollectionRequest{
		CatalogCollectionCrn: state.Crn.ValueStringPointer(),
	})
	result, err := r.client.Df.Operations.DescribeCollection(params)
	if err != nil {
		if strings.Contains(err.Error(), "NOT_FOUND") {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading DataFlow collection", err.Error())
		return
	}

	col := result.GetPayload().CatalogCollection
	if col == nil {
		resp.State.RemoveResource(ctx)
		return
	}
	state.Crn = types.StringValue(col.Crn)
	state.ID = types.StringValue(col.Crn)
	state.Name = types.StringValue(col.Name)
	if col.Description != "" {
		state.Description = types.StringValue(col.Description)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *dfCollectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan collectionModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state collectionModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := &dfmodels.UpdateCollectionRequest{
		CatalogCollectionCrn: state.Crn.ValueStringPointer(),
		Name:                 plan.Name.ValueString(),
	}
	if !plan.Description.IsNull() {
		input.Description = plan.Description.ValueString()
	}

	params := operations.NewUpdateCollectionParamsWithContext(ctx).WithInput(input)
	_, err := r.client.Df.Operations.UpdateCollection(params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating DataFlow collection", err.Error())
		return
	}

	plan.Crn = state.Crn
	plan.ID = state.ID
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *dfCollectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state collectionModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewDeleteCollectionParamsWithContext(ctx).WithInput(&dfmodels.DeleteCollectionRequest{
		CatalogCollectionCrn: state.Crn.ValueStringPointer(),
	})
	if _, err := r.client.Df.Operations.DeleteCollection(params); err != nil {
		if !strings.Contains(err.Error(), "NOT_FOUND") {
			resp.Diagnostics.AddError("Error deleting DataFlow collection", err.Error())
		}
	}
}

func (r *dfCollectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
