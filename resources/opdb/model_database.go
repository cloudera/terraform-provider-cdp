// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package opdb

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

type databaseResourceModel struct {
	Crn               types.String `tfsdk:"crn"`
	DatabaseName      types.String `tfsdk:"database_name"`
	Status            types.String `tfsdk:"status"`
	Environment       types.String `tfsdk:"environment_name"`
	ScaleType         types.String `tfsdk:"scale_type"`
	StorageType       types.String `tfsdk:"storage_type"`
	DisableExternalDB types.Bool   `tfsdk:"disable_external_db"`
	StorageLocation   types.String `tfsdk:"storage_location"`

	PollingOptions *utils.PollingOptions `tfsdk:"polling_options"`

	DisableMultiAz types.Bool   `tfsdk:"disable_multi_az"`
	NumEdgeNodes   types.Int64  `tfsdk:"num_edge_nodes"`
	JavaVersion    types.Int64  `tfsdk:"java_version"`
	SubnetID       types.String `tfsdk:"subnet_id"`
}
