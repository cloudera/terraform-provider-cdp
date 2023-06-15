// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datahub

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client/operations"
	datahubmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
)

func waitForToBeAvailable(datahubName string, client *client.Datahub, ctx context.Context) (string, error) {
	tflog.Info(ctx, fmt.Sprintf("About to poll cluster (name: %s) creation (polling [delay: %s, timeout: %s, interval :%s]).",
		datahubName, pollingDelay, pollingTimeout, pollingInterval))
	status := ""
	stateConf := &retry.StateChangeConf{
		Pending:                   []string{},
		Target:                    []string{"AVAILABLE"},
		Delay:                     pollingDelay,
		Timeout:                   pollingTimeout,
		PollInterval:              pollingInterval,
		ContinuousTargetOccurence: 2,
		Refresh: func() (interface{}, string, error) {
			tflog.Debug(ctx, fmt.Sprintf("About to describe cluster %s", datahubName))
			resp, err := describeWithRecover(datahubName, client, ctx)
			if err != nil {
				if isNotFoundError(err) {
					tflog.Debug(ctx, fmt.Sprintf("Recoverable error describing cluster: %s", err))
					return nil, "", nil
				}
				tflog.Debug(ctx, fmt.Sprintf("Error describing cluster: %s", err))
				return nil, "", err
			}
			tflog.Debug(ctx, fmt.Sprintf("Described cluster: %s", resp.GetPayload().Cluster.Status))
			intf, st, e := checkIfClusterCreationFailed(resp)
			tflog.Debug(ctx, fmt.Sprintf("Updating returning status from '%s' to '%s'", status, st))
			status = st
			return intf, st, e
		},
	}
	_, err := stateConf.WaitForStateContext(ctx)
	return status, err
}

func waitForToBeDeleted(datahubName string, client *client.Datahub, ctx context.Context) error {
	tflog.Info(ctx, fmt.Sprintf("About to poll cluster (name: %s) deletion (polling [delay: %s, timeout: %s, interval :%s]).",
		datahubName, pollingDelay, pollingTimeout, pollingInterval))
	stateConf := &retry.StateChangeConf{
		Target:       []string{},
		Delay:        pollingDelay,
		Timeout:      pollingTimeout,
		PollInterval: pollingInterval,
		Refresh: func() (interface{}, string, error) {
			resp, err := describeWithRecover(datahubName, client, ctx)
			if err != nil {
				tflog.Debug(ctx, fmt.Sprintf("Error describing cluster: %s", err))
				if envErr, ok := err.(*operations.DescribeClusterDefault); ok {
					if cdp.IsDatahubError(envErr.GetPayload(), "NOT_FOUND", "") {
						return nil, "", nil
					}
				}
				return nil, "", err
			}
			if resp.GetPayload().Cluster == nil {
				tflog.Debug(ctx, "Datahub described. No cluster.")
				return nil, "", nil
			}
			tflog.Debug(ctx, fmt.Sprintf("Described cluster: %s", resp.GetPayload().Cluster.Status))
			return resp, resp.GetPayload().Cluster.Status, nil
		},
	}
	_, err := stateConf.WaitForStateContext(ctx)
	return err
}

func describeWithRecover(clusterName string, client *client.Datahub, ctx context.Context) (*operations.DescribeClusterOK, error) {
	tflog.Debug(ctx, fmt.Sprintf("Describing cluster with name: %s", clusterName))
	resp, err := client.Operations.DescribeCluster(operations.NewDescribeClusterParamsWithContext(ctx).WithInput(&datahubmodels.DescribeClusterRequest{ClusterName: &clusterName}))
	for i := 0; i < internalServerErrorRetryQuantity; i++ {
		if err != nil {
			if isInternalServerError(err) {
				tflog.Debug(ctx, fmt.Sprintf("Cluster describe came back with internal server error. "+
					"About to (#%d.) re-attempt to describe cluster '%s'.", i+1, clusterName))
				resp, err = client.Operations.DescribeCluster(operations.NewDescribeClusterParamsWithContext(ctx).WithInput(&datahubmodels.DescribeClusterRequest{ClusterName: &clusterName}))
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
