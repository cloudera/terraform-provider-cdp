// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package environments

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
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/mocks"
)

func createRawProxyConfigResource(resourceID string) tftypes.Value {
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"id":          tftypes.String,
				"name":        tftypes.String,
				"description": tftypes.String,
				"protocol":    tftypes.String,
				"host":        tftypes.String,
				"port":        tftypes.Number,
				"no_proxy_hosts": tftypes.Set{
					ElementType: tftypes.String,
				},
				"user":     tftypes.String,
				"password": tftypes.String,
			},
		},
		map[string]tftypes.Value{
			"id":          tftypes.NewValue(tftypes.String, resourceID),
			"name":        tftypes.NewValue(tftypes.String, "test-name"),
			"description": tftypes.NewValue(tftypes.String, "test-description"),
			"protocol":    tftypes.NewValue(tftypes.String, "test-protocol"),
			"host":        tftypes.NewValue(tftypes.String, "test-host"),
			"port":        tftypes.NewValue(tftypes.Number, 99),
			"no_proxy_hosts": tftypes.NewValue(tftypes.Set{
				ElementType: tftypes.String,
			}, []tftypes.Value{tftypes.NewValue(tftypes.String, "test-npc1"), tftypes.NewValue(tftypes.String, "test-npc2")}),
			"user":     tftypes.NewValue(tftypes.String, "test-user"),
			"password": tftypes.NewValue(tftypes.String, "test-password"),
		},
	)
}

func TestCreateProxyConfiguration(t *testing.T) {
	testCases := map[string]struct {
		expectedResponse      interface{}
		expectedErrorResponse interface{}
		expectedError         bool
		expectedSummary       string
		expectedDetail        string
		expectedID            string
		expectedName          string
		expectedHost          string
	}{
		"OK": {
			expectedResponse: &operations.CreateProxyConfigOK{
				Payload: &models.CreateProxyConfigResponse{
					ProxyConfig: &models.ProxyConfig{
						Crn:             func(s string) *string { return &s }("test-pc-crn"),
						Description:     "test_description",
						Host:            func(s string) *string { return &s }("test-host"),
						NoProxyHosts:    "test-npc1,test-npc2",
						Password:        "test-password",
						Port:            func(i int32) *int32 { return &i }(99),
						Protocol:        func(s string) *string { return &s }("test-protocol"),
						ProxyConfigName: func(s string) *string { return &s }("test-name"),
						User:            "test-user",
					},
				},
			},
			expectedErrorResponse: nil,
			expectedError:         false,
			expectedSummary:       "",
			expectedDetail:        "",
			expectedID:            "test-pc-crn",
			expectedName:          "test-name",
			expectedHost:          "test-host",
		},
		"BadRequest": {
			expectedResponse: nil,
			expectedErrorResponse: &operations.CreateProxyConfigDefault{
				Payload: &models.Error{
					Code:    "BAD_REQUEST",
					Message: "Missing name field",
				},
			},
			expectedError:   true,
			expectedSummary: "Create Proxy Configuration",
			expectedDetail:  "Failed to create proxy configuration, unexpected error: Missing name field",
			expectedID:      "",
			expectedName:    "",
			expectedHost:    "",
		},
		"TransportError": {
			expectedResponse:      nil,
			expectedErrorResponse: errors.New("request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)"),
			expectedError:         true,
			expectedSummary:       "Create Proxy Configuration",
			expectedDetail:        "Failed to create proxy configuration, unexpected error: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)",
			expectedID:            "",
			expectedName:          "",
			expectedHost:          "",
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()

			mockClient := new(mocks.MockEnvironmentClientService)

			createMatcher := func(params *operations.CreateProxyConfigParams) bool {
				match := *params.Input.ProxyConfigName == "test-name"
				match = match && params.Input.Description == "test-description"
				match = match && *params.Input.Protocol == "test-protocol"
				match = match && *params.Input.Host == "test-host"
				match = match && *params.Input.Port == 99
				match = match && params.Input.NoProxyHosts == "test-npc1,test-npc2"
				match = match && params.Input.User == "test-user"
				match = match && params.Input.Password == "test-password"
				return match
			}
			mockClient.On("CreateProxyConfig", mock.MatchedBy(createMatcher)).Return(testCase.expectedResponse, testCase.expectedErrorResponse)

			pcResource := &proxyConfigurationResource{
				client: &cdp.Client{Environments: NewMockEnvironments(mockClient)},
			}

			req := resource.CreateRequest{
				Plan: tfsdk.Plan{
					Raw:    createRawProxyConfigResource(""),
					Schema: ProxyConfigurationSchema,
				},
			}

			resp := &resource.CreateResponse{
				State: tfsdk.State{
					Schema: ProxyConfigurationSchema,
				},
			}

			pcResource.Create(ctx, req, resp)

			assert.Equal(t, testCase.expectedError, resp.Diagnostics.HasError())
			if testCase.expectedError {
				assert.Equal(t, testCase.expectedSummary, resp.Diagnostics.Errors()[0].Summary())
				assert.Equal(t, testCase.expectedDetail, resp.Diagnostics.Errors()[0].Detail())
			}

			var state proxyConfigurationResourceModel
			resp.State.Get(ctx, &state)

			assert.Equal(t, testCase.expectedID, state.ID.ValueString())
			assert.Equal(t, testCase.expectedName, state.Name.ValueString())
			assert.Equal(t, testCase.expectedHost, state.Host.ValueString())

			mockClient.AssertExpectations(t)
		})
	}
}

