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
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type deploymentModel struct {
	ID                          types.String `tfsdk:"id"`
	ServiceCrn                  types.String `tfsdk:"service_crn"`
	FlowVersionCrn              types.String `tfsdk:"flow_version_crn"`
	DeploymentName              types.String `tfsdk:"deployment_name"`
	DeploymentCrn               types.String `tfsdk:"deployment_crn"`
	Name                        types.String `tfsdk:"name"`
	Status                      types.String `tfsdk:"status"`
	StatusMessage               types.String `tfsdk:"status_message"`
	FlowName                    types.String `tfsdk:"flow_name"`
	FlowCrn                     types.String `tfsdk:"flow_crn"`
	ClusterSize                 types.String `tfsdk:"cluster_size"`
	CfmNifiVersion              types.String `tfsdk:"cfm_nifi_version"`
	AutoStartFlow               types.Bool   `tfsdk:"auto_start_flow"`
	ProjectCrn                  types.String `tfsdk:"project_crn"`
	StaticNodeCount             types.Int64  `tfsdk:"static_node_count"`
	AutoScalingEnabled          types.Bool   `tfsdk:"auto_scaling_enabled"`
	AutoScaleMinNodes           types.Int64  `tfsdk:"auto_scale_min_nodes"`
	AutoScaleMaxNodes           types.Int64  `tfsdk:"auto_scale_max_nodes"`
	NifiURL                     types.String `tfsdk:"nifi_url"`
	CurrentNodeCount            types.Int32  `tfsdk:"current_node_count"`
	DeployedByName              types.String `tfsdk:"deployed_by_name"`
	PollingTimeout              types.Int64  `tfsdk:"polling_timeout"`
	Strategy                    types.String `tfsdk:"strategy"`
	WaitForFlowToStopInMinutes  types.Int64  `tfsdk:"wait_for_flow_to_stop_in_minutes"`
	ParameterGroups             types.String `tfsdk:"parameter_groups"`
	ParameterGroupsSha          types.String `tfsdk:"parameter_groups_sha"`
}
