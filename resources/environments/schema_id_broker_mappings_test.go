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

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/cloudera/terraform-provider-cdp/utils/test"
)

var idBrokerMappingsSchemaElements = []test.ResourceSchemaTestCaseStructure{
	{
		Name:          "id field must exist and be valid",
		Field:         "id",
		Computed:      true,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "data_access_role Field must exist and be valid",
		Field:         "data_access_role",
		Computed:      true,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "environment_name Field must exist and be valid",
		Field:         "environment_name",
		Computed:      false,
		Required:      true,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "environment_crn Field must exist and be valid",
		Field:         "environment_crn",
		Computed:      false,
		Required:      true,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "id Field must exist and be valid",
		Field:         "id",
		Computed:      true,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "mappings Field must exist and be valid",
		Field:         "mappings",
		Computed:      false,
		Required:      false,
		AttributeType: schema.SetNestedAttribute{},
	},
	{
		Name:          "ranger_audit_role Field must exist and be valid",
		Field:         "ranger_audit_role",
		Computed:      false,
		Required:      true,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "ranger_cloud_access_authorizer_role Field must exist and be valid",
		Field:         "ranger_cloud_access_authorizer_role",
		Computed:      true,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "set_empty_mappings Field must exist and be valid",
		Field:         "set_empty_mappings",
		Computed:      false,
		Required:      false,
		AttributeType: schema.BoolAttribute{},
	},
	{
		Name:          "mappings_version Field must exist and be valid",
		Field:         "mappings_version",
		Computed:      true,
		Required:      false,
		AttributeType: schema.Int64Attribute{},
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
	test.PerformResourceSchemaValidation(t, IDBrokerMappingSchema.Attributes, idBrokerMappingsSchemaElements)
}
