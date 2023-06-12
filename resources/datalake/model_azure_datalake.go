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

import "github.com/hashicorp/terraform-plugin-framework/types"

type azureDatalakeResourceModel struct {
	ID types.String `tfsdk:"id"`

	ManagedIdentity types.String `tfsdk:"managed_identity"`

	StorageLocation types.String `tfsdk:"storage_location"`

	CertificateExpirationState types.String `tfsdk:"certificate_expiration_state"`

	CloudStorageBaseLocation types.String `tfsdk:"cloud_storage_base_location"`

	CloudbreakVersion types.String `tfsdk:"cloudbreak_version"`

	ClouderaManager types.Object `tfsdk:"cloudera_manager"`

	CreationDate types.String `tfsdk:"creation_date"`

	CredentialCrn types.String `tfsdk:"credential_crn"`

	Crn types.String `tfsdk:"crn"`

	DatalakeName types.String `tfsdk:"datalake_name"`

	EnableRangerRaz types.Bool `tfsdk:"enable_ranger_raz"`

	Endpoints types.Set `tfsdk:"endpoints"`

	EnvironmentCrn types.String `tfsdk:"environment_crn"`

	EnvironmentName types.String `tfsdk:"environment_name"`

	Image *azureDatalakeImage `tfsdk:"image"`

	InstanceGroups types.Set `tfsdk:"instance_groups"`

	ProductVersions types.Set `tfsdk:"product_versions"`

	Region types.String `tfsdk:"region"`

	JavaVersion types.Int64 `tfsdk:"java_version"`

	Recipes []*instanceGroupRecipe `tfsdk:"recipes"`

	Runtime types.String `tfsdk:"runtime"`

	Scale types.String `tfsdk:"scale"`

	Status types.String `tfsdk:"status"`

	StatusReason types.String `tfsdk:"status_reason"`

	Tags types.Map `tfsdk:"tags"`
}

type azureDatalakeImage struct {
	CatalogName types.String `tfsdk:"catalog_name"`

	ID types.String `tfsdk:"id"`
}
