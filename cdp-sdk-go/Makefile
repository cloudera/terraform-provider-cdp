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

SWAGGER_REPO_RAW=https://raw.githubusercontent.com/cloudera/cdp-dev-docs/refs/tags
API_DEFINITION_TAG ?= cdp-api-0.9.148-beta
SWAGGER=$(SWAGGER_REPO_RAW)/$(API_DEFINITION_TAG)/api-docs-beta/swagger

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

mkdir-%:
	mkdir -p gen/$*

# TODO: Manage this via go mod
install-go-swagger:
	go get github.com/go-swagger/go-swagger/cmd/swagger

swagger-gen: install-go-swagger
swagger-gen: swagger-gen-iam
swagger-gen: swagger-gen-environments
swagger-gen: swagger-gen-datalake
swagger-gen: swagger-gen-datahub
swagger-gen: swagger-gen-ml
swagger-gen: swagger-gen-de
swagger-gen: swagger-gen-dw
swagger-gen: swagger-gen-opdb

clone-swaggers: clone-swagger-iam
clone-swaggers: clone-swagger-environments
clone-swaggers: clone-swagger-datalake
clone-swaggers: clone-swagger-datahub
clone-swaggers: clone-swagger-ml
clone-swaggers: clone-swagger-de
clone-swaggers: clone-swagger-dw
clone-swaggers: clone-swagger-opdb

clone-swagger-%:
	wget $(SWAGGER)/$*.yaml -O resources/swagger/$*.yaml

swagger-gen-%: mkdir-%
	go run github.com/go-swagger/go-swagger/cmd/swagger generate client -f resources/swagger/$*.yaml -A $* -t gen/$*/