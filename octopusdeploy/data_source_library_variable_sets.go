package octopusdeploy

import (
	"context"

	"github.com/transactcampus/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLibraryVariableSet() *schema.Resource {
	return &schema.Resource{
		Description: "Provides information about existing library variable sets.",
		ReadContext: dataSourceLibraryVariableSetReadByName,
		Schema:      getLibraryVariableSetDataSchema(),
	}
}

func dataSourceLibraryVariableSetReadByName(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	client := m.(*octopusdeploy.Client)
	libraryVariableSets, err := client.LibraryVariableSets.GetByPartialName(name)
	if err != nil {
		return diag.FromErr(err)
	}
	if len(libraryVariableSets) == 0 {
		return nil
	}

	// NOTE: two or more library variables can have the same name in Octopus.
	// Therefore, a better search criteria needs to be implemented below.

	for _, libraryVariableSet := range libraryVariableSets {
		if libraryVariableSet.Name == name {
			setLibraryVariableSet(ctx, d, libraryVariableSet)
			return nil
		}
	}

	return nil
}
