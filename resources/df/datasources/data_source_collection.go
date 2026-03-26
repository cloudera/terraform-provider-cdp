// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/df/client/operations"
	dfmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/df/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var _ datasource.DataSource = (*dfCollectionDataSource)(nil)

type dfCollectionDataSource struct {
	client *cdp.Client
}

type dfCollectionModel struct {
	Name        types.String `tfsdk:"name"`
	Crn         types.String `tfsdk:"crn"`
	Description types.String `tfsdk:"description"`
}

func NewDfCollectionDataSource() datasource.DataSource {
	return &dfCollectionDataSource{}
}

func (d *dfCollectionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_df_collection"
}

func (d *dfCollectionDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Look up a DataFlow catalog collection by name.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true, MarkdownDescription: "The name of the collection to look up.",
			},
			"crn": schema.StringAttribute{
				Computed: true, MarkdownDescription: "The CRN of the collection.",
			},
			"description": schema.StringAttribute{
				Computed: true, MarkdownDescription: "The description of the collection.",
			},
		},
	}
}

func (d *dfCollectionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = utils.GetCdpClientForDataSource(req, resp)
}

func (d *dfCollectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data dfCollectionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	name := data.Name.ValueString()
	params := operations.NewListCollectionsParamsWithContext(ctx).WithInput(&dfmodels.ListCollectionsRequest{})
	result, err := d.client.Df.Operations.ListCollections(params)
	if err != nil {
		resp.Diagnostics.AddError("Error listing DataFlow collections", err.Error())
		return
	}

	for _, c := range result.GetPayload().Collections {
		if c.Name == name {
			data.Crn = types.StringValue(c.Crn)
			if c.Description != "" {
				data.Description = types.StringValue(c.Description)
			}
			resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
			return
		}
	}

	resp.Diagnostics.AddError("Collection not found", fmt.Sprintf("No DataFlow collection found with name %q", name))
}
