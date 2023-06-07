package utils

import "testing"

func TestContainsAsSubstring(t *testing.T) {
	type input struct {
		description   string
		content       []string
		target        string
		shouldContain bool
	}
	for _, scenario := range []input{
		{
			description:   "input is nil",
			content:       nil,
			target:        "",
			shouldContain: false,
		},
		{
			description:   "input is empty slice",
			content:       []string{},
			target:        "",
			shouldContain: false,
		},
		{
			description:   "input is not empty but does not contain target",
			content:       []string{"one", "two"},
			target:        "three",
			shouldContain: false,
		},
		{
			description:   "input is not empty and contains target",
			content:       []string{"one", "two", "three"},
			target:        "three",
			shouldContain: true,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			if result := ContainsAsSubstring(scenario.content, scenario.target); result != scenario.shouldContain {
				t.Errorf("The expected '%t' value does not match with the result: '%t'.", scenario.shouldContain, result)
			}
		})
	}
}

func TestContainsEitherSubstring(t *testing.T) {
	type input struct {
		description   string
		input         []string
		elements      []string
		shouldContain bool
	}
	for _, scenario := range []input{
		{
			description:   "input and elements are nil",
			input:         nil,
			elements:      nil,
			shouldContain: false,
		},
		{
			description:   "input is empty slice while elements nil",
			input:         []string{},
			elements:      nil,
			shouldContain: false,
		},
		{
			description:   "input is nil and elements is empty",
			input:         nil,
			elements:      []string{},
			shouldContain: false,
		},
		{
			description:   "input is valid but elements nil",
			input:         []string{"one", "two"},
			elements:      nil,
			shouldContain: false,
		},
		{
			description:   "inpiut is valid but elements empty",
			input:         []string{"one", "two"},
			elements:      []string{},
			shouldContain: false,
		},
		{
			description:   "both input and elements valid but has no common element",
			input:         []string{"one", "two"},
			elements:      []string{"three", "four"},
			shouldContain: false,
		},
		{
			description:   "both input and elements valid and has one common element",
			input:         []string{"one", "two"},
			elements:      []string{"two", "three"},
			shouldContain: true,
		},
		{
			description:   "both input and elements valid and has multiple common elements",
			input:         []string{"one", "two"},
			elements:      []string{"one", "two", "three", "four"},
			shouldContain: true,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			if got := ContainsEitherSubstring(scenario.input, scenario.elements); got != scenario.shouldContain {
				t.Errorf("The expected '%t' value does not match with the result: '%t'.", scenario.shouldContain, got)
			}
		})
	}
}
