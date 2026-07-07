// Copyright 2025 Cloudera. All Rights Reserved.
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
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

func updateAzureEnvironment(ctx context.Context, plan *azureEnvironmentResourceModel, state *azureEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return executeUpdateOperations(ctx, plan, state, client, resp,
		updateAzureSubnetIfChanged,
		updateAzureEndpointAccessGatewayIfChanged,
		updateAzureCustomDockerRegistryIfChanged,
		updateAzureProxyConfigurationIfChanged,
		updateAzureAvailabilityZonesIfChanged,
		updateAzureComputeClusterIfChanged,
		updateAzureSecurityAccessIfChanged,
		updateAzureDataServicesIfChanged,
		updateAzureCredentialIfChanged,
		updateAzureEncryptionIfChanged,
		updateAzureCatalogIfChanged,
		updateAzureSshKeyIfChanged,
	)
}

func updateAzureSubnetIfChanged(ctx context.Context, plan *azureEnvironmentResourceModel, state *azureEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	if plan.ExistingNetworkParams.IsNull() || plan.ExistingNetworkParams.IsUnknown() {
		return resp
	}
	var planNetwork, stateNetwork existingAzureNetwork
	resp.Diagnostics.Append(plan.ExistingNetworkParams.As(ctx, &planNetwork, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
	resp.Diagnostics.Append(state.ExistingNetworkParams.As(ctx, &stateNetwork, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
	if resp.Diagnostics.HasError() {
		return resp
	}
	if reflect.DeepEqual(planNetwork.SubnetIds, stateNetwork.SubnetIds) {
		return resp
	}
	tflog.Info(ctx, fmt.Sprintf("Updating subnets for environment '%s'", plan.EnvironmentName.ValueString()))
	params := operations.NewUpdateSubnetParams()
	params.WithInput(&environmentsmodels.UpdateSubnetRequest{
		Environment: plan.EnvironmentName.ValueStringPointer(),
		SubnetIds:   utils.FromSetValueToStringList(planNetwork.SubnetIds),
	})
	if _, err := client.Operations.UpdateSubnetContext(ctx, params); err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "update subnet")
	} else {
		state.ExistingNetworkParams = plan.ExistingNetworkParams
	}
	return resp
}

func updateAzureCredentialIfChanged(ctx context.Context, plan *azureEnvironmentResourceModel, state *azureEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateCredential(ctx, client, plan.CredentialName, &state.CredentialName, plan.EnvironmentName.ValueStringPointer(), resp)
}

func updateAzureSshKeyIfChanged(ctx context.Context, plan *azureEnvironmentResourceModel, state *azureEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateSshKeyIfChanged(ctx, client, plan.PublicKey, &state.PublicKey, plan.EnvironmentName.ValueStringPointer(), resp)
}

func updateAzureCatalogIfChanged(ctx context.Context, plan *azureEnvironmentResourceModel, state *azureEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateCatalogIfChanged(ctx, plan.FreeIpa, &state.FreeIpa, plan.EnvironmentName.ValueString(), client, resp)
}

func updateAzureEndpointAccessGatewayIfChanged(ctx context.Context, plan *azureEnvironmentResourceModel, state *azureEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateEndpointAccessGatewayIfChanged(ctx, client, plan.EndpointAccessGatewayScheme, plan.EndpointAccessGatewaySubnetIds, &state.EndpointAccessGatewayScheme, &state.EndpointAccessGatewaySubnetIds, plan.EnvironmentName.ValueString(), plan.PollingOptions, resp)
}

func updateAzureAvailabilityZonesIfChanged(ctx context.Context, plan *azureEnvironmentResourceModel, state *azureEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateAvailabilityZones(ctx, client, plan.AvailabilityZones, &state.AvailabilityZones, plan.EnvironmentName.ValueStringPointer(), resp)
}

func updateAvailabilityZones(ctx context.Context, client *environmentsclient.Environments, plan types.Set, state *types.Set, env *string, resp *resource.UpdateResponse) *resource.UpdateResponse {
	if plan.IsNull() || plan.IsUnknown() {
		return resp
	}
	if len(plan.Elements()) == 0 {
		resp.Diagnostics.AddError("Invalid availability zone setup", "availability_zones must be a non-empty, known value.")
		return resp
	}
	if !plan.Equal(*state) {
		tflog.Info(ctx, fmt.Sprintf("Updating Azure availability zones for environment '%s'", *env))
		request := environmentsmodels.UpdateAzureAvailabilityZonesRequest{
			AvailabilityZones: utils.FromSetValueToStringList(plan),
			Environment:       env,
		}
		params := operations.NewUpdateAzureAvailabilityZonesParams()
		params.WithInput(&request)
		_, err := client.Operations.UpdateAzureAvailabilityZonesContext(ctx, params)
		if err != nil {
			utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "update azure availability zone")
		} else {
			*state = plan
		}
	}
	return resp
}

func updateAzureProxyConfigurationIfChanged(ctx context.Context, plan *azureEnvironmentResourceModel, state *azureEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateProxyConfigurationIfChanged(ctx, client, &state.ProxyConfigName, &plan.ProxyConfigName, plan.EnvironmentName.ValueStringPointer(), resp)
}

func updateAzureSecurityAccessIfChanged(ctx context.Context, plan *azureEnvironmentResourceModel, state *azureEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateSecurityAccessIfChanged(ctx, client,
		plan.SecurityAccess.DefaultSecurityGroupID,
		plan.SecurityAccess.SecurityGroupIDForKnox,
		&state.SecurityAccess.DefaultSecurityGroupID,
		&state.SecurityAccess.SecurityGroupIDForKnox,
		plan.EnvironmentName.ValueStringPointer(), resp)
}

func updateAzureCustomDockerRegistryIfChanged(ctx context.Context, plan *azureEnvironmentResourceModel, state *azureEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateCustomDockerRegistryIfChanged(ctx, client, state.CustomDockerRegistry, plan.CustomDockerRegistry, plan.EnvironmentName.ValueStringPointer(), resp)
}

func updateAzureEncryptionIfChanged(ctx context.Context, plan *azureEnvironmentResourceModel, state *azureEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	if !azureEncryptionFieldsChanged(plan, state) {
		return resp
	}
	if plan.EncryptionKeyURL.IsNull() || plan.EncryptionKeyURL.IsUnknown() {
		resp.Diagnostics.AddError("update Azure encryption resources", "encryption_key_url must be set to a known value when updating encryption parameters")
		return resp
	}
	if err := updateAzureEncryptionResources(ctx, client, plan.EnvironmentName.ValueStringPointer(), plan); err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "update Azure encryption resources")
		return resp
	}
	state.EncryptionKeyURL = plan.EncryptionKeyURL
	state.EncryptionKeyResourceGroupName = plan.EncryptionKeyResourceGroupName
	state.EncryptionUserManagedIdentity = plan.EncryptionUserManagedIdentity
	return resp
}

