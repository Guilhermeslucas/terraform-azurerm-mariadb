package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformMariaDB(t *testing.T) {
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "./fixture",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"location":               "east us",
			"resource_group_name":    "mariadb-rg",
			"administrator_login":    "awesomeadmin",
			"administrator_password": "awesomepass12345!@#$%",
		},
	}

	//will run in the end of the method, destroying all the resources that
	//terraform has created
	defer terraform.Destroy(t, terraformOptions)

}
