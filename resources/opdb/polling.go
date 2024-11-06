// Copyright 2024 Cloudera. All Rights Reserved.
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
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/client/operations"
	opdbmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

func waitForToBeAvailable(dataBaseName string, environmentName string, client *client.Opdb, ctx context.Context, options *utils.PollingOptions) (string, error) {
	tflog.Info(ctx, fmt.Sprintf("About to poll Database (name: %s) creation (polling [delay: %s, timeout: %s, interval :%s]).",
		dataBaseName, pollingDelay, pollingTimeout, pollingInterval))
	status := ""
	timeout, err := utils.CalculateTimeoutOrDefault(ctx, options, pollingTimeout)
	if err != nil {
		return "", err
	}
	failureThreshold, failureThresholdErr := utils.CalculateCallFailureThresholdOrDefault(ctx, options, callFailureThreshold)
	if failureThresholdErr != nil {
		return "", failureThresholdErr
	}
	callFailedCount := 0
	stateConf := &retry.StateChangeConf{
		Pending:                   []string{},
		Target:                    []string{"AVAILABLE"},
		Delay:                     pollingDelay,
		Timeout:                   *timeout,
		PollInterval:              pollingInterval,
		ContinuousTargetOccurence: 2,
		Refresh: func() (interface{}, string, error) {
			tflog.Debug(ctx, fmt.Sprintf("About to describe Database %s", dataBaseName))
			resp, err := describeWithRecover(dataBaseName, environmentName, client, ctx)
			if err != nil {
				if isNotFoundError(err) {
					tflog.Debug(ctx, fmt.Sprintf("Recoverable error describing Database: %s", err))
					callFailedCount = 0
					return nil, "", nil
				}
				callFailedCount++
				if callFailedCount <= failureThreshold {
					tflog.Warn(ctx, fmt.Sprintf("Error describing cluster with call failure due to [%s] but threshold limit is not reached yet (%d out of %d).", err.Error(), callFailedCount, callFailureThreshold))
					return nil, "", nil
				}
				tflog.Error(ctx, fmt.Sprintf("Error describing Database (due to: %s) and call failure threshold limit exceeded.", err))
				return nil, "", err
			}
			callFailedCount = 0
			tflog.Debug(ctx, fmt.Sprintf("Described Database: %s", resp.GetPayload().DatabaseDetails.Status))
			intf, st, e := checkIfDatabaseCreationFailed(resp)
			tflog.Debug(ctx, fmt.Sprintf("Updating returning status from '%s' to '%s'", status, st))
			status = st
			return intf, st, e
		},
	}
	_, err = stateConf.WaitForStateContext(ctx)
	return status, err
}

func waitForToBeDeleted(dataBaseName string, environmentName string, client *client.Opdb, ctx context.Context, options *utils.PollingOptions) error {
	tflog.Info(ctx, fmt.Sprintf("About to poll Database (name: %s) deletion (polling [delay: %s, timeout: %s, interval :%s]).",
		dataBaseName, pollingDelay, pollingTimeout, pollingInterval))
	timeout, err := utils.CalculateTimeoutOrDefault(ctx, options, pollingTimeout)
	if err != nil {
		return err
	}
	stateConf := &retry.StateChangeConf{
		Target:       []string{},
		Delay:        pollingDelay,
		Timeout:      *timeout,
		PollInterval: pollingInterval,
		Refresh: func() (interface{}, string, error) {
			resp, err := describeWithRecover(dataBaseName, environmentName, client, ctx)
			if err != nil {
				tflog.Debug(ctx, fmt.Sprintf("Error describing Database: %s", err))
				var envErr *operations.DescribeDatabaseDefault
				if errors.As(err, &envErr) {
					if cdp.IsDatabaseError(envErr.GetPayload(), "NOT_FOUND", "") {
						return nil, "", nil
					}
				}
				return nil, "", err
			}
			if resp.GetPayload().DatabaseDetails == nil {
				tflog.Debug(ctx, "Database described. No cluster.")
				return nil, "", nil
			}
			tflog.Debug(ctx, fmt.Sprintf("Described cluster: %s", resp.GetPayload().DatabaseDetails.Status))
			return resp, string(resp.GetPayload().DatabaseDetails.Status), nil
			// TODO maybe there is a better solution for this
		},
	}
	_, err = stateConf.WaitForStateContext(ctx)
	return err
}

func describeWithRecover(dataBaseName string, environmentName string, client *client.Opdb, ctx context.Context) (*operations.DescribeDatabaseOK, error) {
	tflog.Debug(ctx, fmt.Sprintf("Describing Database with name: %s", dataBaseName))
	resp, err := client.Operations.DescribeDatabase(operations.NewDescribeDatabaseParamsWithContext(ctx).WithInput(&opdbmodels.DescribeDatabaseRequest{DatabaseName: &dataBaseName, EnvironmentName: &environmentName}))
	for i := 0; i < internalServerErrorRetryQuantity; i++ {
		if err != nil {
			if isInternalServerError(err) {
				tflog.Debug(ctx, fmt.Sprintf("Database describe came back with internal server error. "+
					"About to (#%d.) re-attempt to describe Database '%s'.", i+1, dataBaseName))
				resp, err = client.Operations.DescribeDatabase(operations.NewDescribeDatabaseParamsWithContext(ctx).WithInput(&opdbmodels.DescribeDatabaseRequest{DatabaseName: &dataBaseName, EnvironmentName: &environmentName}))
				continue
			} else {
				return resp, err
			}
		} else {
			break
		}
	}
	return resp, err
}
