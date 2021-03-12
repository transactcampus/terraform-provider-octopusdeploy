package octopusdeploy

import (
	"github.com/transactcampus/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandCloudRegion(flattenedMap map[string]interface{}) *octopusdeploy.CloudRegionEndpoint {
	endpoint := octopusdeploy.NewCloudRegionEndpoint()
	endpoint.ID = flattenedMap["id"].(string)
	endpoint.DefaultWorkerPoolID = flattenedMap["default_worker_pool_id"].(string)

	return endpoint
}

func flattenCloudRegion(endpoint *octopusdeploy.CloudRegionEndpoint) []interface{} {
	if endpoint == nil {
		return nil
	}

	return []interface{}{map[string]interface{}{
		"default_worker_pool_id": endpoint.DefaultWorkerPoolID,
		"id":                     endpoint.GetID(),
	}}
}

func getCloudRegionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"default_worker_pool_id": {
			Optional: true,
			Type:     schema.TypeString,
		},
		"id": getIDSchema(),
	}
}
