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
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	datahubmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
)

func fromModelToAwsRequest(model awsDatahubResourceModel, ctx context.Context) *datahubmodels.CreateAWSClusterRequest {
	tflog.Debug(ctx, "Conversion from datahubResourceModel to CreateAWSClusterRequest started.")
	req := datahubmodels.CreateAWSClusterRequest{}
	req.ClusterName = model.Name.ValueString()
	req.ClusterTemplate = model.ClusterTemplate.ValueString()
	req.Environment = model.Environment.ValueString()
	req.ClusterDefinition = model.ClusterDefinition.ValueString()
	var igs []*datahubmodels.InstanceGroupRequest
	tflog.Debug(ctx, fmt.Sprintf("%d instance group found in the input model.", len(model.InstanceGroup)))
	for _, group := range model.InstanceGroup {
		tflog.Debug(ctx, fmt.Sprintf("Converting InstanceGroupRequest: %+v.", group))
		var volReqs []*datahubmodels.AttachedVolumeRequest
		tflog.Debug(ctx, fmt.Sprintf("%d attached volume request found in the input model.", len(model.InstanceGroup)))
		for _, vrs := range group.AttachedVolumeConfiguration {
			tflog.Debug(ctx, fmt.Sprintf("Converting AttachedVolumeConfiguration: %+v.", vrs))
			volReqs = append(volReqs, createAttachedVolumeRequest(vrs))
		}
		var igRecipes []string
		if group.Recipes != nil && len(group.Recipes) > 0 {
			for _, recipe := range group.Recipes {
				igRecipes = append(igRecipes, recipe.ValueString())
			}
		}
		ig := &datahubmodels.InstanceGroupRequest{
			AttachedVolumeConfiguration: volReqs,
			InstanceGroupName:           group.InstanceGroupName.ValueStringPointer(),
			InstanceGroupType:           group.InstanceGroupType.ValueStringPointer(),
			InstanceType:                group.InstanceType.ValueStringPointer(),
			NodeCount:                   int64To32Pointer(group.NodeCount),
			RecipeNames:                 igRecipes,
			RecoveryMode:                group.RecoveryMode.ValueString(),
			RootVolumeSize:              int64To32(group.RootVolumeSize),
			VolumeEncryption: &datahubmodels.VolumeEncryptionRequest{
				EnableEncryption: group.VolumeEncryption.Encryption.ValueBoolPointer(),
			},
		}
		igs = append(igs, ig)
	}
	req.InstanceGroups = igs
	tflog.Debug(ctx, fmt.Sprintf("Conversion from datahubResourceModel to CreateAWSClusterRequest has finished with request: %+v.", req))
	return &req
}

func fromModelToGcpRequest(model gcpDatahubResourceModel, ctx context.Context) *datahubmodels.CreateGCPClusterRequest {
	tflog.Debug(ctx, "Conversion from gcpDatahubResourceModel to CreateGCPClusterRequest started.")
	req := datahubmodels.CreateGCPClusterRequest{}
	req.ClusterName = model.Name.ValueString()
	req.EnvironmentName = model.Environment.ValueString()
	req.ClusterTemplateName = model.ClusterTemplate.ValueString()
	req.ClusterDefinitionName = model.ClusterDefinition.ValueString()
	var igs []*datahubmodels.GCPInstanceGroupRequest
	tflog.Debug(ctx, fmt.Sprintf("%d instance group found in the input model.", len(model.InstanceGroup)))
	for _, group := range model.InstanceGroup {
		tflog.Debug(ctx, fmt.Sprintf("Converting GCPInstanceGroupRequest: %+v.", group))
		var volReqs []*datahubmodels.AttachedVolumeRequest
		tflog.Debug(ctx, fmt.Sprintf("%d attached volume request found in the input model.", len(model.InstanceGroup)))
		for _, vrs := range group.AttachedVolumeConfiguration {
			tflog.Debug(ctx, fmt.Sprintf("Converting AttachedVolumeConfiguration: %+v.", vrs))
			volReqs = append(volReqs, createAttachedVolumeRequest(vrs))
		}
		var igRecipes []string
		if group.Recipes != nil && len(group.Recipes) > 0 {
			for _, recipe := range group.Recipes {
				igRecipes = append(igRecipes, recipe.ValueString())
			}
		}
		volumeSize := int64To32(group.RootVolumeSize)
		ig := &datahubmodels.GCPInstanceGroupRequest{
			AttachedVolumeConfiguration: volReqs,
			InstanceGroupName:           group.InstanceGroupName.ValueStringPointer(),
			InstanceGroupType:           group.InstanceGroupType.ValueStringPointer(),
			InstanceType:                group.InstanceType.ValueStringPointer(),
			NodeCount:                   int64To32Pointer(group.NodeCount),
			RecipeNames:                 igRecipes,
			RecoveryMode:                group.RecoveryMode.ValueString(),
			RootVolumeSize:              &volumeSize,
		}
		igs = append(igs, ig)
	}
	req.InstanceGroups = igs
	tflog.Debug(ctx, "Conversion from gcpDatahubResourceModel to CreateGCPClusterRequest has finished.")
	return &req
}

