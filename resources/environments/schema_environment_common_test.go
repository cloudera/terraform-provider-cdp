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

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/cloudera/terraform-provider-cdp/utils/test"
)

var securityAccessSchemaElements = []test.ResourceSchemaTestCaseStructure{
	{
		Name:          "'cidr' should exist",
		Field:         "cidr",
		Computed:      true,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "'default_security_group_id' should exist",
		Field:         "default_security_group_id",
		Computed:      true,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "'default_security_group_ids' should exist",
		Field:         "default_security_group_ids",
		Computed:      false,
		Required:      false,
		AttributeType: schema.SetAttribute{},
	},
	{
		Name:          "'security_group_id_for_knox' should exist",
		Field:         "security_group_id_for_knox",
		Computed:      true,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "'security_group_ids_for_knox' should exist",
		Field:         "security_group_ids_for_knox",
		Computed:      false,
		Required:      false,
		AttributeType: schema.SetAttribute{},
	},
}

func TestSecurityAccessSchemaAttributeNumber(t *testing.T) {
	expected := 5
	if len(securityAccess.Attributes) != expected {
		t.Errorf("The number of fields in the securityAccess schema should be: %d but it is: %d", expected, len(securityAccess.Attributes))
		t.FailNow()
	}
}

func TestSecurityAccessSchemaContainsElements(t *testing.T) {
	test.PerformResourceSchemaValidation(t, securityAccess.Attributes, securityAccessSchemaElements)
}
