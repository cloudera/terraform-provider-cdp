// Copyright 2024 Cloudera. All Rights Reserved.
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
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	datahubmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
	datalakemodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
	dwmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	iammodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
	opdbmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/models"
)

var testNonRetryableCodes = []int{403}

func TestIsNonRetryableErrorWhenCodeIsNonRetryable(t *testing.T) {
	for _, code := range testNonRetryableCodes {
		if !IsNonRetryableError(code) {
			t.Errorf("Code %d is non-retryable but the function returned otherwise!", code)
		}
	}
}

func TestIsNonRetryableErrorWhenCodeIsRetryable(t *testing.T) {
	if IsNonRetryableError(500) {
		t.Error("Code 500 is retryable but the function returned otherwise!")
	}
}

func TestIsRetryableErrorWhenCodeIsRetryable(t *testing.T) {
	for _, code := range testNonRetryableCodes {
		if IsRetryableError(code) {
			t.Errorf("Code %d is non-retryable but the function returned otherwise!", code)
		}
	}
}

func TestIsRetryableErrorWhenCodeIsNonRetryable(t *testing.T) {
	if !IsRetryableError(500) {
		t.Error("Code 500 is retryable but the function returned otherwise!")
	}
}

// --- Centralized error message tests ---

func TestAddDiagnosticsError_401WithEmptyMessage(t *testing.T) {
	var d diag.Diagnostics
	addDiagnosticsErrorFromPayload("401", "", "fallback", &d, "create environment")
	if !d.HasError() {
		t.Fatal("expected error diagnostic")
	}
	detail := d.Errors()[0].Detail()
	if !strings.Contains(detail, authGuidance) {
		t.Errorf("expected auth guidance in detail, got: %s", detail)
	}
}

func TestAddDiagnosticsError_401WithMessage(t *testing.T) {
	var d diag.Diagnostics
	addDiagnosticsErrorFromPayload("401", "token expired", "fallback", &d, "create environment")
	detail := d.Errors()[0].Detail()
	if !strings.Contains(detail, "token expired") {
		t.Errorf("expected API message in detail, got: %s", detail)
	}
	if !strings.Contains(detail, authGuidance) {
		t.Errorf("expected auth guidance appended, got: %s", detail)
	}
}

func TestAddDiagnosticsError_403WithEmptyMessage(t *testing.T) {
	var d diag.Diagnostics
	addDiagnosticsErrorFromPayload("403", "", "fallback", &d, "delete datalake")
	detail := d.Errors()[0].Detail()
	if !strings.Contains(detail, permissionGuidance) {
		t.Errorf("expected permission guidance in detail, got: %s", detail)
	}
}

func TestAddDiagnosticsError_403WithMessage(t *testing.T) {
	var d diag.Diagnostics
	addDiagnosticsErrorFromPayload("403", "access denied for user X", "fallback", &d, "delete datalake")
	detail := d.Errors()[0].Detail()
	if !strings.Contains(detail, "access denied for user X") {
		t.Errorf("expected API message in detail, got: %s", detail)
	}
	if !strings.Contains(detail, permissionGuidance) {
		t.Errorf("expected permission guidance appended, got: %s", detail)
	}
}

func TestAddDiagnosticsError_NotFoundOnRead(t *testing.T) {
	var d diag.Diagnostics
	addDiagnosticsErrorFromPayload("NOT_FOUND", "cluster xyz not found", "fallback", &d, "read datahub")
	detail := d.Errors()[0].Detail()
	if !strings.Contains(detail, "cluster xyz not found") {
		t.Errorf("expected API message in detail, got: %s", detail)
	}
	if !strings.Contains(detail, notFoundGuidance) {
		t.Errorf("expected not-found guidance on read, got: %s", detail)
	}
}

func TestAddDiagnosticsError_NotFoundOnDelete(t *testing.T) {
	var d diag.Diagnostics
	addDiagnosticsErrorFromPayload("NOT_FOUND", "", "fallback", &d, "delete datahub")
	detail := d.Errors()[0].Detail()
	if !strings.Contains(detail, notFoundGuidance) {
		t.Errorf("expected not-found guidance on delete, got: %s", detail)
	}
}

