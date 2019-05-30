# Terraform Module for MariaDB on Azure

Terraform module to create a managed instance of a database in MariaDB on Azure.

## Usage

## Running tests
In order to run the tests, I created two ways of doing that: Native and using Docker,
and I'm explaining how to run each one of these down below.

### Configuring
In order to configure Terraform on your machine and use it against Azure, you can
follow this [official Azure docs](https://docs.microsoft.com/en-us/azure/virtual-machines/linux/terraform-install-configure).

### Local
#### Prerequisites
#### Run test (Native)
### Docker
#### Prerequisites
#### Build the image
#### Run test (Docker)
Pass the tfvars with the variables
### Credits
Following the pattern implemented on the postgresql module, that can be
found [link](https://github.com/Azure/terraform-azurerm-postgresql/)