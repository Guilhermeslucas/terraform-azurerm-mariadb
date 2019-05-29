package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformMariaDB(t *testing.T) {
	var dbNames = "teste"
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "./fixture",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"location": "west us 2",
			"db_name":  dbNames,
		},
	}
	defer terraform.Destroy(t, terraformOptions)
}
