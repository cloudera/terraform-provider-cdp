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
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	cloudprivatelinksclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/cloudprivatelinks/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/cloudprivatelinks/client/operations"
	cloudprivatelinkmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/cloudprivatelinks/models"
	"github.com/cloudera/terraform-provider-cdp/mocks"
)

// overridePollingTimings replaces package-level timing vars with test-friendly values.
// Do NOT combine with t.Parallel() — the vars are shared package state.
func overridePollingTimings(t *testing.T) {
	t.Helper()
	origDelay, origInterval, origTimeout := pollingDelay, pollingInterval, pollingTimeout
	pollingDelay = 0
	pollingInterval = 50 * time.Millisecond
	pollingTimeout = 5 * time.Second
	t.Cleanup(func() {
		pollingDelay = origDelay
		pollingInterval = origInterval
		pollingTimeout = origTimeout
	})
}

func makeEndpointStatusesOK(statuses []*cloudprivatelinkmodels.PrivateLinkEndpointStatus) *operations.ListPrivateLinkEndpointStatusesOK {
	return &operations.ListPrivateLinkEndpointStatusesOK{
		Payload: &cloudprivatelinkmodels.ListPrivateLinkEndpointStatusesResponse{
			PrivatelinkEndpoints: statuses,
		},
	}
}

// ── waitForEndpointToBeCreated ─────────────────────────────────────────────

func TestWaitForEndpointToBeCreated_AllSuccess(t *testing.T) {
	overridePollingTimings(t)

	mockOps := mocks.NewMockCloudPrivateLinksClientService(t)
	mockOps.EXPECT().
		ListPrivateLinkEndpointStatusesContext(mock.Anything, mock.Anything).
		Return(makeEndpointStatusesOK([]*cloudprivatelinkmodels.PrivateLinkEndpointStatus{
			{ServiceComponent: "API", Status: "SUCCESS"},
			{ServiceComponent: "IAMAPI", Status: "SUCCESS"},
		}), nil)

	plClient := &cloudprivatelinksclient.Cloudprivatelinks{Operations: mockOps}

	statuses, err := waitForEndpointToBeCreated("track-123", plClient, context.Background(), nil)

	require.NoError(t, err)
	require.Len(t, statuses, 2)
}

func TestWaitForEndpointToBeCreated_ErrorStatus(t *testing.T) {
	overridePollingTimings(t)

	mockOps := mocks.NewMockCloudPrivateLinksClientService(t)
	mockOps.EXPECT().
		ListPrivateLinkEndpointStatusesContext(mock.Anything, mock.Anything).
		Return(makeEndpointStatusesOK([]*cloudprivatelinkmodels.PrivateLinkEndpointStatus{
			{ServiceComponent: "API", Status: "SUCCESS"},
			{ServiceComponent: "IAMAPI", Status: "ERROR", Error: "endpoint creation failed"},
		}), nil)

	plClient := &cloudprivatelinksclient.Cloudprivatelinks{Operations: mockOps}

	_, err := waitForEndpointToBeCreated("track-123", plClient, context.Background(), nil)

	require.Error(t, err)
	require.Contains(t, err.Error(), "IAMAPI")
	require.Contains(t, err.Error(), "endpoint creation failed")
}

func TestWaitForEndpointToBeCreated_InProgressThenSuccess(t *testing.T) {
	overridePollingTimings(t)

	mockOps := mocks.NewMockCloudPrivateLinksClientService(t)
	mockOps.EXPECT().
		ListPrivateLinkEndpointStatusesContext(mock.Anything, mock.Anything).
		Return(makeEndpointStatusesOK([]*cloudprivatelinkmodels.PrivateLinkEndpointStatus{
			{ServiceComponent: "API", Status: "IN_PROGRESS"},
		}), nil).Once()
	mockOps.EXPECT().
		ListPrivateLinkEndpointStatusesContext(mock.Anything, mock.Anything).
		Return(makeEndpointStatusesOK([]*cloudprivatelinkmodels.PrivateLinkEndpointStatus{
			{ServiceComponent: "API", Status: "SUCCESS"},
		}), nil).Once()

	plClient := &cloudprivatelinksclient.Cloudprivatelinks{Operations: mockOps}

	statuses, err := waitForEndpointToBeCreated("track-123", plClient, context.Background(), nil)

	require.NoError(t, err)
	require.Len(t, statuses, 1)
	require.Equal(t, "SUCCESS", statuses[0].Status)
}

