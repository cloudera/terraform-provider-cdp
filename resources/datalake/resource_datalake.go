package datalake

import (
	"fmt"
	datalakeclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client/operations"
	datalakemodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strings"
	"time"
)

func ResourceDatalake() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatalakeCreate,
		Read:   resourceDatalakeRead,
		Update: resourceDatalakeUpdate,
		Delete: resourceDatalakeDelete,

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
			"data_storage_base": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"id_broker_instance_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceDatalakeCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*utils.CdpClients).Datalake

	cloudPlatform := d.Get("cloud_platform").(string)
	switch cloudPlatform {
	case "AWS":
		if err := resourceAWSDatalakeCreate(d, client); err != nil {
			return err
		}
	case "Azure":
		return fmt.Errorf("azure Not supported yet")
	default:
		return fmt.Errorf("unsupported cloud platform: %s. Must be one of {AWS, AZURE}", "cloud_platform")
	}

	return resourceDatalakeRead(d, m)
}

func resourceAWSDatalakeCreate(d *schema.ResourceData, client *datalakeclient.Datalake) error {
	name := d.Get("name").(string)
	environmentName := d.Get("environment_name").(string)
	dataStorageBase := d.Get("data_storage_base").(string)
	idBrokerInstanceProfile := d.Get("id_broker_instance_profile").(string)

	params := operations.NewCreateAWSDatalakeParams()
	params.WithInput(&datalakemodels.CreateAWSDatalakeRequest{
		DatalakeName:    &name,
		EnvironmentName: &environmentName,
		CloudProviderConfiguration: &datalakemodels.AWSConfigurationRequest{
			InstanceProfile:       &idBrokerInstanceProfile,
			StorageBucketLocation: &dataStorageBase,
		},
	})
	resp, err := client.Operations.CreateAWSDatalake(params)
	if err != nil {
		return err
	}

	// We can also use CRN rather than name, but not supported in API
	d.SetId(*resp.GetPayload().Datalake.DatalakeName)
	d.Set("crn", *resp.GetPayload().Datalake.Crn)

	// TODO: file a JIRA, this shouldn't be necessary.
	time.Sleep(5 * time.Second)

	if err := waitForDatalakeToBeRunning(d.Id(), d.Timeout(schema.TimeoutCreate), client); err != nil {
		return err
	}

	return nil
}

func waitForDatalakeToBeRunning(datalakeName string, timeout time.Duration, client *datalakeclient.Datalake) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{"EXTERNAL_DATABASE_CREATION_IN_PROGRESS"},
		Target:  []string{"RUNNING"},
		Timeout: timeout,
		Refresh: func() (interface{}, string, error) {
			params := operations.NewDescribeDatalakeParams()
			params.WithInput(&datalakemodels.DescribeDatalakeRequest{DatalakeName: &datalakeName})
			resp, err := client.Operations.DescribeDatalake(params)
			if err != nil {
				return 42, "", err
			}

			return resp, resp.GetPayload().Datalake.Status, nil
		},
	}
	_, err := stateConf.WaitForState()

	return err
}

func resourceDatalakeRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*utils.CdpClients).Datalake

	name := d.Id()
	params := operations.NewDescribeDatalakeParams()
	params.WithInput(&datalakemodels.DescribeDatalakeRequest{DatalakeName: &name})
	resp, err := client.Operations.DescribeDatalake(params)
	if err != nil {
		if dlErr, ok := err.(*operations.DescribeDatalakeDefault); ok {
			if matches(dlErr.GetPayload(), "NOT_FOUND", "") {
				d.SetId("") // deleted
				return nil
			}
		}
		return err
	}
	datalake := resp.GetPayload().Datalake
	if datalake == nil {
		d.SetId("") // deleted
		return nil
	}

	d.Set("name", datalake.DatalakeName)
	d.Set("crn", datalake.Crn)
	d.SetId(*datalake.DatalakeName)

	return nil
}

func resourceDatalakeUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceDatalakeRead(d, m)
}

func resourceDatalakeDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*utils.CdpClients).Datalake

	name := d.Id()
	params := operations.NewDeleteDatalakeParams()
	params.WithInput(&datalakemodels.DeleteDatalakeRequest{DatalakeName: &name})
	_, err := client.Operations.DeleteDatalake(params)
	if err != nil {
		return err
	}

	// TODO: file a JIRA, this shouldn't be necessary.
	time.Sleep(5 * time.Second)

	if err := waitForDatalakeToBeDeleted(d.Id(), d.Timeout(schema.TimeoutDelete), client); err != nil {
		return err
	}

	return nil
}

func waitForDatalakeToBeDeleted(datalakeName string, timeout time.Duration, datalake *datalakeclient.Datalake) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{"EXTERNAL_DATABASE_DELETION_IN_PROGRESS"},
		Target:  []string{},
		Timeout: timeout,
		Refresh: func() (interface{}, string, error) {
			params := operations.NewDescribeDatalakeParams()
			params.WithInput(&datalakemodels.DescribeDatalakeRequest{DatalakeName: &datalakeName})
			resp, err := datalake.Operations.DescribeDatalake(params)
			if err != nil {
				if dlErr, ok := err.(*operations.DescribeDatalakeDefault); ok {
					if matches(dlErr.GetPayload(), "NOT_FOUND", "") {
						return nil, "", nil
					}
				}
				return 42, "", err
			}
			if resp.GetPayload().Datalake == nil {
				return nil, "", nil
			}
			return resp, resp.GetPayload().Datalake.Status, nil
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

func matches(err *datalakemodels.Error, code string, message string) bool {
	return err.Code == code && strings.Contains(err.Message, message)
}
