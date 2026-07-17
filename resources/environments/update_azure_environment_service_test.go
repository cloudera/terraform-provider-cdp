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
	"errors"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/mocks"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

const (
	testEncryptionKeyURL               = "https://my-vault.vault.azure.net/keys/my-key/abc123"
	testNewEncryptionKeyURL            = "https://new-vault.vault.azure.net/keys/new-key/xyz789"
	testEncryptionKeyResourceGroupName = "my-encryption-rg"
	testNewEncryptionKeyRG             = "new-rg"
	testEncryptionUserManagedIdentity  = "/subscriptions/sub-id/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/my-identity"
	testAzureSubnet                    = "subnet-a"
	testAzureIpRange                   = "10.0.0.0/24"

	testOldSharedManagedIdentity = "/subscriptions/sub-id/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/old-ds-identity"
	testNewSharedManagedIdentity = "/subscriptions/sub-id/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/new-ds-identity"
	testOldAksPrivateDnsZoneId   = "/subscriptions/sub-id/resourceGroups/rg/providers/Microsoft.Network/privateDnsZones/old-zone"
	testNewAksPrivateDnsZoneId   = "/subscriptions/sub-id/resourceGroups/rg/providers/Microsoft.Network/privateDnsZones/new-zone"
)

func TestConvertNilConfigReturnsNilRequestForAzure(t *testing.T) {
	config := (*AzureComputeClusterConfiguration)(nil)
	fallbackSubnetIds := types.Set{}
	want := (*environmentsmodels.AzureComputeClusterConfigurationRequest)(nil)
	if got := convertConfigToAzureComputeClusterConfigurationRequest(config, fallbackSubnetIds); !reflect.DeepEqual(got, want) {
		t.Errorf("convertConfigToAzureComputeClusterConfigurationRequest() = %v, want %v", got, want)
	}
}

func TestConvertValidConfigReturnsCorrectRequestForAzure(t *testing.T) {
	config := &AzureComputeClusterConfiguration{
		KubeApiAuthorizedIpRanges: utils.ToSetValueFromStringList([]string{"192.168.1.1/32", "10.0.0.0/24"}),
		PrivateCluster:            types.BoolValue(true),
		WorkerNodeSubnets:         utils.ToSetValueFromStringList([]string{"subnet-1", "subnet-2"}),
		OutboundType:              types.StringValue("udr"),
	}
	fallbackSubnetIds := types.Set{}
	want := &environmentsmodels.AzureComputeClusterConfigurationRequest{
		KubeAPIAuthorizedIPRanges: []string{"192.168.1.1/32", "10.0.0.0/24"},
		PrivateCluster:            true,
		WorkerNodeSubnets:         []string{"subnet-1", "subnet-2"},
		OutboundType:              "udr",
	}
	if got := convertConfigToAzureComputeClusterConfigurationRequest(config, fallbackSubnetIds); !reflect.DeepEqual(got, want) {
		t.Errorf("convertConfigToAzureComputeClusterConfigurationRequest() = %v, want %v", got, want)
	}
}

func TestConvertConfigWithEmptyFieldsReturnsCorrectRequestForAzure(t *testing.T) {
	config := &AzureComputeClusterConfiguration{
		KubeApiAuthorizedIpRanges: utils.ToSetValueFromStringList([]string{}),
		PrivateCluster:            types.BoolValue(false),
		WorkerNodeSubnets:         utils.ToSetValueFromStringList([]string{}),
		OutboundType:              types.StringValue(""),
	}
	fallbackSubnetIds := types.Set{}
	want := &environmentsmodels.AzureComputeClusterConfigurationRequest{
		KubeAPIAuthorizedIPRanges: []string{},
		PrivateCluster:            false,
		WorkerNodeSubnets:         []string{},
		OutboundType:              "",
	}
	if got := convertConfigToAzureComputeClusterConfigurationRequest(config, fallbackSubnetIds); !reflect.DeepEqual(got, want) {
		t.Errorf("convertConfigToAzureComputeClusterConfigurationRequest() = %v, want %v", got, want)
	}
}

func TestConvertConfigWithFallbackSubnetsReturnsCorrectRequestForAzure(t *testing.T) {
	config := &AzureComputeClusterConfiguration{
		KubeApiAuthorizedIpRanges: utils.ToSetValueFromStringList([]string{"192.168.1.1/32"}),
		PrivateCluster:            types.BoolValue(true),
		WorkerNodeSubnets:         types.Set{},
		OutboundType:              types.StringValue("udr"),
	}
	fallbackSubnetIds := utils.ToSetValueFromStringList([]string{"subnet-1", "subnet-2"})
	want := &environmentsmodels.AzureComputeClusterConfigurationRequest{
		KubeAPIAuthorizedIPRanges: []string{"192.168.1.1/32"},
		PrivateCluster:            true,
		WorkerNodeSubnets:         []string{"subnet-1", "subnet-2"},
		OutboundType:              "udr",
	}
	if got := convertConfigToAzureComputeClusterConfigurationRequest(config, fallbackSubnetIds); !reflect.DeepEqual(got, want) {
		t.Errorf("convertConfigToAzureComputeClusterConfigurationRequest() = %v, want %v", got, want)
	}
}

func TestUpdateAzureAvailabilityZonesIfChanged_ZonesChanged_CallsAPIAndUpdatesState(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	plan := utils.ToSetValueFromStringList([]string{"1", "2", "3"})
	state := utils.ToSetValueFromStringList([]string{"1", "2"})
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateAzureAvailabilityZonesContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateAzureAvailabilityZonesParams) bool {
		return params.Input != nil &&
			*params.Input.Environment == testEnvName &&
			len(params.Input.AvailabilityZones) == 3
	}), mock.Anything).Return(&operations.UpdateAzureAvailabilityZonesOK{}, nil)

	result := updateAvailabilityZones(ctx, client, plan, &state, new(testEnvName), resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, plan, state)
	mockClient.AssertExpectations(t)
}

