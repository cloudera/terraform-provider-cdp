// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package freeipa

import (
	"context"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/resources/environments/envcommon"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

func ResponseToModel(ipaResp *environmentsmodels.FreeipaDetails, model *types.Object, ctx context.Context) *diag.Diagnostics {
	utils.LogFreeIpaSilently(ctx, ipaResp, "Converting FreeIpa to Model from response: ")
	var diags diag.Diagnostics

	if ipaResp == nil {
		return &diags
	}

	var recipes types.Set
	var recipeDiags diag.Diagnostics
	recipes, recipeDiags = types.SetValueFrom(ctx, types.StringType, ipaResp.Recipes)
	diags.Append(recipeDiags...)

	var freeIpaDetails FreeIpaDetails
	model.As(ctx, &freeIpaDetails, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})

	instSet := make([]FreeIpaInstance, len(ipaResp.Instances))
	var ipaInstances types.Set
	if len(ipaResp.Instances) > 0 {
		for i, v := range ipaResp.Instances {
			volumes, volumesDiags := convertAttachedVolumes(v, ctx, diags)
			diags.Append(volumesDiags...)
			inst := FreeIpaInstance{
				AvailabilityZone:     types.StringValue(v.AvailabilityZone),
				DiscoveryFQDN:        types.StringValue(v.DiscoveryFQDN),
				InstanceGroup:        types.StringValue(v.InstanceGroup),
				InstanceID:           types.StringValue(v.InstanceID),
				InstanceStatus:       types.StringValue(v.InstanceStatus),
				InstanceStatusReason: types.StringValue(v.InstanceStatusReason),
				InstanceType:         types.StringValue(v.InstanceType),
				InstanceVMType:       types.StringValue(v.InstanceVMType),
				LifeCycle:            types.StringValue(v.LifeCycle),
				PrivateIP:            types.StringValue(v.PrivateIP),
				PublicIP:             types.StringValue(v.PublicIP),
				SSHPort:              types.Int64Value(int64(v.SSHPort)),
				SubnetID:             types.StringValue(v.SubnetID),
				AttachedVolumes:      volumes,
			}
			instSet[i] = inst
		}
	}

	var instDiags diag.Diagnostics
	ipaInstances, instDiags = types.SetValueFrom(ctx, FreeIpaInstanceType, instSet)
	diags.Append(instDiags...)

	var instanceCount basetypes.Int32Value
	if val, ok := model.Attributes()["instance_count_by_group"]; ok {
		if val.IsNull() {
			var countErr error
			instanceCount, countErr = envcommon.ConvertIntToInt32IfPossible(len(ipaResp.Instances))
			if countErr != nil {
				diags.AddWarning(fmt.Sprintf("Unable to convert the numerical value of the length of the instances slice. Fallback to %d", ipaResp.InstanceCountByGroup), countErr.Error())
				instanceCount = types.Int32Value(ipaResp.InstanceCountByGroup)
			}
		} else {
			instanceCount = val.(basetypes.Int32Value)
		}
	}
	var ipaDiags diag.Diagnostics
	*model, ipaDiags = types.ObjectValueFrom(ctx, FreeIpaDetailsType.AttrTypes, &FreeIpaDetails{
		Catalog:              freeIpaDetails.Catalog,
		ImageID:              freeIpaDetails.ImageID,
		Os:                   freeIpaDetails.Os,
		InstanceCountByGroup: instanceCount,
		InstanceType:         freeIpaDetails.InstanceType,
		Instances:            ipaInstances,
		MultiAz:              types.BoolValue(ipaResp.MultiAz),
		Recipes:              recipes,
		Architecture:         envcommon.GetStringValueIfNotEmpty(freeIpaDetails.Architecture.ValueString()),
	})
	diags.Append(ipaDiags...)

	return &diags
}

func convertAttachedVolumes(v *environmentsmodels.FreeIpaInstance, ctx context.Context, diags diag.Diagnostics) (types.Set, diag.Diagnostics) {
	var volumes types.Set
	attachedVolumesSet := make([]FreeIpaAttachedVolumes, len(v.AttachedVolumes))
	if len(v.AttachedVolumes) > 0 {
		for j, vol := range v.AttachedVolumes {
			if vol != nil {
				attachedVolumesSet[j] = FreeIpaAttachedVolumes{
					Count:      types.Int32Value(vol.Count),
					VolumeType: types.StringValue(vol.VolumeType),
					Size:       types.Int32Value(vol.Size),
				}
			}
		}
		var attachedVolumesDiag diag.Diagnostics
		volumes, attachedVolumesDiag = types.SetValueFrom(ctx, FreeIpaAttachedVolumesType, attachedVolumesSet)
		diags.Append(attachedVolumesDiag...)
	} else {
		volumes = types.SetNull(FreeIpaAttachedVolumesType)
	}
	return volumes, diags
}

type Transitional struct {
	InstanceCountByGroup int32

	InstanceType string

	MultiAz bool

	Recipes []string

	Architecture string
}

