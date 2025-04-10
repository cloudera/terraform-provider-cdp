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

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var AwsEnvironmentSchema = schema.Schema{
	MarkdownDescription: "The environment is a logical entity that represents the association of your user account with multiple compute resources using which you can provision and manage workloads.",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"crn": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"polling_options": schema.SingleNestedAttribute{
			MarkdownDescription: "Polling related configuration options that could specify various values that will be used during CDP resource creation.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"async": schema.BoolAttribute{
					MarkdownDescription: "Boolean value that specifies if Terraform should wait for resource creation/deletion.",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
					PlanModifiers: []planmodifier.Bool{
						boolplanmodifier.UseStateForUnknown(),
					},
				},
				"polling_timeout": schema.Int64Attribute{
					MarkdownDescription: "Timeout value in minutes that specifies for how long should the polling go for resource creation/deletion.",
					Default:             int64default.StaticInt64(60),
					Computed:            true,
					Optional:            true,
				},
				"call_failure_threshold": schema.Int64Attribute{
					MarkdownDescription: "Threshold value that specifies how many times should a single call failure happen before giving up the polling.",
					Default:             int64default.StaticInt64(3),
					Computed:            true,
					Optional:            true,
				},
			},
		},
		"authentication": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"public_key": schema.StringAttribute{
					Optional: true,
				},
				"public_key_id": schema.StringAttribute{
					Optional: true,
				},
			},
		},
		"compute_cluster": schema.SingleNestedAttribute{
			MarkdownDescription: "Option to set up Externalized compute cluster for the environment.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"enabled": schema.BoolAttribute{
					Required: true,
				},
				"configuration": schema.SingleNestedAttribute{
					MarkdownDescription: "The Externalized k8s configuration for the environment.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"private_cluster": schema.BoolAttribute{
							MarkdownDescription: "If true, creates private cluster. False, if not specified",
							Default:             booldefault.StaticBool(false),
							Computed:            true,
							Optional:            true,
						},
						"kube_api_authorized_ip_ranges": schema.SetAttribute{
							MarkdownDescription: "Kubernetes API authorized IP ranges in CIDR notation. Mutually exclusive with privateCluster.",
							ElementType:         types.StringType,
							Optional:            true,
						},
						"worker_node_subnets": schema.SetAttribute{
							MarkdownDescription: "Specify subnets for Kubernetes Worker Nodes. If not specified, then the environment's subnet(s) will be used.",
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
			Optional: true,
			Computed: true,
			Default:  booldefault.StaticBool(false),
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"create_service_endpoints": schema.BoolAttribute{
			Optional: true,
			Computed: true,
			Default:  booldefault.StaticBool(false),
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"s3_guard_table_name": schema.StringAttribute{
			Optional: true,
			Computed: true,
			Default:  stringdefault.StaticString(""),
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"credential_name": schema.StringAttribute{
			Required: true,
		},
		"description": schema.StringAttribute{
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"enable_tunnel": schema.BoolAttribute{
			Optional: true,
			Computed: true,
			Default:  booldefault.StaticBool(true),
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"encryption_key_arn": schema.StringAttribute{
			Optional: true,
			Computed: true,
			Default:  stringdefault.StaticString(""),
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"endpoint_access_gateway_scheme": schema.StringAttribute{
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"endpoint_access_gateway_subnet_ids": schema.SetAttribute{
			Optional:    true,
			ElementType: types.StringType,
		},
		"environment_name": schema.StringAttribute{
			Required: true,
		},
		"cascading_delete": schema.BoolAttribute{
			Optional:           true,
			Computed:           true,
			Default:            booldefault.StaticBool(true),
			DeprecationMessage: "Cascading_delete is deprecated. Use delete_options.cascading instead. If latter specified, it will override this value.",
		},
		"delete_options": schema.SingleNestedAttribute{
			MarkdownDescription: "Options for deleting the environment.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"cascading": schema.BoolAttribute{
					MarkdownDescription: "If true, all resources in the environment will be deleted.",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"forced": schema.BoolAttribute{
					MarkdownDescription: "Force delete action removes CDP resources and may leave cloud provider resources running even if the deletion did not succeed.",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
			},
		},
		"freeipa": FreeIpaSchema,
		"log_storage": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"instance_profile": schema.StringAttribute{
					Required: true,
				},
				"storage_location_base": schema.StringAttribute{
					Required: true,
				},
				"backup_storage_location_base": schema.StringAttribute{
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
			},
		},
		"region": schema.StringAttribute{
			Required: true,
		},
		"report_deployment_logs": schema.BoolAttribute{
			// report_deployment_logs is a deprecated field and should not be used
			MarkdownDescription: " [Deprecated] When true, this will report additional diagnostic information back to Cloudera.",
			DeprecationMessage:  "report_deployment_logs is a deprecated field and should not be used. ",
			Computed:            true,
			Default:             booldefault.StaticBool(false),
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"proxy_config_name": schema.StringAttribute{
			Optional: true,
			Computed: true,
			Default:  stringdefault.StaticString(""),
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"security_access": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"cidr": schema.StringAttribute{
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
				"default_security_group_id": schema.StringAttribute{
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
				"default_security_group_ids": schema.SetAttribute{
					Optional:    true,
					ElementType: types.StringType,
				},
				"security_group_id_for_knox": schema.StringAttribute{
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
				"security_group_ids_for_knox": schema.SetAttribute{
					Optional:    true,
					ElementType: types.StringType,
				},
			},
		},
		"status": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"status_reason": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"subnet_ids": schema.SetAttribute{
			Required:    true,
			ElementType: types.StringType,
		},
		"tags": schema.MapAttribute{
			Optional:    true,
			Computed:    true,
			ElementType: types.StringType,
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
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"vpc_id": schema.StringAttribute{
			Required: true,
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

	if !model.FreeIpa.IsNull() && !model.FreeIpa.IsUnknown() {
		trans, img := FreeIpaModelToRequest(&model.FreeIpa, ctx)
		req.FreeIpa = &environmentsmodels.AWSFreeIpaCreationRequest{
			InstanceCountByGroup: trans.InstanceCountByGroup,
			InstanceType:         trans.InstanceType,
			MultiAz:              trans.MultiAz,
			Recipes:              trans.Recipes,
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
