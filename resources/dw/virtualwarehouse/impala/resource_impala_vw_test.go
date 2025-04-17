// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package impala

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	dwclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
	mocks "github.com/cloudera/terraform-provider-cdp/mocks/github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/client/operations"
)

var testImpalaSchema = schema.Schema{
	MarkdownDescription: "A Impala Virtual Warehouse is service which is able to run low-latency SQL queries.",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"cluster_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The id of the CDW Cluster which the Impala Virtual Warehouse is attached to.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"database_catalog_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The id of the Database Catalog which the Impala Virtual Warehouse is attached to.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The name of the Impala Virtual Warehouse.",
		},
		"last_updated": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Timestamp of the last Terraform update of the order.",
		},
		"status": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The status of the database catalog.",
		},
		"image_version": schema.StringAttribute{
			Computed:            true,
			Optional:            true,
			MarkdownDescription: "Image version of the impala.",
		},
		"instance_type": schema.StringAttribute{
			Computed:            true,
			Optional:            true,
			MarkdownDescription: "The instance type for the Impala Virtual Warehouse.",
		},
		"tshirt_size": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "T-shirt size of Impala.",
		},
		"node_count": schema.Int32Attribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "Node count of Impala.",
		},
		"availability_zone": schema.StringAttribute{
			Computed:            true,
			Optional:            true,
			MarkdownDescription: "The availability zone for the Impala Virtual Warehouse.",
		},
		"enable_unified_analytics": schema.BoolAttribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "Flag to enable unified analytics.",
		},
		"aws_options": schema.SingleNestedAttribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "Impala-specific configuration options.",
			Attributes: map[string]schema.Attribute{
				"scratch_space_limit": schema.Int32Attribute{
					Optional:            true,
					Computed:            true,
					MarkdownDescription: "Defines the limit for scratch space in GiB needed by Impala for spilling queries. Valid values depend on the platform (AWS or Azure). If set, 'spillToS3Uri' cannot be set.",
				},
				"spill_to_s3_uri": schema.StringAttribute{
					Optional:            true,
					Computed:            true,
					MarkdownDescription: "Set S3 URI in 's3://bucket/path' format to enable spilling to S3. If set, 'scratchSpaceLimit' cannot be set. Not supported on Azure.",
				},
			},
		},
		"ha_settings": schema.SingleNestedAttribute{
			Computed:            true,
			Optional:            true,
			MarkdownDescription: "High availability settings for Impala.",
			Attributes: map[string]schema.Attribute{
				"high_availability_mode": schema.StringAttribute{
					Optional:            true,
					Computed:            true,
					MarkdownDescription: "High Availability mode: DISABLED, ACTIVE_PASSIVE, or ACTIVE_ACTIVE.",
				},
				"enable_shutdown_of_coordinator": schema.BoolAttribute{
					Optional:            true,
					Computed:            true,
					MarkdownDescription: "Enables the shutdown of the coordinator.",
				},
				"shutdown_of_coordinator_delay_secs": schema.Int32Attribute{
					Optional:            true,
					Computed:            true,
					MarkdownDescription: "Delay in seconds before shutting down the coordinator.",
				},
				"num_of_active_coordinators": schema.Int32Attribute{
					Optional:            true,
					Computed:            true,
					MarkdownDescription: "Number of active coordinators.",
				},
				"enable_catalog_high_availability": schema.BoolAttribute{
					Optional:            true,
					Computed:            true,
					MarkdownDescription: "Enables high availability for Impala catalog.",
				},
				"enable_statestore_high_availability": schema.BoolAttribute{
					Optional:            true,
					Computed:            true,
					MarkdownDescription: "Enables high availability for Impala Statestore.",
				},
			},
		},
		"autoscaling": schema.SingleNestedAttribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "Autoscaling configuration options.",
			Attributes: map[string]schema.Attribute{
				"min_clusters": schema.Int32Attribute{
					Computed:            true,
					Optional:            true,
					MarkdownDescription: "Minimum number of available compute groups. Default: 0.",
				},
				"max_clusters": schema.Int32Attribute{
					Computed:            true,
					Optional:            true,
					MarkdownDescription: "Maximum number of available compute groups. Default: 0.",
				},
				"disable_auto_suspend": schema.BoolAttribute{
					Computed:            true,
					Optional:            true,
					MarkdownDescription: "Disable auto-suspend for the Virtual Warehouse.",
				},
				"auto_suspend_timeout_seconds": schema.Int32Attribute{
					Computed:            true,
					Optional:            true,
					MarkdownDescription: "Threshold for auto-suspend in seconds.",
				},
				"scale_up_delay_seconds": schema.Int32Attribute{
					Computed:            true,
					Optional:            true,
					MarkdownDescription: "Scale-up threshold in seconds for Impala.",
				},
				"scale_down_delay_seconds": schema.Int32Attribute{
					Computed:            true,
					Optional:            true,
					MarkdownDescription: "Scale-down threshold in seconds for Impala.",
				},
			},
		},
		"query_isolation_options": schema.SingleNestedAttribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "Query isolation options for Impala.",
			Attributes: map[string]schema.Attribute{
				"max_queries": schema.Int32Attribute{
					Optional:            true,
					Computed:            true,
					MarkdownDescription: "Maximum number of queries for isolation. Default: 0 disables isolation.",
				},
				"max_nodes_per_query": schema.Int32Attribute{
					Computed:            true,
					Optional:            true,
					MarkdownDescription: "Maximum number of nodes per query for isolation. Default: 0 disables isolation.",
				},
			},
		},
		"tags": schema.ListNestedAttribute{
			Optional:            true,
			MarkdownDescription: "Tags associated with the Impala Virtual Warehouse.",
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"key": schema.StringAttribute{
						Required: true,
					},
					"value": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
		"enable_sso": schema.BoolAttribute{
			Computed:            true,
			Optional:            true,
			MarkdownDescription: "Enable sso for Impala VWH",
		},
		"platform_jwt_auth": schema.BoolAttribute{
			Optional:            true,
			MarkdownDescription: "Platform JWT authentication flag.",
		},
		"query_log": schema.BoolAttribute{
			Computed:            true,
			Optional:            true,
			MarkdownDescription: "Enable or disable Impala query logging.",
		},
		"polling_options": schema.SingleNestedAttribute{
			MarkdownDescription: "Polling related configuration options that could specify various values that will be used during CDP resource creation.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"async": schema.BoolAttribute{
					MarkdownDescription: "Boolean value that specifies if Terraform should wait for resource creation/deletion.",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
					PlanModifiers: []planmodifier.Bool{
						boolplanmodifier.UseStateForUnknown(),
					},
				},
				"polling_timeout": schema.Int64Attribute{
					MarkdownDescription: "Timeout value in minutes that specifies for how long should the polling go for resource creation/deletion.",
					Default:             int64default.StaticInt64(40),
					Computed:            true,
					Optional:            true,
				},
				"call_failure_threshold": schema.Int64Attribute{
					MarkdownDescription: "Threshold value that specifies how many times should a single call failure happen before giving up the polling.",
					Default:             int64default.StaticInt64(3),
					Computed:            true,
					Optional:            true,
				},
			},
		},
	},
}

