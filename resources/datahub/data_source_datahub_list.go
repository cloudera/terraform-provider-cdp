// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datahub

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ datasource.DataSource = &datahubListDataSource{}
)

type datahubListDataSource struct {
	client *cdp.Client
}

func NewListDatahubsDataSource() datasource.DataSource {
	return &datahubListDataSource{}
}

func (e *datahubListDataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_datahub_list"
}

func (e *datahubListDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = DatahubListSchema
}

func (e *datahubListDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	e.client = utils.GetCdpClientForDataSource(req, resp)
}

func (e *datahubListDataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var data DatahubListModel
	state := request.Config.Get(ctx, &data)
	response.Diagnostics.Append(state...)
	if response.Diagnostics.HasError() {
		tflog.Error(ctx, "Got Error while trying to get data")
		return
	}

	dls, err := getDatahubs(ctx, e.client.Datahub)

	if err != nil {
		response.Diagnostics.AddError("Error during datahub collection", err.Error())
		return
	}

	data.Datahubs = dls.Datahubs

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func getDatahubs(ctx context.Context, client *client.Datahub) (*DatahubListModel, error) {
	params := operations.NewListClustersParamsWithContext(ctx).WithInput(&models.ListClustersRequest{})
	resp, err := client.Operations.ListClusters(params)
	if err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Error during datahub collection due to : %s", err.Error()))
		return nil, err
	}
	return fromListClustersResponse(resp.GetPayload()), nil
}
