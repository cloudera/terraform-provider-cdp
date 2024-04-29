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

type azureEnvironmentResourceModel struct {
	ID types.String `tfsdk:"id"`

	Crn types.String `tfsdk:"crn"`

	PollingOptions *utils.PollingOptions `tfsdk:"polling_options"`

	CreatePrivateEndpoints types.Bool `tfsdk:"create_private_endpoints"`

	CredentialName types.String `tfsdk:"credential_name"`

	Description types.String `tfsdk:"description"`

	EnableOutboundLoadBalancer types.Bool `tfsdk:"enable_outbound_load_balancer"`

	EnableTunnel types.Bool `tfsdk:"enable_tunnel"`

	EncryptionKeyResourceGroupName types.String `tfsdk:"encryption_key_resource_group_name"`

	EncryptionKeyURL types.String `tfsdk:"encryption_key_url"`

	EncryptionAtHost types.Bool `tfsdk:"encryption_at_host"`

	EnvironmentName types.String `tfsdk:"environment_name"`

	ExistingNetworkParams types.Object `tfsdk:"existing_network_params"`

	FreeIpa types.Object `tfsdk:"freeipa"`

	LogStorage *azureLogStorage `tfsdk:"log_storage"`

	NewNetworkParams types.Object `tfsdk:"new_network_params"`

	ProxyConfigName types.String `tfsdk:"proxy_config_name"`

	PublicKey types.String `tfsdk:"public_key"`

	Region types.String `tfsdk:"region"`

	ReportDeploymentLogs types.Bool `tfsdk:"report_deployment_logs"`

	ResourceGroupName types.String `tfsdk:"resource_group_name"`

	SecurityAccess *SecurityAccess `tfsdk:"security_access"`

	Status types.String `tfsdk:"status"`

	StatusReason types.String `tfsdk:"status_reason"`

	Tags types.Map `tfsdk:"tags"`

	UsePublicIP types.Bool `tfsdk:"use_public_ip"`

	WorkloadAnalytics types.Bool `tfsdk:"workload_analytics"`

	EndpointAccessGatewayScheme types.String `tfsdk:"endpoint_access_gateway_scheme"`

	EndpointAccessGatewaySubnetIds types.Set `tfsdk:"endpoint_access_gateway_subnet_ids"`

	EncryptionUserManagedIdentity types.String `tfsdk:"encryption_user_managed_identity"`
}

type existingAzureNetwork struct {
	AksPrivateDNSZoneID types.String `tfsdk:"aks_private_dns_zone_id"`

	DatabasePrivateDNSZoneID types.String `tfsdk:"database_private_dns_zone_id"`

	NetworkID types.String `tfsdk:"network_id"`

	ResourceGroupName types.String `tfsdk:"resource_group_name"`

	SubnetIds types.Set `tfsdk:"subnet_ids"`

	FlexibleServerSubnetIds types.Set `tfsdk:"flexible_server_subnet_ids"`
}

type azureLogStorage struct {
	ManagedIdentity types.String `tfsdk:"managed_identity"`

	StorageLocationBase types.String `tfsdk:"storage_location_base"`

	BackupStorageLocationBase types.String `tfsdk:"backup_storage_location_base"`
}

type newNetworkParams struct {
	NetworkCidr types.String `tfsdk:"network_cidr"`
}
