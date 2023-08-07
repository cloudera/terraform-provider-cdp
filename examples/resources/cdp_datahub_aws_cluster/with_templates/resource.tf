// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

resource "cdp_datahub_aws_cluster" "aws-cluster" {
  name               = "<value>"
  environment        = "<value>"
  cluster_template   = "7.2.15 - Data Engineering: Apache Spark, Apache Hive, Apache Oozie"
  cluster_definition = "7.2.15 - Data Engineering for AWS"

  destroy_options = {
    force_delete_cluster = false
  }
}

output "cluster" {
  value = cdp_datahub_aws_cluster.aws-cluster.name
}

output "environment" {
  value = cdp_datahub_aws_cluster.aws-cluster.environment
}

output "cluster_template" {
  value = cdp_datahub_aws_cluster.aws-cluster.cluster_template
}

output "cluster_definition" {
  value = cdp_datahub_aws_cluster.aws-cluster.cluster_definition
}

output "destroy_options" {
  value = cdp_datahub_aws_cluster.aws-cluster.destroy_options
}

output "force_delete_cluster" {
  value = cdp_datahub_aws_cluster.aws-cluster.destroy_options.force_delete_cluster
}