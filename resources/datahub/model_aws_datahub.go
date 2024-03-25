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

type awsDatahubResourceModel struct {
	ID                       types.String          `tfsdk:"id"`
	Crn                      types.String          `tfsdk:"crn"`
	Name                     types.String          `tfsdk:"name"`
	Status                   types.String          `tfsdk:"status"`
	InstanceGroup            []InstanceGroup       `tfsdk:"instance_group"`
	PollingOptions           *utils.PollingOptions `tfsdk:"polling_options"`
	DestroyOptions           *DestroyOptions       `tfsdk:"destroy_options"`
	ClusterDefinition        types.String          `tfsdk:"cluster_definition"`
	ClusterTemplate          types.String          `tfsdk:"cluster_template"`
	Environment              types.String          `tfsdk:"environment"`
	CustomConfigurationsName types.String          `tfsdk:"custom_configurations_name"`
	Image                    types.Object          `tfsdk:"image"`
	RequestTemplate          types.String          `tfsdk:"request_template"`
	DatahubDatabase          types.String          `tfsdk:"datahub_database"`
	ClusterExtension         types.Object          `tfsdk:"cluster_extension"`
	SubnetId                 types.String          `tfsdk:"subnet_id"`
	SubnetIds                types.Set             `tfsdk:"subnet_ids"`
	MultiAz                  types.Bool            `tfsdk:"multi_az"`
	JavaVersion              types.Int64           `tfsdk:"java_version"`
	Tags                     types.Map             `tfsdk:"tags"`
	EnableLoadBalancer       types.Bool            `tfsdk:"enable_load_balancer"`
}

type datahubImage struct {
	CatalogName types.String `tfsdk:"catalog"`

	ID types.String `tfsdk:"id"`

	Os types.String `tfsdk:"os"`
}

type clusterExtension struct {
	CustomProperties types.String `tfsdk:"custom_properties"`
}

func (d *awsDatahubResourceModel) forceDeleteRequested() bool {
	return d.DestroyOptions != nil && !d.DestroyOptions.ForceDeleteCluster.IsNull() && d.DestroyOptions.ForceDeleteCluster.ValueBool()
}
