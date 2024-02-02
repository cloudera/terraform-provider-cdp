// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package environments

import (
	"context"

	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func FreeIpaResponseToModel(ipaResp *environmentsmodels.FreeipaDetails, model *types.Object, ctx context.Context) *diag.Diagnostics {
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
			}
			instSet[i] = inst
		}

	}

	var instDiags diag.Diagnostics
	ipaInstances, instDiags = types.SetValueFrom(ctx, FreeIpaInstanceType, instSet)
	diags.Append(instDiags...)

	var ipaDiags diag.Diagnostics
	*model, ipaDiags = types.ObjectValueFrom(ctx, FreeIpaDetailsType.AttrTypes, &FreeIpaDetails{
		Catalog:              freeIpaDetails.Catalog,
		ImageID:              freeIpaDetails.ImageID,
		Os:                   freeIpaDetails.Os,
		InstanceCountByGroup: types.Int64Value(int64(ipaResp.InstanceCountByGroup)),
		InstanceType:         freeIpaDetails.InstanceType,
		Instances:            ipaInstances,
		MultiAz:              freeIpaDetails.MultiAz,
		Recipes:              recipes,
	})
	diags.Append(ipaDiags...)

	return &diags
}

type FreeIpaTransitional struct {
	InstanceCountByGroup int32

	InstanceType string

	MultiAz bool

	Recipes []string
}

func FreeIpaModelToRequest(model *types.Object, ctx context.Context) (*FreeIpaTransitional, *environmentsmodels.FreeIpaImageRequest) {
	var freeIpaDetails FreeIpaDetails
	model.As(ctx, &freeIpaDetails, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
	return &FreeIpaTransitional{
			InstanceCountByGroup: int32(freeIpaDetails.InstanceCountByGroup.ValueInt64()),
			InstanceType:         freeIpaDetails.InstanceType.ValueString(),
			MultiAz:              freeIpaDetails.MultiAz.ValueBool(),
			Recipes:              utils.FromSetValueToStringList(freeIpaDetails.Recipes),
		},
		&environmentsmodels.FreeIpaImageRequest{
			Catalog: freeIpaDetails.Catalog.ValueString(),
			ID:      freeIpaDetails.ImageID.ValueString(),
			Os:      freeIpaDetails.Os.ValueString(),
		}
}
