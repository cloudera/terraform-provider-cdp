// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datahub

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
		name:             "id field must exist",
		field:            "id",
		computed:         true,
		shouldBeRequired: false,
	},
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
	{
		name:             "name must exist",
		field:            "name",
		computed:         false,
		shouldBeRequired: true,
	},
	{
		name:             "destroy_options should exist",
		field:            "destroy_options",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "custom_configurations_name should exist",
		field:            "custom_configurations_name",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "image should exist",
		field:            "image",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "request_template should exist",
		field:            "request_template",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "datahub_database should exist",
		field:            "datahub_database",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "cluster_extension should exist",
		field:            "cluster_extension",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "java_version should exist",
		field:            "java_version",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "tags should exist",
		field:            "tags",
		computed:         false,
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
			t.Errorf("The following field does not exists, however it should: %s", test.field)
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
