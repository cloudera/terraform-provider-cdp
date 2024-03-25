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
		JavaVersion:              types.Int64Value(11),
	}
	got := fromModelToAwsRequest(input, context.TODO())

	compareStrings(got.ClusterName, input.Name.ValueString(), t)
	compareStrings(got.Environment, input.Environment.ValueString(), t)
	compareStrings(got.ClusterTemplate, input.ClusterTemplate.ValueString(), t)
	compareStrings(got.SubnetID, input.SubnetId.ValueString(), t)
	compareStringSlices(got.SubnetIds, input.SubnetIds.Elements(), t)
	compareBools(got.MultiAz, input.MultiAz.ValueBool(), t)
	compareStrings(*got.Tags[0].Value, input.Tags.Elements()["key"].(types.String).ValueString(), t)
	compareStrings(got.CustomConfigurationsName, input.CustomConfigurationsName.ValueString(), t)
	compareStrings(got.Image.CatalogName, input.Image.Attributes()["catalog"].(types.String).ValueString(), t)
	compareStrings(got.Image.ID, input.Image.Attributes()["id"].(types.String).ValueString(), t)
	compareStrings(got.Image.Os, input.Image.Attributes()["os"].(types.String).ValueString(), t)
	compareStrings(got.RequestTemplate, input.RequestTemplate.ValueString(), t)
	compareStrings(string(got.DatahubDatabase), input.DatahubDatabase.ValueString(), t)
	compareStrings(got.ClusterExtension.CustomProperties, input.ClusterExtension.Attributes()["custom_properties"].(types.String).ValueString(), t)
	compareBools(got.EnableLoadBalancer, input.EnableLoadBalancer.ValueBool(), t)
	compareInt32PointerToTypesInt64(&got.JavaVersion, input.JavaVersion, t)
}

func TestFromModelToRequestRecipe(t *testing.T) {
	recipes := []types.String{types.StringValue("recipe1"), types.StringValue("recipe2")}
	igs := []InstanceGroup{{Recipes: recipes}}
	input := awsDatahubResourceModel{InstanceGroup: igs}

	got := fromModelToAwsRequest(input, context.TODO())

	compareInts(len(got.InstanceGroups), len(input.InstanceGroup), t)
	compareInts(len(got.InstanceGroups[0].RecipeNames), len(input.InstanceGroup[0].Recipes), t)

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
		VolumeSize:  types.Int64Value(100),
		VolumeCount: types.Int64Value(1),
		VolumeType:  types.StringValue("ephemeral"),
	}}
	igs := []InstanceGroup{{AttachedVolumeConfiguration: avcs}}
	input := awsDatahubResourceModel{InstanceGroup: igs}

	got := fromModelToAwsRequest(input, context.TODO())

	compareInts(len(got.InstanceGroups), len(input.InstanceGroup), t)
	compareInts(len(got.InstanceGroups[0].AttachedVolumeConfiguration), len(avcs), t)

	resultAvcs := got.InstanceGroups[0].AttachedVolumeConfiguration[0]
	compareInt32PointerToTypesInt64(resultAvcs.VolumeCount, avcs[0].VolumeCount, t)
	compareInt32PointerToTypesInt64(resultAvcs.VolumeSize, avcs[0].VolumeSize, t)
	compareStrings(*resultAvcs.VolumeType, avcs[0].VolumeType.ValueString(), t)
}

func TestFromModelToRequestInstanceGroups(t *testing.T) {
	igs := []InstanceGroup{{
		NodeCount:         types.Int64Value(1),
		InstanceGroupName: types.StringValue("gateway"),
		InstanceGroupType: types.StringValue("CORE"),
		InstanceType:      types.StringValue("m5.xlarge"),
		RootVolumeSize:    types.Int64Value(100),
		RecoveryMode:      types.StringValue("MANUAL"),
	}}

	input := awsDatahubResourceModel{InstanceGroup: igs}

	got := fromModelToAwsRequest(input, context.TODO())

	compareInts(len(got.InstanceGroups), len(igs), t)
	resultIg := got.InstanceGroups[0]
	compareStrings(*resultIg.InstanceGroupName, igs[0].InstanceGroupName.ValueString(), t)
	compareStrings(*resultIg.InstanceGroupType, igs[0].InstanceGroupType.ValueString(), t)
	compareStrings(*resultIg.InstanceType, igs[0].InstanceType.ValueString(), t)
	compareInt32PointerToTypesInt64(&resultIg.RootVolumeSize, igs[0].RootVolumeSize, t)
	compareStrings(resultIg.RecoveryMode, igs[0].RecoveryMode.ValueString(), t)
}

