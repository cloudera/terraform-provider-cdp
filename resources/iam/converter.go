// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package iam

import (
	"context"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func muRespToModel(ctx context.Context, mu *models.MachineUser, data *machineUserResourceModel) {
	data.Id = types.StringPointerValue(mu.Crn)
	data.Status = types.StringValue(mu.Status)
	data.WorkloadUsername = types.StringValue(mu.WorkloadUsername)
	data.CreationDate = types.StringValue(mu.CreationDate.String())
	data.WorkloadPasswordDetails, _ = types.ObjectValueFrom(ctx,
		map[string]attr.Type{
			"is_password_set":   types.BoolType,
			"expiration_date":   types.StringType,
			"min_lifetime_date": types.StringType,
		},
		workloadPasswordDetails{
			IsPasswordSet:   types.BoolPointerValue(mu.WorkloadPasswordDetails.IsPasswordSet),
			ExpirationDate:  types.StringValue(mu.WorkloadPasswordDetails.PasswordExpirationDate.String()),
			MinLifetimeDate: types.StringValue(mu.WorkloadPasswordDetails.PasswordMinLifetimeDate.String()),
		})
	if mu.AzureCloudIdentities != nil {
		aci := make([]azureCloudIdentity, 0, len(mu.AzureCloudIdentities))
		for _, v := range mu.AzureCloudIdentities {
			aci = append(aci, azureCloudIdentity{
				EnvironmentCrn: types.StringValue(v.EnvironmentCrn),
				ObjectId:       types.StringPointerValue(v.ObjectID),
			})
		}
		data.AzureCloudIdentities, _ = types.SetValueFrom(ctx, types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"environment_crn": types.StringType,
				"object_id":       types.StringType,
			},
		}, aci)
	}
}
