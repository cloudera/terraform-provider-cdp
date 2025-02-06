// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var dwClusterSchema = schema.Schema{
	MarkdownDescription: "Creates an AWS Data Warehouse cluster.",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"crn": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The cloudera resource name of the environment that the cluster will read from.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The name of the cluster matches the environment name.",
		},
		"cluster_id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The id of the cluster.",
		},
		"last_updated": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Timestamp of the last Terraform update of the order.",
		},
		"status": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The status of the cluster.",
		},
		"version": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The version of the cluster.",
		},
		"node_role_cdw_managed_policy_arn": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "The managed policy ARN to be attached to the created node instance role.",
		},
		"database_backup_retention_days": schema.Int64Attribute{
			Optional:            true,
			MarkdownDescription: "The number of days to retain database backups.",
		},
		"custom_registry_options": schema.SingleNestedAttribute{
			Optional: true,
			Attributes: map[string]schema.Attribute{
				"registry_type": schema.StringAttribute{
					Required:            true,
					MarkdownDescription: "Registry type, supported values are ECR or ACR.",
				},
				"repository_url": schema.StringAttribute{
					Required:            true,
					MarkdownDescription: "The URL of the registry.",
				},
			},
		},
		"custom_subdomain": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "The custom subdomain to keep compatibility with old URL format.",
		},
		"network_settings": schema.SingleNestedAttribute{
			Required:   true,
			Attributes: networkSettings,
		},
		"instance_settings": schema.SingleNestedAttribute{
			Optional:   true,
			Computed:   true,
			Attributes: instanceSettings,
		},
		"default_database_catalog": schema.SingleNestedAttribute{
			Computed:   true,
			Attributes: defaultDatabaseCatalogProperties,
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

var networkSettings = map[string]schema.Attribute{
	"worker_subnet_ids": schema.ListAttribute{
		Required:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "The list of subnet IDs for worker nodes.",
	},
	"load_balancer_subnet_ids": schema.ListAttribute{
		Required:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "The list of subnet IDs for the load balancer.",
	},
	"use_overlay_network": schema.BoolAttribute{
		Required:            true,
		MarkdownDescription: "Whether to use overlay network.",
	},
	"whitelist_k8s_cluster_access_ip_cidrs": schema.ListAttribute{
		Optional:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "The list of IP CIDRs to allow access for kubernetes cluster API endpoint.",
	},
	"whitelist_workload_access_ip_cidrs": schema.ListAttribute{
		Optional:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "The list of IP CIDRs to allow access for workload endpoints.",
	},
	"use_private_load_balancer": schema.BoolAttribute{
		Required:            true,
		MarkdownDescription: "Whether to use private IP addresses for the load balancer. Determines workload endpoint access.",
	},
	"use_public_worker_node": schema.BoolAttribute{
		Required:            true,
		MarkdownDescription: "Whether to use public IP addresses for worker nodes.",
	},
}

var instanceSettings = map[string]schema.Attribute{
	"custom_ami_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The custom AMI ID to use for worker nodes.",
	},
	"enable_spot_instances": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Whether to use spot instances for worker nodes.",
	},
	"compute_instance_types": schema.ListAttribute{
		Optional:            true,
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "The compute instance types that the environment is restricted to use. This affects the creation of virtual warehouses where this restriction will apply. Select an instance type that meets your computing, memory, networking, or storage needs. As of now, only a single instance type can be listed.",
	},
}

var defaultDatabaseCatalogProperties = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The ID of the database catalog.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the database catalog.",
	},
	"last_updated": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Timestamp of the last Terraform update of the order.",
	},
	"status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The status of the database catalog.",
	},
}
