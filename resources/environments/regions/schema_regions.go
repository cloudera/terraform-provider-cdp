// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package regions

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var regionsSchema = schema.Schema{
	MarkdownDescription: "Lists the available regions for a cloud credential or environment.",
	Attributes: map[string]schema.Attribute{
		"credential": schema.StringAttribute{
			MarkdownDescription: "Name or CRN of the credential. Either this or environment_crn must be provided.",
			Optional:            true,
		},
		"environment_crn": schema.StringAttribute{
			MarkdownDescription: "CRN of the environment. Either this or credential must be provided.",
			Optional:            true,
		},
		"region": schema.StringAttribute{
			MarkdownDescription: "Region filter.",
			Optional:            true,
		},
		"availability_zone": schema.StringAttribute{
			MarkdownDescription: "Availability zone filter.",
			Optional:            true,
		},
		"availability_zones_needed": schema.BoolAttribute{
			MarkdownDescription: "Whether to include availability zone information in the response.",
			Optional:            true,
		},
		"platform_variant": schema.StringAttribute{
			MarkdownDescription: "Cloud platform variant.",
			Optional:            true,
		},
		"names": schema.ListAttribute{
			MarkdownDescription: "Available region codes.",
			Computed:            true,
			ElementType:         types.StringType,
		},
		"default_region": schema.StringAttribute{
			MarkdownDescription: "Default region.",
			Computed:            true,
		},
		"display_names": schema.MapAttribute{
			MarkdownDescription: "Region codes mapped to display names.",
			Computed:            true,
			ElementType:         types.StringType,
		},
		"availability_zones_result": schema.ListNestedAttribute{
			MarkdownDescription: "Availability zones per region.",
			Computed:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"region": schema.StringAttribute{
						MarkdownDescription: "Region code.",
						Computed:            true,
					},
					"availability_zones": schema.ListAttribute{
						MarkdownDescription: "Availability zones within this region.",
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	},
}
