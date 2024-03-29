# Terraform Provider DID Key

This repo is a companion repo to the [Call APIs with Terraform Providers](https://learn.hashicorp.com/collections/terraform/providers) Learn collection.

Visit the [`boilerplate`](https://github.com/hashicorp/terraform-provider-hashicups/tree/boilerplate) branch of this repository for this Terraform provider's specific starter template. The [Terraform Provider Scaffold](https://github.com/hashicorp/terraform-provider-scaffolding) is a quick-start repository for creating a Terraform provider. Use this GitHub template when you're ready to create your own custom provider.



## Build provider

Run the following command to build the provider

```shell
$ go build -o terraform-provider-didkey
```

## Local release build

```shell
$ go install github.com/goreleaser/goreleaser@latest
```

```shell
$ make release
```

You will find the releases in the `/dist` directory. You will need to rename the provider binary to `terraform-provider-didkey` and move the binary into [the appropriate subdirectory within the user plugins directory](https://learn.hashicorp.com/tutorials/terraform/provider-use?in=terraform/providers#install-hashicups-provider).
## Test sample configuration

First, build and install the provider.

```shell
$ make install
```

Then, navigate to the `examples` directory.

```shell
$ cd examples
```

Run the following command to initialize the workspace and apply the sample configuration.

```shell
$ terraform init && terraform apply
```
