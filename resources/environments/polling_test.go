// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package environments

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	envclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/mocks"
)

func TestWaitForEnvironmentToBeAvailable_PendingStatuses(t *testing.T) {
	t.Parallel()

	pendingStatuses := []string{
		"CREATION_INITIATED",
		"NETWORK_CREATION_IN_PROGRESS",
		"PUBLICKEY_CREATE_IN_PROGRESS",
		"ENVIRONMENT_RESOURCE_ENCRYPTION_INITIALIZATION_IN_PROGRESS",
		"ENVIRONMENT_VALIDATION_IN_PROGRESS",
		"ENVIRONMENT_INITIALIZATION_IN_PROGRESS",
		"COMPUTE_CLUSTER_CREATION_IN_PROGRESS",
		"FREEIPA_CREATION_IN_PROGRESS",
	}

	for _, status := range pendingStatuses {
		status := status

		t.Run(status, func(t *testing.T) {
			t.Parallel()

			mockOps := mocks.NewMockEnvironmentClientService(t)
			mockOps.EXPECT().
				DescribeEnvironmentContext(mock.Anything, mock.Anything).
				Return(describeEnvironmentOK(status), nil).
				Maybe()

			client := &envclient.Environments{
				Operations: mockOps,
			}

			err := waitForEnvironmentToBeAvailable(
				"test-env",
				6*time.Second,
				1,
				client,
				context.Background(),
				nil,
				func(*environmentsmodels.Environment) {},
			)

			require.Error(t, err)
			require.Contains(t, err.Error(), "timeout")
		})
	}
}

func TestWaitForEnvironmentToBeAvailable_TargetStatuses(t *testing.T) {
	t.Parallel()

	targetStatuses := []string{
		"AVAILABLE",
		"TRUST_SETUP_REQUIRED",
	}

	for _, status := range targetStatuses {
		status := status

		t.Run(status, func(t *testing.T) {
			t.Parallel()

			mockOps := mocks.NewMockEnvironmentClientService(t)
			mockOps.EXPECT().
				DescribeEnvironmentContext(mock.Anything, mock.Anything).
				Return(describeEnvironmentOK(status), nil)

			client := &envclient.Environments{
				Operations: mockOps,
			}

			var saved *environmentsmodels.Environment

			err := waitForEnvironmentToBeAvailable(
				"test-env",
				6*time.Second,
				1,
				client,
				context.Background(),
				nil,
				func(env *environmentsmodels.Environment) {
					saved = env
				},
			)

			require.NoError(t, err)
			require.NotNil(t, saved)
			require.Equal(t, status, *saved.Status)
		})
	}
}

func TestWaitForEnvironmentToBeAvailable_PendingStatusesAreExpected(t *testing.T) {
	t.Parallel()

	expected := []string{
		"CREATION_INITIATED",
		"NETWORK_CREATION_IN_PROGRESS",
		"PUBLICKEY_CREATE_IN_PROGRESS",
		"ENVIRONMENT_RESOURCE_ENCRYPTION_INITIALIZATION_IN_PROGRESS",
		"ENVIRONMENT_VALIDATION_IN_PROGRESS",
		"ENVIRONMENT_INITIALIZATION_IN_PROGRESS",
		"COMPUTE_CLUSTER_CREATION_IN_PROGRESS",
		"FREEIPA_CREATION_IN_PROGRESS",
		"USER_DEFINED_TAGS_MODIFICATION_IN_PROGRESS",
		"USER_DEFINED_TAGS_MODIFICATION_ON_FREEIPA_IN_PROGRESS",
		"USER_DEFINED_TAGS_MODIFICATION_ON_DATALAKE_IN_PROGRESS",
		"USER_DEFINED_TAGS_MODIFICATION_ON_DATAHUBS_IN_PROGRESS",
		"USER_DEFINED_TAGS_MODIFICATION_ON_DATA_SERVICES_IN_PROGRESS",
	}

	require.Equal(t, expected, envCreatePendingStatuses)
}

func TestWaitForEnvironmentToBeAvailable_TargetStatusesAreExpected(t *testing.T) {
	t.Parallel()

	expected := []string{
		"AVAILABLE",
		"TRUST_SETUP_REQUIRED",
	}

	require.Equal(t, expected, envCreateTargetStatuses)
}

func describeEnvironmentOK(status string) *operations.DescribeEnvironmentOK {
	return &operations.DescribeEnvironmentOK{
		Payload: &environmentsmodels.DescribeEnvironmentResponse{
			Environment: &environmentsmodels.Environment{
				Status: &status,
			},
		},
	}
}

// Tests for checkResponseStatusForError

