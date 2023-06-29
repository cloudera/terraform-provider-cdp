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

import "testing"

func TestFailedStatusKeywords(t *testing.T) {
	expected := [2]string{"FAILED", "DELETED"}
	size := len(failedStatusKeywords)
	expectedSize := len(expected)
	if size != expectedSize {
		t.Errorf("The size of the failed status keywords is not the expected! Expected: %d, got: %d", expectedSize, size)
	}
	for _, expectedElement := range expected {
		found := false
		for _, keyword := range failedStatusKeywords {
			if keyword == expectedElement {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("The given element is not present in the failedStatusKeywords list: %s", expectedElement)
		}
	}
}
