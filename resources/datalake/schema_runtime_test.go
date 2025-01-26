// Copyright 2025 Cloudera. All Rights Reserved.
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
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func TestRuntimeSchemaElements(t *testing.T) {
	cases := []TestCaseStructure{
		{
			name:             "versions should exist",
			field:            "versions",
			computed:         true,
			shouldBeRequired: false,
			attributeType:    schema.SetNestedAttribute{},
		},
	}

	underTestAttributes := createFilledRuntimeTestObject()

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			if underTestAttributes[test.field] == nil {
				t.Errorf("The following field does not exists, however it should: %s", test.field)
				t.FailNow()
			}
			if underTestAttributes[test.field].IsRequired() != test.shouldBeRequired {
				t.Errorf("The '%s' filed's >required< property should be: %t", test.field, test.shouldBeRequired)
			}
			if underTestAttributes[test.field].IsComputed() != test.computed {
				t.Errorf("The '%s' filed's >computed< property should be: %t", test.field, test.computed)
			}
		})
	}
}

func createFilledRuntimeTestObject() map[string]schema.Attribute {
	res := &runtimeDataSource{}
	schemaResponse := &datasource.SchemaResponse{}
	res.Schema(context.TODO(), datasource.SchemaRequest{}, schemaResponse)

	return schemaResponse.Schema.Attributes
}
