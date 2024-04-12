// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package opdb

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/client/operations"
	opdbmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ resource.Resource = &databaseResource{}
)

type databaseResource struct {
	client *cdp.Client
}

func NewDatabaseResource() resource.Resource {
	return &databaseResource{}
}

func (r *databaseResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_operational_database"
}

func (r *databaseResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *databaseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "OPDB cluster creation process requested.")
	var data databaseResourceModel
	tflog.Info(ctx, fmt.Sprintf("Creating OPDB with name: %s", data.DatabaseName.ValueString()))
	diags := req.Plan.Get(ctx, &data)
	tflog.Debug(ctx, fmt.Sprintf("OPDB resource model: %+v", data))
	resp.Diagnostics.Append(diags...)
	tflog.Debug(ctx, fmt.Sprintf("Diags: %+v", resp.Diagnostics))
	if resp.Diagnostics.HasError() {
		tflog.Warn(ctx, "OPDB resource model has error, stopping the creation process.")
		return
	}

	params := operations.NewCreateDatabaseParamsWithContext(ctx)
	params.WithInput(fromModelToDatabaseRequest(data, ctx))

	tflog.Info(ctx, fmt.Sprintf("Sending create request for OPDB with name: %s", data.DatabaseName.ValueString()))
	res, err := r.client.Opdb.Operations.CreateDatabase(params)

	tflog.Info(ctx, fmt.Sprintf("Create request for OPDB with name: %s has been sent with the result of: %+v", data.DatabaseName.ValueString(), res))
	if err != nil {
		utils.AddDatabaseDiagnosticsError(err, &resp.Diagnostics, "create OPDB")
		return
	}

	getCommonDatabaseDetails(&data, res.Payload.DatabaseDetails)

	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	status, err := waitForToBeAvailable(data.DatabaseName.ValueString(), data.Environment.ValueString(), r.client.Opdb, ctx, data.PollingOptions)
	tflog.Debug(ctx, fmt.Sprintf("Database polling finished, setting status from '%s' to '%s'", data.Status.ValueString(), status))
	data.Status = types.StringValue(status)
	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err != nil {
		tflog.Debug(ctx, fmt.Sprintf("Cluster creation has ended up in error: %s", err.Error()))
		utils.AddDatabaseDiagnosticsError(err, &resp.Diagnostics, "create Database")
		return
	}
}

func (r *databaseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state databaseResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewDescribeDatabaseParamsWithContext(ctx)
	params.WithInput(&opdbmodels.DescribeDatabaseRequest{
		DatabaseName:    state.DatabaseName.ValueStringPointer(),
		EnvironmentName: state.Environment.ValueStringPointer(),
	})

	result, err := r.client.Opdb.Operations.DescribeDatabase(params)
	if err != nil {
		if isNotFoundError(err) {
			resp.Diagnostics.AddWarning("Resource not found on provider", "COD not found, removing from state.")
			tflog.Warn(ctx, "COD not found, removing from state", map[string]interface{}{"id": state.DatabaseName.ValueString()})
			resp.State.RemoveResource(ctx)
			return
		}
		utils.AddDatabaseDiagnosticsError(err, &resp.Diagnostics, "read Database")
		return
	}

	databaseDetails := result.Payload.DatabaseDetails

	getCommonDatabaseDetails(&state, databaseDetails)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func getCommonDatabaseDetails(data *databaseResourceModel, databaseDetails *opdbmodels.DatabaseDetails) {
	data.DatabaseName = types.StringPointerValue(databaseDetails.DatabaseName)
	data.Crn = types.StringPointerValue(databaseDetails.Crn)
	data.Environment = types.StringValue(databaseDetails.EnvironmentName)
	data.Status = types.StringValue(string(databaseDetails.Status))

	data.ScaleType = types.StringValue(string(databaseDetails.ScaleType))
	data.StorageLocation = types.StringValue(databaseDetails.StorageLocation)

	data.NumEdgeNodes = types.Int64Value(int64(databaseDetails.DbEdgeNodeCount))

	if len(databaseDetails.StorageDetailsForWorkers) >= 1 {
		data.AttachedStorageForWorkers = createStorageDetailsForWorkers(databaseDetails.StorageDetailsForWorkers[0])
	}
}

func createStorageDetailsForWorkers(storageDetailsForWorker *opdbmodels.StorageDetailsForWorker) *AttachedStorageForWorkersStruct {
	return &AttachedStorageForWorkersStruct{
		VolumeCount: types.Int64Value(int64(storageDetailsForWorker.VolumeCount)),
		VolumeSize:  types.Int64Value(int64(storageDetailsForWorker.VolumeSize)),
		VolumeType:  types.StringValue(string(storageDetailsForWorker.VolumeType)),
	}
}

func (r *databaseResource) Update(ctx context.Context, _ resource.UpdateRequest, _ *resource.UpdateResponse) {
	tflog.Warn(ctx, "Update operation is not implemented yet.")
}

func (r *databaseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state databaseResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewDropDatabaseParamsWithContext(ctx).WithInput(&opdbmodels.DropDatabaseRequest{
		DatabaseName:    state.DatabaseName.ValueStringPointer(),
		EnvironmentName: state.Environment.ValueStringPointer(),
	})

	tflog.Debug(ctx, fmt.Sprintf("Sending drop database request: %s %s", *params.Input.DatabaseName, *params.Input.EnvironmentName))

	_, err := r.client.Opdb.Operations.DropDatabase(params)
	if err != nil {
		if !isNotFoundError(err) {
			utils.AddDatabaseDiagnosticsError(err, &resp.Diagnostics, "delete database")
		}
		return
	}

	err = waitForToBeDeleted(state.DatabaseName.ValueString(), state.Environment.ValueString(), r.client.Opdb, ctx, state.PollingOptions)
	if err != nil {
		utils.AddDatabaseDiagnosticsError(err, &resp.Diagnostics, "delete database")
		return
	}
}
