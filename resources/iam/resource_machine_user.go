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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource = &machineUserResource{}
)

type machineUserResource struct {
	client *cdp.Client
}

type machineUserModel struct {
	ID               types.String `tfsdk:"id"`
	MachineUserName  types.String `tfsdk:"machine_user_name"`
	Crn              types.String `tfsdk:"crn"`
	CreationDate     types.String `tfsdk:"creation_date"`
	Status           types.String `tfsdk:"status"`
	WorkloadUsername types.String `tfsdk:"workload_username"`
}

func NewMachineUserResource() resource.Resource {
	return &machineUserResource{}
}

func (r *machineUserResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_machine_user"
}

func (r *machineUserResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "A machine user account provides programmatic access to CDP.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"machine_user_name": schema.StringAttribute{
				MarkdownDescription: "The machine user name.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"crn": schema.StringAttribute{
				MarkdownDescription: "The CRN of the user.",
				Computed:            true,
			},
			"creation_date": schema.StringAttribute{
				MarkdownDescription: "The date when this machine user was created.",
				Computed:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: "The current status of the machine user.",
				Computed:            true,
			},
			"workload_username": schema.StringAttribute{
				MarkdownDescription: "The username used in all the workload clusters of the machine user.",
				Computed:            true,
			},
		},
	}
}

func (r *machineUserResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *machineUserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from data
	var data machineUserModel
	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Iam

	params := operations.NewCreateMachineUserParamsWithContext(ctx)
	params.WithInput(&iammodels.CreateMachineUserRequest{
		MachineUserName: data.MachineUserName.ValueStringPointer(),
	})

	responseOk, err := client.Operations.CreateMachineUser(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Machine User",
			"Got error while creating Machine User: "+err.Error(),
		)
		return
	}

	data.Crn = types.StringPointerValue(responseOk.Payload.MachineUser.Crn)
	data.ID = data.MachineUserName
	data.Status = types.StringValue(responseOk.Payload.MachineUser.Status)
	data.CreationDate = types.StringValue(responseOk.Payload.MachineUser.CreationDate.String())
	data.WorkloadUsername = types.StringValue(responseOk.Payload.MachineUser.WorkloadUsername)

	// Save data into Terraform state
	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func sharedMachineUserRead(ctx context.Context, client *client.Iam, state *machineUserModel, respState *tfsdk.State, respDiagnostics *diag.Diagnostics) {
	if respDiagnostics.HasError() {
		return
	}

	machineUserName := state.MachineUserName.ValueString()
	params := operations.NewListMachineUsersParamsWithContext(ctx)
	params.WithInput(&iammodels.ListMachineUsersRequest{MachineUserNames: []string{machineUserName}})
	listMachineUsersOk, err := client.Operations.ListMachineUsers(params)
	if err != nil {
		respDiagnostics.AddError(
			"Error Reading Machine User",
			"Could not read Machine User: "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	machineUsers := listMachineUsersOk.GetPayload().MachineUsers
	if len(machineUsers) == 0 || *machineUsers[0].MachineUserName != machineUserName {
		respState.RemoveResource(ctx) // deleted
		return
	}
	mu := machineUsers[0]

	state.Crn = types.StringPointerValue(mu.Crn)
	state.ID = state.MachineUserName
	state.Status = types.StringValue(mu.Status)
	state.CreationDate = types.StringValue(mu.CreationDate.String())
	state.WorkloadUsername = types.StringValue(mu.WorkloadUsername)
}

func (r *machineUserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state machineUserModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	sharedMachineUserRead(ctx, r.client.Iam, &state, &resp.State, &resp.Diagnostics)
}

func (r *machineUserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Get current state
	var plan, state machineUserModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *machineUserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state machineUserModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Iam

	MachineUserName := state.ID.ValueString()
	params := operations.NewDeleteMachineUserParamsWithContext(ctx)
	params.WithInput(&iammodels.DeleteMachineUserRequest{MachineUserName: &MachineUserName})
	_, err := client.Operations.DeleteMachineUser(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting Machine User",
			"Could not delete Machine User, unexpected error: "+err.Error(),
		)
		return
	}
}
