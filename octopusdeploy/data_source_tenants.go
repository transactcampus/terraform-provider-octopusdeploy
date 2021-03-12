package octopusdeploy

import (
	"context"
	"time"

	"github.com/transactcampus/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTenants() *schema.Resource {
	return &schema.Resource{
		Description: "Provides information about existing tenants.",
		ReadContext: dataSourceTenantsRead,
		Schema:      getTenantDataSchema(),
	}
}

func dataSourceTenantsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	query := octopusdeploy.TenantsQuery{
		ClonedFromTenantID: d.Get("cloned_from_tenant_id").(string),
		IDs:                expandArray(d.Get("ids").([]interface{})),
		IsClone:            d.Get("is_clone").(bool),
		Name:               d.Get("name").(string),
		PartialName:        d.Get("partial_name").(string),
		ProjectID:          d.Get("project_id").(string),
		Skip:               d.Get("skip").(int),
		Tags:               expandArray(d.Get("tags").([]interface{})),
		Take:               d.Get("take").(int),
	}

	client := meta.(*octopusdeploy.Client)
	tenants, err := client.Tenants.Get(query)
	if err != nil {
		return diag.FromErr(err)
	}

	flattenedTenants := []interface{}{}
	for _, tenant := range tenants.Items {
		flattenedTenants = append(flattenedTenants, flattenTenant(tenant))
	}

	d.Set("tenants", flattenedTenants)
	d.SetId("Tenants " + time.Now().UTC().String())

	return nil
}
