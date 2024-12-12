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
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

type databaseResourceModel struct {
	Crn               types.String `tfsdk:"crn"`
	DatabaseName      types.String `tfsdk:"database_name"`
	Status            types.String `tfsdk:"status"`
	Environment       types.String `tfsdk:"environment_name"`
	ScaleType         types.String `tfsdk:"scale_type"`
	StorageType       types.String `tfsdk:"storage_type"`
	DisableExternalDB types.Bool   `tfsdk:"disable_external_db"`
	StorageLocation   types.String `tfsdk:"storage_location"`

	PollingOptions *utils.PollingOptions `tfsdk:"polling_options"`

	DisableMultiAz types.Bool   `tfsdk:"disable_multi_az"`
	NumEdgeNodes   types.Int64  `tfsdk:"num_edge_nodes"`
	JavaVersion    types.Int64  `tfsdk:"java_version"`
	SubnetID       types.String `tfsdk:"subnet_id"`

	AutoScalingParameters     *AutoScalingParametersStruct     `tfsdk:"auto_scaling_parameters"`
	AttachedStorageForWorkers *AttachedStorageForWorkersStruct `tfsdk:"attached_storage_for_workers"`

	DisableKerberos types.Bool `tfsdk:"disable_kerberos"`
	DisableJwtAuth  types.Bool `tfsdk:"disable_jwt_auth"`
	Image           *Image     `tfsdk:"image"`
	EnableGrafana   types.Bool `tfsdk:"enable_grafana"`

	CustomUserTags     []KeyValuePair     `tfsdk:"custom_user_tags"`
	EnableRegionCanary types.Bool         `tfsdk:"enable_region_canary"`
	Recipes            []Recipe           `tfsdk:"recipes"`
	VolumeEncryptions  []VolumeEncryption `tfsdk:"volume_encryptions"`
	Architecture       types.String       `tfsdk:"architecture"`
}

type AutoScalingParametersStruct struct {
	TargetedValueForMetric types.Int64 `tfsdk:"targeted_value_for_metric"`
	MaxWorkersForDatabase  types.Int64 `tfsdk:"max_workers_for_database"`
	MaxWorkersPerBatch     types.Int64 `tfsdk:"max_workers_per_batch"`
	MinWorkersForDatabase  types.Int64 `tfsdk:"min_workers_for_database"`
	EvaluationPeriod       types.Int64 `tfsdk:"evaluation_period"`
	MinimumBlockCacheGb    types.Int64 `tfsdk:"minimum_block_cache_gb"`

	MaxHdfsUsagePercentage     types.Int64 `tfsdk:"max_hdfs_usage_percentage"`
	MaxRegionsPerRegionServer  types.Int64 `tfsdk:"max_regions_per_region_server"`
	MaxCPUUtilization          types.Int64 `tfsdk:"max_cpu_utilization"`
	MaxComputeNodesForDatabase types.Int64 `tfsdk:"max_compute_nodes_for_database"`
	MinComputeNodesForDatabase types.Int64 `tfsdk:"min_compute_nodes_for_database"`
}

type AttachedStorageForWorkersStruct struct {
	VolumeCount types.Int64  `tfsdk:"volume_count"`
	VolumeSize  types.Int64  `tfsdk:"volume_size"`
	VolumeType  types.String `tfsdk:"volume_type"`
}

type Image struct {
	ID      types.String `tfsdk:"id"`
	Catalog types.String `tfsdk:"catalog"`
}

type KeyValuePair struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

type Recipe struct {
	Names         types.Set    `tfsdk:"names"`
	InstanceGroup types.String `tfsdk:"instance_group"`
}

type VolumeEncryption struct {
	EncryptionKey types.String `tfsdk:"encryption_key"`
	InstanceGroup types.String `tfsdk:"instance_group"`
}
