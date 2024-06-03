// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package de

import "github.com/hashicorp/terraform-plugin-framework/types"

type serviceResourceModel struct {
	Id                      types.String `tfsdk:"id"`
	ChartValueOverrides     types.Set    `tfsdk:"chart_value_overrides"`
	CPURequests             types.String `tfsdk:"cpu_requests"`
	CustomAzureFilesConfigs types.Object `tfsdk:"custom_azure_files_configs"`
	DeployPreviousVersion   types.Bool   `tfsdk:"deploy_previous_version"`
	EnablePrivateNetwork    types.Bool   `tfsdk:"enable_private_network"`
	EnablePublicEndpoint    types.Bool   `tfsdk:"enable_public_endpoint"`
	EnableWorkloadAnalytics types.Bool   `tfsdk:"enable_workload_analytics"`
	Env                     types.String `tfsdk:"env"`
	GpuRequests             types.String `tfsdk:"gpu_requests"`
	InitialInstances        types.Int64  `tfsdk:"initial_instances"`
	InitialSpotInstances    types.Int64  `tfsdk:"initial_spot_instances"`
	InstanceType            types.String `tfsdk:"instance_type"`
	LoadbalancerAllowlist   types.Set    `tfsdk:"loadbalancer_allowlist"`
	MaximumInstances        types.Int64  `tfsdk:"maximum_instances"`
	MaximumSpotInstances    types.Int64  `tfsdk:"maximum_spot_instances"`
	MemoryRequests          types.String `tfsdk:"memory_requests"`
	MinimumInstances        types.Int64  `tfsdk:"minimum_instances"`
	MinimumSpotInstances    types.Int64  `tfsdk:"minimum_spot_instances"`
	Name                    types.String `tfsdk:"name"`
	NetworkOutboundType     types.String `tfsdk:"network_outbound_type"`
	NfsStorageClass         types.String `tfsdk:"nfs_storage_class"`
	ResourcePool            types.String `tfsdk:"resource_pool"`
	RootVolumeSize          types.Int64  `tfsdk:"root_volume_size"`
	SkipValidation          types.Bool   `tfsdk:"skip_validation"`
	Subnets                 types.Set    `tfsdk:"subnets"`
	Tags                    types.Map    `tfsdk:"tags"`
	UseSsd                  types.Bool   `tfsdk:"use_ssd"`
	WhitelistIps            types.Set    `tfsdk:"whitelist_ips"`
}

type chartValueOverridesRequest struct {
	ChartName types.String `tfsdk:"chart_name"`
	Overrides types.String `tfsdk:"overrides"`
}

type customAzureFilesConfigs struct {
	AzureFilesFQDN     types.String `tfsdk:"azure_files_fqdn"`
	ResourceGroup      types.String `tfsdk:"resource_group"`
	StorageAccountName types.String `tfsdk:"storage_account_name"`
}
