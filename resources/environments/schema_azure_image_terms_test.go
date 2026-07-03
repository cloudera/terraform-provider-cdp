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

var azureImageTermsSchemaElements = []test.ResourceSchemaTestCaseStructure{
	{
		Name:          "'id' should exist",
		Field:         "id",
		Computed:      true,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "'accepted' should exist",
		Field:         "accepted",
		Computed:      false,
		Required:      true,
		AttributeType: schema.BoolAttribute{},
	},
}

func TestAzureImageTermsSchemaAttributeNumber(t *testing.T) {
	expected := 2
	if len(AzureImageTermsPolicySchema.Attributes) != expected {
		t.Errorf("The number of fields in the AzureImageTermsPolicySchema should be: %d but it is: %d", expected, len(AzureImageTermsPolicySchema.Attributes))
		t.FailNow()
	}
}

func TestAzureImageTermsSchemaContainsElements(t *testing.T) {
	test.PerformResourceSchemaValidation(t, AzureImageTermsPolicySchema.Attributes, azureImageTermsSchemaElements)
}
