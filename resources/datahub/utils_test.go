// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datahub

import (
	"testing"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
)

func TestCheckIfClusterCreationFailed(t *testing.T) {
	type testCase struct {
		description        string
		input              operations.DescribeClusterOK
		shouldContainError bool
	}
	for _, scenario := range []testCase{
		{
			description:        "Test with status: CREATION_FAILED",
			input:              createOkInputWithStatus("CREATION_FAILED"),
			shouldContainError: true,
		},
		{
			description:        "Test with status: DELETED_ON_PROVIDER",
			input:              createOkInputWithStatus("DELETED_ON_PROVIDER"),
			shouldContainError: true,
		},
		{
			description:        "Test with status: UPDATE_IN_PROGRESS",
			input:              createOkInputWithStatus("UPDATE_IN_PROGRESS"),
			shouldContainError: false,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			_, _, err := checkIfClusterCreationFailed(&scenario.input)
			if scenario.shouldContainError && err == nil {
				t.Errorf("Test should contain error but the result does not contain it!")
			}
		})
	}
}

func TestIsNotFoundError(t *testing.T) {
	type testCase struct {
		description    string
		input          operations.DescribeClusterDefault
		expectedResult bool
	}
	for _, scenario := range []testCase{
		{
			description:    "Test with status: NOT_FOUND",
			input:          createDefaultInputWithStatus("NOT_FOUND", "Clustar cannot be found."),
			expectedResult: true,
		},
		{
			description:    "Test with status: DELETED_ON_PROVIDER",
			input:          createDefaultInputWithStatus("DELETED_ON_PROVIDER", "Cluster got deleted."),
			expectedResult: false,
		},
		{
			description:    "Test with status: BAD_REQUEST",
			input:          createDefaultInputWithStatus("BAD_REQUEST", "Not a valid request."),
			expectedResult: false,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			result := isNotFoundError(&scenario.input)
			if scenario.expectedResult != result {
				t.Errorf("Test result ('%t') does not match with the expectation ('%t')", result, scenario.expectedResult)
			}
		})
	}
}

func TestIsInternalServerError(t *testing.T) {
	type testCase struct {
		description    string
		input          operations.DescribeClusterDefault
		expectedResult bool
	}
	for _, scenario := range []testCase{
		{
			description:    "Test with status: NOT_FOUND",
			input:          createDefaultInputWithStatus("NOT_FOUND", "Clustar cannot be found."),
			expectedResult: false,
		},
		{
			description:    "Test with status: UNKNOWN",
			input:          createDefaultInputWithStatus("UNKNOWN", "Internal server error occurred."),
			expectedResult: true,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			result := isInternalServerError(&scenario.input)
			if scenario.expectedResult != result {
				t.Errorf("Test result ('%t') does not match with the expectation ('%t')", result, scenario.expectedResult)
			}
		})
	}
}

func createOkInputWithStatus(status string) operations.DescribeClusterOK {
	sum := &models.Cluster{Status: status}
	pl := &models.DescribeClusterResponse{Cluster: sum}
	return operations.DescribeClusterOK{Payload: pl}
}

func createDefaultInputWithStatus(code string, msg string) operations.DescribeClusterDefault {
	return operations.DescribeClusterDefault{
		Payload: &models.Error{
			Code:    code,
			Message: msg,
		},
	}
}
