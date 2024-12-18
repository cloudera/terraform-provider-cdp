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

resource "cdp_dw_data_visualization" "example" {
  cluster_id    = "env-id"
  name          = "data-visualization"
  image_version = "2024.0.18.4-5"

  resource_template = "default"

  user_groups  = ["ugrp0", "ugrp1"]
  admin_groups = ["admgrp0", "admgrp1"]
}
