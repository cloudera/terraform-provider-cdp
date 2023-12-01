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
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

func toGcpEnvironmentRequest(ctx context.Context, model *gcpEnvironmentResourceModel) *environmentsmodels.CreateGCPEnvironmentRequest {
	req := &environmentsmodels.CreateGCPEnvironmentRequest{
		CredentialName:              model.CredentialName.ValueStringPointer(),
		Description:                 model.Description.ValueString(),
		EnableTunnel:                model.EnableTunnel.ValueBool(),
		EncryptionKey:               model.EncryptionKey.ValueString(),
		EndpointAccessGatewayScheme: model.EndpointAccessGatewayScheme.ValueString(),
		EnvironmentName:             model.EnvironmentName.ValueStringPointer(),
		ExistingNetworkParams: &environmentsmodels.ExistingGCPNetworkRequest{
			NetworkName:     model.ExistingNetworkParams.NetworkName.ValueStringPointer(),
			SharedProjectID: model.ExistingNetworkParams.SharedProjectId.ValueStringPointer(),
			SubnetNames:     utils.FromListValueToStringList(model.ExistingNetworkParams.SubnetNames),
		},
		ProxyConfigName:      model.ProxyConfigName.ValueString(),
		PublicKey:            model.PublicKey.ValueStringPointer(),
		Region:               model.Region.ValueStringPointer(),
		ReportDeploymentLogs: model.ReportDeploymentLogs.ValueBool(),
		UsePublicIP:          model.UsePublicIp.ValueBoolPointer(),
		WorkloadAnalytics:    model.WorkloadAnalytics.ValueBool(),
	}
	if model.FreeIpa != nil {
		tflog.Debug(ctx, fmt.Sprintf("model.FreeIpa: %+v\n", model.FreeIpa))
		req.FreeIpa = &environmentsmodels.GCPFreeIpaCreationRequest{
			InstanceCountByGroup: int32(model.FreeIpa.InstanceCountByGroup.ValueInt64()),
			InstanceType:         model.FreeIpa.InstanceType.ValueString(),
			Recipes:              utils.FromSetValueToStringList(model.FreeIpa.Recipes),
		}
		tflog.Debug(ctx, fmt.Sprintf("req.FreeIpa: %+v\n", req.FreeIpa))
	}

	if model.LogStorage != nil {
		req.LogStorage = &environmentsmodels.GcpLogStorageRequest{
			StorageLocationBase:       model.LogStorage.StorageLocationBase.ValueStringPointer(),
			BackupStorageLocationBase: model.LogStorage.BackupStorageLocationBase.ValueString(),
			ServiceAccountEmail:       model.LogStorage.ServiceAccountEmail.ValueStringPointer(),
		}
	}
	if model.SecurityAccess != nil {
		req.SecurityAccess = &environmentsmodels.GcpSecurityAccessRequest{
			DefaultSecurityGroupID: model.SecurityAccess.DefaultSecurityGroupId.ValueString(),
			SecurityGroupIDForKnox: model.SecurityAccess.SecurityGroupIdForKnox.ValueString(),
		}
	}
	req.Tags = ConvertGcpTags(ctx, model.Tags)
	return req
}

func toGcpEnvironmentResource(ctx context.Context, env *environmentsmodels.Environment, model *gcpEnvironmentResourceModel, pollingOptions *utils.PollingOptions, diags *diag.Diagnostics) {
	model.ID = types.StringPointerValue(env.Crn)
	model.CredentialName = types.StringPointerValue(env.CredentialName)
	model.Crn = types.StringPointerValue(env.Crn)
	model.Description = types.StringValue(env.Description)
	model.EnvironmentName = types.StringPointerValue(env.EnvironmentName)
	model.Status = types.StringValue(*env.Status)
	model.StatusReason = types.StringValue(env.StatusReason)
	model.PollingOptions = pollingOptions
	tflog.Info(ctx, "about to convert log storage.")
	if env.LogStorage != nil {
		if env.LogStorage.GcpDetails != nil {
			backupStorageLocationBase := ""
			if model.LogStorage != nil && !model.LogStorage.BackupStorageLocationBase.IsNull() &&
				!model.LogStorage.BackupStorageLocationBase.IsUnknown() {
				backupStorageLocationBase = model.LogStorage.BackupStorageLocationBase.ValueString()
			}
			model.LogStorage = &GcpLogStorage{
				StorageLocationBase: types.StringValue(env.LogStorage.GcpDetails.StorageLocationBase),
				BackupStorageLocationBase: func(base string) types.String {
					if len(base) > 0 {
						return types.StringValue(base)
					}
					return types.StringNull()
				}(backupStorageLocationBase),
			}
		}
	}
	tflog.Info(ctx, "about to convert network.")
	if env.Network != nil {
		if model.ExistingNetworkParams != nil {
			model.ExistingNetworkParams = &ExistingNetworkParams{
				NetworkName:     types.StringValue(*env.Network.Gcp.NetworkName),
				SharedProjectId: types.StringValue(env.Network.Gcp.SharedProjectID),
				SubnetNames:     model.ExistingNetworkParams.SubnetNames,
			}
		} else {
			model.ExistingNetworkParams = &ExistingNetworkParams{
				NetworkName:     types.StringValue(*env.Network.Gcp.NetworkName),
				SharedProjectId: types.StringValue(env.Network.Gcp.SharedProjectID),
			}
		}
	}
	tflog.Info(ctx, "about to convert proxy config.")
	if env.ProxyConfig != nil {
		model.ProxyConfigName = types.StringPointerValue(env.ProxyConfig.ProxyConfigName)
	}
	model.Region = types.StringPointerValue(env.Region)
	model.ReportDeploymentLogs = types.BoolValue(env.ReportDeploymentLogs)
	tflog.Info(ctx, "about to convert security access.")
	if env.SecurityAccess != nil && len(env.SecurityAccess.SecurityGroupIDForKnox) > 0 && len(env.SecurityAccess.DefaultSecurityGroupID) > 0 {
		model.SecurityAccess = &GcpSecurityAccess{
			SecurityGroupIdForKnox: types.StringValue(env.SecurityAccess.SecurityGroupIDForKnox),
			DefaultSecurityGroupId: types.StringValue(env.SecurityAccess.DefaultSecurityGroupID),
		}
	}
	model.Status = types.StringPointerValue(env.Status)
	model.StatusReason = types.StringValue(env.StatusReason)
	tflog.Info(ctx, "about to convert tags.")
	if env.Tags != nil {
		merged := env.Tags.Defaults
		for k, v := range env.Tags.UserDefined {
			merged[k] = v
		}
		var tagDiags diag.Diagnostics
		tagMap, tagDiags := types.MapValueFrom(ctx, types.StringType, merged)
		diags.Append(tagDiags...)
		model.Tags = tagMap
	}
	model.EnableTunnel = types.BoolValue(env.TunnelEnabled)
	model.WorkloadAnalytics = types.BoolValue(env.WorkloadAnalytics)
}
