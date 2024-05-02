## Copyright 2024 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

resource "cdp_iam_machine_user_resource_role_assignment" "example" {
  machine_user      = "example"
  resource_crn      = "crn:cdp:environments:us-west-1:00000000-0000-0000-0000-000000000000:environment:00000000-0000-0000-0000-000000000000"
  resource_role_crn = "crn:altus:iam:us-west-1:altus:resourceRole:EnvironmentUser"
}

output "machine_use" {
  value = cdp_iam_machine_user_resource_role_assignment.example.machine_user
}

output "resource_role_crn" {
  value = cdp_iam_machine_user_resource_role_assignment.example.resource_role_crn
}
