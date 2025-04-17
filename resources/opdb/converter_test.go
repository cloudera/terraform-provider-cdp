// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package opdb

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/utils/test"
)

func TestFromSimplestModelToDatabaseRequestBasicFields(t *testing.T) {
	input := databaseResourceModel{
		DatabaseName: types.StringValue("someName"),
		Environment:  types.StringValue("someEnvironment"),
	}
	got := fromModelToDatabaseRequest(input, context.TODO())

	test.CompareStrings(*got.DatabaseName, input.DatabaseName.ValueString(), t)
	test.CompareStrings(*got.EnvironmentName, input.Environment.ValueString(), t)

	// Default false values
	test.CompareBools(got.DisableExternalDB, false, t)
	test.CompareBools(got.DisableMultiAz, false, t)
	test.CompareBools(got.DisableKerberos, false, t)
	test.CompareBools(got.DisableJwtAuth, false, t)
	test.CompareBools(got.EnableGrafana, false, t)
	test.CompareBools(got.EnableRegionCanary, false, t)
}

func TestFromModelToDatabaseRequestMoreFields(t *testing.T) {
	input := databaseResourceModel{
		DatabaseName:      types.StringValue("someName"),
		Environment:       types.StringValue("someEnvironment"),
		ScaleType:         types.StringValue("MICRO"),
		StorageType:       types.StringValue("SSD"),
		DisableExternalDB: types.BoolValue(true),
		DisableMultiAz:    types.BoolValue(true),
		SubnetID:          types.StringValue("someSubnetID"),
		JavaVersion:       types.Int32Value(11),
		NumEdgeNodes:      types.Int32Value(2),

		DisableKerberos:    types.BoolValue(true),
		DisableJwtAuth:     types.BoolValue(true),
		EnableGrafana:      types.BoolValue(true),
		EnableRegionCanary: types.BoolValue(true),
		StorageLocation:    types.StringValue("someStorageLocation"),
		Architecture:       types.StringValue("ARM64"),
	}
	got := fromModelToDatabaseRequest(input, context.TODO())

	test.CompareStrings(*got.DatabaseName, input.DatabaseName.ValueString(), t)
	test.CompareStrings(*got.EnvironmentName, input.Environment.ValueString(), t)
	test.CompareStrings(string(got.ScaleType), input.ScaleType.ValueString(), t)
	test.CompareStrings(string(got.StorageType), input.StorageType.ValueString(), t)

	test.CompareBools(got.DisableExternalDB, input.DisableExternalDB.ValueBool(), t)
	test.CompareBools(got.DisableMultiAz, input.DisableMultiAz.ValueBool(), t)

	test.CompareStrings(got.SubnetID, input.SubnetID.ValueString(), t)

	test.CompareInt32PointerToTypesInt32(&got.JavaVersion, input.JavaVersion, t)
	test.CompareInt32PointerToTypesInt32(&got.NumEdgeNodes, input.NumEdgeNodes, t)

	test.CompareBools(got.DisableKerberos, input.DisableKerberos.ValueBool(), t)
	test.CompareBools(got.DisableJwtAuth, input.DisableJwtAuth.ValueBool(), t)
	test.CompareBools(got.EnableGrafana, input.EnableGrafana.ValueBool(), t)
	test.CompareBools(got.EnableRegionCanary, input.EnableRegionCanary.ValueBool(), t)

	test.CompareStrings(got.StorageLocation, input.StorageLocation.ValueString(), t)
	test.CompareStrings(string(got.Architecture), input.Architecture.ValueString(), t)
}

