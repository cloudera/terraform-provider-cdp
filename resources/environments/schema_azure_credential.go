// Copyright 2025 Cloudera. All Rights Reserved.
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
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

func (r *azureCredentialResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	attr := map[string]schema.Attribute{}
	utils.Append(attr, generalCredentialSchema)
	utils.Append(attr, map[string]schema.Attribute{
		"subscription_id": schema.StringAttribute{
			Description: "The Azure subscription ID. Required for secret based credentials and should look like the following example: a8d4457d-310v-41p6-sc53-14g8d733e514",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
			Required: true,
		},
		"tenant_id": schema.StringAttribute{
			Description: "The Azure AD tenant ID for the Azure subscription. Required for secret based credentials and should look like the following example: b10u3481-2451-10ba-7sfd-9o2d1v60185d",
			Required:    true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"app_based": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"application_id": schema.StringAttribute{
					Description: "The ID of the application registered in Azure.",
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					},
					Required:  true,
					Sensitive: false,
				},
				"secret_key": schema.StringAttribute{
					Description: "The client secret key (also referred to as application password) for the registered application.",
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					},
					Required:  true,
					Sensitive: true,
				},
			},
		},
	})
	resp.Schema = schema.Schema{
		MarkdownDescription: "The Azure credential is used for authorization  to provision resources such as compute instances within your cloud provider account.",
		Attributes:          attr,
	}
}
