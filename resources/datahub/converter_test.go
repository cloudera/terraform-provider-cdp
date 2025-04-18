// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datahub

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/utils/test"
)

func TestFromModelToRequestBasicFields(t *testing.T) {
	subnetIds, _ := types.SetValue(types.StringType, []attr.Value{types.StringValue("someSubnetIds")})
	tags, _ := types.MapValue(types.StringType, map[string]attr.Value{"key": types.StringValue("value")})
	image, _ := types.ObjectValue(map[string]attr.Type{
		"catalog": types.StringType,
		"id":      types.StringType,
		"os":      types.StringType,
	}, map[string]attr.Value{
		"catalog": types.StringValue("someCatalog"),
		"id":      types.StringValue("someId"),
		"os":      types.StringValue("someOs"),
	})
	clusterExt, _ := types.ObjectValue(map[string]attr.Type{
		"custom_properties": types.StringType,
	}, map[string]attr.Value{
		"custom_properties": types.StringValue("someProps"),
	})
	input := awsDatahubResourceModel{
		Name:                     types.StringValue("someClusterName"),
		Environment:              types.StringValue("someEnvironment"),
		ClusterTemplate:          types.StringValue("someClusterTemplateNameOrCRN"),
		SubnetId:                 types.StringValue("someSubnetID"),
		SubnetIds:                subnetIds,
		MultiAz:                  types.BoolValue(true),
		Tags:                     tags,
		CustomConfigurationsName: types.StringValue("someCustomConfigurationsName"),
		Image:                    image,
		RequestTemplate:          types.StringValue("someRequestTemplate"),
		DatahubDatabase:          types.StringValue("someDatahubDatabase"),
		ClusterExtension:         clusterExt,
		EnableLoadBalancer:       types.BoolValue(true),
		JavaVersion:              types.Int32Value(11),
	}
	got := fromModelToAwsRequest(input, context.TODO())

	test.CompareStrings(got.ClusterName, input.Name.ValueString(), t)
	test.CompareStrings(got.Environment, input.Environment.ValueString(), t)
	test.CompareStrings(got.ClusterTemplate, input.ClusterTemplate.ValueString(), t)
	test.CompareStrings(got.SubnetID, input.SubnetId.ValueString(), t)
	test.CompareStringValueSlices(got.SubnetIds, input.SubnetIds.Elements(), t)
	test.CompareBools(got.MultiAz, input.MultiAz.ValueBool(), t)
	test.CompareStrings(*got.Tags[0].Value, input.Tags.Elements()["key"].(types.String).ValueString(), t)
	test.CompareStrings(got.CustomConfigurationsName, input.CustomConfigurationsName.ValueString(), t)
	test.CompareStrings(got.Image.CatalogName, input.Image.Attributes()["catalog"].(types.String).ValueString(), t)
	test.CompareStrings(got.Image.ID, input.Image.Attributes()["id"].(types.String).ValueString(), t)
	test.CompareStrings(got.Image.Os, input.Image.Attributes()["os"].(types.String).ValueString(), t)
	test.CompareStrings(got.RequestTemplate, input.RequestTemplate.ValueString(), t)
	test.CompareStrings(string(got.DatahubDatabase), input.DatahubDatabase.ValueString(), t)
	test.CompareStrings(got.ClusterExtension.CustomProperties, input.ClusterExtension.Attributes()["custom_properties"].(types.String).ValueString(), t)
	test.CompareBools(got.EnableLoadBalancer, input.EnableLoadBalancer.ValueBool(), t)
	test.CompareInt32PointerToTypesInt32(&got.JavaVersion, input.JavaVersion, t)
}

