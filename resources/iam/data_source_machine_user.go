// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package iam

import (
	"context"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSourceWithConfigure = &machineUserDataSource{}
)

func NewMachineUserDataSource() datasource.DataSource {
	return &machineUserDataSource{}
}

type machineUserDataSource struct {
	client *cdp.Client
}

// Configure implements datasource.DataSourceWithConfigure.
func (m *machineUserDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	m.client = utils.GetCdpClientForDataSource(req, resp)
}

// Metadata implements datasource.DataSourceWithConfigure.
func (m *machineUserDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_machine_user"
}

// Schema implements datasource.DataSourceWithConfigure.
func (m *machineUserDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id":                        schema.StringAttribute{Computed: true},
			"machine_user_name":         schema.StringAttribute{MarkdownDescription: "The machine user name.", Required: true},
			"crn":                       schema.StringAttribute{MarkdownDescription: "The CRN of the user.", Computed: true},
			"creation_date":             schema.StringAttribute{MarkdownDescription: "The date when this machine user was created.", Computed: true},
			"status":                    schema.StringAttribute{MarkdownDescription: "The current status of the machine user.", Computed: true},
			"workload_username":         schema.StringAttribute{MarkdownDescription: "The username used in all the workload clusters of the machine user.", Optional: true},
			"workload_password_details": schema.StringAttribute{MarkdownDescription: "Information about the workload password for the machine user.", Optional: true},
		},
		Blocks:              map[string]schema.Block{},
		Description:         "",
		MarkdownDescription: "A machine user account provides programmatic access to CDP.",
		DeprecationMessage:  "",
	}
}

// Read refreshes the Terraform state with the latest data.
func (m *machineUserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// Read user supplied values
	var data machineUserModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	sharedMachineUserRead(ctx, m.client.Iam, &data, &resp.State, &resp.Diagnostics)
}
