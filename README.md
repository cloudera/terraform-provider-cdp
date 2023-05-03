# Terraform-provider-cdp

Implements a terraform provider for CDP resources. Terraform is the de facto tool for Infrastructure as code. This repo
implements a provider for CDP so that CDP resources (credentials, environments, datalakes, datahubs, etc) can be created using
terraform.

Target terraform version is 1.1+.

## Provider Documentation

Find detailed documentation for the provider in the [docs](./docs) folder.

Provider documentation is maintained according to [terraform guidance](https://www.terraform.io/docs/registry/providers/docs.html).

## Installation

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

## Generating Documentation

We follow the guidelines from https://developer.hashicorp.com/terraform/registry/providers/docs. Mainly we use the tool
[tfplugindocs](https://github.com/hashicorp/terraform-plugin-docs) to automatically generate the documentation from the
schemas from the provider code plus the examples under the `examples` directory.

Please read the docs at https://developer.hashicorp.com/terraform/registry/providers/docs and
https://github.com/hashicorp/terraform-plugin-docs to understand how the templating and examples
are used to render the final content. But in short, for every resource and data-source, we put example code under
`examples/resources/<resource-name>/resource.tf` and the template docs at `templates/resources/<resource-name>.md.tmpl`
(similar for data-sources).

Then you can generate the docs by running:
```
make docs
```

which will override the files under `/docs`. We check in the generated docs for every change that changes resources or
data-sources.

You can use this tool: https://registry.terraform.io/tools/doc-preview to copy-paste the markdown files to see how they
render in the terraform registry public docs.


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

