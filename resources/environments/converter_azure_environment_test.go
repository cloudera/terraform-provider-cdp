// Copyright 2026 Cloudera. All Rights Reserved.
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
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"

	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

const (
	testCrn             = "crn:cdp:environments:us-west-1:tenant:environment:test-crn-123"
	testCredential      = "my-azure-credential"
	testDescription     = "Test environment description"
	testStatus          = "AVAILABLE"
	testStatusReason    = "Environment is running"
	testManagedIdentity = "/subscriptions/sub/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/logger"
	testStorageLocation = "abfs://logs@mystorageaccount.dfs.core.windows.net"
	testBackupLocation  = "abfs://backup@mystorageaccount.dfs.core.windows.net"
	testProxyName       = "my-proxy-config"
	testCidr            = "10.0.0.0/16"
	testDefaultSG       = "sg-default-123"
	testKnoxSG          = "sg-knox-456"
	testNetworkID       = "/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/my-vnet"
	testRGName          = "my-resource-group"
	testAksZoneID       = "/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Network/privateDnsZones/aks-zone"
	testDbZoneID        = "/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Network/privateDnsZones/db-zone"
)

func baseAzureEnvironment() *environmentsmodels.Environment {
	return &environmentsmodels.Environment{
		Crn:               new(testCrn),
		CredentialName:    new(testCredential),
		Description:       testDescription,
		EnvironmentName:   new(testEnvName),
		TunnelEnabled:     true,
		Status:            new(testStatus),
		StatusReason:      testStatusReason,
		WorkloadAnalytics: true,
		Freeipa:           &environmentsmodels.FreeipaDetails{},
		LogStorage: &environmentsmodels.LogStorage{
			AzureDetails: &environmentsmodels.LogStorageAzureDetails{
				ManagedIdentity:     testManagedIdentity,
				StorageLocationBase: testStorageLocation,
			},
		},
		Network: &environmentsmodels.Network{
			SubnetIds:                      []string{"subnet-1", "subnet-2"},
			EndpointAccessGatewaySubnetIds: []string{"eag-subnet-1"},
			Azure: &environmentsmodels.NetworkAzureParams{
				AksPrivateDNSZoneID:      testAksZoneID,
				DatabasePrivateDNSZoneID: testDbZoneID,
				NetworkID:                new(testNetworkID),
				ResourceGroupName:        new(testRGName),
				FlexibleServerSubnetIds:  []string{"flex-subnet-1"},
				UsePublicIP:              new(true),
			},
		},
		ProxyConfig: &environmentsmodels.ProxyConfig{
			ProxyConfigName: new(testProxyName),
		},
		SecurityAccess: &environmentsmodels.SecurityAccess{
			Cidr:                   testCidr,
			DefaultSecurityGroupID: testDefaultSG,
			SecurityGroupIDForKnox: testKnoxSG,
		},
		Tags: &environmentsmodels.EnvironmentTags{
			UserDefined: map[string]string{
				"team": "platform",
				"env":  "dev",
			},
		},
	}
}

func TestToAzureEnvironmentResource_BasicFields(t *testing.T) {
	ctx := context.TODO()
	env := baseAzureEnvironment()
	model := &azureEnvironmentResourceModel{}
	diags := diag.Diagnostics{}

	toAzureEnvironmentResource(ctx, env, model, nil, &diags)

	assert.False(t, diags.HasError())
	assert.Equal(t, testCrn, model.ID.ValueString())
	assert.Equal(t, testCrn, model.Crn.ValueString())
	assert.Equal(t, testCredential, model.CredentialName.ValueString())
	assert.Equal(t, testDescription, model.Description.ValueString())
	assert.Equal(t, testEnvName, model.EnvironmentName.ValueString())
	assert.Equal(t, true, model.EnableTunnel.ValueBool())
	assert.Equal(t, testStatus, model.Status.ValueString())
	assert.Equal(t, testStatusReason, model.StatusReason.ValueString())
	assert.Equal(t, true, model.WorkloadAnalytics.ValueBool())
}

