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
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"testing"
)

func TestAwsSchemaContainsCommonElements(t *testing.T) {
	SchemaContainsCommonElements(t, createFilledAwsDatahubResourceTestObject())
}

func TestAwsSchemaContainsInstanceGroup(t *testing.T) {
	test := TestCaseStructure{
		name:             "instance_group should exist",
		field:            "instance_group",
		computed:         false,
		shouldBeRequired: false,
	}

	underTestAttributes := createFilledAwsDatahubResourceTestObject()

	PerformValidation(t, test, underTestAttributes[test.field])
}

func createFilledAwsDatahubResourceTestObject() map[string]schema.Attribute {
	res := &awsDatahubResource{}
	schemaResponse := &resource.SchemaResponse{}
	res.Schema(context.TODO(), resource.SchemaRequest{}, schemaResponse)

	return schemaResponse.Schema.Attributes
}
