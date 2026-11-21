// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package regions

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ datasource.DataSource                     = &regionsDataSource{}
	_ datasource.DataSourceWithConfigure        = &regionsDataSource{}
	_ datasource.DataSourceWithConfigValidators = &regionsDataSource{}
)

type regionsDataSource struct {
	client *cdp.Client
}

func NewRegionsDataSource() datasource.DataSource {
	return &regionsDataSource{}
}

func (d *regionsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_regions"
}

func (d *regionsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = regionsSchema
}

func (d *regionsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.AtLeastOneOf(
			path.MatchRoot("credential"),
			path.MatchRoot("environment_crn"),
		),
	}
}

func (d *regionsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = utils.GetCdpClientForDataSource(req, resp)
}

func (d *regionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data regionsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := &environmentsmodels.ListRegionsRequest{}
	if !data.Credential.IsNull() {
		input.Credential = data.Credential.ValueString()
	}
	if !data.EnvironmentCrn.IsNull() {
		input.EnvironmentCrn = data.EnvironmentCrn.ValueString()
	}
	if !data.Region.IsNull() {
		input.Region = data.Region.ValueString()
	}
	if !data.AvailabilityZone.IsNull() {
		input.AvailabilityZone = data.AvailabilityZone.ValueString()
	}
	if !data.AvailabilityZonesNeeded.IsNull() {
		input.AvailabilityZonesNeeded = data.AvailabilityZonesNeeded.ValueBool()
	}
	if !data.PlatformVariant.IsNull() {
		input.PlatformVariant = data.PlatformVariant.ValueString()
	}

	params := operations.NewListRegionsParams()
	params.WithInput(input)

	tflog.Debug(ctx, "Listing regions by credential")

	result, err := d.client.Environments.Operations.ListRegionsContext(ctx, params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "list regions")
		return
	}

	data.fromResponse(result.GetPayload())
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