type MockTransport struct {
	runtime.ClientTransport
}

func newDwApi(client *mocks.MockDwClientService) *impalaResource {
	return &impalaResource{
		client: &cdp.Client{
			Dw: &dwclient.Dw{
				Operations: client,
				Transport:  MockTransport{},
			}}}
}

func createRawImpalaResource() tftypes.Value {
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"id":                  tftypes.String,
				"cluster_id":          tftypes.String,
				"database_catalog_id": tftypes.String,
				"name":                tftypes.String,
				"last_updated":        tftypes.String,
				"status":              tftypes.String,
				"image_version":       tftypes.String,
				"tshirt_size":         tftypes.String,
				"autoscaling": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"auto_suspend_timeout_seconds": tftypes.Number,
						"disable_auto_suspend":         tftypes.Bool,
						"scale_down_delay_seconds":     tftypes.Number,
						"scale_up_delay_seconds":       tftypes.Number,
						"max_clusters":                 tftypes.Number,
						"min_clusters":                 tftypes.Number,
					},
				},
				"aws_options": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"scratch_space_limit": tftypes.Number,
						"spill_to_s3_uri":     tftypes.String,
					},
				},
				"ha_settings": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"high_availability_mode":              tftypes.String,
						"enable_shutdown_of_coordinator":      tftypes.Bool,
						"shutdown_of_coordinator_delay_secs":  tftypes.Number,
						"num_of_active_coordinators":          tftypes.Number,
						"enable_catalog_high_availability":    tftypes.Bool,
						"enable_statestore_high_availability": tftypes.Bool,
					},
				},
				"query_isolation_options": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"max_queries":         tftypes.Number,
						"max_nodes_per_query": tftypes.Number,
					},
				},
				"instance_type":            tftypes.String,
				"availability_zone":        tftypes.String,
				"platform_jwt_auth":        tftypes.Bool,
				"enable_unified_analytics": tftypes.Bool,
				"node_count":               tftypes.Number,
				"query_log":                tftypes.Bool,
				"tags": tftypes.List{
					ElementType: tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"key":   tftypes.String,
							"value": tftypes.String,
						},
					},
				},
				"polling_options": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"async":                  tftypes.Bool,
						"polling_timeout":        tftypes.Number,
						"call_failure_threshold": tftypes.Number,
					},
				},
				"enable_sso": tftypes.Bool,
			},
		},
		map[string]tftypes.Value{
			"id":                  tftypes.NewValue(tftypes.String, ""),
			"cluster_id":          tftypes.NewValue(tftypes.String, "cluster-id"),
			"database_catalog_id": tftypes.NewValue(tftypes.String, "database-catalog-id"),
			"name":                tftypes.NewValue(tftypes.String, ""),
			"last_updated":        tftypes.NewValue(tftypes.String, ""),
			"status":              tftypes.NewValue(tftypes.String, "Running"),
			"image_version":       tftypes.NewValue(tftypes.String, nil),
			"tshirt_size":         tftypes.NewValue(tftypes.String, "xsmall"),
			"autoscaling": tftypes.NewValue(
				tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"auto_suspend_timeout_seconds": tftypes.Number,
						"disable_auto_suspend":         tftypes.Bool,
						"scale_down_delay_seconds":     tftypes.Number,
						"scale_up_delay_seconds":       tftypes.Number,
						"max_clusters":                 tftypes.Number,
						"min_clusters":                 tftypes.Number,
					},
				},
				map[string]tftypes.Value{
					"auto_suspend_timeout_seconds": tftypes.NewValue(tftypes.Number, 360),
					"disable_auto_suspend":         tftypes.NewValue(tftypes.Bool, false),
					"scale_down_delay_seconds":     tftypes.NewValue(tftypes.Number, 360),
					"scale_up_delay_seconds":       tftypes.NewValue(tftypes.Number, 40),
					"max_clusters":                 tftypes.NewValue(tftypes.Number, 6),
					"min_clusters":                 tftypes.NewValue(tftypes.Number, 4),
				},
			),
			"aws_options": tftypes.NewValue(
				tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"scratch_space_limit": tftypes.Number,
						"spill_to_s3_uri":     tftypes.String,
					},
				},
				map[string]tftypes.Value{
					"scratch_space_limit": tftypes.NewValue(tftypes.Number, 634),
					"spill_to_s3_uri":     tftypes.NewValue(tftypes.String, ""),
				},
			),
			"ha_settings": tftypes.NewValue(
				tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"high_availability_mode":              tftypes.String,
						"enable_shutdown_of_coordinator":      tftypes.Bool,
						"shutdown_of_coordinator_delay_secs":  tftypes.Number,
						"num_of_active_coordinators":          tftypes.Number,
						"enable_catalog_high_availability":    tftypes.Bool,
						"enable_statestore_high_availability": tftypes.Bool,
					},
				},
				map[string]tftypes.Value{
					"high_availability_mode":              tftypes.NewValue(tftypes.String, "ACTIVE_PASSIVE"),
					"enable_shutdown_of_coordinator":      tftypes.NewValue(tftypes.Bool, false),
					"shutdown_of_coordinator_delay_secs":  tftypes.NewValue(tftypes.Number, 360),
					"num_of_active_coordinators":          tftypes.NewValue(tftypes.Number, 2),
					"enable_catalog_high_availability":    tftypes.NewValue(tftypes.Bool, false),
					"enable_statestore_high_availability": tftypes.NewValue(tftypes.Bool, false),
				},
			),
			"query_isolation_options": tftypes.NewValue(
				tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"max_queries":         tftypes.Number,
						"max_nodes_per_query": tftypes.Number,
					},
				},

				map[string]tftypes.Value{
					"max_queries":         tftypes.NewValue(tftypes.Number, 2),
					"max_nodes_per_query": tftypes.NewValue(tftypes.Number, 2),
				},
			),
			"enable_sso":               tftypes.NewValue(tftypes.Bool, true),
			"instance_type":            tftypes.NewValue(tftypes.String, "r5d.4xlarge"),
			"availability_zone":        tftypes.NewValue(tftypes.String, "us-west-2a"),
			"platform_jwt_auth":        tftypes.NewValue(tftypes.Bool, true),
			"query_log":                tftypes.NewValue(tftypes.Bool, true),
			"enable_unified_analytics": tftypes.NewValue(tftypes.Bool, false),
			"node_count":               tftypes.NewValue(tftypes.Number, 2),
			"tags": tftypes.NewValue(
				tftypes.List{
					ElementType: tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"key":   tftypes.String,
							"value": tftypes.String,
						},
					},
				},
				[]tftypes.Value{
					tftypes.NewValue(
						tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"key":   tftypes.String,
								"value": tftypes.String,
							},
						},
						map[string]tftypes.Value{
							"key":   tftypes.NewValue(tftypes.String, "environment"),
							"value": tftypes.NewValue(tftypes.String, "mow-dev"),
						},
					),
					tftypes.NewValue(
						tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"key":   tftypes.String,
								"value": tftypes.String,
							},
						},
						map[string]tftypes.Value{
							"key":   tftypes.NewValue(tftypes.String, "team"),
							"value": tftypes.NewValue(tftypes.String, "dwx"),
						},
					),
				},
			),
			"polling_options": tftypes.NewValue(
				tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"async":                  tftypes.Bool,
						"polling_timeout":        tftypes.Number,
						"call_failure_threshold": tftypes.Number,
					},
				},
				map[string]tftypes.Value{
					"async":                  tftypes.NewValue(tftypes.Bool, true),
					"polling_timeout":        tftypes.NewValue(tftypes.Number, 600),
					"call_failure_threshold": tftypes.NewValue(tftypes.Number, 3),
				},
			),
		},
	)
}

