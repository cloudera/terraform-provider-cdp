# Terraform-provider-cdp

Implements a terraform provider for CDP resources. Terraform is the de facto tool for "declarative infrastructure". This repo
implements a provider for CDP so that CDP resources (credentials, environments, datalakes, datahubs, etc) can be created using
terraform. Target terraform version is 0.12.

## Usage
Please check sample code under [examples dir](./examples).

```
provider "cdp" {
  cdp_access_key_id = ""  # Optional. Can be read from ~/.cdp/credentials
  cdp_private_key = ""    # Ditto
  cdp_profile = ""        # Optional
}

resource "cdp_environments_aws_credential" "my-credentials" {
  name = "my-cloudera-cdp-credentials"
  role_arn = "arn:aws:iam::11111111111:role/my-cross-account-role"
  description = "testing the terrafrom-provider-cdp"
}
```

## CDP configuration
Using CDP terraform provider requires having an API access key id and private key for the user in CDP hereby referred
to as [CDP credentials](https://docs.cloudera.com/management-console/cloud/cli/topics/mc-cli-generating-an-api-access-key.html).
Terraform provider uses the same mechanism as with the CDP CLI to obtain the CDP credentials. Terraform CDP provider
considers these three options in order of precedence:
### 1. Setting the credentials in the provider block
The easiest way is to specify the access key id and private key in the provider configuration like this:
```
provider "cdp" {
  cdp_access_key_id = "11111111-1111-1111-1111-111111111111"
  cdp_private_key = "......................................"
}
```
However, if the terraform files will be shared, the following are more secure ways for the configuring credentials.

### 2. Using environment variables
Terraform provider respects these environment variables if they are defined:

| Env var           | Description |
| ----------------- | ----------- |
| CDP_ACCESS_KEY_ID | Access key ID for the CDP credential |
| CDP_PRIVATE_KEY   | Private key for the CDP credential |
| CDP_PROFILE       | CDP Profile to use in ~/.cdp/credentials file |

If both `CDP_ACCESS_KEY_ID` and `CDP_PRIVATE_KEY` are defined as environment variables, then the terraform provider automatically picks these up.

### 3. Shared credentials file
CDP respects a credentials file under `~/.cdp/credentials` to read the sensitive data. The credentials file is
INI formatted with sections for each profile.

If you want to specify CDP credentials in the shared file it should be formatted like this:
```
$ cat ~/.cdp/credentials
[dev]
cdp_access_key_id=11111111-1111-1111-1111-111111111111
cdp_private_key=......................................

[default]
cdp_access_key_id=11111111-1111-1111-1111-111111111111
cdp_private_key=......................................
```

You can use [this doc](https://docs.cloudera.com/management-console/cloud/cli/topics/mc-configuring-cdp-client-with-the-api-access-key.html) to configure the credentials file from the CDP CLI.

There can be any number of profiles defined in the file with separate access key ids and private keys. By default the
profile named `default` is used. To select any other profile, you can either set the environment variable `CDP_PROFILE`
or alternatively configure your terraform provider with the `cdp_profile` like this:
 ```
 provider "cdp" {
   cdp_profile = "my-profile"
 }
 ```

## Terraform provider for CDP
The bulk of the code base implements a terraform provider for CDP. CRUD operations on CDP credentials, environments,
datalakes and datahub clusters are supported.

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
