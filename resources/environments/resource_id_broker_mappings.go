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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource = &idBrokerMappingsResource{}
)

type idBrokerMappingsResource struct {
	client *cdp.Client
}

func NewIDBrokerMappingsResource() resource.Resource {
	return &idBrokerMappingsResource{}
}

func (r *idBrokerMappingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_id_broker_mappings"
}

type idBrokerMappingsResourceModel struct {
	ID types.String `tfsdk:"id"`

	BaselineRole types.String `tfsdk:"baseline_role"`

	DataAccessRole types.String `tfsdk:"data_access_role"`

	EnvironmentName types.String `tfsdk:"environment_name"`

	EnvironmentCrn types.String `tfsdk:"environment_crn"`

	Mappings []*idBrokerMapping `tfsdk:"mappings"`

	RangerAuditRole types.String `tfsdk:"ranger_audit_role"`

	RangerCloudAccessAuthorizerRole types.String `tfsdk:"ranger_cloud_access_authorizer_role"`

	SetEmptyMappings types.Bool `tfsdk:"set_empty_mappings"`

	MappingsVersion types.Int64 `tfsdk:"mappings_version"`
}

type idBrokerMapping struct {
	AccessorCrn types.String `tfsdk:"accessor_crn"`

	Role types.String `tfsdk:"role"`
}

func toSetIDBrokerMappingsRequest(model *idBrokerMappingsResourceModel) *environmentsmodels.SetIDBrokerMappingsRequest {
	resp := &environmentsmodels.SetIDBrokerMappingsRequest{}
	resp.BaselineRole = model.BaselineRole.ValueString()
	resp.DataAccessRole = model.DataAccessRole.ValueStringPointer()
	resp.EnvironmentName = model.EnvironmentName.ValueStringPointer()
	resp.RangerAuditRole = model.RangerAuditRole.ValueString()
	resp.RangerCloudAccessAuthorizerRole = model.RangerCloudAccessAuthorizerRole.ValueString()
	resp.SetEmptyMappings = model.SetEmptyMappings.ValueBoolPointer()
	resp.Mappings = make([]*environmentsmodels.IDBrokerMappingRequest, len(model.Mappings))
	for i, v := range model.Mappings {
		resp.Mappings[i] = &environmentsmodels.IDBrokerMappingRequest{
			AccessorCrn: v.AccessorCrn.ValueStringPointer(),
			Role:        v.Role.ValueStringPointer(),
		}
	}
	return resp
}

func (r *idBrokerMappingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "To enable your CDP user to utilize the central authentication features CDP provides and to exchange credentials for AWS or Azure access tokens, you have to map this CDP user to the correct IAM role or Azure Managed Service Identity (MSI).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"baseline_role": schema.StringAttribute{
				Optional: true,
			},
			"data_access_role": schema.StringAttribute{
				Optional: true,
			},
			"environment_name": schema.StringAttribute{
				Required: true,
			},
			"environment_crn": schema.StringAttribute{
				Computed: true,
			},
			"mappings": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"accessor_crn": schema.StringAttribute{
							Required: true,
						},
						"role": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
			"ranger_audit_role": schema.StringAttribute{
				Required: true,
			},
			"ranger_cloud_access_authorizer_role": schema.StringAttribute{
				Optional: true,
			},
			"set_empty_mappings": schema.BoolAttribute{
				Optional: true,
			},
			"mappings_version": schema.Int64Attribute{
				Computed: true,
			},
		},
	}
}

