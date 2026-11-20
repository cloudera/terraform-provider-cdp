// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package platformtypes

import (
	"testing"
)

func TestStringSliceToList(t *testing.T) {
	if r := StringSliceToList(nil); !r.IsNull() {
		t.Error("nil should give null list")
	}

	r := StringSliceToList([]string{})
	if r.IsNull() || len(r.Elements()) != 0 {
		t.Error("empty slice should give empty (non-null) list")
	}

	r = StringSliceToList([]string{"a", "b"})
	if len(r.Elements()) != 2 {
		t.Fatalf("got %d elements, want 2", len(r.Elements()))
	}
}

func TestStringValueOrNull(t *testing.T) {
	if r := StringValueOrNull(""); !r.IsNull() {
		t.Error("empty string should be null")
	}
	if r := StringValueOrNull("hello"); r.ValueString() != "hello" {
		t.Errorf("got %s, want hello", r.ValueString())
	}
}

func TestStringMapToTFMap(t *testing.T) {
	if r := StringMapToTFMap(nil); !r.IsNull() {
		t.Error("nil should give null map")
	}

	r := StringMapToTFMap(map[string]string{"a": "1"})
	if r.IsNull() {
		t.Fatal("non-nil input should give non-null map")
	}
	if len(r.Elements()) != 1 {
		t.Errorf("got %d elements, want 1", len(r.Elements()))
	}
}
