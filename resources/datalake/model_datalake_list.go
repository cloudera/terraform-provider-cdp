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
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
)

type DatalakeListModel struct {
	Datalakes []Datalake `tfsdk:"datalakes"`
}

type Datalake struct {
	Name           types.String `tfsdk:"name"`
	Crn            types.String `tfsdk:"crn"`
	Status         types.String `tfsdk:"status"`
	EnvironmentCrn types.String `tfsdk:"environment_crn"`
}

func fromListDatalakesResponse(response *models.ListDatalakesResponse) *DatalakeListModel {
	if response == nil {
		return nil
	}

	dls := make([]Datalake, len(response.Datalakes))
	for i, dl := range response.Datalakes {
		dls[i] = Datalake{
			Name:           types.StringPointerValue(dl.DatalakeName),
			Crn:            types.StringPointerValue(dl.Crn),
			Status:         types.StringValue(dl.Status),
			EnvironmentCrn: types.StringValue(dl.EnvironmentCrn),
		}
	}

	return &DatalakeListModel{Datalakes: dls}
}
