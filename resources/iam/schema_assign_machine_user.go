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
)

var AssignMachineUserSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"machine_user": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The machine user the role is assigned to. Can be the machine user’s name or CRN.",
		},
		"role": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The role to assign to the machine user. Can be the role’s name or CRN.",
		},
	},
}
