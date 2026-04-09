// Copyright 2023 Cloudera. All Rights Reserved.
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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	datalakemodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

func datalakeDetailsToGcpDatalakeResourceModelFromCreation(resp *datalakemodels.DatalakeDetails, model *gcpDatalakeResourceModel, pollingOptions *utils.PollingOptions) {
	model.ID = types.StringPointerValue(resp.Crn)
	model.CreationDate = types.StringValue(resp.CreationDate.String())
	model.Crn = types.StringPointerValue(resp.Crn)
	model.DatalakeName = types.StringPointerValue(resp.DatalakeName)
	model.EnableRangerRaz = types.BoolValue(resp.EnableRangerRaz)
	model.PollingOptions = pollingOptions
	model.EnvironmentCrn = types.StringValue(resp.EnvironmentCrn)
	model.Scale = types.StringValue(string(resp.Shape))
	model.Status = types.StringValue(resp.Status)
	model.StatusReason = types.StringValue(resp.StatusReason)
}

func datalakeDetailsToGcpDatalakeResourceModel(ctx context.Context, resp *datalakemodels.DatalakeDetails, model *gcpDatalakeResourceModel, pollingOptions *utils.PollingOptions, diags *diag.Diagnostics) {
	model.ID = types.StringPointerValue(resp.Crn)
	if resp.GcpConfiguration != nil {
		model.CloudProviderConfiguration.ServiceAccountEmail = types.StringValue(resp.GcpConfiguration.ServiceAccountEmail)
	}
	model.CreationDate = types.StringValue(resp.CreationDate.String())
	model.Crn = types.StringPointerValue(resp.Crn)
	model.DatalakeName = types.StringPointerValue(resp.DatalakeName)
	model.EnableRangerRaz = types.BoolValue(resp.EnableRangerRaz)
	model.PollingOptions = pollingOptions
	model.EnvironmentCrn = types.StringValue(resp.EnvironmentCrn)
	instanceGroups := make([]*instanceGroup, len(resp.InstanceGroups))
	for i, v := range resp.InstanceGroups {
		instanceGroups[i] = &instanceGroup{
			Name: types.StringPointerValue(v.Name),
		}

		instances := make([]*instance, 0, len(v.Instances))
		for _, ins := range v.Instances {
			if ins == nil || ins.ID == nil || len(*ins.ID) == 0 {
				continue
			}
			instances = append(instances, &instance{
				DiscoveryFQDN:   types.StringValue(ins.DiscoveryFQDN),
				ID:              types.StringPointerValue(ins.ID),
				InstanceGroup:   types.StringValue(ins.InstanceGroup),
				InstanceStatus:  types.StringValue(string(ins.InstanceStatus)),
				InstanceTypeVal: types.StringValue(string(ins.InstanceTypeVal)),
				PrivateIP:       types.StringValue(ins.PrivateIP),
				PublicIP:        types.StringValue(ins.PublicIP),
				SSHPort:         types.Int64Value(int64(ins.SSHPort)),
				State:           types.StringPointerValue(ins.State),
				StatusReason:    types.StringValue(ins.StatusReason),
			})
		}
		var instDiags diag.Diagnostics
		instanceGroups[i].Instances, instDiags = types.SetValueFrom(ctx, types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"discovery_fqdn":    types.StringType,
				"id":                types.StringType,
				"instance_group":    types.StringType,
				"instance_status":   types.StringType,
				"instance_type_val": types.StringType,
				"private_ip":        types.StringType,
				"public_ip":         types.StringType,
				"ssh_port":          types.Int64Type,
				"state":             types.StringType,
				"status_reason":     types.StringType,
			},
		}, instances)
		diags.Append(instDiags...)
	}
	model.Scale = types.StringValue(string(resp.Shape))
	model.Status = types.StringValue(resp.Status)
	model.StatusReason = types.StringValue(resp.StatusReason)
	model.CloudProviderConfiguration.StorageLocation = types.StringValue(resp.CloudStorageBaseLocation)
}

