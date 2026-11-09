// Copyright 2026 Cloudera. All Rights Reserved.
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
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func GetCallFailureThresholdOrDefault(threshold types.Int64, fallback int) int {
	if !threshold.IsNull() && !threshold.IsUnknown() {
		v := int(threshold.ValueInt64())
		if v > 0 {
			return v
		}
	}
	if fallback > 0 {
		return fallback
	}
	return 0
}

func ResolveAsync(po *PollingOptions, topLevelAsync types.Bool) bool {
	if po != nil {
		return po.Async.ValueBool()
	}
	return topLevelAsync.ValueBool()
}

func ResolveTimeout(ctx context.Context, po *PollingOptions, timeoutGetter func(context.Context, time.Duration) (time.Duration, diag.Diagnostics), defaultTimeout time.Duration) (time.Duration, diag.Diagnostics) {
	if po != nil && !po.PollingTimeout.IsNull() && !po.PollingTimeout.IsUnknown() {
		return time.Duration(po.PollingTimeout.ValueInt64()) * time.Minute, nil
	}
	return timeoutGetter(ctx, defaultTimeout)
}

func ResolveCallFailureThreshold(po *PollingOptions, topLevelCFT types.Int64, fallback int) int {
	if po != nil {
		return GetCallFailureThresholdOrDefault(po.CallFailureThreshold, fallback)
	}
	return GetCallFailureThresholdOrDefault(topLevelCFT, fallback)
}
