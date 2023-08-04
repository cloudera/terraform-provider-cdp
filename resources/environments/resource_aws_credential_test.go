// Copyright 2023 Cloudera. All Rights Reserved.
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
	"testing"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/cdpacctest"
	"github.com/cloudera/terraform-provider-cdp/resources/environments"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccAwsCredential_basic(t *testing.T) {
	rName := acctest.RandomWithPrefix(cdpacctest.ResourcePrefix)
	resourceName := "cdp_environments_aws_credential.test"
	var credential models.Credential
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { cdpacctest.PreCheck(t) },
		ProtoV6ProviderFactories: cdpacctest.TestAccProtoV6ProviderFactories,
		ExternalProviders: cdpacctest.ConcatExternalProviders(
			cdpacctest.HttpExternalProvider,
			cdpacctest.AwsExternalProvider,
		),
		CheckDestroy: testAccCheckAwsCredentialDestroy,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: utils.Concat(
					cdpacctest.TestAccCdpProviderConfig(),
					cdpacctest.TestAccAwsProviderConfig(),
					testAccAwsCrossAccountRoleConfig(rName),
					testAccAwsCredentialConfig(rName, "aws_iam_role.cdp_cross_account_role.arn")),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", rName),
					resource.TestCheckResourceAttr(resourceName, "credential_name", rName),
					resource.TestCheckResourceAttrWith(resourceName, "crn", cdpacctest.CheckCrn),
					testAccCheckAwsCredentialExists(resourceName, &credential),
					testAccCheckAwsCredentialValues(&credential, rName, ""),
					resource.TestCheckResourceAttrWith(resourceName, "role_arn", func(value string) error {
						return utils.CheckStringEquals("AwsCredentialProperties.RoleArn", credential.AwsCredentialProperties.RoleArn, value)
					}),
				),
			},
			// ImportState testing
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsCredential_withDescription(t *testing.T) {
	rName := acctest.RandomWithPrefix(cdpacctest.ResourcePrefix)
	resourceName := "cdp_environments_aws_credential.test"
	var credential models.Credential
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { cdpacctest.PreCheck(t) },
		ProtoV6ProviderFactories: cdpacctest.TestAccProtoV6ProviderFactories,
		ExternalProviders: cdpacctest.ConcatExternalProviders(
			cdpacctest.HttpExternalProvider,
			cdpacctest.AwsExternalProvider,
		),
		CheckDestroy: testAccCheckAwsCredentialDestroy,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: utils.Concat(
					cdpacctest.TestAccCdpProviderConfig(),
					cdpacctest.TestAccAwsProviderConfig(),
					testAccAwsCrossAccountRoleConfig(rName),
					testAccAwsCredentialConfigWithDescription(rName, "aws_iam_role.cdp_cross_account_role.arn", rName)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", rName),
					resource.TestCheckResourceAttr(resourceName, "credential_name", rName),
					resource.TestCheckResourceAttr(resourceName, "description", rName),
					resource.TestCheckResourceAttrWith(resourceName, "crn", cdpacctest.CheckCrn),
					testAccCheckAwsCredentialExists(resourceName, &credential),
					testAccCheckAwsCredentialValues(&credential, rName, rName),
					resource.TestCheckResourceAttrWith(resourceName, "role_arn", func(value string) error {
						return utils.CheckStringEquals("AwsCredentialProperties.RoleArn", credential.AwsCredentialProperties.RoleArn, value)
					}),
				),
			},
			// ImportState testing
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccAwsCredentialConfig(rName string, roleArn string) string {
	return fmt.Sprintf(`
resource "cdp_environments_aws_credential" "test" {
  credential_name = %[1]q
  role_arn        = %[2]s
}
`, rName, roleArn)
}

func testAccAwsCredentialConfigWithDescription(rName string, roleArn string, description string) string {
	return fmt.Sprintf(`
resource "cdp_environments_aws_credential" "test" {
  credential_name = %[1]q
  role_arn        = %[2]s
  description     = %[3]q
}
`, rName, roleArn, description)
}

func testAccAwsCrossAccountRoleConfig(rName string) string {
	return fmt.Sprintf(`
data "cdp_environments_aws_credential_prerequisites" "credential_prerequisites" {}

# TODO: Replace this with minimal policy?
data "http" "cdp_cross_account_account_policy_doc" {
  url = "https://raw.githubusercontent.com/hortonworks/cloudbreak/master/cloud-aws-common/src/main/resources/definitions/aws-cb-policy.json"
}

data "aws_iam_policy_document" "cdp_cross_account_assume_role_policy_doc" {
  version = "2012-10-17"

  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"

    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::${data.cdp_environments_aws_credential_prerequisites.credential_prerequisites.account_id}:root"]
    }

    condition {
      test     = "StringEquals"
      variable = "sts:ExternalId"

      values = ["${data.cdp_environments_aws_credential_prerequisites.credential_prerequisites.external_id}"]
    }
  }
}

resource "aws_iam_policy" "cdp_cross_account_policy" {
  name = %[1]q
  policy = data.http.cdp_cross_account_account_policy_doc.response_body
}

resource "aws_iam_role" "cdp_cross_account_role" {
  name = %[1]q
  assume_role_policy = data.aws_iam_policy_document.cdp_cross_account_assume_role_policy_doc.json
}

resource "aws_iam_role_policy_attachment" "cdp_cross_account_policy_attachment" {
  role = aws_iam_role.cdp_cross_account_role.name
  policy_arn = aws_iam_policy.cdp_cross_account_policy.arn
}
`, rName)
}

// testAccCheckAwsCredentialExists queries the API and retrieves the matching AwsCredential via the passed in pointer.
func testAccCheckAwsCredentialExists(resourceName string, credential *models.Credential) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		// TODO: I could not find a way to get the reference to the provider in the new framework. If possible we should
		// avoid creating a new CDP client
		cdpClient := cdpacctest.GetCdpClientForAccTest()

		c, err := environments.FindCredentialByName(context.Background(), cdpClient, rs.Primary.ID)
		if err != nil {
			return nil
		}

		if c == nil {
			return fmt.Errorf("credential %s not found in CDP", rs.Primary.ID)
		}

		// return the value via passed in pointer
		*credential = *c

		return nil
	}
}

func testAccCheckAwsCredentialValues(credential *models.Credential, rName string, description string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if *credential.CredentialName != rName {
			return utils.CheckStringEquals("credential.CredentialName", rName, *credential.CredentialName)
		}

		if *credential.CloudPlatform != "AWS" {
			return utils.CheckStringEquals("credential.CloudPlatform", "AWS", *credential.CloudPlatform)
		}

		if credential.Description != description {
			return utils.CheckStringEquals("credential.Description", description, credential.Description)
		}

		return nil
	}
}

func testAccCheckAwsCredentialDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cdp_environments_aws_credential" {
			continue
		}

		cdpClient := cdpacctest.GetCdpClientForAccTest()

		c, err := environments.FindCredentialByName(context.Background(), cdpClient, rs.Primary.ID)
		if err != nil {
			return nil
		}

		if c != nil {
			return fmt.Errorf("credential %s not deleted in CDP", rs.Primary.ID)
		}
	}
	return nil
}
