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
	_ resource.Resource                = (*dfProjectResource)(nil)
	_ resource.ResourceWithConfigure   = (*dfProjectResource)(nil)
	_ resource.ResourceWithImportState = (*dfProjectResource)(nil)
)

type dfProjectResource struct {
	client *cdp.Client
}

type projectModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	ServiceCrn  types.String `tfsdk:"service_crn"`
	Description types.String `tfsdk:"description"`
	Crn         types.String `tfsdk:"crn"`
	Revision    types.Int32  `tfsdk:"revision"`
}

func NewDfProjectResource() resource.Resource {
	return &dfProjectResource{}
}

func (r *dfProjectResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_df_project"
}

func (r *dfProjectResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *dfProjectResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Creates a DataFlow project within a DataFlow service.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Required: true, MarkdownDescription: "The project name.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"service_crn": schema.StringAttribute{
				Required: true, MarkdownDescription: "The CRN of the DataFlow service.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{
				Optional: true, MarkdownDescription: "The project description.",
			},
			"crn": schema.StringAttribute{
				Computed: true, MarkdownDescription: "The CRN of the project.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"revision": schema.Int32Attribute{
				Computed: true, MarkdownDescription: "The project revision number.",
			},
		},
	}
}

func (r *dfProjectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan projectModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := &dfmodels.CreateProjectRequest{
		Name:       plan.Name.ValueStringPointer(),
		ServiceCrn: plan.ServiceCrn.ValueStringPointer(),
	}
	if !plan.Description.IsNull() {
		input.Description = plan.Description.ValueString()
	}

	params := operations.NewCreateProjectParamsWithContext(ctx).WithInput(input)
	result, err := r.client.Df.Operations.CreateProject(params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating DataFlow project", err.Error())
		return
	}

	project := result.GetPayload().Project
	plan.Crn = types.StringPointerValue(project.Crn)
	plan.ID = types.StringPointerValue(project.Crn)
	plan.Revision = types.Int32PointerValue(project.Revision)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *dfProjectResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state projectModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewDescribeProjectParamsWithContext(ctx).WithInput(&dfmodels.DescribeProjectRequest{
		ProjectCrn: state.Crn.ValueStringPointer(),
	})
	result, err := r.client.Df.Operations.DescribeProject(params)
	if err != nil {
		if strings.Contains(err.Error(), "NOT_FOUND") {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading DataFlow project", err.Error())
		return
	}

	project := result.GetPayload().Project
	if project == nil {
		resp.State.RemoveResource(ctx)
		return
	}
	state.Crn = types.StringPointerValue(project.Crn)
	state.ID = types.StringPointerValue(project.Crn)
	state.Name = types.StringPointerValue(project.Name)
	state.Revision = types.Int32PointerValue(project.Revision)
	if project.Description != "" {
		state.Description = types.StringValue(project.Description)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *dfProjectResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan projectModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state projectModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := &dfmodels.UpdateProjectRequest{
		ProjectCrn: state.Crn.ValueStringPointer(),
	}
	if !plan.Description.IsNull() {
		input.ProjectDescription = plan.Description.ValueString()
	}

	params := operations.NewUpdateProjectParamsWithContext(ctx).WithInput(input)
	_, err := r.client.Df.Operations.UpdateProject(params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating DataFlow project", err.Error())
		return
	}

	plan.Crn = state.Crn
	plan.ID = state.ID
	plan.Revision = types.Int32Value(state.Revision.ValueInt32() + 1)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *dfProjectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state projectModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewDeleteProjectParamsWithContext(ctx).WithInput(&dfmodels.DeleteProjectRequest{
		ProjectCrn:      state.Crn.ValueStringPointer(),
		ProjectRevision: state.Revision.ValueInt32Pointer(),
	})
	if _, err := r.client.Df.Operations.DeleteProject(params); err != nil {
		if !strings.Contains(err.Error(), "NOT_FOUND") {
			resp.Diagnostics.AddError("Error deleting DataFlow project", err.Error())
		}
	}
}

func (r *dfProjectResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
