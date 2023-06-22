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

all: check-go test main

check-go:
ifndef GOPATH
	$(error GOPATH not defined, please define GOPATH. Run "go help gopath" to learn more about GOPATH)
endif

# Run tests
test: generate fmt vet
	go test $(GO_FLAGS) ./...

# Run tests with coverage
test-with-coverage: generate fmt vet
	go test -v -coverprofile coverage.out ./...
	go tool cover -html coverage.out -o coverage.html

# Run terraform provider acceptance tests
testacc:
	TF_ACC=1 go test -count=1 -parallel=4 -timeout 10m -v ./...

# Build main binary
main: build

build: generate fmt vet
	go build $(GO_FLAGS) ./

# See https://golangci-lint.run/
lint:
	golangci-lint run

install: main
	go install .

# for local development
install-terraformrc:
	cp -iv .terraformrc ~/.terraformrc && sed -i -e 's/_USERNAME_/$(USER)/g' ~/.terraformrc

# Make a release
release: test testacc docs
	@goreleaser release --clean

# Make a local snapshot release
release-snapshot: test
	@goreleaser release --snapshot --clean

clean:
	rm -f terraform-provider-cdp
	rm -rf dist

# Run go fmt against code
fmt:
	go fmt ./...
	# terraform fmt -recursive ./examples/  TODO: Re-enable this in CDPCP-9174

# Run go vet against code
vet:
	go vet ./...

# Generate code
generate:
	go generate ./...

docs:
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate \
		--tf-version 1.4.5 \
		--rendered-provider-name CDP \
		--website-source-dir templates

mod-tidy:
	go mod tidy

# Deploy
deploy: all
	cp terraform-provider-cdp ~/.terraform.d/plugins/terraform-provider-cdp

.PHONY: all check-go docs deploy test mod-tidy generate vet fmt clean release release-snapshot install-terraformrc install main build
