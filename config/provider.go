/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/registry/reference"
	"github.com/digitalocean/terraform-provider-digitalocean/digitalocean"
	"github.com/straw-hat-team/provider-digitalocean/config/project"
)

const (
	resourcePrefix = "digitalocean"
	modulePath     = "github.com/straw-hat-team/provider-digitalocean"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("digitalocean.crossplane.io"),
		ujconfig.WithShortName("digitalocean"),
		//ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithIncludeList(ExternalNameConfigured(cliReconciledExternalNameConfigs)),
		ujconfig.WithTerraformPluginSDKIncludeList(ExternalNameConfigured(terraformPluginSDKExternalNameConfigs)),
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithTerraformProvider(digitalocean.Provider()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		project.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
