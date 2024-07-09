// Copyright 2024 Cloudera. All Rights Reserved.
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
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource = &userSyncResource{}
)

var userSyncSchema = schema.Schema{
	MarkdownDescription: "Synchronizes environments with all users and groups state with CDP.",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"environment_names": schema.SetAttribute{
			MarkdownDescription: "List of environments to be synced. If not present, all environments will be synced.",
			ElementType:         types.StringType,
			Optional:            true,
		},
	},
}

type userSyncResource struct {
	client *cdp.Client
}

func NewUserSyncResource() resource.Resource {
	return &userSyncResource{}
}

func (r *userSyncResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_user_sync"
}

type userSyncResourceModel struct {
	ID types.String `tfsdk:"id"`

	EnvironmentNames types.Set `tfsdk:"environment_names"`
}

func toSyncAllUsersRequest(ctx context.Context, model *userSyncResourceModel, diag *diag.Diagnostics) *environmentsmodels.SyncAllUsersRequest {
	req := &environmentsmodels.SyncAllUsersRequest{}
	req.EnvironmentNames = utils.FromSetValueToStringList(model.EnvironmentNames)
	return req
}

func (r *userSyncResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = userSyncSchema
}

func (r *userSyncResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *userSyncResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state userSyncResourceModel
	diags := req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Got Error while trying to set plan")
		return
	}

	client := r.client.Environments

	params := operations.NewSyncAllUsersParamsWithContext(ctx)
	params.WithInput(toSyncAllUsersRequest(ctx, &state, &resp.Diagnostics))
	_, err := client.Operations.SyncAllUsers(params)
	if err != nil {
		if isSyncAllUsersNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Error in sync all users",
				"An environment not found: "+state.EnvironmentNames.String(),
			)
			return
		}
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "sync all users")
		return
	}

	state.ID = types.StringValue(uuid.New().String())

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func isSyncAllUsersNotFoundError(err error) bool {
	if envErr, ok := err.(*operations.SyncAllUsersDefault); ok {
		if cdp.IsEnvironmentsError(envErr.GetPayload(), "NOT_FOUND", "") {
			return true
		}
	}
	return false
}

func (r *userSyncResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Info(ctx, "Nothing to do in read phase")
}

func (r *userSyncResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state userSyncResourceModel
	diags := req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Got Error while trying to set plan")
		return
	}

	client := r.client.Environments

	params := operations.NewSyncAllUsersParamsWithContext(ctx)
	params.WithInput(toSyncAllUsersRequest(ctx, &state, &resp.Diagnostics))
	_, err := client.Operations.SyncAllUsers(params)
	if err != nil {
		if isSyncAllUsersNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Error in sync all users",
				"An environment not found: "+state.EnvironmentNames.String(),
			)
			return
		}
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "sync all users")
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *userSyncResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Info(ctx, "Nothing to do in delete phase")
}
