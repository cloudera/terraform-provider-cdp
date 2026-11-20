// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package disktypes

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/resources/environments/platformtypes"
)

type diskTypesModel struct {
	PlatformDiskTypes []platformDiskTypesEntryModel `tfsdk:"platform_disk_types"`
}

type platformDiskTypesEntryModel struct {
	Platform        types.String `tfsdk:"platform"`
	DefaultDiskType types.String `tfsdk:"default_disk_type"`
	DiskTypes       types.List   `tfsdk:"disk_types"`
	DisplayNames    types.Map    `tfsdk:"display_names"`
	DiskMappings    types.Map    `tfsdk:"disk_mappings"`
}

func diskTypesModelFromResponse(resp *models.ListDiskTypesResponse) diskTypesModel {
	if resp == nil {
		return diskTypesModel{}
	}
	result := make([]platformDiskTypesEntryModel, 0, len(resp.PlatformDiskTypes))
	for _, entry := range resp.PlatformDiskTypes {
		if entry == nil {
			continue
		}
		result = append(result, platformDiskTypesEntryModel{
			Platform:        platformtypes.StringValueOrNull(entry.Platform),
			DefaultDiskType: platformtypes.StringValueOrNull(entry.DefaultDiskType),
			DiskTypes:       platformtypes.StringSliceToList(entry.DiskTypes),
			DisplayNames:    platformtypes.StringMapToTFMap(entry.DisplayNames),
			DiskMappings:    platformtypes.StringMapToTFMap(entry.DiskMappings),
		})
	}
	return diskTypesModel{PlatformDiskTypes: result}
}
