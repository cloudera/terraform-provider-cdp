// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

resource "cdp_datahub_gcp_cluster" "gcp-cluster" {
  name                    = "<value>"
  environment_name        = "<value>"
  cluster_template_name   = "7.2.17 - Data Engineering: Apache Spark, Apache Hive, Apache Oozie"
  cluster_definition_name = "7.2.17 - Data Engineering for Google Cloud"

  destroy_options = {
    force_delete_cluster = false
  }
}

output "cluster" {
  value = cdp_datahub_gcp_cluster.gcp-cluster.name
}

output "environment_name" {
  value = cdp_datahub_gcp_cluster.gcp-cluster.environment_name
}

output "cluster_template_name" {
  value = cdp_datahub_gcp_cluster.gcp-cluster.cluster_template_name
}

output "cluster_definition_name" {
  value = cdp_datahub_gcp_cluster.gcp-cluster.cluster_definition_name
}

output "destroy_options" {
  value = cdp_datahub_gcp_cluster.gcp-cluster.destroy_options
}

output "force_delete_cluster" {
  value = cdp_datahub_gcp_cluster.gcp-cluster.destroy_options.force_delete_cluster
}