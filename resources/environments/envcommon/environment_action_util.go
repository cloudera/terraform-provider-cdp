// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package envcommon

import (
	"context"
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

const (
	DescribeLogPrefix = "Result of describe environment: "
	TimeoutOneHour    = time.Hour * 1
)

func DescribeEnvironmentWithDiagnosticHandle(envName string, id string, ctx context.Context, client *cdp.Client, diags *diag.Diagnostics, state *tfsdk.State) (*environmentsmodels.Environment, error) {
	tflog.Info(ctx, "About to describe environment '"+envName+"'.")
	params := operations.NewDescribeEnvironmentParams()
	params.WithInput(&environmentsmodels.DescribeEnvironmentRequest{
		EnvironmentName: &envName,
	})
	descEnvResp, err := client.Environments.Operations.DescribeEnvironmentContext(ctx, params)
	if err != nil {
		tflog.Warn(ctx, "Something happened during environment fetch: "+err.Error())
		if IsEnvNotFoundError(err) {
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
	return utils.LogEnvironmentSilently(ctx, descEnvResp.GetPayload().Environment, DescribeLogPrefix), nil
}

func DeleteEnvironmentWithDiagnosticHandle(environmentName string, cascading bool, forced bool, ctx context.Context, client *cdp.Client, resp *resource.DeleteResponse, pollingOptions *utils.PollingOptions) error {
	params := operations.NewDeleteEnvironmentParams()
	params.WithInput(&environmentsmodels.DeleteEnvironmentRequest{EnvironmentName: &environmentName, Cascading: cascading, Forced: forced})
	_, err := client.Environments.Operations.DeleteEnvironmentContext(ctx, params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "delete Environment")
		return err
	}

	err = WaitForEnvironmentToBeDeleted(environmentName, TimeoutOneHour, CallFailureThreshold, client.Environments, ctx, pollingOptions)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "delete Environment")
		return err
	}
	return nil
}

func IsEnvNotFoundError(err error) bool {
	if envErr, ok := errors.AsType[*operations.DescribeEnvironmentDefault](err); ok {
		if cdp.IsEnvironmentsError(envErr.GetPayload(), "NOT_FOUND", "") {
			return true
		}
	}
	return false
}

func WaitForCreateEnvironmentWithDiagnosticHandle(ctx context.Context, client *cdp.Client, id string, envName string, resp *resource.CreateResponse, options *utils.PollingOptions,
	stateSaverCb func(*environmentsmodels.Environment)) (*environmentsmodels.Environment, error) {
	if err := WaitForEnvironmentToBeAvailable(id, TimeoutOneHour, CallFailureThreshold, client.Environments, ctx, options, stateSaverCb); err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "create Environment failed")
		return nil, err
	}

	descParams := operations.NewDescribeEnvironmentParams()
	descParams.WithInput(&environmentsmodels.DescribeEnvironmentRequest{
		EnvironmentName: new(envName),
	})
	descEnvResp, err := client.Environments.Operations.DescribeEnvironmentContext(ctx, descParams)
	if err != nil {
		if IsEnvNotFoundError(err) {
			resp.Diagnostics.AddWarning("Resource not found on provider", "Environment not found, removing from state.")
			tflog.Warn(ctx, "Environment not found, removing from state", map[string]interface{}{
				"id": id,
			})
			resp.State.RemoveResource(ctx)
			return nil, err
		}
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "create Environment failed")
		return nil, err
	}
	return descEnvResp.GetPayload().Environment, nil
}