func TestAddDiagnosticsError_NotFoundOnCreate(t *testing.T) {
	var d diag.Diagnostics
	addDiagnosticsErrorFromPayload("NOT_FOUND", "environment xyz not found", "fallback", &d, "create datahub")
	detail := d.Errors()[0].Detail()
	if !strings.Contains(detail, "environment xyz not found") {
		t.Errorf("expected API message in detail, got: %s", detail)
	}
	if strings.Contains(detail, notFoundGuidance) {
		t.Errorf("NOT_FOUND on create should not include reconcile guidance, got: %s", detail)
	}
}

func TestAddDiagnosticsError_NotFoundOnUpdate(t *testing.T) {
	var d diag.Diagnostics
	addDiagnosticsErrorFromPayload("NOT_FOUND", "cluster not found", "fallback", &d, "update datahub")
	detail := d.Errors()[0].Detail()
	if strings.Contains(detail, notFoundGuidance) {
		t.Errorf("NOT_FOUND on update should not include reconcile guidance, got: %s", detail)
	}
}

func TestAddDiagnosticsError_OtherCodeWithMessage(t *testing.T) {
	var d diag.Diagnostics
	addDiagnosticsErrorFromPayload("500", "internal server error", "fallback", &d, "create credential")
	detail := d.Errors()[0].Detail()
	if !strings.Contains(detail, "internal server error") {
		t.Errorf("expected API message in detail, got: %s", detail)
	}
	if strings.Contains(detail, authGuidance) || strings.Contains(detail, permissionGuidance) || strings.Contains(detail, notFoundGuidance) {
		t.Errorf("unexpected guidance in detail for generic error, got: %s", detail)
	}
}

func TestAddDiagnosticsError_NonApiError(t *testing.T) {
	var d diag.Diagnostics
	addDiagnosticsErrorFromPayload("", "", "network timeout", &d, "create credential")
	detail := d.Errors()[0].Detail()
	if !strings.Contains(detail, "network timeout") {
		t.Errorf("expected fallback message in detail, got: %s", detail)
	}
}

func TestAddDiagnosticsError_SummaryIsTitleCased(t *testing.T) {
	var d diag.Diagnostics
	addDiagnosticsErrorFromPayload("", "", "fallback", &d, "create AWS environment")
	summary := d.Errors()[0].Summary()
	if summary != "Create AWS Environment" {
		t.Errorf("expected title-cased summary, got: %s", summary)
	}
}

// --- Per-service function wiring tests ---

type mockEnvError struct {
	payload *environmentsmodels.Error
}

func (m *mockEnvError) Error() string                         { return "raw error" }
func (m *mockEnvError) GetPayload() *environmentsmodels.Error { return m.payload }

func TestAddEnvironmentDiagnosticsError_WithPayload(t *testing.T) {
	var d diag.Diagnostics
	err := &mockEnvError{payload: &environmentsmodels.Error{Code: "401", Message: ""}}
	AddEnvironmentDiagnosticsError(err, &d, "create environment")
	detail := d.Errors()[0].Detail()
	if !strings.Contains(detail, authGuidance) {
		t.Errorf("expected auth guidance, got: %s", detail)
	}
}

func TestAddEnvironmentDiagnosticsError_WithoutPayload(t *testing.T) {
	var d diag.Diagnostics
	AddEnvironmentDiagnosticsError(errors.New("connection refused"), &d, "create environment")
	detail := d.Errors()[0].Detail()
	if !strings.Contains(detail, "connection refused") {
		t.Errorf("expected raw error, got: %s", detail)
	}
}

type mockIamError struct {
	payload *iammodels.Error
}

func (m *mockIamError) Error() string                { return "raw error" }
func (m *mockIamError) GetPayload() *iammodels.Error { return m.payload }

func TestAddIamDiagnosticsError_403(t *testing.T) {
	var d diag.Diagnostics
	err := &mockIamError{payload: &iammodels.Error{Code: "403", Message: "forbidden"}}
	AddIamDiagnosticsError(err, &d, "create machine user")
	detail := d.Errors()[0].Detail()
	if !strings.Contains(detail, "forbidden") || !strings.Contains(detail, permissionGuidance) {
		t.Errorf("expected forbidden + guidance, got: %s", detail)
	}
}

type mockDatalakeError struct {
	payload *datalakemodels.Error
}

func (m *mockDatalakeError) Error() string                     { return "raw error" }
func (m *mockDatalakeError) GetPayload() *datalakemodels.Error { return m.payload }

