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
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

type GcpEnvironmentResourceModelMock struct {
	mock.Mock
}

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
	assert.Equal(t, testObject.ExistingNetworkParams.SharedProjectId.ValueString(), result.ExistingNetworkParams.SharedProjectID)
	assert.Equal(t, len(testObject.ExistingNetworkParams.SubnetNames.Elements()), len(result.ExistingNetworkParams.SubnetNames))
}

func TestToGcpEnvironmentRequestSecurityAccess(t *testing.T) {
	testObject := createFilledGcpEnvironmentResourceModel()

	result := toGcpEnvironmentRequest(context.TODO(), testObject)

	assert.NotNilf(t, result.SecurityAccess, "SecurityAccess is expected to be not nil")
	assert.Equal(t, testObject.SecurityAccess.DefaultSecurityGroupId.ValueString(), result.SecurityAccess.DefaultSecurityGroupID)
	assert.Equal(t, testObject.SecurityAccess.SecurityGroupIdForKnox.ValueString(), result.SecurityAccess.SecurityGroupIDForKnox)
}

func TestToGcpEnvironmentRequestAvailabilityZones(t *testing.T) {
	testObject := createFilledGcpEnvironmentResourceModel()

	result := toGcpEnvironmentRequest(context.TODO(), testObject)

	assert.NotNilf(t, result.AvailabilityZones, "If AvailabilityZones is specified, it is expected to be not nil")
	assert.Equal(t, 1, len(testObject.AvailabilityZones))
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

	ctx := context.TODO()
	result := toGcpEnvironmentRequest(ctx, testObject)

	var freeIpaDetails FreeIpaDetails
	testObject.FreeIpa.As(ctx, &freeIpaDetails, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})

	assert.NotNilf(t, result.FreeIpa, "FreeIpa is expected to be not nil")
	assert.Equal(t, freeIpaDetails.InstanceCountByGroup.ValueInt64(), int64(result.FreeIpa.InstanceCountByGroup))
	assert.Equal(t, len(freeIpaDetails.Recipes.Elements()), len(result.FreeIpa.Recipes))
	assert.Equal(t, freeIpaDetails.InstanceType.ValueString(), result.FreeIpa.InstanceType)
}

func TestToGcpEnvironmentResourceRootFields(t *testing.T) {
	testModel := &gcpEnvironmentResourceModel{}
	testEnv := createFilledEnvironment()

	toGcpEnvironmentResource(context.TODO(), testEnv, testModel, getTestPollingOption(), nil)

	stringExists(t, *testEnv.Crn, testModel.Crn)
	stringExists(t, *testEnv.Region, testModel.Region)
	stringExists(t, *testEnv.Status, testModel.Status)
	stringExists(t, testEnv.Description, testModel.Description)
	stringExists(t, testEnv.StatusReason, testModel.StatusReason)
	stringExists(t, *testEnv.CredentialName, testModel.CredentialName)
	stringExists(t, *testEnv.EnvironmentName, testModel.EnvironmentName)
}

func stringExists(t *testing.T, expected string, underTest types.String) {
	assert.True(t, !underTest.IsNull())
	assert.True(t, !underTest.IsUnknown())
	assert.Equal(t, expected, underTest.ValueString())
}

