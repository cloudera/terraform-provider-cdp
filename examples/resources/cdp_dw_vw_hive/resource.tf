## Copyright 2024 Cloudera. All Rights Reserved.
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
      source = "cloudera/cdp"
    }
  }
}

resource "cdp_dw_vw_hive" "example" {
  cluster_id                      = "env-id"
  database_catalog_id             = "warehouse-id"
  name                            = "default-catalog"
  group_size                      = 2
  platform_jwt_auth               = true
  enable_sso                      = true
  image_version                   = "2024.0.18.4-5"
  min_group_count                 = 1
  max_group_count                 = 3
  disable_auto_suspend            = false
  auto_suspend_timeout_seconds    = 100
  scale_wait_time_seconds         = 230 // either headroom or scale_wait_time_seconds can be configured
  headroom                        = 1
  max_concurrent_isolated_queries = 5
  max_nodes_per_isolated_query    = 2
  aws_options = {
    availability_zone = "us-west-2a"
    ebs_llap_spill_gb = 300
    tags = {
      "key1" = "value1"
    }
  }
}

output "jdbc_url" {
  value = cdp_dw_vw_hive.example.jdbc_url
}

output "kerberos_jdbc_url" {
  value = cdp_dw_vw_hive.example.kerberos_jdbc_url
}

output "hue_url" {
  value = cdp_dw_vw_hive.example.hue_url
}

output "jwt_connection_string" {
  value = cdp_dw_vw_hive.example.jwt_connection_string
}

output "jwt_token_gen_url" {
  value = cdp_dw_vw_hive.example.jwt_token_gen_url
}