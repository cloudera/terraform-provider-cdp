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
	_ resource.ResourceWithConfigure   = (*machineUserGroupAssignmentResource)(nil)
	_ resource.ResourceWithImportState = (*machineUserGroupAssignmentResource)(nil)
)

func NewMachineUserGroupAssignmentResource() resource.Resource {
	return &machineUserGroupAssignmentResource{}
}

type machineUserGroupAssignmentResource struct {
	client *cdp.Client
}

func (r *machineUserGroupAssignmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *machineUserGroupAssignmentResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_machine_user_group_assignment"
}

func (r *machineUserGroupAssignmentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = machineUserGroupAssignmentSchema
}

func (r *machineUserGroupAssignmentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *machineUserGroupAssignmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data machineUserGroupAssignmentResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Iam

	params := operations.NewAddMachineUserToGroupParamsWithContext(ctx)
	params.WithInput(&models.AddMachineUserToGroupRequest{
		MachineUserName: data.MachineUser.ValueStringPointer(),
		GroupName:       data.Group.ValueStringPointer(),
	})

	responseOk, err := client.Operations.AddMachineUserToGroup(params)
	if err != nil {
		utils.AddIamDiagnosticsError(err, &resp.Diagnostics, "create machine user group assignment")
		return
	}

	if responseOk.Payload != nil {
		data.Id = types.StringValue(data.MachineUser.ValueString() + "_" + data.Group.ValueString())

		// Save data into Terraform state
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	}
}

func (r *machineUserGroupAssignmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data machineUserGroupAssignmentResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	client := r.client.Iam

	params := operations.NewListGroupsForMachineUserParamsWithContext(ctx)
	params.WithInput(&models.ListGroupsForMachineUserRequest{
		MachineUserName: data.MachineUser.ValueStringPointer(),
	})

	responseOk, err := client.Operations.ListGroupsForMachineUser(params)
	if err != nil {
		utils.AddIamDiagnosticsError(err, &resp.Diagnostics, "read machine user group assignment")
		if d, ok := err.(*operations.ListGroupsForMachineUserDefault); ok && d.GetPayload() != nil && d.GetPayload().Code == "NOT_FOUND" {
			logAndRemoveNotFoundAssignment(ctx, resp, &data)
		}
		return
	}

	grParams := operations.NewListGroupsParamsWithContext(ctx)
	grParams.WithInput(&models.ListGroupsRequest{
		GroupNames: []string{data.Group.ValueString()},
	})

	grRespOk, err := client.Operations.ListGroups(grParams)
	if err != nil {
		utils.AddIamDiagnosticsError(err, &resp.Diagnostics, "read machine user group assignment")
		if d, ok := err.(*operations.ListGroupsDefault); ok && d.GetPayload() != nil && d.GetPayload().Code == "NOT_FOUND" {
			logAndRemoveNotFoundAssignment(ctx, resp, &data)
		}
		return
	}

	if len(grRespOk.Payload.Groups) == 1 && grRespOk.Payload.Groups[0] != nil {
		grCrn := grRespOk.Payload.Groups[0].Crn
		found := false
		for _, v := range responseOk.Payload.GroupCrns {
			if *grCrn == v {
				found = true
				break
			}
		}
		if !found {
			logAndRemoveNotFoundAssignment(ctx, resp, &data)
		}
	} else {
		logAndRemoveNotFoundAssignment(ctx, resp, &data)
	}
}

func logAndRemoveNotFoundAssignment(ctx context.Context, resp *resource.ReadResponse, data *machineUserGroupAssignmentResourceModel) {
	resp.Diagnostics.AddWarning("Resource not found on provider", "Machine user group assignment not found, removing from state.")
	tflog.Warn(ctx, "Machine user group assignment not found, removing from state", map[string]interface{}{
		"id": data.Id,
	})
	resp.State.RemoveResource(ctx)
}

func (r *machineUserGroupAssignmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Warn(ctx, "Update operation is not implemented yet.")
}

func (r *machineUserGroupAssignmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data machineUserGroupAssignmentResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	client := r.client.Iam

	params := operations.NewRemoveMachineUserFromGroupParamsWithContext(ctx)
	params.WithInput(&models.RemoveMachineUserFromGroupRequest{
		MachineUserName: data.MachineUser.ValueStringPointer(),
		GroupName:       data.Group.ValueStringPointer(),
	})

	_, err := client.Operations.RemoveMachineUserFromGroup(params)
	if err != nil {
		utils.AddIamDiagnosticsError(err, &resp.Diagnostics, "delete machine user group assignment")
		return
	}
}
