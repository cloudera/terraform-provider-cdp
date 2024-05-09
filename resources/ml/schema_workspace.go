// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package ml

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var workspaceSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"authorized_ip_ranges": schema.SetAttribute{
			MarkdownDescription: "The whitelist of CIDR blocks which can access the API server.",
			Optional:            true,
			ElementType:         types.StringType,
		},
		"cdsw_migration_mode": schema.StringAttribute{
			MarkdownDescription: "Toggle for cdsw migration preflight validation",
			Optional:            true,
		},
		"disable_tls": schema.BoolAttribute{
			MarkdownDescription: "The boolean flag to disable TLS setup for workspace. By default, the TLS is enabled.",
			Optional:            true,
		},
		"enable_governance": schema.BoolAttribute{
			MarkdownDescription: "Enables Cloudera Machine Learning governance by integrating with Cloudera Atlas. By default, this flag is disabled.",
			Optional:            true,
		},
		"enable_model_metrics": schema.BoolAttribute{
			MarkdownDescription: "Enables the model metrics service for exporting metrics for models to a metrics store.",
			Optional:            true,
		},
		"enable_monitoring": schema.BoolAttribute{
			MarkdownDescription: "The boolean flag is used to enable monitoring. By default, monitoring is disabled.",
			Optional:            true,
		},
		"environment_name": schema.StringAttribute{
			MarkdownDescription: "The environment for the workspace to create.",
			Required:            true,
		},
		"existing_database_config": schema.SingleNestedAttribute{
			MarkdownDescription: "Optional configurations for an existing Postgres to export model metrics to.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"existing_database_host": schema.StringAttribute{
					MarkdownDescription: "Optionally provide a Postgresql database host to export model metrics to.",
					Optional:            true,
				},
				"existing_database_name": schema.StringAttribute{
					MarkdownDescription: "Optionally provide a Postgresql database name to export model metrics to.",
					Optional:            true,
				},
				"existing_database_password": schema.StringAttribute{
					MarkdownDescription: "Optionally provide a Postgresql database password to use when exporting model metrics.",
					Optional:            true,
				},
				"existing_database_port": schema.StringAttribute{
					MarkdownDescription: "Optionally provide a Postgresql database port to export model metrics to.",
					Optional:            true,
				},
				"existing_database_user": schema.StringAttribute{
					MarkdownDescription: "Optionally provide a Postgresql database user to use when exporting model metrics.",
					Optional:            true,
				},
			},
		},
		"existing_nfs": schema.StringAttribute{
			MarkdownDescription: "Optionally use an existing NFS by providing the hostname and desired path (Azure and Private Cloud only).",
			Optional:            true,
		},
		"load_balancer_ip_whitelists": schema.SetAttribute{
			MarkdownDescription: "The whitelist of IPs for load balancer.",
			Optional:            true,
			ElementType:         types.StringType,
		},
		"ml_version": schema.StringAttribute{
			MarkdownDescription: "The version of ML workload app to install.",
			Optional:            true,
		},
		"nfs_version": schema.StringAttribute{
			MarkdownDescription: "The NFS Protocol version of the NFS server we are using for Azure and Private Cloud.",
			Optional:            true,
		},
		"outbound_types": schema.SetAttribute{
			MarkdownDescription: "Outbound Types provided for the workspace.",
			Optional:            true,
			ElementType:         types.StringType,
		},
		"private_cluster": schema.BoolAttribute{
			MarkdownDescription: "Whether to create a private cluster.",
			Optional:            true,
		},
		"provision_k8s_request": schema.SingleNestedAttribute{
			MarkdownDescription: "The request for Kubernetes workspace provision. Required in public cloud.",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"environment_name": schema.StringAttribute{
					MarkdownDescription: "The name of the environment for the workspace to create.",
					Required:            true,
				},
				"instance_groups": schema.SetNestedAttribute{
					MarkdownDescription: "The instance groups.",
					Required:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"autoscaling": schema.SingleNestedAttribute{
								MarkdownDescription: "The auto scaling configuration.",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										MarkdownDescription: "The boolean flag to enable the auto scaling.",
										Optional:            true,
									},
									"max_instances": schema.Int64Attribute{
										MarkdownDescription: "The maximum number of instance for auto scaling.",
										Required:            true,
									},
									"min_instances": schema.Int64Attribute{
										MarkdownDescription: "The minimum number of instance for auto scaling.",
										Required:            true,
									},
								},
							},
							"ingress_rules": schema.SetAttribute{
								MarkdownDescription: "The networking rules for the ingress.",
								Optional:            true,
								ElementType:         types.StringType,
							},
							"instance_count": schema.Int64Attribute{
								MarkdownDescription: "The initial number of instance node.",
								Optional:            true,
							},
							"instance_tier": schema.StringAttribute{
								MarkdownDescription: "The tier of the instance i.e. on-demand/spot.",
								Optional:            true,
							},
							"instance_type": schema.StringAttribute{
								MarkdownDescription: "The cloud provider instance type for the node instance.",
								Required:            true,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: "The unique name of the instance group.",
								Optional:            true,
							},
							"root_volume": schema.SingleNestedAttribute{
								MarkdownDescription: "The root volume of the instance.",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"size": schema.Int64Attribute{
										MarkdownDescription: "The volume size in GB.",
										Required:            true,
									},
								},
							},
						},
					},
				},
				"network": schema.SingleNestedAttribute{
					MarkdownDescription: "The overlay network for an AWS Kubernetes cluster's CNI.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"plugin": schema.StringAttribute{
							MarkdownDescription: "The plugin specifies specific cni vendor, ex: calico, weave etc.",
							Optional:            true,
						},
						"topology": schema.SingleNestedAttribute{
							MarkdownDescription: "The options for overlay topology.",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"subnets": schema.SetAttribute{
									MarkdownDescription: "The options for subnets.",
									ElementType:         types.StringType,
									Required:            true,
								},
							},
						},
					},
				},
				"tags": schema.SetNestedAttribute{
					MarkdownDescription: "Tags to add to the cloud provider resources created. This is in addition to any tags added by Cloudera.",
					Optional:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								MarkdownDescription: "The name for the tag.",
								Required:            true,
							},
							"value": schema.StringAttribute{
								MarkdownDescription: "The value for the tag.",
								Required:            true,
							},
						},
					},
				},
			},
		},
		"skip_validation": schema.BoolAttribute{
			MarkdownDescription: "Skip pre-flight validations if requested.",
			Optional:            true,
		},
		"static_subdomain": schema.StringAttribute{
			MarkdownDescription: "The static subdomain to be used for the workspace.",
			Optional:            true,
		},
		"subnets_for_load_balancers": schema.SetAttribute{
			MarkdownDescription: "The list of subnets used for the load balancer that CML creates.",
			Optional:            true,
			ElementType:         types.StringType,
		},
		"use_public_load_balancer": schema.BoolAttribute{
			MarkdownDescription: "The boolean flag to request public load balancer. By default, private load balancer is used.",
			Optional:            true,
		},
		"whitelist_authorized_ip_ranges": schema.BoolAttribute{
			MarkdownDescription: "Whether to whitelist only 'authorizedIPRanges' given or all public IPs.",
			Optional:            true,
		},
		"workspace_name": schema.StringAttribute{
			MarkdownDescription: "The name of the workspace to create.",
			Required:            true,
		},
	},
}
