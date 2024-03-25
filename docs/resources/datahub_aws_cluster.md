---
page_title: "cdp_datahub_aws_cluster Resource - terraform-provider-cdp"
subcategory: "datahub"
description: |-
  Creates an AWS Data hub cluster.
---

# cdp_datahub_aws_cluster (Resource)

Creates an AWS Data hub cluster.

## Example Usage
### Datahub with existing Templates and cluster definitions
DataHub's can be created with existing [Cluster Templates](https://docs.cloudera.com/data-hub/cloud/cluster-templates/topics/mc-templates.html) and [Cluster Definitions](https://docs.cloudera.com/data-hub/cloud/cluster-definitions/topics/dh-cluster-definitions.html) by providing the name of the template and the name of the Cluster Definition in the resource configuration.
Below example uses the template named `7.2.15 - Data Engineering: Apache Spark, Apache Hive, Apache Oozie` and Cluster definition called `7.2.15 - Data Engineering for AWS`. More information about templates and cluster definitions can be found in [CDP Datahub documentation](https://docs.cloudera.com/data-hub/cloud/cluster-templates/topics/dh-default-cluster-definitions.html).
```terraform
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
```

### Datahub from custom cluster definition and InstanceGroup configuration
DataHub's can be created with a custom cluster definition and InstanceGroup configuration by providing `instance_group` configuration in the resource configuration.

```terraform
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
  name             = "<value>"
  environment      = "<value>"
  cluster_template = "7.2.15 - Data Engineering: Apache Spark, Apache Hive, Apache Oozie"

  destroy_options = {
    force_delete_cluster = false
  }

  instance_group = [
    {
      node_count          = 0
      instance_group_name = "gateway"
      instance_group_type = "CORE"
      instance_type       = "m5.2xlarge"
      root_volume_size    = 100
      attached_volume_configuration = [
        {
          volume_size  = 100
          volume_count = 1
          volume_type  = "gp3"
        }
      ]
      recovery_mode = "MANUAL"
      volume_encryption = {
        encryption = false
      }
      recipes = []
    },
    {
      node_count          = 1
      instance_group_name = "master"
      instance_group_type = "GATEWAY"
      instance_type       = "m5.4xlarge"
      root_volume_size    = 100
      attached_volume_configuration = [
        {
          volume_size  = 100
          volume_count = 1
          volume_type  = "gp3"
        }
      ]
      recovery_mode = "MANUAL"
      volume_encryption = {
        encryption = false
      }
      recipes = []
    },
    {
      node_count          = 3
      instance_group_name = "worker"
      instance_group_type = "CORE"
      instance_type       = "r5d.2xlarge"
      root_volume_size    = 100
      attached_volume_configuration = [
        {
          volume_size  = 300
          volume_count = 1
          volume_type  = "gp3"
        }
      ]
      recovery_mode = "MANUAL"
      volume_encryption = {
        encryption = false
      }
      recipes = []
    },
    {
      node_count          = 3
      instance_group_name = "compute"
      instance_group_type = "CORE"
      instance_type       = "r5d.2xlarge"
      root_volume_size    = 100
      attached_volume_configuration = [
        {
          volume_size  = 300
          volume_count = 1
          volume_type  = "ephemeral"
        }
      ]
      recovery_mode = "MANUAL"
      volume_encryption = {
        encryption = false
      }
      recipes = []
    }
  ]
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

output "recipes" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].recipes
}

output "instance_group" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group
}

output "node_count" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].node_count
}

output "instance_group_name" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].instance_group_name
}

output "instance_group_type" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].instance_group_type
}

output "instance_type" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].instance_type
}

output "root_volume_size" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].root_volume_size
}

output "attached_volume_configuration" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].attached_volume_configuration
}

output "volume_size" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].attached_volume_configuration[*].volume_size
}

output "volume_count" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].attached_volume_configuration[*].volume_count
}

output "volume_type" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].attached_volume_configuration[*].volume_type
}

output "recovery_mode" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].recovery_mode
}

output "volume_encryption" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].volume_encryption
}

output "encryption" {
  value = cdp_datahub_aws_cluster.aws-cluster.instance_group[*].volume_encryption.encryption
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `environment` (String) The name of the environment where the cluster will belong to.
- `name` (String) The name of the cluster.

### Optional

- `cluster_definition` (String) The name of the cluster definition.
- `cluster_extension` (Attributes) (see [below for nested schema](#nestedatt--cluster_extension))
- `cluster_template` (String) The name of the cluster template.
- `custom_configurations_name` (String) The name of the custom configurations to use for cluster creation.
- `datahub_database` (String) Database type for datahub. Currently supported values: NONE, NON_HA, HA
- `destroy_options` (Attributes) Cluster deletion options. (see [below for nested schema](#nestedatt--destroy_options))
- `enable_load_balancer` (Boolean) Flag that decides whether to provision a load-balancer to front var- ious service endpoints for the given datahub. This will typically be used for HA cluster shapes.
- `image` (Attributes) (see [below for nested schema](#nestedatt--image))
- `instance_group` (Attributes List) (see [below for nested schema](#nestedatt--instance_group))
- `java_version` (Number) Configure the major version of Java on the cluster.
- `multi_az` (Boolean) Flag  that toggles the multi availability zone for the given datahub cluster when you are not sure what subnet IDs can be used. This  way the subnet IDs will be used what the environment suggests.
- `polling_options` (Attributes) Polling related configuration options that could specify various values that will be used during CDP resource creation. (see [below for nested schema](#nestedatt--polling_options))
- `request_template` (String) JSON  template  to  use for cluster creation. This is different from cluster template and would be removed in the future.
- `subnet_id` (String) The subnet id.
- `subnet_ids` (Set of String) The subnet ids.
- `tags` (Map of String)

### Read-Only

- `crn` (String) The CRN of the cluster.
- `id` (String) The ID of this resource.
- `status` (String) The last known state of the cluster

<a id="nestedatt--cluster_extension"></a>
### Nested Schema for `cluster_extension`

Optional:

- `custom_properties` (String)


<a id="nestedatt--destroy_options"></a>
### Nested Schema for `destroy_options`

Optional:

- `force_delete_cluster` (Boolean) An indicator that will take place once the cluster termination will be performed. If it is true, that means if something would go sideways during termination, the operation will proceed, however in such a case no notification would come thus it is advisable to check the cloud provider if there are no leftover resources once the destroy is finished.


<a id="nestedatt--image"></a>
### Nested Schema for `image`

Required:

- `id` (String)

Optional:

- `catalog` (String)
- `os` (String)


<a id="nestedatt--instance_group"></a>
### Nested Schema for `instance_group`

Required:

- `attached_volume_configuration` (Attributes List) Configuration regarding the attached volume to the specific instance group. (see [below for nested schema](#nestedatt--instance_group--attached_volume_configuration))
- `instance_group_name` (String) The name of the instance group.
- `instance_group_type` (String) The type of the instance group.
- `instance_type` (String) The cloud provider-side instance type.
- `node_count` (Number) The cluster node count. Has to be greater or equal than 0 and less than 100,000.
- `recovery_mode` (String) The type of the recovery mode.
- `root_volume_size` (Number) The size of the root volume in GB
- `volume_encryption` (Attributes) The volume encryption related configuration. (see [below for nested schema](#nestedatt--instance_group--volume_encryption))

Optional:

- `recipes` (Set of String) The set of recipe names that are going to be applied on the given instance group.

<a id="nestedatt--instance_group--attached_volume_configuration"></a>
### Nested Schema for `instance_group.attached_volume_configuration`

Required:

- `volume_count` (Number) The number of volumes to be attached.
- `volume_size` (Number) The size of the volume in GB.
- `volume_type` (String) The - cloud provider - type of the volume.


<a id="nestedatt--instance_group--volume_encryption"></a>
### Nested Schema for `instance_group.volume_encryption`

Required:

- `encryption` (Boolean)



<a id="nestedatt--polling_options"></a>
### Nested Schema for `polling_options`

Optional:

- `polling_timeout` (Number) Timeout value in minutes that specifies for how long should the polling go for resource creation/deletion.