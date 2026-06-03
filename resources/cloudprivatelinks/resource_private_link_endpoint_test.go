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
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/cloudprivatelinks/client/operations"
	cloudprivatelinkmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/cloudprivatelinks/models"
	"github.com/cloudera/terraform-provider-cdp/mocks"
)

// strPtr returns a pointer to a string value, used for required string fields in generated models.
func strPtr(s string) *string {
	return &s
}

// =============================================================================
// Create Private Link Endpoint Tests
// =============================================================================

func TestCreatePrivateLinkEndpoint_AWS_Success(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}
	trackingID := "track-aws-123"
	csp := cloudprivatelinkmodels.CloudServiceProviderAWS

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
				Region:    "us-east-1",
				VpcID:     "vpc-abc123",
				SubnetIds: []string{"subnet-111", "subnet-222"},
			},
		},
	)
	result, err := mockOps.CreatePrivateLinkEndpoint(params)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.GetPayload())
	assert.Equal(t, trackingID, result.GetPayload().TrackingID)
	mockOps.AssertExpectations(t)
}

func TestCreatePrivateLinkEndpoint_Azure_Success(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}
	trackingID := "track-azure-456"
	csp := cloudprivatelinkmodels.CloudServiceProviderAZURE

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
				Location: "eastus",
				VNetID:   "vnet-abc123",
				SubnetID: "subnet-azure-111",
			},
		},
	)
	result, err := mockOps.CreatePrivateLinkEndpoint(params)

	assert.NoError(t, err)
	assert.NotNil(t, result.GetPayload())
	assert.Equal(t, trackingID, result.GetPayload().TrackingID)
	mockOps.AssertExpectations(t)
}

func TestCreatePrivateLinkEndpoint_WithResourceTags(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}
	trackingID := "track-tagged-789"
	csp := cloudprivatelinkmodels.CloudServiceProviderAWS

	mockOps.On("CreatePrivateLinkEndpoint", mock.AnythingOfType("*operations.CreatePrivateLinkEndpointParams")).
		Return(&operations.CreatePrivateLinkEndpointOK{
			Payload: &cloudprivatelinkmodels.CreatePrivateLinkEndpointResponse{
				TrackingID: trackingID,
			},
		}, nil)

	params := operations.NewCreatePrivateLinkEndpointParams().WithInput(
		&cloudprivatelinkmodels.CreatePrivateLinkEndpointRequest{
			CloudServiceProvider: &csp,
			ResourceTags: []*cloudprivatelinkmodels.ResourceTag{
				{Key: strPtr("env"), Value: strPtr("production")},
				{Key: strPtr("team"), Value: strPtr("platform")},
			},
			AwsAccountDetails: &cloudprivatelinkmodels.AWSAccountDetails{
				Region: "us-west-2",
				VpcID:  "vpc-tagged",
			},
		},
	)
	result, err := mockOps.CreatePrivateLinkEndpoint(params)

	assert.NoError(t, err)
	assert.Equal(t, trackingID, result.GetPayload().TrackingID)
	mockOps.AssertExpectations(t)
}

func TestCreatePrivateLinkEndpoint_APIError(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}
	csp := cloudprivatelinkmodels.CloudServiceProviderAWS

	mockOps.On("CreatePrivateLinkEndpoint", mock.AnythingOfType("*operations.CreatePrivateLinkEndpointParams")).
		Return(nil, errors.New("API unavailable"))

	params := operations.NewCreatePrivateLinkEndpointParams().WithInput(
		&cloudprivatelinkmodels.CreatePrivateLinkEndpointRequest{
			CloudServiceProvider: &csp,
			AwsAccountDetails: &cloudprivatelinkmodels.AWSAccountDetails{
				Region: "us-east-1",
				VpcID:  "vpc-abc123",
			},
		},
	)
	result, err := mockOps.CreatePrivateLinkEndpoint(params)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "API unavailable")
	mockOps.AssertExpectations(t)
}

// =============================================================================
// Delete Private Link Endpoint Tests
// =============================================================================

