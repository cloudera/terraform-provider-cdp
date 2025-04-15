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
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

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

func CalculateTimeoutOrDefault(ctx context.Context, options *PollingOptions, fallback time.Duration) (*time.Duration, error) {
	tflog.Debug(ctx, fmt.Sprintf("About to calculate polling timeout using the desired timeout (%+v) and the given fallback timeout (%+v)", options, fallback))
	var timeout time.Duration
	if options != nil && !options.PollingTimeout.IsNull() {
		timeout = time.Duration(options.PollingTimeout.ValueInt64()) * time.Minute
	} else {
		tflog.Info(ctx, "No desired polling timeout is given, the fallback value will be used.")
		timeout = fallback
	}
	if timeout.Minutes() <= 0 {
		msg := "no meaningful timeout value can be calculated based on the given parameters, thus operation shall fail immediately"
		tflog.Warn(ctx, msg)
		return nil, fmt.Errorf("%s", msg)
	}
	tflog.Info(ctx, fmt.Sprintf("The following polling timeout calculated: %+v", timeout))
	return &timeout, nil
}

func CalculateCallFailureThresholdOrDefault(ctx context.Context, options *PollingOptions, fallback int) (int, error) {
	tflog.Debug(ctx, fmt.Sprintf("About to calculate call failure threshold using the desired threshold (%+v) and the given fallback threshold (%+v)", options, fallback))
	var threshold int
	if options != nil && !options.CallFailureThreshold.IsNull() {
		threshold = int(options.CallFailureThreshold.ValueInt64())
	} else {
		tflog.Debug(ctx, "No desired call failure threshold is given, the fallback value will be used.")
		threshold = fallback
	}
	if threshold <= 0 {
		msg := "no meaningful threshold value can be calculated based on the given parameters, thus operation shall fail immediately"
		tflog.Warn(ctx, msg)
		return 0, fmt.Errorf("%s", msg)
	}
	tflog.Debug(ctx, fmt.Sprintf("The following call failure threshold calculated: %+v", threshold))
	return threshold, nil
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

type HasPollingOptions interface {
	GetPollingOptions() *PollingOptions
}

// GetPollingTimeout returns the polling timeout from the given polling options or the fallback value. In case a negative
// value is given, it will be treated as one minute.
func GetPollingTimeout[T HasPollingOptions](p T, fallback time.Duration) time.Duration {
	var timeout time.Duration
	if opts := p.GetPollingOptions(); opts != nil && !opts.PollingTimeout.IsNull() {
		timeout = time.Duration(p.GetPollingOptions().PollingTimeout.ValueInt64()) * time.Minute
	} else {
		timeout = fallback
	}
	if timeout.Seconds() <= 0 {
		return time.Minute
	}
	return timeout
}

// GetCallFailureThreshold returns the call failure threshold from the given polling options or the fallback value.
// In case a negative value is given, it will be treated as zero.
func GetCallFailureThreshold[T HasPollingOptions](p T, fallback int) int {
	var threshold int
	if opts := p.GetPollingOptions(); opts != nil && !opts.CallFailureThreshold.IsNull() {
		threshold = int(p.GetPollingOptions().CallFailureThreshold.ValueInt64())
	} else {
		threshold = fallback
	}
	if threshold <= 0 {
		return 0
	}
	return threshold
}
