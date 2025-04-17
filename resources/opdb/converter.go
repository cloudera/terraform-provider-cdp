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
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	opdbmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

func fromModelToDatabaseRequest(model databaseResourceModel, ctx context.Context) *opdbmodels.CreateDatabaseRequest {
	tflog.Debug(ctx, "Conversion from databaseResourceModel to CreateDatabaseRequest started.")
	req := opdbmodels.CreateDatabaseRequest{}
	req.DatabaseName = model.DatabaseName.ValueStringPointer()
	req.EnvironmentName = model.Environment.ValueStringPointer()
	req.ScaleType = opdbmodels.ScaleType(model.ScaleType.ValueString())
	req.StorageType = opdbmodels.StorageType(model.StorageType.ValueString())
	req.DisableExternalDB = model.DisableExternalDB.ValueBool()

	req.DisableMultiAz = model.DisableMultiAz.ValueBool()
	req.SubnetID = model.SubnetID.ValueString()

	req.JavaVersion = int32TypesTo32(model.JavaVersion)
	req.NumEdgeNodes = int32TypesTo32(model.NumEdgeNodes)

	if model.AutoScalingParameters != nil {
		tflog.Info(ctx, fmt.Sprintf("Autoscaling parameters %+v.", model.AutoScalingParameters))
		req.AutoScalingParameters = createAutoScalingParameters(*model.AutoScalingParameters)
	}

	if model.AttachedStorageForWorkers != nil {
		req.AttachedStorageForWorkers = createAttachedStorageForWorkers(*model.AttachedStorageForWorkers, ctx)
	}

	req.DisableKerberos = model.DisableKerberos.ValueBool()
	req.DisableJwtAuth = model.DisableJwtAuth.ValueBool()

	if model.Image != nil {
		req.Image = createImage(*model.Image, ctx)
	}

	req.EnableGrafana = model.EnableGrafana.ValueBool()

	req.CustomUserTags = createCustomUserTags(ctx, model.CustomUserTags)
	req.EnableRegionCanary = model.EnableRegionCanary.ValueBool()

	req.Recipes = createRecipes(ctx, model.Recipes)
	req.StorageLocation = model.StorageLocation.ValueString()

	req.VolumeEncryptions = createVolumeEncryptions(ctx, model.VolumeEncryptions)

	req.Architecture = opdbmodels.ArchitectureType(model.Architecture.ValueString())

	tflog.Debug(ctx, fmt.Sprintf("Conversion from databaseResourceModel to CreateDatabaseRequest has finished with request: %+v.", req))
	return &req
}

func fromModelToUpdateDatabaseRequest(model databaseResourceModel, ctx context.Context) *opdbmodels.UpdateDatabaseRequest {
	tflog.Debug(ctx, "Conversion from databaseResourceModel to UpdateDatabaseRequest started.")
	req := opdbmodels.UpdateDatabaseRequest{}
	req.DatabaseName = model.DatabaseName.ValueStringPointer()
	req.EnvironmentName = model.Environment.ValueStringPointer()

	if model.AutoScalingParameters != nil {
		tflog.Info(ctx, fmt.Sprintf("Autoscaling parameters %+v.", model.AutoScalingParameters))
		req.AutoScalingParameters = createAutoScalingParameters(*model.AutoScalingParameters)
	}

	if model.Image != nil {
		req.Catalog = *model.Image.Catalog.ValueStringPointer()
	}

	tflog.Debug(ctx, fmt.Sprintf("Conversion from databaseResourceModel to UpdateDatabaseRequest has finished with request: %+v.", req))
	return &req
}

