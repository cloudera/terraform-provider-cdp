package environments

import (
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceAWSCredential() *schema.Resource {
	return &schema.Resource{
		Create: resourceAWSCredentialCreate,
		Read:   resourceAWSCredentialRead,
		Delete: resourceAWSCredentialDelete,

		Schema: map[string]*schema.Schema{
			"credential_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"role_arn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"crn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAWSCredentialCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Environments

	credentialName := d.Get("credential_name").(string)
	roleArn := d.Get("role_arn").(string)
	description := d.Get("description").(string)

	params := operations.NewCreateAWSCredentialParams()
	params.WithInput(&environmentsmodels.CreateAWSCredentialRequest{
		CredentialName: &credentialName,
		Description:    description,
		RoleArn:        &roleArn,
	})
	_, err := client.Operations.CreateAWSCredential(params)
	if err != nil {
		return err
	}

	return resourceAWSCredentialRead(d, m)
}

func resourceAWSCredentialRead(d *schema.ResourceData, m interface{}) error {
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

	d.SetId(*c.CredentialName)
	d.Set("credential_name", c.CredentialName)
	d.Set("description", c.Description)
	d.Set("crn", c.Crn)

	return nil
}

func resourceAWSCredentialDelete(d *schema.ResourceData, m interface{}) error {
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
