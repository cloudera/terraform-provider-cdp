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
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"os"
)

func ReadFileContent(ctx context.Context, path string) (*string, error) {
	tflog.Info(ctx, fmt.Sprintf("About to read file on path: %s", path))
	data, err := os.ReadFile(path)
	if err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Error occurred during file read: %s", err.Error()))
		return nil, err
	}
	tflog.Info(ctx, "Reading file was successful.")
	content := string(data)
	return &content, nil
}
