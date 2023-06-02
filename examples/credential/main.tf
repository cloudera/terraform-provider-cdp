## Copyright 2023 Cloudera. All Rights Reserved.
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
      source = "terraform.cloudera.com/cloudera/cdp"
    }
  }
}

provider "cdp" {
  cdp_access_key_id           = "<value>"
  cdp_private_key             = "<value>"
  cdp_profile                 = "mowdev"
  cdp_config_file             = "<value>"
  cdp_shared_credentials_file = "<value>"
}

resource "cdp_environments_aws_credential" "my-credentials" {
  credential_name = "<value>"
  role_arn        = "<value>"
  description     = "testing the terrafrom-provider-cdp"
}

output "credential_name" {
  value = cdp_environments_aws_credential.my-credentials.credential_name
}

output "role_arn" {
  value = cdp_environments_aws_credential.my-credentials.role_arn
}

output "crn" {
  value = cdp_environments_aws_credential.my-credentials.crn
}
