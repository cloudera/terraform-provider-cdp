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

# Basic deployment
resource "cdp_df_deployment" "basic" {
  service_crn      = cdp_df_service.example.crn
  flow_version_crn = cdp_df_flow_definition.example.flow_version_crn
  deployment_name  = "my-deployment"
}

# Full deployment with all options
resource "cdp_df_deployment" "full" {
  service_crn      = cdp_df_service.example.crn
  flow_version_crn = cdp_df_flow_definition.example.flow_version_crn
  deployment_name  = "my-full-deployment"

  cluster_size     = "EXTRA_SMALL"
  cfm_nifi_version = "1.21.0"
  auto_start_flow  = false
  project_crn      = cdp_df_project.example.crn

  static_node_count = 1

  # Change flow version strategy (used when flow_version_crn is updated)
  strategy                       = "STOP_AND_PROCESS_DATA"
  wait_for_flow_to_stop_in_minutes = 15

  # Parameter groups as JSON
  parameter_groups = file("parameters.json")

  polling_timeout = 3600
}

# Deployment with auto-scaling
resource "cdp_df_deployment" "autoscaled" {
  service_crn      = cdp_df_service.example.crn
  flow_version_crn = cdp_df_flow_definition.example.flow_version_crn
  deployment_name  = "my-autoscaled-deployment"

  cluster_size         = "SMALL"
  auto_start_flow      = false
  auto_scaling_enabled = true
  auto_scale_min_nodes = 1
  auto_scale_max_nodes = 3
}
