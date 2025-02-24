// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package impala

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

type impalaResource struct {
	client *cdp.Client
}

var (
	_ resource.Resource              = (*impalaResource)(nil)
	_ resource.ResourceWithConfigure = (*impalaResource)(nil)
)

func NewImpalaResource() resource.Resource {
	return &impalaResource{}
}

func (r *impalaResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *impalaResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dw_vw_impala"
}

func (r *impalaResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = impalaSchema
}

func (r *impalaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan resourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create the VW Request using a helper
	vwhCreateRequest := r.createVwRequestFromPlan(&plan)
	/*
		tflog.Debug(ctx, fmt.Sprintf("CreateVw request: %+v", vwhCreateRequest))
		tflog.Debug(ctx, fmt.Sprintf("AutoscalingOptions: autoSuspendTimeoutSeconds: %d", vwhCreateRequest.Autoscaling.AutoSuspendTimeoutSeconds))
		tflog.Debug(ctx, fmt.Sprintf("AutoscalingOptions: disableAutoSuspend: %t", vwhCreateRequest.Autoscaling.DisableAutoSuspend))
		tflog.Debug(ctx, fmt.Sprintf("AutoscalingOptions: enableUnifiedAnalytics: %t", vwhCreateRequest.Autoscaling.EnableUnifiedAnalytics))
		tflog.Debug(ctx, fmt.Sprintf("AutoscalingOptions: hiveDesiredFreeCapacity: %d", vwhCreateRequest.Autoscaling.HiveDesiredFreeCapacity))
		tflog.Debug(ctx, fmt.Sprintf("AutoscalingOptions: hiveScaleWaitTimeSeconds: %d", vwhCreateRequest.Autoscaling.HiveScaleWaitTimeSeconds))
		tflog.Debug(ctx, fmt.Sprintf("AutoscalingOptions: impalaEnableCatalogHighAvailability: %t", vwhCreateRequest.Autoscaling.ImpalaEnableCatalogHighAvailability))
		tflog.Debug(ctx, fmt.Sprintf("AutoscalingOptions: impalaEnableShutdownOfCoordinator: %t", vwhCreateRequest.Autoscaling.ImpalaEnableShutdownOfCoordinator))
		tflog.Debug(ctx, fmt.Sprintf("AutoscalingOptions: impalaExecutorGroupSets: %+v", vwhCreateRequest.Autoscaling.ImpalaExecutorGroupSets))
		tflog.Debug(ctx, fmt.Sprintf("AutoscalingOptions: impalaHighAvailabilityMode: %s", vwhCreateRequest.Autoscaling.ImpalaHighAvailabilityMode))
		tflog.Debug(ctx, fmt.Sprintf("AutoscalingOptions: impalaNumOfActiveCoordinators: %d", vwhCreateRequest.Autoscaling.ImpalaNumOfActiveCoordinators))
		tflog.Debug(ctx, fmt.Sprintf("AutoscalingOptions: impalaScaleDownDelaySeconds: %d", vwhCreateRequest.Autoscaling.ImpalaScaleDownDelaySeconds))
		tflog.Debug(ctx, fmt.Sprintf("AutoscalingOptions: impalaScaleUpDelaySeconds: %d", vwhCreateRequest.Autoscaling.ImpalaScaleUpDelaySeconds))
		tflog.Debug(ctx, fmt.Sprintf("AutoscalingOptions: impalaShutdownOfCoordinatorDelaySeconds: %d", vwhCreateRequest.Autoscaling.ImpalaShutdownOfCoordinatorDelaySeconds))
		//tflog.Debug(ctx, fmt.Sprintf("AutoscalingOptions: maxClusters: %d", *vwhCreateRequest.Autoscaling.MaxClusters))
		//tflog.Debug(ctx, fmt.Sprintf("AutoscalingOptions: minClusters: %d", *vwhCreateRequest.Autoscaling.MinClusters))
	*/

	// Make API request to create VW
	response, err := r.client.Dw.Operations.CreateVw(
		operations.NewCreateVwParamsWithContext(ctx).WithInput(vwhCreateRequest),
	)
	if err != nil || response.GetPayload() == nil {
		resp.Diagnostics.AddError(
			"Error creating Impala virtual warehouse",
			fmt.Sprintf("Could not create Impala, unexpected error: %v", err),
		)
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("CreateVw response: %+v", response.GetPayload()))

	// Wait for the VW to reach Running state
	if err := r.waitForVwRunning(ctx, &plan, &response.GetPayload().VwID); err != nil {
		resp.Diagnostics.AddError(
			"Error waiting for Data Warehouse Impala virtual warehouse",
			fmt.Sprintf("Could not create Impala, unexpected error: %v", err),
		)
		return
	}

	// Fetch and map the response to plan using a helper
	if err := r.populatePlanFromDescribe(ctx, &plan, &response.GetPayload().VwID); err != nil {
		resp.Diagnostics.AddError(
			"Error creating Impala virtual warehouse",
			fmt.Sprintf("Could not describe Impala, unexpected error: %v", err),
		)
		return
	}

	// Save the updated plan into state
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *impalaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Warn(ctx, "Read operation is not implemented yet.")
}

