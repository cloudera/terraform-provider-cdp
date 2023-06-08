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

type awsDatalakeResourceModel struct {
	ID types.String `tfsdk:"id"`

	InstanceProfile types.String `tfsdk:"instance_profile"`

	StorageBucketLocation types.String `tfsdk:"storage_bucket_location"`

	CertificateExpirationState types.String `tfsdk:"certificate_expiration_state"`

	CloudStorageBaseLocation types.String `tfsdk:"cloud_storage_base_location"`

	ClouderaManager types.Object `tfsdk:"cloudera_manager"`

	CreationDate types.String `tfsdk:"creation_date"`

	CredentialCrn types.String `tfsdk:"credential_crn"`

	Crn types.String `tfsdk:"crn"`

	CustomInstanceGroups []*awsDatalakeInstanceGroup `tfsdk:"custom_instance_groups"`

	DatalakeName types.String `tfsdk:"datalake_name"`

	EnableRangerRaz types.Bool `tfsdk:"enable_ranger_raz"`

	Endpoints types.Set `tfsdk:"endpoints"`

	EnvironmentCrn types.String `tfsdk:"environment_crn"`

	EnvironmentName types.String `tfsdk:"environment_name"`

	Image *awsDatalakeImage `tfsdk:"image"`

	InstanceGroups types.Set `tfsdk:"instance_groups"`

	ProductVersions types.Set `tfsdk:"product_versions"`

	Region types.String `tfsdk:"region"`

	JavaVersion types.Int64 `tfsdk:"java_version"`

	MultiAz types.Bool `tfsdk:"multi_az"`

	Recipes []*instanceGroupRecipe `tfsdk:"recipes"`

	Runtime types.String `tfsdk:"runtime"`

	Scale types.String `tfsdk:"scale"`

	Status types.String `tfsdk:"status"`

	StatusReason types.String `tfsdk:"status_reason"`

	Tags types.Map `tfsdk:"tags"`
}

type awsDatalakeInstanceGroup struct {
	InstanceType types.String `tfsdk:"instance_type"`

	Name types.String `tfsdk:"name"`
}

type awsDatalakeImage struct {
	CatalogName types.String `tfsdk:"catalog_name"`

	ID types.String `tfsdk:"id"`
}

type instanceGroupRecipe struct {
	InstanceGroupName types.String `tfsdk:"instance_group_name"`

	RecipeNames types.Set `tfsdk:"recipe_names"`
}

type instanceGroup struct {
	Instances types.Set `tfsdk:"instances"`

	Name types.String `tfsdk:"name"`
}

type instance struct {
	DiscoveryFQDN types.String `tfsdk:"discovery_fqdn"`

	ID types.String `tfsdk:"id"`

	InstanceGroup types.String `tfsdk:"instance_group"`

	InstanceStatus types.String `tfsdk:"instance_status"`

	InstanceTypeVal types.String `tfsdk:"instance_type_val"`

	PrivateIP types.String `tfsdk:"private_ip"`

	PublicIP types.String `tfsdk:"public_ip"`

	SSHPort types.Int64 `tfsdk:"ssh_port"`

	State types.String `tfsdk:"state"`

	StatusReason types.String `tfsdk:"status_reason"`
}

type clouderaManagerDetails struct {
	ClouderaManagerRepositoryURL types.String `tfsdk:"cloudera_manager_repository_url"`

	ClouderaManagerServerURL types.String `tfsdk:"cloudera_manager_server_url"`

	Version types.String `tfsdk:"version"`
}

type endpoint struct {
	DisplayName types.String `tfsdk:"display_name"`

	KnoxService types.String `tfsdk:"knox_service"`

	Mode types.String `tfsdk:"mode"`

	Open types.Bool `tfsdk:"open"`

	ServiceName types.String `tfsdk:"service_name"`

	ServiceURL types.String `tfsdk:"service_url"`
}

type productVersion struct {
	Name types.String `tfsdk:"name"`

	Version types.String `tfsdk:"version"`
}
