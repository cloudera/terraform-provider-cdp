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
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/common"
)

func TestConcat(t *testing.T) {
	common.AssertEquals(t, Concat(""), "")
	common.AssertEquals(t, Concat("", "a"), "a")
	common.AssertEquals(t, Concat("a", " ", "b"), "a b")
}

func TestToSetValueFromStringList(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "Convert non-empty string slice",
			input:    []string{"apple", "banana", "cherry"},
			expected: []string{"apple", "banana", "cherry"},
		},
		{
			name:     "Convert empty string slice",
			input:    []string{},
			expected: []string{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			setValue := ToSetValueFromStringList(tc.input)

			var result []string
			for _, elem := range setValue.Elements() {
				result = append(result, elem.(types.String).ValueString())
			}

			assert.ElementsMatch(t, tc.expected, result, "The set elements do not match expected values")
		})
	}
}
