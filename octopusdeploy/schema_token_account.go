package octopusdeploy

import (
	"context"
	"fmt"

	"github.com/transactcampus/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandTokenAccount(d *schema.ResourceData) *octopusdeploy.TokenAccount {
	name := d.Get("name").(string)
	token := octopusdeploy.NewSensitiveValue(d.Get("token").(string))

	account, _ := octopusdeploy.NewTokenAccount(name, token)
	account.ID = d.Id()

	if v, ok := d.GetOk("description"); ok {
		account.Description = v.(string)
	}

	if v, ok := d.GetOk("environments"); ok {
		account.EnvironmentIDs = getSliceFromTerraformTypeList(v)
	}

	if v, ok := d.GetOk("space_id"); ok {
		account.SpaceID = v.(string)
	}

	if v, ok := d.GetOk("tenanted_deployment_participation"); ok {
		account.TenantedDeploymentMode = octopusdeploy.TenantedDeploymentMode(v.(string))
	}

	if v, ok := d.GetOk("tenant_tags"); ok {
		account.TenantTags = getSliceFromTerraformTypeList(v)
	}

	if v, ok := d.GetOk("tenants"); ok {
		account.TenantIDs = getSliceFromTerraformTypeList(v)
	}

	return account
}

func getTokenAccountSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description":                       getDescriptionSchema(),
		"environments":                      getEnvironmentsSchema(),
		"id":                                getIDSchema(),
		"name":                              getNameSchema(true),
		"space_id":                          getSpaceIDSchema(),
		"tenanted_deployment_participation": getTenantedDeploymentSchema(),
		"tenants":                           getTenantsSchema(),
		"tenant_tags":                       getTenantTagsSchema(),
		"token":                             getTokenSchema(true),
	}
}

func setTokenAccount(ctx context.Context, d *schema.ResourceData, account *octopusdeploy.TokenAccount) error {
	d.Set("description", account.GetDescription())

	if err := d.Set("environments", account.GetEnvironmentIDs()); err != nil {
		return fmt.Errorf("error setting environments: %s", err)
	}

	d.Set("id", account.GetID())
	d.Set("name", account.GetName())
	d.Set("space_id", account.GetSpaceID())
	d.Set("tenanted_deployment_participation", account.GetTenantedDeploymentMode())

	if err := d.Set("tenants", account.GetTenantIDs()); err != nil {
		return fmt.Errorf("error setting tenants: %s", err)
	}

	if err := d.Set("tenant_tags", account.TenantTags); err != nil {
		return fmt.Errorf("error setting tenant_tags: %s", err)
	}

	return nil
}
