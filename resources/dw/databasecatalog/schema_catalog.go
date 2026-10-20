// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package databasecatalog

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

var dwDefaultDatabaseCatalogSchema = schema.Schema{
	MarkdownDescription: "Creates an AWS Data Warehouse database catalog.",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The name of the database catalog.",
		},
		"cluster_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The unique identifier of the cluster.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"last_updated": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Timestamp of the last Terraform update of the order.",
		},
		"status": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The status of the database catalog.",
		},
		"polling_options": utils.PollingOptionsSchema(40),
		"async": schema.BoolAttribute{
			MarkdownDescription: "When set to `true`, the resource operation returns immediately without waiting for completion. When `false` (default), Terraform waits for the operation to finish, subject to the duration set in the `timeouts` block.",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
		},
		"call_failure_threshold": schema.Int64Attribute{
			MarkdownDescription: "The number of consecutive internal API call failures before the provider gives up waiting. Applies only when `async` is `false`. Defaults to `3`.",
			Optional:            true,
			Computed:            true,
			Default:             int64default.StaticInt64(3),
			Validators: []validator.Int64{
				int64validator.AtLeast(1),
			},
		},
	},
}
