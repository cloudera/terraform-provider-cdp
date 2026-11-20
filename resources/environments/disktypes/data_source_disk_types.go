// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package disktypes

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ datasource.DataSource              = &diskTypesDataSource{}
	_ datasource.DataSourceWithConfigure = &diskTypesDataSource{}
)

type diskTypesDataSource struct {
	client *cdp.Client
}

func NewDiskTypesDataSource() datasource.DataSource {
	return &diskTypesDataSource{}
}

func (d *diskTypesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_disk_types"
}

func (d *diskTypesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = diskTypesSchema
}

func (d *diskTypesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = utils.GetCdpClientForDataSource(req, resp)
}

func (d *diskTypesDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	params := operations.NewListDiskTypesParams()

	tflog.Debug(ctx, "Listing disk types")

	result, err := d.client.Environments.Operations.ListDiskTypesContext(ctx, params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "list disk types")
		return
	}

	data := diskTypesModelFromResponse(result.GetPayload())
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
