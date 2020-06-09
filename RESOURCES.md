# Resource Status

This document exists to give a high level overview of the currently supported data sources and resoures and their levels of completenes.

As of this writing documentation and tests do not exist for any entities, so these are not mentioned as gaps. Once these begin to be added, this document will be updated to indicate which entities need work in those areas.

## Data Sources

### cdp_environments_aws_credential_prerequisites

This is complete with no known issues.

## Resources

### cdp_environments_aws_credential

The provider needs to to be reviewed for the following:

* Are there gaps compared to the API definition? (create, read, update)
 * Are fields missing?
 * Are fields correctly modeled? (optional, required, force new, etc)
* Are there gaps compared to the UI? On Create?  (create, read, update)

This resource needs to be documented.

### cdp_environments_aws_environment

This is relatively fleshed out and should be usable.

The provider needs to to be reviewed for the following:

* Are there gaps compared to the API definition? (create, read, update)
 * Are fields missing?
 * Are fields correctly modeled? (optional, required, force new, etc)
* Are there gaps compared to the UI? On Create?  (create, read, update)

This resource needs to be documented.

### cdp_environments_id_broker_mappings

It is an open question as to whether this is modelled correctly. If we stick with this model, this is complete with no known issues.

This resource needs to be documented.

### cdp_datalake_aws_datalake

This is relatively fleshed out and should be usable.

The provider needs to to be reviewed for the following:

* Are there gaps compared to the API definition? (create, read, update)
 * Are fields missing?
 * Are fields correctly modeled? (optional, required, force new, etc)
* Are there gaps compared to the UI? On Create?  (create, read, update)

This resource needs to be documented.

### cdp_datahub_aws_cluster

This is modelled but not yet fully implemented in the provider. It should not be used yet.

This resource needs to be documented.