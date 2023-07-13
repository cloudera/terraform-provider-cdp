A super bare bones SDK implementation in Golang for CDP. This should be its own project, but hosted here
temporarily until we do that.

# Generating via Swagger
We use go-swagger (for now) to generate the go code for the SDK based on OpenSpec API files (YAMLs).

To generate for all services, run:
```
  make swagger-gen
```

We check in the code for the generated files so that the build does not depend on running go-swagger and running the
code generate which takes some time for every build. We can revisit this decision later.

# Upgrading for newer APIs
From time to time, we need to upgrade the version of the swagger files that we are generating from so that we can
consume new APIs. We do not generate from the latest master, but instead use a version tag which corresponds to a public
API release. These tags correspond to CDP API release versions. The tag can be found at [Makefile](Makefile):
```
API_DEFINITION_TAG ?= cdp-api-0.9.88
```
Then run:
```
    make clone-swaggers
```
which will copy the swagger files from the upstream repo to this repo.

We have a local copy of the swagger files from the upstream repos since we sometimes need the changes to the API that
are yet to be released publicly. The local copies of the swagger files can be found at
[./resources/swagger](./resources/swagger). A PR can be sent to reflect the API changes to these swagger files for
upcoming API changes that will be released soon, but since we have a local copy, we can generate the swagger code before
waiting for the API release. Once the API changes are released publicly, simply update the local copy of the swagger files
by running:
```
    make clone-swaggers
```
which will override any local changes from the upstream repo.

To upgrade the version, simply get a new git tag from https://github.com/cloudera/cdp-dev-docs/, and update it in the
Makefile. Afterwards, run:
```
  make clone-swaggers swagger-gen
```
to generate for the new code. Then send a PR for the new changes to check in the code.

**Note:** Only publicly-released (or soon-to-be-released) APIs can be consumed from the Terraform provider. We use the
Beta API swaggers.

# Supporting a new Service
Adding a new CDP service (CML, CDW, CDE, etc) should be trivial and should follow the existing model from `Makefile`.
Add two new directives like this:
```
swagger-gen: swagger-gen-<service-name>

clone-swaggers: clone-swagger-<service-name>
```