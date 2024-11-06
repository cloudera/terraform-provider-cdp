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

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
	"github.com/cloudera/terraform-provider-cdp/cdpacctest"
	"github.com/cloudera/terraform-provider-cdp/resources/iam"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

func TestAccIamMachineUserRoleAssignment_basic(t *testing.T) {
	cdpRegion, err := iam.GetCdpRegionFromConfig()
	if err != nil {
		t.Fatal(err)
	}
	muName := acctest.RandomWithPrefix(cdpacctest.ResourcePrefix)
	roleCrn := fmt.Sprintf("crn:altus:iam:%s:altus:role:IamViewer", cdpRegion)
	resourceName := "cdp_iam_machine_user_role_assignment.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { cdpacctest.PreCheck(t) },
		ProtoV6ProviderFactories: cdpacctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: utils.Concat(
					cdpacctest.TestAccCdpProviderConfig(),
					testAccIamMachineUserRoleAssignmentConfig(muName, roleCrn)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "machine_user", muName),
					resource.TestCheckResourceAttr(resourceName, "role", roleCrn),
					resource.TestCheckResourceAttr(resourceName, "id", muName+"_"+roleCrn),
					testAccCheckIamMachineUserRoleAssignmentExists(muName, roleCrn),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccIamMachineUserRoleAssignmentConfig(muName string, roleName string) string {
	return fmt.Sprintf(`
resource "cdp_iam_machine_user" "test" {
  name = %[1]q
}

resource "cdp_iam_machine_user_role_assignment" "test" {
  machine_user = %[1]q
  role = %[2]q
  depends_on = [cdp_iam_machine_user.test]
}
`, muName, roleName)
}

// testAccCheckIamMachineUserRoleAssignmentExists queries the API and retrieves the matching IamMachineUserRoleAssignment via the passed in pointer.
func testAccCheckIamMachineUserRoleAssignmentExists(muName, roleName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		cdpClient := cdpacctest.GetCdpClientForAccTest()

		params := operations.NewListMachineUserAssignedRolesParamsWithContext(context.TODO())
		params.WithInput(&models.ListMachineUserAssignedRolesRequest{
			MachineUserName: &muName,
		})

		responseOk, err := cdpClient.Iam.Operations.ListMachineUserAssignedRoles(params)
		if err != nil {
			if d, ok := err.(*operations.ListMachineUserAssignedRolesDefault); ok && d.GetPayload() != nil && d.GetPayload().Code == "NOT_FOUND" {
				return fmt.Errorf("machine user %s not found", muName)
			}
			return nil
		}

		if len(responseOk.Payload.RoleCrns) != 1 || responseOk.Payload.RoleCrns[0] != roleName {
			return fmt.Errorf("machine user role assignment %s not found", roleName)
		}

		return nil
	}
}
