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
	"testing"
	"time"

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
				DescribeEnvironment(mock.Anything).
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
				DescribeEnvironment(mock.Anything).
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