func createAutoScalingParameters(autoScalingParameters AutoScalingParametersStruct) *opdbmodels.AutoScalingParameters {
	return &opdbmodels.AutoScalingParameters{
		TargetedValueForMetric: autoScalingParameters.TargetedValueForMetric.ValueInt64(),
		MaxWorkersForDatabase:  int32TypesTo32(autoScalingParameters.MaxWorkersForDatabase),
		MaxWorkersPerBatch:     int32TypesTo32(autoScalingParameters.MaxWorkersPerBatch),
		MinWorkersForDatabase:  int32TypesTo32(autoScalingParameters.MinWorkersForDatabase),
		EvaluationPeriod:       autoScalingParameters.EvaluationPeriod.ValueInt64(),
		MinimumBlockCacheGb:    int32TypesTo32(autoScalingParameters.MinimumBlockCacheGb),

		MaxCPUUtilization:          int32TypesTo32(autoScalingParameters.MaxCPUUtilization),
		MaxComputeNodesForDatabase: int32TypesTo32Pointer(autoScalingParameters.MaxComputeNodesForDatabase),
		MinComputeNodesForDatabase: int32TypesTo32Pointer(autoScalingParameters.MinComputeNodesForDatabase),
		MaxHdfsUsagePercentage:     int32TypesTo32(autoScalingParameters.MaxHdfsUsagePercentage),
		MaxRegionsPerRegionServer:  int32TypesTo32(autoScalingParameters.MaxRegionsPerRegionServer),
	}
}

func createAttachedStorageForWorkers(attachedStorageForWorkers AttachedStorageForWorkersStruct, ctx context.Context) *opdbmodels.AttachedStorageForWorkers {
	return &opdbmodels.AttachedStorageForWorkers{
		VolumeCount: int32TypesTo32(attachedStorageForWorkers.VolumeCount),
		VolumeSize:  int32TypesTo32(attachedStorageForWorkers.VolumeSize),
		VolumeType:  opdbmodels.VolumeType(attachedStorageForWorkers.VolumeType.ValueString()),
	}
}

func createImage(image Image, ctx context.Context) *opdbmodels.Image {
	return &opdbmodels.Image{
		ID:      image.ID.ValueString(),
		Catalog: image.Catalog.ValueString(),
	}
}

func createCustomUserTags(ctx context.Context, keyValuePairs []KeyValuePair) []*opdbmodels.KeyValuePair {
	var kvList []*opdbmodels.KeyValuePair
	for _, vrs := range keyValuePairs {
		tflog.Debug(ctx, fmt.Sprintf("Converting KeyValuePair: %+v.", vrs))
		kvList = append(kvList, createKeyValuePair(vrs))
	}
	return kvList
}

func createKeyValuePair(keyValuePair KeyValuePair) *opdbmodels.KeyValuePair {
	return &opdbmodels.KeyValuePair{
		Key:   keyValuePair.Key.ValueString(),
		Value: keyValuePair.Value.ValueString(),
	}
}

func createRecipes(ctx context.Context, recipes []Recipe) []*opdbmodels.CustomRecipe {
	var recipeList []*opdbmodels.CustomRecipe
	for _, vrs := range recipes {
		tflog.Debug(ctx, fmt.Sprintf("Converting Recipe: %+v.", vrs))
		recipeList = append(recipeList, createRecipe(vrs))
	}
	return recipeList
}

func createRecipe(customRecipe Recipe) *opdbmodels.CustomRecipe {
	return &opdbmodels.CustomRecipe{
		InstanceGroup: opdbmodels.NewInstanceGroupType(opdbmodels.InstanceGroupType(customRecipe.InstanceGroup.ValueString())),
		Names:         utils.FromSetValueToStringList(customRecipe.Names),
	}
}

func createVolumeEncryptions(ctx context.Context, recipes []VolumeEncryption) []*opdbmodels.VolumeEncryption {
	var volumeEncryptionList []*opdbmodels.VolumeEncryption
	for _, vrs := range recipes {
		tflog.Debug(ctx, fmt.Sprintf("Converting VolumeEncryption: %+v.", vrs))
		volumeEncryptionList = append(volumeEncryptionList, createVolumeEncryption(vrs))
	}
	return volumeEncryptionList
}

func createVolumeEncryption(volumeEncryption VolumeEncryption) *opdbmodels.VolumeEncryption {
	return &opdbmodels.VolumeEncryption{
		EncryptionKey: volumeEncryption.EncryptionKey.ValueStringPointer(),
		InstanceGroup: opdbmodels.NewInstanceGroupType(opdbmodels.InstanceGroupType(volumeEncryption.InstanceGroup.ValueString())),
	}
}

func int32TypesTo32Pointer(in types.Int32) *int32 {
	n64 := in.ValueInt32()
	var n2 = n64
	return &n2
}

func int32TypesTo32(in types.Int32) int32 {
	n64 := in.ValueInt32()
	return n64
}
