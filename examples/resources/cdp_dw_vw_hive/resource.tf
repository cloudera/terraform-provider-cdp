## Copyright 2024 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

resource "cdp_dw_vw_hive" "example" {
  cluster_id          = "env-id"
  database_catalog_id = "warehouse-id"
  name                = "default-catalog"
  node_count = 2
  platform_jwt_auth = true
  enable_sso = true
  image_version = "2024.0.19.0-301"
  autoscaling = {
    min_clusters = 1
    max_clusters = 3
    disable_auto_suspend = false
    auto_suspend_timeout_seconds = 100
    hive_scale_wait_time_seconds = 230
    hive_desired_free_capacity = 1
  }
  aws_options = {
    availability_zone = "us-west-2a"
    ebs_llap_spill_gb = 300
    tags = {
      "key1" = "value1"
    }
  }
  query_isolation_options = {
    max_queries = 100
    max_nodes_per_query = 10
  }
}
