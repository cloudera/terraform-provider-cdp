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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

const (
	testFallbackMinutes      = 5
	testFallbackValue        = time.Minute * testFallbackMinutes
	testCallFailureThreshold = 3
)

func TestCalculateTimeoutOrDefault(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx      context.Context
		options  *PollingOptions
		fallback time.Duration
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "Test when default is given and PollingOptions exists but PollingTimeout is nil, then default should come back.",
			args: args{
				ctx: context.TODO(),
				options: &PollingOptions{
					PollingTimeout: types.Int64Null(),
				},
				fallback: testFallbackValue,
			},
			want:    testFallbackMinutes,
			wantErr: false,
		},
		{
			name: "Test when default is given and PollingOptions is nil then default should come back.",
			args: args{
				ctx:      context.TODO(),
				options:  nil,
				fallback: testFallbackValue,
			},
			want:    testFallbackMinutes,
			wantErr: false,
		},
		{
			name: "Test when nor the PollingOptions nor the fallBackValue is given then error should come back.",
			args: args{
				ctx:      context.TODO(),
				options:  nil,
				fallback: 0,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Test when no default value is given but a valid PollingTimeout is, then its value should come back.",
			args: args{
				ctx: context.TODO(),
				options: &PollingOptions{
					PollingTimeout: types.Int64Value(testFallbackMinutes),
				},
				fallback: 0,
			},
			want:    testFallbackMinutes,
			wantErr: false,
		},
		{
			name: "Test when both default value and PollingTimeout is given but both are zero then error should come.",
			args: args{
				ctx: context.TODO(),
				options: &PollingOptions{
					PollingTimeout: types.Int64Value(0),
				},
				fallback: 0,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateTimeoutOrDefault(tt.args.ctx, tt.args.options, tt.args.fallback)
			if tt.wantErr {
				if err == nil {
					t.Errorf("CalculateTimeoutOrDefault() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			} else if got.Minutes() != tt.want {
				t.Errorf("CalculateTimeoutOrDefault() got = %v, expected %v", got.Minutes(), tt.want)
				return
			}
		})
	}
}

func TestCalculateCallFailureThresholdOrDefault(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx      context.Context
		options  *PollingOptions
		fallback int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Test when default is given and PollingOptions exists but CallFailureThreshold is nil, then default should come back.",
			args: args{
				ctx: context.TODO(),
				options: &PollingOptions{
					CallFailureThreshold: types.Int64Null(),
				},
				fallback: testCallFailureThreshold,
			},
			want:    testCallFailureThreshold,
			wantErr: false,
		},
		{
			name: "Test when default is given and PollingOptions is nil then default should come back.",
			args: args{
				ctx:      context.TODO(),
				options:  nil,
				fallback: testCallFailureThreshold,
			},
			want:    testCallFailureThreshold,
			wantErr: false,
		},
		{
			name: "Test when nor the PollingOptions nor the fallbackValue is given then error should come back.",
			args: args{
				ctx:      context.TODO(),
				options:  nil,
				fallback: 0,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Test when no default value is given but a valid CallFailureThreshold is, then its value should come back.",
			args: args{
				ctx: context.TODO(),
				options: &PollingOptions{
					CallFailureThreshold: types.Int64Value(testCallFailureThreshold),
				},
				fallback: 0,
			},
			want:    testCallFailureThreshold,
			wantErr: false,
		},
		{
			name: "Test when both default value and CallFailureThreshold is given but both are zero then error should come.",
			args: args{
				ctx: context.TODO(),
				options: &PollingOptions{
					CallFailureThreshold: types.Int64Value(0),
				},
				fallback: 0,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateCallFailureThresholdOrDefault(tt.args.ctx, tt.args.options, tt.args.fallback)
			if tt.wantErr {
				if err == nil {
					t.Errorf("CalculateCallFailureThresholdOrDefault() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			} else if got != tt.want {
				t.Errorf("CalculateCallFailureThresholdOrDefault() got = %v, expected %v", got, tt.want)
				return
			}
		})
	}
}

type TestPollingOptions struct {
	PollingOptions *PollingOptions
}

func (p *TestPollingOptions) GetPollingOptions() *PollingOptions {
	return p.PollingOptions
}

func TestGetPollingTimeout(t *testing.T) {
	type args struct {
		p        HasPollingOptions
		fallback time.Duration
	}
	tests := []struct {
		description string
		args        args
		want        time.Duration
	}{
		{
			description: "Test when PollingOptions is nil and fallback is given then fallback should come back.",
			args: args{
				p:        &TestPollingOptions{},
				fallback: 90 * time.Minute,
			},
			want: 90 * time.Minute,
		},
		{
			description: "Test when PollingOptions is nil and fallback is not given then 1 minute should come back.",
			args: args{
				p:        &TestPollingOptions{},
				fallback: 0,
			},
			want: time.Minute,
		},
		{
			description: "Test when PollingOptions is given but PollingTimeout is nil and fallback is given then fallback should come back.",
			args: args{
				p:        &TestPollingOptions{PollingOptions: &PollingOptions{}},
				fallback: 90 * time.Minute,
			},
			want: 90 * time.Minute,
		},
		{
			description: "Test when PollingOptions is given but PollingTimeout is nil and fallback is not given then 1 minute should come back.",
			args: args{
				p:        &TestPollingOptions{PollingOptions: &PollingOptions{}},
				fallback: 0,
			},
			want: time.Minute,
		},
		{
			description: "Test when PollingOptions is given and PollingTimeout is given then its value should come back.",
			args: args{
				p:        &TestPollingOptions{PollingOptions: &PollingOptions{PollingTimeout: types.Int64Value(60)}},
				fallback: 90 * time.Minute,
			},
			want: 60 * time.Minute,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := GetPollingTimeout(tt.args.p, tt.args.fallback)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestGetCallFailureThreshold(t *testing.T) {
	type args struct {
		p        HasPollingOptions
		fallback int
	}
	tests := []struct {
		description string
		args        args
		want        int
	}{
		{
			description: "Test when PollingOptions is nil and fallback is given then fallback should come back.",
			args: args{
				p:        &TestPollingOptions{},
				fallback: 3,
			},
			want: 3,
		},
		{
			description: "Test when PollingOptions is nil and fallback is not given then 0 should come back.",
			args: args{
				p:        &TestPollingOptions{},
				fallback: 0,
			},
			want: 0,
		},
		{
			description: "Test when PollingOptions is given but CallFailureThreshold is nil and fallback is given then fallback should come back.",
			args: args{
				p:        &TestPollingOptions{PollingOptions: &PollingOptions{}},
				fallback: 3,
			},
			want: 3,
		},
		{
			description: "Test when PollingOptions is given but CallFailureThreshold is nil and fallback is not given then 0 should come back.",
			args: args{
				p:        &TestPollingOptions{PollingOptions: &PollingOptions{}},
				fallback: 0,
			},
			want: 0,
		},
		{
			description: "Test when PollingOptions is given and CallFailureThreshold is given then its value should come back.",
			args: args{
				p:        &TestPollingOptions{PollingOptions: &PollingOptions{CallFailureThreshold: types.Int64Value(5)}},
				fallback: 3,
			},
			want: 5,
		},
		{
			description: "Test when PollingOptions is nil and fallback is a negative int then 0 should come back.",
			args: args{
				p:        &TestPollingOptions{},
				fallback: -2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := GetCallFailureThreshold(tt.args.p, tt.args.fallback)
			assert.Equal(t, got, tt.want)
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
