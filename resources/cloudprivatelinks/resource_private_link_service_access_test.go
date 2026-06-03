// Copyright 2024 Cloudera. All Rights Reserved.
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

func TestAuthorizePrivateLinkServiceAccess_AWS(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}
	csp := cloudprivatelinkmodels.CloudServiceProviderAWS

	mockOps.On("AuthorizePrivateLinkServiceAccess", mock.AnythingOfType("*operations.AuthorizePrivateLinkServiceAccessParams")).
		Return(&operations.AuthorizePrivateLinkServiceAccessOK{
			Payload: &cloudprivatelinkmodels.AuthorizePrivateLinkServiceAccessResponse{
				Status: "AUTHORIZED",
			},
		}, nil)

	params := operations.NewAuthorizePrivateLinkServiceAccessParams().WithInput(
		&cloudprivatelinkmodels.AuthorizePrivateLinkServiceAccessRequest{
			CloudServiceProvider: csp,
			CloudAccountID:       "123456789012",
			Region:               "us-east-1",
			ServiceGroup:         "CDP-CONTROL-PLANE",
		},
	)
	result, err := mockOps.AuthorizePrivateLinkServiceAccess(params)

	assert.NoError(t, err)
	assert.Equal(t, "AUTHORIZED", result.GetPayload().Status)
	mockOps.AssertExpectations(t)
}

func TestAuthorizePrivateLinkServiceAccess_Azure(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}
	csp := cloudprivatelinkmodels.CloudServiceProviderAZURE

	mockOps.On("AuthorizePrivateLinkServiceAccess", mock.AnythingOfType("*operations.AuthorizePrivateLinkServiceAccessParams")).
		Return(&operations.AuthorizePrivateLinkServiceAccessOK{
			Payload: &cloudprivatelinkmodels.AuthorizePrivateLinkServiceAccessResponse{
				Status: "AUTHORIZED",
			},
		}, nil)

	params := operations.NewAuthorizePrivateLinkServiceAccessParams().WithInput(
		&cloudprivatelinkmodels.AuthorizePrivateLinkServiceAccessRequest{
			CloudServiceProvider: csp,
			SubscriptionID:       "sub-abc-123",
			Region:               "eastus",
			ServiceGroup:         "CDP-CONTROL-PLANE",
		},
	)
	result, err := mockOps.AuthorizePrivateLinkServiceAccess(params)

	assert.NoError(t, err)
	assert.Equal(t, "AUTHORIZED", result.GetPayload().Status)
	mockOps.AssertExpectations(t)
}

func TestRevokePrivateLinkServiceAccess_AWS(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}
	csp := cloudprivatelinkmodels.CloudServiceProviderAWS

	mockOps.On("RevokePrivateLinkServiceAccess", mock.AnythingOfType("*operations.RevokePrivateLinkServiceAccessParams")).
		Return(&operations.RevokePrivateLinkServiceAccessOK{
			Payload: &cloudprivatelinkmodels.RevokePrivateLinkServiceAccessResponse{
				Status: "REVOKED",
			},
		}, nil)

	params := operations.NewRevokePrivateLinkServiceAccessParams().WithInput(
		&cloudprivatelinkmodels.RevokePrivateLinkServiceAccessRequest{
			CloudServiceProvider: csp,
			CloudAccountID:       "123456789012",
			Region:               "us-east-1",
			ServiceGroup:         "CDP-CONTROL-PLANE",
		},
	)
	result, err := mockOps.RevokePrivateLinkServiceAccess(params)

	assert.NoError(t, err)
	assert.Equal(t, "REVOKED", result.GetPayload().Status)
	mockOps.AssertExpectations(t)
}

func TestRevokePrivateLinkServiceAccess_Azure(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}
	csp := cloudprivatelinkmodels.CloudServiceProviderAZURE

	mockOps.On("RevokePrivateLinkServiceAccess", mock.AnythingOfType("*operations.RevokePrivateLinkServiceAccessParams")).
		Return(&operations.RevokePrivateLinkServiceAccessOK{
			Payload: &cloudprivatelinkmodels.RevokePrivateLinkServiceAccessResponse{
				Status: "REVOKED",
			},
		}, nil)

	params := operations.NewRevokePrivateLinkServiceAccessParams().WithInput(
		&cloudprivatelinkmodels.RevokePrivateLinkServiceAccessRequest{
			CloudServiceProvider: csp,
			SubscriptionID:       "sub-abc-123",
			Region:               "eastus",
			ServiceGroup:         "CDP-CONTROL-PLANE",
		},
	)
	result, err := mockOps.RevokePrivateLinkServiceAccess(params)

	assert.NoError(t, err)
	assert.Equal(t, "REVOKED", result.GetPayload().Status)
	mockOps.AssertExpectations(t)
}
