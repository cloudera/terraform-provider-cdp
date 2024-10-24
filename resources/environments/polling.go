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
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

func waitForEnvironmentToBeDeleted(environmentName string, fallbackTimeout time.Duration, callFailureThresholdDefault int, client *client.Environments, ctx context.Context, options *utils.PollingOptions) error {
	timeout, err := utils.CalculateTimeoutOrDefault(ctx, options, fallbackTimeout)
	callFailureThreshold, failureThresholdError := utils.CalculateCallFailureThresholdOrDefault(ctx, options, callFailureThresholdDefault)
	if failureThresholdError != nil {
		return failureThresholdError
	}
	callFailedCount := 0
	if err != nil {
		return err
	}
	stateConf := &retry.StateChangeConf{
		Pending: []string{"STORAGE_CONSUMPTION_COLLECTION_UNSCHEDULING_IN_PROGRESS",
			"NETWORK_DELETE_IN_PROGRESS",
			"FREEIPA_DELETE_IN_PROGRESS",
			"RDBMS_DELETE_IN_PROGRESS",
			"IDBROKER_MAPPINGS_DELETE_IN_PROGRESS",
			"S3GUARD_TABLE_DELETE_IN_PROGRESS",
			"CLUSTER_DEFINITION_DELETE_PROGRESS",
			"CLUSTER_DEFINITION_CLEANUP_PROGRESS",
			"UMS_RESOURCE_DELETE_IN_PROGRESS",
			"DELETE_INITIATED",
			"DATAHUB_CLUSTERS_DELETE_IN_PROGRESS",
			"DATALAKE_CLUSTERS_DELETE_IN_PROGRESS",
			"PUBLICKEY_DELETE_IN_PROGRESS",
			"EVENT_CLEANUP_IN_PROGRESS",
			"EXPERIENCE_DELETE_IN_PROGRESS",
			"ENVIRONMENT_RESOURCE_ENCRYPTION_DELETE_IN_PROGRESS",
			"ENVIRONMENT_ENCRYPTION_RESOURCES_DELETED"},
		Target:       []string{},
		Delay:        5 * time.Second,
		Timeout:      *timeout,
		PollInterval: 10 * time.Second,
		Refresh: func() (interface{}, string, error) {
			tflog.Debug(ctx, fmt.Sprintf("About to describe environment: %s", environmentName))
			params := operations.NewDescribeEnvironmentParamsWithContext(ctx)
			params.WithInput(&environmentsmodels.DescribeEnvironmentRequest{EnvironmentName: &environmentName})
			resp, err := client.Operations.DescribeEnvironment(params)
			if err != nil {
				tflog.Warn(ctx, fmt.Sprintf("Error describing environment: %s", err))
				var envErr *operations.DescribeEnvironmentDefault
				if errors.As(err, &envErr) {
					if isEnvNotFoundError(envErr) {
						return nil, "", nil
					}
				}
				callFailedCount++
				if callFailedCount <= callFailureThreshold {
					tflog.Warn(ctx, fmt.Sprintf("Error describing environment with call failure due to [%s] but threshold limit is not reached yet (%d out of %d).", err.Error(), callFailedCount, callFailureThreshold))
					return nil, "", nil
				}
				tflog.Error(ctx, fmt.Sprintf("Error describing environment (due to: %s) and call failure threshold limit exceeded.", err))
				return nil, "", err
			}
			if resp.GetPayload().Environment == nil {
				log.Printf("Described environment. No environment.")
				return nil, "", nil
			}
			tflog.Info(ctx, fmt.Sprintf("Described environment's status: %s", *resp.GetPayload().Environment.Status))
			return checkResponseStatusForError(resp)
		},
	}
	_, err = stateConf.WaitForStateContext(ctx)

	return err
}

func waitForEnvironmentToBeAvailable(environmentName string, fallbackTimeout time.Duration, callFailureThresholdDefault int, client *client.Environments, ctx context.Context, pollingOptions *utils.PollingOptions,
	stateSaverCb func(*environmentsmodels.Environment)) error {
	timeout, err := utils.CalculateTimeoutOrDefault(ctx, pollingOptions, fallbackTimeout)
	if err != nil {
		return err
	}
	callFailureThreshold, failureThresholdError := utils.CalculateCallFailureThresholdOrDefault(ctx, pollingOptions, callFailureThresholdDefault)
	if failureThresholdError != nil {
		return failureThresholdError
	}
	callFailedCount := 0
	stateConf := &retry.StateChangeConf{
		Pending: []string{"CREATION_INITIATED",
			"NETWORK_CREATION_IN_PROGRESS",
			"PUBLICKEY_CREATE_IN_PROGRESS",
			"ENVIRONMENT_RESOURCE_ENCRYPTION_INITIALIZATION_IN_PROGRESS",
			"ENVIRONMENT_VALIDATION_IN_PROGRESS",
			"ENVIRONMENT_INITIALIZATION_IN_PROGRESS",
			"FREEIPA_CREATION_IN_PROGRESS"},
		Target:       []string{"AVAILABLE"},
		Delay:        5 * time.Second,
		Timeout:      *timeout,
		PollInterval: 10 * time.Second,
		Refresh: func() (interface{}, string, error) {
			tflog.Debug(ctx, fmt.Sprintf("About to describe environment %s", environmentName))
			params := operations.NewDescribeEnvironmentParamsWithContext(ctx)
			params.WithInput(&environmentsmodels.DescribeEnvironmentRequest{EnvironmentName: &environmentName})
			resp, err := client.Operations.DescribeEnvironment(params)
			if err != nil {
				// Envs that have just been created may not be returned from Describe Environment request because of eventual
				// consistency. We return an empty state to retry.

				if isEnvNotFoundError(err) {
					tflog.Debug(ctx, fmt.Sprintf("Recoverable error describing environment: %s", err))
					callFailedCount = 0
					return nil, "", nil
				}
				callFailedCount++
				if callFailedCount <= callFailureThreshold {
					tflog.Warn(ctx, fmt.Sprintf("Error describing environment with call failure due to [%s] but threshold limit is not reached yet (%d out of %d).", err.Error(), callFailedCount, callFailureThreshold))
					return nil, "", nil
				}
				tflog.Error(ctx, fmt.Sprintf("Error describing environment (due to: %s) and call failure threshold limit exceeded.", err))
				return nil, "", err
			}
			callFailedCount = 0
			stateSaverCb(resp.Payload.Environment)
			tflog.Info(ctx, fmt.Sprintf("Described environment's status: %s", *resp.GetPayload().Environment.Status))
			return checkResponseStatusForError(resp)
		},
	}
	_, err = stateConf.WaitForStateContext(ctx)

	return err
}

func checkResponseStatusForError(resp *operations.DescribeEnvironmentOK) (interface{}, string, error) {
	if utils.ContainsAsSubstring([]string{"FAILED", "ERROR"}, *resp.GetPayload().Environment.Status) {
		return nil, "", fmt.Errorf("unexpected Enviornment status: %s. Reason: %s", *resp.GetPayload().Environment.Status, resp.GetPayload().Environment.StatusReason)
	}
	return resp, *resp.GetPayload().Environment.Status, nil
}
