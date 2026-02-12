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

	"github.com/cloudera/terraform-provider-cdp/utils/test"
)

func TestRuntimeSchemaElements(t *testing.T) {
	cases := []test.ResourceSchemaTestCaseStructure{
		{
			Name:          "versions should exist",
			Field:         "versions",
			Computed:      true,
			Required:      false,
			AttributeType: schema.SetNestedAttribute{},
		},
	}

	underTestAttributes := createFilledRuntimeTestObject()

	for _, toTest := range cases {
		t.Run(toTest.Name, func(t *testing.T) {
			if underTestAttributes[toTest.Field] == nil {
				t.Errorf("The following field does not exists, however it should: %s", toTest.Field)
				t.FailNow()
			}
			if underTestAttributes[toTest.Field].IsRequired() != toTest.Required {
				t.Errorf("The '%s' field's >required< property should be: %t", toTest.Field, toTest.Required)
			}
			if underTestAttributes[toTest.Field].IsComputed() != toTest.Computed {
				t.Errorf("The '%s' field's >Computed< property should be: %t", toTest.Field, toTest.Computed)
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
