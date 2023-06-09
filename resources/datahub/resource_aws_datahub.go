// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datahub

import (
	"context"
	"fmt"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client/operations"
	datahubmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
	envOperations "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	envmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"time"
)

var (
	_ resource.Resource = &awsDatahubResource{}
)

const pollingInterval = 10 * time.Second
const pollingTimeout = 1 * time.Hour
const pollingDelay = 5 * time.Second

type awsDatahubResource struct {
	client *cdp.Client
}

func NewAwsDatahubResource() resource.Resource {
	return &awsDatahubResource{}
}

func (r *awsDatahubResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_datahub_aws_cluster"
}

func (r *awsDatahubResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *awsDatahubResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "AWS cluster creation process requested.")
	var data awsDatahubResourceModel
	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewCreateAWSClusterParamsWithContext(ctx)
	params.WithInput(fromModelToRequest(data, ctx))

	res, err := r.client.Datahub.Operations.CreateAWSCluster(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating AWS Data hub cluster.",
			"Got error while creating AWS Data hub cluster: "+err.Error(),
		)
		return
	}

	data.Crn = types.StringPointerValue(res.Payload.Cluster.Crn)
	data.ID = types.StringPointerValue(res.Payload.Cluster.Crn)
	data.Name = types.StringPointerValue(res.Payload.Cluster.ClusterName)

	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := waitForToBeAvailable(data.ID.ValueString(), r.client.Datahub, ctx); err != nil {
		resp.Diagnostics.AddError(
			"Error creating AWS Data hub cluster",
			"Failure to poll of AWS Data hub cluster creation: "+err.Error(),
		)
		return
	}
}

func (r *awsDatahubResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state awsDatahubResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewDescribeClusterParamsWithContext(ctx)
	params.WithInput(&datahubmodels.DescribeClusterRequest{
		ClusterName: state.Name.ValueStringPointer(),
	})

	result, err := r.client.Datahub.Operations.DescribeCluster(params)
	if err != nil {
		if isNotFoundError(err) {
			resp.Diagnostics.AddWarning("Resource not found on provider", "AWS Data hub cluster not found, removing from state.")
			tflog.Warn(ctx, "AWS Data hub cluster not found, removing from state", map[string]interface{}{"id": state.ID.ValueString()})
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error Reading AWS Data hub cluster",
			"Could not read AWS Data hub cluster: "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	env, err := fetchEnvByName(r, state.Environment.ValueStringPointer(), ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading environment for AWS Data hub cluster",
			"Could not read AWS environment for Data hub cluster: "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	cluster := result.Payload.Cluster

	state.ID = types.StringPointerValue(cluster.ClusterName)
	state.Environment = types.StringPointerValue(env.Payload.Environment.EnvironmentName)
	state.Crn = types.StringPointerValue(cluster.Crn)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *awsDatahubResource) Update(ctx context.Context, _ resource.UpdateRequest, _ *resource.UpdateResponse) {
	tflog.Info(ctx, "Update operation is not implemented yet.")
}

func (r *awsDatahubResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state awsDatahubResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewDeleteClusterParamsWithContext(ctx).WithInput(&datahubmodels.DeleteClusterRequest{
		ClusterName: state.ID.ValueStringPointer(),
	})
	_, err := r.client.Datahub.Operations.DeleteCluster(params)
	if err != nil {
		if !isNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Error Deleting AWS Data hub cluster",
				"Could not delete AWS Data hub cluster due to: "+err.Error(),
			)
		}
		return
	}

	err = waitForToBeDeleted(state.Name.ValueString(), r.client.Datahub, ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting AWS Data hub cluster",
			"Failure to poll AWS Data hub deletion, unexpected error: "+err.Error(),
		)
		return
	}
}

