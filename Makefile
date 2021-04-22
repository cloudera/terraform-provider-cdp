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

# Build main binary
main: generate fmt vet
	go build $(GO_FLAGS) ./

install: main
	mkdir -p $(TF_PLUGIN_DIR)/$(TF_PROVIDER_NAME)/$(VERSION)/$(ARCH)/
	cp terraform-provider-cdp $(TF_PLUGIN_DIR)/$(TF_PROVIDER_NAME)/$(VERSION)/$(ARCH)/terraform-provider-cdp_v$(VERSION)

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

# Run go vet against code
vet:
	go vet . ./provider/... ./resources/... ./utils/...

# Generate code
generate:
	go generate . ./provider/... ./resources/... ./utils/...

# Deploy
deploy: all
	cp terraform-provider-cdp ~/.terraform.d/plugins/terraform-provider-cdp
.PHONY: deploy
