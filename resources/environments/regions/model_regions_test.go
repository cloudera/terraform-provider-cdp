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
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

func TestRegionsFromResponseNil(t *testing.T) {
	var m regionsModel
	m.fromResponse(nil)
	if m.AvailabilityZonesResult != nil {
		t.Error("nil response should leave AZ result nil")
	}
}

func TestRegionsFromResponse(t *testing.T) {
	var m regionsModel
	m.fromResponse(&models.ListRegionsResponse{
		Names:         []string{"us-west-2", "eu-west-1"},
		DefaultRegion: "us-west-2",
		DisplayNames:  map[string]string{"us-west-2": "US West (Oregon)"},
		AvailabilityZones: []*models.RegionAvailabilityZones{
			{
				Region:            "us-west-2",
				AvailabilityZones: []string{"us-west-2a", "us-west-2b"},
			},
		},
	})

	if len(m.Names.Elements()) != 2 {
		t.Fatalf("expected 2 names, got %d", len(m.Names.Elements()))
	}
	if m.DefaultRegion.ValueString() != "us-west-2" {
		t.Errorf("expected us-west-2, got %s", m.DefaultRegion.ValueString())
	}
	if m.DisplayNames.IsNull() {
		t.Fatal("expected non-null display names")
	}

	dnVal, ok := m.DisplayNames.Elements()["us-west-2"]
	if !ok {
		t.Fatal("expected us-west-2 in display names")
	}
	if dnVal.(types.String).ValueString() != "US West (Oregon)" {
		t.Errorf("wrong display name: %s", dnVal.(types.String).ValueString())
	}

	if len(m.AvailabilityZonesResult) != 1 {
		t.Fatalf("expected 1 AZ entry, got %d", len(m.AvailabilityZonesResult))
	}
	az := m.AvailabilityZonesResult[0]
	if az.Region.ValueString() != "us-west-2" {
		t.Errorf("expected us-west-2, got %s", az.Region.ValueString())
	}
	if len(az.AvailabilityZones.Elements()) != 2 {
		t.Errorf("expected 2 AZs, got %d", len(az.AvailabilityZones.Elements()))
	}
}

func TestRegionsFromResponseEmptyRegion(t *testing.T) {
	var m regionsModel
	m.fromResponse(&models.ListRegionsResponse{
		DefaultRegion: "",
	})
	if !m.DefaultRegion.IsNull() {
		t.Error("expected null default region for empty string")
	}
}
