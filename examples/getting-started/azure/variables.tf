## Copyright 2026 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

# --- Naming ---

variable "credential_name" {
  description = "Name for the CDP credential."
  type        = string
  default     = "my-azure-credential"
}

variable "environment_name" {
  description = "Name for the CDP environment."
  type        = string
  default     = "my-azure-environment"
}

variable "datalake_name" {
  description = "Name for the CDP datalake."
  type        = string
  default     = "my-azure-datalake"
}

variable "datahub_name" {
  description = "Name for the CDP DataHub cluster."
  type        = string
  default     = "my-azure-datahub"
}

# --- Azure Identity ---

variable "subscription_id" {
  description = "Azure subscription ID."
  type        = string
}

variable "tenant_id" {
  description = "Azure Active Directory tenant ID."
  type        = string
}

variable "application_id" {
  description = "Azure app registration (service principal) application ID."
  type        = string
}

variable "application_secret" {
  description = "Azure app registration secret key."
  type        = string
  sensitive   = true
}

# --- Azure Infrastructure ---

variable "region" {
  description = "Azure region for the environment (e.g. westus2)."
  type        = string
}

variable "resource_group_name" {
  description = "Azure resource group for CDP resources."
  type        = string
}

variable "vnet_name" {
  description = "Name of the existing Azure VNet."
  type        = string
}

variable "network_resource_group_name" {
  description = "Resource group containing the VNet."
  type        = string
}

variable "subnet_ids" {
  description = "List of subnet IDs within the VNet."
  type        = list(string)
}

variable "security_access_cidr" {
  description = "CIDR range for security group access."
  type        = string
  default     = "10.0.0.0/16"
}

variable "ssh_public_key" {
  description = "SSH public key for cluster node access."
  type        = string
}

variable "use_public_ip" {
  description = "Whether to assign public IPs to environment nodes."
  type        = bool
  default     = true
}

# --- Storage & Identity ---

variable "log_storage_location" {
  description = "ADLS Gen2 location for CDP logs (e.g. abfs://fs@storage.dfs.core.windows.net)."
  type        = string
}

variable "log_managed_identity" {
  description = "Azure managed identity resource ID for log storage access."
  type        = string
}

variable "datalake_managed_identity" {
  description = "Azure managed identity resource ID for datalake admin (IDBroker)."
  type        = string
}

variable "datalake_storage_location" {
  description = "ADLS Gen2 location for datalake data."
  type        = string
}

# --- DataHub ---

variable "datahub_template" {
  description = "Cluster template name for the DataHub."
  type        = string
  default     = "7.2.18 - Data Engineering: Apache Spark, Apache Hive, Apache Oozie"
}

variable "datahub_definition" {
  description = "Cluster definition name for the DataHub."
  type        = string
  default     = "7.2.18 - Data Engineering for Azure"
}
