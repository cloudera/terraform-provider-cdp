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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

func (r *awsDatahubResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	attr := map[string]schema.Attribute{}
	utils.Append(attr, generalAttributes)
	utils.Append(attr, instanceGroupSchemaAttributes)
	utils.Append(attr, map[string]schema.Attribute{
		"cluster_template": schema.StringAttribute{
			MarkdownDescription: "The name of the cluster template.",
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("instance_group")),
			},
		},
		"cluster_definition": schema.StringAttribute{
			MarkdownDescription: "The name of the cluster definition.",
			Optional:            true,
		},
		"environment": schema.StringAttribute{
			MarkdownDescription: "The name of the environment where the cluster will belong to.",
			Required:            true,
		},
	})
	resp.Schema = schema.Schema{
		MarkdownDescription: "Creates an AWS Data hub cluster.",
		Attributes:          attr,
	}
}
