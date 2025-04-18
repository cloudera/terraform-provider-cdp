## Copyright 2023 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

resource "cdp_environments_aws_credential" "example" {
  credential_name = "example-cdp-aws-credential"
  role_arn        = "arn:aws:iam::11111111111:role/example-cross-account-role"
  description     = "Example AWS Credentials"
}

resource "cdp_environments_aws_environment" "example" {
  environment_name = "example-environment"
  credential_name  = cdp_environments_aws_credential.example.credential_name
  region           = "<your-region>"
  security_access = {
    cidr = "0.0.0.0/0"
  }
  authentication = {
    public_key_id = "my-key"
  }
  log_storage = {
    storage_location_base = "s3a://storage-bucket/location"
    instance_profile      = "arn:aws:iam::11111111111:instance-profile/storage-instance-profile"
  }
  vpc_id     = "vpc-1"
  subnet_ids = ["<env-subnet-1>", "<env-subnet-2>", "<env-subnet-3>"]
  compute_cluster = {
    enabled = false
    configuration = {
      kube_api_authorized_ip_ranges = ["0.0.0.0/0"]
      worker_node_subnets           = ["<env-subnet-1>", "<env-subnet-2>", "<env-subnet-3>"]
    }
  }
}

output "credential" {
  value = cdp_environments_aws_credential.example
}

output "environment" {
  value = cdp_environments_aws_environment.example
}