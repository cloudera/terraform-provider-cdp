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