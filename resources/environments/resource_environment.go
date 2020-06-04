package environments

import (
	"fmt"
	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"sort"
	"strings"
	"time"
)

// TODO: set and update telemetry features
func ResourceEnvironment() *schema.Resource {
	return &schema.Resource{
		Create: resourceEnvironmentCreate,
		Read:   resourceEnvironmentRead,
		Update: resourceEnvironmentUpdate,
		Delete: resourceEnvironmentDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(15 * time.Minute),
			Delete: schema.DefaultTimeout(15 * time.Minute),
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
			"credential_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"knox_security_group_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"default_security_group_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"public_key_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vpc_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet_ids": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"log_storage_location_base": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"log_storage_instance_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"s3_guard_table_name": &schema.Schema{
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
			"tunnel_enabled": &schema.Schema{
				Type:     schema.TypeBool,
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

func resourceEnvironmentCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*utils.CdpClients).Environments

	cloudPlatform := d.Get("cloud_platform").(string)
	switch cloudPlatform {
	case "AWS":
		if err := resourceAWSEnvironmentCreate(d, client); err != nil {
			return err
		}
	case "Azure":
		return fmt.Errorf("azure Not supported yet")
	default:
		return fmt.Errorf("unsupported cloud platform: %s. Must be one of {AWS, AZURE}", "cloud_platform")
	}

	return resourceEnvironmentRead(d, m)
}

func resourceAWSEnvironmentCreate(d *schema.ResourceData, client *environmentsclient.Environments) error {
	name := d.Get("name").(string)
	credentialName := d.Get("credential_name").(string)
	region := d.Get("region").(string)
	description := d.Get("description").(string)
	knoxSecurityGroupId := d.Get("knox_security_group_id").(string)
	defaultSecurityGroupId := d.Get("default_security_group_id").(string)
	publicKeyId := d.Get("public_key_id").(string)
	vpcId := d.Get("vpc_id").(string)
	subnetIds := toStringList(d.Get("subnet_ids").([]interface{}))
	logStorageLocationBase := d.Get("log_storage_location_base").(string)
	logStorageInstanceProfile := d.Get("log_storage_instance_profile").(string)
	s3GuardTableName := d.Get("s3_guard_table_name").(string)
	dataAccessRole := d.Get("data_access_role").(string)
	rangerAuditRole := d.Get("ranger_audit_role").(string)
	tunnelEnabled := d.Get("tunnel_enabled").(bool)

	createParams := operations.NewCreateAWSEnvironmentParams()
	createParams.WithInput(&environmentsmodels.CreateAWSEnvironmentRequest{
		EnvironmentName: &name,
		CredentialName:  &credentialName,
		Region:          &region,
		Description:     description,
		Authentication: &environmentsmodels.AuthenticationRequest{
			PublicKeyID: publicKeyId,
		},
		SecurityAccess: &environmentsmodels.SecurityAccessRequest{
			DefaultSecurityGroupID: defaultSecurityGroupId,
			SecurityGroupIDForKnox: knoxSecurityGroupId,
		},
		LogStorage: &environmentsmodels.AwsLogStorageRequest{
			InstanceProfile:     &logStorageInstanceProfile,
			StorageLocationBase: &logStorageLocationBase,
		},
		VpcID:            vpcId,
		SubnetIds:        subnetIds,
		S3GuardTableName: s3GuardTableName,
		EnableTunnel:     tunnelEnabled,
	})
	_, createErr := client.Operations.CreateAWSEnvironment(createParams)
	if createErr != nil {
		return createErr
	}

	// TODO: file a JIRA, this shouldn't be necessary.
	time.Sleep(5 * time.Second)

	if err := waitForEnvironmentToBeAvailable(d.Id(), d.Timeout(schema.TimeoutCreate), client); err != nil {
		return err
	}

	setEmptyMappings := true
	mappingParams := operations.NewSetIDBrokerMappingsParams()
	mappingParams.WithInput(&environmentsmodels.SetIDBrokerMappingsRequest{
		EnvironmentName:  &name,
		DataAccessRole:   &dataAccessRole,
		RangerAuditRole:  rangerAuditRole,
		Mappings:         make([]*environmentsmodels.IDBrokerMappingRequest, 0, 0),
		SetEmptyMappings: &setEmptyMappings,
	})
	_, mappingErr := client.Operations.SetIDBrokerMappings(mappingParams)
	if mappingErr != nil {
		return mappingErr
	}

	return nil
}

func waitForEnvironmentToBeAvailable(environmentName string, timeout time.Duration, client *environmentsclient.Environments) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{"FREEIPA_CREATION_IN_PROGRESS"},
		Target:  []string{"AVAILABLE"},
		Timeout: timeout,
		Refresh: func() (interface{}, string, error) {
			params := operations.NewDescribeEnvironmentParams()
			params.WithInput(&environmentsmodels.DescribeEnvironmentRequest{EnvironmentName: &environmentName})
			resp, err := client.Operations.DescribeEnvironment(params)
			if err != nil {
				return 42, "", err
			}

			return resp, *resp.GetPayload().Environment.Status, nil
		},
	}
	_, err := stateConf.WaitForState()

	return err
}

func resourceEnvironmentRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*utils.CdpClients).Environments

	name := d.Id()

	describeParams := operations.NewDescribeEnvironmentParams()
	describeParams.WithInput(&environmentsmodels.DescribeEnvironmentRequest{EnvironmentName: &name})
	describeEnvironmentResp, err := client.Operations.DescribeEnvironment(describeParams)
	if err != nil {
		if envErr, ok := err.(*operations.DescribeEnvironmentDefault); ok {
			if matches(envErr.GetPayload(), "NOT_FOUND", "") {
				d.SetId("") // deleted
				return nil
			}
		}
		return err
	}
	environment := describeEnvironmentResp.GetPayload().Environment
	if environment == nil {
		d.SetId("") // deleted
		return nil
	}

	getMappingsParams := operations.NewGetIDBrokerMappingsParams()
	getMappingsParams.WithInput(&environmentsmodels.GetIDBrokerMappingsRequest{EnvironmentName: &name})
	getMappingsResp, err := client.Operations.GetIDBrokerMappings(getMappingsParams)
	if err != nil {
		return err
	}
	idBrokerMappingsResponse := getMappingsResp.GetPayload()

	d.SetId(*environment.EnvironmentName)
	d.Set("name", *environment.EnvironmentName)
	d.Set("cloud_platform", *environment.CloudPlatform)
	d.Set("credential_name", *environment.CredentialName)
	d.Set("region", *environment.Region)
	d.Set("knox_security_group_id", environment.SecurityAccess.SecurityGroupIDForKnox)
	d.Set("default_security_group_id", environment.SecurityAccess.DefaultSecurityGroupID)
	d.Set("public_key_id", environment.Authentication.PublicKeyID)
	d.Set("vpc_id", environment.Network.Aws.VpcID)
	sort.Strings(environment.Network.SubnetIds)
	d.Set("subnet_ids", environment.Network.SubnetIds)
	d.Set("log_storage_location_base", environment.LogStorage.AwsDetails.StorageLocationBase)
	d.Set("log_storage_instance_profile", environment.LogStorage.AwsDetails.InstanceProfile)
	d.Set("s3_guard_table_name", environment.AwsDetails.S3GuardTableName)
	d.Set("data_access_role", *idBrokerMappingsResponse.DataAccessRole)
	d.Set("ranger_audit_role", *idBrokerMappingsResponse.RangerAuditRole)
	d.Set("tunnel_enabled", environment.TunnelEnabled)
	d.Set("description", environment.Description)
	d.Set("crn", environment.Crn)

	return nil
}

func resourceEnvironmentUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*utils.CdpClients).Environments

	name := d.Id()

	if d.HasChange("credential_name") {
		credentialName := d.Get("credential_name").(string)

		params := operations.NewChangeEnvironmentCredentialParams()
		params.WithInput(&environmentsmodels.ChangeEnvironmentCredentialRequest{
			EnvironmentName: &name,
			CredentialName:  &credentialName,
		})
		_, err := client.Operations.ChangeEnvironmentCredential(params)
		if err != nil {
			return err
		}
	}

	return resourceEnvironmentRead(d, m)
}

func resourceEnvironmentDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*utils.CdpClients).Environments

	name := d.Id()
	params := operations.NewDeleteEnvironmentParams()
	params.WithInput(&environmentsmodels.DeleteEnvironmentRequest{EnvironmentName: &name})
	_, err := client.Operations.DeleteEnvironment(params)
	if err != nil {
		return err
	}

	// TODO: file a JIRA, this shouldn't be necessary.
	time.Sleep(5 * time.Second)

	if err := waitForEnvironmentToBeDeleted(d.Id(), d.Timeout(schema.TimeoutDelete), client); err != nil {
		return err
	}

	// TODO: file a JIRA, this shouldn't be necessary.
	time.Sleep(5 * time.Second)

	return nil
}

func waitForEnvironmentToBeDeleted(environmentName string, timeout time.Duration, client *environmentsclient.Environments) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{"FREEIPA_DELETE_IN_PROGRESS"},
		Target:  []string{},
		Timeout: timeout,
		Refresh: func() (interface{}, string, error) {
			params := operations.NewDescribeEnvironmentParams()
			params.WithInput(&environmentsmodels.DescribeEnvironmentRequest{EnvironmentName: &environmentName})
			resp, err := client.Operations.DescribeEnvironment(params)
			if err != nil {
				if envErr, ok := err.(*operations.DescribeEnvironmentDefault); ok {
					if matches(envErr.GetPayload(), "NOT_FOUND", "") {
						return nil, "", nil
					}
				}
				return 42, "", err
			}
			if resp.GetPayload().Environment == nil {
				return nil, "", nil
			}
			return resp, *resp.GetPayload().Environment.Status, nil
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

func matches(err *environmentsmodels.Error, code string, message string) bool {
	return err.Code == code && strings.Contains(err.Message, message)
}
