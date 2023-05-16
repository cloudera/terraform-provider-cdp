package utils

import (
	"fmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

func ToBaseTypesStringMap(in map[string]string) map[string]types.String {
	res := map[string]types.String{}
	for k, v := range in {
		res[k] = types.StringValue(v)
	}
	return res
}

func FromListValueToStringList(tl types.List) []string {
	res := make([]string, len(tl.Elements()))
	for i, elem := range tl.Elements() {
		res[i] = elem.(types.String).ValueString()
	}
	return res
}

func FromSetValueToStringList(tl types.Set) []string {
	res := make([]string, len(tl.Elements()))
	for i, elem := range tl.Elements() {
		res[i] = elem.(types.String).ValueString()
	}
	return res
}

func ToListToBaseTypesStringList(in []string) []types.String {
	if in == nil {
		return nil
	}
	res := make([]types.String, 0, len(in))
	for _, v := range in {
		res = append(res, types.StringValue(v))
	}
	return res
}

func ToListToBaseType(in []string) types.List {
	vals := make([]attr.Value, 0, len(in))
	for _, v := range in {
		vals = append(vals, types.StringValue(v))
	}
	res, _ := types.ListValue(types.StringType, vals)
	return res
}
