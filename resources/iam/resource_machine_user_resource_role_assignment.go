// Copyright 2024 Cloudera. All Rights Reserved.
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

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ resource.ResourceWithConfigure   = (*machineUserResourceRoleAssignmentResource)(nil)
	_ resource.ResourceWithImportState = (*machineUserResourceRoleAssignmentResource)(nil)
)

func NewMachineUserResourceRoleAssignmentResource() resource.Resource {
	return &machineUserResourceRoleAssignmentResource{}
}

type machineUserResourceRoleAssignmentResource struct {
	client *cdp.Client
}

func (r *machineUserResourceRoleAssignmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *machineUserResourceRoleAssignmentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = MachineUserResourceRoleAssignmentSchema
}

func (r *machineUserResourceRoleAssignmentResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_machine_user_resource_role_assignment"
}

func (r *machineUserResourceRoleAssignmentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *machineUserResourceRoleAssignmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data machineUserResourceRoleAssignmentResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := operations.NewAssignMachineUserResourceRoleParamsWithContext(ctx).WithInput(
		&models.AssignMachineUserResourceRoleRequest{
			MachineUserName: data.MachineUser.ValueStringPointer(),
			ResourceCrn:     data.ResourceCrn.ValueStringPointer(),
			ResourceRoleCrn: data.ResourceRoleCrn.ValueStringPointer(),
		})

	responseOk, err := r.client.Iam.Operations.AssignMachineUserResourceRole(request)
	if err != nil {
		utils.AddIamDiagnosticsError(err, &resp.Diagnostics, "assign Machine User Resource Role")
		return
	}

	if responseOk.Payload != nil {
		data.Id = types.StringValue(data.MachineUser.ValueString() + "_" + data.ResourceCrn.ValueString() + "_" + data.ResourceRoleCrn.ValueString())

		// Save data into Terraform state
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	}
}

func (r *machineUserResourceRoleAssignmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data machineUserResourceRoleAssignmentResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	params := operations.NewListMachineUserAssignedResourceRolesParamsWithContext(ctx)
	params.WithInput(&models.ListMachineUserAssignedResourceRolesRequest{
		MachineUserName: data.MachineUser.ValueStringPointer(),
	})

	machineUser, err := r.client.Iam.Operations.ListMachineUserAssignedResourceRoles(params)
	if err != nil {
		utils.AddIamDiagnosticsError(err, &resp.Diagnostics, "list Machine User Assigned Resource Roles")
		return
	}

	hasAssignedResourceRole := false
	for _, asgn := range machineUser.Payload.ResourceAssignments {
		if *asgn.ResourceCrn == data.ResourceCrn.ValueString() && *asgn.ResourceRoleCrn == data.ResourceRoleCrn.ValueString() {
			resp.State.Set(ctx, &data)
			hasAssignedResourceRole = true
			break
		}
	}

	if !hasAssignedResourceRole {
		resp.Diagnostics.AddError("Resource Role", "Machine User does not have the specified resource role assigned")
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *machineUserResourceRoleAssignmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Warn(ctx, "Update operation is not supported yet.")
}

func (r *machineUserResourceRoleAssignmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data machineUserResourceRoleAssignmentResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := operations.NewUnassignMachineUserResourceRoleParamsWithContext(ctx).WithInput(
		&models.UnassignMachineUserResourceRoleRequest{
			MachineUserName: data.MachineUser.ValueStringPointer(),
			ResourceCrn:     data.ResourceCrn.ValueStringPointer(),
			ResourceRoleCrn: data.ResourceRoleCrn.ValueStringPointer(),
		},
	)

	_, err := r.client.Iam.Operations.UnassignMachineUserResourceRole(request) // void method, does not have any return value
	if err != nil {
		utils.AddIamDiagnosticsError(err, &resp.Diagnostics, "un-assign Machine User Resource Role")
		return
	}
}
