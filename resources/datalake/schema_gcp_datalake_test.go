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
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func TestCommonSchemaElementsExist(t *testing.T) {
	SchemaContainsCommonElements(t, createFilledTestObject())
}

func TestGcpSpecificElements(t *testing.T) {
	cases := []TestCaseStructure{
		{
			name:             "cloud_provider_configuration should exist",
			field:            "cloud_provider_configuration",
			computed:         false,
			shouldBeRequired: true,
			attributeType:    schema.SingleNestedAttribute{},
		},
	}

	underTestAttributes := createFilledTestObject()

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

func createFilledTestObject() map[string]schema.Attribute {
	res := &gcpDatalakeResource{}
	schemaResponse := &resource.SchemaResponse{}
	res.Schema(context.TODO(), resource.SchemaRequest{}, schemaResponse)

	return schemaResponse.Schema.Attributes
}
