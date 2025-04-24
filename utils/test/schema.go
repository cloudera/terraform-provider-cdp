// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package test

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type ResourceSchemaTestCaseStructure struct {
	Name             string
	Field            string
	Computed         bool
	ShouldBeRequired bool
	AttributeType    schema.Attribute
}

func PerformResourceSchemaValidation(t *testing.T, resourceSchema map[string]schema.Attribute, expectedElements []ResourceSchemaTestCaseStructure) {
	for _, toTest := range expectedElements {
		t.Run(toTest.Name, func(t *testing.T) {
			if resourceSchema[toTest.Field] == nil {
				t.Errorf("The following field does not exists, however it should: %s", toTest.Field)
				t.FailNow()
			}
			if resourceSchema[toTest.Field].IsRequired() != toTest.ShouldBeRequired {
				t.Errorf("The '%s' field's >required< property should be: %t", toTest.Field, toTest.ShouldBeRequired)
			}
			if resourceSchema[toTest.Field].IsComputed() != toTest.Computed {
				t.Errorf("The '%s' field's >Computed< property should be: %t", toTest.Field, toTest.Computed)
			}
			var currentType = reflect.TypeOf(resourceSchema[toTest.Field])
			var expectedType = reflect.TypeOf(toTest.AttributeType)
			if currentType != expectedType {
				t.Errorf("The '%s' field's type should be: %t, instead of %t", toTest.Field, expectedType, currentType)
			}
		})
	}
}