func ModelToRequest(model *types.Object, ctx context.Context) (*Transitional, *environmentsmodels.FreeIpaImageRequest) {
	var freeIpaDetails FreeIpaDetails
	model.As(ctx, &freeIpaDetails, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
	return &Transitional{
			InstanceCountByGroup: freeIpaDetails.InstanceCountByGroup.ValueInt32(),
			InstanceType:         freeIpaDetails.InstanceType.ValueString(),
			MultiAz:              freeIpaDetails.MultiAz.ValueBool(),
			Recipes:              utils.FromSetValueToStringList(freeIpaDetails.Recipes),
			Architecture:         freeIpaDetails.Architecture.ValueString(),
		},
		&environmentsmodels.FreeIpaImageRequest{
			Catalog: freeIpaDetails.Catalog.ValueString(),
			ID:      freeIpaDetails.ImageID.ValueString(),
			Os:      freeIpaDetails.Os.ValueString(),
		}
}

func UpdateCatalogIfChanged(ctx context.Context, planFreeIpa types.Object, stateFreeIpa *types.Object, environmentName string, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	var planDetails FreeIpaDetails
	resp.Diagnostics.Append(planFreeIpa.As(ctx, &planDetails, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
	if resp.Diagnostics.HasError() {
		return resp
	}

	var stateDetails FreeIpaDetails
	resp.Diagnostics.Append(stateFreeIpa.As(ctx, &stateDetails, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
	if resp.Diagnostics.HasError() {
		return resp
	}

	if planDetails.Catalog.IsNull() || planDetails.Catalog.IsUnknown() || planDetails.Catalog.Equal(stateDetails.Catalog) {
		return resp
	}

	tflog.Info(ctx, fmt.Sprintf("Catalog change detected for environment '%s', calling SetCatalog.", environmentName))

	params := operations.NewSetCatalogParams()
	params.WithInput(&environmentsmodels.SetCatalogRequest{
		Catalog:     new(planDetails.Catalog.ValueString()),
		Environment: &environmentName,
	})
	_, err := client.Operations.SetCatalogContext(ctx, params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "set catalog")
		return resp
	}

	updateFreeIpaCatalogInState(ctx, stateFreeIpa, planDetails.Catalog, &resp.Diagnostics)
	return resp
}

func updateFreeIpaCatalogInState(ctx context.Context, freeIpaObj *types.Object, newCatalog types.String, diags *diag.Diagnostics) {
	var details FreeIpaDetails
	diags.Append(freeIpaObj.As(ctx, &details, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
	if diags.HasError() {
		return
	}
	details.Catalog = newCatalog
	newObj, objDiags := types.ObjectValueFrom(ctx, FreeIpaDetailsType.AttrTypes, &details)
	diags.Append(objDiags...)
	*freeIpaObj = newObj
}

func UpdateRecipesIfChanged(ctx context.Context, planFreeIpa types.Object, stateFreeIpa *types.Object, environmentName string, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	var planDetails FreeIpaDetails
	resp.Diagnostics.Append(planFreeIpa.As(ctx, &planDetails, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
	if resp.Diagnostics.HasError() {
		return resp
	}

	var stateDetails FreeIpaDetails
	resp.Diagnostics.Append(stateFreeIpa.As(ctx, &stateDetails, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
	if resp.Diagnostics.HasError() {
		return resp
	}

	if planDetails.Recipes.IsNull() || planDetails.Recipes.IsUnknown() || reflect.DeepEqual(planDetails.Recipes, stateDetails.Recipes) {
		return resp
	}

	tflog.Info(ctx, fmt.Sprintf("Recipe change detected for environment '%s', calling Attach/Detach.", environmentName))

	planRecipes := utils.FromSetValueToStringList(planDetails.Recipes)
	stateRecipes := utils.FromSetValueToStringList(stateDetails.Recipes)

	planMap := make(map[string]struct{}, len(planRecipes))
	for _, r := range planRecipes {
		planMap[r] = struct{}{}
	}
	stateMap := make(map[string]struct{}, len(stateRecipes))
	for _, r := range stateRecipes {
		stateMap[r] = struct{}{}
	}

	var toAttach []string
	for r := range planMap {
		if _, exists := stateMap[r]; !exists {
			toAttach = append(toAttach, r)
		}
	}
	var toDetach []string
	for r := range stateMap {
		if _, exists := planMap[r]; !exists {
			toDetach = append(toDetach, r)
		}
	}

	if len(toAttach) > 0 {
		params := operations.NewAttachFreeIpaRecipesParams()
		params.WithInput(&environmentsmodels.AttachFreeIpaRecipesRequest{
			Environment: &environmentName,
			Recipes:     toAttach,
		})
		if _, err := client.Operations.AttachFreeIpaRecipesContext(ctx, params); err != nil {
			utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "attach freeipa recipes")
			return resp
		}
	}

	if len(toDetach) > 0 {
		params := operations.NewDetachFreeIpaRecipesParams()
		params.WithInput(&environmentsmodels.DetachFreeIpaRecipesRequest{
			Environment: &environmentName,
			Recipes:     toDetach,
		})
		if _, err := client.Operations.DetachFreeIpaRecipesContext(ctx, params); err != nil {
			utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "detach freeipa recipes")
			return resp
		}
	}

	updateFreeIpaRecipesInState(ctx, stateFreeIpa, planDetails.Recipes, &resp.Diagnostics)
	return resp
}

func updateFreeIpaRecipesInState(ctx context.Context, freeIpaObj *types.Object, newRecipes types.Set, diags *diag.Diagnostics) {
	var details FreeIpaDetails
	diags.Append(freeIpaObj.As(ctx, &details, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
	if diags.HasError() {
		return
	}
	details.Recipes = newRecipes
	newObj, objDiags := types.ObjectValueFrom(ctx, FreeIpaDetailsType.AttrTypes, &details)
	diags.Append(objDiags...)
	*freeIpaObj = newObj
}
