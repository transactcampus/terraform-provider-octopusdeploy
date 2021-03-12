package octopusdeploy

import (
	"context"

	"github.com/transactcampus/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandCloudRegionDeploymentTarget(d *schema.ResourceData) *octopusdeploy.DeploymentTarget {
	endpoint := octopusdeploy.NewCloudRegionEndpoint()

	if v, ok := d.GetOk("default_worker_pool_id"); ok {
		endpoint.DefaultWorkerPoolID = v.(string)
	}

	deploymentTarget := expandDeploymentTarget(d)
	deploymentTarget.Endpoint = endpoint
	return deploymentTarget
}

func flattenCloudRegionDeploymentTarget(deploymentTarget *octopusdeploy.DeploymentTarget) map[string]interface{} {
	if deploymentTarget == nil {
		return nil
	}

	flattenedDeploymentTarget := flattenDeploymentTarget(deploymentTarget)
	endpointResource, _ := octopusdeploy.ToEndpointResource(deploymentTarget.Endpoint)
	flattenedDeploymentTarget["default_worker_pool_id"] = endpointResource.DefaultWorkerPoolID
	return flattenedDeploymentTarget
}

func getCloudRegionDeploymentTargetDataSchema() map[string]*schema.Schema {
	dataSchema := getCloudRegionDeploymentTargetSchema()
	setDataSchema(&dataSchema)

	deploymentTargetDataSchema := getDeploymentTargetDataSchema()

	deploymentTargetDataSchema["cloud_region_deployment_targets"] = &schema.Schema{
		Computed:    true,
		Description: "A list of cloud region deployment targets that match the filter(s).",
		Elem:        &schema.Resource{Schema: dataSchema},
		Optional:    true,
		Type:        schema.TypeList,
	}

	delete(deploymentTargetDataSchema, "communication_styles")
	delete(deploymentTargetDataSchema, "deployment_targets")
	deploymentTargetDataSchema["id"] = getDataSchemaID()

	return deploymentTargetDataSchema
}

func getCloudRegionDeploymentTargetSchema() map[string]*schema.Schema {
	cloudRegionDeploymentTargetSchema := getDeploymentTargetSchema()

	delete(cloudRegionDeploymentTargetSchema, "endpoint")

	cloudRegionDeploymentTargetSchema["default_worker_pool_id"] = &schema.Schema{
		Optional: true,
		Type:     schema.TypeString,
	}

	return cloudRegionDeploymentTargetSchema
}

func setCloudRegionDeploymentTarget(ctx context.Context, d *schema.ResourceData, deploymentTarget *octopusdeploy.DeploymentTarget) error {
	endpointResource, err := octopusdeploy.ToEndpointResource(deploymentTarget.Endpoint)
	if err != nil {
		return err
	}

	d.Set("default_worker_pool_id", endpointResource.DefaultWorkerPoolID)

	return setDeploymentTarget(ctx, d, deploymentTarget)
}
