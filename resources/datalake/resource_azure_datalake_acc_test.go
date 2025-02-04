// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.
package datalake_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
	"github.com/cloudera/terraform-provider-cdp/cdpacctest"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

const (
	AzureRegion                          = "ACCEPTANCETEST_AZURE_REGION"
	AzureTenantID                        = "ACCEPTANCETEST_AZURE_TENANT_ID"
	AzureSubscriptionID                  = "ACCEPTANCETEST_AZURE_SUBSCRIPTION_ID"
	AzureApplicationID                   = "ACCEPTANCETEST_AZURE_APPLICATION_ID"
	AzureSecretKey                       = "ACCEPTANCETEST_AZURE_SECRET"
	AzurePublicKey                       = "ACCEPTANCETEST_AZURE_PUBLIC_KEY"
	AzureLoggerManagedIdentity           = "ACCEPTANCETEST_AZURE_LOGGER_MANAGED_IDENTITY"
	AzureLogStorageLocationBase          = "ACCEPTANCETEST_AZURE_LOG_STORAGE_LOCATION_BASE"
	AzureDataStorageLocationBase         = "ACCEPTANCETEST_AZURE_DATA_STORAGE_LOCATION_BASE"
	AzureBackupStorageLocationBase       = "ACCEPTANCETEST_AZURE_BACKUP_STORAGE_LOCATION_BASE"
	AzureVnetID                          = "ACCEPTANCETEST_AZURE_VNET_ID"
	AzureSubnetIDs                       = "ACCEPTANCETEST_AZURE_SUBNET_IDS"
	AzureNetworkResourceGroup            = "ACCEPTANCETEST_AZURE_NETWORK_RESOURCE_GROUP"
	AzureDataAccessRole                  = "ACCEPTANCETEST_AZURE_DATA_ACCESS_ROLE"
	AzureRangerAuditRole                 = "ACCEPTANCETEST_AZURE_RANGER_AUDIT_ROLE"
	AzureRangerCloudAccessAuthorizerRole = "ACCEPTANCETEST_AZURE_RANGER_CLOUD_ACCESS_AUTHORIZER_ROLE"
	AzureManagedIdentity                 = "ACCEPTANCETEST_AZURE_MANAGED_IDENTITY"
	AzureRuntime                         = "ACCEPTANCETEST_AZURE_RUNTIME"
	AzureResourceGroupName               = "ACCEPTANCETEST_AZURE_RESOURCE_GROUP_NAME"
)

func AzureDataLakePreCheck(t *testing.T) {
	envVars := []string{AzureRegion, AzureTenantID, AzureSubscriptionID, AzureApplicationID, AzureSecretKey, AzurePublicKey,
		AzureLoggerManagedIdentity, AzureLogStorageLocationBase, AzureDataStorageLocationBase, AzureBackupStorageLocationBase,
		AzureVnetID, AzureSubnetIDs, AzureNetworkResourceGroup, AzureDataAccessRole, AzureRangerAuditRole,
		AzureRangerCloudAccessAuthorizerRole, AzureManagedIdentity, AzureRuntime, AzureResourceGroupName}
	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			t.Fatalf("%s must be set for acceptance tests", envVar)
		}
	}
}

type azureCredentials struct {
	Name           string
	SubscriptionID string
	TenantID       string
	ApplicationID  string
	SecretKey      string
}

type azureEnvironmentTestParameters struct {
	Name                      string
	Region                    string
	ResourceGroupName         string
	PublicKey                 string // represents an actual ssh key
	LoggerManagedIdentity     string
	LogStorageLocationBase    string
	BackupStorageLocationBase string
	VnetID                    string
	SubnetIDs                 string
	NetworkResourceGroupName  string
}

type azureDataLakeTestParameters struct {
	Name                            string
	ManagedIdentity                 string
	DataStorageLocationBase         string
	Runtime                         string
	DataAccessRole                  string
	RangerAuditRole                 string
	RangerCloudAccessAuthorizerRole string
}

func TestAccAzureDataLake_basic(t *testing.T) {
	credentials := azureCredentials{
		Name:           acctest.RandomWithPrefix(cdpacctest.ResourcePrefix),
		SubscriptionID: os.Getenv(AzureSubscriptionID),
		TenantID:       os.Getenv(AzureTenantID),
		ApplicationID:  os.Getenv(AzureApplicationID),
		SecretKey:      os.Getenv(AzureSecretKey),
	}
	envParams := azureEnvironmentTestParameters{
		Name:                      cdpacctest.RandomShortWithPrefix(cdpacctest.ResourcePrefix),
		Region:                    os.Getenv(AzureRegion),
		ResourceGroupName:         os.Getenv(AzureResourceGroupName),
		PublicKey:                 os.Getenv(AzurePublicKey),
		LoggerManagedIdentity:     os.Getenv(AzureLoggerManagedIdentity),
		LogStorageLocationBase:    os.Getenv(AzureLogStorageLocationBase),
		BackupStorageLocationBase: os.Getenv(AzureBackupStorageLocationBase),
		VnetID:                    os.Getenv(AzureVnetID),
		SubnetIDs:                 os.Getenv(AzureSubnetIDs),
		NetworkResourceGroupName:  os.Getenv(AzureNetworkResourceGroup),
	}
	dlParams := azureDataLakeTestParameters{
		Name:                            cdpacctest.RandomShortWithPrefix(cdpacctest.ResourcePrefix),
		ManagedIdentity:                 os.Getenv(AzureManagedIdentity),
		DataStorageLocationBase:         os.Getenv(AzureDataStorageLocationBase),
		Runtime:                         os.Getenv(AzureRuntime),
		DataAccessRole:                  os.Getenv(AzureDataAccessRole),
		RangerAuditRole:                 os.Getenv(AzureRangerAuditRole),
		RangerCloudAccessAuthorizerRole: os.Getenv(AzureRangerCloudAccessAuthorizerRole),
	}
	resourceName := "cdp_datalake_azure_datalake.test_dl"
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			cdpacctest.PreCheck(t)
			AzureDataLakePreCheck(t)
		},
		ProtoV6ProviderFactories: cdpacctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAzureDataLakeDestroy,
		Steps: []resource.TestStep{
			{
				Config: utils.Concat(
					cdpacctest.TestAccCdpProviderConfig(),
					testAccAzureCredentialBasicConfig(credentials),
					testAccAzureEnvironmentConfig(&envParams),
					testAccAzureDataLakeConfig(&dlParams)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrWith(resourceName, "id", cdpacctest.CheckCrn),
					resource.TestCheckResourceAttr(resourceName, "datalake_name", dlParams.Name),
					resource.TestCheckResourceAttrWith(resourceName, "crn", cdpacctest.CheckCrn),
				),
			},
		},
	})
}