func TestFromModelToRequestVolumeEncryption(t *testing.T) {
	igs := []InstanceGroup{{
		VolumeEncryption: VolumeEncryption{Encryption: types.BoolValue(true)},
	}}

	input := awsDatahubResourceModel{InstanceGroup: igs}

	got := fromModelToAwsRequest(input, context.TODO())

	compareInts(len(got.InstanceGroups), len(igs), t)
	resultVolumeEncryption := got.InstanceGroups[0].VolumeEncryption
	if resultVolumeEncryption == nil {
		t.Errorf("Volume encryption is not filled though it should've been!")
	} else {
		compareBools(*resultVolumeEncryption.EnableEncryption, igs[0].VolumeEncryption.Encryption.ValueBool(), t)
	}
}

func TestFromModelToRequestClusterDefinition(t *testing.T) {
	input := awsDatahubResourceModel{ClusterDefinition: types.StringValue("SomeClusterDef")}
	got := fromModelToAwsRequest(input, context.TODO())

	compareStrings(got.ClusterDefinition, input.ClusterDefinition.ValueString(), t)
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
		JavaVersion:              types.Int64Value(11),
	}
	got := fromModelToGcpRequest(input, context.TODO())

	compareStrings(got.ClusterName, input.Name.ValueString(), t)
	compareStrings(got.EnvironmentName, input.Environment.ValueString(), t)
	compareStrings(got.ClusterTemplateName, input.ClusterTemplate.ValueString(), t)
	compareStrings(got.SubnetName, input.SubnetName.ValueString(), t)
	compareStrings(*got.Tags[0].Value, input.Tags.Elements()["key"].(types.String).ValueString(), t)
	compareStrings(got.CustomConfigurationsName, input.CustomConfigurationsName.ValueString(), t)
	compareStrings(got.Image.CatalogName, input.Image.Attributes()["catalog"].(types.String).ValueString(), t)
	compareStrings(got.Image.ID, input.Image.Attributes()["id"].(types.String).ValueString(), t)
	compareStrings(got.Image.Os, input.Image.Attributes()["os"].(types.String).ValueString(), t)
	compareStrings(got.RequestTemplate, input.RequestTemplate.ValueString(), t)
	compareStrings(string(got.DatahubDatabase), input.DatahubDatabase.ValueString(), t)
	compareStrings(got.ClusterExtension.CustomProperties, input.ClusterExtension.Attributes()["custom_properties"].(types.String).ValueString(), t)
	compareInt32PointerToTypesInt64(&got.JavaVersion, input.JavaVersion, t)
}

func TestFromModelToGcpRequestRecipe(t *testing.T) {
	recipes := []types.String{types.StringValue("recipe1"), types.StringValue("recipe2")}
	igs := []GcpInstanceGroup{{Recipes: recipes}}
	input := gcpDatahubResourceModel{InstanceGroup: igs}

	got := fromModelToGcpRequest(input, context.TODO())

	compareInts(len(got.InstanceGroups), len(input.InstanceGroup), t)
	compareInts(len(got.InstanceGroups[0].RecipeNames), len(input.InstanceGroup[0].Recipes), t)

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
		VolumeSize:  types.Int64Value(100),
		VolumeCount: types.Int64Value(1),
		VolumeType:  types.StringValue("ephemeral"),
	}}
	igs := []GcpInstanceGroup{{AttachedVolumeConfiguration: avcs}}
	input := gcpDatahubResourceModel{InstanceGroup: igs}

	got := fromModelToGcpRequest(input, context.TODO())

	compareInts(len(got.InstanceGroups), len(input.InstanceGroup), t)
	compareInts(len(got.InstanceGroups[0].AttachedVolumeConfiguration), len(avcs), t)

	resultAvcs := got.InstanceGroups[0].AttachedVolumeConfiguration[0]
	compareInt32PointerToTypesInt64(resultAvcs.VolumeCount, avcs[0].VolumeCount, t)
	compareInt32PointerToTypesInt64(resultAvcs.VolumeSize, avcs[0].VolumeSize, t)
	compareStrings(*resultAvcs.VolumeType, avcs[0].VolumeType.ValueString(), t)
}

