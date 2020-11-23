package octopusdeploy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/transactcampus/go-octopusdeploy/octopusdeploy"
)

func expandEnvironment(d *schema.ResourceData) *octopusdeploy.Environment {
	name := d.Get("name").(string)

	environment := octopusdeploy.NewEnvironment(name)
	environment.ID = d.Id()

	if v, ok := d.GetOk("allow_dynamic_infrastructure"); ok {
		environment.AllowDynamicInfrastructure = v.(bool)
	}

	if v, ok := d.GetOk("description"); ok {
		environment.Description = v.(string)
	}

	if v, ok := d.GetOk("sort_order"); ok {
		environment.SortOrder = v.(int)
	}

	if v, ok := d.GetOk("use_guided_failure"); ok {
		environment.UseGuidedFailure = v.(bool)
	}

	if v, ok := d.GetOk("extension_settings"); ok {
		environment.ExtensionSettings = expandEnvironmentExtensionSettingsValues(v.(*schema.Set).List())
	}

	return environment
}

func expandEnvironmentExtensionSettingsValues(extensionSettingsValues []interface{}) []*octopusdeploy.ExtensionSettingsValues {
	expandedExtensionSettingsValues := make([]*octopusdeploy.ExtensionSettingsValues, len(extensionSettingsValues))
	for _, extensionSettingsValue := range extensionSettingsValues {
		extensionSettingsValueMap := extensionSettingsValue.(map[string]interface{})
		expandedExtensionSettingsValues = append(expandedExtensionSettingsValues, &octopusdeploy.ExtensionSettingsValues{
			ExtensionID: extensionSettingsValueMap["extension_id"].(string),
			Values:      extensionSettingsValueMap["values"].([]interface{}),
		})
	}
	return expandedExtensionSettingsValues
}

func flattenEnvironment(ctx context.Context, d *schema.ResourceData, environment *octopusdeploy.Environment) {
	d.Set("allow_dynamic_infrastructure", environment.AllowDynamicInfrastructure)
	d.Set("description", environment.Description)
	d.Set("name", environment.Name)
	d.Set("sort_order", environment.SortOrder)
	d.Set("use_guided_failure", environment.UseGuidedFailure)
	d.Set("extension_settings", environment.ExtensionSettings)

	d.SetId(environment.GetID())
}

func getEnvironmentExtensionSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"extension_id": {
			Computed: true,
			Type:     schema.TypeString,
		},
		"values": {
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Type:     schema.TypeList,
		},
	}
}

func getEnvironmentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_dynamic_infrastructure": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"description": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"name": &schema.Schema{
			Required:     true,
			Type:         schema.TypeString,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"sort_order": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  false,
		},
		"use_guided_failure": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"extension_settings": {
			Optional: true,
			Elem:     &schema.Resource{Schema: getEnvironmentExtensionSettingsSchema()},
			Type:     schema.TypeSet,
		},
	}
}
