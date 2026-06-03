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
resource "cdp_cloudprivatelinks_private_link_endpoint" "aws_example" {
  cloud_service_provider = "AWS"
  service_group          = "CDP-CONTROL-PLANE"
  enable_private_dns     = true

  aws_account_details = {
    credential_crn = "crn:cdp:iam:us-west-1:abc123:credential:my-aws-credential"
    region         = "us-east-1"
    vpc_id         = "vpc-0abc123def456789"
    subnet_ids     = ["subnet-0abc123def456789", "subnet-0def456abc123789"]
  }
}

# Azure example
resource "cdp_cloudprivatelinks_private_link_endpoint" "azure_example" {
  cloud_service_provider = "AZURE"
  service_group          = "CDP-CONTROL-PLANE"
  enable_private_dns     = true

  azure_account_details = {
    credential_crn  = "crn:cdp:iam:us-west-1:abc123:credential:my-azure-credential"
    subscription_id = "00000000-0000-0000-0000-000000000000"
    resource_group  = "my-resource-group"
    location        = "eastus"
    vnet_id         = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/my-rg/providers/Microsoft.Network/virtualNetworks/my-vnet"
    subnet_id       = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/my-rg/providers/Microsoft.Network/virtualNetworks/my-vnet/subnets/my-subnet"
  }
}