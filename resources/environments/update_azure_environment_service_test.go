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
	testEncryptionKeyResourceGroupName = "my-encryption-rg"
	testEncryptionUserManagedIdentity  = "/subscriptions/sub-id/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/my-identity"
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

	result := updateAzureAvailabilityZonesIfChanged(ctx, client, plan, &state, new(testEnvName), resp)

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

	result := updateAzureAvailabilityZonesIfChanged(ctx, client, zones, new(utils.ToSetValueFromStringList([]string{"1", "2"})), new(testEnvName), resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateAzureAvailabilityZonesContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureAvailabilityZonesIfChanged_PlanNull_SkipsAPICall(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	plan := types.SetNull(types.StringType)
	resp := &resource.UpdateResponse{}

	result := updateAzureAvailabilityZonesIfChanged(ctx, client, plan, new(utils.ToSetValueFromStringList([]string{"1", "2"})), new(testEnvName), resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateAzureAvailabilityZonesContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureAvailabilityZonesIfChanged_PlanUnknown_SkipsAPICall(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	plan := types.SetUnknown(types.StringType)
	resp := &resource.UpdateResponse{}

	result := updateAzureAvailabilityZonesIfChanged(ctx, client, plan, new(utils.ToSetValueFromStringList([]string{"1", "2"})), new(testEnvName), resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateAzureAvailabilityZonesContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAzureAvailabilityZonesIfChanged_PlanEmpty_AddsValidationError(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	plan := utils.ToSetValueFromStringList([]string{})
	resp := &resource.UpdateResponse{}

	result := updateAzureAvailabilityZonesIfChanged(ctx, client, plan, new(utils.ToSetValueFromStringList([]string{"1", "2"})), new(testEnvName), resp)

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

	result := updateAzureAvailabilityZonesIfChanged(ctx, client, plan, new(utils.ToSetValueFromStringList([]string{"1", "2"})), new(testEnvName), resp)

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

	updateAzureAvailabilityZonesIfChanged(ctx, client, plan, &state, new(testEnvName), resp)

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

	envName := testEnvName
	err := updateAzureEncryptionResources(ctx, client, &envName, plan)

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

	envName := testEnvName
	err := updateAzureEncryptionResources(ctx, client, &envName, plan)

	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	if err.Error() != testServiceUnavailable {
		t.Errorf("expected error message '%s', got: %s", testServiceUnavailable, err.Error())
	}
}
