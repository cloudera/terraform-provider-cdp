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

	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

func LogEnvironmentSilently(ctx context.Context, content *environmentsmodels.Environment, messagePrefix string) *environmentsmodels.Environment {
	encoded, err := json.Marshal(content)
	if err != nil {
		tflog.Info(ctx, "Logging content as JSON failed due to: "+err.Error())
		return content
	}
	tflog.Debug(ctx, fmt.Sprintf("%s%s", messagePrefix, string(encoded)))
	return content
}

func LogSilently(ctx context.Context, messagePrefix string, in any) {
	encoded, err := json.Marshal(in)
	if err != nil {
		tflog.Info(ctx, "Logging content as JSON failed due to: "+err.Error())
	}
	tflog.Debug(ctx, fmt.Sprintf("%s%s", messagePrefix, string(encoded)))
}