func TestUpdateAzureAvailabilityZonesIfChanged_ZonesUnchanged_SkipsAPICall(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	zones := utils.ToSetValueFromStringList([]string{"1", "2"})
	resp := &resource.UpdateResponse{}

	result := updateAvailabilityZones(ctx, client, zones, new(utils.ToSetValueFromStringList([]string{"1", "2"})), new(testEnvName), resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateAzureAvailabilityZonesContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureAvailabilityZonesIfChanged_PlanNull_SkipsAPICall(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	plan := types.SetNull(types.StringType)
	resp := &resource.UpdateResponse{}

	result := updateAvailabilityZones(ctx, client, plan, new(utils.ToSetValueFromStringList([]string{"1", "2"})), new(testEnvName), resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateAzureAvailabilityZonesContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureAvailabilityZonesIfChanged_PlanUnknown_SkipsAPICall(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	plan := types.SetUnknown(types.StringType)
	resp := &resource.UpdateResponse{}

	result := updateAvailabilityZones(ctx, client, plan, new(utils.ToSetValueFromStringList([]string{"1", "2"})), new(testEnvName), resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateAzureAvailabilityZonesContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureAvailabilityZonesIfChanged_PlanEmpty_AddsValidationError(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	plan := utils.ToSetValueFromStringList([]string{})
	resp := &resource.UpdateResponse{}

	result := updateAvailabilityZones(ctx, client, plan, new(utils.ToSetValueFromStringList([]string{"1", "2"})), new(testEnvName), resp)

	assert.True(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateAzureAvailabilityZonesContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureAvailabilityZonesIfChanged_APIError_AddsDiagnosticError(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	plan := utils.ToSetValueFromStringList([]string{"1", "2", "3"})
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateAzureAvailabilityZonesContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpdateAzureAvailabilityZonesOK)(nil), errors.New("API connection failed"))

	result := updateAvailabilityZones(ctx, client, plan, new(utils.ToSetValueFromStringList([]string{"1", "2"})), new(testEnvName), resp)

	assert.True(t, result.Diagnostics.HasError())
}

func TestUpdateAzureAvailabilityZonesIfChanged_APISuccess_UpdatesStateToPlan(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	plan := utils.ToSetValueFromStringList([]string{"1", "3"})
	state := utils.ToSetValueFromStringList([]string{"1", "2"})
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateAzureAvailabilityZonesContext", mock.Anything, mock.Anything, mock.Anything).
		Return(&operations.UpdateAzureAvailabilityZonesOK{}, nil)

	updateAvailabilityZones(ctx, client, plan, &state, new(testEnvName), resp)

	assert.Equal(t, plan, state)
}

func TestUpdateAzureEncryptionResources_Success(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EncryptionKeyURL:               types.StringValue(testEncryptionKeyURL),
		EncryptionKeyResourceGroupName: types.StringValue(testEncryptionKeyResourceGroupName),
		EncryptionUserManagedIdentity:  types.StringValue(testEncryptionUserManagedIdentity),
	}

	mockClient.On("UpdateAzureEncryptionResourcesContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateAzureEncryptionResourcesParams) bool {
		return params.Input != nil &&
			*params.Input.EncryptionKeyURL == testEncryptionKeyURL &&
			*params.Input.Environment == testEnvName &&
			params.Input.EncryptionKeyResourceGroupName == testEncryptionKeyResourceGroupName &&
			params.Input.EncryptionUserManagedIdentity == testEncryptionUserManagedIdentity
	})).
		Return(&operations.UpdateAzureEncryptionResourcesOK{}, nil)

	err := updateAzureEncryptionResources(ctx, client, new(testEnvName), plan)

	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	mockClient.AssertExpectations(t)
}

func TestUpdateAzureEncryptionResources_ReturnsError(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EncryptionKeyURL:               types.StringValue(testEncryptionKeyURL),
		EncryptionKeyResourceGroupName: types.StringValue(testEncryptionKeyResourceGroupName),
		EncryptionUserManagedIdentity:  types.StringValue(testEncryptionUserManagedIdentity),
	}

	mockClient.On("UpdateAzureEncryptionResourcesContext", mock.Anything, mock.Anything).
		Return((*operations.UpdateAzureEncryptionResourcesOK)(nil), errors.New(testServiceUnavailable))

	err := updateAzureEncryptionResources(ctx, client, new(testEnvName), plan)

	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	if err.Error() != testServiceUnavailable {
		t.Errorf("expected error message '%s', got: %s", testServiceUnavailable, err.Error())
	}
}

// Tests for azureEncryptionFieldsChanged

func TestAzureEncryptionFieldsChanged_AllSame_ReturnsFalse(t *testing.T) {
	plan := &azureEnvironmentResourceModel{
		EncryptionKeyURL:               types.StringValue(testEncryptionKeyURL),
		EncryptionKeyResourceGroupName: types.StringValue(testEncryptionKeyResourceGroupName),
		EncryptionUserManagedIdentity:  types.StringValue(testEncryptionUserManagedIdentity),
	}
	state := &azureEnvironmentResourceModel{
		EncryptionKeyURL:               types.StringValue(testEncryptionKeyURL),
		EncryptionKeyResourceGroupName: types.StringValue(testEncryptionKeyResourceGroupName),
		EncryptionUserManagedIdentity:  types.StringValue(testEncryptionUserManagedIdentity),
	}

	assert.False(t, azureEncryptionFieldsChanged(plan, state))
}

func TestAzureEncryptionFieldsChanged_KeyURLDiffers_ReturnsTrue(t *testing.T) {
	plan := &azureEnvironmentResourceModel{
		EncryptionKeyURL:               types.StringValue("https://new-vault.vault.azure.net/keys/new-key/def456"),
		EncryptionKeyResourceGroupName: types.StringValue(testEncryptionKeyResourceGroupName),
		EncryptionUserManagedIdentity:  types.StringValue(testEncryptionUserManagedIdentity),
	}
	state := &azureEnvironmentResourceModel{
		EncryptionKeyURL:               types.StringValue(testEncryptionKeyURL),
		EncryptionKeyResourceGroupName: types.StringValue(testEncryptionKeyResourceGroupName),
		EncryptionUserManagedIdentity:  types.StringValue(testEncryptionUserManagedIdentity),
	}

	assert.True(t, azureEncryptionFieldsChanged(plan, state))
}

func TestAzureEncryptionFieldsChanged_ResourceGroupDiffers_ReturnsTrue(t *testing.T) {
	plan := &azureEnvironmentResourceModel{
		EncryptionKeyURL:               types.StringValue(testEncryptionKeyURL),
		EncryptionKeyResourceGroupName: types.StringValue("different-rg"),
		EncryptionUserManagedIdentity:  types.StringValue(testEncryptionUserManagedIdentity),
	}
	state := &azureEnvironmentResourceModel{
		EncryptionKeyURL:               types.StringValue(testEncryptionKeyURL),
		EncryptionKeyResourceGroupName: types.StringValue(testEncryptionKeyResourceGroupName),
		EncryptionUserManagedIdentity:  types.StringValue(testEncryptionUserManagedIdentity),
	}

	assert.True(t, azureEncryptionFieldsChanged(plan, state))
}

func TestAzureEncryptionFieldsChanged_ManagedIdentityDiffers_ReturnsTrue(t *testing.T) {
	plan := &azureEnvironmentResourceModel{
		EncryptionKeyURL:               types.StringValue(testEncryptionKeyURL),
		EncryptionKeyResourceGroupName: types.StringValue(testEncryptionKeyResourceGroupName),
		EncryptionUserManagedIdentity:  types.StringValue("/subscriptions/sub-id/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/different-identity"),
	}
	state := &azureEnvironmentResourceModel{
		EncryptionKeyURL:               types.StringValue(testEncryptionKeyURL),
		EncryptionKeyResourceGroupName: types.StringValue(testEncryptionKeyResourceGroupName),
		EncryptionUserManagedIdentity:  types.StringValue(testEncryptionUserManagedIdentity),
	}

	assert.True(t, azureEncryptionFieldsChanged(plan, state))
}

// Tests for updateAzureEncryptionIfChanged

func TestUpdateAzureEncryptionIfChanged_NoChange_Skips(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EncryptionKeyURL:               types.StringValue(testEncryptionKeyURL),
		EncryptionKeyResourceGroupName: types.StringValue(testEncryptionKeyResourceGroupName),
		EncryptionUserManagedIdentity:  types.StringValue(testEncryptionUserManagedIdentity),
		EnvironmentName:                types.StringValue(testEnvName),
	}
	state := &azureEnvironmentResourceModel{
		EncryptionKeyURL:               types.StringValue(testEncryptionKeyURL),
		EncryptionKeyResourceGroupName: types.StringValue(testEncryptionKeyResourceGroupName),
		EncryptionUserManagedIdentity:  types.StringValue(testEncryptionUserManagedIdentity),
	}
	resp := &resource.UpdateResponse{}

	result := updateAzureEncryptionIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateAzureEncryptionResourcesContext", mock.Anything, mock.Anything)
}

func TestUpdateAzureEncryptionIfChanged_Changed_CallsAPIAndUpdatesState(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EncryptionKeyURL:               types.StringValue(testNewEncryptionKeyURL),
		EncryptionKeyResourceGroupName: types.StringValue(testNewEncryptionKeyRG),
		EncryptionUserManagedIdentity:  types.StringValue(testEncryptionUserManagedIdentity),
		EnvironmentName:                types.StringValue(testEnvName),
	}
	state := &azureEnvironmentResourceModel{
		EncryptionKeyURL:               types.StringValue(testEncryptionKeyURL),
		EncryptionKeyResourceGroupName: types.StringValue(testEncryptionKeyResourceGroupName),
		EncryptionUserManagedIdentity:  types.StringValue(testEncryptionUserManagedIdentity),
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateAzureEncryptionResourcesContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateAzureEncryptionResourcesParams) bool {
		return params.Input != nil &&
			*params.Input.EncryptionKeyURL == testNewEncryptionKeyURL &&
			params.Input.EncryptionKeyResourceGroupName == testNewEncryptionKeyRG
	})).Return(&operations.UpdateAzureEncryptionResourcesOK{}, nil)

	result := updateAzureEncryptionIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testNewEncryptionKeyURL), state.EncryptionKeyURL)
	assert.Equal(t, types.StringValue(testNewEncryptionKeyRG), state.EncryptionKeyResourceGroupName)
	mockClient.AssertExpectations(t)
}

