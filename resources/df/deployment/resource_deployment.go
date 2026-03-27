// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package deployment

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	dlog "log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/df/client/operations"
	iamoperations "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/client/operations"
	iammodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
	dfmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/df/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ resource.Resource                = (*dfDeploymentResource)(nil)
	_ resource.ResourceWithConfigure   = (*dfDeploymentResource)(nil)
	_ resource.ResourceWithImportState = (*dfDeploymentResource)(nil)
)

type dfDeploymentResource struct {
	client *cdp.Client
}

func NewDfDeploymentResource() resource.Resource {
	return &dfDeploymentResource{}
}

func (r *dfDeploymentResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_df_deployment"
}

func (r *dfDeploymentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *dfDeploymentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = deploymentSchema
}

func (r *dfDeploymentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan deploymentModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check if a deployment with this name already exists
	existingCrn := r.findDeploymentByName(ctx, plan.DeploymentName.ValueString(), plan.ServiceCrn.ValueString())
	if existingCrn != "" {
		// Adopt the existing deployment and do a change-flow-version
		tflog.Info(ctx, fmt.Sprintf("Deployment %s already exists (crn=%s), adopting and changing flow version", plan.DeploymentName.ValueString(), existingCrn))

		plan.DeploymentCrn = types.StringValue(existingCrn)
		plan.ID = types.StringValue(existingCrn)

		// Read current state to populate service_crn
		state := plan
		state.DeploymentCrn = types.StringValue(existingCrn)
		if err := r.refreshState(ctx, &state); err != nil {
			resp.Diagnostics.AddError("Error reading existing deployment", err.Error())
			return
		}

		// Change flow version if different
		if plan.FlowVersionCrn.ValueString() != state.FlowVersionCrn.ValueString() {
			if err := r.changeFlowVersion(ctx, &state, &plan); err != nil {
				resp.Diagnostics.AddError("Error changing flow version on existing deployment", err.Error())
				return
			}

			timeout := time.Duration(plan.PollingTimeout.ValueInt64()) * time.Second
			callFailedCount := 0
			stateConf := &retry.StateChangeConf{
				Pending:      []string{"DEPLOYING", "STARTING_FLOW", "IMPORTING_FLOW", "UPDATING"},
				Target:       []string{"GOOD_HEALTH", "CONCERNING_HEALTH"},
				Delay:        30 * time.Second,
				Timeout:      timeout,
				PollInterval: 30 * time.Second,
				Refresh:      r.stateRefresh(ctx, &existingCrn, &callFailedCount, 3),
			}
			if _, err := stateConf.WaitForStateContext(ctx); err != nil {
				resp.Diagnostics.AddError("Error waiting for deployment after flow version change", err.Error())
				return
			}
		}

		if err := r.refreshState(ctx, &plan); err != nil {
			resp.Diagnostics.AddError("Error reading deployment after adoption", err.Error())
			return
		}
		computeParameterGroupsSha(&plan)
		resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
		return
	}

	// Step 1: Initiate deployment on control plane
	params := operations.NewInitiateDeploymentParamsWithContext(ctx).WithInput(&dfmodels.InitiateDeploymentRequest{
		ServiceCrn:     plan.ServiceCrn.ValueStringPointer(),
		FlowVersionCrn: plan.FlowVersionCrn.ValueStringPointer(),
	})
	response, err := r.client.Df.Operations.InitiateDeployment(params)
	if err != nil {
		resp.Diagnostics.AddError("Error initiating DataFlow deployment", err.Error())
		return
	}
	deploymentRequestCrn := *response.GetPayload().DeploymentRequestCrn
	tflog.Info(ctx, fmt.Sprintf("Initiated deployment request: %s", deploymentRequestCrn))

	// Step 2: Get environment CRN
	environmentCrn, err := r.getEnvironmentCrn(ctx, plan.ServiceCrn.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error getting environment CRN", err.Error())
		return
	}

	// Step 3: Get workload auth token
	workloadURL, bearerToken, err := r.getWorkloadAuth(ctx, environmentCrn)
	if err != nil {
		resp.Diagnostics.AddError("Error getting workload auth token", err.Error())
		return
	}

	// Step 4: Get deployment request details
	err = r.workloadPost(ctx, workloadURL, bearerToken,
		"/dfx/api/rpc-v1/deployments/get-deployment-request-details",
		map[string]interface{}{
			"environmentCrn":       environmentCrn,
			"deploymentRequestCrn": deploymentRequestCrn,
		}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error getting deployment request details", err.Error())
		return
	}

	// Step 5: Create deployment on workload
	deployConfig := map[string]interface{}{
		"environmentCrn":       environmentCrn,
		"name":                 plan.DeploymentName.ValueString(),
		"deploymentRequestCrn": deploymentRequestCrn,
		"configurationVersion": 0,
		"staticNodeCount":      1,
	}

	// Cluster size
	clusterSize := "EXTRA_SMALL"
	if !plan.ClusterSize.IsNull() && plan.ClusterSize.ValueString() != "" {
		clusterSize = plan.ClusterSize.ValueString()
	}
	deployConfig["clusterSize"] = map[string]string{"name": clusterSize}

	if !plan.CfmNifiVersion.IsNull() && plan.CfmNifiVersion.ValueString() != "" {
		deployConfig["cfmNifiVersion"] = plan.CfmNifiVersion.ValueString()
	}
	autoStart := false
	if !plan.AutoStartFlow.IsNull() {
		autoStart = plan.AutoStartFlow.ValueBool()
	}
	deployConfig["autoStartFlow"] = autoStart
	if !plan.ProjectCrn.IsNull() && plan.ProjectCrn.ValueString() != "" {
		deployConfig["projectCrn"] = plan.ProjectCrn.ValueString()
	}
	if !plan.StaticNodeCount.IsNull() {
		deployConfig["staticNodeCount"] = plan.StaticNodeCount.ValueInt64()
	}
	if !plan.AutoScalingEnabled.IsNull() && plan.AutoScalingEnabled.ValueBool() {
		deployConfig["autoScalingEnabled"] = true
		delete(deployConfig, "staticNodeCount")
		if !plan.AutoScaleMinNodes.IsNull() {
			deployConfig["autoScaleMinNodes"] = plan.AutoScaleMinNodes.ValueInt64()
		}
		if !plan.AutoScaleMaxNodes.IsNull() {
			deployConfig["autoScaleMaxNodes"] = plan.AutoScaleMaxNodes.ValueInt64()
		}
	}
	if !plan.ParameterGroups.IsNull() && plan.ParameterGroups.ValueString() != "" {
		var pg []interface{}
		if err := json.Unmarshal([]byte(plan.ParameterGroups.ValueString()), &pg); err != nil {
			resp.Diagnostics.AddError("Error parsing parameter_groups JSON", err.Error())
			return
		}
		deployConfig["parameterGroups"] = pg
	}

	var createResult struct {
		Deployment struct {
			Crn string `json:"crn"`
		} `json:"deployment"`
	}
	err = r.workloadPost(ctx, workloadURL, bearerToken,
		"/dfx/api/rpc-v1/deployments/create-deployment",
		deployConfig, &createResult)
	if err != nil {
		resp.Diagnostics.AddError("Error creating deployment on workload", err.Error())
		return
	}

	deploymentCrn := createResult.Deployment.Crn
	if deploymentCrn == "" {
		// Fallback: poll ListDeployments to find it
		timeout := time.Duration(plan.PollingTimeout.ValueInt64()) * time.Second
		err = retry.RetryContext(ctx, timeout, func() *retry.RetryError {
			listParams := operations.NewListDeploymentsParamsWithContext(ctx).WithInput(&dfmodels.ListDeploymentsRequest{})
			listResp, listErr := r.client.Df.Operations.ListDeployments(listParams)
			if listErr != nil {
				return retry.NonRetryableError(listErr)
			}
			for _, dep := range listResp.GetPayload().Deployments {
				if dep.Name != nil && *dep.Name == plan.DeploymentName.ValueString() {
					deploymentCrn = *dep.Crn
					return nil
				}
			}
			return retry.RetryableError(fmt.Errorf("deployment not yet available"))
		})
		if err != nil {
			resp.Diagnostics.AddError("Error finding deployment after creation", err.Error())
			return
		}
	}

	plan.DeploymentCrn = types.StringValue(deploymentCrn)
	plan.ID = types.StringValue(deploymentCrn)

	// Poll until deployment reaches a healthy state
	timeout := time.Duration(plan.PollingTimeout.ValueInt64()) * time.Second
	callFailedCount := 0
	stateConf := &retry.StateChangeConf{
		Pending:      []string{"DEPLOYING", "STARTING_FLOW", "IMPORTING_FLOW"},
		Target:       []string{"GOOD_HEALTH", "CONCERNING_HEALTH"},
		Delay:        30 * time.Second,
		Timeout:      timeout,
		PollInterval: 30 * time.Second,
		Refresh:      r.stateRefresh(ctx, &deploymentCrn, &callFailedCount, 3),
	}
	if _, err = stateConf.WaitForStateContext(ctx); err != nil {
		resp.Diagnostics.AddError("Error waiting for DataFlow deployment to become ready", err.Error())
		return
	}

	if err = r.refreshState(ctx, &plan); err != nil {
		resp.Diagnostics.AddError("Error reading DataFlow deployment after create", err.Error())
		return
	}
	computeParameterGroupsSha(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// computeParameterGroupsSha sets the SHA256 hash of parameter_groups content.
func computeParameterGroupsSha(m *deploymentModel) {
	if m.ParameterGroups.IsNull() || m.ParameterGroups.ValueString() == "" {
		m.ParameterGroupsSha = types.StringValue("")
		return
	}
	hash := sha256.Sum256([]byte(m.ParameterGroups.ValueString()))
	m.ParameterGroupsSha = types.StringValue(fmt.Sprintf("%x", hash))
}

func (r *dfDeploymentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state deploymentModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.refreshState(ctx, &state); err != nil {
		if strings.Contains(err.Error(), "NOT_FOUND") {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading DataFlow deployment", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *dfDeploymentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan deploymentModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state deploymentModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Only trigger change-flow-version if flow_version_crn changed
	if plan.FlowVersionCrn.ValueString() != state.FlowVersionCrn.ValueString() {
		if err := r.changeFlowVersion(ctx, &state, &plan); err != nil {
			resp.Diagnostics.AddError("Error changing flow version", err.Error())
			return
		}

		// Poll until deployment reaches a healthy state
		timeout := time.Duration(plan.PollingTimeout.ValueInt64()) * time.Second
		callFailedCount := 0
		stateConf := &retry.StateChangeConf{
			Pending:      []string{"DEPLOYING", "STARTING_FLOW", "IMPORTING_FLOW", "UPDATING"},
			Target:       []string{"GOOD_HEALTH", "CONCERNING_HEALTH"},
			Delay:        30 * time.Second,
			Timeout:      timeout,
			PollInterval: 30 * time.Second,
			Refresh:      r.stateRefresh(ctx, state.DeploymentCrn.ValueStringPointer(), &callFailedCount, 3),
		}
		if _, err := stateConf.WaitForStateContext(ctx); err != nil {
			resp.Diagnostics.AddError("Error waiting for deployment to become ready after flow version change", err.Error())
			return
		}
	}

	plan.DeploymentCrn = state.DeploymentCrn
	plan.ID = state.ID
	if err := r.refreshState(ctx, &plan); err != nil {
		resp.Diagnostics.AddError("Error reading deployment after update", err.Error())
		return
	}
	computeParameterGroupsSha(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// changeFlowVersion orchestrates the multi-step change-flow-version workflow:
// 1. initiateDeployment on control plane (with existing deploymentCrn + new flowVersionCrn)
// 2. describeService to get environmentCrn
// 3. generateWorkloadAuthToken via IAM
// 4. getDeploymentRequestDetails on workload API
// 5. changeFlowVersion on workload API
func (r *dfDeploymentResource) changeFlowVersion(ctx context.Context, state *deploymentModel, plan *deploymentModel) error {
	// Step 1: Initiate deployment (with existing deployment CRN for change-flow-version)
	initiateParams := operations.NewInitiateDeploymentParamsWithContext(ctx).WithInput(&dfmodels.InitiateDeploymentRequest{
		ServiceCrn:     state.ServiceCrn.ValueStringPointer(),
		FlowVersionCrn: plan.FlowVersionCrn.ValueStringPointer(),
		DeploymentCrn:  state.DeploymentCrn.ValueString(),
	})
	initiateResp, err := r.client.Df.Operations.InitiateDeployment(initiateParams)
	if err != nil {
		return fmt.Errorf("initiateDeployment: %w", err)
	}
	deploymentRequestCrn := *initiateResp.GetPayload().DeploymentRequestCrn
	tflog.Info(ctx, fmt.Sprintf("Initiated change-flow-version request: %s", deploymentRequestCrn))

	// Step 2: Get environment CRN
	environmentCrn, err := r.getEnvironmentCrn(ctx, state.ServiceCrn.ValueString())
	if err != nil {
		return fmt.Errorf("getEnvironmentCrn: %w", err)
	}

	// Step 3: Get workload auth token
	workloadURL, bearerToken, err := r.getWorkloadAuth(ctx, environmentCrn)
	if err != nil {
		return fmt.Errorf("getWorkloadAuth: %w", err)
	}

	// Step 4: getDeploymentRequestDetails on workload API
	err = r.workloadPost(ctx, workloadURL, bearerToken,
		"/dfx/api/rpc-v1/deployments/get-deployment-request-details",
		map[string]interface{}{
			"environmentCrn":       environmentCrn,
			"deploymentRequestCrn": deploymentRequestCrn,
		}, nil)
	if err != nil {
		return fmt.Errorf("getDeploymentRequestDetails: %w", err)
	}

	// Step 5: changeFlowVersion on workload API
	cfvBody := map[string]interface{}{
		"environmentCrn":       environmentCrn,
		"deploymentRequestCrn": deploymentRequestCrn,
		"deploymentCrn":        state.DeploymentCrn.ValueString(),
	}
	if !plan.Strategy.IsNull() && plan.Strategy.ValueString() != "" {
		cfvBody["strategy"] = plan.Strategy.ValueString()
	}
	if !plan.WaitForFlowToStopInMinutes.IsNull() {
		cfvBody["waitForFlowToStopInMinutes"] = plan.WaitForFlowToStopInMinutes.ValueInt64()
	}
	if !plan.ParameterGroups.IsNull() && plan.ParameterGroups.ValueString() != "" {
		var pg []interface{}
		if err := json.Unmarshal([]byte(plan.ParameterGroups.ValueString()), &pg); err != nil {
			return fmt.Errorf("parsing parameter_groups: %w", err)
		}
		cfvBody["parameterGroups"] = pg
	}

	err = r.workloadPost(ctx, workloadURL, bearerToken,
		"/dfx/api/rpc-v1/deployments/change-flow-version",
		cfvBody, nil)
	if err != nil {
		return fmt.Errorf("changeFlowVersion: %w", err)
	}

	return nil
}

// findDeploymentByName searches for an existing deployment by name on a service and returns its CRN.
func (r *dfDeploymentResource) findDeploymentByName(ctx context.Context, name string, serviceCrn string) string {
	logFile, _ := os.OpenFile("/tmp/df_deployment_find.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	defer logFile.Close()
	dbg := dlog.New(logFile, "", dlog.LstdFlags)

	dbg.Printf("Looking for deployment name=%s service_crn=%s", name, serviceCrn)

	listParams := operations.NewListDeploymentsParamsWithContext(ctx).WithInput(&dfmodels.ListDeploymentsRequest{})
	listResp, err := r.client.Df.Operations.ListDeployments(listParams)
	if err != nil {
		dbg.Printf("ERROR ListDeployments: %s", err)
		return ""
	}
	for _, dep := range listResp.GetPayload().Deployments {
		depName := ""
		depSvcCrn := ""
		depCrn := ""
		if dep.Name != nil {
			depName = *dep.Name
		}
		if dep.Service != nil && dep.Service.Crn != nil {
			depSvcCrn = *dep.Service.Crn
		}
		if dep.Crn != nil {
			depCrn = *dep.Crn
		}
		dbg.Printf("Found deployment: name=%s service_crn=%s crn=%s", depName, depSvcCrn, depCrn)
		if depName == name {
			if serviceCrn == "" || depSvcCrn == serviceCrn {
				dbg.Printf("MATCH found: %s", depCrn)
				return depCrn
			}
		}
	}
	dbg.Printf("No match found")
	return ""
}

// getEnvironmentCrn resolves the environment CRN from a service CRN.
func (r *dfDeploymentResource) getEnvironmentCrn(ctx context.Context, serviceCrn string) (string, error) {
	descSvcParams := operations.NewDescribeServiceParamsWithContext(ctx).WithInput(&dfmodels.DescribeServiceRequest{
		ServiceCrn: &serviceCrn,
	})
	svcResp, err := r.client.Df.Operations.DescribeService(descSvcParams)
	if err != nil {
		return "", err
	}
	return *svcResp.GetPayload().Service.EnvironmentCrn, nil
}

// getWorkloadAuth gets a workload bearer token and endpoint URL for the DF workload API.
func (r *dfDeploymentResource) getWorkloadAuth(ctx context.Context, environmentCrn string) (workloadURL string, bearerToken string, err error) {
	workloadName := iammodels.WorkloadName("DF")
	tokenParams := iamoperations.NewGenerateWorkloadAuthTokenParamsWithContext(ctx).WithInput(&iammodels.GenerateWorkloadAuthTokenRequest{
		WorkloadName:   &workloadName,
		EnvironmentCrn: environmentCrn,
	})
	tokenResp, err := r.client.Iam.Operations.GenerateWorkloadAuthToken(tokenParams)
	if err != nil {
		return "", "", err
	}
	workloadURL = strings.TrimRight(tokenResp.GetPayload().EndpointURL, "/") + "/"
	bearerToken = "Bearer " + tokenResp.GetPayload().Token
	return workloadURL, bearerToken, nil
}

// workloadPost sends a JSON POST to the DFX workload API with bearer token auth.
func (r *dfDeploymentResource) workloadPost(ctx context.Context, baseURL, bearerToken, apiPath string, body map[string]interface{}, result interface{}) error {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	reqURL := strings.TrimRight(baseURL, "/") + apiPath
	httpReq, err := http.NewRequestWithContext(ctx, "POST", reqURL, bytes.NewReader(jsonBody))
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", bearerToken)
	httpReq.ContentLength = int64(len(jsonBody))

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}

	if httpResp.StatusCode >= 400 {
		return fmt.Errorf("status %d: %s", httpResp.StatusCode, string(respBody))
	}

	if result != nil {
		return json.Unmarshal(respBody, result)
	}
	return nil
}

func (r *dfDeploymentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state deploymentModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	logFile, _ := os.OpenFile("/tmp/df_deployment_delete.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	defer logFile.Close()
	dbg := dlog.New(logFile, "", dlog.LstdFlags)

	dbg.Printf("Starting delete for deployment_crn=%s service_crn=%s", state.DeploymentCrn.ValueString(), state.ServiceCrn.ValueString())

	environmentCrn, err := r.getEnvironmentCrn(ctx, state.ServiceCrn.ValueString())
	if err != nil {
		dbg.Printf("ERROR getEnvironmentCrn: %s", err)
		resp.Diagnostics.AddError("Error getting environment CRN for termination", err.Error())
		return
	}
	dbg.Printf("Got environmentCrn=%s", environmentCrn)

	workloadURL, bearerToken, err := r.getWorkloadAuth(ctx, environmentCrn)
	if err != nil {
		dbg.Printf("ERROR getWorkloadAuth: %s", err)
		resp.Diagnostics.AddError("Error getting workload auth for termination", err.Error())
		return
	}
	dbg.Printf("Got workloadURL=%s", workloadURL)

	dbg.Printf("Calling terminate-deployment")
	err = r.workloadPost(ctx, workloadURL, bearerToken,
		"/dfx/api/rpc-v1/deployments/terminate-deployment",
		map[string]interface{}{
			"environmentCrn": environmentCrn,
			"deploymentCrn":  state.DeploymentCrn.ValueString(),
		}, nil)
	if err != nil {
		dbg.Printf("ERROR terminate-deployment: %s", err)
		if !strings.Contains(err.Error(), "NOT_FOUND") {
			resp.Diagnostics.AddError("Error terminating DataFlow deployment", err.Error())
			return
		}
		dbg.Printf("Deployment already gone")
		return
	}
	dbg.Printf("Terminate sent successfully, polling")

	timeout := time.Duration(state.PollingTimeout.ValueInt64()) * time.Second
	callFailedCount := 0
	stateConf := &retry.StateChangeConf{
		Pending:      []string{"GOOD_HEALTH", "CONCERNING_HEALTH", "BAD_HEALTH", "TERMINATING", "SUSPENDING", "SUSPENDED", "DEPLOYING", "STARTING_FLOW", "UPDATING"},
		Target:       []string{"TERMINATED"},
		Delay:        10 * time.Second,
		Timeout:      timeout,
		PollInterval: 15 * time.Second,
		Refresh:      r.stateRefresh(ctx, state.DeploymentCrn.ValueStringPointer(), &callFailedCount, 3),
	}
	if _, err := stateConf.WaitForStateContext(ctx); err != nil {
		dbg.Printf("ERROR polling: %s", err)
		if !strings.Contains(err.Error(), "NOT_FOUND") {
			resp.Diagnostics.AddError("Error waiting for DataFlow deployment to terminate", err.Error())
		}
	}
	dbg.Printf("Delete complete")
}

func (r *dfDeploymentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *dfDeploymentResource) refreshState(ctx context.Context, state *deploymentModel) error {
	params := operations.NewDescribeDeploymentParamsWithContext(ctx).WithInput(&dfmodels.DescribeDeploymentRequest{
		DeploymentCrn: state.DeploymentCrn.ValueStringPointer(),
	})
	resp, err := r.client.Df.Operations.DescribeDeployment(params)
	if err != nil {
		return err
	}
	dep := resp.GetPayload().Deployment
	if dep == nil {
		return fmt.Errorf("deployment not found")
	}
	state.DeploymentCrn = types.StringPointerValue(dep.Crn)
	state.ID = types.StringPointerValue(dep.Crn)
	state.Name = types.StringPointerValue(dep.Name)
	state.FlowCrn = types.StringPointerValue(dep.FlowCrn)
	state.FlowVersionCrn = types.StringPointerValue(dep.FlowVersionCrn)
	state.ClusterSize = types.StringPointerValue(dep.ClusterSize)
	state.CurrentNodeCount = types.Int32Value(dep.CurrentNodeCount)
	state.Status = types.StringValue(string(*dep.Status.State))
	state.StatusMessage = types.StringPointerValue(dep.Status.Message)
	if dep.FlowName != "" {
		state.FlowName = types.StringValue(dep.FlowName)
	}
	if dep.CfmNifiVersion != "" {
		state.CfmNifiVersion = types.StringValue(dep.CfmNifiVersion)
	}
	if dep.NifiURL != "" {
		state.NifiURL = types.StringValue(dep.NifiURL)
	}
	if dep.DeployedByName != "" {
		state.DeployedByName = types.StringValue(dep.DeployedByName)
	}
	if dep.Service != nil && dep.Service.Crn != nil {
		state.ServiceCrn = types.StringPointerValue(dep.Service.Crn)
	}
	return nil
}

func (r *dfDeploymentResource) stateRefresh(ctx context.Context, crn *string, callFailedCount *int, callFailureThreshold int) func() (any, string, error) {
	return func() (any, string, error) {
		params := operations.NewDescribeDeploymentParamsWithContext(ctx).WithInput(&dfmodels.DescribeDeploymentRequest{
			DeploymentCrn: crn,
		})
		resp, err := r.client.Df.Operations.DescribeDeployment(params)
		if err != nil {
			if strings.Contains(err.Error(), "NOT_FOUND") {
				return &dfmodels.DescribeDeploymentResponse{}, "TERMINATED", nil
			}
			*callFailedCount++
			if *callFailedCount <= callFailureThreshold {
				tflog.Warn(ctx, fmt.Sprintf("could not describe DataFlow deployment due to [%s] but threshold not reached (%d/%d).", err.Error(), *callFailedCount, callFailureThreshold))
				return nil, "", nil
			}
			return nil, "", err
		}
		*callFailedCount = 0
		dep := resp.GetPayload().Deployment
		if dep == nil {
			return &dfmodels.DescribeDeploymentResponse{}, "TERMINATED", nil
		}
		state := string(*dep.Status.State)
		tflog.Debug(ctx, fmt.Sprintf("DataFlow deployment %s status: %s", *crn, state))
		return dep, state, nil
	}
}