func fromModelToAzureRequest(model azureDatahubResourceModel, ctx context.Context) *datahubmodels.CreateAzureClusterRequest {
	tflog.Debug(ctx, "Conversion from datahubResourceModel to CreateAzureClusterRequest started.")
	req := datahubmodels.CreateAzureClusterRequest{}
	req.DatabaseType = model.DatabaseType.ValueString()
	req.ClusterName = model.Name.ValueString()
	req.ClusterTemplateName = model.ClusterTemplate.ValueString()
	req.EnvironmentName = model.Environment.ValueString()
	req.ClusterDefinitionName = model.ClusterDefinition.ValueString()
	var igs []*datahubmodels.AzureInstanceGroupRequest
	tflog.Debug(ctx, fmt.Sprintf("%d instance group found in the input model.", len(model.InstanceGroup)))
	for _, group := range model.InstanceGroup {
		tflog.Debug(ctx, fmt.Sprintf("Converting InstanceGroupRequest: %+v.", group))
		var volReqs []*datahubmodels.AttachedVolumeRequest
		tflog.Debug(ctx, fmt.Sprintf("%d attached volume request found in the input model.", len(model.InstanceGroup)))
		for _, vrs := range group.AttachedVolumeConfiguration {
			tflog.Debug(ctx, fmt.Sprintf("Converting AttachedVolumeConfiguration: %+v.", vrs))
			volReqs = append(volReqs, createAttachedVolumeRequest(vrs))
		}
		var igRecipes []string
		if group.Recipes != nil && len(group.Recipes) > 0 {
			for _, recipe := range group.Recipes {
				igRecipes = append(igRecipes, recipe.ValueString())
			}
		}
		var azs []string
		if group.AvailabilityZones != nil && len(group.AvailabilityZones) > 0 {
			for _, az := range group.AvailabilityZones {
				azs = append(azs, az.ValueString())
			}
		}
		rootVolumeSize := int32(group.RootVolumeSize.ValueInt64())
		ig := &datahubmodels.AzureInstanceGroupRequest{
			AttachedVolumeConfiguration: volReqs,
			InstanceGroupName:           group.InstanceGroupName.ValueStringPointer(),
			InstanceGroupType:           group.InstanceGroupType.ValueStringPointer(),
			InstanceType:                group.InstanceType.ValueStringPointer(),
			NodeCount:                   int64To32Pointer(group.NodeCount),
			RecipeNames:                 igRecipes,
			RecoveryMode:                group.RecoveryMode.ValueString(),
			RootVolumeSize:              &rootVolumeSize,
			AvailabilityZones:           azs,
		}
		igs = append(igs, ig)
	}
	req.InstanceGroups = igs
	return &req
}

func createAttachedVolumeRequest(attachedVolumeConfig AttachedVolumeConfiguration) *datahubmodels.AttachedVolumeRequest {
	return &datahubmodels.AttachedVolumeRequest{
		VolumeCount: int64To32Pointer(attachedVolumeConfig.VolumeCount),
		VolumeSize:  int64To32Pointer(attachedVolumeConfig.VolumeSize),
		VolumeType:  attachedVolumeConfig.VolumeType.ValueStringPointer(),
	}
}

func int64To32(in types.Int64) int32 {
	n64 := in.ValueInt64()
	return int32(n64)
}

func int64To32Pointer(in types.Int64) *int32 {
	n64 := in.ValueInt64()
	var n2 = int32(n64)
	return &n2
}