type ImpalaTestSuite struct {
	suite.Suite
	expectedCreateResponse *operations.CreateVwOK
}

func TestImpala(t *testing.T) {
	suite.Run(t, new(ImpalaTestSuite))
}

func (suite *ImpalaTestSuite) SetupTest() {
	suite.expectedCreateResponse = &operations.CreateVwOK{Payload: &models.CreateVwResponse{
		VwID: "test-id",
	}}

}

func (suite *ImpalaTestSuite) TestImpalaMetadata() {
	dwApi := newDwApi(new(mocks.MockDwClientService))
	resp := resource.MetadataResponse{}

	// Function under test
	dwApi.Metadata(
		context.TODO(),
		resource.MetadataRequest{ProviderTypeName: "dw"},
		&resp,
	)
	suite.Equal("dw_dw_vw_impala", resp.TypeName)
}

func (suite *ImpalaTestSuite) TestImpalaSchema() {
	dwApi := newDwApi(new(mocks.MockDwClientService))
	resp := resource.SchemaResponse{}

	// Function under test
	dwApi.Schema(
		context.TODO(),
		resource.SchemaRequest{},
		&resp,
	)
	suite.Equal(testImpalaSchema, resp.Schema)
}

func (suite *ImpalaTestSuite) TestImpalaCreate_Success() {
	ctx := context.TODO()
	expectedDescribeResponse := &operations.DescribeVwOK{
		Payload: &models.DescribeVwResponse{
			Vw: &models.VwSummary{
				ID:     "test-id",
				DbcID:  "database-catalog-id",
				Name:   "test-name",
				VwType: models.VwTypeImpala,
			}}}

	client := new(mocks.MockDwClientService)
	client.On("CreateVw", mock.Anything).Return(suite.expectedCreateResponse, nil)
	client.On("DescribeVw", mock.Anything).Return(expectedDescribeResponse, nil)
	dwApi := newDwApi(client)

	req := resource.CreateRequest{
		Plan: tfsdk.Plan{
			Raw:    createRawImpalaResource(),
			Schema: testImpalaSchema,
		},
	}

	resp := &resource.CreateResponse{
		State: tfsdk.State{
			Schema: testImpalaSchema,
		},
	}

	// Function under test
	dwApi.Create(ctx, req, resp)
	var result resourceModel
	resp.State.Get(ctx, &result)
	suite.False(resp.Diagnostics.HasError())
	suite.Equal("test-id", result.ID.ValueString())
	suite.Equal("database-catalog-id", result.DatabaseCatalogID.ValueString())
	suite.Equal("cluster-id", result.ClusterID.ValueString())
	suite.Equal("test-name", result.Name.ValueString())
}

