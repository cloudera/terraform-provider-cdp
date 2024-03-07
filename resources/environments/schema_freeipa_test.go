// Copyright 2024 Cloudera. All Rights Reserved.
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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"reflect"
	"testing"
)

type TestCaseStructure struct {
	name          string
	field         string
	computed      bool
	optional      bool
	attributeType schema.Attribute
}

var elements = []TestCaseStructure{
	{
		name:          "catalog field must exists and be valid",
		field:         "catalog",
		optional:      true,
		computed:      false,
		attributeType: schema.StringAttribute{},
	},
	{
		name:          "image_id should exists and be valid",
		field:         "image_id",
		optional:      true,
		computed:      false,
		attributeType: schema.StringAttribute{},
	},
	{
		name:          "os should exists and be valid",
		field:         "os",
		optional:      true,
		computed:      false,
		attributeType: schema.StringAttribute{},
	},
	{
		name:          "instance_count_by_group should exists and be valid",
		field:         "instance_count_by_group",
		optional:      true,
		computed:      true,
		attributeType: schema.Int64Attribute{},
	},
	{
		name:          "instance_type must exists and be valid",
		field:         "instance_type",
		optional:      true,
		computed:      false,
		attributeType: schema.StringAttribute{},
	},
	{
		name:          "instances should exists and be valid",
		field:         "instances",
		optional:      true,
		computed:      false,
		attributeType: schema.SetNestedAttribute{},
	},
	{
		name:          "multi_az should exists and be valid",
		field:         "multi_az",
		optional:      true,
		computed:      true,
		attributeType: schema.BoolAttribute{},
	},
	{
		name:          "recipes should exists and be valid",
		field:         "recipes",
		optional:      true,
		computed:      true,
		attributeType: schema.SetAttribute{},
	},
}

func TestElements(t *testing.T) {
	for _, test := range elements {
		t.Run(test.name, func(t *testing.T) {
			if FreeIpaSchema.Attributes[test.field] == nil {
				t.Errorf("The following field does not exists, however it should: " + test.field)
				t.FailNow()
			}
			if FreeIpaSchema.Attributes[test.field].IsRequired() != !test.optional {
				t.Errorf("The '%s' filed's >required< property should be: %t", test.field, test.optional)
			}
			if FreeIpaSchema.Attributes[test.field].IsComputed() != test.computed {
				t.Errorf("The '%s' filed's >computed< property should be: %t", test.field, test.computed)
			}
			var currentType = reflect.TypeOf(FreeIpaSchema.Attributes[test.field])
			var expectedType = reflect.TypeOf(test.attributeType)
			if currentType != expectedType {
				t.Errorf("The '%s' field's type should be: %t, instead of %t", test.field, expectedType, currentType)
			}
		})
	}
}
