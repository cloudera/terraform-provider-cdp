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
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
)

type SchemaTestCaseStructure struct {
	name             string
	field            string
	computed         bool
	shouldBeRequired bool
}

func PerformSchemaValidationForResource(t *testing.T, test SchemaTestCaseStructure, attr schema.Attribute) {
	t.Run(test.name, func(t *testing.T) {
		if attr == nil {
			t.Errorf("The following field does not exists, however it should: %s", test.field)
			t.FailNow()
		}
		if attr.IsRequired() != test.shouldBeRequired {
			t.Errorf("The '%s' field's >required< property should be: %t", test.field, test.shouldBeRequired)
		}
		if attr.IsComputed() != test.computed {
			t.Errorf("The '%s' field's >Computed< property should be: %t", test.field, test.computed)
		}
	})
}

func GetCdpRegionFromConfig() (string, error) {
	config := cdp.NewConfig()
	if err := config.LoadConfig(); err != nil {
		return "", err
	}
	return config.GetCdpRegion()
}
