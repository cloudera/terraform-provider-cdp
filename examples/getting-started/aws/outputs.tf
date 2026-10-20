## Copyright 2026 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

output "credential_crn" {
  description = "CRN of the CDP credential."
  value       = cdp_environments_aws_credential.cred.crn
}

output "environment_crn" {
  description = "CRN of the CDP environment."
  value       = cdp_environments_aws_environment.env.crn
}

output "environment_status" {
  description = "Status of the CDP environment."
  value       = cdp_environments_aws_environment.env.status
}

output "datalake_crn" {
  description = "CRN of the CDP datalake."
  value       = cdp_datalake_aws_datalake.dl.crn
}

output "datahub_crn" {
  description = "CRN of the DataHub cluster."
  value       = cdp_datahub_aws_cluster.hub.crn
}
