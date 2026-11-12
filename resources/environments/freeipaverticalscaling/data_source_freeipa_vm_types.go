// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package freeipaverticalscaling

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ datasource.DataSource              = &freeipaVMTypesDataSource{}
	_ datasource.DataSourceWithConfigure = &freeipaVMTypesDataSource{}
)

type freeipaVMTypesDataSource struct {
	client *cdp.Client
}

func NewFreeipaVMTypesForVerticalScalingDataSource() datasource.DataSource {
	return &freeipaVMTypesDataSource{}
}

func (d *freeipaVMTypesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_freeipa_vm_types_for_vertical_scaling"
}

func (d *freeipaVMTypesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = freeipaVMTypesSchema
}

func (d *freeipaVMTypesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = utils.GetCdpClientForDataSource(req, resp)
}

func (d *freeipaVMTypesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data freeipaVMTypesModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := &environmentsmodels.ListFreeipaVMTypesForVerticalScalingRequest{
		EnvironmentCrn: data.EnvironmentCrn.ValueStringPointer(),
	}
	if !data.Architecture.IsNull() {
		input.Architecture = data.Architecture.ValueString()
	}
	if !data.InstanceType.IsNull() {
		input.InstanceType = data.InstanceType.ValueString()
	}
	if !data.ResourceType.IsNull() {
		input.ResourceType = data.ResourceType.ValueString()
	}
	if !data.AvailabilityZones.IsNull() {
		var azs []string
		resp.Diagnostics.Append(data.AvailabilityZones.ElementsAs(ctx, &azs, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		input.AvailabilityZones = azs
	}

	params := operations.NewListFreeipaVMTypesForVerticalScalingParams()
	params.WithInput(input)

	tflog.Debug(ctx, "Listing FreeIPA VM types for vertical scaling", map[string]any{
		"environmentCrn": data.EnvironmentCrn.ValueString(),
	})

	result, err := d.client.Environments.Operations.ListFreeipaVMTypesForVerticalScalingContext(ctx, params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "list FreeIPA VM types for vertical scaling")
		return
	}

	data.VMTypes = fromListResponse(result.GetPayload())
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
