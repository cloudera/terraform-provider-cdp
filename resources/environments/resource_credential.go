package environments

import (
	"fmt"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceCredential() *schema.Resource {
	return &schema.Resource{
		Create: resourceCredentialCreate,
		Read:   resourceCredentialRead,
		Update: resourceCredentialUpdate,
		Delete: resourceCredentialDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cloud_platform": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"role_arn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"crn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceCredentialCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Environments

	cloudPlatform := d.Get("cloud_platform").(string)
	switch cloudPlatform {
	case "AWS":
		if err := resourceAWSCredentialCreate(d, client); err != nil {
			return err
		}
	case "Azure":
		return fmt.Errorf("azure Not supported yet")
	default:
		return fmt.Errorf("unsupported cloud platform: %s. Must be one of {AWS, AZURE}", cloudPlatform)
	}

	return resourceCredentialRead(d, m)
}

func resourceAWSCredentialCreate(d *schema.ResourceData, client *environmentsclient.Environments) error {
	name := d.Get("name").(string)
	roleArn := d.Get("role_arn").(string)
	description := d.Get("description").(string)

	params := operations.NewCreateAWSCredentialParams()
	params.WithInput(&environmentsmodels.CreateAWSCredentialRequest{
		CredentialName: &name,
		Description:    description,
		RoleArn:        &roleArn,
	})

	resp, err := client.Operations.CreateAWSCredential(params)
	if err != nil {
		return err
	}

	d.SetId(*resp.GetPayload().Credential.CredentialName)

	return nil
}

func resourceCredentialRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Environments

	name := d.Id()
	params := operations.NewListCredentialsParams()
	params.WithInput(&environmentsmodels.ListCredentialsRequest{CredentialName: name})
	resp, err := client.Operations.ListCredentials(params)
	if err != nil {
		return err
	}
	credentials := resp.GetPayload().Credentials
	if len(credentials) == 0 || *credentials[0].CredentialName != name {
		d.SetId("") // deleted
		return nil
	}

	c := credentials[0]

	d.Set("name", c.CredentialName)
	d.Set("description", c.Description)
	d.Set("crn", c.Crn)
	d.Set("cloud_platform", c.CloudPlatform)
	d.SetId(*c.CredentialName)

	return nil
}

func resourceCredentialUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceCredentialRead(d, m)
}

func resourceCredentialDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Environments

	name := d.Id()
	params := operations.NewDeleteCredentialParams()
	params.WithInput(&environmentsmodels.DeleteCredentialRequest{CredentialName: &name})
	_, err := client.Operations.DeleteCredential(params)
	if err != nil {
		return err
	}

	return nil
}
