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
