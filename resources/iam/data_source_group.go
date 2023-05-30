package iam

import (
	"context"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSourceWithConfigure = &groupDataSource{}
)

func NewGroupDataSource() datasource.DataSource {
	return &groupDataSource{}
}

type groupDataSource struct {
	client *cdp.Client
}

// Configure adds the provider configured client to the data source.
func (d *groupDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = utils.GetCdpClientForDataSource(req, resp)
}

func (d *groupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_group"
}

func (d *groupDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "A group is a named collection of users and machine users. Roles and resource roles can be assigned to a group impacting all members of the group.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"group_name": schema.StringAttribute{
				MarkdownDescription: "The name of the group. This name must be unique. There are certain restrictions on the group name. Refer to the How To > User Management section in the Management Console documentation for the details.",
				Required:            true,
			},
			"sync_membership_on_user_login": schema.BoolAttribute{
				MarkdownDescription: "Whether group membership is synced when a user logs in. The default is to sync group membership.",
				Computed:            true,
			},
			"crn": schema.StringAttribute{
				MarkdownDescription: "The CRN of the resource.",
				Computed:            true,
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *groupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// Read user supplied values
	var data groupModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	sharedGroupRead(ctx, d.client.Iam, &data, &resp.State, &resp.Diagnostics)
}
