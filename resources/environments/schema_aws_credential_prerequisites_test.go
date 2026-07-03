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

var awsCredentialPrerequisitesSchemaElements = []test.DataSourceSchemaTestCaseStructure{
	{
		Name:          "'id' should exist",
		Field:         "id",
		Computed:      true,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "'account_id' should exist",
		Field:         "account_id",
		Computed:      true,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "'external_id' should exist",
		Field:         "external_id",
		Computed:      true,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "'policy' should exist",
		Field:         "policy",
		Computed:      true,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "'policies' should exist",
		Field:         "policies",
		Computed:      true,
		Required:      false,
		AttributeType: schema.MapAttribute{},
	},
}

func TestAwsCredentialPrerequisitesSchemaAttributeNumber(t *testing.T) {
	expected := 5
	attrs := createAwsCredentialPrerequisitesTestSchema()
	if len(attrs) != expected {
		t.Errorf("The number of fields in the awsCredentialPrerequisites schema should be: %d but it is: %d", expected, len(attrs))
		t.FailNow()
	}
}

func TestAwsCredentialPrerequisitesSchemaContainsElements(t *testing.T) {
	test.PerformDataSourceSchemaValidation(t, createAwsCredentialPrerequisitesTestSchema(), awsCredentialPrerequisitesSchemaElements)
}

func createAwsCredentialPrerequisitesTestSchema() map[string]schema.Attribute {
	ds := &awsCredentialPrerequisitesDataSource{}
	resp := &datasource.SchemaResponse{}
	ds.Schema(context.TODO(), datasource.SchemaRequest{}, resp)
	return resp.Schema.Attributes
}
