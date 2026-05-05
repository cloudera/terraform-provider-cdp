// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package validators

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

type imageRuntimeCompatibilityValidator struct{}

func (v imageRuntimeCompatibilityValidator) Description(_ context.Context) string {
	return "validates image fields against runtime"
}

func (v imageRuntimeCompatibilityValidator) MarkdownDescription(_ context.Context) string {
	return "when `runtime` is set, only `image.os` may be provided; otherwise, use `image.catalog_name` and/or `image.id`"
}

func (v imageRuntimeCompatibilityValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	var runtime types.String
	var catalogName types.String
	var imageID types.String
	var os types.String

	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, req.Path.ParentPath().AtName("runtime"), &runtime)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, req.Path.AtName("catalog_name"), &catalogName)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, req.Path.AtName("id"), &imageID)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, req.Path.AtName("os"), &os)...)

	if resp.Diagnostics.HasError() {
		return
	}

	runtimeGiven := !runtime.IsNull() && !runtime.IsUnknown() && runtime.ValueString() != ""
	catalogNameGiven := !catalogName.IsNull() && !catalogName.IsUnknown() && catalogName.ValueString() != ""
	imageIDGiven := !imageID.IsNull() && !imageID.IsUnknown() && imageID.ValueString() != ""
	osGiven := !os.IsNull() && !os.IsUnknown() && os.ValueString() != ""

	runtimeWithCatalogNameOrImageIDInvalid := runtimeGiven && (catalogNameGiven || imageIDGiven)
	osWithoutRuntimeInvalid := !runtimeGiven && osGiven

	validationLogJSON := v.prepareValidationLog(runtime, catalogName, imageID, os)
	utils.LogSilently(ctx, "Validating image runtime compatibility configuration: ", validationLogJSON)

	if runtimeGiven {
		if runtimeWithCatalogNameOrImageIDInvalid {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid image configuration",
				"When `runtime` is set, only `image.os` can be provided. `image.catalog_name` and `image.id` must not be set.",
			)
		}
		return
	}

	if osWithoutRuntimeInvalid {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid image configuration",
			"When `runtime` is not set, `image.os` must not be provided. Use `image.catalog_name` and/or `image.id` for selecting an image.",
		)
	}
}

func (v imageRuntimeCompatibilityValidator) prepareValidationLog(runtime types.String, catalogName types.String, imageID types.String, os types.String) []byte {
	validationLog := map[string]interface{}{
		"values": map[string]string{
			"runtime":            runtime.String(),
			"image.catalog_name": catalogName.String(),
			"image.id":           imageID.String(),
			"image.os":           os.String(),
		},
	}
	validationLogJSON, _ := json.Marshal(validationLog)
	return validationLogJSON
}

func ImageRuntimeCompatibilityValidator() validator.Object {
	return imageRuntimeCompatibilityValidator{}
}
