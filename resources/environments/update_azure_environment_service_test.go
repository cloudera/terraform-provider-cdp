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
