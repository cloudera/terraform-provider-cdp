##
# Copyright 2020 Cloudera, Inc.
##

GO_FLAGS:=""

API_DEFINITION_COMMIT ?= f38bd7b548dafe0a1e19629ff057699814077513

IAM_SWAGGER_YAML=https://github.infra.cloudera.com/raw/thunderhead/thunderhead/$(API_DEFINITION_COMMIT)/services/libs/protocols-yaml/src/main/yaml/iam.yaml
ENVIRONMENTS_SWAGGER_YAML=https://github.infra.cloudera.com/raw/thunderhead/thunderhead/$(API_DEFINITION_COMMIT)/services/libs/protocols-yaml/src/main/yaml/environments2.yaml
DATALAKE_SWAGGER_YAML=https://github.infra.cloudera.com/raw/thunderhead/thunderhead/$(API_DEFINITION_COMMIT)/services/libs/protocols-yaml/src/main/yaml/datalake.yaml
DATAHUB_SWAGGER_YAML=https://github.infra.cloudera.com/raw/thunderhead/thunderhead/$(API_DEFINITION_COMMIT)/services/libs/protocols-yaml/src/main/yaml/datahub.yaml
ML_SWAGGER_YAML=https://github.infra.cloudera.com/raw/thunderhead/thunderhead/$(API_DEFINITION_COMMIT)/services/libs/protocols-yaml/src/main/yaml/ml.yaml
DW_SWAGGER_YAML=https://github.infra.cloudera.com/raw/thunderhead/thunderhead/$(API_DEFINITION_COMMIT)/services/libs/protocols-yaml/src/main/yaml/dw.yaml

all: check-go test build

check-go:
ifndef GOPATH
	$(error GOPATH not defined, please define GOPATH. Run "go help gopath" to learn more about GOPATH)
endif
.PHONY: check-go

# Run tests
test: generate fmt vet
	go test $(GO_FLAGS) ./cdp/... ./common/...

# Build main binary
build: generate fmt vet
	go build $(GO_FLAGS) ./cdp/... ./common/...

# Run main binary
run: generate fmt vet
	go run $(GO_FLAGS) ./main.go ./common/...

# Run go fmt against code
fmt:
	go fmt ./common/... ./cdp/...

# Run go vet against code
vet:
	go vet ./common/... ./cdp/...

# Generate code
generate:
	go generate ./common/... ./cdp/...

mkdirs:
	mkdir -p gen/iam gen/environments gen/datalake gen/datahub gen/ml gen/dw

# TODO: Manage this via go mod
install-go-swagger:
	go get github.com/go-swagger/go-swagger/cmd/swagger

# TODO: should run this as a part of main build
swagger-gen: mkdirs install-go-swagger swagger-gen-iam swagger-gen-environments swagger-gen-datalake swagger-gen-datahub swagger-gen-ml swagger-gen-dw

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

swagger-gen-dw: mkdirs
	go run github.com/go-swagger/go-swagger/cmd/swagger generate client -f $(DW_SWAGGER_YAML) -A dw -t  gen/dw/
