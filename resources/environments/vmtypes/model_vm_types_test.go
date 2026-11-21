// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package vmtypes

import (
	"testing"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/resources/environments/platformtypes"
)

func TestFromListResponse(t *testing.T) {
	tests := []struct {
		name     string
		input    *models.ListVMTypesResponse
		validate func(t *testing.T, result []platformtypes.PlatformVMTypeModel)
	}{
		{
			name:  "nil input",
			input: nil,
			validate: func(t *testing.T, result []platformtypes.PlatformVMTypeModel) {
				t.Helper()
				if result != nil {
					t.Fatalf("expected nil result, got %#v", result)
				}
			},
		},
		{
			name: "empty input",
			input: &models.ListVMTypesResponse{
				VMTypes: []*models.PlatformVMTypes{},
			},
			validate: func(t *testing.T, result []platformtypes.PlatformVMTypeModel) {
				t.Helper()
				if len(result) != 0 {
					t.Fatalf("expected 0 vm types, got %d", len(result))
				}
			},
		},
		{
			name: "populated input",
			input: &models.ListVMTypesResponse{
				VMTypes: []*models.PlatformVMTypes{
					{
						Platform: "AWS",
						DefaultVirtualMachine: &models.VMType{
							Name: "m5.xlarge",
						},
						VirtualMachines: []*models.VMType{
							{
								Name: "m5.xlarge",
								VMTypeMetadata: &models.VMTypeMetadata{
									Configs: []*models.VolumeParameterConfig{
										{
											MaximumNumber:       24,
											MaximumSize:         16384,
											MinimumNumber:       1,
											MinimumSize:         10,
											VolumeParameterType: "SSD",
										},
									},
									Properties: map[string]string{"Cpu": "4"},
								},
							},
							{Name: "m5.2xlarge"},
						},
					},
				},
			},
			validate: func(t *testing.T, result []platformtypes.PlatformVMTypeModel) {
				t.Helper()
				if len(result) != 1 {
					t.Fatalf("expected 1 platform vm type, got %d", len(result))
				}

				pvmt := result[0]
				if pvmt.Platform.ValueString() != "AWS" {
					t.Errorf("expected platform AWS, got %s", pvmt.Platform.ValueString())
				}
				if pvmt.DefaultVirtualMachine == nil {
					t.Fatal("expected non-nil default virtual machine")
				}
				if pvmt.DefaultVirtualMachine.Name.ValueString() != "m5.xlarge" {
					t.Errorf("expected default VM name m5.xlarge, got %s", pvmt.DefaultVirtualMachine.Name.ValueString())
				}
				if len(pvmt.VirtualMachines) != 2 {
					t.Fatalf("expected 2 virtual machines, got %d", len(pvmt.VirtualMachines))
				}
				if pvmt.VirtualMachines[1].VMTypeMetadata != nil {
					t.Error("expected nil metadata for second VM")
				}
			},
		},
		{
			name: "nil default virtual machine",
			input: &models.ListVMTypesResponse{
				VMTypes: []*models.PlatformVMTypes{
					{
						Platform:              "GCP",
						DefaultVirtualMachine: nil,
						VirtualMachines:       []*models.VMType{},
					},
				},
			},
			validate: func(t *testing.T, result []platformtypes.PlatformVMTypeModel) {
				t.Helper()
				if result[0].DefaultVirtualMachine != nil {
					t.Error("expected nil default virtual machine")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fromListResponse(tt.input)
			tt.validate(t, result)
		})
	}
}
