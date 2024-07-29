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
	"github.com/cloudera/terraform-provider-cdp/resources/iam"
	"testing"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
	"github.com/cloudera/terraform-provider-cdp/cdpacctest"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIamMachineUserResourceRoleAssignment_basic(t *testing.T) {
	cdpRegion, err := iam.GetCdpRegionFromConfig()
	if err != nil {
		t.Fatal(err)
	}
	muName := acctest.RandomWithPrefix(cdpacctest.ResourcePrefix)
	resourceRoleCrn := fmt.Sprintf("crn:altus:iam:%s:altus:resourceRole:IamGroupAdmin", cdpRegion)
	resourceName := "cdp_iam_group.test"
	resourceUnderTestName := "cdp_iam_machine_user_resource_role_assignment.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { cdpacctest.PreCheck(t) },
		ProtoV6ProviderFactories: cdpacctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: utils.Concat(
					cdpacctest.TestAccCdpProviderConfig(),
					testAccIamMachineUserResourceRoleAssignmentConfig(muName, resourceRoleCrn)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceUnderTestName, "machine_user", muName),
					testAccCheckResourceAttrWith(resourceUnderTestName, "resource_crn", resourceAttrValue(resourceName, "crn")),
					resource.TestCheckResourceAttr(resourceUnderTestName, "resource_role_crn", resourceRoleCrn),
					testAccCheckResourceAttrWith(resourceUnderTestName, "id", statefValue(muName+"_%s_"+resourceRoleCrn, resourceAttrValue(resourceName, "crn"))),
					testAccCheckIamMachineUserResourceRoleAssignmentExists(muName, resourceAttrValue(resourceName, "crn"), resourceRoleCrn),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccIamMachineUserResourceRoleAssignmentConfig(muName string, resourceRoleCrn string) string {
	return fmt.Sprintf(`
resource "cdp_iam_machine_user" "test" {
  name = %[1]q
}

resource "cdp_iam_group" "test" {
	group_name = %[1]q
}

resource "cdp_iam_machine_user_resource_role_assignment" "test" {
  machine_user = %[1]q
  resource_crn = cdp_iam_group.test.crn
  resource_role_crn = %[2]q
  depends_on = [cdp_iam_machine_user.test]
}
`, muName, resourceRoleCrn)
}

// testAccCheckIamMachineUserRoleAssignmentExists queries the API and retrieves the matching IamMachineUserResourceRoleAssignment via the passed in pointer.
func testAccCheckIamMachineUserResourceRoleAssignmentExists(muName string, resourceCrnFn func(s *terraform.State) (string, error), resourceRoleCrn string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		cdpClient := cdpacctest.GetCdpClientForAccTest()

		params := operations.NewListMachineUserAssignedResourceRolesParamsWithContext(context.TODO())
		params.WithInput(&models.ListMachineUserAssignedResourceRolesRequest{
			MachineUserName: &muName,
		})

		responseOk, err := cdpClient.Iam.Operations.ListMachineUserAssignedResourceRoles(params)
		if err != nil {
			if d, ok := err.(*operations.ListMachineUserAssignedRolesDefault); ok && d.GetPayload() != nil && d.GetPayload().Code == "NOT_FOUND" {
				return fmt.Errorf("machine user %s not found", muName)
			}
			return nil
		}

		resourceCrn, err := resourceCrnFn(s)
		if err != nil {
			return err
		}

		if len(responseOk.Payload.ResourceAssignments) != 1 ||
			(*responseOk.Payload.ResourceAssignments[0].ResourceCrn != resourceCrn &&
				*responseOk.Payload.ResourceAssignments[0].ResourceRoleCrn != resourceRoleCrn) {
			return fmt.Errorf("machine user resource role assignment %s on resource %s not found", resourceRoleCrn, resourceCrn)
		}

		return nil
	}
}

func testAccCheckResourceAttrWith(name, key string, valueFn func(s *terraform.State) (string, error)) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectedValue, err := valueFn(s)
		if err != nil {
			return err
		}

		is, err := modulePrimaryInstanceState(s.RootModule(), name)
		if err != nil {
			return err
		}

		actualValue, ok := is.Attributes[key]
		if !ok {
			return fmt.Errorf("%s: Attribute '%s' not found", name, key)
		}
		if actualValue != expectedValue {
			return fmt.Errorf("%s: Attribute '%s' expected %#v, got %#v", name, key, expectedValue, actualValue)
		}

		return nil
	}
}

func modulePrimaryInstanceState(ms *terraform.ModuleState, name string) (*terraform.InstanceState, error) {
	rs, ok := ms.Resources[name]
	if !ok {
		return nil, fmt.Errorf("Not found: %s in %s", name, ms.Path)
	}

	is := rs.Primary
	if is == nil {
		return nil, fmt.Errorf("No primary instance: %s in %s", name, ms.Path)
	}

	return is, nil
}

func resourceAttrValue(name, key string) func(s *terraform.State) (string, error) {
	return func(s *terraform.State) (string, error) {
		is, err := modulePrimaryInstanceState(s.RootModule(), name)
		if err != nil {
			return "", err
		}

		value, ok := is.Attributes[key]
		if !ok {
			return "", fmt.Errorf("%s: Attribute '%s' not found", name, key)
		}

		return value, nil
	}
}

func statefValue(format string, a ...func(s *terraform.State) (string, error)) func(s *terraform.State) (string, error) {
	return func(s *terraform.State) (string, error) {
		values := make([]any, len(a))
		for i, fn := range a {
			var err error
			values[i], err = fn(s)
			if err != nil {
				return "", err
			}
		}
		return fmt.Sprintf(format, values...), nil
	}
}
