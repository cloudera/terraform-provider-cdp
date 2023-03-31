package environments

import (
	"context"
	"fmt"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSourceWithConfigure = &AWSCredentialPrerequisitesDataSource{}
)

func NewAWSCredentialPrerequisitesDataSource() datasource.DataSource {
	return &AWSCredentialPrerequisitesDataSource{}
}

type AWSCredentialPrerequisitesDataSource struct {
	client *cdp.Client
}

// AWSCredentialPrerequisitesDataSourceModel maps the data source schema data.
type AWSCredentialPrerequisitesDataSourceModel struct {
	AccountID  types.String `tfsdk:"account_id"`
	ExternalID types.String `tfsdk:"external_id"`
}

// Configure adds the provider configured client to the data source.
func (d *AWSCredentialPrerequisitesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cdp.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *cdp.Client, got: %T. Please report this issue to Cloudera.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *AWSCredentialPrerequisitesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_aws_credential_prerequisites"
}

func (d *AWSCredentialPrerequisitesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				MarkdownDescription: "AWS Account Id",
				Computed:            true,
			},
			"external_id": schema.StringAttribute{
				MarkdownDescription: "External Id used for the cross account role",
				Computed:            true,
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *AWSCredentialPrerequisitesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AWSCredentialPrerequisitesDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Reading GetCredentialPrerequisites")

	client := d.client.Environments
	cloudPlatform := "AWS"

	params := operations.NewGetCredentialPrerequisitesParams()
	params.WithInput(&environmentsmodels.GetCredentialPrerequisitesRequest{CloudPlatform: &cloudPlatform})

	response, err := client.Operations.GetCredentialPrerequisites(params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read cdp_environments_aws_credential_prerequisites, got error: %s", err))
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

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