func testAccAzureCredentialBasicConfig(creds azureCredentials) string {
	return fmt.Sprintf(`
        resource "cdp_environments_azure_credential" "test_cred" {
            credential_name = %[1]q
            subscription_id = %[2]q
            tenant_id       = %[3]q
            app_based = {
              application_id = %[4]q
              secret_key     = %[5]q
            }
            description = "Terraform generated test Azure credential"
			}`, creds.Name, creds.SubscriptionID, creds.TenantID, creds.ApplicationID, creds.SecretKey)
}

func testAccAzureEnvironmentConfig(envParams *azureEnvironmentTestParameters) string {
	// TODO: enable CCM tunnel
	return fmt.Sprintf(`
		resource "cdp_environments_azure_environment" "test_env_azure" {
			environment_name = %[1]q
			credential_name = cdp_environments_azure_credential.test_cred.credential_name
			region = %[2]q
			use_public_ip = false
			security_access = {
			  cidr = "0.0.0.0/0"
			}
			endpoint_access_gateway_scheme = "PRIVATE"
			enable_tunnel = false
			public_key = %[3]q
			log_storage = {
			  managed_identity = %[4]q
			  storage_location_base = %[5]q
              backup_storage_location_base = %[6]q
			}
            existing_network_params = {
			  network_id          = %[7]q
			  resource_group_name = %[8]q
			  subnet_ids          = [ %[9]s ]
		    }
		    resource_group_name = %[10]q
			tags = {
			"made-with": "CDP Terraform Provider"
			}
		}`, envParams.Name, envParams.Region, envParams.PublicKey, envParams.LoggerManagedIdentity,
		envParams.LogStorageLocationBase, envParams.BackupStorageLocationBase, envParams.VnetID,
		envParams.NetworkResourceGroupName, envParams.SubnetIDs, envParams.ResourceGroupName)
}

func testAccAzureDataLakeConfig(params *azureDataLakeTestParameters) string {
	return fmt.Sprintf(`
		resource "cdp_environments_id_broker_mappings" "test_idbm_azure" {
			environment_name = cdp_environments_azure_environment.test_env_azure.environment_name
			environment_crn = cdp_environments_azure_environment.test_env_azure.crn
			data_access_role = %[1]q
			ranger_audit_role = %[2]q
			ranger_cloud_access_authorizer_role = %[3]q
			set_empty_mappings = true
		}

		resource "cdp_datalake_azure_datalake" "test_dl" {
			environment_name = cdp_environments_azure_environment.test_env_azure.environment_name
			datalake_name = %[4]q
			managed_identity = %[5]q
			storage_location_base = %[6]q
			scale = "LIGHT_DUTY"
			runtime = %[7]q
			enable_ranger_raz = true
			tags = {
				"made-with": "CDP Terraform Provider"
			}
			depends_on = [ cdp_environments_id_broker_mappings.test_idbm_azure ]
		}
		`,
		params.DataAccessRole, params.RangerAuditRole, params.RangerCloudAccessAuthorizerRole, params.Name, params.ManagedIdentity, params.DataStorageLocationBase,
		params.Runtime)
}

func testAccCheckAzureDataLakeDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cdp_datalake_azure_datalake" {
			continue
		}

		cdpClient := cdpacctest.GetCdpClientForAccTest()
		params := operations.NewDescribeDatalakeParamsWithContext(context.Background())
		params.WithInput(&models.DescribeDatalakeRequest{
			DatalakeName: &rs.Primary.ID,
		})
		_, err := cdpClient.Datalake.Operations.DescribeDatalake(params)
		if err != nil {
			if dlErr, ok := err.(*operations.DescribeDatalakeDefault); ok {
				fmt.Println(dlErr.GetPayload().Code)
				if cdp.IsDatalakeError(dlErr.GetPayload(), "NOT_FOUND", "") {
					return nil
				}
			}
			return err
		}
		return fmt.Errorf("data lake %s not deleted in CDP", rs.Primary.ID)
	}
	return nil
}