func TestDeletePrivateLinkEndpoint_AWS_Success(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}
	trackingID := "delete-track-aws"
	csp := cloudprivatelinkmodels.CloudServiceProviderAWS

	mockOps.On("DeletePrivateLinkEndpoint", mock.AnythingOfType("*operations.DeletePrivateLinkEndpointParams")).
		Return(&operations.DeletePrivateLinkEndpointOK{
			Payload: &cloudprivatelinkmodels.DeletePrivateLinkEndpointResponse{
				TrackingID: trackingID,
			},
		}, nil)

	params := operations.NewDeletePrivateLinkEndpointParams().WithInput(
		&cloudprivatelinkmodels.DeletePrivateLinkEndpointRequest{
			CloudServiceProvider: &csp,
			AwsAccountInfo: &cloudprivatelinkmodels.AWSAccountInfo{
				Region: "us-east-1",
				VpcID:  "vpc-abc123",
			},
		},
	)
	result, err := mockOps.DeletePrivateLinkEndpoint(params)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.GetPayload())
	assert.Equal(t, trackingID, result.GetPayload().TrackingID)
	mockOps.AssertExpectations(t)
}

func TestDeletePrivateLinkEndpoint_Azure_Success(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}
	trackingID := "delete-track-azure"
	csp := cloudprivatelinkmodels.CloudServiceProviderAZURE

	mockOps.On("DeletePrivateLinkEndpoint", mock.AnythingOfType("*operations.DeletePrivateLinkEndpointParams")).
		Return(&operations.DeletePrivateLinkEndpointOK{
			Payload: &cloudprivatelinkmodels.DeletePrivateLinkEndpointResponse{
				TrackingID: trackingID,
			},
		}, nil)

	params := operations.NewDeletePrivateLinkEndpointParams().WithInput(
		&cloudprivatelinkmodels.DeletePrivateLinkEndpointRequest{
			CloudServiceProvider: &csp,
			AzureAccountInfo: &cloudprivatelinkmodels.AzureAccountInfo{
				Location: "eastus",
				VNetID:   "vnet-abc123",
			},
		},
	)
	result, err := mockOps.DeletePrivateLinkEndpoint(params)

	assert.NoError(t, err)
	assert.NotNil(t, result.GetPayload())
	assert.Equal(t, trackingID, result.GetPayload().TrackingID)
	mockOps.AssertExpectations(t)
}

func TestDeletePrivateLinkEndpoint_APIError(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}
	csp := cloudprivatelinkmodels.CloudServiceProviderAWS

	mockOps.On("DeletePrivateLinkEndpoint", mock.AnythingOfType("*operations.DeletePrivateLinkEndpointParams")).
		Return(nil, errors.New("endpoint not found"))

	params := operations.NewDeletePrivateLinkEndpointParams().WithInput(
		&cloudprivatelinkmodels.DeletePrivateLinkEndpointRequest{
			CloudServiceProvider: &csp,
			AwsAccountInfo: &cloudprivatelinkmodels.AWSAccountInfo{
				Region: "us-east-1",
				VpcID:  "vpc-nonexistent",
			},
		},
	)
	result, err := mockOps.DeletePrivateLinkEndpoint(params)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "endpoint not found")
	mockOps.AssertExpectations(t)
}

// =============================================================================
// List Private Link Endpoint Statuses (Polling) Tests
// =============================================================================

func TestListPrivateLinkEndpointStatuses_AllSuccess(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}

	mockOps.On("ListPrivateLinkEndpointStatuses", mock.AnythingOfType("*operations.ListPrivateLinkEndpointStatusesParams")).
		Return(&operations.ListPrivateLinkEndpointStatusesOK{
			Payload: &cloudprivatelinkmodels.ListPrivateLinkEndpointStatusesResponse{
				PrivatelinkEndpoints: []*cloudprivatelinkmodels.PrivateLinkEndpointStatus{
					{ServiceComponent: "API", Status: "SUCCESS", EndpointID: "vpce-api-111", DNSNames: []string{"api.vpce.amazonaws.com"}},
					{ServiceComponent: "IAM", Status: "SUCCESS", EndpointID: "vpce-iam-222", DNSNames: []string{"iam.vpce.amazonaws.com"}},
					{ServiceComponent: "DBUSAPI", Status: "SUCCESS", EndpointID: "vpce-dbus-333", DNSNames: []string{"dbus.vpce.amazonaws.com"}},
				},
			},
		}, nil)

	params := operations.NewListPrivateLinkEndpointStatusesParams().WithInput(
		&cloudprivatelinkmodels.ListPrivateLinkEndpointStatusesRequest{
			TrackingID: strPtr("track-poll-success"),
		},
	)
	result, err := mockOps.ListPrivateLinkEndpointStatuses(params)

	assert.NoError(t, err)
	assert.NotNil(t, result.GetPayload())
	endpoints := result.GetPayload().PrivatelinkEndpoints
	assert.Len(t, endpoints, 3)
	for _, ep := range endpoints {
		assert.Equal(t, "SUCCESS", ep.Status)
		assert.NotEmpty(t, ep.EndpointID)
		assert.NotEmpty(t, ep.DNSNames) // DNSNames, not DnsNames
	}
	mockOps.AssertExpectations(t)
}

