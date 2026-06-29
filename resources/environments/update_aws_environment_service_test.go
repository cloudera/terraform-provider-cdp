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
	testIpRange1            = "192.168.1.1/32"
	testIpRange2            = "10.0.0.0/24"
	testSubnet1             = "subnet-12345"
	testSubnet2             = "subnet-67890"
	testFallbackSubnet1     = "subnet-abcde"
	testFallbackSubnet2     = "subnet-fghij"
	testEncryptionKeyArn    = "arn:aws:kms:us-west-2:123456789012:key/test-key-id"
	testNewEncryptionKeyArn = "arn:aws:kms:us-west-2:123456789012:key/new-key-id"
	testPublicKeyID         = "my-key-id"
)

func TestConvertNilConfigReturnsNilRequestForAws(t *testing.T) {
	config := (*AwsComputeClusterConfiguration)(nil)
	fallbackSubnetIds := types.Set{}
	want := (*environmentsmodels.AWSComputeClusterConfigurationRequest)(nil)
	if got := convertConfigToAwsComputeClusterConfigurationRequest(config, fallbackSubnetIds); !reflect.DeepEqual(got, want) {
		t.Errorf("convertConfigToAwsComputeClusterConfigurationRequest() = %v, want %v", got, want)
	}
}

func TestConvertValidConfigReturnsCorrectRequestForAws(t *testing.T) {
	config := &AwsComputeClusterConfiguration{
		KubeApiAuthorizedIpRanges: utils.ToSetValueFromStringList([]string{testIpRange1, testIpRange2}),
		PrivateCluster:            types.BoolValue(true),
		WorkerNodeSubnets:         utils.ToSetValueFromStringList([]string{testSubnet1, testSubnet2}),
	}
	fallbackSubnetIds := types.Set{}
	want := &environmentsmodels.AWSComputeClusterConfigurationRequest{
		KubeAPIAuthorizedIPRanges: []string{testIpRange1, testIpRange2},
		PrivateCluster:            true,
		WorkerNodeSubnets:         []string{testSubnet1, testSubnet2},
	}
	if got := convertConfigToAwsComputeClusterConfigurationRequest(config, fallbackSubnetIds); !reflect.DeepEqual(got, want) {
		t.Errorf("convertConfigToAwsComputeClusterConfigurationRequest() = %v, want %v", got, want)
	}
}

func TestConvertConfigWithEmptyFieldsReturnsCorrectRequestForAws(t *testing.T) {
	config := &AwsComputeClusterConfiguration{
		KubeApiAuthorizedIpRanges: utils.ToSetValueFromStringList([]string{}),
		PrivateCluster:            types.BoolValue(false),
		WorkerNodeSubnets:         utils.ToSetValueFromStringList([]string{}),
	}
	fallbackSubnetIds := types.Set{}
	want := &environmentsmodels.AWSComputeClusterConfigurationRequest{
		KubeAPIAuthorizedIPRanges: []string{},
		PrivateCluster:            false,
		WorkerNodeSubnets:         []string{},
	}
	if got := convertConfigToAwsComputeClusterConfigurationRequest(config, fallbackSubnetIds); !reflect.DeepEqual(got, want) {
		t.Errorf("convertConfigToAwsComputeClusterConfigurationRequest() = %v, want %v", got, want)
	}
}

func TestConvertConfigWithFallbackSubnetsReturnsCorrectRequestForAws(t *testing.T) {
	config := &AwsComputeClusterConfiguration{
		KubeApiAuthorizedIpRanges: utils.ToSetValueFromStringList([]string{testIpRange1}),
		PrivateCluster:            types.BoolValue(true),
		WorkerNodeSubnets:         types.Set{},
	}
	fallbackSubnetIds := utils.ToSetValueFromStringList([]string{testFallbackSubnet1, testFallbackSubnet2})
	want := &environmentsmodels.AWSComputeClusterConfigurationRequest{
		KubeAPIAuthorizedIPRanges: []string{testIpRange1},
		PrivateCluster:            true,
		WorkerNodeSubnets:         []string{testFallbackSubnet1, testFallbackSubnet2},
	}
	if got := convertConfigToAwsComputeClusterConfigurationRequest(config, fallbackSubnetIds); !reflect.DeepEqual(got, want) {
		t.Errorf("convertConfigToAwsComputeClusterConfigurationRequest() = %v, want %v", got, want)
	}
}

