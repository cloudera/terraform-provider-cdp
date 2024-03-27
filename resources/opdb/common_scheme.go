// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package opdb

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var generalAttributes = map[string]schema.Attribute{
	"polling_options": schema.SingleNestedAttribute{
		MarkdownDescription: "Polling related configuration options that could specify various values that will be used during CDP resource creation.",
		Optional:            true,
		Attributes: map[string]schema.Attribute{
			"polling_timeout": schema.Int64Attribute{
				MarkdownDescription: "Timeout value in minutes that specifies for how long should the polling go for resource creation/deletion.",
				Default:             int64default.StaticInt64(60),
				Computed:            true,
				Optional:            true,
			},
		},
	},
	"crn": schema.StringAttribute{
		MarkdownDescription: "The CRN of the cluster.",
		Computed:            true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"status": schema.StringAttribute{
		MarkdownDescription: "The last known state of the cluster",
		Computed:            true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
}
