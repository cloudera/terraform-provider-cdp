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
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	cloudprivatelinkmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/cloudprivatelinks/models"
)

func fromModelToAwsRequest(model cloudPrivateLinkResourceModel, ctx context.Context) *cloudprivatelinkmodels.CreatePrivateLinkEndpointRequest {
	tflog.Debug(ctx, "Converting model to AWS CreatePrivateLinkEndpoint request")

	csp := cloudprivatelinkmodels.CloudServiceProvider(model.CloudServiceProvider.ValueString())
	req := &cloudprivatelinkmodels.CreatePrivateLinkEndpointRequest{
		CloudServiceProvider: &csp,
		ServiceGroup:         model.ServiceGroup.ValueString(),
		EnablePrivateDNS:     model.EnablePrivateDNS.ValueBool(),
	}

	if !model.AWSAccountDetails.IsNull() {
		awsDetails := &awsAccountDetails{}
		diags := model.AWSAccountDetails.As(ctx, awsDetails, basetypes.ObjectAsOptions{
			UnhandledNullAsEmpty:    true,
			UnhandledUnknownAsEmpty: true,
		})
		if diags.HasError() {
			tflog.Warn(ctx, "Failed to convert AWS account details")
		}

		details := &cloudprivatelinkmodels.AWSAccountDetails{
			CloudAccountID: awsDetails.CloudAccountID.ValueString(),
			CredentialCrn:  awsDetails.CredentialCRN.ValueString(),
			Region:         awsDetails.Region.ValueString(),
			VpcID:          awsDetails.VpcID.ValueString(),
		}

		if !awsDetails.SubnetIDs.IsNull() && !awsDetails.SubnetIDs.IsUnknown() {
			var subnetIDs []string
			_ = awsDetails.SubnetIDs.ElementsAs(ctx, &subnetIDs, false)
			details.SubnetIds = subnetIDs
		}

		req.AwsAccountDetails = details
	}

	return req
}

func fromModelToAzureRequest(model cloudPrivateLinkResourceModel, ctx context.Context) *cloudprivatelinkmodels.CreatePrivateLinkEndpointRequest {
	tflog.Debug(ctx, "Converting model to Azure CreatePrivateLinkEndpoint request")

	csp := cloudprivatelinkmodels.CloudServiceProvider(model.CloudServiceProvider.ValueString())
	req := &cloudprivatelinkmodels.CreatePrivateLinkEndpointRequest{
		CloudServiceProvider: &csp,
		ServiceGroup:         model.ServiceGroup.ValueString(),
		EnablePrivateDNS:     model.EnablePrivateDNS.ValueBool(),
	}

	if !model.AzureAccountDetails.IsNull() {
		azureDetails := &azureAccountDetails{}
		diags := model.AzureAccountDetails.As(ctx, azureDetails, basetypes.ObjectAsOptions{
			UnhandledNullAsEmpty:    true,
			UnhandledUnknownAsEmpty: true,
		})
		if diags.HasError() {
			tflog.Warn(ctx, "Failed to convert Azure account details")
		}

		req.AzureAccountDetails = &cloudprivatelinkmodels.AzureAccountDetails{
			CredentialCrn:  azureDetails.CredentialCRN.ValueString(),
			SubscriptionID: azureDetails.SubscriptionID.ValueString(),
			ResourceGroup:  azureDetails.ResourceGroup.ValueString(),
			Location:       azureDetails.Location.ValueString(),
			VNetID:         azureDetails.VNetID.ValueString(),
			SubnetID:       azureDetails.SubnetID.ValueString(),
		}
	}

	return req
}

func fromResponseToModel(ctx context.Context, resp *cloudprivatelinkmodels.GetPrivateLinkStatusResponse) cloudPrivateLinkResourceModel {
	model := cloudPrivateLinkResourceModel{
		AWSAccountDetails:   types.ObjectNull(awsAccountDetailsAttrTypes()),
		AzureAccountDetails: types.ObjectNull(azureAccountDetailsAttrTypes()),
	}

	if resp == nil {
		return model
	}

	// TODO: populate model fields from resp

	return model
}
