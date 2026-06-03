# Copyright 2025 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

# AWS example
resource "cdp_cloudprivatelinks_private_link_service_access" "aws_example" {
  cloud_service_provider = "AWS"
  cloud_account_id       = "123456789012"
  region                 = "us-east-1"
  service_group          = "CDP-CONTROL-PLANE"
}

# Azure example
resource "cdp_cloudprivatelinks_private_link_service_access" "azure_example" {
  cloud_service_provider = "AZURE"
  subscription_id        = "00000000-0000-0000-0000-000000000000"
  region                 = "eastus"
  service_group          = "CDP-CONTROL-PLANE"
}