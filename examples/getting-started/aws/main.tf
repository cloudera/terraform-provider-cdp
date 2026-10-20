## Copyright 2026 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

# Getting Started — AWS
#
# This example creates a complete CDP stack on AWS:
#   Credential → Environment → Datalake → DataHub cluster
#
# Prerequisites:
#   - A CDP access key (see: https://docs.cloudera.com/cdp-public-cloud/cloud/cli/topics/mc-cli-generating-an-api-access-key.html)
#   - An AWS cross-account role ARN for CDP
#   - A VPC with subnets, an S3 bucket for logs/data, and IAM instance profiles
#
# Usage:
#   terraform init
#   terraform plan -var-file="my-values.tfvars"
#   terraform apply -var-file="my-values.tfvars"

terraform {
  required_providers {
    cdp = {
      source = "cloudera/cdp"
    }
  }
}

provider "cdp" {}

# --- Step 1: Cloud Credential ---
# Register the cross-account IAM role that lets CDP manage AWS resources.

resource "cdp_environments_aws_credential" "cred" {
  credential_name = var.credential_name
  role_arn        = var.cross_account_role_arn
  description     = "CDP cross-account credential for ${var.environment_name}"
}

# --- Step 2: Environment ---
# Create the CDP environment: networking, logging, security, and FreeIPA.

resource "cdp_environments_aws_environment" "env" {
  environment_name = var.environment_name
  credential_name  = cdp_environments_aws_credential.cred.credential_name
  region           = var.region

  security_access = {
    cidr = var.security_access_cidr
  }

  authentication = {
    public_key_id = var.ssh_public_key_id
  }

  log_storage = {
    storage_location_base = var.log_storage_location
    instance_profile      = var.log_instance_profile_arn
  }

  vpc_id     = var.vpc_id
  subnet_ids = var.subnet_ids

  timeouts = {
    create = "60m"
    delete = "60m"
  }
}

# --- Step 3: Datalake ---
# Create the shared data layer (Ranger, Atlas, IDBroker).

resource "cdp_datalake_aws_datalake" "dl" {
  datalake_name         = var.datalake_name
  environment_name      = cdp_environments_aws_environment.env.environment_name
  instance_profile      = var.datalake_instance_profile_arn
  storage_location_base = var.datalake_storage_location

  timeouts = {
    create = "60m"
    delete = "60m"
  }
}

# --- Step 4: DataHub Cluster ---
# Create a DataHub cluster using a built-in template.

resource "cdp_datahub_aws_cluster" "hub" {
  name               = var.datahub_name
  environment        = cdp_environments_aws_environment.env.environment_name
  cluster_template   = var.datahub_template
  cluster_definition = var.datahub_definition

  destroy_options = {
    force_delete_cluster = false
  }
}
