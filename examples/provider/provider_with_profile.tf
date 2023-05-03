# This example shows how to use cdp with a custom profile (other than 'default').
# The profile must be defined in the CDP configuration file (default: ~/.cdp/config) and credentials should be available
# under ~/.cdp/credentials.

provider "cdp" {
  cdp_profile                 = "customprofile"
}