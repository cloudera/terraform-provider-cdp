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
	"testing"

	"github.com/stretchr/testify/assert"

	cloudprivatelinkmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/cloudprivatelinks/models"
)

func TestAllEndpointsSettled(t *testing.T) {
	tests := []struct {
		name     string
		statuses []*cloudprivatelinkmodels.PrivateLinkEndpointStatus
		expected bool
	}{
		{
			name:     "empty list",
			statuses: []*cloudprivatelinkmodels.PrivateLinkEndpointStatus{},
			expected: true,
		},
		{
			name: "all settled",
			statuses: []*cloudprivatelinkmodels.PrivateLinkEndpointStatus{
				{Status: "READY"},
				{Status: "ERROR"},
			},
			expected: true,
		},
		{
			name: "one in progress",
			statuses: []*cloudprivatelinkmodels.PrivateLinkEndpointStatus{
				{Status: "READY"},
				{Status: "IN_PROGRESS"},
			},
			expected: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, allEndpointsSettled(tc.statuses))
		})
	}
}

func TestHasErrors(t *testing.T) {
	tests := []struct {
		name     string
		statuses []*cloudprivatelinkmodels.PrivateLinkEndpointStatus
		expected bool
	}{
		{
			name:     "no errors",
			statuses: []*cloudprivatelinkmodels.PrivateLinkEndpointStatus{{Status: "READY"}},
			expected: false,
		},
		{
			name:     "has error",
			statuses: []*cloudprivatelinkmodels.PrivateLinkEndpointStatus{{Status: "ERROR"}},
			expected: true,
		},
		{
			name:     "empty list",
			statuses: []*cloudprivatelinkmodels.PrivateLinkEndpointStatus{},
			expected: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, hasErrors(tc.statuses))
		})
	}
}
