// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package recipe

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

var schemaElements = []utils.SchemaTestCaseStructure{
	{
		Name:             "crn field must exist and be valid",
		Field:            "crn",
		Computed:         true,
		ShouldBeRequired: false,
	},
	{
		Name:             "id field must exist and be valid",
		Field:            "id",
		Computed:         true,
		ShouldBeRequired: false,
	},
	{
		Name:             "name field must exist and be valid",
		Field:            "name",
		Computed:         false,
		ShouldBeRequired: true,
	},
	{
		Name:             "content field must exist and be valid",
		Field:            "content",
		Computed:         false,
		ShouldBeRequired: true,
	},
	{
		Name:             "type field must exist and be valid",
		Field:            "type",
		Computed:         false,
		ShouldBeRequired: true,
	},
	{
		Name:             "description field must exist and be valid",
		Field:            "description",
		Computed:         false,
		ShouldBeRequired: false,
	},
}

func TestSchemaContainsElements(t *testing.T) {
	for _, test := range schemaElements {
		performValidation(t, test, recipeSchema.Attributes[test.Field])
	}
}

func performValidation(t *testing.T, test utils.SchemaTestCaseStructure, attr schema.Attribute) {
	t.Run(test.Name, func(t *testing.T) {
		if attr == nil {
			t.Errorf("The following field does not exists, however it should: %s", test.Field)
			t.FailNow()
		}
		if attr.IsRequired() != test.ShouldBeRequired {
			t.Errorf("The '%s' filed's >required< property should be: %t", test.Field, test.ShouldBeRequired)
		}
		if attr.IsComputed() != test.Computed {
			t.Errorf("The '%s' filed's >computed< property should be: %t", test.Field, test.Computed)
		}
	})
}
