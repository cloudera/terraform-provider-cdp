// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package ptr

// To creates a pointer value from `v`
func To[T any](v T) *T {
	return &v
}

// DerefDefault dereferences the `v` value if not nil otherwise it returns nil
func DerefDefault[T any](v *T, def T) T {
	if v == nil {
		return def
	}
	return *v
}

// Copy creates a copy of the pointer which will point to the same memory address
func Copy[T any](v *T) *T {
	if v == nil {
		return nil
	}

	cp := *v
	return &cp
}
