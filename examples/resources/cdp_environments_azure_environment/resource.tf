## Copyright 2023 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

resource "cdp_environments_azure_credential" "example-cred" {
  credential_name = "example-cdp-azure-credential"
  subscription_id = "<value>"
  tenant_id       = "<value>"
  app_based = {
    application_id = "<value>"
    secret_key     = "<value>"
  }
  description = "Example Azure Credential"
}

resource "cdp_environments_azure_environment" "example-env" {
  environment_name = "example-cdp-azure-environment"
  credential_name  = cdp_environments_azure_credential.example-cred.credential_name
  region           = "us-west"
  security_access = {
    cidr = "0.0.0.0/0"
  }
  existing_network_params = {
    network_id          = "network-name"
    resource_group_name = "rg-name"
    subnet_ids          = ["subnet.id"]
  }
  public_key = "my-key"
  log_storage = {
    storage_location_base = "abfs://rgname-fs@rgname-storage.dfs.core.windows.net"
    managed_identity      = "/subscriptions/123e4567-e89b-12d3-a456-426614174000/resourcegroups/my-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/logger"
  }
  resource_group_name              = "rg-name"
  encryption_user_managed_identity = "some-identity"
  use_public_ip                    = true
}

output "environment_name" {
  value = cdp_environments_azure_environment.example-env.environment_name
}

output "crn" {
  value = cdp_environments_azure_environment.example-env.crn
}
