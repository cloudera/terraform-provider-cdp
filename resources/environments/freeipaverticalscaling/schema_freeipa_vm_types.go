// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package freeipaverticalscaling

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var freeipaVMTypesSchema = schema.Schema{
	MarkdownDescription: "Lists the available VM types for FreeIPA vertical scaling.",
	Attributes: map[string]schema.Attribute{
		"environment_crn": schema.StringAttribute{
			MarkdownDescription: "The CRN of the environment for which to list the available VM types.",
			Required:            true,
		},
		"architecture": schema.StringAttribute{
			MarkdownDescription: "The architecture to filter VM types. For example, x86_64.",
			Optional:            true,
		},
		"availability_zones": schema.ListAttribute{
			MarkdownDescription: "The availability zones to filter VM types.",
			Optional:            true,
			ElementType:         types.StringType,
		},
		"instance_type": schema.StringAttribute{
			MarkdownDescription: "The current instance type to filter compatible VM types.",
			Optional:            true,
		},
		"resource_type": schema.StringAttribute{
			MarkdownDescription: "The resource type to filter VM types. For example, DEFAULT.",
			Optional:            true,
		},
		"vm_types": schema.ListNestedAttribute{
			MarkdownDescription: "The list of platform VM types grouped by platform.",
			Computed:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"platform": schema.StringAttribute{
						MarkdownDescription: "The cloud platform (e.g., AWS, AZURE, GCP).",
						Computed:            true,
					},
					"default_virtual_machine": schema.SingleNestedAttribute{
						MarkdownDescription: "The default VM type for the platform.",
						Computed:            true,
						Attributes:          vmTypeSchemaAttributes(),
					},
					"virtual_machines": schema.ListNestedAttribute{
						MarkdownDescription: "The list of available VM types for the platform.",
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: vmTypeSchemaAttributes(),
						},
					},
				},
			},
		},
	},
}

func vmTypeSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the VM type.",
			Computed:            true,
		},
		"vm_type_metadata": schema.SingleNestedAttribute{
			MarkdownDescription: "Metadata about the VM type.",
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"configs": schema.ListNestedAttribute{
					MarkdownDescription: "The list of volume parameter configurations.",
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"maximum_number": schema.Int32Attribute{
								MarkdownDescription: "The maximum number of volumes.",
								Computed:            true,
							},
							"maximum_size": schema.Int32Attribute{
								MarkdownDescription: "The maximum volume size in GB.",
								Computed:            true,
							},
							"minimum_number": schema.Int32Attribute{
								MarkdownDescription: "The minimum number of volumes.",
								Computed:            true,
							},
							"minimum_size": schema.Int32Attribute{
								MarkdownDescription: "The minimum volume size in GB.",
								Computed:            true,
							},
							"volume_parameter_type": schema.StringAttribute{
								MarkdownDescription: "The type of volume parameter (e.g., SSD, MAGNETIC).",
								Computed:            true,
							},
						},
					},
				},
				"properties": schema.MapAttribute{
					MarkdownDescription: "Additional properties of the VM type.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	}
}
