---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cdp_environments_config Data Source - terraform-provider-cdp"
subcategory: ""
description: |-
  
---

# cdp_environments_config (Data Source)



## Example Usage

```terraform
// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

terraform {
  required_providers {
    cdp = {
      source = "cloudera/cdp"
    }
  }
}

data "cdp_environments_config" "environment_config" {
  name = "<your_env_name>"
}

resource "cdp_environments_aws_environment" "my-aws-environment" {
  environment_name               = "<your_env_name>"
  credential_name                = data.cdp_environments_config.environment_config.aws.credential_name
  region                         = data.cdp_environments_config.environment_config.aws.region
  security_access                = data.cdp_environments_config.environment_config.aws.security_access
  endpoint_access_gateway_scheme = data.cdp_environments_config.environment_config.aws.endpoint_access_gateway_scheme
  authentication = {
    public_key_id = data.cdp_environments_config.environment_config.aws.authentication.public_key_id
  }
  log_storage              = data.cdp_environments_config.environment_config.aws.log_storage
  vpc_id                   = data.cdp_environments_config.environment_config.aws.vpc_id
  subnet_ids               = data.cdp_environments_config.environment_config.aws.subnet_ids
  create_private_subnets   = true
  create_service_endpoints = false
  freeipa = {
    instance_count_by_group = data.cdp_environments_config.environment_config.aws.freeipa.instance_count_by_group
  }
  tags = {
    "made-with" : "CDP Terraform Provider"
  }
}

output "env" {
  value = cdp_environments_aws_environment.my-aws-environment
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `crn` (String)
- `name` (String)

### Read-Only

- `aws` (Attributes) (see [below for nested schema](#nestedatt--aws))
- `azure` (Attributes) (see [below for nested schema](#nestedatt--azure))
- `gcp` (Attributes) (see [below for nested schema](#nestedatt--gcp))

<a id="nestedatt--aws"></a>
### Nested Schema for `aws`

Read-Only:

- `authentication` (Attributes) (see [below for nested schema](#nestedatt--aws--authentication))
- `cascading_delete` (Boolean)
- `create_private_subnets` (Boolean)
- `create_service_endpoints` (Boolean)
- `credential_name` (String)
- `crn` (String)
- `description` (String)
- `enable_tunnel` (Boolean)
- `encryption_key_arn` (String)
- `endpoint_access_gateway_scheme` (String)
- `endpoint_access_gateway_subnet_ids` (Set of String)
- `environment_name` (String)
- `freeipa` (Attributes) (see [below for nested schema](#nestedatt--aws--freeipa))
- `id` (String)
- `log_storage` (Attributes) (see [below for nested schema](#nestedatt--aws--log_storage))
- `network_cidr` (String)
- `polling_options` (Attributes) (see [below for nested schema](#nestedatt--aws--polling_options))
- `proxy_config_name` (String)
- `region` (String)
- `report_deployment_logs` (Boolean)
- `s3_guard_table_name` (String)
- `security_access` (Attributes) (see [below for nested schema](#nestedatt--aws--security_access))
- `status` (String)
- `status_reason` (String)
- `subnet_ids` (Set of String)
- `tags` (Map of String)
- `tunnel_type` (String)
- `vpc_id` (String)
- `workload_analytics` (Boolean)

<a id="nestedatt--aws--authentication"></a>
### Nested Schema for `aws.authentication`

Optional:

- `public_key` (String)
- `public_key_id` (String)


<a id="nestedatt--aws--freeipa"></a>
### Nested Schema for `aws.freeipa`

Read-Only:

- `catalog` (String)
- `image_id` (String)
- `instance_count_by_group` (Number)
- `instance_type` (String)
- `instances` (Attributes Set) (see [below for nested schema](#nestedatt--aws--freeipa--instances))
- `multi_az` (Boolean)
- `os` (String)
- `recipes` (Set of String)

<a id="nestedatt--aws--freeipa--instances"></a>
### Nested Schema for `aws.freeipa.instances`

Read-Only:

- `availability_zone` (String)
- `discovery_fqdn` (String)
- `instance_group` (String)
- `instance_id` (String)
- `instance_status` (String)
- `instance_status_reason` (String)
- `instance_type` (String)
- `instance_vm_type` (String)
- `life_cycle` (String)
- `private_ip` (String)
- `public_ip` (String)
- `ssh_port` (Number)
- `subnet_id` (String)



<a id="nestedatt--aws--log_storage"></a>
### Nested Schema for `aws.log_storage`

Read-Only:

- `backup_storage_location_base` (String)
- `instance_profile` (String)
- `storage_location_base` (String)


<a id="nestedatt--aws--polling_options"></a>
### Nested Schema for `aws.polling_options`

Read-Only:

- `async` (Boolean)
- `call_failure_threshold` (Number)
- `polling_timeout` (Number)


<a id="nestedatt--aws--security_access"></a>
### Nested Schema for `aws.security_access`

Optional:

- `default_security_group_ids` (Set of String)

Read-Only:

- `cidr` (String)
- `default_security_group_id` (String)
- `security_group_id_for_knox` (String)
- `security_group_ids_for_knox` (Set of String)



<a id="nestedatt--azure"></a>
### Nested Schema for `azure`

Optional:

- `encryption_key_resource_group_name` (String)
- `encryption_key_url` (String)

Read-Only:

- `cascading_delete` (Boolean)
- `create_private_endpoints` (Boolean)
- `credential_name` (String)
- `crn` (String)
- `description` (String)
- `enable_outbound_load_balancer` (Boolean)
- `enable_tunnel` (Boolean)
- `encryption_at_host` (Boolean)
- `encryption_user_managed_identity` (String)
- `endpoint_access_gateway_scheme` (String)
- `endpoint_access_gateway_subnet_ids` (Set of String)
- `environment_name` (String)
- `existing_network_params` (Attributes) (see [below for nested schema](#nestedatt--azure--existing_network_params))
- `freeipa` (Attributes) (see [below for nested schema](#nestedatt--azure--freeipa))
- `log_storage` (Attributes) (see [below for nested schema](#nestedatt--azure--log_storage))
- `proxy_config_name` (String)
- `public_key` (String)
- `region` (String)
- `report_deployment_logs` (Boolean)
- `resource_group_name` (String)
- `security_access` (Attributes) (see [below for nested schema](#nestedatt--azure--security_access))
- `status` (String)
- `status_reason` (String)
- `tags` (Map of String)
- `use_public_ip` (Boolean)
- `workload_analytics` (Boolean)

<a id="nestedatt--azure--existing_network_params"></a>
### Nested Schema for `azure.existing_network_params`

Read-Only:

- `aks_private_dns_zone_id` (String)
- `database_private_dns_zone_id` (String)
- `flexible_server_subnet_ids` (Set of String)
- `network_id` (String)
- `resource_group_name` (String)
- `subnet_ids` (Set of String)


<a id="nestedatt--azure--freeipa"></a>
### Nested Schema for `azure.freeipa`

Read-Only:

- `catalog` (String)
- `image_id` (String)
- `instance_count_by_group` (Number)
- `instance_type` (String)
- `instances` (Attributes Set) (see [below for nested schema](#nestedatt--azure--freeipa--instances))
- `multi_az` (Boolean)
- `os` (String)
- `recipes` (Set of String)

<a id="nestedatt--azure--freeipa--instances"></a>
### Nested Schema for `azure.freeipa.instances`

Read-Only:

- `availability_zone` (String)
- `discovery_fqdn` (String)
- `instance_group` (String)
- `instance_id` (String)
- `instance_status` (String)
- `instance_status_reason` (String)
- `instance_type` (String)
- `instance_vm_type` (String)
- `life_cycle` (String)
- `private_ip` (String)
- `public_ip` (String)
- `ssh_port` (Number)
- `subnet_id` (String)



<a id="nestedatt--azure--log_storage"></a>
### Nested Schema for `azure.log_storage`

Read-Only:

- `backup_storage_location_base` (String)
- `managed_identity` (String)
- `storage_location_base` (String)


<a id="nestedatt--azure--security_access"></a>
### Nested Schema for `azure.security_access`

Optional:

- `default_security_group_ids` (Set of String)

Read-Only:

- `cidr` (String)
- `default_security_group_id` (String)
- `security_group_id_for_knox` (String)
- `security_group_ids_for_knox` (Set of String)



<a id="nestedatt--gcp"></a>
### Nested Schema for `gcp`

Required:

- `environment_name` (String)

Read-Only:

- `availability_zones` (List of String)
- `cascading_delete` (Boolean)
- `credential_name` (String)
- `crn` (String)
- `description` (String)
- `enable_tunnel` (Boolean)
- `encryption_key` (String)
- `endpoint_access_gateway_scheme` (String)
- `existing_network_params` (Attributes) (see [below for nested schema](#nestedatt--gcp--existing_network_params))
- `freeipa` (Attributes) (see [below for nested schema](#nestedatt--gcp--freeipa))
- `id` (String)
- `log_storage` (Attributes) (see [below for nested schema](#nestedatt--gcp--log_storage))
- `polling_options` (Attributes) (see [below for nested schema](#nestedatt--gcp--polling_options))
- `proxy_config_name` (String)
- `public_key` (String)
- `region` (String)
- `report_deployment_logs` (Boolean)
- `security_access` (Attributes) (see [below for nested schema](#nestedatt--gcp--security_access))
- `status` (String)
- `status_reason` (String)
- `tags` (Map of String)
- `use_public_ip` (Boolean)
- `workload_analytics` (Boolean)

<a id="nestedatt--gcp--existing_network_params"></a>
### Nested Schema for `gcp.existing_network_params`

Read-Only:

- `network_name` (String)
- `shared_project_id` (String)
- `subnet_names` (List of String)


<a id="nestedatt--gcp--freeipa"></a>
### Nested Schema for `gcp.freeipa`

Read-Only:

- `catalog` (String)
- `image_id` (String)
- `instance_count_by_group` (Number)
- `instance_type` (String)
- `instances` (Attributes Set) (see [below for nested schema](#nestedatt--gcp--freeipa--instances))
- `multi_az` (Boolean)
- `os` (String)
- `recipes` (Set of String)

<a id="nestedatt--gcp--freeipa--instances"></a>
### Nested Schema for `gcp.freeipa.instances`

Read-Only:

- `availability_zone` (String)
- `discovery_fqdn` (String)
- `instance_group` (String)
- `instance_id` (String)
- `instance_status` (String)
- `instance_status_reason` (String)
- `instance_type` (String)
- `instance_vm_type` (String)
- `life_cycle` (String)
- `private_ip` (String)
- `public_ip` (String)
- `ssh_port` (Number)
- `subnet_id` (String)



<a id="nestedatt--gcp--log_storage"></a>
### Nested Schema for `gcp.log_storage`

Read-Only:

- `backup_storage_location_base` (String)
- `service_account_email` (String)
- `storage_location_base` (String)


<a id="nestedatt--gcp--polling_options"></a>
### Nested Schema for `gcp.polling_options`

Read-Only:

- `async` (Boolean)
- `call_failure_threshold` (Number)
- `polling_timeout` (Number)


<a id="nestedatt--gcp--security_access"></a>
### Nested Schema for `gcp.security_access`

Read-Only:

- `default_security_group_id` (String)
- `security_group_id_for_knox` (String)


