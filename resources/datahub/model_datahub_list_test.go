// Copyright 2026 Cloudera. All Rights Reserved.
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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
)

func TestFromListClustersResponse(t *testing.T) {
	clusterName := "test-cluster"
	clusterCrn := "test-crn"
	tests := []struct {
		name     string
		input    *models.ListClustersResponse
		expected *DatahubListModel
	}{
		{
			name:     "nil response",
			input:    nil,
			expected: nil,
		},
		{
			name: "empty cluster list",
			input: &models.ListClustersResponse{
				Clusters: []*models.ClusterSummary{},
			},
			expected: &DatahubListModel{
				Datahubs: []Datahub{},
			},
		},
		{
			name: "single cluster",
			input: &models.ListClustersResponse{
				Clusters: []*models.ClusterSummary{
					{
						Crn:             &clusterCrn,
						Status:          "available",
						ClusterName:     &clusterName,
						DatalakeCrn:     "datalake-crn",
						CloudPlatform:   "aws",
						EnvironmentCrn:  "env-crn",
						EnvironmentName: "env-name",
					},
				},
			},
			expected: &DatahubListModel{
				Datahubs: []Datahub{
					{
						Crn:             types.StringPointerValue(&clusterCrn),
						Status:          types.StringValue("available"),
						Name:            types.StringPointerValue(&clusterName),
						DatalakeCrn:     types.StringValue("datalake-crn"),
						CloudPlatform:   types.StringValue("aws"),
						EnvironmentCrn:  types.StringValue("env-crn"),
						EnvironmentName: types.StringValue("env-name"),
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := fromListClustersResponse(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
