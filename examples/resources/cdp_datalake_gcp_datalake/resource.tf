// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

resource "cdp_datalake_gcp_datalake" "example" {
  datalake_name    = "<value>"
  environment_name = "<value>"
  cloud_provider_configuration = {
    service_account_email = "<value>"
    storage_location      = "<value>"
  }
}

output "name" {
  value = cdp_datalake_gcp_datalake.example.datalake_name
}

output "environment" {
  value = cdp_datalake_gcp_datalake.example.environment_name
}

output "service_account_email" {
  value = cdp_datalake_gcp_datalake.example.cloud_provider_configuration.service_account_email
}

output "storage_location" {
  value = cdp_datalake_gcp_datalake.example.cloud_provider_configuration.storage_location
}