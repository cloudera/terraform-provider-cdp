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
	"time"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

type networkResourceModel struct {
	WorkerSubnetIDs                  types.List `tfsdk:"worker_subnet_ids"`
	LoadBalancerSubnetIDs            types.List `tfsdk:"load_balancer_subnet_ids"`
	UseOverlayNetwork                types.Bool `tfsdk:"use_overlay_network"`
	WhitelistK8sClusterAccessIPCIDRs types.List `tfsdk:"whitelist_k8s_cluster_access_ip_cidrs"`
	WhitelistWorkloadAccessIPCIDRs   types.List `tfsdk:"whitelist_workload_access_ip_cidrs"`
	UsePrivateLoadBalancer           types.Bool `tfsdk:"use_private_load_balancer"`
	UsePublicWorkerNode              types.Bool `tfsdk:"use_public_worker_node"`
}

type customRegistryOptions struct {
	RegistryType  types.String `tfsdk:"registry_type"`
	RepositoryURL types.String `tfsdk:"repository_url"`
}

type instanceResourceModel struct {
	CustomAmiID             types.String `tfsdk:"custom_ami_id"`
	EnableSpotInstances     types.Bool   `tfsdk:"enable_spot_instances"`
	ComputeInstanceTypes    types.List   `tfsdk:"compute_instance_types"`
	AdditionalInstanceTypes types.List   `tfsdk:"additional_instance_types"`
}

type resourceModel struct {
	ID                          types.String           `tfsdk:"id"`
	Crn                         types.String           `tfsdk:"crn"`
	Name                        types.String           `tfsdk:"name"`
	ClusterID                   types.String           `tfsdk:"cluster_id"`
	LastUpdated                 types.String           `tfsdk:"last_updated"`
	Status                      types.String           `tfsdk:"status"`
	NodeRoleCDWManagedPolicyArn types.String           `tfsdk:"node_role_cdw_managed_policy_arn"`
	DatabaseBackupRetentionDays types.Int64            `tfsdk:"database_backup_retention_days"`
	CustomRegistryOptions       *customRegistryOptions `tfsdk:"custom_registry_options"`
	CustomSubdomain             types.String           `tfsdk:"custom_subdomain"`
	NetworkSettings             *networkResourceModel  `tfsdk:"network_settings"`
	InstanceSettings            *instanceResourceModel `tfsdk:"instance_settings"`
	PollingOptions              *utils.PollingOptions  `tfsdk:"polling_options"`
}

func (p *resourceModel) convertToCreateAwsClusterRequest() *models.CreateAwsClusterRequest {
	return &models.CreateAwsClusterRequest{
		EnvironmentCrn:                   p.Crn.ValueStringPointer(),
		UseOverlayNetwork:                p.NetworkSettings.UseOverlayNetwork.ValueBool(),
		WhitelistK8sClusterAccessIPCIDRs: utils.FromListValueToStringList(p.NetworkSettings.WhitelistK8sClusterAccessIPCIDRs),
		WhitelistWorkloadAccessIPCIDRs:   utils.FromListValueToStringList(p.NetworkSettings.WhitelistWorkloadAccessIPCIDRs),
		UsePrivateLoadBalancer:           p.NetworkSettings.UsePrivateLoadBalancer.ValueBool(),
		UsePublicWorkerNode:              p.NetworkSettings.UsePublicWorkerNode.ValueBool(),
		WorkerSubnetIds:                  utils.FromListValueToStringList(p.NetworkSettings.WorkerSubnetIDs),
		LbSubnetIds:                      utils.FromListValueToStringList(p.NetworkSettings.LoadBalancerSubnetIDs),
		NodeRoleCDWManagedPolicyArn:      p.NodeRoleCDWManagedPolicyArn.ValueString(),
		DatabaseBackupRetentionPeriod:    utils.Int64To32Pointer(p.DatabaseBackupRetentionDays),
		CustomSubdomain:                  p.CustomSubdomain.ValueString(),
		CustomRegistryOptions:            p.getCustomRegistryOptions(),
		EnableSpotInstances:              p.getEnableSpotInstances(),
		CustomAmiID:                      p.getCustomAmiID(),
		ComputeInstanceTypes:             p.getComputeInstanceTypes(),
		AdditionalInstanceTypes:          p.getAdditionalInstanceTypes(),
	}
}

func (p *resourceModel) getEnableSpotInstances() *bool {
	if i := p.InstanceSettings; i != nil {
		return i.EnableSpotInstances.ValueBoolPointer()
	}
	return nil
}

func (p *resourceModel) getCustomAmiID() string {
	if i := p.InstanceSettings; i != nil {
		return p.InstanceSettings.CustomAmiID.ValueString()
	}
	return ""
}

func (p *resourceModel) getComputeInstanceTypes() []string {
	if i := p.InstanceSettings; i != nil {
		return utils.FromListValueToStringList(p.InstanceSettings.ComputeInstanceTypes)
	}
	return nil
}

func (p *resourceModel) getAdditionalInstanceTypes() []string {
	if i := p.InstanceSettings; i != nil {
		return utils.FromListValueToStringList(p.InstanceSettings.AdditionalInstanceTypes)
	}
	return nil
}

func (p *resourceModel) getCustomRegistryOptions() *models.CustomRegistryOptions {
	if cro := p.CustomRegistryOptions; cro != nil {
		return &models.CustomRegistryOptions{
			RegistryType:  p.CustomRegistryOptions.RegistryType.ValueString(),
			RepositoryURL: p.CustomRegistryOptions.RepositoryURL.ValueString(),
		}
	}
	return nil
}

func (p *resourceModel) getPollingTimeout() time.Duration {
	if p.PollingOptions != nil {
		return time.Duration(p.PollingOptions.PollingTimeout.ValueInt64()) * time.Minute
	}
	return 40 * time.Minute
}
