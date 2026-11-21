// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package networks

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var networksSchema = schema.Schema{
	MarkdownDescription: "Lists the available networks for a cloud credential.",
	Attributes: map[string]schema.Attribute{
		"credential": schema.StringAttribute{
			MarkdownDescription: "Name or CRN of the credential.",
			Required:            true,
		},
		"region": schema.StringAttribute{
			MarkdownDescription: "Region filter.",
			Optional:            true,
		},
		"availability_zone": schema.StringAttribute{
			MarkdownDescription: "Availability zone filter.",
			Optional:            true,
		},
		"platform_variant": schema.StringAttribute{
			MarkdownDescription: "Cloud platform variant.",
			Optional:            true,
		},
		"network_id": schema.StringAttribute{
			MarkdownDescription: "Network ID filter.",
			Optional:            true,
		},
		"shared_project_id": schema.StringAttribute{
			MarkdownDescription: "GCP shared project ID.",
			Optional:            true,
		},
		"subnet_ids": schema.StringAttribute{
			MarkdownDescription: "Comma-separated subnet IDs to filter by.",
			Optional:            true,
		},
		"networks": schema.ListNestedAttribute{
			MarkdownDescription: "Networks grouped by platform.",
			Computed:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"platform": schema.StringAttribute{
						MarkdownDescription: "Cloud platform (AWS, AZURE, GCP).",
						Computed:            true,
					},
					"networks": schema.ListNestedAttribute{
						MarkdownDescription: "Available networks for this platform.",
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									MarkdownDescription: "Network ID.",
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: "Network name.",
									Computed:            true,
								},
								"properties": schema.MapAttribute{
									MarkdownDescription: "Platform-specific properties.",
									Computed:            true,
									ElementType:         types.StringType,
								},
								"subnets": schema.MapAttribute{
									MarkdownDescription: "Subnet ID to subnet name mapping.",
									Computed:            true,
									ElementType:         types.StringType,
								},
							},
						},
					},
				},
			},
		},
	},
}