func TestToAzureEnvironmentResource_LogStorage(t *testing.T) {
	ctx := context.TODO()
	env := baseAzureEnvironment()
	model := &azureEnvironmentResourceModel{}
	diags := diag.Diagnostics{}

	toAzureEnvironmentResource(ctx, env, model, nil, &diags)

	assert.False(t, diags.HasError())
	assert.NotNil(t, model.LogStorage)
	assert.Equal(t, testManagedIdentity, model.LogStorage.ManagedIdentity.ValueString())
	assert.Equal(t, testStorageLocation, model.LogStorage.StorageLocationBase.ValueString())
}

func TestToAzureEnvironmentResource_LogStorageWithBackup(t *testing.T) {
	ctx := context.TODO()
	env := baseAzureEnvironment()
	env.BackupStorage = &environmentsmodels.BackupStorage{
		AzureDetails: &environmentsmodels.BackupStorageAzureDetails{
			StorageLocationBase: testBackupLocation,
		},
	}
	model := &azureEnvironmentResourceModel{}
	diags := diag.Diagnostics{}

	toAzureEnvironmentResource(ctx, env, model, nil, &diags)

	assert.False(t, diags.HasError())
	assert.NotNil(t, model.LogStorage)
	assert.Equal(t, testBackupLocation, model.LogStorage.BackupStorageLocationBase.ValueString())
}

func TestToAzureEnvironmentResource_LogStorageNilAzureDetails(t *testing.T) {
	ctx := context.TODO()
	env := baseAzureEnvironment()
	env.LogStorage = &environmentsmodels.LogStorage{AzureDetails: nil}
	model := &azureEnvironmentResourceModel{}
	diags := diag.Diagnostics{}

	toAzureEnvironmentResource(ctx, env, model, nil, &diags)

	assert.False(t, diags.HasError())
	assert.Nil(t, model.LogStorage)
}

func TestToAzureEnvironmentResource_Network(t *testing.T) {
	ctx := context.TODO()
	env := baseAzureEnvironment()
	model := &azureEnvironmentResourceModel{}
	diags := diag.Diagnostics{}

	toAzureEnvironmentResource(ctx, env, model, nil, &diags)

	assert.False(t, diags.HasError())
	assert.True(t, model.UsePublicIP.ValueBool())

	assert.False(t, model.EndpointAccessGatewaySubnetIds.IsNull())
	eagSubnets := utils.FromSetValueToStringList(model.EndpointAccessGatewaySubnetIds)
	assert.Contains(t, eagSubnets, "eag-subnet-1")

	assert.False(t, model.ExistingNetworkParams.IsNull())
}

func TestToAzureEnvironmentResource_NetworkEmptyEAGSubnets(t *testing.T) {
	ctx := context.TODO()
	env := baseAzureEnvironment()
	env.Network.EndpointAccessGatewaySubnetIds = []string{}
	model := &azureEnvironmentResourceModel{}
	diags := diag.Diagnostics{}

	toAzureEnvironmentResource(ctx, env, model, nil, &diags)

	assert.False(t, diags.HasError())
	assert.True(t, model.EndpointAccessGatewaySubnetIds.IsNull())
}

func TestToAzureEnvironmentResource_NetworkNilAzureParams(t *testing.T) {
	ctx := context.TODO()
	env := baseAzureEnvironment()
	env.Network.Azure = nil
	model := &azureEnvironmentResourceModel{}
	diags := diag.Diagnostics{}

	toAzureEnvironmentResource(ctx, env, model, nil, &diags)

	assert.False(t, diags.HasError())
	assert.True(t, model.ExistingNetworkParams.IsNull())
	assert.True(t, model.UsePublicIP.IsNull())
}

func TestToAzureEnvironmentResource_NilNetwork(t *testing.T) {
	ctx := context.TODO()
	env := baseAzureEnvironment()
	env.Network = nil
	model := &azureEnvironmentResourceModel{}
	diags := diag.Diagnostics{}

	toAzureEnvironmentResource(ctx, env, model, nil, &diags)

	assert.False(t, diags.HasError())
	assert.True(t, model.ExistingNetworkParams.IsNull())
	assert.True(t, model.EndpointAccessGatewaySubnetIds.IsNull())
}

