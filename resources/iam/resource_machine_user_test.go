// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package iam_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
	"github.com/cloudera/terraform-provider-cdp/cdpacctest"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIamMachineUser_basic(t *testing.T) {
	rName := acctest.RandomWithPrefix(cdpacctest.ResourcePrefix)
	resourceName := "cdp_iam_machine_user.test"
	var credential models.MachineUser
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { cdpacctest.PreCheck(t) },
		ProtoV6ProviderFactories: cdpacctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIamMachineUserDestroy,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: utils.Concat(
					cdpacctest.TestAccCdpProviderConfig(),
					testAccIamMachineUserConfig(rName)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttrWith(resourceName, "id", cdpacctest.CheckCrn),
					testAccCheckIamMachineUserExists(resourceName, &credential),
					testAccCheckIamMachineUserValues(&credential, rName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccIamMachineUserConfig(rName string) string {
	return fmt.Sprintf(`
resource "cdp_iam_machine_user" "test" {
  name = %[1]q
}
`, rName)
}

// testAccCheckIamMachineUserExists queries the API and retrieves the matching IamMachineUser via the passed in pointer.
func testAccCheckIamMachineUserExists(resourceName string, mu *models.MachineUser) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		cdpClient := cdpacctest.GetCdpClientForAccTest()

		params := operations.NewListMachineUsersParamsWithContext(context.TODO())
		params.WithInput(&models.ListMachineUsersRequest{
			MachineUserNames: []string{rs.Primary.ID},
		})

		responseOk, err := cdpClient.Iam.Operations.ListMachineUsers(params)
		if err != nil {
			return nil
		}

		if len(responseOk.Payload.MachineUsers) != 1 || responseOk.Payload.MachineUsers[0] == nil {
			return fmt.Errorf("machine user %s not found in CDP", rs.Primary.ID)
		}

		// return the value via passed in pointer
		*mu = *responseOk.Payload.MachineUsers[0]

		return nil
	}
}

func testAccCheckIamMachineUserValues(mu *models.MachineUser, rName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if *mu.MachineUserName != rName {
			return utils.CheckStringEquals("machine_user.Name", rName, *mu.MachineUserName)
		}

		return nil
	}
}

func testAccCheckIamMachineUserDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cdp_environments_aws_credential" {
			continue
		}

		cdpClient := cdpacctest.GetCdpClientForAccTest()

		params := operations.NewListMachineUsersParamsWithContext(context.TODO())
		params.WithInput(&models.ListMachineUsersRequest{
			MachineUserNames: []string{rs.Primary.ID},
		})

		responseOk, err := cdpClient.Iam.Operations.ListMachineUsers(params)
		if err != nil {
			return nil
		}

		if len(responseOk.Payload.MachineUsers) == 0 || responseOk.Payload.MachineUsers[0] == nil {
			return nil
		}

		if len(responseOk.Payload.MachineUsers) == 1 || responseOk.Payload.MachineUsers[0] != nil {
			return fmt.Errorf("machine user %s not deleted in CDP", rs.Primary.ID)
		}
	}
	return nil
}
