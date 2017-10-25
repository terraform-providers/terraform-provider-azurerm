package azurerm

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceAzureRMScheduledTime_Daily_Today(t *testing.T) {
	dataSourceName := "data.azurerm_scheduled_time.test"

	now := time.Now().UTC()
	scheduletime := now.Add(time.Duration(1) * time.Hour)

	config := testAccDataSourceAzureRMScheduledTime_Daily(scheduletime)

	expectedTime := time.Date(scheduletime.Year(), scheduletime.Month(), scheduletime.Day(), scheduletime.Hour(), 0, 0, 0, time.UTC)
	formattedExpectedTime := expectedTime.Format(time.RFC3339)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "hour", strconv.Itoa(now.Hour()+1)),
					resource.TestCheckResourceAttr(dataSourceName, "minute", "0"),
					resource.TestCheckResourceAttr(dataSourceName, "second", "0"),
					resource.TestCheckResourceAttr(dataSourceName, "frequency", "Day"),
					resource.TestCheckResourceAttr(dataSourceName, "next_run_time", formattedExpectedTime),
				),
			},
		},
	})
}

func TestAccDataSourceAzureRMScheduledTime_Daily_Tomorrow(t *testing.T) {
	dataSourceName := "data.azurerm_scheduled_time.test"

	now := time.Now().UTC()
	scheduletime := now.Add(time.Duration(-1) * time.Hour)
	config := testAccDataSourceAzureRMScheduledTime_Daily(scheduletime)

	expectedTime := time.Date(scheduletime.Year(), scheduletime.Month(), scheduletime.Day()+1, scheduletime.Hour(), 0, 0, 0, time.UTC)
	formattedExpectedTime := expectedTime.Format(time.RFC3339)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "hour", strconv.Itoa(now.Hour()-1)),
					resource.TestCheckResourceAttr(dataSourceName, "minute", "0"),
					resource.TestCheckResourceAttr(dataSourceName, "second", "0"),
					resource.TestCheckResourceAttr(dataSourceName, "frequency", "Day"),
					resource.TestCheckResourceAttr(dataSourceName, "next_run_time", formattedExpectedTime),
				),
			},
		},
	})
}

func TestAccDataSourceAzureRMScheduledTime_Hourly_CurrentHour(t *testing.T) {
	dataSourceName := "data.azurerm_scheduled_time.test"

	now := time.Now().UTC()
	scheduletime := now.Add(time.Duration(1) * time.Minute)

	config := testAccDataSourceAzureRMScheduledTime_Hourly(scheduletime)

	expectedTime := time.Date(scheduletime.Year(), scheduletime.Month(), scheduletime.Day(), scheduletime.Hour(), scheduletime.Minute(), 0, 0, time.UTC)
	formattedExpectedTime := expectedTime.Format(time.RFC3339)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "hour", strconv.Itoa(scheduletime.Hour())),
					resource.TestCheckResourceAttr(dataSourceName, "minute", strconv.Itoa(scheduletime.Minute())),
					resource.TestCheckResourceAttr(dataSourceName, "second", "0"),
					resource.TestCheckResourceAttr(dataSourceName, "frequency", "Hour"),
					resource.TestCheckResourceAttr(dataSourceName, "next_run_time", formattedExpectedTime),
				),
			},
		},
	})
}

func TestAccDataSourceAzureRMScheduledTime_Hourly_CurrentHour_WithSeconds(t *testing.T) {
	dataSourceName := "data.azurerm_scheduled_time.test"

	now := time.Now().UTC()
	scheduletime := now.Add(time.Duration(1) * time.Minute)

	config := testAccDataSourceAzureRMScheduledTime_Hourly_With_Seconds(scheduletime)

	expectedTime := time.Date(scheduletime.Year(), scheduletime.Month(), scheduletime.Day(), scheduletime.Hour(), scheduletime.Minute(), 24, 0, time.UTC)
	formattedExpectedTime := expectedTime.Format(time.RFC3339)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "hour", strconv.Itoa(scheduletime.Hour())),
					resource.TestCheckResourceAttr(dataSourceName, "minute", strconv.Itoa(scheduletime.Minute())),
					resource.TestCheckResourceAttr(dataSourceName, "second", "24"),
					resource.TestCheckResourceAttr(dataSourceName, "frequency", "Hour"),
					resource.TestCheckResourceAttr(dataSourceName, "next_run_time", formattedExpectedTime),
				),
			},
		},
	})
}

