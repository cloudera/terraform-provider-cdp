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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

type PlatformVMTypeModel struct {
	Platform              types.String  `tfsdk:"platform"`
	DefaultVirtualMachine *VMTypeModel  `tfsdk:"default_virtual_machine"`
	VirtualMachines       []VMTypeModel `tfsdk:"virtual_machines"`
}

type VMTypeModel struct {
	Name           types.String         `tfsdk:"name"`
	VMTypeMetadata *VMTypeMetadataModel `tfsdk:"vm_type_metadata"`
}

type VMTypeMetadataModel struct {
	Configs    []VolumeParameterConfigModel `tfsdk:"configs"`
	Properties types.Map                    `tfsdk:"properties"`
}

type VolumeParameterConfigModel struct {
	MaximumNumber       types.Int32  `tfsdk:"maximum_number"`
	MaximumSize         types.Int32  `tfsdk:"maximum_size"`
	MinimumNumber       types.Int32  `tfsdk:"minimum_number"`
	MinimumSize         types.Int32  `tfsdk:"minimum_size"`
	VolumeParameterType types.String `tfsdk:"volume_parameter_type"`
}

func PlatformVMTypesFromSlice(in []*models.PlatformVMTypes) []PlatformVMTypeModel {
	if in == nil {
		return nil
	}
	result := make([]PlatformVMTypeModel, 0, len(in))
	for _, pvmt := range in {
		if pvmt == nil {
			continue
		}
		result = append(result, fromPlatformVMTypes(pvmt))
	}
	return result
}

func fromPlatformVMTypes(pvmt *models.PlatformVMTypes) PlatformVMTypeModel {
	m := PlatformVMTypeModel{
		Platform: types.StringValue(pvmt.Platform),
	}

	if pvmt.DefaultVirtualMachine != nil {
		dvm := fromVMType(pvmt.DefaultVirtualMachine)
		m.DefaultVirtualMachine = &dvm
	}

	m.VirtualMachines = make([]VMTypeModel, 0, len(pvmt.VirtualMachines))
	for _, vm := range pvmt.VirtualMachines {
		if vm == nil {
			continue
		}
		m.VirtualMachines = append(m.VirtualMachines, fromVMType(vm))
	}

	return m
}

func fromVMType(vm *models.VMType) VMTypeModel {
	m := VMTypeModel{
		Name: types.StringValue(vm.Name),
	}

	if vm.VMTypeMetadata != nil {
		m.VMTypeMetadata = fromVMTypeMetadata(vm.VMTypeMetadata)
	}

	return m
}

func fromVMTypeMetadata(meta *models.VMTypeMetadata) *VMTypeMetadataModel {
	m := &VMTypeMetadataModel{}

	m.Configs = make([]VolumeParameterConfigModel, len(meta.Configs))
	for i, c := range meta.Configs {
		m.Configs[i] = VolumeParameterConfigModel{
			MaximumNumber:       types.Int32Value(c.MaximumNumber),
			MaximumSize:         types.Int32Value(c.MaximumSize),
			MinimumNumber:       types.Int32Value(c.MinimumNumber),
			MinimumSize:         types.Int32Value(c.MinimumSize),
			VolumeParameterType: types.StringValue(c.VolumeParameterType),
		}
	}

	m.Properties = StringMapToTFMap(meta.Properties)

	return m
}

func stringMapToAttrValues(in map[string]string) map[string]attr.Value {
	out := make(map[string]attr.Value, len(in))
	for k, v := range in {
		out[k] = types.StringValue(v)
	}
	return out
}
