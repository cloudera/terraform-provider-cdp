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

resource "cdp_iam_group" "example" {
  group_name = "example"
  sync_membership_on_user_login = true
}

output "cdp_iam_group_example_crn" {
  value = resource.cdp_iam_group.example.crn
}
