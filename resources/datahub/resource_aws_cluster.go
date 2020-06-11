package datahub

import (
	"fmt"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	datahubclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client/operations"
	datahubmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func ResourceAWSCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceAWSClusterCreate,
		Read:   resourceAWSClusterRead,
		Delete: resourceAWSClusterDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"cluster_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"environment_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cluster_template_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"image": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"catalog_name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
			"instance_group": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_group_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"instance_group_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"instance_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"node_count": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},
						"root_volume_size": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},
						"attached_volumes": {
							Type:     schema.TypeSet,
							Required: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"volume_type": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"volume_size": {
										Type:     schema.TypeInt,
										Required: true,
										ForceNew: true,
									},
									"volume_count": {
										Type:     schema.TypeInt,
										Required: true,
										ForceNew: true,
									},
								},
							},
						},
						"recipe_names": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Required: true,
							ForceNew: true,
						},
						"recovery_mode": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"volume_encryption": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_encryption": &schema.Schema{
										Type:     schema.TypeBool,
										Required: true,
										ForceNew: true,
									},
									"encryption_key": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
								},
							},
						},
					},
				},
			},
			"crn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAWSClusterCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Datahub

	clusterName := d.Get("cluster_name").(string)
	environmentName := d.Get("environment_name").(string)
	clusterTemplateName := d.Get("cluster_template_name").(string)

	image, imageErr := utils.GetMapFromSingleItemList(d, "image")
	if imageErr != nil {
		return imageErr
	}
	catalogName := image["catalog_name"].(string)
	id := image["id"].(string)

	instanceGroups := []*datahubmodels.InstanceGroupRequest{}
	instanceGroupSet := d.Get("instance_group").(*schema.Set)
	for _, instanceGroupObject := range instanceGroupSet.List() {
		instanceGroup := instanceGroupObject.(map[string]interface{})

		instanceGroupName := instanceGroup["instance_group_name"].(string)
		instanceGroupType := instanceGroup["instance_group_type"].(string)
		instanceType := instanceGroup["instance_type"].(string)
		nodeCount := int32(instanceGroup["node_count"].(int))
		rootVolumeSize := int32(instanceGroup["root_volume_size"].(int))
		recoveryMode := instanceGroup["recovery_mode"].(string)
		recipeNames := utils.ToStringList(instanceGroup["recipe_names"].([]interface{}))

		attachedVolumes := []*datahubmodels.AttachedVolumeRequest{}
		attachedVolumeSet := instanceGroup["attached_volumes"].(*schema.Set)
		for _, attachedVolumesObject := range attachedVolumeSet.List() {
			attachedVolume := attachedVolumesObject.(map[string]interface{})
			volumeSize := int32(attachedVolume["volume_size"].(int))
			volumeType := attachedVolume["volume_type"].(string)
			volumeCount := int32(attachedVolume["volume_count"].(int))
			attachedVolumes = append(attachedVolumes, &datahubmodels.AttachedVolumeRequest{
				VolumeSize:  &volumeSize,
				VolumeType:  &volumeType,
				VolumeCount: &volumeCount,
			})
		}

		volumeEncryption, volumeEncryptionErr := utils.GetMapFromSingleItemListInMap(instanceGroup, "volume_encryption")
		if volumeEncryptionErr != nil {
			return volumeEncryptionErr
		}

		enableEncryption := volumeEncryption["enable_encryption"].(bool)
		encryptionKey := volumeEncryption["encryption_key"].(string)

		instanceGroups = append(instanceGroups, &datahubmodels.InstanceGroupRequest{
			InstanceGroupName:           &instanceGroupName,
			InstanceGroupType:           &instanceGroupType,
			InstanceType:                &instanceType,
			NodeCount:                   &nodeCount,
			RootVolumeSize:              &rootVolumeSize,
			RecoveryMode:                recoveryMode,
			RecipeNames:                 recipeNames,
			AttachedVolumeConfiguration: attachedVolumes,
			VolumeEncryption: &datahubmodels.VolumeEncryptionRequest{
				EnableEncryption: &enableEncryption,
				EncryptionKey:    encryptionKey,
			},
		})
	}

	params := operations.NewCreateAWSClusterParams()
	params.WithInput(&datahubmodels.CreateAWSClusterRequest{
		ClusterName:         clusterName,
		EnvironmentName:     environmentName,
		ClusterTemplateName: clusterTemplateName,
		Image: &datahubmodels.ImageRequest{
			CatalogName: &catalogName,
			ID:          &id,
		},
		InstanceGroups: instanceGroups,
	})
	resp, err := client.Operations.CreateAWSCluster(params)
	if err != nil {
		return err
	}
	cluster := resp.GetPayload().Cluster
	if cluster == nil {
		d.SetId("") // deleted
		return nil
	}

	d.SetId(*cluster.ClusterName)
	d.Set("cluster_name", *cluster.ClusterName)
	d.Set("crn", *cluster.Crn)

	// TODO: file a JIRA, this shouldn't be necessary.
	time.Sleep(5 * time.Second)

	if err := waitForClusterToBeAvailable(d.Id(), d.Timeout(schema.TimeoutCreate), client); err != nil {
		return err
	}

	return resourceAWSClusterRead(d, m)
}

