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

func CompareStrings(actual, expected []string) (unexpected, missing []string) {
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
