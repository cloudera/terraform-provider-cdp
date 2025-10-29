// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package utils

import "github.com/hashicorp/terraform-plugin-framework/types"

func StringArrayToSlice(input []types.String) []string {
	var result []string
	if len(input) > 0 {
		for _, az := range input {
			result = append(result, az.ValueString())
		}
	}
	return result
}
