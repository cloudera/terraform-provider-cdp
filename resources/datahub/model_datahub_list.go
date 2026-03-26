// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datahub

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
)

type DatahubListModel struct {
	Datahubs []Datahub `tfsdk:"datahubs"`
}

type Datahub struct {
	Crn             types.String `tfsdk:"crn"`
	Status          types.String `tfsdk:"status"`
	Name            types.String `tfsdk:"name"`
	DatalakeCrn     types.String `tfsdk:"datalake_crn"`
	CloudPlatform   types.String `tfsdk:"cloud_platform"`
	EnvironmentCrn  types.String `tfsdk:"environment_crn"`
	EnvironmentName types.String `tfsdk:"environment_name"`
}

func fromListClustersResponse(response *models.ListClustersResponse) *DatahubListModel {
	if response == nil {
		return nil
	}

	clusters := make([]Datahub, len(response.Clusters))
	for i, dh := range response.Clusters {
		clusters[i] = Datahub{
			Crn:             types.StringPointerValue(dh.Crn),
			Status:          types.StringValue(dh.Status),
			Name:            types.StringPointerValue(dh.ClusterName),
			DatalakeCrn:     types.StringValue(dh.DatalakeCrn),
			CloudPlatform:   types.StringValue(dh.CloudPlatform),
			EnvironmentCrn:  types.StringValue(dh.EnvironmentCrn),
			EnvironmentName: types.StringValue(dh.EnvironmentName),
		}
	}

	return &DatahubListModel{Datahubs: clusters}
}
