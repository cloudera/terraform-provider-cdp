## Copyright 2024 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

resource "cdp_dw_aws_cluster" "example" {
  crn                              = cdp_datalake_aws_datalake.example.crn
  name                             = "<value>"
  cluster_id                       = "<value>"
  node_role_cdw_managed_policy_arn = "<value>"
  database_backup_retention_days   = 7
  custom_registry_options = {
    registry_type  = "ECR"
    repository_url = "<value>"
  }
  custom_subdomain = ""
  network_settings = {
    worker_subnet_ids                     = ["<value>", "<value>", "<value>"]
    load_balancer_subnet_ids              = ["<value>", "<value>", "<value>"]
    use_overlay_network                   = false
    whitelist_k8s_cluster_access_ip_cidrs = ["0.0.0.0/0"]
    whitelist_workload_access_ip_cidrs    = ["0.0.0.0/0"]
    use_private_load_balancer             = true
    use_public_worker_node                = false
    enable_private_eks                    = true
  }
  instance_settings = {
    custom_ami_id             = ""
    enable_spot_instances     = false
    additional_instance_types = ["<value>"]
  }
}

output "crn" {
  value = cdp_dw_aws_cluster.example.crn
}

output "cluster_id" {
  value = cdp_dw_aws_cluster.example.cluster_id
}

output "name" {
  value = cdp_dw_aws_cluster.example.name
}
