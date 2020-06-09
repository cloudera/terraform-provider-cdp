# CDP Provider

The Cloudera Data Platform (CDP) provider is used to interact with the resources
supported by CDP. The provider needs to be configured with the proper
credentials before it can be used.

## Example Usage

```
provider "cdp" {
}

resource "cdp_environments_aws_credential" "my-credentials" {
  name = "my-cloudera-cdp-credentials"
  role_arn = "arn:aws:iam::11111111111:role/my-cross-account-role"
  description = "testing the terraform-provider-cdp"
}
```

## Authentication

The CDP provider offers a flexible means of providing credentials for
authentication. The following methods are supported, in this order, and
explained below:

- Static credentials
- Environment variables
- Shared credentials file

### Static credentials

!> **Warning:** Hard-coding credentials into any Terraform configuration is not
recommended, and risks secret leakage should this file ever be committed to a
public version control system.

Static credentials can be provided by adding an `cdp_access_key` and
`cdp_private_key` in-line in the CDP provider block:

Usage:

```hcl
provider "cdp" {
  cdp_access_key = "my-access-key"
  cdp_private_key = "my-private-key"
}
```

### Environment variables

You can provide your credentials via the `CDP_ACCESS_KEY_ID` and
`CDP_PRIVATE_KEY`, environment variables, representing your CDP
access key and CDP private key, respectively.  Note that setting your
CDP credentials using either these environment variables will override
the use of `CDP_PROFILE`.

```hcl
provider "cdp" {}
```

Usage:

```sh
$ export CDP_ACCESS_KEY_ID="anaccesskey"
$ export CDP_PRIVATE_KEY="aprivatekey"
$ terraform plan
```

### Shared Credentials file

You can use a CDP credentials file to specify your credentials. The
default location is `$HOME/.cdp/credentials` on Linux and OS X, or
`"%USERPROFILE%\.cdp\credentials"` for Windows users. If we fail to
detect credentials inline, or in the environment, Terraform will check
this location. This method also supports a `profile` configuration and matching
`CDP_PROFILE` environment variable:

Usage:

```hcl
provider "cdp" {
  cdp_profile                 = "customprofile"
}
```

## Argument Reference

* List any arguments for the provider block.

In addition to [generic `provider` arguments](https://www.terraform.io/docs/configuration/providers.html)
(e.g. `alias` and `version`), the following arguments are supported in the CDP
 `provider` block:

* `cdp_access_key` - (Optional) This is the CDP access key. It must be provided, but
  it can also be sourced from the `CDP_ACCESS_KEY_ID` environment variable, or via
  a shared credentials file if `profile` is specified.

* `cdp_private_key` - (Optional) This is the CDP secret key. It must be provided, but
  it can also be sourced from the `CDP_PRIVATE_KEY` environment variable, or
  via a shared credentials file if `profile` is specified.

* `cdp_profile` - (Optional) This is the CDP profile name as set in the shared credentials
  file. It can also be sourced from the `CDP_PROFILE` environment variable.
