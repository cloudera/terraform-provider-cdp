// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package recipe

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var recipeSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"crn": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The name of the recipe. This name must be unique, must have between 5 and 100 characters, and must contain only lowercase letters, numbers and hyphens. Names are case-sensitive.",
		},
		"content": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The content of the recipe. It can be either plain text or a file path. Please keep in mind that if the content of the file changes then the provider will not detect the it and will not mark the recipe for recreation.",
		},
		"type": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The type of recipe. Supported values are : PRE_SERVICE_DEPLOYMENT, PRE_TERMINATION, POST_SERVICE_DEPLOYMENT, POST_CLOUDERA_MANAGER_START.",
		},
		"description": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "The description of the recipe. The description can have a maximum of 1000 characters.",
		},
	},
}