func TestFromModelToRequestRecipe(t *testing.T) {
	recipes := []types.String{types.StringValue("recipe1"), types.StringValue("recipe2")}
	igs := []InstanceGroup{{Recipes: recipes}}
	input := awsDatahubResourceModel{InstanceGroup: igs}

	got := fromModelToAwsRequest(input, context.TODO())

	test.CompareInts(len(got.InstanceGroups), len(input.InstanceGroup), t)
	test.CompareInts(len(got.InstanceGroups[0].RecipeNames), len(input.InstanceGroup[0].Recipes), t)

	for _, convertedRecipe := range got.InstanceGroups[0].RecipeNames {
		var contains bool
		for _, originalRecipe := range input.InstanceGroup[0].Recipes {
			if originalRecipe.ValueString() == convertedRecipe {
				contains = true
			}
		}
		if !contains {
			t.Errorf("Instance group does not contain recipe: %s", convertedRecipe)
		}
	}
}

func TestFromModelToRequestAttachedVolumeConfiguration(t *testing.T) {
	avcs := []AttachedVolumeConfiguration{{
		VolumeSize:  types.Int32Value(100),
		VolumeCount: types.Int32Value(1),
		VolumeType:  types.StringValue("ephemeral"),
	}}
	igs := []InstanceGroup{{AttachedVolumeConfiguration: avcs}}
	input := awsDatahubResourceModel{InstanceGroup: igs}

	got := fromModelToAwsRequest(input, context.TODO())

	test.CompareInts(len(got.InstanceGroups), len(input.InstanceGroup), t)
	test.CompareInts(len(got.InstanceGroups[0].AttachedVolumeConfiguration), len(avcs), t)

	resultAvcs := got.InstanceGroups[0].AttachedVolumeConfiguration[0]
	test.CompareInt32PointerToTypesInt32(resultAvcs.VolumeCount, avcs[0].VolumeCount, t)
	test.CompareInt32PointerToTypesInt32(resultAvcs.VolumeSize, avcs[0].VolumeSize, t)
	test.CompareStrings(*resultAvcs.VolumeType, avcs[0].VolumeType.ValueString(), t)
}

func TestFromModelToRequestInstanceGroups(t *testing.T) {
	igs := []InstanceGroup{{
		NodeCount:         types.Int32Value(1),
		InstanceGroupName: types.StringValue("gateway"),
		InstanceGroupType: types.StringValue("CORE"),
		InstanceType:      types.StringValue("m5.xlarge"),
		RootVolumeSize:    types.Int32Value(100),
		RecoveryMode:      types.StringValue("MANUAL"),
	}}

	input := awsDatahubResourceModel{InstanceGroup: igs}

	got := fromModelToAwsRequest(input, context.TODO())

	test.CompareInts(len(got.InstanceGroups), len(igs), t)
	resultIg := got.InstanceGroups[0]
	test.CompareStrings(*resultIg.InstanceGroupName, igs[0].InstanceGroupName.ValueString(), t)
	test.CompareStrings(*resultIg.InstanceGroupType, igs[0].InstanceGroupType.ValueString(), t)
	test.CompareStrings(*resultIg.InstanceType, igs[0].InstanceType.ValueString(), t)
	test.CompareInt32PointerToTypesInt32(&resultIg.RootVolumeSize, igs[0].RootVolumeSize, t)
	test.CompareStrings(resultIg.RecoveryMode, igs[0].RecoveryMode.ValueString(), t)
}

func TestFromModelToRequestVolumeEncryption(t *testing.T) {
	igs := []InstanceGroup{{
		VolumeEncryption: VolumeEncryption{Encryption: types.BoolValue(true)},
	}}

	input := awsDatahubResourceModel{InstanceGroup: igs}

	got := fromModelToAwsRequest(input, context.TODO())

	test.CompareInts(len(got.InstanceGroups), len(igs), t)
	resultVolumeEncryption := got.InstanceGroups[0].VolumeEncryption
	if resultVolumeEncryption == nil {
		t.Errorf("Volume encryption is not filled though it should've been!")
	} else {
		test.CompareBools(*resultVolumeEncryption.EnableEncryption, igs[0].VolumeEncryption.Encryption.ValueBool(), t)
	}
}

