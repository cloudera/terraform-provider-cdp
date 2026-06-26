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

# Look up a single DataFlow service by name
data "cdp_df_service" "example" {
  name = "my-df-service"
}

# Use the service CRN directly
resource "cdp_df_deployment" "example" {
  service_crn      = data.cdp_df_service.example.crn
  flow_version_crn = cdp_df_flow_definition.example.flow_version_crn
  deployment_name  = "my-deployment"
}

output "service_crn" {
  value = data.cdp_df_service.example.crn
}

output "environment_crn" {
  value = data.cdp_df_service.example.environment_crn
}