func TestUpdateAzureEncryptionIfChanged_NullKeyURL_AddsDiagnostic(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EncryptionKeyURL:               types.StringNull(),
		EncryptionKeyResourceGroupName: types.StringValue(testNewEncryptionKeyRG),
		EncryptionUserManagedIdentity:  types.StringValue(testEncryptionUserManagedIdentity),
		EnvironmentName:                types.StringValue(testEnvName),
	}
	state := &azureEnvironmentResourceModel{
		EncryptionKeyURL:               types.StringValue(testEncryptionKeyURL),
		EncryptionKeyResourceGroupName: types.StringValue(testEncryptionKeyResourceGroupName),
		EncryptionUserManagedIdentity:  types.StringValue(testEncryptionUserManagedIdentity),
	}
	resp := &resource.UpdateResponse{}

	result := updateAzureEncryptionIfChanged(ctx, plan, state, client, resp)

	assert.True(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateAzureEncryptionResourcesContext", mock.Anything, mock.Anything)
}

func TestUpdateAzureEncryptionIfChanged_APIError_AddsDiagnostic(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EncryptionKeyURL:               types.StringValue(testNewEncryptionKeyURL),
		EncryptionKeyResourceGroupName: types.StringValue(testEncryptionKeyResourceGroupName),
		EncryptionUserManagedIdentity:  types.StringValue(testEncryptionUserManagedIdentity),
		EnvironmentName:                types.StringValue(testEnvName),
	}
	state := &azureEnvironmentResourceModel{
		EncryptionKeyURL:               types.StringValue(testEncryptionKeyURL),
		EncryptionKeyResourceGroupName: types.StringValue(testEncryptionKeyResourceGroupName),
		EncryptionUserManagedIdentity:  types.StringValue(testEncryptionUserManagedIdentity),
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateAzureEncryptionResourcesContext", mock.Anything, mock.Anything).
		Return((*operations.UpdateAzureEncryptionResourcesOK)(nil), errors.New("forbidden"))

	result := updateAzureEncryptionIfChanged(ctx, plan, state, client, resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testEncryptionKeyURL), state.EncryptionKeyURL)
}

// Tests for enableComputeClusterForAzure

func TestEnableComputeClusterForAzure_Success(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	config := &AzureComputeClusterConfiguration{
		KubeApiAuthorizedIpRanges: utils.ToSetValueFromStringList([]string{testAzureIpRange}),
		PrivateCluster:            types.BoolValue(true),
		WorkerNodeSubnets:         utils.ToSetValueFromStringList([]string{testAzureSubnet}),
		OutboundType:              types.StringValue("udr"),
	}
	envSubnets := utils.ToSetValueFromStringList([]string{"subnet-fallback"})

	mockClient.On("InitializeAzureComputeClusterContext", mock.Anything, mock.MatchedBy(func(params *operations.InitializeAzureComputeClusterParams) bool {
		return params.Input != nil &&
			*params.Input.EnvironmentName == testEnvName &&
			params.Input.ComputeClusterConfiguration != nil &&
			params.Input.ComputeClusterConfiguration.PrivateCluster == true &&
			params.Input.ComputeClusterConfiguration.OutboundType == "udr"
	}), mock.Anything).Return(&operations.InitializeAzureComputeClusterOK{}, nil)

	err := enableComputeClusterForAzure(ctx, config, testEnvName, envSubnets, client)

	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}

func TestEnableComputeClusterForAzure_APIError_ReturnsError(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	config := &AzureComputeClusterConfiguration{
		KubeApiAuthorizedIpRanges: utils.ToSetValueFromStringList([]string{}),
		PrivateCluster:            types.BoolValue(false),
		WorkerNodeSubnets:         utils.ToSetValueFromStringList([]string{testAzureSubnet}),
		OutboundType:              types.StringValue(""),
	}
	envSubnets := utils.ToSetValueFromStringList([]string{testAzureSubnet})

	mockClient.On("InitializeAzureComputeClusterContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.InitializeAzureComputeClusterOK)(nil), errors.New(testClusterInitFailed))

	err := enableComputeClusterForAzure(ctx, config, testEnvName, envSubnets, client)

	assert.Error(t, err)
	assert.Equal(t, testClusterInitFailed, err.Error())
}

// Tests for updateAzureSecurityAccessIfChanged

func TestUpdateAzureSecurityAccessIfChanged_Changed_CallsAPI(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		SecurityAccess: &SecurityAccess{
			DefaultSecurityGroupID: types.StringValue(testNewDefaultSG),
			SecurityGroupIDForKnox: types.StringValue(testNewKnoxSG),
		},
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &azureEnvironmentResourceModel{
		SecurityAccess: &SecurityAccess{
			DefaultSecurityGroupID: types.StringValue(testOldDefaultSG),
			SecurityGroupIDForKnox: types.StringValue(testOldKnoxSG),
		},
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateSecurityAccessContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateSecurityAccessParams) bool {
		return params.Input != nil &&
			*params.Input.DefaultSecurityGroupID == testNewDefaultSG &&
			*params.Input.GatewayNodeSecurityGroupID == testNewKnoxSG &&
			*params.Input.Environment == testEnvName
	}), mock.Anything).Return(&operations.UpdateSecurityAccessOK{}, nil)

	result := updateAzureSecurityAccessIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testNewDefaultSG), state.SecurityAccess.DefaultSecurityGroupID)
	assert.Equal(t, types.StringValue(testNewKnoxSG), state.SecurityAccess.SecurityGroupIDForKnox)
	mockClient.AssertExpectations(t)
}

func TestUpdateAzureSecurityAccessIfChanged_Unchanged_SkipsAPICall(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		SecurityAccess: &SecurityAccess{
			DefaultSecurityGroupID: types.StringValue(testSameDefaultSG),
			SecurityGroupIDForKnox: types.StringValue(testSameKnoxSG),
		},
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &azureEnvironmentResourceModel{
		SecurityAccess: &SecurityAccess{
			DefaultSecurityGroupID: types.StringValue(testSameDefaultSG),
			SecurityGroupIDForKnox: types.StringValue(testSameKnoxSG),
		},
	}
	resp := &resource.UpdateResponse{}

	result := updateAzureSecurityAccessIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateSecurityAccessContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureSecurityAccessIfChanged_OnlySetFieldsDiffer_SkipsAPICall(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		SecurityAccess: &SecurityAccess{
			DefaultSecurityGroupID:  types.StringValue(testSameDefaultSG),
			SecurityGroupIDForKnox:  types.StringValue(testSameKnoxSG),
			DefaultSecurityGroupIDs: utils.ToSetValueFromStringList([]string{"sg-b", "sg-a"}),
		},
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &azureEnvironmentResourceModel{
		SecurityAccess: &SecurityAccess{
			DefaultSecurityGroupID:  types.StringValue(testSameDefaultSG),
			SecurityGroupIDForKnox:  types.StringValue(testSameKnoxSG),
			DefaultSecurityGroupIDs: utils.ToSetValueFromStringList([]string{"sg-a", "sg-b"}),
		},
	}
	resp := &resource.UpdateResponse{}

	result := updateAzureSecurityAccessIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateSecurityAccessContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureSecurityAccessIfChanged_APIError_AddsDiagnostic(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		SecurityAccess: &SecurityAccess{
			DefaultSecurityGroupID: types.StringValue(testNewDefaultSG),
			SecurityGroupIDForKnox: types.StringValue(testNewKnoxSG),
		},
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &azureEnvironmentResourceModel{
		SecurityAccess: &SecurityAccess{
			DefaultSecurityGroupID: types.StringValue(testOldDefaultSG),
			SecurityGroupIDForKnox: types.StringValue(testOldKnoxSG),
		},
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateSecurityAccessContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpdateSecurityAccessOK)(nil), errors.New(testServiceUnavailable))

	result := updateAzureSecurityAccessIfChanged(ctx, plan, state, client, resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testOldDefaultSG), state.SecurityAccess.DefaultSecurityGroupID)
	assert.Equal(t, types.StringValue(testOldKnoxSG), state.SecurityAccess.SecurityGroupIDForKnox)
}

// Tests for updateAzureDataServicesIfChanged

