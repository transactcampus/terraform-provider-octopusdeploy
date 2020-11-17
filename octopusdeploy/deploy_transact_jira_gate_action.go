package octopusdeploy

import (
	"encoding/json"
	"github.com/transactcampus/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getDeployTransactJiraGateActionSchema() *schema.Schema {

	actionSchema, element := getCommonDeploymentActionSchema()
	addExecutionLocationSchema(element)
	element.Schema["auth_key"] = &schema.Schema{
		Type:        schema.TypeString,
		Description: "The name of the secret resource",
		Required:    true,
	}

	element.Schema["required_status"] = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"key": {
					Type:     schema.TypeString,
					Required: true,
				},
				"value": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	}

	return actionSchema
}

func buildDeployTransactJiraGateActionResource(tfAction map[string]interface{}) octopusdeploy.DeploymentAction {
	resource := buildDeploymentActionResource(tfAction)

	resource.ActionType = "Octopus.Script"

	resource.Properties["Octopus.Action.JiraGate.AuthKey"] = tfAction["auth_key"].(string)

	//resource.Properties["Octopus.Action.Script.ScriptBody"] = "Write-Host 'hi'".(string)


	if tfSecretValues, ok := tfAction["required_status"]; ok {

		secretValues := make(map[string]string)

		for _, tfSecretValue := range tfSecretValues.([]interface{}) {
			tfSecretValueTyped := tfSecretValue.(map[string]interface{})
			secretValues[tfSecretValueTyped["key"].(string)] = tfSecretValueTyped["value"].(string)
		}

		j, _ := json.Marshal(secretValues)

		resource.Properties["Octopus.Action.JiraGate.RequiredStatus"] = string(j)
	}

	return resource
}