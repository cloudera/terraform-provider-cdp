// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package opdb

import (
	"fmt"
	"testing"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/models"
)

func newDescribeDatabaseError(code string) *operations.DescribeDatabaseDefault {
	e := operations.NewDescribeDatabaseDefault(500)
	e.Payload = &models.Error{Code: code}
	return e
}

func TestIsNotFoundError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"direct NOT_FOUND", newDescribeDatabaseError("NOT_FOUND"), true},
		{"direct UNKNOWN", newDescribeDatabaseError("UNKNOWN"), false},
		{"wrapped NOT_FOUND", fmt.Errorf("context: %w", newDescribeDatabaseError("NOT_FOUND")), true},
		{"nil payload", operations.NewDescribeDatabaseDefault(500), false},
		{"unrelated error", fmt.Errorf("something else"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNotFoundError(tt.err); got != tt.want {
				t.Errorf("isNotFoundError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsInternalServerError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"direct UNKNOWN", newDescribeDatabaseError("UNKNOWN"), true},
		{"direct NOT_FOUND", newDescribeDatabaseError("NOT_FOUND"), false},
		{"wrapped UNKNOWN", fmt.Errorf("context: %w", newDescribeDatabaseError("UNKNOWN")), true},
		{"unrelated error", fmt.Errorf("something else"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInternalServerError(tt.err); got != tt.want {
				t.Errorf("isInternalServerError() = %v, want %v", got, tt.want)
			}
		})
	}
}
