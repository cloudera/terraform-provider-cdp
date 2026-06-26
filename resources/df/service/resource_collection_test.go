// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package service

import (
	"context"
	"errors"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	dfclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/df/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/df/client/operations"
	dfmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/df/models"
	"github.com/cloudera/terraform-provider-cdp/mocks"
)

type MockTransport struct {
	runtime.ClientTransport
}

func NewMockDf(mockClient *mocks.MockDfClientService) *dfclient.Df {
	return &dfclient.Df{
		Operations: mockClient,
		Transport:  &MockTransport{},
	}
}

func createRawCollectionResource(id string) tftypes.Value {
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"id":          tftypes.String,
				"name":        tftypes.String,
				"description": tftypes.String,
				"crn":         tftypes.String,
			},
		},
		map[string]tftypes.Value{
			"id":          tftypes.NewValue(tftypes.String, id),
			"name":        tftypes.NewValue(tftypes.String, "test-collection"),
			"description": tftypes.NewValue(tftypes.String, "test description"),
			"crn":         tftypes.NewValue(tftypes.String, ""),
		},
	)
}

func collectionSchema() resource.SchemaResponse {
	r := &dfCollectionResource{}
	resp := &resource.SchemaResponse{}
	r.Schema(context.TODO(), resource.SchemaRequest{}, resp)
	return *resp
}

func TestCreateCollection_OK(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDfClientService(t)

	crn := "crn:cdp:df:us-west-1:tenant:collection:test-id"
	mockClient.On("CreateCollection", mock.Anything).Return(
		&operations.CreateCollectionOK{
			Payload: &dfmodels.CreateCollectionResponse{
				CatalogCollection: &dfmodels.CatalogCollection{
					Crn:  crn,
					Name: "test-collection",
				},
			},
		}, nil)

	r := &dfCollectionResource{client: &cdp.Client{Df: NewMockDf(mockClient)}}
	s := collectionSchema()

	req := resource.CreateRequest{
		Plan: tfsdk.Plan{Raw: createRawCollectionResource(""), Schema: s.Schema},
	}
	resp := &resource.CreateResponse{
		State: tfsdk.State{Schema: s.Schema},
	}

	r.Create(ctx, req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	var state collectionModel
	resp.State.Get(ctx, &state)
	assert.Equal(t, crn, state.Crn.ValueString())
	assert.Equal(t, crn, state.ID.ValueString())
}

func TestCreateCollection_Error(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDfClientService(t)

	mockClient.On("CreateCollection", mock.Anything).Return(nil, errors.New("service unavailable"))

	r := &dfCollectionResource{client: &cdp.Client{Df: NewMockDf(mockClient)}}
	s := collectionSchema()

	req := resource.CreateRequest{
		Plan: tfsdk.Plan{Raw: createRawCollectionResource(""), Schema: s.Schema},
	}
	resp := &resource.CreateResponse{
		State: tfsdk.State{Schema: s.Schema},
	}

	r.Create(ctx, req, resp)

	assert.True(t, resp.Diagnostics.HasError())
	assert.Contains(t, resp.Diagnostics.Errors()[0].Detail(), "service unavailable")
}

func TestReadCollection_OK(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDfClientService(t)

	crn := "crn:cdp:df:us-west-1:tenant:collection:test-id"
	mockClient.On("DescribeCollection", mock.Anything).Return(
		&operations.DescribeCollectionOK{
			Payload: &dfmodels.DescribeCollectionResponse{
				CatalogCollection: &dfmodels.CatalogCollection{
					Crn:         crn,
					Name:        "test-collection",
					Description: "test description",
				},
			},
		}, nil)

	r := &dfCollectionResource{client: &cdp.Client{Df: NewMockDf(mockClient)}}
	s := collectionSchema()

	raw := createRawCollectionResource(crn)
	raw, _ = tftypes.Transform(raw, func(path *tftypes.AttributePath, v tftypes.Value) (tftypes.Value, error) {
		if path.Equal(tftypes.NewAttributePath().WithAttributeName("crn")) {
			return tftypes.NewValue(tftypes.String, crn), nil
		}
		return v, nil
	})

	req := resource.ReadRequest{
		State: tfsdk.State{Raw: raw, Schema: s.Schema},
	}
	resp := &resource.ReadResponse{
		State: tfsdk.State{Schema: s.Schema},
	}

	r.Read(ctx, req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	var state collectionModel
	resp.State.Get(ctx, &state)
	assert.Equal(t, crn, state.Crn.ValueString())
	assert.Equal(t, "test-collection", state.Name.ValueString())
}

func TestReadCollection_NotFound(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDfClientService(t)

	mockClient.On("DescribeCollection", mock.Anything).Return(nil, errors.New("NOT_FOUND"))

	r := &dfCollectionResource{client: &cdp.Client{Df: NewMockDf(mockClient)}}
	s := collectionSchema()

	crn := "crn:cdp:df:us-west-1:tenant:collection:test-id"
	req := resource.ReadRequest{
		State: tfsdk.State{Raw: createRawCollectionResource(crn), Schema: s.Schema},
	}
	resp := &resource.ReadResponse{
		State: tfsdk.State{Schema: s.Schema},
	}

	r.Read(ctx, req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	// State should be removed
	assert.True(t, resp.State.Raw.IsNull())
}

func TestDeleteCollection_OK(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDfClientService(t)

	mockClient.On("DeleteCollection", mock.Anything).Return(&operations.DeleteCollectionOK{}, nil)

	r := &dfCollectionResource{client: &cdp.Client{Df: NewMockDf(mockClient)}}
	s := collectionSchema()

	crn := "crn:cdp:df:us-west-1:tenant:collection:test-id"
	req := resource.DeleteRequest{
		State: tfsdk.State{Raw: createRawCollectionResource(crn), Schema: s.Schema},
	}
	resp := &resource.DeleteResponse{
		State: tfsdk.State{Schema: s.Schema},
	}

	r.Delete(ctx, req, resp)

	assert.False(t, resp.Diagnostics.HasError())
}
