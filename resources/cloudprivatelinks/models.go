// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package cloudprivatelinks

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type cloudPrivateLinkResourceModel struct {
	TrackingID           types.String `tfsdk:"id"`
	CloudServiceProvider types.String `tfsdk:"cloud_service_provider"`
	ServiceGroup         types.String `tfsdk:"service_group"`
	EnablePrivateDNS     types.Bool   `tfsdk:"enable_private_dns"`
	AWSAccountDetails    types.Object `tfsdk:"aws_account_details"`
	AzureAccountDetails  types.Object `tfsdk:"azure_account_details"`
}

type awsAccountDetails struct {
	CloudAccountID types.String `tfsdk:"cloud_account_id"`
	CredentialCRN  types.String `tfsdk:"credential_crn"`
	Region         types.String `tfsdk:"region"`
	VpcID          types.String `tfsdk:"vpc_id"`
	SubnetIDs      types.List   `tfsdk:"subnet_ids"`
}

type azureAccountDetails struct {
	CredentialCRN  types.String `tfsdk:"credential_crn"`
	SubscriptionID types.String `tfsdk:"subscription_id"`
	ResourceGroup  types.String `tfsdk:"resource_group"`
	Location       types.String `tfsdk:"location"`
	VNetID         types.String `tfsdk:"vnet_id"`
	SubnetID       types.String `tfsdk:"subnet_id"`
}

func awsAccountDetailsAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"cloud_account_id": types.StringType,
		"credential_crn":   types.StringType,
		"region":           types.StringType,
		"vpc_id":           types.StringType,
		"subnet_ids":       types.ListType{ElemType: types.StringType},
	}
}

func azureAccountDetailsAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"credential_crn":  types.StringType,
		"subscription_id": types.StringType,
		"resource_group":  types.StringType,
		"location":        types.StringType,
		"vnet_id":         types.StringType,
		"subnet_id":       types.StringType,
	}
}
