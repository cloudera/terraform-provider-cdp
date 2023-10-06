// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datalake

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type TestCaseStructure struct {
	name             string
	field            string
	computed         bool
	shouldBeRequired bool
}

var commonElementCaseSet = []TestCaseStructure{
	{
		name:             "'id' field must exist",
		field:            "id",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "'polling_options' should exist",
		field:            "polling_options",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "'creation_date' should exist",
		field:            "creation_date",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "'crn' should exist",
		field:            "crn",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "'datalake_name' must exist",
		field:            "datalake_name",
		computed:         false,
		shouldBeRequired: true,
	},
	{
		name:             "'enable_ranger_raz' should exist",
		field:            "enable_ranger_raz",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "'environment_crn' should exist",
		field:            "environment_crn",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "'environment_name' must exist",
		field:            "environment_name",
		computed:         false,
		shouldBeRequired: true,
	},
	{
		name:             "'image' should exist",
		field:            "image",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "'java_version' should exist",
		field:            "java_version",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "'recipes' should exist",
		field:            "recipes",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "'runtime' should exist",
		field:            "runtime",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "'scale' should exist",
		field:            "scale",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "'status' should exist",
		field:            "status",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "'status_reason' should exist",
		field:            "status_reason",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "'multi_az' should exist",
		field:            "multi_az",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "'tags' should exist",
		field:            "tags",
		computed:         false,
		shouldBeRequired: false,
	},
}

func TestRootElements(t *testing.T) {
	SchemaContainsCommonElements(t, generalAttributes)
}

func SchemaContainsCommonElements(t *testing.T, providerSpecificSchema map[string]schema.Attribute) {
	for _, test := range commonElementCaseSet {
		t.Run(test.name, func(t *testing.T) {
			if providerSpecificSchema[test.field] == nil {
				t.Errorf("The following field does not exists, however it should: " + test.field)
				t.FailNow()
			}
			if providerSpecificSchema[test.field].IsRequired() != test.shouldBeRequired {
				t.Errorf("The '%s' filed's >required< property should be: %t", test.field, test.shouldBeRequired)
			}
			if providerSpecificSchema[test.field].IsComputed() != test.computed {
				t.Errorf("The '%s' filed's >computed< property should be: %t", test.field, test.computed)
			}
		})
	}
}