func TestFromModelToRequestClusterDefinition(t *testing.T) {
	input := awsDatahubResourceModel{ClusterDefinition: types.StringValue("SomeClusterDef")}
	got := fromModelToAwsRequest(input, context.TODO())

	test.CompareStrings(got.ClusterDefinition, input.ClusterDefinition.ValueString(), t)
}

func TestFromModelToRequestTags(t *testing.T) {
	tags, _ := types.MapValue(types.StringType, map[string]attr.Value{
		"key1": types.StringValue("value1"),
		"key2": types.StringValue("value2"),
		"key3": types.StringValue("value3"),
	})
	input := awsDatahubResourceModel{Tags: tags}
	got := fromModelToAwsRequest(input, context.TODO())

	expectedKeys := [3]string{"key1", "key2", "key1"}
	gotKeys := [3]string{*got.Tags[0].Key, *got.Tags[1].Key, *got.Tags[2].Key}
	test.CompareInts(len(got.Tags), 3, t)
	test.CompareStringSlices(gotKeys[:], expectedKeys[:])
	test.CompareStrings(*got.Tags[0].Value, input.Tags.Elements()[*got.Tags[0].Key].(types.String).ValueString(), t)
	test.CompareStrings(*got.Tags[1].Value, input.Tags.Elements()[*got.Tags[1].Key].(types.String).ValueString(), t)
	test.CompareStrings(*got.Tags[2].Value, input.Tags.Elements()[*got.Tags[2].Key].(types.String).ValueString(), t)
}

func TestFromModelToGcpRequestBasicFields(t *testing.T) {
	tags, _ := types.MapValue(types.StringType, map[string]attr.Value{"key": types.StringValue("value")})
	image, _ := types.ObjectValue(map[string]attr.Type{
		"catalog": types.StringType,
		"id":      types.StringType,
		"os":      types.StringType,
	}, map[string]attr.Value{
		"catalog": types.StringValue("someCatalog"),
		"id":      types.StringValue("someId"),
		"os":      types.StringValue("someOs"),
	})
	clusterExt, _ := types.ObjectValue(map[string]attr.Type{
		"custom_properties": types.StringType,
	}, map[string]attr.Value{
		"custom_properties": types.StringValue("someProps"),
	})
	input := gcpDatahubResourceModel{
		Name:                     types.StringValue("someClusterName"),
		Environment:              types.StringValue("someEnvironment"),
		ClusterTemplate:          types.StringValue("someClusterTemplateNameOrCRN"),
		SubnetName:               types.StringValue("someSubnetID"),
		Tags:                     tags,
		CustomConfigurationsName: types.StringValue("someCustomConfigurationsName"),
		Image:                    image,
		RequestTemplate:          types.StringValue("someRequestTemplate"),
		DatahubDatabase:          types.StringValue("someDatahubDatabase"),
		ClusterExtension:         clusterExt,
		JavaVersion:              types.Int32Value(11),
	}
	got := fromModelToGcpRequest(input, context.TODO())

	test.CompareStrings(got.ClusterName, input.Name.ValueString(), t)
	test.CompareStrings(got.EnvironmentName, input.Environment.ValueString(), t)
	test.CompareStrings(got.ClusterTemplateName, input.ClusterTemplate.ValueString(), t)
	test.CompareStrings(got.SubnetName, input.SubnetName.ValueString(), t)
	test.CompareStrings(*got.Tags[0].Value, input.Tags.Elements()["key"].(types.String).ValueString(), t)
	test.CompareStrings(got.CustomConfigurationsName, input.CustomConfigurationsName.ValueString(), t)
	test.CompareStrings(got.Image.CatalogName, input.Image.Attributes()["catalog"].(types.String).ValueString(), t)
	test.CompareStrings(got.Image.ID, input.Image.Attributes()["id"].(types.String).ValueString(), t)
	test.CompareStrings(got.Image.Os, input.Image.Attributes()["os"].(types.String).ValueString(), t)
	test.CompareStrings(got.RequestTemplate, input.RequestTemplate.ValueString(), t)
	test.CompareStrings(string(got.DatahubDatabase), input.DatahubDatabase.ValueString(), t)
	test.CompareStrings(got.ClusterExtension.CustomProperties, input.ClusterExtension.Attributes()["custom_properties"].(types.String).ValueString(), t)
	test.CompareInt32PointerToTypesInt32(&got.JavaVersion, input.JavaVersion, t)
}

