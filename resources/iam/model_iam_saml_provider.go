// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package iam

import "github.com/hashicorp/terraform-plugin-framework/types"

type samlProviderModel struct {
	GenerateWorkloadUsernameByEmail types.Bool   `tfsdk:"generate_workload_username_by_email"`
	SamlMetadataDocument            types.String `tfsdk:"saml_metadata_document"`
	SyncGroupsOnLogin               types.Bool   `tfsdk:"sync_groups_on_login"`
	SamlProviderName                types.String `tfsdk:"saml_provider_name"`
	SamlProviderId                  types.String `tfsdk:"saml_provider_id"`
	CdpSpMetadata                   types.String `tfsdk:"cdp_sp_metadata"`
	CreationDate                    types.String `tfsdk:"creation_date"`
	EnableScim                      types.Bool   `tfsdk:"enable_scim"`
	ScimURL                         types.String `tfsdk:"scim_url"`
	Crn                             types.String `tfsdk:"crn"`
}
