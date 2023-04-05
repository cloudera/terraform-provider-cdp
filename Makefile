##
# Copyright 2020 Cloudera, Inc.
##

GO_FLAGS:=""

VERSION ?= 0.0.4
ARCH := $(shell uname -s | tr A-Z a-z)_amd64
TF_PLUGIN_DIR ?= ~/.terraform.d/plugins
TF_PROVIDER_NAME ?= terraform.cloudera.com/cloudera/cdp

all: check-go test main

check-go:
ifndef GOPATH
	$(error GOPATH not defined, please define GOPATH. Run "go help gopath" to learn more about GOPATH)
endif
.PHONY: check-go

# Run tests
test: generate fmt vet
	go test $(GO_FLAGS) . ./provider/... ./resources/... ./utils/...

# Run terraform provider acceptance tests
testacc:
	TF_ACC=1 go test -count=1 -parallel=4 -timeout 10m -v ./...

# Build main binary
main: generate fmt vet
	go build $(GO_FLAGS) ./

install: main
	go install .

# for local development
install-terraformrc:
	cp -iv .terraformrc ~/.terraformrc && sed -i -e 's/_USERNAME_/$(USER)/g' ~/.terraformrc

# Build main binary
dist: test
	@build-tools/make-release.sh dist terraform-provider-cdp $(GO_FLAGS)
.PHONY: dist

clean:
	rm -f terraform-provider-cdp
	rm -rf dist
.PHONY: clean

# Run go fmt against code
fmt:
	go fmt . ./provider/... ./resources/... ./utils/...
	terraform fmt -recursive ./examples/

# Run go vet against code
vet:
	go vet . ./provider/... ./resources/... ./utils/...

# Generate code
generate:
	go generate . ./provider/... ./resources/... ./utils/...

tfplugindocs:
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name cdp

mod-tidy:
	go mod tidy

# Deploy
deploy: all
	cp terraform-provider-cdp ~/.terraform.d/plugins/terraform-provider-cdp
.PHONY: deploy
