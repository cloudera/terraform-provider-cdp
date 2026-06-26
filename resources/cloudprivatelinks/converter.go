// Copyright 2026 Cloudera. All Rights Reserved.
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

	"github.com/hashicorp/terraform-plugin-framework/attr"
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

		if !awsDetails.CrossAccountRoleDetails.IsNull() && !awsDetails.CrossAccountRoleDetails.IsUnknown() {
			crossRoleAttrs := awsDetails.CrossAccountRoleDetails.Attributes()
			details.CrossAccountRoleDetails = &cloudprivatelinkmodels.CrossAccountRoleDetails{
				CrossAccountRole: crossRoleAttrs["cross_account_role"].(types.String).ValueString(),
				ExternalID:       crossRoleAttrs["external_id"].(types.String).ValueString(),
			}
		}

		req.AwsAccountDetails = details
	}
	req.ResourceTags = fromModelToResourceTags(ctx, model.ResourceTags)
	return req
}

func fromEndpointStatusesToModel(ctx context.Context, statuses []*cloudprivatelinkmodels.PrivateLinkEndpointStatus) types.List {
	attrTypes := endpointStatusAttrTypes()
	objType := types.ObjectType{AttrTypes: attrTypes}

	if len(statuses) == 0 {
		emptyList, _ := types.ListValue(objType, []attr.Value{})
		return emptyList
	}

	items := make([]attr.Value, 0, len(statuses))
	for _, s := range statuses {
		dnsVals := make([]attr.Value, len(s.DNSNames))
		for i, d := range s.DNSNames {
			dnsVals[i] = types.StringValue(d)
		}
		dnsNames, _ := types.ListValue(types.StringType, dnsVals)

		obj, _ := types.ObjectValue(attrTypes, map[string]attr.Value{
			"service_component":  types.StringValue(s.ServiceComponent),
			"status":             types.StringValue(s.Status),
			"error":              types.StringValue(s.Error),
			"endpoint_id":        types.StringValue(s.EndpointID),
			"dns_names":          dnsNames,
			"creation_timestamp": types.StringValue(s.CreationTimestamp),
		})
		items = append(items, obj)
	}

	list, _ := types.ListValue(objType, items)
	return list
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

		azureReqDetails := &cloudprivatelinkmodels.AzureAccountDetails{
			CredentialCrn:  azureDetails.CredentialCRN.ValueString(),
			SubscriptionID: azureDetails.SubscriptionID.ValueString(),
			ResourceGroup:  azureDetails.ResourceGroup.ValueString(),
			Location:       azureDetails.Location.ValueString(),
			VNetID:         azureDetails.VNetID.ValueString(),
			SubnetID:       azureDetails.SubnetID.ValueString(),
		}

		if !azureDetails.AzureClientSecretCredential.IsNull() && !azureDetails.AzureClientSecretCredential.IsUnknown() {
			clientSecretAttrs := azureDetails.AzureClientSecretCredential.Attributes()
			azureReqDetails.AzureClientSecretCredential = &cloudprivatelinkmodels.AzureClientSecretCredential{
				ClientID:     clientSecretAttrs["client_id"].(types.String).ValueString(),
				ClientSecret: clientSecretAttrs["client_secret"].(types.String).ValueString(),
				TenantID:     clientSecretAttrs["tenant_id"].(types.String).ValueString(),
			}
		}

		req.AzureAccountDetails = azureReqDetails
	}
	req.ResourceTags = fromModelToResourceTags(ctx, model.ResourceTags)
	return req
}

func fromModelToResourceTags(ctx context.Context, tagsList types.List) []*cloudprivatelinkmodels.ResourceTag {
	if tagsList.IsNull() || tagsList.IsUnknown() {
		return nil
	}
	var tags []resourceTag
	_ = tagsList.ElementsAs(ctx, &tags, false)
	result := make([]*cloudprivatelinkmodels.ResourceTag, 0, len(tags))
	for _, t := range tags {
		key := t.Key.ValueString()
		value := t.Value.ValueString()
		result = append(result, &cloudprivatelinkmodels.ResourceTag{Key: &key, Value: &value})
	}
	return result
}
