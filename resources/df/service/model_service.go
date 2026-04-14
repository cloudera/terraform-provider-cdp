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
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type serviceModel struct {
	ID                             types.String `tfsdk:"id"`
	EnvironmentCrn                 types.String `tfsdk:"environment_crn"`
	Crn                            types.String `tfsdk:"crn"`
	Name                           types.String `tfsdk:"name"`
	CloudPlatform                  types.String `tfsdk:"cloud_platform"`
	Region                         types.String `tfsdk:"region"`
	Status                         types.String `tfsdk:"status"`
	StatusMessage                  types.String `tfsdk:"status_message"`
	WorkloadVersion                types.String `tfsdk:"workload_version"`
	MinK8sNodeCount                types.Int32  `tfsdk:"min_k8s_node_count"`
	MaxK8sNodeCount                types.Int32  `tfsdk:"max_k8s_node_count"`
	UsePublicLoadBalancer          types.Bool   `tfsdk:"use_public_load_balancer"`
	PrivateCluster                 types.Bool   `tfsdk:"private_cluster"`
	ClusterSubnets                 types.List   `tfsdk:"cluster_subnets"`
	LoadBalancerSubnets            types.List   `tfsdk:"load_balancer_subnets"`
	KubeAPIAuthorizedIPRanges      types.List   `tfsdk:"kube_api_authorized_ip_ranges"`
	LoadBalancerAuthorizedIPRanges types.List   `tfsdk:"load_balancer_authorized_ip_ranges"`
	Tags                           types.Map    `tfsdk:"tags"`
	InstanceType                   types.String `tfsdk:"instance_type"`
	SkipPreflightChecks            types.Bool   `tfsdk:"skip_preflight_checks"`
	UserDefinedRouting             types.Bool   `tfsdk:"user_defined_routing"`
	PodCidr                        types.String `tfsdk:"pod_cidr"`
	ServiceCidr                    types.String `tfsdk:"service_cidr"`
	TerminateDeploymentsOnDisable  types.Bool   `tfsdk:"terminate_deployments_on_disable"`
	PollingTimeout                 types.Int64  `tfsdk:"polling_timeout"`
}
