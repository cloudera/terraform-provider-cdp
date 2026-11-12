// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package freeipaverticalscaling

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

func TestFromListResponse(t *testing.T) {
	tests := []struct {
		name     string
		input    *models.ListFreeipaVMTypesForVerticalScalingResponse
		validate func(t *testing.T, result []platformVMTypeModel)
	}{
		{
			name:  "nil input",
			input: nil,
			validate: func(t *testing.T, result []platformVMTypeModel) {
				t.Helper()
				if result != nil {
					t.Fatalf("expected nil result, got %#v", result)
				}
			},
		},
		{
			name: "empty input",
			input: &models.ListFreeipaVMTypesForVerticalScalingResponse{
				VMTypes: []*models.PlatformVMTypes{},
			},
			validate: func(t *testing.T, result []platformVMTypeModel) {
				t.Helper()
				if len(result) != 0 {
					t.Fatalf("expected 0 vm types, got %d", len(result))
				}
			},
		},
		{
			name: "populated input",
			input: &models.ListFreeipaVMTypesForVerticalScalingResponse{
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
									Properties: map[string]string{
										"Cpu":    "4",
										"Memory": "16",
									},
								},
							},
							{
								Name: "m5.2xlarge",
							},
						},
					},
				},
			},
			validate: func(t *testing.T, result []platformVMTypeModel) {
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

				vm0 := pvmt.VirtualMachines[0]
				if vm0.Name.ValueString() != "m5.xlarge" {
					t.Errorf("expected VM name m5.xlarge, got %s", vm0.Name.ValueString())
				}
				if vm0.VMTypeMetadata == nil {
					t.Fatal("expected non-nil metadata for first VM")
				}
				if len(vm0.VMTypeMetadata.Configs) != 1 {
					t.Fatalf("expected 1 config, got %d", len(vm0.VMTypeMetadata.Configs))
				}
				cfg := vm0.VMTypeMetadata.Configs[0]
				if cfg.MaximumNumber.ValueInt32() != 24 {
					t.Errorf("expected max number 24, got %d", cfg.MaximumNumber.ValueInt32())
				}
				if cfg.MaximumSize.ValueInt32() != 16384 {
					t.Errorf("expected max size 16384, got %d", cfg.MaximumSize.ValueInt32())
				}
				if cfg.MinimumNumber.ValueInt32() != 1 {
					t.Errorf("expected min number 1, got %d", cfg.MinimumNumber.ValueInt32())
				}
				if cfg.MinimumSize.ValueInt32() != 10 {
					t.Errorf("expected min size 10, got %d", cfg.MinimumSize.ValueInt32())
				}
				if cfg.VolumeParameterType.ValueString() != "SSD" {
					t.Errorf("expected volume type SSD, got %s", cfg.VolumeParameterType.ValueString())
				}

				propMap := vm0.VMTypeMetadata.Properties
				if propMap.IsNull() {
					t.Fatal("expected non-null properties map")
				}
				elements := propMap.Elements()
				cpuVal, ok := elements["Cpu"]
				if !ok {
					t.Fatal("expected Cpu property")
				}
				if cpuVal.(types.String).ValueString() != "4" {
					t.Errorf("expected Cpu=4, got %s", cpuVal.(types.String).ValueString())
				}

				vm1 := pvmt.VirtualMachines[1]
				if vm1.VMTypeMetadata != nil {
					t.Error("expected nil metadata for second VM")
				}
			},
		},
		{
			name: "empty properties map",
			input: &models.ListFreeipaVMTypesForVerticalScalingResponse{
				VMTypes: []*models.PlatformVMTypes{
					{
						Platform: "AZURE",
						VirtualMachines: []*models.VMType{
							{
								Name: "Standard_D4s_v3",
								VMTypeMetadata: &models.VMTypeMetadata{
									Configs:    []*models.VolumeParameterConfig{},
									Properties: map[string]string{},
								},
							},
						},
					},
				},
			},
			validate: func(t *testing.T, result []platformVMTypeModel) {
				t.Helper()
				meta := result[0].VirtualMachines[0].VMTypeMetadata
				if meta == nil {
					t.Fatal("expected non-nil metadata")
				}
				if len(meta.Configs) != 0 {
					t.Errorf("expected 0 configs, got %d", len(meta.Configs))
				}
				if meta.Properties.IsNull() {
					t.Error("expected non-null (empty) properties map, got null")
				}
				if len(meta.Properties.Elements()) != 0 {
					t.Errorf("expected 0 properties, got %d", len(meta.Properties.Elements()))
				}
			},
		},
		{
			name: "nil default virtual machine",
			input: &models.ListFreeipaVMTypesForVerticalScalingResponse{
				VMTypes: []*models.PlatformVMTypes{
					{
						Platform:              "GCP",
						DefaultVirtualMachine: nil,
						VirtualMachines:       []*models.VMType{},
					},
				},
			},
			validate: func(t *testing.T, result []platformVMTypeModel) {
				t.Helper()
				if len(result) != 1 {
					t.Fatalf("expected 1 platform vm type, got %d", len(result))
				}
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
