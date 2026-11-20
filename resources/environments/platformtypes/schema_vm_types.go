// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package platformtypes

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func VMTypesNestedAttribute() schema.ListNestedAttribute {
	return schema.ListNestedAttribute{
		MarkdownDescription: "VM types grouped by platform.",
		Computed:            true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"platform": schema.StringAttribute{
					MarkdownDescription: "Cloud platform (AWS, AZURE, GCP).",
					Computed:            true,
				},
				"default_virtual_machine": schema.SingleNestedAttribute{
					MarkdownDescription: "Default VM type for this platform.",
					Computed:            true,
					Attributes:          VMTypeSchemaAttributes(),
				},
				"virtual_machines": schema.ListNestedAttribute{
					MarkdownDescription: "Available VM types for this platform.",
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: VMTypeSchemaAttributes(),
					},
				},
			},
		},
	}
}

func VMTypeSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"name": schema.StringAttribute{
			MarkdownDescription: "VM type name.",
			Computed:            true,
		},
		"vm_type_metadata": schema.SingleNestedAttribute{
			MarkdownDescription: "VM type metadata.",
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"configs": schema.ListNestedAttribute{
					MarkdownDescription: "Volume parameter configurations.",
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"maximum_number": schema.Int32Attribute{
								MarkdownDescription: "Maximum number of volumes.",
								Computed:            true,
							},
							"maximum_size": schema.Int32Attribute{
								MarkdownDescription: "Maximum volume size in GB.",
								Computed:            true,
							},
							"minimum_number": schema.Int32Attribute{
								MarkdownDescription: "Minimum number of volumes.",
								Computed:            true,
							},
							"minimum_size": schema.Int32Attribute{
								MarkdownDescription: "Minimum volume size in GB.",
								Computed:            true,
							},
							"volume_parameter_type": schema.StringAttribute{
								MarkdownDescription: "Volume parameter type (SSD, MAGNETIC).",
								Computed:            true,
							},
						},
					},
				},
				"properties": schema.MapAttribute{
					MarkdownDescription: "VM type properties.",
					Computed:            true,
					ElementType:         types.StringType,
				},
			},
		},
	}
}
