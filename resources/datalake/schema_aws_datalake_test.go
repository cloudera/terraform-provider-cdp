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

func TestAwsCommonSchemaElementsExist(t *testing.T) {
	SchemaContainsCommonElements(t, createFilledAwsTestObject())
}

func TestAwsSpecificElements(t *testing.T) {
	cases := []test.ResourceSchemaTestCaseStructure{
		{
			Name:          "'certificate_expiration_state' should exist",
			Field:         "certificate_expiration_state",
			Computed:      true,
			Required:      false,
			AttributeType: schema.StringAttribute{},
		},
		{
			Name:          "'storage_location_base' should exist",
			Field:         "storage_location_base",
			Computed:      false,
			Required:      true,
			AttributeType: schema.StringAttribute{},
		},
		{
			Name:          "'instance_profile' should exist",
			Field:         "instance_profile",
			Computed:      false,
			Required:      true,
			AttributeType: schema.StringAttribute{},
		},
	}

	test.PerformResourceSchemaValidation(t, createFilledAwsTestObject(), cases)
}

func TestAwsDatalakeSchemaAttributeNumber(t *testing.T) {
	expected := 21
	attrs := createFilledAwsTestObject()
	if len(attrs) != expected {
		t.Errorf("The number of fields in the AwsDatalakeSchema schema should be: %d but it is: %d", expected, len(attrs))
		t.FailNow()
	}
}

func createFilledAwsTestObject() map[string]schema.Attribute {
	res := &awsDatalakeResource{}
	schemaResponse := &resource.SchemaResponse{}
	res.Schema(context.TODO(), resource.SchemaRequest{}, schemaResponse)

	return schemaResponse.Schema.Attributes
}
