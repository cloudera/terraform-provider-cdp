// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package cloudprivatelinks

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/cloudprivatelinks/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/cloudprivatelinks/client/operations"
	cloudprivatelinkmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/cloudprivatelinks/models"
)

var (
	pollingDelay    = 10 * time.Second
	pollingInterval = 15 * time.Second
)

func waitForEndpointToBeCreated(trackingID string, plClient *client.Cloudprivatelinks, ctx context.Context, timeout time.Duration, callFailureThreshold int) ([]*cloudprivatelinkmodels.PrivateLinkEndpointStatus, error) {
	tflog.Info(ctx, fmt.Sprintf("About to poll private link endpoint creation (trackingId: %s, polling [delay: %s, timeout: %s, interval: %s]).",
		trackingID, pollingDelay, timeout, pollingInterval))

	callFailedCount := 0
	var resultStatuses []*cloudprivatelinkmodels.PrivateLinkEndpointStatus

	stateConf := &retry.StateChangeConf{
		Pending:      []string{"IN_PROGRESS"},
		Target:       []string{"SUCCESS"},
		Delay:        pollingDelay,
		Timeout:      timeout,
		PollInterval: pollingInterval,
		Refresh: func() (interface{}, string, error) {
			tflog.Debug(ctx, fmt.Sprintf("Polling private link endpoint statuses for trackingId: %s", trackingID))
			resp, err := plClient.Operations.ListPrivateLinkEndpointStatusesContext(ctx,
				operations.NewListPrivateLinkEndpointStatusesParams().WithInput(
					&cloudprivatelinkmodels.ListPrivateLinkEndpointStatusesRequest{TrackingID: &trackingID},
				),
			)
			if err != nil {
				callFailedCount++
				if callFailedCount <= callFailureThreshold {
					tflog.Warn(ctx, fmt.Sprintf("Error listing private link endpoint statuses [%s], threshold not reached yet (%d/%d).", err.Error(), callFailedCount, callFailureThreshold))
					return nil, "IN_PROGRESS", nil
				}
				tflog.Error(ctx, fmt.Sprintf("Error listing private link endpoint statuses [%s], threshold exceeded.", err.Error()))
				return nil, "ERROR", err
			}
			callFailedCount = 0
			resultStatuses = resp.GetPayload().PrivatelinkEndpoints

			for _, s := range resultStatuses {
				if s.Status == "ERROR" {
					return nil, "ERROR", fmt.Errorf("private link endpoint for component %s failed: %s", s.ServiceComponent, s.Error)
				}
				if s.Status == "IN_PROGRESS" {
					tflog.Debug(ctx, fmt.Sprintf("Private link endpoint for component %s is still IN_PROGRESS.", s.ServiceComponent))
					return resp, "IN_PROGRESS", nil
				}
			}
			return resp, "SUCCESS", nil
		},
	}

	_, err := stateConf.WaitForStateContext(ctx)
	return resultStatuses, err
}

func waitForEndpointToBeDeleted(trackingID string, plClient *client.Cloudprivatelinks, ctx context.Context, timeout time.Duration, callFailureThreshold int) error {
	tflog.Info(ctx, fmt.Sprintf("About to poll private link endpoint deletion (trackingId: %s, polling [delay: %s, timeout: %s, interval: %s]).",
		trackingID, pollingDelay, timeout, pollingInterval))

	callFailedCount := 0

	stateConf := &retry.StateChangeConf{
		Pending:      []string{"IN_PROGRESS"},
		Target:       []string{"SUCCESS"},
		Delay:        pollingDelay,
		Timeout:      timeout,
		PollInterval: pollingInterval,
		Refresh: func() (interface{}, string, error) {
			tflog.Debug(ctx, fmt.Sprintf("Polling private link endpoint statuses for trackingId: %s", trackingID))
			resp, err := plClient.Operations.ListPrivateLinkEndpointStatusesContext(ctx,
				operations.NewListPrivateLinkEndpointStatusesParams().WithInput(
					&cloudprivatelinkmodels.ListPrivateLinkEndpointStatusesRequest{TrackingID: &trackingID},
				),
			)
			if err != nil {
				callFailedCount++
				if callFailedCount <= callFailureThreshold {
					tflog.Warn(ctx, fmt.Sprintf("Error listing private link endpoint statuses [%s], threshold not reached yet (%d/%d).", err.Error(), callFailedCount, callFailureThreshold))
					return nil, "IN_PROGRESS", nil
				}
				tflog.Error(ctx, fmt.Sprintf("Error listing private link endpoint statuses [%s], threshold exceeded.", err.Error()))
				return nil, "ERROR", err
			}
			callFailedCount = 0

			for _, s := range resp.GetPayload().PrivatelinkEndpoints {
				if s.Status == "ERROR" {
					return nil, "ERROR", fmt.Errorf("private link endpoint deletion for component %s failed: %s", s.ServiceComponent, s.Error)
				}
				if s.Status == "IN_PROGRESS" {
					tflog.Debug(ctx, fmt.Sprintf("Private link endpoint deletion for component %s is still IN_PROGRESS.", s.ServiceComponent))
					return resp, "IN_PROGRESS", nil
				}
			}
			return resp, "SUCCESS", nil
		},
	}

	_, err := stateConf.WaitForStateContext(ctx)
	return err
}