func createFilledGcpEnvironmentResourceModel() *gcpEnvironmentResourceModel {
	var elems []attr.Value
	instances, _ := types.SetValue(FreeIpaInstanceType, elems)
	freeIpaObj, _ := basetypes.NewObjectValue(FreeIpaDetailsType.AttrTypes, map[string]attr.Value{
		"instance_count_by_group": types.Int64Value(123),
		"recipes":                 createSetOfString(),
		"instance_type":           types.StringValue("someInstanceType"),
		"catalog":                 types.StringValue(""),
		"image_id":                types.StringValue(""),
		"os":                      types.StringValue(""),
		"instances":               instances,
		"multi_az":                types.BoolValue(false),
	})
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
		Description:                 types.StringValue("someDescription"),
		EnableTunnel:                types.BoolValue(true),
		WorkloadAnalytics:           types.BoolValue(true),
		ReportDeploymentLogs:        types.BoolValue(true),
		FreeIpa:                     freeIpaObj,
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

func createFilledEnvironment() *environmentsmodels.Environment {
	falseVal := false
	return &environmentsmodels.Environment{
		Authentication: &environmentsmodels.Authentication{
			LoginUserName: "someLoginUserName",
			PublicKey:     "somePublicKey",
			PublicKeyID:   "somePublicKeyID",
		},
		AwsDetails: &environmentsmodels.EnvironmentAwsDetails{},
		BackupStorage: &environmentsmodels.BackupStorage{
			AwsDetails: &environmentsmodels.BackupStorageAwsDetails{
				InstanceProfile:     "someInstanceProfile",
				StorageLocationBase: "someStorageLocationBase",
			},
			AzureDetails: &environmentsmodels.BackupStorageAzureDetails{
				ManagedIdentity:     "someManagedIdentity",
				StorageLocationBase: "someStorageLocationBase",
			},
			Enabled: true,
			GcpDetails: &environmentsmodels.BackupStorageGcpDetails{
				ServiceAccountEmail: "someServiceAccountEmail",
				StorageLocationBase: "someStorageLocationBase",
			},
		},
		CloudPlatform:   func(s string) *string { return &s }("someCloudPlatform"),
		Created:         strfmt.NewDateTime(),
		Creator:         "someone",
		CredentialName:  func(s string) *string { return &s }("someCredential"),
		Crn:             func(s string) *string { return &s }("someCrn"),
		Description:     "someDescription",
		EnvironmentName: func(s string) *string { return &s }("someEnvironmentName"),
		Freeipa: &environmentsmodels.FreeipaDetails{
			Crn:      "someFreeIpaCrn",
			Domain:   "someDomain",
			Hostname: "someHostname",
			Instances: []*environmentsmodels.FreeIpaInstance{{
				AttachedVolumes: []*environmentsmodels.AttachedVolumeDetail{
					{
						Count:      1,
						Size:       1,
						VolumeType: "someVolumeType",
					},
				},
				AvailabilityZone:     "someAvailabilityZone",
				DiscoveryFQDN:        "someDiscoveryFQDN",
				InstanceGroup:        "someInstanceGroup",
				InstanceID:           "someInstanceID",
				InstanceStatus:       "someInstanceStatus",
				InstanceStatusReason: "someInstanceStatusReason",
				InstanceType:         "someInstanceType",
				InstanceVMType:       "someInstanceVMType",
				LifeCycle:            "someLifeCycle",
				PrivateIP:            "somePrivateIP",
				PublicIP:             "somePublicIP",
				SSHPort:              1,
				SubnetID:             "someSubnetID",
			}},
			Recipes:  []string{"someFreeIpaRecipe"},
			ServerIP: []string{"someServerIp"},
		},
		GcpDetails: &environmentsmodels.EnvironmentGcpDetails{},
		LogStorage: &environmentsmodels.LogStorage{
			AwsDetails: &environmentsmodels.LogStorageAwsDetails{
				InstanceProfile:     "someInstanceProfile",
				StorageLocationBase: "someStorageLocationBase",
			},
			AzureDetails: &environmentsmodels.LogStorageAzureDetails{
				ManagedIdentity:     "someManagedIdentity",
				StorageLocationBase: "someStorageLocationBase",
			},
			Enabled: func(b bool) *bool { return &b }(true),
			GcpDetails: &environmentsmodels.LogStorageGcpDetails{
				ServiceAccountEmail: "someServiceAccountEmail",
				StorageLocationBase: "someStorageLocationBase",
			},
		},
		Network: &environmentsmodels.Network{
			Aws: &environmentsmodels.NetworkAwsParams{},
			Azure: &environmentsmodels.NetworkAzureParams{
				AksPrivateDNSZoneID:        "someAksPrivateDNSZoneID",
				DatabasePrivateDNSZoneID:   "someDatabasePrivateDNSZoneID",
				EnableOutboundLoadBalancer: true,
				NetworkID:                  func(s string) *string { return &s }("someNetworkID"),
				ResourceGroupName:          func(s string) *string { return &s }("someResourceGroupName"),
				UsePublicIP:                func(b bool) *bool { return &b }(true),
			},
			EndpointAccessGatewayScheme:    "someEndpointAccessGatewayScheme",
			EndpointAccessGatewaySubnetIds: []string{"someEndpointAccessGatewaySubnetId"},
			Gcp: &environmentsmodels.NetworkGcpParams{
				NetworkName:     func(s string) *string { return &s }("someNetworkName"),
				SharedProjectID: "someSharedProjectID",
				UsePublicIP:     func(b bool) *bool { return &b }(true),
			},
			NetworkCidr: "someNetworkCidr",
			NetworkName: func(s string) *string { return &s }("someNetworkName"),
			SubnetIds:   []string{"someSubnetId"},
			SubnetMetadata: map[string]environmentsmodels.CloudSubnet{
				"someSubnetMetadata": {
					AvailabilityZone: "someAvailabilityZone",
					Cidr:             "someCidr",
					SubnetID:         "someSubnetID",
					SubnetName:       "someSubnetName",
				},
			},
		},
		ProxyConfig: &environmentsmodels.ProxyConfig{
			Crn:             func(s string) *string { return &s }("someProxyCrn"),
			Description:     "someDescription",
			Host:            func(s string) *string { return &s }("someHost"),
			NoProxyHosts:    "someNoProxyHosts",
			Password:        "somePassword",
			Port:            func(i int32) *int32 { return &i }(int32(1)),
			Protocol:        func(s string) *string { return &s }("someProtocol"),
			ProxyConfigName: func(s string) *string { return &s }("someProxyConfigName"),
			User:            "someUser",
		},
		Region:               func(s string) *string { return &s }("someRegion"),
		ReportDeploymentLogs: &falseVal,
		SecurityAccess: &environmentsmodels.SecurityAccess{
			Cidr:                   "someCidr",
			DefaultSecurityGroupID: "someDefaultSecurityGroupID",
			SecurityGroupIDForKnox: "someSecurityGroupIDForKnox",
		},
		Status:       func(s string) *string { return &s }("someStatus"),
		StatusReason: "someStatusReason",
		Tags: &environmentsmodels.EnvironmentTags{
			Defaults: map[string]string{
				"someDefaultKey": "someDefaultValue",
			},
			UserDefined: map[string]string{
				"someUserDefinedKey": "someUserDefinedValue",
			},
		},
		TunnelEnabled:     true,
		TunnelType:        environmentsmodels.TunnelTypeCCMV2,
		WorkloadAnalytics: true,
	}
}

func getTestPollingOption() *utils.PollingOptions {
	return &utils.PollingOptions{PollingTimeout: types.Int64Value(1234)}
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
