---
page_title: "Cloudera Data Platform (CDP) Provider"
subcategory: ""
description: |-
  The Cloudera Data Platform (CDP) provider manages the lifecycle of resources supported by CDP like Credentials, Environment, Datalake, Datahub and other data services.
---

# CDP Provider

The Cloudera Data Platform (CDP) provider manages the lifecycle of resources supported by CDP like Credentials, Environment, Datalake, Datahub and other data services.

The provider needs to be configured with the proper credentials before it can be used (see Authentication section below).

## Example Usage
```terraform
## Copyright 2023 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

# This example is mimimum configuration for CDP provider and creates an aws credential in CDP
terraform {
  required_providers {
    cdp = {
      source = "registry.terraform.io/cloudera/cdp"
    }
  }
}

provider "cdp" {
}

resource "cdp_environments_aws_credential" "example" {
  name        = "example-cdp-aws-credential"
  role_arn    = "arn:aws:iam::11111111111:role/example-cross-account-role"
  description = "Example AWS Credential"
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
```terraform
## Copyright 2023 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

# This example shows how to use cdp access key id and secret key to manually configure the credentials in the provider
# configuration block.
#
# You can follow the guide at
# https://docs.cloudera.com/cdp-public-cloud/cloud/cli/topics/mc-cli-generating-an-api-access-key.html
# to generate your API access credentials.
#
# WARNING:  Hard-coding credentials into any Terraform configuration is NOT
# recommended, and risks secret leakage should this file ever be committed to a
# public version control system.
#
# You can also specify the credentials as environment variables.

terraform {
  required_providers {
    cdp = {
      source = "registry.terraform.io/cloudera/cdp"
    }
  }
}

provider "cdp" {
  cdp_access_key_id = var.cdp_access_key_id
  cdp_private_key   = var.cdp_private_key
}

variable "cdp_access_key_id" {
  description = "The access key id for the CDP credentials."
}

variable "cdp_private_key" {
  description = "The private key for the CDP credentials."
}
```


### Environment variables

You can provide your credentials via the `CDP_ACCESS_KEY_ID` and
`CDP_PRIVATE_KEY`, environment variables, representing your CDP
access key and CDP private key, respectively.  Note that setting your
CDP credentials using either these environment variables will override
the use of `CDP_PROFILE`.

```terraform
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

```terraform
## Copyright 2023 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

# This example shows how to use cdp with a custom profile (other than 'default').
# The profile must be defined in the CDP configuration file (default: ~/.cdp/config) and credentials should be available
# under ~/.cdp/credentials.

terraform {
  required_providers {
    cdp = {
      source = "registry.terraform.io/cloudera/cdp"
    }
  }
}

provider "cdp" {
  cdp_profile = "customprofile"
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

You can also override the config and credentials file locations that CDP provider. Usage:

```terraform
## Copyright 2023 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

# This example shows how to use cdp with custom config and credentials files.

provider "cdp" {
  cdp_config_file             = "/Users/myuser/.cdp/config"
  cdp_shared_credentials_file = "/Users/myuser/.cdp/credentials"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `cdp_access_key_id` (String) CDP access key id to authenticate the requests. It can be provided in the provider config (not recommended!), or it can be sourced from the `CDP_ACCESS_KEY_ID` environment variable, or via a shared credentials file. If `cdp_profile` is specified credentials for the specific profile will be used.
- `cdp_config_file` (String) CDP configuration file. Defaults to ~/.cdp/config.
- `cdp_endpoint_url` (String) CDP Endpoint URL to use. Customize the endpoint URL format for connecting to alternate endpoints for CDP services. See the Custom [Service Endpoints Guide](guides/custom-service-endpoints.md) for more information about connecting to alternate CDP endpoints.
- `cdp_private_key` (String, Sensitive) CDP private key associated with the given access key. It can be provided in the provider config(not recommended!), or it can also be sourced from the `CDP_PRIVATE_KEY` environment variable, or via a shared credentials file. If `cdp_profile` is specified credentials for the specific profile will be used.
- `cdp_profile` (String) CDP Profile to use for the configuration in shared credentials file (~/.cdp/credentials). It can also be sourced from the `CDP_PROFILE` environment variable.
- `cdp_shared_credentials_file` (String) CDP shared credentials file. Defaults to ~/.cdp/credentials.
- `endpoint_url` (String) Endpoint URL to use. Customize the endpoint URL format for connecting to alternate endpoints for IAM and Workload Management services. See the Custom [Service Endpoints Guide](guides/custom-service-endpoints.md) for more information about connecting to alternate CDP endpoints.
- `local_environment` (Boolean) Defines wether CDP CP runs locally. Defaults to false.