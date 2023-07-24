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
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"log"
	"time"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

func waitForEnvironmentToBeDeleted(environmentName string, timeout time.Duration, client *client.Environments, ctx context.Context) error {
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
		Timeout:      timeout,
		PollInterval: 10 * time.Second,
		Refresh: func() (interface{}, string, error) {
			log.Printf("About to describe environment")
			params := operations.NewDescribeEnvironmentParamsWithContext(ctx)
			params.WithInput(&environmentsmodels.DescribeEnvironmentRequest{EnvironmentName: &environmentName})
			resp, err := client.Operations.DescribeEnvironment(params)
			if err != nil {
				log.Printf("Error describing environment: %s", err)
				if envErr, ok := err.(*operations.DescribeEnvironmentDefault); ok {
					if isEnvNotFoundError(envErr) {
						return nil, "", nil
					}
				}
				return nil, "", err
			}
			if resp.GetPayload().Environment == nil {
				log.Printf("Described environment. No environment.")
				return nil, "", nil
			}
			log.Printf("Described environment's status: %s", *resp.GetPayload().Environment.Status)
			return checkResponseStatusForError(resp)
		},
	}
	_, err := stateConf.WaitForStateContext(ctx)

	return err
}

func waitForEnvironmentToBeAvailable(environmentName string, timeout time.Duration, client *client.Environments, ctx context.Context) error {
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
		Timeout:      timeout,
		PollInterval: 10 * time.Second,
		Refresh: func() (interface{}, string, error) {
			log.Printf("[DEBUG] About to describe environment %s", environmentName)
			params := operations.NewDescribeEnvironmentParamsWithContext(ctx)
			params.WithInput(&environmentsmodels.DescribeEnvironmentRequest{EnvironmentName: &environmentName})
			resp, err := client.Operations.DescribeEnvironment(params)
			if err != nil {
				// Envs that have just been created may not be returned from Describe Environment request because of eventual
				// consistency. We return an empty state to retry.

				if isEnvNotFoundError(err) {
					log.Printf("[DEBUG] Recoverable error describing environment: %s", err)
					return nil, "", nil
				}
				log.Printf("Error describing environment: %s", err)
				return nil, "", err
			}
			log.Printf("Described environment's status: %s", *resp.GetPayload().Environment.Status)
			return checkResponseStatusForError(resp)
		},
	}
	_, err := stateConf.WaitForStateContext(ctx)

	return err
}

func checkResponseStatusForError(resp *operations.DescribeEnvironmentOK) (interface{}, string, error) {
	if utils.ContainsAsSubstring([]string{"FAILED", "ERROR"}, *resp.GetPayload().Environment.Status) {
		return nil, "", fmt.Errorf("unexpected Enviornment status: %s. Reason: %s", *resp.GetPayload().Environment.Status, resp.GetPayload().Environment.StatusReason)
	}
	return resp, *resp.GetPayload().Environment.Status, nil
}
