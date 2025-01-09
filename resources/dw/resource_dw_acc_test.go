// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package dw

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
	"github.com/cloudera/terraform-provider-cdp/cdpacctest"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

const (
	AwsAccessKeyID         = "ACCEPTANCETEST_AWS_ACCESS_KEY_ID"
	AwsSecretAccessKey     = "ACCEPTANCETEST_AWS_SECRET_ACCESS_KEY"
	AwsXAccRoleArn         = "ACCEPTANCETEST_AWS_X_ACC_ROLE_ARN"
	AwsRegion              = "ACCEPTANCETEST_AWS_REGION"
	AwsPublicKeyID         = "ACCEPTANCETEST_AWS_PUBLIC_KEY_ID"
	AwsInstanceProfile     = "ACCEPTANCETEST_AWS_INSTANCE_PROFILE"
	AwsStorageLocationBase = "ACCEPTANCETEST_AWS_STORAGE_LOCATION_BASE"
	AwsVpcId               = "ACCEPTANCETEST_AWS_VPC_ID"
	AwsSubnetIds           = "ACCEPTANCETEST_AWS_SUBNET_IDS"
	AwsDataAccessRole      = "ACCEPTANCETEST_AWS_DATA_ACCESS_ROLE"
	AwsRangerAuditRole     = "ACCEPTANCETEST_AWS_RANGER_AUDIT_ROLE"
	AwsAssumerRole         = "ACCEPTANCETEST_AWS_ASSUMER_ROLE"
	AwsRuntime             = "ACCEPTANCETEST_AWS_RUNTIME"
)

type awsEnvironmentTestParameters struct {
	Name                string
	Region              string
	PublicKeyId         string
	InstanceProfile     string
	StorageLocationBase string
	VpcId               string
	SubnetIds           string
}

type awsDataLakeTestParameters struct {
	Name                string
	DataAccessRole      string
	RangerAuditRole     string
	AssumerRole         string
	StorageLocationBase string
	Runtime             string
}

func AwsDataLakePreCheck(t *testing.T) {
	errMsg := "AWS CDW Terraform acceptance testing requires environment variable %s to be set"
	if _, ok := os.LookupEnv(AwsAccessKeyID); !ok {
		t.Fatalf(errMsg, AwsAccessKeyID)
	}
	if _, ok := os.LookupEnv(AwsSecretAccessKey); !ok {
		t.Fatalf(errMsg, AwsSecretAccessKey)
	}
	if _, ok := os.LookupEnv(AwsXAccRoleArn); !ok {
		t.Fatalf(errMsg, AwsXAccRoleArn)
	}
	if _, ok := os.LookupEnv(AwsRegion); !ok {
		t.Fatalf(errMsg, AwsRegion)
	}
	if _, ok := os.LookupEnv(AwsPublicKeyID); !ok {
		t.Fatalf(errMsg, AwsPublicKeyID)
	}
	if _, ok := os.LookupEnv(AwsInstanceProfile); !ok {
		t.Fatalf(errMsg, AwsInstanceProfile)
	}
	if _, ok := os.LookupEnv(AwsStorageLocationBase); !ok {
		t.Fatalf(errMsg, AwsStorageLocationBase)
	}
	if _, ok := os.LookupEnv(AwsVpcId); !ok {
		t.Fatalf(errMsg, AwsVpcId)
	}
	if _, ok := os.LookupEnv(AwsSubnetIds); !ok {
		t.Fatalf(errMsg, AwsSubnetIds)
	}
	if _, ok := os.LookupEnv(AwsDataAccessRole); !ok {
		t.Fatalf(errMsg, AwsDataAccessRole)
	}
	if _, ok := os.LookupEnv(AwsRangerAuditRole); !ok {
		t.Fatalf(errMsg, AwsRangerAuditRole)
	}
	if _, ok := os.LookupEnv(AwsAssumerRole); !ok {
		t.Fatalf(errMsg, AwsAssumerRole)
	}
	if _, ok := os.LookupEnv(AwsRuntime); !ok {
		t.Fatalf(errMsg, AwsRuntime)
	}
}

func PreCheck(t *testing.T) {
	if _, ok := os.LookupEnv(AwsSecretAccessKey); !ok {
		t.Skipf("Terraform acceptance testing requires environment variable %s to be set", AwsSecretAccessKey)
	}

	if os.Getenv(cdp.CdpProfileEnvVar) == "" && os.Getenv(cdp.CdpAccessKeyIdEnvVar) == "" {
		t.Skipf("Terraform acceptance testing requires either %s or %s environment variables to be set", cdp.CdpProfileEnvVar, cdp.CdpAccessKeyIdEnvVar)
	}

	if os.Getenv(cdp.CdpAccessKeyIdEnvVar) != "" {
		if _, ok := os.LookupEnv(cdp.CdpPrivateKeyEnvVar); !ok {
			t.Skipf("Environment variable %s should be set together with %s", cdp.CdpPrivateKeyEnvVar, cdp.CdpAccessKeyIdEnvVar)
		}
	}
}

