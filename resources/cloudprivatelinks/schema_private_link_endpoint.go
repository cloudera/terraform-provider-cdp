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
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *privateLinkEndpointResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Creates a Private Link endpoint for CDP services.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The ID of the Private Link endpoint.",
			},
			"cloud_service_provider": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The cloud service provider. Accepted values are AWS and AZURE.",
			},
			"service_group": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The service group for the Private Link endpoint.",
			},
			"enable_private_dns": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Whether to enable private DNS for the endpoint.",
			},
			"aws_account_details": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "AWS account details for the Private Link endpoint.",
				Attributes: map[string]schema.Attribute{
					"cloud_account_id": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The AWS account ID.",
					},
					"credential_crn": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The CRN of the CDP credential.",
					},
					"region": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "The AWS region.",
					},
					"vpc_id": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "The VPC ID.",
					},
					"subnet_ids": schema.SetAttribute{
						Optional:            true,
						MarkdownDescription: "List of subnet IDs.",
						ElementType:         types.StringType,
					},
				},
			},
			"azure_account_details": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Azure account details for the Private Link endpoint.",
				Attributes: map[string]schema.Attribute{
					"credential_crn": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The CRN of the CDP credential.",
					},
					"subscription_id": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "The Azure subscription ID.",
					},
					"resource_group": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The Azure resource group.",
					},
					"location": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "The Azure location.",
					},
					"vnet_id": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "The VNet ID.",
					},
					"subnet_id": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The subnet ID.",
					},
				},
			},
		},
	}
}
