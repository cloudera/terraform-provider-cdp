// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package auditcredential

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

func (r *azureAuditCredentialResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	attr := map[string]schema.Attribute{}
	utils.AppendToResourceSchema(attr, auditCredentialCommonSchema)
	utils.AppendToResourceSchema(attr, map[string]schema.Attribute{
		"subscription_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The Azure subscription ID.",
		},
		"tenant_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The Azure AD tenant ID for the Azure subscription.",
		},
		"app_based": schema.SingleNestedAttribute{
			Required:            true,
			MarkdownDescription: "Additional configurations needed for app-based authentication.",
			Attributes: map[string]schema.Attribute{
				"application_id": schema.StringAttribute{
					Required:            true,
					MarkdownDescription: "The ID of the application registered in Azure.",
				},
				"secret_key": schema.StringAttribute{
					Required:            true,
					Sensitive:           true,
					MarkdownDescription: "The client secret key (also referred to as application password) for the registered application.",
				},
			},
		},
	})
	resp.Schema = schema.Schema{
		MarkdownDescription: "The Azure audit credential is used for MSA (Monitoring, SPI, and Auditing) access to your Azure account.",
		Attributes:          attr,
	}
}
