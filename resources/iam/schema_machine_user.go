// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package iam

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var machineUserSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"name": schema.StringAttribute{
			Required: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"status": schema.StringAttribute{
			Computed: true,
		},
		"workload_username": schema.StringAttribute{
			Computed: true,
		},
		"workload_password": schema.StringAttribute{
			Optional:  true,
			Sensitive: true,
		},
		"creation_date": schema.StringAttribute{
			Computed: true,
		},
		"workload_password_details": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"is_password_set": schema.BoolAttribute{
					Computed: true,
				},
				"expiration_date": schema.StringAttribute{
					Computed: true,
				},
				"min_lifetime_date": schema.StringAttribute{
					Computed: true,
				},
			},
		},
		"azure_cloud_identities": schema.SetNestedAttribute{
			Computed: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"environment_crn": schema.StringAttribute{
						Computed: true,
					},
					"object_id": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	},
}
