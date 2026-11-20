// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package regions

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/resources/environments/platformtypes"
)

type regionsModel struct {
	Credential              types.String                   `tfsdk:"credential"`
	EnvironmentCrn          types.String                   `tfsdk:"environment_crn"`
	Region                  types.String                   `tfsdk:"region"`
	AvailabilityZone        types.String                   `tfsdk:"availability_zone"`
	AvailabilityZonesNeeded types.Bool                     `tfsdk:"availability_zones_needed"`
	PlatformVariant         types.String                   `tfsdk:"platform_variant"`
	Names                   types.List                     `tfsdk:"names"`
	DefaultRegion           types.String                   `tfsdk:"default_region"`
	DisplayNames            types.Map                      `tfsdk:"display_names"`
	AvailabilityZonesResult []regionAvailabilityZonesModel `tfsdk:"availability_zones_result"`
}

type regionAvailabilityZonesModel struct {
	Region            types.String `tfsdk:"region"`
	AvailabilityZones types.List   `tfsdk:"availability_zones"`
}

func (m *regionsModel) fromResponse(resp *models.ListRegionsResponse) {
	if resp == nil {
		return
	}

	m.Names = platformtypes.StringSliceToList(resp.Names)
	m.DefaultRegion = platformtypes.StringValueOrNull(resp.DefaultRegion)
	m.DisplayNames = platformtypes.StringMapToTFMap(resp.DisplayNames)

	m.AvailabilityZonesResult = make([]regionAvailabilityZonesModel, len(resp.AvailabilityZones))
	for i, az := range resp.AvailabilityZones {
		m.AvailabilityZonesResult[i] = regionAvailabilityZonesModel{
			Region:            platformtypes.StringValueOrNull(az.Region),
			AvailabilityZones: platformtypes.StringSliceToList(az.AvailabilityZones),
		}
	}
}
