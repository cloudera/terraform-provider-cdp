## Copyright 2023 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

resource "cdp_environments_azure_credential" "example" {
  credential_name = "example-cdp-azure-credential"
  subscription_id = "<value>"
  tenant_id       = "<value>"
  app_based = {
    application_id = "<value>"
    secret_key     = "<value>"
  }
  description = "Example Azure Credential"
}

output "credential_name" {
  value = cdp_environments_azure_credential.example.credential_name
}

output "subscription_id" {
  value = cdp_environments_azure_credential.example.subscription_id
}

output "tenant_id" {
  value = cdp_environments_azure_credential.example.tenant_id
}

output "application_id" {
  value = cdp_environments_azure_credential.example.app_based.application_id
}

output "secret_key" {
  value = cdp_environments_azure_credential.example.app_based.secret_key
}

output "crn" {
  value = cdp_environments_azure_credential.example.crn
}
