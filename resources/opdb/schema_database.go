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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int32default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int32planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

func (r *databaseResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	attr := map[string]schema.Attribute{}
	utils.Append(attr, generalAttributes)
	utils.Append(attr, map[string]schema.Attribute{
		"database_name": schema.StringAttribute{
			MarkdownDescription: "The name of the database.",
			Required:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"environment_name": schema.StringAttribute{
			MarkdownDescription: "The name of the environment where the cluster will belong to.",
			Required:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"scale_type": schema.StringAttribute{
			MarkdownDescription: "Scale type, MICRO, LIGHT or HEAVY",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("LIGHT"),
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"storage_type": schema.StringAttribute{
			MarkdownDescription: "Storage type for clusters, CLOUD_WITH_EPHEMERAL, CLOUD or HDFS",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"disable_external_db": schema.BoolAttribute{
			MarkdownDescription: "Disable external database creation or not. It is only available in the BETA api.",
			Optional:            true,
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.RequiresReplace(),
			},
		},
		"disable_multi_az": schema.BoolAttribute{
			MarkdownDescription: "Disable deployment to multiple availability zones or not",
			Optional:            true,
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.RequiresReplace(),
			},
		},
		"subnet_id": schema.StringAttribute{
			MarkdownDescription: "ID of the subnet to deploy to",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"num_edge_nodes": schema.Int32Attribute{
			MarkdownDescription: "Number of edge nodes",
			Optional:            true,
			Computed:            true,
			Default:             int32default.StaticInt32(0),
			PlanModifiers: []planmodifier.Int32{
				int32planmodifier.RequiresReplace(),
			},
		},
		"java_version": schema.Int32Attribute{
			MarkdownDescription: "Java version. It is only available in the BETA api.",
			Optional:            true,
			PlanModifiers: []planmodifier.Int32{
				int32planmodifier.RequiresReplace(),
			},
		},

		"storage_location": schema.StringAttribute{
			MarkdownDescription: "Storage Location for OPDB. It is only available in the BETA api.",
			Computed:            true,
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			},
		},
		"auto_scaling_parameters": schema.SingleNestedAttribute{
			Optional: true,
			Attributes: map[string]schema.Attribute{
				"targeted_value_for_metric": schema.Int64Attribute{
					Optional:            true,
					MarkdownDescription: "The target value of the metric a user expect to maintain for the cluster",
				},
				"max_workers_for_database": schema.Int32Attribute{
					Optional:            true,
					MarkdownDescription: "Maximum number of worker nodes as per this metrics can be scaled up to.",
				},
				"max_workers_per_batch": schema.Int32Attribute{
					Optional:            true,
					MarkdownDescription: "Maximum number of worker nodes as per this metrics can be scaled up to in one batch.",
				},
				"min_workers_for_database": schema.Int32Attribute{
					Optional:            true,
					MarkdownDescription: "Minimum number of worker nodes as per this metrics can be scaled down to.",
				},
				"evaluation_period": schema.Int64Attribute{
					Optional:            true,
					MarkdownDescription: "Period of metrics(in seconds) needs to be considered.",
				},
				"minimum_block_cache_gb": schema.Int32Attribute{
					Optional:            true,
					MarkdownDescription: "The amount of block cache, in Gigabytes, which the database should have.",
				},
				"max_hdfs_usage_percentage": schema.Int32Attribute{
					Optional:            true,
					MarkdownDescription: "The maximum percentage of HDFS utilization for the database before we trigger the scaling. It is only available in the BETA api.",
				},
				"max_regions_per_region_server": schema.Int32Attribute{
					Optional:            true,
					MarkdownDescription: "The maximum number of regions per region server. It is only available in the BETA api.",
				},
				"max_cpu_utilization": schema.Int32Attribute{
					Optional:            true,
					MarkdownDescription: "The maximum percentage threshold for the CPU utilization of the worker nodes. The CPU utilization is obtained from the Cloudera Manager metric ‘cpu_percent’ across worker nodes. Set 100 or more to disable the CPU metrics. It is only available in the BETA api.",
				},
				"max_compute_nodes_for_database": schema.Int32Attribute{
					Optional:            true,
					MarkdownDescription: "The maximum number of compute nodes, as per these metrics, that can be scaled up to. It is only available in the BETA api.",
				},
				"min_compute_nodes_for_database": schema.Int32Attribute{
					Optional:            true,
					MarkdownDescription: "The minimum number of compute nodes, as per these metrics, that can be scaled down to. It is only available in the BETA api.",
				},
			},
		},
		"attached_storage_for_workers": schema.SingleNestedAttribute{
			Optional:            true,
			MarkdownDescription: "Attached storage for the worker nodes for AWS, Azure, and GCP cloud providers.",
			PlanModifiers: []planmodifier.Object{
				objectplanmodifier.RequiresReplace(),
			},
			Attributes: map[string]schema.Attribute{
				"volume_count": schema.Int32Attribute{
					Optional:            true,
					Computed:            true,
					MarkdownDescription: "The number of Volumes. Default is 4. Valid Range: Minimum value of 1, maximum value 8.",
					Default:             int32default.StaticInt32(4),
				},
				"volume_size": schema.Int32Attribute{
					Optional:            true,
					Computed:            true,
					MarkdownDescription: "The target size of the volume, in GiB. Default is 2048.",
					Default:             int32default.StaticInt32(2048),
				},
				"volume_type": schema.StringAttribute{
					Optional:            true,
					Computed:            true,
					MarkdownDescription: "Volume Type. HDD - Hard disk drives (HDD) volume type. Default is HDD. SSD - Solid disk drives (SSD) volume type. LOCAL_SSD - Local SSD volume type.",
					Default:             stringdefault.StaticString("HDD"),
				},
			},
		},
		"disable_kerberos": schema.BoolAttribute{
			MarkdownDescription: "Disable Kerberos authentication. ",
			Optional:            true,
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.RequiresReplace(),
			},
		},
		"disable_jwt_auth": schema.BoolAttribute{
			MarkdownDescription: "Disable OAuth Bearer (JWT) authentication scheme. ",
			Optional:            true,
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.RequiresReplace(),
			},
		},
		"image": schema.SingleNestedAttribute{
			Optional:            true,
			MarkdownDescription: "Details of an Image.",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					Required:            true,
					MarkdownDescription: "Image ID for the database.",
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					},
				},
				"catalog": schema.StringAttribute{
					Required:            true,
					MarkdownDescription: "Catalog name for the image.",
				},
			},
		},

		"enable_grafana": schema.BoolAttribute{
			MarkdownDescription: "To enable grafana server for the database.",
			Optional:            true,
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.RequiresReplace(),
			},
		},
		"custom_user_tags": schema.SetNestedAttribute{
			Optional:            true,
			MarkdownDescription: "Optional tags to apply to launched infrastructure resources",
			PlanModifiers: []planmodifier.Set{
				setplanmodifier.RequiresReplace(),
			},
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"key": schema.StringAttribute{
						Required: true,
					},
					"value": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
		"enable_region_canary": schema.BoolAttribute{
			MarkdownDescription: "To enable the region canary for the database.",
			Optional:            true,
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.RequiresReplace(),
			},
		},
		"recipes": schema.SetNestedAttribute{
			Optional:            true,
			MarkdownDescription: "Custom recipes for the database.",
			PlanModifiers: []planmodifier.Set{
				setplanmodifier.RequiresReplace(),
			},
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"names": schema.SetAttribute{
						MarkdownDescription: "The set of recipe names that are going to be applied on the given instance group.",
						ElementType:         types.StringType,
						Required:            true,
					},
					"instance_group": schema.StringAttribute{
						MarkdownDescription: "The name of the designated instance group.",
						Required:            true,
					},
				},
			},
		},
		"volume_encryptions": schema.SetNestedAttribute{
			Optional:            true,
			MarkdownDescription: "Specifies encryption key to encrypt volume for instance group. It is currently supported for AWS cloud provider only. It is only available in the BETA api.",
			PlanModifiers: []planmodifier.Set{
				setplanmodifier.RequiresReplace(),
			},
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"encryption_key": schema.StringAttribute{
						MarkdownDescription: "Encryption key to encrypt volume.",
						Required:            true,
					},
					"instance_group": schema.StringAttribute{
						MarkdownDescription: "The name of the designated instance group.",
						Required:            true,
					},
				},
			},
		},
		"architecture": schema.StringAttribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "CPU Architecture is used for the cluster",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			},
		},
	})
	resp.Schema = schema.Schema{
		MarkdownDescription: "Creates an Operational DataBase.",
		Attributes:          attr,
	}
}
