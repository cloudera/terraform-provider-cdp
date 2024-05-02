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

func TestAccIamMachineUserGroupAssignment_basic(t *testing.T) {
	rName := acctest.RandomWithPrefix(cdpacctest.ResourcePrefix)
	grName := acctest.RandomWithPrefix(cdpacctest.ResourcePrefix)
	resourceName := "cdp_iam_machine_user_group_assignment.test"
	var credential models.ListGroupsForMachineUserResponse
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { cdpacctest.PreCheck(t) },
		ProtoV6ProviderFactories: cdpacctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: utils.Concat(
					cdpacctest.TestAccCdpProviderConfig(),
					testAccIamMachineUserConfig(rName),
					testAccIamMachineUserGroupAssignmentConfig(rName, grName)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "id", rName+"_"+grName),
					testAccCheckIamMachineUserGroupAssignmentExists(rName, grName, &credential),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccIamMachineUserGroupAssignmentConfig(rName string, grName string) string {
	return fmt.Sprintf(`
resource "cdp_iam_machine_user_group_assignment" "test" {
  machine_user = %[1]q
  group = %[2]q
}
`, rName, grName)
}

// testAccCheckIamMachineUserGroupAssignmentExists queries the API and retrieves the matching IamMachineUserGroupAssignment via the passed in pointer.
func testAccCheckIamMachineUserGroupAssignmentExists(rName string, grName string, mu *models.ListGroupsForMachineUserResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		cdpClient := cdpacctest.GetCdpClientForAccTest()

		params := operations.NewListGroupsForMachineUserParamsWithContext(context.TODO())
		params.WithInput(&models.ListGroupsForMachineUserRequest{
			MachineUserName: &rName,
		})

		responseOk, err := cdpClient.Iam.Operations.ListGroupsForMachineUser(params)
		if err != nil {
			if d, ok := err.(*operations.ListGroupsForMachineUserDefault); ok && d.GetPayload() != nil && d.GetPayload().Code == "NOT_FOUND" {
				return fmt.Errorf("machine user %s not found", rName)
			}
			return nil
		}

		grParams := operations.NewListGroupsParamsWithContext(context.TODO())
		grParams.WithInput(&models.ListGroupsRequest{
			GroupNames: []string{grName},
		})

		grRespOk, err := cdpClient.Iam.Operations.ListGroups(grParams)
		if err != nil {
			if d, ok := err.(*operations.ListGroupsDefault); ok && d.GetPayload() != nil && d.GetPayload().Code == "NOT_FOUND" {
				return fmt.Errorf("group %s not found", grName)
			}
			return nil
		}

		if len(grRespOk.Payload.Groups) == 1 && grRespOk.Payload.Groups[0] != nil {
			grCrn := grRespOk.Payload.Groups[0].Crn
			found := false
			for _, v := range responseOk.Payload.GroupCrns {
				if *grCrn == v {
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("machine user group assignment %s not found", grName)
			}
		} else {
			return fmt.Errorf("group %s not found", grName)
		}

		return nil
	}
}
