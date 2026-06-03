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
)

func TestFromModelToAwsRequest(t *testing.T) {
	ctx := context.Background()

	subnetIDs, _ := types.ListValue(types.StringType, []attr.Value{
		types.StringValue("subnet-111"),
		types.StringValue("subnet-222"),
	})

	awsObj, _ := types.ObjectValue(awsAccountDetailsAttrTypes(), map[string]attr.Value{
		"cloud_account_id": types.StringValue("123456789"),
		"credential_crn":   types.StringValue("crn:cdp:iam:us-west-1:abc:credential:def"),
		"region":           types.StringValue("us-east-1"),
		"vpc_id":           types.StringValue("vpc-abc123"),
		"subnet_ids":       subnetIDs,
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
	assert.Equal(t, "us-east-1", req.AwsAccountDetails.Region)
	assert.Equal(t, "vpc-abc123", req.AwsAccountDetails.VpcID)
	assert.Equal(t, []string{"subnet-111", "subnet-222"}, req.AwsAccountDetails.SubnetIds)
}

func TestFromModelToAwsRequest_NoSubnets(t *testing.T) {
	ctx := context.Background()

	awsObj, _ := types.ObjectValue(awsAccountDetailsAttrTypes(), map[string]attr.Value{
		"cloud_account_id": types.StringValue("123456789"),
		"credential_crn":   types.StringValue("crn:cdp:iam:us-west-1:abc:credential:def"),
		"region":           types.StringValue("us-east-1"),
		"vpc_id":           types.StringValue("vpc-abc123"),
		"subnet_ids":       types.ListNull(types.StringType),
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

func TestFromModelToAzureRequest(t *testing.T) {
	ctx := context.Background()

	azureObj, _ := types.ObjectValue(azureAccountDetailsAttrTypes(), map[string]attr.Value{
		"credential_crn":  types.StringValue("crn:cdp:iam:us-west-1:abc:credential:xyz"),
		"subscription_id": types.StringValue("sub-111"),
		"resource_group":  types.StringValue("my-rg"),
		"location":        types.StringValue("eastus"),
		"vnet_id":         types.StringValue("vnet-abc"),
		"subnet_id":       types.StringValue("subnet-abc"),
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
	assert.Equal(t, "sub-111", req.AzureAccountDetails.SubscriptionID)
	assert.Equal(t, "my-rg", req.AzureAccountDetails.ResourceGroup)
	assert.Equal(t, "eastus", req.AzureAccountDetails.Location)
	assert.Equal(t, "vnet-abc", req.AzureAccountDetails.VNetID)
	assert.Equal(t, "subnet-abc", req.AzureAccountDetails.SubnetID)
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
