// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package networks

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ datasource.DataSource              = &networksDataSource{}
	_ datasource.DataSourceWithConfigure = &networksDataSource{}
)

type networksDataSource struct {
	client *cdp.Client
}

func NewNetworksDataSource() datasource.DataSource {
	return &networksDataSource{}
}

func (d *networksDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_networks"
}

func (d *networksDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = networksSchema
}

func (d *networksDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = utils.GetCdpClientForDataSource(req, resp)
}

func (d *networksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data networksModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := &environmentsmodels.ListNetworksRequest{
		Credential: data.Credential.ValueStringPointer(),
	}
	if !data.Region.IsNull() {
		input.Region = data.Region.ValueString()
	}
	if !data.AvailabilityZone.IsNull() {
		input.AvailabilityZone = data.AvailabilityZone.ValueString()
	}
	if !data.PlatformVariant.IsNull() {
		input.PlatformVariant = data.PlatformVariant.ValueString()
	}
	if !data.NetworkID.IsNull() {
		input.NetworkID = data.NetworkID.ValueString()
	}
	if !data.SharedProjectID.IsNull() {
		input.SharedProjectID = data.SharedProjectID.ValueString()
	}
	if !data.SubnetIds.IsNull() {
		var ids []string
		data.SubnetIds.ElementsAs(ctx, &ids, false)
		input.SubnetIds = strings.Join(ids, ",")
	}

	params := operations.NewListNetworksParams()
	params.WithInput(input)

	tflog.Debug(ctx, "Listing networks by credential")

	result, err := d.client.Environments.Operations.ListNetworksContext(ctx, params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "list networks")
		return
	}

	data.fromResponse(result.GetPayload())
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
