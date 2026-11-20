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
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/resources/environments/platformtypes"
)

type vmTypesModel struct {
	CredentialName   types.String                        `tfsdk:"credential_name"`
	CredentialCrn    types.String                        `tfsdk:"credential_crn"`
	Region           types.String                        `tfsdk:"region"`
	AvailabilityZone types.String                        `tfsdk:"availability_zone"`
	Architecture     types.String                        `tfsdk:"architecture"`
	PlatformVariant  types.String                        `tfsdk:"platform_variant"`
	ResourceType     types.String                        `tfsdk:"resource_type"`
	VMTypes          []platformtypes.PlatformVMTypeModel `tfsdk:"vm_types"`
}

func fromListResponse(response *models.ListVMTypesResponse) []platformtypes.PlatformVMTypeModel {
	if response == nil {
		return nil
	}
	return platformtypes.PlatformVMTypesFromSlice(response.VMTypes)
}
