// Copyright 2026 Cloudera. All Rights Reserved.
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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *privateLinkEndpointResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Creates a Private Link endpoint for CDP services.",
		Attributes: map[string]schema.Attribute{
			"tracking_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Tracking ID of the create endpoint request.",
			},
			"resource_tags": schema.ListNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Custom tags for cloud resources created during Private Links creation.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key":   schema.StringAttribute{Required: true},
						"value": schema.StringAttribute{Required: true},
					},
				},
			},
			"endpoint_statuses": schema.ListNestedAttribute{
				Computed:            true,
				MarkdownDescription: "Statuses of the Private Link endpoints after creation.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"service_component":  schema.StringAttribute{Computed: true},
						"status":             schema.StringAttribute{Computed: true},
						"error":              schema.StringAttribute{Computed: true},
						"endpoint_id":        schema.StringAttribute{Computed: true},
						"dns_names":          schema.ListAttribute{Computed: true, ElementType: types.StringType},
						"creation_timestamp": schema.StringAttribute{Computed: true},
					},
				},
			},
			"cloud_service_provider": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The cloud service provider. Accepted values are AWS and AZURE.",
				Validators: []validator.String{
					stringvalidator.OneOf("AWS", "AZURE"),
				},
			},
			"service_group": schema.StringAttribute{
				Optional:            true,
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
						MarkdownDescription: "AWS account ID to authorize access for Private Link.",
					},
					"cross_account_role_details": schema.SingleNestedAttribute{
						Optional:            true,
						MarkdownDescription: "Cross-account role details. Either this or `credential_crn` must be provided.",
						Attributes: map[string]schema.Attribute{
							"cross_account_role": schema.StringAttribute{
								Optional:            true,
								MarkdownDescription: "Cross-account role ARN.",
							},
							"external_id": schema.StringAttribute{
								Optional:            true,
								MarkdownDescription: "External ID associated with the cross-account role.",
							},
						},
					},
					"credential_crn": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "CDP Credential CRN to fetch the AWS cross-account RoleArn. Either this or `cross_account_role_details` must be provided.",
					},
					"region": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "AWS region in which the VPC exists.",
					},
					"vpc_id": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "VPC ID in which the Private Link endpoint should be created.",
					},
					"subnet_ids": schema.ListAttribute{
						Optional:            true,
						ElementType:         types.StringType,
						MarkdownDescription: "IDs of the private subnets in which the Private Link should be created.",
					},
				},
			},
			"azure_account_details": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Azure account details for the Private Link endpoint.",
				Attributes: map[string]schema.Attribute{
					"azure_client_secret_credential": schema.SingleNestedAttribute{
						Optional:            true,
						MarkdownDescription: "Azure client secret credential. Either this or `credential_crn` must be provided.",
						Attributes: map[string]schema.Attribute{
							"client_id": schema.StringAttribute{
								Optional:            true,
								MarkdownDescription: "The client (application) ID of the service principal.",
							},
							"client_secret": schema.StringAttribute{
								Optional:            true,
								Sensitive:           true,
								MarkdownDescription: "A client secret for the App Registration.",
							},
							"tenant_id": schema.StringAttribute{
								Optional:            true,
								MarkdownDescription: "The Azure AD tenant ID of the service principal.",
							},
						},
					},
					"credential_crn": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "CDP Credential CRN to fetch the Azure client secret credentials. Either this or `azure_client_secret_credential` must be provided.",
					},
					"subscription_id": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The Azure subscription ID for which the Private Link service is to be allowed/accessible.",
					},
					"resource_group": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The resource group under which the private endpoint is to be created.",
					},
					"location": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Azure location where the endpoint should be created.",
					},
					"vnet_id": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "vNet ID in which the Private Link should be created.",
					},
					"subnet_id": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "ID of the private subnet in which the Private Link should be created.",
					},
				},
			},
		},
	}
}
