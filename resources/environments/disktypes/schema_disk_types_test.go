// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package disktypes

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/cloudera/terraform-provider-cdp/utils/test"
)

var diskTypesSchemaElements = []test.DataSourceSchemaTestCaseStructure{
	{
		Name:          "platform_disk_types field must exist and be computed",
		Field:         "platform_disk_types",
		Computed:      true,
		Required:      false,
		AttributeType: schema.ListNestedAttribute{},
	},
}

func TestDiskTypesSchemaAttributeNumber(t *testing.T) {
	expected := 1
	if len(diskTypesSchema.Attributes) != expected {
		t.Errorf("expected %d attributes, got %d", expected, len(diskTypesSchema.Attributes))
	}
}

func TestDiskTypesSchemaContainsElements(t *testing.T) {
	test.PerformDataSourceSchemaValidation(t, diskTypesSchema.Attributes, diskTypesSchemaElements)
}