func TestFromModelToGcpRequestRecipe(t *testing.T) {
	recipes := []types.String{types.StringValue("recipe1"), types.StringValue("recipe2")}
	igs := []GcpInstanceGroup{{Recipes: recipes}}
	input := gcpDatahubResourceModel{InstanceGroup: igs}

	got := fromModelToGcpRequest(input, context.TODO())

	test.CompareInts(len(got.InstanceGroups), len(input.InstanceGroup), t)
	test.CompareInts(len(got.InstanceGroups[0].RecipeNames), len(input.InstanceGroup[0].Recipes), t)

	for _, convertedRecipe := range got.InstanceGroups[0].RecipeNames {
		var contains bool
		for _, originalRecipe := range input.InstanceGroup[0].Recipes {
			if originalRecipe.ValueString() == convertedRecipe {
				contains = true
			}
		}
		if !contains {
			t.Errorf("Instance group does not contain recipe: %s", convertedRecipe)
		}
	}
}

func TestFromModelToGcpRequestAttachedVolumeConfiguration(t *testing.T) {
	avcs := []AttachedVolumeConfiguration{{
		VolumeSize:  types.Int32Value(100),
		VolumeCount: types.Int32Value(1),
		VolumeType:  types.StringValue("ephemeral"),
	}}
	igs := []GcpInstanceGroup{{AttachedVolumeConfiguration: avcs}}
	input := gcpDatahubResourceModel{InstanceGroup: igs}

	got := fromModelToGcpRequest(input, context.TODO())

	test.CompareInts(len(got.InstanceGroups), len(input.InstanceGroup), t)
	test.CompareInts(len(got.InstanceGroups[0].AttachedVolumeConfiguration), len(avcs), t)

	resultAvcs := got.InstanceGroups[0].AttachedVolumeConfiguration[0]
	test.CompareInt32PointerToTypesInt32(resultAvcs.VolumeCount, avcs[0].VolumeCount, t)
	test.CompareInt32PointerToTypesInt32(resultAvcs.VolumeSize, avcs[0].VolumeSize, t)
	test.CompareStrings(*resultAvcs.VolumeType, avcs[0].VolumeType.ValueString(), t)
}

func TestFromModelToGcpRequestInstanceGroups(t *testing.T) {
	igs := []GcpInstanceGroup{{
		NodeCount:         types.Int32Value(1),
		InstanceGroupName: types.StringValue("gateway"),
		InstanceGroupType: types.StringValue("CORE"),
		InstanceType:      types.StringValue("m5.xlarge"),
		RootVolumeSize:    types.Int32Value(100),
		RecoveryMode:      types.StringValue("MANUAL"),
	}}

	input := gcpDatahubResourceModel{InstanceGroup: igs}

	got := fromModelToGcpRequest(input, context.TODO())

	test.CompareInts(len(got.InstanceGroups), len(igs), t)
	resultIg := got.InstanceGroups[0]
	test.CompareStrings(*resultIg.InstanceGroupName, igs[0].InstanceGroupName.ValueString(), t)
	test.CompareStrings(*resultIg.InstanceGroupType, igs[0].InstanceGroupType.ValueString(), t)
	test.CompareStrings(*resultIg.InstanceType, igs[0].InstanceType.ValueString(), t)
	test.CompareInt32PointerToTypesInt32(resultIg.RootVolumeSize, igs[0].RootVolumeSize, t)
	test.CompareStrings(resultIg.RecoveryMode, igs[0].RecoveryMode.ValueString(), t)
}

