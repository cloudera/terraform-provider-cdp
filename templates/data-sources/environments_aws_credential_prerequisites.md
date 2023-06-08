---
subcategory: "environments"
---

# cdp_environments_aws_credential_prerequisites Data Source

Use this data source to get information required to set up a delegated access
role that can be used to create a CDP credential.

## Example Usage

```hcl
data "cdp_environments_aws_credential_prerequisites" "credential_prerequisites" {
}

resource "aws_iam_role" "delegated_access_role" {
  name = "${var.prefix}delegated-access"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::${data.cdp_environments_aws_credential_prerequisites.credential_prerequisites.account_id}:root"
      },
      "Action": "sts:AssumeRole",
      "Condition": {
        "StringEquals": {
          "sts:ExternalId": "${data.cdp_environments_aws_credential_prerequisites.credential_prerequisites.external_id}"
        }
      }
    }
  ]
}
    
EOF

}
```

## Argument Reference

This data source has no arguments.

## Attribute Reference

`id` is set to the external ID value. In addition, the following attributes are
exported:

* `account_id` - The AWS account ID of the identity used by CDP when assuming a
  delegated access role associated with a CDP credential.
* `external_id` - The external ID that will be used when assuming a delegated
  access role associated with a CDP credential.
