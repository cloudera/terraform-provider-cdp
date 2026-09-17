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
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var testAttrTypes = map[string]attr.Type{"testattr": types.StringType}

func TestUseStateOrNullForUnknown_KnownPlan(t *testing.T) {
	request := planmodifier.ObjectRequest{
		State: tfsdk.State{
			Raw: tftypes.NewValue(
				tftypes.Object{AttributeTypes: map[string]tftypes.Type{
					"attr": tftypes.Object{AttributeTypes: map[string]tftypes.Type{"testattr": tftypes.String}},
				}},
				map[string]tftypes.Value{
					"attr": tftypes.NewValue(
						tftypes.Object{AttributeTypes: map[string]tftypes.Type{"testattr": tftypes.String}},
						map[string]tftypes.Value{"testattr": tftypes.NewValue(tftypes.String, "old")},
					),
				},
			),
		},
		StateValue:  types.ObjectValueMust(testAttrTypes, map[string]attr.Value{"testattr": types.StringValue("old")}),
		PlanValue:   types.ObjectValueMust(testAttrTypes, map[string]attr.Value{"testattr": types.StringValue("new")}),
		ConfigValue: types.ObjectValueMust(testAttrTypes, map[string]attr.Value{"testattr": types.StringValue("new")}),
	}
	expected := types.ObjectValueMust(testAttrTypes, map[string]attr.Value{"testattr": types.StringValue("new")})

	resp := &planmodifier.ObjectResponse{PlanValue: request.PlanValue}
	useStateOrNullForUnknown().PlanModifyObject(context.Background(), request, resp)

	if diff := cmp.Diff(expected, resp.PlanValue); diff != "" {
		t.Errorf("unexpected plan value (-want +got):\n%s", diff)
	}
}

func TestUseStateOrNullForUnknown_UnknownPlan_NullState_ResolvesToNull(t *testing.T) {
	request := planmodifier.ObjectRequest{
		State: tfsdk.State{
			Raw: tftypes.NewValue(
				tftypes.Object{AttributeTypes: map[string]tftypes.Type{
					"attr": tftypes.Object{AttributeTypes: map[string]tftypes.Type{"testattr": tftypes.String}},
				}},
				nil,
			),
		},
		StateValue:  types.ObjectNull(testAttrTypes),
		PlanValue:   types.ObjectUnknown(testAttrTypes),
		ConfigValue: types.ObjectNull(testAttrTypes),
	}
	expected := types.ObjectNull(testAttrTypes)

	resp := &planmodifier.ObjectResponse{PlanValue: request.PlanValue}
	useStateOrNullForUnknown().PlanModifyObject(context.Background(), request, resp)

	if diff := cmp.Diff(expected, resp.PlanValue); diff != "" {
		t.Errorf("unexpected plan value (-want +got):\n%s", diff)
	}
}

func TestUseStateOrNullForUnknown_UnknownPlan_WithState_ResolvesToState(t *testing.T) {
	stateValue := types.ObjectValueMust(testAttrTypes, map[string]attr.Value{"testattr": types.StringValue("from-state")})

	request := planmodifier.ObjectRequest{
		State: tfsdk.State{
			Raw: tftypes.NewValue(
				tftypes.Object{AttributeTypes: map[string]tftypes.Type{
					"attr": tftypes.Object{AttributeTypes: map[string]tftypes.Type{"testattr": tftypes.String}},
				}},
				map[string]tftypes.Value{
					"attr": tftypes.NewValue(
						tftypes.Object{AttributeTypes: map[string]tftypes.Type{"testattr": tftypes.String}},
						map[string]tftypes.Value{"testattr": tftypes.NewValue(tftypes.String, "from-state")},
					),
				},
			),
		},
		StateValue:  stateValue,
		PlanValue:   types.ObjectUnknown(testAttrTypes),
		ConfigValue: types.ObjectNull(testAttrTypes),
	}

	resp := &planmodifier.ObjectResponse{PlanValue: request.PlanValue}
	useStateOrNullForUnknown().PlanModifyObject(context.Background(), request, resp)

	if diff := cmp.Diff(stateValue, resp.PlanValue); diff != "" {
		t.Errorf("unexpected plan value (-want +got):\n%s", diff)
	}
}

func TestUseStateOrNullForUnknown_UnknownConfig_NoChange(t *testing.T) {
	unknownPlan := types.ObjectUnknown(testAttrTypes)

	request := planmodifier.ObjectRequest{
		State: tfsdk.State{
			Raw: tftypes.NewValue(
				tftypes.Object{AttributeTypes: map[string]tftypes.Type{
					"attr": tftypes.Object{AttributeTypes: map[string]tftypes.Type{"testattr": tftypes.String}},
				}},
				map[string]tftypes.Value{
					"attr": tftypes.NewValue(
						tftypes.Object{AttributeTypes: map[string]tftypes.Type{"testattr": tftypes.String}},
						map[string]tftypes.Value{"testattr": tftypes.NewValue(tftypes.String, "from-state")},
					),
				},
			),
		},
		StateValue:  types.ObjectValueMust(testAttrTypes, map[string]attr.Value{"testattr": types.StringValue("from-state")}),
		PlanValue:   unknownPlan,
		ConfigValue: types.ObjectUnknown(testAttrTypes),
	}

	resp := &planmodifier.ObjectResponse{PlanValue: request.PlanValue}
	useStateOrNullForUnknown().PlanModifyObject(context.Background(), request, resp)

	if diff := cmp.Diff(unknownPlan, resp.PlanValue); diff != "" {
		t.Errorf("unexpected plan value (-want +got):\n%s", diff)
	}
}

func TestUseStateOrNullForUnknown_NullPlan_NoChange(t *testing.T) {
	nullPlan := types.ObjectNull(testAttrTypes)

	request := planmodifier.ObjectRequest{
		State: tfsdk.State{
			Raw: tftypes.NewValue(
				tftypes.Object{AttributeTypes: map[string]tftypes.Type{
					"attr": tftypes.Object{AttributeTypes: map[string]tftypes.Type{"testattr": tftypes.String}},
				}},
				map[string]tftypes.Value{
					"attr": tftypes.NewValue(
						tftypes.Object{AttributeTypes: map[string]tftypes.Type{"testattr": tftypes.String}},
						map[string]tftypes.Value{"testattr": tftypes.NewValue(tftypes.String, "from-state")},
					),
				},
			),
		},
		StateValue:  types.ObjectValueMust(testAttrTypes, map[string]attr.Value{"testattr": types.StringValue("from-state")}),
		PlanValue:   nullPlan,
		ConfigValue: types.ObjectNull(testAttrTypes),
	}

	resp := &planmodifier.ObjectResponse{PlanValue: request.PlanValue}
	useStateOrNullForUnknown().PlanModifyObject(context.Background(), request, resp)

	if diff := cmp.Diff(nullPlan, resp.PlanValue); diff != "" {
		t.Errorf("unexpected plan value (-want +got):\n%s", diff)
	}
}
