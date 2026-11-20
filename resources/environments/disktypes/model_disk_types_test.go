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
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

func TestDiskTypesModelFromResponseNil(t *testing.T) {
	if r := diskTypesModelFromResponse(nil); r.PlatformDiskTypes != nil {
		t.Error("nil response should give nil disk types")
	}
}

func TestDiskTypesModelFromResponseNilElement(t *testing.T) {
	result := diskTypesModelFromResponse(&models.ListDiskTypesResponse{
		PlatformDiskTypes: []*models.PlatformDiskTypesEntry{
			{Platform: "AWS"},
			nil,
			{Platform: "GCP"},
		},
	})
	if len(result.PlatformDiskTypes) != 2 {
		t.Fatalf("expected 2 entries (nil skipped), got %d", len(result.PlatformDiskTypes))
	}
	if result.PlatformDiskTypes[1].Platform.ValueString() != "GCP" {
		t.Errorf("expected GCP at index 1, got %s", result.PlatformDiskTypes[1].Platform.ValueString())
	}
}

func TestDiskTypesModelFromResponseEmpty(t *testing.T) {
	result := diskTypesModelFromResponse(&models.ListDiskTypesResponse{})
	if len(result.PlatformDiskTypes) != 0 {
		t.Errorf("expected 0 entries, got %d", len(result.PlatformDiskTypes))
	}
}

func TestDiskTypesModelFromResponse(t *testing.T) {
	result := diskTypesModelFromResponse(&models.ListDiskTypesResponse{
		PlatformDiskTypes: []*models.PlatformDiskTypesEntry{
			{
				Platform:        "AWS",
				DefaultDiskType: "gp2",
				DiskTypes:       []string{"gp2", "gp3", "io1"},
				DisplayNames:    map[string]string{"gp2": "General Purpose SSD"},
				DiskMappings:    map[string]string{"gp2": "SSD"},
			},
		},
	})

	if len(result.PlatformDiskTypes) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(result.PlatformDiskTypes))
	}

	entry := result.PlatformDiskTypes[0]
	if entry.Platform.ValueString() != "AWS" {
		t.Errorf("expected AWS, got %s", entry.Platform.ValueString())
	}
	if entry.DefaultDiskType.ValueString() != "gp2" {
		t.Errorf("expected gp2, got %s", entry.DefaultDiskType.ValueString())
	}
	if len(entry.DiskTypes.Elements()) != 3 {
		t.Errorf("expected 3 disk types, got %d", len(entry.DiskTypes.Elements()))
	}

	dnVal, ok := entry.DisplayNames.Elements()["gp2"]
	if !ok {
		t.Fatal("expected gp2 in display names")
	}
	if dnVal.(types.String).ValueString() != "General Purpose SSD" {
		t.Errorf("wrong display name: %s", dnVal.(types.String).ValueString())
	}

	dmVal, ok := entry.DiskMappings.Elements()["gp2"]
	if !ok {
		t.Fatal("expected gp2 in disk mappings")
	}
	if dmVal.(types.String).ValueString() != "SSD" {
		t.Errorf("wrong disk mapping: %s", dmVal.(types.String).ValueString())
	}
}
