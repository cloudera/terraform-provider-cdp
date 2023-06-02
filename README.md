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

which copies the binary to `$GOPATH/bin`. After that you need to run:
```
install-terraformrc
```
which installs a `.terraformrc` file under your home directory to point to the locally
installed version of the provider binary.

If you have downloaded a binary release, you can execute these steps to install:
```
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/cloudera/cdp/$VERSION/$ARCH
cp terraform-provider-cdp ~/.terraform.d/plugins/registry.terraform.io/cloudera/cdp/$VERSION/$ARCH/terraform-provider-cdp_v$VERSION
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
on versioning our provider. See the [change log](./CHANGELOG.md) for the release history.

We use the [goreleaser](https://goreleaser.com/) tool to build and publish official releases. Please follow the 
[Quick Start Guide](https://goreleaser.com/quick-start/) to familiarize yourself with the tool. 

We publish new releases following the [Terraform Publishing Guide](https://developer.hashicorp.com/terraform/registry/providers/publishing)
to the Terraform Registry.

#### One time setup
1. If you have not done so, create a personal github token here: https://github.infra.cloudera.com/settings/tokens
2. `export GITHUB_TOKEN=<PLACE_THE_TOKEN>`
3. https://developer.hashicorp.com/terraform/registry/providers/publishing#preparing-and-adding-a-signing-key
4. https://docs.github.com/en/authentication/managing-commit-signature-verification/generating-a-new-gpg-key

#### Publishing new releases
1. Make sure that the build is fine, and unit tests and Terraform acceptance tests are running fine.
2. Make sure that the NOTICE file is up to date and re-export and re-commit the NOTICE using the tool.
3. Review the change log and update it as necessary. Ideally the content is
  up-to-date as it has been maintained along the way. Note the data of the
  release and create a new, empty unreleased entry at the top.
4. Set the GPG fingerprint to use to sign the release `export GPG_FINGERPRINT=<YOUR_CODE_SIGNING_GPG_KEY_ID>`. Use (`gpg --list-secret-keys --keyid-format=long` to find out).
5. Cache the password for your GPG private key with `echo "foo" | gpg --armor --detach-sign --default-key $GPG_FINGERPRINT` (GoReleaser does not support signing binaries with a GPG key that requires a passphrase. Some systems may cache your GPG passphrase for a few minutes).
6. Tag the commit with an appropriate semantic version, e.g. `git tag v0.0.1`.
   1. We use [Semantic Versioning](https://semver.org/) to mark the releases and `v` prefix is mandatory for terraform providers
   2. A release-candidate can be pushed by adding `-rc1` suffix like `v0.0.1-rc1`.
   3. You can find the next version to use by looking at the existing releases / tags.
7. Push the tag: `git push origin v0.0.1`.
8. Run `goreleaser`: `make release`
9. If goreleaser runs successfully, it will automatically:
   1. Cross-compile against all platforms and create binaries under `dist/`
   2. Create zip archives for all binaries.
   3. Checksums all of the binaries using sha256 and saves the checksums under `dist/terraform-provider-cdp_<VERSION>_SHA256SUMS`.
   4. Signs the checksums file with the gpg keys of the user.
   5. Creates other metadata files for the build and release.
   6. Creates a release **as a draft** in Github (we are intentionally doing this. Once we get the other mechanics working we can do non-draft releases).
   7. Uploads artifacts and release notes to the Github release.
10. Until otherwise noted, select the pre-release checkbox to indicate that we
  are not yet production ready.
11. The release that is pushed by goreleaser is a draft release. Go to the release page in Github, and double check the release notes, artifacts and the version. If everything is fine, click on "Edit" and then "Publish Release" button.
12. Once the release is done, send a PR to update the `CHANGELOG.md` with the new release section, and update the release date.

#### Publishing new releases to Terraform Registry
Above staps only publish the artifacts to github. We need to futher publish the artifacts to Terraform Registry. The steps will be documented here.
