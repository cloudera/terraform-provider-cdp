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
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
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
		"polling_options": schema.SingleNestedAttribute{
			MarkdownDescription: "Polling related configuration options that could specify various values that will be used during CDP resource creation.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"async": schema.BoolAttribute{
					MarkdownDescription: "Boolean value that specifies if Terraform should wait for resource creation/deletion.",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
					PlanModifiers: []planmodifier.Bool{
						boolplanmodifier.UseStateForUnknown(),
					},
				},
				"polling_timeout": schema.Int64Attribute{
					MarkdownDescription: "Timeout value in minutes that specifies for how long should the polling go for resource creation/deletion.",
					Default:             int64default.StaticInt64(90),
					Computed:            true,
					Optional:            true,
				},
				"call_failure_threshold": schema.Int64Attribute{
					MarkdownDescription: "Threshold value that specifies how many times should a single call failure happen before giving up the polling.",
					Default:             int64default.StaticInt64(3),
					Computed:            true,
					Optional:            true,
				},
			},
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
	ID               types.String          `tfsdk:"id"`
	EnvironmentNames types.Set             `tfsdk:"environment_names"`
	PollingOptions   *utils.PollingOptions `tfsdk:"polling_options"`
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
	res, err := client.Operations.SyncAllUsers(params)
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

	opID := res.Payload.OperationID
	tflog.Debug(ctx, fmt.Sprintf("User sync operation ID: %s", *opID))
	if !(state.PollingOptions != nil && state.PollingOptions.Async.ValueBool()) {
		tflog.Debug(ctx, "User sync polling starts")
		err = waitForUserSync(*opID, time.Hour*1, callFailureThreshold, r.client.Environments, ctx, state.PollingOptions)
		if err != nil {
			return
		}
	}
}

func waitForUserSync(opID string, fallbackTimeout time.Duration, callFailureThresholdDefault int, client *client.Environments, ctx context.Context, pollingOptions *utils.PollingOptions) error {
	timeout, err := utils.CalculateTimeoutOrDefault(ctx, pollingOptions, fallbackTimeout)
	if err != nil {
		return err
	}
	callFailureThreshold, failureThresholdError := utils.CalculateCallFailureThresholdOrDefault(ctx, pollingOptions, callFailureThresholdDefault)
	if failureThresholdError != nil {
		return failureThresholdError
	}
	callFailedCount := 0
	stateConf := &retry.StateChangeConf{
		Pending: []string{"NEVER_RUN",
			"REQUESTED",
			"REJECTED",
			"RUNNING",
			"COMPLETED",
			"FAILED",
			"TIMEDOUT"},
		Target:       []string{"COMPLETED"},
		Delay:        5 * time.Second,
		Timeout:      *timeout,
		PollInterval: 10 * time.Second,
		Refresh: func() (interface{}, string, error) {
			tflog.Debug(ctx, fmt.Sprintf("About to get sync status for operationID %s", opID))
			params := operations.NewSyncStatusParamsWithContext(ctx)
			params.WithInput(&environmentsmodels.SyncStatusRequest{OperationID: &opID})
			resp, err := client.Operations.SyncStatus(params)
			if err != nil {
				if isEnvNotFoundError(err) {
					tflog.Debug(ctx, fmt.Sprintf("Recoverable error getting user sync status: %s", err))
					callFailedCount = 0
					return nil, "", nil
				}
				callFailedCount++
				if callFailedCount <= callFailureThreshold {
					tflog.Warn(ctx, fmt.Sprintf("Error getting user sync status with call failure due to [%s] but threshold limit is not reached yet (%d out of %d).", err.Error(), callFailedCount, callFailureThreshold))
					return nil, "", nil
				}
				tflog.Error(ctx, fmt.Sprintf("Error getting user sync status (due to: %s) and call failure threshold limit exceeded.", err))
				return nil, "", err
			}
			callFailedCount = 0
			tflog.Info(ctx, fmt.Sprintf("User sync status: %s", resp.GetPayload().Status))
			return checkUserSyncResponseStatusForError(resp)
		},
	}
	_, err = stateConf.WaitForStateContext(ctx)

	return err
}

func checkUserSyncResponseStatusForError(resp *operations.SyncStatusOK) (interface{}, string, error) {
	if utils.ContainsAsSubstring([]string{"FAILED", "ERROR"}, string(resp.GetPayload().Status)) {
		return nil, "", fmt.Errorf("unexpected user sync status status: %s. ", resp.GetPayload().Status)
	}
	return resp, string(resp.GetPayload().Status), nil
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
