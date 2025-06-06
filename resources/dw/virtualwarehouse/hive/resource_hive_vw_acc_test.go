// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

//go:build hive

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

func TestAccHive_Basic(t *testing.T) {
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
					resource.TestCheckResourceAttr("cdp_dw_vw_hive.test_hive", "name", params.Name),
					resource.TestCheckResourceAttr("cdp_dw_vw_hive.test_hive", "cluster_id", params.ClusterID),
					resource.TestCheckResourceAttr("cdp_dw_vw_hive.test_hive", "database_catalog_id", params.DatabaseCatalogID),
					resource.TestCheckResourceAttrSet("cdp_dw_vw_hive.test_hive", "compactor"),
					resource.TestCheckResourceAttrSet("cdp_dw_vw_hive.test_hive", "jdbc_url"),
					resource.TestCheckResourceAttrSet("cdp_dw_vw_hive.test_hive", "hue_url"),
					resource.TestCheckResourceAttrSet("cdp_dw_vw_hive.test_hive", "jwt_token_gen_url"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccHiveBasicConfig(params hiveTestParameters) string {
	return fmt.Sprintf(`
		resource "cdp_dw_vw_hive" "test_hive" {
		  cluster_id    = %[1]q
		  database_catalog_id = %[2]q
		  name = %[3]q
		  group_size = 2
		  platform_jwt_auth = true
		  enable_sso = true
		  min_group_count = 2
		  max_group_count = 5
		  disable_auto_suspend = false
		  auto_suspend_timeout_seconds = 1200
		  scale_wait_time_seconds = 230
		  max_concurrent_isolated_queries = 10
		  max_nodes_per_isolated_query = 10
		  aws_options = {
			ebs_llap_spill_gb = 300
			tags = {
			  "made-with": "CDP-Terraform-Provider"
			}
		  }
		}
	`, params.ClusterID, params.DatabaseCatalogID, params.Name)
}

func testCheckHiveDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cdp_dw_vw_hive" {
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
