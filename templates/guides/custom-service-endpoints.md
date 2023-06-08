---
subcategory: ""
page_title: "Terraform CDP Provider Custom Service Endpoint Configuration"
description: |-
  Configuring the Terraform CDP Provider to connect to custom CDP service endpoints.
---

# Custom Service Endpoint Configuration

The Terraform CDP Provider configuration can be customized to connect to non-default CDP service endpoints. This may be
useful for local testing.

!> **Warning:** This is an advanced configuration option, and most Terraform users do not need to configure the
endpoints.

This guide outlines how to get started with customizing endpoints, the available endpoint configurations, and offers example configurations for working with certain local development and testing solutions.

<!-- TOC depthFrom:2 -->

- [Getting Started with Custom Endpoints](#getting-started-with-custom-endpoints)
- [Available Endpoint Customizations](#available-endpoint-customizations)
- [CDP Configuration File](#cdp-configuration-file)

<!-- /TOC -->

## Getting Started with Custom Endpoints

To configure the Terraform CDP Provider to use customized endpoints, it can be done within `provider` declarations using
the `endpoint_url` and `cdp_endpoint_url` arguments, e.g.

```hcl
provider "cdp" {
  # ... potentially other provider configuration ...

  endpoint_url = "..."
  cdp_endpoint_url = "..."
}
```

If multiple, different Terraform CDP Provider configurations are required, see the [Terraform documentation on multiple provider instances](https://www.terraform.io/docs/configuration/providers.html#alias-multiple-provider-instances) for additional information about the `alias` provider configuration and its usage.

## Available Endpoint Customizations

The Terraform CDP Provider allows the following endpoints to be customized:

<!-- markdownlint-disable MD033 -->
<div style="column-width: 14em;">
<ul>
  <li><code>endpoint_url</code> - Allows IAM and Workload Manager service endpoints to be customized. Defaults to <code>https://%sapi.us-west-1.altus.cloudera.com/</code></li>
  <li><code>cdp_endpoint_url</code> - Allows CDP service endpoints to be customized. Defaults to <code>https://api.us-west-1.cdp.cloudera.com/</code></li></li>
</ul>
</div>
<!-- markdownlint-enable MD033 -->

## CDP Configuration File
Alternative to setting arguments in the terraform provider, you can configure the CDP endpoints under a specific profile
or the default profile using the CDP configuration file in `$HOME/.cdp/config` and use the `cdp_profile` argument
in the provider or use the `CDP_PROFILE` environment variable to select the active profile.

```bash
$ cat $HOME/.cdp/config
[default]
endpoint_url = https://%sapi.us-west-1.altus.cloudera.com/
cdp_endpoint_url = https://api.us-west-1.cdp.cloudera.com/

[profile dev]
endpoint_url = ...
cdp_endpoint_url = ...
```

```hcl
provider "cdp" {
  cdp_profile                 = "customprofile"
}
```

