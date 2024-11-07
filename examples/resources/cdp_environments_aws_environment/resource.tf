## Copyright 2023 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

resource "cdp_environments_aws_credential" "example" {
  credential_name = "example-cdp-aws-credential"
  role_arn        = "arn:aws:iam::11111111111:role/example-cross-account-role"
  description     = "Example AWS Credentials"
}

resource "cdp_environments_aws_environment" "example" {
  environment_name = "example-environment"
  credential_name  = cdp_environments_aws_credential.example.credential_name
  region           = "us-west"
  security_access = {
    cidr = "0.0.0.0/0"
  }
  network_cidr = "10.10.0.0/16"
  authentication = {
    public_key_id = "my-key"
  }
  log_storage = {
    storage_location_base = "s3a://storage-bucket/location"
    instance_profile      = "arn:aws:iam::11111111111:instance-profile/storage-instance-profile"
  }
}

output "environment_name" {
  value = cdp_environments_aws_environment.example.environment_name
}

output "crn" {
  value = cdp_environments_aws_environment.example.crn
}

output "status" {
  value = cdp_environments_aws_environment.example.status
}

output "status_reason" {
  value = cdp_environments_aws_environment.example.status_reason
}

output "cloud_platform" {
  value = cdp_environments_aws_environment.example.region
}

output "security_access" {
  value = cdp_environments_aws_environment.example.security_access
}

output "network_cidr" {
  value = cdp_environments_aws_environment.example.network_cidr
}

output "authentication" {
  value = cdp_environments_aws_environment.example.authentication
}

output "log_storage" {
  value = cdp_environments_aws_environment.example.log_storage
}

output "proxy_config_name" {
  value = cdp_environments_aws_environment.example.proxy_config_name
}

output "tags" {
  value = cdp_environments_aws_environment.example.tags
}

output "create_private_subnets" {
  value = cdp_environments_aws_environment.example.create_private_subnets
}

output "create_service_endpoints" {
  value = cdp_environments_aws_environment.example.create_service_endpoints
}

output "s3_guard_table_name" {
  value = cdp_environments_aws_environment.example.s3_guard_table_name
}

output "credential_name" {
  value = cdp_environments_aws_environment.example.credential_name
}

output "description" {
  value = cdp_environments_aws_environment.example.description
}

output "enable_tunnel" {
  value = cdp_environments_aws_environment.example.enable_tunnel
}

output "encryption_key_arn" {
  value = cdp_environments_aws_environment.example.encryption_key_arn
}

output "endpoint_access_gateway_scheme" {
  value = cdp_environments_aws_environment.example.endpoint_access_gateway_scheme
}

output "endpoint_access_gateway_subnet_ids" {
  value = cdp_environments_aws_environment.example.endpoint_access_gateway_subnet_ids
}

output "freeipa" {
  value = cdp_environments_aws_environment.example.freeipa
}

output "report_deployment_logs" {
  value = cdp_environments_aws_environment.example.report_deployment_logs
}

output "subnet_ids" {
  value = cdp_environments_aws_environment.example.subnet_ids
}

output "tunnel_type" {
  value = cdp_environments_aws_environment.example.tunnel_type
}

output "workload_analytics" {
  value = cdp_environments_aws_environment.example.workload_analytics
}

output "vpc_id" {
  value = cdp_environments_aws_environment.example.vpc_id
}
