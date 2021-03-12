package octopusdeploy

import (
	"github.com/transactcampus/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandExternalSecurityGroups(externalSecurityGroups []interface{}) []octopusdeploy.NamedReferenceItem {
	expandedExternalSecurityGroups := make([]octopusdeploy.NamedReferenceItem, 0, len(externalSecurityGroups))
	for _, externalSecurityGroup := range externalSecurityGroups {
		if externalSecurityGroup != nil {
			rawExternalSecurityGroup := externalSecurityGroup.(map[string]interface{})

			displayIDAndName := false
			if rawExternalSecurityGroup["display_id_and_name"] != nil {
				displayIDAndName = rawExternalSecurityGroup["display_id_and_name"].(bool)
			}

			displayName := ""
			if rawExternalSecurityGroup["display_name"] != nil {
				displayName = rawExternalSecurityGroup["display_name"].(string)
			}

			id := ""
			if rawExternalSecurityGroup["id"] != nil {
				id = rawExternalSecurityGroup["id"].(string)
			}

			item := octopusdeploy.NamedReferenceItem{
				DisplayIDAndName: displayIDAndName,
				DisplayName:      displayName,
				ID:               id,
			}
			expandedExternalSecurityGroups = append(expandedExternalSecurityGroups, item)
		}
	}
	return expandedExternalSecurityGroups
}

func flattenExternalSecurityGroups(externalSecurityGroups []octopusdeploy.NamedReferenceItem) []interface{} {
	if externalSecurityGroups == nil {
		return nil
	}

	flattenedExternalSecurityGroups := make([]interface{}, len(externalSecurityGroups))
	for i, externalSecurityGroup := range externalSecurityGroups {
		rawExternalSecurityGroup := map[string]interface{}{
			"display_id_and_name": externalSecurityGroup.DisplayIDAndName,
			"display_name":        externalSecurityGroup.DisplayName,
			"id":                  externalSecurityGroup.ID,
		}

		flattenedExternalSecurityGroups[i] = rawExternalSecurityGroup
	}

	return flattenedExternalSecurityGroups
}

func getExternalSecurityGroupsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"display_id_and_name": {
			Computed: true,
			Optional: true,
			Type:     schema.TypeBool,
		},
		"display_name": {
			Computed: true,
			Optional: true,
			Type:     schema.TypeString,
		},
		"id": getIDSchema(),
	}
}