func TestFromModelToGcpRequestInstanceGroups(t *testing.T) {
	igs := []GcpInstanceGroup{{
		NodeCount:         types.Int64Value(1),
		InstanceGroupName: types.StringValue("gateway"),
		InstanceGroupType: types.StringValue("CORE"),
		InstanceType:      types.StringValue("m5.xlarge"),
		RootVolumeSize:    types.Int64Value(100),
		RecoveryMode:      types.StringValue("MANUAL"),
	}}

	input := gcpDatahubResourceModel{InstanceGroup: igs}

	got := fromModelToGcpRequest(input, context.TODO())

	compareInts(len(got.InstanceGroups), len(igs), t)
	resultIg := got.InstanceGroups[0]
	compareStrings(*resultIg.InstanceGroupName, igs[0].InstanceGroupName.ValueString(), t)
	compareStrings(*resultIg.InstanceGroupType, igs[0].InstanceGroupType.ValueString(), t)
	compareStrings(*resultIg.InstanceType, igs[0].InstanceType.ValueString(), t)
	compareInt32PointerToTypesInt64(resultIg.RootVolumeSize, igs[0].RootVolumeSize, t)
	compareStrings(resultIg.RecoveryMode, igs[0].RecoveryMode.ValueString(), t)
}

func TestFromModelToGcpRequestClusterDefinition(t *testing.T) {
	input := gcpDatahubResourceModel{ClusterDefinition: types.StringValue("SomeClusterDef")}
	got := fromModelToGcpRequest(input, context.TODO())

	compareStrings(got.ClusterDefinitionName, input.ClusterDefinition.ValueString(), t)
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
		JavaVersion:                     types.Int64Value(11),
		FlexibleServerDelegatedSubnetId: types.StringValue("someFlexibleServerDelegatedSubnetId"),
	}
	got := fromModelToAzureRequest(input, context.TODO())

	compareStrings(got.ClusterName, input.Name.ValueString(), t)
	compareStrings(got.EnvironmentName, input.Environment.ValueString(), t)
	compareStrings(got.ClusterTemplateName, input.ClusterTemplate.ValueString(), t)
	compareStrings(got.SubnetID, input.SubnetId.ValueString(), t)
	compareBools(*got.MultiAz, input.MultiAz.ValueBool(), t)
	compareStrings(*got.Tags[0].Value, input.Tags.Elements()["key"].(types.String).ValueString(), t)
	compareStrings(got.CustomConfigurationsName, input.CustomConfigurationsName.ValueString(), t)
	compareStrings(got.Image.CatalogName, input.Image.Attributes()["catalog"].(types.String).ValueString(), t)
	compareStrings(got.Image.ID, input.Image.Attributes()["id"].(types.String).ValueString(), t)
	compareStrings(got.Image.Os, input.Image.Attributes()["os"].(types.String).ValueString(), t)
	compareStrings(got.RequestTemplate, input.RequestTemplate.ValueString(), t)
	compareStrings(string(got.DatahubDatabase), input.DatahubDatabase.ValueString(), t)
	compareStrings(got.ClusterExtension.CustomProperties, input.ClusterExtension.Attributes()["custom_properties"].(types.String).ValueString(), t)
	compareBools(got.EnableLoadBalancer, input.EnableLoadBalancer.ValueBool(), t)
	compareInt32PointerToTypesInt64(&got.JavaVersion, input.JavaVersion, t)
	compareStrings(got.FlexibleServerDelegatedSubnetID, input.FlexibleServerDelegatedSubnetId.ValueString(), t)
}

