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
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

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
	CustomAmiID          types.String `tfsdk:"custom_ami_id"`
	EnableSpotInstances  types.Bool   `tfsdk:"enable_spot_instances"`
	ComputeInstanceTypes types.List   `tfsdk:"compute_instance_types"`
}

type resourceModel struct {
	ID                          types.String           `tfsdk:"id"`
	Crn                         types.String           `tfsdk:"crn"`
	Name                        types.String           `tfsdk:"name"`
	ClusterID                   types.String           `tfsdk:"cluster_id"`
	LastUpdated                 types.String           `tfsdk:"last_updated"`
	Status                      types.String           `tfsdk:"status"`
	Version                     types.String           `tfsdk:"version"`
	NodeRoleCDWManagedPolicyArn types.String           `tfsdk:"node_role_cdw_managed_policy_arn"`
	DatabaseBackupRetentionDays types.Int64            `tfsdk:"database_backup_retention_days"`
	CustomRegistryOptions       *customRegistryOptions `tfsdk:"custom_registry_options"`
	CustomSubdomain             types.String           `tfsdk:"custom_subdomain"`
	NetworkSettings             *networkResourceModel  `tfsdk:"network_settings"`
	InstanceSettings            types.Object           `tfsdk:"instance_settings"`
	DefaultDatabaseCatalog      types.Object           `tfsdk:"default_database_catalog"`
	PollingOptions              *utils.PollingOptions  `tfsdk:"polling_options"`
}

func (p *resourceModel) convertToCreateAwsClusterRequest(ctx context.Context) (*models.CreateAwsClusterRequest, diag.Diagnostics) {
	enableSpotInstances, diags := p.getEnableSpotInstances(ctx)
	if diags.HasError() {
		return nil, diags
	}
	customAmiID, diags := p.getCustomAmiID(ctx)
	if diags.HasError() {
		return nil, diags
	}
	computeInstanceTypes, diags := p.getComputeInstanceTypes(ctx)
	if diags.HasError() {
		return nil, diags
	}

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
		EnableSpotInstances:              enableSpotInstances,
		CustomAmiID:                      customAmiID,
		ComputeInstanceTypes:             computeInstanceTypes,
	}, diags
}

func (p *resourceModel) getEnableSpotInstances(ctx context.Context) (*bool, diag.Diagnostics) {
	var irm instanceResourceModel
	diags := p.InstanceSettings.As(ctx, &irm, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
	return irm.EnableSpotInstances.ValueBoolPointer(), diags
}

func (p *resourceModel) getCustomAmiID(ctx context.Context) (string, diag.Diagnostics) {
	var irm instanceResourceModel
	diags := p.InstanceSettings.As(ctx, &irm, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
	return irm.CustomAmiID.ValueString(), diags
}

func (p *resourceModel) getComputeInstanceTypes(ctx context.Context) ([]string, diag.Diagnostics) {
	var irm instanceResourceModel
	diags := p.InstanceSettings.As(ctx, &irm, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
	return utils.FromListValueToStringList(irm.ComputeInstanceTypes), diags
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

func (p *resourceModel) GetPollingOptions() *utils.PollingOptions {
	return p.PollingOptions
}

func (p *resourceModel) setResourceModel(ctx context.Context, resp *models.DescribeClusterResponse) diag.Diagnostics {
	p.ID = types.StringValue(resp.Cluster.EnvironmentCrn)
	p.Crn = types.StringValue(resp.Cluster.EnvironmentCrn)
	p.Name = types.StringValue(resp.Cluster.Name)
	p.Status = types.StringValue(resp.Cluster.Status)
	p.Version = types.StringValue(resp.Cluster.Version)
	p.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
	var irm instanceResourceModel
	diags := p.InstanceSettings.As(ctx, &irm, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
	if diags.HasError() {
		return diags
	}
	attributeTypes := map[string]attr.Type{
		"custom_ami_id":          types.StringType,
		"enable_spot_instances":  types.BoolType,
		"compute_instance_types": types.ListType{ElemType: types.StringType},
	}
	attributes := map[string]attr.Value{
		"custom_ami_id":          irm.CustomAmiID,
		"enable_spot_instances":  basetypes.NewBoolValue(resp.Cluster.EnableSpotInstances),
		"compute_instance_types": utils.FromStringListToListValue(resp.Cluster.ComputeInstanceTypes),
	}
	p.InstanceSettings, diags = basetypes.NewObjectValue(attributeTypes, attributes)
	return diags
}

func (p *resourceModel) setDefaultDatabaseCatalog(catalog *models.DbcSummary) diag.Diagnostics {
	attributeTypes := map[string]attr.Type{
		"id":           types.StringType,
		"name":         types.StringType,
		"last_updated": types.StringType,
		"status":       types.StringType,
	}
	attributes := map[string]attr.Value{
		"id":           basetypes.NewStringValue(catalog.ID),
		"name":         basetypes.NewStringValue(catalog.Name),
		"last_updated": basetypes.NewStringValue(time.Now().Format(time.RFC850)),
		"status":       basetypes.NewStringValue(catalog.Status),
	}
	dbc, diags := basetypes.NewObjectValue(attributeTypes, attributes)
	p.DefaultDatabaseCatalog = dbc
	return diags
}