func (r *impalaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Warn(ctx, "Update operation is not implemented yet.")
}

func (r *impalaResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	clusterID := state.ClusterID.ValueStringPointer()
	vwID := state.ID.ValueStringPointer()

	if clusterID == nil || vwID == nil {
		resp.Diagnostics.AddError(
			"Invalid State",
			"ClusterID or Virtual Warehouse ID is missing.",
		)
		return
	}

	err := r.deleteVirtualWarehouse(ctx, clusterID, vwID)
	if err != nil {
		if strings.Contains(err.Error(), "Virtual Warehouse not found") {
			return
		}
		resp.Diagnostics.AddError(
			"Error deleting Impala Virtual Warehouse",
			"Could not delete Impala Virtual Warehouse, unexpected error: "+err.Error(),
		)
		return
	}

	// Handle polling for not async call
	if opts := state.PollingOptions; opts == nil || !opts.Async.ValueBool() {
		err = r.pollForDeletion(ctx, state, clusterID, vwID)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error waiting for Impala Virtual Warehouse deletion",
				"Could not delete Impala Virtual Warehouse, unexpected error: "+err.Error(),
			)
		}
	}
}

func (r *impalaResource) stateRefresh(ctx context.Context, clusterID *string, vwID *string, callFailedCount *int, callFailureThreshold int) func() (any, string, error) {
	return func() (any, string, error) {
		tflog.Debug(ctx, "About to describe Impala")
		params := operations.NewDescribeVwParamsWithContext(ctx).
			WithInput(&models.DescribeVwRequest{ClusterID: clusterID, VwID: vwID})
		resp, err := r.client.Dw.Operations.DescribeVw(params)
		if err != nil {
			if strings.Contains(err.Error(), "Virtual Warehouse not found") {
				return &models.DescribeVwResponse{}, "Deleted", nil
			}
			*callFailedCount++
			if *callFailedCount <= callFailureThreshold {
				tflog.Warn(ctx, fmt.Sprintf("could not describe Data Warehouse Impala Virtual Warehouse "+
					"due to [%s] but threshold limit is not reached yet (%d out of %d).", err.Error(), callFailedCount, callFailureThreshold))
				return nil, "", nil
			}
			tflog.Error(ctx, fmt.Sprintf("error describing Data Warehouse Impala Virtual Warehouse due to [%s] "+
				"failure threshold limit exceeded.", err.Error()))
			return nil, "", err
		}
		*callFailedCount = 0
		vw := resp.GetPayload()
		tflog.Debug(ctx, fmt.Sprintf("Described Impala %s with status %s", vw.Vw.ID, vw.Vw.Status))
		// Debug statement to delete
		if vw.Vw.Status == "Running" {
			if vw == nil {
				tflog.Debug(ctx, "Pillow: VW is nil")

			}

			if vw.Vw == nil {
				tflog.Debug(ctx, "Pillow: VW.Vw is nil")

			}

			if vw.Vw.Status != "Running" {
				tflog.Debug(ctx, fmt.Sprintf("Pillow: VW Status is not Running: %s", vw.Vw.Status))

			}

			tflog.Debug(ctx, fmt.Sprintf("==============================================================Pillow: Basic VW Details:"+
				"\nID: %s"+
				"\nName: %s"+
				"\nStatus: %s"+
				"\nAvailability Zone: %s"+
				"\nCDH Version: %s"+
				"\nCompactor: %v"+
				"\nConfig ID: %s"+
				"\nCRN: %s"+
				"\nDBC ID: %s"+
				"\nEBS LLAP Spill GB: %d"+
				"\nEnable Unified Analytics: %v"+
				"\nInstance Type: %s"+
				"\nMemory Capacity: %d"+
				"\nNode Count: %d"+
				"\nNumber of Cores: %d"+
				"\nResource Pool: %s"+
				"\nViz Enabled: %v"+
				"\nVW Type: %s"+
				"\nImpala Query Log: %v"+
				"\n==============================================================",
				vw.Vw.ID,
				vw.Vw.Name,
				vw.Vw.Status,
				vw.Vw.AvailabilityZone,
				vw.Vw.CdhVersion,
				vw.Vw.Compactor,
				vw.Vw.ConfigID,
				vw.Vw.Crn,
				vw.Vw.DbcID,
				vw.Vw.EbsLLAPSpillGB,
				vw.Vw.EnableUnifiedAnalytics,
				vw.Vw.InstanceType,
				vw.Vw.MemoryCapacity,
				vw.Vw.NodeCount,
				vw.Vw.NumOfCores,
				vw.Vw.ResourcePool,
				vw.Vw.Viz,
				vw.Vw.VwType,
				vw.Vw.ImpalaQueryLog))

			// Replica Status
			if vw.Vw.ReplicaStatus != nil {
				tflog.Debug(ctx, fmt.Sprintf("==============================================================Pillow: Replica Status:"+
					"\nReady Coordinator Replicas: %d"+
					"\nReady Executor Replicas: %d"+
					"\nTotal Coordinator Replicas: %d"+
					"\nTotal Executor Replicas: %d"+
					"\n==============================================================",
					vw.Vw.ReplicaStatus.ReadyCoordinatorReplicas,
					vw.Vw.ReplicaStatus.ReadyExecutorReplicas,
					vw.Vw.ReplicaStatus.TotalCoordinatorReplicas,
					vw.Vw.ReplicaStatus.TotalExecutorReplicas))
			} else {
				tflog.Debug(ctx, "Pillow: Replica Status is nil")
			}

			// Query Isolation Options
			if vw.Vw.QueryIsolationOptions != nil {
				tflog.Debug(ctx, fmt.Sprintf("==============================================================Pillow: Query Isolation Options:"+
					"\nMax Nodes Per Query: %d"+
					"\nMax Queries: %d"+
					"==============================================================",
					vw.Vw.QueryIsolationOptions.MaxNodesPerQuery,
					vw.Vw.QueryIsolationOptions.MaxQueries))
			} else {
				tflog.Debug(ctx, "Pillow: Query Isolation Options is nil")
			}

			// Impala Options
			if vw.Vw.ImpalaOptions != nil {
				tflog.Debug(ctx, fmt.Sprintf("Pillow: Impala Options:"+
					"\nScratch Space Limit: %d"+
					"\nSpill to S3 URI: %s"+
					"\n==============================================================",
					vw.Vw.ImpalaOptions.ScratchSpaceLimit,
					vw.Vw.ImpalaOptions.SpillToS3URI))
			} else {
				tflog.Debug(ctx, "Pillow: Impala Options is nil")
			}

			// Impala HA Settings
			if vw.Vw.ImpalaHaSettingsOptions != nil {
				tflog.Debug(ctx, fmt.Sprintf("==============================================================Pillow: Impala HA Settings:"+
					"\nEnable Catalog HA: %v"+
					"\nEnable Shutdown Of Coordinator: %v"+
					"\nEnable Statestore HA: %v"+
					"\nHA Mode: %v"+
					"\nNum Of Active Coordinators: %d"+
					"\n==============================================================",
					vw.Vw.ImpalaHaSettingsOptions.EnableCatalogHighAvailability,
					vw.Vw.ImpalaHaSettingsOptions.EnableShutdownOfCoordinator,
					vw.Vw.ImpalaHaSettingsOptions.EnableStatestoreHighAvailability,
					vw.Vw.ImpalaHaSettingsOptions.HighAvailabilityMode,
					vw.Vw.ImpalaHaSettingsOptions.NumOfActiveCoordinators))
			} else {
				tflog.Debug(ctx, "Pillow: Impala HA Settings is nil")
			}

			// JWT Auth
			if vw.Vw.JwtAuth != nil {
				tflog.Debug(ctx, fmt.Sprintf("Pillow: JWT Auth Provider: %s", vw.Vw.JwtAuth.Provider))
			} else {
				tflog.Debug(ctx, "Pillow: JWT Auth is nil")
			}

			// Supported Auth Methods
			if vw.Vw.SupportedAuthMethods != nil {
				if vw.Vw.SupportedAuthMethods.Jwt != nil && vw.Vw.SupportedAuthMethods.Ldap != nil && vw.Vw.SupportedAuthMethods.Sso != nil {
					tflog.Debug(ctx, fmt.Sprintf("==============================================================Pillow: Supported Auth Methods:"+
						"\nJWT: %v"+
						"\nLDAP: %v"+
						"\nSSO: %v"+
						"\n==============================================================",
						*vw.Vw.SupportedAuthMethods.Jwt,
						*vw.Vw.SupportedAuthMethods.Ldap,
						*vw.Vw.SupportedAuthMethods.Sso))
				} else {
					tflog.Debug(ctx, "Pillow: One or more Supported Auth Methods values are nil")
				}
			} else {
				tflog.Debug(ctx, "Pillow: Supported Auth Methods is nil")
			}

			// Autoscaling Options
			if vw.Vw.AutoscalingOptions != nil {
				tflog.Debug(ctx, fmt.Sprintf("==============================================================Pillow: Autoscaling Options:"+
					"\nMin Clusters: %d"+
					"\nMax Clusters: %d"+
					"\nDisable Auto Suspend: %v"+
					"\nAuto Suspend Timeout: %d"+
					"\nHive Scale Wait Time: %d"+
					"\nHive Desired Free Capacity: %d"+
					"\nImpala Scale Up Delay: %d"+
					"\nImpala Scale Down Delay: %d"+
					"\n==============================================================",
					vw.Vw.AutoscalingOptions.MinClusters,
					vw.Vw.AutoscalingOptions.MaxClusters,
					vw.Vw.AutoscalingOptions.DisableAutoSuspend,
					vw.Vw.AutoscalingOptions.AutoSuspendTimeoutSeconds,
					vw.Vw.AutoscalingOptions.HiveScaleWaitTimeSeconds,
					vw.Vw.AutoscalingOptions.HiveDesiredFreeCapacity,
					vw.Vw.AutoscalingOptions.ImpalaScaleUpDelaySeconds,
					vw.Vw.AutoscalingOptions.ImpalaScaleDownDelaySeconds))

				// Executor Group Sets
				if vw.Vw.AutoscalingOptions.ImpalaExecutorGroupSets != nil {
					if vw.Vw.AutoscalingOptions.ImpalaExecutorGroupSets.Small != nil {
						tflog.Debug(ctx, fmt.Sprintf("Pillow: Small Executor Group Set:"+
							"\nMin Groups: %d"+
							"\nMax Groups: %d"+
							"\nExec Group Size: %d"+
							"\nAuto Suspend Timeout: %d"+
							"\nDisable Auto Suspend: %v"+
							"\n==============================================================",
							vw.Vw.AutoscalingOptions.ImpalaExecutorGroupSets.Small.MinExecutorGroups,
							vw.Vw.AutoscalingOptions.ImpalaExecutorGroupSets.Small.MaxExecutorGroups,
							vw.Vw.AutoscalingOptions.ImpalaExecutorGroupSets.Small.ExecGroupSize,
							vw.Vw.AutoscalingOptions.ImpalaExecutorGroupSets.Small.AutoSuspendTimeoutSeconds,
							vw.Vw.AutoscalingOptions.ImpalaExecutorGroupSets.Small.DisableAutoSuspend))
					} else {
						tflog.Debug(ctx, "Pillow: Small Executor Group Set is nil")
					}

					if vw.Vw.AutoscalingOptions.ImpalaExecutorGroupSets.Large != nil {
						tflog.Debug(ctx, fmt.Sprintf("Pillow: Large Executor Group Set:"+
							"\nMin Groups: %d"+
							"\nMax Groups: %d"+
							"\nExec Group Size: %d"+
							"\nAuto Suspend Timeout: %d"+
							"\nDisable Auto Suspend: %v",
							vw.Vw.AutoscalingOptions.ImpalaExecutorGroupSets.Large.MinExecutorGroups,
							vw.Vw.AutoscalingOptions.ImpalaExecutorGroupSets.Large.MaxExecutorGroups,
							vw.Vw.AutoscalingOptions.ImpalaExecutorGroupSets.Large.ExecGroupSize,
							vw.Vw.AutoscalingOptions.ImpalaExecutorGroupSets.Large.AutoSuspendTimeoutSeconds,
							vw.Vw.AutoscalingOptions.ImpalaExecutorGroupSets.Large.DisableAutoSuspend))
					} else {
						tflog.Debug(ctx, "Pillow: Large Executor Group Set is nil")
					}
				} else {
					tflog.Debug(ctx, "Pillow: Impala Executor Group Sets is nil")
				}
			} else {
				tflog.Debug(ctx, "Pillow: Autoscaling Options is nil")
			}
		}
		return vw, vw.Vw.Status, nil
	}
}

