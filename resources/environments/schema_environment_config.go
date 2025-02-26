// Copyright 2025 Cloudera. All Rights Reserved.
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
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var EnvironmentConfigSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Optional: true,
		},
		"crn": schema.StringAttribute{
			Optional: true,
		},
		"aws": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					Computed: true,
				},
				"crn": schema.StringAttribute{
					Computed: true,
				},
				"polling_options": schema.SingleNestedAttribute{
					Computed: true,
					Attributes: map[string]schema.Attribute{
						"async": schema.BoolAttribute{
							Computed: true,
						},
						"polling_timeout": schema.Int64Attribute{
							Computed: true,
						},
						"call_failure_threshold": schema.Int64Attribute{
							Computed: true,
						},
					},
				},
				"authentication": schema.SingleNestedAttribute{
					Computed: true,
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
					Computed: true,
				},
				"create_service_endpoints": schema.BoolAttribute{
					Computed: true,
				},
				"s3_guard_table_name": schema.StringAttribute{
					Computed: true,
				},
				"credential_name": schema.StringAttribute{
					Computed: true,
				},
				"description": schema.StringAttribute{
					Computed: true,
				},
				"enable_tunnel": schema.BoolAttribute{
					Computed: true,
				},
				"encryption_key_arn": schema.StringAttribute{
					Computed: true,
				},
				"endpoint_access_gateway_scheme": schema.StringAttribute{
					Computed: true,
				},
				"endpoint_access_gateway_subnet_ids": schema.SetAttribute{
					Computed:    true,
					ElementType: types.StringType,
				},
				"environment_name": schema.StringAttribute{
					Computed: true,
				},
				"cascading_delete": schema.BoolAttribute{
					Computed: true,
				},
				"freeipa": FreeIpaConfigSchema,
				"log_storage": schema.SingleNestedAttribute{
					Computed: true,
					Attributes: map[string]schema.Attribute{
						"instance_profile": schema.StringAttribute{
							Computed: true,
						},
						"storage_location_base": schema.StringAttribute{
							Computed: true,
						},
						"backup_storage_location_base": schema.StringAttribute{
							Computed: true,
						},
					},
				},
				"region": schema.StringAttribute{
					Computed: true,
				},
				"report_deployment_logs": schema.BoolAttribute{
					Computed: true,
				},
				"network_cidr": schema.StringAttribute{
					Computed: true,
				},
				"proxy_config_name": schema.StringAttribute{
					Computed: true,
				},
				"security_access": schema.SingleNestedAttribute{
					Computed: true,
					Attributes: map[string]schema.Attribute{
						"cidr": schema.StringAttribute{
							Computed: true,
						},
						"default_security_group_id": schema.StringAttribute{
							Computed: true,
						},
						"default_security_group_ids": schema.SetAttribute{
							Computed:    true,
							Optional:    true,
							ElementType: types.StringType,
						},
						"security_group_id_for_knox": schema.StringAttribute{
							Computed: true,
						},
						"security_group_ids_for_knox": schema.SetAttribute{
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
					Computed:    true,
					ElementType: types.StringType,
				},
				"tags": schema.MapAttribute{
					Computed:    true,
					ElementType: types.StringType,
				},
				"tunnel_type": schema.StringAttribute{
					Computed: true,
				},
				"workload_analytics": schema.BoolAttribute{
					Computed: true,
				},
				"vpc_id": schema.StringAttribute{
					Computed: true,
				},
			},
		},
		"azure": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"crn": schema.StringAttribute{
					Computed: true,
				},
				"create_private_endpoints": schema.BoolAttribute{
					Computed: true,
				},
				"credential_name": schema.StringAttribute{
					Computed: true,
				},
				"description": schema.StringAttribute{
					Computed: true,
				},
				"enable_outbound_load_balancer": schema.BoolAttribute{
					Computed: true,
				},
				"enable_tunnel": schema.BoolAttribute{
					Computed: true,
				},
				"endpoint_access_gateway_scheme": schema.StringAttribute{
					Computed: true,
				},
				"endpoint_access_gateway_subnet_ids": schema.SetAttribute{
					Computed:    true,
					ElementType: types.StringType,
				},
				"encryption_key_resource_group_name": schema.StringAttribute{
					Optional: true,
				},
				"encryption_key_url": schema.StringAttribute{
					Optional: true,
				},
				"encryption_at_host": schema.BoolAttribute{
					Computed: true,
				},
				"environment_name": schema.StringAttribute{
					Computed: true,
				},
				"cascading_delete": schema.BoolAttribute{
					Computed: true,
				},
				"existing_network_params": schema.SingleNestedAttribute{
					Computed: true,
					Attributes: map[string]schema.Attribute{
						"aks_private_dns_zone_id": schema.StringAttribute{
							Computed: true,
						},
						"database_private_dns_zone_id": schema.StringAttribute{
							Computed: true,
						},
						"network_id": schema.StringAttribute{
							Computed: true,
						},
						"resource_group_name": schema.StringAttribute{
							Computed: true,
						},
						"subnet_ids": schema.SetAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
						"flexible_server_subnet_ids": schema.SetAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
					},
				},
				"freeipa": FreeIpaConfigSchema,
				"log_storage": schema.SingleNestedAttribute{
					Computed: true,
					Attributes: map[string]schema.Attribute{
						"managed_identity": schema.StringAttribute{
							Computed: true,
						},
						"storage_location_base": schema.StringAttribute{
							Computed: true,
						},
						"backup_storage_location_base": schema.StringAttribute{
							Computed: true,
						},
					},
				},
				"proxy_config_name": schema.StringAttribute{
					Computed: true,
				},
				"public_key": schema.StringAttribute{
					Computed: true,
				},
				"region": schema.StringAttribute{
					Computed: true,
				},
				"report_deployment_logs": schema.BoolAttribute{
					Computed: true,
				},
				"resource_group_name": schema.StringAttribute{
					Computed: true,
				},
				"security_access": schema.SingleNestedAttribute{
					Computed: true,
					Attributes: map[string]schema.Attribute{
						"cidr": schema.StringAttribute{
							Computed: true,
						},
						"default_security_group_id": schema.StringAttribute{
							Computed: true,
						},
						"default_security_group_ids": schema.SetAttribute{
							Optional:    true,
							ElementType: types.StringType,
						},
						"security_group_id_for_knox": schema.StringAttribute{
							Computed: true,
						},
						"security_group_ids_for_knox": schema.SetAttribute{
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
				"encryption_user_managed_identity": schema.StringAttribute{
					Computed: true,
				},
				"tags": schema.MapAttribute{
					Computed:    true,
					ElementType: types.StringType,
				},
				"use_public_ip": schema.BoolAttribute{
					Computed: true,
				},
				"workload_analytics": schema.BoolAttribute{
					Computed: true,
				},
			},
		},
		"gcp": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"environment_name": schema.StringAttribute{
					Required: true,
				},
				"cascading_delete": schema.BoolAttribute{
					Computed: true,
				},
				"polling_options": schema.SingleNestedAttribute{
					Computed: true,
					Attributes: map[string]schema.Attribute{
						"async": schema.BoolAttribute{
							Computed: true,
						},
						"polling_timeout": schema.Int64Attribute{
							Computed: true,
						},
						"call_failure_threshold": schema.Int64Attribute{
							Computed: true,
						},
					},
				},
				"credential_name": schema.StringAttribute{
					Computed: true,
				},
				"region": schema.StringAttribute{
					Computed: true,
				},
				"public_key": schema.StringAttribute{
					Computed: true,
				},
				"use_public_ip": schema.BoolAttribute{
					Computed: true,
				},
				"existing_network_params": schema.SingleNestedAttribute{
					Computed: true,
					Attributes: map[string]schema.Attribute{
						"network_name": schema.StringAttribute{
							Computed: true,
						},
						"subnet_names": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
						"shared_project_id": schema.StringAttribute{
							Computed: true,
						},
					},
				},
				"security_access": schema.SingleNestedAttribute{
					Computed: true,
					Attributes: map[string]schema.Attribute{
						"security_group_id_for_knox": schema.StringAttribute{
							Computed: true,
						},
						"default_security_group_id": schema.StringAttribute{
							Computed: true,
						},
					},
				},
				"log_storage": schema.SingleNestedAttribute{
					Computed: true,
					Attributes: map[string]schema.Attribute{
						"storage_location_base": schema.StringAttribute{
							Computed: true,
						},
						"service_account_email": schema.StringAttribute{
							Computed: true,
						},
						"backup_storage_location_base": schema.StringAttribute{
							Computed: true,
						},
					},
				},
				"description": schema.StringAttribute{
					Computed: true,
				},
				"enable_tunnel": schema.BoolAttribute{
					Computed: true,
				},
				"workload_analytics": schema.BoolAttribute{
					Computed: true,
				},
				"report_deployment_logs": schema.BoolAttribute{
					Computed: true,
				},
				"freeipa": FreeIpaConfigSchema,
				"endpoint_access_gateway_scheme": schema.StringAttribute{
					Computed: true,
				},
				"tags": schema.MapAttribute{
					Computed:    true,
					ElementType: types.StringType,
				},
				"proxy_config_name": schema.StringAttribute{
					Computed: true,
				},
				"encryption_key": schema.StringAttribute{
					Computed: true,
				},
				"availability_zones": schema.ListAttribute{
					Computed:    true,
					ElementType: types.StringType,
				},
				"id": schema.StringAttribute{
					Computed: true,
				},
				"crn": schema.StringAttribute{
					Computed: true,
				},
				"status": schema.StringAttribute{
					Computed: true,
				},
				"status_reason": schema.StringAttribute{
					Computed: true,
				},
			},
		},
	},
}

