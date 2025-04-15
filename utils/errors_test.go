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
	"testing"

	datahubmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
	datalakemodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
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

func TestDecorateEnvironmentUnAuthWhenErrIsNil(t *testing.T) {
	if decorateEnvironmentUnauthorizedErrorIfMessageNotExists(nil) != nil {
		t.Error("Result is not nil but it should be!")
	}
}

func TestDecorateEnvironmentUnAuthWhenErrMessageIsNotEmpty(t *testing.T) {
	msg := "something"
	result := decorateEnvironmentUnauthorizedErrorIfMessageNotExists(&environmentsmodels.Error{
		Message: msg,
	})
	if result.Message != msg {
		t.Errorf("Result message is not the expected one! Expected: '%s', got: '%s'.", msg, result.Message)
	}
}

func TestDecorateEnvironmentUnAuthWhenErrMessageIsEmpty(t *testing.T) {
	result := decorateEnvironmentUnauthorizedErrorIfMessageNotExists(&environmentsmodels.Error{})
	if result.Message != authFailMsg {
		t.Errorf("Result message is not the expected one! Expected: '%s', got: '%s'.", authFailMsg, result.Message)
	}
}

func TestDecorateIamUnAuthWhenErrIsNil(t *testing.T) {
	if decorateIamUnauthorizedErrorIfMessageNotExists(nil) != nil {
		t.Error("Result is not nil but it should be!")
	}
}

func TestDecorateIamUnAuthWhenErrMessageIsNotEmpty(t *testing.T) {
	msg := "something"
	result := decorateIamUnauthorizedErrorIfMessageNotExists(&iammodels.Error{
		Message: msg,
	})
	if result.Message != msg {
		t.Errorf("Result message is not the expected one! Expected: '%s', got: '%s'.", msg, result.Message)
	}
}

func TestDecorateIamUnAuthWhenErrMessageIsEmpty(t *testing.T) {
	result := decorateIamUnauthorizedErrorIfMessageNotExists(&iammodels.Error{})
	if result.Message != authFailMsg {
		t.Errorf("Result message is not the expected one! Expected: '%s', got: '%s'.", authFailMsg, result.Message)
	}
}

func TestDecorateDatalakeUnAuthWhenErrIsNil(t *testing.T) {
	if decorateDatalakeUnauthorizedErrorIfMessageNotExists(nil) != nil {
		t.Error("Result is not nil but it should be!")
	}
}

func TestDecorateDatalakeUnAuthWhenErrMessageIsNotEmpty(t *testing.T) {
	msg := "something"
	result := decorateDatalakeUnauthorizedErrorIfMessageNotExists(&datalakemodels.Error{
		Message: msg,
	})
	if result.Message != msg {
		t.Errorf("Result message is not the expected one! Expected: '%s', got: '%s'.", msg, result.Message)
	}
}

func TestDecorateDatalakeUnAuthWhenErrMessageIsEmpty(t *testing.T) {
	result := decorateDatalakeUnauthorizedErrorIfMessageNotExists(&datalakemodels.Error{})
	if result.Message != authFailMsg {
		t.Errorf("Result message is not the expected one! Expected: '%s', got: '%s'.", authFailMsg, result.Message)
	}
}

func TestDecorateDatahubUnAuthWhenErrIsNil(t *testing.T) {
	if decorateDatahubUnauthorizedErrorIfMessageNotExists(nil) != nil {
		t.Error("Result is not nil but it should be!")
	}
}

func TestDecorateDatahubUnAuthWhenErrMessageIsNotEmpty(t *testing.T) {
	msg := "something"
	result := decorateDatahubUnauthorizedErrorIfMessageNotExists(&datahubmodels.Error{
		Message: msg,
	})
	if result.Message != msg {
		t.Errorf("Result message is not the expected one! Expected: '%s', got: '%s'.", msg, result.Message)
	}
}

func TestDecorateDatahubUnAuthWhenErrMessageIsEmpty(t *testing.T) {
	result := decorateDatahubUnauthorizedErrorIfMessageNotExists(&datahubmodels.Error{})
	if result.Message != authFailMsg {
		t.Errorf("Result message is not the expected one! Expected: '%s', got: '%s'.", authFailMsg, result.Message)
	}
}

func TestDecorateDatabaseUnAuthWhenErrIsNil(t *testing.T) {
	if decorateDatabaseUnauthorizedErrorIfMessageNotExists(nil) != nil {
		t.Error("Result is not nil but it should be!")
	}
}

func TestDecorateDatabaseUnAuthWhenErrMessageIsNotEmpty(t *testing.T) {
	msg := "something"
	result := decorateDatabaseUnauthorizedErrorIfMessageNotExists(&opdbmodels.Error{
		Message: msg,
	})
	if result.Message != msg {
		t.Errorf("Result message is not the expected one! Expected: '%s', got: '%s'.", msg, result.Message)
	}
}

func TestDecorateDatabaseUnAuthWhenErrMessageIsEmpty(t *testing.T) {
	result := decorateDatabaseUnauthorizedErrorIfMessageNotExists(&opdbmodels.Error{})
	if result.Message != authFailMsg {
		t.Errorf("Result message is not the expected one! Expected: '%s', got: '%s'.", authFailMsg, result.Message)
	}
}
