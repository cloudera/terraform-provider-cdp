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

func TestAzureCommonSchemaElementsExist(t *testing.T) {
	SchemaContainsCommonElements(t, createFilledTestObject())
}

func TestAzureSpecificElements(t *testing.T) {
	cases := []test.ResourceSchemaTestCaseStructure{
		{
			Name:          "'database_type' should exist",
			Field:         "database_type",
			Computed:      false,
			Required:      false,
			AttributeType: schema.StringAttribute{},
		},
		{
			Name:          "'load_balancer_sku' should exist",
			Field:         "load_balancer_sku",
			Computed:      false,
			Required:      false,
			AttributeType: schema.StringAttribute{},
		},
		{
			Name:          "'flexible_server_delegated_subnet_id' should exist",
			Field:         "flexible_server_delegated_subnet_id",
			Computed:      false,
			Required:      false,
			AttributeType: schema.StringAttribute{},
		},
		{
			Name:          "'multi_az' should exist",
			Field:         "multi_az",
			Computed:      false,
			Required:      false,
			AttributeType: schema.BoolAttribute{},
		},
	}

	underTestAttributes := createFilledAzureTestObject()

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

func TestAzureDatalakeSchemaAttributeNumber(t *testing.T) {
	expected := 24
	attrs := createFilledAzureTestObject()
	if len(attrs) != expected {
		t.Errorf("The number of fields in the AzureDatalakeSchema schema should be: %d but it is: %d", expected, len(attrs))
		t.FailNow()
	}
}

func createFilledAzureTestObject() map[string]schema.Attribute {
	res := &azureDatalakeResource{}
	schemaResponse := &resource.SchemaResponse{}
	res.Schema(context.TODO(), resource.SchemaRequest{}, schemaResponse)

	return schemaResponse.Schema.Attributes
}