func TestUpdateAzureDataServicesIfChanged_Changed_CallsAPIAndUpdatesState(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName: types.StringValue(testEnvName),
		DataServices: &DataServices{
			SharedManagedIdentity: types.StringValue(testNewSharedManagedIdentity),
			AksPrivateDnsZoneId:   types.StringValue(testNewAksPrivateDnsZoneId),
		},
	}
	state := &azureEnvironmentResourceModel{
		DataServices: &DataServices{
			SharedManagedIdentity: types.StringValue(testOldSharedManagedIdentity),
			AksPrivateDnsZoneId:   types.StringValue(testOldAksPrivateDnsZoneId),
		},
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateDataServiceResourcesContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateDataServiceResourcesParams) bool {
		return params.Input != nil &&
			*params.Input.Environment == testEnvName &&
			params.Input.DataServices != nil &&
			params.Input.DataServices.Azure != nil &&
			*params.Input.DataServices.Azure.SharedManagedIdentity == testNewSharedManagedIdentity &&
			params.Input.DataServices.Azure.AksPrivateDNSZoneID == testNewAksPrivateDnsZoneId
	}), mock.Anything).Return(&operations.UpdateDataServiceResourcesOK{}, nil)

	result := updateAzureDataServicesIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testNewSharedManagedIdentity), state.DataServices.SharedManagedIdentity)
	assert.Equal(t, types.StringValue(testNewAksPrivateDnsZoneId), state.DataServices.AksPrivateDnsZoneId)
	mockClient.AssertExpectations(t)
}

func TestUpdateAzureDataServicesIfChanged_Unchanged_SkipsAPICall(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	ds := &DataServices{
		SharedManagedIdentity: types.StringValue(testOldSharedManagedIdentity),
		AksPrivateDnsZoneId:   types.StringValue(testOldAksPrivateDnsZoneId),
	}
	plan := &azureEnvironmentResourceModel{
		EnvironmentName: types.StringValue(testEnvName),
		DataServices:    ds,
	}
	state := &azureEnvironmentResourceModel{
		DataServices: &DataServices{
			SharedManagedIdentity: types.StringValue(testOldSharedManagedIdentity),
			AksPrivateDnsZoneId:   types.StringValue(testOldAksPrivateDnsZoneId),
		},
	}
	resp := &resource.UpdateResponse{}

	result := updateAzureDataServicesIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateDataServiceResourcesContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureDataServicesIfChanged_PlanNilStateNonNil_ReturnsRemovalError(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName: types.StringValue(testEnvName),
		DataServices:    nil,
	}
	state := &azureEnvironmentResourceModel{
		DataServices: &DataServices{
			SharedManagedIdentity: types.StringValue(testOldSharedManagedIdentity),
			AksPrivateDnsZoneId:   types.StringValue(testOldAksPrivateDnsZoneId),
		},
	}
	resp := &resource.UpdateResponse{}

	result := updateAzureDataServicesIfChanged(ctx, plan, state, client, resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Contains(t, result.Diagnostics.Errors()[0].Detail(), "cannot be removed once set")
	mockClient.AssertNotCalled(t, "UpdateDataServiceResourcesContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureDataServicesIfChanged_PlanNilStateNil_NoOp(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName: types.StringValue(testEnvName),
		DataServices:    nil,
	}
	state := &azureEnvironmentResourceModel{
		DataServices: nil,
	}
	resp := &resource.UpdateResponse{}

	result := updateAzureDataServicesIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateDataServiceResourcesContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureDataServicesIfChanged_StateNil_CallsAPIAndSetsState(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName: types.StringValue(testEnvName),
		DataServices: &DataServices{
			SharedManagedIdentity: types.StringValue(testNewSharedManagedIdentity),
			AksPrivateDnsZoneId:   types.StringValue(testNewAksPrivateDnsZoneId),
		},
	}
	state := &azureEnvironmentResourceModel{
		DataServices: nil,
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateDataServiceResourcesContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateDataServiceResourcesParams) bool {
		return params.Input != nil &&
			*params.Input.Environment == testEnvName &&
			params.Input.DataServices != nil &&
			params.Input.DataServices.Azure != nil &&
			*params.Input.DataServices.Azure.SharedManagedIdentity == testNewSharedManagedIdentity &&
			params.Input.DataServices.Azure.AksPrivateDNSZoneID == testNewAksPrivateDnsZoneId
	}), mock.Anything).Return(&operations.UpdateDataServiceResourcesOK{}, nil)

	result := updateAzureDataServicesIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.NotNil(t, state.DataServices)
	assert.Equal(t, types.StringValue(testNewSharedManagedIdentity), state.DataServices.SharedManagedIdentity)
	assert.Equal(t, types.StringValue(testNewAksPrivateDnsZoneId), state.DataServices.AksPrivateDnsZoneId)
	mockClient.AssertExpectations(t)
}

func TestUpdateAzureDataServicesIfChanged_APIError_AddsDiagnosticAndPreservesState(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName: types.StringValue(testEnvName),
		DataServices: &DataServices{
			SharedManagedIdentity: types.StringValue(testNewSharedManagedIdentity),
			AksPrivateDnsZoneId:   types.StringValue(testNewAksPrivateDnsZoneId),
		},
	}
	state := &azureEnvironmentResourceModel{
		DataServices: &DataServices{
			SharedManagedIdentity: types.StringValue(testOldSharedManagedIdentity),
			AksPrivateDnsZoneId:   types.StringValue(testOldAksPrivateDnsZoneId),
		},
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateDataServiceResourcesContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpdateDataServiceResourcesOK)(nil), errors.New(testServiceUnavailable))

	result := updateAzureDataServicesIfChanged(ctx, plan, state, client, resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testOldSharedManagedIdentity), state.DataServices.SharedManagedIdentity)
	assert.Equal(t, types.StringValue(testOldAksPrivateDnsZoneId), state.DataServices.AksPrivateDnsZoneId)
}

func TestUpdateAzureDataServicesIfChanged_OnlySharedManagedIdentityChanged_CallsAPI(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName: types.StringValue(testEnvName),
		DataServices: &DataServices{
			SharedManagedIdentity: types.StringValue(testNewSharedManagedIdentity),
			AksPrivateDnsZoneId:   types.StringValue(testOldAksPrivateDnsZoneId),
		},
	}
	state := &azureEnvironmentResourceModel{
		DataServices: &DataServices{
			SharedManagedIdentity: types.StringValue(testOldSharedManagedIdentity),
			AksPrivateDnsZoneId:   types.StringValue(testOldAksPrivateDnsZoneId),
		},
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateDataServiceResourcesContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateDataServiceResourcesParams) bool {
		return params.Input != nil &&
			*params.Input.DataServices.Azure.SharedManagedIdentity == testNewSharedManagedIdentity &&
			params.Input.DataServices.Azure.AksPrivateDNSZoneID == testOldAksPrivateDnsZoneId
	}), mock.Anything).Return(&operations.UpdateDataServiceResourcesOK{}, nil)

	result := updateAzureDataServicesIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testNewSharedManagedIdentity), state.DataServices.SharedManagedIdentity)
	mockClient.AssertExpectations(t)
}

