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

import (
	"fmt"
	"math"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func StringArrayToSlice(input []types.String) []string {
	var result []string
	if len(input) > 0 {
		for _, az := range input {
			result = append(result, az.ValueString())
		}
	}
	return result
}

func FromStringListToListValue(s []string) types.List {
	if s == nil {
		return types.ListNull(types.StringType)
	}

	var elems []attr.Value
	elems = make([]attr.Value, 0, len(s))
	for _, v := range s {
		elems = append(elems, types.StringValue(v))
	}
	var list types.List
	list, _ = types.ListValue(types.StringType, elems)
	return list
}

func FromListValueToStringList(tl types.List) []string {
	if tl.IsNull() {
		return []string{}
	}
	res := make([]string, len(tl.Elements()))
	for i, elem := range tl.Elements() {
		res[i] = elem.(types.String).ValueString()
	}
	return res
}

func FromTfStringSliceToStringSlice(tss []types.String) []string {
	result := make([]string, 0)
	for _, az := range tss {
		result = append(result, az.ValueString())
	}
	return result
}

func FromSetValueToStringList(tl types.Set) []string {
	if tl.IsNull() {
		return []string{}
	}
	res := make([]string, len(tl.Elements()))
	for i, elem := range tl.Elements() {
		res[i] = elem.(types.String).ValueString()
	}
	return res
}

func Int64To32Pointer(in types.Int64) *int32 {
	n32 := Int64To32(in)
	return &n32
}

func Int64To32(in types.Int64) int32 {
	n64 := in.ValueInt64()
	if n64 >= math.MinInt32 && n64 <= math.MaxInt32 {
		return int32(n64)
	}
	panic(fmt.Sprintf("int64 value %d is out of range for int32", n64))
}

func IsStringTypeHasValue(value types.String) bool {
	return !value.IsNull() && !value.IsUnknown() && len(strings.TrimSpace(value.ValueString())) > 0
}