func TestFromModelToUpdateDatabaseRequestAutoScaling(t *testing.T) {
	autoScalingParameters := AutoScalingParametersStruct{
		TargetedValueForMetric: types.Int64Value(234),
		MaxWorkersForDatabase:  types.Int32Value(4),
		MaxWorkersPerBatch:     types.Int32Value(2),
		MinWorkersForDatabase:  types.Int32Value(3),
		EvaluationPeriod:       types.Int64Value(2400),
		MinimumBlockCacheGb:    types.Int32Value(1),

		MaxCPUUtilization:          types.Int32Value(-1),
		MaxComputeNodesForDatabase: types.Int32Value(-1),
		MinComputeNodesForDatabase: types.Int32Value(-1),
		MaxHdfsUsagePercentage:     types.Int32Value(80),
		MaxRegionsPerRegionServer:  types.Int32Value(200),
	}

	input := databaseResourceModel{
		DatabaseName:          types.StringValue("someName"),
		Environment:           types.StringValue("someEnvironment"),
		AutoScalingParameters: &autoScalingParameters,
	}
	got := fromModelToUpdateDatabaseRequest(input, context.TODO())

	test.CompareStrings(*got.DatabaseName, input.DatabaseName.ValueString(), t)
	test.CompareStrings(*got.EnvironmentName, input.Environment.ValueString(), t)

	gotAutoscalingParametersRequest := *got.AutoScalingParameters

	test.CompareInt64PointerToTypesInt64(&gotAutoscalingParametersRequest.TargetedValueForMetric, autoScalingParameters.TargetedValueForMetric, t)
	test.CompareInt32PointerToTypesInt32(&gotAutoscalingParametersRequest.MaxWorkersForDatabase, autoScalingParameters.MaxWorkersForDatabase, t)
	test.CompareInt32PointerToTypesInt32(&gotAutoscalingParametersRequest.MaxWorkersPerBatch, autoScalingParameters.MaxWorkersPerBatch, t)
	test.CompareInt32PointerToTypesInt32(&gotAutoscalingParametersRequest.MinWorkersForDatabase, autoScalingParameters.MinWorkersForDatabase, t)
	test.CompareInt64PointerToTypesInt64(&gotAutoscalingParametersRequest.EvaluationPeriod, autoScalingParameters.EvaluationPeriod, t)
	test.CompareInt32PointerToTypesInt32(&gotAutoscalingParametersRequest.MinimumBlockCacheGb, autoScalingParameters.MinimumBlockCacheGb, t)

	test.CompareInt32PointerToTypesInt32(&gotAutoscalingParametersRequest.MaxCPUUtilization, autoScalingParameters.MaxCPUUtilization, t)
	test.CompareInt32PointerToTypesInt32(gotAutoscalingParametersRequest.MaxComputeNodesForDatabase, autoScalingParameters.MaxComputeNodesForDatabase, t)
	test.CompareInt32PointerToTypesInt32(gotAutoscalingParametersRequest.MinComputeNodesForDatabase, autoScalingParameters.MinComputeNodesForDatabase, t)
	test.CompareInt32PointerToTypesInt32(&gotAutoscalingParametersRequest.MaxHdfsUsagePercentage, autoScalingParameters.MaxHdfsUsagePercentage, t)
	test.CompareInt32PointerToTypesInt32(&gotAutoscalingParametersRequest.MaxRegionsPerRegionServer, autoScalingParameters.MaxRegionsPerRegionServer, t)

}

func TestFromModelToUpdateDatabaseRequestImage(t *testing.T) {
	image := Image{
		ID:      types.StringValue("someID"),
		Catalog: types.StringValue("someCatalog"),
	}

	input := databaseResourceModel{
		DatabaseName: types.StringValue("someName"),
		Environment:  types.StringValue("someEnvironment"),
		Image:        &image,
	}
	got := fromModelToUpdateDatabaseRequest(input, context.TODO())

	test.CompareStrings(*got.DatabaseName, input.DatabaseName.ValueString(), t)
	test.CompareStrings(*got.EnvironmentName, input.Environment.ValueString(), t)
	test.CompareStrings(got.Catalog, input.Image.Catalog.ValueString(), t)
}

func TestCreateAutoScalingParams(t *testing.T) {
	autoScalingParameters := AutoScalingParametersStruct{
		TargetedValueForMetric: types.Int64Value(234),
		MaxWorkersForDatabase:  types.Int32Value(4),
		MaxWorkersPerBatch:     types.Int32Value(2),
		MinWorkersForDatabase:  types.Int32Value(3),
		EvaluationPeriod:       types.Int64Value(2400),
		MinimumBlockCacheGb:    types.Int32Value(1),

		MaxCPUUtilization:          types.Int32Value(-1),
		MaxComputeNodesForDatabase: types.Int32Value(-1),
		MinComputeNodesForDatabase: types.Int32Value(-1),
		MaxHdfsUsagePercentage:     types.Int32Value(80),
		MaxRegionsPerRegionServer:  types.Int32Value(200),
	}

	got := createAutoScalingParameters(autoScalingParameters)

	test.CompareInt64PointerToTypesInt64(&got.TargetedValueForMetric, autoScalingParameters.TargetedValueForMetric, t)
	test.CompareInt32PointerToTypesInt32(&got.MaxWorkersForDatabase, autoScalingParameters.MaxWorkersForDatabase, t)
	test.CompareInt32PointerToTypesInt32(&got.MaxWorkersPerBatch, autoScalingParameters.MaxWorkersPerBatch, t)
	test.CompareInt32PointerToTypesInt32(&got.MinWorkersForDatabase, autoScalingParameters.MinWorkersForDatabase, t)
	test.CompareInt64PointerToTypesInt64(&got.EvaluationPeriod, autoScalingParameters.EvaluationPeriod, t)
	test.CompareInt32PointerToTypesInt32(&got.MinimumBlockCacheGb, autoScalingParameters.MinimumBlockCacheGb, t)

	test.CompareInt32PointerToTypesInt32(&got.MaxCPUUtilization, autoScalingParameters.MaxCPUUtilization, t)
	test.CompareInt32PointerToTypesInt32(got.MaxComputeNodesForDatabase, autoScalingParameters.MaxComputeNodesForDatabase, t)
	test.CompareInt32PointerToTypesInt32(got.MinComputeNodesForDatabase, autoScalingParameters.MinComputeNodesForDatabase, t)
	test.CompareInt32PointerToTypesInt32(&got.MaxHdfsUsagePercentage, autoScalingParameters.MaxHdfsUsagePercentage, t)
	test.CompareInt32PointerToTypesInt32(&got.MaxRegionsPerRegionServer, autoScalingParameters.MaxRegionsPerRegionServer, t)
}

