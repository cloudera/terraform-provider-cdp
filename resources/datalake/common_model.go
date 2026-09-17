// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datalake

import "github.com/hashicorp/terraform-plugin-framework/types"

type Security struct {
	SeLinux types.String `tfsdk:"se_linux"`
}

type DeleteOptions struct {
	Forced types.Bool `tfsdk:"forced"`
}

type UpgradeOptions struct {
	DryRun                types.Bool `tfsdk:"dry_run"`
	LockComponents        types.Bool `tfsdk:"lock_components"`
	RollingUpgradeEnabled types.Bool `tfsdk:"rolling_upgrade_enabled"`
	SkipBackup            types.Bool `tfsdk:"skip_backup"`
	SkipBackupValidation  types.Bool `tfsdk:"skip_backup_validation"`
	SkipAtlasMetadata     types.Bool `tfsdk:"skip_atlas_metadata"`
	SkipRangerAudits      types.Bool `tfsdk:"skip_ranger_audits"`
	SkipRangerHmsMetadata types.Bool `tfsdk:"skip_ranger_hms_metadata"`
	SkipDatahubValidation types.Bool `tfsdk:"skip_datahub_validation"`
}
