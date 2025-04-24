// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datalake

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/cloudera/terraform-provider-cdp/utils/test"
)

var (
	commonElementCaseSet = []test.ResourceSchemaTestCaseStructure{
		{
			Name:             "'id' field must exist",
			Field:            "id",
			Computed:         true,
			ShouldBeRequired: false,
			AttributeType:    schema.StringAttribute{},
		},
		{
			Name:             "'polling_options' should exist",
			Field:            "polling_options",
			Computed:         false,
			ShouldBeRequired: false,
			AttributeType:    schema.SingleNestedAttribute{},
		},
		{
			Name:             "'creation_date' should exist",
			Field:            "creation_date",
			Computed:         true,
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
		{
			Name:             "'datalake_name' must exist",
			Field:            "datalake_name",
			Computed:         false,
			ShouldBeRequired: true,
			AttributeType:    schema.StringAttribute{},
		},
		{
			Name:             "'enable_ranger_raz' should exist",
			Field:            "enable_ranger_raz",
			Computed:         true,
			ShouldBeRequired: false,
			AttributeType:    schema.BoolAttribute{},
		},
		{
			Name:             "'environment_crn' should exist",
			Field:            "environment_crn",
			Computed:         true,
			ShouldBeRequired: false,
			AttributeType:    schema.StringAttribute{},
		},
		{
			Name:             "'environment_name' must exist",
			Field:            "environment_name",
			Computed:         false,
			ShouldBeRequired: true,
			AttributeType:    schema.StringAttribute{},
		},
		{
			Name:             "'image' should exist",
			Field:            "image",
			Computed:         false,
			ShouldBeRequired: false,
			AttributeType:    schema.SingleNestedAttribute{},
		},
		{
			Name:             "'java_version' should exist",
			Field:            "java_version",
			Computed:         false,
			ShouldBeRequired: false,
			AttributeType:    schema.Int32Attribute{},
		},
		{
			Name:             "'recipes' should exist",
			Field:            "recipes",
			Computed:         false,
			ShouldBeRequired: false,
			AttributeType:    schema.SetNestedAttribute{},
		},
		{
			Name:             "'runtime' should exist",
			Field:            "runtime",
			Computed:         false,
			ShouldBeRequired: false,
			AttributeType:    schema.StringAttribute{},
		},
		{
			Name:             "'scale' should exist",
			Field:            "scale",
			Computed:         true,
			ShouldBeRequired: false,
			AttributeType:    schema.StringAttribute{},
		},
		{
			Name:             "'status' should exist",
			Field:            "status",
			Computed:         true,
			ShouldBeRequired: false,
			AttributeType:    schema.StringAttribute{},
		},
		{
			Name:             "'status_reason' should exist",
			Field:            "status_reason",
			Computed:         true,
			ShouldBeRequired: false,
			AttributeType:    schema.StringAttribute{},
		},
		{
			Name:             "'multi_az' should exist",
			Field:            "multi_az",
			Computed:         true,
			ShouldBeRequired: false,
			AttributeType:    schema.BoolAttribute{},
		},
		{
			Name:             "'tags' should exist",
			Field:            "tags",
			Computed:         false,
			ShouldBeRequired: false,
			AttributeType:    schema.MapAttribute{},
		},
	}
)

func TestRootElements(t *testing.T) {
	SchemaContainsCommonElements(t, generalAttributes)
}

func SchemaContainsCommonElements(t *testing.T, providerSpecificSchema map[string]schema.Attribute) {
	for _, toTest := range commonElementCaseSet {
		t.Run(toTest.Name, func(t *testing.T) {
			if providerSpecificSchema[toTest.Field] == nil {
				t.Errorf("The following field does not exists, however it should: %s", toTest.Field)
				t.FailNow()
			}
			if providerSpecificSchema[toTest.Field].IsRequired() != toTest.ShouldBeRequired {
				t.Errorf("The '%s' field's >required< property should be: %t", toTest.Field, toTest.ShouldBeRequired)
			}
			if providerSpecificSchema[toTest.Field].IsComputed() != toTest.Computed {
				t.Errorf("The '%s' field's >Computed< property should be: %t", toTest.Field, toTest.Computed)
			}
			var currentType = reflect.TypeOf(providerSpecificSchema[toTest.Field])
			var expectedType = reflect.TypeOf(toTest.AttributeType)
			if currentType != expectedType {
				t.Errorf("The '%s' field's type should be: %t, instead of %t", toTest.Field, expectedType, currentType)
			}
		})
	}
}