func TestFromModelToAzureRequestRecipe(t *testing.T) {
	recipes := []types.String{types.StringValue("recipe1"), types.StringValue("recipe2")}
	igs := []InstanceGroup{{Recipes: recipes}}
	input := azureDatahubResourceModel{InstanceGroup: igs}

	got := fromModelToAzureRequest(input, context.TODO())

	compareInts(len(got.InstanceGroups), len(input.InstanceGroup), t)
	compareInts(len(got.InstanceGroups[0].RecipeNames), len(input.InstanceGroup[0].Recipes), t)

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
		VolumeSize:  types.Int64Value(100),
		VolumeCount: types.Int64Value(1),
		VolumeType:  types.StringValue("ephemeral"),
	}}
	igs := []InstanceGroup{{AttachedVolumeConfiguration: avcs}}
	input := azureDatahubResourceModel{InstanceGroup: igs}

	got := fromModelToAzureRequest(input, context.TODO())

	compareInts(len(got.InstanceGroups), len(input.InstanceGroup), t)
	compareInts(len(got.InstanceGroups[0].AttachedVolumeConfiguration), len(avcs), t)

	resultAvcs := got.InstanceGroups[0].AttachedVolumeConfiguration[0]
	compareInt32PointerToTypesInt64(resultAvcs.VolumeCount, avcs[0].VolumeCount, t)
	compareInt32PointerToTypesInt64(resultAvcs.VolumeSize, avcs[0].VolumeSize, t)
	compareStrings(*resultAvcs.VolumeType, avcs[0].VolumeType.ValueString(), t)
}

func TestFromModelToAzureRequestInstanceGroups(t *testing.T) {
	igs := []InstanceGroup{{
		NodeCount:         types.Int64Value(1),
		InstanceGroupName: types.StringValue("gateway"),
		InstanceGroupType: types.StringValue("CORE"),
		InstanceType:      types.StringValue("m5.xlarge"),
		RootVolumeSize:    types.Int64Value(100),
		RecoveryMode:      types.StringValue("MANUAL"),
	}}

	input := azureDatahubResourceModel{InstanceGroup: igs}

	got := fromModelToAzureRequest(input, context.TODO())

	compareInts(len(got.InstanceGroups), len(igs), t)
	resultIg := got.InstanceGroups[0]
	compareStrings(*resultIg.InstanceGroupName, igs[0].InstanceGroupName.ValueString(), t)
	compareStrings(*resultIg.InstanceGroupType, igs[0].InstanceGroupType.ValueString(), t)
	compareStrings(*resultIg.InstanceType, igs[0].InstanceType.ValueString(), t)
	compareInt32PointerToTypesInt64(resultIg.RootVolumeSize, igs[0].RootVolumeSize, t)
	compareStrings(resultIg.RecoveryMode, igs[0].RecoveryMode.ValueString(), t)
}

func TestFromModelToAzureRequestClusterDefinition(t *testing.T) {
	input := azureDatahubResourceModel{ClusterDefinition: types.StringValue("SomeClusterDef")}
	got := fromModelToAzureRequest(input, context.TODO())

	compareStrings(got.ClusterDefinitionName, input.ClusterDefinition.ValueString(), t)
}

func compareStrings(got string, expected string, t *testing.T) {
	if got != expected {
		t.Errorf("Assertion error! Expected: %s, got: %s", expected, got)
	}
}

func compareStringSlices(got []string, expected []attr.Value, t *testing.T) {
	if len(got) != len(expected) {
		t.Errorf("Assertion error! Expected length: %d, got length: %d", len(expected), len(got))
		return
	}

	for i, exp := range expected {
		if got[i] != exp.(types.String).ValueString() {
			t.Errorf("Assertion error! Expected: %s, got: %s", expected, got)
		}
	}
}

func compareInt32PointerToTypesInt64(got *int32, expected types.Int64, t *testing.T) {
	if *got != *int64To32Pointer(expected) {
		t.Errorf("Assertion error! Expected: %d, got: %d", expected.ValueInt64(), *got)
	}
}

func compareInts(got int, expected int, t *testing.T) {
	if got != expected {
		t.Errorf("Assertion error! Expected: %d, got: %d", expected, got)
	}
}

func compareBools(got bool, expected bool, t *testing.T) {
	if got != expected {
		t.Errorf("Assertion error! Expected: %t, got: %t", expected, got)
	}
}
