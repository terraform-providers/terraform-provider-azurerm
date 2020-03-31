package tests

import (
	"context"
	"fmt"
	"regexp"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/features"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/advisor/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

var (
	recommendationId string
	once             sync.Once
)

func TestAccAzureRMAdvisorSuppression_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_advisor_suppression", "test")
	recommendationId = buildAzureRMAdvisorRecommendationData(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMAdvisorSuppressionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMAdvisorSuppression_basic(data, recommendationId),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMAdvisorSuppressionExists(data.ResourceName),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccAzureRMAdvisorSuppression_requiresImport(t *testing.T) {
	if !features.ShouldResourcesBeImported() {
		t.Skip("Skipping since resources aren't required to be imported")
		return
	}

	data := acceptance.BuildTestData(t, "azurerm_advisor_suppression", "test")
	recommendationId = buildAzureRMAdvisorRecommendationData(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMAdvisorSuppressionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMAdvisorSuppression_basic(data, recommendationId),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMAdvisorSuppressionExists(data.ResourceName),
				),
			},
			{
				Config:      testAccAzureRMAdvisorSuppression_requiresImport(data, recommendationId),
				ExpectError: acceptance.RequiresImportError("azurerm_advisor_suppression"),
			},
		},
	})
}

func TestAccAzureRMAdvisorSuppression_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_advisor_suppression", "test")
	recommendationId = buildAzureRMAdvisorRecommendationData(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMAdvisorSuppressionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMAdvisorSuppression_complete(data, recommendationId),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMAdvisorSuppressionExists(data.ResourceName),
					resource.TestMatchResourceAttr(data.ResourceName, "name", regexp.MustCompile("acctest-sp-/*")),
					resource.TestCheckResourceAttr(data.ResourceName, "suppressed_duration", "14.00:00:00"),
				),
			},
			data.ImportStep(),
			{
				Config: testAccAzureRMAdvisorSuppression_update(data, recommendationId),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMAdvisorSuppressionExists(data.ResourceName),
					resource.TestMatchResourceAttr(data.ResourceName, "name", regexp.MustCompile("acctest-sp2-/*")),
					resource.TestCheckResourceAttr(data.ResourceName, "suppressed_duration", "7.10:30:00"),
				),
			},
			data.ImportStep(),
		},
	})
}

func testCheckAzureRMAdvisorSuppressionExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Advisor Suppression not found: %s", resourceName)
		}

		id, err := parse.AdvisorSuppressionID(rs.Primary.ID)
		if err != nil {
			return err
		}

		client := acceptance.AzureProvider.Meta().(*clients.Client).Advisor.SuppressionsClient
		ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

		if resp, err := client.Get(ctx, id.ResourceUri, id.RecommendationName, id.Name); err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Advisor Suppression %q does not exist", id.Name)
			}
			return fmt.Errorf("Bad: Get on Advisor Suppression Client: %+v", err)
		}

		return nil
	}
}

func testCheckAzureRMAdvisorSuppressionDestroy(s *terraform.State) error {
	client := acceptance.AzureProvider.Meta().(*clients.Client).Advisor.SuppressionsClient
	ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azurerm_advisor_suppression" {
			continue
		}

		id, err := parse.AdvisorSuppressionID(rs.Primary.ID)
		if err != nil {
			return err
		}

		if resp, err := client.Get(ctx, id.ResourceUri, id.RecommendationName, id.Name); err != nil {
			if !utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Get on Advisor Suppression Client: %+v", err)
			}
		}

		return nil
	}

	return nil
}

func buildAzureRMAdvisorRecommendationData(t *testing.T) string {
	once.Do(func() {
		config := acceptance.GetAuthConfig(t)
		if config == nil {
			t.SkipNow()
			t.Fatalf("Bad: Failure in building ARM Client")
		}

		builder := clients.ClientBuilder{
			AuthConfig:                  config,
			TerraformVersion:            "0.0.0",
			PartnerId:                   "",
			DisableCorrelationRequestID: true,
			DisableTerraformPartnerID:   false,
			SkipProviderRegistration:    false,
		}
		client, err := clients.Build(context.Background(), builder)
		if err != nil {
			t.Fatal(fmt.Errorf("Bad: Failure in building ARM Client: %+v", err))
		}

		client.StopContext = acceptance.AzureProvider.StopContext()

		rclient := client.Advisor.RecommendationsClient
		ctx := client.StopContext
		recommendationIterator, err := rclient.ListComplete(ctx, "", nil, "")
		if err != nil {
			t.Fatalf("Failure in retrieving Advisor Recommendations: %+v", err)
		}

		if !recommendationIterator.NotDone() {
			t.Fatalf("Bad: Advisor Recommendations are empty")
		}

		recommendationId = *recommendationIterator.Value().ID
		if recommendationId == "" {
			t.Fatalf("Advisor Recommendation ID is empty")
		}
	})
	return recommendationId
}

func testAccAzureRMAdvisorSuppression_basic(data acceptance.TestData, recommendationId string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_advisor_suppression" "test" {
  name              = "acctest-sp-%d"
  recommendation_id = "%s"
}
`, data.RandomInteger, recommendationId)
}

func testAccAzureRMAdvisorSuppression_requiresImport(data acceptance.TestData, recommendationId string) string {
	template := testAccAzureRMAdvisorSuppression_basic(data, recommendationId)
	return fmt.Sprintf(`
%s

resource "azurerm_advisor_suppression" "import" {
  name              = azurerm_advisor_suppression.test.name
  recommendation_id = azurerm_advisor_suppression.test.recommendation_id
}
`, template)
}

func testAccAzureRMAdvisorSuppression_complete(data acceptance.TestData, recommendationId string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_advisor_suppression" "test" {
  name                = "acctest-sp-%d"
  recommendation_id   = "%s"
  suppressed_duration = "14"
}
`, data.RandomInteger, recommendationId)
}

func testAccAzureRMAdvisorSuppression_update(data acceptance.TestData, recommendationId string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_advisor_suppression" "test" {
  name                = "acctest-sp2-%d"
  recommendation_id   = "%s"
  suppressed_duration = "7.10:30"
}
`, data.RandomInteger, recommendationId)
}
