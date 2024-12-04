# Contributing

## How-To

### Getting Started

1. Clone the repository.
2. Run `make submodules` to initialize the submodules.
3. Run `make create-secret-yaml` to create the secret yaml file.
4. Fix the configuration of `./examples/providerconfig/secret.yaml` manifest.
5. Run `make k-apply-requires` to apply the required resources.

### Run the local DigitalOcean Provider

1. Run `make run-all` to run the local DigitalOcean provider.

### Setting Up Local Crossplane with Helm

1. Run `make helm-install-crossplane` to install Crossplane.

### Add new resource

1. Open `config/provider.go` file.
   1. Find `ExternalNameConfigs` variable.
   1. Add new resource terraform resource name to the `ExternalNameConfigs` map.
   1. Find `GetProvider` function.
   1. Add a new `AddResourceConfigurator` function call with the new resource name using
      1. Add the `r.ShortGroup` to the resource.
      1. Add (if any) all the `r.References` to the resource.
1. Run `make generate` to generate the new resource configuration.
1. Run `make k-apply-crds` to apply the new CRDS.

#### Test the new resource

1. Add Kubernetes resources under `./tmp/manifests` directory.
2. Run `make k-apply-tmp` to apply all the resources under `./tmp/manifests` directory.

Or to be safe, you call apply all the required resources at once:

1. `k-apply-all`

### Releasing a New Version

1. [Create Release Branch](#creating-a-release-branch) or [Patch a Release Branch](#patching-a-release-branch) the
  [Release Branch](#release-branches).
1. Visit [Tag Workflow Action](https://github.com/crossplane-contrib/provider-upjet-digitalocean/actions/workflows/tag.yaml)
and click on the `Run workflow` button. Make sure the "Branch" is set to the release branch, e.g., `release-0.2`.
Enter the version number in the "Release version" field, e.g., `v0.2.1`, and a sensible value for the "Tag message"
field. Verify the information and click the "Run workflow" button.
1. Wait for the Tag Workflow Action to complete.
1. Run the [CI Workflow Action](https://github.com/crossplane-contrib/provider-upjet-digitalocean/actions/workflows/ci.yml)
   by clicking on the "Run workflow" button. Make sure the "Branch" is set to the release branch, e.g., `release-0.2`.
1. Create a new GitHub Release
   at [New Release](https://github.com/crossplane-contrib/provider-upjet-digitalocean/releases/new).
   Select the tag that was created by the Tag Workflow Action, e.g., `v0.0.1`, and enter the release notes. Make sure the Target branch is set to the release branch, e.g., `release-0.2`. Click the "Publish release" button.

## Explanations

### Release Branches

Release branches are prefixed with `release-` followed by the version number, e.g., `release-0.2`. The version number
should be composed by the major and minor version numbers, e.g., `0.2`. The patch version number should never be part
of the release branch name.

Reuse the existing release branch, backport fixes to the existing major/minor branch.

#### Creating a Release Branch

Create or patch the [Release Branch](#release-branches) from the `main` branch. You could use visit [Branches](https://github.com/crossplane-contrib/provider-upjet-digitalocean/branches)
and click on the `New branch` button, or use the following commands:

```shell
git checkout -b release-<MAJOR>.<MINOR>
git push origin release-<MAJOR>.<MINOR>
```

#### Patching a Release Branch

To apply a patch to a release branch, you should tag the pull request with the
backport label, eg. `Backport release-0.2`. The pull request will be merged into
the release branch.

### Make Commands

Run code-generation pipeline:

```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

### Local Development with IntelliJ IDEA

When developing locally with IntelliJ IDEA, verify that the following environment variables are set:

```shell
TERRAFORM_PROVIDER_SOURCE=<copy values from Makefile>
TERRAFORM_PROVIDER_VERSION=<copy values from Makefile>
TERRAFORM_VERSION=<copy values from Makefile>
UPBOUND_CONTEXT=local
```
