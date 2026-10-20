// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package azure

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/resources/environments/envcommon"
	"github.com/cloudera/terraform-provider-cdp/resources/environments/freeipa"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ resource.ResourceWithConfigure        = &azureEnvironmentResource{}
	_ resource.ResourceWithImportState      = &azureEnvironmentResource{}
	_ resource.ResourceWithConfigValidators = &azureEnvironmentResource{}
)

type azureEnvironmentResource struct {
	client *cdp.Client
}

func (r *azureEnvironmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func NewAzureEnvironmentResource() resource.Resource {
	return &azureEnvironmentResource{}
}

func (r *azureEnvironmentResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_azure_environment"
}

func (r *azureEnvironmentResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	s := AzureEnvironmentSchema
	s.Attributes = make(map[string]schema.Attribute, len(AzureEnvironmentSchema.Attributes)+1)
	for k, v := range AzureEnvironmentSchema.Attributes {
		s.Attributes[k] = v
	}
	s.Attributes["timeouts"] = timeouts.Attributes(ctx, timeouts.Opts{Create: true, Delete: true, Update: true})
	resp.Schema = s
}

func (r *azureEnvironmentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *azureEnvironmentResource) ConfigValidators(context.Context) []resource.ConfigValidator {
	return utils.PollingOptionsConflictValidators()
}

func (r *azureEnvironmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AzureEnvironmentResourceModel
	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Got Error while trying to set plan")
		return
	}

	client := r.client.Environments

	params := operations.NewCreateAzureEnvironmentParams()
	params.WithInput(ToAzureEnvironmentRequest(ctx, &data))

	responseOk, err := client.Operations.CreateAzureEnvironmentContext(ctx, params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "create Azure Environment")
		return
	}

	envResp := responseOk.Payload.Environment
	ToAzureEnvironmentResource(ctx, envResp, &data, &resp.Diagnostics)

	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	descEnvResp, err := envcommon.DescribeEnvironmentWithDiagnosticHandle(data.EnvironmentName.ValueString(), data.ID.ValueString(), ctx, r.client, &resp.Diagnostics, &resp.State)
	if err != nil {
		return
	}
	if !utils.ResolveAsync(data.PollingOptions, data.Async) {
		createTimeout, diagsT := utils.ResolveTimeout(ctx, data.PollingOptions, data.Timeouts.Create, 60*time.Minute)
		resp.Diagnostics.Append(diagsT...)
		if resp.Diagnostics.HasError() {
			return
		}
		callFailureThreshold := utils.ResolveCallFailureThreshold(data.PollingOptions, data.CallFailureThreshold, envcommon.CallFailureThreshold)
		stateSaver := func(env *environmentsmodels.Environment) {
			ToAzureEnvironmentResource(ctx, utils.LogEnvironmentSilently(ctx, env, envcommon.DescribeLogPrefix), &data, &resp.Diagnostics)
			diags = resp.State.Set(ctx, data)
			resp.Diagnostics.Append(diags...)
		}
		descEnvResp, err = envcommon.WaitForCreateEnvironmentWithDiagnosticHandle(ctx, r.client, data.ID.ValueString(), data.EnvironmentName.ValueString(), resp, createTimeout, callFailureThreshold, stateSaver)
		if err != nil {
			return
		}
	}

	ToAzureEnvironmentResource(ctx, descEnvResp, &data, &resp.Diagnostics)

	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *azureEnvironmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AzureEnvironmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	envName := state.EnvironmentName.ValueString()
	if len(envName) == 0 {
		envName = state.ID.ValueString()
	}
	descEnvResp, err := envcommon.DescribeEnvironmentWithDiagnosticHandle(envName, state.ID.ValueString(), ctx, r.client, &resp.Diagnostics, &resp.State)
	if err != nil {
		return
	}
	ToAzureEnvironmentResource(ctx, descEnvResp, &state, &resp.Diagnostics)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func ToAzureEnvironmentResource(ctx context.Context, env *environmentsmodels.Environment, model *AzureEnvironmentResourceModel, diags *diag.Diagnostics) {
	utils.LogEnvironmentSilently(ctx, env, "Converting environment: ")
	model.ID = types.StringPointerValue(env.Crn)
	model.Crn = types.StringPointerValue(env.Crn)
	model.CredentialName = types.StringPointerValue(env.CredentialName)
	model.Description = types.StringValue(env.Description)
	model.EnableTunnel = types.BoolValue(env.TunnelEnabled)
	model.EnvironmentName = types.StringPointerValue(env.EnvironmentName)
	if env.LogStorage != nil {
		if env.LogStorage.AzureDetails != nil {
			model.LogStorage = &azureLogStorage{
				ManagedIdentity:     types.StringValue(env.LogStorage.AzureDetails.ManagedIdentity),
				StorageLocationBase: types.StringValue(env.LogStorage.AzureDetails.StorageLocationBase),
			}
			if env.BackupStorage != nil {
				if env.BackupStorage.AzureDetails != nil {
					model.LogStorage.BackupStorageLocationBase = types.StringValue(env.BackupStorage.AzureDetails.StorageLocationBase)
				}

			}
		}
	}
	diags.Append(*freeipa.ResponseToModel(env.Freeipa, &model.FreeIpa, ctx)...)
	if env.Network != nil {
		var npDiags diag.Diagnostics
		if env.Network.EndpointAccessGatewaySubnetIds != nil {
			var gatewaySubnetIds types.Set
			if len(env.Network.EndpointAccessGatewaySubnetIds) > 0 {
				var eagSnDiags diag.Diagnostics
				gatewaySubnetIds, eagSnDiags = types.SetValueFrom(ctx, types.StringType, env.Network.EndpointAccessGatewaySubnetIds)
				diags.Append(eagSnDiags...)
			} else {
				gatewaySubnetIds = types.SetNull(types.StringType)
			}
			model.EndpointAccessGatewaySubnetIds = gatewaySubnetIds
		}
		diags.Append(npDiags...)
		if env.Network.Azure != nil {
			subnetIds, snDiags := types.SetValueFrom(ctx, types.StringType, env.Network.SubnetIds)
			diags.Append(snDiags...)
			flexSubnetIds, fsDiags := types.SetValueFrom(ctx, types.StringType, env.Network.Azure.FlexibleServerSubnetIds)
			diags.Append(fsDiags...)
			var enpDiags diag.Diagnostics
			model.ExistingNetworkParams, enpDiags = types.ObjectValueFrom(ctx, map[string]attr.Type{
				"aks_private_dns_zone_id":      types.StringType,
				"database_private_dns_zone_id": types.StringType,
				"network_id":                   types.StringType,
				"resource_group_name":          types.StringType,
				"subnet_ids": types.SetType{
					ElemType: types.StringType,
				},
				"flexible_server_subnet_ids": types.SetType{
					ElemType: types.StringType,
				},
			}, &existingAzureNetwork{
				AksPrivateDNSZoneID:      types.StringValue(env.Network.Azure.AksPrivateDNSZoneID),
				DatabasePrivateDNSZoneID: types.StringValue(env.Network.Azure.DatabasePrivateDNSZoneID),
				NetworkID:                types.StringPointerValue(env.Network.Azure.NetworkID),
				ResourceGroupName:        types.StringPointerValue(env.Network.Azure.ResourceGroupName),
				SubnetIds:                subnetIds,
				FlexibleServerSubnetIds:  flexSubnetIds,
			})
			diags.Append(enpDiags...)
			model.UsePublicIP = types.BoolPointerValue(env.Network.Azure.UsePublicIP)
		}
	}
	if env.ProxyConfig != nil {
		model.ProxyConfigName = types.StringPointerValue(env.ProxyConfig.ProxyConfigName)
	}
	if env.SecurityAccess != nil {
		var dsgIDs types.Set
		if model.SecurityAccess == nil || model.SecurityAccess.DefaultSecurityGroupIDs.IsUnknown() {
			dsgIDs = types.SetNull(types.StringType)
		} else {
			dsgIDs = model.SecurityAccess.DefaultSecurityGroupIDs
		}
		var sgIDsknox types.Set
		if model.SecurityAccess == nil || model.SecurityAccess.SecurityGroupIDsForKnox.IsUnknown() {
			sgIDsknox = types.SetNull(types.StringType)
		} else {
			sgIDsknox = model.SecurityAccess.DefaultSecurityGroupIDs
		}
		model.SecurityAccess = &SecurityAccess{
			Cidr:                    types.StringValue(env.SecurityAccess.Cidr),
			DefaultSecurityGroupID:  types.StringValue(env.SecurityAccess.DefaultSecurityGroupID),
			DefaultSecurityGroupIDs: dsgIDs,
			SecurityGroupIDForKnox:  types.StringValue(env.SecurityAccess.SecurityGroupIDForKnox),
			SecurityGroupIDsForKnox: sgIDsknox,
		}
	}
	model.Status = types.StringPointerValue(env.Status)
	model.StatusReason = types.StringValue(env.StatusReason)
	if env.Tags != nil {
		tagMap, tagDiags := types.MapValueFrom(ctx, types.StringType, env.Tags.UserDefined)
		diags.Append(tagDiags...)
		model.Tags = tagMap
	}
	model.WorkloadAnalytics = types.BoolValue(env.WorkloadAnalytics)
	utils.LogEnvironmentSilently(ctx, env, "Environment conversion finished: ")
}

func (r *azureEnvironmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	envcommon.PerformEnvironmentUpdate(ctx, req, resp, r.client.Environments, updateAzureEnvironment)
}

func (r *azureEnvironmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state AzureEnvironmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	cascading := state.cascadeDelete()
	forced := state.forceDelete()

	deleteTimeout, diags := utils.ResolveTimeout(ctx, state.PollingOptions, state.Timeouts.Delete, 60*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	callFailureThreshold := utils.ResolveCallFailureThreshold(state.PollingOptions, state.CallFailureThreshold, envcommon.CallFailureThreshold)
	if err := envcommon.DeleteEnvironmentWithDiagnosticHandle(state.EnvironmentName.ValueString(), cascading, forced, ctx, r.client, resp, deleteTimeout, callFailureThreshold); err != nil {
		return
	}
}
