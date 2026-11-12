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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

type freeipaVMTypesModel struct {
	EnvironmentCrn    types.String          `tfsdk:"environment_crn"`
	Architecture      types.String          `tfsdk:"architecture"`
	AvailabilityZones types.List            `tfsdk:"availability_zones"`
	InstanceType      types.String          `tfsdk:"instance_type"`
	ResourceType      types.String          `tfsdk:"resource_type"`
	VMTypes           []platformVMTypeModel `tfsdk:"vm_types"`
}

type platformVMTypeModel struct {
	Platform              types.String  `tfsdk:"platform"`
	DefaultVirtualMachine *vmTypeModel  `tfsdk:"default_virtual_machine"`
	VirtualMachines       []vmTypeModel `tfsdk:"virtual_machines"`
}

type vmTypeModel struct {
	Name           types.String         `tfsdk:"name"`
	VMTypeMetadata *vmTypeMetadataModel `tfsdk:"vm_type_metadata"`
}

type vmTypeMetadataModel struct {
	Configs    []volumeParameterConfigModel `tfsdk:"configs"`
	Properties types.Map                    `tfsdk:"properties"`
}

type volumeParameterConfigModel struct {
	MaximumNumber       types.Int32  `tfsdk:"maximum_number"`
	MaximumSize         types.Int32  `tfsdk:"maximum_size"`
	MinimumNumber       types.Int32  `tfsdk:"minimum_number"`
	MinimumSize         types.Int32  `tfsdk:"minimum_size"`
	VolumeParameterType types.String `tfsdk:"volume_parameter_type"`
}

func fromListResponse(response *models.ListFreeipaVMTypesForVerticalScalingResponse) []platformVMTypeModel {
	if response == nil {
		return nil
	}

	result := make([]platformVMTypeModel, len(response.VMTypes))
	for i, pvmt := range response.VMTypes {
		result[i] = fromPlatformVMTypes(pvmt)
	}
	return result
}

func fromPlatformVMTypes(pvmt *models.PlatformVMTypes) platformVMTypeModel {
	m := platformVMTypeModel{
		Platform: types.StringValue(pvmt.Platform),
	}

	if pvmt.DefaultVirtualMachine != nil {
		dvm := fromVMType(pvmt.DefaultVirtualMachine)
		m.DefaultVirtualMachine = &dvm
	}

	m.VirtualMachines = make([]vmTypeModel, len(pvmt.VirtualMachines))
	for i, vm := range pvmt.VirtualMachines {
		m.VirtualMachines[i] = fromVMType(vm)
	}

	return m
}

func fromVMType(vm *models.VMType) vmTypeModel {
	m := vmTypeModel{
		Name: types.StringValue(vm.Name),
	}

	if vm.VMTypeMetadata != nil {
		m.VMTypeMetadata = fromVMTypeMetadata(vm.VMTypeMetadata)
	}

	return m
}

func fromVMTypeMetadata(meta *models.VMTypeMetadata) *vmTypeMetadataModel {
	m := &vmTypeMetadataModel{}

	m.Configs = make([]volumeParameterConfigModel, len(meta.Configs))
	for i, c := range meta.Configs {
		m.Configs[i] = volumeParameterConfigModel{
			MaximumNumber:       types.Int32Value(c.MaximumNumber),
			MaximumSize:         types.Int32Value(c.MaximumSize),
			MinimumNumber:       types.Int32Value(c.MinimumNumber),
			MinimumSize:         types.Int32Value(c.MinimumSize),
			VolumeParameterType: types.StringValue(c.VolumeParameterType),
		}
	}

	if meta.Properties != nil {
		m.Properties = types.MapValueMust(types.StringType, stringMapToAttrValues(meta.Properties))
	} else {
		m.Properties = types.MapNull(types.StringType)
	}

	return m
}

func stringMapToAttrValues(in map[string]string) map[string]attr.Value {
	out := make(map[string]attr.Value, len(in))
	for k, v := range in {
		out[k] = types.StringValue(v)
	}
	return out
}
