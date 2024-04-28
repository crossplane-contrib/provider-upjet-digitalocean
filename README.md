# Provider DigitalOcean

`provider-digitalocean` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
DigitalOcean API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/straw-hat-team/provider-digitalocean):
```
up ctp provider install straw-hat-team/provider-digitalocean:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-digitalocean
spec:
  package: straw-hat-team/provider-digitalocean:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/straw-hat-team/provider-digitalocean).

## Developing

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

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/straw-hat-team/provider-digitalocean/issues).

## Completeness List

- [ ] digitalocean_app
- [x] digitalocean_cdn
- [x] digitalocean_certificate
- [ ] digitalocean_container_registry
- [ ] digitalocean_container_registry_docker_credentials
- [x] digitalocean_custom_image
- [x] digitalocean_database_cluster
- [ ] digitalocean_database_connection_pool
- [ ] digitalocean_database_db
- [ ] digitalocean_database_firewall
- [ ] digitalocean_database_kafka_topic
- [ ] digitalocean_database_mysql_config
- [ ] digitalocean_database_redis_config
- [x] digitalocean_database_replica
- [x] digitalocean_database_user
- [x] digitalocean_domain
- [x] digitalocean_droplet
- [x] digitalocean_droplet_snapshot
- [x] digitalocean_firewall
- [x] digitalocean_kubernetes_cluster
- [x] digitalocean_kubernetes_node_pool
- [x] digitalocean_loadbalancer
- [x] digitalocean_monitor_alert
- [x] digitalocean_project
- [ ] digitalocean_project_resources
- [x] digitalocean_record
- [x] digitalocean_reserved_ip
- [x] digitalocean_reserved_ip_assignment
- [x] digitalocean_spaces_bucket
- [x] digitalocean_spaces_bucket_cors_configuration
- [x] digitalocean_spaces_bucket_object
- [x] digitalocean_spaces_bucket_policy
- [x] digitalocean_ssh_key
- [x] digitalocean_tag
- [x] digitalocean_uptime_alert
- [x] digitalocean_uptime_check
- [x] digitalocean_volume
- [x] digitalocean_volume_attachment
- [x] digitalocean_volume_snapshot
- [x] digitalocean_vpc
