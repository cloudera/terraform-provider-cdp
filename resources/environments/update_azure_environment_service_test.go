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
