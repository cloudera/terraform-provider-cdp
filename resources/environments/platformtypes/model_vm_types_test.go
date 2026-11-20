// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package platformtypes

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

func TestPlatformVMTypesFromSliceNil(t *testing.T) {
	if r := PlatformVMTypesFromSlice(nil); r != nil {
		t.Fatalf("expected nil, got %#v", r)
	}
	if r := PlatformVMTypesFromSlice([]*models.PlatformVMTypes{}); len(r) != 0 {
		t.Fatalf("expected 0, got %d", len(r))
	}
}

func TestPlatformVMTypesFromSlice(t *testing.T) {
	input := []*models.PlatformVMTypes{
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
	}

	result := PlatformVMTypesFromSlice(input)
	if len(result) != 1 {
		t.Fatalf("expected 1, got %d", len(result))
	}

	pvmt := result[0]
	if pvmt.Platform.ValueString() != "AWS" {
		t.Errorf("expected AWS, got %s", pvmt.Platform.ValueString())
	}
	if pvmt.DefaultVirtualMachine == nil {
		t.Fatal("expected non-nil default VM")
	}
	if pvmt.DefaultVirtualMachine.Name.ValueString() != "m5.xlarge" {
		t.Errorf("expected m5.xlarge, got %s", pvmt.DefaultVirtualMachine.Name.ValueString())
	}
	if len(pvmt.VirtualMachines) != 2 {
		t.Fatalf("expected 2 VMs, got %d", len(pvmt.VirtualMachines))
	}

	vm0 := pvmt.VirtualMachines[0]
	if vm0.VMTypeMetadata == nil {
		t.Fatal("expected non-nil metadata")
	}
	cfg := vm0.VMTypeMetadata.Configs[0]
	if cfg.MaximumNumber.ValueInt32() != 24 {
		t.Errorf("expected 24, got %d", cfg.MaximumNumber.ValueInt32())
	}
	if cfg.VolumeParameterType.ValueString() != "SSD" {
		t.Errorf("expected SSD, got %s", cfg.VolumeParameterType.ValueString())
	}

	cpuVal, ok := vm0.VMTypeMetadata.Properties.Elements()["Cpu"]
	if !ok {
		t.Fatal("expected Cpu property")
	}
	if cpuVal.(types.String).ValueString() != "4" {
		t.Errorf("expected 4, got %s", cpuVal.(types.String).ValueString())
	}

	vm1 := pvmt.VirtualMachines[1]
	if vm1.VMTypeMetadata != nil {
		t.Error("expected nil metadata for second VM")
	}
}

func TestPlatformVMTypesNilDefaultVM(t *testing.T) {
	input := []*models.PlatformVMTypes{
		{
			Platform:              "GCP",
			DefaultVirtualMachine: nil,
			VirtualMachines:       []*models.VMType{},
		},
	}
	result := PlatformVMTypesFromSlice(input)
	if result[0].DefaultVirtualMachine != nil {
		t.Error("expected nil default VM")
	}
}

func TestPlatformVMTypesNilProperties(t *testing.T) {
	input := []*models.PlatformVMTypes{
		{
			Platform: "AZURE",
			VirtualMachines: []*models.VMType{
				{
					Name: "Standard_D4s_v3",
					VMTypeMetadata: &models.VMTypeMetadata{
						Properties: nil,
					},
				},
			},
		},
	}
	result := PlatformVMTypesFromSlice(input)
	meta := result[0].VirtualMachines[0].VMTypeMetadata
	if !meta.Properties.IsNull() {
		t.Error("expected null properties for nil input")
	}
}