func TestAccDataSourceAzureRMScheduledTime_Hourly_NextHour(t *testing.T) {
	dataSourceName := "data.azurerm_scheduled_time.test"

	now := time.Now().UTC()
	scheduletime := now.Add(time.Duration(-1) * time.Minute)

	config := testAccDataSourceAzureRMScheduledTime_Hourly(scheduletime)

	expectedTime := time.Date(scheduletime.Year(), scheduletime.Month(), scheduletime.Day(), scheduletime.Hour()+1, scheduletime.Minute(), 0, 0, time.UTC)
	formattedExpectedTime := expectedTime.Format(time.RFC3339)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "hour", strconv.Itoa(scheduletime.Hour())),
					resource.TestCheckResourceAttr(dataSourceName, "minute", strconv.Itoa(scheduletime.Minute())),
					resource.TestCheckResourceAttr(dataSourceName, "second", "0"),
					resource.TestCheckResourceAttr(dataSourceName, "frequency", "Hour"),
					resource.TestCheckResourceAttr(dataSourceName, "next_run_time", formattedExpectedTime),
				),
			},
		},
	})
}

func TestAccDataSourceAzureRMScheduledTime_Weekly_Tomorrow(t *testing.T) {
        dataSourceName := "data.azurerm_scheduled_time.test"

        now := time.Now().UTC()
        scheduletime := now.Add(time.Duration(24) * time.Hour) //1 day
        dayofweek := int(scheduletime.Weekday())

        config := testAccDataSourceAzureRMScheduledTime_Weekly(dayofweek, scheduletime)

        expectedTime := time.Date(scheduletime.Year(), scheduletime.Month(), scheduletime.Day(), scheduletime.Hour(), 0, 0, 0, time.UTC)
        formattedExpectedTime := expectedTime.Format(time.RFC3339)

        resource.Test(t, resource.TestCase{
                PreCheck:  func() { testAccPreCheck(t) },
                Providers: testAccProviders,
                Steps: []resource.TestStep{
                        {
                                Config: config,
                                Check: resource.ComposeTestCheckFunc(
                                        resource.TestCheckResourceAttr(dataSourceName, "day_of_week", strconv.Itoa(dayofweek)),
                                        resource.TestCheckResourceAttr(dataSourceName, "hour", strconv.Itoa(now.Hour())),
                                        resource.TestCheckResourceAttr(dataSourceName, "minute", "0"),
                                        resource.TestCheckResourceAttr(dataSourceName, "second", "0"),
                                        resource.TestCheckResourceAttr(dataSourceName, "frequency", "Week"),
                                        resource.TestCheckResourceAttr(dataSourceName, "next_run_time", formattedExpectedTime),
                                ),
                        },
                },
        })
}

func TestAccDataSourceAzureRMScheduledTime_Weekly_Yesterday(t *testing.T) {
        dataSourceName := "data.azurerm_scheduled_time.test"

        now := time.Now().UTC()
        scheduletime := now.Add(time.Duration(144) * time.Hour) //6 days
        dayofweek := int(scheduletime.Weekday())

        config := testAccDataSourceAzureRMScheduledTime_Weekly(dayofweek, scheduletime)

        expectedTime := time.Date(scheduletime.Year(), scheduletime.Month(), scheduletime.Day(), scheduletime.Hour(), 0, 0, 0, time.UTC)
        formattedExpectedTime := expectedTime.Format(time.RFC3339)

        resource.Test(t, resource.TestCase{
                PreCheck:  func() { testAccPreCheck(t) },
                Providers: testAccProviders,
                Steps: []resource.TestStep{
                        {
                                Config: config,
                                Check: resource.ComposeTestCheckFunc(
                                        resource.TestCheckResourceAttr(dataSourceName, "day_of_week", strconv.Itoa(dayofweek)),
                                        resource.TestCheckResourceAttr(dataSourceName, "hour", strconv.Itoa(now.Hour())),
                                        resource.TestCheckResourceAttr(dataSourceName, "minute", "0"),
                                        resource.TestCheckResourceAttr(dataSourceName, "second", "0"),
                                        resource.TestCheckResourceAttr(dataSourceName, "frequency", "Week"),
                                        resource.TestCheckResourceAttr(dataSourceName, "next_run_time", formattedExpectedTime),
                                ),
                        },
                },
        })
}