func TestUpdateDiskEncryption_Success(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	keyArn := types.StringValue(testEncryptionKeyArn)

	mockClient.On("UpdateAwsDiskEncryptionParametersContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateAwsDiskEncryptionParametersParams) bool {
		return params.Input != nil &&
			*params.Input.EncryptionKeyArn == testEncryptionKeyArn &&
			*params.Input.Environment == testEnvName
	}), mock.Anything).
		Return(&operations.UpdateAwsDiskEncryptionParametersOK{}, nil)

	err := updateDiskEncryption(ctx, client, new(testEnvName), keyArn)

	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	mockClient.AssertExpectations(t)
}

func TestUpdateDiskEncryption_ReturnsError(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	keyArn := types.StringValue(testEncryptionKeyArn)

	mockClient.On("UpdateAwsDiskEncryptionParametersContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpdateAwsDiskEncryptionParametersOK)(nil), errors.New(testServiceUnavailable))

	err := updateDiskEncryption(ctx, client, new(testEnvName), keyArn)

	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	if err.Error() != testServiceUnavailable {
		t.Errorf("expected error message '%s', got: %s", testServiceUnavailable, err.Error())
	}
}

// Tests for updateSshKeyForAws

func TestUpdateSshKeyForAws_PublicKey_CallsUpdateSSHKeyWithNewPublicKey(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	auth := &Authentication{
		PublicKey:   types.StringValue(testNewKey),
		PublicKeyID: types.StringNull(),
	}

	mockClient.On("UpdateSSHKeyContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateSSHKeyParams) bool {
		return params.Input != nil &&
			params.Input.NewPublicKey == testNewKey &&
			*params.Input.Environment == testEnvName
	}), mock.Anything).Return(&operations.UpdateSSHKeyOK{}, nil)

	err := updateSshKeyForAws(ctx, client, auth, new(testEnvName))

	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}

func TestUpdateSshKeyForAws_PublicKeyID_CallsUpdateSSHKeyWithExistingPublicKeyID(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	auth := &Authentication{
		PublicKey:   types.StringNull(),
		PublicKeyID: types.StringValue(testPublicKeyID),
	}

	mockClient.On("UpdateSSHKeyContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateSSHKeyParams) bool {
		return params.Input != nil &&
			params.Input.ExistingPublicKeyID == testPublicKeyID &&
			*params.Input.Environment == testEnvName
	}), mock.Anything).Return(&operations.UpdateSSHKeyOK{}, nil)

	err := updateSshKeyForAws(ctx, client, auth, new(testEnvName))

	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}

