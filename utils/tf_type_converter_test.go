// Copyright 2025 Cloudera. All Rights Reserved.
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
	"slices"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestStringArrayToSlice(t *testing.T) {
	tests := []struct {
		name string
		in   []types.String
		want []string
	}{
		{
			name: "nil input",
			in:   nil,
			want: []string{},
		},
		{
			name: "empty input",
			in:   []types.String{},
			want: []string{},
		},
		{
			name: "single value",
			in:   []types.String{types.StringValue("foo")},
			want: []string{"foo"},
		},
		{
			name: "multiple values",
			in:   []types.String{types.StringValue("a"), types.StringValue("b"), types.StringValue("c")},
			want: []string{"a", "b", "c"},
		},
		{
			name: "includes null and unknown",
			in:   []types.String{types.StringNull(), types.StringUnknown(), types.StringValue("x")},
			want: []string{"", "", "x"},
		},
		{
			name: "includes empty string value",
			in:   []types.String{types.StringValue("")},
			want: []string{""},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := StringArrayToSlice(tc.in)
			if !slices.Equal(got, tc.want) {
				t.Fatalf("StringArrayToSlice(%v) = %v; want %v", tc.in, got, tc.want)
			}
		})
	}
}