func TestAddDatalakeDiagnosticsError_NotFound(t *testing.T) {
	var d diag.Diagnostics
	err := &mockDatalakeError{payload: &datalakemodels.Error{Code: "NOT_FOUND", Message: "datalake not found"}}
	AddDatalakeDiagnosticsError(err, &d, "read datalake")
	detail := d.Errors()[0].Detail()
	if !strings.Contains(detail, "datalake not found") || !strings.Contains(detail, notFoundGuidance) {
		t.Errorf("expected not-found + guidance, got: %s", detail)
	}
}

type mockDatahubError struct {
	payload *datahubmodels.Error
}

func (m *mockDatahubError) Error() string                    { return "raw error" }
func (m *mockDatahubError) GetPayload() *datahubmodels.Error { return m.payload }

func TestAddDatahubDiagnosticsError_WithPayload(t *testing.T) {
	var d diag.Diagnostics
	err := &mockDatahubError{payload: &datahubmodels.Error{Code: "500", Message: "cluster creation failed"}}
	AddDatahubDiagnosticsError(err, &d, "create datahub")
	detail := d.Errors()[0].Detail()
	if !strings.Contains(detail, "cluster creation failed") {
		t.Errorf("expected API message, got: %s", detail)
	}
}

func TestAddRecipeDiagnosticsError_WithPayload(t *testing.T) {
	var d diag.Diagnostics
	err := &mockDatahubError{payload: &datahubmodels.Error{Code: "401", Message: ""}}
	AddRecipeDiagnosticsError(err, &d, "create recipe")
	detail := d.Errors()[0].Detail()
	if !strings.Contains(detail, authGuidance) {
		t.Errorf("expected auth guidance, got: %s", detail)
	}
}

type mockDatabaseError struct {
	payload *opdbmodels.Error
}

func (m *mockDatabaseError) Error() string                 { return "raw error" }
func (m *mockDatabaseError) GetPayload() *opdbmodels.Error { return m.payload }

func TestAddDatabaseDiagnosticsError_WithPayload(t *testing.T) {
	var d diag.Diagnostics
	err := &mockDatabaseError{payload: &opdbmodels.Error{Code: "403", Message: ""}}
	AddDatabaseDiagnosticsError(err, &d, "create database")
	detail := d.Errors()[0].Detail()
	if !strings.Contains(detail, permissionGuidance) {
		t.Errorf("expected permission guidance, got: %s", detail)
	}
}

type mockDwError struct {
	payload *dwmodels.Error
}

func (m *mockDwError) Error() string               { return "raw error" }
func (m *mockDwError) GetPayload() *dwmodels.Error { return m.payload }

func TestAddDwDiagnosticsError_WithPayload(t *testing.T) {
	var d diag.Diagnostics
	err := &mockDwError{payload: &dwmodels.Error{Code: "500", Message: "cluster creation failed"}}
	AddDwDiagnosticsError(err, &d, "create DW cluster")
	detail := d.Errors()[0].Detail()
	if !strings.Contains(detail, "cluster creation failed") {
		t.Errorf("expected API message, got: %s", detail)
	}
}

func TestAddDwDiagnosticsError_WithoutPayload(t *testing.T) {
	var d diag.Diagnostics
	AddDwDiagnosticsError(errors.New("connection refused"), &d, "create DW cluster")
	detail := d.Errors()[0].Detail()
	if !strings.Contains(detail, "connection refused") {
		t.Errorf("expected raw error, got: %s", detail)
	}
}

func TestAddDwDiagnosticsError_401(t *testing.T) {
	var d diag.Diagnostics
	err := &mockDwError{payload: &dwmodels.Error{Code: "401", Message: ""}}
	AddDwDiagnosticsError(err, &d, "create DW cluster")
	detail := d.Errors()[0].Detail()
	if !strings.Contains(detail, authGuidance) {
		t.Errorf("expected auth guidance, got: %s", detail)
	}
}

func TestAddEnvironmentDiagnosticsError_WrappedError(t *testing.T) {
	var d diag.Diagnostics
	inner := &mockEnvError{payload: &environmentsmodels.Error{Code: "403", Message: "access denied"}}
	wrapped := fmt.Errorf("operation failed: %w", inner)
	AddEnvironmentDiagnosticsError(wrapped, &d, "create environment")
	detail := d.Errors()[0].Detail()
	if !strings.Contains(detail, "access denied") {
		t.Errorf("expected API message from wrapped error, got: %s", detail)
	}
	if !strings.Contains(detail, permissionGuidance) {
		t.Errorf("expected permission guidance from wrapped error, got: %s", detail)
	}
}