func TestAccDataSourceAzureRMScheduledTime_Weekly_Next_Week(t *testing.T) {
	dataSourceName := "data.azurerm_scheduled_time.test"

	now := time.Now().UTC()
	scheduletime := now.Add(time.Duration(168) * time.Hour) //1 week
	dayofweek := int(scheduletime.Weekday())

	config := testAccDataSourceAzureRMScheduledTime_Weekly(dayofweek, scheduletime)

	expectedTime := time.Date(scheduletime.Year(), scheduletime.Month(), scheduletime.Day(), scheduletime.Hour(), 0, 0, 0, time.UTC)
	formattedExpectedTime := expectedTime.Format(time.RFC3339)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "day_of_week", strconv.Itoa(dayofweek)),
					resource.TestCheckResourceAttr(dataSourceName, "hour", strconv.Itoa(now.Hour())),
					resource.TestCheckResourceAttr(dataSourceName, "minute", "0"),
					resource.TestCheckResourceAttr(dataSourceName, "second", "0"),
					resource.TestCheckResourceAttr(dataSourceName, "frequency", "Week"),
					resource.TestCheckResourceAttr(dataSourceName, "next_run_time", formattedExpectedTime),
				),
			},
		},
	})
}


func TestAccDataSourceAzureRMScheduledTime_Monthly_Tomorrow(t *testing.T) {
        dataSourceName := "data.azurerm_scheduled_time.test"

        now := time.Now().UTC()
        scheduletime := now.Add(time.Duration(24) * time.Hour) //1 day
        dayofmonth := scheduletime.Day()

        config := testAccDataSourceAzureRMScheduledTime_Monthly(dayofmonth, scheduletime)

        expectedTime := time.Date(scheduletime.Year(), scheduletime.Month(), scheduletime.Day(), scheduletime.Hour(), 0, 0, 0, time.UTC)
        formattedExpectedTime := expectedTime.Format(time.RFC3339)

        resource.Test(t, resource.TestCase{
                PreCheck:  func() { testAccPreCheck(t) },
                Providers: testAccProviders,
                Steps: []resource.TestStep{
                        {
                                Config: config,
                                Check: resource.ComposeTestCheckFunc(
                                        resource.TestCheckResourceAttr(dataSourceName, "day_of_month", strconv.Itoa(dayofmonth)),
                                        resource.TestCheckResourceAttr(dataSourceName, "hour", strconv.Itoa(now.Hour())),
                                        resource.TestCheckResourceAttr(dataSourceName, "minute", "0"),
                                        resource.TestCheckResourceAttr(dataSourceName, "second", "0"),
                                        resource.TestCheckResourceAttr(dataSourceName, "frequency", "Month"),
                                        resource.TestCheckResourceAttr(dataSourceName, "next_run_time", formattedExpectedTime),
                                ),
                        },
                },
        })
}

func TestAccDataSourceAzureRMScheduledTime_Monthly_Next_Month(t *testing.T) {
        dataSourceName := "data.azurerm_scheduled_time.test"

        now := time.Now().UTC()
        scheduletime := now.Add(25 * time.Duration(24) * time.Hour) //25 days
        dayofmonth := scheduletime.Day()

        config := testAccDataSourceAzureRMScheduledTime_Monthly(dayofmonth, scheduletime)

        expectedTime := time.Date(scheduletime.Year(), scheduletime.Month(), scheduletime.Day(), scheduletime.Hour(), 0, 0, 0, time.UTC)
        formattedExpectedTime := expectedTime.Format(time.RFC3339)

        resource.Test(t, resource.TestCase{
                PreCheck:  func() { testAccPreCheck(t) },
                Providers: testAccProviders,
                Steps: []resource.TestStep{
                        {
                                Config: config,
                                Check: resource.ComposeTestCheckFunc(
                                        resource.TestCheckResourceAttr(dataSourceName, "day_of_month", strconv.Itoa(dayofmonth)),
                                        resource.TestCheckResourceAttr(dataSourceName, "hour", strconv.Itoa(now.Hour())),
                                        resource.TestCheckResourceAttr(dataSourceName, "minute", "0"),
                                        resource.TestCheckResourceAttr(dataSourceName, "second", "0"),
                                        resource.TestCheckResourceAttr(dataSourceName, "frequency", "Month"),
                                        resource.TestCheckResourceAttr(dataSourceName, "next_run_time", formattedExpectedTime),
                                ),
                        },
                },
        })
}

