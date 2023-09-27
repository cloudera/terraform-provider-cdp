// // Copyright (c) HashiCorp, Inc.
// // SPDX-License-Identifier: MPL-2.0
package names

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

func TestProviderPackageForAlias(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		TestName string
		Input    string
		Expected string
		Error    bool
	}{
		{
			TestName: "empty",
			Input:    "",
			Expected: "",
			Error:    true,
		},
		{
			TestName: "IAM",
			Input:    "iam",
			Expected: "iam",
			Error:    false,
		},
		{
			TestName: "cml",
			Input:    "cml",
			Expected: "ml",
			Error:    false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.TestName, func(t *testing.T) {
			t.Parallel()

			got, err := ProviderPackageForAlias(testCase.Input)

			if err != nil && !testCase.Error {
				t.Errorf("got error (%s), expected no error", err)
			}

			if err == nil && testCase.Error {
				t.Errorf("got (%s) and no error, expected error", got)
			}

			if got != testCase.Expected {
				t.Errorf("got %s, expected %s", got, testCase.Expected)
			}
		})
	}
}

func TestServicesForDirectories(t *testing.T) {
	t.Parallel()

	for _, testCase := range ProviderPackages() {
		testCase := testCase
		t.Run(testCase, func(t *testing.T) {
			t.Parallel()

			wd, err := os.Getwd()
			if err != nil {
				t.Errorf("error reading working directory: %s", err)
			}

			if _, err := os.Stat(fmt.Sprintf("%s/../resources/%s", wd, testCase)); errors.Is(err, fs.ErrNotExist) {
				t.Errorf("expected %s/../resources/%s to exist %s", wd, testCase, err)
			}
		})
	}
}

func TestProviderNameUpper(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		TestName string
		Input    string
		Expected string
		Error    bool
	}{
		{
			TestName: "empty",
			Input:    "",
			Expected: "",
			Error:    true,
		},
		{
			TestName: "iam",
			Input:    "iam",
			Expected: "IAM",
			Error:    false,
		},
		{
			TestName: "cml",
			Input:    "ml",
			Expected: "MachineLearning",
			Error:    false,
		},
		{
			TestName: "doesnotexist",
			Input:    "doesnotexist",
			Expected: "",
			Error:    true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.TestName, func(t *testing.T) {
			t.Parallel()

			got, err := ProviderNameUpper(testCase.Input)

			if err != nil && !testCase.Error {
				t.Errorf("got error (%s), expected no error", err)
			}

			if err == nil && testCase.Error {
				t.Errorf("got (%s) and no error, expected error", got)
			}

			if got != testCase.Expected {
				t.Errorf("got %s, expected %s", got, testCase.Expected)
			}
		})
	}
}

func TestFullHumanFriendly(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		TestName string
		Input    string
		Expected string
		Error    bool
	}{
		{
			TestName: "empty",
			Input:    "",
			Expected: "",
			Error:    true,
		},
		{
			TestName: "iam",
			Input:    "iam",
			Expected: "Cloudera IAM Service",
			Error:    false,
		},
		{
			TestName: "cml",
			Input:    "ml",
			Expected: "Cloudera Machine Learning",
			Error:    false,
		},
		{
			TestName: "doesnotexist",
			Input:    "doesnotexist",
			Expected: "",
			Error:    true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.TestName, func(t *testing.T) {
			t.Parallel()

			got, err := FullHumanFriendly(testCase.Input)

			if err != nil && !testCase.Error {
				t.Errorf("got error (%s), expected no error", err)
			}

			if err == nil && testCase.Error {
				t.Errorf("got (%s) and no error, expected error", got)
			}

			if got != testCase.Expected {
				t.Errorf("got %s, expected %s", got, testCase.Expected)
			}
		})
	}
}

func TestCDPGoPackage(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		TestName string
		Input    string
		Expected string
		Error    bool
	}{
		{
			TestName: "empty",
			Input:    "",
			Expected: "",
			Error:    true,
		},
		{
			TestName: "iam",
			Input:    "iam",
			Expected: "iam",
			Error:    false,
		},
		{
			TestName: "ml",
			Input:    "ml",
			Expected: "ml",
			Error:    false,
		},
		{
			TestName: "doesnotexist",
			Input:    "doesnotexist",
			Expected: "",
			Error:    true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.TestName, func(t *testing.T) {
			t.Parallel()

			got, err := CDPGoPackage(testCase.Input)

			if err != nil && !testCase.Error {
				t.Errorf("got error (%s), expected no error", err)
			}

			if err == nil && testCase.Error {
				t.Errorf("got (%s) and no error, expected error", got)
			}

			if got != testCase.Expected {
				t.Errorf("got %s, expected %s", got, testCase.Expected)
			}
		})
	}
}
