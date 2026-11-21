// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package vmtypes

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
	_ datasource.DataSource                     = &vmTypesDataSource{}
	_ datasource.DataSourceWithConfigure        = &vmTypesDataSource{}
	_ datasource.DataSourceWithConfigValidators = &vmTypesDataSource{}
)

type vmTypesDataSource struct {
	client *cdp.Client
}

func NewVMTypesDataSource() datasource.DataSource {
	return &vmTypesDataSource{}
}

func (d *vmTypesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_vm_types"
}

func (d *vmTypesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = vmTypesSchema
}

func (d *vmTypesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.AtLeastOneOf(
			path.MatchRoot("credential_name"),
			path.MatchRoot("credential_crn"),
		),
	}
}

func (d *vmTypesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = utils.GetCdpClientForDataSource(req, resp)
}

func (d *vmTypesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data vmTypesModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := &environmentsmodels.ListVMTypesRequest{}
	if !data.CredentialName.IsNull() {
		input.CredentialName = data.CredentialName.ValueString()
	}
	if !data.CredentialCrn.IsNull() {
		input.CredentialCrn = data.CredentialCrn.ValueString()
	}
	if !data.Region.IsNull() {
		input.Region = data.Region.ValueString()
	}
	if !data.AvailabilityZone.IsNull() {
		input.AvailabilityZone = data.AvailabilityZone.ValueString()
	}
	if !data.Architecture.IsNull() {
		input.Architecture = data.Architecture.ValueString()
	}
	if !data.PlatformVariant.IsNull() {
		input.PlatformVariant = data.PlatformVariant.ValueString()
	}
	if !data.ResourceType.IsNull() {
		input.ResourceType = data.ResourceType.ValueString()
	}

	params := operations.NewListVMTypesParams()
	params.WithInput(input)

	tflog.Debug(ctx, "Listing VM types by credential")

	result, err := d.client.Environments.Operations.ListVMTypesContext(ctx, params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "list VM types")
		return
	}

	data.VMTypes = fromListResponse(result.GetPayload())
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
