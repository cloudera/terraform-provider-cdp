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
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
)

// TODO vcsomor: This file contains utility methods for multiple purposes, such as
// - credential handling
// - TF <-> GO Data Object transformations
// - Timeout handling
// TODO vcsomor: I'd good to separate these by purpose

func GetCdpClientForResource(req resource.ConfigureRequest, resp *resource.ConfigureResponse) *cdp.Client {
	if req.ProviderData == nil {
		return nil
	}

	client, ok := req.ProviderData.(*cdp.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *cdp.Client, got: %T. Please report this issue to Cloudera.", req.ProviderData),
		)
		return nil
	}
	return client
}

func GetCdpClientForDataSource(req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) *cdp.Client {
	if req.ProviderData == nil {
		return nil
	}

	client, ok := req.ProviderData.(*cdp.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *cdp.Client, got: %T. Please report this issue to Cloudera.", req.ProviderData),
		)
		return nil
	}
	return client
}

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

func FromStringListToListValue(s []string) types.List {
	if s == nil {
		return types.ListNull(types.StringType)
	}

	var elems []attr.Value
	elems = make([]attr.Value, 0, len(s))
	for _, v := range s {
		elems = append(elems, types.StringValue(v))
	}
	var list types.List
	list, _ = types.ListValue(types.StringType, elems)
	return list
}

func FromListValueToStringList(tl types.List) []string {
	if tl.IsNull() {
		return []string{}
	}
	res := make([]string, len(tl.Elements()))
	for i, elem := range tl.Elements() {
		res[i] = elem.(types.String).ValueString()
	}
	return res
}

func FromTfStringSliceToStringSlice(tss []types.String) []string {
	result := make([]string, 0)
	for _, az := range tss {
		result = append(result, az.ValueString())
	}
	return result
}

func FromSetValueToStringList(tl types.Set) []string {
	if tl.IsNull() {
		return []string{}
	}
	res := make([]string, len(tl.Elements()))
	for i, elem := range tl.Elements() {
		res[i] = elem.(types.String).ValueString()
	}
	return res
}

func Int64To32Pointer(in types.Int64) *int32 {
	n32 := Int64To32(in)
	return &n32
}

func Int64To32(in types.Int64) int32 {
	n64 := in.ValueInt64()
	if n64 >= math.MinInt32 && n64 <= math.MaxInt32 {
		return int32(n64)
	}
	panic(fmt.Sprintf("int64 value %d is out of range for int32", n64))
}

func IsStringTypeHasValue(value types.String) bool {
	return !value.IsNull() && !value.IsUnknown() && len(strings.TrimSpace(value.ValueString())) > 0
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
