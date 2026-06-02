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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/resources/environments/validators"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

const computeClusterOutboundTypeDefaultValue = "udr"

var AzureEnvironmentSchema = schema.Schema{
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
			MarkdownDescription: "When this is enabled, then Azure Postgres will be configured with Private Endpoint and a Private DNS Zone. When this is disabled, then Azure Service Endpoints will be created. The default value is disabled.",
			Description:         "When this is enabled, then Azure Postgres will be configured with Private Endpoint and a Private DNS Zone. When this is disabled, then Azure Service Endpoints will be created. The default value is disabled.",
			Optional:            true,
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
		"enable_outbound_load_balancer": schema.BoolAttribute{
			MarkdownDescription: "Whether or not outbound load balancers should be created for Azure environments. The default behavior is to not create the outbound load balancer.",
			Description:         "Whether or not outbound load balancers should be created for Azure environments. The default behavior is to not create the outbound load balancer.",
			Optional:            true,
		},
		"enable_tunnel": schema.BoolAttribute{
			Description:         "Whether to enable SSH tunneling for the environment.",
			MarkdownDescription: "Whether to enable SSH tunneling for the environment.",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(true),
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"endpoint_access_gateway_scheme": schema.StringAttribute{
			Description:         "The scheme for the endpoint gateway. PUBLIC creates an external endpoint that can be accessed over the Internet. Defaults to PRIVATE which restricts the traffic to be internal to the VPC.",
			MarkdownDescription: "The scheme for the endpoint gateway. PUBLIC creates an external endpoint that can be accessed over the Internet. Defaults to PRIVATE which restricts the traffic to be internal to the VPC.",
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("PUBLIC", "PRIVATE"),
			},
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"endpoint_access_gateway_subnet_ids": schema.SetAttribute{
			Optional:            true,
			ElementType:         types.StringType,
			Description:         "The subnets to use for endpoint access gateway.",
			MarkdownDescription: "The subnets to use for endpoint access gateway.",
		},
		"encryption_key_resource_group_name": schema.StringAttribute{
			MarkdownDescription: "Name of the existing Azure resource group hosting the Azure Key Vault containing customer managed key which will be used to encrypt the Azure Managed Disks. It is required only when the entitlement is granted and the resource group of the key vault is different from the resource group in which the environment is to be created. Omitting it implies that, the key vault containing the encryption key is present in the same resource group where the environment would be created.",
			Description:         "Name of the existing Azure resource group hosting the Azure Key Vault containing customer managed key which will be used to encrypt the Azure Managed Disks. It is required only when the entitlement is granted and the resource group of the key vault is different from the resource group in which the environment is to be created. Omitting it implies that, the key vault containing the encryption key is present in the same resource group where the environment would be created.",
			Optional:            true,
		},
		"encryption_key_url": schema.StringAttribute{
			MarkdownDescription: "URL of the key which will be used to encrypt the Azure Managed Disks, if entitlement has been granted.",
			Description:         "URL of the key which will be used to encrypt the Azure Managed Disks, if entitlement has been granted.",
			Optional:            true,
		},
		"encryption_at_host": schema.BoolAttribute{
			MarkdownDescription: "When this is enabled, we will provision resources with host encrypted true flag.",
			Description:         "When this is enabled, we will provision resources with host encrypted true flag.",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
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
			MarkdownDescription: "Request object for creating an Azure environment using existing VNet and subnets.",
			Description:         "Request object for creating an Azure environment using existing VNet and subnets.",
			Required:            true,
			PlanModifiers: []planmodifier.Object{
				objectplanmodifier.UseStateForUnknown(),
			},
			Attributes: map[string]schema.Attribute{
				"aks_private_dns_zone_id": schema.StringAttribute{
					MarkdownDescription: "The full Azure resource ID of an existing Private DNS zone used for the AKS.",
					Description:         "The full Azure resource ID of an existing Private DNS zone used for the AKS.",
					Optional:            true,
					Computed:            true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
				"database_private_dns_zone_id": schema.StringAttribute{
					MarkdownDescription: "The full Azure resource ID of the existing Private DNS Zone used for Flexible Server and Single Server Databases.",
					Description:         "The full Azure resource ID of the existing Private DNS Zone used for Flexible Server and Single Server Databases.",
					Optional:            true,
					Computed:            true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
				"network_id": schema.StringAttribute{
					MarkdownDescription: "The id of the Azure VNet.",
					Description:         "The id of the Azure VNet.",
					Required:            true,
				},
				"resource_group_name": schema.StringAttribute{
					MarkdownDescription: "The name of the resource group associated with the VNet.",
					Description:         "The name of the resource group associated with the VNet.",
					Required:            true,
				},
				"subnet_ids": schema.SetAttribute{
					MarkdownDescription: "One or more subnet ids within the VNet.",
					Description:         "One or more subnet ids within the VNet.",
					Required:            true,
					ElementType:         types.StringType,
				},
				"flexible_server_subnet_ids": schema.SetAttribute{
					MarkdownDescription: "The subnets delegated for Flexible Server database. Accepts either the name or the full resource id.",
					Description:         "The subnets delegated for Flexible Server database. Accepts either the name or the full resource id.",
					Optional:            true,
					Computed:            true,
					ElementType:         types.StringType,
					PlanModifiers: []planmodifier.Set{
						setplanmodifier.UseStateForUnknown(),
					},
				},
			},
		},
		"freeipa": FreeIpaSchema,
		"log_storage": schema.SingleNestedAttribute{
			MarkdownDescription: "Azure storage configuration for cluster and audit logs.",
			Description:         "Azure storage configuration for cluster and audit logs.",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"managed_identity": schema.StringAttribute{
					MarkdownDescription: "The managed identity associated with the logger. This identity should have Storage Blob Data Contributor role on the given storage account.",
					Description:         "The managed identity associated with the logger. This identity should have Storage Blob Data Contributor role on the given storage account.",
					Required:            true,
				},
				"storage_location_base": schema.StringAttribute{
					MarkdownDescription: "The storage location to use. The location has to be in the following format abfs://filesystem@storage-account-name.dfs.core.windows.net. The filesystem must already exist and the storage account must be StorageV2.",
					Description:         "The storage location to use. The location has to be in the following format abfs://filesystem@storage-account-name.dfs.core.windows.net. The filesystem must already exist and the storage account must be StorageV2.",
					Required:            true,
				},
				"backup_storage_location_base": schema.StringAttribute{
					MarkdownDescription: "The storage location to use. The location has to be in the following format abfs://filesystem@storage-account-name.dfs.core.windows.net. The filesystem must already exist and the storage account must be StorageV2.",
					Description:         "The storage location to use. The location has to be in the following format abfs://filesystem@storage-account-name.dfs.core.windows.net. The filesystem must already exist and the storage account must be StorageV2.",
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
						"outbound_type": schema.StringAttribute{
							MarkdownDescription: "Customize cluster egress with defined outbound type in Azure Kubernetes Service. Possible value(s): udr",
							Description:         "Customize cluster egress with defined outbound type in Azure Kubernetes Service. Possible value(s): udr",
							Optional:            true,
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
		"proxy_config_name": schema.StringAttribute{
			MarkdownDescription: "Name of the proxy config to use for the environment.",
			Description:         "Name of the proxy config to use for the environment.",
			Optional:            true,
		},
		"public_key": schema.StringAttribute{
			MarkdownDescription: "Public SSH key string. The associated private key can be used to get root-level access to the Data Lake instance and Data Hub cluster instances.",
			Description:         "Public SSH key string. The associated private key can be used to get root-level access to the Data Lake instance and Data Hub cluster instances.",
			Required:            true,
		},
		"region": schema.StringAttribute{
			MarkdownDescription: "The region of the environment.",
			Description:         "The region of the environment.",
			Required:            true,
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
			MarkdownDescription: "Name of an existing Azure resource group to be used for the environment. If it is not specified then new resource groups will be generated.",
			Description:         "Name of an existing Azure resource group to be used for the environment. If it is not specified then new resource groups will be generated.",
			Optional:            true,
			Computed:            true,
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
		"encryption_user_managed_identity": schema.StringAttribute{
			MarkdownDescription: "User managed identity for encryption.",
			Description:         "User managed identity for encryption.",
			Optional:            true,
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
		"use_public_ip": schema.BoolAttribute{
			MarkdownDescription: "Whether to associate public ip's to the resources within the network.",
			Description:         "Whether to associate public ip's to the resources within the network.",
			Required:            true,
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
					Default:             stringdefault.StaticString("PERMISSIVE"),
					MarkdownDescription: "Override default SELinux configuration which is PERMISSIVE by default. Available values: PERMISSIVE, ENFORCING",
					Description:         "Override default SELinux configuration which is PERMISSIVE by default. Available values: PERMISSIVE, ENFORCING",
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
		"data_services": schema.SingleNestedAttribute{
			Optional:            true,
			MarkdownDescription: "Azure-specific Data Service parameters request.",
			Description:         "Azure-specific Data Service parameters request.",
			Attributes: map[string]schema.Attribute{
				"shared_managed_identity": schema.StringAttribute{
					Required:            true,
					MarkdownDescription: "User-assigned managed identity used by the AKS control plane.",
					Description:         "User-assigned managed identity used by the AKS control plane.",
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
				"aks_private_dns_zone_id": schema.StringAttribute{
					Optional:            true,
					MarkdownDescription: "The full Azure resource ID of an existing Private DNS zone used for the AKS.",
					Description:         "The full Azure resource ID of an existing Private DNS zone used for the AKS.",
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
			},
		},
		"flexible_server_subnet_ids": schema.SetAttribute{
			Optional:            true,
			ElementType:         types.StringType,
			MarkdownDescription: "The subnets delegated for Flexible Server database. Accepts either the name or the full resource id.",
			Description:         "The subnets delegated for Flexible Server database. Accepts either the name or the full resource id.",
		},
		"availability_zones": schema.SetAttribute{
			Optional:            true,
			ElementType:         types.StringType,
			MarkdownDescription: "The zones of the environment in the given region.",
			Description:         "The zones of the environment in the given region.",
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
	req.EnvironmentType = model.EnvironmentType.ValueString()

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
	if !model.FlexibleServerSubnetIds.IsNull() && !model.FlexibleServerSubnetIds.IsUnknown() {
		req.FlexibleServerSubnetIds = utils.FromSetValueToStringList(model.FlexibleServerSubnetIds)
	}
	if !model.AvailabilityZones.IsNull() && !model.AvailabilityZones.IsUnknown() {
		req.AvailabilityZones = utils.FromSetValueToStringList(model.AvailabilityZones)
	}
	if model.DataServices != nil {
		req.DataServices = &environmentsmodels.DataServicesRequest{Azure: &environmentsmodels.AzureDataServicesParametersRequest{
			AksPrivateDNSZoneID:   model.DataServices.AksPrivateDnsZoneId.ValueString(),
			SharedManagedIdentity: model.DataServices.SharedManagedIdentity.ValueStringPointer(),
		}}
	}

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
				outboundType = computeClusterOutboundTypeDefaultValue
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
