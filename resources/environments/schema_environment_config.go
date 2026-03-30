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
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

func (e *environmentConfigDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attr := map[string]schema.Attribute{}
	utils.AppendToDataSourceSchema(attr, generalAttributes)
	utils.AppendToDataSourceSchema(attr, map[string]schema.Attribute{
		"aws":   awsSchema,
		"azure": azureSchema,
		"gcp":   gcpSchema,
	})
	resp.Schema = schema.Schema{
		MarkdownDescription: "Collects environment information for re-creation purposes.",
		Attributes:          attr,
	}
}

var awsSchema = schema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"crn": schema.StringAttribute{
			Computed: true,
		},
		"polling_options": pollingOptionsSchema,
		"authentication": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"public_key": schema.StringAttribute{
					Computed: true,
					Optional: true,
				},
				"public_key_id": schema.StringAttribute{
					Computed: true,
					Optional: true,
				},
			},
		},
		"compute_cluster": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"enabled": schema.BoolAttribute{
					Computed: true,
				},
				"configuration": schema.SingleNestedAttribute{
					Computed: true,
					Attributes: map[string]schema.Attribute{
						"private_cluster": schema.BoolAttribute{
							Computed: true,
						},
						"kube_api_authorized_ip_ranges": schema.SetAttribute{
							ElementType: types.StringType,
							Computed:    true,
						},
						"worker_node_subnets": schema.SetAttribute{
							ElementType: types.StringType,
							Computed:    true,
						},
					},
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
		"delete_options": deleteOptionsSchema,
		"freeipa":        freeIpaConfigSchema,
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
					ElementType: types.StringType,
					Computed:    true,
					Optional:    true,
				},
				"security_group_id_for_knox": schema.StringAttribute{
					Computed: true,
				},
				"security_group_ids_for_knox": schema.SetAttribute{
					ElementType: types.StringType,
					Computed:    true,
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
		"custom_docker_registry": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"crn": schema.StringAttribute{
					Computed: true,
				},
			},
		},
		"security": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"se_linux": schema.StringAttribute{
					Computed: true,
				},
			},
		},
		"environment_type": schema.StringAttribute{
			Computed: true,
		},
	},
}

var azureSchema = schema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"crn": schema.StringAttribute{
			Computed: true,
		},
		"polling_options": pollingOptionsSchema,
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
		"delete_options": deleteOptionsSchema,
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
		"freeipa": freeIpaConfigSchema,
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
		"compute_cluster": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"enabled": schema.BoolAttribute{
					Computed: true,
				},
				"configuration": schema.SingleNestedAttribute{
					Computed: true,
					Attributes: map[string]schema.Attribute{
						"private_cluster": schema.BoolAttribute{
							Computed: true,
						},
						"kube_api_authorized_ip_ranges": schema.SetAttribute{
							ElementType: types.StringType,
							Computed:    true,
						},
						"outbound_type": schema.StringAttribute{
							Computed: true,
						},
						"worker_node_subnets": schema.SetAttribute{
							ElementType: types.StringType,
							Computed:    true,
						},
					},
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
		"custom_docker_registry": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"crn": schema.StringAttribute{
					Computed: true,
				},
			},
		},
		"security": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"se_linux": schema.StringAttribute{
					Computed: true,
				},
			},
		},
		"environment_type": schema.StringAttribute{
			Computed: true,
		},
		"data_services": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"shared_managed_identity": schema.StringAttribute{
					Computed: true,
				},
				"aks_private_dns_zone_id": schema.StringAttribute{
					Computed: true,
				},
			},
		},
		"flexible_server_subnet_ids": schema.SetAttribute{
			Computed:    true,
			ElementType: types.StringType,
		},
		"availability_zones": schema.SetAttribute{
			Computed:    true,
			ElementType: types.StringType,
		},
	},
}

var gcpSchema = schema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]schema.Attribute{
		"environment_name": schema.StringAttribute{
			Computed: true,
		},
		"cascading_delete": schema.BoolAttribute{
			Computed: true,
		},
		"delete_options":  deleteOptionsSchema,
		"polling_options": pollingOptionsSchema,
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
		"freeipa": freeIpaConfigSchema,
		"endpoint_access_gateway_scheme": schema.StringAttribute{
			Computed: true,
		},
		"endpoint_access_gateway_subnet_ids": schema.SetAttribute{
			Computed:    true,
			ElementType: types.StringType,
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
		"custom_docker_registry": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"crn": schema.StringAttribute{
					Computed: true,
				},
			},
		},
		"security": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"se_linux": schema.StringAttribute{
					Computed: true,
				},
			},
		},
		"environment_type": schema.StringAttribute{
			Computed: true,
		},
	},
}

var freeIpaConfigSchema = schema.SingleNestedAttribute{
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
		"instance_count_by_group": schema.Int32Attribute{
			Computed: true,
		},
		"instance_type": schema.StringAttribute{
			Computed: true,
		},
		"architecture": schema.StringAttribute{
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

var pollingOptionsSchema = schema.SingleNestedAttribute{
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
}

var deleteOptionsSchema = schema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]schema.Attribute{
		"cascading": schema.BoolAttribute{
			Computed: true,
		},
		"forced": schema.BoolAttribute{
			Computed: true,
		},
	},
}

var generalAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Description: "The name of the environment. Either this or the CRN must be specified - but not both -. In case of hybrid environments this can be only CRN.",
		Optional:    true,
	},
	"crn": schema.StringAttribute{
		Description: "The CRN of the environment. Either this or the name must be specified - but not both -. In case of hybrid environments this can be only CRN.",
		Optional:    true,
	},
}