func (suite *ImpalaTestSuite) TestImpalaCreate_CreationError() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("CreateVw", mock.Anything).Return(&operations.CreateVwOK{}, fmt.Errorf("create error"))
	dwApi := newDwApi(client)

	req := resource.CreateRequest{
		Plan: tfsdk.Plan{
			Raw:    createRawImpalaResource(),
			Schema: testImpalaSchema,
		},
	}

	resp := &resource.CreateResponse{
		State: tfsdk.State{
			Schema: testImpalaSchema,
		},
	}

	// Function under test
	dwApi.Create(ctx, req, resp)
	suite.True(resp.Diagnostics.HasError())
	suite.Contains(resp.Diagnostics.Errors()[0].Detail(), "create error")
}

func (suite *ImpalaTestSuite) TestImpalaDeletion_Success() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("DeleteVw", mock.Anything).Return(&operations.DeleteVwOK{}, nil)
	dwApi := newDwApi(client)

	req := resource.DeleteRequest{
		State: tfsdk.State{
			Schema: testImpalaSchema,
			Raw:    createRawImpalaResource(),
		},
	}
	resp := resource.DeleteResponse{}

	// Function under test
	dwApi.Delete(ctx, req, &resp)
	suite.False(resp.Diagnostics.HasError())
}

func (suite *ImpalaTestSuite) TestImpalaDeletion_ReturnsError() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("DeleteVw", mock.Anything).Return(&operations.DeleteVwOK{}, fmt.Errorf("deletion error"))
	dwApi := newDwApi(client)

	req := resource.DeleteRequest{
		State: tfsdk.State{
			Schema: testImpalaSchema,
			Raw:    createRawImpalaResource(),
		},
	}
	resp := resource.DeleteResponse{}

	// Function under test
	dwApi.Delete(ctx, req, &resp)
	suite.True(resp.Diagnostics.HasError())
	suite.Contains(resp.Diagnostics.Errors()[0].Detail(), "deletion error")
}

