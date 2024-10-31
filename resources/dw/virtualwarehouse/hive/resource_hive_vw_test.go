// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package hive

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

var testHiveSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"cluster_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The id of the CDW Cluster which the Hive Virtual Warehouse is attached to.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"database_catalog_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The id of the Database Catalog which the Hive Virtual Warehouse is attached to.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The name of the Hive Virtual Warehouse.",
		},
		"last_updated": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Timestamp of the last Terraform update of the order.",
		},
		"status": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The status of the database catalog.",
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

func NewDwApi(client *mocks.MockDwClientService) *hiveResource {
	return &hiveResource{
		client: &cdp.Client{
			Dw: &dwclient.Dw{
				Operations: client,
				Transport:  MockTransport{},
			}}}
}

func createRawHiveResource() tftypes.Value {
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"id":                  tftypes.String,
				"cluster_id":          tftypes.String,
				"database_catalog_id": tftypes.String,
				"name":                tftypes.String,
				"last_updated":        tftypes.String,
				"status":              tftypes.String,
				"polling_options": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"async":                  tftypes.Bool,
						"polling_timeout":        tftypes.Number,
						"call_failure_threshold": tftypes.Number,
					},
				},
			}},
		map[string]tftypes.Value{
			"id":                  tftypes.NewValue(tftypes.String, ""),
			"cluster_id":          tftypes.NewValue(tftypes.String, "cluster-id"),
			"database_catalog_id": tftypes.NewValue(tftypes.String, "database-catalog-id"),
			"name":                tftypes.NewValue(tftypes.String, ""),
			"last_updated":        tftypes.NewValue(tftypes.String, ""),
			"status":              tftypes.NewValue(tftypes.String, "Running"),
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

type HiveTestSuite struct {
	suite.Suite
	expectedCreateResponse *operations.CreateVwOK
}

func TestHive(t *testing.T) {
	suite.Run(t, new(HiveTestSuite))
}

func (suite *HiveTestSuite) SetupTest() {
	suite.expectedCreateResponse = &operations.CreateVwOK{Payload: &models.CreateVwResponse{
		VwID: "test-id",
	}}

}

func (suite *HiveTestSuite) TestHiveMetadata() {
	dwApi := NewDwApi(new(mocks.MockDwClientService))
	resp := resource.MetadataResponse{}

	// Function under test
	dwApi.Metadata(
		context.TODO(),
		resource.MetadataRequest{ProviderTypeName: "dw"},
		&resp,
	)
	suite.Equal("dw_vw_hive", resp.TypeName)
}

func (suite *HiveTestSuite) TestHiveSchema() {
	dwApi := NewDwApi(new(mocks.MockDwClientService))
	resp := resource.SchemaResponse{}

	// Function under test
	dwApi.Schema(
		context.TODO(),
		resource.SchemaRequest{},
		&resp,
	)
	suite.Equal(testHiveSchema, resp.Schema)
}

func (suite *HiveTestSuite) TestHiveCreate_Success() {
	ctx := context.TODO()
	expectedDescribeResponse := &operations.DescribeVwOK{
		Payload: &models.DescribeVwResponse{
			Vw: &models.VwSummary{
				ID:     "test-id",
				DbcID:  "database-catalog-id",
				Name:   "test-name",
				VwType: models.VwTypeHive,
			}}}

	client := new(mocks.MockDwClientService)
	client.On("CreateVw", mock.Anything).Return(suite.expectedCreateResponse, nil)
	client.On("DescribeVw", mock.Anything).Return(expectedDescribeResponse, nil)
	dwApi := NewDwApi(client)

	req := resource.CreateRequest{
		Plan: tfsdk.Plan{
			Raw:    createRawHiveResource(),
			Schema: testHiveSchema,
		},
	}

	resp := &resource.CreateResponse{
		State: tfsdk.State{
			Schema: testHiveSchema,
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

func (suite *HiveTestSuite) TestHiveCreate_CreationError() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("CreateVw", mock.Anything).Return(&operations.CreateVwOK{}, fmt.Errorf("create failed"))
	client.On("DescribeVw", mock.Anything).Return(&operations.DescribeVwOK{}, nil)
	dwApi := NewDwApi(client)

	req := resource.CreateRequest{
		Plan: tfsdk.Plan{
			Raw:    createRawHiveResource(),
			Schema: testHiveSchema,
		},
	}

	resp := &resource.CreateResponse{
		State: tfsdk.State{
			Schema: testHiveSchema,
		},
	}

	// Function under test
	dwApi.Create(ctx, req, resp)
	var result resourceModel
	resp.State.Get(ctx, &result)
	suite.True(resp.Diagnostics.HasError())
	suite.Contains(resp.Diagnostics.Errors()[0].Summary(), "Error creating hive virtual warehouse")
	suite.Contains(resp.Diagnostics.Errors()[0].Detail(), "Could not create hive")
}

func (suite *HiveTestSuite) TestHiveCreate_DescribeError() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("CreateVw", mock.Anything).Return(suite.expectedCreateResponse, nil)
	client.On("DescribeVw", mock.Anything).Return(&operations.DescribeVwOK{}, fmt.Errorf("describe failed"))
	dwApi := NewDwApi(client)

	req := resource.CreateRequest{
		Plan: tfsdk.Plan{
			Raw:    createRawHiveResource(),
			Schema: testHiveSchema,
		},
	}

	resp := &resource.CreateResponse{
		State: tfsdk.State{
			Schema: testHiveSchema,
		},
	}

	// Function under test
	dwApi.Create(ctx, req, resp)
	var result resourceModel
	resp.State.Get(ctx, &result)
	suite.True(resp.Diagnostics.HasError())
	suite.Contains(resp.Diagnostics.Errors()[0].Summary(), "Error creating hive virtual warehouse")
	suite.Contains(resp.Diagnostics.Errors()[0].Detail(), "Could not describe hive")
}

func (suite *HiveTestSuite) TestHiveDeletion_Success() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("DeleteVw", mock.Anything).Return(&operations.DeleteVwOK{}, nil)
	dwApi := NewDwApi(client)

	req := resource.DeleteRequest{
		State: tfsdk.State{
			Schema: testHiveSchema,
			Raw:    createRawHiveResource(),
		},
	}

	resp := &resource.DeleteResponse{}

	// Function under test
	dwApi.Delete(ctx, req, resp)
	suite.False(resp.Diagnostics.HasError())
}

func (suite *HiveTestSuite) TestHiveDeletion_ReturnsError() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("DeleteVw", mock.Anything).Return(&operations.DeleteVwOK{}, fmt.Errorf("delete failed"))
	dwApi := NewDwApi(client)

	req := resource.DeleteRequest{
		State: tfsdk.State{
			Schema: testHiveSchema,
			Raw:    createRawHiveResource(),
		},
	}

	resp := &resource.DeleteResponse{}

	// Function under test
	dwApi.Delete(ctx, req, resp)
	suite.True(resp.Diagnostics.HasError())
}

func (suite *HiveTestSuite) TestStateRefresh_Success() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("DescribeVw", mock.Anything).Return(
		&operations.DescribeVwOK{
			Payload: &models.DescribeVwResponse{
				Vw: &models.VwSummary{
					ID:     "hive-id",
					Status: "Running",
				},
			},
		},
		nil)
	dwApi := NewDwApi(client)

	clusterID := "cluster-id"
	vwID := "hive-id"
	callFailedCount := 0
	callFailureThreshold := 3

	// Function under test
	refresh := dwApi.stateRefresh(ctx, &clusterID, &vwID, &callFailedCount, callFailureThreshold)
	_, status, err := refresh()
	suite.NoError(err)
	suite.Equal("Running", status)
}

func (suite *HiveTestSuite) TestStateRefresh_FailureThresholdReached() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("DescribeVw", mock.Anything).Return(
		&operations.DescribeVwOK{}, fmt.Errorf("unknown error"))
	dwApi := NewDwApi(client)

	clusterID := "cluster-id"
	vwID := "hive-id"
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
