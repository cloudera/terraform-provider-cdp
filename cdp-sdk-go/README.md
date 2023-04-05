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
consume new APIs from the thunderhead repository. We do not generate from the latest master, but instead use a git sha
as our commit id. The commit id can be found at [Makefile](Makefile):
```
API_DEFINITION_COMMIT ?= f38bd7b548dafe0a1e19629ff057699814077513
```

To upgrade the version, simply get a new git sha from thunderhead, and update it in the Makefile. Afterwards, run:
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