// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package environments

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

type EnvironmentListModel struct {
	Environments []Environment `tfsdk:"environments"`
}

type Environment struct {
	Name           types.String `tfsdk:"name"`
	Crn            types.String `tfsdk:"crn"`
	Status         types.String `tfsdk:"status"`
	Region         types.String `tfsdk:"region"`
	CloudPlatform  types.String `tfsdk:"cloud_platform"`
	CredentialName types.String `tfsdk:"credential_name"`
}

func fromListEnvironmentsResponse(response *models.ListEnvironmentsResponse) *EnvironmentListModel {
	if response == nil {
		return nil
	}

	environments := make([]Environment, len(response.Environments))
	for i, env := range response.Environments {
		environments[i] = Environment{
			Name:           types.StringPointerValue(env.EnvironmentName),
			Crn:            types.StringPointerValue(env.Crn),
			Status:         types.StringPointerValue(env.Status),
			Region:         types.StringPointerValue(env.Region),
			CloudPlatform:  types.StringPointerValue(env.CloudPlatform),
			CredentialName: types.StringPointerValue(env.CredentialName),
		}
	}

	return &EnvironmentListModel{Environments: environments}
}
