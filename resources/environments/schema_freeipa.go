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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var FreeIpaSchema = schema.SingleNestedAttribute{
	Optional: true,
	Computed: true,
	PlanModifiers: []planmodifier.Object{
		objectplanmodifier.UseStateForUnknown(),
	},
	Attributes: map[string]schema.Attribute{
		"catalog": schema.StringAttribute{
			Optional: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"image_id": schema.StringAttribute{
			Optional: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"os": schema.StringAttribute{
			Optional: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"instance_count_by_group": schema.Int64Attribute{
			Optional: true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		"instance_type": schema.StringAttribute{
			Optional: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"instances": schema.SetNestedAttribute{
			Computed: true,
			Optional: true,
			PlanModifiers: []planmodifier.Set{
				setplanmodifier.UseStateForUnknown(),
			},
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"availability_zone": schema.StringAttribute{
						Computed: true,
					},
					"discovery_fqdn": schema.StringAttribute{
						Computed: true,
					},
					"instance_group": schema.StringAttribute{
						Computed: true,
					},
					"instance_id": schema.StringAttribute{
						Computed: true,
					},
					"instance_status": schema.StringAttribute{
						Computed: true,
					},
					"instance_status_reason": schema.StringAttribute{
						Computed: true,
					},
					"instance_type": schema.StringAttribute{
						Computed: true,
					},
					"instance_vm_type": schema.StringAttribute{
						Computed: true,
					},
					"life_cycle": schema.StringAttribute{
						Computed: true,
					},
					"private_ip": schema.StringAttribute{
						Computed: true,
					},
					"public_ip": schema.StringAttribute{
						Computed: true,
					},
					"ssh_port": schema.Int64Attribute{
						Computed: true,
					},
					"subnet_id": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
		"multi_az": schema.BoolAttribute{
			Optional: true,
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"recipes": schema.SetAttribute{
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Set{
				setplanmodifier.UseStateForUnknown(),
			},
			ElementType: types.StringType,
		},
	},
}
