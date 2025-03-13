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
	resp = enableComputeClustersForAzureIfNecessary(ctx, plan, state, client, resp)
	if resp.Diagnostics.HasError() {
		return resp
	}
	// further update operations shall come here
	return resp
}

func enableComputeClustersForAzureIfNecessary(ctx context.Context, plan *azureEnvironmentResourceModel, state *azureEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
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
	return nil
}

func enableComputeClusterForAzure(ctx context.Context, config *AzureComputeClusterConfiguration, environmentName string, envSubnets types.Set, envClient *environmentsclient.Environments) error {
	request := environmentsmodels.InitializeAzureComputeClusterRequest{
		ComputeClusterConfiguration: convertConfigToAzureComputeClusterConfigurationRequest(config, envSubnets),
		EnvironmentName:             &environmentName,
	}
	params := operations.NewInitializeAzureComputeClusterParamsWithContext(ctx)
	params.WithInput(&request)
	tflog.Info(ctx, fmt.Sprintf("Initializing Azure compute cluster for environment '%s'", environmentName))
	_, err := envClient.Operations.InitializeAzureComputeCluster(params)
	return err
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
