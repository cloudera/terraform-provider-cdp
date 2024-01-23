// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package environments

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

const (
	describeLogPrefix = "Result of describe environment: "
	timeoutOneHour    = time.Hour * 1
)

func describeEnvironmentWithDiagnosticHandle(envName string, id string, ctx context.Context, client *cdp.Client, resp *resource.ReadResponse) (*environmentsmodels.Environment, error) {
	tflog.Info(ctx, "About to describe environment '"+envName+"'.")
	params := operations.NewDescribeEnvironmentParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.DescribeEnvironmentRequest{
		EnvironmentName: &envName,
	})
	descEnvResp, err := client.Environments.Operations.DescribeEnvironment(params)
	if err != nil {
		tflog.Warn(ctx, "Something happened during environment fetch: "+err.Error())
		if isEnvNotFoundError(err) {
			resp.Diagnostics.AddWarning("Resource not found on provider", "Environment not found, removing from state.")
			tflog.Warn(ctx, "Environment not found, removing from state", map[string]interface{}{
				"id": id,
			})
			resp.State.RemoveResource(ctx)
			return nil, err
		}
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "read Environment")
		return nil, err
	}
	return utils.LogEnvironmentSilently(ctx, descEnvResp.GetPayload().Environment, describeLogPrefix), nil
}

func deleteEnvironmentWithDiagnosticHandle(environmentName string, ctx context.Context, client *cdp.Client, resp *resource.DeleteResponse, pollingTimeout *utils.PollingOptions) error {
	params := operations.NewDeleteEnvironmentParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.DeleteEnvironmentRequest{EnvironmentName: &environmentName})
	_, err := client.Environments.Operations.DeleteEnvironment(params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "delete Environment")
		return err
	}

	err = waitForEnvironmentToBeDeleted(environmentName, timeoutOneHour, client.Environments, ctx, pollingTimeout)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "delete Environment")
		return err
	}
	return nil
}

func isEnvNotFoundError(err error) bool {
	if envErr, ok := err.(*operations.DescribeEnvironmentDefault); ok {
		if cdp.IsEnvironmentsError(envErr.GetPayload(), "NOT_FOUND", "") {
			return true
		}
	}
	return false
}

func waitForCreateEnvironmentWithDiagnosticHandle(ctx context.Context, client *cdp.Client, id string, envName string, resp *resource.CreateResponse, options *utils.PollingOptions) (*operations.DescribeEnvironmentOK, error) {
	if err := waitForEnvironmentToBeAvailable(id, timeoutOneHour, client.Environments, ctx, options); err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "create Environment failed")
		return nil, err
	}

	environmentName := envName
	descParams := operations.NewDescribeEnvironmentParamsWithContext(ctx)
	descParams.WithInput(&environmentsmodels.DescribeEnvironmentRequest{
		EnvironmentName: &environmentName,
	})
	descEnvResp, err := client.Environments.Operations.DescribeEnvironment(descParams)
	if err != nil {
		if isEnvNotFoundError(err) {
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
	return descEnvResp, nil
}
