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
  default     = "my-gcp-credential"
}

variable "environment_name" {
  description = "Name for the CDP environment."
  type        = string
  default     = "my-gcp-environment"
}

variable "datalake_name" {
  description = "Name for the CDP datalake."
  type        = string
  default     = "my-gcp-datalake"
}

variable "datahub_name" {
  description = "Name for the CDP DataHub cluster."
  type        = string
  default     = "my-gcp-datahub"
}

# --- GCP Identity ---

variable "gcp_credential_key" {
  description = "Base64-encoded GCP service account key JSON. Use filebase64(\"key.json\")."
  type        = string
  sensitive   = true
}

# --- GCP Infrastructure ---

variable "region" {
  description = "GCP region for the environment (e.g. us-central1)."
  type        = string
}

variable "network_name" {
  description = "Name of the existing GCP VPC network."
  type        = string
}

variable "subnet_names" {
  description = "List of subnet names within the VPC network."
  type        = list(string)
}

variable "shared_project_id" {
  description = "GCP project ID that owns the shared VPC (if using shared VPC)."
  type        = string
}

variable "ssh_public_key" {
  description = "SSH public key for cluster node access."
  type        = string
}

variable "use_public_ip" {
  description = "Whether to assign public IPs to environment nodes."
  type        = bool
  default     = false
}

# --- Storage & Identity ---

variable "datalake_service_account_email" {
  description = "GCP service account email for datalake admin (IDBroker)."
  type        = string
}

variable "datalake_storage_location" {
  description = "GCS location for datalake data (e.g. gs://my-bucket/datalake)."
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
  default     = "7.2.18 - Data Engineering for Google Cloud"
}
