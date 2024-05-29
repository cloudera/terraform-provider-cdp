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

import "github.com/hashicorp/terraform-plugin-framework/types"

type PollingOptions struct {
	Async                types.Bool  `tfsdk:"async"`
	PollingTimeout       types.Int64 `tfsdk:"polling_timeout"`
	CallFailureThreshold types.Int64 `tfsdk:"call_failure_threshold"`
}
