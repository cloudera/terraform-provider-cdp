// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package dw

import (
	"context"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type hiveResource struct {
	client *cdp.Client
}

var (
	_ resource.Resource              = (*hiveResource)(nil)
	_ resource.ResourceWithConfigure = (*hiveResource)(nil)
)

func NewHiveResource() resource.Resource {
	return &hiveResource{}
}

func (r *hiveResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *hiveResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vw_hive"
}

func (r *hiveResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = hiveSchema
}

func (r *hiveResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan hiveResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	vw := operations.NewCreateVwParamsWithContext(ctx).
		WithInput(&models.CreateVwRequest{
			Name:      plan.Name.ValueStringPointer(),
			ClusterID: plan.ClusterID.ValueStringPointer(),
			DbcID:     plan.DbCatalogID.ValueStringPointer(),
			VwType:    models.VwTypeHive.Pointer(),
		})

	// Create new virtual warehouse
	response, err := r.client.Dw.Operations.CreateVw(vw)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating hive virtual warehouse",
			"Could not create hive, unexpected error: "+err.Error(),
		)
		return
	}

	payload := response.GetPayload()
	desc := operations.NewDescribeVwParamsWithContext(ctx).
		WithInput(&models.DescribeVwRequest{VwID: &payload.VwID, ClusterID: plan.ClusterID.ValueStringPointer()})
	describe, err := r.client.Dw.Operations.DescribeVw(desc)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating hive virtual warehouse",
			"Could not describe hive, unexpected error: "+err.Error(),
		)
		return
	}

	hive := describe.GetPayload()

	// Map response body to schema and populate Computed attribute values
	plan.ID = types.StringValue(hive.Vw.ID)
	plan.DbCatalogID = types.StringValue(hive.Vw.DbcID)
	plan.Name = types.StringValue(hive.Vw.Name)
	// TODO why is this not accepted with error: An unexpected error was encountered trying to convert tftypes.Value into dw.hiveResourceModel. This is always an error in the provider. Please report the following to the provider developer:
	//plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *hiveResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	//TODO implement me
	tflog.Warn(ctx, "Read operation is not implemented yet.")
}

func (r *hiveResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	//TODO implement me
	tflog.Warn(ctx, "Update operation is not implemented yet.")
}

func (r *hiveResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state hiveResourceModel

	// Read Terraform prior state into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	op := operations.NewDeleteVwParamsWithContext(ctx).
		WithInput(&models.DeleteVwRequest{
			ClusterID: state.ClusterID.ValueStringPointer(),
			VwID:      state.ID.ValueStringPointer(),
		})

	if _, err := r.client.Dw.Operations.DeleteVw(op); err != nil {
		resp.Diagnostics.AddError(
			"Error deleting Hive Virtual Warehouse",
			"Could not delete Hive Virtual Warehouse, unexpected error: "+err.Error(),
		)
		return
	}
}
