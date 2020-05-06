package environments

import (
	"fmt"
	environmentclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	nameField          = "name"
	roleArnField       = "role_arn"
	descriptionField   = "description"
	crnField           = "crn"
	cloudPlatformField = "cloud_platform"
)

func ResourceCredential() *schema.Resource {
	return &schema.Resource{
		Create: resourceCredentialCreate,
		Read:   resourceCredentialRead,
		Update: resourceCredentialUpdate,
		Delete: resourceCredentialDelete,

		Schema: map[string]*schema.Schema{
			nameField: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			cloudPlatformField: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			roleArnField: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			descriptionField: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			crnField: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceCredentialCreate(d *schema.ResourceData, m interface{}) error {
	environments := m.(*environmentclient.Environments)

	cloudPlatform := d.Get(cloudPlatformField).(string)
	switch cloudPlatform {
	case "AWS":
		if err := resourceAWSCredentialCreate(d, environments); err != nil {
			return err
		}
	case "Azure":
		return fmt.Errorf("azure Not supported yet")
	default:
		return fmt.Errorf("unsupported cloud platform: %s. Must be one of {AWS, AZURE}", cloudPlatform)
	}

	return resourceCredentialRead(d, m)
}

func resourceAWSCredentialCreate(d *schema.ResourceData, environments *environmentclient.Environments) error {
	params := operations.NewCreateAWSCredentialParams()
	name := d.Get(nameField).(string)
	roleArn, ok := d.GetOk(roleArnField)
	if !ok {
		return fmt.Errorf("%s field is required", roleArnField)
	}
	roleArnStr := roleArn.(string)

	description := d.Get(descriptionField).(string)

	params.WithInput(&environmentmodels.CreateAWSCredentialRequest{
		CredentialName: &name,
		Description:    description,
		RoleArn:        &roleArnStr,
	})
	resp, err := environments.Operations.CreateAWSCredential(params)
	if err != nil {
		return err
	}

	// We can also use CRN rather than name, but not supported in API
	d.SetId(*resp.GetPayload().Credential.CredentialName)
	return nil
}

func resourceCredentialRead(d *schema.ResourceData, m interface{}) error {
	environments := m.(*environmentclient.Environments)

	name := d.Id()
	params := operations.NewListCredentialsParams()
	params.WithInput(&environmentmodels.ListCredentialsRequest{CredentialName: name})
	resp, err := environments.Operations.ListCredentials(params)
	if err != nil {
		return err
	}
	credentials := resp.GetPayload().Credentials
	if len(credentials) == 0 || *credentials[0].CredentialName != name {
		d.SetId("") // deleted
		return nil
	}

	c := credentials[0]

	d.Set(nameField, c.CredentialName)
	d.Set(descriptionField, c.Description)
	d.Set(crnField, c.Crn)
	d.Set(cloudPlatformField, c.CloudPlatform)
	d.SetId(*c.CredentialName)

	return nil
}

func resourceCredentialUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceCredentialRead(d, m)
}

func resourceCredentialDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
