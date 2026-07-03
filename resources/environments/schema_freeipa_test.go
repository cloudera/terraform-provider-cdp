// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package environments

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/cloudera/terraform-provider-cdp/utils/test"
)

var freeIpaSchemaElements = []test.ResourceSchemaTestCaseStructure{
	{
		Name:          "'catalog' should exist",
		Field:         "catalog",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "'image_id' should exist",
		Field:         "image_id",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "'os' should exist",
		Field:         "os",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "'instance_count_by_group' should exist",
		Field:         "instance_count_by_group",
		Computed:      true,
		Required:      false,
		AttributeType: schema.Int32Attribute{},
	},
	{
		Name:          "'instance_type' should exist",
		Field:         "instance_type",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "'architecture' should exist",
		Field:         "architecture",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "'instances' should exist",
		Field:         "instances",
		Computed:      true,
		Required:      false,
		AttributeType: schema.SetNestedAttribute{},
	},
	{
		Name:          "'multi_az' should exist",
		Field:         "multi_az",
		Computed:      true,
		Required:      false,
		AttributeType: schema.BoolAttribute{},
	},
	{
		Name:          "'recipes' should exist",
		Field:         "recipes",
		Computed:      true,
		Required:      false,
		AttributeType: schema.SetAttribute{},
	},
}

func TestFreeIpaSchemaAttributeNumber(t *testing.T) {
	expected := 9
	if len(FreeIpaSchema.Attributes) != expected {
		t.Errorf("The number of fields in the FreeIpaSchema should be: %d but it is: %d", expected, len(FreeIpaSchema.Attributes))
		t.FailNow()
	}
}

func TestFreeIpaSchemaContainsElements(t *testing.T) {
	test.PerformResourceSchemaValidation(t, FreeIpaSchema.Attributes, freeIpaSchemaElements)
}
