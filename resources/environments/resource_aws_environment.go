package environments

import (
	"fmt"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"sort"
	"strings"
	"time"
)

func ResourceAWSEnvironment() *schema.Resource {
	return &schema.Resource{
		Create: resourceAWSEnvironmentCreate,
		Read:   resourceAWSEnvironmentRead,
		Update: resourceAWSEnvironmentUpdate,
		Delete: resourceAWSEnvironmentDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"environment_name": &schema.Schema{
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
			"security_access": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_security_group_id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"security_group_id_for_knox": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"authentication": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"public_key_id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
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
			"log_storage": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"storage_location_base": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"instance_profile": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"s3_guard_table_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"enable_tunnel": &schema.Schema{
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

func resourceAWSEnvironmentCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Environments

	environmentName := d.Get("environment_name").(string)
	credentialName := d.Get("credential_name").(string)
	region := d.Get("region").(string)

	authentication, authenticationErr := utils.GetMapFromSingleItemList(d, "authentication")
	if authenticationErr != nil {
		return authenticationErr
	}
	publicKeyId := authentication["public_key_id"].(string)

	securityAccess, securityAccessErr := utils.GetMapFromSingleItemList(d, "security_access")
	if securityAccessErr != nil {
		return securityAccessErr
	}
	defaultSecurityGroupId := securityAccess["default_security_group_id"].(string)
	securityGroupIdForKnox := securityAccess["security_group_id_for_knox"].(string)

	logStorage, logStorageErr := utils.GetMapFromSingleItemList(d, "log_storage")
	if logStorageErr != nil {
		return logStorageErr
	}
	storageLocationBase := logStorage["storage_location_base"].(string)
	instanceProfile := logStorage["instance_profile"].(string)

	params := operations.NewCreateAWSEnvironmentParams()
	params.WithInput(&environmentsmodels.CreateAWSEnvironmentRequest{
		EnvironmentName: &environmentName,
		CredentialName:  &credentialName,
		Region:          &region,
		Description:     d.Get("description").(string),
		Authentication: &environmentsmodels.AuthenticationRequest{
			PublicKeyID: publicKeyId,
		},
		SecurityAccess: &environmentsmodels.SecurityAccessRequest{
			DefaultSecurityGroupID: defaultSecurityGroupId,
			SecurityGroupIDForKnox: securityGroupIdForKnox,
		},
		LogStorage: &environmentsmodels.AwsLogStorageRequest{
			InstanceProfile:     &instanceProfile,
			StorageLocationBase: &storageLocationBase,
		},
		VpcID:            d.Get("vpc_id").(string),
		SubnetIds:        utils.ToStringList(d.Get("subnet_ids").([]interface{})),
		S3GuardTableName: d.Get("s3_guard_table_name").(string),
		EnableTunnel:     d.Get("enable_tunnel").(bool),
	})

	var resp *operations.CreateAWSEnvironmentOK
	err := resource.Retry(30*time.Second, func() *resource.RetryError {
		var err error
		resp, err = client.Operations.CreateAWSEnvironment(params)

		// There is an eventual consistency issue when the env is deleted and re-created. We check for this condition and
		// retry for a short time if it is the case.
		if err != nil {
			if isAlreadyExistsError(err) {
				log.Printf("[DEBUG] Got recoverable error while creating environment %v", err)
				return resource.RetryableError(err)
			} else {
				return resource.NonRetryableError(err)
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	d.SetId(*resp.GetPayload().Environment.EnvironmentName)

	if err := waitForEnvironmentToBeAvailable(d.Id(), d.Timeout(schema.TimeoutCreate), client); err != nil {
		return err
	}

	return resourceAWSEnvironmentRead(d, m)
}

func isAlreadyExistsError(err error) bool {
	if d, ok := err.(*operations.CreateAWSEnvironmentDefault); ok && d.GetPayload() != nil {
		return d.GetPayload().Code == "INVALID_ARGUMENT" && strings.Contains(d.GetPayload().Message, "already exist")
	}
	return false
}

func isNotFoundError(err error) bool {
	if d, ok := err.(*operations.DescribeEnvironmentDefault); ok && d.GetPayload() != nil {
		return d.GetPayload().Code == "NOT_FOUND"
	}
	return false
}

func waitForEnvironmentToBeAvailable(environmentName string, timeout time.Duration, client *environmentsclient.Environments) error {
	stateConf := &resource.StateChangeConf{
		Pending:      []string{"FREEIPA_CREATION_IN_PROGRESS"},
		Target:       []string{"AVAILABLE"},
		Delay:        5 * time.Second,
		Timeout:      timeout,
		PollInterval: 10 * time.Second,
		Refresh: func() (interface{}, string, error) {
			log.Printf("[DEBUG] About to describe environment %s", environmentName)
			params := operations.NewDescribeEnvironmentParams()
			params.WithInput(&environmentsmodels.DescribeEnvironmentRequest{EnvironmentName: &environmentName})
			resp, err := client.Operations.DescribeEnvironment(params)
			if err != nil {
				// Envs that have just been created may not be returned from Describe Environment request because of eventual
				// consistency. We return an empty state to retry.

				if isNotFoundError(err) {
					log.Printf("[DEBUG] Recoverable error describing environment: %s", err)
					return nil, "", nil
				}
				log.Printf("Error describing environment: %s", err)
				return nil, "", err
			}
			log.Printf("Described environment: %s", *resp.GetPayload().Environment.Status)
			return resp, *resp.GetPayload().Environment.Status, nil
		},
	}
	_, err := stateConf.WaitForState()

	return err
}

func resourceAWSEnvironmentRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Environments

	environmentName := d.Id()

	describeParams := operations.NewDescribeEnvironmentParams()
	describeParams.WithInput(&environmentsmodels.DescribeEnvironmentRequest{EnvironmentName: &environmentName})
	describeEnvironmentResp, err := client.Operations.DescribeEnvironment(describeParams)
	if err != nil {
		if envErr, ok := err.(*operations.DescribeEnvironmentDefault); ok {
			if cdp.IsEnvironmentsError(envErr.GetPayload(), "NOT_FOUND", "") {
				d.SetId("") // deleted
				return nil
			}
		}
		return err
	}
	e := describeEnvironmentResp.GetPayload().Environment
	if e == nil {
		d.SetId("") // deleted
		return nil
	}

	d.SetId(*e.EnvironmentName)
	d.Set("environment_name", *e.EnvironmentName)
	d.Set("credential_name", *e.CredentialName)
	d.Set("region", *e.Region)
	d.Set("vpc_id", e.Network.Aws.VpcID)
	sort.Strings(e.Network.SubnetIds)
	d.Set("subnet_ids", e.Network.SubnetIds)
	d.Set("s3_guard_table_name", e.AwsDetails.S3GuardTableName)
	d.Set("tunnel_enabled", e.TunnelEnabled)
	d.Set("description", e.Description)
	d.Set("crn", e.Crn)

	authentication := []interface{}{
		map[string]interface{}{
			"public_key_id": e.Authentication.PublicKeyID,
		},
	}
	if err := d.Set("authentication", authentication); err != nil {
		return fmt.Errorf("error setting authentication: %s", err)
	}

	logStorage := []interface{}{
		map[string]interface{}{
			"storage_location_base": e.LogStorage.AwsDetails.StorageLocationBase,
			"instance_profile":      e.LogStorage.AwsDetails.InstanceProfile,
		},
	}
	if err := d.Set("log_storage", logStorage); err != nil {
		return fmt.Errorf("error setting log_storage: %s", err)
	}

	securityAccess := []interface{}{
		map[string]interface{}{
			"default_security_group_id":  e.SecurityAccess.DefaultSecurityGroupID,
			"security_group_id_for_knox": e.SecurityAccess.SecurityGroupIDForKnox,
		},
	}
	if err := d.Set("security_access", securityAccess); err != nil {
		return fmt.Errorf("error setting security_access: %s", err)
	}

	return nil
}

func resourceAWSEnvironmentUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Environments

	environmentName := d.Id()

	if d.HasChange("credential_name") {
		credentialName := d.Get("credential_name").(string)

		params := operations.NewChangeEnvironmentCredentialParams()
		params.WithInput(&environmentsmodels.ChangeEnvironmentCredentialRequest{
			EnvironmentName: &environmentName,
			CredentialName:  &credentialName,
		})
		_, err := client.Operations.ChangeEnvironmentCredential(params)
		if err != nil {
			return err
		}
	}

	return resourceAWSEnvironmentRead(d, m)
}

func resourceAWSEnvironmentDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Environments

	environmentName := d.Id()
	params := operations.NewDeleteEnvironmentParams()
	params.WithInput(&environmentsmodels.DeleteEnvironmentRequest{EnvironmentName: &environmentName})
	_, err := client.Operations.DeleteEnvironment(params)
	if err != nil {
		return err
	}

	if err := waitForEnvironmentToBeDeleted(d.Id(), d.Timeout(schema.TimeoutDelete), client); err != nil {
		return err
	}

	return nil
}

func waitForEnvironmentToBeDeleted(environmentName string, timeout time.Duration, client *environmentsclient.Environments) error {
	stateConf := &resource.StateChangeConf{
		Pending:      []string{"FREEIPA_DELETE_IN_PROGRESS"},
		Target:       []string{},
		Delay:        5 * time.Second,
		Timeout:      timeout,
		PollInterval: 10 * time.Second,
		Refresh: func() (interface{}, string, error) {
			log.Printf("About to describe environment")
			params := operations.NewDescribeEnvironmentParams()
			params.WithInput(&environmentsmodels.DescribeEnvironmentRequest{EnvironmentName: &environmentName})
			resp, err := client.Operations.DescribeEnvironment(params)
			if err != nil {
				log.Printf("Error describing environment: %s", err)
				if envErr, ok := err.(*operations.DescribeEnvironmentDefault); ok {
					if cdp.IsEnvironmentsError(envErr.GetPayload(), "NOT_FOUND", "") {
						return nil, "", nil
					}
				}
				return nil, "", err
			}
			if resp.GetPayload().Environment == nil {
				log.Printf("Described environment. No environment.")
				return nil, "", nil
			}
			log.Printf("Described environment: %s", *resp.GetPayload().Environment.Status)
			return resp, *resp.GetPayload().Environment.Status, nil
		},
	}
	_, err := stateConf.WaitForState()

	return err
}
