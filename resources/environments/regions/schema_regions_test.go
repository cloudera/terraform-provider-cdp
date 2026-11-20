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

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/cloudera/terraform-provider-cdp/utils/test"
)

var regionsSchemaElements = []test.DataSourceSchemaTestCaseStructure{
	{
		Name:          "credential field must exist and be optional",
		Field:         "credential",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "environment_crn field must exist and be optional",
		Field:         "environment_crn",
		Computed:      false,
		Required:      false,
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
		Name:          "availability_zones_needed field must exist and be optional",
		Field:         "availability_zones_needed",
		Computed:      false,
		Required:      false,
		AttributeType: schema.BoolAttribute{},
	},
	{
		Name:          "platform_variant field must exist and be optional",
		Field:         "platform_variant",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "names field must exist and be computed",
		Field:         "names",
		Computed:      true,
		Required:      false,
		AttributeType: schema.ListAttribute{},
	},
	{
		Name:          "default_region field must exist and be computed",
		Field:         "default_region",
		Computed:      true,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "display_names field must exist and be computed",
		Field:         "display_names",
		Computed:      true,
		Required:      false,
		AttributeType: schema.MapAttribute{},
	},
	{
		Name:          "availability_zones_result field must exist and be computed",
		Field:         "availability_zones_result",
		Computed:      true,
		Required:      false,
		AttributeType: schema.ListNestedAttribute{},
	},
}

func TestRegionsSchemaAttributeNumber(t *testing.T) {
	expected := 10
	if len(regionsSchema.Attributes) != expected {
		t.Errorf("expected %d attributes, got %d", expected, len(regionsSchema.Attributes))
	}
}

func TestRegionsSchemaContainsElements(t *testing.T) {
	test.PerformDataSourceSchemaValidation(t, regionsSchema.Attributes, regionsSchemaElements)
}
