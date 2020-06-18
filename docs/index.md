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

Static credentials can be provided by adding an `cdp_access_key_id` and
`cdp_private_key` in-line in the CDP provider block:

Usage:

```hcl
provider "cdp" {
  cdp_access_key_id = "my-access-key"
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
this location. You can optionally specify a different location in the
configuration by providing the `cdp_shared_credentials_file` attribute, or
in the environment with the `CDP_SHARED_CREDENTIALS_FILE` variable.
This method also supports a `cdp_profile` configuration and matching
`CDP_PROFILE` environment variable:

Usage:

```hcl
provider "cdp" {
  cdp_profile                 = "customprofile"
}
```

## Configuration
You can use a CDP configuration file to specify other CDP related configuration.
The default location is `$HOME/.cdp/config` on Linux and OS X, or
`"%USERPROFILE%\.cdp\config"` for Windows users. You can optionally specify a
different location in the configuration by providing the `cdp_config_file`
attribute, or in the environment with the `CDP_CONFIG_FILE` variable. This method
also supports a `cdp_profile` configuration and matching `CDP_PROFILE` environment
variable to read configuration under different profiles:

```bash
$ cat $HOME/.cdp/config
[default]
endpoint_url = https://%sapi.us-west-1.altus.cloudera.com/
cdp_endpoint_url = https://api.us-west-1.cdp.cloudera.com/

[profile dev]
endpoint_url = ...
cdp_endpoint_url = ...
```

## Argument Reference

* List any arguments for the provider block.

In addition to [generic `provider` arguments](https://www.terraform.io/docs/configuration/providers.html)
(e.g. `alias` and `version`), the following arguments are supported in the CDP
 `provider` block:

* `cdp_access_key` - (Optional) This is the CDP access key. It must be provided, but
  it can also be sourced from the `CDP_ACCESS_KEY_ID` environment variable, or via
  a shared credentials file if `cdp_profile` is specified.

* `cdp_private_key` - (Optional) This is the CDP secret key. It must be provided, but
  it can also be sourced from the `CDP_PRIVATE_KEY` environment variable, or
  via a shared credentials file if `cdp_profile` is specified.

* `cdp_profile` - (Optional) This is the CDP profile name as set in the shared credentials
  file. It can also be sourced from the `CDP_PROFILE` environment variable.

* `cdp_config_file` - (Optional) This is the path to the configuration file.
  If this is not set and a profile is specified, `$HOME/.cdp/config` will be used.

* `cdp_shared_credentials_file` - (Optional) This is the path to the shared credentials file.
  If this is not set and a profile is specified, `$HOME/.cdp/credentials` will be used.

* `endpoint_url` - (Optional) Customize the endpoint URL format for connecting to alternate
  endpoints for IAM and Workload Management services. See the Custom [Service Endpoints Guide](guides/custom-service-endpoints.md)
  for more information about connecting to alternate CDP endpoints.

* `cdp_endpoint_url` - (Optional) Customize the endpoint URL format for connecting to alternate
  endpoints for CDP services. See the Custom [Service Endpoints Guide](guides/custom-service-endpoints.md)
  for more information about connecting to alternate CDP endpoints.