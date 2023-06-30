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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_             resource.Resource = &idBrokerMappingsResource{}
	emptyMappings                   = true
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

	DataAccessRole types.String `tfsdk:"data_access_role"`

	EnvironmentName types.String `tfsdk:"environment_name"`

	EnvironmentCrn types.String `tfsdk:"environment_crn"`

	Mappings types.Set `tfsdk:"mappings"`

	RangerAuditRole types.String `tfsdk:"ranger_audit_role"`

	RangerCloudAccessAuthorizerRole types.String `tfsdk:"ranger_cloud_access_authorizer_role"`

	SetEmptyMappings types.Bool `tfsdk:"set_empty_mappings"`

	MappingsVersion types.Int64 `tfsdk:"mappings_version"`
}

type idBrokerMapping struct {
	AccessorCrn types.String `tfsdk:"accessor_crn"`

	Role types.String `tfsdk:"role"`
}

func toSetIDBrokerMappingsRequest(ctx context.Context, model *idBrokerMappingsResourceModel) *environmentsmodels.SetIDBrokerMappingsRequest {
	resp := &environmentsmodels.SetIDBrokerMappingsRequest{}
	resp.DataAccessRole = model.DataAccessRole.ValueStringPointer()
	resp.EnvironmentName = model.EnvironmentName.ValueStringPointer()
	resp.RangerAuditRole = model.RangerAuditRole.ValueString()
	resp.RangerCloudAccessAuthorizerRole = model.RangerCloudAccessAuthorizerRole.ValueString()
	resp.SetEmptyMappings = model.SetEmptyMappings.ValueBoolPointer()
	resp.Mappings = make([]*environmentsmodels.IDBrokerMappingRequest, len(model.Mappings.Elements()))
	model.Mappings.ElementsAs(ctx, resp.Mappings, true)
	return resp
}

func (r *idBrokerMappingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "To enable your CDP user to utilize the central authentication features CDP provides and to exchange credentials for AWS or Azure access tokens, you have to map this CDP user to the correct IAM role or Azure Managed Service Identity (MSI).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"data_access_role": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"environment_name": schema.StringAttribute{
				Required: true,
			},
			"environment_crn": schema.StringAttribute{
				Required: true,
			},
			"mappings": schema.SetNestedAttribute{
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
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"set_empty_mappings": schema.BoolAttribute{
				Optional: true,
			},
			"mappings_version": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
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

	params := operations.NewSetIDBrokerMappingsParamsWithContext(ctx)
	params.WithInput(toSetIDBrokerMappingsRequest(ctx, &state))
	responseOk, err := client.Operations.SetIDBrokerMappings(params)
	if err != nil {
		if isSetIDBEnvNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Error applying ID Broker mappings",
				"Environment not found: "+state.EnvironmentCrn.ValueString(),
			)
			return
		}
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "creating ID Broker mapping")
		return
	}

	idBrokerResp := responseOk.Payload
	state.ID = state.EnvironmentCrn
	state.RangerCloudAccessAuthorizerRole = types.StringValue(idBrokerResp.RangerCloudAccessAuthorizerRole)
	state.MappingsVersion = types.Int64PointerValue(idBrokerResp.MappingsVersion)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func isSetIDBEnvNotFoundError(err error) bool {
	if envErr, ok := err.(*operations.SetIDBrokerMappingsDefault); ok {
		if cdp.IsEnvironmentsError(envErr.GetPayload(), "NOT_FOUND", "") {
			return true
		}
	}
	return false
}

func (r *idBrokerMappingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state idBrokerMappingsResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Environments
	params := operations.NewGetIDBrokerMappingsParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.GetIDBrokerMappingsRequest{
		EnvironmentName: state.EnvironmentName.ValueStringPointer(),
	})
	responseOk, err := client.Operations.GetIDBrokerMappings(params)
	if err != nil {
		if envErr, ok := err.(*operations.GetIDBrokerMappingsDefault); ok {
			if cdp.IsEnvironmentsError(envErr.GetPayload(), "NOT_FOUND", "") {
				resp.Diagnostics.AddWarning("Resource not found on provider", "Environment not found, removing from state.")
				tflog.Warn(ctx, "Environment not found, removing from state", map[string]interface{}{
					"id": state.ID.ValueString(),
				})
				resp.State.RemoveResource(ctx)
				return
			}
		}
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "reading ID Broker mapping")
		return
	}

	idBrokerResp := responseOk.Payload
	toIdBrokerMappingsResourceModel(ctx, idBrokerResp, &state)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func toIdBrokerMappingsResourceModel(ctx context.Context, mapping *environmentsmodels.GetIDBrokerMappingsResponse, out *idBrokerMappingsResourceModel) {
	out.DataAccessRole = types.StringPointerValue(mapping.DataAccessRole)
	out.MappingsVersion = types.Int64PointerValue(mapping.MappingsVersion)
	out.RangerAuditRole = types.StringPointerValue(mapping.RangerAuditRole)
	out.RangerCloudAccessAuthorizerRole = types.StringValue(mapping.RangerCloudAccessAuthorizerRole)
	if len(mapping.Mappings) == 0 {
		out.Mappings = types.SetNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"accessor_crn": types.StringType,
				"role":         types.StringType,
			},
		})
	} else {
		mappings := make([]*idBrokerMapping, len(mapping.Mappings))
		for i, v := range mapping.Mappings {
			mappings[i] = &idBrokerMapping{
				AccessorCrn: types.StringPointerValue(v.AccessorCrn),
				Role:        types.StringPointerValue(v.Role),
			}
		}
		out.Mappings, _ = types.SetValueFrom(ctx, types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"accessor_crn": types.StringType,
				"role":         types.StringType,
			},
		}, mappings)
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

	params := operations.NewSetIDBrokerMappingsParamsWithContext(ctx)
	params.WithInput(toSetIDBrokerMappingsRequest(ctx, &state))
	responseOk, err := client.Operations.SetIDBrokerMappings(params)
	if err != nil {
		if isSetIDBEnvNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Error applying ID Broker mappings",
				"Environment not found: "+state.EnvironmentCrn.ValueString(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Error setting ID Broker mappings",
			"Got the following error setting ID Broker mappings: "+err.Error(),
		)
		return
	}

	idBrokerResp := responseOk.Payload
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

	params := operations.NewSetIDBrokerMappingsParamsWithContext(ctx)
	input := &environmentsmodels.SetIDBrokerMappingsRequest{}
	input.SetEmptyMappings = &emptyMappings
	params.WithInput(input)
	_, err := client.Operations.SetIDBrokerMappings(params)
	if err != nil {
		if isSetIDBEnvNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Error deleting ID Broker mappings",
				"Environment not found: "+state.EnvironmentCrn.ValueString(),
			)
			return
		}
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "deleting ID Broker mapping")
		return
	}
}
