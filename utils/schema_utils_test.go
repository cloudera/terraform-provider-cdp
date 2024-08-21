// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package utils

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"testing"
)

func TestWithNoErrorAndAction(t *testing.T) {
	type testCase struct {
		name   string
		target map[string]schema.Attribute
		source map[string]schema.Attribute
	}
	for _, test := range []testCase{
		{
			name:   "Both inputs are nil",
			target: nil,
			source: nil,
		},
		{
			name:   "Target map is nil.",
			target: nil,
			source: map[string]schema.Attribute{},
		},
		{
			name:   "Source is nil.",
			target: map[string]schema.Attribute{},
			source: nil,
		},
		{
			name:   "Both are empty.",
			target: map[string]schema.Attribute{},
			source: map[string]schema.Attribute{},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			Append(test.target, test.source)

			if len(test.target) > 0 {
				t.Errorf("Target got extended eventhough it should've left untouched.")
			}
		})
	}
}

func TestAppendSchemaWhenNoOverlapThenSimpleCopyHappens(t *testing.T) {
	sourceKey := "some_other_key"
	target := map[string]schema.Attribute{
		"id": schema.StringAttribute{},
	}
	initialLength := len(target)
	source := map[string]schema.Attribute{
		sourceKey: schema.StringAttribute{},
	}

	Append(target, source)

	if len(target) != initialLength+len(source) {
		t.Errorf("Map did not get updated!")
	}
	if target["some_other_key"] == nil {
		t.Errorf("the new value with the key of '%s' did not get into the target map.", sourceKey)
	}
}

func TestAppendSchemaWhenOverlapThenOverwriteHappens(t *testing.T) {
	key := "keyValue"
	originalDescr := "original description"
	target := map[string]schema.Attribute{
		key: schema.StringAttribute{
			Description: originalDescr,
		},
	}
	initialLength := len(target)
	source := map[string]schema.Attribute{
		key: schema.StringAttribute{
			Description: "some other description",
		},
	}

	Append(target, source)

	if len(target) != initialLength {
		t.Errorf("Map got extended but should not be!")
	}
	if target[key].GetDescription() == originalDescr {
		t.Errorf("The target map did not get updated properly.")
	}
}
