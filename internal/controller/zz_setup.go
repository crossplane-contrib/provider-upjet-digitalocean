// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	droplet "github.com/straw-hat-team/provider-digitalocean/internal/controller/compute/droplet"
	cluster "github.com/straw-hat-team/provider-digitalocean/internal/controller/database/cluster"
	replica "github.com/straw-hat-team/provider-digitalocean/internal/controller/database/replica"
	user "github.com/straw-hat-team/provider-digitalocean/internal/controller/database/user"
	tag "github.com/straw-hat-team/provider-digitalocean/internal/controller/digitalocean/tag"
	snapshot "github.com/straw-hat-team/provider-digitalocean/internal/controller/droplet/snapshot"
	clusterkubernetes "github.com/straw-hat-team/provider-digitalocean/internal/controller/kubernetes/cluster"
	nodepool "github.com/straw-hat-team/provider-digitalocean/internal/controller/kubernetes/nodepool"
	domain "github.com/straw-hat-team/provider-digitalocean/internal/controller/networking/domain"
	loadbalancer "github.com/straw-hat-team/provider-digitalocean/internal/controller/networking/loadbalancer"
	record "github.com/straw-hat-team/provider-digitalocean/internal/controller/networking/record"
	vpc "github.com/straw-hat-team/provider-digitalocean/internal/controller/networking/vpc"
	project "github.com/straw-hat-team/provider-digitalocean/internal/controller/project/project"
	providerconfig "github.com/straw-hat-team/provider-digitalocean/internal/controller/providerconfig"
	bucket "github.com/straw-hat-team/provider-digitalocean/internal/controller/storage/bucket"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		droplet.Setup,
		cluster.Setup,
		replica.Setup,
		user.Setup,
		tag.Setup,
		snapshot.Setup,
		clusterkubernetes.Setup,
		nodepool.Setup,
		domain.Setup,
		loadbalancer.Setup,
		record.Setup,
		vpc.Setup,
		project.Setup,
		providerconfig.Setup,
		bucket.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
