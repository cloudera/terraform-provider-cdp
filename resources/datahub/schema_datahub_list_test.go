// Copyright 2024 Cloudera. All Rights Reserved.
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
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/cloudera/terraform-provider-cdp/utils/test"
)

var datahubListSchemaElements = []test.DataSourceSchemaTestCaseStructure{
	{
		Name:          "datahubs field must exist and be valid",
		Field:         "datahubs",
		Computed:      true,
		Required:      false,
		AttributeType: schema.SetNestedAttribute{},
	},
}

func TestDatahubListSchemaAttributeNumber(t *testing.T) {
	expected := 1
	if len(DatahubListSchema.Attributes) != expected {
		t.Errorf("The number of fields in the datahub list schema should be: %d but it is: %d", expected, len(DatahubListSchema.Attributes))
		t.FailNow()
	}
}

func TestDatahubListSchemaContainsElements(t *testing.T) {
	test.PerformDataSourceSchemaValidation(t, DatahubListSchema.Attributes, datahubListSchemaElements)
}
