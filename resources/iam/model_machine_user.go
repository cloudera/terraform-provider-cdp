// Copyright 2024 Cloudera. All Rights Reserved.
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

type machineUserResourceModel struct {
	Id                      types.String `tfsdk:"id"`
	Name                    types.String `tfsdk:"name"`
	Status                  types.String `tfsdk:"status"`
	WorkloadUsername        types.String `tfsdk:"workload_username"`
	WorkloadPassword        types.String `tfsdk:"workload_password"`
	CreationDate            types.String `tfsdk:"creation_date"`
	WorkloadPasswordDetails types.Object `tfsdk:"workload_password_details"`
	AzureCloudIdentities    types.Set    `tfsdk:"azure_cloud_identities"`
}

type workloadPasswordDetails struct {
	IsPasswordSet   types.Bool   `tfsdk:"is_password_set"`
	ExpirationDate  types.String `tfsdk:"expiration_date"`
	MinLifetimeDate types.String `tfsdk:"min_lifetime_date"`
}

type azureCloudIdentity struct {
	EnvironmentCrn types.String `tfsdk:"environment_crn"`
	ObjectId       types.String `tfsdk:"object_id"`
}