func (r *idBrokerMappingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *idBrokerMappingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state idBrokerMappingsResourceModel
	diags := req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Got Error while trying to set plan")
		return
	}

	client := r.client.Environments

	descParams := operations.NewDescribeEnvironmentParams()
	descParams.WithInput(&environmentsmodels.DescribeEnvironmentRequest{
		EnvironmentName: state.EnvironmentName.ValueStringPointer(),
	})
	descEnvResp, err := r.client.Environments.Operations.DescribeEnvironment(descParams)
	if err != nil {
		if strings.Contains(err.Error(), "Code:NOT_FOUND") {
			resp.Diagnostics.AddWarning("Environment resource not found on provider", "Environment not found.")
			tflog.Warn(ctx, "Environment not found")
			return
		}
	}

	params := operations.NewSetIDBrokerMappingsParams()
	params.WithInput(toSetIDBrokerMappingsRequest(&state))
	responseOk, err := client.Operations.SetIDBrokerMappings(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error setting ID Broker mappings",
			"Got the following error setting ID Broker mappings: "+err.Error(),
		)
		return
	}

	idBrokerResp := responseOk.Payload
	state.ID = types.StringPointerValue(descEnvResp.Payload.Environment.Crn)
	state.EnvironmentCrn = types.StringPointerValue(descEnvResp.Payload.Environment.Crn)
	state.EnvironmentName = types.StringPointerValue(descEnvResp.Payload.Environment.EnvironmentName)
	state.MappingsVersion = types.Int64PointerValue(idBrokerResp.MappingsVersion)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *idBrokerMappingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state idBrokerMappingsResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Environments

	params := operations.NewGetIDBrokerMappingsParams()
	params.WithInput(&environmentsmodels.GetIDBrokerMappingsRequest{
		EnvironmentName: state.EnvironmentName.ValueStringPointer(),
	})
	responseOk, err := client.Operations.GetIDBrokerMappings(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting ID Broker mappings",
			"Got the following error getting ID Broker mappings: "+err.Error(),
		)
		return
	}

	idBrokerResp := responseOk.Payload
	toIdBrokerMappingsResourceModel(idBrokerResp, &state)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func toIdBrokerMappingsResourceModel(mapping *environmentsmodels.GetIDBrokerMappingsResponse, out *idBrokerMappingsResourceModel) {
	out.BaselineRole = types.StringValue(mapping.BaselineRole)
	out.DataAccessRole = types.StringPointerValue(mapping.DataAccessRole)
	out.MappingsVersion = types.Int64PointerValue(mapping.MappingsVersion)
	out.RangerAuditRole = types.StringPointerValue(mapping.RangerAuditRole)
	out.RangerCloudAccessAuthorizerRole = types.StringValue(mapping.RangerCloudAccessAuthorizerRole)
	out.Mappings = make([]*idBrokerMapping, len(mapping.Mappings))
	for i, v := range mapping.Mappings {
		out.Mappings[i] = &idBrokerMapping{
			AccessorCrn: types.StringPointerValue(v.AccessorCrn),
			Role:        types.StringPointerValue(v.Role),
		}
	}
}

func (r *idBrokerMappingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state idBrokerMappingsResourceModel
	diags := req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Got Error while trying to set plan")
		return
	}

	client := r.client.Environments

	params := operations.NewSetIDBrokerMappingsParams()
	params.WithInput(toSetIDBrokerMappingsRequest(&state))
	responseOk, err := client.Operations.SetIDBrokerMappings(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error setting ID Broker mappings",
			"Got the following error setting ID Broker mappings: "+err.Error(),
		)
		return
	}

	idBrokerResp := responseOk.Payload
	state.ID = state.EnvironmentCrn
	state.MappingsVersion = types.Int64PointerValue(idBrokerResp.MappingsVersion)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *idBrokerMappingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state idBrokerMappingsResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Environments

	params := operations.NewDeleteIDBrokerMappingsParams()
	params.WithInput(&environmentsmodels.DeleteIDBrokerMappingsRequest{EnvironmentCrn: state.EnvironmentCrn.ValueStringPointer()})
	_, err := client.Operations.DeleteIDBrokerMappings(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting ID Broker mapping",
			"Could not delete ID Broker mapping unexpected error: "+err.Error(),
		)
		return
	}
}
