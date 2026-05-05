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
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int32default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int32planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var FreeIpaSchema = schema.SingleNestedAttribute{
	MarkdownDescription: "Request object for creating FreeIPA in the environment.",
	Description:         "Request object for creating FreeIPA in the environment.",
	Optional:            true,
	Computed:            true,
	PlanModifiers: []planmodifier.Object{
		objectplanmodifier.UseStateForUnknown(),
	},
	Attributes: map[string]schema.Attribute{
		"catalog": schema.StringAttribute{
			MarkdownDescription: "Image catalog to use for FreeIPA image selection.",
			Description:         "Image catalog to use for FreeIPA image selection.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"image_id": schema.StringAttribute{
			MarkdownDescription: "Image ID to use for creating FreeIPA instances.",
			Description:         "Image ID to use for creating FreeIPA instances.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"os": schema.StringAttribute{
			MarkdownDescription: "The OS to use for creating FreeIPA instances.",
			Description:         "The OS to use for creating FreeIPA instances.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"instance_count_by_group": schema.Int32Attribute{
			MarkdownDescription: "The number of FreeIPA instances to create per group when creating FreeIPA in the environment",
			Description:         "The number of FreeIPA instances to create per group when creating FreeIPA in the environment",
			Optional:            true,
			Computed:            true,
			Default:             int32default.StaticInt32(1),
			PlanModifiers: []planmodifier.Int32{
				int32planmodifier.UseStateForUnknown(),
			},
		},
		"instance_type": schema.StringAttribute{
			MarkdownDescription: "Custom instance type of FreeIPA instances.",
			Description:         "Custom instance type of FreeIPA instances.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"architecture": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "CPU architecture of the FreeIPA instance. Can be either X86_64 or ARM64.",
			Validators: []validator.String{
				stringvalidator.OneOf("X86_64", "ARM64"),
			},
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"instances": schema.SetNestedAttribute{
			MarkdownDescription: "The instances of the FreeIPA cluster.",
			Description:         "The instances of the FreeIPA cluster.",
			Computed:            true,
			PlanModifiers: []planmodifier.Set{
				setplanmodifier.UseStateForUnknown(),
			},
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"availability_zone": schema.StringAttribute{
						MarkdownDescription: "The availability zone of the instance.",
						Description:         "The availability zone of the instance.",
						Computed:            true,
					},
					"discovery_fqdn": schema.StringAttribute{
						MarkdownDescription: "The fully qualified domain name of the instance in the service discovery cluster.",
						Description:         "The fully qualified domain name of the instance in the service discovery cluster.",
						Computed:            true,
					},
					"instance_group": schema.StringAttribute{
						MarkdownDescription: "The instance group that contains the instance.",
						Description:         "The instance group that contains the instance.",
						Computed:            true,
					},
					"instance_id": schema.StringAttribute{
						MarkdownDescription: "The instance ID for the instance.",
						Description:         "The instance ID for the instance.",
						Computed:            true,
					},
					"instance_status": schema.StringAttribute{
						MarkdownDescription: "The status of the instance.",
						Description:         "The status of the instance.",
						Computed:            true,
					},
					"instance_status_reason": schema.StringAttribute{
						MarkdownDescription: "The status reason for the instance.",
						Description:         "The status reason for the instance.",
						Computed:            true,
					},
					"instance_type": schema.StringAttribute{
						MarkdownDescription: "The type of the instance (either GATEWAY or GATEWAY_PRIMARY).",
						Description:         "The type of the instance (either GATEWAY or GATEWAY_PRIMARY).",
						Computed:            true,
					},
					"instance_vm_type": schema.StringAttribute{
						MarkdownDescription: "The VM type of the instance. Supported values depend on the cloud platform.",
						Description:         "The VM type of the instance. Supported values depend on the cloud platform.",
						Computed:            true,
					},
					"life_cycle": schema.StringAttribute{
						MarkdownDescription: "The life cycle type for the instance (either NORMAL or SPOT).",
						Description:         "The life cycle type for the instance (either NORMAL or SPOT).",
						Computed:            true,
					},
					"private_ip": schema.StringAttribute{
						MarkdownDescription: "The private IP of the instance.",
						Description:         "The private IP of the instance.",
						Computed:            true,
					},
					"public_ip": schema.StringAttribute{
						MarkdownDescription: "The public IP of the instance.",
						Description:         "The public IP of the instance.",
						Computed:            true,
					},
					"ssh_port": schema.Int64Attribute{
						MarkdownDescription: "The SSH port of the instance.",
						Description:         "The SSH port of the instance.",
						Computed:            true,
					},
					"subnet_id": schema.StringAttribute{
						MarkdownDescription: "The subnet ID of the instance.",
						Description:         "The subnet ID of the instance.",
						Computed:            true,
					},
				},
			},
		},
		"multi_az": schema.BoolAttribute{
			MarkdownDescription: "Flag that enables deployment of the FreeIPA in a multi-availability zone.",
			Description:         "Flag that enables deployment of the FreeIPA in a multi-availability zone.",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"recipes": schema.SetAttribute{
			MarkdownDescription: "The recipes for the FreeIPA cluster.",
			Description:         "The recipes for the FreeIPA cluster.",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.Set{
				setplanmodifier.UseStateForUnknown(),
			},
			ElementType: types.StringType,
		},
	},
}
