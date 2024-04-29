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

resource "cdp_operational_database" "opdb" {
  environment_name = "<value>"
  database_name    = "<value>"

  // scale_type   = "MICRO" // valid options are "MICRO","LIGHT","HEAVY"
  // storage_type = "HDFS"  // valid options are "CLOUD_WITH_EPHEMERAL","CLOUD","HDFS"

  // num_edge_nodes   = 1
}