// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package environments

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

func ConvertTags(ctx context.Context, tagsIn types.Map) []*environmentsmodels.TagRequest {
	if !tagsIn.IsNull() && len(tagsIn.Elements()) > 0 {
		var tags []*environmentsmodels.TagRequest
		for k, v := range tagsIn.Elements() {
			val, diag := v.(basetypes.StringValuable).ToStringValue(ctx)
			if !diag.HasError() {
				key := string(k)
				tags = append(tags, &environmentsmodels.TagRequest{
					Key:   &key,
					Value: val.ValueStringPointer(),
				})
			}
		}
		return tags
	}
	return nil
}

func ConvertGcpTags(ctx context.Context, tagsIn types.Map) []*environmentsmodels.GcpTagRequest {
	if !tagsIn.IsNull() && len(tagsIn.Elements()) > 0 {
		var tags []*environmentsmodels.GcpTagRequest
		for k, v := range tagsIn.Elements() {
			val, diag := v.(basetypes.StringValuable).ToStringValue(ctx)
			if !diag.HasError() {
				key := string(k)
				tags = append(tags, &environmentsmodels.GcpTagRequest{
					Key:   &key,
					Value: val.ValueStringPointer(),
				})
			}
		}
		return tags
	}
	return nil
}

func ConvertIntToInt32IfPossible(value int) (basetypes.Int32Value, error) {
	i32, err := safeIntToInt32(value)
	if err != nil {
		return basetypes.Int32Value{}, err
	}
	return types.Int32Value(i32), nil
}

func safeIntToInt32(n int) (int32, error) {
	if n > math.MaxInt32 || n < math.MinInt32 {
		return 0, fmt.Errorf("value %d out of int32 range", n)
	}
	return int32(n), nil
}

func getStringValueIfNotEmpty(s string) types.String {
	val := strings.TrimSpace(s)
	if len(val) > 0 {
		return types.StringValue(val)
	}
	return types.StringNull()
}
