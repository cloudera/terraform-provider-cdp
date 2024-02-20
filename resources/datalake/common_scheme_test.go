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
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type TestCaseStructure struct {
	name             string
	field            string
	computed         bool
	shouldBeRequired bool
	attributeType    schema.Attribute
}

var (
	commonElementCaseSet = []TestCaseStructure{
		{
			name:             "'id' field must exist",
			field:            "id",
			computed:         true,
			shouldBeRequired: false,
			attributeType:    schema.StringAttribute{},
		},
		{
			name:             "'polling_options' should exist",
			field:            "polling_options",
			computed:         false,
			shouldBeRequired: false,
			attributeType:    schema.SingleNestedAttribute{},
		},
		{
			name:             "'creation_date' should exist",
			field:            "creation_date",
			computed:         true,
			shouldBeRequired: false,
			attributeType:    schema.StringAttribute{},
		},
		{
			name:             "'crn' should exist",
			field:            "crn",
			computed:         true,
			shouldBeRequired: false,
			attributeType:    schema.StringAttribute{},
		},
		{
			name:             "'datalake_name' must exist",
			field:            "datalake_name",
			computed:         false,
			shouldBeRequired: true,
			attributeType:    schema.StringAttribute{},
		},
		{
			name:             "'enable_ranger_raz' should exist",
			field:            "enable_ranger_raz",
			computed:         true,
			shouldBeRequired: false,
			attributeType:    schema.BoolAttribute{},
		},
		{
			name:             "'environment_crn' should exist",
			field:            "environment_crn",
			computed:         true,
			shouldBeRequired: false,
			attributeType:    schema.StringAttribute{},
		},
		{
			name:             "'environment_name' must exist",
			field:            "environment_name",
			computed:         false,
			shouldBeRequired: true,
			attributeType:    schema.StringAttribute{},
		},
		{
			name:             "'image' should exist",
			field:            "image",
			computed:         false,
			shouldBeRequired: false,
			attributeType:    schema.SingleNestedAttribute{},
		},
		{
			name:             "'java_version' should exist",
			field:            "java_version",
			computed:         false,
			shouldBeRequired: false,
			attributeType:    schema.Int64Attribute{},
		},
		{
			name:             "'recipes' should exist",
			field:            "recipes",
			computed:         false,
			shouldBeRequired: false,
			attributeType:    schema.SetNestedAttribute{},
		},
		{
			name:             "'runtime' should exist",
			field:            "runtime",
			computed:         false,
			shouldBeRequired: false,
			attributeType:    schema.StringAttribute{},
		},
		{
			name:             "'scale' should exist",
			field:            "scale",
			computed:         true,
			shouldBeRequired: false,
			attributeType:    schema.StringAttribute{},
		},
		{
			name:             "'status' should exist",
			field:            "status",
			computed:         true,
			shouldBeRequired: false,
			attributeType:    schema.StringAttribute{},
		},
		{
			name:             "'status_reason' should exist",
			field:            "status_reason",
			computed:         true,
			shouldBeRequired: false,
			attributeType:    schema.StringAttribute{},
		},
		{
			name:             "'multi_az' should exist",
			field:            "multi_az",
			computed:         true,
			shouldBeRequired: false,
			attributeType:    schema.BoolAttribute{},
		},
		{
			name:             "'tags' should exist",
			field:            "tags",
			computed:         false,
			shouldBeRequired: false,
			attributeType:    schema.MapAttribute{},
		},
	}
)

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
			var currentType = reflect.TypeOf(providerSpecificSchema[test.field])
			var expectedType = reflect.TypeOf(test.attributeType)
			if currentType != expectedType {
				t.Errorf("The '%s' field's type should be: %t, instead of %t", test.field, expectedType, currentType)
			}
		})
	}
}
