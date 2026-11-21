// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package freeipaverticalscaling

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/resources/environments/platformtypes"
)

var freeipaVMTypesSchema = schema.Schema{
	MarkdownDescription: "Lists the available VM types for FreeIPA vertical scaling.",
	Attributes: map[string]schema.Attribute{
		"environment_crn": schema.StringAttribute{
			MarkdownDescription: "The CRN of the environment for which to list the available VM types.",
			Required:            true,
		},
		"architecture": schema.StringAttribute{
			MarkdownDescription: "The architecture to filter VM types. For example, x86_64.",
			Optional:            true,
		},
		"availability_zones": schema.ListAttribute{
			MarkdownDescription: "The availability zones to filter VM types.",
			Optional:            true,
			ElementType:         types.StringType,
		},
		"instance_type": schema.StringAttribute{
			MarkdownDescription: "The current instance type to filter compatible VM types.",
			Optional:            true,
		},
		"resource_type": schema.StringAttribute{
			MarkdownDescription: "The resource type to filter VM types. For example, DEFAULT.",
			Optional:            true,
		},
		"vm_types": platformtypes.VMTypesNestedAttribute(),
	},
}
