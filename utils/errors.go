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
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	datahubmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
	datalakemodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	iammodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
	opdbmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/models"
)

const authFailMsg = "authentication failure, no access key provided or the key is no longer valid"

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

type EnvironmentErrorPayload interface {
	GetPayload() *environmentsmodels.Error
}

func decorateEnvironmentUnauthorizedErrorIfMessageNotExists(err *environmentsmodels.Error) *environmentsmodels.Error {
	if err != nil && len(err.Message) == 0 {
		return &environmentsmodels.Error{
			Message: authFailMsg,
		}
	}
	return err
}

func AddEnvironmentDiagnosticsError(err error, diagnostics *diag.Diagnostics, errMsg string) {
	msg := err.Error()
	if d, ok := err.(EnvironmentErrorPayload); ok && d.GetPayload() != nil {
		if d.GetPayload().Code == "401" {
			msg = decorateEnvironmentUnauthorizedErrorIfMessageNotExists(d.GetPayload()).Message
		} else {
			msg = d.GetPayload().Message
		}
	}
	caser := cases.Title(language.English)
	diagnostics.AddError(
		caser.String(errMsg),
		"Failed to "+errMsg+", unexpected error: "+msg,
	)
}

type IamErrorPayload interface {
	GetPayload() *iammodels.Error
}

func decorateIamUnauthorizedErrorIfMessageNotExists(err *iammodels.Error) *iammodels.Error {
	if err != nil && len(err.Message) == 0 {
		return &iammodels.Error{
			Message: authFailMsg,
		}
	}
	return err
}

func AddIamDiagnosticsError(err error, diagnostics *diag.Diagnostics, errMsg string) {
	msg := err.Error()
	if d, ok := err.(IamErrorPayload); ok && d.GetPayload() != nil {
		if d.GetPayload().Code == "401" {
			msg = decorateIamUnauthorizedErrorIfMessageNotExists(d.GetPayload()).Message
		} else {
			msg = d.GetPayload().Message
		}
	}
	caser := cases.Title(language.English)
	diagnostics.AddError(
		caser.String(errMsg),
		"Failed to "+errMsg+", unexpected error: "+msg,
	)
}

type DatalakeErrorPayload interface {
	GetPayload() *datalakemodels.Error
}

func decorateDatalakeUnauthorizedErrorIfMessageNotExists(err *datalakemodels.Error) *datalakemodels.Error {
	if err != nil && len(err.Message) == 0 {
		return &datalakemodels.Error{
			Message: authFailMsg,
		}
	}
	return err
}

func AddDatalakeDiagnosticsError(err error, diagnostics *diag.Diagnostics, errMsg string) {
	msg := err.Error()
	if d, ok := err.(DatalakeErrorPayload); ok && d.GetPayload() != nil {
		if d.GetPayload().Code == "401" {
			msg = decorateDatalakeUnauthorizedErrorIfMessageNotExists(d.GetPayload()).Message
		} else {
			msg = d.GetPayload().Message
		}
	}
	caser := cases.Title(language.English)
	diagnostics.AddError(
		caser.String(errMsg),
		"Failed to "+errMsg+", unexpected error: "+msg,
	)
}

type DatahubErrorPayload interface {
	GetPayload() *datahubmodels.Error
}

func decorateDatahubUnauthorizedErrorIfMessageNotExists(err *datahubmodels.Error) *datahubmodels.Error {
	if err != nil && len(err.Message) == 0 {
		return &datahubmodels.Error{
			Message: authFailMsg,
		}
	}
	return err
}

func decorateRecipeUnauthorizedErrorIfMessageNotExists(err *datahubmodels.Error) *datahubmodels.Error {
	if err != nil && len(err.Message) == 0 {
		return &datahubmodels.Error{
			Message: authFailMsg,
		}
	}
	return err
}

func AddDatahubDiagnosticsError(err error, diagnostics *diag.Diagnostics, errMsg string) {
	msg := err.Error()
	if d, ok := err.(DatahubErrorPayload); ok && d.GetPayload() != nil {
		if d.GetPayload().Code == "401" {
			msg = decorateDatahubUnauthorizedErrorIfMessageNotExists(d.GetPayload()).Message
		} else {
			msg = d.GetPayload().Message
		}
	}
	caser := cases.Title(language.English)
	diagnostics.AddError(
		caser.String(errMsg),
		"Failed to "+errMsg+", unexpected error: "+msg,
	)
}

func AddRecipeDiagnosticsError(err error, diagnostics *diag.Diagnostics, errMsg string) {
	msg := err.Error()
	if d, ok := err.(DatahubErrorPayload); ok && d.GetPayload() != nil {
		if d.GetPayload().Code == "401" {
			msg = decorateRecipeUnauthorizedErrorIfMessageNotExists(d.GetPayload()).Message
		} else {
			msg = d.GetPayload().Message
		}
	}
	caser := cases.Title(language.English)
	diagnostics.AddError(
		caser.String(errMsg),
		"Failed to "+errMsg+", unexpected error: "+msg,
	)
}

type DatabaseErrorPayload interface {
	GetPayload() *opdbmodels.Error
}

func decorateDatabaseUnauthorizedErrorIfMessageNotExists(err *opdbmodels.Error) *opdbmodels.Error {
	if err != nil && len(err.Message) == 0 {
		return &opdbmodels.Error{
			Message: authFailMsg,
		}
	}
	return err
}

func AddDatabaseDiagnosticsError(err error, diagnostics *diag.Diagnostics, errMsg string) {
	msg := err.Error()
	if d, ok := err.(DatabaseErrorPayload); ok && d.GetPayload() != nil {
		if d.GetPayload().Code == "401" {
			msg = decorateDatabaseUnauthorizedErrorIfMessageNotExists(d.GetPayload()).Message
		} else {
			msg = d.GetPayload().Message
		}
	}
	caser := cases.Title(language.English)
	diagnostics.AddError(
		caser.String(errMsg),
		"Failed to "+errMsg+", unexpected error: "+msg,
	)
}
