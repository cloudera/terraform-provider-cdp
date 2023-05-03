# This example shows how to use cdp access key id and secret key to manually configure the credentials in the provider
# configuration block.
#
# WARNING:  Hard-coding credentials into any Terraform configuration is NOT
# recommended, and risks secret leakage should this file ever be committed to a
# public version control system.

provider "cdp" {
  cdp_access_key_id = "my-access-key"
  cdp_private_key = "my-private-key"
}