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
