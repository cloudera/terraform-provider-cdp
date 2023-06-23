## Copyright 2023 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

# This example is mimimum configuration for CDP provider and creates an aws credential in CDP
terraform {
  required_providers {
    cdp = {
      source = "registry.terraform.io/cloudera/cdp"
    }
  }
}

provider "cdp" {
}

resource "cdp_environments_aws_credential" "example" {
  name        = "example-cdp-aws-credential"
  role_arn    = "arn:aws:iam::11111111111:role/example-cross-account-role"
  description = "Example AWS Credential"
}
