package octopusdeploy

import (
	"context"

	"github.com/transactcampus/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandAmazonWebServicesAccount(d *schema.ResourceData) *octopusdeploy.AmazonWebServicesAccount {
	name := d.Get("name").(string)
	accessKey := d.Get("access_key").(string)
	password := d.Get("secret_key").(string)
	secretKey := octopusdeploy.NewSensitiveValue(password)

	account, _ := octopusdeploy.NewAmazonWebServicesAccount(name, accessKey, secretKey)
	account.ID = d.Id()

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

func flattenAmazonWebServicesAccount(ctx context.Context, d *schema.ResourceData, account *octopusdeploy.AmazonWebServicesAccount) {
	flattenAccount(ctx, d, account)

	d.Set("account_type", "AmazonWebServicesAccount")
	d.Set("access_key", account.AccessKey)
	d.Set("secret_key", account.SecretKey)

	d.SetId(account.GetID())
}

func getAmazonWebServicesAccountSchema() map[string]*schema.Schema {
	schemaMap := getAccountSchema()
	schemaMap["access_key"] = &schema.Schema{
		Required:  true,
		Sensitive: true,
		Type:      schema.TypeString,
	}
	schemaMap["account_type"] = &schema.Schema{
		Optional: true,
		Default:  "AmazonWebServicesAccount",
		Type:     schema.TypeString,
	}
	schemaMap["secret_key"] = &schema.Schema{
		Required:  true,
		Sensitive: true,
		Type:      schema.TypeString,
	}
	return schemaMap
}
