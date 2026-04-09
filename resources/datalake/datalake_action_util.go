// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.ó

package datalake

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client/operations"
	datalakemodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

const (
	describeLogPrefix = "Result of describe datalake: "
)

func describeDatalakeWithDiagnosticHandle(datalake string, id string, ctx context.Context, client *cdp.Client, diags *diag.Diagnostics, state *tfsdk.State) (*datalakemodels.DatalakeDetails, error) {
	tflog.Info(ctx, "About to describe datalake '"+datalake+"'.")
	params := operations.NewDescribeDatalakeParamsWithContext(ctx)
	params.WithInput(&datalakemodels.DescribeDatalakeRequest{
		DatalakeName: &datalake,
	})
	descDlResp, err := client.Datalake.Operations.DescribeDatalake(params)
	if err != nil {
		tflog.Warn(ctx, "Something happened during environment fetch: "+err.Error())
		if isDatalakeNotFoundError(err) {
			diags.AddWarning("Resource not found on provider", "Environment not found, removing from state.")
			tflog.Warn(ctx, "Environment not found, removing from state", map[string]interface{}{
				"id": id,
			})
			state.RemoveResource(ctx)
			return nil, err
		}
		utils.AddEnvironmentDiagnosticsError(err, diags, "read Environment")
		return nil, err
	}
	return utils.LogDatalakeSilently(ctx, descDlResp.GetPayload().Datalake, describeLogPrefix), nil
}

func isDatalakeNotFoundError(err error) bool {
	var envErr *operations.DescribeDatalakeDefault
	if errors.As(err, &envErr) {
		if cdp.IsDatalakeError(envErr.GetPayload(), "NOT_FOUND", "") {
			return true
		}
	}
	return false
}
