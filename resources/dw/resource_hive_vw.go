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
	"github.com/cloudera/terraform-provider-cdp/utils/ptr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type hiveVwResource struct {
	client *cdp.Client
}

var (
	_ resource.Resource              = (*hiveVwResource)(nil)
	_ resource.ResourceWithConfigure = (*hiveVwResource)(nil)
)

func (r *hiveVwResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *hiveVwResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ml_workspace"
}

func (r *hiveVwResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = hiveVwSchema
}

func (r *hiveVwResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	//TODO implement me
	panic("implement me")
}

func (r *hiveVwResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	//TODO implement me
	panic("implement me")
}

func (r *hiveVwResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	//TODO implement me
	panic("implement me")
}

func (r *hiveVwResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data hiveVwResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	op := operations.NewDeleteVwParamsWithContext(ctx).
		WithInput(&models.DeleteVwRequest{
			ClusterID: ptr.To(data.ClusterID.ValueString()),
			VwID:      ptr.To(data.ID.ValueString()),
		})

	if _, err := r.client.Dw.Operations.DeleteVw(op); err != nil {
		resp.Diagnostics.AddError(
			"Error deleting Hive Virtual Warehouse",
			"Could not deleting Hive Virtual Warehouse, unexpected error: "+err.Error(),
		)
		return
	}
}
