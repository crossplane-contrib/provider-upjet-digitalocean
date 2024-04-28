// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	droplet "github.com/straw-hat-team/provider-digitalocean/internal/controller/compute/droplet"
	registry "github.com/straw-hat-team/provider-digitalocean/internal/controller/container/registry"
	registrydockercredentials "github.com/straw-hat-team/provider-digitalocean/internal/controller/container/registrydockercredentials"
	image "github.com/straw-hat-team/provider-digitalocean/internal/controller/custom/image"
	cluster "github.com/straw-hat-team/provider-digitalocean/internal/controller/database/cluster"
	connectionpool "github.com/straw-hat-team/provider-digitalocean/internal/controller/database/connectionpool"
	db "github.com/straw-hat-team/provider-digitalocean/internal/controller/database/db"
	replica "github.com/straw-hat-team/provider-digitalocean/internal/controller/database/replica"
	user "github.com/straw-hat-team/provider-digitalocean/internal/controller/database/user"
	app "github.com/straw-hat-team/provider-digitalocean/internal/controller/digitalocean/app"
	cdn "github.com/straw-hat-team/provider-digitalocean/internal/controller/digitalocean/cdn"
	certificate "github.com/straw-hat-team/provider-digitalocean/internal/controller/digitalocean/certificate"
	tag "github.com/straw-hat-team/provider-digitalocean/internal/controller/digitalocean/tag"
	snapshot "github.com/straw-hat-team/provider-digitalocean/internal/controller/droplet/snapshot"
	clusterkubernetes "github.com/straw-hat-team/provider-digitalocean/internal/controller/kubernetes/cluster"
	nodepool "github.com/straw-hat-team/provider-digitalocean/internal/controller/kubernetes/nodepool"
	alert "github.com/straw-hat-team/provider-digitalocean/internal/controller/monitor/alert"
	domain "github.com/straw-hat-team/provider-digitalocean/internal/controller/networking/domain"
	firewall "github.com/straw-hat-team/provider-digitalocean/internal/controller/networking/firewall"
	ip "github.com/straw-hat-team/provider-digitalocean/internal/controller/networking/ip"
	ipassignment "github.com/straw-hat-team/provider-digitalocean/internal/controller/networking/ipassignment"
	loadbalancer "github.com/straw-hat-team/provider-digitalocean/internal/controller/networking/loadbalancer"
	record "github.com/straw-hat-team/provider-digitalocean/internal/controller/networking/record"
	vpc "github.com/straw-hat-team/provider-digitalocean/internal/controller/networking/vpc"
	project "github.com/straw-hat-team/provider-digitalocean/internal/controller/project/project"
	providerconfig "github.com/straw-hat-team/provider-digitalocean/internal/controller/providerconfig"
	bucketcorsconfiguration "github.com/straw-hat-team/provider-digitalocean/internal/controller/spaces/bucketcorsconfiguration"
	bucketobject "github.com/straw-hat-team/provider-digitalocean/internal/controller/spaces/bucketobject"
	bucketpolicy "github.com/straw-hat-team/provider-digitalocean/internal/controller/spaces/bucketpolicy"
	key "github.com/straw-hat-team/provider-digitalocean/internal/controller/ssh/key"
	bucket "github.com/straw-hat-team/provider-digitalocean/internal/controller/storage/bucket"
	alertuptime "github.com/straw-hat-team/provider-digitalocean/internal/controller/uptime/alert"
	check "github.com/straw-hat-team/provider-digitalocean/internal/controller/uptime/check"
	attachment "github.com/straw-hat-team/provider-digitalocean/internal/controller/volume/attachment"
	snapshotvolume "github.com/straw-hat-team/provider-digitalocean/internal/controller/volume/snapshot"
	volume "github.com/straw-hat-team/provider-digitalocean/internal/controller/volume/volume"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		droplet.Setup,
		registry.Setup,
		registrydockercredentials.Setup,
		image.Setup,
		cluster.Setup,
		connectionpool.Setup,
		db.Setup,
		replica.Setup,
		user.Setup,
		app.Setup,
		cdn.Setup,
		certificate.Setup,
		tag.Setup,
		snapshot.Setup,
		clusterkubernetes.Setup,
		nodepool.Setup,
		alert.Setup,
		domain.Setup,
		firewall.Setup,
		ip.Setup,
		ipassignment.Setup,
		loadbalancer.Setup,
		record.Setup,
		vpc.Setup,
		project.Setup,
		providerconfig.Setup,
		bucketcorsconfiguration.Setup,
		bucketobject.Setup,
		bucketpolicy.Setup,
		key.Setup,
		bucket.Setup,
		alertuptime.Setup,
		check.Setup,
		attachment.Setup,
		snapshotvolume.Setup,
		volume.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