func TestUpdateSshKeyForAws_NeitherSet_ReturnsError(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	auth := &Authentication{
		PublicKey:   types.StringNull(),
		PublicKeyID: types.StringNull(),
	}

	err := updateSshKeyForAws(ctx, client, auth, new(testEnvName))

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "either authentication.public_key or authentication.public_key_id must be set")
	mockClient.AssertNotCalled(t, "UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateSshKeyForAws_NilAuth_ReturnsNil(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	err := updateSshKeyForAws(ctx, client, nil, new(testEnvName))

	assert.NoError(t, err)
	mockClient.AssertNotCalled(t, "UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateSshKeyForAws_APIError_ReturnsError(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	auth := &Authentication{
		PublicKey:   types.StringValue(testNewKey),
		PublicKeyID: types.StringNull(),
	}

	mockClient.On("UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpdateSSHKeyOK)(nil), errors.New("connection refused"))

	err := updateSshKeyForAws(ctx, client, auth, new(testEnvName))

	assert.Error(t, err)
	assert.Equal(t, "connection refused", err.Error())
}

// Tests for updateAwsAuthenticationIfChanged

func TestUpdateAwsAuthenticationIfChanged_NilPlan_Skips(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &awsEnvironmentResourceModel{
		Authentication:  nil,
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &awsEnvironmentResourceModel{
		Authentication: &Authentication{PublicKey: types.StringValue(testOldKey)},
	}
	resp := &resource.UpdateResponse{}

	result := updateAwsAuthenticationIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAwsAuthenticationIfChanged_Unchanged_Skips(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	auth := &Authentication{PublicKey: types.StringValue(testSameKey), PublicKeyID: types.StringNull()}
	plan := &awsEnvironmentResourceModel{
		Authentication:  auth,
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &awsEnvironmentResourceModel{
		Authentication: auth,
	}
	resp := &resource.UpdateResponse{}

	result := updateAwsAuthenticationIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAwsAuthenticationIfChanged_Changed_CallsUpdateAndUpdatesState(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	planAuth := &Authentication{PublicKey: types.StringValue(testNewKey), PublicKeyID: types.StringNull()}
	stateAuth := &Authentication{PublicKey: types.StringValue(testOldKey), PublicKeyID: types.StringNull()}
	plan := &awsEnvironmentResourceModel{
		Authentication:  planAuth,
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &awsEnvironmentResourceModel{
		Authentication: stateAuth,
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything).
		Return(&operations.UpdateSSHKeyOK{}, nil)

	result := updateAwsAuthenticationIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, planAuth, state.Authentication)
	mockClient.AssertExpectations(t)
}

func TestUpdateAwsAuthenticationIfChanged_APIError_AddsDiagnostic(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	planAuth := &Authentication{PublicKey: types.StringValue(testNewKey), PublicKeyID: types.StringNull()}
	stateAuth := &Authentication{PublicKey: types.StringValue(testOldKey), PublicKeyID: types.StringNull()}
	plan := &awsEnvironmentResourceModel{
		Authentication:  planAuth,
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &awsEnvironmentResourceModel{
		Authentication: stateAuth,
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpdateSSHKeyOK)(nil), errors.New("API error"))

	result := updateAwsAuthenticationIfChanged(ctx, plan, state, client, resp)

	assert.True(t, result.Diagnostics.HasError())
}

// Tests for updateAwsEncryptionKeyIfChanged

func TestUpdateAwsEncryptionKeyIfChanged_Unchanged_Skips(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &awsEnvironmentResourceModel{
		EncryptionKeyArn: types.StringValue(testEncryptionKeyArn),
		EnvironmentName:  types.StringValue(testEnvName),
	}
	state := &awsEnvironmentResourceModel{
		EncryptionKeyArn: types.StringValue(testEncryptionKeyArn),
	}
	resp := &resource.UpdateResponse{}

	result := updateAwsEncryptionKeyIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateAwsDiskEncryptionParametersContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAwsEncryptionKeyIfChanged_Changed_CallsAPIAndUpdatesState(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &awsEnvironmentResourceModel{
		EncryptionKeyArn: types.StringValue(testNewEncryptionKeyArn),
		EnvironmentName:  types.StringValue(testEnvName),
	}
	state := &awsEnvironmentResourceModel{
		EncryptionKeyArn: types.StringValue(testEncryptionKeyArn),
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateAwsDiskEncryptionParametersContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateAwsDiskEncryptionParametersParams) bool {
		return params.Input != nil && *params.Input.EncryptionKeyArn == testNewEncryptionKeyArn
	}), mock.Anything).Return(&operations.UpdateAwsDiskEncryptionParametersOK{}, nil)

	result := updateAwsEncryptionKeyIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testNewEncryptionKeyArn), state.EncryptionKeyArn)
	mockClient.AssertExpectations(t)
}

func TestUpdateAwsEncryptionKeyIfChanged_APIError_AddsDiagnostic(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &awsEnvironmentResourceModel{
		EncryptionKeyArn: types.StringValue(testNewEncryptionKeyArn),
		EnvironmentName:  types.StringValue(testEnvName),
	}
	state := &awsEnvironmentResourceModel{
		EncryptionKeyArn: types.StringValue(testEncryptionKeyArn),
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateAwsDiskEncryptionParametersContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpdateAwsDiskEncryptionParametersOK)(nil), errors.New("access denied"))

	result := updateAwsEncryptionKeyIfChanged(ctx, plan, state, client, resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testEncryptionKeyArn), state.EncryptionKeyArn)
}

// Tests for enableComputeClusterForAws

func TestEnableComputeClusterForAws_Success(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	config := &AwsComputeClusterConfiguration{
		KubeApiAuthorizedIpRanges: utils.ToSetValueFromStringList([]string{testIpRange1}),
		PrivateCluster:            types.BoolValue(true),
		WorkerNodeSubnets:         utils.ToSetValueFromStringList([]string{testSubnet1}),
	}
	envSubnets := utils.ToSetValueFromStringList([]string{testFallbackSubnet1})

	mockClient.On("InitializeAWSComputeClusterContext", mock.Anything, mock.MatchedBy(func(params *operations.InitializeAWSComputeClusterParams) bool {
		return params.Input != nil &&
			*params.Input.EnvironmentName == testEnvName &&
			params.Input.ComputeClusterConfiguration != nil &&
			params.Input.ComputeClusterConfiguration.PrivateCluster == true
	}), mock.Anything).Return(&operations.InitializeAWSComputeClusterOK{}, nil)

	err := enableComputeClusterForAws(ctx, config, testEnvName, envSubnets, client)

	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}