func (suite *ImpalaTestSuite) TestStateRefresh_Success() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("DescribeVw", mock.Anything).Return(&operations.DescribeVwOK{
		Payload: &models.DescribeVwResponse{
			Vw: &models.VwSummary{
				ID:     "impala-id",
				Status: "Running",
			},
		}}, nil)
	dwApi := newDwApi(client)

	clusterID := "cluster-id"
	vwID := "impala-id"
	callFailedCount := 0
	callFailureThreshold := 3

	// Function under test
	refresh := dwApi.stateRefresh(ctx, &clusterID, &vwID, &callFailedCount, callFailureThreshold)
	_, status, err := refresh()
	suite.NoError(err)
	suite.Equal("Running", status)
}

func (suite *ImpalaTestSuite) TestStateRefresh_FailureThresholdReached() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("DescribeVw", mock.Anything).Return(
		&operations.DescribeVwOK{}, fmt.Errorf("unknown error"))
	dwApi := newDwApi(client)

	clusterID := "cluster-id"
	vwID := "impala-id"
	callFailedCount := 0
	callFailureThreshold := 3

	// Function under test
	refresh := dwApi.stateRefresh(ctx, &clusterID, &vwID, &callFailedCount, callFailureThreshold)
	var err error
	for i := 0; i <= callFailureThreshold; i++ {
		_, _, err = refresh()
	}
	suite.Error(err, "unknown error")
}
