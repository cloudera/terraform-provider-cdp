// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datalake

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/utils/test"
)

func TestFromModelToAzureRequestBasicFields(t *testing.T) {
	tags, _ := types.MapValue(types.StringType, map[string]attr.Value{"key": types.StringValue("value")})
	image := &azureDatalakeImage{
		CatalogName: types.StringValue("someCatalog"),
		ID:          types.StringValue("someId"),
		Os:          types.StringValue("someOs"),
	}
	recipes, _ := types.SetValue(types.StringType, []attr.Value{types.StringValue("recipe")})
	input := azureDatalakeResourceModel{
		DatalakeName:        types.StringValue("someClusterName"),
		EnvironmentName:     types.StringValue("someEnvironment"),
		DatabaseType:        types.StringValue("someClusterTemplateNameOrCRN"),
		ManagedIdentity:     types.StringValue("someManagedIdentity"),
		StorageLocationBase: types.StringValue("someStorageLocationBase"),
		MultiAz:             types.BoolValue(true),
		Tags:                tags,
		Recipes: []*instanceGroupRecipe{&instanceGroupRecipe{
			InstanceGroupName: types.StringValue("someSubnetID"),
			RecipeNames:       recipes,
		}},
		Runtime:                         types.StringValue("someRuntime"),
		Image:                           image,
		Scale:                           types.StringValue("someScale"),
		LoadBalancerSku:                 types.StringValue("someLoadBalancerSku"),
		EnableRangerRaz:                 types.BoolValue(true),
		JavaVersion:                     types.Int64Value(11),
		FlexibleServerDelegatedSubnetId: types.StringValue("someFlexibleServerDelegatedSubnetId"),
	}
	got := toAzureDatalakeRequest(context.TODO(), &input)

	test.CompareStrings(*got.DatalakeName, input.DatalakeName.ValueString(), t)
	test.CompareStrings(*got.EnvironmentName, input.EnvironmentName.ValueString(), t)
	test.CompareStrings(got.DatabaseType, input.DatabaseType.ValueString(), t)
	test.CompareStrings(*got.CloudProviderConfiguration.ManagedIdentity, input.ManagedIdentity.ValueString(), t)
	test.CompareStrings(*got.CloudProviderConfiguration.StorageLocation, input.StorageLocationBase.ValueString(), t)
	test.CompareBools(*got.MultiAz, input.MultiAz.ValueBool(), t)
	test.CompareStrings(*got.Tags[0].Value, input.Tags.Elements()["key"].(types.String).ValueString(), t)
	test.CompareStrings(*got.Recipes[0].InstanceGroupName, input.Recipes[0].InstanceGroupName.ValueString(), t)
	test.CompareStrings(got.Recipes[0].RecipeNames[0], input.Recipes[0].RecipeNames.Elements()[0].(types.String).ValueString(), t)
	test.CompareStrings(got.Runtime, input.Runtime.ValueString(), t)
	test.CompareStrings(*got.Image.CatalogName, input.Image.CatalogName.ValueString(), t)
	test.CompareStrings(got.Image.ID, input.Image.ID.ValueString(), t)
	test.CompareStrings(got.Image.Os, input.Image.Os.ValueString(), t)
	test.CompareStrings(string(got.Scale), input.Scale.ValueString(), t)
	test.CompareStrings(string(got.LoadBalancerSku), input.LoadBalancerSku.ValueString(), t)
	test.CompareBools(got.EnableRangerRaz, input.EnableRangerRaz.ValueBool(), t)
	test.CompareInt32PointerToTypesInt64(&got.JavaVersion, input.JavaVersion, t)
	test.CompareStrings(got.FlexibleServerDelegatedSubnetID, input.FlexibleServerDelegatedSubnetId.ValueString(), t)
}
