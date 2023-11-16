// Copyright 2023 Cloudera. All Rights Reserved.
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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

func TestToGcpEnvironmentRequestRootFields(t *testing.T) {
	testObject := createFilledGcpEnvironmentResourceModel()

	result := toGcpEnvironmentRequest(context.TODO(), testObject)

	assert.Equal(t, testObject.EnvironmentName.ValueString(), *result.EnvironmentName)
	assert.Equal(t, testObject.CredentialName.ValueString(), *result.CredentialName)
	assert.Equal(t, testObject.Region.ValueString(), *result.Region)
	assert.Equal(t, testObject.PublicKey.ValueString(), *result.PublicKey)
	assert.Equal(t, testObject.Description.ValueString(), result.Description)
	assert.Equal(t, testObject.EndpointAccessGatewayScheme.ValueString(), result.EndpointAccessGatewayScheme)
	assert.Equal(t, testObject.ProxyConfigName.ValueString(), result.ProxyConfigName)
	assert.Equal(t, testObject.EncryptionKey.ValueString(), result.EncryptionKey)
	assert.Equal(t, len(testObject.Tags.Elements()), len(result.Tags))
}

func TestToGcpEnvironmentRequestExistingNetworkParams(t *testing.T) {
	testObject := createFilledGcpEnvironmentResourceModel()

	result := toGcpEnvironmentRequest(context.TODO(), testObject)

	assert.NotNilf(t, result.ExistingNetworkParams, "ExistingNetworkParams is expected to be not nil")
	assert.Equal(t, testObject.ExistingNetworkParams.NetworkName.ValueString(), *result.ExistingNetworkParams.NetworkName)
	assert.Equal(t, testObject.ExistingNetworkParams.SharedProjectId.ValueString(), *result.ExistingNetworkParams.SharedProjectID)
	assert.Equal(t, len(testObject.ExistingNetworkParams.SubnetNames.Elements()), len(result.ExistingNetworkParams.SubnetNames))
}

func TestToGcpEnvironmentRequestSecurityAccess(t *testing.T) {
	testObject := createFilledGcpEnvironmentResourceModel()

	result := toGcpEnvironmentRequest(context.TODO(), testObject)

	assert.NotNilf(t, result.SecurityAccess, "SecurityAccess is expected to be not nil")
	assert.Equal(t, testObject.SecurityAccess.DefaultSecurityGroupId.ValueString(), result.SecurityAccess.DefaultSecurityGroupID)
	assert.Equal(t, testObject.SecurityAccess.SecurityGroupIdForKnox.ValueString(), result.SecurityAccess.SecurityGroupIDForKnox)
}

func TestToGcpEnvironmentRequestLogStorage(t *testing.T) {
	testObject := createFilledGcpEnvironmentResourceModel()

	result := toGcpEnvironmentRequest(context.TODO(), testObject)

	assert.NotNilf(t, result.LogStorage, "LogStorage is expected to be not nil")
	assert.Equal(t, testObject.LogStorage.StorageLocationBase.ValueString(), *result.LogStorage.StorageLocationBase)
	assert.Equal(t, testObject.LogStorage.BackupStorageLocationBase.ValueString(), result.LogStorage.BackupStorageLocationBase)
	assert.Equal(t, testObject.LogStorage.ServiceAccountEmail.ValueString(), *result.LogStorage.ServiceAccountEmail)
}

func TestToGcpEnvironmentRequestFreeIpa(t *testing.T) {
	testObject := createFilledGcpEnvironmentResourceModel()

	result := toGcpEnvironmentRequest(context.TODO(), testObject)

	assert.NotNilf(t, result.FreeIpa, "FreeIpa is expected to be not nil")
	assert.Equal(t, testObject.FreeIpa.InstanceCountByGroup.ValueInt64(), int64(result.FreeIpa.InstanceCountByGroup))
	assert.Equal(t, len(testObject.FreeIpa.Recipes.Elements()), len(result.FreeIpa.Recipes))
	assert.Equal(t, testObject.FreeIpa.InstanceType.ValueString(), result.FreeIpa.InstanceType)
}

func createFilledGcpEnvironmentResourceModel() *gcpEnvironmentResourceModel {
	return &gcpEnvironmentResourceModel{
		EnvironmentName: types.StringValue("someEnvironmentName"),
		PollingOptions:  &utils.PollingOptions{PollingTimeout: types.Int64Value(123)},
		CredentialName:  types.StringValue("someCredentialName"),
		Region:          types.StringValue("someRegion"),
		PublicKey:       types.StringValue("somePublicKey"),
		UsePublicIp:     types.BoolValue(true),
		ExistingNetworkParams: &ExistingNetworkParams{
			NetworkName:     types.StringValue("someNetworkName"),
			SubnetNames:     createListOfStrings(),
			SharedProjectId: types.StringValue("someSharedProjectId"),
		},
		SecurityAccess: &GcpSecurityAccess{
			SecurityGroupIdForKnox: types.StringValue("someSecurityGroupIdForKnox"),
			DefaultSecurityGroupId: types.StringValue("someDefaultSecurityGroupId"),
		},
		LogStorage: &GcpLogStorage{
			StorageLocationBase:       types.StringValue("someStorageLocationBase"),
			ServiceAccountEmail:       types.StringValue("someServiceAccountEmail"),
			BackupStorageLocationBase: types.StringValue("someBackupStorageLocationBase"),
		},
		Description:          types.StringValue("someDescription"),
		EnableTunnel:         types.BoolValue(true),
		WorkloadAnalytics:    types.BoolValue(true),
		ReportDeploymentLogs: types.BoolValue(true),
		FreeIpa: &GcpFreeIpa{
			InstanceCountByGroup: types.Int64Value(123),
			Recipes:              createSetOfString(),
			InstanceType:         types.StringValue("someInstanceType"),
		},
		EndpointAccessGatewayScheme: types.StringValue("someEndpointAccessGatewayScheme"),
		Tags:                        createMapOfString(),
		ProxyConfigName:             types.StringValue("someProxyConfigName"),
		EncryptionKey:               types.StringValue("someEncryptionKey"),
		AvailabilityZones:           []types.String{types.StringValue("someAvailabilityZone1")},
		ID:                          types.StringValue("someId"),
		Crn:                         types.StringValue("someCrn"),
		Status:                      types.StringValue("someStatus"),
		StatusReason:                types.StringValue("someStatusReason"),
	}
}

func createListOfStrings() types.List {
	var elems []attr.Value
	elems = make([]attr.Value, 0)
	elems = append(elems, types.StringValue("someListElement"))
	var list types.List
	list, _ = types.ListValue(types.StringType, elems)
	return list
}

func createSetOfString() types.Set {
	var elems []attr.Value
	elems = make([]attr.Value, 0)
	elems = append(elems, types.StringValue("someSetElement"))
	var set types.Set
	set, _ = types.SetValue(types.StringType, elems)
	return set
}

func createMapOfString() types.Map {
	elems := make(map[string]attr.Value)
	elems["someMapKey"] = types.StringValue("someMapValue")
	var mapType types.Map
	mapType, _ = types.MapValue(types.StringType, elems)
	return mapType
}
