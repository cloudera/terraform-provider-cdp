// Copyright 2025 Cloudera. All Rights Reserved.
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
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
)

type RuntimeModel struct {
	Versions []Runtime `tfsdk:"versions"`
}

type Runtime struct {
	Version types.String `tfsdk:"version"`
	Default types.Bool   `tfsdk:"default"`
}

func fromListRuntimesResponse(response *models.ListRuntimesResponse) *RuntimeModel {
	if response == nil {
		return nil
	}

	runtimes := make([]Runtime, len(response.Versions))
	for i, runtime := range response.Versions {
		runtimes[i] = Runtime{
			Version: types.StringPointerValue(runtime.RuntimeVersion),
			Default: types.BoolPointerValue(runtime.DefaultRuntimeVersion),
		}
	}

	return &RuntimeModel{
		Versions: runtimes,
	}
}
