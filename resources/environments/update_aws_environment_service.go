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

	"github.com/hashicorp/terraform-plugin-log/tflog"

	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

func updateAwsDiskEncryptionParameters(ctx context.Context, client *environmentsclient.Environments, plan awsEnvironmentResourceModel) error {
	params := operations.NewUpdateAwsDiskEncryptionParametersParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.UpdateAwsDiskEncryptionParametersRequest{
		EncryptionKeyArn: plan.EncryptionKeyArn.ValueStringPointer(),
		Environment:      plan.EnvironmentName.ValueStringPointer(),
	})
	tflog.Info(ctx, "Updating disk encryption parameters in the environment")
	_, err := client.Operations.UpdateAwsDiskEncryptionParameters(params)
	return err
}

func updateSshKey(ctx context.Context, client *environmentsclient.Environments, authPlan *Authentication, env *string) error {
	params := operations.NewUpdateSSHKeyParamsWithContext(ctx)
	if !authPlan.PublicKey.IsNull() && authPlan.PublicKey.ValueString() != "" {
		params.WithInput(&environmentsmodels.UpdateSSHKeyRequest{
			Environment:         env,
			NewPublicKey:        authPlan.PublicKey.ValueString(),
			ExistingPublicKeyID: "",
		})
	} else {
		params.WithInput(&environmentsmodels.UpdateSSHKeyRequest{
			Environment:         env,
			ExistingPublicKeyID: authPlan.PublicKeyID.ValueString(),
			NewPublicKey:        "",
		})
	}
	tflog.Info(ctx, "Updating SSH key in the environment")
	_, err := client.Operations.UpdateSSHKey(params)
	return err
}

func updateSubnet(ctx context.Context, client *environmentsclient.Environments, plan awsEnvironmentResourceModel) error {
	params := operations.NewUpdateSubnetParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.UpdateSubnetRequest{
		Environment:                    plan.EnvironmentName.ValueStringPointer(),
		SubnetIds:                      utils.FromSetValueToStringList(plan.SubnetIds),
		EndpointAccessGatewaySubnetIds: utils.FromSetValueToStringList(plan.EndpointAccessGatewaySubnetIds),
	})
	tflog.Info(ctx, "Updating subnet(s) in the environment")
	_, err := client.Operations.UpdateSubnet(params)
	return err
}

func updateSecurityAccess(ctx context.Context, client *environmentsclient.Environments, plan awsEnvironmentResourceModel) error {
	params := operations.NewUpdateSecurityAccessParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.UpdateSecurityAccessRequest{
		DefaultSecurityGroupID:     plan.SecurityAccess.DefaultSecurityGroupID.ValueStringPointer(),
		Environment:                plan.EnvironmentName.ValueStringPointer(),
		GatewayNodeSecurityGroupID: plan.SecurityAccess.GatewayNodeSecurityGroupID.ValueStringPointer(),
	})
	tflog.Info(ctx, "Updating security access in the environment")
	_, err := client.Operations.UpdateSecurityAccess(params)
	return err
}

func updateTags(ctx context.Context, _ *environmentsclient.Environments, _ awsEnvironmentResourceModel) error {
	tflog.Error(ctx, "UpdateTags is not implemented yet, it has to be present in BETA SDK beforehand")
	return nil
}

func updateProxyConfig(ctx context.Context, client *environmentsclient.Environments, plan awsEnvironmentResourceModel) error {
	params := operations.NewUpdateProxyConfigParamsWithContext(ctx)
	if plan.ProxyConfigName.IsNull() || plan.ProxyConfigName.ValueString() == "" {
		params.WithInput(&environmentsmodels.UpdateProxyConfigRequest{
			Environment: plan.EnvironmentName.ValueStringPointer(),
			RemoveProxy: true,
		})
		tflog.Info(ctx, "Removing proxy configuration from the environment")
	} else {
		params.WithInput(&environmentsmodels.UpdateProxyConfigRequest{
			Environment:     plan.EnvironmentName.ValueStringPointer(),
			ProxyConfigName: plan.ProxyConfigName.ValueString(),
		})
		tflog.Info(ctx, "Updating proxy configuration in the environment")
	}
	_, err := client.Operations.UpdateProxyConfig(params)
	return err
}
