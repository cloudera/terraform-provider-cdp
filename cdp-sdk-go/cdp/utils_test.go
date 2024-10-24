// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package cdp

import (
	"os"
	"testing"
)

func TestContainsInt(t *testing.T) {
	type input struct {
		description   string
		content       []int
		target        int
		shouldContain bool
	}
	for _, scenario := range []input{
		{
			description:   "input is empty slice",
			content:       []int{},
			target:        1,
			shouldContain: false,
		},
		{
			description:   "input is not empty but does not contain target",
			content:       []int{1, 2},
			target:        3,
			shouldContain: false,
		},
		{
			description:   "input is not empty and contains target",
			content:       []int{1, 2, 3},
			target:        3,
			shouldContain: true,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			if result := ContainsInt(scenario.content, scenario.target); result != scenario.shouldContain {
				t.Errorf("The expected '%t' value does not match with the result: '%t'.", scenario.shouldContain, result)
			}
		})
	}
}

func ContainsInt(slice []int, element int) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func TestEnvironmentVariableIsSet(t *testing.T) {
	_ = os.Setenv("TEST_ENV_VAR", "42")
	defer func() {
		_ = os.Unsetenv("TEST_ENV_VAR")
	}()

	result := intFromEnvOrDefault("TEST_ENV_VAR", 10)
	expected := 42
	if result != expected {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

func TestEnvironmentVariableIsNotSet(t *testing.T) {
	_ = os.Unsetenv("TEST_ENV_VAR")

	result := intFromEnvOrDefault("TEST_ENV_VAR", 10)
	expected := 10
	if result != expected {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

func TestEnvironmentVariableIsInvalid(t *testing.T) {
	_ = os.Setenv("TEST_ENV_VAR", "invalid")
	defer func() {
		_ = os.Unsetenv("TEST_ENV_VAR")
	}()

	result := intFromEnvOrDefault("TEST_ENV_VAR", 10)
	expected := 10
	if result != expected {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

func TestSliceContainsElement(t *testing.T) {
	slice := []int{1, 2, 3}
	element := 2

	result := sliceContains(slice, element)
	if result != true {
		t.Fatalf("Expected %v, got %v", true, result)
	}
}

func TestSliceDoesNotContainElement(t *testing.T) {
	slice := []int{1, 2, 3}
	element := 4

	result := sliceContains(slice, element)
	if result != false {
		t.Fatalf("Expected %v, got %v", false, result)
	}
}

func TestSliceIsEmpty(t *testing.T) {
	var slice []int
	element := 1

	result := sliceContains(slice, element)
	if result != false {
		t.Fatalf("Expected %v, got %v", false, result)
	}
}
