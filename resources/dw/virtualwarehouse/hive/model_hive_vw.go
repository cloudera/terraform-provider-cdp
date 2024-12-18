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
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

type awsOptions struct {
	AvailabilityZone types.String `tfsdk:"availability_zone"`
	EbsLLAPSpillGb   types.Int64  `tfsdk:"ebs_llap_spill_gb"`
	Tags             types.Map    `tfsdk:"tags"`
}

type resourceModel struct {
	ID                           types.String          `tfsdk:"id"`
	ClusterID                    types.String          `tfsdk:"cluster_id"`
	DatabaseCatalogID            types.String          `tfsdk:"database_catalog_id"`
	Name                         types.String          `tfsdk:"name"`
	ImageVersion                 types.String          `tfsdk:"image_version"`
	GroupSize                    types.Int64           `tfsdk:"group_size"`
	PlatformJwtAuth              types.Bool            `tfsdk:"platform_jwt_auth"`
	LdapGroups                   types.List            `tfsdk:"ldap_groups"`
	EnableSSO                    types.Bool            `tfsdk:"enable_sso"`
	Compactor                    types.Bool            `tfsdk:"compactor"`
	JdbcUrl                      types.String          `tfsdk:"jdbc_url"`
	KerberosJdbcUrl              types.String          `tfsdk:"kerberos_jdbc_url"`
	HueUrl                       types.String          `tfsdk:"hue_url"`
	JwtConnectionString          types.String          `tfsdk:"jwt_connection_string"`
	JwtTokenGenUrl               types.String          `tfsdk:"jwt_token_gen_url"`
	MinGroupCount                types.Int64           `tfsdk:"min_group_count"`
	MaxGroupCount                types.Int64           `tfsdk:"max_group_count"`
	DisableAutoSuspend           types.Bool            `tfsdk:"disable_auto_suspend"`
	AutoSuspendTimeoutSeconds    types.Int64           `tfsdk:"auto_suspend_timeout_seconds"`
	ScaleWaitTimeSeconds         types.Int64           `tfsdk:"scale_wait_time_seconds"`
	Headroom                     types.Int64           `tfsdk:"headroom"`
	MaxConcurrentIsolatedQueries types.Int64           `tfsdk:"max_concurrent_isolated_queries"`
	MaxNodesPerIsolatedQuery     types.Int64           `tfsdk:"max_nodes_per_isolated_query"`
	AwsOptions                   *awsOptions           `tfsdk:"aws_options"`
	LastUpdated                  types.String          `tfsdk:"last_updated"`
	Status                       types.String          `tfsdk:"status"`
	PollingOptions               *utils.PollingOptions `tfsdk:"polling_options"`
}

func (p *resourceModel) GetPollingOptions() *utils.PollingOptions {
	return p.PollingOptions
}

func (p *resourceModel) convertToCreateVwRequest(ctx context.Context) (*models.CreateVwRequest, diag.Diagnostics) {
	vwType := models.VwType("hive")
	tags, diags := p.getTags(ctx)
	return &models.CreateVwRequest{
		ClusterID:             p.ClusterID.ValueStringPointer(),
		DbcID:                 p.DatabaseCatalogID.ValueStringPointer(),
		EbsLLAPSpillGB:        p.getEbsLLAPSpillGB(),
		ImageVersion:          p.ImageVersion.ValueString(),
		Name:                  p.Name.ValueStringPointer(),
		NodeCount:             utils.Int64To32(p.GroupSize),
		PlatformJwtAuth:       p.PlatformJwtAuth.ValueBoolPointer(),
		QueryIsolationOptions: p.getQueryIsolationOptions(),
		Autoscaling:           p.getAutoscaling(),
		AvailabilityZone:      p.getAvailabilityZone(),
		Tags:                  tags,
		Config:                p.getServiceConfig(),
		VwType:                &vwType,
	}, diags
}

func (p *resourceModel) getServiceConfig() *models.ServiceConfigReq {
	return &models.ServiceConfigReq{
		ApplicationConfigs: nil,
		CommonConfigs:      nil,
		EnableSSO:          p.EnableSSO.ValueBool(),
		LdapGroups:         utils.FromListValueToStringList(p.LdapGroups),
	}
}

func (p *resourceModel) getTags(ctx context.Context) ([]*models.TagRequest, diag.Diagnostics) {
	if p.AwsOptions.Tags.IsNull() {
		return nil, diag.Diagnostics{}
	}
	tags := make([]*models.TagRequest, 0, len(p.AwsOptions.Tags.Elements()))
	elements := make(map[string]string, len(p.AwsOptions.Tags.Elements()))
	diags := p.AwsOptions.Tags.ElementsAs(ctx, &elements, false)
	for k, v := range elements {
		if v == "" {
			continue
		}
		tags = append(tags, &models.TagRequest{
			Key:   &k,
			Value: &v,
		})
	}
	return tags, diags
}

func (p *resourceModel) getQueryIsolationOptions() *models.QueryIsolationOptionsRequest {
	return &models.QueryIsolationOptionsRequest{
		MaxQueries:       utils.Int64To32(p.MaxConcurrentIsolatedQueries),
		MaxNodesPerQuery: utils.Int64To32(p.MaxNodesPerIsolatedQuery),
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
	return &models.AutoscalingOptionsCreateRequest{
		MinClusters:               utils.Int64To32Pointer(p.MinGroupCount),
		MaxClusters:               utils.Int64To32Pointer(p.MaxGroupCount),
		DisableAutoSuspend:        p.DisableAutoSuspend.ValueBool(),
		AutoSuspendTimeoutSeconds: utils.Int64To32(p.AutoSuspendTimeoutSeconds),
		HiveScaleWaitTimeSeconds:  utils.Int64To32(p.ScaleWaitTimeSeconds),
		HiveDesiredFreeCapacity:   utils.Int64To32(p.Headroom),
	}
}

func (p *resourceModel) setFromDescribeVwResponse(resp *models.DescribeVwResponse) {
	if resp.Vw == nil {
		return
	}
	p.ID = types.StringValue(resp.Vw.ID)
	p.DatabaseCatalogID = types.StringValue(resp.Vw.DbcID)
	p.Name = types.StringValue(resp.Vw.Name)
	p.Status = types.StringValue(resp.Vw.Status)
	p.ImageVersion = types.StringValue(resp.Vw.CdhVersion)
	p.Compactor = types.BoolValue(resp.Vw.Compactor)
	p.JdbcUrl = types.StringValue(resp.Vw.Endpoints.HiveJdbc)
	p.KerberosJdbcUrl = types.StringValue(resp.Vw.Endpoints.HiveKerberosJdbc)
	p.HueUrl = types.StringValue(resp.Vw.Endpoints.Hue)
	p.JwtConnectionString = types.StringValue(resp.Vw.Endpoints.JwtConnectionString)
	p.JwtTokenGenUrl = types.StringValue(resp.Vw.Endpoints.JwtTokenGenURL)
	p.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
}
