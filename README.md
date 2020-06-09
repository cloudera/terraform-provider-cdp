# Terraform-provider-cdp

Implements a terraform provider for CDP resources. Terraform is the de facto tool for "declarative infrastructure". This repo
implements a provider for CDP so that CDP resources (credentials, environments, datalakes, datahubs, etc) can be created using
terraform. Target terraform version is 0.12.

## Provider Documentation

Find detailed documentation for the provider in the [docs](./docs) folder.

Provider documentation is maintained according to [terraform guidance](https://www.terraform.io/docs/registry/providers/docs.html).

## Development

Find documentation about writing terraform providers is here: https://www.terraform.io/docs/extend/writing-custom-providers.html.

### Compiling
```
make
```

### Execute example terraform
```
cd examples/credential
make terraform-apply
```

## cdp-sdk-go
Since CDP does not have an official golang SDK, this repo contains a subdirectory named "cdp-sdk-go" that implements a pure
golang SDK using the swagger definitions from the public API. go-swagger is used to generate the client code. This SDK is
independent of any terraform related logic, and ideally should be hosted elsewhere as a stand alone library.

### Compiling cdp-sdk-go
```
cd cdp-sdk-go
make
```

### Running swagger code gen
```
cd cdp-sdk-go
make swagger-gen
```
