// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datalake

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// useStateOrNullForUnknown returns a plan modifier for Object attributes that
// resolves unknown values: to null during resource creation (no prior state),
// or to the prior state value during updates. This allows an Optional+Computed
// SingleNestedAttribute to be mapped to a Go struct pointer without triggering
// the framework's "cannot handle unknown values" error.
func useStateOrNullForUnknown() planmodifier.Object {
	return useStateOrNullForUnknownModifier{}
}

type useStateOrNullForUnknownModifier struct{}

func (m useStateOrNullForUnknownModifier) Description(_ context.Context) string {
	return "Resolves unknown values to null during creation or to the prior state value during updates."
}

func (m useStateOrNullForUnknownModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m useStateOrNullForUnknownModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	if !req.PlanValue.IsUnknown() {
		return
	}

	if req.ConfigValue.IsUnknown() {
		return
	}

	if req.State.Raw.IsNull() {
		resp.PlanValue = types.ObjectNull(req.PlanValue.AttributeTypes(ctx))
		return
	}

	resp.PlanValue = req.StateValue
}
