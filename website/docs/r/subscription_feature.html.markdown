---
subcategory: "Base"
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_subscription_feature"
description: |-
  Manages a Feature.
---

# azurerm_subscription_feature

Register and unregister a preview feature for the subscription. Features which `approvalType` is `AutoApproval` can be managed by terraform.  [More information can be found in this document](https://docs.microsoft.com/en-us/rest/api/resources/features).

## Example Usage

```hcl
resource "azurerm_subscription_feature" "example" {
  name               = "AutoApproveFeature"
  provider_namespace = "Microsoft.CognitiveServices"
}
```

## Arguments Reference

The following arguments are supported:

* `name` - (Required) The name of the feature to register. Changing this forces a new Feature to be created.

* `provider_namespace` - (Required) The namespace of the resource provider. Changing this forces a new Feature to be created.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported: 

* `id` - The ID of the Feature.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 30 minutes) Used when creating the Feature.
* `read` - (Defaults to 5 minutes) Used when retrieving the Feature.
* `update` - (Defaults to 30 minutes) Used when updating the Feature.
* `delete` - (Defaults to 30 minutes) Used when deleting the Feature.

## Import

Features can be imported using the `resource id`, e.g.

```shell
terraform import azurerm_subscription_feature.example /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Features/providers/Microsoft.Compute/features/AllowManagedDisksReplaceOSDisk
```