func TestReadProxyConfiguration(t *testing.T) {
	testCases := map[string]struct {
		expectedResponse      interface{}
		expectedErrorResponse interface{}
		expectedError         bool
		expectedSummary       string
		expectedDetail        string
		expectedID            string
		expectedName          string
		expectedHost          string
	}{
		"OK": {
			expectedResponse: &operations.ListProxyConfigsOK{
				Payload: &models.ListProxyConfigsResponse{
					ProxyConfigs: []*models.ProxyConfig{
						{
							Crn:             func(s string) *string { return &s }("test-pc-crn"),
							Description:     "test_description",
							Host:            func(s string) *string { return &s }("test-host"),
							NoProxyHosts:    "test-npc1,test-npc2",
							Password:        "test-password",
							Port:            func(i int32) *int32 { return &i }(99),
							Protocol:        func(s string) *string { return &s }("test-protocol"),
							ProxyConfigName: func(s string) *string { return &s }("test-name"),
							User:            "test-user",
						},
					},
				},
			},
			expectedErrorResponse: nil,
			expectedError:         false,
			expectedSummary:       "",
			expectedDetail:        "",
			expectedID:            "test-pc-crn",
			expectedName:          "test-name",
			expectedHost:          "test-host",
		},
		"BadRequest": {
			expectedResponse: nil,
			expectedErrorResponse: &operations.ListProxyConfigsDefault{
				Payload: &models.Error{
					Code:    "BAD_REQUEST",
					Message: "Missing name field",
				},
			},
			expectedError:   true,
			expectedSummary: "Read Proxy Configuration",
			expectedDetail:  "Failed to read proxy configuration, unexpected error: Missing name field",
			expectedID:      "",
			expectedName:    "",
			expectedHost:    "",
		},
		"TransportError": {
			expectedResponse:      nil,
			expectedErrorResponse: errors.New("request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)"),
			expectedError:         true,
			expectedSummary:       "Read Proxy Configuration",
			expectedDetail:        "Failed to read proxy configuration, unexpected error: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)",
			expectedID:            "",
			expectedName:          "",
			expectedHost:          "",
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()

			mockClient := new(mocks.MockEnvironmentClientService)

			createMatcher := func(params *operations.ListProxyConfigsParams) bool {
				return params.Input.ProxyConfigName == "test-name"
			}
			mockClient.On("ListProxyConfigs", mock.MatchedBy(createMatcher)).Return(testCase.expectedResponse, testCase.expectedErrorResponse)

			pcResource := &proxyConfigurationResource{
				client: &cdp.Client{Environments: NewMockEnvironments(mockClient)},
			}

			req := resource.ReadRequest{
				State: tfsdk.State{
					Raw:    createRawProxyConfigResource("test-pc-crn"),
					Schema: ProxyConfigurationSchema,
				},
			}

			resp := &resource.ReadResponse{
				State: tfsdk.State{
					Schema: ProxyConfigurationSchema,
				},
			}

			pcResource.Read(ctx, req, resp)

			assert.Equal(t, testCase.expectedError, resp.Diagnostics.HasError())
			if testCase.expectedError {
				assert.Equal(t, testCase.expectedSummary, resp.Diagnostics.Errors()[0].Summary())
				assert.Equal(t, testCase.expectedDetail, resp.Diagnostics.Errors()[0].Detail())
			}

			var state proxyConfigurationResourceModel
			resp.State.Get(ctx, &state)

			assert.Equal(t, testCase.expectedID, state.ID.ValueString())
			assert.Equal(t, testCase.expectedName, state.Name.ValueString())
			assert.Equal(t, testCase.expectedHost, state.Host.ValueString())

			mockClient.AssertExpectations(t)
		})
	}
}

