// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package vmtypes

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/cloudera/terraform-provider-cdp/resources/environments/platformtypes"
)

var vmTypesSchema = schema.Schema{
	MarkdownDescription: "Lists the available VM types for a cloud credential.",
	Attributes: map[string]schema.Attribute{
		"credential_name": schema.StringAttribute{
			MarkdownDescription: "Name of the credential. Either this or credential_crn must be provided.",
			Optional:            true,
		},
		"credential_crn": schema.StringAttribute{
			MarkdownDescription: "CRN of the credential. Either this or credential_name must be provided.",
			Optional:            true,
		},
		"region": schema.StringAttribute{
			MarkdownDescription: "Region to list VM types for.",
			Optional:            true,
		},
		"availability_zone": schema.StringAttribute{
			MarkdownDescription: "Availability zone to filter VM types for.",
			Optional:            true,
		},
		"architecture": schema.StringAttribute{
			MarkdownDescription: "CPU architecture filter (e.g. x86_64).",
			Optional:            true,
		},
		"platform_variant": schema.StringAttribute{
			MarkdownDescription: "Cloud platform variant.",
			Optional:            true,
		},
		"resource_type": schema.StringAttribute{
			MarkdownDescription: "Resource type filter.",
			Optional:            true,
		},
		"vm_types": platformtypes.VMTypesNestedAttribute(),
	},
}
