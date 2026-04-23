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

var _ datasource.DataSource = (*dfServicesDataSource)(nil)

type dfServicesDataSource struct {
	client *cdp.Client
}

type dfServicesModel struct {
	Name     types.String   `tfsdk:"name"`
	Services []dfServiceItem `tfsdk:"services"`
}

type dfServiceItem struct {
	Crn             types.String `tfsdk:"crn"`
	Name            types.String `tfsdk:"name"`
	EnvironmentCrn  types.String `tfsdk:"environment_crn"`
	CloudPlatform   types.String `tfsdk:"cloud_platform"`
	Region          types.String `tfsdk:"region"`
	Status          types.String `tfsdk:"status"`
	WorkloadVersion types.String `tfsdk:"workload_version"`
	DeploymentCount types.Int32  `tfsdk:"deployment_count"`
}

func NewDfServicesDataSource() datasource.DataSource {
	return &dfServicesDataSource{}
}

func (d *dfServicesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_df_services"
}

func (d *dfServicesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Lists DataFlow services. Optionally filter by name.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Optional: true, MarkdownDescription: "Filter services by name. If specified, only the matching service is returned.",
			},
			"services": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"crn":              schema.StringAttribute{Computed: true},
						"name":             schema.StringAttribute{Computed: true},
						"environment_crn":  schema.StringAttribute{Computed: true},
						"cloud_platform":   schema.StringAttribute{Computed: true},
						"region":           schema.StringAttribute{Computed: true},
						"status":           schema.StringAttribute{Computed: true},
						"workload_version": schema.StringAttribute{Computed: true},
						"deployment_count": schema.Int32Attribute{Computed: true},
					},
				},
			},
		},
	}
}

func (d *dfServicesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = utils.GetCdpClientForDataSource(req, resp)
}

func (d *dfServicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data dfServicesModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewListServicesParamsWithContext(ctx).WithInput(&dfmodels.ListServicesRequest{})
	result, err := d.client.Df.Operations.ListServices(params)
	if err != nil {
		resp.Diagnostics.AddError("Error listing DataFlow services", err.Error())
		return
	}

	nameFilter := data.Name.ValueString()
	for _, svc := range result.GetPayload().Services {
		if nameFilter != "" && (svc.Name == nil || *svc.Name != nameFilter) {
			continue
		}
		data.Services = append(data.Services, dfServiceItem{
			Crn:             types.StringPointerValue(svc.Crn),
			Name:            types.StringPointerValue(svc.Name),
			EnvironmentCrn:  types.StringPointerValue(svc.EnvironmentCrn),
			CloudPlatform:   types.StringPointerValue(svc.CloudPlatform),
			Region:          types.StringPointerValue(svc.Region),
			Status:          types.StringValue(string(*svc.Status.State)),
			WorkloadVersion: types.StringPointerValue(svc.WorkloadVersion),
			DeploymentCount: types.Int32PointerValue(svc.DeploymentCount),
		})
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}
