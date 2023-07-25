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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

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
		},
		"encryption_key_resource_group_name": schema.StringAttribute{
			Optional: true,
		},
		"encryption_key_url": schema.StringAttribute{
			Optional: true,
		},
		"environment_name": schema.StringAttribute{
			Required: true,
		},
		"existing_network_params": schema.SingleNestedAttribute{
			Optional: true,
			Computed: true,
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
			},
		},
		"freeipa": schema.SingleNestedAttribute{
			Optional: true,
			PlanModifiers: []planmodifier.Object{
				objectplanmodifier.UseStateForUnknown(),
			},
			Attributes: map[string]schema.Attribute{
				"catalog": schema.StringAttribute{
					Optional: true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
				"image_id": schema.StringAttribute{
					Optional: true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
				"instance_count_by_group": schema.Int64Attribute{
					Optional: true,
					PlanModifiers: []planmodifier.Int64{
						int64planmodifier.UseStateForUnknown(),
					},
				},
				"instance_type": schema.StringAttribute{
					Optional: true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
				"recipes": schema.SetAttribute{
					Optional:    true,
					ElementType: types.StringType,
				},
			},
		},
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
		"new_network_params": schema.SingleNestedAttribute{
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{
				objectplanmodifier.UseStateForUnknown(),
			},
			Attributes: map[string]schema.Attribute{
				"network_cidr": schema.StringAttribute{
					Required: true,
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
			Optional: true,
			Computed: true,
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
	req.EnableTunnel = model.EnableTunnel.ValueBool()
	req.EncryptionKeyResourceGroupName = model.EncryptionKeyResourceGroupName.ValueString()
	req.EncryptionKeyURL = model.EncryptionKeyURL.ValueString()
	req.EnvironmentName = model.EnvironmentName.ValueStringPointer()
	req.EndpointAccessGatewayScheme = model.EndpointAccessGatewayScheme.ValueString()
	req.EndpointAccessGatewaySubnetIds = utils.FromSetValueToStringList(model.EndpointAccessGatewaySubnetIds)
	if !model.ExistingNetworkParams.IsNull() && !model.ExistingNetworkParams.IsUnknown() {
		tflog.Debug(ctx, "existing network params")
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
	}

	if !model.FreeIpa.IsNull() && !model.FreeIpa.IsUnknown() {
		var freeIpaDetails azureFreeIpaDetails
		model.FreeIpa.As(ctx, &freeIpaDetails, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
		req.FreeIpa = &environmentsmodels.AzureFreeIpaCreationRequest{
			InstanceCountByGroup: int32(freeIpaDetails.InstanceCountByGroup.ValueInt64()),
			InstanceType:         freeIpaDetails.InstanceType.ValueString(),
			Recipes:              utils.FromSetValueToStringList(freeIpaDetails.Recipes),
		}
		req.Image = &environmentsmodels.FreeIpaImageRequest{
			Catalog: freeIpaDetails.Catalog.ValueString(),
			ID:      freeIpaDetails.ImageID.ValueString(),
		}
	}

	if model.LogStorage != nil {
		req.LogStorage = &environmentsmodels.AzureLogStorageRequest{
			ManagedIdentity:           model.LogStorage.ManagedIdentity.ValueStringPointer(),
			StorageLocationBase:       model.LogStorage.StorageLocationBase.ValueStringPointer(),
			BackupStorageLocationBase: model.LogStorage.BackupStorageLocationBase.ValueString(),
		}
	}
	if !model.NewNetworkParams.IsNull() && !model.NewNetworkParams.IsUnknown() {
		var newNetworkParams newNetworkParams
		model.NewNetworkParams.As(ctx, &newNetworkParams, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
		req.NewNetworkParams = &environmentsmodels.CreateAzureEnvironmentRequestNewNetworkParams{
			NetworkCidr: newNetworkParams.NetworkCidr.ValueStringPointer(),
		}
	}
	req.ProxyConfigName = model.ProxyConfigName.ValueString()
	req.PublicKey = model.PublicKey.ValueStringPointer()
	req.Region = model.Region.ValueStringPointer()
	req.ReportDeploymentLogs = model.ReportDeploymentLogs.ValueBool()
	req.ResourceGroupName = model.ResourceGroupName.ValueString()
	req.SecurityAccess = &environmentsmodels.SecurityAccessRequest{
		Cidr:                    model.SecurityAccess.Cidr.ValueString(),
		DefaultSecurityGroupIDs: utils.FromSetValueToStringList(model.SecurityAccess.DefaultSecurityGroupIDs),
		DefaultSecurityGroupID:  model.SecurityAccess.DefaultSecurityGroupID.ValueString(),
		SecurityGroupIDsForKnox: utils.FromSetValueToStringList(model.SecurityAccess.SecurityGroupIDsForKnox),
		SecurityGroupIDForKnox:  model.SecurityAccess.SecurityGroupIDForKnox.ValueString(),
	}
	if !model.Tags.IsNull() {
		req.Tags = make([]*environmentsmodels.TagRequest, len(model.Tags.Elements()))
		i := 0
		for k, v := range model.Tags.Elements() {
			val, diag := v.(basetypes.StringValuable).ToStringValue(ctx)
			if !diag.HasError() {
				req.Tags[i] = &environmentsmodels.TagRequest{
					Key:   &k,
					Value: val.ValueStringPointer(),
				}
			}
			i++
		}
	}
	req.UsePublicIP = model.UsePublicIP.ValueBoolPointer()
	req.WorkloadAnalytics = model.WorkloadAnalytics.ValueBool()
	return req
}
