// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package hive

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

type autoscaling struct {
	MinClusters               types.Int64 `tfsdk:"min_clusters"`
	MaxClusters               types.Int64 `tfsdk:"max_clusters"`
	DisableAutoSuspend        types.Bool  `tfsdk:"disable_auto_suspend"`
	AutoSuspendTimeoutSeconds types.Int64 `tfsdk:"auto_suspend_timeout_seconds"`
	HiveScaleWaitTimeSeconds  types.Int64 `tfsdk:"hive_scale_wait_time_seconds"`
	HiveDesiredFreeCapacity   types.Int64 `tfsdk:"hive_desired_free_capacity"`
}

type awsOptions struct {
	AvailabilityZone types.String `tfsdk:"availability_zone"`
	EbsLLAPSpillGb   types.Int64  `tfsdk:"ebs_llap_spill_gb"`
	Tags             types.Map    `tfsdk:"tags"`
}

type queryIsolationOptions struct {
	MaxQueries       types.Int64 `tfsdk:"max_queries"`
	MaxNodesPerQuery types.Int64 `tfsdk:"max_nodes_per_query"`
}

type resourceModel struct {
	ID                    types.String           `tfsdk:"id"`
	ClusterID             types.String           `tfsdk:"cluster_id"`
	DatabaseCatalogID     types.String           `tfsdk:"database_catalog_id"`
	Name                  types.String           `tfsdk:"name"`
	ImageVersion          types.String           `tfsdk:"image_version"`
	NodeCount             types.Int64            `tfsdk:"node_count"`
	PlatformJwtAuth       types.Bool             `tfsdk:"platform_jwt_auth"`
	LdapGroups            types.List             `tfsdk:"ldap_groups"`
	EnableSSO             types.Bool             `tfsdk:"enable_sso"`
	Compactor             types.Bool             `tfsdk:"compactor"`
	JdbcUrl               types.String           `tfsdk:"jdbc_url"`
	KerberosJdbcUrl       types.String           `tfsdk:"kerberos_jdbc_url"`
	HueUrl                types.String           `tfsdk:"hue_url"`
	JwtConnectionString   types.String           `tfsdk:"jwt_connection_string"`
	JwtTokenGenUrl        types.String           `tfsdk:"jwt_token_gen_url"`
	Autoscaling           *autoscaling           `tfsdk:"autoscaling"`
	AwsOptions            *awsOptions            `tfsdk:"aws_options"`
	QueryIsolationOptions *queryIsolationOptions `tfsdk:"query_isolation_options"`
	LastUpdated           types.String           `tfsdk:"last_updated"`
	Status                types.String           `tfsdk:"status"`
	PollingOptions        *utils.PollingOptions  `tfsdk:"polling_options"`
}

func (p *resourceModel) GetPollingOptions() *utils.PollingOptions {
	return p.PollingOptions
}

func (p *resourceModel) convertToCreateVwRequest() *models.CreateVwRequest {
	vwType := models.VwType("hive")
	return &models.CreateVwRequest{
		ClusterID:      p.ClusterID.ValueStringPointer(),
		DbcID:          p.DatabaseCatalogID.ValueStringPointer(),
		EbsLLAPSpillGB: p.getEbsLLAPSpillGB(),
		//ImageVersion:          p.getImageVersion(),
		Name:                  p.Name.ValueStringPointer(),
		NodeCount:             utils.Int64To32(p.NodeCount),
		PlatformJwtAuth:       p.PlatformJwtAuth.ValueBoolPointer(),
		QueryIsolationOptions: p.getQueryIsolationOptions(),
		Autoscaling:           p.getAutoscaling(),
		AvailabilityZone:      p.getAvailabilityZone(),
		//Tags:                  p.getTags(),
		Config: p.getServiceConfig(),
		VwType: &vwType,
	}
}

func (p *resourceModel) getImageVersion() string {
	if p.ImageVersion.IsNull() || p.ImageVersion.String() == "unknown" {
		return ""
	}
	return p.ImageVersion.String()
}

func (p *resourceModel) getServiceConfig() *models.ServiceConfigReq {
	return &models.ServiceConfigReq{
		ApplicationConfigs: nil,
		CommonConfigs:      nil,
		EnableSSO:          p.EnableSSO.ValueBool(),
		LdapGroups:         utils.FromListValueToStringList(p.LdapGroups),
	}
}

func (p *resourceModel) getTags() []*models.TagRequest {
	if p.AwsOptions.Tags.IsNull() {
		return nil
	}
	tags := make([]*models.TagRequest, len(p.AwsOptions.Tags.Elements()))
	for k, v := range p.AwsOptions.Tags.Elements() {
		if v.IsNull() {
			continue
		}
		value := v.String()
		tags = append(tags, &models.TagRequest{
			Key:   &k,
			Value: &value,
		})
	}
	return tags
}

func (p *resourceModel) getQueryIsolationOptions() *models.QueryIsolationOptionsRequest {
	if p.QueryIsolationOptions == nil {
		return nil
	}
	return &models.QueryIsolationOptionsRequest{
		MaxQueries:       utils.Int64To32(p.QueryIsolationOptions.MaxQueries),
		MaxNodesPerQuery: utils.Int64To32(p.QueryIsolationOptions.MaxNodesPerQuery),
	}
}

func (p *resourceModel) getAvailabilityZone() string {
	if p.AwsOptions == nil {
		return ""
	}
	return p.AwsOptions.AvailabilityZone.ValueString()
}

func (p *resourceModel) getEbsLLAPSpillGB() int32 {
	if p.AwsOptions == nil {
		return 0
	}
	return utils.Int64To32(p.AwsOptions.EbsLLAPSpillGb)
}

func (p *resourceModel) getAutoscaling() *models.AutoscalingOptionsCreateRequest {
	if p.Autoscaling == nil {
		return nil
	}
	return &models.AutoscalingOptionsCreateRequest{
		MinClusters:               utils.Int64To32Pointer(p.Autoscaling.MinClusters),
		MaxClusters:               utils.Int64To32Pointer(p.Autoscaling.MaxClusters),
		DisableAutoSuspend:        p.Autoscaling.DisableAutoSuspend.ValueBool(),
		AutoSuspendTimeoutSeconds: utils.Int64To32(p.Autoscaling.AutoSuspendTimeoutSeconds),
		HiveScaleWaitTimeSeconds:  utils.Int64To32(p.Autoscaling.HiveScaleWaitTimeSeconds),
		HiveDesiredFreeCapacity:   utils.Int64To32(p.Autoscaling.HiveDesiredFreeCapacity),
	}
}
