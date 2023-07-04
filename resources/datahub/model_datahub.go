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
)

type datahubResourceModel struct {
	ID                types.String    `tfsdk:"id"`
	Crn               types.String    `tfsdk:"crn"`
	Name              types.String    `tfsdk:"name"`
	Status            types.String    `tfsdk:"status"`
	DestroyOptions    *DestroyOptions `tfsdk:"destroy_options"`
	Environment       types.String    `tfsdk:"environment"`
	ClusterTemplate   types.String    `tfsdk:"cluster_template"`
	ClusterDefinition types.String    `tfsdk:"cluster_definition"`
	InstanceGroup     []InstanceGroup `tfsdk:"instance_group"`
}

type InstanceGroup struct {
	NodeCount                   types.Int64                   `tfsdk:"node_count"`
	InstanceGroupName           types.String                  `tfsdk:"instance_group_name"`
	InstanceGroupType           types.String                  `tfsdk:"instance_group_type"`
	InstanceType                types.String                  `tfsdk:"instance_type"`
	RootVolumeSize              types.Int64                   `tfsdk:"root_volume_size"`
	AttachedVolumeConfiguration []AttachedVolumeConfiguration `tfsdk:"attached_volume_configuration"`
	RecoveryMode                types.String                  `tfsdk:"recovery_mode"`
	VolumeEncryption            VolumeEncryption              `tfsdk:"volume_encryption"`
	Recipes                     []types.String                `tfsdk:"recipes"`
}

type AttachedVolumeConfiguration struct {
	VolumeSize  types.Int64  `tfsdk:"volume_size"`
	VolumeCount types.Int64  `tfsdk:"volume_count"`
	VolumeType  types.String `tfsdk:"volume_type"`
}

type VolumeEncryption struct {
	Encryption types.Bool `tfsdk:"encryption"`
}

type DestroyOptions struct {
	ForceDeleteCluster bool `tfsdk:"force_delete_cluster"`
}

func (d *datahubResourceModel) forceDeleteRequested() bool {
	return d.DestroyOptions != nil && d.DestroyOptions.ForceDeleteCluster
}
