// Copyright 2026 Cloudera. All Rights Reserved.
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
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ datasource.DataSource = &datalakeListDataSource{}
)

type datalakeListDataSource struct {
	client *cdp.Client
}

func NewListDatalakesDataSource() datasource.DataSource {
	return &datalakeListDataSource{}
}

func (e *datalakeListDataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_datalake_list"
}

func (e *datalakeListDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = DatalakeListSchema
}

func (e *datalakeListDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	e.client = utils.GetCdpClientForDataSource(req, resp)
}

func (e *datalakeListDataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var data DatalakeListModel
	state := request.Config.Get(ctx, &data)
	response.Diagnostics.Append(state...)
	if response.Diagnostics.HasError() {
		tflog.Error(ctx, "Got Error while trying to get data")
		return
	}

	dls, err := getDatalakes(ctx, e.client.Datalake)

	if err != nil {
		response.Diagnostics.AddError("Error during datalake collection", err.Error())
		return
	}

	data.Datalakes = dls.Datalakes

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func getDatalakes(ctx context.Context, client *client.Datalake) (*DatalakeListModel, error) {
	params := operations.NewListDatalakesParamsWithContext(ctx).WithInput(&models.ListDatalakesRequest{})
	resp, err := client.Operations.ListDatalakes(params)
	if err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Error during datalake collection due to : %s", err.Error()))
		return nil, err
	}
	return fromListDatalakesResponse(resp.GetPayload()), nil
}
