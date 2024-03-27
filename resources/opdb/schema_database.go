// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package opdb

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

func (r *databaseResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	attr := map[string]schema.Attribute{}
	utils.Append(attr, generalAttributes)
	utils.Append(attr, map[string]schema.Attribute{
		"database_name": schema.StringAttribute{
			MarkdownDescription: "The name of the database.",
			Required:            true,
		},
		"environment_name": schema.StringAttribute{
			MarkdownDescription: "The name of the environment where the cluster will belong to.",
			Required:            true,
		},
		"scale_type": schema.StringAttribute{
			MarkdownDescription: "Scale type, MICRO, LIGHT or HEAVY",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("LIGHT"),
		},
		"storage_type": schema.StringAttribute{
			MarkdownDescription: "Storage type for clusters, CLOUD_WITH_EPHEMERAL, CLOUD or HDFS",
			Optional:            true,
		},
		"disable_external_db": schema.BoolAttribute{
			MarkdownDescription: "Disable external database creation or not, it is only available in the BETA cdpcli",
			Optional:            true,
		},
		"disable_multi_az": schema.BoolAttribute{
			MarkdownDescription: "Disable deployment to multiple availability zones or not",
			Optional:            true,
		},
		"subnet_id": schema.StringAttribute{
			MarkdownDescription: "ID of the subnet to deploy to",
			Optional:            true,
		},
		"num_edge_nodes": schema.Int64Attribute{
			MarkdownDescription: "Number of edge nodes",
			Optional:            true,
			Computed:            true,
			Default:             int64default.StaticInt64(0),
		},
		"java_version": schema.Int64Attribute{
			MarkdownDescription: "Java version, it is only available in the BETA cdpcli",
			Optional:            true,
		},

		"storage_location": schema.StringAttribute{
			MarkdownDescription: "Storage Location for OPDB",
			Computed:            true,
		},
	})
	resp.Schema = schema.Schema{
		MarkdownDescription: "Creates an Operational DataBase.",
		Attributes:          attr,
	}
}
