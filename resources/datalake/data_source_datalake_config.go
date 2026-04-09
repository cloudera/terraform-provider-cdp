// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datalake

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ datasource.DataSource = &datalakeConfigDataSource{}
)

type datalakeConfigDataSource struct {
	client *cdp.Client
}

func NewDatalakeConfigDataSource() datasource.DataSource {
	return &datalakeConfigDataSource{}
}

func (e *datalakeConfigDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_datalake_config"
}

func (e *datalakeConfigDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	e.client = utils.GetCdpClientForDataSource(req, resp)
}

func (e *datalakeConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DatalakeConfigModel
	config := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(config...)

	var datalake string
	if utils.IsStringTypeHasValue(data.Crn) && utils.IsStringTypeHasValue(data.Name) {
		resp.Diagnostics.AddError("Too many identifier provided.", "Only the datalake\\'s name or its CRN should be given but not both.")
		return
	}
	if !utils.IsStringTypeHasValue(data.Crn) && !utils.IsStringTypeHasValue(data.Name) {
		resp.Diagnostics.AddError("No identifier provided", "Either the CRN or the name of a datalake has to be given to be able to successfully collect data.")
		return
	}
	if utils.IsStringTypeHasValue(data.Crn) {
		datalake = data.Crn.ValueString()
	} else {
		datalake = data.Name.ValueString()
	}
	dl, err := describeDatalakeWithDiagnosticHandle(datalake, "", ctx, e.client, &resp.Diagnostics, &resp.State)
	if err != nil {
		return
	}

	fillDatalakePlatformValid(ctx, dl, resp, &data)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func fillDatalakePlatformValid(ctx context.Context, dl *models.DatalakeDetails, resp *datasource.ReadResponse, data *DatalakeConfigModel) {
	if len(dl.CloudPlatform) > 0 {
		switch dl.CloudPlatform {
		case "AWS":
			{
				mappedDatalake := &awsDatalakeResourceModel{}
				datalakeDetailsToAwsDatalakeResourceModel(ctx, dl, mappedDatalake, &utils.PollingOptions{}, &resp.Diagnostics)
				data.Aws = mappedDatalake
				break
			}
		case "AZURE":
			{
				mappedDatalake := &azureDatalakeResourceModel{}
				datalakeDetailsToAzureDatalakeResourceModel(ctx, dl, mappedDatalake, &utils.PollingOptions{}, &resp.Diagnostics)
				data.Azure = mappedDatalake
				break
			}
		case "GCP":
			{
				mappedDatalake := &gcpDatalakeResourceModel{}
				datalakeDetailsToGcpDatalakeResourceModel(ctx, dl, mappedDatalake, &utils.PollingOptions{}, &resp.Diagnostics)
				data.Gcp = mappedDatalake
				break
			}
		default:
			resp.Diagnostics.AddError("Unknown cloud platform", "Unknown cloud platform")
		}
	} else {
		resp.Diagnostics.AddError("Cloud platform not set", "Cloud platform not set")
	}
	fmt.Println("Datalake config data: ", data)
}