func TestListPrivateLinkEndpointStatuses_PartialInProgress(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}

	mockOps.On("ListPrivateLinkEndpointStatuses", mock.AnythingOfType("*operations.ListPrivateLinkEndpointStatusesParams")).
		Return(&operations.ListPrivateLinkEndpointStatusesOK{
			Payload: &cloudprivatelinkmodels.ListPrivateLinkEndpointStatusesResponse{
				PrivatelinkEndpoints: []*cloudprivatelinkmodels.PrivateLinkEndpointStatus{
					{ServiceComponent: "API", Status: "SUCCESS", EndpointID: "vpce-api-111"},
					{ServiceComponent: "IAM", Status: "IN_PROGRESS"},
					{ServiceComponent: "DBUSAPI", Status: "IN_PROGRESS"},
				},
			},
		}, nil)

	params := operations.NewListPrivateLinkEndpointStatusesParams().WithInput(
		&cloudprivatelinkmodels.ListPrivateLinkEndpointStatusesRequest{
			TrackingID: strPtr("track-poll-mixed"),
		},
	)
	result, err := mockOps.ListPrivateLinkEndpointStatuses(params)

	assert.NoError(t, err)
	statusMap := make(map[string]string)
	for _, ep := range result.GetPayload().PrivatelinkEndpoints {
		statusMap[ep.ServiceComponent] = ep.Status
	}
	assert.Equal(t, "SUCCESS", statusMap["API"])
	assert.Equal(t, "IN_PROGRESS", statusMap["IAM"])
	assert.Equal(t, "IN_PROGRESS", statusMap["DBUSAPI"])
	mockOps.AssertExpectations(t)
}

func TestListPrivateLinkEndpointStatuses_EndpointError(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}
	errMsg := "subnet capacity exceeded in availability zone"

	mockOps.On("ListPrivateLinkEndpointStatuses", mock.AnythingOfType("*operations.ListPrivateLinkEndpointStatusesParams")).
		Return(&operations.ListPrivateLinkEndpointStatusesOK{
			Payload: &cloudprivatelinkmodels.ListPrivateLinkEndpointStatusesResponse{
				PrivatelinkEndpoints: []*cloudprivatelinkmodels.PrivateLinkEndpointStatus{
					{ServiceComponent: "API", Status: "SUCCESS", EndpointID: "vpce-api-111"},
					{ServiceComponent: "IAM", Status: "ERROR", Error: errMsg},
				},
			},
		}, nil)

	params := operations.NewListPrivateLinkEndpointStatusesParams().WithInput(
		&cloudprivatelinkmodels.ListPrivateLinkEndpointStatusesRequest{
			TrackingID: strPtr("track-poll-endpoint-error"),
		},
	)
	result, err := mockOps.ListPrivateLinkEndpointStatuses(params)

	assert.NoError(t, err) // API call succeeded; the error is in the endpoint status
	var errEndpoints []*cloudprivatelinkmodels.PrivateLinkEndpointStatus
	for _, ep := range result.GetPayload().PrivatelinkEndpoints {
		if ep.Status == "ERROR" {
			errEndpoints = append(errEndpoints, ep)
		}
	}
	assert.Len(t, errEndpoints, 1)
	assert.Equal(t, "IAM", errEndpoints[0].ServiceComponent)
	assert.Equal(t, errMsg, errEndpoints[0].Error)
	mockOps.AssertExpectations(t)
}

