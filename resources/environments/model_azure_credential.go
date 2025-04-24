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

type azureCredentialResourceModel struct {
	ID             types.String `tfsdk:"id"`
	CredentialName types.String `tfsdk:"credential_name"`
	SubscriptionID types.String `tfsdk:"subscription_id"`
	TenantID       types.String `tfsdk:"tenant_id"`
	AppBased       *AppBased    `tfsdk:"app_based"`
	Crn            types.String `tfsdk:"crn"`
	Description    types.String `tfsdk:"description"`
}

type AppBased struct {
	ApplicationID types.String `tfsdk:"application_id"`
	SecretKey     types.String `tfsdk:"secret_key"`
}
