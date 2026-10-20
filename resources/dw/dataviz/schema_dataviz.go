// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package dataviz

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

var datavizSchema = schema.Schema{
	MarkdownDescription: "Cloudera Data Warehouse (CDW) integrates [Data Visualization](https://docs.cloudera.com/data-warehouse/cloud/managing-warehouses/topics/dw-use-data-visualization.html) for building graphic representations of data, dashboards, and visual applications based on CDW data.",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"cluster_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The id of the CDW Cluster which the Data Visualization is attached to.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The name of the Data Visualization.",
		},

		"image_version": schema.StringAttribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "The version of the Data Visualization.",
		},
		"resource_template": schema.StringAttribute{
			Optional: true,
			Computed: true,
			// TODO vcsomor add validation logic to the allowed types
			MarkdownDescription: "The name of the resource template being used. Available options: viz-default, viz-low, viz-medium, viz-large. Empty means the default resources template will be assigned.",
		},

		"user_groups": schema.ListAttribute{
			Required:            true,
			ElementType:         types.StringType,
			MarkdownDescription: "List of the LDAP groups which have access to this Data Visualization instance. It might be an empty list.",
		},
		"admin_groups": schema.ListAttribute{
			Required:            true,
			ElementType:         types.StringType,
			MarkdownDescription: "List of the LDAP groups which can administer this Data Visualization instance. At least one valid group is required.",
		},

		// TODO vcsomor add missing Tags to the API

		"last_updated": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Timestamp of the last Terraform update of the order.",
		},
		"status": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The status of the Data Visualization.",
		},
		"polling_options": utils.PollingOptionsSchema(20),
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