type deleteProxyConfigResponseBody struct{}

func TestDeleteProxyConfiguration(t *testing.T) {
	testCases := map[string]struct {
		expectedResponse      interface{}
		expectedErrorResponse interface{}
		expectedError         bool
		expectedSummary       string
		expectedDetail        string
		expectedID            string
		expectedName          string
		expectedHost          string
	}{
		"OK": {
			expectedResponse: &operations.DeleteProxyConfigOK{
				Payload: deleteProxyConfigResponseBody{},
			},
			expectedErrorResponse: nil,
			expectedError:         false,
			expectedSummary:       "",
			expectedDetail:        "",
			expectedID:            "",
			expectedName:          "",
			expectedHost:          "",
		},
		"BadRequest": {
			expectedResponse: nil,
			expectedErrorResponse: &operations.DeleteProxyConfigDefault{
				Payload: &models.Error{
					Code:    "BAD_REQUEST",
					Message: "Missing name field",
				},
			},
			expectedError:   true,
			expectedSummary: "Delete Proxy Configuration",
			expectedDetail:  "Failed to delete proxy configuration: Missing name field",
			expectedID:      "",
			expectedName:    "",
			expectedHost:    "",
		},
		"TransportError": {
			expectedResponse:      nil,
			expectedErrorResponse: errors.New("request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)"),
			expectedError:         true,
			expectedSummary:       "Delete Proxy Configuration",
			expectedDetail:        "Failed to delete proxy configuration, unexpected error: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)",
			expectedID:            "",
			expectedName:          "",
			expectedHost:          "",
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()

			mockClient := new(mocks.MockEnvironmentClientService)

			createMatcher := func(params *operations.DeleteProxyConfigParams) bool {
				return *params.Input.ProxyConfigName == "test-name"
			}
			mockClient.On("DeleteProxyConfig", mock.MatchedBy(createMatcher)).Return(testCase.expectedResponse, testCase.expectedErrorResponse)

			pcResource := &proxyConfigurationResource{
				client: &cdp.Client{Environments: NewMockEnvironments(mockClient)},
			}

			req := resource.DeleteRequest{
				State: tfsdk.State{
					Raw:    createRawProxyConfigResource("test-pc-crn"),
					Schema: ProxyConfigurationSchema,
				},
			}

			resp := &resource.DeleteResponse{
				State: tfsdk.State{
					Schema: ProxyConfigurationSchema,
				},
			}

			pcResource.Delete(ctx, req, resp)

			assert.Equal(t, testCase.expectedError, resp.Diagnostics.HasError())
			if testCase.expectedError {
				assert.Equal(t, testCase.expectedSummary, resp.Diagnostics.Errors()[0].Summary())
				assert.Equal(t, testCase.expectedDetail, resp.Diagnostics.Errors()[0].Detail())
			}

			var state proxyConfigurationResourceModel
			resp.State.Get(ctx, &state)

			assert.Equal(t, testCase.expectedID, state.ID.ValueString())
			assert.Equal(t, testCase.expectedName, state.Name.ValueString())
			assert.Equal(t, testCase.expectedHost, state.Host.ValueString())
			assert.True(t, resp.State.Raw.IsNull())

			mockClient.AssertExpectations(t)
		})
	}
}
