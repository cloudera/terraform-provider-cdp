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

import "github.com/hashicorp/terraform-plugin-framework/datasource/schema"

var KeytabSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"environment": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The name or CRN of the environment.",
		},
		"actor_crn": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "The CRN of the user or machine user to retrieve the keytab for. If it is not included, it defaults to the user making the request.",
		},
		"keytab": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The contents of the keytab encoded as a base64 string.",
		},
	},
}
