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

var _ datasource.DataSource = (*dfFlowDefinitionsDataSource)(nil)

type dfFlowDefinitionsDataSource struct {
	client *cdp.Client
}

type dfFlowDefinitionsModel struct {
	Flows []dfFlowItem `tfsdk:"flows"`
}

type dfFlowItem struct {
	Crn          types.String `tfsdk:"crn"`
	Name         types.String `tfsdk:"name"`
	ArtifactType types.String `tfsdk:"artifact_type"`
	VersionCount types.Int32  `tfsdk:"version_count"`
}

func NewDfFlowDefinitionsDataSource() datasource.DataSource {
	return &dfFlowDefinitionsDataSource{}
}

func (d *dfFlowDefinitionsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_df_flow_definitions"
}

func (d *dfFlowDefinitionsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Lists all DataFlow flow definitions in the catalog.",
		Attributes: map[string]schema.Attribute{
			"flows": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"crn":           schema.StringAttribute{Computed: true},
						"name":          schema.StringAttribute{Computed: true},
						"artifact_type": schema.StringAttribute{Computed: true},
						"version_count": schema.Int32Attribute{Computed: true},
					},
				},
			},
		},
	}
}

func (d *dfFlowDefinitionsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = utils.GetCdpClientForDataSource(req, resp)
}

func (d *dfFlowDefinitionsDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	params := operations.NewListFlowDefinitionsParamsWithContext(ctx).WithInput(&dfmodels.ListFlowDefinitionsRequest{})
	result, err := d.client.Df.Operations.ListFlowDefinitions(params)
	if err != nil {
		resp.Diagnostics.AddError("Error listing DataFlow flow definitions", err.Error())
		return
	}

	var data dfFlowDefinitionsModel
	for _, f := range result.GetPayload().Flows {
		data.Flows = append(data.Flows, dfFlowItem{
			Crn:          types.StringPointerValue(f.Crn),
			Name:         types.StringPointerValue(f.Name),
			ArtifactType: types.StringPointerValue(f.ArtifactType),
			VersionCount: types.Int32PointerValue(f.VersionCount),
		})
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}
