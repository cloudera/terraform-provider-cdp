// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package dw

import (
	"context"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	dwclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
	mocks "github.com/cloudera/terraform-provider-cdp/mocks/github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/client/operations"
	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

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
			}},
		map[string]tftypes.Value{
			"id":                  tftypes.NewValue(tftypes.String, ""),
			"cluster_id":          tftypes.NewValue(tftypes.String, "cluster-id"),
			"database_catalog_id": tftypes.NewValue(tftypes.String, "database-catalog-id"),
			"name":                tftypes.NewValue(tftypes.String, ""),
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
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	dwApi := NewDwApi(client)

	req := resource.MetadataRequest{ProviderTypeName: "dw"}
	resp := &resource.MetadataResponse{}

	// Function under test
	dwApi.Metadata(ctx, req, resp)
	suite.Equal("dw_vw_hive", resp.TypeName)
}

func (suite *HiveTestSuite) TestHiveSchema() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	dwApi := NewDwApi(client)

	req := resource.SchemaRequest{}
	resp := &resource.SchemaResponse{}

	// Function under test
	dwApi.Schema(ctx, req, resp)
	suite.Equal(hiveSchema, resp.Schema)
}

func (suite *HiveTestSuite) TestHiveCreate_basic() {
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
			Schema: hiveSchema,
		},
	}

	resp := &resource.CreateResponse{
		State: tfsdk.State{
			Schema: hiveSchema,
		},
	}

	// Function under test
	dwApi.Create(ctx, req, resp)
	var result hiveResourceModel
	resp.State.Get(ctx, &result)
	suite.False(resp.Diagnostics.HasError())
	suite.Equal("test-id", result.ID.ValueString())
	suite.Equal("database-catalog-id", result.DbCatalogID.ValueString())
	suite.Equal("cluster-id", result.ClusterID.ValueString())
	suite.Equal("test-name", result.Name.ValueString())
}

func (suite *HiveTestSuite) TestHiveDeletion_basic() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("DeleteVw", mock.Anything).Return(&operations.DeleteVwOK{}, nil)
	dwApi := NewDwApi(client)

	req := resource.DeleteRequest{
		State: tfsdk.State{
			Schema: hiveSchema,
			Raw:    createRawHiveResource(),
		},
	}

	resp := &resource.DeleteResponse{}

	// Function under test
	dwApi.Delete(ctx, req, resp)
	suite.False(resp.Diagnostics.HasError())
}
