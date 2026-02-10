// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package databasecatalog

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
	"github.com/cloudera/terraform-provider-cdp/mocks"
)

var testDatabaseCatalogSchema = schema.Schema{
	MarkdownDescription: "Creates an AWS Data Warehouse database catalog.",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The name of the database catalog.",
		},
		"cluster_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The unique identifier of the cluster.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
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

func NewDwApi(client *mocks.MockDwClientService) *dwDatabaseCatalogResource {
	return &dwDatabaseCatalogResource{
		client: &cdp.Client{
			Dw: &dwclient.Dw{
				Operations: client,
				Transport:  MockTransport{},
			}}}
}

func createRawCatalogResource() tftypes.Value {
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"id":           tftypes.String,
				"name":         tftypes.String,
				"cluster_id":   tftypes.String,
				"last_updated": tftypes.String,
				"status":       tftypes.String,
				"polling_options": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"async":                  tftypes.Bool,
						"polling_timeout":        tftypes.Number,
						"call_failure_threshold": tftypes.Number,
					},
				},
			}}, map[string]tftypes.Value{
			"id":           tftypes.NewValue(tftypes.String, "id"),
			"name":         tftypes.NewValue(tftypes.String, "name"),
			"cluster_id":   tftypes.NewValue(tftypes.String, "cluster-id"),
			"last_updated": tftypes.NewValue(tftypes.String, ""),
			"status":       tftypes.NewValue(tftypes.String, "Accepted"),
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

type DwDatabaseCatalogTestSuite struct {
	suite.Suite
	expectedListResponse *operations.ListDbcsOK
}

func TestDwDatabaseCatalogTestSuite(t *testing.T) {
	suite.Run(t, new(DwDatabaseCatalogTestSuite))
}

func (suite *DwDatabaseCatalogTestSuite) SetupTest() {
	suite.expectedListResponse = &operations.ListDbcsOK{
		Payload: &models.ListDbcsResponse{
			Dbcs: []*models.DbcSummary{
				{
					ID:     "dbc-id",
					Name:   "name",
					Status: "Accepted",
				},
			},
		},
	}
}

func (suite *DwDatabaseCatalogTestSuite) TestDwDatabaseCatalogMetadata() {
	dwApi := NewDwApi(new(mocks.MockDwClientService))
	resp := resource.MetadataResponse{}

	// Function under test
	dwApi.Metadata(
		context.TODO(),
		resource.MetadataRequest{ProviderTypeName: "cdp"},
		&resp,
	)
	suite.Equal("cdp_dw_database_catalog", resp.TypeName)
}

func (suite *DwDatabaseCatalogTestSuite) TestDwDatabaseCatalogSchema() {
	dwApi := NewDwApi(new(mocks.MockDwClientService))
	resp := resource.SchemaResponse{}

	// Function under test
	dwApi.Schema(
		context.TODO(),
		resource.SchemaRequest{},
		&resp,
	)
	suite.Equal(testDatabaseCatalogSchema, resp.Schema)
}

func (suite *DwDatabaseCatalogTestSuite) TestDwDatabaseCatalogCreate_Success() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("ListDbcs", mock.Anything).Return(suite.expectedListResponse, nil)
	dwApi := NewDwApi(client)

	req := resource.CreateRequest{
		Plan: tfsdk.Plan{
			Raw:    createRawCatalogResource(),
			Schema: testDatabaseCatalogSchema,
		},
	}

	resp := &resource.CreateResponse{
		State: tfsdk.State{
			Schema: testDatabaseCatalogSchema,
		},
	}

	// Function under test
	dwApi.Create(ctx, req, resp)
	var result resourceModel
	resp.State.Get(ctx, &result)
	suite.False(resp.Diagnostics.HasError())
	suite.Equal("dbc-id", result.ID.ValueString())
	suite.Equal("cluster-id", result.ClusterID.ValueString())
	suite.Equal("name", result.Name.ValueString())
	suite.Equal("Accepted", result.Status.ValueString())
}

