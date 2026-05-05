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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/resources/environments/validators"
)

var securityAccess = schema.SingleNestedAttribute{
	MarkdownDescription: "Security control configuration for FreeIPA and Datalake deployment. Choosing a CIDR will automatically create security groups. Alternatively existing security groups can be specified.",
	Description:         "Security control configuration for FreeIPA and Datalake deployment. Choosing a CIDR will automatically create security groups. Alternatively existing security groups can be specified.",
	Required:            true,
	Attributes: map[string]schema.Attribute{
		"cidr": schema.StringAttribute{
			MarkdownDescription: "CIDR range which is allowed for inbound traffic. Either IPv4 or IPv6 is allowed.",
			Description:         "CIDR range which is allowed for inbound traffic. Either IPv4 or IPv6 is allowed.",
			Optional:            true,
			Computed:            true,
			Validators: []validator.String{
				validators.CIDROrSecurityGroupsValidator(),
			},
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"default_security_group_id": schema.StringAttribute{
			MarkdownDescription: "Security group where all other hosts are placed. Mutually exclusive with cidr.",
			Description:         "Security group where all other hosts are placed. Mutually exclusive with cidr.",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"default_security_group_ids": schema.SetAttribute{
			MarkdownDescription: "Security group IDs where all other hosts are placed. Mutually exclusive with CIDR.",
			Description:         "Security group IDs where all other hosts are placed. Mutually exclusive with CIDR.",
			Optional:            true,
			ElementType:         types.StringType,
		},
		"security_group_id_for_knox": schema.StringAttribute{
			MarkdownDescription: "Security group where Knox-enabled hosts are placed. Mutually exclusive with cidr.",
			Description:         "Security group where Knox-enabled hosts are placed. Mutually exclusive with cidr.",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"security_group_ids_for_knox": schema.SetAttribute{
			MarkdownDescription: "Security group IDs where Knox-enabled hosts are placed. Mutually exclusive with CIDR.",
			Description:         "Security group IDs where Knox-enabled hosts are placed. Mutually exclusive with CIDR.",
			Optional:            true,
			ElementType:         types.StringType,
		},
	},
}
