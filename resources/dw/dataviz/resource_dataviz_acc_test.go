// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

//go:build dataviz

package dataviz_test

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

type datavizTestParameters struct {
	Name      string
	ClusterID string
}

func dataVisualizationTestPreCheck(t *testing.T) {
	errMsg := "AWS CDW Data Visualization Terraform acceptance testing requires environment variable %s to be set"
	if _, ok := os.LookupEnv("CDW_CLUSTER_ID"); !ok {
		t.Skipf(errMsg, "CDW_CLUSTER_ID")
	}
}

func TestAccDataViz_Low(t *testing.T) {
	params := datavizTestParameters{
		Name:      cdpacctest.RandomShortWithPrefix(cdpacctest.ResourcePrefix),
		ClusterID: os.Getenv("CDW_CLUSTER_ID"),
	}

	const testResource = "cdp_dw_data_visualization.test-dataviz-low"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			cdpacctest.PreCheck(t)
			dataVisualizationTestPreCheck(t)
		},
		ProtoV6ProviderFactories: cdpacctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testCheckDataVizDestroy,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: utils.Concat(
					cdpacctest.TestAccCdpProviderConfig(),
					testAccDtaVizLowConfig(params)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(testResource, "name", params.Name),
					resource.TestCheckResourceAttr(testResource, "cluster_id", params.ClusterID),
					resource.TestCheckResourceAttrSet(testResource, "image_version"),
					resource.TestCheckResourceAttr(testResource, "admin_groups.0", "dwx-viz"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccDtaVizLowConfig(p datavizTestParameters) string {
	return fmt.Sprintf(`
		resource "cdp_dw_data_visualization" "test-dataviz-low" {
		  cluster_id    = %[1]q
		  name          = %[2]q
		
		  resource_template = "viz-low"
		
		  admin_groups = ["dwx-viz"]
		  user_groups  = []
		}
	`, p.ClusterID, p.Name)
}

func testCheckDataVizDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cdp_dw_data_visualization" {
			continue
		}

		cdpClient := cdpacctest.GetCdpClientForAccTest()
		clusterID := rs.Primary.Attributes["cluster_id"]
		_, err := cdpClient.Dw.Operations.
			DescribeDataVisualization(
				operations.NewDescribeDataVisualizationParamsWithContext(context.Background()).
					WithInput(&models.DescribeDataVisualizationRequest{
						ClusterID:           &clusterID,
						DataVisualizationID: &rs.Primary.ID,
					}))
		if err != nil {
			if strings.Contains(err.Error(), "unable to get viz-webapp") {
				continue
			}
			return err
		}
	}
	return nil
}
