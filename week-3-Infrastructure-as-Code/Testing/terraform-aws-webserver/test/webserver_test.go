package test

import (
	"fmt"
	"time"
	http_hepler "github.com/gruntwork-io/terratest/modules/http-helper"
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformWebserverExample(t *testing.T){

	// The values to pass into the Terraform CLI
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		// The path to where the example Terraform code is located
		TerraformDir: "../examples/webserver",

		// Variables to pass to the Terraform code using -var options
		Vars: map[string]interface{}{
			"region": "ap-southeast-2",
			"servername": "cstestwebserver2",
		},
	})

	// Run a Terraform init and apply with the Terraform options
	terraform.InitAndApply(t, terraformOptions)

	// Run a Terraform Destroy at the end of the test
	defer terraform.Destroy(t, terraformOptions)

	publicIp := terraform.Output(t, terraformOptions, "public_ip")

	url := fmt.Sprintf("http://%s:8080", publicIp)

	http_hepler.HttpGetWithRetry(t, url, nil, 200, "I made a Terraform Module!", 30, 5*time.Second)
}