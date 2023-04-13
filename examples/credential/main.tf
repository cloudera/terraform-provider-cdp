terraform {
  required_providers {
    cdp = {
      source = "terraform.cloudera.com/cloudera/cdp"
    }
  }
}

provider "cdp" {
  # cdp_access_key_id = ""
  # cdp_private_key = ""
  # cdp_profile = "dev"
  # cdp_config_file = "/Users/spongebob/.cdp/config"
  # cdp_shared_credentials_file = "/Users/spongebob/.cdp/credentials"
}

resource "cdp_environments_aws_credential" "my-credentials" {
  credential_name = "my-cdp-credentials"
  role_arn        = "arn:aws:iam::111111111111:role/my-cdp-cross-account-role"
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
