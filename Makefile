##
# Copyright 2020 Cloudera, Inc.
##

GO_FLAGS:=""

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

# Run main binary
run: generate fmt vet
	go run $(GO_FLAGS) ./main.go

# Run go fmt against code
fmt:
	go fmt . ./provider/... ./resources/... ./utils/...

# Run go vet against code
vet:
	go vet . ./provider/... ./resources/... ./utils/...

# Generate code
generate:
	go generate . ./provider/... ./resources/... ./utils/...

