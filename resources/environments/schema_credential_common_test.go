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

func CredentialSchemaContainsCommonElements(t *testing.T, underTestAttributes map[string]schema.Attribute) {
	cases := []test.ResourceSchemaTestCaseStructure{
		{
			Name:             "'id' should exist",
			Field:            "id",
			Computed:         true,
			ShouldBeRequired: false,
			AttributeType:    schema.StringAttribute{},
		},
		{
			Name:             "'credential_name' should exist",
			Field:            "credential_name",
			Computed:         false,
			ShouldBeRequired: true,
			AttributeType:    schema.StringAttribute{},
		},
		{
			Name:             "'description' should exist",
			Field:            "description",
			Computed:         false,
			ShouldBeRequired: false,
			AttributeType:    schema.StringAttribute{},
		},
		{
			Name:             "'crn' should exist",
			Field:            "crn",
			Computed:         true,
			ShouldBeRequired: false,
			AttributeType:    schema.StringAttribute{},
		},
	}

	test.PerformResourceSchemaValidation(t, underTestAttributes, cases)
}
