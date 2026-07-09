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

func TestGcpAuditCredentialCommonSchemaElementsExist(t *testing.T) {
	AuditCredentialSchemaContainsCommonElements(t, createFilledGcpAuditCredentialTestObject())
}

func TestGcpAuditCredentialSpecificElements(t *testing.T) {
	cases := []test.ResourceSchemaTestCaseStructure{
		{
			Name:          "'credential_key' should exist",
			Field:         "credential_key",
			Computed:      false,
			Required:      true,
			AttributeType: schema.StringAttribute{},
		},
	}

	test.PerformResourceSchemaValidation(t, createFilledGcpAuditCredentialTestObject(), cases)
}

func TestGcpAuditCredentialSchemaAttributeNumber(t *testing.T) {
	expected := 5
	attrs := createFilledGcpAuditCredentialTestObject()
	if len(attrs) != expected {
		t.Errorf("The number of fields in the GCP Audit Credential schema should be: %d but it is: %d", expected, len(attrs))
	}
}

func createFilledGcpAuditCredentialTestObject() map[string]schema.Attribute {
	res := &gcpAuditCredentialResource{}
	schemaResponse := &resource.SchemaResponse{}
	res.Schema(context.TODO(), resource.SchemaRequest{}, schemaResponse)
	return schemaResponse.Schema.Attributes
}
