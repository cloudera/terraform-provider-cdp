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
)

var AzureImageTermsPolicySchema = schema.Schema{
	MarkdownDescription: "Updates account level Azure Marketplace image policy. CDP is capable to automatically accept Azure Marketplace image terms during cluster deployment. You can use this setting in your account to opt in or opt out this behaviour.",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"accepted": schema.BoolAttribute{
			Required:    true,
			Description: "Flag to enable or disable automatic acceptance of Azure Marketplace image terms.",
		},
	},
}
