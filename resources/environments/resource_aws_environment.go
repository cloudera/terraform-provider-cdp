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
	"log"
	"time"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

var (
	_ resource.Resource = &awsEnvironmentResource{}
)

type awsEnvironmentResource struct {
	client *cdp.Client
}

func NewAwsEnvironmentResource() resource.Resource {
	return &awsEnvironmentResource{}
}

func (r *awsEnvironmentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_aws_environment"
}

func (r *awsEnvironmentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = AwsEnvironmentSchema
}

func (r *awsEnvironmentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *awsEnvironmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data awsEnvironmentResourceModel
	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Got Error while trying to set plan")
		return
	}

	client := r.client.Environments

	params := operations.NewCreateAWSEnvironmentParamsWithContext(ctx)
	params.WithInput(ToAwsEnvrionmentRequest(ctx, &data))

	responseOk, err := client.Operations.CreateAWSEnvironment(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating AWS Environment",
			"Got error while creating AWS Environment: "+err.Error(),
		)
		return
	}

	envResp := responseOk.Payload.Environment
	refreshedState := toAwsEnvrionmentResource(ctx, envResp)

	diags = resp.State.Set(ctx, refreshedState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	timeout := time.Hour * 1
	if err := waitForEnvironmentToBeAvailable(refreshedState.ID.ValueString(), timeout, client, ctx); err != nil {
		resp.Diagnostics.AddError(
			"Error creating AWS Environment",
			"Failed to poll creating AWS Environment: "+err.Error(),
		)
		return
	}

	environmentName := data.EnvironmentName.ValueString()
	descParams := operations.NewDescribeEnvironmentParamsWithContext(ctx)
	descParams.WithInput(&environmentsmodels.DescribeEnvironmentRequest{
		EnvironmentName: &environmentName,
	})
	descEnvResp, err := r.client.Environments.Operations.DescribeEnvironment(descParams)
	if err != nil {
		if isEnvNotFoundError(err) {
			resp.Diagnostics.AddWarning("Resource not found on provider", "Environment not found, removing from state.")
			tflog.Warn(ctx, "Environment not found, removing from state", map[string]interface{}{
				"id": data.ID.ValueString(),
			})
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error creating AWS Environment",
			"Could not read AWS Environment: "+data.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	refreshedState = toAwsEnvrionmentResource(ctx, descEnvResp.GetPayload().Environment)
	diags = resp.State.Set(ctx, refreshedState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func isEnvNotFoundError(err error) bool {
	if envErr, ok := err.(*operations.DescribeEnvironmentDefault); ok {
		if cdp.IsEnvironmentsError(envErr.GetPayload(), "NOT_FOUND", "") {
			return true
		}
	}
	return false
}

func waitForEnvironmentToBeAvailable(environmentName string, timeout time.Duration, client *client.Environments, ctx context.Context) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"CREATION_INITIATED",
			"NETWORK_CREATION_IN_PROGRESS",
			"PUBLICKEY_CREATE_IN_PROGRESS",
			"ENVIRONMENT_RESOURCE_ENCRYPTION_INITIALIZATION_IN_PROGRESS",
			"ENVIRONMENT_VALIDATION_IN_PROGRESS",
			"ENVIRONMENT_INITIALIZATION_IN_PROGRESS",
			"FREEIPA_CREATION_IN_PROGRESS"},
		Target:       []string{"AVAILABLE"},
		Delay:        5 * time.Second,
		Timeout:      timeout,
		PollInterval: 10 * time.Second,
		Refresh: func() (interface{}, string, error) {
			log.Printf("[DEBUG] About to describe environment %s", environmentName)
			params := operations.NewDescribeEnvironmentParamsWithContext(ctx)
			params.WithInput(&environmentsmodels.DescribeEnvironmentRequest{EnvironmentName: &environmentName})
			resp, err := client.Operations.DescribeEnvironment(params)
			if err != nil {
				// Envs that have just been created may not be returned from Describe Environment request because of eventual
				// consistency. We return an empty state to retry.

				if isEnvNotFoundError(err) {
					log.Printf("[DEBUG] Recoverable error describing environment: %s", err)
					return nil, "", nil
				}
				log.Printf("Error describing environment: %s", err)
				return nil, "", err
			}
			log.Printf("Described environment: %s", *resp.GetPayload().Environment.Status)
			return resp, *resp.GetPayload().Environment.Status, nil
		},
	}
	_, err := stateConf.WaitForStateContext(ctx)

	return err
}

func (r *awsEnvironmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state awsEnvironmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	environmentName := state.EnvironmentName.ValueString()
	params := operations.NewDescribeEnvironmentParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.DescribeEnvironmentRequest{
		EnvironmentName: &environmentName,
	})
	descEnvResp, err := r.client.Environments.Operations.DescribeEnvironment(params)
	if err != nil {
		if isEnvNotFoundError(err) {
			resp.Diagnostics.AddWarning("Resource not found on provider", "Environment not found, removing from state.")
			tflog.Warn(ctx, "Environment not found, removing from state", map[string]interface{}{
				"id": state.ID.ValueString(),
			})
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error Reading AWS Environment",
			"Could not read AWS Environment: "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	refreshedState := toAwsEnvrionmentResource(ctx, descEnvResp.GetPayload().Environment)

	diags = resp.State.Set(ctx, &refreshedState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func nilForEmptyString(in string) types.String {
	if len(in) == 0 {
		return types.StringNull()
	} else {
		return types.StringValue(in)
	}
}

func toAwsEnvrionmentResource(ctx context.Context, env *environmentsmodels.Environment) *awsEnvironmentResourceModel {
	var model awsEnvironmentResourceModel

	model.ID = types.StringPointerValue(env.Crn)
	if env.Authentication != nil {
		model.Authentication = &Authentication{
			PublicKey:   nilForEmptyString(env.Authentication.PublicKey),
			PublicKeyID: types.StringValue(env.Authentication.PublicKeyID),
		}
	}
	if env.AwsDetails != nil {
		model.S3GuardTableName = types.StringValue(env.AwsDetails.S3GuardTableName)
	}
	model.CredentialName = types.StringPointerValue(env.CredentialName)
	model.Crn = types.StringPointerValue(env.Crn)
	model.Description = types.StringValue(env.Description)
	model.EnvironmentName = types.StringPointerValue(env.EnvironmentName)
	var freeIpaRecipes types.Set
	if env.Freeipa != nil {
		freeIpaRecipes, _ = types.SetValueFrom(ctx, types.StringType, env.Freeipa.Recipes)
	} else {
		freeIpaRecipes = types.SetNull(types.StringType)
	}
	model.FreeIpa, _ = types.ObjectValueFrom(ctx, map[string]attr.Type{
		"catalog":                 types.StringType,
		"image_id":                types.StringType,
		"instance_count_by_group": types.Int64Type,
		"instance_type":           types.StringType,
		"multi_az":                types.BoolType,
		"recipes":                 types.SetType{ElemType: types.StringType},
	}, &AWSFreeIpaDetails{
		Recipes: freeIpaRecipes,
	})
	if env.LogStorage != nil {
		if env.LogStorage.AwsDetails != nil {
			model.LogStorage = &AWSLogStorage{
				InstanceProfile:     types.StringValue(env.LogStorage.AwsDetails.InstanceProfile),
				StorageLocationBase: types.StringValue(env.LogStorage.AwsDetails.StorageLocationBase),
			}
			if env.BackupStorage != nil {
				if env.BackupStorage.AwsDetails != nil {
					model.LogStorage.BackupStorageLocationBase = types.StringValue(env.BackupStorage.AwsDetails.StorageLocationBase)
				}

			}
		}
	}
	if env.Network != nil {
		model.NetworkCidr = types.StringValue(env.Network.NetworkCidr)
		if env.Network.EndpointAccessGatewaySubnetIds != nil {
			var eagSubnetids types.Set
			if len(env.Network.EndpointAccessGatewaySubnetIds) > 0 {
				eagSubnetids, _ = types.SetValueFrom(ctx, types.StringType, env.Network.EndpointAccessGatewaySubnetIds)
			} else {
				eagSubnetids = types.SetNull(types.StringType)
			}
			model.EndpointAccessGatewaySubnetIds = eagSubnetids
		}
		if env.Network.Aws != nil {
			model.VpcID = types.StringPointerValue(env.Network.Aws.VpcID)
		}
		var subnetids types.Set
		if len(env.Network.SubnetIds) > 0 {
			subnetids, _ = types.SetValueFrom(ctx, types.StringType, env.Network.SubnetIds)
		} else {
			subnetids = types.SetNull(types.StringType)
		}
		model.SubnetIds = subnetids

	}
	if env.ProxyConfig != nil {
		model.ProxyConfigName = types.StringPointerValue(env.ProxyConfig.ProxyConfigName)
	}
	model.Region = types.StringPointerValue(env.Region)
	model.ReportDeploymentLogs = types.BoolValue(env.ReportDeploymentLogs)
	if env.SecurityAccess != nil {
		var dsgIDs types.Set
		if model.SecurityAccess == nil || model.SecurityAccess.DefaultSecurityGroupIDs.IsUnknown() {
			dsgIDs = types.SetNull(types.StringType)
		} else {
			dsgIDs = model.SecurityAccess.DefaultSecurityGroupIDs
		}
		var sgIDsknox types.Set
		if model.SecurityAccess == nil || model.SecurityAccess.SecurityGroupIDsForKnox.IsUnknown() {
			sgIDsknox = types.SetNull(types.StringType)
		} else {
			sgIDsknox = model.SecurityAccess.DefaultSecurityGroupIDs
		}
		model.SecurityAccess = &SecurityAccess{
			Cidr:                    types.StringValue(env.SecurityAccess.Cidr),
			DefaultSecurityGroupID:  types.StringValue(env.SecurityAccess.DefaultSecurityGroupID),
			DefaultSecurityGroupIDs: dsgIDs,
			SecurityGroupIDForKnox:  types.StringValue(env.SecurityAccess.SecurityGroupIDForKnox),
			SecurityGroupIDsForKnox: sgIDsknox,
		}
	}
	model.Status = types.StringPointerValue(env.Status)
	model.StatusReason = types.StringValue(env.StatusReason)
	if env.Tags != nil {
		merged := env.Tags.Defaults
		for k, v := range env.Tags.UserDefined {
			merged[k] = v
		}
		tagMap, _ := types.MapValueFrom(ctx, types.StringType, merged)
		model.Tags = tagMap
	}
	model.EnableTunnel = types.BoolValue(env.TunnelEnabled)
	model.TunnelType = types.StringValue(string(env.TunnelType))
	model.WorkloadAnalytics = types.BoolValue(env.WorkloadAnalytics)

	return &model
}

func (r *awsEnvironmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

}

func (r *awsEnvironmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state awsEnvironmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	environmentName := state.EnvironmentName.ValueString()
	params := operations.NewDeleteEnvironmentParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.DeleteEnvironmentRequest{EnvironmentName: &environmentName})
	_, err := r.client.Environments.Operations.DeleteEnvironment(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting AWS Environment",
			"Could not delete AWS Environment, unexpected error: "+err.Error(),
		)
		return
	}

	timeout := time.Hour * 1
	err = waitForEnvironmentToBeDeleted(environmentName, timeout, r.client.Environments, ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting AWS Environment",
			"Failed to poll delete AWS Environment, unexpected error: "+err.Error(),
		)
		return
	}
}

func waitForEnvironmentToBeDeleted(environmentName string, timeout time.Duration, client *client.Environments, ctx context.Context) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"STORAGE_CONSUMPTION_COLLECTION_UNSCHEDULING_IN_PROGRESS",
			"NETWORK_DELETE_IN_PROGRESS",
			"FREEIPA_DELETE_IN_PROGRESS",
			"RDBMS_DELETE_IN_PROGRESS",
			"IDBROKER_MAPPINGS_DELETE_IN_PROGRESS",
			"S3GUARD_TABLE_DELETE_IN_PROGRESS",
			"CLUSTER_DEFINITION_DELETE_PROGRESS",
			"CLUSTER_DEFINITION_CLEANUP_PROGRESS",
			"UMS_RESOURCE_DELETE_IN_PROGRESS",
			"DELETE_INITIATED",
			"DATAHUB_CLUSTERS_DELETE_IN_PROGRESS",
			"DATALAKE_CLUSTERS_DELETE_IN_PROGRESS",
			"PUBLICKEY_DELETE_IN_PROGRESS",
			"EVENT_CLEANUP_IN_PROGRESS",
			"EXPERIENCE_DELETE_IN_PROGRESS",
			"ENVIRONMENT_RESOURCE_ENCRYPTION_DELETE_IN_PROGRESS",
			"ENVIRONMENT_ENCRYPTION_RESOURCES_DELETED"},
		Target:       []string{},
		Delay:        5 * time.Second,
		Timeout:      timeout,
		PollInterval: 10 * time.Second,
		Refresh: func() (interface{}, string, error) {
			log.Printf("About to describe environment")
			params := operations.NewDescribeEnvironmentParamsWithContext(ctx)
			params.WithInput(&environmentsmodels.DescribeEnvironmentRequest{EnvironmentName: &environmentName})
			resp, err := client.Operations.DescribeEnvironment(params)
			if err != nil {
				log.Printf("Error describing environment: %s", err)
				if envErr, ok := err.(*operations.DescribeEnvironmentDefault); ok {
					if isEnvNotFoundError(envErr) {
						return nil, "", nil
					}
				}
				return nil, "", err
			}
			if resp.GetPayload().Environment == nil {
				log.Printf("Described environment. No environment.")
				return nil, "", nil
			}
			log.Printf("Described environment: %s", *resp.GetPayload().Environment.Status)
			return resp, *resp.GetPayload().Environment.Status, nil
		},
	}
	_, err := stateConf.WaitForStateContext(ctx)

	return err
}
