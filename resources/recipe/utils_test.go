// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package recipe

import (
	"os"
	"testing"
)

func TestProcessInputWithRegularText(t *testing.T) {
	input := "regular text"
	output, err := processInput(input)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if output != input {
		t.Fatalf("Expected %v, got %v", input, output)
	}
}

func TestProcessInputWithExistingFile(t *testing.T) {
	content := "file content"
	file, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer func(name string) {
		_ = os.Remove(name)
	}(file.Name())

	_, err = file.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	_ = file.Close()

	output, err := processInput(file.Name())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if output != content {
		t.Fatalf("Expected %v, got %v", content, output)
	}
}

func TestProcessInputWithNonExistingPath(t *testing.T) {
	input := "/non/existing/file/path"
	output, err := processInput(input)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if output != input {
		t.Fatalf("Expected %v, got %v", input, output)
	}
}

func TestProcessInputWithDirectory(t *testing.T) {
	dir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(dir)

	_, err = processInput(dir)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestIsPathWithExistingFile(t *testing.T) {
	file, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer func(name string) {
		_ = os.Remove(name)
	}(file.Name())

	if !isPath(file.Name()) {
		t.Fatalf("Expected true for existing file, got false")
	}
}

func TestIsPathWithNonExistingFile(t *testing.T) {
	if isPath("/non/existing/file/path") {
		t.Fatalf("Expected false for non-existing file, got true")
	}
}

func TestIsPathWithExistingDirectory(t *testing.T) {
	dir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(dir)

	if !isPath(dir) {
		t.Fatalf("Expected true for existing directory, got false")
	}
}

func TestIsPathWithEmptyString(t *testing.T) {
	if isPath("") {
		t.Fatalf("Expected false for empty string, got true")
	}
}
