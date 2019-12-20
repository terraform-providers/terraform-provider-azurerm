package azurerm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/provider"
)

func Provider() terraform.ResourceProvider {
	dataSources := map[string]*schema.Resource{
		"azurerm_app_service_plan":                          dataSourceAppServicePlan(),
		"azurerm_app_service_certificate":                   dataSourceAppServiceCertificate(),
		"azurerm_app_service":                               dataSourceArmAppService(),
		"azurerm_app_service_certificate_order":             dataSourceArmAppServiceCertificateOrder(),
		"azurerm_application_security_group":                dataSourceArmApplicationSecurityGroup(),
		"azurerm_builtin_role_definition":                   dataSourceArmBuiltInRoleDefinition(),
		"azurerm_client_config":                             dataSourceArmClientConfig(),
		"azurerm_express_route_circuit":                     dataSourceArmExpressRouteCircuit(),
		"azurerm_image":                                     dataSourceArmImage(),
		"azurerm_key_vault_access_policy":                   dataSourceArmKeyVaultAccessPolicy(),
		"azurerm_key_vault_key":                             dataSourceArmKeyVaultKey(),
		"azurerm_key_vault_secret":                          dataSourceArmKeyVaultSecret(),
		"azurerm_key_vault":                                 dataSourceArmKeyVault(),
		"azurerm_lb":                                        dataSourceArmLoadBalancer(),
		"azurerm_lb_backend_address_pool":                   dataSourceArmLoadBalancerBackendAddressPool(),
		"azurerm_mssql_elasticpool":                         dataSourceArmMsSqlElasticpool(),
		"azurerm_platform_image":                            dataSourceArmPlatformImage(),
		"azurerm_private_link_endpoint_connection":          dataSourceArmPrivateLinkEndpointConnection(),
		"azurerm_private_endpoint_connection":               dataSourceArmPrivateEndpointConnection(),
		"azurerm_private_link_service":                      dataSourceArmPrivateLinkService(),
		"azurerm_private_link_service_endpoint_connections": dataSourceArmPrivateLinkServiceEndpointConnections(),
		"azurerm_proximity_placement_group":                 dataSourceArmProximityPlacementGroup(),
		"azurerm_public_ip":                                 dataSourceArmPublicIP(),
		"azurerm_public_ips":                                dataSourceArmPublicIPs(),
		"azurerm_public_ip_prefix":                          dataSourceArmPublicIpPrefix(),
		"azurerm_resources":                                 dataSourceArmResources(),
		"azurerm_resource_group":                            dataSourceArmResourceGroup(),
		"azurerm_stream_analytics_job":                      dataSourceArmStreamAnalyticsJob(),
		"azurerm_storage_account_blob_container_sas":        dataSourceArmStorageAccountBlobContainerSharedAccessSignature(),
		"azurerm_storage_account_sas":                       dataSourceArmStorageAccountSharedAccessSignature(),
		"azurerm_storage_account":                           dataSourceArmStorageAccount(),
		"azurerm_storage_management_policy":                 dataSourceArmStorageManagementPolicy(),
		"azurerm_subscription":                              dataSourceArmSubscription(),
		"azurerm_subscriptions":                             dataSourceArmSubscriptions(),
		"azurerm_traffic_manager_geographical_location":     dataSourceArmTrafficManagerGeographicalLocation(),
	}

	resources := map[string]*schema.Resource{
		"azurerm_app_service_active_slot":                               resourceArmAppServiceActiveSlot(),
		"azurerm_app_service_certificate":                               resourceArmAppServiceCertificate(),
		"azurerm_app_service_certificate_order":                         resourceArmAppServiceCertificateOrder(),
		"azurerm_app_service_custom_hostname_binding":                   resourceArmAppServiceCustomHostnameBinding(),
		"azurerm_app_service_plan":                                      resourceArmAppServicePlan(),
		"azurerm_app_service_slot":                                      resourceArmAppServiceSlot(),
		"azurerm_app_service_source_control_token":                      resourceArmAppServiceSourceControlToken(),
		"azurerm_app_service_virtual_network_swift_connection":          resourceArmAppServiceVirtualNetworkSwiftConnection(),
		"azurerm_app_service":                                           resourceArmAppService(),
		"azurerm_application_gateway":                                   resourceArmApplicationGateway(),
		"azurerm_application_security_group":                            resourceArmApplicationSecurityGroup(),
		"azurerm_autoscale_setting":                                     resourceArmAutoScaleSetting(),
		"azurerm_bastion_host":                                          resourceArmBastionHost(),
		"azurerm_connection_monitor":                                    resourceArmConnectionMonitor(),
		"azurerm_dashboard":                                             resourceArmDashboard(),
		"azurerm_express_route_circuit_authorization":                   resourceArmExpressRouteCircuitAuthorization(),
		"azurerm_express_route_circuit_peering":                         resourceArmExpressRouteCircuitPeering(),
		"azurerm_express_route_circuit":                                 resourceArmExpressRouteCircuit(),
		"azurerm_function_app":                                          resourceArmFunctionApp(),
		"azurerm_image":                                                 resourceArmImage(),
		"azurerm_key_vault_access_policy":                               resourceArmKeyVaultAccessPolicy(),
		"azurerm_key_vault_certificate":                                 resourceArmKeyVaultCertificate(),
		"azurerm_key_vault_key":                                         resourceArmKeyVaultKey(),
		"azurerm_key_vault_secret":                                      resourceArmKeyVaultSecret(),
		"azurerm_key_vault":                                             resourceArmKeyVault(),
		"azurerm_lb_backend_address_pool":                               resourceArmLoadBalancerBackendAddressPool(),
		"azurerm_lb_nat_pool":                                           resourceArmLoadBalancerNatPool(),
		"azurerm_lb_nat_rule":                                           resourceArmLoadBalancerNatRule(),
		"azurerm_lb_probe":                                              resourceArmLoadBalancerProbe(),
		"azurerm_lb_outbound_rule":                                      resourceArmLoadBalancerOutboundRule(),
		"azurerm_lb_rule":                                               resourceArmLoadBalancerRule(),
		"azurerm_lb":                                                    resourceArmLoadBalancer(),
		"azurerm_management_lock":                                       resourceArmManagementLock(),
		"azurerm_marketplace_agreement":                                 resourceArmMarketplaceAgreement(),
		"azurerm_mssql_elasticpool":                                     resourceArmMsSqlElasticPool(),
		"azurerm_mssql_database_vulnerability_assessment_rule_baseline": resourceArmMssqlDatabaseVulnerabilityAssessmentRuleBaseline(),
		"azurerm_mssql_server_security_alert_policy":                    resourceArmMssqlServerSecurityAlertPolicy(),
		"azurerm_mssql_server_vulnerability_assessment":                 resourceArmMssqlServerVulnerabilityAssessment(),
		"azurerm_packet_capture":                                        resourceArmPacketCapture(),
		"azurerm_point_to_site_vpn_gateway":                             resourceArmPointToSiteVPNGateway(),
		"azurerm_private_link_endpoint":                                 resourceArmPrivateLinkEndpoint(),
		"azurerm_private_endpoint":                                      resourceArmPrivateEndpoint(),
		"azurerm_private_link_service":                                  resourceArmPrivateLinkService(),
		"azurerm_proximity_placement_group":                             resourceArmProximityPlacementGroup(),
		"azurerm_public_ip":                                             resourceArmPublicIp(),
		"azurerm_public_ip_prefix":                                      resourceArmPublicIpPrefix(),
		"azurerm_resource_group":                                        resourceArmResourceGroup(),
		"azurerm_shared_image_gallery":                                  resourceArmSharedImageGallery(),
		"azurerm_shared_image_version":                                  resourceArmSharedImageVersion(),
		"azurerm_shared_image":                                          resourceArmSharedImage(),
		"azurerm_storage_account":                                       resourceArmStorageAccount(),
		"azurerm_storage_account_network_rules":                         resourceArmStorageAccountNetworkRules(),
		"azurerm_storage_blob":                                          resourceArmStorageBlob(),
		"azurerm_storage_container":                                     resourceArmStorageContainer(),
		"azurerm_storage_data_lake_gen2_filesystem":                     resourceArmStorageDataLakeGen2FileSystem(),
		"azurerm_storage_management_policy":                             resourceArmStorageManagementPolicy(),
		"azurerm_storage_queue":                                         resourceArmStorageQueue(),
		"azurerm_storage_share":                                         resourceArmStorageShare(),
		"azurerm_storage_share_directory":                               resourceArmStorageShareDirectory(),
		"azurerm_storage_table":                                         resourceArmStorageTable(),
		"azurerm_storage_table_entity":                                  resourceArmStorageTableEntity(),
		"azurerm_stream_analytics_job":                                  resourceArmStreamAnalyticsJob(),
		"azurerm_stream_analytics_function_javascript_udf":              resourceArmStreamAnalyticsFunctionUDF(),
		"azurerm_stream_analytics_output_blob":                          resourceArmStreamAnalyticsOutputBlob(),
		"azurerm_stream_analytics_output_mssql":                         resourceArmStreamAnalyticsOutputSql(),
		"azurerm_stream_analytics_output_eventhub":                      resourceArmStreamAnalyticsOutputEventHub(),
		"azurerm_stream_analytics_output_servicebus_queue":              resourceArmStreamAnalyticsOutputServiceBusQueue(),
		"azurerm_stream_analytics_output_servicebus_topic":              resourceArmStreamAnalyticsOutputServiceBusTopic(),
		"azurerm_stream_analytics_reference_input_blob":                 resourceArmStreamAnalyticsReferenceInputBlob(),
		"azurerm_stream_analytics_stream_input_blob":                    resourceArmStreamAnalyticsStreamInputBlob(),
		"azurerm_stream_analytics_stream_input_eventhub":                resourceArmStreamAnalyticsStreamInputEventHub(),
		"azurerm_stream_analytics_stream_input_iothub":                  resourceArmStreamAnalyticsStreamInputIoTHub(),
		"azurerm_template_deployment":                                   resourceArmTemplateDeployment(),
		"azurerm_traffic_manager_endpoint":                              resourceArmTrafficManagerEndpoint(),
		"azurerm_traffic_manager_profile":                               resourceArmTrafficManagerProfile(),
	}

	return provider.AzureProvider(dataSources, resources)
}
