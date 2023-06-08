---
subcategory: "iam"
---

# cdp_iam_group Data Source

Provides an IAM group data source.

## Example Usage

```hcl
data "cdp_iam_group" "my_group" {
  group_name = "my_group"
}
```

## Argument Reference

The following arguments are supported:

* `group_name` - (Required) The name of the group.

## Attribute Reference

`id` is set to the group name. In addition, the following attributes are exported:

* `sync_membership_on_user_login` - Whether group membership is synced when a user logs in.
* `crn` - The CRN of the group.