var FreeIpaConfigSchema = schema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]schema.Attribute{
		"catalog": schema.StringAttribute{
			Computed: true,
		},
		"image_id": schema.StringAttribute{
			Computed: true,
		},
		"os": schema.StringAttribute{
			Computed: true,
		},
		"instance_count_by_group": schema.Int64Attribute{
			Computed: true,
		},
		"instance_type": schema.StringAttribute{
			Computed: true,
		},
		"instances": schema.SetNestedAttribute{
			Computed: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"availability_zone": schema.StringAttribute{
						Computed: true,
					},
					"discovery_fqdn": schema.StringAttribute{
						Computed: true,
					},
					"instance_group": schema.StringAttribute{
						Computed: true,
					},
					"instance_id": schema.StringAttribute{
						Computed: true,
					},
					"instance_status": schema.StringAttribute{
						Computed: true,
					},
					"instance_status_reason": schema.StringAttribute{
						Computed: true,
					},
					"instance_type": schema.StringAttribute{
						Computed: true,
					},
					"instance_vm_type": schema.StringAttribute{
						Computed: true,
					},
					"life_cycle": schema.StringAttribute{
						Computed: true,
					},
					"private_ip": schema.StringAttribute{
						Computed: true,
					},
					"public_ip": schema.StringAttribute{
						Computed: true,
					},
					"ssh_port": schema.Int64Attribute{
						Computed: true,
					},
					"subnet_id": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
		"multi_az": schema.BoolAttribute{
			Computed: true,
		},
		"recipes": schema.SetAttribute{
			Computed:    true,
			ElementType: types.StringType,
		},
	},
}
