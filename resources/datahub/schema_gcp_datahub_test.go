// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datahub

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var exceptionalKeys = []string{"environment", "cluster_template", "cluster_definition"}

func TestGcpCommonAttributesExceptExceptionalOnes(t *testing.T) {
	commonsExceptExceptions := make([]TestCaseStructure, 0)
	for _, element := range commonElementCaseSet {
		exceptional := false
		for _, exception := range exceptionalKeys {
			if element.field == exception {
				exceptional = true
				break
			}
		}
		if !exceptional {
			commonsExceptExceptions = append(commonsExceptExceptions, element)
		}
	}

	underTestAttributes := createFilledGcpDatahubResourceTestObject()

	for _, test := range commonsExceptExceptions {
		t.Run(test.name, func(t *testing.T) {
			PerformValidation(t, test, underTestAttributes[test.field])
		})
	}
}

func TestGcpSchemaContainsGcpSpecificFields(t *testing.T) {
	cases := []TestCaseStructure{
		{
			name:             "cluster_template_name must exist",
			field:            "cluster_template_name",
			computed:         false,
			shouldBeRequired: false,
		},
		{
			name:             "cluster_definition_name must exist",
			field:            "cluster_definition_name",
			computed:         false,
			shouldBeRequired: false,
		},
		{
			name:             "environment_name must exist",
			field:            "environment_name",
			computed:         false,
			shouldBeRequired: true,
		},
	}

	underTestAttributes := createFilledGcpDatahubResourceTestObject()

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			PerformValidation(t, test, underTestAttributes[test.field])
		})
	}
}

func TestGcpSchemaContainsInstanceGroup(t *testing.T) {
	test := TestCaseStructure{
		name:             "instance_group should exist",
		field:            "instance_group",
		computed:         false,
		shouldBeRequired: false,
	}

	underTestAttributes := createFilledGcpDatahubResourceTestObject()

	PerformValidation(t, test, underTestAttributes[test.field])
}

func createFilledGcpDatahubResourceTestObject() map[string]schema.Attribute {
	res := &gcpDatahubResource{}
	schemaResponse := &resource.SchemaResponse{}
	res.Schema(context.TODO(), resource.SchemaRequest{}, schemaResponse)

	return schemaResponse.Schema.Attributes
}
