##
# Copyright 2020 Cloudera, Inc.
##

GO_FLAGS:=""

IAM_SWAGGER_YAML=https://github.infra.cloudera.com/raw/thunderhead/thunderhead/master/services/libs/protocols-yaml/src/main/yaml/iam.yaml
ENVIRONMENTS2_SWAGGER_YAML=https://github.infra.cloudera.com/raw/thunderhead/thunderhead/master/services/libs/protocols-yaml/src/main/yaml/environments2.yaml

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
	go fmt ./authn/... ./common/... ./config/...

# Run go vet against code
vet:
	go vet ./authn/... ./common/... ./config/...

# Generate code
generate:
	go generate ./authn/... ./common/... ./config/...

mkdirs:
	mkdir -p gen/iam gen/environments

# TODO: should run this as a part of main build
swagger-gen: mkdirs swagger-gen-iam swagger-gen-environments2

swagger-gen-iam: mkdirs
	go run github.com/go-swagger/go-swagger/cmd/swagger generate client -f $(IAM_SWAGGER_YAML) -A iam -t  gen/iam/

swagger-gen-environments2:
	go run github.com/go-swagger/go-swagger/cmd/swagger generate client -f $(ENVIRONMENTS2_SWAGGER_YAML) -A environments -t  gen/environments/