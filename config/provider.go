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
}

const networkingGroup = "networking"
const databaseGroup = "database"

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

	pc.AddResourceConfigurator("digitalocean_database_redis_config", func(r *ujconfig.Resource) {
		r.ShortGroup = databaseGroup
		r.UseAsync = false
		r.References["cluster_id"] = ujconfig.Reference{
			Type:          "Cluster",
			TerraformName: "digitalocean_database_cluster",
		}
	})
	pc.AddResourceConfigurator("digitalocean_database_mysql_config", func(r *ujconfig.Resource) {
		r.ShortGroup = databaseGroup
		r.UseAsync = false
		r.References["cluster_id"] = ujconfig.Reference{
			Type:          "Cluster",
			TerraformName: "digitalocean_database_cluster",
		}
	})
	pc.AddResourceConfigurator("digitalocean_database_kafka_topic", func(r *ujconfig.Resource) {
		r.ShortGroup = databaseGroup
		r.UseAsync = false
		r.References["cluster_id"] = ujconfig.Reference{
			Type:          "Cluster",
			TerraformName: "digitalocean_database_cluster",
		}
	})
	pc.AddResourceConfigurator("digitalocean_database_db", func(r *ujconfig.Resource) {
		r.ShortGroup = databaseGroup
		r.UseAsync = false
		r.References["cluster_id"] = ujconfig.Reference{
			Type:          "Cluster",
			TerraformName: "digitalocean_database_cluster",
		}
	})
	pc.AddResourceConfigurator("digitalocean_database_connection_pool", func(r *ujconfig.Resource) {
		r.ShortGroup = databaseGroup
		r.UseAsync = false
		r.References["cluster_id"] = ujconfig.Reference{
			Type:          "Cluster",
			TerraformName: "digitalocean_database_cluster",
		}
	})
	pc.AddResourceConfigurator("digitalocean_app", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{
			Type:          referenceType(pc, "project", "v1alpha1", "Project"),
			TerraformName: "digitalocean_project",
		}
	})
	pc.AddResourceConfigurator("digitalocean_reserved_ip_assignment", func(r *ujconfig.Resource) {
		r.ShortGroup = networkingGroup
		r.UseAsync = false
		r.References["droplet_id"] = ujconfig.Reference{
			Type:          referenceType(pc, "compute", "v1alpha1", "Droplet"),
			TerraformName: "digitalocean_droplet",
		}
		r.References["ip_address"] = ujconfig.Reference{
			Type:          "IP",
			TerraformName: "digitalocean_reserved_ip",
		}
	})
	pc.AddResourceConfigurator("digitalocean_volume_attachment", func(r *ujconfig.Resource) {
		r.ShortGroup = "volume"
		r.UseAsync = false
		r.References["volume_id"] = ujconfig.Reference{
			Type:          "Volume",
			TerraformName: "digitalocean_volume",
		}
		r.References["droplet_id"] = ujconfig.Reference{
			Type:          referenceType(pc, "compute", "v1alpha1", "Droplet"),
			TerraformName: "digitalocean_droplet",
		}
	})
	pc.AddResourceConfigurator("digitalocean_volume_snapshot", func(r *ujconfig.Resource) {
		r.UseAsync = false
		r.References["volume_id"] = ujconfig.Reference{
			Type:          "Volume",
			TerraformName: "digitalocean_volume",
		}
	})
	pc.AddResourceConfigurator("digitalocean_volume", func(r *ujconfig.Resource) {
		r.ShortGroup = "volume"
		r.UseAsync = false
		r.References["snapshot_id"] = ujconfig.Reference{
			Type:          "Snapshot",
			TerraformName: "digitalocean_volume_snapshot",
		}
	})
	pc.AddResourceConfigurator("digitalocean_uptime_alert", func(r *ujconfig.Resource) {
		r.UseAsync = false
		r.References["check_id"] = ujconfig.Reference{
			Type:          "Check",
			TerraformName: "digitalocean_uptime_check",
		}
	})
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
		r.ShortGroup = databaseGroup
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
		r.ShortGroup = databaseGroup
		r.References["cluster_id"] = ujconfig.Reference{
			Type:          "Cluster",
			TerraformName: "digitalocean_database_cluster",
		}
	})
	pc.AddResourceConfigurator("digitalocean_database_cluster", func(r *ujconfig.Resource) {
		r.ShortGroup = databaseGroup
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
		r.References["ssh_keys"] = ujconfig.Reference{
			Type:          referenceType(pc, "ssh", "v1alpha1", "Key"),
			TerraformName: "digitalocean_ssh_key",
		}
		r.References["image"] = ujconfig.Reference{
			Type:          referenceType(pc, "custom", "v1alpha1", "Image"),
			TerraformName: "digitalocean_custom_image",
		}
		// r.References["vpc_uuid"] = ujconfig.Reference{
		// 	Type:          referenceType(pc, "networking", "v1alpha1", "VPC"),
		// 	TerraformName: "digitalocean_vpc",
		// }
		// r.References["volume_ids"] = ujconfig.Reference{
		// 	Type:          referenceType(pc, "volume", "v1alpha1", "Volume"),
		// 	TerraformName: "digitalocean_volume",
		// }
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
