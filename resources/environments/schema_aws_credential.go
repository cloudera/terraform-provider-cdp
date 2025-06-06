// Copyright 2024 Cloudera. All Rights Reserved.
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
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

func (r *awsCredentialResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	attr := map[string]schema.Attribute{}
	utils.Append(attr, generalCredentialSchema)
	utils.Append(attr, map[string]schema.Attribute{
		"role_arn": schema.StringAttribute{
			Required: true,
		},
		"skip_org_policy_decisions": schema.BoolAttribute{
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
			MarkdownDescription: "Whether to skip organizational policy decision checks or not.",
		},
		"verify_permissions": schema.BoolAttribute{
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
			MarkdownDescription: "Whether to verify permissions upon saving or not.",
		},
	})
	resp.Schema = schema.Schema{
		MarkdownDescription: "The AWS credential is used for authorization to provision resources such as compute instances within your cloud provider account.",
		Attributes:          attr,
	}
}
