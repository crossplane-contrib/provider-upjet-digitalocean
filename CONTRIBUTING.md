# Contributing

## How-To

### Run the local DigitalOcean Provider

1. Run `make run-all` to run the local DigitalOcean provider.

### Setting Up Local Crossplane with Helm

1. Run `make helm-install-crossplane` to install Crossplane.

### Add new resource

1. Open `config/provider.go` file.
   1. Find `ExternalNameConfigs` variable.
   2. Add new resource terraform resource name to the `ExternalNameConfigs` map.
   3. Find `GetProvider` function.
   4. Add a new `AddResourceConfigurator` function call with the new resource name using
      1. Add the `r.ShortGroup` to the resource.
      2. Add (if any) all the `r.References` to the resource.
2. Run `make generate` to generate the new resource configuration.
3. Run `make k-apply-crds` to apply the new CRDS.

#### Test the new resource

1. Add Kubernetes resources under `./tmp` directory.
2. Run `make k-apply-tmp` to apply all the resources under `./tmp` directory.

Or to be safe, you call apply all the required resources at once:

1. `k-apply-all`

## Explanations

### Local Development with IntelliJ IDEA

When developing locally with IntelliJ IDEA, verify that the following environment variables are set:

```shell
TERRAFORM_PROVIDER_SOURCE=<copy values from Makefile>
TERRAFORM_PROVIDER_VERSION=<copy values from Makefile>
TERRAFORM_VERSION=<copy values from Makefile>
UPBOUND_CONTEXT=local
```
