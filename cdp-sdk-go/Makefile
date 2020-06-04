##
# Copyright 2020 Cloudera, Inc.
##

GO_FLAGS:=""

API_DEFINITION_COMMIT ?= f3d16d71b576f179f3d7df6407e12635ca077f8c

IAM_SWAGGER_YAML=https://github.infra.cloudera.com/raw/thunderhead/thunderhead/$(API_DEFINITION_COMMIT)/services/libs/protocols-yaml/src/main/yaml/iam.yaml
ENVIRONMENTS_SWAGGER_YAML=https://github.infra.cloudera.com/raw/thunderhead/thunderhead/$(API_DEFINITION_COMMIT)/services/libs/protocols-yaml/src/main/yaml/environments2.yaml
DATALAKE_SWAGGER_YAML=https://github.infra.cloudera.com/raw/thunderhead/thunderhead/$(API_DEFINITION_COMMIT)/services/libs/protocols-yaml/src/main/yaml/datalake.yaml
DATAHUB_SWAGGER_YAML=https://github.infra.cloudera.com/raw/thunderhead/thunderhead/$(API_DEFINITION_COMMIT)/services/libs/protocols-yaml/src/main/yaml/datahub.yaml
ML_SWAGGER_YAML=https://github.infra.cloudera.com/raw/thunderhead/thunderhead/$(API_DEFINITION_COMMIT)/services/libs/protocols-yaml/src/main/yaml/ml.yaml

all: check-go test build

check-go:
ifndef GOPATH
	$(error GOPATH not defined, please define GOPATH. Run "go help gopath" to learn more about GOPATH)
endif
.PHONY: check-go

# Run tests
test: generate fmt vet
	go test $(GO_FLAGS) ./authn/... ./common/...

# Build main binary
build: generate fmt vet
	go build $(GO_FLAGS) ./authn/... ./common/...

# Run main binary
run: generate fmt vet
	go run $(GO_FLAGS) ./main.go ./common/...

# Run go fmt against code
fmt:
	go fmt ./authn/... ./common/... ./cdp/...

# Run go vet against code
vet:
	go vet ./authn/... ./common/... ./cdp/...

# Generate code
generate:
	go generate ./authn/... ./common/... ./cdp/...

mkdirs:
	mkdir -p gen/iam gen/environments gen/datalake gen/datahub gen/ml

# TODO: should run this as a part of main build
swagger-gen: mkdirs swagger-gen-iam swagger-gen-environments swagger-gen-datalake swagger-gen-datahub swagger-gen-ml

swagger-gen-iam: mkdirs
	go run github.com/go-swagger/go-swagger/cmd/swagger generate client -f $(IAM_SWAGGER_YAML) -A iam -t gen/iam/

swagger-gen-environments: mkdirs
	go run github.com/go-swagger/go-swagger/cmd/swagger generate client -f $(ENVIRONMENTS_SWAGGER_YAML) -A environments -t gen/environments/

swagger-gen-datalake: mkdirs
	go run github.com/go-swagger/go-swagger/cmd/swagger generate client -f $(DATALAKE_SWAGGER_YAML) -A datalake -t gen/datalake/

swagger-gen-datahub: mkdirs
	go run github.com/go-swagger/go-swagger/cmd/swagger generate client -f $(DATAHUB_SWAGGER_YAML) -A datahub -t  gen/datahub/

swagger-gen-ml: mkdirs
	go run github.com/go-swagger/go-swagger/cmd/swagger generate client -f $(ML_SWAGGER_YAML) -A ml -t  gen/ml/
