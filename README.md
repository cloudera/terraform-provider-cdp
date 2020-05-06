# terraform-provider-cdp

Implements a terraform provider for CDP resources. Terraform is the de facto tool for "declerative infrastructure". This repo
implements a provider for CDP so that CDP resources (credentials, environments, datalakes, datahubs, etc) can be created using
terraform. Target terraform version is 0.12.

Example
```
provider "cdp" {
  access_key_id = ""  # Optional. Can be read from ~/.cdp/credentials
  private_key = ""    # Ditto
  profile = ""        # Ditto
}

resource "cdp_environments_credential" "my-credentials" {
  name = "my-cloudera-cdp-credentials"
  cloud_platform = "AWS"
  role_arn = "arn:aws:iam::11111111111:role/my-cross-account-role"
  description = "testing the terrafrom-provider-cdp"
}
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

## terraform provider for CDP
Rest of this code base implements a terraform provider for CDP. There is only 1 resouce so far which is the credentials
resource.

### Compiling
```
make
```

### Execute example terraform
```
cd examples/credential
make terraform-apply
```

More documentation about writing terraform providers is here: https://www.terraform.io/docs/extend/writing-custom-providers.html.