func TestUpdateAzureDataServicesIfChanged_SharedManagedIdentityNull_ReturnsValidationError(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName: types.StringValue(testEnvName),
		DataServices: &DataServices{
			SharedManagedIdentity: types.StringNull(),
			AksPrivateDnsZoneId:   types.StringValue(testNewAksPrivateDnsZoneId),
		},
	}
	state := &azureEnvironmentResourceModel{
		DataServices: &DataServices{
			SharedManagedIdentity: types.StringValue(testOldSharedManagedIdentity),
			AksPrivateDnsZoneId:   types.StringValue(testOldAksPrivateDnsZoneId),
		},
	}
	resp := &resource.UpdateResponse{}

	result := updateAzureDataServicesIfChanged(ctx, plan, state, client, resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Contains(t, result.Diagnostics.Errors()[0].Detail(), "shared_managed_identity must be set")
	mockClient.AssertNotCalled(t, "UpdateDataServiceResourcesContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureDataServicesIfChanged_SharedManagedIdentityUnknown_ReturnsValidationError(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName: types.StringValue(testEnvName),
		DataServices: &DataServices{
			SharedManagedIdentity: types.StringUnknown(),
			AksPrivateDnsZoneId:   types.StringValue(testNewAksPrivateDnsZoneId),
		},
	}
	state := &azureEnvironmentResourceModel{
		DataServices: &DataServices{
			SharedManagedIdentity: types.StringValue(testOldSharedManagedIdentity),
			AksPrivateDnsZoneId:   types.StringValue(testOldAksPrivateDnsZoneId),
		},
	}
	resp := &resource.UpdateResponse{}

	result := updateAzureDataServicesIfChanged(ctx, plan, state, client, resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Contains(t, result.Diagnostics.Errors()[0].Detail(), "shared_managed_identity must be set")
	mockClient.AssertNotCalled(t, "UpdateDataServiceResourcesContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureDataServicesIfChanged_AksPrivateDnsZoneIdUnset_ReturnsError(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName: types.StringValue(testEnvName),
		DataServices: &DataServices{
			SharedManagedIdentity: types.StringValue(testNewSharedManagedIdentity),
			AksPrivateDnsZoneId:   types.StringNull(),
		},
	}
	state := &azureEnvironmentResourceModel{
		DataServices: &DataServices{
			SharedManagedIdentity: types.StringValue(testOldSharedManagedIdentity),
			AksPrivateDnsZoneId:   types.StringValue(testOldAksPrivateDnsZoneId),
		},
	}
	resp := &resource.UpdateResponse{}

	result := updateAzureDataServicesIfChanged(ctx, plan, state, client, resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Contains(t, result.Diagnostics.Errors()[0].Detail(), "aks_private_dns_zone_id cannot be unset")
	mockClient.AssertNotCalled(t, "UpdateDataServiceResourcesContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureDataServicesIfChanged_AksPrivateDnsZoneIdNullWithStateNil_CallsAPI(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName: types.StringValue(testEnvName),
		DataServices: &DataServices{
			SharedManagedIdentity: types.StringValue(testNewSharedManagedIdentity),
			AksPrivateDnsZoneId:   types.StringNull(),
		},
	}
	state := &azureEnvironmentResourceModel{
		DataServices: nil,
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateDataServiceResourcesContext", mock.Anything, mock.Anything, mock.Anything).
		Return(&operations.UpdateDataServiceResourcesOK{}, nil)

	result := updateAzureDataServicesIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertExpectations(t)
}

// Tests for updateAzureNetworkParamsIfChanged

func buildExistingNetworkParamsObject(ctx context.Context, network existingAzureNetwork) types.Object {
	obj, _ := types.ObjectValueFrom(ctx, map[string]attr.Type{
		"aks_private_dns_zone_id":      types.StringType,
		"database_private_dns_zone_id": types.StringType,
		"network_id":                   types.StringType,
		"resource_group_name":          types.StringType,
		"subnet_ids":                   types.SetType{ElemType: types.StringType},
		"flexible_server_subnet_ids":   types.SetType{ElemType: types.StringType},
	}, &network)
	return obj
}

func TestUpdateAzureNetworkParamsIfChanged_SubnetChanged_CallsSubnetAPI(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	planNetwork := existingAzureNetwork{
		DatabasePrivateDNSZoneID: types.StringValue("dns-zone-id"),
		NetworkID:                types.StringValue("net-id"),
		ResourceGroupName:        types.StringValue("rg"),
		SubnetIds:                utils.ToSetValueFromStringList([]string{"subnet-a", "subnet-b"}),
		FlexibleServerSubnetIds:  utils.ToSetValueFromStringList([]string{"flex-subnet"}),
	}
	stateNetwork := existingAzureNetwork{
		DatabasePrivateDNSZoneID: types.StringValue("dns-zone-id"),
		NetworkID:                types.StringValue("net-id"),
		ResourceGroupName:        types.StringValue("rg"),
		SubnetIds:                utils.ToSetValueFromStringList([]string{"subnet-a"}),
		FlexibleServerSubnetIds:  utils.ToSetValueFromStringList([]string{"flex-subnet"}),
	}

	plan := &azureEnvironmentResourceModel{
		EnvironmentName:         types.StringValue(testEnvName),
		ExistingNetworkParams:   buildExistingNetworkParamsObject(ctx, planNetwork),
		FlexibleServerSubnetIds: utils.ToSetValueFromStringList([]string{"flex-subnet"}),
	}
	state := &azureEnvironmentResourceModel{
		ExistingNetworkParams:   buildExistingNetworkParamsObject(ctx, stateNetwork),
		FlexibleServerSubnetIds: utils.ToSetValueFromStringList([]string{"flex-subnet"}),
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateSubnetContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateSubnetParams) bool {
		return params.Input != nil && *params.Input.Environment == testEnvName && len(params.Input.SubnetIds) == 2
	}), mock.Anything).Return(&operations.UpdateSubnetOK{}, nil)

	result := updateAzureNetworkParamsIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, plan.ExistingNetworkParams, state.ExistingNetworkParams)
	mockClient.AssertExpectations(t)
	mockClient.AssertNotCalled(t, "UpdateAzureDatabaseResourcesContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureNetworkParamsIfChanged_DatabaseResourcesChanged_CallsDatabaseAPI(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	planNetwork := existingAzureNetwork{
		DatabasePrivateDNSZoneID: types.StringValue("new-dns-zone-id"),
		NetworkID:                types.StringValue("net-id"),
		ResourceGroupName:        types.StringValue("rg"),
		SubnetIds:                utils.ToSetValueFromStringList([]string{"subnet-a"}),
		FlexibleServerSubnetIds:  utils.ToSetValueFromStringList([]string{"flex-subnet-new"}),
	}
	stateNetwork := existingAzureNetwork{
		DatabasePrivateDNSZoneID: types.StringValue("old-dns-zone-id"),
		NetworkID:                types.StringValue("net-id"),
		ResourceGroupName:        types.StringValue("rg"),
		SubnetIds:                utils.ToSetValueFromStringList([]string{"subnet-a"}),
		FlexibleServerSubnetIds:  utils.ToSetValueFromStringList([]string{"flex-subnet-old"}),
	}

	plan := &azureEnvironmentResourceModel{
		EnvironmentName:         types.StringValue(testEnvName),
		ExistingNetworkParams:   buildExistingNetworkParamsObject(ctx, planNetwork),
		FlexibleServerSubnetIds: utils.ToSetValueFromStringList([]string{"flex-subnet-new"}),
	}
	state := &azureEnvironmentResourceModel{
		ExistingNetworkParams:   buildExistingNetworkParamsObject(ctx, stateNetwork),
		FlexibleServerSubnetIds: utils.ToSetValueFromStringList([]string{"flex-subnet-old"}),
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateAzureDatabaseResourcesContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateAzureDatabaseResourcesParams) bool {
		return params.Input != nil &&
			*params.Input.Environment == testEnvName &&
			params.Input.DatabasePrivateDNSZoneID == "new-dns-zone-id" &&
			len(params.Input.FlexibleServerSubnetIds) == 1 &&
			params.Input.FlexibleServerSubnetIds[0] == "flex-subnet-new"
	}), mock.Anything).Return(&operations.UpdateAzureDatabaseResourcesOK{}, nil)

	result := updateAzureNetworkParamsIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, plan.FlexibleServerSubnetIds, state.FlexibleServerSubnetIds)
	assert.Equal(t, plan.ExistingNetworkParams, state.ExistingNetworkParams)
	mockClient.AssertExpectations(t)
	mockClient.AssertNotCalled(t, "UpdateSubnetContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureNetworkParamsIfChanged_SubnetAndDatabaseBothChanged_CallsBothAPIs(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	planNetwork := existingAzureNetwork{
		DatabasePrivateDNSZoneID: types.StringValue("new-dns-zone-id"),
		NetworkID:                types.StringValue("net-id"),
		ResourceGroupName:        types.StringValue("rg"),
		SubnetIds:                utils.ToSetValueFromStringList([]string{"subnet-a", "subnet-b"}),
		FlexibleServerSubnetIds:  utils.ToSetValueFromStringList([]string{"flex-subnet-new"}),
	}
	stateNetwork := existingAzureNetwork{
		DatabasePrivateDNSZoneID: types.StringValue("old-dns-zone-id"),
		NetworkID:                types.StringValue("net-id"),
		ResourceGroupName:        types.StringValue("rg"),
		SubnetIds:                utils.ToSetValueFromStringList([]string{"subnet-a"}),
		FlexibleServerSubnetIds:  utils.ToSetValueFromStringList([]string{"flex-subnet-old"}),
	}

	plan := &azureEnvironmentResourceModel{
		EnvironmentName:         types.StringValue(testEnvName),
		ExistingNetworkParams:   buildExistingNetworkParamsObject(ctx, planNetwork),
		FlexibleServerSubnetIds: utils.ToSetValueFromStringList([]string{"flex-subnet-new"}),
	}
	state := &azureEnvironmentResourceModel{
		ExistingNetworkParams:   buildExistingNetworkParamsObject(ctx, stateNetwork),
		FlexibleServerSubnetIds: utils.ToSetValueFromStringList([]string{"flex-subnet-old"}),
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateSubnetContext", mock.Anything, mock.Anything, mock.Anything).
		Return(&operations.UpdateSubnetOK{}, nil)
	mockClient.On("UpdateAzureDatabaseResourcesContext", mock.Anything, mock.Anything, mock.Anything).
		Return(&operations.UpdateAzureDatabaseResourcesOK{}, nil)

	result := updateAzureNetworkParamsIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, plan.ExistingNetworkParams, state.ExistingNetworkParams)
	assert.Equal(t, plan.FlexibleServerSubnetIds, state.FlexibleServerSubnetIds)
	mockClient.AssertExpectations(t)
}

