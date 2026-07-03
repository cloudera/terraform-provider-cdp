// Copyright 2026 Cloudera. All Rights Reserved.
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

var keytabSchemaElements = []test.DataSourceSchemaTestCaseStructure{
	{
		Name:          "'environment' should exist",
		Field:         "environment",
		Computed:      false,
		Required:      true,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "'actor_crn' should exist",
		Field:         "actor_crn",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "'keytab' should exist",
		Field:         "keytab",
		Computed:      true,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
}

func TestKeytabSchemaAttributeNumber(t *testing.T) {
	expected := 3
	if len(KeytabSchema.Attributes) != expected {
		t.Errorf("The number of fields in the KeytabSchema should be: %d but it is: %d", expected, len(KeytabSchema.Attributes))
		t.FailNow()
	}
}

func TestKeytabSchemaContainsElements(t *testing.T) {
	test.PerformDataSourceSchemaValidation(t, KeytabSchema.Attributes, keytabSchemaElements)
}
