// Copyright 2024 Cloudera. All Rights Reserved.
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

resource "cdp_datalake_azure_datalake" "example" {
  datalake_name         = "<value>"
  environment_name      = "<value>"
  managed_identity      = "<value>"
  storage_location_base = "<value>"
}

output "name" {
  value = cdp_datalake_azure_datalake.example.datalake_name
}

output "environment" {
  value = cdp_datalake_azure_datalake.example.environment_name
}

output "instance_profile" {
  value = cdp_datalake_azure_datalake.example.managed_identity
}

output "storage_bucket_location" {
  value = cdp_datalake_azure_datalake.example.storage_location_base
}

output "recipes" {
  value = cdp_datalake_azure_datalake.example.recipes
}