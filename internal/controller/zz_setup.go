// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	droplet "github.com/straw-hat-team/provider-digitalocean/internal/controller/compute/droplet"
	cluster "github.com/straw-hat-team/provider-digitalocean/internal/controller/database/cluster"
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
