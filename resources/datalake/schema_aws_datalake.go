// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datalake

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

var awsDatalakeSchema = getAwsResourceSchema()

func getAwsResourceSchema() schema.Schema {
	attr := map[string]schema.Attribute{}
	utils.AppendToResourceSchema(attr, generalAttributes)
	utils.AppendToResourceSchema(attr, map[string]schema.Attribute{
		"certificate_expiration_state": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"storage_location_base": schema.StringAttribute{
			Required: true,
		},
		"instance_profile": schema.StringAttribute{
			Required: true,
		},
		"enable_ranger_rms": schema.BoolAttribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "Whether to enable Ranger RMS for the datalake. Defaults to not being enabled.",
			Default:             booldefault.StaticBool(false),
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"architecture": schema.StringAttribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "Specifies the CPU architecture of the data lake cluster. Accepted values are `ARM64` and `X86_64`.",
			Validators: []validator.String{
				stringvalidator.OneOf("ARM64", "X86_64"),
			},
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
	})
	return schema.Schema{
		MarkdownDescription: "A Data Lake is a service which provides a protective ring around the data stored in a cloud object store, including authentication, authorization, and governance support.",
		Attributes:          attr,
	}
}
