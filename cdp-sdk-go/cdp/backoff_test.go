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
	"time"
)

func TestLinearBackoffWithPositiveRetries(t *testing.T) {
	retries := 3
	step := 2
	expected := 8 * time.Second

	result := linearBackoff(retries, step)
	if result != expected {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

func TestLinearBackoffWithZeroRetries(t *testing.T) {
	retries := 0
	step := 2
	expected := 2 * time.Second

	result := linearBackoff(retries, step)
	if result != expected {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

func TestLinearBackoffWithNegativeRetries(t *testing.T) {
	retries := -1
	step := 2
	expected := 0 * time.Second

	result := linearBackoff(retries, step)
	if result != expected {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

func TestLinearBackoffWithLargeStep(t *testing.T) {
	retries := 2
	step := 1000
	expected := 3000 * time.Second

	result := linearBackoff(retries, step)
	if result != expected {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

func TestExponentialBackoffWithPositiveRetries(t *testing.T) {
	retries := 3
	expectedMin := 6 * time.Second
	expectedMax := 8 * time.Second

	result := exponentialBackoff(retries)
	if result < expectedMin || result > expectedMax {
		t.Fatalf("Expected between %v and %v, got %v", expectedMin, expectedMax, result)
	}
}

func TestExponentialBackoffWithZeroRetries(t *testing.T) {
	retries := 0
	expectedMin := 0.75 * float64(time.Second)
	expectedMax := 1 * time.Second

	result := exponentialBackoff(retries)
	if result < time.Duration(int(expectedMin)) || result > expectedMax {
		t.Fatalf("Expected between %v and %v, got %v", time.Duration(int(expectedMin)), expectedMax, result)
	}
}

func TestExponentialBackoffWithHighRetries(t *testing.T) {
	retries := 10
	expectedMin := 768 * time.Second
	expectedMax := 1024 * time.Second

	result := exponentialBackoff(retries)
	if result < expectedMin || result > expectedMax {
		t.Fatalf("Expected between %v and %v, got %v", expectedMin, expectedMax, result)
	}
}

func TestLinearBackoffStrategyWithDefaultStep(t *testing.T) {
	_ = os.Setenv("CDP_TF_BACKOFF_STRATEGY", "linear")
	defer func() {
		_ = os.Unsetenv("CDP_TF_BACKOFF_STRATEGY")
	}()

	result := backoff(2)
	expected := defaultLinearBackoffStep * 3 * time.Second
	if result != expected {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

func TestExponentialBackoffStrategyWithDefault(t *testing.T) {
	_ = os.Unsetenv("CDP_TF_BACKOFF_STRATEGY")

	result := backoff(3)
	expectedMin := 6 * time.Second
	expectedMax := 8 * time.Second
	if result < expectedMin || result > expectedMax {
		t.Fatalf("Expected between %v and %v, got %v", expectedMin, expectedMax, result)
	}
}

func TestLinearBackoffStrategyWithCustomStep(t *testing.T) {
	_ = os.Setenv("CDP_TF_BACKOFF_STRATEGY", "linear")
	_ = os.Setenv("CDP_TF_BACKOFF_STEP", "5")
	defer func() {
		_ = os.Unsetenv("CDP_TF_BACKOFF_STRATEGY")
		_ = os.Unsetenv("CDP_TF_BACKOFF_STEP")
	}()

	result := backoff(1)
	expected := 10 * time.Second
	if result != expected {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

func TestExponentialBackoffStrategyWithHighRetries(t *testing.T) {
	_ = os.Unsetenv("CDP_TF_BACKOFF_STRATEGY")

	result := backoff(10)
	expectedMin := 768 * time.Second
	expectedMax := 1024 * time.Second
	if result < expectedMin || result > expectedMax {
		t.Fatalf("Expected between %v and %v, got %v", expectedMin, expectedMax, result)
	}
}