func TestWaitForEndpointToBeCreated_EmptyEndpoints(t *testing.T) {
	overridePollingTimings(t)

	mockOps := mocks.NewMockCloudPrivateLinksClientService(t)
	mockOps.EXPECT().
		ListPrivateLinkEndpointStatusesContext(mock.Anything, mock.Anything).
		Return(makeEndpointStatusesOK(nil), nil)

	plClient := &cloudprivatelinksclient.Cloudprivatelinks{Operations: mockOps}

	statuses, err := waitForEndpointToBeCreated("track-123", plClient, context.Background(), nil)

	require.NoError(t, err)
	require.Empty(t, statuses)
}

func TestWaitForEndpointToBeCreated_Timeout(t *testing.T) {
	origDelay, origInterval, origTimeout := pollingDelay, pollingInterval, pollingTimeout
	pollingDelay = 0
	pollingInterval = 50 * time.Millisecond
	pollingTimeout = 200 * time.Millisecond
	t.Cleanup(func() {
		pollingDelay = origDelay
		pollingInterval = origInterval
		pollingTimeout = origTimeout
	})

	mockOps := mocks.NewMockCloudPrivateLinksClientService(t)
	mockOps.EXPECT().
		ListPrivateLinkEndpointStatusesContext(mock.Anything, mock.Anything).
		Return(makeEndpointStatusesOK([]*cloudprivatelinkmodels.PrivateLinkEndpointStatus{
			{ServiceComponent: "API", Status: "IN_PROGRESS"},
		}), nil).Maybe()

	plClient := &cloudprivatelinksclient.Cloudprivatelinks{Operations: mockOps}

	_, err := waitForEndpointToBeCreated("track-123", plClient, context.Background(), nil)

	require.Error(t, err)
	require.Contains(t, err.Error(), "timeout")
}

func TestWaitForEndpointToBeCreated_APIErrorExceedsThreshold(t *testing.T) {
	overridePollingTimings(t)

	apiErr := errors.New("service unavailable")

	mockOps := mocks.NewMockCloudPrivateLinksClientService(t)
	// threshold=3: first 3 calls return IN_PROGRESS, 4th exceeds and returns error
	mockOps.EXPECT().
		ListPrivateLinkEndpointStatusesContext(mock.Anything, mock.Anything).
		Return(nil, apiErr).Times(callFailureThreshold + 1)

	plClient := &cloudprivatelinksclient.Cloudprivatelinks{Operations: mockOps}

	_, err := waitForEndpointToBeCreated("track-123", plClient, context.Background(), nil)

	require.Error(t, err)
	require.Contains(t, err.Error(), "service unavailable")
}

func TestWaitForEndpointToBeCreated_APIErrorBelowThreshold(t *testing.T) {
	overridePollingTimings(t)

	apiErr := errors.New("transient error")

	mockOps := mocks.NewMockCloudPrivateLinksClientService(t)
	// 3 failures reach but do not exceed threshold (callFailedCount <= failureThreshold)
	mockOps.EXPECT().
		ListPrivateLinkEndpointStatusesContext(mock.Anything, mock.Anything).
		Return(nil, apiErr).Times(callFailureThreshold)
	mockOps.EXPECT().
		ListPrivateLinkEndpointStatusesContext(mock.Anything, mock.Anything).
		Return(makeEndpointStatusesOK([]*cloudprivatelinkmodels.PrivateLinkEndpointStatus{
			{ServiceComponent: "API", Status: "SUCCESS"},
		}), nil).Once()

	plClient := &cloudprivatelinksclient.Cloudprivatelinks{Operations: mockOps}

	statuses, err := waitForEndpointToBeCreated("track-123", plClient, context.Background(), nil)

	require.NoError(t, err)
	require.Len(t, statuses, 1)
}

