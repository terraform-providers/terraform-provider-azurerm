package azurerm

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-05-10/resources"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/resourceproviders"
)

// requiredResourceProviders returns all of the Resource Providers used by the AzureRM Provider
// whilst all may not be used by every user - the intention is that we determine which should be
// registered such that we can avoid obscure errors where Resource Providers aren't registered.
// new Resource Providers should be added to this list as they're used in the Provider
// (this is the approach used by Microsoft in their tooling)
func requiredResourceProviders() map[string]struct{} {
	// NOTE: Resource Providers in this list are case sensitive
	return map[string]struct{}{
		"Microsoft.ApiManagement":       {},
		"Microsoft.Authorization":       {},
		"Microsoft.Automation":          {},
		"Microsoft.Cache":               {},
		"Microsoft.Cdn":                 {},
		"Microsoft.CognitiveServices":   {},
		"Microsoft.Compute":             {},
		"Microsoft.ContainerInstance":   {},
		"Microsoft.ContainerRegistry":   {},
		"Microsoft.ContainerService":    {},
		"Microsoft.Databricks":          {},
		"Microsoft.DataLakeStore":       {},
		"Microsoft.DBforMySQL":          {},
		"Microsoft.DBforPostgreSQL":     {},
		"Microsoft.Devices":             {},
		"Microsoft.DevTestLab":          {},
		"Microsoft.DocumentDB":          {},
		"Microsoft.EventGrid":           {},
		"Microsoft.EventHub":            {},
		"Microsoft.KeyVault":            {},
		"microsoft.insights":            {},
		"Microsoft.Logic":               {},
		"Microsoft.ManagedIdentity":     {},
		"Microsoft.Management":          {},
		"Microsoft.Network":             {},
		"Microsoft.NotificationHubs":    {},
		"Microsoft.OperationalInsights": {},
		"Microsoft.Relay":               {},
		"Microsoft.Resources":           {},
		"Microsoft.Search":              {},
		"Microsoft.ServiceBus":          {},
		"Microsoft.ServiceFabric":       {},
		"Microsoft.Sql":                 {},
		"Microsoft.Storage":             {},
	}
}

func ensureResourceProvidersAreRegistered(ctx context.Context, client resources.ProvidersClient, availableRPs []resources.Provider, requiredRPs map[string]struct{}) error {
	log.Printf("[DEBUG] Determining which Resource Providers require Registration")
	providersToRegister := resourceproviders.DetermineResourceProvidersRequiringRegistration(availableRPs, requiredRPs)

	if len(providersToRegister) > 0 {
		log.Printf("[DEBUG] Registering %d Resource Providers", len(providersToRegister))
		err := resourceproviders.RegisterForSubscription(ctx, client, providersToRegister)
		if err != nil {
			return err
		}
	} else {
		log.Printf("[DEBUG] All required Resource Providers are registered")
	}

	return nil
}
