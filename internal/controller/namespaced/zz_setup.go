/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	registry "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/container/registry"
	registrydockercredentials "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/container/registrydockercredentials"
	image "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/custom/image"
	cluster "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/database/cluster"
	connectionpool "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/database/connectionpool"
	db "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/database/db"
	firewall "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/database/firewall"
	kafkaconfig "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/database/kafkaconfig"
	kafkaschemaregistry "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/database/kafkaschemaregistry"
	kafkatopic "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/database/kafkatopic"
	logsinkopensearch "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/database/logsinkopensearch"
	logsinkrsyslog "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/database/logsinkrsyslog"
	mongodbconfig "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/database/mongodbconfig"
	mysqlconfig "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/database/mysqlconfig"
	onlinemigration "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/database/onlinemigration"
	opensearchconfig "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/database/opensearchconfig"
	postgresqlconfig "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/database/postgresqlconfig"
	redisconfig "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/database/redisconfig"
	replica "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/database/replica"
	user "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/database/user"
	valkeyconfig "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/database/valkeyconfig"
	app "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/digitalocean/app"
	tag "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/digitalocean/tag"
	domain "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/dns/domain"
	record "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/dns/record"
	autoscalepool "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/droplet/autoscalepool"
	droplet "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/droplet/droplet"
	snapshot "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/droplet/snapshot"
	agent "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/gradientai/agent"
	agentknowledgebaseattachment "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/gradientai/agentknowledgebaseattachment"
	agentroute "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/gradientai/agentroute"
	function "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/gradientai/function"
	indexingjobcancel "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/gradientai/indexingjobcancel"
	knowledgebase "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/gradientai/knowledgebase"
	knowledgebasedatasource "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/gradientai/knowledgebasedatasource"
	openaiapikey "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/gradientai/openaiapikey"
	clusterkubernetes "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/kubernetes/cluster"
	nodepool "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/kubernetes/nodepool"
	alert "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/monitor/alert"
	byoipprefix "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/networking/byoipprefix"
	certificate "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/networking/certificate"
	firewallnetworking "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/networking/firewall"
	ip "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/networking/ip"
	ipassignment "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/networking/ipassignment"
	ipv6 "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/networking/ipv6"
	ipv6assignment "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/networking/ipv6assignment"
	loadbalancer "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/networking/loadbalancer"
	attachment "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/nfs/attachment"
	nfsshare "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/nfs/nfsshare"
	snapshotnfs "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/nfs/snapshot"
	project "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/project/project"
	resources "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/project/resources"
	providerconfig "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/providerconfig"
	bucket "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/spaces/bucket"
	bucketcorsconfiguration "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/spaces/bucketcorsconfiguration"
	bucketlogging "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/spaces/bucketlogging"
	bucketobject "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/spaces/bucketobject"
	bucketpolicy "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/spaces/bucketpolicy"
	cdn "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/spaces/cdn"
	key "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/spaces/key"
	keyssh "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/ssh/key"
	alertuptime "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/uptime/alert"
	check "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/uptime/check"
	attachmentvolume "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/volume/attachment"
	snapshotvolume "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/volume/snapshot"
	volume "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/volume/volume"
	natgateway "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/vpc/natgateway"
	peering "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/vpc/peering"
	vpc "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/namespaced/vpc/vpc"
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
		kafkaconfig.Setup,
		kafkaschemaregistry.Setup,
		kafkatopic.Setup,
		logsinkopensearch.Setup,
		logsinkrsyslog.Setup,
		mongodbconfig.Setup,
		mysqlconfig.Setup,
		onlinemigration.Setup,
		opensearchconfig.Setup,
		postgresqlconfig.Setup,
		redisconfig.Setup,
		replica.Setup,
		user.Setup,
		valkeyconfig.Setup,
		app.Setup,
		tag.Setup,
		domain.Setup,
		record.Setup,
		autoscalepool.Setup,
		droplet.Setup,
		snapshot.Setup,
		agent.Setup,
		agentknowledgebaseattachment.Setup,
		agentroute.Setup,
		function.Setup,
		indexingjobcancel.Setup,
		knowledgebase.Setup,
		knowledgebasedatasource.Setup,
		openaiapikey.Setup,
		clusterkubernetes.Setup,
		nodepool.Setup,
		alert.Setup,
		byoipprefix.Setup,
		certificate.Setup,
		firewallnetworking.Setup,
		ip.Setup,
		ipassignment.Setup,
		ipv6.Setup,
		ipv6assignment.Setup,
		loadbalancer.Setup,
		attachment.Setup,
		nfsshare.Setup,
		snapshotnfs.Setup,
		project.Setup,
		resources.Setup,
		providerconfig.Setup,
		bucket.Setup,
		bucketcorsconfiguration.Setup,
		bucketlogging.Setup,
		bucketobject.Setup,
		bucketpolicy.Setup,
		cdn.Setup,
		key.Setup,
		keyssh.Setup,
		alertuptime.Setup,
		check.Setup,
		attachmentvolume.Setup,
		snapshotvolume.Setup,
		volume.Setup,
		natgateway.Setup,
		peering.Setup,
		vpc.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		registry.SetupGated,
		registrydockercredentials.SetupGated,
		image.SetupGated,
		cluster.SetupGated,
		connectionpool.SetupGated,
		db.SetupGated,
		firewall.SetupGated,
		kafkaconfig.SetupGated,
		kafkaschemaregistry.SetupGated,
		kafkatopic.SetupGated,
		logsinkopensearch.SetupGated,
		logsinkrsyslog.SetupGated,
		mongodbconfig.SetupGated,
		mysqlconfig.SetupGated,
		onlinemigration.SetupGated,
		opensearchconfig.SetupGated,
		postgresqlconfig.SetupGated,
		redisconfig.SetupGated,
		replica.SetupGated,
		user.SetupGated,
		valkeyconfig.SetupGated,
		app.SetupGated,
		tag.SetupGated,
		domain.SetupGated,
		record.SetupGated,
		autoscalepool.SetupGated,
		droplet.SetupGated,
		snapshot.SetupGated,
		agent.SetupGated,
		agentknowledgebaseattachment.SetupGated,
		agentroute.SetupGated,
		function.SetupGated,
		indexingjobcancel.SetupGated,
		knowledgebase.SetupGated,
		knowledgebasedatasource.SetupGated,
		openaiapikey.SetupGated,
		clusterkubernetes.SetupGated,
		nodepool.SetupGated,
		alert.SetupGated,
		byoipprefix.SetupGated,
		certificate.SetupGated,
		firewallnetworking.SetupGated,
		ip.SetupGated,
		ipassignment.SetupGated,
		ipv6.SetupGated,
		ipv6assignment.SetupGated,
		loadbalancer.SetupGated,
		attachment.SetupGated,
		nfsshare.SetupGated,
		snapshotnfs.SetupGated,
		project.SetupGated,
		resources.SetupGated,
		providerconfig.SetupGated,
		bucket.SetupGated,
		bucketcorsconfiguration.SetupGated,
		bucketlogging.SetupGated,
		bucketobject.SetupGated,
		bucketpolicy.SetupGated,
		cdn.SetupGated,
		key.SetupGated,
		keyssh.SetupGated,
		alertuptime.SetupGated,
		check.SetupGated,
		attachmentvolume.SetupGated,
		snapshotvolume.SetupGated,
		volume.SetupGated,
		natgateway.SetupGated,
		peering.SetupGated,
		vpc.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