func TestAccDwCluster_Basic(t *testing.T) {
	PreCheck(t)
	credName := acctest.RandomWithPrefix(cdpacctest.ResourcePrefix)
	awsProvider := cdpacctest.NewAwsProvider(os.Getenv(AwsAccessKeyID), os.Getenv(AwsSecretAccessKey), os.Getenv(AwsRegion))
	accountParams := cdpacctest.NewAwsAccountCredentials(cdpacctest.RandomShortWithPrefix(cdpacctest.ResourcePrefix)).
		WithAccountID(t).
		WithExternalID(t).
		WithPolicy(t)
	envParams := awsEnvironmentTestParameters{
		Name:                cdpacctest.RandomShortWithPrefix(cdpacctest.ResourcePrefix),
		Region:              os.Getenv(AwsRegion),
		PublicKeyId:         os.Getenv(AwsPublicKeyID),
		InstanceProfile:     os.Getenv(AwsInstanceProfile),
		StorageLocationBase: os.Getenv(AwsStorageLocationBase),
		VpcId:               os.Getenv(AwsVpcId),
		SubnetIds:           os.Getenv(AwsSubnetIds),
	}
	dlParams := awsDataLakeTestParameters{
		Name:                cdpacctest.RandomShortWithPrefix(cdpacctest.ResourcePrefix),
		DataAccessRole:      os.Getenv(AwsDataAccessRole),
		RangerAuditRole:     os.Getenv(AwsRangerAuditRole),
		AssumerRole:         os.Getenv(AwsAssumerRole),
		StorageLocationBase: os.Getenv(AwsStorageLocationBase),
		Runtime:             os.Getenv(AwsRuntime),
	}

	const (
		awsClusterResource        = "cdp_dw_aws_cluster.test_data_warehouse_aws"
		hiveResource              = "cdp_dw_vw_hive.test_hive"
		dataVisualizationResource = "cdp_dw_data_visualization.test_dataviz"
	)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			cdpacctest.PreCheck(t)
			AwsDataLakePreCheck(t)
		},
		ProtoV6ProviderFactories: cdpacctest.TestAccProtoV6ProviderFactories,
		ExternalProviders: cdpacctest.ConcatExternalProviders(
			cdpacctest.AwsExternalProvider,
			cdpacctest.TimeExternalProvider,
		),
		CheckDestroy: testCheckClusterDestroy,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: utils.Concat(
					cdpacctest.CreateDefaultRoleAndPolicy(accountParams),
					cdpacctest.TestAccAwsProviderConfig(awsProvider),
					testAccAwsCredentialBasicConfig(credName),
					testAccAwsEnvironmentConfig(&envParams),
					testAccAwsDataLakeConfig(&dlParams),
					testAccAwsClusterBasicConfig(&envParams),
					testAccDwCatalog(),
					testAccHiveVirtualWarehouse(cdpacctest.RandomShortWithPrefix("tf-hive")),
					testAccImpalaVirtualWarehouse(cdpacctest.RandomShortWithPrefix("tf-impala")),
					testAccDataVisualization(cdpacctest.RandomShortWithPrefix("tf-dataviz"))),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(awsClusterResource, "name", envParams.Name),
					resource.TestCheckResourceAttr(awsClusterResource, "status", "Accepted"),

					resource.TestCheckResourceAttrSet(hiveResource, "compactor"),
					resource.TestCheckResourceAttrSet(hiveResource, "jdbc_url"),
					resource.TestCheckResourceAttrSet(hiveResource, "hue_url"),
					resource.TestCheckResourceAttrSet(hiveResource, "jwt_token_gen_url"),

					// TODO vcsomor add checks for Impala!

					resource.TestCheckResourceAttrSet(dataVisualizationResource, "image_version"),
					resource.TestCheckResourceAttr(dataVisualizationResource, "admin_groups.0", "dwx-dummy-ldap-group"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccAwsCredentialBasicConfig(name string) string {
	// Wait for the IAM policy attachment to be created before creating the credential, after a couple of seconds,
	// the CDP credential creation fails, the privileges are not yet available.
	return fmt.Sprintf(`
		resource "time_sleep" "wait_10_seconds" {
		  depends_on = [aws_iam_policy_attachment.test-attach]
		  create_duration = "10s"
		}
		
		resource "cdp_environments_aws_credential" "test_cred" {
		  credential_name = "%[1]s-cred"
		  role_arn        = aws_iam_role.cdp_test_role.arn
		  depends_on = [time_sleep.wait_10_seconds]
		}
`, name)
}

func testAccAwsEnvironmentConfig(envParams *awsEnvironmentTestParameters) string {
	return fmt.Sprintf(`
		resource "cdp_environments_aws_environment" "test_env_dw_aws" {
			environment_name = %[1]q
			credential_name = cdp_environments_aws_credential.test_cred.credential_name
			region = %[2]q
			security_access = {
			cidr = "0.0.0.0/0"
			}
			endpoint_access_gateway_scheme = "PRIVATE"
			enable_tunnel = false
			authentication = {
			public_key_id = %[3]q
			}
			log_storage = {
			instance_profile = %[4]q
			storage_location_base = %[5]q
			}
			vpc_id = %[6]q
			subnet_ids = [ %[7]s ]
			create_private_subnets = true
			create_service_endpoints = false
			tags = {
			"made-with": "CDP Terraform Provider"
			}
		}`, envParams.Name, envParams.Region, envParams.PublicKeyId, envParams.InstanceProfile,
		envParams.StorageLocationBase, envParams.VpcId, envParams.SubnetIds)
}

func testAccAwsDataLakeConfig(params *awsDataLakeTestParameters) string {
	return fmt.Sprintf(`
		resource "cdp_environments_id_broker_mappings" "test_idbm_dw_aws" {
			environment_name = cdp_environments_aws_environment.test_env_dw_aws.environment_name
			environment_crn = cdp_environments_aws_environment.test_env_dw_aws.crn
			data_access_role = %[1]q
			ranger_audit_role = %[2]q
			set_empty_mappings = true
			}
			
		resource "cdp_datalake_aws_datalake" "test_dl_dw_aws" {
			datalake_name = %[3]q
			environment_name = cdp_environments_aws_environment.test_env_dw_aws.environment_name
			instance_profile = %[4]q
			storage_location_base = %[5]q
			scale = "LIGHT_DUTY"
			runtime = %[6]q
			enable_ranger_raz = false
			
			tags = {
				"made-with": "CDP Terraform Provider"
			}
			
			depends_on = [ cdp_environments_id_broker_mappings.test_idbm_dw_aws ]
		}
		`,
		params.DataAccessRole, params.RangerAuditRole, params.Name, params.AssumerRole, params.StorageLocationBase,
		params.Runtime)
}

func testAccAwsClusterBasicConfig(params *awsEnvironmentTestParameters) string {
	return fmt.Sprintf(`
		resource "cdp_dw_aws_cluster" "test_data_warehouse_aws" {
		  crn = cdp_environments_aws_environment.test_env_dw_aws.crn
		  network_settings = {
			worker_subnet_ids = [ %[1]s ]
			load_balancer_subnet_ids = [ %[1]s ]
			use_overlay_network = true
			use_private_load_balancer = true
			use_public_worker_node = false
		  }
          depends_on = [ cdp_datalake_aws_datalake.test_dl_dw_aws ]
		}
	`, params.SubnetIds)
}

func testAccDwCatalog() string {
	return `
		resource "cdp_dw_database_catalog" "test_catalog" {
			cluster_id = cdp_dw_aws_cluster.test_data_warehouse_aws.cluster_id
		}
	`
}

func testAccHiveVirtualWarehouse(name string) string {
	return fmt.Sprintf(`
		resource "cdp_dw_vw_hive" "test_hive" {
			cluster_id = cdp_dw_aws_cluster.test_data_warehouse_aws.cluster_id
			database_catalog_id = cdp_dw_database_catalog.test_catalog.id
			name = %[1]q
			group_size = 2
			platform_jwt_auth = true
			enable_sso = true
			min_group_count = 2
			max_group_count = 5
			disable_auto_suspend = false
			auto_suspend_timeout_seconds = 100
			scale_wait_time_seconds = 230
			max_concurrent_isolated_queries = 10
			max_nodes_per_isolated_query = 10
			aws_options = {
			  availability_zone = "us-west-2a"
			  ebs_llap_spill_gb = 300
			  tags = {
				owner = "cdw-terraform@cloudera.com"
			  }
			}
		}
	`, name)
}

func testCheckClusterDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cdp_dw_aws_cluster" {
			continue
		}

		cdpClient := cdpacctest.GetCdpClientForAccTest()
		params := operations.NewDescribeClusterParamsWithContext(context.Background())
		clusterID := rs.Primary.Attributes["cluster_id"]
		params.WithInput(&models.DescribeClusterRequest{
			ClusterID: &clusterID,
		})

		_, err := cdpClient.Dw.Operations.DescribeCluster(params)
		if err != nil {
			if strings.Contains(err.Error(), "NOT_FOUND") {
				continue
			}
			return err
		}
	}
	return nil
}

func testAccImpalaVirtualWarehouse(name string) string {
	return fmt.Sprintf(`
		resource "cdp_dw_vw_impala" "test_impala" {
			cluster_id = cdp_dw_aws_cluster.test_data_warehouse_aws.cluster_id
			database_catalog_id = cdp_dw_database_catalog.test_catalog.id
			name = %[1]q
		}
	`, name)
}

func testAccDataVisualization(name string) string {
	// NOTE: the LDAP groups vary by CDP environment, the admin_groups has to be injected via an ENV var or some other way, we use now a dummy one
	return fmt.Sprintf(`
		resource "cdp_dw_data_visualization" "test_dataviz" {
			cluster_id = cdp_dw_aws_cluster.test_data_warehouse_aws.cluster_id
			name          = %[1]q
			
			resource_template = "viz-low"
			
			admin_groups = ["dwx-dummy-ldap-group"]
			user_groups  = []
		}
	`, name)
}