func (r *impalaResource) createVwRequestFromPlan(plan *resourceModel) *models.CreateVwRequest {
	req := &models.CreateVwRequest{
		Name:      plan.Name.ValueStringPointer(),
		ClusterID: plan.ClusterID.ValueStringPointer(),
		DbcID:     plan.DatabaseCatalogID.ValueStringPointer(),
		VwType:    models.VwTypeImpala.Pointer(),
	}
	if imageVersion := plan.ImageVersion.ValueString(); imageVersion != "" {
		req.ImageVersion = imageVersion
	}

	if instanceType := plan.InstanceType.ValueString(); instanceType != "" {
		req.InstanceType = instanceType
	}

	if tshirtSize := plan.TShirtSize.ValueString(); tshirtSize != "" {
		req.TShirtSize = tshirtSize
	}

	if nodeCount := plan.NodeCount.ValueInt32(); nodeCount > 0 {
		req.NodeCount = nodeCount
	}

	if availabilityZone := plan.AvailabilityZone.ValueString(); availabilityZone != "" {
		req.AvailabilityZone = availabilityZone
	}

	if enableUnifiedAnalytics := plan.EnableUnifiedAnalytics.ValueBool(); enableUnifiedAnalytics != false {
		req.EnableUnifiedAnalytics = enableUnifiedAnalytics
	}

	if resourcePool := plan.ResourcePool.ValueString(); resourcePool != "" {
		req.ResourcePool = resourcePool
	}

	if platformJwtAuth := plan.PlatformJwtAuth.ValueBool(); platformJwtAuth != false {
		req.PlatformJwtAuth = &platformJwtAuth
	}

	if impalaQueryLog := plan.ImpalaQueryLog.ValueBool(); impalaQueryLog != false {
		req.ImpalaQueryLog = impalaQueryLog
	}

	if ebsLLAPSpillGB := plan.EbsLLAPSpillGB.ValueInt64(); ebsLLAPSpillGB > 0 {
		req.EbsLLAPSpillGB = int32(ebsLLAPSpillGB)
	}

	/*if hiveServerHaMode := plan.HiveServerHaMode.ValueStringPointer(); *hiveServerHaMode != "" {
		req.HiveServerHaMode = hiveServerHaMode
	}*/

	if !plan.ImpalaOptions.IsNull() {
		req.ImpalaOptions = convertToAPIImpalaOptions(plan.ImpalaOptions)
	}

	if !plan.Autoscaling.IsNull() {
		req.Autoscaling = convertToAPIAutoscaling(plan.Autoscaling)
	}

	if !plan.ImpalaHASettings.IsNull() {
		req.ImpalaHaSettings = convertToAPIImpalaHASettings(plan.ImpalaHASettings)
	}

	if !plan.QueryIsolationOptions.IsNull() {
		req.QueryIsolationOptions = convertToAPIQueryIsolationOptions(plan.QueryIsolationOptions)
	}

	/*if !plan.Config.IsNull() {
		req.Config = ConvertToServiceConfigReqModel(plan.Config)
	}*/

	if len(plan.Tags.Elements()) > 0 {
		req.Tags = convertToAPITagRequests(plan.Tags)
	}

	return req
}

