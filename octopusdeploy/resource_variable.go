package octopusdeploy

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/transactcampus/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var mutex = &sync.Mutex{}

func resourceVariable() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVariableCreate,
		DeleteContext: resourceVariableDelete,
		Importer: &schema.ResourceImporter{
			State: resourceVariableImport,
		},
		ReadContext:   resourceVariableRead,
		Schema:        getVariableSchema(),
		UpdateContext: resourceVariableUpdate,
	}
}

func resourceVariableImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	importStrings := strings.Split(d.Id(), ":")
	if len(importStrings) != 2 {
		return nil, fmt.Errorf("octopusdeploy_variable import must be in the form of ProjectID:VariableID (e.g. Projects-62:0906031f-68ba-4a15-afaa-657c1564e07b")
	}

	d.Set("project_id", importStrings[0])
	d.SetId(importStrings[1])

	return []*schema.ResourceData{d}, nil
}

func resourceVariableRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[INFO] reading variable (%s)", d.Id())

	id := d.Id()
	projectID := d.Get("project_id").(string)

	client := m.(*octopusdeploy.Client)
	variable, err := client.Variables.GetByID(projectID, id)
	if err != nil {
		apiError := err.(*octopusdeploy.APIError)
		if apiError.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	d.Set("name", variable.Name)
	d.Set("type", variable.Type)

	isSensitive := d.Get("is_sensitive").(bool)
	if isSensitive {
		d.Set("value", nil)
	} else {
		d.Set("value", variable.Value)
	}

	d.Set("description", variable.Description)

	log.Printf("[INFO] variable read (%s)", d.Id())
	return nil
}

func resourceVariableCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	mutex.Lock()
	defer mutex.Unlock()
	if err := validateVariable(d); err != nil {
		return diag.FromErr(err)
	}

	projID := d.Get("project_id").(string)
	newVariable := expandVariable(d)

	log.Printf("[INFO] creating variable: %#v", newVariable)

	client := m.(*octopusdeploy.Client)
	tfVar, err := client.Variables.AddSingle(projID, newVariable)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, v := range tfVar.Variables {
		if v.Name == newVariable.Name && v.Type == newVariable.Type && (v.IsSensitive || v.Value == newVariable.Value) && v.Description == newVariable.Description && v.IsSensitive == newVariable.IsSensitive {
			scopeMatches, _, err := client.Variables.MatchesScope(v.Scope, newVariable.Scope)
			if err != nil {
				return diag.FromErr(err)
			}
			if scopeMatches {
				d.SetId(v.ID)
				log.Printf("[INFO] variable created (%s)", d.Id())
				return nil
			}
		}
	}

	d.SetId("")
	return diag.Errorf("unable to locate variable in project %s", projID)
}

func resourceVariableUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[INFO] updating variable (%s)", d.Id())

	mutex.Lock()
	defer mutex.Unlock()

	if err := validateVariable(d); err != nil {
		return diag.FromErr(err)
	}

	tfVar := expandVariable(d)
	projID := d.Get("project_id").(string)

	client := m.(*octopusdeploy.Client)
	updatedVars, err := client.Variables.UpdateSingle(projID, tfVar)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, v := range updatedVars.Variables {
		if v.Name == tfVar.Name && v.Type == tfVar.Type && (v.IsSensitive || v.Value == tfVar.Value) && v.Description == tfVar.Description && v.IsSensitive == tfVar.IsSensitive {
			scopeMatches, _, _ := client.Variables.MatchesScope(v.Scope, tfVar.Scope)
			if scopeMatches {
				d.SetId(v.ID)
				log.Printf("[INFO] variable updated (%s)", d.Id())
				return nil
			}
		}
	}

	d.SetId("")
	return diag.Errorf("unable to locate variable in project %s", projID)
}

func resourceVariableDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[INFO] deleting variable (%s)", d.Id())

	mutex.Lock()
	defer mutex.Unlock()

	projID := d.Get("project_id").(string)

	client := m.(*octopusdeploy.Client)
	_, err := client.Variables.DeleteSingle(projID, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	log.Printf("[INFO] variable deleted")
	return nil
}

// Validating is done in its own function as we need to compare options once the entire
// schema has been parsed, which as far as I can tell we can't do in a normal validation
// function.
func validateVariable(d *schema.ResourceData) error {
	tfSensitive := d.Get("is_sensitive").(bool)
	tfType := d.Get("type").(string)

	if tfSensitive && tfType != "Sensitive" {
		return fmt.Errorf("when is_sensitive is set to true, type needs to be 'Sensitive'")
	}

	if !tfSensitive && tfType == "Sensitive" {
		return fmt.Errorf("when type is set to 'Sensitive', is_sensitive needs to be true")
	}

	return nil
}
