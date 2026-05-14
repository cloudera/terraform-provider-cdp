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
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	datalakevalidators "github.com/cloudera/terraform-provider-cdp/resources/datalake/validators"
)

var generalAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"creation_date": schema.StringAttribute{
		MarkdownDescription: "The date when the datalake was created.",
		Description:         "The date when the datalake was created.",
		Computed:            true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"crn": schema.StringAttribute{
		MarkdownDescription: "The CRN of the datalake.",
		Description:         "The CRN of the datalake.",
		Computed:            true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"datalake_name": schema.StringAttribute{
		MarkdownDescription: "The datalake name. This name must be unique, must have between 5 and 40 characters, and must contain only lowercase letters, numbers and hyphens. Names are case-sensitive.",
		Description:         "The datalake name. This name must be unique, must have between 5 and 40 characters, and must contain only lowercase letters, numbers and hyphens. Names are case-sensitive.",
		Required:            true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(5, 40),
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[a-z0-9-]+$`),
				"must contain only lowercase letters, numbers and hyphens",
			),
		},
	},
	"enable_ranger_raz": schema.BoolAttribute{
		MarkdownDescription: "Whether to enable Ranger RAZ for the datalake. Defaults to not being enabled.",
		Description:         "Whether to enable Ranger RAZ for the datalake. Defaults to not being enabled.",
		Optional:            true,
		Computed:            true,
		PlanModifiers: []planmodifier.Bool{
			boolplanmodifier.UseStateForUnknown(),
		},
	},
	"environment_crn": schema.StringAttribute{
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"environment_name": schema.StringAttribute{
		MarkdownDescription: "The name or CRN of the environment where the datalake will be created. Either this - old - or the new `environment` field must be provided, but not both or neither of them.",
		Description:         "The name or CRN of the environment where the datalake will be created.",
		DeprecationMessage:  "This attribute is deprecated. Use 'environment' instead.",
		Validators: []validator.String{
			datalakevalidators.EnvironmentReferenceValidator(),
		},
		Optional: true,
	},
	"environment": schema.StringAttribute{
		MarkdownDescription: "The name or CRN of the environment where the datalake will be created.",
		Description:         "The name or CRN of the environment where the datalake will be created.",
		Optional:            true,
	},
	"image": schema.SingleNestedAttribute{
		MarkdownDescription: "The image to use for the datalake. This must not be set if the runtime parameter is provided. When the 'runtime' parameter is set, only the 'os' parameter can be provided. Otherwise, you can use 'catalog name' and/or 'id' for selecting an image.",
		Description:         "The image to use for the datalake. This must not be set if the runtime parameter is provided. When the 'runtime' parameter is set, only the 'os' parameter can be provided. Otherwise, you can use 'catalog name' and/or 'id' for selecting an image.",
		Optional:            true,
		Validators: []validator.Object{
			datalakevalidators.ImageRuntimeCompatibilityValidator(),
		},
		Attributes: map[string]schema.Attribute{
			"catalog_name": schema.StringAttribute{
				MarkdownDescription: "The name of the custom image catalog to use, defaulting to 'cdp-default' if not present.",
				Description:         "The name of the custom image catalog to use, defaulting to 'cdp-default' if not present.",
				Default:             stringdefault.StaticString("cdp-default"),
				Computed:            true,
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The image ID from the catalog. The corresponding image will be used for the created cluster machines.",
				Description:         "The image ID from the catalog. The corresponding image will be used for the created cluster machines.",
				Optional:            true,
			},
			"os": schema.StringAttribute{
				MarkdownDescription: "The OS of the image used for cluster instances.",
				Description:         "The OS of the image used for cluster instances.",
				Optional:            true,
			},
		},
	},
	"java_version": schema.Int32Attribute{
		MarkdownDescription: "Configure the major version of Java on the cluster.",
		Description:         "Configure the major version of Java on the cluster.",
		Optional:            true,
	},
	"recipes": schema.SetNestedAttribute{
		MarkdownDescription: "Additional recipes that will be attached on the datalake instances (by instance groups, most common ones are like 'master' or 'idbroker').",
		Description:         "Additional recipes that will be attached on the datalake instances (by instance groups, most common ones are like 'master' or 'idbroker').",
		Optional:            true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"instance_group_name": schema.StringAttribute{
					MarkdownDescription: "The name of the designated instance group.",
					Description:         "The name of the designated instance group.",
					Required:            true,
				},
				"recipe_names": schema.SetAttribute{
					MarkdownDescription: "The set of recipe names that are going to be applied on the given instance group.",
					Description:         "The set of recipe names that are going to be applied on the given instance group.",
					ElementType:         types.StringType,
					Required:            true,
				},
			},
		},
	},
	"custom_instance_groups": schema.SetNestedAttribute{
		MarkdownDescription: "Request object for host group level custom configurations.",
		Description:         "Request object for host group level custom configurations.",
		Optional:            true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"name": schema.StringAttribute{
					MarkdownDescription: "The name of the custom instance group.",
					Description:         "The name of the custom instance group.",
					Required:            true,
				},
				"instance_type": schema.StringAttribute{
					MarkdownDescription: "The instance type for the custom instance group.",
					Description:         "The instance type for the custom instance group.",
					Required:            true,
				},
			},
		},
	},
	"runtime": schema.StringAttribute{
		MarkdownDescription: "Cloudera Runtime version.",
		Description:         "Cloudera Runtime version.",
		Optional:            true,
	},
	"scale": schema.StringAttribute{
		MarkdownDescription: "Represents the available datalake scales. Defaults to LIGHT_DUTY if not set.",
		Description:         "Represents the available datalake scales. Defaults to LIGHT_DUTY if not set.",
		Validators: []validator.String{
			stringvalidator.OneOf("LIGHT_DUTY", "MEDIUM_DUTY_HA", "ENTERPRISE"),
		},
		Default:  stringdefault.StaticString("LIGHT_DUTY"),
		Computed: true,
		Optional: true,
	},
	"status": schema.StringAttribute{
		MarkdownDescription: "The status of the datalake.",
		Description:         "The status of the datalake.",
		Computed:            true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"status_reason": schema.StringAttribute{
		MarkdownDescription: "The reason for the status of the datalake.",
		Description:         "The reason for the status of the datalake.",
		Computed:            true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"multi_az": schema.BoolAttribute{
		MarkdownDescription: "Controls if the datalake is deployed in a multi-availability zone way.",
		Description:         "Controls if the datalake is deployed in a multi-availability zone way.",
		Optional:            true,
		Computed:            true,
		PlanModifiers: []planmodifier.Bool{
			boolplanmodifier.UseStateForUnknown(),
		},
	},
	"tags": schema.MapAttribute{
		MarkdownDescription: "Tags to be added to Data Lake related resources.",
		Description:         "Tags to be added to Data Lake related resources.",
		Optional:            true,
		ElementType:         types.StringType,
	},
	"security": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Security related configuration for Datalake.",
		Attributes: map[string]schema.Attribute{
			"se_linux": schema.StringAttribute{
				MarkdownDescription: "Override default SELinux configuration which is PERMISSIVE by default. Values are: PERMISSIVE, ENFORCING",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("PERMISSIVE"),
				Validators: []validator.String{
					stringvalidator.OneOf("PERMISSIVE", "ENFORCING"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	},
	"delete_options": schema.SingleNestedAttribute{
		MarkdownDescription: "Options for deleting the Datalake.",
		Optional:            true,
		Attributes: map[string]schema.Attribute{
			"forced": schema.BoolAttribute{
				MarkdownDescription: "Whether the datalake should be force deleted. This option can be used when cluster deletion fails. This removes the entry from Cloudera Datalake service. Any lingering resources have to be deleted from the cloud provider manually. The default is false.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	},
	"polling_options": schema.SingleNestedAttribute{
		MarkdownDescription: "Polling related configuration options that could specify various values that will be used during CDP resource creation.",
		Description:         "Polling related configuration options that could specify various values that will be used during CDP resource creation.",
		Optional:            true,
		Attributes: map[string]schema.Attribute{
			"async": schema.BoolAttribute{
				MarkdownDescription: "Boolean value that specifies if Terraform should wait for resource creation/deletion.",
				Description:         "Boolean value that specifies if Terraform should wait for resource creation/deletion.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"polling_timeout": schema.Int64Attribute{
				MarkdownDescription: "Timeout value in minutes that specifies for how long should the polling go for resource creation/deletion.",
				Description:         "Timeout value in minutes that specifies for how long should the polling go for resource creation/deletion.",
				Default:             int64default.StaticInt64(60),
				Computed:            true,
				Optional:            true,
			},
			"call_failure_threshold": schema.Int64Attribute{
				MarkdownDescription: "Threshold value that specifies how many times should a single call failure happen before giving up the polling.",
				Description:         "Threshold value that specifies how many times should a single call failure happen before giving up the polling.",
				Default:             int64default.StaticInt64(3),
				Computed:            true,
				Optional:            true,
			},
		},
	},
}
