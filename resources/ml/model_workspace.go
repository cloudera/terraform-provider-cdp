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

import "github.com/hashicorp/terraform-plugin-framework/types"

type workspaceResourceModel struct {
	Id                          types.String `tfsdk:"id"`
	AuthorizedIPRanges          types.Set    `tfsdk:"authorized_ip_ranges"`
	CdswMigrationMode           types.String `tfsdk:"cdsw_migration_mode"`
	DisableTLS                  types.Bool   `tfsdk:"disable_tls"`
	EnableGovernance            types.Bool   `tfsdk:"enable_governance"`
	EnableModelMetrics          types.Bool   `tfsdk:"enable_model_metrics"`
	EnableMonitoring            types.Bool   `tfsdk:"enable_monitoring"`
	EnvironmentName             types.String `tfsdk:"environment_name"`
	ExistingDatabaseConfig      types.Object `tfsdk:"existing_database_config"`
	ExistingNFS                 types.String `tfsdk:"existing_nfs"`
	LoadBalancerIPWhitelists    types.Set    `tfsdk:"load_balancer_ip_whitelists"`
	MlVersion                   types.String `tfsdk:"ml_version"`
	NfsVersion                  types.String `tfsdk:"nfs_version"`
	OutboundTypes               types.Set    `tfsdk:"outbound_types"`
	PrivateCluster              types.Bool   `tfsdk:"private_cluster"`
	ProvisionK8sRequest         types.Object `tfsdk:"provision_k8s_request"`
	SkipValidation              types.Bool   `tfsdk:"skip_validation"`
	StaticSubdomain             types.String `tfsdk:"static_subdomain"`
	SubnetsForLoadBalancers     types.Set    `tfsdk:"subnets_for_load_balancers"`
	UsePublicLoadBalancer       types.Bool   `tfsdk:"use_public_load_balancer"`
	WhitelistAuthorizedIPRanges types.Bool   `tfsdk:"whitelist_authorized_ip_ranges"`
	WorkspaceName               types.String `tfsdk:"workspace_name"`
}

type ExistingDatabaseConfig struct {
	ExistingDatabaseHost     types.String `tfsdk:"existing_database_host"`
	ExistingDatabaseName     types.String `tfsdk:"existing_database_name"`
	ExistingDatabasePassword types.String `tfsdk:"existing_database_password"`
	ExistingDatabasePort     types.String `tfsdk:"existing_database_port"`
	ExistingDatabaseUser     types.String `tfsdk:"existing_database_user"`
}

type ProvisionK8sRequest struct {
	EnvironmentName types.String `tfsdk:"environment_name"`
	InstanceGroups  types.Set    `tfsdk:"instance_groups"`
	Network         types.Object `tfsdk:"network"`
	Tags            types.Set    `tfsdk:"tags"`
}

type InstanceGroup struct {
	Autoscaling   types.Object `tfsdk:"autoscaling"`
	IngressRules  types.Set    `tfsdk:"ingress_rules"`
	InstanceCount types.Int64  `tfsdk:"instance_count"`
	InstanceTier  types.String `tfsdk:"instance_tier"`
	InstanceType  types.String `tfsdk:"instance_type"`
	Name          types.String `tfsdk:"name"`
	RootVolume    types.Object `tfsdk:"root_volume"`
}

type Autoscaling struct {
	Enabled      types.Bool  `tfsdk:"enabled"`
	MaxInstances types.Int64 `tfsdk:"max_instances"`
	MinInstances types.Int64 `tfsdk:"min_instances"`
}

type RootVolume struct {
	Size types.Int64 `tfsdk:"size"`
}

type OverlayNetwork struct {
	Plugin   types.String `tfsdk:"plugin"`
	Topology types.Object `tfsdk:"topology"`
}

type Topology struct {
	Subnets types.Set `tfsdk:"subnets"`
}

type ProvisionTag struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}
