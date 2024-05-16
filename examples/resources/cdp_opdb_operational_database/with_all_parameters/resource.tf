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

provider "cdp" {
  cdp_config_file             = "/Users/<value>/.cdp/config"
  cdp_shared_credentials_file = "/Users/<value>/.cdp/credentials"
}

resource "cdp_opdb_operational_database" "opdb-detailed-example" {
  environment_name = "<value>"
  database_name    = "<value>"

  scale_type   = "HEAVY"                // valid options are "MICRO","LIGHT","HEAVY"
  storage_type = "CLOUD_WITH_EPHEMERAL" // valid options are "CLOUD_WITH_EPHEMERAL","CLOUD","HDFS"

  java_version = 8

  disable_external_db = true

  disable_multi_az = true
  subnet_id        = "<value>"

  num_edge_nodes = 1

  auto_scaling_parameters = {
    targeted_value_for_metric = 249
    max_workers_for_database  = 50
    max_workers_per_batch     = 4
    min_workers_for_database  = 15
    evaluation_period         = 2400
    minimum_block_cache_gb    = 1
    # beta only
    max_cpu_utilization            = -1
    max_compute_nodes_for_database = -1
    min_compute_nodes_for_database = -1
    max_hdfs_usage_percentage      = 80
    max_regions_per_region_server  = 200
  }

  attached_storage_for_workers = {
    volume_count = 3 // min 1 max 8
    volume_size  = 1024
    volume_type  = "SSD" // valid options are "HDD", "SSD", "LOCAL_SSD"
  }

  image = {
    id      = "<value>"
    catalog = "<value>"
  }

  disable_kerberos = true
  disable_jwt_auth = true

  enable_grafana = true

  custom_user_tags = [
    {
      key   = "key1"
      value = "value1"
    },
    {
      key   = "key2"
      value = "value2"
    },
    {
      key   = "key3"
      value = "value3"
    }
  ]

  enable_region_canary = true

  recipes = [
    {
      names          = ["<value>"],
      instance_group = "<value>"
    },
    {
      names          = ["<value>", "<value>"],
      instance_group = "<value>"
    }
  ]
  storage_location = "s3a://<value>/"

  volume_encryptions = [
    {
      encryption_key = "<value>",
      instance_group = "<value>"
    },
    {
      encryption_key = "<value>",
      instance_group = "<value>"
    }
  ]

}