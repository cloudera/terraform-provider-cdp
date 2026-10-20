---
subcategory: ""
page_title: "CDP Environment Lifecycle"
description: |-
  Understanding the CDP resource hierarchy: from cloud credentials to running data services.
---

# CDP Environment Lifecycle

CDP resources form a dependency chain: you create cloud credentials, then an environment, then a datalake, and finally
data services like DataHub, Data Warehouse, or Operational Database on top. This guide walks through each layer with
concrete Terraform examples.

If you have used the AWS Terraform provider, think of the environment as your VPC, the datalake as your shared storage
and security layer, and data services as the compute clusters you run inside them.

<!-- TOC depthFrom:2 -->

- [Resource Hierarchy](#resource-hierarchy)
- [Step 1: Cloud Credentials](#step-1-cloud-credentials)
- [Step 2: Environment](#step-2-environment)
- [Step 3: Datalake](#step-3-datalake)
- [Step 4: Data Services](#step-4-data-services)
- [Putting It All Together](#putting-it-all-together)
- [Cross-Cloud Reference](#cross-cloud-reference)

<!-- /TOC -->

## Resource Hierarchy

```
Cloud Credential
  └── Environment (references a credential)
        └── Datalake (references an environment)
              ├── DataHub cluster
              ├── Data Warehouse
              └── Operational Database
```

Each layer depends on the one above it. Terraform determines the creation order from attribute references
between resources (e.g., passing the credential name to the environment). However, data services such as
DataHub depend on a running datalake without directly referencing it, so you must add an explicit
`depends_on` for that link (see [Putting It All Together](#putting-it-all-together) below).

## Step 1: Cloud Credentials

A CDP credential tells the platform how to access your cloud account. You create a cross-account role (AWS), app
registration (Azure), or service account (GCP) in your cloud provider, then register it in CDP.

**AWS** — register a cross-account IAM role ARN:

```hcl
resource "cdp_environments_aws_credential" "my_cred" {
  credential_name = "my-aws-credential"
  role_arn        = "arn:aws:iam::111111111111:role/my-cross-account-role"
  description     = "Cross-account role for CDP"
}
```

**Azure** — register an app registration with its service principal:

```hcl
resource "cdp_environments_azure_credential" "my_cred" {
  credential_name = "my-azure-credential"
  subscription_id = "00000000-0000-0000-0000-000000000000"
  tenant_id       = "00000000-0000-0000-0000-000000000000"
  app_based = {
    application_id = "00000000-0000-0000-0000-000000000000"
    secret_key     = var.azure_app_secret
  }
}
```

**GCP** — register a service account key:

```hcl
resource "cdp_environments_gcp_credential" "my_cred" {
  credential_name = "my-gcp-credential"
  credential_key  = filebase64("gcp-service-account-key.json")
}
```

-> **Tip:** You can use the `cdp_environments_aws_credential_prerequisites` data source to discover the policies and
trust relationship your cross-account role needs before creating the credential.

## Step 2: Environment

An environment ties a CDP workspace to a specific cloud region and network. It references the credential from step 1
and specifies networking, logging, and security settings.

**AWS example:**

```hcl
resource "cdp_environments_aws_environment" "my_env" {
  environment_name = "my-environment"
  credential_name  = cdp_environments_aws_credential.my_cred.credential_name
  region           = "us-west-2"

  security_access = {
    cidr = "10.0.0.0/16"
  }

  authentication = {
    public_key_id = "my-ssh-key"
  }

  log_storage = {
    storage_location_base = "s3a://my-cdp-logs/logs"
    instance_profile      = "arn:aws:iam::111111111111:instance-profile/my-log-instance-profile"
  }

  vpc_id     = "vpc-0123456789abcdef0"
  subnet_ids = ["subnet-aaa", "subnet-bbb", "subnet-ccc"]
}
```

Key concepts:
- **`region`** is the cloud provider region (e.g., `us-west-2`), not the CDP control plane region
- **`security_access`** controls network access — you can use a CIDR range or reference existing security groups
- **`log_storage`** configures where CDP writes audit and operational logs
- **`authentication`** specifies the SSH public key for cluster node access

The environment creates the network foundation, a FreeIPA instance (for identity management), and the security
infrastructure that all data services share.

## Step 3: Datalake

A datalake provides the shared storage, security (Apache Ranger), and governance (Apache Atlas) layer for an
environment. Each environment has exactly one datalake.

**AWS example:**

```hcl
resource "cdp_datalake_aws_datalake" "my_dl" {
  datalake_name         = "my-datalake"
  environment           = cdp_environments_aws_environment.my_env.environment_name
  instance_profile      = "arn:aws:iam::111111111111:instance-profile/my-datalake-admin-profile"
  storage_location_base = "s3a://my-cdp-data/datalake"
}
```

The `instance_profile` is an IDBroker profile that maps CDP users to cloud storage permissions — this is how CDP
provides fine-grained access control to your data lake storage.

~> **Note:** Datalake creation typically takes 20-30 minutes. Use the `timeouts` block to adjust the default timeout
if needed:

```hcl
resource "cdp_datalake_aws_datalake" "my_dl" {
  # ... attributes ...

  timeouts = {
    create = "60m"
  }
}
```

## Step 4: Data Services

With an environment and datalake in place, you can create data services. These are the compute workloads that run on
top of your data platform.

**DataHub** — a general-purpose cluster (Spark, Hive, Kafka, etc.):

```hcl
resource "cdp_datahub_aws_cluster" "my_cluster" {
  name               = "my-datahub"
  environment        = cdp_environments_aws_environment.my_env.environment_name
  cluster_template   = "7.2.18 - Data Engineering: Apache Spark, Apache Hive, Apache Oozie"
  cluster_definition = "7.2.18 - Data Engineering for AWS"

  destroy_options = {
    force_delete_cluster = false
  }

  depends_on = [cdp_datalake_aws_datalake.my_dl]
}
```

The template-based approach (shown above) is the simplest way to create a DataHub — CDP applies a predefined cluster
layout. For full control over instance types and node counts, use the `instance_group` attribute instead (see the
[`cdp_datahub_aws_cluster` resource documentation](../resources/datahub_aws_cluster.md)).

Other data services follow the same pattern: they reference the environment and require a running datalake
(declared via `depends_on`):

- **Data Warehouse** — `cdp_dw_aws_cluster` → `cdp_dw_database_catalog` → `cdp_dw_hive_vw` / `cdp_dw_impala_vw`
- **Operational Database** — `cdp_opdb_operational_database`

## Putting It All Together

Here is a complete AWS example that creates the full stack:

```hcl
terraform {
  required_providers {
    cdp = {
      source = "cloudera/cdp"
    }
  }
}

provider "cdp" {}

# 1. Register cloud credentials
resource "cdp_environments_aws_credential" "cred" {
  credential_name = "my-credential"
  role_arn        = "arn:aws:iam::111111111111:role/cross-account-role"
}

# 2. Create an environment
resource "cdp_environments_aws_environment" "env" {
  environment_name = "my-environment"
  credential_name  = cdp_environments_aws_credential.cred.credential_name
  region           = "us-west-2"

  security_access = {
    cidr = "10.0.0.0/16"
  }

  authentication = {
    public_key_id = "my-ssh-key"
  }

  log_storage = {
    storage_location_base = "s3a://my-cdp-logs/logs"
    instance_profile      = "arn:aws:iam::111111111111:instance-profile/log-profile"
  }

  vpc_id     = "vpc-0123456789abcdef0"
  subnet_ids = ["subnet-aaa", "subnet-bbb", "subnet-ccc"]
}

# 3. Create a datalake
resource "cdp_datalake_aws_datalake" "dl" {
  datalake_name         = "my-datalake"
  environment           = cdp_environments_aws_environment.env.environment_name
  instance_profile      = "arn:aws:iam::111111111111:instance-profile/datalake-admin-profile"
  storage_location_base = "s3a://my-cdp-data/datalake"
}

# 4. Create a DataHub cluster
resource "cdp_datahub_aws_cluster" "datahub" {
  name               = "my-datahub"
  environment        = cdp_environments_aws_environment.env.environment_name
  cluster_template   = "7.2.18 - Data Engineering: Apache Spark, Apache Hive, Apache Oozie"
  cluster_definition = "7.2.18 - Data Engineering for AWS"

  depends_on = [cdp_datalake_aws_datalake.dl]
}
```

Terraform determines the creation order from attribute references (credential → environment → datalake).
Data services like DataHub do not reference the datalake directly, so the `depends_on` on the datalake is
required — without it Terraform would create the DataHub in parallel with the datalake, which fails.

Destroying this configuration tears down the resources in reverse order.

## Cross-Cloud Reference

The resource hierarchy is the same across all three cloud providers, but some attribute names differ:

### Credential Resources

| Cloud | Resource | Key Attributes |
|-------|----------|----------------|
| AWS | `cdp_environments_aws_credential` | `credential_name`, `role_arn` |
| Azure | `cdp_environments_azure_credential` | `credential_name`, `subscription_id`, `tenant_id`, `app_based` |
| GCP | `cdp_environments_gcp_credential` | `credential_name`, `credential_key` |

### Environment Resources

| Attribute | AWS | Azure | GCP |
|-----------|-----|-------|-----|
| SSH key | `authentication.public_key_id` | `public_key` | `public_key` |
| Network | `vpc_id` + `subnet_ids` | `existing_network_params` | `existing_network_params` |
| Log storage identity | `instance_profile` (IAM) | `managed_identity` | _(via service account)_ |

### Datalake Resources

| Cloud | Resource | Storage Identity |
|-------|----------|-----------------|
| AWS | `cdp_datalake_aws_datalake` | `instance_profile` |
| Azure | `cdp_datalake_azure_datalake` | `managed_identity` |
| GCP | `cdp_datalake_gcp_datalake` | `cloud_provider_configuration.service_account_email` |

### DataHub Naming Differences

| Attribute | AWS / Azure | GCP |
|-----------|-------------|-----|
| Environment ref | `environment` | `environment_name` |
| Template | `cluster_template` | `cluster_template_name` |
| Definition | `cluster_definition` | `cluster_definition_name` |

-> **Note:** These naming differences across cloud providers are a known inconsistency. When writing modules that
support multiple clouds, use variables to abstract over them.
