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
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/client/operations"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

func checkIfDatabaseCreationFailed(resp *operations.DescribeDatabaseOK) (interface{}, string, error) {
	if utils.ContainsAsSubstring(failedStatusKeywords, string(resp.GetPayload().DatabaseDetails.Status)) {
		return nil, "", fmt.Errorf("cluster status became unacceptable: %s", types.StringValue(string(resp.GetPayload().DatabaseDetails.Status)))
	}
	return resp, string(resp.GetPayload().DatabaseDetails.Status), nil
}

func isNotFoundError(err error) bool {
	if d, ok := err.(*operations.DescribeDatabaseDefault); ok && d.GetPayload() != nil {
		return d.GetPayload().Code == "NOT_FOUND"
	}
	return false
}

func isInternalServerError(err error) bool {
	if d, ok := err.(*operations.DescribeDatabaseDefault); ok && d.GetPayload() != nil {
		return d.GetPayload().Code == "UNKNOWN"
	}
	return false
}
