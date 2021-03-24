package test

import (
  "github.com/gruntwork-io/terratest/modules/terraform"
  "github.com/stretchr/testify/assert"
  "testing"
)

func vnet_test(t *testing.T) {

    terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../appserviceterraform",
	})

    defer terraform.Destroy(t, terraformOptions)

    terraform.InitAndApply(t, terraformOptions)
    output := terraform.Output(t, terraformOptions, "hello_world")
    appname :=terraform.Output(t, terraformOptions, "ServicePlanName")
	  assert.Equal(t, "Hello, World!", output)
    assert.Equal(t, "", appname)
}