// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package service

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var serviceSchema = schema.Schema{
	MarkdownDescription: "Enables a CDP DataFlow service on a given CDP environment. The service manages the underlying Kubernetes cluster and allows deploying NiFi flow deployments.",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The ID of the service, same as the CRN.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"environment_crn": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The CDP environment CRN.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"crn": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The CRN of the DataFlow service.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The name of the DataFlow service.",
		},
		"cloud_platform": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The cloud platform of the service.",
		},
		"region": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The region of the service.",
		},
		"status": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The status of the DataFlow service.",
		},
		"status_message": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The status message of the DataFlow service.",
		},
		"workload_version": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The workload version of the DataFlow service.",
		},
		"min_k8s_node_count": schema.Int32Attribute{
			Required:            true,
			MarkdownDescription: "The minimum number of Kubernetes nodes needed for the service.",
		},
		"max_k8s_node_count": schema.Int32Attribute{
			Required:            true,
			MarkdownDescription: "The maximum number of Kubernetes nodes that the service may scale up to.",
		},
		"use_public_load_balancer": schema.BoolAttribute{
			Required:            true,
			MarkdownDescription: "Whether to use a public load balancer when deploying dependencies stack.",
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.RequiresReplace(),
			},
		},
		"private_cluster": schema.BoolAttribute{
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
			MarkdownDescription: "Whether to provision a private Kubernetes cluster. Defaults to `false`.",
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.RequiresReplace(),
			},
		},
		"cluster_subnets": schema.ListAttribute{
			Optional:            true,
			ElementType:         types.StringType,
			MarkdownDescription: "Subnets to use for the Kubernetes cluster.",
		},
		"load_balancer_subnets": schema.ListAttribute{
			Optional:            true,
			ElementType:         types.StringType,
			MarkdownDescription: "Subnets to use for the Load Balancer.",
		},
		"kube_api_authorized_ip_ranges": schema.ListAttribute{
			Optional:            true,
			ElementType:         types.StringType,
			MarkdownDescription: "The IP ranges authorized to connect to the Kubernetes API server.",
		},
		"load_balancer_authorized_ip_ranges": schema.ListAttribute{
			Optional:            true,
			ElementType:         types.StringType,
			MarkdownDescription: "IP ranges authorized to access DF local endpoints.",
		},
		"tags": schema.MapAttribute{
			Optional:            true,
			ElementType:         types.StringType,
			MarkdownDescription: "Tags to apply to service-related resources.",
		},
		"instance_type": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "Custom instance type to be used.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"skip_preflight_checks": schema.BoolAttribute{
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
			MarkdownDescription: "Whether to skip Liftie's pre-flight checks. Defaults to `false`.",
		},
		"user_defined_routing": schema.BoolAttribute{
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
			MarkdownDescription: "Whether User Defined Routing (UDR) mode is enabled for AKS clusters. Defaults to `false`.",
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.RequiresReplace(),
			},
		},
		"pod_cidr": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "CIDR range from which to assign IPs to pods in the Kubernetes cluster.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"service_cidr": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "CIDR range from which to assign IPs to internal services in the Kubernetes cluster.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"terminate_deployments_on_disable": schema.BoolAttribute{
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(true),
			MarkdownDescription: "Whether to terminate all deployments when disabling the service. Defaults to `true`.",
		},
		"polling_timeout": schema.Int64Attribute{
			Optional:            true,
			Computed:            true,
			Default:             int64default.StaticInt64(3600),
			MarkdownDescription: "Timeout in seconds for polling the service status. Defaults to `3600` (1 hour).",
		},
	},
}
