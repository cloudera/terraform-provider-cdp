// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package iam

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func TestMachineUserResourceRoleAssignmentResourceSchemaContainsExpectedFields(t *testing.T) {
	cases := []SchemaTestCaseStructure{
		{
			name:             "id must exist",
			field:            "id",
			computed:         true,
			shouldBeRequired: false,
		},
		{
			name:             "machine_user must exist",
			field:            "machine_user",
			computed:         false,
			shouldBeRequired: true,
		},
		{
			name:             "resource_crn must exist",
			field:            "resource_crn",
			computed:         false,
			shouldBeRequired: true,
		},
		{
			name:             "resource_role_crn must exist",
			field:            "resource_role_crn",
			computed:         false,
			shouldBeRequired: true,
		},
	}

	underTestAttributes := createFilledMachineUserResourceRoleAssignmentResourceTestObject()

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			PerformSchemaValidationForResource(t, test, underTestAttributes[test.field])
		})
	}
}

func createFilledMachineUserResourceRoleAssignmentResourceTestObject() map[string]schema.Attribute {
	res := &machineUserResourceRoleAssignmentResource{}
	schemaResponse := &resource.SchemaResponse{}
	res.Schema(context.TODO(), resource.SchemaRequest{}, schemaResponse)

	return schemaResponse.Schema.Attributes
}