func toGcpDatalakeRequest(ctx context.Context, model *gcpDatalakeResourceModel) *datalakemodels.CreateGCPDatalakeRequest {
	req := &datalakemodels.CreateGCPDatalakeRequest{
		EnvironmentName: model.EnvironmentName.ValueStringPointer(),
		DatalakeName:    model.DatalakeName.ValueStringPointer(),
		EnableRangerRaz: model.EnableRangerRaz.ValueBool(),
		JavaVersion:     model.JavaVersion.ValueInt32(),
		MultiAz:         model.MultiAz.ValueBoolPointer(),
		Runtime:         model.Runtime.ValueString(),
	}
	if model.Security != nil {
		req.Security = &datalakemodels.SecurityRequest{
			SeLinux: model.Security.SeLinux.ValueString(),
		}
	}
	if model.CloudProviderConfiguration != nil {
		req.CloudProviderConfiguration = &datalakemodels.GCPConfigurationRequest{
			ServiceAccountEmail: model.CloudProviderConfiguration.ServiceAccountEmail.ValueStringPointer(),
			StorageLocation:     model.CloudProviderConfiguration.StorageLocation.ValueStringPointer(),
		}
	}
	if model.Image != nil {
		req.Image = &datalakemodels.ImageRequest{
			CatalogName: model.Image.CatalogName.ValueStringPointer(),
			ID:          model.Image.ID.ValueString(),
			Os:          model.Image.Os.ValueString(),
		}
	}
	req.Recipes = make([]*datalakemodels.InstanceGroupRecipeRequest, len(model.Recipes))
	for i, v := range model.Recipes {
		req.Recipes[i] = &datalakemodels.InstanceGroupRecipeRequest{
			InstanceGroupName: v.InstanceGroupName.ValueStringPointer(),
			RecipeNames:       utils.FromSetValueToStringList(v.RecipeNames),
		}
	}
	req.CustomInstanceGroups = make([]*datalakemodels.SdxInstanceGroupRequest, len(model.CustomInstanceGroups))
	for i, v := range model.CustomInstanceGroups {
		req.CustomInstanceGroups[i] = &datalakemodels.SdxInstanceGroupRequest{
			Name:         v.Name.ValueStringPointer(),
			InstanceType: v.InstanceType.ValueString(),
		}
	}
	req.Scale = datalakemodels.DatalakeScaleType(model.Scale.ValueString())
	if !model.Tags.IsNull() {
		req.Tags = make([]*datalakemodels.DatalakeResourceGCPTagRequest, len(model.Tags.Elements()))
		i := 0
		for k, v := range model.Tags.Elements() {
			key := k
			val, diag := v.(basetypes.StringValuable).ToStringValue(ctx)
			if !diag.HasError() {
				req.Tags[i] = &datalakemodels.DatalakeResourceGCPTagRequest{
					Key:   &key,
					Value: val.ValueStringPointer(),
				}
			}
			i++
		}
	}
	return req
}

func toGcpDatalakeResourceModel(resp *datalakemodels.CreateGCPDatalakeResponse, model *gcpDatalakeResourceModel) {
	model.ID = types.StringPointerValue(resp.Datalake.DatalakeName)
	model.CreationDate = types.StringValue(resp.Datalake.CreationDate.String())
	model.Crn = types.StringPointerValue(resp.Datalake.Crn)
	model.DatalakeName = types.StringPointerValue(resp.Datalake.DatalakeName)
	model.EnableRangerRaz = types.BoolValue(resp.Datalake.EnableRangerRaz)
	model.EnvironmentCrn = types.StringValue(resp.Datalake.EnvironmentCrn)
	model.MultiAz = types.BoolValue(resp.Datalake.MultiAz)
	model.Status = types.StringValue(resp.Datalake.Status)
	model.StatusReason = types.StringValue(resp.Datalake.StatusReason)
}
