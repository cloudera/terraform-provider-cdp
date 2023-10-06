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
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/types"

	datalakemodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

func datalakeDetailsToGcpDatalakeResourceModel(ctx context.Context, resp *datalakemodels.DatalakeDetails, model *gcpDatalakeResourceModel, pollingOptions *utils.PollingOptions, diags *diag.Diagnostics) {
	model.ID = types.StringPointerValue(resp.Crn)
	model.CreationDate = types.StringValue(resp.CreationDate.String())
	model.Crn = types.StringPointerValue(resp.Crn)
	model.DatalakeName = types.StringPointerValue(resp.DatalakeName)
	model.EnableRangerRaz = types.BoolValue(resp.EnableRangerRaz)
	model.PollingOptions = pollingOptions
	endpoints := make([]*endpoint, len(resp.Endpoints.Endpoints))
	for i, v := range resp.Endpoints.Endpoints {
		endpoints[i] = &endpoint{
			DisplayName: types.StringPointerValue(v.DisplayName),
			KnoxService: types.StringPointerValue(v.KnoxService),
			Mode:        types.StringPointerValue(v.Mode),
			Open:        types.BoolPointerValue(v.Open),
			ServiceName: types.StringPointerValue(v.ServiceName),
			ServiceURL:  types.StringPointerValue(v.ServiceURL),
		}
	}
	model.EnvironmentCrn = types.StringValue(resp.EnvironmentCrn)
	productVersions := make([]*productVersion, len(resp.ProductVersions))
	for i, v := range resp.ProductVersions {
		productVersions[i] = &productVersion{
			Name:    types.StringPointerValue(v.Name),
			Version: types.StringPointerValue(v.Version),
		}
	}
	model.Scale = types.StringValue(string(resp.Shape))
	model.Status = types.StringValue(resp.Status)
	model.StatusReason = types.StringValue(resp.StatusReason)
}

func toGcpDatalakeRequest(ctx context.Context, model *gcpDatalakeResourceModel) *datalakemodels.CreateGCPDatalakeRequest {
	req := &datalakemodels.CreateGCPDatalakeRequest{}
	if model.CloudProviderConfiguration != nil {
		req.CloudProviderConfiguration = &datalakemodels.GCPConfigurationRequest{
			ServiceAccountEmail: model.CloudProviderConfiguration.ServiceAccountEmail.ValueStringPointer(),
			StorageLocation:     model.CloudProviderConfiguration.StorageLocation.ValueStringPointer(),
		}
	}
	req.CustomInstanceGroups = make([]*datalakemodels.SdxInstanceGroupRequest, len(model.CustomInstanceGroups))
	for i, v := range model.CustomInstanceGroups {
		req.CustomInstanceGroups[i] = &datalakemodels.SdxInstanceGroupRequest{
			InstanceType: v.InstanceType.ValueString(),
			Name:         v.Name.ValueStringPointer(),
		}
	}
	req.DatalakeName = model.DatalakeName.ValueStringPointer()
	req.EnableRangerRaz = model.EnableRangerRaz.ValueBool()
	req.EnvironmentName = model.EnvironmentName.ValueStringPointer()
	if model.Image != nil {
		req.Image = &datalakemodels.ImageRequest{
			CatalogName: model.Image.CatalogName.ValueStringPointer(),
			ID:          model.Image.ID.ValueStringPointer(),
		}
	}
	req.JavaVersion = int32(model.JavaVersion.ValueInt64())
	req.Recipes = make([]*datalakemodels.InstanceGroupRecipeRequest, len(model.Recipes))
	for i, v := range model.Recipes {
		req.Recipes[i] = &datalakemodels.InstanceGroupRecipeRequest{
			InstanceGroupName: v.InstanceGroupName.ValueStringPointer(),
			RecipeNames:       utils.FromSetValueToStringList(v.RecipeNames),
		}
	}
	req.Runtime = model.Runtime.ValueString()
	req.Scale = datalakemodels.DatalakeScaleType(model.Scale.ValueString())
	if !model.Tags.IsNull() {
		req.Tags = make([]*datalakemodels.DatalakeResourceGCPTagRequest, len(model.Tags.Elements()))
		i := 0
		for k, v := range model.Tags.Elements() {
			val, diag := v.(basetypes.StringValuable).ToStringValue(ctx)
			if !diag.HasError() {
				req.Tags[i] = &datalakemodels.DatalakeResourceGCPTagRequest{
					Key:   &k,
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
