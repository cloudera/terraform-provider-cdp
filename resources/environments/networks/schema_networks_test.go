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

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/cloudera/terraform-provider-cdp/utils/test"
)

var networksSchemaElements = []test.DataSourceSchemaTestCaseStructure{
	{
		Name:          "credential field must exist and be required",
		Field:         "credential",
		Computed:      false,
		Required:      true,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "region field must exist and be optional",
		Field:         "region",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "availability_zone field must exist and be optional",
		Field:         "availability_zone",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "platform_variant field must exist and be optional",
		Field:         "platform_variant",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "network_id field must exist and be optional",
		Field:         "network_id",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "shared_project_id field must exist and be optional",
		Field:         "shared_project_id",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "subnet_ids field must exist and be optional",
		Field:         "subnet_ids",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "networks field must exist and be computed",
		Field:         "networks",
		Computed:      true,
		Required:      false,
		AttributeType: schema.ListNestedAttribute{},
	},
}

func TestNetworksSchemaAttributeNumber(t *testing.T) {
	expected := 8
	if len(networksSchema.Attributes) != expected {
		t.Errorf("expected %d attributes, got %d", expected, len(networksSchema.Attributes))
	}
}

func TestNetworksSchemaContainsElements(t *testing.T) {
	test.PerformDataSourceSchemaValidation(t, networksSchema.Attributes, networksSchemaElements)
}
