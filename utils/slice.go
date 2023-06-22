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

import "strings"

func ContainsAsSubstring(slice []string, element string) bool {
	if len(slice) > 0 {
		for _, e := range slice {
			if strings.Contains(element, e) {
				return true
			}
		}
	}
	return false
}

func ContainsEitherSubstring(slice []string, elements []string) bool {
	if len(slice) > 0 && len(elements) > 0 {
		for _, e := range slice {
			for _, substring := range elements {
				if strings.Contains(e, substring) {
					return true
				}
			}
		}
	}
	return false
}
