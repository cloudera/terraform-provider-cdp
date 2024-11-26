// Copyright 2023 Cloudera. All Rights Reserved.
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
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ resource.ResourceWithConfigure   = &awsEnvironmentResource{}
	_ resource.ResourceWithImportState = &awsEnvironmentResource{}
)

type awsEnvironmentResource struct {
	client *cdp.Client
}

func (r *awsEnvironmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func NewAwsEnvironmentResource() resource.Resource {
	return &awsEnvironmentResource{}
}

func (r *awsEnvironmentResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_aws_environment"
}

func (r *awsEnvironmentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = AwsEnvironmentSchema
}

func (r *awsEnvironmentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *awsEnvironmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data awsEnvironmentResourceModel
	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Got Error while trying to set plan")
		return
	}

	client := r.client.Environments

	params := operations.NewCreateAWSEnvironmentParamsWithContext(ctx)
	params.WithInput(ToAwsEnvironmentRequest(ctx, &data))

	responseOk, err := client.Operations.CreateAWSEnvironment(params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "create AWS Environment")
		return
	}

	envResp := responseOk.Payload.Environment
	toAwsEnvironmentResource(ctx, envResp, &data, data.PollingOptions, &resp.Diagnostics)

	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	descEnvResp, err := describeEnvironmentWithDiagnosticHandle(data.EnvironmentName.ValueString(), data.ID.ValueString(), ctx, r.client, &resp.Diagnostics, &resp.State)
	if err != nil {
		return
	}
	if !(data.PollingOptions != nil && data.PollingOptions.Async.ValueBool()) {
		stateSaver := func(env *environmentsmodels.Environment) {
			toAwsEnvironmentResource(ctx, utils.LogEnvironmentSilently(ctx, env, describeLogPrefix), &data, data.PollingOptions, &resp.Diagnostics)
			diags = resp.State.Set(ctx, data)
			resp.Diagnostics.Append(diags...)
		}
		descEnvResp, err = waitForCreateEnvironmentWithDiagnosticHandle(ctx, r.client, data.ID.ValueString(), data.EnvironmentName.ValueString(), resp, data.PollingOptions, stateSaver)
		if err != nil {
			return
		}
	}

	toAwsEnvironmentResource(ctx, utils.LogEnvironmentSilently(ctx, descEnvResp, describeLogPrefix), &data, data.PollingOptions, &resp.Diagnostics)
	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *awsEnvironmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state awsEnvironmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	envName := state.EnvironmentName.ValueString()
	if len(envName) == 0 {
		envName = state.ID.ValueString()
	}
	env, err := describeEnvironmentWithDiagnosticHandle(envName, state.ID.ValueString(), ctx, r.client, &resp.Diagnostics, &resp.State)
	if err != nil {
		return
	}
	toAwsEnvironmentResource(ctx, env, &state, state.PollingOptions, &resp.Diagnostics)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *awsEnvironmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan awsEnvironmentResourceModel
	var state awsEnvironmentResourceModel
	planDiags := req.Plan.Get(ctx, &plan)
	var stateDiags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(planDiags...)
	resp.Diagnostics.Append(stateDiags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Got Error while trying to set plan")
		return
	}

	updateEnvironment(ctx, &plan, &state, r.client.Environments, resp)
	updateFreeIpa(ctx, &plan, &state, r.client.Environments, resp)

	stateDiags = resp.State.Set(ctx, state)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.State.Set(ctx, state)
}

func (r *awsEnvironmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state awsEnvironmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	cascading := state.Cascading.ValueBool()
	if state.Cascading.IsNull() {
		cascading = true
	}
	if err := deleteEnvironmentWithDiagnosticHandle(state.EnvironmentName.ValueString(), cascading, ctx, r.client, resp, state.PollingOptions); err != nil {
		return
	}
}

