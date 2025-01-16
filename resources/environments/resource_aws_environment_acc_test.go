// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package environments_test

import (
	"context"
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/cdpacctest"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

const (
	AwsXAccRoleArn         = "ACCEPTANCETEST_AWS_X_ACC_ROLE_ARN"
	AwsRegion              = "ACCEPTANCETEST_AWS_REGION"
	AwsPublicKeyId         = "ACCEPTANCETEST_AWS_PUBLIC_KEY_ID"
	AwsInstanceProfile     = "ACCEPTANCETEST_AWS_INSTANCE_PROFILE"
	AwsStorageLocationBase = "ACCEPTANCETEST_AWS_STORAGE_LOCATION_BASE"
	AwsVpcId               = "ACCEPTANCETEST_AWS_VPC_ID"
	AwsSubnetIds           = "ACCEPTANCETEST_AWS_SUBNET_IDS"
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

var (
	preCheckOnce sync.Once
)

func TestAccAwsEnvironment_basic(t *testing.T) {
	credName := acctest.RandomWithPrefix(cdpacctest.ResourcePrefix)
	params := awsEnvironmentTestParameters{
		Name:                cdpacctest.RandomShortWithPrefix(cdpacctest.ResourcePrefix),
		Region:              os.Getenv(AwsRegion),
		PublicKeyId:         os.Getenv(AwsPublicKeyId),
		InstanceProfile:     os.Getenv(AwsInstanceProfile),
		StorageLocationBase: os.Getenv(AwsStorageLocationBase),
		VpcId:               os.Getenv(AwsVpcId),
		SubnetIds:           os.Getenv(AwsSubnetIds),
	}
	resourceName := "cdp_environments_aws_environment.test_env"
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			cdpacctest.PreCheck(t)
			AwsEnvironmentPreCheck(t)
		},
		ProtoV6ProviderFactories: cdpacctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAwsEnvironmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: utils.Concat(
					cdpacctest.TestAccCdpProviderConfig(),
					testAccAwsCredentialBasicConfig(credName, os.Getenv(AwsXAccRoleArn)),
					testAccRecipeConfig(fmt.Sprintf("%s_recipe", params.Name)),
					testAccAwsEnvironmentConfig("cdp_environments_aws_credential.test_cred.credential_name", &params)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrWith(resourceName, "id", cdpacctest.CheckCrn),
					resource.TestCheckResourceAttr(resourceName, "environment_name", params.Name),
					resource.TestCheckResourceAttrWith(resourceName, "crn", cdpacctest.CheckCrn),
				),
			},
		},
	})
}

func AwsEnvironmentPreCheck(t *testing.T) {
	preCheckOnce.Do(func() {
		errMsg := "AWS Terraform acceptance testing requires environment variable %s to be set"
		if _, ok := os.LookupEnv(AwsXAccRoleArn); !ok {
			t.Fatalf(errMsg, AwsXAccRoleArn)
		}
		if _, ok := os.LookupEnv(AwsRegion); !ok {
			t.Fatalf(errMsg, AwsRegion)
		}
		if _, ok := os.LookupEnv(AwsPublicKeyId); !ok {
			t.Fatalf(errMsg, AwsPublicKeyId)
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
	})
}

func testAccAwsCredentialBasicConfig(rName string, roleArn string) string {
	return fmt.Sprintf(`
resource "cdp_environments_aws_credential" "test_cred" {
  credential_name = %[1]q
  role_arn        = %[2]q
}
`, rName, roleArn)
}

func testAccRecipeConfig(recipeName string) string {
	return fmt.Sprintf(`
resource "cdp_recipe" "test_recipe" {
  name = %[1]q
  content = <<EOF
#!/bin/bash
echo 'some content'
EOF
  type = "PRE_SERVICE_DEPLOYMENT"
}
`, recipeName)
}

func testAccAwsEnvironmentConfig(credName string, envParams *awsEnvironmentTestParameters) string {
	return fmt.Sprintf(`
	resource "cdp_environments_aws_environment" "test_env" {
		environment_name = %[1]q
		credential_name = %[2]s
		region = %[3]q
		security_access = {
		  cidr = "0.0.0.0/0"
		}
		endpoint_access_gateway_scheme = "PRIVATE"
		enable_tunnel = false
		authentication = {
		  public_key_id = %[4]q
		}
		log_storage = {
		  instance_profile = %[5]q
		  storage_location_base = %[6]q
		}
		vpc_id = %[7]q
		subnet_ids = [ %[8]s ]
		create_private_subnets = true
		create_service_endpoints = false
		tags = {
		  "made-with": "CDP Terraform Provider"
		}
		polling_options = {
		  async = false
		}
		freeipa = {
    		recipes = [cdp_recipe.test_recipe.name]
  		}
	  }
`, envParams.Name, credName, envParams.Region, envParams.PublicKeyId, envParams.InstanceProfile, envParams.StorageLocationBase, envParams.VpcId, envParams.SubnetIds)
}

func testAccCheckAwsEnvironmentDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cdp_environments_aws_environment" {
			continue
		}

		cdpClient := cdpacctest.GetCdpClientForAccTest()
		params := operations.NewDescribeEnvironmentParamsWithContext(context.Background())
		params.WithInput(&models.DescribeEnvironmentRequest{
			EnvironmentName: &rs.Primary.ID,
		})
		_, err := cdpClient.Environments.Operations.DescribeEnvironment(params)
		if err != nil {
			if envErr, ok := err.(*operations.DescribeEnvironmentDefault); ok {
				if cdp.IsEnvironmentsError(envErr.GetPayload(), "NOT_FOUND", "") {
					return nil
				}
			}
			return err
		}
		return fmt.Errorf("environment %s not deleted in CDP", rs.Primary.ID)
	}
	return nil
}