func TestFromModelToGcpRequestClusterDefinition(t *testing.T) {
	input := gcpDatahubResourceModel{ClusterDefinition: types.StringValue("SomeClusterDef")}
	got := fromModelToGcpRequest(input, context.TODO())

	test.CompareStrings(got.ClusterDefinitionName, input.ClusterDefinition.ValueString(), t)
}

func TestFromModelToAzureRequestBasicFields(t *testing.T) {
	tags, _ := types.MapValue(types.StringType, map[string]attr.Value{"key": types.StringValue("value")})
	image, _ := types.ObjectValue(map[string]attr.Type{
		"catalog": types.StringType,
		"id":      types.StringType,
		"os":      types.StringType,
	}, map[string]attr.Value{
		"catalog": types.StringValue("someCatalog"),
		"id":      types.StringValue("someId"),
		"os":      types.StringValue("someOs"),
	})
	clusterExt, _ := types.ObjectValue(map[string]attr.Type{
		"custom_properties": types.StringType,
	}, map[string]attr.Value{
		"custom_properties": types.StringValue("someProps"),
	})
	input := azureDatahubResourceModel{
		Name:                            types.StringValue("someClusterName"),
		Environment:                     types.StringValue("someEnvironment"),
		ClusterTemplate:                 types.StringValue("someClusterTemplateNameOrCRN"),
		SubnetId:                        types.StringValue("someSubnetID"),
		MultiAz:                         types.BoolValue(true),
		Tags:                            tags,
		CustomConfigurationsName:        types.StringValue("someCustomConfigurationsName"),
		Image:                           image,
		RequestTemplate:                 types.StringValue("someRequestTemplate"),
		DatahubDatabase:                 types.StringValue("someDatahubDatabase"),
		ClusterExtension:                clusterExt,
		EnableLoadBalancer:              types.BoolValue(true),
		JavaVersion:                     types.Int32Value(11),
		FlexibleServerDelegatedSubnetId: types.StringValue("someFlexibleServerDelegatedSubnetId"),
	}
	got := fromModelToAzureRequest(input, context.TODO())

	test.CompareStrings(got.ClusterName, input.Name.ValueString(), t)
	test.CompareStrings(got.EnvironmentName, input.Environment.ValueString(), t)
	test.CompareStrings(got.ClusterTemplateName, input.ClusterTemplate.ValueString(), t)
	test.CompareStrings(got.SubnetID, input.SubnetId.ValueString(), t)
	test.CompareBools(*got.MultiAz, input.MultiAz.ValueBool(), t)
	test.CompareStrings(*got.Tags[0].Value, input.Tags.Elements()["key"].(types.String).ValueString(), t)
	test.CompareStrings(got.CustomConfigurationsName, input.CustomConfigurationsName.ValueString(), t)
	test.CompareStrings(got.Image.CatalogName, input.Image.Attributes()["catalog"].(types.String).ValueString(), t)
	test.CompareStrings(got.Image.ID, input.Image.Attributes()["id"].(types.String).ValueString(), t)
	test.CompareStrings(got.Image.Os, input.Image.Attributes()["os"].(types.String).ValueString(), t)
	test.CompareStrings(got.RequestTemplate, input.RequestTemplate.ValueString(), t)
	test.CompareStrings(string(got.DatahubDatabase), input.DatahubDatabase.ValueString(), t)
	test.CompareStrings(got.ClusterExtension.CustomProperties, input.ClusterExtension.Attributes()["custom_properties"].(types.String).ValueString(), t)
	test.CompareBools(got.EnableLoadBalancer, input.EnableLoadBalancer.ValueBool(), t)
	test.CompareInt32PointerToTypesInt32(&got.JavaVersion, input.JavaVersion, t)
	test.CompareStrings(got.FlexibleServerDelegatedSubnetID, input.FlexibleServerDelegatedSubnetId.ValueString(), t)
}

