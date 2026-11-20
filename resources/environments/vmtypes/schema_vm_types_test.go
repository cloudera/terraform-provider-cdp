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
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/cloudera/terraform-provider-cdp/utils/test"
)

var vmTypesSchemaElements = []test.DataSourceSchemaTestCaseStructure{
	{
		Name:          "credential_name field must exist and be optional",
		Field:         "credential_name",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "credential_crn field must exist and be optional",
		Field:         "credential_crn",
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
		Name:          "architecture field must exist and be optional",
		Field:         "architecture",
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
		Name:          "resource_type field must exist and be optional",
		Field:         "resource_type",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "vm_types field must exist and be computed",
		Field:         "vm_types",
		Computed:      true,
		Required:      false,
		AttributeType: schema.ListNestedAttribute{},
	},
}

func TestVMTypesSchemaAttributeNumber(t *testing.T) {
	expected := 8
	if len(vmTypesSchema.Attributes) != expected {
		t.Errorf("expected %d attributes, got %d", expected, len(vmTypesSchema.Attributes))
	}
}

func TestVMTypesSchemaContainsElements(t *testing.T) {
	test.PerformDataSourceSchemaValidation(t, vmTypesSchema.Attributes, vmTypesSchemaElements)
}
