// Copyright 2023 Cloudera. All Rights Reserved.
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
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var ProxyConfigurationSchema = schema.Schema{
	MarkdownDescription: "",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Required: true,
		},
		"description": schema.StringAttribute{
			Optional: true,
		},
		"protocol": schema.StringAttribute{
			Required: true,
		},
		"host": schema.StringAttribute{
			Required: true,
		},
		"port": schema.Int32Attribute{
			Required: true,
		},
		"no_proxy_hosts": schema.SetAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"user": schema.StringAttribute{
			Optional: true,
		},
		"password": schema.StringAttribute{
			Optional: true,
		},
	},
}
