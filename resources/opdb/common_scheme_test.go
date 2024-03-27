// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package opdb

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
		name:             "polling_options should exist",
		field:            "polling_options",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "crn should exist",
		field:            "crn",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "status should exist",
		field:            "status",
		computed:         true,
		shouldBeRequired: false,
	},
}

func TestRootCommonElements(t *testing.T) {
	SchemaContainsCommonElements(t, generalAttributes)
}

func SchemaContainsCommonElements(t *testing.T, providerSpecificSchema map[string]schema.Attribute) {
	for _, test := range commonElementCaseSet {
		PerformValidation(t, test, providerSpecificSchema[test.field])
	}
}

func PerformValidation(t *testing.T, test TestCaseStructure, attr schema.Attribute) {
	t.Run(test.name, func(t *testing.T) {
		if attr == nil {
			t.Errorf("The following field does not exists, however it should: " + test.field)
			t.FailNow()
		}
		if attr.IsRequired() != test.shouldBeRequired {
			t.Errorf("The '%s' filed's >required< property should be: %t", test.field, test.shouldBeRequired)
		}
		if attr.IsComputed() != test.computed {
			t.Errorf("The '%s' filed's >computed< property should be: %t", test.field, test.computed)
		}
	})
}
