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

import (
	"context"
	"log"
	"os"
	"runtime"
	"testing"

	"github.com/hectane/go-acl"
)

const FILE_NO_PERMISSIOM = os.FileMode(000)

func TestWhenFileDoesNotExists(t *testing.T) {
	content, err := ReadFileContent(context.TODO(), "someNonExistingStuff")
	checkFailure(t, content, err)
}

func TestReadFileContentWhenNoPermissionToRead(t *testing.T) {
	file, err := os.CreateTemp(os.TempDir(), "gcp_read_temp")
	if err != nil {
		t.Errorf("Unable to prepare temprary file for testing due to: " + err.Error())
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Default().Printf("unable to clean up file ('%s') due to: %s", file.Name(), err.Error())
		}
	}(file)
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			log.Default().Printf("unable to clean up file ('%s') due to: %s", file.Name(), err.Error())
		}
	}(file.Name())
	if runtime.GOOS == `windows` {
		err = acl.Chmod(file.Name(), FILE_NO_PERMISSIOM)
	} else {
		err = os.Chmod(file.Name(), FILE_NO_PERMISSIOM)
	}
	if err != nil {
		t.Errorf("Unable to update temprary file's permission for testing due to: " + err.Error())
	}

	content, err := ReadFileContent(context.TODO(), file.Name())

	checkFailure(t, content, err)
}

func TestReadFileContentWhenFileExistsAndHavePermission(t *testing.T) {
	file, err := os.CreateTemp(os.TempDir(), "gcp_read_temp")
	if err != nil {
		t.Errorf("Unable to prepare temprary file for testing due to: " + err.Error())
	}
	originalContent := []byte(`{"some":"amazing","content":true}`)
	if _, err = file.Write(originalContent); err != nil {
		t.Errorf("Unable to update temp file ('%s') content due to: %s\n", file.Name(), err.Error())
	}

	resultContent, err := ReadFileContent(context.TODO(), file.Name())

	if err != nil {
		t.Errorf("File read failed due to: %s", err.Error())
	}
	originalContentToCompare := string(originalContent)
	if *resultContent != originalContentToCompare {
		t.Errorf("After file read it did not return the expected content! Expected: %s, got: %s", originalContentToCompare, *resultContent)
	}
}

func checkFailure(t *testing.T, content *string, readError error) {
	if readError == nil {
		t.Error("Expected read failure did not happen!")
	}
	if content != nil {
		t.Error("Content should not be filled when error happen!")
	}
}
