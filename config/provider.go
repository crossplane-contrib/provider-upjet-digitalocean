/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"fmt"

	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "digitalocean"
	modulePath     = "github.com/straw-hat-team/provider-digitalocean"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]ujconfig.ExternalName{
	"digitalocean_project":              ujconfig.IdentifierFromProvider,
	"digitalocean_droplet":              ujconfig.IdentifierFromProvider,
	"digitalocean_kubernetes_cluster":   ujconfig.IdentifierFromProvider,
	"digitalocean_kubernetes_node_pool": ujconfig.IdentifierFromProvider,
	"digitalocean_domain":               ujconfig.IdentifierFromProvider,
	"digitalocean_record":               ujconfig.IdentifierFromProvider,
	"digitalocean_vpc":                  ujconfig.IdentifierFromProvider,
	"digitalocean_spaces_bucket":        ujconfig.NameAsIdentifier,
	"digitalocean_loadbalancer":         ujconfig.IdentifierFromProvider,
	"digitalocean_database_cluster":     ujconfig.IdentifierFromProvider,
	"digitalocean_database_user":        ujconfig.NameAsIdentifier,
	"digitalocean_database_replica":     ujconfig.IdentifierFromProvider,
	"digitalocean_droplet_snapshot":     ujconfig.IdentifierFromProvider,
	"digitalocean_tag":                  ujconfig.IdentifierFromProvider,
	"digitalocean_firewall":             ujconfig.IdentifierFromProvider,
	"digitalocean_reserved_ip":          ujconfig.IdentifierFromProvider,
	"digitalocean_monitor_alert":        ujconfig.IdentifierFromProvider,
	"digitalocean_uptime_check":         ujconfig.IdentifierFromProvider,
}

const networkingGroup = "networking"

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("digitalocean.crossplane.io"),
		ujconfig.WithShortName("do"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	pc.AddResourceConfigurator("digitalocean_monitor_alert", func(r *ujconfig.Resource) {
		r.UseAsync = false
		r.References["entities"] = ujconfig.Reference{
			Type:          referenceType(pc, "compute", "v1alpha1", "Droplet"),
			TerraformName: "digitalocean_droplet",
		}
	})
	pc.AddResourceConfigurator("digitalocean_reserved_ip", func(r *ujconfig.Resource) {
		r.ShortGroup = networkingGroup
		r.UseAsync = false
		r.References["droplet_id"] = ujconfig.Reference{
			Type:          referenceType(pc, "compute", "v1alpha1", "Droplet"),
			TerraformName: "digitalocean_droplet",
		}
	})
	pc.AddResourceConfigurator("digitalocean_firewall", func(r *ujconfig.Resource) {
		r.ShortGroup = networkingGroup
		r.UseAsync = false
		r.References["droplet_ids"] = ujconfig.Reference{
			Type:          referenceType(pc, "compute", "v1alpha1", "Droplet"),
			TerraformName: "digitalocean_droplet",
		}
	})
	pc.AddResourceConfigurator("digitalocean_droplet_snapshot", func(r *ujconfig.Resource) {
		r.UseAsync = false
		r.References["droplet_id"] = ujconfig.Reference{
			Type:          referenceType(pc, "compute", "v1alpha1", "Droplet"),
			TerraformName: "digitalocean_droplet",
		}
	})
	pc.AddResourceConfigurator("digitalocean_database_replica", func(r *ujconfig.Resource) {
		r.ShortGroup = "database"
		r.UseAsync = false
		r.References["cluster_id"] = ujconfig.Reference{
			Type:          "Cluster",
			TerraformName: "digitalocean_database_cluster",
		}
		r.References["private_network_uuid"] = ujconfig.Reference{
			Type:          referenceType(pc, "networking", "v1alpha1", "VPC"),
			TerraformName: "digitalocean_vpc",
		}
	})
	pc.AddResourceConfigurator("digitalocean_database_user", func(r *ujconfig.Resource) {
		r.ShortGroup = "database"
		r.References["cluster_id"] = ujconfig.Reference{
			Type:          "Cluster",
			TerraformName: "digitalocean_database_cluster",
		}
	})
	pc.AddResourceConfigurator("digitalocean_database_cluster", func(r *ujconfig.Resource) {
		r.ShortGroup = "database"
		r.References["private_network_uuid"] = ujconfig.Reference{
			Type:          referenceType(pc, "networking", "v1alpha1", "VPC"),
			TerraformName: "digitalocean_vpc",
		}
		r.References["project_id"] = ujconfig.Reference{
			Type:          referenceType(pc, "project", "v1alpha1", "Project"),
			TerraformName: "digitalocean_project",
		}
	})
	pc.AddResourceConfigurator("digitalocean_spaces_bucket", func(r *ujconfig.Resource) {
		r.ShortGroup = "storage"
	})
	pc.AddResourceConfigurator("digitalocean_project", func(r *ujconfig.Resource) {
		r.ShortGroup = "project"
	})
	pc.AddResourceConfigurator("digitalocean_loadbalancer", func(r *ujconfig.Resource) {
		r.ShortGroup = networkingGroup
		r.References["droplet_ids"] = ujconfig.Reference{
			Type:          referenceType(pc, "compute", "v1alpha1", "Droplet"),
			TerraformName: "digitalocean_droplet",
		}
		r.References["vpc_uuid"] = ujconfig.Reference{
			Type:          "VPC",
			TerraformName: "digitalocean_vpc",
		}
	})
	pc.AddResourceConfigurator("digitalocean_vpc", func(r *ujconfig.Resource) {
		r.ShortGroup = networkingGroup
	})
	pc.AddResourceConfigurator("digitalocean_domain", func(r *ujconfig.Resource) {
		r.ShortGroup = networkingGroup
	})
	pc.AddResourceConfigurator("digitalocean_record", func(r *ujconfig.Resource) {
		r.ShortGroup = networkingGroup
		r.References["domain"] = ujconfig.Reference{
			Type:          "Domain",
			TerraformName: "digitalocean_domain",
		}
	})
	pc.AddResourceConfigurator("digitalocean_droplet", func(r *ujconfig.Resource) {
		r.ShortGroup = "compute"
		r.UseAsync = false
		r.References["tags"] = ujconfig.Reference{
			Type:          referenceType(pc, "digitalocean", "v1alpha1", "Tag"),
			TerraformName: "digitalocean_tag",
		}
	})
	pc.AddResourceConfigurator("digitalocean_kubernetes_cluster", func(r *ujconfig.Resource) {
		r.ShortGroup = "kubernetes"
		r.UseAsync = false
	})
	pc.AddResourceConfigurator("digitalocean_kubernetes_node_pool", func(r *ujconfig.Resource) {
		r.ShortGroup = "kubernetes"
		r.UseAsync = false
		r.References["cluster_id"] = ujconfig.Reference{
			Type:          "Cluster",
			TerraformName: "digitalocean_kubernetes_cluster",
		}
	})

	pc.ConfigureResources()
	return pc
}

//nolint:unparam
func referenceType(pc *ujconfig.Provider, shortGroup string, version string, structTypeName string) string {
	return fmt.Sprintf("%s/apis/%s/%s.%s", pc.ModulePath, shortGroup, version, structTypeName)
}