func TestEnableComputeClusterForAws_APIError_ReturnsError(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	config := &AwsComputeClusterConfiguration{
		KubeApiAuthorizedIpRanges: utils.ToSetValueFromStringList([]string{}),
		PrivateCluster:            types.BoolValue(false),
		WorkerNodeSubnets:         utils.ToSetValueFromStringList([]string{testSubnet1}),
	}
	envSubnets := utils.ToSetValueFromStringList([]string{testSubnet1})

	mockClient.On("InitializeAWSComputeClusterContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.InitializeAWSComputeClusterOK)(nil), errors.New(testClusterInitFailed))

	err := enableComputeClusterForAws(ctx, config, testEnvName, envSubnets, client)

	assert.Error(t, err)
	assert.Equal(t, testClusterInitFailed, err.Error())
}

// Tests for updateAwsSecurityAccessIfChanged

func TestUpdateAwsSecurityAccessIfChanged_Changed_CallsAPI(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &awsEnvironmentResourceModel{
		SecurityAccess: &SecurityAccess{
			DefaultSecurityGroupID: types.StringValue(testNewDefaultSG),
			SecurityGroupIDForKnox: types.StringValue(testNewKnoxSG),
		},
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &awsEnvironmentResourceModel{
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

	result := updateAwsSecurityAccessIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testNewDefaultSG), state.SecurityAccess.DefaultSecurityGroupID)
	assert.Equal(t, types.StringValue(testNewKnoxSG), state.SecurityAccess.SecurityGroupIDForKnox)
	mockClient.AssertExpectations(t)
}

func TestUpdateAwsSecurityAccessIfChanged_Unchanged_SkipsAPICall(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &awsEnvironmentResourceModel{
		SecurityAccess: &SecurityAccess{
			DefaultSecurityGroupID: types.StringValue(testSameDefaultSG),
			SecurityGroupIDForKnox: types.StringValue(testSameKnoxSG),
		},
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &awsEnvironmentResourceModel{
		SecurityAccess: &SecurityAccess{
			DefaultSecurityGroupID: types.StringValue(testSameDefaultSG),
			SecurityGroupIDForKnox: types.StringValue(testSameKnoxSG),
		},
	}
	resp := &resource.UpdateResponse{}

	result := updateAwsSecurityAccessIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateSecurityAccessContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAwsSecurityAccessIfChanged_OnlySetFieldsDiffer_SkipsAPICall(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &awsEnvironmentResourceModel{
		SecurityAccess: &SecurityAccess{
			DefaultSecurityGroupID:  types.StringValue(testSameDefaultSG),
			SecurityGroupIDForKnox:  types.StringValue(testSameKnoxSG),
			DefaultSecurityGroupIDs: utils.ToSetValueFromStringList([]string{"sg-b", "sg-a"}),
		},
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &awsEnvironmentResourceModel{
		SecurityAccess: &SecurityAccess{
			DefaultSecurityGroupID:  types.StringValue(testSameDefaultSG),
			SecurityGroupIDForKnox:  types.StringValue(testSameKnoxSG),
			DefaultSecurityGroupIDs: utils.ToSetValueFromStringList([]string{"sg-a", "sg-b"}),
		},
	}
	resp := &resource.UpdateResponse{}

	result := updateAwsSecurityAccessIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateSecurityAccessContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateAwsSecurityAccessIfChanged_APIError_AddsDiagnostic(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &awsEnvironmentResourceModel{
		SecurityAccess: &SecurityAccess{
			DefaultSecurityGroupID: types.StringValue(testNewDefaultSG),
			SecurityGroupIDForKnox: types.StringValue(testNewKnoxSG),
		},
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &awsEnvironmentResourceModel{
		SecurityAccess: &SecurityAccess{
			DefaultSecurityGroupID: types.StringValue(testOldDefaultSG),
			SecurityGroupIDForKnox: types.StringValue(testOldKnoxSG),
		},
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateSecurityAccessContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpdateSecurityAccessOK)(nil), errors.New(testServiceUnavailable))

	result := updateAwsSecurityAccessIfChanged(ctx, plan, state, client, resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testOldDefaultSG), state.SecurityAccess.DefaultSecurityGroupID)
	assert.Equal(t, types.StringValue(testOldKnoxSG), state.SecurityAccess.SecurityGroupIDForKnox)
}
