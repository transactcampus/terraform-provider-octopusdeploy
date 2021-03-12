package octopusdeploy

import (
	"context"
	"time"

	"github.com/transactcampus/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceKubernetesClusterDeploymentTargets() *schema.Resource {
	return &schema.Resource{
		Description: "Provides information about existing Kubernetes cluster deployment targets.",
		ReadContext: dataSourceKubernetesClusterDeploymentTargetsRead,
		Schema:      getKubernetesClusterDeploymentTargetDataSchema(),
	}
}

func dataSourceKubernetesClusterDeploymentTargetsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	query := octopusdeploy.MachinesQuery{
		CommunicationStyles: []string{"Kubernetes"},
		DeploymentID:        d.Get("deployment_id").(string),
		EnvironmentIDs:      expandArray(d.Get("environments").([]interface{})),
		HealthStatuses:      expandArray(d.Get("health_statuses").([]interface{})),
		IDs:                 expandArray(d.Get("ids").([]interface{})),
		IsDisabled:          d.Get("is_disabled").(bool),
		Name:                d.Get("name").(string),
		PartialName:         d.Get("partial_name").(string),
		Roles:               expandArray(d.Get("roles").([]interface{})),
		ShellNames:          expandArray(d.Get("shell_names").([]interface{})),
		Skip:                d.Get("skip").(int),
		Take:                d.Get("take").(int),
		TenantIDs:           expandArray(d.Get("tenants").([]interface{})),
		TenantTags:          expandArray(d.Get("tenant_tags").([]interface{})),
		Thumbprint:          d.Get("thumbprint").(string),
	}

	client := m.(*octopusdeploy.Client)
	deploymentTargets, err := client.Machines.Get(query)
	if err != nil {
		return diag.FromErr(err)
	}

	flattenedKubernetesClusterDeploymentTargets := []interface{}{}
	for _, deploymentTarget := range deploymentTargets.Items {
		flattenedKubernetesClusterDeploymentTargets = append(flattenedKubernetesClusterDeploymentTargets, flattenKubernetesClusterDeploymentTarget(deploymentTarget))
	}

	d.Set("kubernetes_cluster_deployment_targets", flattenedKubernetesClusterDeploymentTargets)
	d.SetId("KubernetesClusterDeploymentTargets " + time.Now().UTC().String())

	return nil
}
