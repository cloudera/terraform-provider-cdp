## Copyright 2026 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

# Getting Started — GCP
#
# This example creates a complete CDP stack on GCP:
#   Credential → Environment → Datalake → DataHub cluster
#
# Prerequisites:
#   - A CDP access key
#   - A GCP service account key (JSON) for CDP
#   - A VPC network with subnets, a GCS bucket, and service accounts
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
# Register the GCP service account that lets CDP manage GCP resources.

resource "cdp_environments_gcp_credential" "cred" {
  credential_name = var.credential_name
  credential_key  = var.gcp_credential_key
  description     = "CDP credential for ${var.environment_name}"
}

# --- Step 2: Environment ---
# Create the CDP environment with existing GCP networking.

resource "cdp_environments_gcp_environment" "env" {
  environment_name = var.environment_name
  credential_name  = cdp_environments_gcp_credential.cred.credential_name
  region           = var.region

  public_key    = var.ssh_public_key
  use_public_ip = var.use_public_ip

  existing_network_params = {
    network_name      = var.network_name
    subnet_names      = var.subnet_names
    shared_project_id = var.shared_project_id
  }

  timeouts = {
    create = "60m"
    delete = "60m"
  }
}

# --- Step 3: Datalake ---
# Create the shared data layer (Ranger, Atlas, IDBroker).

resource "cdp_datalake_gcp_datalake" "dl" {
  datalake_name    = var.datalake_name
  environment_name = cdp_environments_gcp_environment.env.environment_name

  cloud_provider_configuration = {
    service_account_email = var.datalake_service_account_email
    storage_location      = var.datalake_storage_location
  }

  timeouts = {
    create = "60m"
    delete = "60m"
  }
}

# --- Step 4: DataHub Cluster ---
# Create a DataHub cluster using a built-in template.
# Note: GCP DataHub uses _name suffixed attributes (environment_name,
# cluster_template_name, cluster_definition_name) unlike AWS/Azure.

resource "cdp_datahub_gcp_cluster" "hub" {
  name                    = var.datahub_name
  environment_name        = cdp_environments_gcp_environment.env.environment_name
  cluster_template_name   = var.datahub_template
  cluster_definition_name = var.datahub_definition

  destroy_options = {
    force_delete_cluster = false
  }
}
