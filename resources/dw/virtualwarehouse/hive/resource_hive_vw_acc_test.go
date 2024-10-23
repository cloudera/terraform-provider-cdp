// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package hive_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
	"github.com/cloudera/terraform-provider-cdp/cdpacctest"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

type hiveTestParameters struct {
	Name              string
	ClusterID         string
	DatabaseCatalogID string
}

func HivePreCheck(t *testing.T) {
	errMsg := "AWS CDW Hive Terraform acceptance testing requires environment variable %s to be set"
	if _, ok := os.LookupEnv("CDW_CLUSTER_ID"); !ok {
		t.Skipf(errMsg, "CDW_CLUSTER_ID")
	}
	if _, ok := os.LookupEnv("CDW_DATABASE_CATALOG_ID"); !ok {
		t.Skipf(errMsg, "CDW_DATABASE_CATALOG_ID")
	}
}

func TestAccHive_basic(t *testing.T) {
	params := hiveTestParameters{
		Name:              cdpacctest.RandomShortWithPrefix(cdpacctest.ResourcePrefix),
		ClusterID:         os.Getenv("CDW_CLUSTER_ID"),
		DatabaseCatalogID: os.Getenv("CDW_DATABASE_CATALOG_ID"),
	}
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			cdpacctest.PreCheck(t)
			HivePreCheck(t)
		},
		ProtoV6ProviderFactories: cdpacctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testCheckHiveDestroy,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: utils.Concat(
					cdpacctest.TestAccCdpProviderConfig(),
					testAccHiveBasicConfig(params)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("cdp_vw_hive.test_hive", "name", params.Name),
					resource.TestCheckResourceAttr("cdp_vw_hive.test_hive", "cluster_id", params.ClusterID),
					resource.TestCheckResourceAttr("cdp_vw_hive.test_hive", "database_catalog_id", params.DatabaseCatalogID),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccHiveBasicConfig(params hiveTestParameters) string {
	return fmt.Sprintf(`
		resource "cdp_vw_hive" "test_hive" {
		  cluster_id = %[1]q
		  database_catalog_id = %[2]q
		  name = %[3]q
		}
	`, params.ClusterID, params.DatabaseCatalogID, params.Name)
}

func testCheckHiveDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cdp_vw_hive" {
			continue
		}

		cdpClient := cdpacctest.GetCdpClientForAccTest()
		params := operations.NewDescribeVwParamsWithContext(context.Background())
		clusterID := rs.Primary.Attributes["cluster_id"]
		params.WithInput(&models.DescribeVwRequest{
			VwID:      &rs.Primary.ID,
			ClusterID: &clusterID,
		})

		_, err := cdpClient.Dw.Operations.DescribeVw(params)
		if err != nil {
			if strings.Contains(err.Error(), "Virtual Warehouse not found") {
				continue
			}
			return err
		}
	}
	return nil
}
