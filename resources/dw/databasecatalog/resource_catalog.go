// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package databasecatalog

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

type dwDatabaseCatalogResource struct {
	client *cdp.Client
}

var (
	_ resource.Resource              = (*dwDatabaseCatalogResource)(nil)
	_ resource.ResourceWithConfigure = (*dwDatabaseCatalogResource)(nil)
)

func NewDwDatabaseCatalogResource() resource.Resource {
	return &dwDatabaseCatalogResource{}
}

func (r *dwDatabaseCatalogResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *dwDatabaseCatalogResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dw_database_catalog"
}

func (r *dwDatabaseCatalogResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = dwDefaultDatabaseCatalogSchema
}

func (r *dwDatabaseCatalogResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan resourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	clusterID := plan.ClusterID.ValueStringPointer()

	if opts := plan.PollingOptions; !(opts != nil && opts.Async.ValueBool()) {
		callFailedCount := 0
		stateConf := &retry.StateChangeConf{
			Pending:      []string{"Accepted", "Creating", "Created", "Loading", "Starting"},
			Target:       []string{"Running"},
			Delay:        30 * time.Second,
			Timeout:      plan.getPollingTimeout(),
			PollInterval: 30 * time.Second,
			Refresh:      r.stateRefresh(ctx, clusterID, &callFailedCount, plan.getCallFailureThreshold()),
		}
		if _, err := stateConf.WaitForStateContext(ctx); err != nil {
			resp.Diagnostics.AddError(
				"Error waiting for Data Warehouse database catalog",
				fmt.Sprintf("Could not create database catalog, unexpected error: %v", err),
			)
			return
		}
	}

	catalog, err := r.getDatabaseCatalog(ctx, clusterID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error finding Data Warehouse database catalog", fmt.Sprintf("unexpected error: %v", err),
		)
		return
	}
	plan.ID = types.StringValue(catalog.ID)
	plan.Name = types.StringValue(catalog.Name)
	plan.Status = types.StringValue(catalog.Status)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *dwDatabaseCatalogResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Warn(ctx, "Read operation is not implemented yet.")
}

func (r *dwDatabaseCatalogResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Warn(ctx, "Update operation is not implemented yet.")
}

func (r *dwDatabaseCatalogResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Catalog is deleted as part of environment deletion, no need to delete it explicitly
	state.Status = types.StringValue("Deleted")
	resp.State.Set(ctx, state)
}

func (r *dwDatabaseCatalogResource) stateRefresh(ctx context.Context, clusterID *string, callFailedCount *int, callFailureThreshold int) func() (any, string, error) {
	return func() (any, string, error) {
		tflog.Debug(ctx, "About to get DBCs")
		catalogParams := operations.NewListDbcsParamsWithContext(ctx).WithInput(&models.ListDbcsRequest{ClusterID: clusterID})
		// List existing catalogs
		response, err := r.client.Dw.Operations.ListDbcs(catalogParams)
		if err != nil {
			tflog.Error(ctx,
				fmt.Sprintf("could not list database catalogs, unexpected error: %s", err.Error()),
			)
			return nil, "", err
		}
		resp := response.GetPayload()
		if len(resp.Dbcs) == 0 {
			*callFailedCount++
			if *callFailedCount <= callFailureThreshold {
				tflog.Warn(ctx, fmt.Sprintf("could not find Data Warehouse database catalog "+
					"due to [%s] but threshold limit is not reached yet (%d out of %d).", err.Error(), callFailedCount, callFailureThreshold))
				return nil, "", nil
			}
			tflog.Error(ctx, fmt.Sprintf("error describing Data Warehouse database catalog due to [%s] "+
				"failure threshold limit exceeded.", err.Error()))
			return nil, "", err
		}
		if len(resp.Dbcs) > 1 {
			err = fmt.Errorf("found more than one Data Warehouse database catalog for cluster %s", *clusterID)
			tflog.Error(ctx, fmt.Sprintf("error describing Data Warehouse database catalog due to [%s] ", err.Error()))
			return nil, "", err
		}
		*callFailedCount = 0

		tflog.Debug(ctx, fmt.Sprintf("Found database catalog %s with status %s", resp.Dbcs[0].ID, resp.Dbcs[0].Status))
		return resp.Dbcs[0], resp.Dbcs[0].Status, nil
	}
}

func (r *dwDatabaseCatalogResource) getDatabaseCatalog(ctx context.Context, clusterID *string) (*models.DbcSummary, error) {
	catalogParams := operations.NewListDbcsParamsWithContext(ctx).WithInput(&models.ListDbcsRequest{ClusterID: clusterID})
	// List existing catalogs
	response, err := r.client.Dw.Operations.ListDbcs(catalogParams)
	if err != nil {
		err = fmt.Errorf("could not list database catalogs, unexpected error: %s", err.Error())
		return nil, err
	}
	resp := response.GetPayload()
	if len(resp.Dbcs) != 1 {
		err = fmt.Errorf("exactly one Data Warehouse database catalog should be deployed for cluster %s", *clusterID)
		return nil, err
	}
	return resp.Dbcs[0], nil
}
