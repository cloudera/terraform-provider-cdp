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
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/resources/environments/platformtypes"
)

type freeipaVMTypesModel struct {
	EnvironmentCrn    types.String                        `tfsdk:"environment_crn"`
	Architecture      types.String                        `tfsdk:"architecture"`
	AvailabilityZones types.List                          `tfsdk:"availability_zones"`
	InstanceType      types.String                        `tfsdk:"instance_type"`
	ResourceType      types.String                        `tfsdk:"resource_type"`
	VMTypes           []platformtypes.PlatformVMTypeModel `tfsdk:"vm_types"`
}

func fromListResponse(response *models.ListFreeipaVMTypesForVerticalScalingResponse) []platformtypes.PlatformVMTypeModel {
	if response == nil {
		return nil
	}
	return platformtypes.PlatformVMTypesFromSlice(response.VMTypes)
}
