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

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client/operations"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ datasource.DataSource = &runtimeDataSource{}
)

type runtimeDataSource struct {
	client *cdp.Client
}

func NewListRuntimeDataSource() datasource.DataSource {
	return &runtimeDataSource{}
}

func (p *runtimeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_datalake_list_runtimes"
}

func (p *runtimeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = RuntimeSchema
}

func (p *runtimeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	p.client = utils.GetCdpClientForDataSource(req, resp)
}

func (p *runtimeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data RuntimeModel
	stuff := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(stuff...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Got Error while trying to get data")
		return
	}

	runtimes, err := fetchRuntimes(ctx, p.client.Datalake)

	if err != nil {
		resp.Diagnostics.AddError("Error during runtime collection", err.Error())
		return
	}

	data.Versions = runtimes.Versions

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func fetchRuntimes(ctx context.Context, client *client.Datalake) (*RuntimeModel, error) {
	params := operations.NewListRuntimesParamsWithContext(ctx)
	resp, err := client.Operations.ListRuntimes(params)
	if err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Error during runtime collection due to : %s", err.Error()))
		return nil, err
	}
	return fromListRuntimesResponse(resp.GetPayload()), nil
}
