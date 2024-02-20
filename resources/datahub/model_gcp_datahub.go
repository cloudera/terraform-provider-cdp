// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datahub

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

type gcpDatahubResourceModel struct {
	ID                types.String          `tfsdk:"id"`
	Crn               types.String          `tfsdk:"crn"`
	Name              types.String          `tfsdk:"name"`
	Status            types.String          `tfsdk:"status"`
	Environment       types.String          `tfsdk:"environment_name"`
	InstanceGroup     []GcpInstanceGroup    `tfsdk:"instance_group"`
	PollingOptions    *utils.PollingOptions `tfsdk:"polling_options"`
	DestroyOptions    *DestroyOptions       `tfsdk:"destroy_options"`
	ClusterTemplate   types.String          `tfsdk:"cluster_template_name"`
	ClusterDefinition types.String          `tfsdk:"cluster_definition_name"`
}

type GcpInstanceGroup struct {
	NodeCount                   types.Int64                   `tfsdk:"node_count"`
	InstanceGroupName           types.String                  `tfsdk:"instance_group_name"`
	InstanceGroupType           types.String                  `tfsdk:"instance_group_type"`
	InstanceType                types.String                  `tfsdk:"instance_type"`
	Recipes                     []types.String                `tfsdk:"recipes"`
	AttachedVolumeConfiguration []AttachedVolumeConfiguration `tfsdk:"attached_volume_configuration"`
	RootVolumeSize              types.Int64                   `tfsdk:"root_volume_size"`
	RecoveryMode                types.String                  `tfsdk:"recovery_mode"`
}

func (d *gcpDatahubResourceModel) forceDeleteRequested() bool {
	return d.DestroyOptions != nil && !d.DestroyOptions.ForceDeleteCluster.IsNull() && d.DestroyOptions.ForceDeleteCluster.ValueBool()
}