func updateAzureComputeClusterIfChanged(ctx context.Context, plan *azureEnvironmentResourceModel, state *azureEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	if state.ComputeCluster == nil && plan.ComputeCluster != nil && plan.ComputeCluster.Enabled.ValueBool() {
		tflog.Info(ctx, fmt.Sprintf("Request for compute cluster enablement for environment '%s' is detected.", plan.EnvironmentName.ValueString()))
		var existingNetwork existingAzureNetwork
		diags := state.ExistingNetworkParams.As(ctx, &existingNetwork, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
		if err := enableComputeClusterForAzure(ctx, plan.ComputeCluster.Configuration, plan.EnvironmentName.ValueString(), existingNetwork.SubnetIds, client); err != nil {
			tflog.Warn(ctx, "Failed to enable compute cluster", map[string]interface{}{
				"error": err.Error(),
			})
			utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "enable compute cluster")
			return resp
		}
		stateSaver := func(env *environmentsmodels.Environment) {
			toAzureEnvironmentResource(ctx, utils.LogEnvironmentSilently(ctx, env, describeLogPrefix), state, state.PollingOptions, &resp.Diagnostics)
			resp.Diagnostics = resp.State.Set(ctx, state)
			resp.Diagnostics.Append(diags...)
		}
		if err := waitForEnvironmentToBeAvailable(state.ID.ValueString(), timeoutOneHour, callFailureThreshold, client, ctx, state.PollingOptions, stateSaver); err != nil {
			utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "wait for environment to be available")
		}
	}
	return resp
}

func enableComputeClusterForAzure(ctx context.Context, config *AzureComputeClusterConfiguration, environmentName string, envSubnets types.Set, envClient *environmentsclient.Environments) error {
	request := environmentsmodels.InitializeAzureComputeClusterRequest{
		ComputeClusterConfiguration: convertConfigToAzureComputeClusterConfigurationRequest(config, envSubnets),
		EnvironmentName:             &environmentName,
	}
	params := operations.NewInitializeAzureComputeClusterParams()
	params.WithInput(&request)
	tflog.Info(ctx, fmt.Sprintf("Initializing Azure compute cluster for environment '%s'", environmentName))
	_, err := envClient.Operations.InitializeAzureComputeClusterContext(ctx, params)
	return err
}