// TestListPrivateLinkEndpointStatuses_PollProgression simulates the polling loop:
// first response is IN_PROGRESS, second response is SUCCESS.
func TestListPrivateLinkEndpointStatuses_PollProgression(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}

	mockOps.On("ListPrivateLinkEndpointStatuses", mock.AnythingOfType("*operations.ListPrivateLinkEndpointStatusesParams")).
		Return(&operations.ListPrivateLinkEndpointStatusesOK{
			Payload: &cloudprivatelinkmodels.ListPrivateLinkEndpointStatusesResponse{
				PrivatelinkEndpoints: []*cloudprivatelinkmodels.PrivateLinkEndpointStatus{
					{ServiceComponent: "API", Status: "IN_PROGRESS"},
					{ServiceComponent: "IAM", Status: "IN_PROGRESS"},
				},
			},
		}, nil).Once()

	mockOps.On("ListPrivateLinkEndpointStatuses", mock.AnythingOfType("*operations.ListPrivateLinkEndpointStatusesParams")).
		Return(&operations.ListPrivateLinkEndpointStatusesOK{
			Payload: &cloudprivatelinkmodels.ListPrivateLinkEndpointStatusesResponse{
				PrivatelinkEndpoints: []*cloudprivatelinkmodels.PrivateLinkEndpointStatus{
					{ServiceComponent: "API", Status: "SUCCESS", EndpointID: "vpce-api-111"},
					{ServiceComponent: "IAM", Status: "SUCCESS", EndpointID: "vpce-iam-222"},
				},
			},
		}, nil).Once()

	params := operations.NewListPrivateLinkEndpointStatusesParams().WithInput(
		&cloudprivatelinkmodels.ListPrivateLinkEndpointStatusesRequest{
			TrackingID: strPtr("track-poll-progression"),
		},
	)

	firstPoll, err := mockOps.ListPrivateLinkEndpointStatuses(params)
	assert.NoError(t, err)
	for _, ep := range firstPoll.GetPayload().PrivatelinkEndpoints {
		assert.Equal(t, "IN_PROGRESS", ep.Status)
		assert.Empty(t, ep.EndpointID)
	}

	secondPoll, err := mockOps.ListPrivateLinkEndpointStatuses(params)
	assert.NoError(t, err)
	for _, ep := range secondPoll.GetPayload().PrivatelinkEndpoints {
		assert.Equal(t, "SUCCESS", ep.Status)
		assert.NotEmpty(t, ep.EndpointID)
	}

	mockOps.AssertExpectations(t)
}

func TestListPrivateLinkEndpointStatuses_APIError(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}

	mockOps.On("ListPrivateLinkEndpointStatuses", mock.AnythingOfType("*operations.ListPrivateLinkEndpointStatusesParams")).
		Return(nil, errors.New("service unavailable"))

	params := operations.NewListPrivateLinkEndpointStatusesParams().WithInput(
		&cloudprivatelinkmodels.ListPrivateLinkEndpointStatusesRequest{
			TrackingID: strPtr("track-poll-apierr"),
		},
	)
	result, err := mockOps.ListPrivateLinkEndpointStatuses(params)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "service unavailable")
	mockOps.AssertExpectations(t)
}

// =============================================================================
// Get Private Link Status Tests
// =============================================================================

func TestGetPrivateLinkStatus_AWS_Enabled(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}
	csp := cloudprivatelinkmodels.CloudServiceProviderAWS

	mockOps.On("GetPrivateLinkStatus", mock.AnythingOfType("*operations.GetPrivateLinkStatusParams")).
		Return(&operations.GetPrivateLinkStatusOK{
			Payload: &cloudprivatelinkmodels.GetPrivateLinkStatusResponse{
				Status: cloudprivatelinkmodels.PrivateLinkStatusENABLED,
			},
		}, nil)

	params := operations.NewGetPrivateLinkStatusParams().WithInput(
		&cloudprivatelinkmodels.GetPrivateLinkStatusRequest{
			CloudServiceProvider: csp, // value type: cloudServiceProvider is not required in swagger
			VpcID:                "vpc-abc123",
		},
	)
	result, err := mockOps.GetPrivateLinkStatus(params)

	assert.NoError(t, err)
	assert.NotNil(t, result.GetPayload())
	assert.Equal(t, cloudprivatelinkmodels.PrivateLinkStatusENABLED, result.GetPayload().Status)
	mockOps.AssertExpectations(t)
}

