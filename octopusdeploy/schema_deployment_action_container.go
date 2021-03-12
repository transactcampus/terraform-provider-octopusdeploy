package octopusdeploy

import (
	"github.com/transactcampus/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandDeploymentActionContainer(values interface{}) octopusdeploy.DeploymentActionContainer {
	if values == nil {
		return octopusdeploy.DeploymentActionContainer{}
	}

	flattenedValues := values.([]interface{})
	if len(flattenedValues) == 0 || flattenedValues[0] == nil {
		return octopusdeploy.DeploymentActionContainer{}
	}

	flattenedMap := flattenedValues[0].(map[string]interface{})

	deploymentActionContainer := octopusdeploy.DeploymentActionContainer{}

	if feedID := flattenedMap["feed_id"]; feedID != nil {
		deploymentActionContainer.FeedID = feedID.(string)
	}

	if image := flattenedMap["image"]; image != nil {
		deploymentActionContainer.Image = image.(string)
	}

	return deploymentActionContainer
}

func flattenDeploymentActionContainer(deploymentActionContainer octopusdeploy.DeploymentActionContainer) []interface{} {
	return []interface{}{map[string]interface{}{
		"feed_id": deploymentActionContainer.FeedID,
		"image":   deploymentActionContainer.Image,
	}}
}

func getDeploymentActionContainerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"feed_id": {
			Optional: true,
			Type:     schema.TypeString,
		},
		"image": {
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}
