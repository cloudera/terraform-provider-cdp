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
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

type awsEnvironmentResourceModel struct {
	ID types.String `tfsdk:"id"`

	Crn types.String `tfsdk:"crn"`

	PollingOptions *utils.PollingOptions `tfsdk:"polling_options"`

	Authentication *Authentication `tfsdk:"authentication"`

	CreatePrivateSubnets types.Bool `tfsdk:"create_private_subnets"`

	CreateServiceEndpoints types.Bool `tfsdk:"create_service_endpoints"`

	CredentialName types.String `tfsdk:"credential_name"`

	Description types.String `tfsdk:"description"`

	EnableTunnel types.Bool `tfsdk:"enable_tunnel"`

	EncryptionKeyArn types.String `tfsdk:"encryption_key_arn"`

	EndpointAccessGatewayScheme types.String `tfsdk:"endpoint_access_gateway_scheme"`

	EndpointAccessGatewaySubnetIds types.Set `tfsdk:"endpoint_access_gateway_subnet_ids"`

	EnvironmentName types.String `tfsdk:"environment_name"`

	Cascading types.Bool `tfsdk:"cascading_delete"`

	FreeIpa types.Object `tfsdk:"freeipa"`

	LogStorage *AWSLogStorage `tfsdk:"log_storage"`

	NetworkCidr types.String `tfsdk:"network_cidr"`

	ProxyConfigName types.String `tfsdk:"proxy_config_name"`

	Region types.String `tfsdk:"region"`

	ReportDeploymentLogs types.Bool `tfsdk:"report_deployment_logs"`

	S3GuardTableName types.String `tfsdk:"s3_guard_table_name"`

	SecurityAccess *SecurityAccess `tfsdk:"security_access"`

	Status types.String `tfsdk:"status"`

	StatusReason types.String `tfsdk:"status_reason"`

	SubnetIds types.Set `tfsdk:"subnet_ids"`

	Tags types.Map `tfsdk:"tags"`

	TunnelType types.String `tfsdk:"tunnel_type"`

	VpcID types.String `tfsdk:"vpc_id"`

	WorkloadAnalytics types.Bool `tfsdk:"workload_analytics"`
}

type Authentication struct {
	PublicKey types.String `tfsdk:"public_key"`

	PublicKeyID types.String `tfsdk:"public_key_id"`
}

type AWSLogStorage struct {
	InstanceProfile types.String `tfsdk:"instance_profile"`

	StorageLocationBase types.String `tfsdk:"storage_location_base"`

	BackupStorageLocationBase types.String `tfsdk:"backup_storage_location_base"`
}

type SecurityAccess struct {
	Cidr types.String `tfsdk:"cidr"`

	DefaultSecurityGroupID types.String `tfsdk:"default_security_group_id"`

	DefaultSecurityGroupIDs types.Set `tfsdk:"default_security_group_ids"`

	SecurityGroupIDForKnox types.String `tfsdk:"security_group_id_for_knox"`

	SecurityGroupIDsForKnox types.Set `tfsdk:"security_group_ids_for_knox"`
}
