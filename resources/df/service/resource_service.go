// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/df/client/operations"
	dfmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/df/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ resource.Resource                = (*dfServiceResource)(nil)
	_ resource.ResourceWithConfigure   = (*dfServiceResource)(nil)
	_ resource.ResourceWithImportState = (*dfServiceResource)(nil)
)

type dfServiceResource struct {
	client *cdp.Client
}

func NewDfServiceResource() resource.Resource {
	return &dfServiceResource{}
}

func (r *dfServiceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_df_service"
}

func (r *dfServiceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *dfServiceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = serviceSchema
}

func (r *dfServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan serviceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	clusterSubnets := utils.FromListValueToStringList(plan.ClusterSubnets)
	lbSubnets := utils.FromListValueToStringList(plan.LoadBalancerSubnets)
	kubeAPIRanges := utils.FromListValueToStringList(plan.KubeAPIAuthorizedIPRanges)
	lbRanges := utils.FromListValueToStringList(plan.LoadBalancerAuthorizedIPRanges)
	tags := utils.FromMapValueToStringMap(plan.Tags)

	input := &dfmodels.EnableServiceRequest{
		EnvironmentCrn:                 plan.EnvironmentCrn.ValueStringPointer(),
		MinK8sNodeCount:                plan.MinK8sNodeCount.ValueInt32Pointer(),
		MaxK8sNodeCount:                plan.MaxK8sNodeCount.ValueInt32Pointer(),
		UsePublicLoadBalancer:          plan.UsePublicLoadBalancer.ValueBoolPointer(),
		PrivateCluster:                 plan.PrivateCluster.ValueBool(),
		ClusterSubnets:                 clusterSubnets,
		LoadBalancerSubnets:            lbSubnets,
		KubeAPIAuthorizedIPRanges:      kubeAPIRanges,
		LoadBalancerAuthorizedIPRanges: lbRanges,
		Tags:                           tags,
		SkipPreflightChecks:            plan.SkipPreflightChecks.ValueBool(),
		UserDefinedRouting:             plan.UserDefinedRouting.ValueBool(),
	}
	if !plan.InstanceType.IsNull() && !plan.InstanceType.IsUnknown() {
		input.InstanceType = plan.InstanceType.ValueString()
	}
	if !plan.PodCidr.IsNull() && !plan.PodCidr.IsUnknown() {
		input.PodCidr = plan.PodCidr.ValueString()
	}
	if !plan.ServiceCidr.IsNull() && !plan.ServiceCidr.IsUnknown() {
		input.ServiceCidr = plan.ServiceCidr.ValueString()
	}

	params := operations.NewEnableServiceParamsWithContext(ctx).WithInput(input)
	response, err := r.client.Df.Operations.EnableService(params)
	if err != nil {
		resp.Diagnostics.AddError("Error enabling DataFlow service", err.Error())
		return
	}

	svc := response.GetPayload().Service
	plan.Crn = types.StringPointerValue(svc.Crn)
	plan.ID = types.StringPointerValue(svc.Crn)
	plan.Name = types.StringPointerValue(svc.Name)
	plan.CloudPlatform = types.StringPointerValue(svc.CloudPlatform)
	plan.Region = types.StringPointerValue(svc.Region)
	plan.Status = types.StringValue(string(*svc.Status.State))
	plan.StatusMessage = types.StringValue(svc.Status.Message)
	plan.WorkloadVersion = types.StringPointerValue(svc.WorkloadVersion)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	timeout := time.Duration(plan.PollingTimeout.ValueInt64()) * time.Second
	callFailedCount := 0
	stateConf := &retry.StateChangeConf{
		Pending:      []string{"ENABLING"},
		Target:       []string{"GOOD_HEALTH", "CONCERNING_HEALTH", "BAD_HEALTH"},
		Delay:        30 * time.Second,
		Timeout:      timeout,
		PollInterval: 30 * time.Second,
		Refresh:      r.stateRefresh(ctx, svc.Crn, &callFailedCount, 3),
	}
	if _, err = stateConf.WaitForStateContext(ctx); err != nil {
		resp.Diagnostics.AddError("Error waiting for DataFlow service to become ready", err.Error())
		return
	}

	if err = r.refreshState(ctx, &plan); err != nil {
		resp.Diagnostics.AddError("Error reading DataFlow service after create", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *dfServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state serviceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.refreshState(ctx, &state); err != nil {
		if strings.Contains(err.Error(), "NOT_FOUND") {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading DataFlow service", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *dfServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan serviceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state serviceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := &dfmodels.UpdateServiceRequest{
		ServiceCrn:      state.Crn.ValueStringPointer(),
		MinK8sNodeCount: plan.MinK8sNodeCount.ValueInt32Pointer(),
		MaxK8sNodeCount: plan.MaxK8sNodeCount.ValueInt32Pointer(),
	}
	if !plan.KubeAPIAuthorizedIPRanges.IsNull() {
		input.KubeAPIAuthorizedIPRanges = utils.FromListValueToStringList(plan.KubeAPIAuthorizedIPRanges)
	}
	if !plan.LoadBalancerAuthorizedIPRanges.IsNull() {
		input.LoadBalancerAuthorizedIPRanges = utils.FromListValueToStringList(plan.LoadBalancerAuthorizedIPRanges)
	}

	params := operations.NewUpdateServiceParamsWithContext(ctx).WithInput(input)
	if _, err := r.client.Df.Operations.UpdateService(params); err != nil {
		resp.Diagnostics.AddError("Error updating DataFlow service", err.Error())
		return
	}

	if err := r.refreshState(ctx, &plan); err != nil {
		resp.Diagnostics.AddError("Error reading DataFlow service after update", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *dfServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state serviceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewDisableServiceParamsWithContext(ctx).WithInput(&dfmodels.DisableServiceRequest{
		ServiceCrn:           state.Crn.ValueStringPointer(),
		TerminateDeployments: state.TerminateDeploymentsOnDisable.ValueBool(),
	})
	if _, err := r.client.Df.Operations.DisableService(params); err != nil {
		if strings.Contains(err.Error(), "NOT_FOUND") {
			return
		}
		resp.Diagnostics.AddError("Error disabling DataFlow service", err.Error())
		return
	}

	timeout := time.Duration(state.PollingTimeout.ValueInt64()) * time.Second
	callFailedCount := 0
	stateConf := &retry.StateChangeConf{
		Pending:      []string{"DISABLING", "GOOD_HEALTH", "CONCERNING_HEALTH", "BAD_HEALTH", "ENABLING", "UPDATING"},
		Target:       []string{"DISABLED"},
		Delay:        30 * time.Second,
		Timeout:      timeout,
		PollInterval: 30 * time.Second,
		Refresh:      r.stateRefresh(ctx, state.Crn.ValueStringPointer(), &callFailedCount, 3),
	}
	if _, err := stateConf.WaitForStateContext(ctx); err != nil {
		resp.Diagnostics.AddError("Error waiting for DataFlow service to be disabled", err.Error())
	}
}

func (r *dfServiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *dfServiceResource) refreshState(ctx context.Context, state *serviceModel) error {
	params := operations.NewDescribeServiceParamsWithContext(ctx).WithInput(&dfmodels.DescribeServiceRequest{
		ServiceCrn: state.Crn.ValueStringPointer(),
	})
	resp, err := r.client.Df.Operations.DescribeService(params)
	if err != nil {
		return err
	}
	svc := resp.GetPayload().Service
	if svc == nil {
		return fmt.Errorf("service not found")
	}
	state.Crn = types.StringPointerValue(svc.Crn)
	state.ID = types.StringPointerValue(svc.Crn)
	state.Name = types.StringPointerValue(svc.Name)
	state.CloudPlatform = types.StringPointerValue(svc.CloudPlatform)
	state.Region = types.StringPointerValue(svc.Region)
	state.Status = types.StringValue(string(*svc.Status.State))
	state.StatusMessage = types.StringValue(svc.Status.Message)
	state.WorkloadVersion = types.StringPointerValue(svc.WorkloadVersion)
	state.MinK8sNodeCount = types.Int32PointerValue(svc.MinK8sNodeCount)
	state.MaxK8sNodeCount = types.Int32PointerValue(svc.MaxK8sNodeCount)
	state.UsePublicLoadBalancer = types.BoolValue(svc.UsePublicLoadBalancer)
	state.PrivateCluster = types.BoolValue(svc.PrivateCluster)
	state.UserDefinedRouting = types.BoolValue(svc.UserDefinedRouting)
	if svc.InstanceType != "" {
		state.InstanceType = types.StringValue(svc.InstanceType)
	}
	if svc.PodCidr != "" {
		state.PodCidr = types.StringValue(svc.PodCidr)
	}
	if svc.ServiceCidr != "" {
		state.ServiceCidr = types.StringValue(svc.ServiceCidr)
	}
	return nil
}

func (r *dfServiceResource) stateRefresh(ctx context.Context, crn *string, callFailedCount *int, callFailureThreshold int) func() (any, string, error) {
	return func() (any, string, error) {
		params := operations.NewDescribeServiceParamsWithContext(ctx).WithInput(&dfmodels.DescribeServiceRequest{
			ServiceCrn: crn,
		})
		resp, err := r.client.Df.Operations.DescribeService(params)
		if err != nil {
			if strings.Contains(err.Error(), "NOT_FOUND") {
				return &dfmodels.DescribeServiceResponse{}, "DISABLED", nil
			}
			*callFailedCount++
			if *callFailedCount <= callFailureThreshold {
				tflog.Warn(ctx, fmt.Sprintf("could not describe DataFlow service due to [%s] but threshold not reached (%d/%d).", err.Error(), *callFailedCount, callFailureThreshold))
				return nil, "", nil
			}
			return nil, "", err
		}
		*callFailedCount = 0
		svc := resp.GetPayload().Service
		if svc == nil {
			return &dfmodels.DescribeServiceResponse{}, "DISABLED", nil
		}
		state := string(*svc.Status.State)
		tflog.Debug(ctx, fmt.Sprintf("DataFlow service %s status: %s", *crn, state))
		return svc, state, nil
	}
}
