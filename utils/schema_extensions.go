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
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PollingOptions struct {
	Async                types.Bool  `tfsdk:"async"`
	PollingTimeout       types.Int64 `tfsdk:"polling_timeout"`
	CallFailureThreshold types.Int64 `tfsdk:"call_failure_threshold"`
}

func PollingOptionsConflictValidators() []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.Conflicting(
			path.MatchRoot("polling_options"),
			path.MatchRoot("async"),
		),
		resourcevalidator.Conflicting(
			path.MatchRoot("polling_options"),
			path.MatchRoot("call_failure_threshold"),
		),
		resourcevalidator.Conflicting(
			path.MatchRoot("polling_options"),
			path.MatchRoot("timeouts"),
		),
	}
}

func PollingOptionsSchema(pollingTimeoutDefault int64) schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		DeprecationMessage: "Use the top-level `async`, `call_failure_threshold` attributes and the `timeouts` block instead. " +
			"This block will be removed in a future major version.",
		MarkdownDescription: "Polling options for the resource operation. Deprecated: use the top-level `async`, `call_failure_threshold` attributes and the `timeouts` block instead.",
		Attributes: map[string]schema.Attribute{
			"async": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "When set to `true`, the resource operation returns immediately without waiting for completion.",
			},
			"polling_timeout": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(pollingTimeoutDefault),
				MarkdownDescription: fmt.Sprintf("Timeout in minutes for how long to wait for the operation to complete. Defaults to `%d`.", pollingTimeoutDefault),
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"call_failure_threshold": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(3),
				MarkdownDescription: "The number of consecutive internal API call failures before the provider gives up waiting. Defaults to `3`.",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
		},
	}
}
