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
To run the tests locally in your machine, there are two things you need to install?

- [Terraform **(~> 0.11.7)**](https://www.terraform.io/downloads.html)
- [Golang **(~> 1.10.3)**](https://golang.org/dl/)

#### Run test (Native)
I added a script called ```test.sh``` in order to make testing the module easier.

There are two ways of running this script:
- **validate**: this mode will check if the terraform you have is a valid one. Besides,
it will check for go dependencies that we are using for testing the module.
- **full**: will execute all the validate steps and will run the integration tests using
terratest. It will create the infrastructure and do some testing with it. Note: when 
running the full command, be sure that you have followed the official docs mentioned above
and authorized your command line to create resources on Azure.

So, to test, make the scrpit executabe

``` bash
chmod +x test.sh
```

Then, just run with the desired option. It can be one of both ways below:

##### Validate

``` bash
./test.sh validate
```

##### Full

``` bash
./test.sh full
```

### Docker

#### Prerequisites

#### Build the image

#### Run test (Docker)

## Credits
To create this reposiotry, I followed the patterns and practices from the 
[postgresql module for Azure](https://github.com/Azure/terraform-azurerm-postgresql/).
If you are willing to develop a terraform module, this repo is a nice way to start. 