func azureEncryptionFieldsChanged(plan, state *azureEnvironmentResourceModel) bool {
	return !reflect.DeepEqual(plan.EncryptionKeyURL, state.EncryptionKeyURL) ||
		!reflect.DeepEqual(plan.EncryptionKeyResourceGroupName, state.EncryptionKeyResourceGroupName) ||
		!reflect.DeepEqual(plan.EncryptionUserManagedIdentity, state.EncryptionUserManagedIdentity)
}

func updateAzureEncryptionResources(ctx context.Context, client *environmentsclient.Environments, env *string, plan *azureEnvironmentResourceModel) error {
	params := operations.NewUpdateAzureEncryptionResourcesParams().WithInput(&environmentsmodels.UpdateAzureEncryptionResourcesRequest{
		EncryptionKeyURL:               plan.EncryptionKeyURL.ValueStringPointer(),
		Environment:                    env,
		EncryptionKeyResourceGroupName: plan.EncryptionKeyResourceGroupName.ValueString(),
		EncryptionUserManagedIdentity:  plan.EncryptionUserManagedIdentity.ValueString(),
	})
	_, err := client.Operations.UpdateAzureEncryptionResourcesContext(ctx, params)
	return err
}

func updateAzureDataServicesIfChanged(ctx context.Context, plan *azureEnvironmentResourceModel, state *azureEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	if plan.DataServices == nil {
		// The backing API requires shared_managed_identity, so the provider cannot apply a removal of the entire block.
		if state.DataServices != nil {
			resp.Diagnostics.AddError("update data service resources", "data_services cannot be removed once set; please keep it configured or recreate the environment")
		}
		return resp
	}

	if state.DataServices != nil && reflect.DeepEqual(*plan.DataServices, *state.DataServices) {
		return resp
	}

	if plan.DataServices.SharedManagedIdentity.IsNull() || plan.DataServices.SharedManagedIdentity.IsUnknown() {
		resp.Diagnostics.AddError("update data service resources", "shared_managed_identity must be set to a known value when updating data service resources")
		return resp
	}

	// Avoid silently drifting state when the optional zone ID is unset; the API model omits empty string values.
	if plan.DataServices.AksPrivateDnsZoneId.IsNull() && state.DataServices != nil && !state.DataServices.AksPrivateDnsZoneId.IsNull() {
		resp.Diagnostics.AddError("update data service resources", "aks_private_dns_zone_id cannot be unset via update; please keep the existing value or recreate the environment")
		return resp
	}

	params := operations.NewUpdateDataServiceResourcesParams().WithInput(&environmentsmodels.UpdateDataServiceResourcesRequest{
		Environment: plan.EnvironmentName.ValueStringPointer(),
		DataServices: &environmentsmodels.DataServicesRequest{
			Azure: &environmentsmodels.AzureDataServicesParametersRequest{
				SharedManagedIdentity: plan.DataServices.SharedManagedIdentity.ValueStringPointer(),
				AksPrivateDNSZoneID:   plan.DataServices.AksPrivateDnsZoneId.ValueString(),
			},
		},
	})
	tflog.Info(ctx, fmt.Sprintf("Updating data service resources for environment '%s'", plan.EnvironmentName.ValueString()))
	if _, err := client.Operations.UpdateDataServiceResourcesContext(ctx, params); err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "update data service resources")
		return resp
	}

	if state.DataServices == nil {
		state.DataServices = &DataServices{}
	}
	*state.DataServices = *plan.DataServices
	return resp
}

func convertConfigToAzureComputeClusterConfigurationRequest(config *AzureComputeClusterConfiguration, fallbackSubnetIds types.Set) *environmentsmodels.AzureComputeClusterConfigurationRequest {
	if config == nil {
		return nil
	}
	var subnetIds types.Set
	if !config.WorkerNodeSubnets.IsNull() && len(config.WorkerNodeSubnets.Elements()) > 0 {
		subnetIds = config.WorkerNodeSubnets
	} else {
		subnetIds = fallbackSubnetIds
	}
	return &environmentsmodels.AzureComputeClusterConfigurationRequest{
		KubeAPIAuthorizedIPRanges: utils.FromSetValueToStringList(config.KubeApiAuthorizedIpRanges),
		PrivateCluster:            config.PrivateCluster.ValueBool(),
		WorkerNodeSubnets:         utils.FromSetValueToStringList(subnetIds),
		OutboundType:              config.OutboundType.ValueString(),
	}
}
