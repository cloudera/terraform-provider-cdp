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
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

var (
	skippedPrefixes = []string{".", "_"}
	skippedPaths    = []string{"/testdata/", "/gen/", "/dist/"}

	copyrightRe = regexp.MustCompile(`[[:graph:]]+ Copyright \d{4} Cloudera\. All Rights Reserved\.\s+[[:graph:]]+\s+[[:graph:]]+ This file is licensed under the Apache License Version 2\.0 \(the "License"\)\.\s+[[:graph:]]+ You may not use this file except in compliance with the License.\s+[[:graph:]]+ You may obtain a copy of the License at http:\/\/www\.apache\.org\/licenses\/LICENSE-2\.0\.\s+[[:graph:]]+\s+[[:graph:]]+ This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS\s+[[:graph:]]+ OF ANY KIND, either express or implied. Refer to the License for the specific\s+[[:graph:]]+ permissions and limitations governing your use of the file\.`)
)

func TestAllLicenseHeaders(t *testing.T) {
	path, err := filepath.Abs("../")
	if err != nil {
		t.Errorf("Failed with err: %v", err)
	}
	files, err := checkHeader(path)
	if err != nil {
		t.Errorf("Failed with err: %v", err)
	}
	if len(files) > 0 {
		t.Fatalf("Found files without copyright header: %v. To fix the test, copy the header from a similar file"+
			"and fix the year.", files)
	}
}

func checkHeader(dir string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			for _, prefix := range skippedPrefixes {
				if strings.HasPrefix(d.Name(), prefix) {
					return filepath.SkipDir
				}
			}
			return nil
		}
		normalized := strings.TrimPrefix(filepath.ToSlash(path), filepath.ToSlash(dir))
		for _, skippedPath := range skippedPaths {
			if strings.Contains(normalized, skippedPath) {
				return filepath.SkipDir
			}
		}

		if shouldIncludeHeader(path) {
			needsHeader, err := checkFile(path)
			if err != nil {
				return err
			}
			if needsHeader {
				files = append(files, path)
			}
		}

		return nil
	})
	return files, err
}

func shouldIncludeHeader(path string) bool {
	return strings.HasSuffix(path, ".go") ||
		strings.HasSuffix(path, ".tf") ||
		strings.HasSuffix(path, ".yaml") ||
		strings.HasSuffix(path, "Makefile")
}

func checkFile(filename string) (bool, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return false, err
	}

	return !copyrightRe.MatchString(string(content)), nil
}
