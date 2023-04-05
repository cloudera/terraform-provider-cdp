package environments

import (
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceAzureCredential() *schema.Resource {
	return &schema.Resource{
		Create: resourceAzureCredentialCreate,
		Read:   resourceAzureCredentialRead,
		Delete: resourceAzureCredentialDelete,

		Schema: map[string]*schema.Schema{
			"credential_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subscription_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"app_based": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"application_id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"secret_key": &schema.Schema{
							Type:      schema.TypeString,
							Required:  true,
							ForceNew:  true,
							Sensitive: true,
						},
					},
				},
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

func resourceAzureCredentialCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Environments

	credentialName := d.Get("credential_name").(string)
	subscriptionId := d.Get("subscription_id").(string)
	tenantId := d.Get("tenant_id").(string)
	description := d.Get("description").(string)

	appBasedMap, err := utils.GetMapFromSingleItemList(d, "app_based")
	if err != nil {
		return err
	}
	applicationId := appBasedMap["application_id"].(string)
	secretKey := appBasedMap["secret_key"].(string)

	appBased := &environmentsmodels.CreateAzureCredentialRequestAppBased{
		ApplicationID: applicationId,
		SecretKey:     secretKey,
	}

	params := operations.NewCreateAzureCredentialParams()
	params.WithInput(&environmentsmodels.CreateAzureCredentialRequest{
		CredentialName: &credentialName,
		Description:    description,
		SubscriptionID: subscriptionId,
		TenantID:       tenantId,
		AppBased:       appBased,
	})
	_, err = client.Operations.CreateAzureCredential(params)
	if err != nil {
		return err
	}

	d.SetId(credentialName)

	return resourceAzureCredentialRead(d, m)
}

func resourceAzureCredentialRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Environments

	credentialName := d.Id()
	params := operations.NewListCredentialsParams()
	params.WithInput(&environmentsmodels.ListCredentialsRequest{CredentialName: credentialName})
	resp, err := client.Operations.ListCredentials(params)
	if err != nil {
		return err
	}
	credentials := resp.GetPayload().Credentials
	if len(credentials) == 0 || *credentials[0].CredentialName != credentialName {
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

func resourceAzureCredentialDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Environments

	credentialName := d.Id()
	params := operations.NewDeleteCredentialParams()
	params.WithInput(&environmentsmodels.DeleteCredentialRequest{CredentialName: &credentialName})
	_, err := client.Operations.DeleteCredential(params)
	if err != nil {
		return err
	}

	return nil
}
