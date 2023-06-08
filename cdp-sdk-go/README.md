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

**Note:** Only publicly-released APIs can be consumed from the Terraform provider. We use the Beta API swaggers.

To upgrade the version, simply get a new git tag from https://github.com/cloudera/cdp-dev-docs/, and update it in the
Makefile. Afterwards, run:
```
  make swagger-gen
```
to generate for the new code. Then send a PR for the new changes to check in the code.

# Supporting a new Service
Adding a new CDP service (CML, CDW, CDE, etc) should be trivial and should follow the existing model from `Makefile`:
```
swagger-gen-<newservice>: mkdirs
	go run github.com/go-swagger/go-swagger/cmd/swagger generate client -f $(NEWSERVICE_SWAGGER_YAML) -A dw -t  gen/<newservice>/
```