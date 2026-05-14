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

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

type environmentReferenceValidator struct{}

func (v environmentReferenceValidator) Description(_ context.Context) string {
	return "validates environment_name and environment fields"
}

func (v environmentReferenceValidator) MarkdownDescription(_ context.Context) string {
	return "exactly one of `environment_name` and `environment` must be provided"
}

func (v environmentReferenceValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	var environmentName types.String
	var environment types.String

	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("environment_name"), &environmentName)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("environment"), &environment)...)

	if resp.Diagnostics.HasError() {
		return
	}

	environmentNameGiven := !environmentName.IsNull() && !environmentName.IsUnknown() && environmentName.ValueString() != ""
	environmentGiven := !environment.IsNull() && !environment.IsUnknown() && environment.ValueString() != ""

	validationLogJSON := v.prepareValidationLog(environmentName, environment)
	utils.LogSilently(ctx, "Validating environment name/CRN configuration: ", validationLogJSON)

	if environmentNameGiven && environmentGiven {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid environment configuration",
			"When `environment_name` is set, `environment` must not be set, and vice versa.",
		)
		return
	}

	anyUnknown := environmentName.IsUnknown() || environment.IsUnknown()
	if anyUnknown {
		tflog.Debug(ctx, "Bypassing validation due to unknown environment fields")
		return
	}

	if !environmentNameGiven && !environmentGiven {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid environment configuration",
			"Exactly one of `environment_name` and `environment` must be provided.",
		)
	}
}

func (v environmentReferenceValidator) prepareValidationLog(environmentName types.String, environment types.String) []byte {
	validationLog := map[string]interface{}{
		"values": map[string]string{
			"environment_name": environmentName.String(),
			"environment":      environment.String(),
		},
	}
	validationLogJSON, _ := json.Marshal(validationLog)
	return validationLogJSON
}

func EnvironmentReferenceValidator() validator.String {
	return environmentReferenceValidator{}
}
