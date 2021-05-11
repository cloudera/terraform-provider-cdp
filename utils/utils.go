package utils

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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
