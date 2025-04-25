// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

//go:build impala

package impala_test

import (
	"context"
	"fmt"
	"net/http"
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

type impalaTestParameters struct {
	Name                             string
	ClusterID                        string
	DatabaseCatalogID                string
	ImageVersion                     string
	TshirtSize                       string
	AutoSuspendTimeoutSeconds        int
	DisableAutoSuspend               bool
	ImpalaScaleDownDelaySeconds      int
	ImpalaScaleUpDelaySeconds        int
	MaxClusters                      int
	MinClusters                      int
	ScratchSpaceLimit                int
	HighAvailabilityMode             string
	EnableShutdownOfCoordinator      bool
	ShutdownOfCoordinatorDelaySecs   int
	NumOfActiveCoordinators          int
	EnableCatalogHighAvailability    bool
	EnableStatestoreHighAvailability bool
	EnableUnifiedAnalytics           bool
	MaxQueries                       int
	MaxNodesPerQuery                 int
	InstanceType                     string
	AvailabilityZone                 string
	HiveAuthenticationMode           string
	PlatformJwtAuth                  bool
	ImpalaQueryLog                   bool
	EbsLlapSpillGb                   int
	Tags                             []Tag
	EnableSso                        bool
}

type Tag struct {
	Key   string
	Value string
}

func ImpalaPreCheck(t *testing.T) {
	errMsg := "AWS CDW Impala Terraform acceptance testing requires environment variable %s to be set"
	if _, ok := os.LookupEnv("CDW_CLUSTER_ID"); !ok {
		t.Skipf(errMsg, "CDW_CLUSTER_ID")
	}
	if _, ok := os.LookupEnv("CDW_DATABASE_CATALOG_ID"); !ok {
		t.Skipf(errMsg, "CDW_DATABASE_CATALOG_ID")
	}
}

func TestAccImpala_basic(t *testing.T) {
	params := impalaTestParameters{
		Name:              cdpacctest.RandomShortWithPrefix(cdpacctest.ResourcePrefix),
		ClusterID:         os.Getenv("CDW_CLUSTER_ID"),
		DatabaseCatalogID: os.Getenv("CDW_DATABASE_CATALOG_ID"),
	}
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			cdpacctest.PreCheck(t)
			ImpalaPreCheck(t)
		},
		ProtoV6ProviderFactories: cdpacctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testCheckImpalaDestroy,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: utils.Concat(
					cdpacctest.TestAccCdpProviderConfig(),
					testAccImpalaBasicConfig(params)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("cdp_dw_vw_impala.test_impala", "name", params.Name),
					resource.TestCheckResourceAttr("cdp_dw_vw_impala.test_impala", "cluster_id", params.ClusterID),
					resource.TestCheckResourceAttr("cdp_dw_vw_impala.test_impala", "database_catalog_id", params.DatabaseCatalogID),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func formatTags(tags []Tag) string {
	var tagStrings []string
	for _, tag := range tags {
		tagStrings = append(tagStrings, fmt.Sprintf("    {\n      key   = %q\n      value = %q\n    }", tag.Key, tag.Value))
	}
	return strings.Join(tagStrings, ",\n")
}

func testAccImpalaBasicConfig(params impalaTestParameters) string {
	var config strings.Builder
	config.WriteString(fmt.Sprintf(`
		resource "cdp_dw_vw_impala" "test_impala" {
		  cluster_id = %q
		  database_catalog_id = %q
		  name = %q
	`, params.ClusterID, params.DatabaseCatalogID, params.Name))

	if params.ImageVersion != "" {
		config.WriteString(fmt.Sprintf("\n  image_version = %q", params.ImageVersion))
	}

	if params.TshirtSize != "" {
		config.WriteString(fmt.Sprintf("\n  tshirt_size = %q", params.TshirtSize))
	}

	if params.AutoSuspendTimeoutSeconds > 0 ||
		params.DisableAutoSuspend ||
		params.ImpalaScaleDownDelaySeconds > 0 ||
		params.ImpalaScaleUpDelaySeconds > 0 ||
		params.MaxClusters > 0 ||
		params.MinClusters > 0 {

		config.WriteString("\n  autoscaling = {")

		if params.AutoSuspendTimeoutSeconds > 0 {
			config.WriteString(fmt.Sprintf("\n    auto_suspend_timeout_seconds = %d", params.AutoSuspendTimeoutSeconds))
		}
		if params.DisableAutoSuspend {
			config.WriteString("\n    disable_auto_suspend = true")
		} else {
			config.WriteString("\n    disable_auto_suspend = false")
		}
		if params.ImpalaScaleDownDelaySeconds > 0 {
			config.WriteString(fmt.Sprintf("\n    scale_down_delay_seconds = %d", params.ImpalaScaleDownDelaySeconds))
		}
		if params.ImpalaScaleUpDelaySeconds > 0 {
			config.WriteString(fmt.Sprintf("\n    scale_up_delay_seconds = %d", params.ImpalaScaleUpDelaySeconds))
		}
		if params.MaxClusters > 0 {
			config.WriteString(fmt.Sprintf("\n    max_clusters = %d", params.MaxClusters))
		}
		if params.MinClusters > 0 {
			config.WriteString(fmt.Sprintf("\n    min_clusters = %d", params.MinClusters))
		}
		config.WriteString("\n  }")
	}

	if params.ScratchSpaceLimit > 0 {
		config.WriteString(fmt.Sprintf("\n  aws_options = {\n    scratch_space_limit = %d\n  }", params.ScratchSpaceLimit))
	}

	if params.HighAvailabilityMode != "" ||
		params.EnableShutdownOfCoordinator ||
		params.ShutdownOfCoordinatorDelaySecs > 0 ||
		params.NumOfActiveCoordinators > 0 ||
		params.EnableCatalogHighAvailability ||
		params.EnableStatestoreHighAvailability {
		config.WriteString("\n  ha_settings = {")

		if params.HighAvailabilityMode != "" {
			config.WriteString(fmt.Sprintf("\n    high_availability_mode = %q", params.HighAvailabilityMode))
		}
		if params.EnableShutdownOfCoordinator {
			config.WriteString("\n    enable_shutdown_of_coordinator = true")
		} else {
			config.WriteString("\n    enable_shutdown_of_coordinator = false")
		}
		if params.ShutdownOfCoordinatorDelaySecs > 0 {
			config.WriteString(fmt.Sprintf("\n    shutdown_of_coordinator_delay_secs = %d", params.ShutdownOfCoordinatorDelaySecs))
		}
		if params.NumOfActiveCoordinators > 0 {
			config.WriteString(fmt.Sprintf("\n    num_of_active_coordinators = %d", params.NumOfActiveCoordinators))
		}
		if params.EnableCatalogHighAvailability {
			config.WriteString("\n    enable_catalog_high_availability = true")
		} else {
			config.WriteString("\n    enable_catalog_high_availability = false")
		}
		if params.EnableStatestoreHighAvailability {
			config.WriteString("\n    enable_statestore_high_availability = true")
		} else {
			config.WriteString("\n    enable_statestore_high_availability = false")
		}
		config.WriteString("\n  }")
	}

	if params.EnableUnifiedAnalytics {
		config.WriteString("\n  enable_unified_analytics = true")
	}

	if params.MaxQueries > 0 || params.MaxNodesPerQuery > 0 {
		config.WriteString("\n  query_isolation_options = {")
		if params.MaxQueries > 0 {
			config.WriteString(fmt.Sprintf("\n    max_queries = %d", params.MaxQueries))
		}
		if params.MaxNodesPerQuery > 0 {
			config.WriteString(fmt.Sprintf("\n    max_nodes_per_query = %d", params.MaxNodesPerQuery))
		}
		config.WriteString("\n  }")
	}

	if params.InstanceType != "" {
		config.WriteString(fmt.Sprintf("\n  instance_type = %q", params.InstanceType))
	}
	if params.AvailabilityZone != "" {
		config.WriteString(fmt.Sprintf("\n  availability_zone = %q", params.AvailabilityZone))
	}
	if params.PlatformJwtAuth {
		config.WriteString("\n  platform_jwt_auth = true")
	}
	if params.ImpalaQueryLog {
		config.WriteString("\n  query_log = true")
	}

	if len(params.Tags) > 0 {
		config.WriteString("\n  tags = [\n    " + formatTags(params.Tags) + "\n  ]")
	}

	if params.EnableSso {
		config.WriteString("\n  enable_sso = true")
	}

	config.WriteString(`

  lifecycle {
    ignore_changes = [
      aws_options["spill_to_s3_uri"],
      last_updated,
      node_count,
      status
    ]
  }
  `)

	config.WriteString("\n}")

	result := config.String()
	fmt.Println("Generated Config:", result)
	return result
}

func testCheckImpalaDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cdp_dw_vw_impala" {
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

func TestAccImpalaImageVersion(t *testing.T) {
	clusterID := os.Getenv("CDW_CLUSTER_ID")
	if clusterID == "" {
		t.Fatal("Environment variable CDW_CLUSTER_ID must be set")
	}

	databaseCatalogID := os.Getenv("CDW_DATABASE_CATALOG_ID")
	if databaseCatalogID == "" {
		t.Fatal("Environment variable CDW_DATABASE_CATALOG_ID must be set")
	}

	latestImageVersion, err := fetchLatestImageVersion(clusterID)
	if err != nil {
		t.Fatalf("Error fetching latest image version: %v", err)
	}

	name := cdpacctest.RandomShortWithPrefix(cdpacctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			cdpacctest.PreCheck(t)
			ImpalaPreCheck(t)
		},
		ProtoV6ProviderFactories: cdpacctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testCheckImpalaDestroy,
		Steps: []resource.TestStep{
			{
				Config: utils.Concat(
					cdpacctest.TestAccCdpProviderConfig(),
					fmt.Sprintf(`
resource "cdp_dw_vw_impala" "test_impala" {
  cluster_id         = %q
  database_catalog_id = %q
  name               = %q
  image_version      = %q
  tshirt_size        = "xsmall"

  autoscaling = {
    auto_suspend_timeout_seconds = 350
    disable_auto_suspend         = false
    scale_down_delay_seconds     = 330
    scale_up_delay_seconds       = 40
    max_clusters                 = 4
    min_clusters                 = 2
  }

  aws_options = {
    scratch_space_limit = 300
  }

  ha_settings = {
    high_availability_mode               = "ACTIVE_PASSIVE"
    enable_shutdown_of_coordinator      = false
    shutdown_of_coordinator_delay_secs  = 360
    num_of_active_coordinators          = 2
    enable_catalog_high_availability    = false
    enable_statestore_high_availability = false
  }

  enable_unified_analytics = true

  query_isolation_options = {
    max_queries           = 2
    max_nodes_per_query   = 2
  }

  instance_type      = "r5d.4xlarge"
  availability_zone  = "us-west-2a"
  platform_jwt_auth  = true
  query_log          = true
  enable_sso         = true

  tags = [
    {
      key   = "environment"
      value = "mow-dev"
    },
    {
      key   = "team"
      value = "dwx"
    }
  ]

  lifecycle {
    ignore_changes = [
      aws_options["spill_to_s3_uri"],
      last_updated,
      node_count,
      status
    ]
  }
}
`, clusterID, databaseCatalogID, name, latestImageVersion),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("cdp_dw_vw_impala.test_impala", "name", name),
					resource.TestCheckResourceAttr("cdp_dw_vw_impala.test_impala", "cluster_id", clusterID),
					resource.TestCheckResourceAttr("cdp_dw_vw_impala.test_impala", "database_catalog_id", databaseCatalogID),
					resource.TestCheckResourceAttr("cdp_dw_vw_impala.test_impala", "image_version", latestImageVersion),
				),
			},
		},
	})
}

func fetchLatestImageVersion(clusterID string) (string, error) {
	cdpClient := cdpacctest.GetCdpClientForAccTest()

	params := &operations.ListLatestVersionsParams{
		Input: &models.ListLatestVersionsRequest{
			ClusterID: &clusterID,
		},
		Context:    context.Background(),
		HTTPClient: http.DefaultClient,
	}

	response, err := cdpClient.Dw.Operations.ListLatestVersions(params)
	if err != nil {
		return "", fmt.Errorf("failed to fetch versions: %w", err)
	}

	if len(response.Payload.IDToLatestVersionsMap) == 0 {
		return "", fmt.Errorf("no versions found for cluster ID: %s", clusterID)
	}

	versions := response.Payload.IDToLatestVersionsMap

	// Returns only one entry for WH
	for _, versionInfo := range versions {
		return versionInfo, nil
	}

	return "", fmt.Errorf("no versions found for cluster ID: %s", clusterID)
}
