---
subcategory: "ml"
---

# ml_workspace Resource

Provides an ML workspace resource.

## Example Usage

```hcl
resource "cdp_ml_workspace" "my_workspace" {
  workspace_name   = "my-workspace"
  environment_name = "my-environment"

  kubernetes {
    instance_group {
      instance_type = "m5.2xlarge"
      autoscaling {
        min_instances = 0
        max_instances = 3
      }
    }
    instance_group {
      instance_type = "p2.8xlarge"
      autoscaling {
        min_instances = 0
        max_instances = 1
      }
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `workspace_name` - (Required, Forces new resource) The name of the workspace.
* `environment_name` - (Required, Forces new resource) The name of the environment.
* `use_public_load_balancer` - (Optional, Forces new resource) Whether to request public load balancer. By default a private load balancer is used.
* `disable_tls` - (Optional, Forces new resource) Whether to disable TLS setup for the workspace. By default TLS is enabled.
* `kubernetes` - (Required, Forces new resource) Parameters for the kubernetes that will host the workspace.

### Nested fields

#### `kubernetes`

* `instance_group` - (Required, Forces new resource) An instance group for the kubernetes cluster.

#### `instance_group`

* `instance_type` - (Required, Forces new resource) The instance type for the group.
* `autoscaling` - (Required, Forces new resource) The autoscaling values for the group.

#### `autoscaling`

* `min_instances` - (Required,  Forces new resource) The minimum number of instances.
* `max_instances` - (Required,  Forces new resource) The maximum number of instances.

## Attribute Reference

`id` is set to the workspace name appended to the environment name after a '::'.
In addition, the following attributes are exported:

* `crn` - The CRN of the workspace.
