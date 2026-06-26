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
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/df/client/operations"
	dfmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/df/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ resource.Resource                = (*dfAddedReadyflowResource)(nil)
	_ resource.ResourceWithConfigure   = (*dfAddedReadyflowResource)(nil)
	_ resource.ResourceWithImportState = (*dfAddedReadyflowResource)(nil)
)

type dfAddedReadyflowResource struct {
	client *cdp.Client
}

type addedReadyflowModel struct {
	ID                types.String `tfsdk:"id"`
	ReadyflowCrn      types.String `tfsdk:"readyflow_crn"`
	AddedReadyflowCrn types.String `tfsdk:"added_readyflow_crn"`
}

func NewDfAddedReadyflowResource() resource.Resource {
	return &dfAddedReadyflowResource{}
}

func (r *dfAddedReadyflowResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_df_added_readyflow"
}

func (r *dfAddedReadyflowResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *dfAddedReadyflowResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Adds a ReadyFlow to the current CDP account's DataFlow catalog.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"readyflow_crn": schema.StringAttribute{
				Required: true, MarkdownDescription: "The CRN of the ReadyFlow to add.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"added_readyflow_crn": schema.StringAttribute{
				Computed: true, MarkdownDescription: "The CRN of the added ReadyFlow in the account.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *dfAddedReadyflowResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan addedReadyflowModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewAddReadyflowParamsWithContext(ctx).WithInput(&dfmodels.AddReadyflowRequest{
		ReadyflowCrn: plan.ReadyflowCrn.ValueStringPointer(),
	})
	result, err := r.client.Df.Operations.AddReadyflow(params)
	if err != nil {
		resp.Diagnostics.AddError("Error adding ReadyFlow", err.Error())
		return
	}

	detail := result.GetPayload().AddedReadyflowDetail
	plan.AddedReadyflowCrn = types.StringPointerValue(detail.AddedReadyflowCrn)
	plan.ID = types.StringPointerValue(detail.AddedReadyflowCrn)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *dfAddedReadyflowResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state addedReadyflowModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewDescribeAddedReadyflowParamsWithContext(ctx).WithInput(&dfmodels.DescribeAddedReadyflowRequest{
		AddedReadyflowCrn: state.AddedReadyflowCrn.ValueStringPointer(),
	})
	result, err := r.client.Df.Operations.DescribeAddedReadyflow(params)
	if err != nil {
		if strings.Contains(err.Error(), "NOT_FOUND") {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading added ReadyFlow", err.Error())
		return
	}

	detail := result.GetPayload().AddedReadyflowDetail
	state.AddedReadyflowCrn = types.StringPointerValue(detail.AddedReadyflowCrn)
	state.ID = types.StringPointerValue(detail.AddedReadyflowCrn)

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *dfAddedReadyflowResource) Update(ctx context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Warn(ctx, "Update not supported for df_added_readyflow; changes require replacement.")
}

func (r *dfAddedReadyflowResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state addedReadyflowModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewDeleteAddedReadyflowParamsWithContext(ctx).WithInput(&dfmodels.DeleteAddedReadyflowRequest{
		AddedReadyflowCrn: state.AddedReadyflowCrn.ValueStringPointer(),
	})
	if _, err := r.client.Df.Operations.DeleteAddedReadyflow(params); err != nil {
		if !strings.Contains(err.Error(), "NOT_FOUND") {
			resp.Diagnostics.AddError("Error deleting added ReadyFlow", err.Error())
		}
	}
}

func (r *dfAddedReadyflowResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
