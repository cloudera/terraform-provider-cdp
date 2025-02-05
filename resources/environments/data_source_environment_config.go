// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package environments

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ datasource.DataSource = &environmentConfigDataSource{}
)

type environmentConfigDataSource struct {
	client *cdp.Client
}

func NewEnvironmentConfigDataSource() datasource.DataSource {
	return &environmentConfigDataSource{}
}

func (e *environmentConfigDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_config"
}

func (e *environmentConfigDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = EnvironmentConfigSchema
}

func (e *environmentConfigDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	e.client = utils.GetCdpClientForDataSource(req, resp)
}

func (e *environmentConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data EnvironmentConfigModel
	config := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(config...)

	envName := data.Name.ValueString()
	env, err := describeEnvironmentWithDiagnosticHandle(envName, "", ctx, e.client, &resp.Diagnostics, &resp.State)
	if err != nil {
		return
	}

	fillEnvironmentPlatformValid(ctx, env, resp, &data)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func fillEnvironmentPlatformValid(ctx context.Context, env *models.Environment, resp *datasource.ReadResponse, data *EnvironmentConfigModel) {
	if env.CloudPlatform != nil {
		switch *env.CloudPlatform {
		case "AWS":
			{
				mappedEnv := &awsEnvironmentResourceModel{}
				toAwsEnvironmentResource(ctx, env, mappedEnv, &utils.PollingOptions{}, &resp.Diagnostics)
				data.Aws = mappedEnv
				break
			}
		case "AZURE":
			{
				mappedEnv := &azureEnvironmentResourceModel{}
				toAzureEnvironmentResource(ctx, env, mappedEnv, &utils.PollingOptions{}, &resp.Diagnostics)
				data.Azure = mappedEnv
				break
			}
		case "GCP":
			{
				mappedEnv := &gcpEnvironmentResourceModel{}
				toGcpEnvironmentResource(ctx, env, mappedEnv, &utils.PollingOptions{}, &resp.Diagnostics)
				data.Gcp = mappedEnv
				break
			}
		default:
			resp.Diagnostics.AddError("Unknown cloud platform", "Unknown cloud platform")
		}
	} else {
		resp.Diagnostics.AddError("Cloud platform not set", "Cloud platform not set")
	}
	fmt.Println("Environment config data: ", data)
}
