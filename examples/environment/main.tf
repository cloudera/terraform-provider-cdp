## Copyright 2023 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

terraform {
  required_providers {
    cdp = {
      source = "terraform.cloudera.com/cloudera/cdp"
    }
  }
}

provider "cdp" {
  # cdp_profile                 = "default"
  # cdp_config_file             = "/Users/spongebob/.cdp/config"
  # cdp_shared_credentials_file = "/Users/spongebob/.cdp/credentials"
}

resource "cdp_environments_aws_environment" "my-environment" {
  environment_name = "my-environment"
  credential_name  = "my-credential"
  region           = "us-west"
  security_access = {
    cidr = "0.0.0.0/0"
  }
  network_cidr = "10.10.0.0/16"
  authentication = {
    public_key_id = "my-key"
  }
  log_storage = {
    storage_location_base = "s3a://storage-location/bucket"
    instance_profile      = "arn:aws:iam::000000000000:instance-profile/my-instance-profile "
  }
}

output "environment_name" {
  value = cdp_environments_aws_environment.my-environment.environment_name
}

output "crn" {
  value = cdp_environments_aws_environment.my-environment.crn
}
