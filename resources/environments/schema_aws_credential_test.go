// Copyright 2023 Cloudera. All Rights Reserved.
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
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/cloudera/terraform-provider-cdp/utils/test"
)

func TestAwsCredentialCommonSchemaElementsExist(t *testing.T) {
	CredentialSchemaContainsCommonElements(t, createFilledAwsCredentialTestObject())
}

func TestAwsCredentialSpecificElements(t *testing.T) {
	cases := []test.ResourceSchemaTestCaseStructure{
		{
			Name:             "'role_arn' should exist",
			Field:            "role_arn",
			Computed:         false,
			ShouldBeRequired: true,
			AttributeType:    schema.StringAttribute{},
		},
		{
			Name:             "'skip_org_policy_decisions' should exist",
			Field:            "skip_org_policy_decisions",
			Computed:         true,
			ShouldBeRequired: false,
			AttributeType:    schema.BoolAttribute{},
		},
		{
			Name:             "'verify_permissions' should exist",
			Field:            "verify_permissions",
			Computed:         true,
			ShouldBeRequired: false,
			AttributeType:    schema.BoolAttribute{},
		},
	}

	test.PerformResourceSchemaValidation(t, createFilledAwsCredentialTestObject(), cases)
}

func TestAwsCredentialSchemaAttributeNumber(t *testing.T) {
	expected := 7
	attrs := createFilledAwsCredentialTestObject()
	if len(attrs) != expected {
		t.Errorf("The number of fields in the AwsEnvironmentSchema schema should be: %d but it is: %d", expected, len(attrs))
		t.FailNow()
	}
}

func createFilledAwsCredentialTestObject() map[string]schema.Attribute {
	res := &awsCredentialResource{}
	schemaResponse := &resource.SchemaResponse{}
	res.Schema(context.TODO(), resource.SchemaRequest{}, schemaResponse)

	return schemaResponse.Schema.Attributes
}
