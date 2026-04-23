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

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/df/client/operations"
	dfmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/df/models"
	"github.com/cloudera/terraform-provider-cdp/mocks"
)

func createRawProjectResource(id string) tftypes.Value {
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"id":          tftypes.String,
				"name":        tftypes.String,
				"service_crn": tftypes.String,
				"description": tftypes.String,
				"crn":         tftypes.String,
				"revision":    tftypes.Number,
			},
		},
		map[string]tftypes.Value{
			"id":          tftypes.NewValue(tftypes.String, id),
			"name":        tftypes.NewValue(tftypes.String, "test-project"),
			"service_crn": tftypes.NewValue(tftypes.String, "crn:cdp:df:us-west-1:tenant:service:svc-id"),
			"description": tftypes.NewValue(tftypes.String, "test description"),
			"crn":         tftypes.NewValue(tftypes.String, ""),
			"revision":    tftypes.NewValue(tftypes.Number, 0),
		},
	)
}

func projectSchema() resource.SchemaResponse {
	r := &dfProjectResource{}
	resp := &resource.SchemaResponse{}
	r.Schema(context.TODO(), resource.SchemaRequest{}, resp)
	return *resp
}

func TestCreateProject_OK(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDfClientService(t)

	crn := "crn:cdp:df:us-west-1:tenant:project:test-id"
	name := "test-project"
	var revision int32 = 1
	mockClient.On("CreateProject", mock.Anything).Return(
		&operations.CreateProjectOK{
			Payload: &dfmodels.CreateProjectResponse{
				Project: &dfmodels.Project{
					Crn:      &crn,
					Name:     &name,
					Revision: &revision,
				},
			},
		}, nil)

	r := &dfProjectResource{client: &cdp.Client{Df: NewMockDf(mockClient)}}
	s := projectSchema()

	req := resource.CreateRequest{
		Plan: tfsdk.Plan{Raw: createRawProjectResource(""), Schema: s.Schema},
	}
	resp := &resource.CreateResponse{
		State: tfsdk.State{Schema: s.Schema},
	}

	r.Create(ctx, req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	var state projectModel
	resp.State.Get(ctx, &state)
	assert.Equal(t, crn, state.Crn.ValueString())
	assert.Equal(t, crn, state.ID.ValueString())
}

func TestCreateProject_Error(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDfClientService(t)

	mockClient.On("CreateProject", mock.Anything).Return(nil, errors.New("permission denied"))

	r := &dfProjectResource{client: &cdp.Client{Df: NewMockDf(mockClient)}}
	s := projectSchema()

	req := resource.CreateRequest{
		Plan: tfsdk.Plan{Raw: createRawProjectResource(""), Schema: s.Schema},
	}
	resp := &resource.CreateResponse{
		State: tfsdk.State{Schema: s.Schema},
	}

	r.Create(ctx, req, resp)

	assert.True(t, resp.Diagnostics.HasError())
	assert.Contains(t, resp.Diagnostics.Errors()[0].Detail(), "permission denied")
}

func TestReadProject_OK(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDfClientService(t)

	crn := "crn:cdp:df:us-west-1:tenant:project:test-id"
	name := "test-project"
	var revision int32 = 1
	mockClient.On("DescribeProject", mock.Anything).Return(
		&operations.DescribeProjectOK{
			Payload: &dfmodels.DescribeProjectResponse{
				Project: &dfmodels.Project{
					Crn:         &crn,
					Name:        &name,
					Description: "test description",
					Revision:    &revision,
				},
			},
		}, nil)

	r := &dfProjectResource{client: &cdp.Client{Df: NewMockDf(mockClient)}}
	s := projectSchema()

	req := resource.ReadRequest{
		State: tfsdk.State{Raw: createRawProjectResource(crn), Schema: s.Schema},
	}
	resp := &resource.ReadResponse{
		State: tfsdk.State{Schema: s.Schema},
	}

	r.Read(ctx, req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	var state projectModel
	resp.State.Get(ctx, &state)
	assert.Equal(t, crn, state.Crn.ValueString())
	assert.Equal(t, "test-project", state.Name.ValueString())
}

func TestReadProject_NotFound(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDfClientService(t)

	mockClient.On("DescribeProject", mock.Anything).Return(nil, errors.New("NOT_FOUND"))

	r := &dfProjectResource{client: &cdp.Client{Df: NewMockDf(mockClient)}}
	s := projectSchema()

	crn := "crn:cdp:df:us-west-1:tenant:project:test-id"
	req := resource.ReadRequest{
		State: tfsdk.State{Raw: createRawProjectResource(crn), Schema: s.Schema},
	}
	resp := &resource.ReadResponse{
		State: tfsdk.State{Schema: s.Schema},
	}

	r.Read(ctx, req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.True(t, resp.State.Raw.IsNull())
}

func TestDeleteProject_OK(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDfClientService(t)

	mockClient.On("DeleteProject", mock.Anything).Return(&operations.DeleteProjectOK{}, nil)

	r := &dfProjectResource{client: &cdp.Client{Df: NewMockDf(mockClient)}}
	s := projectSchema()

	crn := "crn:cdp:df:us-west-1:tenant:project:test-id"
	req := resource.DeleteRequest{
		State: tfsdk.State{Raw: createRawProjectResource(crn), Schema: s.Schema},
	}
	resp := &resource.DeleteResponse{
		State: tfsdk.State{Schema: s.Schema},
	}

	r.Delete(ctx, req, resp)

	assert.False(t, resp.Diagnostics.HasError())
}
