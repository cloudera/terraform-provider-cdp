// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package auditcredential

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/cloudera/terraform-provider-cdp/utils/test"
)

func TestAwsGovCloudAuditCredentialCommonSchemaElementsExist(t *testing.T) {
	AuditCredentialSchemaContainsCommonElements(t, createFilledAwsGovCloudAuditCredentialTestObject())
}

func TestAwsGovCloudAuditCredentialSpecificElements(t *testing.T) {
	cases := []test.ResourceSchemaTestCaseStructure{
		{
			Name:          "'role_arn' should exist",
			Field:         "role_arn",
			Computed:      false,
			Required:      true,
			AttributeType: schema.StringAttribute{},
		},
	}

	test.PerformResourceSchemaValidation(t, createFilledAwsGovCloudAuditCredentialTestObject(), cases)
}

func TestAwsGovCloudAuditCredentialSchemaAttributeNumber(t *testing.T) {
	expected := 5
	attrs := createFilledAwsGovCloudAuditCredentialTestObject()
	if len(attrs) != expected {
		t.Errorf("The number of fields in the AWS GovCloud Audit Credential schema should be: %d but it is: %d", expected, len(attrs))
	}
}

func createFilledAwsGovCloudAuditCredentialTestObject() map[string]schema.Attribute {
	res := &awsGovCloudAuditCredentialResource{}
	schemaResponse := &resource.SchemaResponse{}
	res.Schema(context.TODO(), resource.SchemaRequest{}, schemaResponse)
	return schemaResponse.Schema.Attributes
}
