// Copyright 2023 Cloudera. All Rights Reserved.
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
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

func (r *azureDatahubResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	attr := map[string]schema.Attribute{}
	utils.Append(attr, generalAttributes)
	utils.Append(attr, azureInstanceGroupSchemaAttributes)
	utils.Append(attr, map[string]schema.Attribute{
		"cluster_template": schema.StringAttribute{
			MarkdownDescription: "The name of the cluster template.",
			Optional:            true,
		},
		"cluster_definition": schema.StringAttribute{
			MarkdownDescription: "The name of the cluster definition.",
			Optional:            true,
		},
		"environment": schema.StringAttribute{
			MarkdownDescription: "The name of the environment where the cluster will belong to.",
			Required:            true,
		},
		"database_type": schema.StringAttribute{
			Optional: true,
			Computed: false,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"subnet_id": schema.StringAttribute{
			MarkdownDescription: "The subnet id.",
			Optional:            true,
		},
		"multi_az": schema.BoolAttribute{
			MarkdownDescription: "Flag  that toggles the multi availability zone for the given datahub cluster when you are not sure what subnet IDs can be used. This  way the subnet IDs will be used what the environment suggests.",
			Optional:            true,
		},
		"enable_load_balancer": schema.BoolAttribute{
			MarkdownDescription: "Flag that decides whether to provision a load-balancer to front var- ious service endpoints for the given datahub. This will typically be used for HA cluster shapes.",
			Optional:            true,
		},
		"load_balancer_sku": schema.StringAttribute{
			MarkdownDescription: "Represents  the Azure load balancer SKU type. The current default is BASIC. To disable the load balancer, use type NONE.",
			Optional:            true,
		},
		"flexible_server_delegated_subnet_id": schema.StringAttribute{
			MarkdownDescription: "Allows you to specify the subnet ID for the subnet within which  you want to configure your Azure Flexible Server.",
			Optional:            true,
		},
	})
	resp.Schema = schema.Schema{
		MarkdownDescription: "Creates an Azure Data hub cluster.",
		Attributes:          attr,
	}
}
