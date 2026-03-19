// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package environments

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/cloudera/terraform-provider-cdp/utils/test"
)

var environmentListSchemaElements = []test.DataSourceSchemaTestCaseStructure{
	{
		Name:          "environments field must exist and be valid",
		Field:         "environments",
		Computed:      true,
		Required:      false,
		AttributeType: schema.SetNestedAttribute{},
	},
}

func TestEnvironmentListSchemaAttributeNumber(t *testing.T) {
	expected := 1
	if len(EnvironmentListSchema.Attributes) != expected {
		t.Errorf("The number of fields in the environment list schema should be: %d but it is: %d", expected, len(EnvironmentListSchema.Attributes))
		t.FailNow()
	}
}

func TestEnvironmentListSchemaContainsElements(t *testing.T) {
	test.PerformDataSourceSchemaValidation(t, EnvironmentListSchema.Attributes, environmentListSchemaElements)
}
