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
  default     = "my-aws-credential"
}

variable "environment_name" {
  description = "Name for the CDP environment."
  type        = string
  default     = "my-aws-environment"
}

variable "datalake_name" {
  description = "Name for the CDP datalake."
  type        = string
  default     = "my-aws-datalake"
}

variable "datahub_name" {
  description = "Name for the CDP DataHub cluster."
  type        = string
  default     = "my-aws-datahub"
}

# --- AWS Infrastructure ---

variable "cross_account_role_arn" {
  description = "ARN of the AWS cross-account role for CDP."
  type        = string
}

variable "region" {
  description = "AWS region for the environment (e.g. us-west-2)."
  type        = string
}

variable "vpc_id" {
  description = "ID of the VPC to deploy into."
  type        = string
}

variable "subnet_ids" {
  description = "List of subnet IDs within the VPC."
  type        = list(string)
}

variable "security_access_cidr" {
  description = "CIDR range for security group access."
  type        = string
  default     = "10.0.0.0/16"
}

variable "ssh_public_key_id" {
  description = "ID of the SSH public key registered in CDP."
  type        = string
}

# --- Storage & IAM ---

variable "log_storage_location" {
  description = "S3 location for CDP logs (e.g. s3a://my-bucket/logs)."
  type        = string
}

variable "log_instance_profile_arn" {
  description = "ARN of the IAM instance profile for log storage access."
  type        = string
}

variable "datalake_instance_profile_arn" {
  description = "ARN of the IAM instance profile for datalake admin (IDBroker)."
  type        = string
}

variable "datalake_storage_location" {
  description = "S3 location for datalake data (e.g. s3a://my-bucket/datalake)."
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
  default     = "7.2.18 - Data Engineering for AWS"
}
