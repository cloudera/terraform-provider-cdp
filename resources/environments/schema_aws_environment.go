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
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/resources/environments/validators"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var AwsEnvironmentSchema = schema.Schema{
	MarkdownDescription: "The environment is a logical entity that represents the association of your user account with multiple compute resources using which you can provision and manage workloads.",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			MarkdownDescription: "The id of the environment associated by Terraform",
			Description:         "The id of the environment associated by Terraform",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"crn": schema.StringAttribute{
			MarkdownDescription: "The CRN of the environment.",
			Description:         "The CRN of the environment.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"polling_options": schema.SingleNestedAttribute{
			MarkdownDescription: "Polling related configuration options that could specify various values that will be used during CDP resource creation.",
			Description:         "Polling related configuration options that could specify various values that will be used during CDP resource creation.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"async": schema.BoolAttribute{
					MarkdownDescription: "Boolean value that specifies if Terraform should wait for resource creation/deletion.",
					Description:         "Boolean value that specifies if Terraform should wait for resource creation/deletion.",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
					PlanModifiers: []planmodifier.Bool{
						boolplanmodifier.UseStateForUnknown(),
					},
				},
				"polling_timeout": schema.Int64Attribute{
					MarkdownDescription: "Timeout value in minutes that specifies for how long should the polling go for resource creation/deletion.",
					Description:         "Timeout value in minutes that specifies for how long should the polling go for resource creation/deletion.",
					Default:             int64default.StaticInt64(60),
					Computed:            true,
					Optional:            true,
				},
				"call_failure_threshold": schema.Int64Attribute{
					MarkdownDescription: "Threshold value that specifies how many times should a single call failure happen before giving up the polling.",
					Description:         "Threshold value that specifies how many times should a single call failure happen before giving up the polling.",
					Default:             int64default.StaticInt64(3),
					Computed:            true,
					Optional:            true,
				},
			},
		},
		"authentication": schema.SingleNestedAttribute{
			MarkdownDescription: "Authentication configuration for the environment.",
			Description:         "Authentication configuration for the environment.",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"public_key": schema.StringAttribute{
					MarkdownDescription: "Public SSH key string. The associated private key can be used to get root-level access to the Data Lake instance and Data Hub cluster instances.",
					Description:         "Public SSH key string. The associated private key can be used to get root-level access to the Data Lake instance and Data Hub cluster instances.",
					Optional:            true,
				},
				"public_key_id": schema.StringAttribute{
					MarkdownDescription: "Identifier of the uploaded public SSH key.",
					Description:         "Identifier of the uploaded public SSH key.",
					Optional:            true,
					Computed:            true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
			},
		},
		"compute_cluster": schema.SingleNestedAttribute{
			MarkdownDescription: "Option to set up Externalized compute cluster for the environment.",
			Description:         "Option to set up Externalized compute cluster for the environment.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"enabled": schema.BoolAttribute{
					Required: true,
				},
				"configuration": schema.SingleNestedAttribute{
					MarkdownDescription: "The Externalized k8s configuration for the environment.",
					Description:         "The Externalized k8s configuration for the environment.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"private_cluster": schema.BoolAttribute{
							MarkdownDescription: "If true, creates private cluster. False, if not specified",
							Description:         "If true, creates private cluster. False, if not specified",
							Default:             booldefault.StaticBool(false),
							Computed:            true,
							Optional:            true,
						},
						"kube_api_authorized_ip_ranges": schema.SetAttribute{
							MarkdownDescription: "Kubernetes API authorized IP ranges in CIDR notation. Mutually exclusive with privateCluster.",
							Description:         "Kubernetes API authorized IP ranges in CIDR notation. Mutually exclusive with privateCluster.",
							ElementType:         types.StringType,
							Validators: []validator.Set{
								validators.KubeAPIAuthorizedIPRangesMustBeEmptyWhenPrivateClusterTrue(),
							},
							Optional: true,
						},
						"worker_node_subnets": schema.SetAttribute{
							MarkdownDescription: "Specify subnets for Kubernetes Worker Nodes. If not specified, then the environment's subnet(s) will be used.",
							Description:         "Specify subnets for Kubernetes Worker Nodes. If not specified, then the environment's subnet(s) will be used.",
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},
							ElementType: types.StringType,
							Computed:    true,
							Optional:    true,
						},
					},
				},
			},
		},
		"create_private_subnets": schema.BoolAttribute{
			MarkdownDescription: "When this is enabled, private subnets will be created for the environment.",
			Description:         "When this is enabled, private subnets will be created for the environment.",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"create_service_endpoints": schema.BoolAttribute{
			MarkdownDescription: "Whether or not service endpoints should be created for the environment.",
			Description:         "Whether or not service endpoints should be created for the environment.",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"s3_guard_table_name": schema.StringAttribute{
			MarkdownDescription: "Name of the DynamoDB table to be used for S3Guard.",
			Description:         "Name of the DynamoDB table to be used for S3Guard.",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString(""),
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"credential_name": schema.StringAttribute{
			MarkdownDescription: "Name of the credential to use for the environment.",
			Description:         "Name of the credential to use for the environment.",
			Required:            true,
		},
		"description": schema.StringAttribute{
			MarkdownDescription: "Description of the environment.",
			Description:         "Description of the environment.",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"enable_tunnel": schema.BoolAttribute{
			MarkdownDescription: "Whether to enable SSH tunneling for the environment.",
			Description:         "Whether to enable SSH tunneling for the environment.",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(true),
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"encryption_key_arn": schema.StringAttribute{
			MarkdownDescription: "ARN of the key which will be used to encrypt cloud resources, if entitlement has been granted.",
			Description:         "ARN of the key which will be used to encrypt cloud resources, if entitlement has been granted.",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString(""),
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"endpoint_access_gateway_scheme": schema.StringAttribute{
			MarkdownDescription: "The scheme for the endpoint gateway. PUBLIC creates an external endpoint that can be accessed over the Internet. Defaults to PRIVATE which restricts the traffic to be internal to the VPC.",
			Description:         "The scheme for the endpoint gateway. PUBLIC creates an external endpoint that can be accessed over the Internet. Defaults to PRIVATE which restricts the traffic to be internal to the VPC.",
			Optional:            true,
			Computed:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("PUBLIC", "PRIVATE"),
			},
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"endpoint_access_gateway_subnet_ids": schema.SetAttribute{
			MarkdownDescription: "The subnets to use for endpoint access gateway.",
			Description:         "The subnets to use for endpoint access gateway.",
			Optional:            true,
			ElementType:         types.StringType,
		},
		"environment_name": schema.StringAttribute{
			MarkdownDescription: "The name of the environment. Must contain only lowercase letters, numbers and hyphens.",
			Description:         "The name of the environment. Must contain only lowercase letters, numbers and hyphens.",
			Validators: []validator.String{
				stringvalidator.RegexMatches(
					regexp.MustCompile(`^[a-z0-9-]+$`),
					"must contain only lowercase letters, numbers and hyphens",
				),
			},
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
			Required: true,
		},
		"delete_options": schema.SingleNestedAttribute{
			MarkdownDescription: "Options for deleting the environment.",
			Description:         "Options for deleting the environment.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"cascading": schema.BoolAttribute{
					MarkdownDescription: "If true, all resources in the environment will be deleted.",
					Description:         "If true, all resources in the environment will be deleted.",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"forced": schema.BoolAttribute{
					MarkdownDescription: "Force delete action removes CDP resources and may leave cloud provider resources running even if the deletion did not succeed.",
					Description:         "Force delete action removes CDP resources and may leave cloud provider resources running even if the deletion did not succeed.",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
			},
		},
		"freeipa": FreeIpaSchema,
		"log_storage": schema.SingleNestedAttribute{
			MarkdownDescription: "AWS storage configuration for cluster and audit logs.",
			Description:         "AWS storage configuration for cluster and audit logs.",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"instance_profile": schema.StringAttribute{
					MarkdownDescription: "The instance profile associated with the logger.",
					Description:         "The instance profile associated with the logger.",
					Required:            true,
				},
				"storage_location_base": schema.StringAttribute{
					MarkdownDescription: "The storage location to use for cluster and audit logs.",
					Description:         "The storage location to use for cluster and audit logs.",
					Required:            true,
				},
				"backup_storage_location_base": schema.StringAttribute{
					MarkdownDescription: "The storage location to use for backups.",
					Description:         "The storage location to use for backups.",
					Optional:            true,
					Computed:            true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
			},
		},
		"region": schema.StringAttribute{
			MarkdownDescription: "The region of the environment.",
			Description:         "The region of the environment.",
			Required:            true,
		},
		"proxy_config_name": schema.StringAttribute{
			MarkdownDescription: "Name of the proxy config to use for the environment.",
			Description:         "Name of the proxy config to use for the environment.",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString(""),
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"security_access": securityAccess,
		"status": schema.StringAttribute{
			MarkdownDescription: "The current status of the environment.",
			Description:         "The current status of the environment.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"status_reason": schema.StringAttribute{
			MarkdownDescription: "The detailed status reason of the environment.",
			Description:         "The detailed status reason of the environment.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"subnet_ids": schema.SetAttribute{
			MarkdownDescription: "One or more subnet ids within the VPC.",
			Description:         "One or more subnet ids within the VPC.",
			Required:            true,
			ElementType:         types.StringType,
		},
		"tags": schema.MapAttribute{
			MarkdownDescription: "Tags associated with the resources.",
			Description:         "Tags associated with the resources.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers: []planmodifier.Map{
				mapplanmodifier.UseStateForUnknown(),
			},
		},
		"tunnel_type": schema.StringAttribute{
			// tunnel_type is read only.
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"workload_analytics": schema.BoolAttribute{
			MarkdownDescription: "When this is enabled, diagnostic information about job and query execution is sent to Workload Manager for Data Hub clusters created within this environment.",
			Description:         "When this is enabled, diagnostic information about job and query execution is sent to Workload Manager for Data Hub clusters created within this environment.",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"vpc_id": schema.StringAttribute{
			MarkdownDescription: "The id of the AWS VPC.",
			Description:         "The id of the AWS VPC.",
			Required:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"custom_docker_registry": schema.SingleNestedAttribute{
			Optional:            true,
			MarkdownDescription: "The desired custom docker registry for data services to be used.",
			Description:         "The desired custom docker registry for data services to be used.",
			Attributes: map[string]schema.Attribute{
				"crn": schema.StringAttribute{
					Required:            true,
					MarkdownDescription: "The CRN of the desired custom docker registry for data services to be used.",
					Description:         "The CRN of the desired custom docker registry for data services to be used.",
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
			},
		},
		"security": schema.SingleNestedAttribute{
			Optional:            true,
			MarkdownDescription: "Security related configuration for Data Hub cluster.",
			Description:         "Security related configuration for Data Hub cluster.",
			Attributes: map[string]schema.Attribute{
				"se_linux": schema.StringAttribute{
					Optional:            true,
					Computed:            true,
					MarkdownDescription: "Override default SELinux configuration which is PERMISSIVE by default. Available values: PERMISSIVE, ENFORCING",
					Description:         "Override default SELinux configuration which is PERMISSIVE by default. Available values: PERMISSIVE, ENFORCING",
					Default:             stringdefault.StaticString("PERMISSIVE"),
					Validators: []validator.String{
						stringvalidator.OneOf("PERMISSIVE", "ENFORCING"),
					},
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
			},
		},
		"environment_type": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "Environment type which can be hybrid or public cloud. Available values: PUBLIC_CLOUD, HYBRID",
			Description:         "Environment type which can be hybrid or public cloud. Available values: PUBLIC_CLOUD, HYBRID",
			Validators: []validator.String{
				stringvalidator.OneOf("PUBLIC_CLOUD", "HYBRID"),
			},
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
	},
}

func ToAwsEnvironmentRequest(ctx context.Context, model *awsEnvironmentResourceModel) *environmentsmodels.CreateAWSEnvironmentRequest {
	req := &environmentsmodels.CreateAWSEnvironmentRequest{}
	req.Authentication = &environmentsmodels.AuthenticationRequest{
		PublicKey:   model.Authentication.PublicKey.ValueString(),
		PublicKeyID: model.Authentication.PublicKeyID.ValueString(),
	}
	req.CreatePrivateSubnets = model.CreatePrivateSubnets.ValueBool()
	req.CreateServiceEndpoints = model.CreateServiceEndpoints.ValueBool()
	req.CredentialName = model.CredentialName.ValueStringPointer()
	req.Description = model.Description.ValueString()
	req.EnableTunnel = model.EnableTunnel.ValueBoolPointer()
	req.EncryptionKeyArn = model.EncryptionKeyArn.ValueString()
	req.EndpointAccessGatewayScheme = model.EndpointAccessGatewayScheme.ValueString()
	req.EndpointAccessGatewaySubnetIds = utils.FromSetValueToStringList(model.EndpointAccessGatewaySubnetIds)
	req.EnvironmentName = model.EnvironmentName.ValueStringPointer()
	req.EnvironmentType = model.EnvironmentType.ValueString()

	if model.Security != nil {
		req.Security = &environmentsmodels.SecurityRequest{
			SeLinux: model.Security.Crn.ValueString(),
		}
	}
	if model.CustomDockerRegistry != nil {
		req.CustomDockerRegistry = &environmentsmodels.CustomDockerRegistryRequest{
			Crn: model.CustomDockerRegistry.Crn.ValueStringPointer(),
		}
	}
	if !model.EndpointAccessGatewaySubnetIds.IsNull() && !model.EndpointAccessGatewaySubnetIds.IsUnknown() {
		req.EndpointAccessGatewaySubnetIds = utils.FromSetValueToStringList(model.EndpointAccessGatewaySubnetIds)
	}

	if !model.FreeIpa.IsNull() && !model.FreeIpa.IsUnknown() {
		trans, img := FreeIpaModelToRequest(&model.FreeIpa, ctx)
		req.FreeIpa = &environmentsmodels.AWSFreeIpaCreationRequest{
			InstanceCountByGroup: trans.InstanceCountByGroup,
			InstanceType:         trans.InstanceType,
			MultiAz:              trans.MultiAz,
			Recipes:              trans.Recipes,
			Architecture:         trans.Architecture,
		}
		req.Image = img
	}

	if model.LogStorage != nil {
		req.LogStorage = &environmentsmodels.AwsLogStorageRequest{
			InstanceProfile:           model.LogStorage.InstanceProfile.ValueStringPointer(),
			StorageLocationBase:       model.LogStorage.StorageLocationBase.ValueStringPointer(),
			BackupStorageLocationBase: model.LogStorage.BackupStorageLocationBase.ValueString(),
		}
	}
	req.ProxyConfigName = model.ProxyConfigName.ValueString()
	req.Region = model.Region.ValueStringPointer()
	req.S3GuardTableName = model.S3GuardTableName.ValueString()
	req.SecurityAccess = &environmentsmodels.SecurityAccessRequest{
		Cidr:                    model.SecurityAccess.Cidr.ValueString(),
		DefaultSecurityGroupIDs: utils.FromSetValueToStringList(model.SecurityAccess.DefaultSecurityGroupIDs),
		DefaultSecurityGroupID:  model.SecurityAccess.DefaultSecurityGroupID.ValueString(),
		SecurityGroupIDsForKnox: utils.FromSetValueToStringList(model.SecurityAccess.SecurityGroupIDsForKnox),
		SecurityGroupIDForKnox:  model.SecurityAccess.SecurityGroupIDForKnox.ValueString(),
	}
	if !model.SubnetIds.IsNull() && !model.SubnetIds.IsUnknown() {
		req.SubnetIds = utils.FromSetValueToStringList(model.SubnetIds)
	}
	req.Tags = ConvertTags(ctx, model.Tags)
	req.VpcID = model.VpcID.ValueString()
	req.WorkloadAnalytics = model.WorkloadAnalytics.ValueBool()
	if model.ComputeCluster != nil && model.ComputeCluster.Enabled.ValueBool() {
		var subnets []string
		var ipRanges []string
		var privateCluster bool
		if model.ComputeCluster.Configuration != nil {
			privateCluster = model.ComputeCluster.Configuration.PrivateCluster.ValueBool()
			if !model.ComputeCluster.Configuration.WorkerNodeSubnets.IsNull() && !model.ComputeCluster.Configuration.WorkerNodeSubnets.IsUnknown() {
				subnets = utils.FromSetValueToStringList(model.ComputeCluster.Configuration.WorkerNodeSubnets)
			} else {
				subnets = utils.FromSetValueToStringList(model.SubnetIds)
				model.ComputeCluster.Configuration.WorkerNodeSubnets = model.SubnetIds
			}
			if !model.ComputeCluster.Configuration.KubeApiAuthorizedIpRanges.IsNull() && !model.ComputeCluster.Configuration.KubeApiAuthorizedIpRanges.IsUnknown() {
				ipRanges = utils.FromSetValueToStringList(model.ComputeCluster.Configuration.KubeApiAuthorizedIpRanges)
			} else {
				ipRanges = nil
			}
			model.ComputeCluster.Configuration.KubeApiAuthorizedIpRanges = utils.ToSetValueFromStringList(ipRanges)
		} else {
			subnets = utils.FromSetValueToStringList(model.SubnetIds)
			privateCluster = true
		}
		req.EnableComputeCluster = true
		req.ComputeClusterConfiguration = &environmentsmodels.AWSComputeClusterConfigurationRequest{
			KubeAPIAuthorizedIPRanges: ipRanges,
			PrivateCluster:            privateCluster,
			WorkerNodeSubnets:         subnets,
		}
	}
	utils.LogSilently(ctx, "CreateAWSEnvironmentRequest has been created: ", req)
	return req
}
