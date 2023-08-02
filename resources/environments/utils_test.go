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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"golang.org/x/net/context"
	"testing"
)

func TestConvertTagsWhenInputIsNil(t *testing.T) {
	if ConvertTags(context.TODO(), types.MapNull(types.StringType)) != nil {
		t.Error("Result list is not nil but it should be!")
	}
}

func TestConvertTagsWhenInputIsEmpty(t *testing.T) {
	inMap, _ := types.MapValue(types.StringType, map[string]attr.Value{})
	if ConvertTags(context.TODO(), inMap) != nil {
		t.Error("Result list is not nil but it should be!")
	}
}

func TestConvertTagsWhenInputIsNotEmpty(t *testing.T) {
	key, value := "someKey-1", "someValue-1"
	inMap, _ := types.MapValue(types.StringType, map[string]attr.Value{key: types.StringValue(value)})
	result := ConvertTags(context.TODO(), inMap)

	if len(result) != 1 {
		t.Errorf("After tag conversion, not the expected amount came back. Expected: %d, got: %d.",
			1, len(result))
	}
	for _, tag := range result {
		if *tag.Key != key {
			t.Errorf("The provided key (%s) is not present in the result map!", key)
		}
		if *tag.Value != value {
			t.Errorf("The provided value (%s) is not present in the result map!", value)
		}
	}
}

func TestConvertTagsWhenInputHasMoreElements(t *testing.T) {
	keys := []string{"someKey-1", "someKey-2"}
	values := []string{"someValue-1", "someValue-2"}
	inMap, _ := types.MapValue(types.StringType, map[string]attr.Value{
		keys[0]: types.StringValue(values[0]),
		keys[1]: types.StringValue(values[1]),
	})
	result := ConvertTags(context.TODO(), inMap)

	if len(result) != len(keys) {
		t.Errorf("After tag conversion, not the expected amount came back. Expected: %d, got: %d.",
			len(keys), len(result))
	}

	var has bool
	for _, inputKey := range keys {
		has = false
		for _, tag := range result {
			if *tag.Key == inputKey {
				has = true
				break
			}
		}
		if !has {
			t.Errorf("The following key is not present in the result: %s", inputKey)
		}
	}
	for _, inputValue := range values {
		has = false
		for _, tag := range result {
			if *tag.Value == inputValue {
				has = true
				break
			}
		}
		if !has {
			t.Errorf("The following value is not present in the result: %s", inputValue)
		}
	}
}

func TestConvertGcpTagsWhenInputIsNil(t *testing.T) {
	if ConvertGcpTags(context.TODO(), types.MapNull(types.StringType)) != nil {
		t.Error("Result list is not nil but it should be!")
	}
}

func TestConvertGcpTagsWhenInputIsEmpty(t *testing.T) {
	inMap, _ := types.MapValue(types.StringType, map[string]attr.Value{})
	if ConvertGcpTags(context.TODO(), inMap) != nil {
		t.Error("Result list is not nil but it should be!")
	}
}

func TestConvertGcpTagsWhenInputIsNotEmpty(t *testing.T) {
	key, value := "someKey-1", "someValue-1"
	inMap, _ := types.MapValue(types.StringType, map[string]attr.Value{key: types.StringValue(value)})
	result := ConvertGcpTags(context.TODO(), inMap)

	if len(result) != 1 {
		t.Errorf("After tag conversion, not the expected amount came back. Expected: %d, got: %d.",
			1, len(result))
	}
	for _, tag := range result {
		if *tag.Key != key {
			t.Errorf("The provided key (%s) is not present in the result map!", key)
		}
		if *tag.Value != value {
			t.Errorf("The provided value (%s) is not present in the result map!", value)
		}
	}
}

func TestConvertGcpTagsWhenInputHasMoreElements(t *testing.T) {
	keys := []string{"someKey-1", "someKey-2"}
	values := []string{"someValue-1", "someValue-2"}
	inMap, _ := types.MapValue(types.StringType, map[string]attr.Value{
		keys[0]: types.StringValue(values[0]),
		keys[1]: types.StringValue(values[1]),
	})
	result := ConvertGcpTags(context.TODO(), inMap)

	if len(result) != len(keys) {
		t.Errorf("After tag conversion, not the expected amount came back. Expected: %d, got: %d.",
			len(keys), len(result))
	}

	var has bool
	for _, inputKey := range keys {
		has = false
		for _, tag := range result {
			if *tag.Key == inputKey {
				has = true
				break
			}
		}
		if !has {
			t.Errorf("The following key is not present in the result: %s", inputKey)
		}
	}
	for _, inputValue := range values {
		has = false
		for _, tag := range result {
			if *tag.Value == inputValue {
				has = true
				break
			}
		}
		if !has {
			t.Errorf("The following value is not present in the result: %s", inputValue)
		}
	}
}
