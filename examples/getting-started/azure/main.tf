## Copyright 2026 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

# Getting Started — Azure
#
# This example creates a complete CDP stack on Azure:
#   Credential → Environment → Datalake → DataHub cluster
#
# Prerequisites:
#   - A CDP access key
#   - An Azure app registration (service principal) with Contributor role
#   - A resource group, VNet with subnets, ADLS Gen2 storage, and managed identities
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
# Register the Azure app registration that lets CDP manage Azure resources.

resource "cdp_environments_azure_credential" "cred" {
  credential_name = var.credential_name
  subscription_id = var.subscription_id
  tenant_id       = var.tenant_id
  app_based = {
    application_id = var.application_id
    secret_key     = var.application_secret
  }
  description = "CDP credential for ${var.environment_name}"
}

# --- Step 2: Environment ---
# Create the CDP environment with existing Azure networking.

resource "cdp_environments_azure_environment" "env" {
  environment_name = var.environment_name
  credential_name  = cdp_environments_azure_credential.cred.credential_name
  region           = var.region

  security_access = {
    cidr = var.security_access_cidr
  }

  public_key = var.ssh_public_key

  existing_network_params = {
    network_id          = var.vnet_name
    resource_group_name = var.network_resource_group_name
    subnet_ids          = var.subnet_ids
  }

  log_storage = {
    storage_location_base = var.log_storage_location
    managed_identity      = var.log_managed_identity
  }

  resource_group_name = var.resource_group_name
  use_public_ip       = var.use_public_ip

  timeouts = {
    create = "60m"
    delete = "60m"
  }
}

# --- Step 3: Datalake ---
# Create the shared data layer (Ranger, Atlas, IDBroker).

resource "cdp_datalake_azure_datalake" "dl" {
  datalake_name         = var.datalake_name
  environment_name      = cdp_environments_azure_environment.env.environment_name
  managed_identity      = var.datalake_managed_identity
  storage_location_base = var.datalake_storage_location

  timeouts = {
    create = "60m"
    delete = "60m"
  }
}

# --- Step 4: DataHub Cluster ---
# Create a DataHub cluster using a built-in template.

resource "cdp_datahub_azure_cluster" "hub" {
  name               = var.datahub_name
  environment        = cdp_environments_azure_environment.env.environment_name
  cluster_template   = var.datahub_template
  cluster_definition = var.datahub_definition

  destroy_options = {
    force_delete_cluster = false
  }
}
