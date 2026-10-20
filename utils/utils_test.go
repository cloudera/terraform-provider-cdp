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
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestGetCallFailureThresholdOrDefault(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		threshold types.Int64
		fallback  int
		want      int
	}{
		{
			name:      "Null threshold returns fallback",
			threshold: types.Int64Null(),
			fallback:  3,
			want:      3,
		},
		{
			name:      "Unknown threshold returns fallback",
			threshold: types.Int64Unknown(),
			fallback:  3,
			want:      3,
		},
		{
			name:      "Null threshold and zero fallback returns 0",
			threshold: types.Int64Null(),
			fallback:  0,
			want:      0,
		},
		{
			name:      "Positive threshold value is returned",
			threshold: types.Int64Value(5),
			fallback:  3,
			want:      5,
		},
		{
			name:      "Zero threshold falls through to fallback",
			threshold: types.Int64Value(0),
			fallback:  3,
			want:      3,
		},
		{
			name:      "Negative threshold falls through to fallback",
			threshold: types.Int64Value(-1),
			fallback:  3,
			want:      3,
		},
		{
			name:      "Zero threshold and zero fallback returns 0",
			threshold: types.Int64Value(0),
			fallback:  0,
			want:      0,
		},
		{
			name:      "Negative fallback returns 0",
			threshold: types.Int64Null(),
			fallback:  -2,
			want:      0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetCallFailureThresholdOrDefault(tt.threshold, tt.fallback)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFromTfStringSliceToStringSliceIfEmpty(t *testing.T) {
	assert.Equal(t, 0, len(FromTfStringSliceToStringSlice([]types.String{})))
}

func TestFromTfStringSliceToStringSliceIfNotEmpty(t *testing.T) {
	val := "someValue"
	input := []types.String{
		types.StringValue(val),
	}

	result := FromTfStringSliceToStringSlice(input)

	assert.Equal(t, len(input), len(result))
	assert.Equal(t, val, result[0])
}

func TestFromStringListToListValue(t *testing.T) {
	assert.Equal(
		t,
		types.ListNull(types.StringType),
		FromStringListToListValue(nil),
	)

	expected, err := types.ListValue(types.StringType, []attr.Value{})
	assert.Nil(t, err)
	assert.Equal(
		t,
		expected,
		FromStringListToListValue([]string{}),
	)

	expected, err = types.ListValue(
		types.StringType,
		[]attr.Value{
			types.StringValue("test0"),
			types.StringValue("test1"),
		},
	)
	assert.Nil(t, err)
	assert.Equal(
		t,
		expected,
		FromStringListToListValue([]string{"test0", "test1"}),
	)
}

func TestResolveAsync(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		po            *PollingOptions
		topLevelAsync types.Bool
		want          bool
	}{
		{
			name:          "nil polling options uses top-level true",
			po:            nil,
			topLevelAsync: types.BoolValue(true),
			want:          true,
		},
		{
			name:          "nil polling options uses top-level false",
			po:            nil,
			topLevelAsync: types.BoolValue(false),
			want:          false,
		},
		{
			name:          "polling options async true overrides top-level false",
			po:            &PollingOptions{Async: types.BoolValue(true)},
			topLevelAsync: types.BoolValue(false),
			want:          true,
		},
		{
			name:          "polling options async false overrides top-level true",
			po:            &PollingOptions{Async: types.BoolValue(false)},
			topLevelAsync: types.BoolValue(true),
			want:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ResolveAsync(tt.po, tt.topLevelAsync)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestResolveTimeout(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	defaultTimeout := 60 * time.Minute
	getterTimeout := 45 * time.Minute

	getter := func(_ context.Context, d time.Duration) (time.Duration, diag.Diagnostics) {
		return d, nil
	}
	customGetter := func(_ context.Context, _ time.Duration) (time.Duration, diag.Diagnostics) {
		return getterTimeout, nil
	}

	tests := []struct {
		name    string
		po      *PollingOptions
		getter  func(context.Context, time.Duration) (time.Duration, diag.Diagnostics)
		want    time.Duration
		wantErr bool
	}{
		{
			name:   "nil polling options delegates to getter with default",
			po:     nil,
			getter: getter,
			want:   defaultTimeout,
		},
		{
			name:   "nil polling options uses getter return value",
			po:     nil,
			getter: customGetter,
			want:   getterTimeout,
		},
		{
			name:   "polling options with set timeout returns it in minutes",
			po:     &PollingOptions{PollingTimeout: types.Int64Value(90)},
			getter: customGetter,
			want:   90 * time.Minute,
		},
		{
			name:   "polling options with null timeout delegates to getter",
			po:     &PollingOptions{PollingTimeout: types.Int64Null()},
			getter: customGetter,
			want:   getterTimeout,
		},
		{
			name:   "polling options with unknown timeout delegates to getter",
			po:     &PollingOptions{PollingTimeout: types.Int64Unknown()},
			getter: customGetter,
			want:   getterTimeout,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, diags := ResolveTimeout(ctx, tt.po, tt.getter, defaultTimeout)
			assert.False(t, diags.HasError())
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestResolveCallFailureThreshold(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		po          *PollingOptions
		topLevelCFT types.Int64
		fallback    int
		want        int
	}{
		{
			name:        "nil polling options uses top-level value",
			po:          nil,
			topLevelCFT: types.Int64Value(5),
			fallback:    3,
			want:        5,
		},
		{
			name:        "nil polling options with null top-level uses fallback",
			po:          nil,
			topLevelCFT: types.Int64Null(),
			fallback:    3,
			want:        3,
		},
		{
			name:        "polling options value overrides top-level",
			po:          &PollingOptions{CallFailureThreshold: types.Int64Value(7)},
			topLevelCFT: types.Int64Value(5),
			fallback:    3,
			want:        7,
		},
		{
			name:        "polling options null uses fallback",
			po:          &PollingOptions{CallFailureThreshold: types.Int64Null()},
			topLevelCFT: types.Int64Value(5),
			fallback:    3,
			want:        3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ResolveCallFailureThreshold(tt.po, tt.topLevelCFT, tt.fallback)
			assert.Equal(t, tt.want, got)
		})
	}
}