func TestCheckResponseStatusForError_AvailableStatus_ReturnsStatus(t *testing.T) {
	status := "AVAILABLE"
	resp := describeEnvironmentOK(status)

	result, statusStr, err := checkResponseStatusForError(resp)

	assert.NoError(t, err)
	assert.Equal(t, status, statusStr)
	assert.Equal(t, resp, result)
}

func TestCheckResponseStatusForError_FailedStatus_ReturnsError(t *testing.T) {
	resp := &operations.DescribeEnvironmentOK{
		Payload: &environmentsmodels.DescribeEnvironmentResponse{
			Environment: &environmentsmodels.Environment{
				Status:       new("CREATE_FAILED"),
				StatusReason: "something went wrong",
			},
		},
	}

	result, _, err := checkResponseStatusForError(resp)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "unexpected Environment status")
	assert.Contains(t, err.Error(), "CREATE_FAILED")
	assert.Contains(t, err.Error(), "something went wrong")
}

func TestCheckResponseStatusForError_ErrorStatus_ReturnsError(t *testing.T) {
	resp := &operations.DescribeEnvironmentOK{
		Payload: &environmentsmodels.DescribeEnvironmentResponse{
			Environment: &environmentsmodels.Environment{
				Status:       new("ERROR"),
				StatusReason: "internal error",
			},
		},
	}

	result, _, err := checkResponseStatusForError(resp)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "ERROR")
}

func TestCheckResponseStatusForError_PendingStatus_ReturnsStatusWithoutError(t *testing.T) {
	status := "CREATION_INITIATED"
	resp := describeEnvironmentOK(status)

	result, statusStr, err := checkResponseStatusForError(resp)

	assert.NoError(t, err)
	assert.Equal(t, "CREATION_INITIATED", statusStr)
	assert.NotNil(t, result)
}

// Tests for waitForEnvironmentToBeAvailable - error scenarios

func TestWaitForEnvironmentToBeAvailable_FailedStatus_ReturnsError(t *testing.T) {
	t.Parallel()

	mockOps := mocks.NewMockEnvironmentClientService(t)
	mockOps.EXPECT().
		DescribeEnvironmentContext(mock.Anything, mock.Anything).
		Return(&operations.DescribeEnvironmentOK{
			Payload: &environmentsmodels.DescribeEnvironmentResponse{
				Environment: &environmentsmodels.Environment{
					Status:       new("CREATE_FAILED"),
					StatusReason: "infra failure",
				},
			},
		}, nil)

	client := &envclient.Environments{Operations: mockOps}

	err := waitForEnvironmentToBeAvailable(
		"test-env",
		6*time.Second,
		3,
		client,
		context.Background(),
		nil,
		func(*environmentsmodels.Environment) {},
	)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "CREATE_FAILED")
}

func TestWaitForEnvironmentToBeAvailable_NotFoundRecovery_RetriesUntilAvailable(t *testing.T) {
	t.Parallel()

	mockOps := mocks.NewMockEnvironmentClientService(t)

	notFoundErr := &operations.DescribeEnvironmentDefault{Payload: &environmentsmodels.Error{Code: "NOT_FOUND", Message: "env not found"}}

	mockOps.EXPECT().
		DescribeEnvironmentContext(mock.Anything, mock.Anything).
		Return((*operations.DescribeEnvironmentOK)(nil), notFoundErr).
		Once()

	mockOps.EXPECT().
		DescribeEnvironmentContext(mock.Anything, mock.Anything).
		Return(describeEnvironmentOK("AVAILABLE"), nil).
		Once()

	client := &envclient.Environments{Operations: mockOps}

	err := waitForEnvironmentToBeAvailable(
		"test-env",
		30*time.Second,
		3,
		client,
		context.Background(),
		nil,
		func(*environmentsmodels.Environment) {},
	)

	require.NoError(t, err)
}

func TestWaitForEnvironmentToBeAvailable_CallFailureThresholdExceeded_ReturnsError(t *testing.T) {
	t.Parallel()

	mockOps := mocks.NewMockEnvironmentClientService(t)

	genericErr := errors.New("connection refused")

	mockOps.EXPECT().
		DescribeEnvironmentContext(mock.Anything, mock.Anything).
		Return((*operations.DescribeEnvironmentOK)(nil), genericErr).
		Maybe()

	client := &envclient.Environments{Operations: mockOps}

	err := waitForEnvironmentToBeAvailable(
		"test-env",
		30*time.Second,
		2,
		client,
		context.Background(),
		nil,
		func(*environmentsmodels.Environment) {},
	)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "connection refused")
}

