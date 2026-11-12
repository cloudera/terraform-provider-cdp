// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package freeipaverticalscaling

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/cloudera/terraform-provider-cdp/utils/test"
)

var freeipaVMTypesSchemaElements = []test.DataSourceSchemaTestCaseStructure{
	{
		Name:          "environment_crn field must exist and be required",
		Field:         "environment_crn",
		Computed:      false,
		Required:      true,
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
		Name:          "availability_zones field must exist and be optional",
		Field:         "availability_zones",
		Computed:      false,
		Required:      false,
		AttributeType: schema.ListAttribute{},
	},
	{
		Name:          "instance_type field must exist and be optional",
		Field:         "instance_type",
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

func TestFreeipaVMTypesSchemaAttributeNumber(t *testing.T) {
	expected := 6
	if len(freeipaVMTypesSchema.Attributes) != expected {
		t.Errorf("The number of fields in the freeipa VM types schema should be: %d but it is: %d", expected, len(freeipaVMTypesSchema.Attributes))
		t.FailNow()
	}
}

func TestFreeipaVMTypesSchemaContainsElements(t *testing.T) {
	test.PerformDataSourceSchemaValidation(t, freeipaVMTypesSchema.Attributes, freeipaVMTypesSchemaElements)
}
