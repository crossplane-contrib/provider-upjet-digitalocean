/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	registry "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/container/registry"
	registrydockercredentials "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/container/registrydockercredentials"
	image "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/custom/image"
	cluster "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/database/cluster"
	connectionpool "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/database/connectionpool"
	db "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/database/db"
	firewall "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/database/firewall"
	kafkaconfig "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/database/kafkaconfig"
	kafkaschemaregistry "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/database/kafkaschemaregistry"
	kafkatopic "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/database/kafkatopic"
	logsinkopensearch "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/database/logsinkopensearch"
	logsinkrsyslog "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/database/logsinkrsyslog"
	mongodbconfig "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/database/mongodbconfig"
	mysqlconfig "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/database/mysqlconfig"
	onlinemigration "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/database/onlinemigration"
	opensearchconfig "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/database/opensearchconfig"
	postgresqlconfig "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/database/postgresqlconfig"
	redisconfig "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/database/redisconfig"
	replica "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/database/replica"
	user "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/database/user"
	valkeyconfig "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/database/valkeyconfig"
	app "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/digitalocean/app"
	tag "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/digitalocean/tag"
	domain "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/dns/domain"
	record "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/dns/record"
	autoscalepool "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/droplet/autoscalepool"
	droplet "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/droplet/droplet"
	snapshot "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/droplet/snapshot"
	agent "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/gradientai/agent"
	agentknowledgebaseattachment "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/gradientai/agentknowledgebaseattachment"
	agentroute "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/gradientai/agentroute"
	function "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/gradientai/function"
	indexingjobcancel "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/gradientai/indexingjobcancel"
	knowledgebase "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/gradientai/knowledgebase"
	knowledgebasedatasource "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/gradientai/knowledgebasedatasource"
	openaiapikey "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/gradientai/openaiapikey"
	clusterkubernetes "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/kubernetes/cluster"
	nodepool "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/kubernetes/nodepool"
	alert "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/monitor/alert"
	byoipprefix "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/networking/byoipprefix"
	certificate "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/networking/certificate"
	firewallnetworking "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/networking/firewall"
	ip "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/networking/ip"
	ipassignment "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/networking/ipassignment"
	ipv6 "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/networking/ipv6"
	ipv6assignment "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/networking/ipv6assignment"
	loadbalancer "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/networking/loadbalancer"
	attachment "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/nfs/attachment"
	nfsshare "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/nfs/nfsshare"
	snapshotnfs "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/nfs/snapshot"
	project "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/project/project"
	resources "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/project/resources"
	providerconfig "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/providerconfig"
	bucket "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/spaces/bucket"
	bucketcorsconfiguration "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/spaces/bucketcorsconfiguration"
	bucketlogging "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/spaces/bucketlogging"
	bucketobject "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/spaces/bucketobject"
	bucketpolicy "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/spaces/bucketpolicy"
	cdn "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/spaces/cdn"
	key "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/spaces/key"
	keyssh "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/ssh/key"
	alertuptime "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/uptime/alert"
	check "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/uptime/check"
	attachmentvolume "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/volume/attachment"
	snapshotvolume "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/volume/snapshot"
	volume "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/volume/volume"
	natgateway "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/vpc/natgateway"
	peering "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/vpc/peering"
	vpc "github.com/crossplane-contrib/provider-upjet-digitalocean/internal/controller/cluster/vpc/vpc"
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
