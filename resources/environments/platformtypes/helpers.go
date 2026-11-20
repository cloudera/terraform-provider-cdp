// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package platformtypes

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func StringSliceToList(in []string) types.List {
	if in == nil {
		return types.ListNull(types.StringType)
	}
	vals := make([]attr.Value, len(in))
	for i, s := range in {
		vals[i] = types.StringValue(s)
	}
	return types.ListValueMust(types.StringType, vals)
}

func StringValueOrNull(s string) types.String {
	if s == "" {
		return types.StringNull()
	}
	return types.StringValue(s)
}

func StringMapToTFMap(in map[string]string) types.Map {
	if in == nil {
		return types.MapNull(types.StringType)
	}
	return types.MapValueMust(types.StringType, stringMapToAttrValues(in))
}
