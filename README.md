# Terraform-provider-cdp

Implements a terraform provider for CDP resources. Terraform is the de facto tool for "declarative infrastructure". This repo
implements a provider for CDP so that CDP resources (credentials, environments, datalakes, datahubs, etc) can be created using
terraform.

Target terraform version is 0.13 although it should also work with 0.12. Older versions are not supported.

## Provider Documentation

Find detailed documentation for the provider in the [docs](./docs) folder.

Provider documentation is maintained according to [terraform guidance](https://www.terraform.io/docs/registry/providers/docs.html).

## Installation

The installation instructions for custom terraform providers was changed in a backwards incompatible way in
Terraform-0.13. See [Upgrading to Terraform v0.13](https://www.terraform.io/upgrade-guides/0-13.html).

## Terraform-0.13
The easiest way to install the CDP provider from source code is by
```
  make install
```

which copies the binary to the directory: `~/.terraform.d/plugins/terraform.cloudera.com/cloudera/cdp/$VERSION/$ARCH/terraform-provider-cdp_v$VERSION`.

If you have downloaded a binary release, you can execute these steps to install:
```
mkdir -p ~/.terraform.d/plugins/terraform.cloudera.com/cloudera/cdp/$VERSION/$ARCH
cp terraform-provider-cdp ~/.terraform.d/plugins/terraform.cloudera.com/cloudera/cdp/$VERSION/$ARCH/terraform-provider-cdp_v$VERSION
```

where VERSION should be replaced with something like `0.0.3`
and ARCH should be replaced with something like `darwin_amd64` or `linux_amd64`

### Terraform-0.12
The plugin binary must be placed in ~/.terraform.d/plugins/terraform-provider-cdp.

If downloading the binary from a github release, make sure to select the platform appropriate binary and to rename it to terraform-provider-cdp when placing it in ~/.terraform.d/plugins/terraform-provider-cdp.

Any time the plugin binary is changed, you will need to re-run `terraform init`.

## Filing Bugs

File JIRAs using the CDPCP JIRA project using the cdp terraform provider component.

Generate logs by setting the TF_LOG environment variable to any value and capturing the output, for example by running `TF_LOG=true terraform apply plan.txt > tf.log 2>&1`. Please attach the output file to the filed bug. Please also attach the crash.log file from any terraform crash. See the [terraform docs](https://www.terraform.io/docs/internals/debugging.html) for more debugging information.

## Development

Find documentation about writing terraform providers is here: https://www.terraform.io/docs/extend/writing-custom-providers.html.

## Contributing

All code and documentation contributions are welcome. Please fork the repository and send a pull request for your code
change. If you want to review patches you can add yourself in the [Code Owners file](.github/CODEOWNERS). You can also
use that file to ping people about the code reviews.

Pull requests are expected to include appropriate updates to the [change log](./CHANGELOG.md) as detailed [here](https://www.terraform.io/docs/extend/best-practices/versioning.html#changelog-specification).

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

## Releases

We aim to follow [official guidance](https://www.terraform.io/docs/extend/best-practices/versioning.html)
on versioning our provider. See the [change log](./CHANGELOG.md) for the release
history.

To make a new release:

* Review the change log and update it as necessary. Ideally the content is
  up-to-date as it has been maintained along the way. Note the data of the
  release and create a new, empty unreleased entry at the top.
* Tag the commit with an appropriate semantic version, e.g. `git tag v0.0.1`.
* Push the tag: `git push --tags`.
* Create a github release off the tag.
 * The title of the release should be `tag (Month Day, Year)`.
 * The change log entry for the release should be copied as the description.
 * Build distributable artifacts by running `make clean dist`. Attach the binaries and
   checksum artifacts to the release for each supported platform.
 * Until otherwise noted, select the pre-release checkbox to indicate that we
   are not yet production ready.
 * Once the release is done, send a PR to update the `CHANGELOG.md` with the new
   release section, and update the release date.

