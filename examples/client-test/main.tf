terraform {
  required_providers {
    cdp = {
      source = "terraform.cloudera.com/cloudera/cdp"
    }
  }
}

provider "cdp" {
  cdp_profile = "default"
}

data "cdp_environments_aws_credential_prerequisites" "example" {}

output "cdp_environments_aws_credential_prerequisites" {
  value = data.cdp_environments_aws_credential_prerequisites.example
}
