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
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"

	cloudprivatelinkmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/cloudprivatelinks/models"
)

// ── AWS ──────────────────────────────────────────────────────────────────────

func TestFromModelToAwsRequest_WithCredentialCrn(t *testing.T) {
	ctx := context.Background()

	subnetIDs, _ := types.ListValue(types.StringType, []attr.Value{
		types.StringValue("subnet-111"),
		types.StringValue("subnet-222"),
	})

	awsObj, _ := types.ObjectValue(awsAccountDetailsAttrTypes(), map[string]attr.Value{
		"cloud_account_id":           types.StringValue("123456789"),
		"cross_account_role_details": types.ObjectNull(crossAccountRoleDetailsAttrTypes()),
		"credential_crn":             types.StringValue("crn:cdp:iam:us-west-1:abc:credential:def"),
		"region":                     types.StringValue("us-east-1"),
		"vpc_id":                     types.StringValue("vpc-abc123"),
		"subnet_ids":                 subnetIDs,
	})

	model := cloudPrivateLinkResourceModel{
		CloudServiceProvider: types.StringValue("AWS"),
		ServiceGroup:         types.StringValue("COMPUTE"),
		EnablePrivateDNS:     types.BoolValue(true),
		AWSAccountDetails:    awsObj,
		AzureAccountDetails:  types.ObjectNull(azureAccountDetailsAttrTypes()),
	}

	req := fromModelToAwsRequest(model, ctx)

	assert.NotNil(t, req)
	assert.Equal(t, "AWS", string(*req.CloudServiceProvider))
	assert.Equal(t, "COMPUTE", req.ServiceGroup)
	assert.True(t, req.EnablePrivateDNS)
	assert.NotNil(t, req.AwsAccountDetails)
	assert.Equal(t, "123456789", req.AwsAccountDetails.CloudAccountID)
	assert.Equal(t, "crn:cdp:iam:us-west-1:abc:credential:def", req.AwsAccountDetails.CredentialCrn)
	assert.Nil(t, req.AwsAccountDetails.CrossAccountRoleDetails)
	assert.Equal(t, "us-east-1", req.AwsAccountDetails.Region)
	assert.Equal(t, "vpc-abc123", req.AwsAccountDetails.VpcID)
	assert.Equal(t, []string{"subnet-111", "subnet-222"}, req.AwsAccountDetails.SubnetIds)
}

func TestFromModelToAwsRequest_WithCrossAccountRole(t *testing.T) {
	ctx := context.Background()

	subnetIDs, _ := types.ListValue(types.StringType, []attr.Value{
		types.StringValue("subnet-111"),
	})

	crossAccountRoleObj, _ := types.ObjectValue(crossAccountRoleDetailsAttrTypes(), map[string]attr.Value{
		"cross_account_role": types.StringValue("arn:aws:iam::123456789:role/MyRole"),
		"external_id":        types.StringValue("ext-id-123"),
	})

	awsObj, _ := types.ObjectValue(awsAccountDetailsAttrTypes(), map[string]attr.Value{
		"cloud_account_id":           types.StringValue("123456789"),
		"cross_account_role_details": crossAccountRoleObj,
		"credential_crn":             types.StringNull(),
		"region":                     types.StringValue("us-east-1"),
		"vpc_id":                     types.StringValue("vpc-abc123"),
		"subnet_ids":                 subnetIDs,
	})

	model := cloudPrivateLinkResourceModel{
		CloudServiceProvider: types.StringValue("AWS"),
		ServiceGroup:         types.StringValue("COMPUTE"),
		EnablePrivateDNS:     types.BoolValue(true),
		AWSAccountDetails:    awsObj,
		AzureAccountDetails:  types.ObjectNull(azureAccountDetailsAttrTypes()),
	}

	req := fromModelToAwsRequest(model, ctx)

	assert.NotNil(t, req)
	assert.NotNil(t, req.AwsAccountDetails)
	assert.NotNil(t, req.AwsAccountDetails.CrossAccountRoleDetails)
	assert.Equal(t, "arn:aws:iam::123456789:role/MyRole", req.AwsAccountDetails.CrossAccountRoleDetails.CrossAccountRole)
	assert.Equal(t, "ext-id-123", req.AwsAccountDetails.CrossAccountRoleDetails.ExternalID)
	assert.Empty(t, req.AwsAccountDetails.CredentialCrn)
}

