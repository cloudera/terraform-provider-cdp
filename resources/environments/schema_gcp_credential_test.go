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
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/cloudera/terraform-provider-cdp/utils/test"
)

func TestGcpCredentialCommonSchemaElementsExist(t *testing.T) {
	CredentialSchemaContainsCommonElements(t, createFilledGcpCredentialTestObject())
}

func TestGcpCredentialSpecificElements(t *testing.T) {
	cases := []test.ResourceSchemaTestCaseStructure{
		{
			Name:             "'credential_key' should exist",
			Field:            "credential_key",
			Computed:         false,
			ShouldBeRequired: true,
			AttributeType:    schema.StringAttribute{},
		},
	}

	test.PerformResourceSchemaValidation(t, createFilledGcpCredentialTestObject(), cases)
}

func TestGcpCredentialSchemaAttributeNumber(t *testing.T) {
	expected := 5
	attrs := createFilledGcpCredentialTestObject()
	if len(attrs) != expected {
		t.Errorf("The number of fields in the AzureEnvironmentSchema schema should be: %d but it is: %d", expected, len(attrs))
		t.FailNow()
	}
}

func createFilledGcpCredentialTestObject() map[string]schema.Attribute {
	res := &gcpCredentialResource{}
	schemaResponse := &resource.SchemaResponse{}
	res.Schema(context.TODO(), resource.SchemaRequest{}, schemaResponse)

	return schemaResponse.Schema.Attributes
}
