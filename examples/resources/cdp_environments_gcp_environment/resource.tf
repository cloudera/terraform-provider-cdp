## Copyright 2023 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

terraform {
  required_providers {
    cdp = {
      source = "registry.terraform.io/cloudera/cdp"
    }
  }
}

resource "cdp_environments_gcp_environment" "example" {
  environment_name        = "<value>"
  credential_name         = "<value>"
  region                  = "<value>"
  public_key              = "<value>"
  use_public_ip           = false // or true, depends on the requirements/configurations
  existing_network_params = {
    network_name      = "<value>"
    subnet_names      = ["<value>", "<value2>", "..."] // one or more entries accepted
    shared_project_id = "<value>"
  }
}

output "environment_name" {
  value = cdp_environments_gcp_environment.example.environment_name
}

output "credential_name" {
  value = cdp_environments_gcp_environment.example.credential_name
}

output "region" {
  value = cdp_environments_gcp_environment.example.region
}

output "public_key" {
  value = cdp_environments_gcp_environment.example.public_key
}

output "use_public_ip" {
  value = cdp_environments_gcp_environment.example.use_public_ip
}

output "existing_network_params" {
  value = cdp_environments_gcp_environment.example.existing_network_params
}

output "network_name" {
  value = cdp_environments_gcp_environment.example.existing_network_params.network_name
}

output "subnet_names" {
  value = cdp_environments_gcp_environment.example.existing_network_params.subnet_names
}

output "shared_project_id" {
  value = cdp_environments_gcp_environment.example.existing_network_params.shared_project_id
}