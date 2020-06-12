package datalake

import (
	"fmt"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	datalakeclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client/operations"
	datalakemodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"time"
)

func ResourceAWSDatalake() *schema.Resource {
	return &schema.Resource{
		Create: resourceAWSDatalakeCreate,
		Read:   resourceAWSDatalakeRead,
		Delete: resourceAWSDatalakeDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"datalake_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"environment_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cloud_provider_configuration": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"storage_bucket_location": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"instance_profile": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
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

func resourceAWSDatalakeCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Datalake

	datalakeName := d.Get("datalake_name").(string)
	environmentName := d.Get("environment_name").(string)

	storageBucketLocation, instanceProfile, getCloudProviderConfigurationErr := getCloudProviderConfiguration(d)
	if getCloudProviderConfigurationErr != nil {
		return getCloudProviderConfigurationErr
	}

	params := operations.NewCreateAWSDatalakeParams()
	params.WithInput(&datalakemodels.CreateAWSDatalakeRequest{
		DatalakeName:    &datalakeName,
		EnvironmentName: &environmentName,
		CloudProviderConfiguration: &datalakemodels.AWSConfigurationRequest{
			InstanceProfile:       &instanceProfile,
			StorageBucketLocation: &storageBucketLocation,
		},
	})
	resp, err := client.Operations.CreateAWSDatalake(params)
	if err != nil {
		return err
	}

	d.SetId(*resp.GetPayload().Datalake.DatalakeName)

	if err := waitForDatalakeToBeRunning(d.Id(), d.Timeout(schema.TimeoutCreate), client); err != nil {
		return err
	}

	return nil
}

func waitForDatalakeToBeRunning(datalakeName string, timeout time.Duration, client *datalakeclient.Datalake) error {
	stateConf := &resource.StateChangeConf{
		Pending:      []string{"REQUESTED", "EXTERNAL_DATABASE_CREATION_IN_PROGRESS", "EXTERNAL_DATABASE_CREATED", "STACK_CREATION_IN_PROGRESS"},
		Target:       []string{"RUNNING"},
		Delay:        5 * time.Second,
		Timeout:      timeout,
		PollInterval: 10 * time.Second,
		Refresh: func() (interface{}, string, error) {
			log.Printf("About to describe datalake")
			params := operations.NewDescribeDatalakeParams()
			params.WithInput(&datalakemodels.DescribeDatalakeRequest{DatalakeName: &datalakeName})
			resp, err := client.Operations.DescribeDatalake(params)
			if err != nil {
				log.Printf("Error describing datalake: %s", err)
				return nil, "", err
			}
			log.Printf("Described datalake: %s", resp.GetPayload().Datalake.Status)
			return resp, resp.GetPayload().Datalake.Status, nil
		},
	}
	_, err := stateConf.WaitForState()

	return err
}

func resourceAWSDatalakeRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Datalake

	datalakeName := d.Id()
	params := operations.NewDescribeDatalakeParams()
	params.WithInput(&datalakemodels.DescribeDatalakeRequest{DatalakeName: &datalakeName})
	resp, err := client.Operations.DescribeDatalake(params)
	if err != nil {
		if dlErr, ok := err.(*operations.DescribeDatalakeDefault); ok {
			if cdp.IsDatalakeError(dlErr.GetPayload(), "NOT_FOUND", "") {
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

	d.SetId(*datalake.DatalakeName)
	d.Set("datalake_name", datalake.DatalakeName)
	d.Set("crn", datalake.Crn)
	// TODO: file a JIRA, we can't get the environment name from the describe datalake call, and creating via an environment CRN does not seem to work.

	// TODO: file a JIRA, we can't get the storage location from the describe datalake call.
	storageBucketLocation, _, getCloudProviderConfigurationErr := getCloudProviderConfiguration(d)
	if getCloudProviderConfigurationErr != nil {
		return getCloudProviderConfigurationErr
	}

	cloudProviderConfiguration := []interface{}{
		map[string]interface{}{
			"storage_bucket_location": storageBucketLocation,
			"instance_profile":        datalake.AwsConfiguration.InstanceProfile,
		},
	}
	if err := d.Set("cloud_provider_configuration", cloudProviderConfiguration); err != nil {
		return fmt.Errorf("error setting authentication: %s", err)
	}

	return nil
}

func resourceAWSDatalakeDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Datalake

	datalakeName := d.Id()
	params := operations.NewDeleteDatalakeParams()
	params.WithInput(&datalakemodels.DeleteDatalakeRequest{DatalakeName: &datalakeName})
	_, err := client.Operations.DeleteDatalake(params)
	if err != nil {
		return err
	}

	if err := waitForDatalakeToBeDeleted(d.Id(), d.Timeout(schema.TimeoutDelete), client); err != nil {
		return err
	}

	return nil
}

func waitForDatalakeToBeDeleted(datalakeName string, timeout time.Duration, client *datalakeclient.Datalake) error {
	stateConf := &resource.StateChangeConf{
		Pending:      []string{"EXTERNAL_DATABASE_DELETION_IN_PROGRESS", "STACK_DELETION_IN_PROGRESS", "STACK_DELETED"},
		Target:       []string{},
		Delay:        5 * time.Second,
		Timeout:      timeout,
		PollInterval: 10 * time.Second,
		Refresh: func() (interface{}, string, error) {
			log.Printf("About to describe datalake")
			params := operations.NewDescribeDatalakeParams()
			params.WithInput(&datalakemodels.DescribeDatalakeRequest{DatalakeName: &datalakeName})
			resp, err := client.Operations.DescribeDatalake(params)
			if err != nil {
				log.Printf("Error describing datalake: %s", err)
				if dlErr, ok := err.(*operations.DescribeDatalakeDefault); ok {
					if cdp.IsDatalakeError(dlErr.GetPayload(), "NOT_FOUND", "") {
						return nil, "", nil
					}
				}
				return nil, "", err
			}
			if resp.GetPayload().Datalake == nil {
				log.Printf("Described datalake. No datalake.")
				return nil, "", nil
			}
			log.Printf("Described datalake: %s", resp.GetPayload().Datalake.Status)
			return resp, resp.GetPayload().Datalake.Status, nil
		},
	}
	_, err := stateConf.WaitForState()

	return err
}

func getCloudProviderConfiguration(d *schema.ResourceData) (string, string, error) {
	cloudProviderConfiguration, cloudProviderConfigurationErr := utils.GetMapFromSingleItemList(d, "cloud_provider_configuration")
	if cloudProviderConfigurationErr != nil {
		return "", "", cloudProviderConfigurationErr
	}
	return cloudProviderConfiguration["storage_bucket_location"].(string), cloudProviderConfiguration["instance_profile"].(string), nil
}
