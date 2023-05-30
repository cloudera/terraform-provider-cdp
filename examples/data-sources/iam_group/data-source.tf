terraform {
  required_providers {
    cdp = {
      source = "terraform.cloudera.com/cloudera/cdp"
    }
  }
}

provider "cdp" {
  cdp_profile = "prod-us"
}

data "cdp_iam_group" "example" {
  group_name = "example"
}

output "cdp_iam_group_crn" {
  value = data.cdp_iam_group.example.crn
}