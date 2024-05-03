// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	registry "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/container/registry"
	registrydockercredentials "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/container/registrydockercredentials"
	image "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/custom/image"
	cluster "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/database/cluster"
	connectionpool "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/database/connectionpool"
	db "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/database/db"
	firewall "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/database/firewall"
	kafkatopic "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/database/kafkatopic"
	mysqlconfig "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/database/mysqlconfig"
	redisconfig "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/database/redisconfig"
	replica "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/database/replica"
	user "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/database/user"
	app "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/digitalocean/app"
	tag "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/digitalocean/tag"
	domain "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/dns/domain"
	record "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/dns/record"
	droplet "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/droplet/droplet"
	snapshot "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/droplet/snapshot"
	clusterkubernetes "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/kubernetes/cluster"
	nodepool "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/kubernetes/nodepool"
	alert "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/monitor/alert"
	certificate "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/networking/certificate"
	firewallnetworking "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/networking/firewall"
	ip "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/networking/ip"
	ipassignment "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/networking/ipassignment"
	loadbalancer "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/networking/loadbalancer"
	project "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/project/project"
	resources "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/project/resources"
	providerconfig "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/providerconfig"
	bucket "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/spaces/bucket"
	bucketcorsconfiguration "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/spaces/bucketcorsconfiguration"
	bucketobject "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/spaces/bucketobject"
	bucketpolicy "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/spaces/bucketpolicy"
	cdn "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/spaces/cdn"
	key "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/ssh/key"
	alertuptime "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/uptime/alert"
	check "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/uptime/check"
	attachment "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/volume/attachment"
	snapshotvolume "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/volume/snapshot"
	volume "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/volume/volume"
	vpc "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/vpc/vpc"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		registry.Setup,
		registrydockercredentials.Setup,
		image.Setup,
		cluster.Setup,
		connectionpool.Setup,
		db.Setup,
		firewall.Setup,
		kafkatopic.Setup,
		mysqlconfig.Setup,
		redisconfig.Setup,
		replica.Setup,
		user.Setup,
		app.Setup,
		tag.Setup,
		domain.Setup,
		record.Setup,
		droplet.Setup,
		snapshot.Setup,
		clusterkubernetes.Setup,
		nodepool.Setup,
		alert.Setup,
		certificate.Setup,
		firewallnetworking.Setup,
		ip.Setup,
		ipassignment.Setup,
		loadbalancer.Setup,
		project.Setup,
		resources.Setup,
		providerconfig.Setup,
		bucket.Setup,
		bucketcorsconfiguration.Setup,
		bucketobject.Setup,
		bucketpolicy.Setup,
		cdn.Setup,
		key.Setup,
		alertuptime.Setup,
		check.Setup,
		attachment.Setup,
		snapshotvolume.Setup,
		volume.Setup,
		vpc.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
