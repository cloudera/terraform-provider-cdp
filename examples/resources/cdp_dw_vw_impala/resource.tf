// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

resource "cdp_dw_vw_impala" "test-terraform" {
  cluster_id          = "env-fwkk6k"
  database_catalog_id = "warehouse-1738648902-2tjt"
  image_version       = "2025.0.20.0-16"
  name                = "test-terraform"
  tshirt_size         = "xsmall"

  autoscaling = {
    auto_suspend_timeout_seconds = 360
    disable_auto_suspend         = false
    scale_down_delay_seconds     = 360
    scale_up_delay_seconds       = 40
    max_clusters                 = 6
    min_clusters                 = 4
  }

  aws_options = {
    scratch_space_limit = 634
  }

  ha_settings = {
    high_availability_mode              = "ACTIVE_PASSIVE"
    enable_shutdown_of_coordinator      = false
    shutdown_of_coordinator_delay_secs  = 360
    num_of_active_coordinators          = 2
    enable_catalog_high_availability    = false
    enable_statestore_high_availability = false
  }

  enable_unified_analytics = true

  query_isolation_options = {
    max_queries         = 2
    max_nodes_per_query = 2
  }

  instance_type = "r5d.4xlarge"
  // either use node_count or autoscaling options, not both
  // node_count              = 3
  availability_zone = "us-west-2a"

  platform_jwt_auth = true
  query_log         = true

  tags = [
    {
      key   = "environment"
      value = "mow-dev"
    },
    {
      key   = "team"
      value = "dwx"
    },
  ]

  enable_sso = true

}
