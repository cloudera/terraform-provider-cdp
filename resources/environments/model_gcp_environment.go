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

type gcpEnvironmentResourceModel struct {
	EnvironmentName types.String `tfsdk:"environment_name"`

	PollingOptions *utils.PollingOptions `tfsdk:"polling_options"`

	CredentialName types.String `tfsdk:"credential_name"`

	Region types.String `tfsdk:"region"`

	PublicKey types.String `tfsdk:"public_key"`

	UsePublicIp types.Bool `tfsdk:"use_public_ip"`

	ExistingNetworkParams *ExistingNetworkParams `tfsdk:"existing_network_params"`

	SecurityAccess *GcpSecurityAccess `tfsdk:"security_access"`

	LogStorage *GcpLogStorage `tfsdk:"log_storage"`

	Description types.String `tfsdk:"description"`

	EnableTunnel types.Bool `tfsdk:"enable_tunnel"`

	WorkloadAnalytics types.Bool `tfsdk:"workload_analytics"`

	ReportDeploymentLogs types.Bool `tfsdk:"report_deployment_logs"`

	FreeIpa *GcpFreeIpa `tfsdk:"freeipa"`

	EndpointAccessGatewayScheme types.String `tfsdk:"endpoint_access_gateway_scheme"`

	Tags types.Map `tfsdk:"tags"`

	ProxyConfigName types.String `tfsdk:"proxy_config_name"`

	EncryptionKey types.String `tfsdk:"encryption_key"`

	AvailabilityZones []types.String `tfsdk:"availability_zones"`

	ID types.String `tfsdk:"id"`

	Crn types.String `tfsdk:"crn"`

	Status types.String `tfsdk:"status"`

	StatusReason types.String `tfsdk:"status_reason"`
}

type GcpFreeIpa struct {
	InstanceCountByGroup types.Int64        `tfsdk:"instance_count_by_group"`
	Recipes              types.Set          `tfsdk:"recipes"`
	InstanceType         types.String       `tfsdk:"instance_type"`
	Instances            *[]FreeIpaInstance `tfsdk:"instances"`
}

type ExistingNetworkParams struct {
	NetworkName     types.String `tfsdk:"network_name"`
	SubnetNames     types.List   `tfsdk:"subnet_names"`
	SharedProjectId types.String `tfsdk:"shared_project_id"`
}

type GcpLogStorage struct {
	StorageLocationBase       types.String `tfsdk:"storage_location_base"`
	ServiceAccountEmail       types.String `tfsdk:"service_account_email"`
	BackupStorageLocationBase types.String `tfsdk:"backup_storage_location_base"`
}

type GcpSecurityAccess struct {
	SecurityGroupIdForKnox types.String `tfsdk:"security_group_id_for_knox"`
	DefaultSecurityGroupId types.String `tfsdk:"default_security_group_id"`
}
