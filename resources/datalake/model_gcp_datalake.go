// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datalake

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

type gcpDatalakeResourceModel struct {
	ID types.String `tfsdk:"id"`

	PollingOptions *utils.PollingOptions `tfsdk:"polling_options"`

	CloudProviderConfiguration *gcpConfiguration `tfsdk:"cloud_provider_configuration"`

	CreationDate types.String `tfsdk:"creation_date"`

	Crn types.String `tfsdk:"crn"`

	CustomInstanceGroups []*gcpDatalakeInstanceGroup `tfsdk:"custom_instance_groups"`

	DatalakeName types.String `tfsdk:"datalake_name"`

	EnableRangerRaz types.Bool `tfsdk:"enable_ranger_raz"`

	EnvironmentCrn types.String `tfsdk:"environment_crn"`

	EnvironmentName types.String `tfsdk:"environment_name"`

	Image *gcpDatalakeImage `tfsdk:"image"`

	JavaVersion types.Int64 `tfsdk:"java_version"`

	MultiAz types.Bool `tfsdk:"multi_az"`

	Recipes []*instanceGroupRecipe `tfsdk:"recipes"`

	Runtime types.String `tfsdk:"runtime"`

	Scale types.String `tfsdk:"scale"`

	Status types.String `tfsdk:"status"`

	StatusReason types.String `tfsdk:"status_reason"`

	Tags types.Map `tfsdk:"tags"`
}

type gcpConfiguration struct {
	ServiceAccountEmail types.String `tfsdk:"service_account_email"`
	StorageLocation     types.String `tfsdk:"storage_location"`
}

type gcpDatalakeInstanceGroup struct {
	InstanceType types.String `tfsdk:"instance_type"`

	Name types.String `tfsdk:"name"`
}

type gcpDatalakeImage struct {
	CatalogName types.String `tfsdk:"catalog_name"`

	ID types.String `tfsdk:"id"`

	Os types.String `tfsdk:"os"`
}
