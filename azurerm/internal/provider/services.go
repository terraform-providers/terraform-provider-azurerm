package provider

import (
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/advisor"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/analysisservices"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/apimanagement"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/appconfiguration"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/applicationinsights"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/appplatform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/attestation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/authorization"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/automation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/azurestackhci"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/batch"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/blueprints"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/bot"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/cdn"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/cognitive"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/common"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/compute"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/containers"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/cosmos"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/costmanagement"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/customproviders"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/databasemigration"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/databricks"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/datafactory"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/datalake"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/datashare"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/desktopvirtualization"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/devspace"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/devtestlabs"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/dns"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/eventgrid"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/eventhub"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/frontdoor"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/hdinsight"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/healthcare"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/hpccache"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/hsm"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/iotcentral"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/iothub"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/iottimeseriesinsights"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/keyvault"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/kusto"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/lighthouse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/loganalytics"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/logic"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/machinelearning"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/maintenance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/managedapplications"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/managementgroup"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/maps"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/mariadb"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/media"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/mixedreality"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/monitor"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/msi"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/mssql"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/mysql"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/netapp"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/network"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/notificationhub"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/policy"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/portal"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/postgres"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/powerbi"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/privatedns"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/recoveryservices"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/redis"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/relay"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/search"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/securitycenter"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/sentinel"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/servicebus"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/servicefabric"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/servicefabricmesh"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/signalr"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/sql"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/storage"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/streamanalytics"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/subscription"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/synapse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/trafficmanager"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/web"
)

//go:generate go run ../tools/generator-services/main.go -path=../../../

func SupportedServices() []common.ServiceRegistration {
	return []common.ServiceRegistration{
		advisor.Registration{},
		analysisservices.Registration{},
		apimanagement.Registration{},
		appconfiguration.Registration{},
		appplatform.Registration{},
		applicationinsights.Registration{},
		attestation.Registration{},
		authorization.Registration{},
		automation.Registration{},
		azurestackhci.Registration{},
		batch.Registration{},
		blueprints.Registration{},
		bot.Registration{},
		cdn.Registration{},
		cognitive.Registration{},
		compute.Registration{},
		containers.Registration{},
		cosmos.Registration{},
		costmanagement.Registration{},
		customproviders.Registration{},
		databricks.Registration{},
		datafactory.Registration{},
		datalake.Registration{},
		databasemigration.Registration{},
		datashare.Registration{},
		desktopvirtualization.Registration{},
		devspace.Registration{},
		devtestlabs.Registration{},
		dns.Registration{},
		eventgrid.Registration{},
		eventhub.Registration{},
		frontdoor.Registration{},
		hpccache.Registration{},
		hsm.Registration{},
		hdinsight.Registration{},
		healthcare.Registration{},
		iothub.Registration{},
		iotcentral.Registration{},
		keyvault.Registration{},
		kusto.Registration{},
		loganalytics.Registration{},
		logic.Registration{},
		machinelearning.Registration{},
		maintenance.Registration{},
		managedapplications.Registration{},
		lighthouse.Registration{},
		managementgroup.Registration{},
		maps.Registration{},
		mariadb.Registration{},
		media.Registration{},
		mixedreality.Registration{},
		monitor.Registration{},
		msi.Registration{},
		mssql.Registration{},
		mysql.Registration{},
		netapp.Registration{},
		network.Registration{},
		notificationhub.Registration{},
		policy.Registration{},
		portal.Registration{},
		postgres.Registration{},
		powerbi.Registration{},
		privatedns.Registration{},
		recoveryservices.Registration{},
		redis.Registration{},
		relay.Registration{},
		resource.Registration{},
		search.Registration{},
		securitycenter.Registration{},
		sentinel.Registration{},
		servicebus.Registration{},
		servicefabric.Registration{},
		servicefabricmesh.Registration{},
		signalr.Registration{},
		sql.Registration{},
		storage.Registration{},
		streamanalytics.Registration{},
		subscription.Registration{},
		synapse.Registration{},
		iottimeseriesinsights.Registration{},
		trafficmanager.Registration{},
		web.Registration{},
	}
}
