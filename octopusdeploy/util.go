package octopusdeploy

import (
	"fmt"
	"log"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Validate a value against a set of possible values
func validateValueFunc(values []string) schema.SchemaValidateDiagFunc {

	return func(v interface{}, c cty.Path) diag.Diagnostics {
		var diags diag.Diagnostics

		value := v.(string)
		valid := false
		for _, val := range values {
			if value == val {
				valid = true
				break
			}
		}

		if !valid {
			diags = append(diag.Errorf("unexpected: %s", values))
		}
		return diags
	}
}

// validateStringInSlice checks if a string is in the given slice
func validateStringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}

	return false
}

func validateAllSliceItemsInSlice(givenSlice, validationSlice []string) (string, bool) {
	for _, v := range givenSlice {
		if !validateStringInSlice(v, validationSlice) {
			return v, false
		}
	}

	return constEmptyString, true
}

func getSliceFromTerraformTypeList(inputTypeList interface{}) []string {
	var newSlice []string
	terraformList := inputTypeList.([]interface{})
	for _, item := range terraformList {
		newSlice = append(newSlice, item.(string))
	}
	return newSlice
}

func isEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func logResource(name string, resource interface{}) {
	log.Printf("[DEBUG] %s: %v", name, resource)
}

func getStringOrEmpty(tfAttr interface{}) string {
	if tfAttr == nil {
		return constEmptyString
	}
	return tfAttr.(string)
}

func getTenantedDeploymentSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Default:  "Untenanted",
		ValidateDiagFunc: validateValueFunc([]string{
			"Untenanted",
			"TenantedOrUntenanted",
			"Tenanted",
		}),
	}
}

func destroyFeedHelper(s *terraform.State, apiClient *client.Client) error {
	for _, r := range s.RootModule().Resources {
		if _, err := apiClient.Feeds.GetByID(r.Primary.ID); err != nil {
			return fmt.Errorf("Received an error retrieving feed %s", err)
		}
		return fmt.Errorf("Feed still exists")
	}
	return nil
}

func feedExistsHelper(s *terraform.State, apiClient *client.Client) error {
	for _, r := range s.RootModule().Resources {
		if _, err := apiClient.Feeds.GetByID(r.Primary.ID); err != nil {
			return fmt.Errorf("Received an error retrieving feed %s", err)
		}
	}
	return nil
}