func TestWaitForEnvironmentToBeAvailable_StateSaverCbCalledOnSuccess(t *testing.T) {
	t.Parallel()

	mockOps := mocks.NewMockEnvironmentClientService(t)
	mockOps.EXPECT().
		DescribeEnvironmentContext(mock.Anything, mock.Anything).
		Return(describeEnvironmentOK("AVAILABLE"), nil)

	client := &envclient.Environments{Operations: mockOps}

	callCount := 0
	err := waitForEnvironmentToBeAvailable(
		"test-env",
		6*time.Second,
		3,
		client,
		context.Background(),
		nil,
		func(env *environmentsmodels.Environment) {
			callCount++
			assert.Equal(t, "AVAILABLE", *env.Status)
		},
	)

	require.NoError(t, err)
	assert.Equal(t, 1, callCount)
}

// Tests for waitForEnvironmentToBeDeleted

func TestWaitForEnvironmentToBeDeleted_EnvNotFound_ReturnsSuccess(t *testing.T) {
	t.Parallel()

	mockOps := mocks.NewMockEnvironmentClientService(t)

	notFoundErr := &operations.DescribeEnvironmentDefault{Payload: &environmentsmodels.Error{Code: "NOT_FOUND", Message: "env not found"}}
	mockOps.EXPECT().
		DescribeEnvironmentContext(mock.Anything, mock.Anything).
		Return((*operations.DescribeEnvironmentOK)(nil), notFoundErr)

	client := &envclient.Environments{Operations: mockOps}

	err := waitForEnvironmentToBeDeleted(
		"test-env",
		6*time.Second,
		3,
		client,
		context.Background(),
		nil,
	)

	require.NoError(t, err)
}

func TestWaitForEnvironmentToBeDeleted_DeleteInProgress_WaitsUntilGone(t *testing.T) {
	t.Parallel()

	mockOps := mocks.NewMockEnvironmentClientService(t)

	mockOps.EXPECT().
		DescribeEnvironmentContext(mock.Anything, mock.Anything).
		Return(describeEnvironmentOK("FREEIPA_DELETE_IN_PROGRESS"), nil).
		Once()

	notFoundErr := &operations.DescribeEnvironmentDefault{Payload: &environmentsmodels.Error{Code: "NOT_FOUND", Message: "env not found"}}
	mockOps.EXPECT().
		DescribeEnvironmentContext(mock.Anything, mock.Anything).
		Return((*operations.DescribeEnvironmentOK)(nil), notFoundErr).
		Once()

	client := &envclient.Environments{Operations: mockOps}

	err := waitForEnvironmentToBeDeleted(
		"test-env",
		30*time.Second,
		3,
		client,
		context.Background(),
		nil,
	)

	require.NoError(t, err)
}

func TestWaitForEnvironmentToBeDeleted_FailedStatus_ReturnsError(t *testing.T) {
	t.Parallel()

	mockOps := mocks.NewMockEnvironmentClientService(t)
	mockOps.EXPECT().
		DescribeEnvironmentContext(mock.Anything, mock.Anything).
		Return(&operations.DescribeEnvironmentOK{
			Payload: &environmentsmodels.DescribeEnvironmentResponse{
				Environment: &environmentsmodels.Environment{
					Status:       new("DELETE_FAILED"),
					StatusReason: "delete failed reason",
				},
			},
		}, nil).
		Maybe()

	client := &envclient.Environments{Operations: mockOps}

	err := waitForEnvironmentToBeDeleted(
		"test-env",
		6*time.Second,
		3,
		client,
		context.Background(),
		nil,
	)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "DELETE_FAILED")
}

func TestWaitForEnvironmentToBeDeleted_NilEnvironment_ReturnsSuccess(t *testing.T) {
	t.Parallel()

	mockOps := mocks.NewMockEnvironmentClientService(t)
	mockOps.EXPECT().
		DescribeEnvironmentContext(mock.Anything, mock.Anything).
		Return(&operations.DescribeEnvironmentOK{
			Payload: &environmentsmodels.DescribeEnvironmentResponse{
				Environment: nil,
			},
		}, nil)

	client := &envclient.Environments{Operations: mockOps}

	err := waitForEnvironmentToBeDeleted(
		"test-env",
		6*time.Second,
		3,
		client,
		context.Background(),
		nil,
	)

	require.NoError(t, err)
}

func TestWaitForEnvironmentToBeDeleted_NonNotFoundError_BelowThreshold_TreatsAsGone(t *testing.T) {
	t.Parallel()

	mockOps := mocks.NewMockEnvironmentClientService(t)

	genericErr := errors.New("network timeout")
	mockOps.EXPECT().
		DescribeEnvironmentContext(mock.Anything, mock.Anything).
		Return((*operations.DescribeEnvironmentOK)(nil), genericErr)

	client := &envclient.Environments{Operations: mockOps}

	err := waitForEnvironmentToBeDeleted(
		"test-env",
		6*time.Second,
		3,
		client,
		context.Background(),
		nil,
	)

	require.NoError(t, err)
}

