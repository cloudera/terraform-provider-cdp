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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
)

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

func GetMapFromSingleItemList(d *schema.ResourceData, name string) (map[string]interface{}, error) {
	if l := d.Get(name).([]interface{}); len(l) == 1 && l[0] != nil {
		return l[0].(map[string]interface{}), nil
	}
	return nil, fmt.Errorf("error getting %s", name)
}

func GetMapFromSingleItemListInMap(d map[string]interface{}, name string) (map[string]interface{}, error) {
	if l := d[name].([]interface{}); len(l) == 1 && l[0] != nil {
		return l[0].(map[string]interface{}), nil
	}
	return nil, fmt.Errorf("error getting %s", name)
}

func ToStringList(configured []interface{}) []string {
	vs := make([]string, 0, len(configured))
	for _, v := range configured {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, v.(string))
		}
	}
	return vs
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
		return nil, fmt.Errorf(msg)
	}
	tflog.Info(ctx, fmt.Sprintf("The following polling timeout calculated: %+v", timeout))
	return &timeout, nil
}

func ToBaseTypesStringMap(in map[string]string) map[string]types.String {
	res := map[string]types.String{}
	for k, v := range in {
		res[k] = types.StringValue(v)
	}
	return res
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

func ConvertStringSliceToTypesSet(input []string) types.Set {
	var elems []attr.Value
	elems = make([]attr.Value, len(input))
	for _, str := range input {
		elems = append(elems, types.StringValue(str))
	}
	var set types.Set
	set, _ = types.SetValue(types.StringType, elems)
	return set
}

func ConvertIntToTypesInt64(input int) types.Int64 {
	upgraded := int64(input)
	return types.Int64Value(upgraded)
}

func ConvertInt32ToTypesInt64(input int32) types.Int64 {
	upgraded := int64(input)
	return types.Int64Value(upgraded)
}
