## 0.2.0 (Unreleased)

NOTES:
* Added comprehensive DataFlow (DF) resource and data source support for managing NiFi flow definitions, deployments, collections, and projects via Terraform.
* No breaking changes to existing resources or data sources.

FEATURES:
* **New Resource:** `cdp_df_flow_definition` — Import NiFi flow definitions into the DataFlow catalog. Supports uploading flow JSON, assigning to collections, and automatic version management. Exposes `flow_version_crn` for use with deployments. Uses the two-step CDP API upload protocol (redirect + raw body) matching the CDP CLI behavior.
* **New Resource:** `cdp_df_collection` — Create and manage DataFlow catalog collections for organizing flow definitions.
* **New Resource:** `cdp_df_project` — Create and manage DataFlow projects within a DataFlow service.

* **New Data Source:** `cdp_df_service` — Look up a single DataFlow service by name. Returns CRN, environment CRN, status, and other fields directly (no array indexing).
* **New Data Source:** `cdp_df_project` — Look up a DataFlow project by name and return its CRN.

IMPROVEMENTS:
* **`cdp_df_deployment`:** Added full workload API integration for deployment creation, matching the CDP CLI `create-deployment` workflow (initiate → workload auth → get request details → create deployment).
* **`cdp_df_deployment`:** New optional attributes: `deployment_name`, `cluster_size`, `cfm_nifi_version`, `auto_start_flow` (defaults to false), `project_crn`, `static_node_count`, `auto_scaling_enabled`, `auto_scale_min_nodes`, `auto_scale_max_nodes`, `parameter_groups`.
* **`cdp_df_deployment`:** In-place flow version changes — changing `flow_version_crn` no longer forces replacement. Instead, it triggers the multi-step `change-flow-version` workflow (initiate → workload auth → get request details → change flow version) with configurable `strategy` and `wait_for_flow_to_stop_in_minutes`.
* **`cdp_df_deployment`:** `parameter_groups` is marked sensitive to hide large JSON diffs from plan output. A computed `parameter_groups_sha` attribute provides a clean one-line change indicator.
* **`cdp_df_flow_definition`:** `file` attribute is marked sensitive to hide large flow JSON diffs. A computed `file_sha` attribute provides change detection.
* **`data.cdp_df_services`:** Added optional `name` filter to narrow results.
* Added examples for all new resources and data sources under `examples/`.
* Added unit tests for `cdp_df_collection`, `cdp_df_project`, `cdp_df_flow_definition` (helper functions and SHA computation), `data.cdp_df_project`, and `data.cdp_df_service`.
* Generated mock interface for the DF client service (`MockDfClientService`) via mockery.

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
