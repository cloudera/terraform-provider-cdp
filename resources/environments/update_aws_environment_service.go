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
	return executeUpdateOperations(ctx, plan, state, client, resp,
		updateAwsEndpointAccessGatewayIfChanged,
		updateAwsCustomDockerRegistryIfChanged,
		updateAwsProxyConfigurationIfChanged,
		updateAwsComputeClusterIfChanged,
		updateAwsSecurityAccessIfChanged,
		updateAwsAuthenticationIfChanged,
		updateAwsEncryptionKeyIfChanged,
		updateAwsCredentialIfChanged,
		updateAwsCatalogIfChanged,
		updateAwsSubnetIfChanged,
		updateAwsTagsIfChanged,
	)
}

func updateAwsTagsIfChanged(ctx context.Context, plan *awsEnvironmentResourceModel, state *awsEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateTagsIfChanged(ctx, client, plan.Tags, &state.Tags, plan.EnvironmentName.ValueStringPointer(), plan.PollingOptions, resp)
}

func updateAwsSubnetIfChanged(ctx context.Context, plan *awsEnvironmentResourceModel, state *awsEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateSubnetIfChanged(ctx, client, plan.SubnetIds, &state.SubnetIds, plan.EnvironmentName.ValueStringPointer(), resp)
}

func updateAwsComputeClusterIfChanged(ctx context.Context, plan *awsEnvironmentResourceModel, state *awsEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
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
	return resp
}

func updateAwsCredentialIfChanged(ctx context.Context, plan *awsEnvironmentResourceModel, state *awsEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateCredential(ctx, client, plan.CredentialName, &state.CredentialName, plan.EnvironmentName.ValueStringPointer(), resp)
}

func updateAwsAuthenticationIfChanged(ctx context.Context, plan *awsEnvironmentResourceModel, state *awsEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	if plan.Authentication == nil {
		return resp
	}
	if reflect.DeepEqual(plan.Authentication, state.Authentication) {
		return resp
	}
	if err := updateSshKeyForAws(ctx, client, plan.Authentication, plan.EnvironmentName.ValueStringPointer()); err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "update SSH key")
		return resp
	}
	state.Authentication = plan.Authentication
	return resp
}

func updateAwsEncryptionKeyIfChanged(ctx context.Context, plan *awsEnvironmentResourceModel, state *awsEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	if reflect.DeepEqual(plan.EncryptionKeyArn, state.EncryptionKeyArn) {
		return resp
	}
	if err := updateDiskEncryption(ctx, client, plan.EnvironmentName.ValueStringPointer(), plan.EncryptionKeyArn); err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "update disk encryption param(s)")
		return resp
	}
	state.EncryptionKeyArn = plan.EncryptionKeyArn
	return resp
}

func updateAwsCatalogIfChanged(ctx context.Context, plan *awsEnvironmentResourceModel, state *awsEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateCatalogIfChanged(ctx, plan.FreeIpa, &state.FreeIpa, plan.EnvironmentName.ValueString(), client, resp)
}

func updateAwsProxyConfigurationIfChanged(ctx context.Context, plan *awsEnvironmentResourceModel, state *awsEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateProxyConfigurationIfChanged(ctx, client, &state.ProxyConfigName, &plan.ProxyConfigName, plan.EnvironmentName.ValueStringPointer(), resp)
}

func updateAwsEndpointAccessGatewayIfChanged(ctx context.Context, plan *awsEnvironmentResourceModel, state *awsEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateEndpointAccessGatewayIfChanged(ctx, client, plan.EndpointAccessGatewayScheme, plan.EndpointAccessGatewaySubnetIds, &state.EndpointAccessGatewayScheme, &state.EndpointAccessGatewaySubnetIds, plan.EnvironmentName.ValueString(), plan.PollingOptions, resp)
}

func updateAwsSecurityAccessIfChanged(ctx context.Context, plan *awsEnvironmentResourceModel, state *awsEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateSecurityAccessIfChanged(ctx, client,
		plan.SecurityAccess.DefaultSecurityGroupID,
		plan.SecurityAccess.SecurityGroupIDForKnox,
		&state.SecurityAccess.DefaultSecurityGroupID,
		&state.SecurityAccess.SecurityGroupIDForKnox,
		plan.EnvironmentName.ValueStringPointer(), resp)
}

func updateAwsCustomDockerRegistryIfChanged(ctx context.Context, plan *awsEnvironmentResourceModel, state *awsEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateCustomDockerRegistryIfChanged(ctx, client, state.CustomDockerRegistry, plan.CustomDockerRegistry, plan.EnvironmentName.ValueStringPointer(), resp)
}

func enableComputeClusterForAws(ctx context.Context, config *AwsComputeClusterConfiguration, environmentName string, envSubnets types.Set, envClient *environmentsclient.Environments) error {
	request := environmentsmodels.InitializeAWSComputeClusterRequest{
		ComputeClusterConfiguration: convertConfigToAwsComputeClusterConfigurationRequest(config, envSubnets),
		EnvironmentName:             &environmentName,
	}
	params := operations.NewInitializeAWSComputeClusterParams()
	params.WithInput(&request)
	tflog.Info(ctx, fmt.Sprintf("Initializing AWS compute cluster for environment '%s'", environmentName))
	_, err := envClient.Operations.InitializeAWSComputeClusterContext(ctx, params)
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

func updateSshKeyForAws(ctx context.Context, client *environmentsclient.Environments, authPlan *Authentication, env *string) error {
	if authPlan == nil {
		return nil
	}
	params := operations.NewUpdateSSHKeyParams()
	if !authPlan.PublicKey.IsNull() && !authPlan.PublicKey.IsUnknown() && authPlan.PublicKey.ValueString() != "" {
		params.WithInput(&environmentsmodels.UpdateSSHKeyRequest{
			Environment:  env,
			NewPublicKey: authPlan.PublicKey.ValueString(),
		})
	} else if !authPlan.PublicKeyID.IsNull() && !authPlan.PublicKeyID.IsUnknown() && authPlan.PublicKeyID.ValueString() != "" {
		params.WithInput(&environmentsmodels.UpdateSSHKeyRequest{
			Environment:         env,
			ExistingPublicKeyID: authPlan.PublicKeyID.ValueString(),
		})
	} else {
		return fmt.Errorf("either authentication.public_key or authentication.public_key_id must be set")
	}
	tflog.Info(ctx, "Updating SSH key in the environment")
	_, err := client.Operations.UpdateSSHKeyContext(ctx, params)
	return err
}

func updateDiskEncryption(ctx context.Context, client *environmentsclient.Environments, env *string, keyArn types.String) error {
	params := operations.NewUpdateAwsDiskEncryptionParametersParams().WithInput(&environmentsmodels.UpdateAwsDiskEncryptionParametersRequest{
		EncryptionKeyArn: keyArn.ValueStringPointer(),
		Environment:      env,
	})
	_, err := client.Operations.UpdateAwsDiskEncryptionParametersContext(ctx, params)
	return err
}
