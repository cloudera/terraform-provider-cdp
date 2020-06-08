provider "cdp" {
  # cdp_access_key_id = ""
  # cdp_private_key = ""
  # cdp_profile = "default"
}

resource "cdp_environments_credential" "my-credentials" {
  name = "my-cloudera-cdp-credentials"
  cloud_platform = "AWS"
  role_arn = "arn:aws:iam::643594445606:role/EGRESS_TEST_cdp-cross-account-role"
  description = "testing the terrafrom-provider-cdp"
}

output "name" {
  value = cdp_environments_credential.my-credentials.name
}

output "role_arn" {
  value = cdp_environments_credential.my-credentials.role_arn
}

output "crn" {
  value = cdp_environments_credential.my-credentials.crn
}