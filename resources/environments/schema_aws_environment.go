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

	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
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
		"create_private_subnets": schema.BoolAttribute{
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"create_service_endpoints": schema.BoolAttribute{
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"s3_guard_table_name": schema.StringAttribute{
			Optional: true,
			Computed: true,
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
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"encryption_key_arn": schema.StringAttribute{
			Optional: true,
			Computed: true,
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
			Computed:    true,
			ElementType: types.StringType,
		},
		"environment_name": schema.StringAttribute{
			Required: true,
		},
		"freeipa": schema.SingleNestedAttribute{
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{
				objectplanmodifier.UseStateForUnknown(),
			},
			Attributes: map[string]schema.Attribute{
				"catalog": schema.StringAttribute{
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
				"image_id": schema.StringAttribute{
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
				"instance_count_by_group": schema.Int64Attribute{
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Int64{
						int64planmodifier.UseStateForUnknown(),
					},
				},
				"instance_type": schema.StringAttribute{
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
				"multi_az": schema.BoolAttribute{
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Bool{
						boolplanmodifier.UseStateForUnknown(),
					},
				},
				"recipes": schema.SetAttribute{
					Optional:    true,
					Computed:    true,
					ElementType: types.StringType,
				},
			},
		},
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
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"network_cidr": schema.StringAttribute{
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"proxy_config_name": schema.StringAttribute{
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
					Computed:    true,
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
					Computed:    true,
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
			Optional:    true,
			Computed:    true,
			ElementType: types.StringType,
		},
		"tags": schema.MapAttribute{
			Optional:    true,
			Computed:    true,
			ElementType: types.StringType,
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
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
	},
}

func ToAwsEnvrionmentRequest(ctx context.Context, model *awsEnvironmentResourceModel) *environmentsmodels.CreateAWSEnvironmentRequest {
	res := &environmentsmodels.CreateAWSEnvironmentRequest{}
	res.Authentication = &environmentsmodels.AuthenticationRequest{
		PublicKey:   model.Authentication.PublicKey.ValueString(),
		PublicKeyID: model.Authentication.PublicKeyID.ValueString(),
	}
	res.CreatePrivateSubnets = model.CreatePrivateSubnets.ValueBool()
	res.CreateServiceEndpoints = model.CreateServiceEndpoints.ValueBool()
	res.CredentialName = model.CredentialName.ValueStringPointer()
	res.Description = model.Description.ValueString()
	res.EnableTunnel = model.EnableTunnel.ValueBool()
	res.EncryptionKeyArn = model.EncryptionKeyArn.ValueString()
	res.EndpointAccessGatewayScheme = model.EndpointAccessGatewayScheme.ValueString()
	res.EndpointAccessGatewaySubnetIds = utils.FromSetValueToStringList(model.EndpointAccessGatewaySubnetIds)
	res.EnvironmentName = model.EnvironmentName.ValueStringPointer()

	if !model.FreeIpa.IsNull() && !model.FreeIpa.IsUnknown() {
		var freeIpaDetails AWSFreeIpaDetails
		model.FreeIpa.As(ctx, &freeIpaDetails, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
		res.FreeIpa = &environmentsmodels.AWSFreeIpaCreationRequest{
			InstanceCountByGroup: int32(freeIpaDetails.InstanceCountByGroup.ValueInt64()),
			InstanceType:         freeIpaDetails.InstanceType.ValueString(),
			MultiAz:              freeIpaDetails.MultiAz.ValueBool(),
			Recipes:              utils.FromSetValueToStringList(freeIpaDetails.Recipes),
		}
		res.Image = &environmentsmodels.FreeIpaImageRequest{
			Catalog: freeIpaDetails.Catalog.ValueStringPointer(),
			ID:      freeIpaDetails.ImageID.ValueStringPointer(),
		}
	}

	if model.LogStorage != nil {
		res.LogStorage = &environmentsmodels.AwsLogStorageRequest{
			InstanceProfile:           model.LogStorage.InstanceProfile.ValueStringPointer(),
			StorageLocationBase:       model.LogStorage.StorageLocationBase.ValueStringPointer(),
			BackupStorageLocationBase: model.LogStorage.BackupStorageLocationBase.ValueString(),
		}
	}
	if !model.NetworkCidr.IsNull() && !model.NetworkCidr.IsUnknown() {
		res.NetworkCidr = model.NetworkCidr.ValueString()
	}
	res.ProxyConfigName = model.ProxyConfigName.ValueString()
	res.Region = model.Region.ValueStringPointer()
	res.ReportDeploymentLogs = model.ReportDeploymentLogs.ValueBool()
	res.S3GuardTableName = model.S3GuardTableName.ValueString()
	res.SecurityAccess = &environmentsmodels.SecurityAccessRequest{
		Cidr:                    model.SecurityAccess.Cidr.ValueString(),
		DefaultSecurityGroupIDs: utils.FromSetValueToStringList(model.SecurityAccess.DefaultSecurityGroupIDs),
		DefaultSecurityGroupID:  model.SecurityAccess.DefaultSecurityGroupID.ValueString(),
		SecurityGroupIDsForKnox: utils.FromSetValueToStringList(model.SecurityAccess.SecurityGroupIDsForKnox),
		SecurityGroupIDForKnox:  model.SecurityAccess.SecurityGroupIDForKnox.ValueString(),
	}
	if !model.SubnetIds.IsNull() && !model.SubnetIds.IsUnknown() {
		res.SubnetIds = utils.FromSetValueToStringList(model.SubnetIds)
	}
	if !model.Tags.IsNull() {
		res.Tags = make([]*environmentsmodels.TagRequest, len(model.Tags.Elements()))
		i := 0
		for k, v := range model.Tags.Elements() {
			val, diag := v.(basetypes.StringValuable).ToStringValue(ctx)
			if !diag.HasError() {
				res.Tags[i] = &environmentsmodels.TagRequest{
					Key:   &k,
					Value: val.ValueStringPointer(),
				}
			}
			i++
		}
	}
	res.VpcID = model.VpcID.ValueString()
	res.WorkloadAnalytics = model.WorkloadAnalytics.ValueBool()
	return res
}
