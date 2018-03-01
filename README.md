# AzureRM Terraform Provider

* Website: https://www.terraform.io
* [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
* Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

## General Requirements

* [Terraform](https://www.terraform.io/downloads.html) 0.10.x
* [Go](https://golang.org/doc/install) 1.9 (to build the provider plugin)

## Windows Specific Requirements

* [Make for Windows](http://gnuwin32.sourceforge.net/packages/make.htm)
* [Git Bash for Windows](https://git-scm.com/download/win)

For _GNU32 Make_, make sure its bin path is added to PATH environment variable.\*

For _Git Bash for Windows_, at the step of "Adjusting your PATH environment", please choose "Use Git and optional Unix tools from Windows Command Prompt".\*

## Building The Provider

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-azurerm`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:terraform-providers/terraform-provider-azurerm
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-azurerm
$ make build
```

## Using the provider

```
# Configure the Microsoft Azure Provider
provider "azurerm" {
  subscription_id = "..."
  client_id       = "..."
  client_secret   = "..."
  tenant_id       = "..."
}

# Create a resource group
resource "azurerm_resource_group" "production" {
  name     = "production"
  location = "West US"
}

# Create a virtual network in the web_servers resource group
resource "azurerm_virtual_network" "network" {
  name                = "productionNetwork"
  address_space       = ["10.0.0.0/16"]
  location            = "West US"
  resource_group_name = "${azurerm_resource_group.production.name}"

  subnet {
    name           = "subnet1"
    address_prefix = "10.0.1.0/24"
  }

  subnet {
    name           = "subnet2"
    address_prefix = "10.0.2.0/24"
  }

  subnet {
    name           = "subnet3"
    address_prefix = "10.0.3.0/24"
  }
}
```

Further [usage documentation is available on the Terraform website](https://www.terraform.io/docs/providers/azurerm/index.html).

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.9+ is _required_). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-azurerm
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

_Note:_ Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

In order to run a specific Acceptance test with your own subscription info, run

```sh
TF_ACC=1 \
ARM_SUBSCRIPTION_ID="$ARM_SUBSCRIPTION_ID" \
ARM_CLIENT_ID="$ARM_CLIENT_ID" \
ARM_CLIENT_SECRET="$ARM_CLIENT_SECRET" \
ARM_TENANT_ID="$ARM_TENANT_ID" \
ARM_TEST_LOCATION="$ARM_TEST_LOCATION" \
ARM_TEST_LOCATION_ALT="$ARM_TEST_LOCATION_ALT" \
go test $(go list ./... |grep -v 'vendor') -v -timeout 30m -run TestAccAzureRMVirtualMachineScaleSet_applicationGateway
```
