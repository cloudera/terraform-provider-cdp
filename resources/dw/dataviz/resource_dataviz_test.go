// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package dataviz

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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	dwclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
	mocks "github.com/cloudera/terraform-provider-cdp/mocks/github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/client/operations"
)

var testDatavizSchema = schema.Schema{
	MarkdownDescription: "Cloudera Data Warehouse (CDW) integrates [Data Visualization](https://docs.cloudera.com/data-warehouse/cloud/managing-warehouses/topics/dw-use-data-visualization.html) for building graphic representations of data, dashboards, and visual applications based on CDW data.",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"cluster_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The id of the CDW Cluster which the Data Visualization is attached to.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The name of the Data Visualization.",
		},

		"image_version": schema.StringAttribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "The version of the Data Visualization.",
		},
		"resource_template": schema.StringAttribute{
			Optional: true,
			Computed: true,
			// TODO vcsomor add validation logic to the allowed types
			MarkdownDescription: "The name of the resource template being used. Available options: reduced, medium, large. Empty means the default resources template will be assigned.",
		},

		"user_groups": schema.ListAttribute{
			Optional:            true,
			ElementType:         types.StringType,
			MarkdownDescription: "List of the LDAP groups which have access to this Data Visualization instance.",
		},
		"admin_groups": schema.ListAttribute{
			Required:            true,
			ElementType:         types.StringType,
			MarkdownDescription: "List of the LDAP groups which can administer this Data Visualization instance. At least one valid group is required.",
		},

		// TODO vcsomor add missing Tags to the API

		"last_updated": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Timestamp of the last Terraform update of the order.",
		},
		"status": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The status of the Data Visualization.",
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
					Default:             int64default.StaticInt64(20),
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

func newDwApi(client *mocks.MockDwClientService) *datavizResource {
	return &datavizResource{
		client: &cdp.Client{
			Dw: &dwclient.Dw{
				Operations: client,
				Transport:  MockTransport{},
			}}}
}

func createRawDatavizResource() tftypes.Value {
	return tftypes.NewValue(
		// schema --------------------------
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"id":         tftypes.String,
				"cluster_id": tftypes.String,
				"name":       tftypes.String,

				"image_version":     tftypes.String,
				"resource_template": tftypes.String,

				"user_groups":  tftypes.List{ElementType: tftypes.String},
				"admin_groups": tftypes.List{ElementType: tftypes.String},

				"last_updated": tftypes.String,
				"status":       tftypes.String,
				"polling_options": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"async":                  tftypes.Bool,
						"polling_timeout":        tftypes.Number,
						"call_failure_threshold": tftypes.Number,
					},
				},
			},
		},

		// value --------------------------
		map[string]tftypes.Value{
			"id":         tftypes.NewValue(tftypes.String, ""),
			"cluster_id": tftypes.NewValue(tftypes.String, "cluster-id"),
			"name":       tftypes.NewValue(tftypes.String, "test-name"),

			"image_version":     tftypes.NewValue(tftypes.String, "test-version"),
			"resource_template": tftypes.NewValue(tftypes.String, "test-template"),

			"user_groups": tftypes.NewValue(tftypes.List{ElementType: tftypes.String},
				[]tftypes.Value{
					tftypes.NewValue(tftypes.String, "user-group1"),
					tftypes.NewValue(tftypes.String, "user-group2"),
				}),
			"admin_groups": tftypes.NewValue(tftypes.List{ElementType: tftypes.String},
				[]tftypes.Value{
					tftypes.NewValue(tftypes.String, "admin-group1"),
					tftypes.NewValue(tftypes.String, "admin-group2"),
				}),
			"last_updated": tftypes.NewValue(tftypes.String, ""),
			"status":       tftypes.NewValue(tftypes.String, "Running"),

			"polling_options": tftypes.NewValue(
				tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"async":                  tftypes.Bool,
						"polling_timeout":        tftypes.Number,
						"call_failure_threshold": tftypes.Number,
					}}, map[string]tftypes.Value{
					"async":                  tftypes.NewValue(tftypes.Bool, true),
					"polling_timeout":        tftypes.NewValue(tftypes.Number, 90),
					"call_failure_threshold": tftypes.NewValue(tftypes.Number, 3),
				}),
		})
}

type DataVizTestSuite struct {
	suite.Suite
}

