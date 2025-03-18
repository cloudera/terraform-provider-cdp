// Copyright 2025 Cloudera. All Rights Reserved.
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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var IDBrokerMappingSchema = schema.Schema{
	MarkdownDescription: "To enable your CDP user to utilize the central authentication features CDP provides and to exchange credentials for AWS or Azure access tokens, you have to map this CDP user to the correct IAM role or Azure Managed Service Identity (MSI).",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"data_access_role": schema.StringAttribute{
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"environment_name": schema.StringAttribute{
			Required: true,
		},
		"environment_crn": schema.StringAttribute{
			Required: true,
		},
		"mappings": schema.SetNestedAttribute{
			Optional: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"accessor_crn": schema.StringAttribute{
						Required: true,
					},
					"role": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
		"ranger_audit_role": schema.StringAttribute{
			Required: true,
		},
		"ranger_cloud_access_authorizer_role": schema.StringAttribute{
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"set_empty_mappings": schema.BoolAttribute{
			Optional: true,
		},
		"mappings_version": schema.Int64Attribute{
			Computed: true,
		},
	},
}
