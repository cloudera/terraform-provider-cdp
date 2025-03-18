// Copyright 2025 Cloudera. All Rights Reserved.
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
)

var idBrokerMappingsSchemaElements = []SchemaTestCaseStructure{
	{
		name:             "id field must exist and be valid",
		field:            "id",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "data_access_role field must exist and be valid",
		field:            "data_access_role",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "environment_name field must exist and be valid",
		field:            "environment_name",
		computed:         false,
		shouldBeRequired: true,
	},
	{
		name:             "environment_crn field must exist and be valid",
		field:            "environment_crn",
		computed:         false,
		shouldBeRequired: true,
	},
	{
		name:             "id field must exist and be valid",
		field:            "id",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "mappings field must exist and be valid",
		field:            "mappings",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "ranger_audit_role field must exist and be valid",
		field:            "ranger_audit_role",
		computed:         false,
		shouldBeRequired: true,
	},
	{
		name:             "ranger_cloud_access_authorizer_role field must exist and be valid",
		field:            "ranger_cloud_access_authorizer_role",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "set_empty_mappings field must exist and be valid",
		field:            "set_empty_mappings",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "mappings_version field must exist and be valid",
		field:            "mappings_version",
		computed:         true,
		shouldBeRequired: false,
	},
}

func TestIdBrokerMappingsSchemaAttributeNumber(t *testing.T) {
	expected := 9
	if len(IDBrokerMappingSchema.Attributes) != expected {
		t.Errorf("The number of fields in the IDBrokerMapping schema should be: %d but it is: %d", expected, len(IDBrokerMappingSchema.Attributes))
		t.FailNow()
	}
}

func TestIdBrokerMappingsSchemaContainsElements(t *testing.T) {
	for _, test := range idBrokerMappingsSchemaElements {
		performResourceSchemaValidation(t, test, IDBrokerMappingSchema.Attributes[test.field])
	}
}
