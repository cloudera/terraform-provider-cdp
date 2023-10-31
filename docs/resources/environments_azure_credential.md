---
page_title: "cdp_environments_azure_credential Resource - terraform-provider-cdp"
subcategory: "environments"
description: |-
  The Azure credential is used for authorization  to provision resources such as compute instances within your cloud provider account.
---

# cdp_environments_azure_credential (Resource)

The Azure credential is used for authorization  to provision resources such as compute instances within your cloud provider account.

## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_based` (Attributes) (see [below for nested schema](#nestedatt--app_based))
- `credential_name` (String) The name of the CDP credential.
- `subscription_id` (String) The Azure subscription ID. Required for secret based credentials and should look like the following example: a8d4457d-310v-41p6-sc53-14g8d733e514
- `tenant_id` (String) The Azure AD tenant ID for the Azure subscription. Required for secret based credentials and should look like the following example: b10u3481-2451-10ba-7sfd-9o2d1v60185d

### Optional

- `description` (String) A description for the credential.

### Read-Only

- `crn` (String) The CRN of the credential.
- `id` (String) The ID of this resource.

<a id="nestedatt--app_based"></a>
### Nested Schema for `app_based`

Required:

- `application_id` (String) The ID of the application registered in Azure.
- `secret_key` (String, Sensitive) The client secret key (also referred to as application password) for the registered application.