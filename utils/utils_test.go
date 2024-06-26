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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"testing"
	"time"
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
