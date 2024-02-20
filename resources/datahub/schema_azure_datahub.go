// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datahub

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

func (r *azureDatahubResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	attr := map[string]schema.Attribute{}
	utils.Append(attr, generalAttributes)
	utils.Append(attr, instanceGroupSchemaAttributes)
	utils.Append(attr, map[string]schema.Attribute{
		"cluster_template": schema.StringAttribute{
			MarkdownDescription: "The name of the cluster template.",
			Required:            true,
		},
		"cluster_definition": schema.StringAttribute{
			MarkdownDescription: "The name of the cluster definition.",
			Required:            true,
		},
		"environment": schema.StringAttribute{
			MarkdownDescription: "The name of the environment where the cluster will belong to.",
			Required:            true,
		},
		"database_type": schema.StringAttribute{
			Optional: true,
			Computed: false,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
	})
	resp.Schema = schema.Schema{
		MarkdownDescription: "Creates an Azure Data hub cluster.",
		Attributes:          attr,
	}
}
