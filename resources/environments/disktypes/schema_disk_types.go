// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package disktypes

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var diskTypesSchema = schema.Schema{
	MarkdownDescription: "Lists available disk types for all cloud platforms.",
	Attributes: map[string]schema.Attribute{
		"platform_disk_types": schema.ListNestedAttribute{
			MarkdownDescription: "Disk type information grouped by platform.",
			Computed:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"platform": schema.StringAttribute{
						MarkdownDescription: "Cloud platform (AWS, AZURE, GCP).",
						Computed:            true,
					},
					"default_disk_type": schema.StringAttribute{
						MarkdownDescription: "Default disk type for this platform.",
						Computed:            true,
					},
					"disk_types": schema.ListAttribute{
						MarkdownDescription: "Supported disk types for this platform.",
						Computed:            true,
						ElementType:         types.StringType,
					},
					"display_names": schema.MapAttribute{
						MarkdownDescription: "Disk type to display name mapping.",
						Computed:            true,
						ElementType:         types.StringType,
					},
					"disk_mappings": schema.MapAttribute{
						MarkdownDescription: "Disk type to volume parameter type mapping.",
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	},
}
