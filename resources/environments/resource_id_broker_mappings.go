package environments

import (
	"fmt"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func ResourceIDBrokerMappings() *schema.Resource {
	return &schema.Resource{
		Create: resourceIDBrokerMappingsCreate,
		Read:   resourceIDBrokerMappingsRead,
		Update: resourceIDBrokerMappingsUpdate,
		Delete: resourceIDBrokerMappingsDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(15 * time.Minute),
			Delete: schema.DefaultTimeout(15 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"environment_crn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"data_access_role": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ranger_audit_role": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mapping": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accessor_crn": {
							Type:     schema.TypeString,
							Required: true,
						},
						"role": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"mappings_version": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceIDBrokerMappingsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Environments

	environmentCrn := d.Get("environment_crn").(string)
	dataAccessRole := d.Get("data_access_role").(string)
	rangerAuditRole := d.Get("ranger_audit_role").(string)

	mappings := []*environmentsmodels.IDBrokerMappingRequest{}
	for _, mappingObject := range d.Get("mapping").([]interface{}) {
		mapping := mappingObject.(map[string]interface{})
		accessorCrn := mapping["accessor_crn"].(string)
		role := mapping["role"].(string)
		mappings = append(mappings, &environmentsmodels.IDBrokerMappingRequest{
			AccessorCrn: &accessorCrn,
			Role:        &role,
		})
	}

	setEmptyMappings := len(mappings) == 0

	mappingParams := operations.NewSetIDBrokerMappingsParams()
	mappingParams.WithInput(&environmentsmodels.SetIDBrokerMappingsRequest{
		EnvironmentName:  &environmentCrn,
		DataAccessRole:   &dataAccessRole,
		RangerAuditRole:  rangerAuditRole,
		Mappings:         mappings,
		SetEmptyMappings: &setEmptyMappings,
	})
	_, mappingErr := client.Operations.SetIDBrokerMappings(mappingParams)
	if mappingErr != nil {
		return mappingErr
	}

	d.SetId(environmentCrn)

	return resourceIDBrokerMappingsRead(d, m)
}

func resourceIDBrokerMappingsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Environments

	environmentCrn := d.Id()

	getMappingsParams := operations.NewGetIDBrokerMappingsParams()
	getMappingsParams.WithInput(&environmentsmodels.GetIDBrokerMappingsRequest{EnvironmentName: &environmentCrn})
	getMappingsResp, err := client.Operations.GetIDBrokerMappings(getMappingsParams)
	if err != nil {
		return err
	}
	idBrokerMappingsResponse := getMappingsResp.GetPayload()

	d.SetId(environmentCrn)
	d.Set("environment_crn", environmentCrn)
	d.Set("data_access_role", idBrokerMappingsResponse.DataAccessRole)
	d.Set("ranger_audit_role", idBrokerMappingsResponse.RangerAuditRole)
	d.Set("mappings_version", idBrokerMappingsResponse.MappingsVersion)

	mappings := make([]map[string]interface{}, 0)
	for _, mapping := range idBrokerMappingsResponse.Mappings {
		item := map[string]interface{}{}
		item["accessor_crn"] = mapping.AccessorCrn
		item["role"] = mapping.Role
		mappings = append(mappings, item)
	}
	if err := d.Set("mapping", mappings); err != nil {
		return fmt.Errorf("error setting mappings: %s", err)
	}

	return nil
}

func resourceIDBrokerMappingsUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceIDBrokerMappingsCreate(d, m)
}

func resourceIDBrokerMappingsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Environments

	environmentCrn := d.Id()

	params := operations.NewDeleteIDBrokerMappingsParams()
	params.WithInput(&environmentsmodels.DeleteIDBrokerMappingsRequest{EnvironmentCrn: &environmentCrn})
	_, err := client.Operations.DeleteIDBrokerMappings(params)
	if err != nil {
		return err
	}

	return nil
}
