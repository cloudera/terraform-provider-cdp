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

	"github.com/cloudera/terraform-provider-cdp/utils/test"
)

func TestCommonSchemaElementsExist(t *testing.T) {
	SchemaContainsCommonElements(t, createFilledTestObject())
}

func TestGcpSpecificElements(t *testing.T) {
	cases := []test.ResourceSchemaTestCaseStructure{
		{
			Name:             "cloud_provider_configuration should exist",
			Field:            "cloud_provider_configuration",
			Computed:         false,
			ShouldBeRequired: true,
			AttributeType:    schema.SingleNestedAttribute{},
		},
	}

	underTestAttributes := createFilledTestObject()

	for _, toTest := range cases {
		t.Run(toTest.Name, func(t *testing.T) {
			if underTestAttributes[toTest.Field] == nil {
				t.Errorf("The following field does not exists, however it should: %s", toTest.Field)
				t.FailNow()
			}
			if underTestAttributes[toTest.Field].IsRequired() != toTest.ShouldBeRequired {
				t.Errorf("The '%s' field's >required< property should be: %t", toTest.Field, toTest.ShouldBeRequired)
			}
			if underTestAttributes[toTest.Field].IsComputed() != toTest.Computed {
				t.Errorf("The '%s' field's >Computed< property should be: %t", toTest.Field, toTest.Computed)
			}
		})
	}
}

func TestGcpDatalakeSchemaAttributeNumber(t *testing.T) {
	expected := 19
	attrs := createFilledTestObject()
	if len(attrs) != expected {
		t.Errorf("The number of fields in the GcpDatalakeSchema schema should be: %d but it is: %d", expected, len(attrs))
		t.FailNow()
	}
}

func createFilledTestObject() map[string]schema.Attribute {
	res := &gcpDatalakeResource{}
	schemaResponse := &resource.SchemaResponse{}
	res.Schema(context.TODO(), resource.SchemaRequest{}, schemaResponse)

	return schemaResponse.Schema.Attributes
}
