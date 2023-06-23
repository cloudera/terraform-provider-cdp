## 0.1.0 (Unreleased)

NOTES:
* The provider is in a alpha state. Expect some breaking changes.
* The provider is dev-tested only.

IMPROVEMENTS:
* Rebased the terraform provider code on top of the recommended terraform-plugin-framework.
* All existing resources and data sources have been either rewritten or dropped as a result.
* Migrated to tfplugindocs for the document generation. Drastically improved docs from auto-generated content.
* Migrated the code to github.com and open sourced it
* Switched to goreleaser
* Switched to github actions for builds and releases

FEATURES:
* Reimplemented `cdp` provider support using the new terraform-plugin-framework
* Reimplemented the resources for
    * `cdp_environments_aws_credential`
    * `cdp_environments_aws_environment`
    * `cdp_environments_id_broker_mappings`
    * `cdp_datalake_aws_datalake`
    * `cdp_datahub_aws_cluster`
    * `cdp_iam_group`
    * `cdp_environments_azure_credential`
* Reimplemented the data-sources for
    * `cdp_environments_aws_credential_prerequisites`
    * `cdp_iam_group`
* New Resouce: `cdp_datalake_azure_datalake`
* New Resouce: `cdp_environments_azure_environment`

BREAKING CHANGES:
* Removed support for `cdp_ml_workspace` resource (will be added back at a later release).

## 0.0.3 (Sep 28, 2020)

NOTES:

* The provider is in a pre-release state. Expect many breaking changes. The provider supports resource creation and destruction, but is not a good fit yet ongoing resource management.
* Documentation of existing resources is in progress but far from complete


IMPROVEMENTS:
* provider: Add `cdp_config_file`, `cdp_shared_credentials_file`, `cdp_endpoint_url` and `endpoint_url` arguments
* resource/cdp_environments_aws_environment: Add retries with exponential backoff to ride over consistency issues.

FEATURES:

* **New Resource:** `cdp_environments_azure_credential`

## 0.0.2 (June 12, 2020)

NOTES:

* The provider is in a pre-release state. Expect many breaking changes. The provider supports resource creation and destruction, but is not a good fit yet ongoing resource management.
* Documentation of existing resources is in progress but far from complete

FEATURES:

* **New Resource:** `cdp_iam_group`
* **New Data Source:** `cdp_iam_group`
* **New Resource:** `cdp_ml_workspace`

## 0.0.1 (June 8, 2020)

NOTES:

* The provider is in a pre-release state. Expect many breaking changes. The provider supports resource creation and destruction, but is not a good fit yet ongoing resource management.
* Documentation of existing resources is in progress but far from complete

FEATURES:

* **New Data Source:** `cdp_environments_aws_credential_prerequisites`
* **New Resource:** `cdp_environments_aws_credential`
* **New Resource:** `cdp_environments_aws_environment`
* **New Resource:** `cdp_environments_id_broker_mappings`
* **New Resource:** `cdp_datalake_aws_datalake`
* **New Resource:** `cdp_datahub_aws_cluster`