func toAwsEnvironmentResource(ctx context.Context, env *environmentsmodels.Environment, model *awsEnvironmentResourceModel, pollingOptions *utils.PollingOptions, diags *diag.Diagnostics) {
	utils.LogEnvironmentSilently(ctx, env, "")
	model.ID = types.StringPointerValue(env.Crn)
	if env.AwsDetails != nil {
		model.S3GuardTableName = types.StringValue(env.AwsDetails.S3GuardTableName)
	}
	model.CredentialName = types.StringPointerValue(env.CredentialName)
	model.Crn = types.StringPointerValue(env.Crn)
	model.Description = types.StringValue(env.Description)
	model.EnvironmentName = types.StringPointerValue(env.EnvironmentName)
	model.PollingOptions = pollingOptions
	if env.LogStorage != nil {
		if env.LogStorage.AwsDetails != nil {
			model.LogStorage = &AWSLogStorage{
				InstanceProfile:     types.StringValue(env.LogStorage.AwsDetails.InstanceProfile),
				StorageLocationBase: types.StringValue(env.LogStorage.AwsDetails.StorageLocationBase),
			}
			if env.BackupStorage != nil {
				if env.BackupStorage.AwsDetails != nil {
					model.LogStorage.BackupStorageLocationBase = types.StringValue(env.BackupStorage.AwsDetails.StorageLocationBase)
				}

			}
		}
	}
	diags.Append(*FreeIpaResponseToModel(env.Freeipa, &model.FreeIpa, ctx)...)
	if env.Network != nil {
		model.EndpointAccessGatewayScheme = types.StringValue(env.Network.EndpointAccessGatewayScheme)
		model.NetworkCidr = types.StringValue(env.Network.NetworkCidr)
		if env.Network.EndpointAccessGatewaySubnetIds != nil {
			var eagSubnetids types.Set
			if len(env.Network.EndpointAccessGatewaySubnetIds) > 0 {
				var eagSnDiags diag.Diagnostics
				eagSubnetids, eagSnDiags = types.SetValueFrom(ctx, types.StringType, env.Network.EndpointAccessGatewaySubnetIds)
				diags.Append(eagSnDiags...)
			} else {
				eagSubnetids = types.SetNull(types.StringType)
			}
			model.EndpointAccessGatewaySubnetIds = eagSubnetids
		}
		if env.Network.Aws != nil {
			model.VpcID = types.StringPointerValue(env.Network.Aws.VpcID)
		}
		var subnetids types.Set
		if len(env.Network.SubnetIds) > 0 {
			var snDiags diag.Diagnostics
			subnetids, snDiags = types.SetValueFrom(ctx, types.StringType, env.Network.SubnetIds)
			diags.Append(snDiags...)
		} else {
			subnetids = types.SetNull(types.StringType)
		}
		model.SubnetIds = subnetids

	}
	if env.ProxyConfig != nil {
		model.ProxyConfigName = types.StringPointerValue(env.ProxyConfig.ProxyConfigName)
	}
	model.Region = types.StringPointerValue(env.Region)
	if env.SecurityAccess != nil {
		var dsgIDs types.Set
		if model.SecurityAccess != nil && !model.SecurityAccess.DefaultSecurityGroupIDs.IsUnknown() {
			dsgIDs = model.SecurityAccess.DefaultSecurityGroupIDs
		}
		var sgIDsknox types.Set
		if model.SecurityAccess != nil && !model.SecurityAccess.SecurityGroupIDsForKnox.IsUnknown() {
			sgIDsknox = model.SecurityAccess.SecurityGroupIDsForKnox
		}
		model.SecurityAccess = &SecurityAccess{
			Cidr:                       types.StringValue(env.SecurityAccess.Cidr),
			DefaultSecurityGroupID:     types.StringValue(env.SecurityAccess.DefaultSecurityGroupID),
			DefaultSecurityGroupIDs:    dsgIDs,
			SecurityGroupIDForKnox:     types.StringValue(env.SecurityAccess.SecurityGroupIDForKnox),
			SecurityGroupIDsForKnox:    sgIDsknox,
			GatewayNodeSecurityGroupID: model.SecurityAccess.GatewayNodeSecurityGroupID,
		}
	}
	model.Status = types.StringPointerValue(env.Status)
	model.StatusReason = types.StringValue(env.StatusReason)
	if env.Tags != nil {
		var tagDiags diag.Diagnostics
		tagMap, tagDiags := types.MapValueFrom(ctx, types.StringType, env.Tags.UserDefined)
		diags.Append(tagDiags...)
		model.Tags = tagMap
	}
	model.EnableTunnel = types.BoolValue(env.TunnelEnabled)
	model.TunnelType = types.StringValue(string(env.TunnelType))
	model.WorkloadAnalytics = types.BoolValue(env.WorkloadAnalytics)
}

