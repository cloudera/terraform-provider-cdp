## Copyright 2024 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

resource "cdp_ml_workspace" "example" {
  environment_name = var.environment_name
  workspace_name   = "example-ws"
  provision_k8s_request = {
    instance_groups = [
      {
        instance_type = var.instance_type
        root_volume = {
          size = 128
        }
        autoscaling = {
          min_instances = 1
          max_instances = 5
        }
      }
    ],
    environment_name = var.environment_name
  }

  disable_tls              = false
  use_public_load_balancer = false
  private_cluster          = false

  enable_monitoring    = true
  enable_governance    = false
  enable_model_metrics = true

  whitelist_authorized_ip_ranges = false

  skip_validation = false
}