func waitForClusterToBeAvailable(datahubName string, timeout time.Duration, client *datahubclient.Datahub) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{"REQUESTED", "UPDATE_IN_PROGRESS"},
		Target:  []string{"AVAILABLE"},
		Timeout: timeout,
		Refresh: func() (interface{}, string, error) {
			params := operations.NewDescribeClusterParams()
			params.WithInput(&datahubmodels.DescribeClusterRequest{ClusterName: &datahubName})
			resp, err := client.Operations.DescribeCluster(params)
			if err != nil {
				return nil, "", err
			}

			return resp, resp.GetPayload().Cluster.Status, nil
		},
	}
	_, err := stateConf.WaitForState()

	return err
}

func resourceAWSClusterRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Datahub

	clusterName := d.Id()
	params := operations.NewDescribeClusterParams()
	params.WithInput(&datahubmodels.DescribeClusterRequest{ClusterName: &clusterName})
	resp, err := client.Operations.DescribeCluster(params)
	if err != nil {
		if dlErr, ok := err.(*operations.DescribeClusterDefault); ok {
			if cdp.IsDatahubError(dlErr.GetPayload(), "NOT_FOUND", "") {
				d.SetId("") // deleted
				return nil
			}
		}
		return err
	}
	cluster := resp.GetPayload().Cluster
	if cluster == nil {
		d.SetId("") // deleted
		return nil
	}

	d.SetId(*cluster.ClusterName)
	d.Set("cluster_name", *cluster.ClusterName)

	image := []interface{}{
		map[string]interface{}{
			"catalog_name": cluster.ImageDetails.CatalogName,
			"id":           cluster.ImageDetails.ID,
		},
	}
	if err := d.Set("image", image); err != nil {
		return fmt.Errorf("error setting image: %s", err)
	}

	d.Set("crn", *cluster.Crn)

	return nil
}

func resourceAWSClusterDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Datahub

	name := d.Id()
	params := operations.NewDeleteClusterParams()
	params.WithInput(&datahubmodels.DeleteClusterRequest{ClusterName: &name})
	_, err := client.Operations.DeleteCluster(params)
	if err != nil {
		return err
	}

	// TODO: file a JIRA, this shouldn't be necessary.
	time.Sleep(5 * time.Second)

	if err := waitForClusterToBeDeleted(d.Id(), d.Timeout(schema.TimeoutDelete), client); err != nil {
		return err
	}

	return nil
}

func waitForClusterToBeDeleted(datahubName string, timeout time.Duration, datahub *datahubclient.Datahub) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{"DELETE_IN_PROGRESS"},
		Target:  []string{},
		Timeout: timeout,
		Refresh: func() (interface{}, string, error) {
			params := operations.NewDescribeClusterParams()
			params.WithInput(&datahubmodels.DescribeClusterRequest{ClusterName: &datahubName})
			resp, err := datahub.Operations.DescribeCluster(params)
			if err != nil {
				if dlErr, ok := err.(*operations.DescribeClusterDefault); ok {
					if cdp.IsDatahubError(dlErr.GetPayload(), "NOT_FOUND", "") {
						return nil, "", nil
					}
				}
				return nil, "", err
			}
			if resp.GetPayload().Cluster == nil {
				return nil, "", nil
			}
			return resp, resp.GetPayload().Cluster.Status, nil
		},
	}
	_, err := stateConf.WaitForState()

	return err
}
