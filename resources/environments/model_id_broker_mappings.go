// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package environments

import "github.com/hashicorp/terraform-plugin-framework/types"

type idBrokerMappingsResourceModel struct {
	ID types.String `tfsdk:"id"`

	DataAccessRole types.String `tfsdk:"data_access_role"`

	EnvironmentName types.String `tfsdk:"environment_name"`

	EnvironmentCrn types.String `tfsdk:"environment_crn"`

	Mappings types.Set `tfsdk:"mappings"`

	RangerAuditRole types.String `tfsdk:"ranger_audit_role"`

	RangerCloudAccessAuthorizerRole types.String `tfsdk:"ranger_cloud_access_authorizer_role"`

	SetEmptyMappings types.Bool `tfsdk:"set_empty_mappings"`

	MappingsVersion types.Int64 `tfsdk:"mappings_version"`
}