func TestGetPrivateLinkStatus_AWS_Disabled(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}
	csp := cloudprivatelinkmodels.CloudServiceProviderAWS

	mockOps.On("GetPrivateLinkStatus", mock.AnythingOfType("*operations.GetPrivateLinkStatusParams")).
		Return(&operations.GetPrivateLinkStatusOK{
			Payload: &cloudprivatelinkmodels.GetPrivateLinkStatusResponse{
				Status: cloudprivatelinkmodels.PrivateLinkStatusDISABLED,
			},
		}, nil)

	params := operations.NewGetPrivateLinkStatusParams().WithInput(
		&cloudprivatelinkmodels.GetPrivateLinkStatusRequest{
			CloudServiceProvider: csp,
			VpcID:                "vpc-abc123",
		},
	)
	result, err := mockOps.GetPrivateLinkStatus(params)

	assert.NoError(t, err)
	assert.Equal(t, cloudprivatelinkmodels.PrivateLinkStatusDISABLED, result.GetPayload().Status)
	mockOps.AssertExpectations(t)
}

func TestGetPrivateLinkStatus_Azure_Enabled(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}
	csp := cloudprivatelinkmodels.CloudServiceProviderAZURE

	mockOps.On("GetPrivateLinkStatus", mock.AnythingOfType("*operations.GetPrivateLinkStatusParams")).
		Return(&operations.GetPrivateLinkStatusOK{
			Payload: &cloudprivatelinkmodels.GetPrivateLinkStatusResponse{
				Status: cloudprivatelinkmodels.PrivateLinkStatusENABLED,
			},
		}, nil)

	params := operations.NewGetPrivateLinkStatusParams().WithInput(
		&cloudprivatelinkmodels.GetPrivateLinkStatusRequest{
			CloudServiceProvider: csp,
			VNetID:               "vnet-abc123",
		},
	)
	result, err := mockOps.GetPrivateLinkStatus(params)

	assert.NoError(t, err)
	assert.Equal(t, cloudprivatelinkmodels.PrivateLinkStatusENABLED, result.GetPayload().Status)
	mockOps.AssertExpectations(t)
}

// TestGetPrivateLinkStatus_ErrorStatus verifies an ERROR status is returned by the API
// without a transport-level error — the caller must inspect Status separately.
func TestGetPrivateLinkStatus_ErrorStatus(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}
	csp := cloudprivatelinkmodels.CloudServiceProviderAWS

	mockOps.On("GetPrivateLinkStatus", mock.AnythingOfType("*operations.GetPrivateLinkStatusParams")).
		Return(&operations.GetPrivateLinkStatusOK{
			Payload: &cloudprivatelinkmodels.GetPrivateLinkStatusResponse{
				Status: cloudprivatelinkmodels.PrivateLinkStatusERROR,
			},
		}, nil)

	params := operations.NewGetPrivateLinkStatusParams().WithInput(
		&cloudprivatelinkmodels.GetPrivateLinkStatusRequest{
			CloudServiceProvider: csp,
			VpcID:                "vpc-abc123",
		},
	)
	result, err := mockOps.GetPrivateLinkStatus(params)

	assert.NoError(t, err)
	assert.Equal(t, cloudprivatelinkmodels.PrivateLinkStatusERROR, result.GetPayload().Status)
	mockOps.AssertExpectations(t)
}

func TestGetPrivateLinkStatus_APIError(t *testing.T) {
	mockOps := &mocks.MockCloudPrivateLinksClientService{}
	csp := cloudprivatelinkmodels.CloudServiceProviderAWS

	mockOps.On("GetPrivateLinkStatus", mock.AnythingOfType("*operations.GetPrivateLinkStatusParams")).
		Return(nil, errors.New("internal server error"))

	params := operations.NewGetPrivateLinkStatusParams().WithInput(
		&cloudprivatelinkmodels.GetPrivateLinkStatusRequest{
			CloudServiceProvider: csp,
			VpcID:                "vpc-abc123",
		},
	)
	result, err := mockOps.GetPrivateLinkStatus(params)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "internal server error")
	mockOps.AssertExpectations(t)
}