func TestWaitForEnvironmentToBeDeleted_DescribeEnvironmentDefault_NonNotFound_BelowThreshold_TreatsAsGone(t *testing.T) {
	t.Parallel()

	mockOps := mocks.NewMockEnvironmentClientService(t)

	serverErr := &operations.DescribeEnvironmentDefault{Payload: &environmentsmodels.Error{Code: "INTERNAL_ERROR", Message: "server exploded"}}
	mockOps.EXPECT().
		DescribeEnvironmentContext(mock.Anything, mock.Anything).
		Return((*operations.DescribeEnvironmentOK)(nil), serverErr)

	client := &envclient.Environments{Operations: mockOps}

	err := waitForEnvironmentToBeDeleted(
		"test-env",
		6*time.Second,
		3,
		client,
		context.Background(),
		nil,
	)

	require.NoError(t, err)
}

func TestWaitForEnvironmentToBeDeleted_InvalidThreshold_ReturnsErrorImmediately(t *testing.T) {
	t.Parallel()

	mockOps := mocks.NewMockEnvironmentClientService(t)
	client := &envclient.Environments{Operations: mockOps}

	err := waitForEnvironmentToBeDeleted(
		"test-env",
		6*time.Second,
		0,
		client,
		context.Background(),
		nil,
	)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "no meaningful threshold value")
}

func TestWaitForEnvironmentToBeAvailable_InvalidThreshold_ReturnsErrorImmediately(t *testing.T) {
	t.Parallel()

	mockOps := mocks.NewMockEnvironmentClientService(t)
	client := &envclient.Environments{Operations: mockOps}

	err := waitForEnvironmentToBeAvailable(
		"test-env",
		6*time.Second,
		0,
		client,
		context.Background(),
		nil,
		func(*environmentsmodels.Environment) {},
	)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "no meaningful threshold value")
}

func TestWaitForEnvironmentToBeAvailable_NonNotFoundError_ThresholdExceeded_ReturnsError(t *testing.T) {
	t.Parallel()

	mockOps := mocks.NewMockEnvironmentClientService(t)

	serverErr := &operations.DescribeEnvironmentDefault{Payload: &environmentsmodels.Error{Code: "INTERNAL_ERROR", Message: "server exploded"}}
	mockOps.EXPECT().
		DescribeEnvironmentContext(mock.Anything, mock.Anything).
		Return((*operations.DescribeEnvironmentOK)(nil), serverErr).
		Maybe()

	client := &envclient.Environments{Operations: mockOps}

	err := waitForEnvironmentToBeAvailable(
		"test-env",
		30*time.Second,
		1,
		client,
		context.Background(),
		nil,
		func(*environmentsmodels.Environment) {},
	)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "describeEnvironment default")
}

func TestWaitForEnvironmentToBeDeleted_PendingStatuses(t *testing.T) {
	t.Parallel()

	pendingStatuses := []string{
		"CLUSTER_DEFINITION_CLEANUP_PROGRESS",
		"CLUSTER_DEFINITION_DELETE_PROGRESS",
		"COMPUTE_CLUSTERS_DELETE_IN_PROGRESS",
		"DATAHUB_CLUSTERS_DELETE_IN_PROGRESS",
		"DATALAKE_CLUSTERS_DELETE_IN_PROGRESS",
		"DELETE_INITIATED",
		"ENVIRONMENT_ENCRYPTION_RESOURCES_DELETED",
		"ENVIRONMENT_RESOURCE_ENCRYPTION_DELETE_IN_PROGRESS",
		"EVENT_CLEANUP_IN_PROGRESS",
		"EXPERIENCE_DELETE_IN_PROGRESS",
		"FREEIPA_DELETE_IN_PROGRESS",
		"IDBROKER_MAPPINGS_DELETE_IN_PROGRESS",
		"NETWORK_DELETE_IN_PROGRESS",
		"PUBLICKEY_DELETE_IN_PROGRESS",
		"RDBMS_DELETE_IN_PROGRESS",
		"S3GUARD_TABLE_DELETE_IN_PROGRESS",
		"STORAGE_CONSUMPTION_COLLECTION_UNSCHEDULING_IN_PROGRESS",
		"UMS_RESOURCE_DELETE_IN_PROGRESS",
	}

	for _, status := range pendingStatuses {
		status := status
		t.Run(status, func(t *testing.T) {
			t.Parallel()

			mockOps := mocks.NewMockEnvironmentClientService(t)
			mockOps.EXPECT().
				DescribeEnvironmentContext(mock.Anything, mock.Anything).
				Return(describeEnvironmentOK(status), nil).
				Maybe()

			client := &envclient.Environments{Operations: mockOps}

			err := waitForEnvironmentToBeDeleted(
				"test-env",
				6*time.Second,
				3,
				client,
				context.Background(),
				nil,
			)

			require.Error(t, err)
			assert.Contains(t, err.Error(), "timeout")
		})
	}
}
