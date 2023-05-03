---
subcategory: "iam"
---

# cdp_iam_group Resource

Provides an IAM group resource.

## Example Usage

```hcl
resource "cdp_iam_group" "my_group" {
  group_name = "my_group"
  sync_membership_on_user_login = true
}
```

## Argument Reference

The following arguments are supported:

* `group_name` - (Required, Forces new resource) The name of the group.
* `sync_membership_on_user_login` - (Optional) Whether group membership is synced when a user logs in. The default is to sync group membership.

## Attribute Reference

`id` is set to the group name. In addition, the following attributes are exported:

* `crn` - The CRN of the group.