func Int64PointerTo32Pointer(in *int64) *int32 {
	var n2 = int32(*in)
	return &n2
}

func TestCreateAttachedStorageForWorkers(t *testing.T) {
	attachedStorageForWorkers := AttachedStorageForWorkersStruct{
		VolumeCount: types.Int32Value(2),
		VolumeSize:  types.Int32Value(2024),
		VolumeType:  types.StringValue("LOCAL_SSD"),
	}

	got := createAttachedStorageForWorkers(attachedStorageForWorkers, context.TODO())

	test.CompareInt32PointerToTypesInt32(&got.VolumeCount, attachedStorageForWorkers.VolumeCount, t)
	test.CompareInt32PointerToTypesInt32(&got.VolumeSize, attachedStorageForWorkers.VolumeSize, t)
	test.CompareStrings(string(got.VolumeType), attachedStorageForWorkers.VolumeType.ValueString(), t)
}

func TestCreateImage(t *testing.T) {
	image := Image{
		ID:      types.StringValue("FOO"),
		Catalog: types.StringValue("BAR"),
	}

	got := createImage(image, context.TODO())

	test.CompareStrings(got.ID, image.ID.ValueString(), t)
	test.CompareStrings(got.Catalog, image.Catalog.ValueString(), t)
}

func TestCreateCustomUserTags(t *testing.T) {
	var tags []KeyValuePair

	a := KeyValuePair{
		Key:   types.StringValue("k1"),
		Value: types.StringValue("v1"),
	}
	b := KeyValuePair{
		Key:   types.StringValue("k2"),
		Value: types.StringValue("v2"),
	}
	tags = append(tags, a)
	tags = append(tags, b)

	got := createCustomUserTags(context.TODO(), tags)

	test.CompareStrings(got[0].Key, a.Key.ValueString(), t)
	test.CompareStrings(got[0].Value, a.Value.ValueString(), t)

	test.CompareStrings(got[1].Key, b.Key.ValueString(), t)
	test.CompareStrings(got[1].Value, b.Value.ValueString(), t)
}

func TestCreateRecipes(t *testing.T) {
	var recipes []Recipe
	recipesA, _ := types.SetValue(types.StringType, []attr.Value{types.StringValue("recipeA1"), types.StringValue("recipeA2")})
	recipesB, _ := types.SetValue(types.StringType, []attr.Value{types.StringValue("recipeB1"), types.StringValue("recipeB2")})

	a := Recipe{
		Names:         recipesA,
		InstanceGroup: types.StringValue("i1"),
	}
	b := Recipe{
		Names:         recipesB,
		InstanceGroup: types.StringValue("i2"),
	}
	recipes = append(recipes, a)
	recipes = append(recipes, b)

	got := createRecipes(context.TODO(), recipes)

	test.CompareStrings(string(*got[0].InstanceGroup), a.InstanceGroup.ValueString(), t)
	test.CompareStringValueSlices(got[0].Names, a.Names.Elements(), t)
	test.CompareStrings(string(*got[1].InstanceGroup), b.InstanceGroup.ValueString(), t)
	test.CompareStringValueSlices(got[1].Names, b.Names.Elements(), t)
}

func TestVolumeEncryptions(t *testing.T) {
	var volumeEncryptions []VolumeEncryption

	a := VolumeEncryption{
		EncryptionKey: types.StringValue("k1"),
		InstanceGroup: types.StringValue("i1"),
	}
	b := VolumeEncryption{
		EncryptionKey: types.StringValue("k2"),
		InstanceGroup: types.StringValue("i2"),
	}
	volumeEncryptions = append(volumeEncryptions, a)
	volumeEncryptions = append(volumeEncryptions, b)

	got := createVolumeEncryptions(context.TODO(), volumeEncryptions)

	test.CompareStrings(string(*got[0].InstanceGroup), a.InstanceGroup.ValueString(), t)
	test.CompareStrings(string(*got[0].EncryptionKey), a.EncryptionKey.ValueString(), t)
	test.CompareStrings(string(*got[1].InstanceGroup), b.InstanceGroup.ValueString(), t)
	test.CompareStrings(string(*got[1].EncryptionKey), b.EncryptionKey.ValueString(), t)
}