func TestUpdateAzureNetworkParamsIfChanged_Unchanged_SkipsAllAPICalls(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	network := existingAzureNetwork{
		DatabasePrivateDNSZoneID: types.StringValue("dns-zone-id"),
		NetworkID:                types.StringValue("net-id"),
		ResourceGroupName:        types.StringValue("rg"),
		SubnetIds:                utils.ToSetValueFromStringList([]string{"subnet-a"}),
		FlexibleServerSubnetIds:  utils.ToSetValueFromStringList([]string{"flex-subnet"}),
	}
	networkObj := buildExistingNetworkParamsObject(ctx, network)
	flexSubnets := utils.ToSetValueFromStringList([]string{"flex-subnet"})

	plan := &azureEnvironmentResourceModel{
		EnvironmentName:         types.StringValue(testEnvName),
		ExistingNetworkParams:   networkObj,
		FlexibleServerSubnetIds: flexSubnets,
	}
	state := &azureEnvironmentResourceModel{
		ExistingNetworkParams:   networkObj,
		FlexibleServerSubnetIds: flexSubnets,
	}
	resp := &resource.UpdateResponse{}

	result := updateAzureNetworkParamsIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateSubnetContext", mock.Anything, mock.Anything, mock.Anything)
	mockClient.AssertNotCalled(t, "UpdateAzureDatabaseResourcesContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureNetworkParamsIfChanged_DnsZoneRemoval_ReturnsError(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	planNetwork := existingAzureNetwork{
		DatabasePrivateDNSZoneID: types.StringNull(),
		NetworkID:                types.StringValue("net-id"),
		ResourceGroupName:        types.StringValue("rg"),
		SubnetIds:                utils.ToSetValueFromStringList([]string{"subnet-a"}),
		FlexibleServerSubnetIds:  utils.ToSetValueFromStringList([]string{"flex-subnet"}),
	}
	stateNetwork := existingAzureNetwork{
		DatabasePrivateDNSZoneID: types.StringValue("old-dns-zone-id"),
		NetworkID:                types.StringValue("net-id"),
		ResourceGroupName:        types.StringValue("rg"),
		SubnetIds:                utils.ToSetValueFromStringList([]string{"subnet-a"}),
		FlexibleServerSubnetIds:  utils.ToSetValueFromStringList([]string{"flex-subnet"}),
	}

	plan := &azureEnvironmentResourceModel{
		EnvironmentName:         types.StringValue(testEnvName),
		ExistingNetworkParams:   buildExistingNetworkParamsObject(ctx, planNetwork),
		FlexibleServerSubnetIds: utils.ToSetValueFromStringList([]string{"flex-subnet"}),
	}
	state := &azureEnvironmentResourceModel{
		ExistingNetworkParams:   buildExistingNetworkParamsObject(ctx, stateNetwork),
		FlexibleServerSubnetIds: utils.ToSetValueFromStringList([]string{"flex-subnet"}),
	}
	resp := &resource.UpdateResponse{}

	result := updateAzureNetworkParamsIfChanged(ctx, plan, state, client, resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Contains(t, result.Diagnostics.Errors()[0].Detail(), "database_private_dns_zone_id cannot be unset")
	mockClient.AssertNotCalled(t, "UpdateAzureDatabaseResourcesContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureNetworkParamsIfChanged_DatabaseAPIError_PreservesState(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	planNetwork := existingAzureNetwork{
		DatabasePrivateDNSZoneID: types.StringValue("new-dns-zone-id"),
		NetworkID:                types.StringValue("net-id"),
		ResourceGroupName:        types.StringValue("rg"),
		SubnetIds:                utils.ToSetValueFromStringList([]string{"subnet-a"}),
		FlexibleServerSubnetIds:  utils.ToSetValueFromStringList([]string{"flex-subnet-new"}),
	}
	stateNetwork := existingAzureNetwork{
		DatabasePrivateDNSZoneID: types.StringValue("old-dns-zone-id"),
		NetworkID:                types.StringValue("net-id"),
		ResourceGroupName:        types.StringValue("rg"),
		SubnetIds:                utils.ToSetValueFromStringList([]string{"subnet-a"}),
		FlexibleServerSubnetIds:  utils.ToSetValueFromStringList([]string{"flex-subnet-old"}),
	}
	origStateNetworkObj := buildExistingNetworkParamsObject(ctx, stateNetwork)
	origFlexSubnets := utils.ToSetValueFromStringList([]string{"flex-subnet-old"})

	plan := &azureEnvironmentResourceModel{
		EnvironmentName:         types.StringValue(testEnvName),
		ExistingNetworkParams:   buildExistingNetworkParamsObject(ctx, planNetwork),
		FlexibleServerSubnetIds: utils.ToSetValueFromStringList([]string{"flex-subnet-new"}),
	}
	state := &azureEnvironmentResourceModel{
		ExistingNetworkParams:   origStateNetworkObj,
		FlexibleServerSubnetIds: origFlexSubnets,
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateAzureDatabaseResourcesContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpdateAzureDatabaseResourcesOK)(nil), errors.New(testServiceUnavailable))

	result := updateAzureNetworkParamsIfChanged(ctx, plan, state, client, resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Equal(t, origFlexSubnets, state.FlexibleServerSubnetIds)
	assert.Equal(t, origStateNetworkObj, state.ExistingNetworkParams)
}

func TestUpdateAzureNetworkParamsIfChanged_TopLevelFlexNull_FallsBackToNested(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	planNetwork := existingAzureNetwork{
		DatabasePrivateDNSZoneID: types.StringValue("new-dns-zone-id"),
		NetworkID:                types.StringValue("net-id"),
		ResourceGroupName:        types.StringValue("rg"),
		SubnetIds:                utils.ToSetValueFromStringList([]string{"subnet-a"}),
		FlexibleServerSubnetIds:  utils.ToSetValueFromStringList([]string{"flex-subnet-nested"}),
	}
	stateNetwork := existingAzureNetwork{
		DatabasePrivateDNSZoneID: types.StringValue("old-dns-zone-id"),
		NetworkID:                types.StringValue("net-id"),
		ResourceGroupName:        types.StringValue("rg"),
		SubnetIds:                utils.ToSetValueFromStringList([]string{"subnet-a"}),
		FlexibleServerSubnetIds:  utils.ToSetValueFromStringList([]string{"flex-subnet-nested"}),
	}

	plan := &azureEnvironmentResourceModel{
		EnvironmentName:         types.StringValue(testEnvName),
		ExistingNetworkParams:   buildExistingNetworkParamsObject(ctx, planNetwork),
		FlexibleServerSubnetIds: types.SetNull(types.StringType),
	}
	state := &azureEnvironmentResourceModel{
		ExistingNetworkParams:   buildExistingNetworkParamsObject(ctx, stateNetwork),
		FlexibleServerSubnetIds: types.SetNull(types.StringType),
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateAzureDatabaseResourcesContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateAzureDatabaseResourcesParams) bool {
		return params.Input != nil &&
			*params.Input.Environment == testEnvName &&
			params.Input.DatabasePrivateDNSZoneID == "new-dns-zone-id" &&
			len(params.Input.FlexibleServerSubnetIds) == 1 &&
			params.Input.FlexibleServerSubnetIds[0] == "flex-subnet-nested"
	}), mock.Anything).Return(&operations.UpdateAzureDatabaseResourcesOK{}, nil)

	result := updateAzureNetworkParamsIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertExpectations(t)
}

