---
subcategory: "Key Vault"
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_key_vault_managed_storage_account"
description: |-
  Manages a Key Vault Managed Storage Account.
---

# azurerm_key_vault_managed_storage_account

Manages a Key Vault Managed Storage Account.

## Example Usage

```hcl
data "azurerm_client_config" "example" {}

resource "azurerm_resource_group" "example" {
  name     = "example-resources"
  location = "West Europe"
}

resource "azurerm_storage_account" "example" {
  name                     = "storageaccountname"
  resource_group_name      = azurerm_resource_group.example.name
  location                 = azurerm_resource_group.example.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
}

data "azurerm_storage_account_sas" "example" {
  connection_string = azurerm_storage_account.example.primary_connection_string
  https_only        = true

  resource_types {
    service   = true
    container = false
    object    = false
  }

  services {
    blob  = true
    queue = false
    table = false
    file  = false
  }

  start  = "2021-04-30T00:00:00Z"
  expiry = "2023-04-30T00:00:00Z"

  permissions {
    read    = true
    write   = true
    delete  = false
    list    = false
    add     = true
    create  = true
    update  = false
    process = false
  }
}

resource "azurerm_key_vault" "example" {
  name                = ""
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  tenant_id           = data.azurerm_client_config.current.tenant_id
  sku_name            = "standard"

  access_policy {
    tenant_id = data.azurerm_client_config.current.tenant_id
    object_id = data.azurerm_client_config.current.object_id

    secret_permissions = [
      "Get",
      "Delete"
    ]

    storage_permissions = [
      "Get",
      "List",
      "Set",
      "SetSAS",
      "GetSAS",
      "DeleteSAS",
      "Update",
      "RegenerateKey"
    ]
  }
}

resource "azurerm_key_vault_managed_storage_account" "example" {
  name                = "examplemanagedstorage"
  key_vault_id        = azurerm_key_vault.example.id
  storage_account_id  = azurerm_storage_account.example.id
  storage_account_key = "key1"
  auto_regenerate_key = false
}
```

## Example Usage (automatically regenerate Storage Account access key)

```hcl
data "azurerm_client_config" "example" {}

resource "azurerm_resource_group" "example" {
  name     = "example-resources"
  location = "West Europe"
}

resource "azurerm_storage_account" "example" {
  name                     = "storageaccountname"
  resource_group_name      = azurerm_resource_group.example.name
  location                 = azurerm_resource_group.example.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
}

data "azurerm_storage_account_sas" "example" {
  connection_string = azurerm_storage_account.example.primary_connection_string
  https_only        = true

  resource_types {
    service   = true
    container = false
    object    = false
  }

  services {
    blob  = true
    queue = false
    table = false
    file  = false
  }

  start  = "2021-04-30T00:00:00Z"
  expiry = "2023-04-30T00:00:00Z"

  permissions {
    read    = true
    write   = true
    delete  = false
    list    = false
    add     = true
    create  = true
    update  = false
    process = false
  }
}

resource "azurerm_key_vault" "example" {
  name                = ""
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  tenant_id           = data.azurerm_client_config.current.tenant_id
  sku_name            = "standard"

  access_policy {
    tenant_id = data.azurerm_client_config.current.tenant_id
    object_id = data.azurerm_client_config.current.object_id

    secret_permissions = [
      "Get",
      "Delete"
    ]

    storage_permissions = [
      "Get",
      "List",
      "Set",
      "SetSAS",
      "GetSAS",
      "DeleteSAS",
      "Update",
      "RegenerateKey"
    ]
  }
}

resource "azurerm_role_assignment" "example" {
  scope                = azurerm_storage_account.example.id
  role_definition_name = "Storage Account Key Operator Service Role"
  principal_id         = "727055f9-0386-4ccb-bcf1-9237237ee102"
}

resource "azurerm_key_vault_managed_storage_account" "example" {
  name                = "examplemanagedstorage"
  key_vault_id        = azurerm_key_vault.example.id
  storage_account_id  = azurerm_storage_account.example.id
  storage_account_key = "key1"
  auto_regenerate_key = true
  regeneration_period = "P1D"
}
```

## Arguments Reference

The following arguments are supported:

* `name` - (Required) The name which should be used for this Key Vault Managed Storage Account. Changing this forces a new Key Vault Managed Storage Account to be created.

* `key_vault_id` - (Required) The ID of the Key Vault where the Managed Storage Account should be created. Changing this forces a new resource to be created.

* `storage_account_id` - (Required) The ID of the Storage Account.

* `storage_account_key` - (Required) Which Storage Account access key that is managed by Key Vault. Possible values are `key1` and `key2`.

---

* `auto_regenerate_key` - (Optional) Should Storage Account access key be regenerated periodically?

~> **NOTE:** Azure Key Vault application needs to have access to Storage Account for auto regeneration to work. Example can be found above.

* `regeneration_period` - (Optional) How often Storage Account access key should be regenerated. Value needs to be in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations).

* `tags` - (Optional) A mapping of tags which should be assigned to the Key Vault Managed Storage Account.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported: 

* `id` - The ID of the Key Vault Managed Storage Account.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 30 minutes) Used when creating the Key Vault Managed Storage Account.
* `read` - (Defaults to 5 minutes) Used when retrieving the Key Vault Managed Storage Account.
* `update` - (Defaults to 30 minutes) Used when updating the Key Vault Managed Storage Account.
* `delete` - (Defaults to 30 minutes) Used when deleting the Key Vault Managed Storage Account.

## Import

Key Vault Managed Storage Accounts can be imported using the `resource id`, e.g.

```shell
terraform import azurerm_key_vault_managed_storage_account.example https://example-keyvault.vault.azure.net/storage/exampleStorageAcc01
```
