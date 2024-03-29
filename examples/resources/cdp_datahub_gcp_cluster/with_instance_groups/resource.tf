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
  name                  = "<value>"
  environment_name      = "<value>"
  cluster_template_name = "7.2.17 - Data Engineering: Apache Spark, Apache Hive, Apache Oozie"

  destroy_options = {
    force_delete_cluster = false
  }

  instance_group = [
    {
      node_count          = 0
      instance_group_name = "<value>"
      instance_group_type = "<value>"
      instance_type       = "<value>"
      root_volume_size    = 100
      attached_volume_configuration = [
        {
          volume_size  = 100
          volume_count = 1
          volume_type  = "<value>"
        }
      ]
      recovery_mode = "<value>"
      recipes       = []
    },
    {
      node_count          = 1
      instance_group_name = "<value>"
      instance_group_type = "<value>"
      instance_type       = "<value>"
      root_volume_size    = 100
      attached_volume_configuration = [
        {
          volume_size  = 100
          volume_count = 1
          volume_type  = "<value>"
        }
      ]
      recovery_mode = "<value>"
      recipes       = []
    },
    {
      node_count          = 3
      instance_group_name = "<value>"
      instance_group_type = "<value>"
      instance_type       = "<value>"
      root_volume_size    = 100
      attached_volume_configuration = [
        {
          volume_size  = 300
          volume_count = 1
          volume_type  = "<value>"
        }
      ]
      recovery_mode = "<value>"
      recipes       = []
    },
    {
      node_count          = 3
      instance_group_name = "<value>"
      instance_group_type = "<value>"
      instance_type       = "<value>"
      root_volume_size    = 100
      attached_volume_configuration = [
        {
          volume_size  = 300
          volume_count = 1
          volume_type  = "<value>"
        }
      ]
      recovery_mode = "<value>"
      recipes       = []
    }
  ]
}

output "cluster" {
  value = cdp_datahub_gcp_cluster.gcp-cluster.name
}

output "environment" {
  value = cdp_datahub_gcp_cluster.gcp-cluster.environment_name
}

output "cluster_template" {
  value = cdp_datahub_gcp_cluster.gcp-cluster.cluster_template_name
}

output "destroy_options" {
  value = cdp_datahub_gcp_cluster.gcp-cluster.destroy_options
}

output "force_delete_cluster" {
  value = cdp_datahub_gcp_cluster.gcp-cluster.destroy_options.force_delete_cluster
}

output "recipes" {
  value = cdp_datahub_gcp_cluster.gcp-cluster.instance_group[*].recipes
}

output "instance_group" {
  value = cdp_datahub_gcp_cluster.gcp-cluster.instance_group
}

output "node_count" {
  value = cdp_datahub_gcp_cluster.gcp-cluster.instance_group[*].node_count
}

output "instance_group_name" {
  value = cdp_datahub_gcp_cluster.gcp-cluster.instance_group[*].instance_group_name
}

output "instance_group_type" {
  value = cdp_datahub_gcp_cluster.gcp-cluster.instance_group[*].instance_group_type
}

output "instance_type" {
  value = cdp_datahub_gcp_cluster.gcp-cluster.instance_group[*].instance_type
}

output "root_volume_size" {
  value = cdp_datahub_gcp_cluster.gcp-cluster.instance_group[*].root_volume_size
}

output "attached_volume_configuration" {
  value = cdp_datahub_gcp_cluster.gcp-cluster.instance_group[*].attached_volume_configuration
}

output "volume_size" {
  value = cdp_datahub_gcp_cluster.gcp-cluster.instance_group[*].attached_volume_configuration[*].volume_size
}

output "volume_count" {
  value = cdp_datahub_gcp_cluster.gcp-cluster.instance_group[*].attached_volume_configuration[*].volume_count
}

output "volume_type" {
  value = cdp_datahub_gcp_cluster.gcp-cluster.instance_group[*].attached_volume_configuration[*].volume_type
}

output "recovery_mode" {
  value = cdp_datahub_gcp_cluster.gcp-cluster.instance_group[*].recovery_mode
}