func TestUpdateAzureNetworkParamsIfChanged_SubnetAPIError_StopsBeforeDatabaseCall(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	planNetwork := existingAzureNetwork{
		DatabasePrivateDNSZoneID: types.StringValue("new-dns-zone-id"),
		NetworkID:                types.StringValue("net-id"),
		ResourceGroupName:        types.StringValue("rg"),
		SubnetIds:                utils.ToSetValueFromStringList([]string{"subnet-a", "subnet-b"}),
		FlexibleServerSubnetIds:  utils.ToSetValueFromStringList([]string{"flex-subnet-new"}),
	}
	stateNetwork := existingAzureNetwork{
		DatabasePrivateDNSZoneID: types.StringValue("old-dns-zone-id"),
		NetworkID:                types.StringValue("net-id"),
		ResourceGroupName:        types.StringValue("rg"),
		SubnetIds:                utils.ToSetValueFromStringList([]string{"subnet-a"}),
		FlexibleServerSubnetIds:  utils.ToSetValueFromStringList([]string{"flex-subnet-old"}),
	}

	plan := &azureEnvironmentResourceModel{
		EnvironmentName:         types.StringValue(testEnvName),
		ExistingNetworkParams:   buildExistingNetworkParamsObject(ctx, planNetwork),
		FlexibleServerSubnetIds: utils.ToSetValueFromStringList([]string{"flex-subnet-new"}),
	}
	state := &azureEnvironmentResourceModel{
		ExistingNetworkParams:   buildExistingNetworkParamsObject(ctx, stateNetwork),
		FlexibleServerSubnetIds: utils.ToSetValueFromStringList([]string{"flex-subnet-old"}),
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateSubnetContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpdateSubnetOK)(nil), errors.New(testServiceUnavailable))

	result := updateAzureNetworkParamsIfChanged(ctx, plan, state, client, resp)

	assert.True(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateAzureDatabaseResourcesContext", mock.Anything, mock.Anything, mock.Anything)
}

// Tests for updateAzureCredentialIfChanged

func TestUpdateAzureCredentialIfChanged_Changed_CallsAPI(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName: types.StringValue(testEnvName),
		CredentialName:  types.StringValue(testNewCredentialName),
	}
	state := &azureEnvironmentResourceModel{
		CredentialName: types.StringValue(testOldCredentialName),
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("ChangeEnvironmentCredentialContext", mock.Anything, mock.MatchedBy(func(params *operations.ChangeEnvironmentCredentialParams) bool {
		return params.Input != nil &&
			*params.Input.CredentialName == testNewCredentialName &&
			*params.Input.EnvironmentName == testEnvName
	}), mock.Anything).Return(&operations.ChangeEnvironmentCredentialOK{}, nil)

	result := updateAzureCredentialIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testNewCredentialName), state.CredentialName)
	mockClient.AssertExpectations(t)
}

func TestUpdateAzureCredentialIfChanged_Unchanged_SkipsAPICall(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName: types.StringValue(testEnvName),
		CredentialName:  types.StringValue(testSameCredentialName),
	}
	state := &azureEnvironmentResourceModel{
		CredentialName: types.StringValue(testSameCredentialName),
	}
	resp := &resource.UpdateResponse{}

	result := updateAzureCredentialIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "ChangeEnvironmentCredentialContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureCredentialIfChanged_APIError_AddsDiagnostic(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName: types.StringValue(testEnvName),
		CredentialName:  types.StringValue(testNewCredentialName),
	}
	state := &azureEnvironmentResourceModel{
		CredentialName: types.StringValue(testOldCredentialName),
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("ChangeEnvironmentCredentialContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.ChangeEnvironmentCredentialOK)(nil), errors.New(testServiceUnavailable))

	result := updateAzureCredentialIfChanged(ctx, plan, state, client, resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testOldCredentialName), state.CredentialName)
}

// Tests for updateAzureSshKeyIfChanged

func TestUpdateAzureSshKeyIfChanged_Changed_CallsAPI(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName: types.StringValue(testEnvName),
		PublicKey:       types.StringValue(testNewKey),
	}
	state := &azureEnvironmentResourceModel{
		PublicKey: types.StringValue(testOldKey),
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateSSHKeyContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateSSHKeyParams) bool {
		return params.Input != nil &&
			params.Input.NewPublicKey == testNewKey &&
			*params.Input.Environment == testEnvName
	}), mock.Anything).Return(&operations.UpdateSSHKeyOK{}, nil)

	result := updateAzureSshKeyIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testNewKey), state.PublicKey)
	mockClient.AssertExpectations(t)
}

func TestUpdateAzureSshKeyIfChanged_Unchanged_SkipsAPICall(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName: types.StringValue(testEnvName),
		PublicKey:       types.StringValue(testSameKey),
	}
	state := &azureEnvironmentResourceModel{
		PublicKey: types.StringValue(testSameKey),
	}
	resp := &resource.UpdateResponse{}

	result := updateAzureSshKeyIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureSshKeyIfChanged_APIError_AddsDiagnostic(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName: types.StringValue(testEnvName),
		PublicKey:       types.StringValue(testNewKey),
	}
	state := &azureEnvironmentResourceModel{
		PublicKey: types.StringValue(testOldKey),
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpdateSSHKeyOK)(nil), errors.New(testServiceUnavailable))

	result := updateAzureSshKeyIfChanged(ctx, plan, state, client, resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testOldKey), state.PublicKey)
}

// Tests for updateAzureProxyConfigurationIfChanged

func TestUpdateAzureProxyConfigurationIfChanged_Changed_CallsAPI(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName: types.StringValue(testEnvName),
		ProxyConfigName: types.StringValue(testNewProxyConfigName),
	}
	state := &azureEnvironmentResourceModel{
		ProxyConfigName: types.StringValue(testOldProxyConfigName),
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateProxyConfigContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateProxyConfigParams) bool {
		return params.Input != nil &&
			*params.Input.Environment == testEnvName &&
			params.Input.ProxyConfigName == testNewProxyConfigName &&
			!params.Input.RemoveProxy
	}), mock.Anything).Return(&operations.UpdateProxyConfigOK{}, nil)

	result := updateAzureProxyConfigurationIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testNewProxyConfigName), state.ProxyConfigName)
	mockClient.AssertExpectations(t)
}

func TestUpdateAzureProxyConfigurationIfChanged_Unchanged_SkipsAPICall(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName: types.StringValue(testEnvName),
		ProxyConfigName: types.StringValue(testOldProxyConfigName),
	}
	state := &azureEnvironmentResourceModel{
		ProxyConfigName: types.StringValue(testOldProxyConfigName),
	}
	resp := &resource.UpdateResponse{}

	result := updateAzureProxyConfigurationIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateProxyConfigContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureProxyConfigurationIfChanged_APIError_AddsDiagnostic(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName: types.StringValue(testEnvName),
		ProxyConfigName: types.StringValue(testNewProxyConfigName),
	}
	state := &azureEnvironmentResourceModel{
		ProxyConfigName: types.StringValue(testOldProxyConfigName),
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateProxyConfigContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpdateProxyConfigOK)(nil), errors.New(testServiceUnavailable))

	result := updateAzureProxyConfigurationIfChanged(ctx, plan, state, client, resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testOldProxyConfigName), state.ProxyConfigName)
}

// Tests for updateAzureCustomDockerRegistryIfChanged

func TestUpdateAzureCustomDockerRegistryIfChanged_Changed_CallsAPI(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName:      types.StringValue(testEnvName),
		CustomDockerRegistry: &CustomDockerRegistry{Crn: types.StringValue(testNewDockerRegistryCrn)},
	}
	state := &azureEnvironmentResourceModel{
		CustomDockerRegistry: &CustomDockerRegistry{Crn: types.StringValue(testOldDockerRegistryCrn)},
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateCustomDockerRegistryContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateCustomDockerRegistryParams) bool {
		return params.Input != nil &&
			*params.Input.CustomDockerRegistry == testNewDockerRegistryCrn &&
			*params.Input.Environment == testEnvName
	}), mock.Anything).Return(&operations.UpdateCustomDockerRegistryOK{}, nil)

	result := updateAzureCustomDockerRegistryIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testNewDockerRegistryCrn), state.CustomDockerRegistry.Crn)
	mockClient.AssertExpectations(t)
}

