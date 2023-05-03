# This example is mimimum configuration for CDP provider and creates an aws credential in CDP

provider "cdp" {
}

resource "cdp_environments_aws_credential" "my-credentials" {
  name = "my-cloudera-cdp-credentials"
  role_arn = "arn:aws:iam::11111111111:role/my-cross-account-role"
  description = "Testing the CDP Terraform Provider"
}
