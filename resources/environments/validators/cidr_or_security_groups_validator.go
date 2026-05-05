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
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

const (
	validationErrorSummary           = "Invalid security access configuration"
	validationLogPrefix              = "Validating security access configuration: "
	cidrFieldName                    = "cidr"
	defaultSecurityGroupIdFieldName  = "default_security_group_id"
	defaultSecurityGroupIdsFieldName = "default_security_group_ids"
	securityGroupIdForKnoxFieldName  = "security_group_id_for_knox"
	securityGroupIdsForKnoxFieldName = "security_group_ids_for_knox"
)

type cidrOrSecurityGroupsValidator struct{}

func (v cidrOrSecurityGroupsValidator) Description(_ context.Context) string {
	return "either cidr or security group fields must be configured, but not both"
}

func (v cidrOrSecurityGroupsValidator) MarkdownDescription(_ context.Context) string {
	return "`cidr` is mutually exclusive with the security group fields. If `cidr` is set, security group fields must not be set. If `cidr` is not set, security group fields must be set."
}

func (v cidrOrSecurityGroupsValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	var defaultSecurityGroupID types.String
	var defaultSecurityGroupIDs types.Set
	var securityGroupIDForKnox types.String
	var securityGroupIDsForKnox types.Set

	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, req.Path.ParentPath().AtName(defaultSecurityGroupIdFieldName), &defaultSecurityGroupID)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, req.Path.ParentPath().AtName(defaultSecurityGroupIdsFieldName), &defaultSecurityGroupIDs)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, req.Path.ParentPath().AtName(securityGroupIdForKnoxFieldName), &securityGroupIDForKnox)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, req.Path.ParentPath().AtName(securityGroupIdsForKnoxFieldName), &securityGroupIDsForKnox)...)

	if resp.Diagnostics.HasError() {
		return
	}

	cidrGiven := !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown() && req.ConfigValue.ValueString() != ""

	defaultSecurityGroupIDGiven := knownNonEmptyString(defaultSecurityGroupID)
	defaultSecurityGroupIDsGiven := knownNonEmptySet(defaultSecurityGroupIDs)
	securityGroupIDForKnoxGiven := knownNonEmptyString(securityGroupIDForKnox)
	securityGroupIDsForKnoxGiven := knownNonEmptySet(securityGroupIDsForKnox)

	validationLogJSON := v.prepareValidationLog(req, defaultSecurityGroupID, defaultSecurityGroupIDs, securityGroupIDForKnox, securityGroupIDsForKnox)
	utils.LogSilently(ctx, validationLogPrefix, validationLogJSON)

	anySecurityGroupFieldGiven := defaultSecurityGroupIDGiven ||
		defaultSecurityGroupIDsGiven ||
		securityGroupIDForKnoxGiven ||
		securityGroupIDsForKnoxGiven

	if cidrGiven && anySecurityGroupFieldGiven {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			validationErrorSummary,
			"When `cidr` is set, `default_security_group_id`, `default_security_group_ids`, `security_group_id_for_knox`, and `security_group_ids_for_knox` must not be set.",
		)
		return
	}

	anyUnknown := req.ConfigValue.IsUnknown() ||
		defaultSecurityGroupID.IsUnknown() ||
		defaultSecurityGroupIDs.IsUnknown() ||
		securityGroupIDForKnox.IsUnknown() ||
		securityGroupIDsForKnox.IsUnknown()

	if anyUnknown {
		tflog.Debug(ctx, "Bypassing validation due to unknown security group fields")
		return
	}

	if !cidrGiven && !anySecurityGroupFieldGiven {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			validationErrorSummary,
			"When `cidr` is not set, security group fields must be provided.",
		)
	}
}

func CIDROrSecurityGroupsValidator() validator.String {
	return cidrOrSecurityGroupsValidator{}
}

func (v cidrOrSecurityGroupsValidator) prepareValidationLog(req validator.StringRequest, defaultSecurityGroupID types.String, defaultSecurityGroupIDs types.Set,
	securityGroupIDForKnox types.String, securityGroupIDsForKnox types.Set) []byte {
	validationLog := map[string]interface{}{
		"security_config": map[string]string{
			cidrFieldName:                    req.ConfigValue.String(),
			defaultSecurityGroupIdFieldName:  defaultSecurityGroupID.String(),
			defaultSecurityGroupIdsFieldName: defaultSecurityGroupIDs.String(),
			securityGroupIdForKnoxFieldName:  securityGroupIDForKnox.String(),
			securityGroupIdsForKnoxFieldName: securityGroupIDsForKnox.String(),
		},
	}
	validationLogJSON, _ := json.Marshal(validationLog)
	return validationLogJSON
}

func knownNonEmptyString(v types.String) bool {
	return !v.IsNull() && !v.IsUnknown() && v.ValueString() != ""
}

func knownNonEmptySet(v types.Set) bool {
	return !v.IsNull() && !v.IsUnknown() && len(v.Elements()) > 0
}
