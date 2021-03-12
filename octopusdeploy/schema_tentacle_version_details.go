package octopusdeploy

import (
	"github.com/transactcampus/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandTentacleVersionDetails(values interface{}) *octopusdeploy.TentacleVersionDetails {
	if values == nil {
		return nil
	}

	flattenedValues := values.([]interface{})
	flattenedTentacleVersionDetails := flattenedValues[0].(map[string]interface{})

	version := flattenedTentacleVersionDetails["version"].(string)

	return &octopusdeploy.TentacleVersionDetails{
		UpgradeLocked:    flattenedTentacleVersionDetails["upgrade_locked"].(bool),
		UpgradeRequired:  flattenedTentacleVersionDetails["upgrade_required"].(bool),
		UpgradeSuggested: flattenedTentacleVersionDetails["upgrade_suggested"].(bool),
		Version:          &version,
	}
}

func flattenTentacleVersionDetails(tentacleVersionDetails *octopusdeploy.TentacleVersionDetails) []interface{} {
	if tentacleVersionDetails == nil {
		return nil
	}

	return []interface{}{map[string]interface{}{
		"upgrade_locked":    tentacleVersionDetails.UpgradeLocked,
		"upgrade_required":  tentacleVersionDetails.UpgradeRequired,
		"upgrade_suggested": tentacleVersionDetails.UpgradeSuggested,
		"version":           tentacleVersionDetails.Version,
	}}
}

func getTentacleVersionDetailsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"upgrade_locked": {
			Computed: true,
			Optional: true,
			Type:     schema.TypeBool,
		},
		"upgrade_required": {
			Computed: true,
			Optional: true,
			Type:     schema.TypeBool,
		},
		"upgrade_suggested": {
			Computed: true,
			Optional: true,
			Type:     schema.TypeBool,
		},
		"version": {
			Computed: true,
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}
