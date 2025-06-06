/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "digitalocean"
	modulePath     = "github.com/crossplane-contrib/provider-upjet-digitalocean"
)

const networkingShortGroup = "networking"

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]ujconfig.ExternalName{
	"digitalocean_project":                               ujconfig.IdentifierFromProvider,
	"digitalocean_droplet":                               ujconfig.IdentifierFromProvider,
	"digitalocean_kubernetes_cluster":                    ujconfig.IdentifierFromProvider,
	"digitalocean_kubernetes_node_pool":                  ujconfig.IdentifierFromProvider,
	"digitalocean_domain":                                ujconfig.IdentifierFromProvider,
	"digitalocean_record":                                ujconfig.IdentifierFromProvider,
	"digitalocean_vpc":                                   ujconfig.IdentifierFromProvider,
	"digitalocean_spaces_bucket":                         ujconfig.NameAsIdentifier,
	"digitalocean_cdn":                                   ujconfig.IdentifierFromProvider,
	"digitalocean_loadbalancer":                          ujconfig.IdentifierFromProvider,
	"digitalocean_database_cluster":                      ujconfig.IdentifierFromProvider,
	"digitalocean_database_user":                         ujconfig.NameAsIdentifier,
	"digitalocean_database_replica":                      ujconfig.IdentifierFromProvider,
	"digitalocean_droplet_snapshot":                      ujconfig.IdentifierFromProvider,
	"digitalocean_tag":                                   ujconfig.IdentifierFromProvider,
	"digitalocean_firewall":                              ujconfig.IdentifierFromProvider,
	"digitalocean_reserved_ip":                           ujconfig.IdentifierFromProvider,
	"digitalocean_reserved_ip_assignment":                ujconfig.IdentifierFromProvider,
	"digitalocean_monitor_alert":                         ujconfig.IdentifierFromProvider,
	"digitalocean_uptime_check":                          ujconfig.IdentifierFromProvider,
	"digitalocean_uptime_alert":                          ujconfig.IdentifierFromProvider,
	"digitalocean_volume":                                ujconfig.IdentifierFromProvider,
	"digitalocean_volume_snapshot":                       ujconfig.IdentifierFromProvider,
	"digitalocean_volume_attachment":                     ujconfig.IdentifierFromProvider,
	"digitalocean_ssh_key":                               ujconfig.IdentifierFromProvider,
	"digitalocean_spaces_bucket_cors_configuration":      ujconfig.IdentifierFromProvider,
	"digitalocean_spaces_bucket_object":                  ujconfig.IdentifierFromProvider,
	"digitalocean_spaces_bucket_policy":                  ujconfig.IdentifierFromProvider,
	"digitalocean_custom_image":                          ujconfig.IdentifierFromProvider,
	"digitalocean_certificate":                           ujconfig.IdentifierFromProvider,
	"digitalocean_app":                                   ujconfig.IdentifierFromProvider,
	"digitalocean_container_registry":                    ujconfig.IdentifierFromProvider,
	"digitalocean_container_registry_docker_credentials": ujconfig.IdentifierFromProvider,
	"digitalocean_database_connection_pool":              ujconfig.IdentifierFromProvider,
	"digitalocean_database_db":                           ujconfig.IdentifierFromProvider,
	"digitalocean_database_kafka_topic":                  ujconfig.IdentifierFromProvider,
	"digitalocean_database_mysql_config":                 ujconfig.IdentifierFromProvider,
	"digitalocean_database_redis_config":                 ujconfig.IdentifierFromProvider,
	"digitalocean_database_firewall":                     ujconfig.IdentifierFromProvider,
	"digitalocean_project_resources":                     ujconfig.IdentifierFromProvider,
}

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

	pc.AddResourceConfigurator("digitalocean_spaces_cors_configuration", func(r *ujconfig.Resource) {
		r.References["region"] = ujconfig.Reference{
			TerraformName:     "digitalocean_spaces_bucket",
			SelectorFieldName: "Region",
		}
		r.References["bucket"] = ujconfig.Reference{
			TerraformName: "digitalocean_spaces_bucket",
		}
	})
	pc.AddResourceConfigurator("digitalocean_certificate", func(r *ujconfig.Resource) {
		r.ShortGroup = networkingShortGroup
		r.References["domains"] = ujconfig.Reference{
			SelectorFieldName: "ID",
			TerraformName:     "digitalocean_domain",
		}
	})
	pc.AddResourceConfigurator("digitalocean_project", func(r *ujconfig.Resource) {
		r.ShortGroup = "project"
	})
	pc.AddResourceConfigurator("digitalocean_cdn", func(r *ujconfig.Resource) {
		r.ShortGroup = "spaces"
		r.References["custom_domain"] = ujconfig.Reference{
			SelectorFieldName: "ID",
			TerraformName:     "digitalocean_domain",
		}
		r.References["certificate_name"] = ujconfig.Reference{
			SelectorFieldName: "Name",
			TerraformName:     "digitalocean_certificate",
		}
		r.References["origin"] = ujconfig.Reference{
			SelectorFieldName: "BucketDomainName",
			TerraformName:     "digitalocean_spaces_bucket",
		}
	})
	pc.AddResourceConfigurator("digitalocean_project_resources", func(r *ujconfig.Resource) {
		r.References["project"] = ujconfig.Reference{
			TerraformName: "digitalocean_project",
		}
	})
	pc.AddResourceConfigurator("digitalocean_database_firewall", func(r *ujconfig.Resource) {
		r.References["cluster_id"] = ujconfig.Reference{
			TerraformName: "digitalocean_database_cluster",
		}
	})
	pc.AddResourceConfigurator("digitalocean_database_redis_config", func(r *ujconfig.Resource) {
		r.References["cluster_id"] = ujconfig.Reference{
			TerraformName: "digitalocean_database_cluster",
		}
	})
	pc.AddResourceConfigurator("digitalocean_database_mysql_config", func(r *ujconfig.Resource) {
		r.References["cluster_id"] = ujconfig.Reference{
			TerraformName: "digitalocean_database_cluster",
		}
	})
	pc.AddResourceConfigurator("digitalocean_database_kafka_topic", func(r *ujconfig.Resource) {
		r.References["cluster_id"] = ujconfig.Reference{
			TerraformName: "digitalocean_database_cluster",
		}
	})
	pc.AddResourceConfigurator("digitalocean_database_db", func(r *ujconfig.Resource) {
		r.References["cluster_id"] = ujconfig.Reference{
			TerraformName: "digitalocean_database_cluster",
		}
	})
	pc.AddResourceConfigurator("digitalocean_database_connection_pool", func(r *ujconfig.Resource) {
		r.References["cluster_id"] = ujconfig.Reference{
			TerraformName: "digitalocean_database_cluster",
		}
	})
	pc.AddResourceConfigurator("digitalocean_app", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{
			TerraformName: "digitalocean_project",
		}
	})
	pc.AddResourceConfigurator("digitalocean_reserved_ip_assignment", func(r *ujconfig.Resource) {
		r.ShortGroup = networkingShortGroup
		r.References["droplet_id"] = ujconfig.Reference{
			TerraformName: "digitalocean_droplet",
		}
		r.References["ip_address"] = ujconfig.Reference{
			TerraformName: "digitalocean_reserved_ip",
		}
	})
	pc.AddResourceConfigurator("digitalocean_volume_attachment", func(r *ujconfig.Resource) {
		r.References["volume_id"] = ujconfig.Reference{
			TerraformName: "digitalocean_volume",
		}
		r.References["droplet_id"] = ujconfig.Reference{
			TerraformName: "digitalocean_droplet",
		}
	})
	pc.AddResourceConfigurator("digitalocean_volume_snapshot", func(r *ujconfig.Resource) {
		r.References["volume_id"] = ujconfig.Reference{
			TerraformName: "digitalocean_volume",
		}
	})
	pc.AddResourceConfigurator("digitalocean_volume", func(r *ujconfig.Resource) {
		r.ShortGroup = "volume"
		r.References["snapshot_id"] = ujconfig.Reference{
			TerraformName: "digitalocean_volume_snapshot",
		}
	})
	pc.AddResourceConfigurator("digitalocean_uptime_alert", func(r *ujconfig.Resource) {
		r.References["check_id"] = ujconfig.Reference{
			TerraformName: "digitalocean_uptime_check",
		}
	})
	pc.AddResourceConfigurator("digitalocean_monitor_alert", func(r *ujconfig.Resource) {
		r.References["entities"] = ujconfig.Reference{
			TerraformName: "digitalocean_droplet",
		}
	})
	pc.AddResourceConfigurator("digitalocean_reserved_ip", func(r *ujconfig.Resource) {
		r.ShortGroup = networkingShortGroup
		r.References["droplet_id"] = ujconfig.Reference{
			TerraformName: "digitalocean_droplet",
		}
	})
	pc.AddResourceConfigurator("digitalocean_vpc", func(r *ujconfig.Resource) {
		r.ShortGroup = "vpc"
	})
	pc.AddResourceConfigurator("digitalocean_firewall", func(r *ujconfig.Resource) {
		r.ShortGroup = networkingShortGroup
		r.References["droplet_ids"] = ujconfig.Reference{
			TerraformName: "digitalocean_droplet",
		}
	})
	pc.AddResourceConfigurator("digitalocean_droplet_snapshot", func(r *ujconfig.Resource) {
		r.References["droplet_id"] = ujconfig.Reference{
			TerraformName: "digitalocean_droplet",
		}
	})
	pc.AddResourceConfigurator("digitalocean_database_replica", func(r *ujconfig.Resource) {
		r.References["cluster_id"] = ujconfig.Reference{
			TerraformName: "digitalocean_database_cluster",
		}
		r.References["private_network_uuid"] = ujconfig.Reference{
			TerraformName: "digitalocean_vpc",
		}
	})
	pc.AddResourceConfigurator("digitalocean_database_user", func(r *ujconfig.Resource) {
		r.References["cluster_id"] = ujconfig.Reference{
			TerraformName: "digitalocean_database_cluster",
		}
	})
	pc.AddResourceConfigurator("digitalocean_database_cluster", func(r *ujconfig.Resource) {
		r.References["private_network_uuid"] = ujconfig.Reference{
			TerraformName: "digitalocean_vpc",
		}
		r.References["project_id"] = ujconfig.Reference{
			TerraformName: "digitalocean_project",
		}
	})
	pc.AddResourceConfigurator("digitalocean_loadbalancer", func(r *ujconfig.Resource) {
		r.ShortGroup = networkingShortGroup
		r.References["droplet_ids"] = ujconfig.Reference{
			TerraformName: "digitalocean_droplet",
		}
		r.References["vpc_uuid"] = ujconfig.Reference{
			TerraformName: "digitalocean_vpc",
		}
		r.References["project_id"] = ujconfig.Reference{
			TerraformName: "digitalocean_project",
		}
		r.References["forwarding_rule.certificate_name"] = ujconfig.Reference{
			TerraformName: "digitalocean_certificate",
		}
	})
	pc.AddResourceConfigurator("digitalocean_domain", func(r *ujconfig.Resource) {
		r.ShortGroup = "dns"
	})
	pc.AddResourceConfigurator("digitalocean_record", func(r *ujconfig.Resource) {
		r.ShortGroup = "dns"
		r.References["domain"] = ujconfig.Reference{
			TerraformName: "digitalocean_domain",
		}
	})
	pc.AddResourceConfigurator("digitalocean_droplet", func(r *ujconfig.Resource) {
		r.ShortGroup = "droplet"
		r.References["tags"] = ujconfig.Reference{
			TerraformName: "digitalocean_tag",
		}
		r.References["ssh_keys"] = ujconfig.Reference{
			TerraformName: "digitalocean_ssh_key",
		}
		r.References["image"] = ujconfig.Reference{
			TerraformName: "digitalocean_custom_image",
		}
		r.References["vpc_uuid"] = ujconfig.Reference{
			TerraformName: "digitalocean_vpc",
		}
		// README: We can not enable a reference to volume_ids because it would create a circular dependency.
		// What that means in practice is that users must use `digitalocean_volume_attachment` to attach volumes to
		// droplets. Terraform documentation already mention that you can not use ``digitalocean_volume_attachment` and
		// `digitalocean_droplet.volume_ids` at the same time since it will cause constantly drift.
		// TODO: ask Crossplane team if there is a way to hide a given resource attribute. That way we remove the ability
		//  to set `volume_ids` attribute.
		// r.References["volume_ids"] = ujconfig.Reference{}
	})
	pc.AddResourceConfigurator("digitalocean_kubernetes_cluster", func(r *ujconfig.Resource) {
		r.References["vpc_uuid"] = ujconfig.Reference{
			TerraformName: "digitalocean_vpc",
		}
		// r.References["tags"] = ujconfig.Reference{
		// 	TerraformName: "digitalocean_tag",
		// 	SelectorFieldName: "Name",
		// }
	})
	pc.AddResourceConfigurator("digitalocean_kubernetes_node_pool", func(r *ujconfig.Resource) {
		r.References["cluster_id"] = ujconfig.Reference{
			TerraformName: "digitalocean_kubernetes_cluster",
		}
	})

	pc.ConfigureResources()
	return pc
}