func updateEnvironment(ctx context.Context, plan *awsEnvironmentResourceModel, state *awsEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	if plan.CredentialName.ValueString() != state.CredentialName.ValueString() {
		params := operations.NewChangeEnvironmentCredentialParamsWithContext(ctx)
		params.WithInput(&environmentsmodels.ChangeEnvironmentCredentialRequest{
			CredentialName:  plan.CredentialName.ValueStringPointer(),
			EnvironmentName: state.EnvironmentName.ValueStringPointer(),
		})
		_, err := client.Operations.ChangeEnvironmentCredential(params)
		if err != nil {
			utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "change AWS Environment credential")
			return resp
		}
	}

	if state.EncryptionKeyArn != plan.EncryptionKeyArn {
		if err := updateAwsDiskEncryptionParameters(ctx, client, *plan); err != nil {
			utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "update disk encryption parameters")
			return resp
		}
		state.EncryptionKeyArn = plan.EncryptionKeyArn
	}

	if plan.Authentication != nil && !reflect.DeepEqual(plan.Authentication, state.Authentication) {
		if err := updateSshKey(ctx, client, plan.Authentication, plan.EnvironmentName.ValueStringPointer()); err != nil {
			utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "update SSH key")
			return resp
		}
		state.Authentication = plan.Authentication
	}

	if !reflect.DeepEqual(utils.FromSetValueToStringList(plan.SubnetIds), utils.FromSetValueToStringList(state.SubnetIds)) ||
		!reflect.DeepEqual(plan.EndpointAccessGatewaySubnetIds, state.EndpointAccessGatewaySubnetIds) {
		if err := updateSubnet(ctx, client, *plan); err != nil {
			utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "update subnet")
			return resp
		}
		state.SubnetIds = plan.SubnetIds
		state.EndpointAccessGatewaySubnetIds = plan.EndpointAccessGatewaySubnetIds
	}

	if plan.SecurityAccess != nil && !reflect.DeepEqual(plan.SecurityAccess, state.SecurityAccess) {
		if err := updateSecurityAccess(ctx, client, *plan); err != nil {
			utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "update security access")
			return resp
		}
		state.SecurityAccess = plan.SecurityAccess
	}

	if !plan.Tags.IsNull() && !reflect.DeepEqual(plan.Tags, state.Tags) {
		if err := updateTags(ctx, client, *plan); err != nil {
			utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "update tags")
			return resp
		}
		state.Tags = plan.Tags
	}

	if plan.ProxyConfigName != state.ProxyConfigName {
		if err := updateProxyConfig(ctx, client, *plan); err != nil {
			utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "update proxy config")
			return resp
		}
		state.ProxyConfigName = plan.ProxyConfigName
	}
	return resp
}

func startEnvironment(ctx context.Context, state *awsEnvironmentResourceModel, resp *resource.UpdateResponse, client *environmentsclient.Environments) error {
	if !(state.PollingOptions != nil && state.PollingOptions.Async.ValueBool()) {
		stateSaver := func(env *environmentsmodels.Environment) {
			toAwsEnvironmentResource(ctx, utils.LogEnvironmentSilently(ctx, env, describeLogPrefix), state, state.PollingOptions, &resp.Diagnostics)
		}
		_, err := waitForStartEnvironmentWithDiagnosticHandle(ctx, client, state.ID.ValueString(), state.EnvironmentName.ValueString(), resp, state.PollingOptions, stateSaver)
		if err != nil {
			return err
		}
	}
	return nil
}

func stopAndWaitForEnvironment(ctx context.Context, environment string, pollingOptions *utils.PollingOptions, resp *resource.UpdateResponse, client *environmentsclient.Environments) error {
	params := operations.NewStopEnvironmentParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.StopEnvironmentRequest{
		EnvironmentName: &environment,
	})
	_, err := client.Operations.StopEnvironment(params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "stop Environment")
		return err
	}
	if err := waitForEnvironmentToBeStopped(environment, timeoutOneHour, callFailureThreshold, client, ctx, pollingOptions); err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "create Environment failed")
		return err
	}
	return nil
}
