// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datahub

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ datasource.DataSource = &datahubConfigDataSource{}
)

type datahubConfigDataSource struct {
	client *cdp.Client
}

func NewDatahubConfigDataSource() datasource.DataSource {
	return &datahubConfigDataSource{}
}

func (e *datahubConfigDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datahubConfigSchema
}

func (e *datahubConfigDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_datahub_config"
}

func (e *datahubConfigDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	e.client = utils.GetCdpClientForDataSource(req, resp)
}

func (e *datahubConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DatahubConfigModel
	config := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(config...)

	var datahub string
	if utils.IsStringTypeHasValue(data.Crn) && utils.IsStringTypeHasValue(data.Name) {
		resp.Diagnostics.AddError("Too many identifiers provided.", "Only the datahub's name or its CRN should be given but not both.")
		return
	}
	if !utils.IsStringTypeHasValue(data.Crn) && !utils.IsStringTypeHasValue(data.Name) {
		resp.Diagnostics.AddError("No identifier provided", "Either the CRN or the name of a datahub has to be given to be able to successfully collect data.")
		return
	}
	if utils.IsStringTypeHasValue(data.Crn) {
		datahub = data.Crn.ValueString()
	} else {
		datahub = data.Name.ValueString()
	}
	dh, err := describeDatahubWithDiagnosticHandle(datahub, "", ctx, e.client, &resp.Diagnostics, &resp.State)
	if err != nil {
		return
	}

	fillDatahubPlatformValid(ctx, dh, resp, &data)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func fillDatahubPlatformValid(ctx context.Context, dh *models.Cluster, resp *datasource.ReadResponse, data *DatahubConfigModel) {
	if len(strings.TrimSpace(dh.CloudPlatform)) > 0 {
		switch dh.CloudPlatform {
		case "AWS":
			{
				mappedDh := &awsDatahubResourceModel{}
				toAwsDatahubResource(ctx, dh, mappedDh, &utils.PollingOptions{})
				data.Aws = mappedDh
				break
			}
		case "AZURE":
			{
				mappedDh := &azureDatahubResourceModel{}
				toAzureDatahubResource(ctx, dh, mappedDh, &utils.PollingOptions{})
				data.Azure = mappedDh
				break
			}
		case "GCP":
			{
				mappedDh := &gcpDatahubResourceModel{}
				toGcpDatahubResource(ctx, dh, mappedDh, &utils.PollingOptions{})
				data.Gcp = mappedDh
				break
			}
		default:
			resp.Diagnostics.AddError("Unknown cloud platform", "Unknown cloud platform")
		}
	} else {
		resp.Diagnostics.AddError("Cloud platform not set", "Cloud platform not set")
	}
	tflog.Debug(ctx, fmt.Sprintf("Datahub config data: %+v", data))
}
