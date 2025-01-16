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
	"errors"
	"os"
)

func processInput(input string) (string, error) {
	// Check if the path exists
	info, err := os.Stat(input)
	if err != nil {
		// Input is not a valid path, so treat it as regular text
		return input, nil
	}

	// Check if the path points to a regular file
	if info.Mode().IsRegular() {
		// Read the file's content
		content, err := os.ReadFile(input)
		if err != nil {
			return "", err // Return an error if reading fails
		}
		return string(content), nil
	}

	// Path exists but is not a regular file
	return "", errors.New("path exists but is not a regular file")
}

func isPath(input string) bool {
	_, err := os.Stat(input)
	return err == nil
}
