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
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	opdbmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/models"
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

	req.JavaVersion = int64To32(model.JavaVersion)
	req.NumEdgeNodes = int64To32(model.NumEdgeNodes)

	if model.AutoScalingParameters != nil {
		tflog.Info(ctx, fmt.Sprintf("Autoscaling parameters %+v.", model.AutoScalingParameters))
		req.AutoScalingParameters = createAutoScalingParameters(*model.AutoScalingParameters, ctx)
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

	tflog.Debug(ctx, fmt.Sprintf("Conversion from databaseResourceModel to CreateDatabaseRequest has finished with request: %+v.", req))
	return &req
}

func createAutoScalingParameters(autoScalingParameters AutoScalingParametersStruct, ctx context.Context) *opdbmodels.AutoScalingParameters {
	return &opdbmodels.AutoScalingParameters{
		TargetedValueForMetric: autoScalingParameters.TargetedValueForMetric.ValueInt64(),
		MaxWorkersForDatabase:  int64To32(autoScalingParameters.MaxWorkersForDatabase),
		MaxWorkersPerBatch:     int64To32(autoScalingParameters.MaxWorkersPerBatch),
		MinWorkersForDatabase:  int64To32(autoScalingParameters.MinWorkersForDatabase),
		EvaluationPeriod:       autoScalingParameters.EvaluationPeriod.ValueInt64(),
		MinimumBlockCacheGb:    int64To32(autoScalingParameters.MinimumBlockCacheGb),

		MaxCPUUtilization:          int64To32(autoScalingParameters.MaxCPUUtilization),
		MaxComputeNodesForDatabase: int64To32Pointer(autoScalingParameters.MaxComputeNodesForDatabase),
		MinComputeNodesForDatabase: int64To32Pointer(autoScalingParameters.MinComputeNodesForDatabase),
		MaxHdfsUsagePercentage:     int64To32(autoScalingParameters.MaxHdfsUsagePercentage),
		MaxRegionsPerRegionServer:  int64To32(autoScalingParameters.MaxRegionsPerRegionServer),
	}
}

func createAttachedStorageForWorkers(attachedStorageForWorkers AttachedStorageForWorkersStruct, ctx context.Context) *opdbmodels.AttachedStorageForWorkers {
	return &opdbmodels.AttachedStorageForWorkers{
		VolumeCount: int64To32(attachedStorageForWorkers.VolumeCount),
		VolumeSize:  int64To32(attachedStorageForWorkers.VolumeSize),
		VolumeType:  opdbmodels.VolumeType(attachedStorageForWorkers.VolumeType.ValueString()),
	}
}

func createImage(image Image, ctx context.Context) *opdbmodels.Image {
	return &opdbmodels.Image{
		ID:      image.ID.ValueStringPointer(),
		Catalog: image.Catalog.ValueStringPointer(),
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
		tflog.Debug(ctx, fmt.Sprintf("Converting KeyValuePair: %+v.", vrs))
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

func int64To32Pointer(in types.Int64) *int32 {
	n64 := in.ValueInt64()
	var n2 = int32(n64)
	return &n2
}

func int64To32(in types.Int64) int32 {
	n64 := in.ValueInt64()
	return int32(n64)
}
