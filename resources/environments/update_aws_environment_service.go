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

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

func updateAwsEnvironment(ctx context.Context, plan *awsEnvironmentResourceModel, state *awsEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	resp = enableComputeClustersForAwsIfNecessary(ctx, plan, state, client, resp)
	if resp.Diagnostics.HasError() {
		return resp
	}
	// further update operations shall come here
	return resp
}

func enableComputeClustersForAwsIfNecessary(ctx context.Context, plan *awsEnvironmentResourceModel, state *awsEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	if state.ComputeCluster == nil && plan.ComputeCluster != nil && plan.ComputeCluster.Enabled.ValueBool() {
		tflog.Info(ctx, fmt.Sprintf("Request for compute cluster enablement for environment '%s' is detected.", plan.EnvironmentName.ValueString()))
		if err := enableComputeClusterForAws(ctx, plan.ComputeCluster.Configuration, plan.EnvironmentName.ValueString(), state.SubnetIds, client); err != nil {
			tflog.Warn(ctx, "Failed to enable compute cluster", map[string]interface{}{
				"error": err.Error(),
			})
			utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "enable compute cluster")
			return resp
		}
		if err := waitForUpdateToFinish(ctx, plan.EnvironmentName.ValueString(), client, resp, plan.PollingOptions, nil, state); err != nil {
			utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "wait for environment to be available")
			return resp
		}
	}
	return nil
}

func enableComputeClusterForAws(ctx context.Context, config *AwsComputeClusterConfiguration, environmentName string, envSubnets types.Set, envClient *environmentsclient.Environments) error {
	request := environmentsmodels.InitializeAWSComputeClusterRequest{
		ComputeClusterConfiguration: convertConfigToAwsComputeClusterConfigurationRequest(config, envSubnets),
		EnvironmentName:             &environmentName,
	}
	params := operations.NewInitializeAWSComputeClusterParamsWithContext(ctx)
	params.WithInput(&request)
	tflog.Info(ctx, fmt.Sprintf("Initializing AWS compute cluster for environment '%s'", environmentName))
	_, err := envClient.Operations.InitializeAWSComputeCluster(params)
	return err
}

func convertConfigToAwsComputeClusterConfigurationRequest(config *AwsComputeClusterConfiguration, fallbackSubnetIds types.Set) *environmentsmodels.AWSComputeClusterConfigurationRequest {
	if config == nil {
		return nil
	}
	var subnetIds types.Set
	if !config.WorkerNodeSubnets.IsNull() && len(config.WorkerNodeSubnets.Elements()) > 0 {
		subnetIds = config.WorkerNodeSubnets
	} else {
		subnetIds = fallbackSubnetIds
	}
	return &environmentsmodels.AWSComputeClusterConfigurationRequest{
		KubeAPIAuthorizedIPRanges: utils.FromSetValueToStringList(config.KubeApiAuthorizedIpRanges),
		PrivateCluster:            config.PrivateCluster.ValueBool(),
		WorkerNodeSubnets:         utils.FromSetValueToStringList(subnetIds),
	}
}

func waitForUpdateToFinish(ctx context.Context, id string, client *environmentsclient.Environments, resp *resource.UpdateResponse, options *utils.PollingOptions, diags diag.Diagnostics, data *awsEnvironmentResourceModel) error {
	stateSaver := func(env *environmentsmodels.Environment) {
		toAwsEnvironmentResource(ctx, utils.LogEnvironmentSilently(ctx, env, describeLogPrefix), data, data.PollingOptions, &resp.Diagnostics)
		diags = resp.State.Set(ctx, data)
		resp.Diagnostics.Append(diags...)
	}
	if err := waitForEnvironmentToBeAvailable(id, timeoutOneHour, callFailureThreshold, client, ctx, options, stateSaver); err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "create Environment failed")
		return err
	}
	return nil
}
