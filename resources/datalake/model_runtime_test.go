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
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
)

func TestNilResponseReturnsNilRuntimeModel(t *testing.T) {
	response := (*models.ListRuntimesResponse)(nil)
	want := (*RuntimeModel)(nil)
	if got := fromListRuntimesResponse(response); !reflect.DeepEqual(got, want) {
		t.Errorf("fromListRuntimesResponse() = %v, want %v", got, want)
	}
}

func TestEmptyResponseReturnsEmptyRuntimeModel(t *testing.T) {
	response := &models.ListRuntimesResponse{Versions: make([]*models.Runtime, 0)}
	want := &RuntimeModel{Versions: []Runtime{}}
	if got := fromListRuntimesResponse(response); !reflect.DeepEqual(got, want) {
		t.Errorf("fromListRuntimesResponse() = %v, want %v", got, want)
	}
}

func TestValidResponseReturnsCorrectRuntimeModel(t *testing.T) {
	versions := make([]*models.Runtime, 1)
	def := true
	ver := "1.0"
	versions[0] = &models.Runtime{RuntimeVersion: &ver, DefaultRuntimeVersion: &def}
	response := &models.ListRuntimesResponse{
		Versions: versions,
	}
	want := &RuntimeModel{
		Versions: []Runtime{
			{Version: types.StringPointerValue(&ver), Default: types.BoolPointerValue(&def)},
		},
	}
	if got := fromListRuntimesResponse(response); !reflect.DeepEqual(got, want) {
		t.Errorf("fromListRuntimesResponse() = %v, want %v", got, want)
	}
}