func TestDataViz(t *testing.T) {
	suite.Run(t, new(DataVizTestSuite))
}

func (suite *DataVizTestSuite) SetupTest() {
}

func (suite *DataVizTestSuite) TestDataVizMetadata() {
	dwApi := newDwApi(new(mocks.MockDwClientService))
	resp := resource.MetadataResponse{}

	// Function under test
	dwApi.Metadata(
		context.TODO(),
		resource.MetadataRequest{ProviderTypeName: "cdp"},
		&resp,
	)
	suite.Equal("cdp_dw_data_visualization", resp.TypeName)
}

func (suite *DataVizTestSuite) TestDatavizSchema() {
	dwApi := newDwApi(new(mocks.MockDwClientService))
	resp := resource.SchemaResponse{}

	// Function under test
	dwApi.Schema(
		context.TODO(),
		resource.SchemaRequest{},
		&resp,
	)
	suite.Equal(testDatavizSchema, resp.Schema)
}

func (suite *DataVizTestSuite) TestDatavizCreate_Success() {
	ctx := context.TODO()

	client := new(mocks.MockDwClientService)
	client.On("CreateDataVisualization", mock.Anything).
		Return(&operations.CreateDataVisualizationOK{
			Payload: &models.CreateDataVisualizationResponse{
				DataVisualizationID: "test-id",
			}}, nil)
	client.On("DescribeDataVisualization", mock.Anything).
		Return(&operations.DescribeDataVisualizationOK{
			Payload: &models.DescribeDataVisualizationResponse{
				DataVisualization: &models.DataVisualizationSummary{
					ID:           "test-id",
					Name:         "test-name",
					ImageVersion: "test-version",
					AdminGroups:  []string{"admin-group1", "admin-group2"},
					UserGroups:   []string{"user-group1", "user-group2"},
				}}}, nil)
	dwApi := newDwApi(client)

	resp := &resource.CreateResponse{
		State: tfsdk.State{
			Schema: testDatavizSchema,
		},
	}

	// Function under test
	dwApi.Create(
		ctx,
		resource.CreateRequest{
			Plan: tfsdk.Plan{
				Raw:    createRawDatavizResource(),
				Schema: testDatavizSchema,
			},
		},
		resp,
	)
	var result resourceModel
	resp.State.Get(ctx, &result)
	suite.False(resp.Diagnostics.HasError())
	suite.Equal("test-id", result.ID.ValueString())
	suite.Equal("cluster-id", result.ClusterID.ValueString())
	suite.Equal("test-name", result.Name.ValueString())
}

func (suite *DataVizTestSuite) TestDatavizCreate_CreationError() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("CreateDataVisualization", mock.Anything).
		Return(nil, fmt.Errorf("create failed"))
	client.On("DescribeDataVisualization", mock.Anything).
		Return(&operations.DescribeDataVisualizationOK{}, nil)

	dwApi := newDwApi(client)

	resp := &resource.CreateResponse{
		State: tfsdk.State{
			Schema: testDatavizSchema,
		},
	}

	// Function under test
	dwApi.Create(ctx, resource.CreateRequest{
		Plan: tfsdk.Plan{
			Raw:    createRawDatavizResource(),
			Schema: testDatavizSchema,
		},
	}, resp)
	var result resourceModel
	resp.State.Get(ctx, &result)
	suite.True(resp.Diagnostics.HasError())

	err0 := resp.Diagnostics.Errors()[0]
	suite.Contains(err0.Summary(), "Error creating Data Visualization")
	suite.Contains(err0.Detail(), "Could not create Data Visualization, unexpected error: create failed")
}

func (suite *DataVizTestSuite) TestDatavizCreate_DescribeError() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("CreateDataVisualization", mock.Anything).
		Return(&operations.CreateDataVisualizationOK{
			Payload: &models.CreateDataVisualizationResponse{
				DataVisualizationID: "test-id",
			}}, nil)
	client.On("DescribeDataVisualization", mock.Anything).
		Return(nil, fmt.Errorf("describe failed"))
	dwApi := newDwApi(client)

	resp := &resource.CreateResponse{
		State: tfsdk.State{
			Schema: testDatavizSchema,
		},
	}

	// Function under test
	dwApi.Create(ctx, resource.CreateRequest{
		Plan: tfsdk.Plan{
			Raw:    createRawDatavizResource(),
			Schema: testDatavizSchema,
		},
	}, resp)
	var result resourceModel
	resp.State.Get(ctx, &result)
	suite.True(resp.Diagnostics.HasError())
	err0 := resp.Diagnostics.Errors()[0]
	suite.Contains(err0.Summary(), "Error creating Data Visualization")
	suite.Contains(err0.Detail(), "Could not describe Data Visualization, unexpected error: describe failed")
}

