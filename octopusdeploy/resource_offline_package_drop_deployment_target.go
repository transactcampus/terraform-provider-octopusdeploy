package octopusdeploy

import (
	"context"
	"log"

	"github.com/transactcampus/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOfflinePackageDropDeploymentTarget() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOfflinePackageDropDeploymentTargetCreate,
		DeleteContext: resourceOfflinePackageDropDeploymentTargetDelete,
		Description:   "This resource manages offline package drop deployment targets in Octopus Deploy.",
		Importer:      getImporter(),
		ReadContext:   resourceOfflinePackageDropDeploymentTargetRead,
		Schema:        getOfflinePackageDropDeploymentTargetSchema(),
		UpdateContext: resourceOfflinePackageDropDeploymentTargetUpdate,
	}
}

func resourceOfflinePackageDropDeploymentTargetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	deploymentTarget := expandOfflinePackageDropDeploymentTarget(d)

	log.Printf("[INFO] creating offline package drop deployment target: %#v", deploymentTarget)

	client := m.(*octopusdeploy.Client)
	createdDeploymentTarget, err := client.Machines.Add(deploymentTarget)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setOfflinePackageDropDeploymentTarget(ctx, d, createdDeploymentTarget); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(createdDeploymentTarget.GetID())

	log.Printf("[INFO] offline package drop deployment target created (%s)", d.Id())
	return nil
}

func resourceOfflinePackageDropDeploymentTargetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[INFO] deleting offline package drop deployment target (%s)", d.Id())

	client := m.(*octopusdeploy.Client)
	if err := client.Machines.DeleteByID(d.Id()); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	log.Printf("[INFO] offline package drop deployment target deleted")
	return nil
}

func resourceOfflinePackageDropDeploymentTargetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[INFO] reading offline package drop deployment target (%s)", d.Id())

	client := m.(*octopusdeploy.Client)
	deploymentTarget, err := client.Machines.GetByID(d.Id())
	if err != nil {
		apiError := err.(*octopusdeploy.APIError)
		if apiError.StatusCode == 404 {
			log.Printf("[INFO] offline package drop deployment target (%s) not found; deleting from state", d.Id())
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	if err := setOfflinePackageDropDeploymentTarget(ctx, d, deploymentTarget); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[INFO] offline package drop deployment target read (%s)", d.Id())
	return nil
}

func resourceOfflinePackageDropDeploymentTargetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[INFO] updating offline package drop deployment target (%s)", d.Id())

	deploymentTarget := expandOfflinePackageDropDeploymentTarget(d)
	client := m.(*octopusdeploy.Client)
	updatedDeploymentTarget, err := client.Machines.Update(deploymentTarget)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setOfflinePackageDropDeploymentTarget(ctx, d, updatedDeploymentTarget); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[INFO] offline package drop deployment target updated (%s)", d.Id())
	return nil
}
