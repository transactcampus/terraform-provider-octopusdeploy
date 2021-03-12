package octopusdeploy

import (
	"context"
	"log"

	"github.com/transactcampus/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSSHConnectionDeploymentTarget() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSSHConnectionDeploymentTargetCreate,
		DeleteContext: resourceSSHConnectionDeploymentTargetDelete,
		Description:   "This resource manages SSH connection deployment targets in Octopus Deploy.",
		Importer:      getImporter(),
		ReadContext:   resourceSSHConnectionDeploymentTargetRead,
		Schema:        getSSHConnectionDeploymentTargetSchema(),
		UpdateContext: resourceSSHConnectionDeploymentTargetUpdate,
	}
}

func resourceSSHConnectionDeploymentTargetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	deploymentTarget := expandSSHConnectionDeploymentTarget(d)

	log.Printf("[INFO] creating SSH connection deployment target: %#v", deploymentTarget)

	client := m.(*octopusdeploy.Client)
	createdDeploymentTarget, err := client.Machines.Add(deploymentTarget)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setSSHConnectionDeploymentTarget(ctx, d, createdDeploymentTarget); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(createdDeploymentTarget.GetID())

	log.Printf("[INFO] SSH connection deployment target created (%s)", d.Id())
	return nil
}

func resourceSSHConnectionDeploymentTargetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[INFO] deleting SSH connection deployment target (%s)", d.Id())

	client := m.(*octopusdeploy.Client)
	if err := client.Machines.DeleteByID(d.Id()); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	log.Printf("[INFO] SSH connection deployment target deleted")
	return nil
}

func resourceSSHConnectionDeploymentTargetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[INFO] reading SSH connection deployment target (%s)", d.Id())

	client := m.(*octopusdeploy.Client)
	deploymentTarget, err := client.Machines.GetByID(d.Id())
	if err != nil {
		apiError := err.(*octopusdeploy.APIError)
		if apiError.StatusCode == 404 {
			log.Printf("[INFO] SSH connection deployment target (%s) not found; deleting from state", d.Id())
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	if err := setSSHConnectionDeploymentTarget(ctx, d, deploymentTarget); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[INFO] SSH connection deployment target read (%s)", d.Id())
	return nil
}

func resourceSSHConnectionDeploymentTargetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[INFO] updating SSH connection deployment target (%s)", d.Id())

	deploymentTarget := expandSSHConnectionDeploymentTarget(d)
	client := m.(*octopusdeploy.Client)
	updatedDeploymentTarget, err := client.Machines.Update(deploymentTarget)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setSSHConnectionDeploymentTarget(ctx, d, updatedDeploymentTarget); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[INFO] SSH connection deployment target updated (%s)", d.Id())
	return nil
}
