## Copyright 2023 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

# This example shows how to use cdp access key id and secret key to manually configure the credentials in the provider
# configuration block.
#
# WARNING:  Hard-coding credentials into any Terraform configuration is NOT
# recommended, and risks secret leakage should this file ever be committed to a
# public version control system.

provider "cdp" {
  cdp_access_key_id = "my-access-key"
  cdp_private_key = "my-private-key"
}