func (r *impalaResource) waitForVwRunning(ctx context.Context, plan *resourceModel, vwID *string) error {
	clusterID := plan.ClusterID.ValueStringPointer()

	if opts := plan.PollingOptions; opts == nil || !opts.Async.ValueBool() {
		callFailedCount := 0
		stateConf := &retry.StateChangeConf{
			Pending:      []string{"Accepted", "Creating", "Created", "Starting"},
			Target:       []string{"Running"},
			Delay:        30 * time.Second,
			Timeout:      utils.GetPollingTimeout(plan, 20*time.Minute),
			PollInterval: 30 * time.Second,
			Refresh:      r.stateRefresh(ctx, clusterID, vwID, &callFailedCount, utils.GetCallFailureThreshold(plan, 3)),
		}
		_, err := stateConf.WaitForStateContext(ctx)
		return err
	}
	return nil
}

func (r *impalaResource) populatePlanFromDescribe(ctx context.Context, plan *resourceModel, vwID *string) error {
	desc := operations.NewDescribeVwParamsWithContext(ctx).
		WithInput(&models.DescribeVwRequest{VwID: vwID, ClusterID: plan.ClusterID.ValueStringPointer()})
	describe, err := r.client.Dw.Operations.DescribeVw(desc)
	if err != nil {
		return err
	}

	impala := describe.GetPayload()
	tflog.Info(context.Background(), fmt.Sprintf("Prateek API Response: %+v", impala))
	plan.setFromDescribeVwResponse(impala)

	return nil
}

func (r *impalaResource) deleteVirtualWarehouse(ctx context.Context, clusterID, vwID *string) error {
	op := operations.NewDeleteVwParamsWithContext(ctx).
		WithInput(&models.DeleteVwRequest{
			ClusterID: clusterID,
			VwID:      vwID,
		})

	_, err := r.client.Dw.Operations.DeleteVw(op)
	return err
}

func (r *impalaResource) pollForDeletion(ctx context.Context, state resourceModel, clusterID, vwID *string) error {
	callFailedCount := 0
	stateConf := &retry.StateChangeConf{
		Pending:      []string{"Deleting", "Running", "Stopping", "Stopped", "Creating", "Created", "Starting", "Updating"},
		Target:       []string{"Deleted"},
		Delay:        30 * time.Second,
		Timeout:      utils.GetPollingTimeout(&state, 20*time.Minute),
		PollInterval: 30 * time.Second,
		Refresh:      r.stateRefresh(ctx, clusterID, vwID, &callFailedCount, utils.GetCallFailureThreshold(&state, 3)),
	}

	_, err := stateConf.WaitForStateContext(ctx)
	return err
}
