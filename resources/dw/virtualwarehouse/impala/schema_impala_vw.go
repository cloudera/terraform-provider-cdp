// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package impala

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var impalaSchema = schema.Schema{
	MarkdownDescription: "A Impala Virtual Warehouse is service which is able to run low-latency SQL queries.",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"cluster_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The id of the CDW Cluster which the Impala Virtual Warehouse is attached to.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"database_catalog_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The id of the Database Catalog which the Impala Virtual Warehouse is attached to.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The name of the Impala Virtual Warehouse.",
		},
		"last_updated": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Timestamp of the last Terraform update of the order.",
		},
		"status": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The status of the database catalog.",
		},
		"image_version": schema.StringAttribute{
			Computed:            true,
			Optional:            true,
			MarkdownDescription: "Image version of the impala.",
		},
		"instance_type": schema.StringAttribute{
			Computed:            true,
			Optional:            true,
			MarkdownDescription: "The instance type for the Impala Virtual Warehouse.",
		},
		"tshirt_size": schema.StringAttribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "T-shirt size of Impala.",
		},
		"node_count": schema.Int32Attribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "Node count of Impala.",
		},
		"availability_zone": schema.StringAttribute{
			Computed:            true,
			Optional:            true,
			MarkdownDescription: "The availability zone for the Impala Virtual Warehouse.",
		},
		"enable_unified_analytics": schema.BoolAttribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "Flag to enable unified analytics.",
		},
		"impala_options": schema.SingleNestedAttribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "Impala-specific configuration options.",
			Attributes: map[string]schema.Attribute{
				"scratch_space_limit": schema.Int32Attribute{
					Optional:            true,
					Computed:            true,
					MarkdownDescription: "Defines the limit for scratch space in GiB needed by Impala for spilling queries. Valid values depend on the platform (AWS or Azure). If set, 'spillToS3Uri' cannot be set.",
				},
				"spill_to_s3_uri": schema.StringAttribute{
					Optional:            true,
					Computed:            true,
					MarkdownDescription: "Set S3 URI in 's3://bucket/path' format to enable spilling to S3. If set, 'scratchSpaceLimit' cannot be set. Not supported on Azure.",
				},
			},
		},
		"impala_ha_settings": schema.SingleNestedAttribute{
			Computed:            true,
			Optional:            true,
			MarkdownDescription: "High availability settings for Impala.",
			Attributes: map[string]schema.Attribute{
				"high_availability_mode": schema.StringAttribute{
					Optional:            true,
					Computed:            true,
					MarkdownDescription: "High Availability mode: DISABLED, ACTIVE_PASSIVE, or ACTIVE_ACTIVE.",
				},
				"enable_shutdown_of_coordinator": schema.BoolAttribute{
					Optional:            true,
					Computed:            true,
					MarkdownDescription: "Enables the shutdown of the coordinator.",
				},
				"shutdown_of_coordinator_delay_secs": schema.Int32Attribute{
					Optional:            true,
					Computed:            true,
					MarkdownDescription: "Delay in seconds before shutting down the coordinator.",
				},
				"num_of_active_coordinators": schema.Int32Attribute{
					Optional:            true,
					Computed:            true,
					MarkdownDescription: "Number of active coordinators.",
				},
				"enable_catalog_high_availability": schema.BoolAttribute{
					Optional:            true,
					Computed:            true,
					MarkdownDescription: "Enables high availability for Impala catalog.",
				},
				"enable_statestore_high_availability": schema.BoolAttribute{
					Optional:            true,
					Computed:            true,
					MarkdownDescription: "Enables high availability for Impala Statestore.",
				},
			},
		},
		"autoscaling": schema.SingleNestedAttribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "Autoscaling configuration options.",
			Attributes: map[string]schema.Attribute{
				"min_clusters": schema.Int32Attribute{
					Computed:            true,
					Optional:            true,
					MarkdownDescription: "Minimum number of available compute groups. Default: 0.",
				},
				"max_clusters": schema.Int32Attribute{
					Computed:            true,
					Optional:            true,
					MarkdownDescription: "Maximum number of available compute groups. Default: 0.",
				},
				"disable_auto_suspend": schema.BoolAttribute{
					Computed:            true,
					Optional:            true,
					MarkdownDescription: "Disable auto-suspend for the Virtual Warehouse.",
				},
				"auto_suspend_timeout_seconds": schema.Int32Attribute{
					Computed:            true,
					Optional:            true,
					MarkdownDescription: "Threshold for auto-suspend in seconds.",
				},
				"impala_scale_up_delay_seconds": schema.Int32Attribute{
					Computed:            true,
					Optional:            true,
					MarkdownDescription: "Scale-up threshold in seconds for Impala.",
				},
				"impala_scale_down_delay_seconds": schema.Int32Attribute{
					Computed:            true,
					Optional:            true,
					MarkdownDescription: "Scale-down threshold in seconds for Impala.",
				},
				// TODO Prateek
				// Come back to enable these once you fix these in Impala APIs
				// The problem here is that the computed flags in autoscaling can't be partial
				// You have to send  either all or none
				/*"impala_shutdown_of_coordinator_delay_seconds": schema.Int32Attribute{
					Optional:            true,
					MarkdownDescription: "Delay in seconds before shutting down Impala coordinator. Deprecated.",
				},
				"impala_num_of_active_coordinators": schema.Int32Attribute{
					Optional:            true,
					MarkdownDescription: "Number of active Impala coordinators. Deprecated.",
				},*/
				/*"impala_executor_group_sets": schema.ListNestedAttribute{
					Optional:            true,
					Computed:            false,
					MarkdownDescription: "Reconfigure executor group sets for workload-aware autoscaling.",
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"small": schema.SingleNestedAttribute{
								Optional:            true,
								MarkdownDescription: "Configure small executor group set.",
								Attributes: map[string]schema.Attribute{
									"exec_group_size": schema.Int32Attribute{
										Required:            true,
										MarkdownDescription: "Number of executors per executor group.",
									},
									"min_executor_groups": schema.Int32Attribute{
										Required:            true,
										MarkdownDescription: "Minimum number of executor groups.",
									},
									"max_executor_groups": schema.Int32Attribute{
										Required:            true,
										MarkdownDescription: "Maximum number of executor groups.",
									},
									"auto_suspend_timeout_seconds": schema.NumberAttribute{
										Optional:            true,
										MarkdownDescription: "Auto suspend timeout seconds.",
									},
									"disable_auto_suspend": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Disable auto suspend.",
									},
									"trigger_scale_up_delay": schema.Int32Attribute{
										Optional:            true,
										MarkdownDescription: "Scale-up delay trigger in seconds.",
									},
									"trigger_scale_down_delay": schema.Int32Attribute{
										Optional:            true,
										MarkdownDescription: "Scale-down delay trigger in seconds.",
									},
								},
							},
							"custom1": schema.SingleNestedAttribute{
								Optional:            true,
								MarkdownDescription: "Configure first custom executor group set.",
								Attributes:          map[string]schema.Attribute{
									// Same as 'small' executor group, define attributes here.
								},
							},
							"custom2": schema.SingleNestedAttribute{
								Optional:            true,
								MarkdownDescription: "Configure second custom executor group set.",
								Attributes: map[string]schema.Attribute{
									"exec_group_size": schema.Int32Attribute{
										Required:            true,
										MarkdownDescription: "Number of executors per executor group.",
									},
									"min_executor_groups": schema.Int32Attribute{
										Required:            true,
										MarkdownDescription: "Minimum number of executor groups.",
									},
									"max_executor_groups": schema.Int32Attribute{
										Required:            true,
										MarkdownDescription: "Maximum number of executor groups.",
									},
									"auto_suspend_timeout_seconds": schema.NumberAttribute{
										Optional:            true,
										MarkdownDescription: "Auto suspend timeout seconds.",
									},
									"disable_auto_suspend": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Disable auto suspend.",
									},
									"trigger_scale_up_delay": schema.Int32Attribute{
										Optional:            true,
										MarkdownDescription: "Scale-up delay trigger in seconds.",
									},
									"trigger_scale_down_delay": schema.Int32Attribute{
										Optional:            true,
										MarkdownDescription: "Scale-down delay trigger in seconds.",
									},
								},
							},
							"custom3": schema.SingleNestedAttribute{
								Optional:            true,
								MarkdownDescription: "Configure third custom executor group set.",
								Attributes: map[string]schema.Attribute{
									"exec_group_size": schema.Int32Attribute{
										Required:            true,
										MarkdownDescription: "Number of executors per executor group.",
									},
									"min_executor_groups": schema.Int32Attribute{
										Required:            true,
										MarkdownDescription: "Minimum number of executor groups.",
									},
									"max_executor_groups": schema.Int32Attribute{
										Required:            true,
										MarkdownDescription: "Maximum number of executor groups.",
									},
									"auto_suspend_timeout_seconds": schema.NumberAttribute{
										Optional:            true,
										MarkdownDescription: "Auto suspend timeout seconds.",
									},
									"disable_auto_suspend": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Disable auto suspend.",
									},
									"trigger_scale_up_delay": schema.Int32Attribute{
										Optional:            true,
										MarkdownDescription: "Scale-up delay trigger in seconds.",
									},
									"trigger_scale_down_delay": schema.Int32Attribute{
										Optional:            true,
										MarkdownDescription: "Scale-down delay trigger in seconds.",
									},
								},
							},
							"large": schema.SingleNestedAttribute{
								Optional:            true,
								MarkdownDescription: "Configure large executor group set.",
								Attributes: map[string]schema.Attribute{
									"exec_group_size": schema.Int32Attribute{
										Required:            true,
										MarkdownDescription: "Number of executors per executor group.",
									},
									"min_executor_groups": schema.Int32Attribute{
										Required:            true,
										MarkdownDescription: "Minimum number of executor groups.",
									},
									"max_executor_groups": schema.Int32Attribute{
										Required:            true,
										MarkdownDescription: "Maximum number of executor groups.",
									},
									"auto_suspend_timeout_seconds": schema.NumberAttribute{
										Optional:            true,
										MarkdownDescription: "Auto suspend timeout seconds.",
									},
									"disable_auto_suspend": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Disable auto suspend.",
									},
									"trigger_scale_up_delay": schema.Int32Attribute{
										Optional:            true,
										MarkdownDescription: "Scale-up delay trigger in seconds.",
									},
									"trigger_scale_down_delay": schema.Int32Attribute{
										Optional:            true,
										MarkdownDescription: "Scale-down delay trigger in seconds.",
									},
								},
							},
						},
					},
				},*/
			},
		},
		/*"config": schema.SingleNestedAttribute{
			Optional:            true,
			MarkdownDescription: "Service configuration settings.",
			Attributes: map[string]schema.Attribute{
				"common_configs": schema.SingleNestedAttribute{
					Optional:            true,
					MarkdownDescription: "Common configuration settings.",
					Attributes: map[string]schema.Attribute{
						"config_blocks": schema.ListNestedAttribute{
							Optional:            true,
							MarkdownDescription: "List of config blocks.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Required: true,
									},
									"value": schema.StringAttribute{
										Required: true,
									},
									"type": schema.StringAttribute{
										Required: true,
									},
									"enabled": schema.BoolAttribute{
										Required: true,
									},
								},
							},
						},
					},
				},
				"application_configs": schema.MapNestedAttribute{
					Optional:            true,
					MarkdownDescription: "Map of application-specific configuration settings.",
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"config_blocks": schema.ListNestedAttribute{
								Optional:            true,
								MarkdownDescription: "List of config blocks for application configurations.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Required: true,
										},
										"value": schema.StringAttribute{
											Required: true,
										},
										"type": schema.StringAttribute{
											Required: true,
										},
										"enabled": schema.BoolAttribute{
											Required: true,
										},
									},
								},
							},
						},
					},
				},
				"ldap_groups": schema.ListAttribute{
					Optional:            true,
					MarkdownDescription: "LDAP groups for SSO configuration.",
					ElementType:         types.StringType,
				},
				"enable_sso": schema.BoolAttribute{
					Optional:            true,
					MarkdownDescription: "Enable Single Sign-On (SSO).",
				},
			},
		},*/
		"query_isolation_options": schema.SingleNestedAttribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "Query isolation options for Impala.",
			Attributes: map[string]schema.Attribute{
				"max_queries": schema.Int32Attribute{
					Optional:            true,
					Computed:            true,
					MarkdownDescription: "Maximum number of queries for isolation. Default: 0 disables isolation.",
				},
				"max_nodes_per_query": schema.Int32Attribute{
					Computed:            true,
					Optional:            true,
					MarkdownDescription: "Maximum number of nodes per query for isolation. Default: 0 disables isolation.",
				},
			},
		},
		"tags": schema.ListNestedAttribute{
			Optional:            true,
			MarkdownDescription: "Tags associated with the Impala Virtual Warehouse.",
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
		"resource_pool": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "Resource pool for the Impala Virtual Warehouse.",
		},
		"hive_authentication_mode": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "Hive authentication mode.",
		},
		"platform_jwt_auth": schema.BoolAttribute{
			Optional:            true,
			MarkdownDescription: "Platform JWT authentication flag.",
		},
		"impala_query_log": schema.BoolAttribute{
			Computed:            true,
			Optional:            true,
			MarkdownDescription: "Enable or disable Impala query logging.",
		},
		"ebs_llap_spill_gb": schema.Int64Attribute{
			Optional:            true,
			MarkdownDescription: "EBS LLAP spill size in GB.",
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