func TestFromModelToAwsRequest_NoSubnets(t *testing.T) {
	ctx := context.Background()

	awsObj, _ := types.ObjectValue(awsAccountDetailsAttrTypes(), map[string]attr.Value{
		"cloud_account_id":           types.StringValue("123456789"),
		"cross_account_role_details": types.ObjectNull(crossAccountRoleDetailsAttrTypes()),
		"credential_crn":             types.StringValue("crn:cdp:iam:us-west-1:abc:credential:def"),
		"region":                     types.StringValue("us-east-1"),
		"vpc_id":                     types.StringValue("vpc-abc123"),
		"subnet_ids":                 types.ListNull(types.StringType),
	})

	model := cloudPrivateLinkResourceModel{
		CloudServiceProvider: types.StringValue("AWS"),
		ServiceGroup:         types.StringValue("COMPUTE"),
		EnablePrivateDNS:     types.BoolValue(false),
		AWSAccountDetails:    awsObj,
		AzureAccountDetails:  types.ObjectNull(azureAccountDetailsAttrTypes()),
	}

	req := fromModelToAwsRequest(model, ctx)

	assert.NotNil(t, req)
	assert.NotNil(t, req.AwsAccountDetails)
	assert.Nil(t, req.AwsAccountDetails.SubnetIds)
}

func TestFromModelToAwsRequest_NullAccountDetails(t *testing.T) {
	ctx := context.Background()

	model := cloudPrivateLinkResourceModel{
		CloudServiceProvider: types.StringValue("AWS"),
		ServiceGroup:         types.StringValue("COMPUTE"),
		EnablePrivateDNS:     types.BoolValue(false),
		AWSAccountDetails:    types.ObjectNull(awsAccountDetailsAttrTypes()),
		AzureAccountDetails:  types.ObjectNull(azureAccountDetailsAttrTypes()),
	}

	req := fromModelToAwsRequest(model, ctx)

	assert.NotNil(t, req)
	assert.Nil(t, req.AwsAccountDetails)
}

// ── Azure ────────────────────────────────────────────────────────────────────

func TestFromModelToAzureRequest_WithCredentialCrn(t *testing.T) {
	ctx := context.Background()

	azureObj, _ := types.ObjectValue(azureAccountDetailsAttrTypes(), map[string]attr.Value{
		"azure_client_secret_credential": types.ObjectNull(azureClientSecretCredentialAttrTypes()),
		"credential_crn":                 types.StringValue("crn:cdp:iam:us-west-1:abc:credential:xyz"),
		"subscription_id":                types.StringValue("sub-111"),
		"resource_group":                 types.StringValue("my-rg"),
		"location":                       types.StringValue("eastus"),
		"vnet_id":                        types.StringValue("vnet-abc"),
		"subnet_id":                      types.StringValue("subnet-abc"),
	})

	model := cloudPrivateLinkResourceModel{
		CloudServiceProvider: types.StringValue("AZURE"),
		ServiceGroup:         types.StringValue("COMPUTE"),
		EnablePrivateDNS:     types.BoolValue(true),
		AWSAccountDetails:    types.ObjectNull(awsAccountDetailsAttrTypes()),
		AzureAccountDetails:  azureObj,
	}

	req := fromModelToAzureRequest(model, ctx)

	assert.NotNil(t, req)
	assert.Equal(t, "AZURE", string(*req.CloudServiceProvider))
	assert.Equal(t, "COMPUTE", req.ServiceGroup)
	assert.True(t, req.EnablePrivateDNS)
	assert.NotNil(t, req.AzureAccountDetails)
	assert.Equal(t, "crn:cdp:iam:us-west-1:abc:credential:xyz", req.AzureAccountDetails.CredentialCrn)
	assert.Nil(t, req.AzureAccountDetails.AzureClientSecretCredential)
	assert.Equal(t, "sub-111", req.AzureAccountDetails.SubscriptionID)
	assert.Equal(t, "my-rg", req.AzureAccountDetails.ResourceGroup)
	assert.Equal(t, "eastus", req.AzureAccountDetails.Location)
	assert.Equal(t, "vnet-abc", req.AzureAccountDetails.VNetID)
	assert.Equal(t, "subnet-abc", req.AzureAccountDetails.SubnetID)
}

func TestFromModelToAzureRequest_WithClientSecretCredential(t *testing.T) {
	ctx := context.Background()

	clientSecretObj, _ := types.ObjectValue(azureClientSecretCredentialAttrTypes(), map[string]attr.Value{
		"client_id":     types.StringValue("client-id-123"),
		"client_secret": types.StringValue("super-secret"),
		"tenant_id":     types.StringValue("tenant-id-456"),
	})

	azureObj, _ := types.ObjectValue(azureAccountDetailsAttrTypes(), map[string]attr.Value{
		"azure_client_secret_credential": clientSecretObj,
		"credential_crn":                 types.StringNull(),
		"subscription_id":                types.StringValue("sub-111"),
		"resource_group":                 types.StringValue("my-rg"),
		"location":                       types.StringValue("eastus"),
		"vnet_id":                        types.StringValue("vnet-abc"),
		"subnet_id":                      types.StringValue("subnet-abc"),
	})

	model := cloudPrivateLinkResourceModel{
		CloudServiceProvider: types.StringValue("AZURE"),
		ServiceGroup:         types.StringValue("COMPUTE"),
		EnablePrivateDNS:     types.BoolValue(true),
		AWSAccountDetails:    types.ObjectNull(awsAccountDetailsAttrTypes()),
		AzureAccountDetails:  azureObj,
	}

	req := fromModelToAzureRequest(model, ctx)

	assert.NotNil(t, req)
	assert.NotNil(t, req.AzureAccountDetails)
	assert.NotNil(t, req.AzureAccountDetails.AzureClientSecretCredential)
	assert.Equal(t, "client-id-123", req.AzureAccountDetails.AzureClientSecretCredential.ClientID)
	assert.Equal(t, "super-secret", req.AzureAccountDetails.AzureClientSecretCredential.ClientSecret)
	assert.Equal(t, "tenant-id-456", req.AzureAccountDetails.AzureClientSecretCredential.TenantID)
	assert.Empty(t, req.AzureAccountDetails.CredentialCrn)
}

