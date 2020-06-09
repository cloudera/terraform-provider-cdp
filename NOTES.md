. Lots of name vs. CRN stuff to resolve. 
. Need to add unit testing.
. Need to fill out properties of the resources. Review for optional, force new, etc.
. File JIRAs to make more properties mutable.
. Lots of validation to add.
. We ought to have a script to produce terraform schema and docs from our service models.
. Should thinks like datahubs depend on the datalake and not the environment? In the terraform if not the API?
. Deleting ID broker mappings requires a CRN while the rest of the ID broker mapping calls take the name or CRN.
. We should revisit the CRN format, and general CRN guidance, for the environment, datalake, etc.
. The modeling of the ID broker mappings, and the interaction with the datalake should be reviewed. 
. Do we need debug logging of our calls in TF_LOG like AWS has? Do we have that?
. Do we want to add additional provider arguments from the AWS list that would make sense for us?
. Enis's point on data source in addition to resource.
. This is worth investigating.a