func TestToAzureEnvironmentResource_ProxyConfig(t *testing.T) {
	ctx := context.TODO()
	env := baseAzureEnvironment()
	model := &azureEnvironmentResourceModel{}
	diags := diag.Diagnostics{}

	toAzureEnvironmentResource(ctx, env, model, nil, &diags)

	assert.False(t, diags.HasError())
	assert.Equal(t, testProxyName, model.ProxyConfigName.ValueString())
}

func TestToAzureEnvironmentResource_NilProxyConfig(t *testing.T) {
	ctx := context.TODO()
	env := baseAzureEnvironment()
	env.ProxyConfig = nil
	model := &azureEnvironmentResourceModel{}
	diags := diag.Diagnostics{}

	toAzureEnvironmentResource(ctx, env, model, nil, &diags)

	assert.False(t, diags.HasError())
	assert.True(t, model.ProxyConfigName.IsNull())
}

func TestToAzureEnvironmentResource_SecurityAccess(t *testing.T) {
	ctx := context.TODO()
	env := baseAzureEnvironment()
	model := &azureEnvironmentResourceModel{}
	diags := diag.Diagnostics{}

	toAzureEnvironmentResource(ctx, env, model, nil, &diags)

	assert.False(t, diags.HasError())
	assert.NotNil(t, model.SecurityAccess)
	assert.Equal(t, testCidr, model.SecurityAccess.Cidr.ValueString())
	assert.Equal(t, testDefaultSG, model.SecurityAccess.DefaultSecurityGroupID.ValueString())
	assert.Equal(t, testKnoxSG, model.SecurityAccess.SecurityGroupIDForKnox.ValueString())
}

func TestToAzureEnvironmentResource_SecurityAccessPreservesExistingGroupIDs(t *testing.T) {
	ctx := context.TODO()
	env := baseAzureEnvironment()
	existingGroupIDs := utils.ToSetValueFromStringList([]string{"sg-a", "sg-b"})
	model := &azureEnvironmentResourceModel{
		SecurityAccess: &SecurityAccess{
			DefaultSecurityGroupIDs: existingGroupIDs,
			SecurityGroupIDsForKnox: existingGroupIDs,
		},
	}
	diags := diag.Diagnostics{}

	toAzureEnvironmentResource(ctx, env, model, nil, &diags)

	assert.False(t, diags.HasError())
	assert.Equal(t, existingGroupIDs, model.SecurityAccess.DefaultSecurityGroupIDs)
}

func TestToAzureEnvironmentResource_SecurityAccessNilModelSetsNullGroupIDs(t *testing.T) {
	ctx := context.TODO()
	env := baseAzureEnvironment()
	model := &azureEnvironmentResourceModel{
		SecurityAccess: nil,
	}
	diags := diag.Diagnostics{}

	toAzureEnvironmentResource(ctx, env, model, nil, &diags)

	assert.False(t, diags.HasError())
	assert.NotNil(t, model.SecurityAccess)
	assert.True(t, model.SecurityAccess.DefaultSecurityGroupIDs.IsNull())
	assert.True(t, model.SecurityAccess.SecurityGroupIDsForKnox.IsNull())
}

func TestToAzureEnvironmentResource_NilSecurityAccess(t *testing.T) {
	ctx := context.TODO()
	env := baseAzureEnvironment()
	env.SecurityAccess = nil
	model := &azureEnvironmentResourceModel{}
	diags := diag.Diagnostics{}

	toAzureEnvironmentResource(ctx, env, model, nil, &diags)

	assert.False(t, diags.HasError())
	assert.Nil(t, model.SecurityAccess)
}

