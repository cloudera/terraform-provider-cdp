// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datalake

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var DatalakeListSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"datalakes": schema.SetNestedAttribute{
			Computed:            true,
			MarkdownDescription: "The list of existing datalakes with reduced information of them.",
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						MarkdownDescription: "The name of the datalake.",
						Computed:            true,
					},
					"crn": schema.StringAttribute{
						MarkdownDescription: "The CRN of the datalake.",
						Computed:            true,
					},
					"status": schema.StringAttribute{
						MarkdownDescription: "The status of the datalake.",
						Computed:            true,
					},
					"environment_crn": schema.StringAttribute{
						MarkdownDescription: "The CRN of the environment where the datalake belongs.",
						Computed:            true,
					},
				},
			},
		},
	},
}
