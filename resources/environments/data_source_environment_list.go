// Copyright 2026 Cloudera. All Rights Reserved.
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
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ datasource.DataSource = &environmentListDataSource{}
)

type environmentListDataSource struct {
	client *cdp.Client
}

func NewListEnvironmentsDataSource() datasource.DataSource {
	return &environmentListDataSource{}
}

func (e *environmentListDataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_environments_list"
}

func (e *environmentListDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = EnvironmentListSchema
}

func (e *environmentListDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	e.client = utils.GetCdpClientForDataSource(req, resp)
}

func (e *environmentListDataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var data EnvironmentListModel
	state := request.Config.Get(ctx, &data)
	response.Diagnostics.Append(state...)
	if response.Diagnostics.HasError() {
		tflog.Error(ctx, "Got Error while trying to get data")
		return
	}

	environments, err := getEnvironments(ctx, e.client.Environments)

	if err != nil {
		response.Diagnostics.AddError("Error during environment collection", err.Error())
		return
	}

	data.Environments = environments.Environments

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func getEnvironments(ctx context.Context, client *client.Environments) (*EnvironmentListModel, error) {
	params := operations.NewListEnvironmentsParamsWithContext(ctx).WithInput(map[string]any{})
	resp, err := client.Operations.ListEnvironments(params)
	if err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Error during environment collection due to : %s", err.Error()))
		return nil, err
	}
	return fromListEnvironmentsResponse(resp.GetPayload()), nil
}