func waitForToBeAvailable(datahubName string, client *client.Datahub, ctx context.Context) error {
	tflog.Info(ctx, fmt.Sprintf("About to poll cluster (name: %s) creation (polling [delay: %s, timeout: %s, interval :%s]).",
		datahubName, pollingDelay, pollingTimeout, pollingInterval))
	stateConf := &retry.StateChangeConf{
		Pending:                   []string{},
		Target:                    []string{"AVAILABLE"},
		Delay:                     pollingDelay,
		Timeout:                   pollingTimeout,
		PollInterval:              pollingInterval,
		ContinuousTargetOccurence: 2,
		Refresh: func() (interface{}, string, error) {
			tflog.Debug(ctx, fmt.Sprintf("About to describe cluster %s", datahubName))
			params := operations.NewDescribeClusterParamsWithContext(ctx)
			params.WithInput(&datahubmodels.DescribeClusterRequest{ClusterName: &datahubName})
			resp, err := client.Operations.DescribeCluster(params)
			if err != nil {
				if isNotFoundError(err) {
					tflog.Debug(ctx, fmt.Sprintf("Recoverable error describing cluster: %s", err))
					return nil, "", nil
				}
				tflog.Debug(ctx, fmt.Sprintf("Error describing cluster: %s", err))
				return nil, "", err
			}
			tflog.Debug(ctx, fmt.Sprintf("Described cluster: %s", resp.GetPayload().Cluster.Status))
			return checkIfClusterCreationFailed(resp)
		},
	}
	_, err := stateConf.WaitForStateContext(ctx)
	return err
}

func isNotFoundError(err error) bool {
	if d, ok := err.(*operations.DescribeClusterDefault); ok && d.GetPayload() != nil {
		return d.GetPayload().Code == "NOT_FOUND"
	}
	return false
}

func waitForToBeDeleted(datahubName string, client *client.Datahub, ctx context.Context) error {
	stateConf := &retry.StateChangeConf{
		Target:       []string{},
		Delay:        pollingDelay,
		Timeout:      pollingTimeout,
		PollInterval: pollingInterval,
		Refresh: func() (interface{}, string, error) {
			tflog.Debug(ctx, fmt.Sprintf("About to describe cluster %s", datahubName))
			params := operations.NewDescribeClusterParamsWithContext(ctx)
			params.WithInput(&datahubmodels.DescribeClusterRequest{ClusterName: &datahubName})
			resp, err := client.Operations.DescribeCluster(params)
			if err != nil {
				tflog.Debug(ctx, fmt.Sprintf("Error describing cluster: %s", err))
				if envErr, ok := err.(*operations.DescribeClusterDefault); ok {
					if cdp.IsDatahubError(envErr.GetPayload(), "NOT_FOUND", "") {
						return nil, "", nil
					}
				}
				return nil, "", err
			}
			if resp.GetPayload().Cluster == nil {
				tflog.Debug(ctx, "Datahub described. No cluster.")
				return nil, "", nil
			}
			tflog.Debug(ctx, fmt.Sprintf("Described cluster: %s", resp.GetPayload().Cluster.Status))
			return resp, resp.GetPayload().Cluster.Status, nil
		},
	}
	_, err := stateConf.WaitForStateContext(ctx)
	return err
}

func checkIfClusterCreationFailed(resp *operations.DescribeClusterOK) (interface{}, string, error) {
	if utils.ContainsAsSubstring([]string{"FAILED", "DELETED"}, resp.GetPayload().Cluster.Status) {
		return nil, "", fmt.Errorf("cluster status became unacceptable: %s", resp.GetPayload().Cluster.Status)
	}
	return resp, resp.GetPayload().Cluster.Status, nil
}

func fetchEnvByName(r *awsDatahubResource, envName *string, ctx context.Context) (*envOperations.DescribeEnvironmentOK, error) {
	tflog.Debug(ctx, fmt.Sprintf("About to fetch environment based on its name: %s", *envName))
	return r.client.Environments.Operations.DescribeEnvironment(envOperations.NewDescribeEnvironmentParamsWithContext(ctx).WithInput(&envmodels.DescribeEnvironmentRequest{EnvironmentName: envName}))
}
