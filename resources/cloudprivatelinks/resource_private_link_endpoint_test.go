// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package cloudprivatelinks_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/cloudprivatelinks/client/operations"
	cloudprivatelinkmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/cloudprivatelinks/models"
	"github.com/cloudera/terraform-provider-cdp/mocks"
)

func TestCreatePrivateLinkEndpoint_AWS(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}
	trackingID := "track-123"
	csp := cloudprivatelinkmodels.CloudServiceProviderAWS
	region := "us-east-1"
	vpcID := "vpc-abc123"

	mockOps.On("CreatePrivateLinkEndpoint", mock.AnythingOfType("*operations.CreatePrivateLinkEndpointParams")).
		Return(&operations.CreatePrivateLinkEndpointOK{
			Payload: &cloudprivatelinkmodels.CreatePrivateLinkEndpointResponse{
				TrackingID: trackingID,
			},
		}, nil)

	params := operations.NewCreatePrivateLinkEndpointParams().WithInput(
		&cloudprivatelinkmodels.CreatePrivateLinkEndpointRequest{
			CloudServiceProvider: &csp,
			AwsAccountDetails: &cloudprivatelinkmodels.AWSAccountDetails{
				Region: region,
				VpcID:  vpcID,
			},
		},
	)
	result, err := mockOps.CreatePrivateLinkEndpoint(params)

	assert.NoError(t, err)
	assert.Equal(t, trackingID, result.GetPayload().TrackingID)
	mockOps.AssertExpectations(t)
}

func TestCreatePrivateLinkEndpoint_Azure(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}
	trackingID := "track-456"
	csp := cloudprivatelinkmodels.CloudServiceProviderAZURE
	location := "eastus"
	vnetID := "vnet-abc123"

	mockOps.On("CreatePrivateLinkEndpoint", mock.AnythingOfType("*operations.CreatePrivateLinkEndpointParams")).
		Return(&operations.CreatePrivateLinkEndpointOK{
			Payload: &cloudprivatelinkmodels.CreatePrivateLinkEndpointResponse{
				TrackingID: trackingID,
			},
		}, nil)

	params := operations.NewCreatePrivateLinkEndpointParams().WithInput(
		&cloudprivatelinkmodels.CreatePrivateLinkEndpointRequest{
			CloudServiceProvider: &csp,
			AzureAccountDetails: &cloudprivatelinkmodels.AzureAccountDetails{
				Location: location,
				VNetID:   vnetID,
			},
		},
	)
	result, err := mockOps.CreatePrivateLinkEndpoint(params)

	assert.NoError(t, err)
	assert.Equal(t, trackingID, result.GetPayload().TrackingID)
	mockOps.AssertExpectations(t)
}

func TestCreatePrivateLinkEndpoint_UnsupportedCSP(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}
	mockOps.AssertNotCalled(t, "CreatePrivateLinkEndpoint")
	mockOps.AssertExpectations(t)
}

func TestDeletePrivateLinkEndpoint_AWS(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}
	csp := cloudprivatelinkmodels.CloudServiceProviderAWS

	mockOps.On("DeletePrivateLinkEndpoint", mock.AnythingOfType("*operations.DeletePrivateLinkEndpointParams")).
		Return(&operations.DeletePrivateLinkEndpointOK{}, nil)

	params := operations.NewDeletePrivateLinkEndpointParams().WithInput(
		&cloudprivatelinkmodels.DeletePrivateLinkEndpointRequest{
			CloudServiceProvider: &csp,
		},
	)
	_, err := mockOps.DeletePrivateLinkEndpoint(params)

	assert.NoError(t, err)
	mockOps.AssertExpectations(t)
}

func TestDeletePrivateLinkEndpoint_Azure(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}
	csp := cloudprivatelinkmodels.CloudServiceProviderAZURE

	mockOps.On("DeletePrivateLinkEndpoint", mock.AnythingOfType("*operations.DeletePrivateLinkEndpointParams")).
		Return(&operations.DeletePrivateLinkEndpointOK{}, nil)

	params := operations.NewDeletePrivateLinkEndpointParams().WithInput(
		&cloudprivatelinkmodels.DeletePrivateLinkEndpointRequest{
			CloudServiceProvider: &csp,
		},
	)
	_, err := mockOps.DeletePrivateLinkEndpoint(params)

	assert.NoError(t, err)
	mockOps.AssertExpectations(t)
}
