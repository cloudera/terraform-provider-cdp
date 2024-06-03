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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTo(t *testing.T) {
	intVal := 10
	assert.Equal(t, intVal, *To(intVal))

	stringValVal := "str"
	assert.Equal(t, stringValVal, *To(stringValVal))

	type myStruct struct {
		x int
		y int
	}
	assert.Equal(t, myStruct{1, 2}, *To(myStruct{1, 2}))
}

func TestDerefDefault(t *testing.T) {
	intVal := 10
	assert.Equal(t, 200, DerefDefault(nil, 200))
	assert.Equal(t, 10, DerefDefault(&intVal, 200))

	stringValVal := "str"
	assert.Equal(t, "default", DerefDefault(nil, "default"))
	assert.Equal(t, "str", DerefDefault(&stringValVal, "default"))

	type myType struct {
		x int
		y int
	}

	structVal := myType{1, 1}
	assert.Equal(t, myType{2, 3}, DerefDefault(nil, myType{2, 3}))
	assert.Equal(t, myType{1, 1}, DerefDefault(&structVal, myType{2, 3}))
}

func TestCopy(t *testing.T) {
	orig := &[]bool{true}[0]
	cp := Copy(orig)

	*orig = false
	assert.True(t, *cp)
}

func TestCopyNil(t *testing.T) {
	var input *bool = nil
	assert.Nil(t, Copy(input))
}
