// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package deployment

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var deploymentSchema = schema.Schema{
	MarkdownDescription: "Initiates a CDP DataFlow deployment on an enabled DataFlow service. A deployment runs a NiFi flow definition on the DataFlow service's Kubernetes cluster.",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The ID of the deployment, same as the deployment CRN.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"service_crn": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The CRN of the DataFlow service where the deployment will be created.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplaceIfConfigured(),
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"flow_version_crn": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The CRN of the flow definition version to deploy. Changing this triggers an in-place flow version change.",
		},
		"deployment_name": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The name of the deployment.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"deployment_crn": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The CRN of the deployment.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The name of the deployment (from API).",
		},
		"status": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The status of the deployment.",
		},
		"status_message": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The status message of the deployment.",
		},
		"flow_name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The name of the flow.",
		},
		"flow_crn": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The CRN of the flow definition.",
		},
		"cluster_size": schema.StringAttribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "The cluster size. Valid values: `EXTRA_SMALL`, `SMALL`, `MEDIUM`, `LARGE`. Defaults to `EXTRA_SMALL`.",
		},
		"cfm_nifi_version": schema.StringAttribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "The CFM NiFi version to use. Defaults to the latest version.",
		},
		"auto_start_flow": schema.BoolAttribute{
			Optional:            true,
			MarkdownDescription: "Whether to automatically start the flow after deployment. Defaults to true.",
		},
		"project_crn": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "The CRN of the project to assign this deployment to.",
		},
		"static_node_count": schema.Int64Attribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "The static number of nodes. Defaults to 1. Ignored when auto-scaling is enabled.",
		},
		"auto_scaling_enabled": schema.BoolAttribute{
			Optional:            true,
			MarkdownDescription: "Whether auto-scaling is enabled. Defaults to false.",
		},
		"auto_scale_min_nodes": schema.Int64Attribute{
			Optional:            true,
			MarkdownDescription: "Minimum number of nodes for auto-scaling.",
		},
		"auto_scale_max_nodes": schema.Int64Attribute{
			Optional:            true,
			MarkdownDescription: "Maximum number of nodes for auto-scaling.",
		},
		"nifi_url": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The URL to open the deployed flow in NiFi.",
		},
		"current_node_count": schema.Int32Attribute{
			Computed:            true,
			MarkdownDescription: "The current node count.",
		},
		"deployed_by_name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The name of the person who deployed the flow.",
		},
		"polling_timeout": schema.Int64Attribute{
			Optional:            true,
			Computed:            true,
			Default:             int64default.StaticInt64(3600),
			MarkdownDescription: "Timeout in seconds for polling the deployment status. Defaults to `3600` (1 hour).",
		},
		"strategy": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "The strategy to use during change flow version. Valid values: `STOP_AND_PROCESS_DATA`, `STOP_AND_EMPTY_QUEUES`, `ONLY_RESTART_AFFECTED_COMPONENTS`. Defaults to `STOP_AND_PROCESS_DATA`.",
		},
		"wait_for_flow_to_stop_in_minutes": schema.Int64Attribute{
			Optional:            true,
			MarkdownDescription: "The max time in minutes to wait for the flow to bleed out during change flow version. Only relevant with `STOP_AND_PROCESS_DATA` strategy. Defaults to 15.",
		},
		"parameter_groups": schema.StringAttribute{
			Optional:            true,
			Sensitive:           true,
			MarkdownDescription: "The parameter groups JSON for the deployment. Pass as a JSON string (e.g. `file(\"parameters.json\")`). The content is hidden from plan output; use `parameter_groups_sha` to detect changes.",
		},
		"parameter_groups_sha": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "SHA256 hash of the parameter_groups content. Changes when the parameter groups file is modified.",
		},
	},
}
