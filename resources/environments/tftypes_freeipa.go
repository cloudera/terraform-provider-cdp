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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var FreeIpaDetailsObject = tftypes.Object{
	AttributeTypes: map[string]tftypes.Type{
		"catalog":                 tftypes.String,
		"image_id":                tftypes.String,
		"os":                      tftypes.String,
		"instance_count_by_group": tftypes.Number,
		"instance_type":           tftypes.String,
		"instances": tftypes.Set{
			ElementType: FreeIpaInstanceObject,
		},
		"multi_az": tftypes.Bool,
		"recipes": tftypes.Set{
			ElementType: tftypes.String,
		},
		"architecture": tftypes.String,
	},
}

var FreeIpaInstanceObject = tftypes.Object{
	AttributeTypes: map[string]tftypes.Type{
		"availability_zone":      tftypes.String,
		"discovery_fqdn":         tftypes.String,
		"instance_group":         tftypes.String,
		"instance_id":            tftypes.String,
		"instance_status":        tftypes.String,
		"instance_status_reason": tftypes.String,
		"instance_type":          tftypes.String,
		"instance_vm_type":       tftypes.String,
		"life_cycle":             tftypes.String,
		"private_ip":             tftypes.String,
		"public_ip":              tftypes.String,
		"ssh_port":               tftypes.Number,
		"subnet_id":              tftypes.String,
	},
}

var FreeIpaDetailsType = types.ObjectType{
	AttrTypes: map[string]attr.Type{
		"catalog":                 types.StringType,
		"image_id":                types.StringType,
		"os":                      types.StringType,
		"instance_count_by_group": types.Int64Type,
		"instance_type":           types.StringType,
		"instances": types.SetType{
			ElemType: FreeIpaInstanceType,
		},
		"multi_az": types.BoolType,
		"recipes": types.SetType{
			ElemType: types.StringType,
		},
		"architecture": types.StringType,
	},
}

var FreeIpaInstanceType = types.ObjectType{
	AttrTypes: map[string]attr.Type{
		"availability_zone":      types.StringType,
		"discovery_fqdn":         types.StringType,
		"instance_group":         types.StringType,
		"instance_id":            types.StringType,
		"instance_status":        types.StringType,
		"instance_status_reason": types.StringType,
		"instance_type":          types.StringType,
		"instance_vm_type":       types.StringType,
		"life_cycle":             types.StringType,
		"private_ip":             types.StringType,
		"public_ip":              types.StringType,
		"ssh_port":               types.Int64Type,
		"subnet_id":              types.StringType,
	},
}
