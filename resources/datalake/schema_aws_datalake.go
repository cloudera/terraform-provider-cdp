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
			MarkdownDescription: "Indicates the certificate status on the cluster.",
			Description:         "Indicates the certificate status on the cluster.",
			Validators: []validator.String{
				stringvalidator.OneOf("VALID", "HOST_CERT_EXPIRING"),
			},
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"storage_location_base": schema.StringAttribute{
			MarkdownDescription: "The location of the S3 bucket to be used as storage. The location has to start with s3a:// followed by the bucket name.",
			Description:         "The location of the S3 bucket to be used as storage. The location has to start with s3a:// followed by the bucket name.",
			Required:            true,
		},
		"instance_profile": schema.StringAttribute{
			MarkdownDescription: "The ARN of an IAM instance profile.",
			Description:         "The ARN of an IAM instance profile.",
			Required:            true,
		},
		"enable_ranger_rms": schema.BoolAttribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "Whether to enable Ranger RMS for the datalake. Defaults to not being enabled.",
			Description:         "Whether to enable Ranger RMS for the datalake. Defaults to not being enabled.",
			Default:             booldefault.StaticBool(false),
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"architecture": schema.StringAttribute{
			Optional:            true,
			Computed:            false,
			MarkdownDescription: "Specifies the CPU architecture of the data lake cluster. Accepted values are `ARM64` and `X86_64`.",
			Description:         "Specifies the CPU architecture of the data lake cluster. Accepted values are `ARM64` and `X86_64`.",
			Validators: []validator.String{
				stringvalidator.OneOf("ARM64", "X86_64"),
			},
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"upgrade_options": schema.SingleNestedAttribute{
			MarkdownDescription: "Options controlling datalake upgrade behavior. An upgrade is triggered when `runtime` or `image` changes. Set `dry_run = true` to preview available upgrades without applying them.",
			Description:         "Options controlling datalake upgrade behavior. An upgrade is triggered when runtime or image changes. Set dry_run = true to preview available upgrades without applying them.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"dry_run": schema.BoolAttribute{
					MarkdownDescription: "Checks the eligibility of an image to upgrade without performing the upgrade. When set, the provider reports available upgrade candidates as a warning and preserves the current runtime and image values in state.",
					Description:         "Checks the eligibility of an image to upgrade without performing the upgrade. When set, the provider reports available upgrade candidates as a warning and preserves the current runtime and image values in state.",
					Optional:            true,
				},
				"lock_components": schema.BoolAttribute{
					MarkdownDescription: "Perform an OS upgrade only, locking Cloudera component versions.",
					Description:         "Perform an OS upgrade only, locking Cloudera component versions.",
					Optional:            true,
				},
				"rolling_upgrade_enabled": schema.BoolAttribute{
					MarkdownDescription: "Enables the ability to perform a rolling runtime upgrade.",
					Description:         "Enables the ability to perform a rolling runtime upgrade.",
					Optional:            true,
				},
				"skip_backup": schema.BoolAttribute{
					MarkdownDescription: "Skip the backup step before upgrade.",
					Description:         "Skip the backup step before upgrade.",
					Optional:            true,
				},
				"skip_backup_validation": schema.BoolAttribute{
					MarkdownDescription: "Skip validation steps that run before backup. Redundant if skip_backup is set.",
					Description:         "Skip validation steps that run before backup. Redundant if skip_backup is set.",
					Optional:            true,
				},
				"skip_atlas_metadata": schema.BoolAttribute{
					MarkdownDescription: "Skip backup of Atlas metadata. Redundant if skip_backup is set.",
					Description:         "Skip backup of Atlas metadata. Redundant if skip_backup is set.",
					Optional:            true,
				},
				"skip_ranger_audits": schema.BoolAttribute{
					MarkdownDescription: "Skip backup of Ranger audit logs. Redundant if skip_backup is set.",
					Description:         "Skip backup of Ranger audit logs. Redundant if skip_backup is set.",
					Optional:            true,
				},
				"skip_ranger_hms_metadata": schema.BoolAttribute{
					MarkdownDescription: "Skip backup of HMS/Ranger databases. Redundant if skip_backup is set.",
					Description:         "Skip backup of HMS/Ranger databases. Redundant if skip_backup is set.",
					Optional:            true,
				},
				"skip_datahub_validation": schema.BoolAttribute{
					MarkdownDescription: "Allow upgrade with running DataHub clusters. This may cause issues on running DataHub clusters during upgrade.",
					Description:         "Allow upgrade with running DataHub clusters. This may cause issues on running DataHub clusters during upgrade.",
					Optional:            true,
				},
			},
		},
	})
	return schema.Schema{
		MarkdownDescription: "A Data Lake is a service which provides a protective ring around the data stored in a cloud object store, including authentication, authorization, and governance support.",
		Description:         "A Data Lake is a service which provides a protective ring around the data stored in a cloud object store, including authentication, authorization, and governance support.",
		Attributes:          attr,
	}
}
