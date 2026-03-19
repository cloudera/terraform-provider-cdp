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
	"testing"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

func TestFromListEnvironmentsResponse(t *testing.T) {
	environmentName := "test-env"
	crn := "crn:cdp:environments:us-west-1:tenant:environment:1234"
	status := "AVAILABLE"
	region := "us-west-1"
	cloudPlatform := "AWS"
	credentialName := "test-credential"

	tests := []struct {
		name     string
		input    *models.ListEnvironmentsResponse
		validate func(t *testing.T, result *EnvironmentListModel)
	}{
		{
			name:  "nil input",
			input: nil,
			validate: func(t *testing.T, result *EnvironmentListModel) {
				t.Helper()

				if result != nil {
					t.Fatalf("expected nil result, got %#v", result)
				}
			},
		},
		{
			name: "empty input",
			input: &models.ListEnvironmentsResponse{
				Environments: []*models.EnvironmentSummary{},
			},
			validate: func(t *testing.T, result *EnvironmentListModel) {
				t.Helper()

				if result == nil {
					t.Fatal("expected non-nil result")
					return
				}

				if result.Environments == nil {
					t.Fatal("expected non-nil environments slice")
					return
				}

				if len(result.Environments) != 0 {
					t.Fatalf("expected 0 environments, got %d", len(result.Environments))
				}
			},
		},
		{
			name: "populated input",
			input: &models.ListEnvironmentsResponse{
				Environments: []*models.EnvironmentSummary{
					{
						EnvironmentName: &environmentName,
						Crn:             &crn,
						Status:          &status,
						Region:          &region,
						CloudPlatform:   &cloudPlatform,
						CredentialName:  &credentialName,
					},
				},
			},
			validate: func(t *testing.T, result *EnvironmentListModel) {
				t.Helper()

				if result == nil {
					t.Fatal("expected non-nil result")
					return
				}

				if len(result.Environments) != 1 {
					t.Fatalf("expected 1 environment, got %d", len(result.Environments))
					return
				}

				env := result.Environments[0]

				if env.Name.ValueString() != environmentName {
					t.Errorf("expected Name to be %q, got %q", environmentName, env.Name.ValueString())
				}

				if env.Crn.ValueString() != crn {
					t.Errorf("expected Crn to be %q, got %q", crn, env.Crn.ValueString())
				}

				if env.Status.ValueString() != status {
					t.Errorf("expected Status to be %q, got %q", status, env.Status.ValueString())
				}

				if env.Region.ValueString() != region {
					t.Errorf("expected Region to be %q, got %q", region, env.Region.ValueString())
				}

				if env.CloudPlatform.ValueString() != cloudPlatform {
					t.Errorf("expected CloudPlatform to be %q, got %q", cloudPlatform, env.CloudPlatform.ValueString())
				}

				if env.CredentialName.ValueString() != credentialName {
					t.Errorf("expected CredentialName to be %q, got %q", credentialName, env.CredentialName.ValueString())
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fromListEnvironmentsResponse(tt.input)
			tt.validate(t, result)
		})
	}
}
