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
	mlmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/models"
	opdbmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/models"
)

type EnvironmentErrorPayload interface {
	GetPayload() *environmentsmodels.Error
}

func AddEnvironmentDiagnosticsError(err error, diagnostics *diag.Diagnostics, errMsg string) {
	msg := err.Error()
	if d, ok := err.(EnvironmentErrorPayload); ok && d.GetPayload() != nil {
		msg = d.GetPayload().Message
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

func AddIamDiagnosticsError(err error, diagnostics *diag.Diagnostics, errMsg string) {
	msg := err.Error()
	if d, ok := err.(IamErrorPayload); ok && d.GetPayload() != nil {
		msg = d.GetPayload().Message
	}
	caser := cases.Title(language.English)
	diagnostics.AddError(
		caser.String(errMsg),
		"Failed to "+errMsg+", unexpected error: "+msg,
	)
}

type MlErrorPayload interface {
	GetPayload() *mlmodels.Error
}

func AddMlDiagnosticsError(err error, diagnostics *diag.Diagnostics, errMsg string) {
	msg := err.Error()
	if d, ok := err.(MlErrorPayload); ok && d.GetPayload() != nil {
		msg = d.GetPayload().Message
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

func AddDatalakeDiagnosticsError(err error, diagnostics *diag.Diagnostics, errMsg string) {
	msg := err.Error()
	if d, ok := err.(DatalakeErrorPayload); ok && d.GetPayload() != nil {
		msg = d.GetPayload().Message
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

func AddDatahubDiagnosticsError(err error, diagnostics *diag.Diagnostics, errMsg string) {
	msg := err.Error()
	if d, ok := err.(DatahubErrorPayload); ok && d.GetPayload() != nil {
		msg = d.GetPayload().Message
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

func AddDatabaseDiagnosticsError(err error, diagnostics *diag.Diagnostics, errMsg string) {
	msg := err.Error()
	if d, ok := err.(DatabaseErrorPayload); ok && d.GetPayload() != nil {
		msg = d.GetPayload().Message
	}
	caser := cases.Title(language.English)
	diagnostics.AddError(
		caser.String(errMsg),
		"Failed to "+errMsg+", unexpected error: "+msg,
	)
}
