package datahub

import (
	"fmt"
	datahubclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client/operations"
	datahubmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strings"
	"time"
)

func ResourceCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceClusterCreate,
		Read:   resourceClusterRead,
		Update: resourceClusterUpdate,
		Delete: resourceClusterDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cloud_platform": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"environment_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cluster_template_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"image_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"image_catalog_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_groups": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"instance_group_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"node_count": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"instance_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"root_volume_size": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"attached_volumes": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"volume_type": {
										Type:     schema.TypeString,
										Required: true,
									},
									"volume_size": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"volume_count": {
										Type:     schema.TypeInt,
										Required: true,
									},
								},
							},
						},
						"recipe_names": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Required: true,
						},
						"recovery_mode": {
							Type:     schema.TypeString,
							Required: true,
						},
						"volume_encryption_enabled": &schema.Schema{
							Type:     schema.TypeBool,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceClusterCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*utils.CdpClients).Datahub

	cloudPlatform := d.Get("cloud_platform").(string)
	switch cloudPlatform {
	case "AWS":
		if err := resourceAWSClusterCreate(d, client); err != nil {
			return err
		}
	case "Azure":
		return fmt.Errorf("azure Not supported yet")
	default:
		return fmt.Errorf("unsupported cloud platform: %s. Must be one of {AWS, AZURE}", "cloud_platform")
	}

	return resourceClusterRead(d, m)
}

func resourceAWSClusterCreate(d *schema.ResourceData, client *datahubclient.Datahub) error {
	name := d.Get("name").(string)
	environmentName := d.Get("environment_name").(string)
	clusterTemplateName := d.Get("cluster_template_name").(string)

	params := operations.NewCreateAWSClusterParams()
	params.WithInput(&datahubmodels.CreateAWSClusterRequest{
		ClusterName:         name,
		EnvironmentName:     environmentName,
		ClusterTemplateName: clusterTemplateName,
	})
	resp, err := client.Operations.CreateAWSCluster(params)
	if err != nil {
		return err
	}

	// We can also use CRN rather than name, but not supported in API
	d.SetId(*resp.GetPayload().Cluster.ClusterName)
	d.Set("crn", *resp.GetPayload().Cluster.Crn)

	// TODO: file a JIRA, this shouldn't be necessary.
	time.Sleep(5 * time.Second)

	if err := waitForClusterToBeAvailable(d.Id(), d.Timeout(schema.TimeoutCreate), client); err != nil {
		return err
	}

	return nil
}

func waitForClusterToBeAvailable(datahubName string, timeout time.Duration, client *datahubclient.Datahub) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{"EXTERNAL_DATABASE_CREATION_IN_PROGRESS"},
		Target:  []string{"AVAILABLE"},
		Timeout: timeout,
		Refresh: func() (interface{}, string, error) {
			params := operations.NewDescribeClusterParams()
			params.WithInput(&datahubmodels.DescribeClusterRequest{ClusterName: &datahubName})
			resp, err := client.Operations.DescribeCluster(params)
			if err != nil {
				return 42, "", err
			}

			return resp, resp.GetPayload().Cluster.Status, nil
		},
	}
	_, err := stateConf.WaitForState()

	return err
}

func resourceClusterRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*utils.CdpClients).Datahub

	name := d.Id()
	params := operations.NewDescribeClusterParams()
	params.WithInput(&datahubmodels.DescribeClusterRequest{ClusterName: &name})
	resp, err := client.Operations.DescribeCluster(params)
	if err != nil {
		if dlErr, ok := err.(*operations.DescribeClusterDefault); ok {
			if matches(dlErr.GetPayload(), "NOT_FOUND", "") {
				d.SetId("") // deleted
				return nil
			}
		}
		return err
	}
	datahub := resp.GetPayload().Cluster
	if datahub == nil {
		d.SetId("") // deleted
		return nil
	}

	d.Set("name", datahub.ClusterName)
	d.Set("crn", datahub.Crn)
	d.SetId(*datahub.ClusterName)

	return nil
}

func resourceClusterUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceClusterRead(d, m)
}

func resourceClusterDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*utils.CdpClients).Datahub

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
		Pending: []string{"EXTERNAL_DATABASE_DELETION_IN_PROGRESS"},
		Target:  []string{},
		Timeout: timeout,
		Refresh: func() (interface{}, string, error) {
			params := operations.NewDescribeClusterParams()
			params.WithInput(&datahubmodels.DescribeClusterRequest{ClusterName: &datahubName})
			resp, err := datahub.Operations.DescribeCluster(params)
			if err != nil {
				if dlErr, ok := err.(*operations.DescribeClusterDefault); ok {
					if matches(dlErr.GetPayload(), "NOT_FOUND", "") {
						return nil, "", nil
					}
				}
				return 42, "", err
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

func toStringList(configured []interface{}) []string {
	vs := make([]string, 0, len(configured))
	for _, v := range configured {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, v.(string))
		}
	}
	return vs
}

func matches(err *datahubmodels.Error, code string, message string) bool {
	return err.Code == code && strings.Contains(err.Message, message)
}
