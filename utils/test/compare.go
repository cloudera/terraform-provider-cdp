// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func CompareStringSlices(actual, expected []string) (unexpected, missing []string) {
	for _, res := range actual {
		if !containsString(expected, res) {
			unexpected = append(unexpected, res)
		}
	}
	for _, expRes := range expected {
		if !containsString(actual, expRes) {
			missing = append(missing, expRes)
		}
	}
	return
}

func CompareStrings(got string, expected string, t *testing.T) {
	if got != expected {
		t.Errorf("Assertion error! Expected: %s, got: %s", expected, got)
	}
}

func containsString(resources []string, target string) bool {
	for _, res := range resources {
		if res == target {
			return true
		}
	}
	return false
}

func ToStringSliceFunc[T any](elements []T, f func(T) string) []string {
	var result []string
	for _, e := range elements {
		result = append(result, f(e))
	}
	return result
}

func CompareStringValueSlices(got []string, expected []attr.Value, t *testing.T) {
	if len(got) != len(expected) {
		t.Errorf("Assertion error! Expected length: %d, got length: %d", len(expected), len(got))
		return
	}

	for i, exp := range expected {
		if got[i] != exp.(types.String).ValueString() {
			t.Errorf("Assertion error! Expected: %s, got: %s", expected, got)
		}
	}
}

func CompareInt32PointerToTypesInt32(got *int32, expected types.Int32, t *testing.T) {
	if *got != expected.ValueInt32() {
		t.Errorf("Assertion error! Expected: %d, got: %d", expected.ValueInt32(), *got)
	}
}

func CompareInt64PointerToTypesInt64(got *int64, expected types.Int64, t *testing.T) {
	if *got != expected.ValueInt64() {
		t.Errorf("Assertion error! Expected: %d, got: %d", expected.ValueInt64(), *got)
	}
}

func CompareInts(got int, expected int, t *testing.T) {
	if got != expected {
		t.Errorf("Assertion error! Expected: %d, got: %d", expected, got)
	}
}

func CompareBools(got bool, expected bool, t *testing.T) {
	if got != expected {
		t.Errorf("Assertion error! Expected: %t, got: %t", expected, got)
	}
}
