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
      source = "registry.terraform.io/cloudera/cdp"
    }
  }
}

data "cdp_environments_aws_credential_prerequisites" "credential_prerequisites" {}

resource "aws_iam_role" "cdp-cross-account-role" {
  name = "cdp-cross-account-role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::${data.cdp_environments_aws_credential_prerequisites.credential_prerequisites.account_id}:root"
      },
      "Action": "sts:AssumeRole",
      "Condition": {
        "StringEquals": {
          "sts:ExternalId": "${data.cdp_environments_aws_credential_prerequisites.credential_prerequisites.external_id}"
        }
      }
    }
  ]
}

EOF

}

output "account_id" {
  value = data.cdp_environments_aws_credential_prerequisites.credential_prerequisites.account_id
}

output "external_id" {
  value = data.cdp_environments_aws_credential_prerequisites.credential_prerequisites.external_id
}
