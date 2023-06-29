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
	"fmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client/operations"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

func checkIfClusterCreationFailed(resp *operations.DescribeClusterOK) (interface{}, string, error) {
	if utils.ContainsAsSubstring(failedStatusKeywords, resp.GetPayload().Cluster.Status) {
		return nil, "", fmt.Errorf("Cluster status became unacceptable: %s", resp.GetPayload().Cluster.Status)
	}
	return resp, resp.GetPayload().Cluster.Status, nil
}

func isNotFoundError(err error) bool {
	if d, ok := err.(*operations.DescribeClusterDefault); ok && d.GetPayload() != nil {
		return d.GetPayload().Code == "NOT_FOUND"
	}
	return false
}

func isInternalServerError(err error) bool {
	if d, ok := err.(*operations.DescribeClusterDefault); ok && d.GetPayload() != nil {
		return d.GetPayload().Code == "UNKNOWN"
	}
	return false
}
