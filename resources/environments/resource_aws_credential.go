package environments

import (
	"context"
	"strings"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource = &awsCredentialResource{}
)

type awsCredentialResource struct {
	client *cdp.Client
}

type awsCredentialResourceModel struct {
	ID             types.String `tfsdk:"id"`
	CredentialName types.String `tfsdk:"credential_name"`
	RoleArn        types.String `tfsdk:"role_arn"`
	Crn            types.String `tfsdk:"crn"`
	Description    types.String `tfsdk:"description"`
}

func NewAwsCredentialResource() resource.Resource {
	return &awsCredentialResource{}
}

func (r *awsCredentialResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_aws_credential"
}

func (r *awsCredentialResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The AWS credential is used for authorization  to provision resources such as compute instances within your cloud provider account.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"credential_name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"role_arn": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"crn": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *awsCredentialResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *awsCredentialResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from data
	var data awsCredentialResourceModel
	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Environments

	params := operations.NewCreateAWSCredentialParams()
	params.WithInput(&environmentsmodels.CreateAWSCredentialRequest{
		CredentialName: data.CredentialName.ValueStringPointer(),
		Description:    data.Description.ValueString(),
		RoleArn:        data.RoleArn.ValueStringPointer(),
	})

	// TODO: find out how to do retries. There is an eventual consistency issue when the AWS cross account credential
	// is just created but is not "synced up" in AWS. We should retry for a short time if it is the case.
	responseOk, err := client.Operations.CreateAWSCredential(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating AWS Credentials",
			"Got error while creating AWS Credentials: "+err.Error(),
		)
		return
	}

	data.Crn = types.StringPointerValue(responseOk.Payload.Credential.Crn)
	data.ID = data.CredentialName

	// Save data into Terraform state
	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *awsCredentialResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state awsCredentialResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed value from CDP
	credentialName := state.ID.ValueString()
	params := operations.NewListCredentialsParams()
	params.WithInput(&environmentsmodels.ListCredentialsRequest{CredentialName: credentialName})
	listCredentialsResp, err := r.client.Environments.Operations.ListCredentials(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading AWS Credentials",
			"Could not read AWS Credentials: "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	// Overwrite items with refreshed state
	credentials := listCredentialsResp.GetPayload().Credentials
	if len(credentials) == 0 || *credentials[0].CredentialName != credentialName {
		resp.State.RemoveResource(ctx) // deleted
		return
	}
	c := credentials[0]

	state.ID = types.StringPointerValue(c.CredentialName)
	state.CredentialName = types.StringPointerValue(c.CredentialName)
	state.Crn = types.StringPointerValue(c.Crn)
	state.Description = types.StringValue(c.Description)
	state.RoleArn = types.StringValue(c.AwsCredentialProperties.RoleArn)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *awsCredentialResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *awsCredentialResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state awsCredentialResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	credentialName := state.ID.ValueString()
	params := operations.NewDeleteCredentialParams()
	params.WithInput(&environmentsmodels.DeleteCredentialRequest{CredentialName: &credentialName})
	_, err := r.client.Environments.Operations.DeleteCredential(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting AWS Credential",
			"Could not delete AWS Credential, unexpected error: "+err.Error(),
		)
		return
	}
}

func isNotAuthorizedError(err error) bool {
	if d, ok := err.(*operations.CreateAWSEnvironmentDefault); ok && d.GetPayload() != nil {
		return d.GetPayload().Code == "INVALID_ARGUMENT" &&
			strings.Contains(d.GetPayload().Message, "You are not authorized")
	}
	return false
}
