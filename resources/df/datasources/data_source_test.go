// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datasources

import (
	"context"
	"errors"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
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

// --- cdp_df_project data source tests ---

func projectDsSchema() datasource.SchemaResponse {
	d := &dfProjectDataSource{}
	resp := &datasource.SchemaResponse{}
	d.Schema(context.TODO(), datasource.SchemaRequest{}, resp)
	return *resp
}

func createRawProjectDs(name string) tftypes.Value {
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"name": tftypes.String,
				"crn":  tftypes.String,
			},
		},
		map[string]tftypes.Value{
			"name": tftypes.NewValue(tftypes.String, name),
			"crn":  tftypes.NewValue(tftypes.String, ""),
		},
	)
}

func TestReadProject_Found(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDfClientService(t)

	crn := "crn:cdp:df:us-west-1:tenant:project:p1"
	name := "my-project"
	mockClient.On("ListProjects", mock.Anything).Return(
		&operations.ListProjectsOK{
			Payload: &dfmodels.ListProjectsResponse{
				Projects: []*dfmodels.Project{
					{Crn: &crn, Name: &name},
				},
			},
		}, nil)

	d := &dfProjectDataSource{client: &cdp.Client{Df: NewMockDf(mockClient)}}
	s := projectDsSchema()

	req := datasource.ReadRequest{
		Config: tfsdk.Config{Raw: createRawProjectDs("my-project"), Schema: s.Schema},
	}
	resp := &datasource.ReadResponse{
		State: tfsdk.State{Schema: s.Schema},
	}

	d.Read(ctx, req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	var state dfProjectModel
	resp.State.Get(ctx, &state)
	assert.Equal(t, crn, state.Crn.ValueString())
}

func TestReadProject_NotFound(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDfClientService(t)

	otherName := "other-project"
	otherCrn := "crn:other"
	mockClient.On("ListProjects", mock.Anything).Return(
		&operations.ListProjectsOK{
			Payload: &dfmodels.ListProjectsResponse{
				Projects: []*dfmodels.Project{
					{Crn: &otherCrn, Name: &otherName},
				},
			},
		}, nil)

	d := &dfProjectDataSource{client: &cdp.Client{Df: NewMockDf(mockClient)}}
	s := projectDsSchema()

	req := datasource.ReadRequest{
		Config: tfsdk.Config{Raw: createRawProjectDs("my-project"), Schema: s.Schema},
	}
	resp := &datasource.ReadResponse{
		State: tfsdk.State{Schema: s.Schema},
	}

	d.Read(ctx, req, resp)

	assert.True(t, resp.Diagnostics.HasError())
	assert.Contains(t, resp.Diagnostics.Errors()[0].Detail(), "my-project")
}

func TestReadProject_APIError(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDfClientService(t)

	mockClient.On("ListProjects", mock.Anything).Return(nil, errors.New("connection refused"))

	d := &dfProjectDataSource{client: &cdp.Client{Df: NewMockDf(mockClient)}}
	s := projectDsSchema()

	req := datasource.ReadRequest{
		Config: tfsdk.Config{Raw: createRawProjectDs("my-project"), Schema: s.Schema},
	}
	resp := &datasource.ReadResponse{
		State: tfsdk.State{Schema: s.Schema},
	}

	d.Read(ctx, req, resp)

	assert.True(t, resp.Diagnostics.HasError())
	assert.Contains(t, resp.Diagnostics.Errors()[0].Detail(), "connection refused")
}

// --- cdp_df_service data source tests ---

func serviceDsSchema() datasource.SchemaResponse {
	d := &dfServiceDataSource{}
	resp := &datasource.SchemaResponse{}
	d.Schema(context.TODO(), datasource.SchemaRequest{}, resp)
	return *resp
}

func createRawServiceDs(name string) tftypes.Value {
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"name":             tftypes.String,
				"crn":              tftypes.String,
				"environment_crn":  tftypes.String,
				"cloud_platform":   tftypes.String,
				"region":           tftypes.String,
				"status":           tftypes.String,
				"workload_version": tftypes.String,
				"deployment_count": tftypes.Number,
			},
		},
		map[string]tftypes.Value{
			"name":             tftypes.NewValue(tftypes.String, name),
			"crn":              tftypes.NewValue(tftypes.String, ""),
			"environment_crn":  tftypes.NewValue(tftypes.String, ""),
			"cloud_platform":   tftypes.NewValue(tftypes.String, ""),
			"region":           tftypes.NewValue(tftypes.String, ""),
			"status":           tftypes.NewValue(tftypes.String, ""),
			"workload_version": tftypes.NewValue(tftypes.String, ""),
			"deployment_count": tftypes.NewValue(tftypes.Number, 0),
		},
	)
}

func TestReadService_Found(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDfClientService(t)

	crn := "crn:cdp:df:us-west-1:tenant:service:s1"
	name := "my-service"
	envCrn := "crn:cdp:environments:us-west-1:tenant:environment:e1"
	platform := "AWS"
	region := "us-west-1"
	version := "2.5.0"
	var depCount int32 = 3
	state := dfmodels.ServiceStateGOODHEALTH

	mockClient.On("ListServices", mock.Anything).Return(
		&operations.ListServicesOK{
			Payload: &dfmodels.ListServicesResponse{
				Services: []*dfmodels.ServiceSummary{
					{
						Crn:             &crn,
						Name:            &name,
						EnvironmentCrn:  &envCrn,
						CloudPlatform:   &platform,
						Region:          &region,
						WorkloadVersion: &version,
						DeploymentCount: &depCount,
						Status:          &dfmodels.ServiceStatus{State: &state},
					},
				},
			},
		}, nil)

	d := &dfServiceDataSource{client: &cdp.Client{Df: NewMockDf(mockClient)}}
	s := serviceDsSchema()

	req := datasource.ReadRequest{
		Config: tfsdk.Config{Raw: createRawServiceDs("my-service"), Schema: s.Schema},
	}
	resp := &datasource.ReadResponse{
		State: tfsdk.State{Schema: s.Schema},
	}

	d.Read(ctx, req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	var result dfServiceModel
	resp.State.Get(ctx, &result)
	assert.Equal(t, crn, result.Crn.ValueString())
	assert.Equal(t, envCrn, result.EnvironmentCrn.ValueString())
}

func TestReadService_NotFound(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDfClientService(t)

	otherName := "other-service"
	otherCrn := "crn:other"
	state := dfmodels.ServiceStateGOODHEALTH

	mockClient.On("ListServices", mock.Anything).Return(
		&operations.ListServicesOK{
			Payload: &dfmodels.ListServicesResponse{
				Services: []*dfmodels.ServiceSummary{
					{
						Crn:    &otherCrn,
						Name:   &otherName,
						Status: &dfmodels.ServiceStatus{State: &state},
					},
				},
			},
		}, nil)

	d := &dfServiceDataSource{client: &cdp.Client{Df: NewMockDf(mockClient)}}
	s := serviceDsSchema()

	req := datasource.ReadRequest{
		Config: tfsdk.Config{Raw: createRawServiceDs("my-service"), Schema: s.Schema},
	}
	resp := &datasource.ReadResponse{
		State: tfsdk.State{Schema: s.Schema},
	}

	d.Read(ctx, req, resp)

	assert.True(t, resp.Diagnostics.HasError())
	assert.Contains(t, resp.Diagnostics.Errors()[0].Detail(), "my-service")
}