// ── waitForEndpointToBeDeleted ─────────────────────────────────────────────

func TestWaitForEndpointToBeDeleted_AllSuccess(t *testing.T) {
	overridePollingTimings(t)

	mockOps := mocks.NewMockCloudPrivateLinksClientService(t)
	mockOps.EXPECT().
		ListPrivateLinkEndpointStatusesContext(mock.Anything, mock.Anything).
		Return(makeEndpointStatusesOK([]*cloudprivatelinkmodels.PrivateLinkEndpointStatus{
			{ServiceComponent: "API", Status: "SUCCESS"},
		}), nil)

	plClient := &cloudprivatelinksclient.Cloudprivatelinks{Operations: mockOps}

	err := waitForEndpointToBeDeleted("track-123", plClient, context.Background(), nil)

	require.NoError(t, err)
}

func TestWaitForEndpointToBeDeleted_ErrorStatus(t *testing.T) {
	overridePollingTimings(t)

	mockOps := mocks.NewMockCloudPrivateLinksClientService(t)
	mockOps.EXPECT().
		ListPrivateLinkEndpointStatusesContext(mock.Anything, mock.Anything).
		Return(makeEndpointStatusesOK([]*cloudprivatelinkmodels.PrivateLinkEndpointStatus{
			{ServiceComponent: "API", Status: "ERROR", Error: "deletion failed"},
		}), nil)

	plClient := &cloudprivatelinksclient.Cloudprivatelinks{Operations: mockOps}

	err := waitForEndpointToBeDeleted("track-123", plClient, context.Background(), nil)

	require.Error(t, err)
	require.Contains(t, err.Error(), "API")
	require.Contains(t, err.Error(), "deletion failed")
}

func TestWaitForEndpointToBeDeleted_Timeout(t *testing.T) {
	origDelay, origInterval, origTimeout := pollingDelay, pollingInterval, pollingTimeout
	pollingDelay = 0
	pollingInterval = 50 * time.Millisecond
	pollingTimeout = 200 * time.Millisecond
	t.Cleanup(func() {
		pollingDelay = origDelay
		pollingInterval = origInterval
		pollingTimeout = origTimeout
	})

	mockOps := mocks.NewMockCloudPrivateLinksClientService(t)
	mockOps.EXPECT().
		ListPrivateLinkEndpointStatusesContext(mock.Anything, mock.Anything).
		Return(makeEndpointStatusesOK([]*cloudprivatelinkmodels.PrivateLinkEndpointStatus{
			{ServiceComponent: "API", Status: "IN_PROGRESS"},
		}), nil).Maybe()

	plClient := &cloudprivatelinksclient.Cloudprivatelinks{Operations: mockOps}

	err := waitForEndpointToBeDeleted("track-123", plClient, context.Background(), nil)

	require.Error(t, err)
	require.Contains(t, err.Error(), "timeout")
}

func TestWaitForEndpointToBeDeleted_APIErrorExceedsThreshold(t *testing.T) {
	overridePollingTimings(t)

	apiErr := errors.New("service unavailable")

	mockOps := mocks.NewMockCloudPrivateLinksClientService(t)
	mockOps.EXPECT().
		ListPrivateLinkEndpointStatusesContext(mock.Anything, mock.Anything).
		Return(nil, apiErr).Times(callFailureThreshold + 1)

	plClient := &cloudprivatelinksclient.Cloudprivatelinks{Operations: mockOps}

	err := waitForEndpointToBeDeleted("track-123", plClient, context.Background(), nil)

	require.Error(t, err)
	require.Contains(t, err.Error(), "service unavailable")
}

// ── polling defaults ───────────────────────────────────────────────────────

func TestPollingDefaultValues(t *testing.T) {
	require.Equal(t, 10*time.Second, pollingDelay)
	require.Equal(t, 30*time.Minute, pollingTimeout)
	require.Equal(t, 15*time.Second, pollingInterval)
	require.Equal(t, 3, callFailureThreshold)
}
