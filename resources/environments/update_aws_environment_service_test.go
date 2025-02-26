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
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"

	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
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
		KubeApiAuthorizedIpRanges: utils.ToSetValueFromStringList([]string{"192.168.1.1/32", "10.0.0.0/24"}),
		PrivateCluster:            types.BoolValue(true),
		WorkerNodeSubnets:         utils.ToSetValueFromStringList([]string{"subnet-12345", "subnet-67890"}),
	}
	fallbackSubnetIds := types.Set{}
	want := &environmentsmodels.AWSComputeClusterConfigurationRequest{
		KubeAPIAuthorizedIPRanges: []string{"192.168.1.1/32", "10.0.0.0/24"},
		PrivateCluster:            true,
		WorkerNodeSubnets:         []string{"subnet-12345", "subnet-67890"},
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
		KubeApiAuthorizedIpRanges: utils.ToSetValueFromStringList([]string{"192.168.1.1/32"}),
		PrivateCluster:            types.BoolValue(true),
		WorkerNodeSubnets:         types.Set{},
	}
	fallbackSubnetIds := utils.ToSetValueFromStringList([]string{"subnet-abcde", "subnet-fghij"})
	want := &environmentsmodels.AWSComputeClusterConfigurationRequest{
		KubeAPIAuthorizedIPRanges: []string{"192.168.1.1/32"},
		PrivateCluster:            true,
		WorkerNodeSubnets:         []string{"subnet-abcde", "subnet-fghij"},
	}
	if got := convertConfigToAwsComputeClusterConfigurationRequest(config, fallbackSubnetIds); !reflect.DeepEqual(got, want) {
		t.Errorf("convertConfigToAwsComputeClusterConfigurationRequest() = %v, want %v", got, want)
	}
}
