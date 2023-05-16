terraform {
  required_providers {
    cdp = {
      source = "terraform.cloudera.com/cloudera/cdp"
    }
  }
}

provider "cdp" {
  cdp_profile                 = "<value>"
  cdp_endpoint_url            = "<value>"
  cdp_config_file             = "<value>"
  cdp_shared_credentials_file = "<value>"
  //local_environment           = true
}

resource "cdp_environments_azure_credential" "my-credentials" {
  credential_name = "example-cdp-azure-credential"
  subscription_id = "<value>"
  tenant_id       = "<value>"
  app_based = {
    application_id  = "<value>"
    secret_key      = "<value>"
  }
  description     = "testing the terrafrom-provider-cdp"
}

output "credential_name" {
  value = cdp_environments_azure_credential.my-credentials.credential_name
}

output "subscription_id" {
  value = cdp_environments_azure_credential.my-credentials.subscription_id
}

output "tenant_id" {
  value = cdp_environments_azure_credential.my-credentials.tenant_id
}

output "application_id" {
  value = cdp_environments_azure_credential.my-credentials.app_based.application_id
}

output "secret_key" {
  value = cdp_environments_azure_credential.my-credentials.app_based.secret_key
}

output "crn" {
  value = cdp_environments_azure_credential.my-credentials.crn
}
