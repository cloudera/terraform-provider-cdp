## Copyright 2024 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

variable "password" {
  type = string
}

resource "cdp_iam_machine_user" "example" {
  name = "example"

  # Optional
  workload_password = var.password
}

output "machine_user" {
  value = cdp_iam_machine_user.example.name
}

output "password_expiration_date" {
  value = cdp_iam_machine_user.example.workload_password_details.expiration_date
}
