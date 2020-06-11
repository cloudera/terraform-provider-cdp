package environments

import (
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceAWSCredentialPrerequisites() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAWSCredentialPrerequisitesRead,

		Schema: map[string]*schema.Schema{
			"account_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAWSCredentialPrerequisitesRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Environments

	cloudPlatform := "AWS"

	params := operations.NewGetCredentialPrerequisitesParams()
	params.WithInput(&environmentsmodels.GetCredentialPrerequisitesRequest{CloudPlatform: &cloudPlatform})
	resp, err := client.Operations.GetCredentialPrerequisites(params)
	if err != nil {
		return err
	}
	prerequisites := resp.GetPayload()
	if prerequisites == nil || prerequisites.Aws == nil {
		d.SetId("") // deleted
		return nil
	}

	d.SetId(*prerequisites.Aws.ExternalID)
	d.Set("account_id", prerequisites.AccountID)
	d.Set("external_id", prerequisites.Aws.ExternalID)

	return nil
}
