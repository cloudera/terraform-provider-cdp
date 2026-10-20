---
subcategory: ""
page_title: "Terraform CDP Provider Authentication"
description: |-
  Configuring credentials and authentication for the Terraform CDP Provider.
---

# Authentication and Configuration

The CDP Terraform provider needs credentials to communicate with the CDP control plane API. This guide covers all
supported authentication methods, from simplest to most advanced.

If you are familiar with the [AWS Terraform provider](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#authentication-and-configuration)
authentication model, the CDP provider follows the same conventions: static credentials, environment variables, and
shared credential files with named profiles.

<!-- TOC depthFrom:2 -->

- [Quick Start](#quick-start)
- [Authentication Methods](#authentication-methods)
  - [Static Credentials](#static-credentials)
  - [Environment Variables](#environment-variables)
  - [Shared Credentials File](#shared-credentials-file)
  - [Custom File Paths](#custom-file-paths)
- [Credential Resolution Order](#credential-resolution-order)
- [Named Profiles](#named-profiles)
- [CDP Region](#cdp-region)

<!-- /TOC -->

## Quick Start

If you have already configured the CDP CLI with `cdp configure`, your credentials are stored in `~/.cdp/credentials`
and the provider picks them up automatically:

```hcl
terraform {
  required_providers {
    cdp = {
      source = "cloudera/cdp"
    }
  }
}

provider "cdp" {}
```

No additional configuration is needed. The provider reads your default profile from `~/.cdp/credentials`.

## Authentication Methods

### Static Credentials

You can provide credentials directly in the provider block. Use Terraform variables to avoid hardcoding secrets:

```hcl
variable "cdp_access_key_id" {
  type      = string
  sensitive = true
}

variable "cdp_private_key" {
  type      = string
  sensitive = true
}

provider "cdp" {
  cdp_access_key_id = var.cdp_access_key_id
  cdp_private_key   = var.cdp_private_key
}
```

!> **Warning:** Never commit credentials to version control. Use variables, environment variables, or a shared
credentials file instead.

### Environment Variables

The provider reads credentials from the following environment variables:

- `CDP_ACCESS_KEY_ID` — Your CDP access key
- `CDP_PRIVATE_KEY` — Your CDP private key

```bash
export CDP_ACCESS_KEY_ID="your-access-key-id"
export CDP_PRIVATE_KEY="your-private-key"
```

```hcl
provider "cdp" {}
```

This approach works well in CI/CD pipelines where secrets are injected as environment variables.

### Shared Credentials File

The default location is `~/.cdp/credentials`, the same file used by the CDP CLI. The file format uses INI-style
profiles:

```ini
[default]
cdp_access_key_id = your-access-key-id
cdp_private_key = your-private-key

[profile staging]
cdp_access_key_id = staging-access-key-id
cdp_private_key = staging-private-key
```

To use a specific profile:

```hcl
provider "cdp" {
  cdp_profile = "staging"
}
```

### Custom File Paths

If your credentials or configuration files are in non-default locations:

```hcl
provider "cdp" {
  cdp_config_file             = "/path/to/config"
  cdp_shared_credentials_file = "/path/to/credentials"
}
```

## Credential Resolution Order

The provider resolves credentials in this order (first match wins):

1. **Static credentials** in the provider block (`cdp_access_key_id` + `cdp_private_key`)
2. **Environment variables** (`CDP_ACCESS_KEY_ID` + `CDP_PRIVATE_KEY`)
3. **Shared credentials file** (`~/.cdp/credentials`, or custom path via `cdp_shared_credentials_file`)

This mirrors the AWS provider's resolution chain, so teams already using AWS Terraform workflows can apply the same
patterns (e.g., static credentials for local dev, environment variables in CI).

## Named Profiles

Named profiles let you manage multiple CDP accounts or environments from a single credentials file. Select a profile
with the `cdp_profile` argument or the `CDP_PROFILE` environment variable:

```hcl
provider "cdp" {
  cdp_profile = "production"
}
```

```bash
export CDP_PROFILE="production"
```

You can also use [provider aliases](https://developer.hashicorp.com/terraform/language/providers/configuration#alias-multiple-provider-configurations)
to work with multiple CDP accounts in the same configuration:

```hcl
provider "cdp" {
  alias       = "dev"
  cdp_profile = "dev"
}

provider "cdp" {
  alias       = "prod"
  cdp_profile = "prod"
}

resource "cdp_environments_aws_credential" "dev_cred" {
  provider        = cdp.dev
  credential_name = "dev-credential"
  role_arn        = "arn:aws:iam::111111111111:role/dev-cross-account-role"
}
```

## CDP Region

The `cdp_region` controls which CDP control plane your provider talks to. This is **not** the cloud provider region
(e.g., `us-east-1` in AWS) — it is the region of the CDP management service itself.

Available CDP regions:

| Region | Description |
|--------|-------------|
| `us-west-1` | United States (default) |
| `eu-1` | Europe |
| `ap-1` | Asia Pacific |
| `usg-1` | US Government |

```hcl
provider "cdp" {
  cdp_region = "eu-1"
}
```

Or via environment variable:

```bash
export CDP_REGION="eu-1"
```

-> **Note:** The `cdp_region` determines where CDP metadata and control plane operations run. The actual cloud
resources (VPCs, storage, compute) are created in the cloud provider region you specify in each environment resource.