func TestFromModelToAzureRequestRecipe(t *testing.T) {
	recipes := []types.String{types.StringValue("recipe1"), types.StringValue("recipe2")}
	igs := []InstanceGroup{{Recipes: recipes}}
	input := azureDatahubResourceModel{InstanceGroup: igs}

	got := fromModelToAzureRequest(input, context.TODO())

	test.CompareInts(len(got.InstanceGroups), len(input.InstanceGroup), t)
	test.CompareInts(len(got.InstanceGroups[0].RecipeNames), len(input.InstanceGroup[0].Recipes), t)

	for _, convertedRecipe := range got.InstanceGroups[0].RecipeNames {
		var contains bool
		for _, originalRecipe := range input.InstanceGroup[0].Recipes {
			if originalRecipe.ValueString() == convertedRecipe {
				contains = true
			}
		}
		if !contains {
			t.Errorf("Instance group does not contain recipe: %s", convertedRecipe)
		}
	}
}

func TestFromModelToAzureRequestAttachedVolumeConfiguration(t *testing.T) {
	avcs := []AttachedVolumeConfiguration{{
		VolumeSize:  types.Int32Value(100),
		VolumeCount: types.Int32Value(1),
		VolumeType:  types.StringValue("ephemeral"),
	}}
	igs := []InstanceGroup{{AttachedVolumeConfiguration: avcs}}
	input := azureDatahubResourceModel{InstanceGroup: igs}

	got := fromModelToAzureRequest(input, context.TODO())

	test.CompareInts(len(got.InstanceGroups), len(input.InstanceGroup), t)
	test.CompareInts(len(got.InstanceGroups[0].AttachedVolumeConfiguration), len(avcs), t)

	resultAvcs := got.InstanceGroups[0].AttachedVolumeConfiguration[0]
	test.CompareInt32PointerToTypesInt32(resultAvcs.VolumeCount, avcs[0].VolumeCount, t)
	test.CompareInt32PointerToTypesInt32(resultAvcs.VolumeSize, avcs[0].VolumeSize, t)
	test.CompareStrings(*resultAvcs.VolumeType, avcs[0].VolumeType.ValueString(), t)
}

func TestFromModelToAzureRequestInstanceGroups(t *testing.T) {
	igs := []InstanceGroup{{
		NodeCount:         types.Int32Value(1),
		InstanceGroupName: types.StringValue("gateway"),
		InstanceGroupType: types.StringValue("CORE"),
		InstanceType:      types.StringValue("m5.xlarge"),
		RootVolumeSize:    types.Int32Value(100),
		RecoveryMode:      types.StringValue("MANUAL"),
	}}

	input := azureDatahubResourceModel{InstanceGroup: igs}

	got := fromModelToAzureRequest(input, context.TODO())

	test.CompareInts(len(got.InstanceGroups), len(igs), t)
	resultIg := got.InstanceGroups[0]
	test.CompareStrings(*resultIg.InstanceGroupName, igs[0].InstanceGroupName.ValueString(), t)
	test.CompareStrings(*resultIg.InstanceGroupType, igs[0].InstanceGroupType.ValueString(), t)
	test.CompareStrings(*resultIg.InstanceType, igs[0].InstanceType.ValueString(), t)
	test.CompareInt32PointerToTypesInt32(resultIg.RootVolumeSize, igs[0].RootVolumeSize, t)
	test.CompareStrings(resultIg.RecoveryMode, igs[0].RecoveryMode.ValueString(), t)
}

func TestFromModelToAzureRequestClusterDefinition(t *testing.T) {
	input := azureDatahubResourceModel{ClusterDefinition: types.StringValue("SomeClusterDef")}
	got := fromModelToAzureRequest(input, context.TODO())

	test.CompareStrings(got.ClusterDefinitionName, input.ClusterDefinition.ValueString(), t)
}
