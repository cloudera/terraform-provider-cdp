// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package networks

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/resources/environments/platformtypes"
)

type networksModel struct {
	Credential       types.String           `tfsdk:"credential"`
	Region           types.String           `tfsdk:"region"`
	AvailabilityZone types.String           `tfsdk:"availability_zone"`
	PlatformVariant  types.String           `tfsdk:"platform_variant"`
	NetworkID        types.String           `tfsdk:"network_id"`
	SharedProjectID  types.String           `tfsdk:"shared_project_id"`
	SubnetIds        types.List             `tfsdk:"subnet_ids"`
	Networks         []platformNetworkModel `tfsdk:"networks"`
}

type platformNetworkModel struct {
	Platform types.String        `tfsdk:"platform"`
	Networks []cloudNetworkModel `tfsdk:"networks"`
}

type cloudNetworkModel struct {
	ID         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	Properties types.Map    `tfsdk:"properties"`
	Subnets    types.Map    `tfsdk:"subnets"`
}

func (m *networksModel) fromResponse(resp *models.ListNetworksResponse) {
	if resp == nil {
		return
	}
	m.Networks = make([]platformNetworkModel, 0, len(resp.Networks))
	for _, pn := range resp.Networks {
		if pn == nil {
			continue
		}
		nets := make([]cloudNetworkModel, 0, len(pn.Networks))
		for _, cn := range pn.Networks {
			if cn == nil {
				continue
			}
			nets = append(nets, cloudNetworkModel{
				ID:         platformtypes.StringValueOrNull(cn.ID),
				Name:       platformtypes.StringValueOrNull(cn.Name),
				Properties: platformtypes.StringMapToTFMap(cn.Properties),
				Subnets:    platformtypes.StringMapToTFMap(cn.Subnets),
			})
		}
		m.Networks = append(m.Networks, platformNetworkModel{
			Platform: platformtypes.StringValueOrNull(pn.Platform),
			Networks: nets,
		})
	}
}
