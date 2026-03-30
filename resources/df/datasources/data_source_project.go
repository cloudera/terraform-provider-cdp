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

var _ datasource.DataSource = (*dfProjectDataSource)(nil)

type dfProjectDataSource struct {
	client *cdp.Client
}

type dfProjectModel struct {
	Name types.String `tfsdk:"name"`
	Crn  types.String `tfsdk:"crn"`
}

func NewDfProjectDataSource() datasource.DataSource {
	return &dfProjectDataSource{}
}

func (d *dfProjectDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_df_project"
}

func (d *dfProjectDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Look up a DataFlow project by name.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true, MarkdownDescription: "The name of the project to look up.",
			},
			"crn": schema.StringAttribute{
				Computed: true, MarkdownDescription: "The CRN of the project.",
			},
		},
	}
}

func (d *dfProjectDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = utils.GetCdpClientForDataSource(req, resp)
}

func (d *dfProjectDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data dfProjectModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	name := data.Name.ValueString()
	params := operations.NewListProjectsParamsWithContext(ctx).WithInput(&dfmodels.ListProjectsRequest{})
	result, err := d.client.Df.Operations.ListProjects(params)
	if err != nil {
		resp.Diagnostics.AddError("Error listing DataFlow projects", err.Error())
		return
	}

	for _, p := range result.GetPayload().Projects {
		if p.Name != nil && *p.Name == name {
			data.Crn = types.StringPointerValue(p.Crn)
			resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
			return
		}
	}

	resp.Diagnostics.AddError("Project not found", fmt.Sprintf("No DataFlow project found with name %q", name))
}
