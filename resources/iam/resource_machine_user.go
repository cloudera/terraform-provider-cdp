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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*machineUserResource)(nil)

func NewMachineUserResource() resource.Resource {
	return &machineUserResource{}
}

type machineUserResource struct {
	client *cdp.Client
}

func (r *machineUserResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_machine_user"
}

func (r *machineUserResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = machineUserSchema
}

func (r *machineUserResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *machineUserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data machineUserResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Iam

	params := operations.NewCreateMachineUserParamsWithContext(ctx)
	params.WithInput(&models.CreateMachineUserRequest{
		MachineUserName: data.Name.ValueStringPointer(),
	})

	responseOk, err := client.Operations.CreateMachineUser(params)
	if err != nil {
		utils.AddIamDiagnosticsError(err, &resp.Diagnostics, "create Machine User")
		return
	}

	if responseOk.Payload.MachineUser != nil {
		mu := responseOk.Payload.MachineUser
		muRespToModel(ctx, mu, &data)

		// Save data into Terraform state
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	}
}

func (r *machineUserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data machineUserResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	client := r.client.Iam

	params := operations.NewListMachineUsersParamsWithContext(ctx)
	params.WithInput(&models.ListMachineUsersRequest{
		MachineUserNames: []string{data.Name.ValueString()},
	})

	responseOk, err := client.Operations.ListMachineUsers(params)
	if err != nil {
		utils.AddIamDiagnosticsError(err, &resp.Diagnostics, "read Machine User")
		if d, ok := err.(*operations.CreateMachineUserDefault); ok && d.GetPayload() != nil && d.GetPayload().Code == "NOT_FOUND" {
			resp.Diagnostics.AddWarning("Resource not found on provider", "Machine User not found, removing from state.")
			tflog.Warn(ctx, "Machine User not found, removing from state", map[string]interface{}{
				"id": data.Id,
			})
			resp.State.RemoveResource(ctx)
		}
		return
	}

	if len(responseOk.Payload.MachineUsers) == 1 && responseOk.Payload.MachineUsers[0] != nil {
		mu := responseOk.Payload.MachineUsers[0]
		muRespToModel(ctx, mu, &data)

		// Save data into Terraform state
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	} else {
		if d, ok := err.(*operations.CreateMachineUserDefault); ok && d.GetPayload() != nil && d.GetPayload().Code == "NOT_FOUND" {
			resp.Diagnostics.AddWarning("Resource not found on provider", "Machine User not found, removing from state.")
			tflog.Warn(ctx, "Machine User not found, removing from state", map[string]interface{}{
				"id": data.Id,
			})
			resp.State.RemoveResource(ctx)
			return
		}
	}
}

func (r *machineUserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Warn(ctx, "Update operation is not implemented yet.")
}

func (r *machineUserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data machineUserResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	client := r.client.Iam

	params := operations.NewDeleteMachineUserParamsWithContext(ctx)
	params.WithInput(&models.DeleteMachineUserRequest{
		MachineUserName: data.Name.ValueStringPointer(),
	})

	_, err := client.Operations.DeleteMachineUser(params)
	if err != nil {
		utils.AddIamDiagnosticsError(err, &resp.Diagnostics, "delete Machine User")
		return
	}
}
