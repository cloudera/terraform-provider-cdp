// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package environments

import (
	"context"
	"math"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

type tagKV struct {
	Key   string
	Value string
}

func convertTagsToKV(input types.Map) []tagKV {
	result := ConvertTags(context.TODO(), input)
	if result == nil {
		return nil
	}
	kvs := make([]tagKV, len(result))
	for i, tag := range result {
		kvs[i] = tagKV{Key: *tag.Key, Value: *tag.Value}
	}
	return kvs
}

func convertGcpTagsToKV(input types.Map) []tagKV {
	result := ConvertGcpTags(context.TODO(), input)
	if result == nil {
		return nil
	}
	kvs := make([]tagKV, len(result))
	for i, tag := range result {
		kvs[i] = tagKV{Key: *tag.Key, Value: *tag.Value}
	}
	return kvs
}

func testConvertTagsNilReturnsNil(t *testing.T, name string, convertFn func(types.Map) []tagKV) {
	t.Run(name, func(t *testing.T) {
		result := convertFn(types.MapNull(types.StringType))
		assert.Nil(t, result)
	})
}

func testConvertTagsEmptyReturnsNil(t *testing.T, name string, convertFn func(types.Map) []tagKV) {
	t.Run(name, func(t *testing.T) {
		inMap, _ := types.MapValue(types.StringType, map[string]attr.Value{})
		result := convertFn(inMap)
		assert.Nil(t, result)
	})
}

func testConvertTagsSingleElement(t *testing.T, name string, convertFn func(types.Map) []tagKV) {
	t.Run(name, func(t *testing.T) {
		inMap, _ := types.MapValue(types.StringType, map[string]attr.Value{
			"someKey-1": types.StringValue("someValue-1"),
		})
		result := convertFn(inMap)
		assert.Len(t, result, 1)
		assert.Equal(t, "someKey-1", result[0].Key)
		assert.Equal(t, "someValue-1", result[0].Value)
	})
}

func testConvertTagsMultipleElements(t *testing.T, name string, convertFn func(types.Map) []tagKV) {
	t.Run(name, func(t *testing.T) {
		inMap, _ := types.MapValue(types.StringType, map[string]attr.Value{
			"someKey-1": types.StringValue("someValue-1"),
			"someKey-2": types.StringValue("someValue-2"),
		})
		result := convertFn(inMap)
		assert.Len(t, result, 2)

		tagMap := make(map[string]string)
		for _, kv := range result {
			tagMap[kv.Key] = kv.Value
		}
		assert.Equal(t, "someValue-1", tagMap["someKey-1"])
		assert.Equal(t, "someValue-2", tagMap["someKey-2"])
	})
}

func testConvertTagsPreservesKeyValuePairing(t *testing.T, name string, convertFn func(types.Map) []tagKV) {
	t.Run(name, func(t *testing.T) {
		inMap, _ := types.MapValue(types.StringType, map[string]attr.Value{
			"env":     types.StringValue("prod"),
			"team":    types.StringValue("platform"),
			"project": types.StringValue("cdp"),
		})
		result := convertFn(inMap)
		assert.Len(t, result, 3)

		tagMap := make(map[string]string)
		for _, kv := range result {
			tagMap[kv.Key] = kv.Value
		}
		assert.Equal(t, "prod", tagMap["env"])
		assert.Equal(t, "platform", tagMap["team"])
		assert.Equal(t, "cdp", tagMap["project"])
	})
}

func TestConvertTags(t *testing.T) {
	testConvertTagsNilReturnsNil(t, "nil input returns nil", convertTagsToKV)
	testConvertTagsEmptyReturnsNil(t, "empty input returns nil", convertTagsToKV)
	testConvertTagsSingleElement(t, "single element", convertTagsToKV)
	testConvertTagsMultipleElements(t, "multiple elements", convertTagsToKV)
	testConvertTagsPreservesKeyValuePairing(t, "preserves key-value pairing", convertTagsToKV)
}

func TestConvertGcpTags(t *testing.T) {
	testConvertTagsNilReturnsNil(t, "nil input returns nil", convertGcpTagsToKV)
	testConvertTagsEmptyReturnsNil(t, "empty input returns nil", convertGcpTagsToKV)
	testConvertTagsSingleElement(t, "single element", convertGcpTagsToKV)
	testConvertTagsMultipleElements(t, "multiple elements", convertGcpTagsToKV)
	testConvertTagsPreservesKeyValuePairing(t, "preserves key-value pairing", convertGcpTagsToKV)
}

func TestGetStringValueIfNotEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected types.String
	}{
		{"empty string", "", types.StringNull()},
		{"normal content", "arm", types.StringValue("arm")},
		{"spaces only", "   ", types.StringNull()},
		{"leading/trailing spaces", "  hello  ", types.StringValue("hello")},
		{"tab only", "\t", types.StringNull()},
		{"newline only", "\n", types.StringNull()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, getStringValueIfNotEmpty(tt.input))
		})
	}
}

func TestSafeIntToInt32_ValidValues(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int32
	}{
		{"zero", 0, 0},
		{"positive", 42, 42},
		{"negative", -100, -100},
		{"max int32", math.MaxInt32, math.MaxInt32},
		{"min int32", math.MinInt32, math.MinInt32},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safeIntToInt32(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSafeIntToInt32_OutOfRange(t *testing.T) {
	tests := []struct {
		name  string
		input int
	}{
		{"overflow", math.MaxInt32 + 1},
		{"underflow", math.MinInt32 - 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safeIntToInt32(tt.input)
			assert.Error(t, err)
			assert.Equal(t, int32(0), result)
			assert.Contains(t, err.Error(), "out of int32 range")
		})
	}
}

func TestConvertIntToInt32IfPossible_ValidValues(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int32
	}{
		{"zero", 0, 0},
		{"positive", 123, 123},
		{"negative", -456, -456},
		{"max int32", math.MaxInt32, math.MaxInt32},
		{"min int32", math.MinInt32, math.MinInt32},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertIntToInt32IfPossible(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, types.Int32Value(tt.expected), result)
		})
	}
}

func TestConvertIntToInt32IfPossible_OutOfRange(t *testing.T) {
	tests := []struct {
		name  string
		input int
	}{
		{"overflow", math.MaxInt32 + 1},
		{"underflow", math.MinInt32 - 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ConvertIntToInt32IfPossible(tt.input)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "out of int32 range")
		})
	}
}
