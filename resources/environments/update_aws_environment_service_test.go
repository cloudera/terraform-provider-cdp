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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/mock"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/mocks"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

const (
	testIpRange1         = "192.168.1.1/32"
	testIpRange2         = "10.0.0.0/24"
	testSubnet1          = "subnet-12345"
	testSubnet2          = "subnet-67890"
	testFallbackSubnet1  = "subnet-abcde"
	testFallbackSubnet2  = "subnet-fghij"
	testEncryptionKeyArn = "arn:aws:kms:us-west-2:123456789012:key/test-key-id"
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

	mockClient.On("UpdateAwsDiskEncryptionParameters", mock.MatchedBy(func(params *operations.UpdateAwsDiskEncryptionParametersParams) bool {
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

	mockClient.On("UpdateAwsDiskEncryptionParameters", mock.Anything, mock.Anything).
		Return((*operations.UpdateAwsDiskEncryptionParametersOK)(nil), errors.New(testServiceUnavailable))

	err := updateDiskEncryption(ctx, client, new(testEnvName), keyArn)

	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	if err.Error() != testServiceUnavailable {
		t.Errorf("expected error message '%s', got: %s", testServiceUnavailable, err.Error())
	}
}
