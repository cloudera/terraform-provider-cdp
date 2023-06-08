## Copyright 2023 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

GO_FLAGS:=""

SWAGGER_REPO_RAW=https://raw.githubusercontent.com/cloudera/cdp-dev-docs/
API_DEFINITION_TAG ?= cdp-api-0.9.88
SWAGGER=$(SWAGGER_REPO_RAW)/$(API_DEFINITION_TAG)/api-docs-beta/swagger

IAM_SWAGGER_YAML=$(SWAGGER)/iam.yaml
ENVIRONMENTS_SWAGGER_YAML=$(SWAGGER)/environments.yaml
DATALAKE_SWAGGER_YAML=$(SWAGGER)/datalake.yaml
DATAHUB_SWAGGER_YAML=$(SWAGGER)/datahub.yaml
ML_SWAGGER_YAML=$(SWAGGER)/ml.yaml
DW_SWAGGER_YAML=$(SWAGGER)/dw.yaml

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
