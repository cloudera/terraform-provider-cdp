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

var _ datasource.DataSource = (*dfServiceDataSource)(nil)

type dfServiceDataSource struct {
	client *cdp.Client
}

type dfServiceModel struct {
	Name            types.String `tfsdk:"name"`
	Crn             types.String `tfsdk:"crn"`
	EnvironmentCrn  types.String `tfsdk:"environment_crn"`
	CloudPlatform   types.String `tfsdk:"cloud_platform"`
	Region          types.String `tfsdk:"region"`
	Status          types.String `tfsdk:"status"`
	WorkloadVersion types.String `tfsdk:"workload_version"`
	DeploymentCount types.Int32  `tfsdk:"deployment_count"`
}

func NewDfServiceDataSource() datasource.DataSource {
	return &dfServiceDataSource{}
}

func (d *dfServiceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_df_service"
}

func (d *dfServiceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Look up a DataFlow service by name.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true, MarkdownDescription: "The name of the DataFlow service.",
			},
			"crn":              schema.StringAttribute{Computed: true, MarkdownDescription: "The service CRN."},
			"environment_crn":  schema.StringAttribute{Computed: true, MarkdownDescription: "The CDP environment CRN."},
			"cloud_platform":   schema.StringAttribute{Computed: true, MarkdownDescription: "The cloud platform."},
			"region":           schema.StringAttribute{Computed: true, MarkdownDescription: "The region."},
			"status":           schema.StringAttribute{Computed: true, MarkdownDescription: "The service status."},
			"workload_version": schema.StringAttribute{Computed: true, MarkdownDescription: "The workload version."},
			"deployment_count": schema.Int32Attribute{Computed: true, MarkdownDescription: "The number of deployments."},
		},
	}
}

func (d *dfServiceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = utils.GetCdpClientForDataSource(req, resp)
}

func (d *dfServiceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data dfServiceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	name := data.Name.ValueString()
	params := operations.NewListServicesParamsWithContext(ctx).WithInput(&dfmodels.ListServicesRequest{})
	result, err := d.client.Df.Operations.ListServices(params)
	if err != nil {
		resp.Diagnostics.AddError("Error listing DataFlow services", err.Error())
		return
	}

	for _, svc := range result.GetPayload().Services {
		if svc.Name != nil && *svc.Name == name {
			data.Crn = types.StringPointerValue(svc.Crn)
			data.EnvironmentCrn = types.StringPointerValue(svc.EnvironmentCrn)
			data.CloudPlatform = types.StringPointerValue(svc.CloudPlatform)
			data.Region = types.StringPointerValue(svc.Region)
			data.Status = types.StringValue(string(*svc.Status.State))
			data.WorkloadVersion = types.StringPointerValue(svc.WorkloadVersion)
			data.DeploymentCount = types.Int32PointerValue(svc.DeploymentCount)
			resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
			return
		}
	}

	resp.Diagnostics.AddError("Service not found", fmt.Sprintf("No DataFlow service found with name %q", name))
}