func TestAccDataSourceAzureRMScheduledTime_Monthly_Last_Month(t *testing.T) {
        dataSourceName := "data.azurerm_scheduled_time.test"

        now := time.Now().UTC()
        scheduletime := now.Add(-25 * time.Duration(24) * time.Hour) //25 days
        dayofmonth := scheduletime.Day()

        config := testAccDataSourceAzureRMScheduledTime_Monthly(dayofmonth, scheduletime)
	var expectedTime time.Time

	if now.Day() < 25 {
		expectedTime = time.Date(scheduletime.Year(), now.Month() + 1, scheduletime.Day(), scheduletime.Hour(), 0, 0, 0, time.UTC)
	} else {
        	expectedTime = time.Date(scheduletime.Year(), now.Month(), scheduletime.Day(), scheduletime.Hour(), 0, 0, 0, time.UTC)
	}
        formattedExpectedTime := expectedTime.Format(time.RFC3339)

        resource.Test(t, resource.TestCase{
                PreCheck:  func() { testAccPreCheck(t) },
                Providers: testAccProviders,
                Steps: []resource.TestStep{
                        {
                                Config: config,
                                Check: resource.ComposeTestCheckFunc(
                                        resource.TestCheckResourceAttr(dataSourceName, "day_of_month", strconv.Itoa(dayofmonth)),
                                        resource.TestCheckResourceAttr(dataSourceName, "hour", strconv.Itoa(now.Hour())),
                                        resource.TestCheckResourceAttr(dataSourceName, "minute", "0"),
                                        resource.TestCheckResourceAttr(dataSourceName, "second", "0"),
                                        resource.TestCheckResourceAttr(dataSourceName, "frequency", "Month"),
                                        resource.TestCheckResourceAttr(dataSourceName, "next_run_time", formattedExpectedTime),
                                ),
                        },
                },
        })
}

func testAccDataSourceAzureRMScheduledTime_Monthly(dayofmonth int, scheduletime time.Time) string {
        return fmt.Sprintf(`
data "azurerm_scheduled_time" "test" {
        "day_of_month" = "%d"
        "hour" = "%d"
        "minute" = "0"
        "second" = "0"
        "frequency" = "Month"
}
`, dayofmonth, scheduletime.Hour())
}

func testAccDataSourceAzureRMScheduledTime_Weekly(dayofweek int, scheduletime time.Time) string {
	return fmt.Sprintf(`
data "azurerm_scheduled_time" "test" {
	"day_of_week" = "%d"
        "hour" = "%d"
        "minute" = "0"
        "second" = "0"
        "frequency" = "Week"
}
`, dayofweek, scheduletime.Hour())
}

func testAccDataSourceAzureRMScheduledTime_Daily(scheduletime time.Time) string {
	return fmt.Sprintf(`
data "azurerm_scheduled_time" "test" {
	"hour" = "%d"		
        "minute" = "0"		
        "second" = "0"
	"frequency" = "Day"
}
`, scheduletime.Hour())
}

func testAccDataSourceAzureRMScheduledTime_Hourly(scheduletime time.Time) string {
	return fmt.Sprintf(`
data "azurerm_scheduled_time" "test" {
        "hour" = "%d"
        "minute" = "%d"
        "second" = "0"
        "frequency" = "Hour"
}
`, scheduletime.Hour(), scheduletime.Minute())
}

func testAccDataSourceAzureRMScheduledTime_Hourly_With_Seconds(scheduletime time.Time) string {
	return fmt.Sprintf(`
data "azurerm_scheduled_time" "test" {
        "hour" = "%d"
        "minute" = "%d"
        "second" = "24"
        "frequency" = "Hour"
}
`, scheduletime.Hour(), scheduletime.Minute())
}
