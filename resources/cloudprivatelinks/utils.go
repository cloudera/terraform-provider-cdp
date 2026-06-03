// Copyright 2023 Cloudera. All Rights Reserved.
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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/cloudprivatelinks/client/operations"
	cloudprivatelinkmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/cloudprivatelinks/models"
)

const (
	pollInterval = 15 * time.Second
	pollTimeout  = 30 * time.Minute
)

func waitForEndpointReady(ctx context.Context, client *operations.Client, trackingID string) ([]*cloudprivatelinkmodels.PrivateLinkEndpointStatus, error) {
	tflog.Debug(ctx, "Waiting for Private Link endpoint to be ready", map[string]interface{}{
		"trackingId": trackingID,
	})

	deadline := time.Now().Add(pollTimeout)
	for time.Now().Before(deadline) {
		params := operations.NewListPrivateLinkEndpointStatusesParams().
			WithInput(&cloudprivatelinkmodels.ListPrivateLinkEndpointStatusesRequest{
				TrackingID: &trackingID,
			})

		resp, err := client.ListPrivateLinkEndpointStatuses(params)
		if err != nil {
			return nil, fmt.Errorf("error polling Private Link endpoint status: %w", err)
		}

		statuses := resp.GetPayload().PrivatelinkEndpoints
		if allEndpointsSettled(statuses) {
			if hasErrors(statuses) {
				return statuses, fmt.Errorf("one or more Private Link endpoints failed")
			}
			return statuses, nil
		}

		tflog.Debug(ctx, "Endpoints still in progress, waiting...", map[string]interface{}{
			"trackingId": trackingID,
		})
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(pollInterval):
		}
	}

	return nil, fmt.Errorf("timed out waiting for Private Link endpoints to be ready (trackingId: %s)", trackingID)
}

func allEndpointsSettled(statuses []*cloudprivatelinkmodels.PrivateLinkEndpointStatus) bool {
	for _, s := range statuses {
		if s.Status == "IN_PROGRESS" {
			return false
		}
	}
	return true
}

func hasErrors(statuses []*cloudprivatelinkmodels.PrivateLinkEndpointStatus) bool {
	for _, s := range statuses {
		if s.Status == "ERROR" {
			return true
		}
	}
	return false
}