func TestToAzureEnvironmentResource_Tags(t *testing.T) {
	ctx := context.TODO()
	env := baseAzureEnvironment()
	model := &azureEnvironmentResourceModel{}
	diags := diag.Diagnostics{}

	toAzureEnvironmentResource(ctx, env, model, nil, &diags)

	assert.False(t, diags.HasError())
	assert.False(t, model.Tags.IsNull())
	tags := make(map[string]string)
	model.Tags.ElementsAs(ctx, &tags, false)
	assert.Equal(t, "platform", tags["team"])
	assert.Equal(t, "dev", tags["env"])
}

func TestToAzureEnvironmentResource_NilTags(t *testing.T) {
	ctx := context.TODO()
	env := baseAzureEnvironment()
	env.Tags = nil
	model := &azureEnvironmentResourceModel{}
	diags := diag.Diagnostics{}

	toAzureEnvironmentResource(ctx, env, model, nil, &diags)

	assert.False(t, diags.HasError())
	assert.True(t, model.Tags.IsNull())
}

func TestToAzureEnvironmentResource_PollingOptions(t *testing.T) {
	ctx := context.TODO()
	env := baseAzureEnvironment()
	pollingOpts := &utils.PollingOptions{
		Async: types.BoolValue(true),
	}
	model := &azureEnvironmentResourceModel{}
	diags := diag.Diagnostics{}

	toAzureEnvironmentResource(ctx, env, model, pollingOpts, &diags)

	assert.False(t, diags.HasError())
	assert.Equal(t, pollingOpts, model.PollingOptions)
}

func TestToAzureEnvironmentResource_NilLogStorage(t *testing.T) {
	ctx := context.TODO()
	env := baseAzureEnvironment()
	env.LogStorage = nil
	model := &azureEnvironmentResourceModel{}
	diags := diag.Diagnostics{}

	toAzureEnvironmentResource(ctx, env, model, nil, &diags)

	assert.False(t, diags.HasError())
	assert.Nil(t, model.LogStorage)
}

func TestToAzureEnvironmentResource_FullEnvironment(t *testing.T) {
	ctx := context.TODO()
	env := baseAzureEnvironment()
	env.BackupStorage = &environmentsmodels.BackupStorage{
		AzureDetails: &environmentsmodels.BackupStorageAzureDetails{
			StorageLocationBase: testBackupLocation,
		},
	}
	model := &azureEnvironmentResourceModel{}
	diags := diag.Diagnostics{}

	toAzureEnvironmentResource(ctx, env, model, nil, &diags)

	assert.False(t, diags.HasError())
	assert.Equal(t, testCrn, model.ID.ValueString())
	assert.Equal(t, testCrn, model.Crn.ValueString())
	assert.Equal(t, testCredential, model.CredentialName.ValueString())
	assert.Equal(t, testEnvName, model.EnvironmentName.ValueString())
	assert.Equal(t, testDescription, model.Description.ValueString())
	assert.True(t, model.EnableTunnel.ValueBool())
	assert.True(t, model.WorkloadAnalytics.ValueBool())
	assert.Equal(t, testStatus, model.Status.ValueString())
	assert.Equal(t, testStatusReason, model.StatusReason.ValueString())
	assert.NotNil(t, model.LogStorage)
	assert.Equal(t, testManagedIdentity, model.LogStorage.ManagedIdentity.ValueString())
	assert.Equal(t, testStorageLocation, model.LogStorage.StorageLocationBase.ValueString())
	assert.Equal(t, testBackupLocation, model.LogStorage.BackupStorageLocationBase.ValueString())
	assert.NotNil(t, model.SecurityAccess)
	assert.Equal(t, testCidr, model.SecurityAccess.Cidr.ValueString())
	assert.Equal(t, testDefaultSG, model.SecurityAccess.DefaultSecurityGroupID.ValueString())
	assert.Equal(t, testKnoxSG, model.SecurityAccess.SecurityGroupIDForKnox.ValueString())
	assert.Equal(t, testProxyName, model.ProxyConfigName.ValueString())
	assert.True(t, model.UsePublicIP.ValueBool())
}
