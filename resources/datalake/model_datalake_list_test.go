// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datalake

import (
	"testing"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
)

func TestFromListDatalakesResponse(t *testing.T) {
	datalakeName := "test-datalake"
	crn := "crn:cdp:datalake:us-west-1:tenant:datalake:1234"
	status := "RUNNING"
	environmentCrn := "crn:cdp:environments:us-west-1:tenant:environment:5678"

	tests := []struct {
		name     string
		input    *models.ListDatalakesResponse
		validate func(t *testing.T, result *DatalakeListModel)
	}{
		{
			name:  "nil input",
			input: nil,
			validate: func(t *testing.T, result *DatalakeListModel) {
				t.Helper()

				if result != nil {
					t.Fatalf("expected nil result, got %#v", result)
				}
			},
		},
		{
			name: "empty input",
			input: &models.ListDatalakesResponse{
				Datalakes: []*models.Datalake{},
			},
			validate: func(t *testing.T, result *DatalakeListModel) {
				t.Helper()

				if result == nil {
					t.Fatal("expected non-nil result")
					return
				}

				if result.Datalakes == nil {
					t.Fatal("expected non-nil datalakes slice")
					return
				}

				if len(result.Datalakes) != 0 {
					t.Fatalf("expected 0 datalakes, got %d", len(result.Datalakes))
					return
				}
			},
		},
		{
			name: "populated input",
			input: &models.ListDatalakesResponse{
				Datalakes: []*models.Datalake{
					{
						DatalakeName:   &datalakeName,
						Crn:            &crn,
						Status:         status,
						EnvironmentCrn: environmentCrn,
					},
				},
			},
			validate: func(t *testing.T, result *DatalakeListModel) {
				t.Helper()

				if result == nil {
					t.Fatal("expected non-nil result")
					return
				}

				if len(result.Datalakes) != 1 {
					t.Fatalf("expected 1 datalake, got %d", len(result.Datalakes))
					return
				}

				dl := result.Datalakes[0]

				if dl.Name.ValueString() != datalakeName {
					t.Errorf("expected Name to be %q, got %q", datalakeName, dl.Name.ValueString())
				}

				if dl.Crn.ValueString() != crn {
					t.Errorf("expected Crn to be %q, got %q", crn, dl.Crn.ValueString())
				}

				if dl.Status.ValueString() != status {
					t.Errorf("expected Status to be %q, got %q", status, dl.Status.ValueString())
				}

				if dl.EnvironmentCrn.ValueString() != environmentCrn {
					t.Errorf("expected EnvironmentCrn to be %q, got %q", environmentCrn, dl.EnvironmentCrn.ValueString())
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fromListDatalakesResponse(tt.input)
			tt.validate(t, result)
		})
	}
}
