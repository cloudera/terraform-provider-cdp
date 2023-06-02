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
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/client/operations"
	iammodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource = &groupResource{}
)

type groupResource struct {
	client *cdp.Client
}

type groupModel struct {
	ID                        types.String `tfsdk:"id"`
	GroupName                 types.String `tfsdk:"group_name"`
	SyncMembershipOnUserLogin types.Bool   `tfsdk:"sync_membership_on_user_login"`
	Crn                       types.String `tfsdk:"crn"`
}

func NewGroupResource() resource.Resource {
	return &groupResource{}
}

func (r *groupResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_group"
}

func (r *groupResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "A group is a named collection of users and machine users. Roles and resource roles can be assigned to a group impacting all members of the group.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"group_name": schema.StringAttribute{
				MarkdownDescription: "The name of the group. This name must be unique. There are certain restrictions on the group name. Refer to the How To > User Management section in the Management Console documentation for the details.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"sync_membership_on_user_login": schema.BoolAttribute{
				MarkdownDescription: "Whether group membership is synced when a user logs in. The default is to sync group membership.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"crn": schema.StringAttribute{
				MarkdownDescription: "The CRN of the resource.",
				Computed:            true,
			},
		},
	}
}

func (r *groupResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *groupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from data
	var data groupModel
	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Iam

	params := operations.NewCreateGroupParamsWithContext(ctx)
	params.WithInput(&iammodels.CreateGroupRequest{
		GroupName:                 data.GroupName.ValueStringPointer(),
		SyncMembershipOnUserLogin: data.SyncMembershipOnUserLogin.ValueBoolPointer(),
	})

	responseOk, err := client.Operations.CreateGroup(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Group",
			"Got error while creating Group: "+err.Error(),
		)
		return
	}

	data.Crn = types.StringPointerValue(responseOk.Payload.Group.Crn)
	data.ID = data.GroupName

	// Save data into Terraform state
	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *groupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state groupModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	sharedGroupRead(ctx, r.client.Iam, &state, &resp.State, &resp.Diagnostics)
}

func sharedGroupRead(ctx context.Context, client *client.Iam, state *groupModel, respState *tfsdk.State, respDiagnostics *diag.Diagnostics) {
	if respDiagnostics.HasError() {
		return
	}

	groupName := state.GroupName.ValueString()
	params := operations.NewListGroupsParamsWithContext(ctx)
	params.WithInput(&iammodels.ListGroupsRequest{GroupNames: []string{groupName}})
	listGroupsOk, err := client.Operations.ListGroups(params)
	if err != nil {
		respDiagnostics.AddError(
			"Error Reading Group",
			"Could not read Group: "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	// Overwrite items with refreshed state
	groups := listGroupsOk.GetPayload().Groups
	if len(groups) == 0 || *groups[0].GroupName != groupName {
		respState.RemoveResource(ctx) // deleted
		return
	}
	g := groups[0]

	state.ID = types.StringPointerValue(g.GroupName)
	state.GroupName = types.StringPointerValue(g.GroupName)
	state.Crn = types.StringPointerValue(g.Crn)

	// Set refreshed state
	respDiagnostics.Append(respState.Set(ctx, &state)...)
	if respDiagnostics.HasError() {
		return
	}
}

func (r *groupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Get current state
	var plan, state groupModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Iam

	if !plan.SyncMembershipOnUserLogin.Equal(state.SyncMembershipOnUserLogin) {
		params := operations.NewUpdateGroupParamsWithContext(ctx)
		// TODO: Below works for false -> true, but does not work for true -> false since swagger generates the
		// the UpdateGroupRequest.SyncMembershipOnUserLogin with `omitempty` which then gets omitted in the request
		// resulting in the server side not seeing the intended change to this field at all. We need to take a look
		// at x-omitempty and maybe change the swagger generation behavior.
		params.WithInput(&iammodels.UpdateGroupRequest{
			GroupName:                 plan.GroupName.ValueStringPointer(),
			SyncMembershipOnUserLogin: plan.SyncMembershipOnUserLogin.ValueBool(),
		})

		_, err := client.Operations.UpdateGroup(params)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Updating Group",
				"Could not update Group: "+state.ID.ValueString()+": "+err.Error(),
			)
			return
		}
		// Save updated data into Terraform state
		resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
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
