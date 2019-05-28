resource "azurerm_resource_group" "test" {
  name     = "api-rg-pro"
  location = "West Europe"
}

resource "azurerm_mariadb_server" "test" {
  name                = "mariadb-server-1"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"

  sku {
    name     = "B_Gen5_2"
    capacity = 2
    tier     = "Basic"
    family   = "Gen5"
  }

  storage_profile {
    storage_mb            = 5120
    backup_retention_days = 7
    geo_redundant_backup  = "Disabled"
  }

  administrator_login          = "mariadbadmin"
  administrator_login_password = "H@Sh1CoR3!"
  version                      = "10.2"
  ssl_enforcement              = "Enabled"
}

resource "azurerm_mariadb_database" "example" {
  name                = "mariadb_database"
  resource_group_name = "${azurerm_resource_group.example.name}"
  server_name         = "${azurerm_mariadb_server.example.name}"
  charset             = "utf8"
  collation           = "utf8_general_ci"
}