// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package cloudprivatelinks

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var PrivateLinkServiceAccessSchema = schema.Schema{
	MarkdownDescription: "Authorizes Private Link service access for a cloud account. On destroy, access is revoked.",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Unique identifier composed of cloud_service_provider, region, service_group and cloud account.",
		},
		"cloud_service_provider": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "Cloud Service Provider. Currently supporting `AWS` and `AZURE`.",
		},
		"cloud_account_id": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "AWS account ID to authorize access for Private Link.",
		},
		"subscription_id": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "Azure subscription ID to authorize access for Private Link.",
		},
		"region": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "Region where the Private Link service exists.",
		},
		"service_group": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: `CDP service group. Currently supported "CDP-CONTROL-PLANE" for PaaS.`,
		},
		"status": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Status of Private Link service access authorization.",
		},
	},
}

func (r *privateLinkServiceAccessResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = PrivateLinkServiceAccessSchema
}
