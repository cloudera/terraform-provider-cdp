## Copyright 2026 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

resource "cdp_environments_aws_gov_cloud_audit_credential" "example" {
  role_arn = "arn:aws-us-gov:iam::11111111111:role/example-govcloud-audit-role"
}

output "credential_name" {
  value = cdp_environments_aws_gov_cloud_audit_credential.example.credential_name
}

output "crn" {
  value = cdp_environments_aws_gov_cloud_audit_credential.example.crn
}