func (suite *DataVizTestSuite) TestDatavizDeletion_Success() {
	client := new(mocks.MockDwClientService)
	client.On("DeleteDataVisualization", mock.Anything).
		Return(&operations.DeleteDataVisualizationOK{}, nil)
	dwApi := newDwApi(client)

	resp := &resource.DeleteResponse{}

	// Function under test
	dwApi.Delete(context.TODO(), resource.DeleteRequest{
		State: tfsdk.State{
			Schema: testDatavizSchema,
			Raw:    createRawDatavizResource(),
		},
	}, resp)
	suite.False(resp.Diagnostics.HasError())
}

func (suite *DataVizTestSuite) TestDatavizDeletion_ReturnsError() {
	client := new(mocks.MockDwClientService)
	client.On("DeleteDataVisualization", mock.Anything).
		Return(nil, fmt.Errorf("delete failed"))
	dwApi := newDwApi(client)

	resp := &resource.DeleteResponse{}

	// Function under test
	dwApi.Delete(context.TODO(), resource.DeleteRequest{
		State: tfsdk.State{
			Schema: testDatavizSchema,
			Raw:    createRawDatavizResource(),
		},
	}, resp)
	suite.True(resp.Diagnostics.HasError())
}

func (suite *DataVizTestSuite) TestStateRefresh_Success() {
	client := new(mocks.MockDwClientService)
	client.On("DescribeDataVisualization", mock.Anything).Return(
		&operations.DescribeDataVisualizationOK{
			Payload: &models.DescribeDataVisualizationResponse{
				DataVisualization: &models.DataVisualizationSummary{
					ID:     "dataviz-id",
					Status: "Running",
				},
			},
		},
		nil)
	dwApi := newDwApi(client)

	callFailedCount := 0
	callFailureThreshold := 3

	// Function under test
	refresh := dwApi.stateRefresh(
		context.TODO(),
		ptrOf("cluster-id"),
		ptrOf("dataviz-id"),
		&callFailedCount,
		callFailureThreshold,
	)
	_, status, err := refresh()
	suite.NoError(err)
	suite.Equal("Running", status)
}

func (suite *DataVizTestSuite) TestStateRefresh_FailureThresholdReached() {
	client := new(mocks.MockDwClientService)
	client.On("DescribeDataVisualization", mock.Anything).Return(
		nil, fmt.Errorf("unknown error"))
	dwApi := newDwApi(client)

	callFailedCount := 0
	callFailureThreshold := 3

	// Function under test
	refresh := dwApi.stateRefresh(
		context.TODO(),
		ptrOf("cluster-id"),
		ptrOf("dataviz-id"),
		&callFailedCount,
		callFailureThreshold,
	)
	var err error
	for i := 0; i <= callFailureThreshold; i++ {
		_, _, err = refresh()
	}
	suite.Error(err, "unknown error")
}

func TestRetryConfigs(t *testing.T) {
	assert.Equal(
		t,
		&retryStateCfg{
			clusterID: ptrOf("cluster-id"),
			vizID:     ptrOf("dataviz-id"),
			pending:   []string{"Accepted", "Creating", "Created", "Starting"},
			target:    []string{"Running"},
		},
		setupRetryCfg(ptrOf("cluster-id"), ptrOf("dataviz-id")),
	)

	assert.Equal(
		t,
		&retryStateCfg{
			clusterID: ptrOf("cluster-id"),
			vizID:     ptrOf("dataviz-id"),
			pending:   []string{"Deleting", "Running", "Stopping", "Stopped", "Creating", "Created", "Starting", "Updating"},
			target:    []string{"Deleted"},
		},
		teardownRetryCfg(ptrOf("cluster-id"), ptrOf("dataviz-id")),
	)
}

func ptrOf[T any](v T) *T {
	return &v
}
