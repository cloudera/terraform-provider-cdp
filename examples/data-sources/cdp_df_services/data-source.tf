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

# List all DataFlow services
data "cdp_df_services" "all" {}

output "all_services" {
  value = data.cdp_df_services.all.services
}

# Filter by name
data "cdp_df_services" "filtered" {
  name = "my-df-service"
}

output "filtered_service_crn" {
  value = data.cdp_df_services.filtered.services[0].crn
}
