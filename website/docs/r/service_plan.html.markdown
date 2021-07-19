---
subcategory: "App Service (Web Apps)"
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_service_plan"
description: |-
  Manages an App Service: Service Plan.
---

# azurerm_service_plan

Manages an App Service: Service Plan.

!> **NOTE:** This is a 3.0 beta resource and is intended to replace the older `azurerm_app_service_plan`. This older resource will be removed in 3.0 and this resource will be renamed to supersede it.

## Example Usage

```hcl
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "example" {
  name     = "example-resources"
  location = "West Europe"
}

resource "azurerm_service_plan" "example" {
  name                = "example"
  resource_group_name = azurerm_resource_group.example.name
  location            = "West Europe"
  sku_name            = "P1V2"
}
```

## Arguments Reference

The following arguments are supported:

* `name` - (Required) The name which should be used for this Service Plan. Changing this forces a new AppService to be created.

* `location` - (Required) The Azure Region where the Service Plan should exist. Changing this forces a new AppService to be created.

* `resource_group_name` - (Required) The name of the Resource Group where the AppService should exist. Changing this forces a new AppService to be created.

* `sku_name` - (Required) The SKU for the plan. Possible values include `B1`, `B2`, `B3`, `D1`, `F1`, `FREE`, `I1`, `I2`, `I3`, `I1v2`, `I2v2`, `I3v2`, `P1v2`, `P2v2`, `P3v2`, `P1v3`, `P2v3`, `P3v3`, `S1`, `S2`, `S3`, `SHARED`, `PC2`, `PC3`, `PC4`, `EP1`, `EP2`, and `EP3`. 

~> **NOTE:** Isolated SKUs (`I1`, `I2`, `I3`, `I1v2`, `I2v2`, and `I3v2`) can only be used with App Service Environments

~> **NOTE:** Elastic and Consumption SKUs (`PC2`, `PC3`, `PC4`, `EP1`, `EP2`, and `EP3`) are for use with Function Apps.

---

* `app_service_environment_id` - (Optional) The ID of the App Service Plan to create this plan in.

~> **NOTE:** Requires an Isolated SKU. Use one of `I1`, `I2`, `I3` for `azurerm_app_service_environment`, or `I1v2`, `I2v2`, `I3v2` for `azurerm_app_service_environment_v3`

* `maximum_elastic_worker_count` - (Optional) The maximum number of workers to use in an Elastic SKU Plan. Cannot be set unless using an Elastic SKU.

* `number_of_workers` - (Optional) The number of Workers (instances) to be allocated. 

* `os_type` - (Optional) The O/S type for the App Services to be hosted in this plan. Possible values include `Windows` (default), `Linux`, and `WindowsContainer`.

* `per_site_scaling` - (Optional) Should Per Site Scaling be enabled. Defaults to `false`.

* `tags` - (Optional) A mapping of tags which should be assigned to the AppService.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported: 

* `id` - The ID of the Service Plan.

* `kind` - A string representing the Kind of Service Plan.

* `reserved` - Whether this is a reserved Service Plan Type. `true` if `os_type` is `Linux`, otherwise `false`. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 1 hour) Used when creating the Service Plan.
* `read` - (Defaults to 5 minutes) Used when retrieving the Service Plan.
* `update` - (Defaults to 1 hour) Used when updating the Service Plan.
* `delete` - (Defaults to 1 hour) Used when deleting the Service Plan.

## Import

AppServices can be imported using the `resource id`, e.g.

```shell
terraform import azurerm_service_plan.example /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/serverfarms/farm1
```
