// Copyright 2023 Cloudera. All Rights Reserved.
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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSourceWithConfigure = &awsCredentialPrerequisitesDataSource{}
)

func NewAWSCredentialPrerequisitesDataSource() datasource.DataSource {
	return &awsCredentialPrerequisitesDataSource{}
}

type awsCredentialPrerequisitesDataSource struct {
	client *cdp.Client
}

// awsCredentialPrerequisitesDataSourceModel maps the data source schema data.
type awsCredentialPrerequisitesDataSourceModel struct {
	ID         types.String `tfsdk:"id"`
	AccountID  types.String `tfsdk:"account_id"`
	ExternalID types.String `tfsdk:"external_id"`
}

// Configure adds the provider configured client to the data source.
func (d *awsCredentialPrerequisitesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = utils.GetCdpClientForDataSource(req, resp)
}

func (d *awsCredentialPrerequisitesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_aws_credential_prerequisites"
}

func (d *awsCredentialPrerequisitesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This data source is used to get information required to set up a delegated access role in AWS that can be used to create a CDP credential.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				MarkdownDescription: "The AWS account ID of the identity used by CDP when assuming a delegated access role associated with a CDP credential.",
				Computed:            true,
			},
			"external_id": schema.StringAttribute{
				MarkdownDescription: "The external ID that will be used when assuming a delegated access role associated with a CDP credential.",
				Computed:            true,
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *awsCredentialPrerequisitesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data awsCredentialPrerequisitesDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Reading GetCredentialPrerequisites")

	client := d.client.Environments
	cloudPlatform := "AWS"

	params := operations.NewGetCredentialPrerequisitesParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.GetCredentialPrerequisitesRequest{CloudPlatform: &cloudPlatform})

	response, err := client.Operations.GetCredentialPrerequisites(params)
	if err != nil {
		msg := err.Error()
		if d, ok := err.(*operations.GetCredentialPrerequisitesDefault); ok && d.GetPayload() != nil {
			msg = d.GetPayload().Message
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read cdp_environments_aws_credential_prerequisites, got error: %s", msg))
		return
	}
	prerequisites := response.GetPayload()
	if prerequisites == nil || prerequisites.Aws == nil {
		resp.State.RemoveResource(ctx) // deleted
		return
	}

	ctx = tflog.SetField(ctx, "Client info:", client)
	tflog.Info(ctx, "Read GetCredentialPrerequisites")

	data.AccountID = types.StringValue(prerequisites.AccountID)
	data.ExternalID = types.StringValue(*prerequisites.Aws.ExternalID)
	data.ID = types.StringValue(prerequisites.AccountID + ":" + *prerequisites.Aws.ExternalID)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
