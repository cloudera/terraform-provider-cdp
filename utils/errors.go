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
	"errors"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	datahubmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
	datalakemodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
	dwmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	iammodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
	opdbmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/models"
)

const authGuidance = "Verify that your CDP access key is valid and has not expired. " +
	"Check your credentials in ~/.cdp/credentials or the CDP_ACCESS_KEY_ID and CDP_PRIVATE_KEY environment variables."

const permissionGuidance = "Verify that your CDP user or machine user has the required role assignments for this operation."

const notFoundGuidance = "The resource may have been deleted outside of Terraform. " +
	"Run 'terraform plan' to reconcile the state."

var NonRetryableErrorCodes = [...]int{403}

func IsNonRetryableError(code int) bool {
	for _, nonRetryableCode := range NonRetryableErrorCodes {
		if nonRetryableCode == code {
			return true
		}
	}
	return false
}

func IsRetryableError(code int) bool {
	return !IsNonRetryableError(code)
}

func isReadOrDeleteOperation(errMsg string) bool {
	op, _, _ := strings.Cut(errMsg, " ")
	switch strings.ToLower(op) {
	case "read", "reading", "delete", "deleting", "describe":
		return true
	}
	return false
}

func addDiagnosticsErrorFromPayload(code string, apiMessage string, fallbackMsg string, diagnostics *diag.Diagnostics, errMsg string) {
	detail := fallbackMsg

	switch code {
	case "401":
		if apiMessage != "" {
			detail = apiMessage + "\n\n" + authGuidance
		} else {
			detail = authGuidance
		}
	case "403":
		if apiMessage != "" {
			detail = apiMessage + "\n\n" + permissionGuidance
		} else {
			detail = permissionGuidance
		}
	case "NOT_FOUND":
		if apiMessage != "" {
			detail = apiMessage
		}
		if isReadOrDeleteOperation(errMsg) {
			detail += "\n\n" + notFoundGuidance
		}
	default:
		if apiMessage != "" {
			detail = apiMessage
		}
	}

	caser := cases.Title(language.English)
	diagnostics.AddError(
		caser.String(errMsg),
		"Failed to "+errMsg+", unexpected error: "+detail,
	)
}

func addServiceDiagnosticsError(err error, diagnostics *diag.Diagnostics, errMsg string, extract func(error) (string, string)) {
	code, apiMessage := extract(err)
	addDiagnosticsErrorFromPayload(code, apiMessage, err.Error(), diagnostics, errMsg)
}

type EnvironmentErrorPayload interface {
	GetPayload() *environmentsmodels.Error
}

func AddEnvironmentDiagnosticsError(err error, diagnostics *diag.Diagnostics, errMsg string) {
	addServiceDiagnosticsError(err, diagnostics, errMsg, func(e error) (string, string) {
		var d EnvironmentErrorPayload
		if errors.As(e, &d) && d.GetPayload() != nil {
			return d.GetPayload().Code, d.GetPayload().Message
		}
		return "", ""
	})
}

type IamErrorPayload interface {
	GetPayload() *iammodels.Error
}

func AddIamDiagnosticsError(err error, diagnostics *diag.Diagnostics, errMsg string) {
	addServiceDiagnosticsError(err, diagnostics, errMsg, func(e error) (string, string) {
		var d IamErrorPayload
		if errors.As(e, &d) && d.GetPayload() != nil {
			return d.GetPayload().Code, d.GetPayload().Message
		}
		return "", ""
	})
}

type DatalakeErrorPayload interface {
	GetPayload() *datalakemodels.Error
}

func AddDatalakeDiagnosticsError(err error, diagnostics *diag.Diagnostics, errMsg string) {
	addServiceDiagnosticsError(err, diagnostics, errMsg, func(e error) (string, string) {
		var d DatalakeErrorPayload
		if errors.As(e, &d) && d.GetPayload() != nil {
			return d.GetPayload().Code, d.GetPayload().Message
		}
		return "", ""
	})
}

type DatahubErrorPayload interface {
	GetPayload() *datahubmodels.Error
}

func AddDatahubDiagnosticsError(err error, diagnostics *diag.Diagnostics, errMsg string) {
	addServiceDiagnosticsError(err, diagnostics, errMsg, func(e error) (string, string) {
		var d DatahubErrorPayload
		if errors.As(e, &d) && d.GetPayload() != nil {
			return d.GetPayload().Code, d.GetPayload().Message
		}
		return "", ""
	})
}

func AddRecipeDiagnosticsError(err error, diagnostics *diag.Diagnostics, errMsg string) {
	AddDatahubDiagnosticsError(err, diagnostics, errMsg)
}

type DatabaseErrorPayload interface {
	GetPayload() *opdbmodels.Error
}

func AddDatabaseDiagnosticsError(err error, diagnostics *diag.Diagnostics, errMsg string) {
	addServiceDiagnosticsError(err, diagnostics, errMsg, func(e error) (string, string) {
		var d DatabaseErrorPayload
		if errors.As(e, &d) && d.GetPayload() != nil {
			return d.GetPayload().Code, d.GetPayload().Message
		}
		return "", ""
	})
}

type DwErrorPayload interface {
	GetPayload() *dwmodels.Error
}

func AddDwDiagnosticsError(err error, diagnostics *diag.Diagnostics, errMsg string) {
	addServiceDiagnosticsError(err, diagnostics, errMsg, func(e error) (string, string) {
		var d DwErrorPayload
		if errors.As(e, &d) && d.GetPayload() != nil {
			return d.GetPayload().Code, d.GetPayload().Message
		}
		return "", ""
	})
}
