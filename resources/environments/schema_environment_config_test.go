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
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/cloudera/terraform-provider-cdp/utils/test"
)

var environmentConfigSchemaElements = []test.DataSourceSchemaTestCaseStructure{
	{
		Name:          "'name' should exist",
		Field:         "name",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "'crn' should exist",
		Field:         "crn",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "'aws' should exist",
		Field:         "aws",
		Computed:      true,
		Required:      false,
		AttributeType: schema.SingleNestedAttribute{},
	},
	{
		Name:          "'azure' should exist",
		Field:         "azure",
		Computed:      true,
		Required:      false,
		AttributeType: schema.SingleNestedAttribute{},
	},
	{
		Name:          "'gcp' should exist",
		Field:         "gcp",
		Computed:      true,
		Required:      false,
		AttributeType: schema.SingleNestedAttribute{},
	},
}

func TestEnvironmentConfigSchemaAttributeNumber(t *testing.T) {
	expected := 5
	attrs := createEnvironmentConfigTestSchema()
	if len(attrs) != expected {
		t.Errorf("The number of fields in the environmentConfig schema should be: %d but it is: %d", expected, len(attrs))
		t.FailNow()
	}
}

func TestEnvironmentConfigSchemaContainsElements(t *testing.T) {
	test.PerformDataSourceSchemaValidation(t, createEnvironmentConfigTestSchema(), environmentConfigSchemaElements)
}

func TestEnvironmentConfigAwsSchemaAttributeNumber(t *testing.T) {
	expected := 31
	attrs := createEnvironmentConfigTestSchema()
	awsAttr := attrs["aws"].(schema.SingleNestedAttribute)
	if len(awsAttr.Attributes) != expected {
		t.Errorf("The number of fields in the aws nested schema should be: %d but it is: %d", expected, len(awsAttr.Attributes))
		t.FailNow()
	}
}

func TestEnvironmentConfigAzureSchemaAttributeNumber(t *testing.T) {
	expected := 36
	attrs := createEnvironmentConfigTestSchema()
	azureAttr := attrs["azure"].(schema.SingleNestedAttribute)
	if len(azureAttr.Attributes) != expected {
		t.Errorf("The number of fields in the azure nested schema should be: %d but it is: %d", expected, len(azureAttr.Attributes))
		t.FailNow()
	}
}

func TestEnvironmentConfigGcpSchemaAttributeNumber(t *testing.T) {
	expected := 27
	attrs := createEnvironmentConfigTestSchema()
	gcpAttr := attrs["gcp"].(schema.SingleNestedAttribute)
	if len(gcpAttr.Attributes) != expected {
		t.Errorf("The number of fields in the gcp nested schema should be: %d but it is: %d", expected, len(gcpAttr.Attributes))
		t.FailNow()
	}
}

func createEnvironmentConfigTestSchema() map[string]schema.Attribute {
	ds := &environmentConfigDataSource{}
	resp := &datasource.SchemaResponse{}
	ds.Schema(context.TODO(), datasource.SchemaRequest{}, resp)
	return resp.Schema.Attributes
}