func TestUpdateAzureCustomDockerRegistryIfChanged_Unchanged_SkipsAPICall(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName:      types.StringValue(testEnvName),
		CustomDockerRegistry: &CustomDockerRegistry{Crn: types.StringValue(testOldDockerRegistryCrn)},
	}
	state := &azureEnvironmentResourceModel{
		CustomDockerRegistry: &CustomDockerRegistry{Crn: types.StringValue(testOldDockerRegistryCrn)},
	}
	resp := &resource.UpdateResponse{}

	result := updateAzureCustomDockerRegistryIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateCustomDockerRegistryContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureCustomDockerRegistryIfChanged_APIError_AddsDiagnostic(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName:      types.StringValue(testEnvName),
		CustomDockerRegistry: &CustomDockerRegistry{Crn: types.StringValue(testNewDockerRegistryCrn)},
	}
	state := &azureEnvironmentResourceModel{
		CustomDockerRegistry: &CustomDockerRegistry{Crn: types.StringValue(testOldDockerRegistryCrn)},
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateCustomDockerRegistryContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpdateCustomDockerRegistryOK)(nil), errors.New(testServiceUnavailable))

	result := updateAzureCustomDockerRegistryIfChanged(ctx, plan, state, client, resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testOldDockerRegistryCrn), state.CustomDockerRegistry.Crn)
}

// Tests for updateAzureEndpointAccessGatewayIfChanged

func TestUpdateAzureEndpointAccessGatewayIfChanged_SchemeChanged_CallsAPI(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	planSubnets := utils.ToSetValueFromStringList([]string{testAzureSubnet})
	plan := &azureEnvironmentResourceModel{
		EnvironmentName:                types.StringValue(testEnvName),
		EndpointAccessGatewayScheme:    types.StringValue(testGatewaySchemePrivate),
		EndpointAccessGatewaySubnetIds: planSubnets,
	}
	state := &azureEnvironmentResourceModel{
		EndpointAccessGatewayScheme:    types.StringValue(testGatewaySchemePublic),
		EndpointAccessGatewaySubnetIds: planSubnets,
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("SetEndpointAccessGatewayContext", mock.Anything, mock.MatchedBy(func(params *operations.SetEndpointAccessGatewayParams) bool {
		return params.Input != nil &&
			*params.Input.EndpointAccessGatewayScheme == testGatewaySchemePrivate &&
			*params.Input.Environment == testEnvName
	}), mock.Anything).Return(&operations.SetEndpointAccessGatewayOK{
		Payload: &environmentsmodels.SetEndpointAccessGatewayResponse{},
	}, nil)

	result := updateAzureEndpointAccessGatewayIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testGatewaySchemePrivate), state.EndpointAccessGatewayScheme)
	mockClient.AssertExpectations(t)
}

func TestUpdateAzureEndpointAccessGatewayIfChanged_Unchanged_SkipsAPICall(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	subnets := utils.ToSetValueFromStringList([]string{testAzureSubnet})
	plan := &azureEnvironmentResourceModel{
		EnvironmentName:                types.StringValue(testEnvName),
		EndpointAccessGatewayScheme:    types.StringValue(testGatewaySchemePublic),
		EndpointAccessGatewaySubnetIds: subnets,
	}
	state := &azureEnvironmentResourceModel{
		EndpointAccessGatewayScheme:    types.StringValue(testGatewaySchemePublic),
		EndpointAccessGatewaySubnetIds: subnets,
	}
	resp := &resource.UpdateResponse{}

	result := updateAzureEndpointAccessGatewayIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "SetEndpointAccessGatewayContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureEndpointAccessGatewayIfChanged_APIError_AddsDiagnostic(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	planSubnets := utils.ToSetValueFromStringList([]string{testAzureSubnet})
	plan := &azureEnvironmentResourceModel{
		EnvironmentName:                types.StringValue(testEnvName),
		EndpointAccessGatewayScheme:    types.StringValue(testGatewaySchemePrivate),
		EndpointAccessGatewaySubnetIds: planSubnets,
	}
	state := &azureEnvironmentResourceModel{
		EndpointAccessGatewayScheme:    types.StringValue(testGatewaySchemePublic),
		EndpointAccessGatewaySubnetIds: planSubnets,
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("SetEndpointAccessGatewayContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.SetEndpointAccessGatewayOK)(nil), errors.New(testServiceUnavailable))

	result := updateAzureEndpointAccessGatewayIfChanged(ctx, plan, state, client, resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testGatewaySchemePublic), state.EndpointAccessGatewayScheme)
}

// Tests for updateAzureAvailabilityZonesIfChanged

func TestUpdateAzureAvailabilityZonesIfChanged_Changed_DelegatesToUpdateAvailabilityZones(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName:   types.StringValue(testEnvName),
		AvailabilityZones: utils.ToSetValueFromStringList([]string{"1", "2", "3"}),
	}
	state := &azureEnvironmentResourceModel{
		AvailabilityZones: utils.ToSetValueFromStringList([]string{"1", "2"}),
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateAzureAvailabilityZonesContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateAzureAvailabilityZonesParams) bool {
		return params.Input != nil &&
			*params.Input.Environment == testEnvName &&
			len(params.Input.AvailabilityZones) == 3
	}), mock.Anything).Return(&operations.UpdateAzureAvailabilityZonesOK{}, nil)

	result := updateAzureAvailabilityZonesIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, plan.AvailabilityZones, state.AvailabilityZones)
	mockClient.AssertExpectations(t)
}

func TestUpdateAzureAvailabilityZonesIfChanged_Unchanged_SkipsAPICall(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	zones := utils.ToSetValueFromStringList([]string{"1", "2"})
	plan := &azureEnvironmentResourceModel{
		EnvironmentName:   types.StringValue(testEnvName),
		AvailabilityZones: zones,
	}
	state := &azureEnvironmentResourceModel{
		AvailabilityZones: utils.ToSetValueFromStringList([]string{"1", "2"}),
	}
	resp := &resource.UpdateResponse{}

	result := updateAzureAvailabilityZonesIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateAzureAvailabilityZonesContext", mock.Anything, mock.Anything, mock.Anything)
}

// Tests for updateAzureCatalogIfChanged

func TestUpdateAzureCatalogIfChanged_Changed_CallsAPI(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName: types.StringValue(testEnvName),
		FreeIpa:         newFreeIpaObject(testNewCatalogURL),
	}
	state := &azureEnvironmentResourceModel{
		FreeIpa: newFreeIpaObject(testOldCatalogURL),
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("SetCatalogContext", mock.Anything, mock.MatchedBy(func(params *operations.SetCatalogParams) bool {
		return params.Input != nil &&
			*params.Input.Catalog == testNewCatalogURL &&
			*params.Input.Environment == testEnvName
	}), mock.Anything).Return(&operations.SetCatalogOK{}, nil)

	result := updateAzureCatalogIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertExpectations(t)
}

func TestUpdateAzureCatalogIfChanged_Unchanged_SkipsAPICall(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName: types.StringValue(testEnvName),
		FreeIpa:         newFreeIpaObject(testSameCatalogURL),
	}
	state := &azureEnvironmentResourceModel{
		FreeIpa: newFreeIpaObject(testSameCatalogURL),
	}
	resp := &resource.UpdateResponse{}

	result := updateAzureCatalogIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "SetCatalogContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureCatalogIfChanged_APIError_AddsDiagnostic(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &azureEnvironmentResourceModel{
		EnvironmentName: types.StringValue(testEnvName),
		FreeIpa:         newFreeIpaObject(testNewCatalogURL),
	}
	state := &azureEnvironmentResourceModel{
		FreeIpa: newFreeIpaObject(testOldCatalogURL),
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("SetCatalogContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.SetCatalogOK)(nil), errors.New(testServiceUnavailable))

	result := updateAzureCatalogIfChanged(ctx, plan, state, client, resp)

	assert.True(t, result.Diagnostics.HasError())
}
