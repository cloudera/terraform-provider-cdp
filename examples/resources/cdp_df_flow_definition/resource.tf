// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

terraform {
  required_providers {
    cdp = {
      source = "cloudera/cdp"
    }
  }
}

# Import a NiFi flow definition into the DataFlow catalog
resource "cdp_df_flow_definition" "example" {
  name        = "my-nifi-flow"
  file        = file("flow.json")
  description = "My NiFi flow definition"
  comments    = "Initial version"
}

# Import a flow definition and assign it to a collection
resource "cdp_df_flow_definition" "with_collection" {
  name           = "my-nifi-flow-in-collection"
  file           = file("flow.json")
  description    = "Flow assigned to a collection"
  comments       = "Initial version"
  collection_crn = cdp_df_collection.example.crn
}

# The flow_version_crn output can be used to create a deployment
output "flow_version_crn" {
  value = cdp_df_flow_definition.example.flow_version_crn
}
