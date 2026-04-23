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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/df/client/operations"
	dfmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/df/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var _ datasource.DataSource = (*dfReadyflowsDataSource)(nil)

type dfReadyflowsDataSource struct {
	client *cdp.Client
}

type dfReadyflowsModel struct {
	Readyflows []dfReadyflowItem `tfsdk:"readyflows"`
}

type dfReadyflowItem struct {
	ReadyflowCrn types.String `tfsdk:"readyflow_crn"`
	Name         types.String `tfsdk:"name"`
	Summary      types.String `tfsdk:"summary"`
	Source       types.String `tfsdk:"source"`
	Destination  types.String `tfsdk:"destination"`
	Imported     types.Bool   `tfsdk:"imported"`
}

func NewDfReadyflowsDataSource() datasource.DataSource {
	return &dfReadyflowsDataSource{}
}

func (d *dfReadyflowsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_df_readyflows"
}

func (d *dfReadyflowsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Lists all DataFlow ReadyFlows available in the catalog.",
		Attributes: map[string]schema.Attribute{
			"readyflows": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"readyflow_crn": schema.StringAttribute{Computed: true},
						"name":          schema.StringAttribute{Computed: true},
						"summary":       schema.StringAttribute{Computed: true},
						"source":        schema.StringAttribute{Computed: true},
						"destination":   schema.StringAttribute{Computed: true},
						"imported":      schema.BoolAttribute{Computed: true},
					},
				},
			},
		},
	}
}

func (d *dfReadyflowsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = utils.GetCdpClientForDataSource(req, resp)
}

func (d *dfReadyflowsDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	params := operations.NewListReadyflowsParamsWithContext(ctx).WithInput(&dfmodels.ListReadyflowsRequest{})
	result, err := d.client.Df.Operations.ListReadyflows(params)
	if err != nil {
		resp.Diagnostics.AddError("Error listing DataFlow ReadyFlows", err.Error())
		return
	}

	var data dfReadyflowsModel
	for _, rf := range result.GetPayload().Readyflows {
		data.Readyflows = append(data.Readyflows, dfReadyflowItem{
			ReadyflowCrn: types.StringPointerValue(rf.ReadyflowCrn),
			Name:         types.StringPointerValue(rf.Name),
			Summary:      types.StringValue(rf.Summary),
			Source:       types.StringValue(rf.Source),
			Destination:  types.StringValue(rf.Destination),
			Imported:     types.BoolValue(rf.Imported),
		})
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}
