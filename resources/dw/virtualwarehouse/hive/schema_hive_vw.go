// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package hive

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var hiveSchema = schema.Schema{
	MarkdownDescription: "A Hive Virtual Warehouse is service which is able to run big SQL queries.",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"cluster_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The id of the CDW Cluster which the Hive Virtual Warehouse is attached to.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"database_catalog_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The id of the Database Catalog which the Hive Virtual Warehouse is attached to.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The name of the Hive Virtual Warehouse.",
		},
		"image_version": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The version of the Hive Virtual Warehouse image.",
		},
		"node_count": schema.Int64Attribute{
			Optional:            true,
			MarkdownDescription: "Nodes per compute cluster. If specified, forces ‘template’ to be ‘custom’.",
		},
		"platform_jwt_auth": schema.BoolAttribute{
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
			MarkdownDescription: "Value of ‘true’ automatically configures the Virtual Warehouse to support JWTs issued by the CDP JWT token provider. Value of ‘false’ does not enable JWT auth on the Virtual Warehouse. If this field is not specified, it defaults to ‘false’.",
		},
		"ldap_groups": schema.ListAttribute{
			Optional:            true,
			ElementType:         types.StringType,
			MarkdownDescription: "LDAP group names to be enabled for auth.",
		},
		"enable_sso": schema.BoolAttribute{
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
			MarkdownDescription: "Enable SSO for the Virtual Warehouse. If this field is not specified, it defaults to ‘false’.",
		},
		"autoscaling": schema.SingleNestedAttribute{
			MarkdownDescription: "Autoscaling related configuration options that could specify various values that will be used during CDW resource creation.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"min_clusters": schema.Int64Attribute{
					Required:            true,
					MarkdownDescription: "Minimum number of available compute groups.",
				},
				"max_clusters": schema.Int64Attribute{
					Required:            true,
					MarkdownDescription: "Maximum number of available compute groups.",
				},
				"disable_auto_suspend": schema.BoolAttribute{
					Required:            true,
					MarkdownDescription: "Boolean value that specifies if auto-suspend should be disabled.",
				},
				"auto_suspend_timeout_seconds": schema.Int64Attribute{
					Required:            true,
					MarkdownDescription: "The time in seconds after which the compute group should be suspended.",
				},
				"hive_scale_wait_time_seconds": schema.Int64Attribute{
					Optional:            true,
					MarkdownDescription: "Set wait time before a scale event happens. Either “hiveScaleWaitTimeSeconds” or “hiveDesiredFreeCapacity” can be provided.",
				},
				"hive_desired_free_capacity": schema.Int64Attribute{
					Optional:            true,
					MarkdownDescription: "Set Desired free capacity. Either “hiveScaleWaitTimeSeconds” or “hiveDesiredFreeCapacity” can be provided.",
				},
			},
		},
		"aws_options": schema.SingleNestedAttribute{
			MarkdownDescription: "AWS related configuration options that could specify various values that will be used during CDW resource creation.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"availability_zone": schema.StringAttribute{
					Optional:            true,
					Computed:            true,
					MarkdownDescription: "This feature works only for AWS cluster type. An availability zone to host compute instances. If not specified, defaults to a randomly selected availability zone inferred from available subnets.",
				},
				"ebs_llap_spill_gb": schema.Int64Attribute{
					Optional:            true,
					MarkdownDescription: "This feature works only for AWS cluster type. The size of the EBS volume in GB to be used for LLAP spill storage. If not specified, defaults to no extra spill disk.",
				},
				"tags": schema.ListAttribute{
					Optional:            true,
					ElementType:         types.StringType,
					MarkdownDescription: "This feature works only for AWS cluster type. Tags to be applied to the underlying compute nodes.",
				},
			},
		},
		"query_isolation_options": schema.SingleNestedAttribute{
			MarkdownDescription: "Query isolation related configuration options.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"max_queries": schema.Int64Attribute{
					Optional:            true,
					MarkdownDescription: "Maximum number of concurrent isolated queries. If not provided, 0 will be applied. The 0 value means the query isolation functionality will be disabled.",
				},
				"max_nodes_per_query": schema.Int64Attribute{
					Optional:            true,
					MarkdownDescription: "Maximum number of nodes per isolated query. If not provided, 0 will be applied. The 0 value means the query isolation functionality will be disabled.",
				},
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
		"polling_options": schema.SingleNestedAttribute{
			MarkdownDescription: "Polling related configuration options that could specify various values that will be used during CDP resource creation.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"async": schema.BoolAttribute{
					MarkdownDescription: "Boolean value that specifies if Terraform should wait for resource creation/deletion.",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
					PlanModifiers: []planmodifier.Bool{
						boolplanmodifier.UseStateForUnknown(),
					},
				},
				"polling_timeout": schema.Int64Attribute{
					MarkdownDescription: "Timeout value in minutes that specifies for how long should the polling go for resource creation/deletion.",
					Default:             int64default.StaticInt64(40),
					Computed:            true,
					Optional:            true,
				},
				"call_failure_threshold": schema.Int64Attribute{
					MarkdownDescription: "Threshold value that specifies how many times should a single call failure happen before giving up the polling.",
					Default:             int64default.StaticInt64(3),
					Computed:            true,
					Optional:            true,
				},
			},
		},
	},
}
