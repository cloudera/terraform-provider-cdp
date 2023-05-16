package environments

import (
	"context"
	"log"
	"sort"
	"strings"
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
	_ resource.ResourceWithValidateConfig = &awsEnvironmentResource{}
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

func (r *awsEnvironmentResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data awsEnvironmentResourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.VpcID.IsNull() || data.VpcID.IsUnknown() {
		if data.NetworkCidr.IsNull() || data.NetworkCidr.IsUnknown() {
			resp.Diagnostics.AddError("Missing Network CIDR", "Network CIDR is required for new networks")
			if data.SecurityAccess == nil || data.SecurityAccess.Cidr.IsNull() || data.SecurityAccess.Cidr.IsUnknown() {
				resp.Diagnostics.AddError("Missing Security Access CIDR", "Security Access CIDR is required for new networks")
			}
		}
	} else {
		if data.SubnetIds.IsNull() || data.SubnetIds.IsUnknown() {
			resp.Diagnostics.AddError("Missing Subnet IDs", "Subnet IDs are required when a VPC ID is specified")
		}
		if data.SecurityAccess != nil {
			defaultSecurityGroupProvided := !data.SecurityAccess.DefaultSecurityGroupID.IsNull() && !data.SecurityAccess.DefaultSecurityGroupID.IsUnknown()
			securityGroupForGatewayNodesProvided := !data.SecurityAccess.SecurityGroupIDForKnox.IsNull() && !data.SecurityAccess.SecurityGroupIDForKnox.IsUnknown()
			if defaultSecurityGroupProvided && !securityGroupForGatewayNodesProvided {
				resp.Diagnostics.AddError("Missing Security Group IDs for Knox", "Security Group IDs for Knox are required when Security Group IDs are provided")
			} else if !defaultSecurityGroupProvided && securityGroupForGatewayNodesProvided {
				resp.Diagnostics.AddError("Missing Security Group IDs", "Security Group IDs for Knox are required when Security Group IDs for Knox are provided")
			} else if !defaultSecurityGroupProvided && !securityGroupForGatewayNodesProvided {
				if data.SecurityAccess.Cidr.IsNull() || data.SecurityAccess.Cidr.IsUnknown() {
					resp.Diagnostics.AddError("Missing Access CIDR", "Access CIDR is required to create new security groups")
				}
			}
		} else {
			resp.Diagnostics.AddError("Missing Security Access Fields", "Either provide an access CIDR of security group ids")
		}
	}
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

	descParams := operations.NewDescribeEnvironmentParamsWithContext(ctx)
	descParams.WithInput(&environmentsmodels.DescribeEnvironmentRequest{EnvironmentName: data.EnvironmentName.ValueStringPointer()})
	_, err := client.Operations.DescribeEnvironment(descParams)
	if err != nil {
		if strings.Contains(err.Error(), "Code:NOT_FOUND") {
			tflog.Debug(ctx, "Environment not found with this name. Proceeding with environment creation.", map[string]interface{}{
				"id": data.EnvironmentName.ValueString(),
			})
		} else {
			resp.Diagnostics.AddError(
				"Error Reading AWS Environment",
				"Could not read AWS Environment: "+data.EnvironmentName.ValueString()+": "+err.Error(),
			)
			return
		}
	} else {
		resp.Diagnostics.AddError(
			"AWS Environment already exists",
			"An environment with the name: "+data.EnvironmentName.ValueString()+" already exists. ",
		)
		return
	}

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
	descParams = operations.NewDescribeEnvironmentParams()
	descParams.WithInput(&environmentsmodels.DescribeEnvironmentRequest{
		EnvironmentName: &environmentName,
	})
	descEnvResp, err := r.client.Environments.Operations.DescribeEnvironment(descParams)
	if err != nil {
		if strings.Contains(err.Error(), "Code:NOT_FOUND") {
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

func isNotFoundError(err error) bool {
	if d, ok := err.(*operations.DescribeEnvironmentDefault); ok && d.GetPayload() != nil {
		return d.GetPayload().Code == "NOT_FOUND"
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
			params := operations.NewDescribeEnvironmentParams()
			params.WithInput(&environmentsmodels.DescribeEnvironmentRequest{EnvironmentName: &environmentName})
			resp, err := client.Operations.DescribeEnvironment(params)
			if err != nil {
				// Envs that have just been created may not be returned from Describe Environment request because of eventual
				// consistency. We return an empty state to retry.

				if isNotFoundError(err) {
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

	environmentName := state.ID.ValueString()
	params := operations.NewDescribeEnvironmentParams()
	params.WithInput(&environmentsmodels.DescribeEnvironmentRequest{
		EnvironmentName: &environmentName,
	})
	descEnvResp, err := r.client.Environments.Operations.DescribeEnvironment(params)
	if err != nil {
		if strings.Contains(err.Error(), "Code:NOT_FOUND") {
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

	model.ID = types.StringPointerValue(env.EnvironmentName)
	if env.Authentication != nil {
		model.Authentication = &Authentication{
			PublicKey:   nilForEmptyString(env.Authentication.PublicKey),
			PublicKeyID: types.StringValue(env.Authentication.PublicKeyID),
		}
	}
	if env.AwsDetails != nil {
		model.S3GuardTableName = types.StringValue(env.AwsDetails.S3GuardTableName)
	}
	model.CloudStorageLogging = types.BoolValue(env.CloudStorageLogging)
	model.CredentialName = types.StringPointerValue(env.CredentialName)
	model.Crn = types.StringPointerValue(env.Crn)
	model.Description = types.StringValue(env.Description)
	model.EnableWorkloadAnalytics = types.BoolValue(env.EnableWorkloadAnalytics)
	model.EnvironmentName = types.StringPointerValue(env.EnvironmentName)
	var freeIpaRecipes []string
	if env.Freeipa != nil {
		freeIpaRecipes = env.Freeipa.Recipes
	}
	model.FreeIpa, _ = types.ObjectValueFrom(ctx, map[string]attr.Type{
		"catalog":                 types.StringType,
		"image_id":                types.StringType,
		"instance_count_by_group": types.Int64Type,
		"instance_type":           types.StringType,
		"multi_az":                types.BoolType,
		"recipes":                 types.ListType{ElemType: types.StringType},
	}, &AWSFreeIpaDetails{
		Recipes: utils.ToListToBaseType(freeIpaRecipes),
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
			sort.Strings(env.Network.EndpointAccessGatewaySubnetIds)
			subnetid, _ := types.ListValueFrom(ctx, types.StringType, env.Network.EndpointAccessGatewaySubnetIds)
			model.EndpointAccessGatewaySubnetIds = subnetid
		}
		if env.Network.Aws != nil {
			model.VpcID = types.StringPointerValue(env.Network.Aws.VpcID)
		}
		if model.SubnetIds.IsNull() {
			sort.Strings(env.Network.SubnetIds)
			for i, v := range env.Network.SubnetIds {
				env.Network.SubnetIds[i] = strings.ReplaceAll(v, "\"", "")
			}
			subnetid, _ := types.ListValueFrom(ctx, types.StringType, env.Network.SubnetIds)
			model.SubnetIds = subnetid
		}
	}
	if env.ProxyConfig != nil {
		model.ProxyConfigName = types.StringPointerValue(env.ProxyConfig.ProxyConfigName)
	}
	model.Region = types.StringPointerValue(env.Region)
	model.ReportDeploymentLogs = types.BoolValue(env.ReportDeploymentLogs)
	if env.SecurityAccess != nil {
		model.SecurityAccess = &SecurityAccess{
			Cidr:                    types.StringValue(env.SecurityAccess.Cidr),
			DefaultSecurityGroupID:  types.StringValue(env.SecurityAccess.DefaultSecurityGroupID),
			DefaultSecurityGroupIDs: types.ListNull(types.StringType),
			SecurityGroupIDForKnox:  types.StringValue(env.SecurityAccess.SecurityGroupIDForKnox),
			SecurityGroupIDsForKnox: types.ListNull(types.StringType),
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

	environmentName := state.ID.ValueString()
	params := operations.NewDeleteEnvironmentParams()
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
			params := operations.NewDescribeEnvironmentParams()
			params.WithInput(&environmentsmodels.DescribeEnvironmentRequest{EnvironmentName: &environmentName})
			resp, err := client.Operations.DescribeEnvironment(params)
			if err != nil {
				log.Printf("Error describing environment: %s", err)
				if envErr, ok := err.(*operations.DescribeEnvironmentDefault); ok {
					if cdp.IsEnvironmentsError(envErr.GetPayload(), "NOT_FOUND", "") {
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
