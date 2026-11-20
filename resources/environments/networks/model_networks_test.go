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
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

func TestNetworksFromResponseNil(t *testing.T) {
	var m networksModel
	m.fromResponse(nil)
	if m.Networks != nil {
		t.Error("nil response should give nil networks")
	}
}

func TestNetworksFromResponse(t *testing.T) {
	var m networksModel
	m.fromResponse(&models.ListNetworksResponse{
		Networks: []*models.PlatformNetwork{
			{
				Platform: "AWS",
				Networks: []*models.CloudNetwork{
					{
						ID:         "vpc-123",
						Name:       "my-vpc",
						Properties: map[string]string{"vpcId": "vpc-123"},
						Subnets:    map[string]string{"subnet-1": "public-a"},
					},
				},
			},
		},
	})

	if len(m.Networks) != 1 {
		t.Fatalf("expected 1 platform network, got %d", len(m.Networks))
	}

	pn := m.Networks[0]
	if pn.Platform.ValueString() != "AWS" {
		t.Errorf("expected AWS, got %s", pn.Platform.ValueString())
	}
	if len(pn.Networks) != 1 {
		t.Fatalf("expected 1 cloud network, got %d", len(pn.Networks))
	}

	cn := pn.Networks[0]
	if cn.ID.ValueString() != "vpc-123" {
		t.Errorf("expected vpc-123, got %s", cn.ID.ValueString())
	}
	if cn.Name.ValueString() != "my-vpc" {
		t.Errorf("expected my-vpc, got %s", cn.Name.ValueString())
	}

	propVal, ok := cn.Properties.Elements()["vpcId"]
	if !ok {
		t.Fatal("expected vpcId in properties")
	}
	if propVal.(types.String).ValueString() != "vpc-123" {
		t.Errorf("wrong property: %s", propVal.(types.String).ValueString())
	}

	subVal, ok := cn.Subnets.Elements()["subnet-1"]
	if !ok {
		t.Fatal("expected subnet-1 in subnets")
	}
	if subVal.(types.String).ValueString() != "public-a" {
		t.Errorf("wrong subnet: %s", subVal.(types.String).ValueString())
	}
}

func TestNetworksFromResponseNilPlatformNetwork(t *testing.T) {
	var m networksModel
	m.fromResponse(&models.ListNetworksResponse{
		Networks: []*models.PlatformNetwork{
			{Platform: "AWS"},
			nil,
			{Platform: "GCP"},
		},
	})
	if len(m.Networks) != 2 {
		t.Fatalf("expected 2 networks (nil skipped), got %d", len(m.Networks))
	}
	if m.Networks[1].Platform.ValueString() != "GCP" {
		t.Errorf("expected GCP at index 1, got %s", m.Networks[1].Platform.ValueString())
	}
}

func TestNetworksFromResponseNilCloudNetwork(t *testing.T) {
	var m networksModel
	m.fromResponse(&models.ListNetworksResponse{
		Networks: []*models.PlatformNetwork{
			{
				Platform: "AWS",
				Networks: []*models.CloudNetwork{
					{ID: "vpc-1"},
					nil,
					{ID: "vpc-3"},
				},
			},
		},
	})
	if len(m.Networks) != 1 {
		t.Fatalf("expected 1 platform network, got %d", len(m.Networks))
	}
	if len(m.Networks[0].Networks) != 2 {
		t.Fatalf("expected 2 cloud networks (nil skipped), got %d", len(m.Networks[0].Networks))
	}
	if m.Networks[0].Networks[1].ID.ValueString() != "vpc-3" {
		t.Errorf("expected vpc-3 at index 1, got %s", m.Networks[0].Networks[1].ID.ValueString())
	}
}

func TestNetworksFromResponseEmpty(t *testing.T) {
	var m networksModel
	m.fromResponse(&models.ListNetworksResponse{
		Networks: []*models.PlatformNetwork{},
	})
	if len(m.Networks) != 0 {
		t.Errorf("expected 0 networks, got %d", len(m.Networks))
	}
}
