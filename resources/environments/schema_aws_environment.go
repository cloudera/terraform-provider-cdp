package environments

import (
	"context"

	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var AwsEnvironmentSchema = schema.Schema{
	MarkdownDescription: "The environment is a logical entity that represents the association of your user account with multiple compute resources using which you can provision and manage workloads.",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"crn": schema.StringAttribute{
			Computed: true,
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
		"cloud_storage_logging": schema.BoolAttribute{
			Optional: true,
			Computed: true,
		},
		"create_private_subnets": schema.BoolAttribute{
			Optional: true,
			Computed: true,
		},
		"create_service_endpoints": schema.BoolAttribute{
			Optional: true,
			Computed: true,
		},
		"s3_guard_table_name": schema.StringAttribute{
			Optional: true,
			Computed: true,
		},
		"credential_name": schema.StringAttribute{
			Required: true,
		},
		"description": schema.StringAttribute{
			Optional: true,
			Computed: true,
		},
		"enable_workload_analytics": schema.BoolAttribute{
			Optional: true,
			Computed: true,
		},
		"enable_tunnel": schema.BoolAttribute{
			Optional: true,
			Computed: true,
		},
		"encryption_key_arn": schema.StringAttribute{
			Optional: true,
			Computed: true,
		},
		"endpoint_access_gateway_scheme": schema.StringAttribute{
			Optional: true,
			Computed: true,
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
			Attributes: map[string]schema.Attribute{
				"catalog": schema.StringAttribute{
					Optional: true,
					Computed: true,
				},
				"image_id": schema.StringAttribute{
					Optional: true,
					Computed: true,
				},
				"instance_count_by_group": schema.Int64Attribute{
					Optional: true,
					Computed: true,
				},
				"instance_type": schema.StringAttribute{
					Optional: true,
					Computed: true,
				},
				"multi_az": schema.BoolAttribute{
					Optional: true,
					Computed: true,
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
				},
			},
		},
		"region": schema.StringAttribute{
			Required: true,
		},
		"report_deployment_logs": schema.BoolAttribute{
			Optional: true,
			Computed: true,
		},
		"network_cidr": schema.StringAttribute{
			Optional: true,
			Computed: true,
		},
		"proxy_config_name": schema.StringAttribute{
			Optional: true,
			Computed: true,
		},
		"security_access": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"cidr": schema.StringAttribute{
					Optional: true,
					Computed: true,
				},
				"default_security_group_id": schema.StringAttribute{
					Optional: true,
					Computed: true,
				},
				"default_security_group_ids": schema.SetAttribute{
					Optional:    true,
					Computed:    true,
					ElementType: types.StringType,
				},
				"security_group_id_for_knox": schema.StringAttribute{
					Optional: true,
					Computed: true,
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
		},
		"status_reason": schema.StringAttribute{
			Computed: true,
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
			Optional: true,
			Computed: true,
		},
		"workload_analytics": schema.BoolAttribute{
			Optional: true,
			Computed: true,
		},
		"vpc_id": schema.StringAttribute{
			Optional: true,
			Computed: true,
		},
	},
}

func ToAwsEnvrionmentRequest(ctx context.Context, model *awsEnvironmentResourceModel) *environmentsmodels.CreateAWSEnvironmentRequest {
	res := &environmentsmodels.CreateAWSEnvironmentRequest{}
	res.Authentication = &environmentsmodels.AuthenticationRequest{
		PublicKey:   model.Authentication.PublicKey.ValueString(),
		PublicKeyID: model.Authentication.PublicKeyID.ValueString(),
	}
	res.CloudStorageLogging = model.CloudStorageLogging.ValueBool()
	res.CreatePrivateSubnets = model.CreatePrivateSubnets.ValueBool()
	res.CreateServiceEndpoints = model.CreateServiceEndpoints.ValueBool()
	res.CredentialName = model.CredentialName.ValueStringPointer()
	res.Description = model.Description.ValueString()
	res.EnableTunnel = model.EnableTunnel.ValueBool()
	res.EnableWorkloadAnalytics = model.EnableWorkloadAnalytics.ValueBool()
	res.EncryptionKeyArn = model.EncryptionKeyArn.ValueString()
	res.EndpointAccessGatewayScheme = model.EndpointAccessGatewayScheme.ValueString()
	res.EndpointAccessGatewaySubnetIds = utils.FromSetValueToStringList(model.EndpointAccessGatewaySubnetIds)
	res.EnvironmentName = model.EnvironmentName.ValueStringPointer()

	if !model.FreeIpa.IsNull() && !model.FreeIpa.IsUnknown() {
		var freeIpaDetails AWSFreeIpaDetails
		model.FreeIpa.As(ctx, freeIpaDetails, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
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
		}
	}
	res.TunnelType = environmentsmodels.TunnelType(model.TunnelType.ValueString())
	res.VpcID = model.VpcID.ValueString()
	res.WorkloadAnalytics = model.WorkloadAnalytics.ValueBool()
	return res
}