func (suite *DwDatabaseCatalogTestSuite) TestDwDatabaseCatalogCreate_CreationError() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	expectedListResponse := &operations.ListDbcsOK{
		Payload: &models.ListDbcsResponse{
			Dbcs: []*models.DbcSummary{},
		},
	}
	client.On("ListDbcs", mock.Anything).Return(expectedListResponse, nil)
	dwApi := NewDwApi(client)

	req := resource.CreateRequest{
		Plan: tfsdk.Plan{
			Raw:    createRawCatalogResource(),
			Schema: testDatabaseCatalogSchema,
		},
	}

	resp := &resource.CreateResponse{
		State: tfsdk.State{
			Schema: testDatabaseCatalogSchema,
		},
	}

	// Function under test
	dwApi.Create(ctx, req, resp)
	var result resourceModel
	resp.State.Get(ctx, &result)
	suite.True(resp.Diagnostics.HasError())
	suite.Contains(resp.Diagnostics.Errors()[0].Summary(),
		"Error finding Data Warehouse database catalog")
	suite.Contains(resp.Diagnostics.Errors()[0].Detail(),
		"exactly one Data Warehouse database catalog should be deployed for cluster")
}

func (suite *DwDatabaseCatalogTestSuite) TestDwDatabaseCatalogCreate_TooManyDbcsError() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	expectedListResponse := &operations.ListDbcsOK{
		Payload: &models.ListDbcsResponse{
			Dbcs: []*models.DbcSummary{
				{
					ID:     "dbc-id",
					Name:   "name",
					Status: "Accepted",
				},
				{
					ID:     "custom-dbc-id",
					Name:   "custom-name",
					Status: "Accepted",
				},
			},
		},
	}
	client.On("ListDbcs", mock.Anything).Return(expectedListResponse, nil)
	dwApi := NewDwApi(client)

	req := resource.CreateRequest{
		Plan: tfsdk.Plan{
			Raw:    createRawCatalogResource(),
			Schema: testDatabaseCatalogSchema,
		},
	}

	resp := &resource.CreateResponse{
		State: tfsdk.State{
			Schema: testDatabaseCatalogSchema,
		},
	}

	// Function under test
	dwApi.Create(ctx, req, resp)
	var result resourceModel
	resp.State.Get(ctx, &result)
	suite.True(resp.Diagnostics.HasError())
	suite.Contains(resp.Diagnostics.Errors()[0].Summary(), "Error finding Data Warehouse database catalog")
	suite.Contains(resp.Diagnostics.Errors()[0].Detail(), "exactly one Data Warehouse database catalog should be deployed for cluster")
}

func (suite *DwDatabaseCatalogTestSuite) TestDwDatabaseCatalogDeletion_Success() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	dwApi := NewDwApi(client)

	req := resource.DeleteRequest{
		State: tfsdk.State{
			Schema: testDatabaseCatalogSchema,
			Raw:    createRawCatalogResource(),
		},
	}
	resp := &resource.DeleteResponse{
		State: tfsdk.State{
			Schema: testDatabaseCatalogSchema,
		},
	}

	// Function under test
	dwApi.Delete(ctx, req, resp)
	var result resourceModel
	resp.State.Get(ctx, &result)
	suite.False(resp.Diagnostics.HasError())
	suite.Equal("Deleted", result.Status.ValueString())
}

func (suite *DwDatabaseCatalogTestSuite) TestStateRefresh_Success() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("ListDbcs", mock.Anything).Return(
		&operations.ListDbcsOK{
			Payload: &models.ListDbcsResponse{
				Dbcs: []*models.DbcSummary{
					{
						ID:     "dbc-id",
						Name:   "name",
						Status: "Running",
					},
				},
			},
		},
		nil)
	dwApi := NewDwApi(client)

	clusterID := "cluster-id"
	callFailedCount := 0
	callFailureThreshold := 3

	// Function under test
	refresh := dwApi.stateRefresh(ctx, &clusterID, &callFailedCount, callFailureThreshold)
	_, status, err := refresh()
	suite.NoError(err)
	suite.Equal("Running", status)
}

func (suite *DwDatabaseCatalogTestSuite) TestStateRefresh_FailureThresholdReached() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("ListDbcs", mock.Anything).Return(
		&operations.ListDbcsOK{}, fmt.Errorf("unknown error"))
	dwApi := NewDwApi(client)

	clusterID := "cluster-id"
	callFailedCount := 0
	callFailureThreshold := 3

	// Function under test
	refresh := dwApi.stateRefresh(ctx, &clusterID, &callFailedCount, callFailureThreshold)
	var err error
	for i := 0; i <= callFailureThreshold; i++ {
		_, _, err = refresh()
	}
	suite.Error(err, "unknown error")
}

func (suite *DwDatabaseCatalogTestSuite) TestGetDatabaseCatalog_Success() {

}
