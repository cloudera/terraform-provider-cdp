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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

const compute_cluster_outbound_type_default_value = "udr"

var AzureEnvironmentSchema = schema.Schema{
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
		"create_private_endpoints": schema.BoolAttribute{
			Optional: true,
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
		"enable_outbound_load_balancer": schema.BoolAttribute{
			Optional: true,
		},
		"enable_tunnel": schema.BoolAttribute{
			Optional: true,
			Computed: true,
			Default:  booldefault.StaticBool(true),
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"endpoint_access_gateway_scheme": schema.StringAttribute{
			Description:         "The scheme for the endpoint gateway. PUBLIC creates an external endpoint that can be accessed over the Internet. Defaults to PRIVATE which restricts the traffic to be internal to the VPC.",
			MarkdownDescription: "The scheme for the endpoint gateway. PUBLIC creates an external endpoint that can be accessed over the Internet. Defaults to PRIVATE which restricts the traffic to be internal to the VPC.",
			Optional:            true,
		},
		"endpoint_access_gateway_subnet_ids": schema.SetAttribute{
			Optional:    true,
			ElementType: types.StringType,
			Description: "The subnets to use for endpoint access gateway.",
		},
		"encryption_key_resource_group_name": schema.StringAttribute{
			Optional: true,
		},
		"encryption_key_url": schema.StringAttribute{
			Optional: true,
		},
		"encryption_at_host": schema.BoolAttribute{
			Optional: true,
			Computed: true,
			Default:  booldefault.StaticBool(false),
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
		"existing_network_params": schema.SingleNestedAttribute{
			Required: true,
			PlanModifiers: []planmodifier.Object{
				objectplanmodifier.UseStateForUnknown(),
			},
			Attributes: map[string]schema.Attribute{
				"aks_private_dns_zone_id": schema.StringAttribute{
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
				"database_private_dns_zone_id": schema.StringAttribute{
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
				"network_id": schema.StringAttribute{
					Required: true,
				},
				"resource_group_name": schema.StringAttribute{
					Required: true,
				},
				"subnet_ids": schema.SetAttribute{
					Required:    true,
					ElementType: types.StringType,
				},
				"flexible_server_subnet_ids": schema.SetAttribute{
					Optional:    true,
					Computed:    true,
					ElementType: types.StringType,
					PlanModifiers: []planmodifier.Set{
						setplanmodifier.UseStateForUnknown(),
					},
				},
			},
		},
		"freeipa": FreeIpaSchema,
		"log_storage": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"managed_identity": schema.StringAttribute{
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
						"outbound_type": schema.StringAttribute{
							MarkdownDescription: "Customize cluster egress with defined outbound type in Azure Kubernetes Service. Possible value(s): udr",
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
		"proxy_config_name": schema.StringAttribute{
			Optional: true,
		},
		"public_key": schema.StringAttribute{
			Required: true,
		},
		"region": schema.StringAttribute{
			Required: true,
		},
		"report_deployment_logs": schema.BoolAttribute{
			// report_deployment_logs is a deprecated field and should not be used
			MarkdownDescription: " [Deprecated] When true, this will report additional diagnostic information back to Cloudera.",
			Computed:            true,
			Default:             booldefault.StaticBool(false),
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"resource_group_name": schema.StringAttribute{
			Optional: true,
			Computed: true,
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
		"encryption_user_managed_identity": schema.StringAttribute{
			Optional: true,
		},
		"tags": schema.MapAttribute{
			Optional:    true,
			Computed:    true,
			ElementType: types.StringType,
			PlanModifiers: []planmodifier.Map{
				mapplanmodifier.UseStateForUnknown(),
			},
		},
		"use_public_ip": schema.BoolAttribute{
			Required: true,
		},
		"workload_analytics": schema.BoolAttribute{
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
	},
}

func ToAzureEnvironmentRequest(ctx context.Context, model *azureEnvironmentResourceModel) *environmentsmodels.CreateAzureEnvironmentRequest {
	req := &environmentsmodels.CreateAzureEnvironmentRequest{}
	req.CreatePrivateEndpoints = model.CreatePrivateEndpoints.ValueBool()
	req.CredentialName = model.CredentialName.ValueStringPointer()
	req.Description = model.Description.ValueString()
	req.EnableOutboundLoadBalancer = model.EnableOutboundLoadBalancer.ValueBool()
	req.EnableTunnel = model.EnableTunnel.ValueBoolPointer()
	req.EncryptionKeyResourceGroupName = model.EncryptionKeyResourceGroupName.ValueString()
	req.EncryptionKeyURL = model.EncryptionKeyURL.ValueString()
	req.EnvironmentName = model.EnvironmentName.ValueStringPointer()
	req.EndpointAccessGatewayScheme = model.EndpointAccessGatewayScheme.ValueString()
	req.EndpointAccessGatewaySubnetIds = utils.FromSetValueToStringList(model.EndpointAccessGatewaySubnetIds)
	req.EncryptionAtHost = model.EncryptionAtHost.ValueBool()
	req.UserManagedIdentity = model.EncryptionUserManagedIdentity.ValueString()
	var existingNetworkParams existingAzureNetwork
	diag := model.ExistingNetworkParams.As(ctx, &existingNetworkParams, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
	if diag.HasError() {
		for _, v := range diag.Errors() {
			tflog.Debug(ctx, "ERROR: "+v.Detail())
		}
	}
	tflog.Debug(ctx, "network id: "+existingNetworkParams.NetworkID.ValueString())
	tflog.Debug(ctx, "network id cast: "+model.ExistingNetworkParams.Attributes()["network_id"].(types.String).ValueString())
	req.ExistingNetworkParams = &environmentsmodels.ExistingAzureNetworkRequest{
		AksPrivateDNSZoneID:      existingNetworkParams.AksPrivateDNSZoneID.ValueString(),
		DatabasePrivateDNSZoneID: existingNetworkParams.DatabasePrivateDNSZoneID.ValueString(),
		NetworkID:                existingNetworkParams.NetworkID.ValueStringPointer(),
		ResourceGroupName:        existingNetworkParams.ResourceGroupName.ValueStringPointer(),
		SubnetIds:                utils.FromSetValueToStringList(existingNetworkParams.SubnetIds),
	}
	req.FlexibleServerSubnetIds = utils.FromSetValueToStringList(existingNetworkParams.FlexibleServerSubnetIds)

	if !model.FreeIpa.IsNull() && !model.FreeIpa.IsUnknown() {
		trans, img := FreeIpaModelToRequest(&model.FreeIpa, ctx)
		req.FreeIpa = &environmentsmodels.AzureFreeIpaCreationRequest{
			InstanceCountByGroup: trans.InstanceCountByGroup,
			InstanceType:         trans.InstanceType,
			MultiAz:              &trans.MultiAz,
			Recipes:              trans.Recipes,
		}
		req.Image = img
	}

	if model.LogStorage != nil {
		req.LogStorage = &environmentsmodels.AzureLogStorageRequest{
			ManagedIdentity:           model.LogStorage.ManagedIdentity.ValueStringPointer(),
			StorageLocationBase:       model.LogStorage.StorageLocationBase.ValueStringPointer(),
			BackupStorageLocationBase: model.LogStorage.BackupStorageLocationBase.ValueString(),
		}
	}
	req.ProxyConfigName = model.ProxyConfigName.ValueString()
	req.PublicKey = model.PublicKey.ValueStringPointer()
	req.Region = model.Region.ValueStringPointer()
	req.ResourceGroupName = model.ResourceGroupName.ValueString()
	req.SecurityAccess = &environmentsmodels.SecurityAccessRequest{
		Cidr:                    model.SecurityAccess.Cidr.ValueString(),
		DefaultSecurityGroupIDs: utils.FromSetValueToStringList(model.SecurityAccess.DefaultSecurityGroupIDs),
		DefaultSecurityGroupID:  model.SecurityAccess.DefaultSecurityGroupID.ValueString(),
		SecurityGroupIDsForKnox: utils.FromSetValueToStringList(model.SecurityAccess.SecurityGroupIDsForKnox),
		SecurityGroupIDForKnox:  model.SecurityAccess.SecurityGroupIDForKnox.ValueString(),
	}
	req.Tags = ConvertTags(ctx, model.Tags)
	req.UsePublicIP = model.UsePublicIP.ValueBoolPointer()
	req.WorkloadAnalytics = model.WorkloadAnalytics.ValueBool()
	if model.ComputeCluster != nil && model.ComputeCluster.Enabled.ValueBool() {
		var subnets []string
		var ipRanges []string
		var outboundType string
		var privateCluster bool
		if model.ComputeCluster.Configuration != nil {
			privateCluster = model.ComputeCluster.Configuration.PrivateCluster.ValueBool()
			if !model.ComputeCluster.Configuration.WorkerNodeSubnets.IsNull() && !model.ComputeCluster.Configuration.WorkerNodeSubnets.IsUnknown() {
				subnets = utils.FromSetValueToStringList(model.ComputeCluster.Configuration.WorkerNodeSubnets)
			} else {
				subnets = utils.FromSetValueToStringList(existingNetworkParams.SubnetIds)
				model.ComputeCluster.Configuration.WorkerNodeSubnets = existingNetworkParams.SubnetIds
			}
			if !model.ComputeCluster.Configuration.KubeApiAuthorizedIpRanges.IsNull() {
				ipRanges = utils.FromSetValueToStringList(model.ComputeCluster.Configuration.KubeApiAuthorizedIpRanges)
			} else {
				ipRanges = nil
			}
			model.ComputeCluster.Configuration.KubeApiAuthorizedIpRanges = utils.ToSetValueFromStringList(ipRanges)
			if !model.ComputeCluster.Configuration.OutboundType.IsNull() {
				outboundType = model.ComputeCluster.Configuration.OutboundType.ValueString()
			} else {
				outboundType = compute_cluster_outbound_type_default_value
			}
		} else {
			subnets = utils.FromSetValueToStringList(existingNetworkParams.SubnetIds)
			privateCluster = true
		}
		req.EnableComputeCluster = true
		req.ComputeClusterConfiguration = &environmentsmodels.AzureComputeClusterConfigurationRequest{
			KubeAPIAuthorizedIPRanges: ipRanges,
			PrivateCluster:            privateCluster,
			WorkerNodeSubnets:         subnets,
			OutboundType:              outboundType,
		}
	}
	return req
}