func TestFromModelToAzureRequest_NullAccountDetails(t *testing.T) {
	ctx := context.Background()

	model := cloudPrivateLinkResourceModel{
		CloudServiceProvider: types.StringValue("AZURE"),
		ServiceGroup:         types.StringValue("COMPUTE"),
		EnablePrivateDNS:     types.BoolValue(false),
		AWSAccountDetails:    types.ObjectNull(awsAccountDetailsAttrTypes()),
		AzureAccountDetails:  types.ObjectNull(azureAccountDetailsAttrTypes()),
	}

	req := fromModelToAzureRequest(model, ctx)

	assert.NotNil(t, req)
	assert.Nil(t, req.AzureAccountDetails)
}

// ── fromEndpointStatusesToModel ───────────────────────────────────────────────

func TestFromEndpointStatusesToModel_Empty(t *testing.T) {
	ctx := context.Background()

	result := fromEndpointStatusesToModel(ctx, nil)

	assert.False(t, result.IsNull())
	assert.False(t, result.IsUnknown())
	assert.Equal(t, 0, len(result.Elements()))
}

func TestFromEndpointStatusesToModel(t *testing.T) {
	ctx := context.Background()

	statuses := []*cloudprivatelinkmodels.PrivateLinkEndpointStatus{
		{
			ServiceComponent:  "API",
			Status:            "SUCCESS",
			Error:             "",
			EndpointID:        "vpce-abc123",
			DNSNames:          []string{"vpce-abc123.us-east-1.vpce.amazonaws.com"},
			CreationTimestamp: "2024-01-01T00:00:00Z",
		},
		{
			ServiceComponent:  "IAMAPI",
			Status:            "ERROR",
			Error:             "endpoint creation failed",
			EndpointID:        "",
			DNSNames:          []string{},
			CreationTimestamp: "2024-01-01T00:01:00Z",
		},
	}

	result := fromEndpointStatusesToModel(ctx, statuses)

	assert.False(t, result.IsNull())
	assert.Equal(t, 2, len(result.Elements()))

	elems := result.Elements()

	obj0 := elems[0].(types.Object)
	attrs0 := obj0.Attributes()
	assert.Equal(t, types.StringValue("API"), attrs0["service_component"])
	assert.Equal(t, types.StringValue("SUCCESS"), attrs0["status"])
	assert.Equal(t, types.StringValue(""), attrs0["error"])
	assert.Equal(t, types.StringValue("vpce-abc123"), attrs0["endpoint_id"])
	assert.Equal(t, types.StringValue("2024-01-01T00:00:00Z"), attrs0["creation_timestamp"])
	dnsNames0 := attrs0["dns_names"].(types.List)
	assert.Equal(t, 1, len(dnsNames0.Elements()))

	obj1 := elems[1].(types.Object)
	attrs1 := obj1.Attributes()
	assert.Equal(t, types.StringValue("IAMAPI"), attrs1["service_component"])
	assert.Equal(t, types.StringValue("ERROR"), attrs1["status"])
	assert.Equal(t, types.StringValue("endpoint creation failed"), attrs1["error"])
	assert.Equal(t, types.StringValue(""), attrs1["endpoint_id"])
	dnsNames1 := attrs1["dns_names"].(types.List)
	assert.Equal(t, 0, len(dnsNames1.Elements()))
}

func crossAccountRoleDetailsAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"cross_account_role": types.StringType,
		"external_id":        types.StringType,
	}
}

func awsAccountDetailsAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"cloud_account_id":           types.StringType,
		"cross_account_role_details": types.ObjectType{AttrTypes: crossAccountRoleDetailsAttrTypes()},
		"credential_crn":             types.StringType,
		"region":                     types.StringType,
		"vpc_id":                     types.StringType,
		"subnet_ids":                 types.ListType{ElemType: types.StringType},
	}
}

func azureClientSecretCredentialAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"client_id":     types.StringType,
		"client_secret": types.StringType,
		"tenant_id":     types.StringType,
	}
}

func azureAccountDetailsAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"azure_client_secret_credential": types.ObjectType{AttrTypes: azureClientSecretCredentialAttrTypes()},
		"credential_crn":                 types.StringType,
		"subscription_id":                types.StringType,
		"resource_group":                 types.StringType,
		"location":                       types.StringType,
		"vnet_id":                        types.StringType,
		"subnet_id":                      types.StringType,
	}
}
