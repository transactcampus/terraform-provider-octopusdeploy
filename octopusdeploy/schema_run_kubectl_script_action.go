package octopusdeploy

import (
	"github.com/transactcampus/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getRunKubectlScriptSchema() *schema.Schema {
	actionSchema, element := getCommonDeploymentActionSchema()
	addExecutionLocationSchema(element)
	addScriptFromPackageSchema(element)
	addPackagesSchema(element, false)
	return actionSchema
}

func expandRunKubectlScriptAction(tfAction map[string]interface{}) octopusdeploy.DeploymentAction {
	resource := expandDeploymentAction(tfAction)
	resource.ActionType = "Octopus.KubernetesRunScript"
	resource.Properties = merge(resource.Properties, buildScriptFromPackageProperties(tfAction))
	return resource
}

func merge(map1 map[string]string, map2 map[string]string) map[string]string {
	result := make(map[string]string)

	for k, v := range map1 {
		result[k] = v
	}

	for k, v := range map2 {
		result[k] = v
	}

	return result
}

func buildScriptFromPackageProperties(tfAction map[string]interface{}) map[string]string {
	properties := make(map[string]string)
	properties["Octopus.Action.Script.ScriptFileName"] = tfAction["script_file_name"].(string)
	properties["Octopus.Action.Script.ScriptParameters"] = tfAction["script_parameters"].(string)
	properties["Octopus.Action.Script.ScriptSource"] = "Package"
	return properties
}
