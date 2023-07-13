// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	environmentOperations "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

const Descr_log_prefix = "Result of describe environment: "

func LogEnvironmentSilently(ctx context.Context, content *environmentsmodels.Environment, messagePrefix string) *environmentsmodels.Environment {
	encoded, err := json.Marshal(content)
	if err != nil {
		tflog.Info(ctx, "Logging content as JSON failed due to: "+err.Error())
		return content
	}
	tflog.Debug(ctx, fmt.Sprintf("%s%s", messagePrefix, string(encoded)))
	return content
}

func LogEnvironmentResponseSilently(ctx context.Context, envResponse *environmentOperations.DescribeEnvironmentOK, msgPrefix string) *environmentOperations.DescribeEnvironmentOK {
	encoded, err := json.Marshal(&envResponse)
	if err != nil {
		tflog.Info(ctx, "Logging content as JSON failed due to: "+err.Error())
		return envResponse
	}
	tflog.Debug(ctx, fmt.Sprintf("%s%s", msgPrefix, string(encoded)))
	return envResponse
}
