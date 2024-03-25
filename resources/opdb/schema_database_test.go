// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package opdb

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"testing"
)

func TestSchemaContainsCommonElements(t *testing.T) {
	SchemaContainsCommonElements(t, createFilledDatabaseResourceTestObject())
}

func TestSchemaContainsStorageType(t *testing.T) {
	test := TestCaseStructure{
		name:             "storage_type should exist",
		field:            "storage_type",
		computed:         false,
		shouldBeRequired: false,
	}

	underTestAttributes := createFilledDatabaseResourceTestObject()

	PerformValidation(t, test, underTestAttributes[test.field])
}

func TestSchemaContainsScaleType(t *testing.T) {
	test := TestCaseStructure{
		name:             "scale_type should exist",
		field:            "scale_type",
		computed:         false,
		shouldBeRequired: false,
	}

	underTestAttributes := createFilledDatabaseResourceTestObject()

	PerformValidation(t, test, underTestAttributes[test.field])
}

func TestSchemaContainsDisableExternalDB(t *testing.T) {
	test := TestCaseStructure{
		name:             "disable_external_db should exist",
		field:            "disable_external_db",
		computed:         false,
		shouldBeRequired: false,
	}

	underTestAttributes := createFilledDatabaseResourceTestObject()

	PerformValidation(t, test, underTestAttributes[test.field])
}

func createFilledDatabaseResourceTestObject() map[string]schema.Attribute {
	res := &databaseResource{}
	schemaResponse := &resource.SchemaResponse{}
	res.Schema(context.TODO(), resource.SchemaRequest{}, schemaResponse)

	return schemaResponse.Schema.Attributes
}
