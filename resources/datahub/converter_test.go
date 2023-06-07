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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"testing"
)

func TestFromModelToRequestBasicFields(t *testing.T) {
	input := awsDatahubResourceModel{
		Name:            types.StringValue("someClusterName"),
		Environment:     types.StringValue("someEnvironment"),
		ClusterTemplate: types.StringValue("someClusterTemplateNameOrCRN"),
	}
	got := fromModelToRequest(input, nil)

	compareStrings(got.ClusterName, input.Name.ValueString(), t)
	compareStrings(got.Environment, input.Environment.ValueString(), t)
	compareStrings(got.ClusterTemplate, input.ClusterTemplate.ValueString(), t)
}

func TestFromModelToRequestRecipe(t *testing.T) {
	recipes := []types.String{types.StringValue("recipe1"), types.StringValue("recipe2")}
	igs := []InstanceGroup{{Recipes: recipes}}
	input := awsDatahubResourceModel{InstanceGroup: igs}

	got := fromModelToRequest(input, nil)

	compareBasicInts(len(got.InstanceGroups), len(input.InstanceGroup), t)
	compareBasicInts(len(got.InstanceGroups[0].RecipeNames), len(input.InstanceGroup[0].Recipes), t)

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

	got := fromModelToRequest(input, nil)

	compareBasicInts(len(got.InstanceGroups), len(input.InstanceGroup), t)
	compareBasicInts(len(got.InstanceGroups[0].AttachedVolumeConfiguration), len(avcs), t)

	resultAvcs := got.InstanceGroups[0].AttachedVolumeConfiguration[0]
	compareInts(resultAvcs.VolumeCount, avcs[0].VolumeCount, t)
	compareInts(resultAvcs.VolumeSize, avcs[0].VolumeSize, t)
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

	got := fromModelToRequest(input, nil)

	compareBasicInts(len(got.InstanceGroups), len(igs), t)
	resultIg := got.InstanceGroups[0]
	compareStrings(*resultIg.InstanceGroupName, igs[0].InstanceGroupName.ValueString(), t)
	compareStrings(*resultIg.InstanceGroupType, igs[0].InstanceGroupType.ValueString(), t)
	compareStrings(*resultIg.InstanceType, igs[0].InstanceType.ValueString(), t)
	compareInts(&resultIg.RootVolumeSize, igs[0].RootVolumeSize, t)
	compareStrings(resultIg.RecoveryMode, igs[0].RecoveryMode.ValueString(), t)
}

func TestFromModelToRequestVolumeEncryption(t *testing.T) {
	igs := []InstanceGroup{{
		VolumeEncryption: VolumeEncryption{Encryption: types.BoolValue(true)},
	}}

	input := awsDatahubResourceModel{InstanceGroup: igs}

	got := fromModelToRequest(input, nil)

	compareBasicInts(len(got.InstanceGroups), len(igs), t)
	resultVolumeEncryption := got.InstanceGroups[0].VolumeEncryption
	if resultVolumeEncryption == nil {
		t.Errorf("Volume encryption is not filled though it should've been!")
	} else {
		compareBools(*resultVolumeEncryption.EnableEncryption, igs[0].VolumeEncryption.Encryption.ValueBool(), t)
	}
}

func TestFromModelToRequestClusterDefinition(t *testing.T) {
	input := awsDatahubResourceModel{ClusterDefinition: types.StringValue("SomeClusterDef")}
	got := fromModelToRequest(input, nil)

	compareStrings(got.ClusterDefinition, input.ClusterDefinition.ValueString(), t)
}

func compareStrings(got string, expected string, t *testing.T) {
	if got != expected {
		t.Errorf("Assertion error! Expected: %s, got: %s", expected, got)
	}
}

func compareInts(got *int32, expected types.Int64, t *testing.T) {
	if *got != *int64To32Pointer(expected) {
		t.Errorf("Assertion error! Expected: %d, got: %d", expected.ValueInt64(), *got)
	}
}

func compareBasicInts(got int, expected int, t *testing.T) {
	if got != expected {
		t.Errorf("Assertion error! Expected: %d, got: %d", expected, got)
	}
}

func compareBools(got bool, expected bool, t *testing.T) {
	if got != expected {
		t.Errorf("Assertion error! Expected: %t, got: %t", expected, got)
	}
}
