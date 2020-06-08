. you were curious about what to do for resource fields that are only used on input but are not available after -- e.g. the volume configuration stuff
. datalake creation timed out, and you were curious if it was going to make you re-create the resource or would pick it up and why the answer was what it was (do you understand taint?)
. Lots of name vs. CRN stuff to resolve. 
. Need to add unit testing.
. Need to add documentation.
. Need to add a release plan and establish versioning.
. Need to fill out properties of the resources. Review for optional, force new, etc.
. File JIRAs to make more properties mutable.
. Lots of validation to add.
. We ought to have a script to produce terraform schema and docs from our service models.
. Should thinks like datahubs depend on the datalake and not the environment? In the terraform if not the API?
. Deleting ID broker mappings requires a CRN while the rest of the ID broker mapping calls take the name or CRN.
. We should revisit the CRN format, and general CRN guidance, for the environment, datalake, etc.
. The modeling of the ID broker mappings, and the interaction with the datalake should be reviewed. 
. Do we need debug logging of our calls in TF_LOG like AWS has? Do we ahve that?