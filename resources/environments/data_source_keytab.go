// Copyright 2024 Cloudera. All Rights Reserved.
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
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ datasource.DataSource = &keytabDataSource{}
)

type keytabDataSource struct {
	client *cdp.Client
}

func NewKeytabDataSource() datasource.DataSource {
	return &keytabDataSource{}
}

func (p *keytabDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_keytab"
}

func (p *keytabDataSource) Schema(_ context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = KeytabSchema
}

func (p *keytabDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	p.client = utils.GetCdpClientForDataSource(req, resp)
}

func (p *keytabDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data KeytabModel
	stuff := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(stuff...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Got Error while trying to get data")
		return
	}

	keytabs, err := fetchKeytabs(ctx, p.client.Environments, data)

	if err != nil {
		resp.Diagnostics.AddError("Error during keytab collection", err.Error())
		return
	}

	data.Keytab = keytabs.Keytab

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func fetchKeytabs(ctx context.Context, env2Client *client.Environments, state KeytabModel) (KeytabModel, error) {
	paramReq := models.GetKeytabRequest{
		EnvironmentName: state.Environment.ValueStringPointer(),
	}

	if !state.ActorCrn.IsNull() {
		paramReq.ActorCrn = state.ActorCrn.ValueString()
	}

	params := operations.NewGetKeytabParamsWithContext(ctx)
	params.WithInput(&paramReq)

	keytab, err := env2Client.Operations.GetKeytab(params)
	if err != nil {
		if err.Error() == "not found" {
			tflog.Warn(ctx, "Error during keytab collection!", map[string]interface{}{})
		}
		return state, err
	}

	state.Keytab = types.StringValue(keytab.Payload.Contents)
	return state, nil
}
