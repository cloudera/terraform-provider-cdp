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

	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/mocks"
)

type MockOperations struct {
	operations.ClientService
}

func (m *MockOperations) SetTransport(transport runtime.ClientTransport) {
}

type MockTransport struct {
	runtime.ClientTransport
}

func NewMockEnvironments(mockClient *mocks.MockEnvironmentClientService) *environmentsclient.Environments {
	return &environmentsclient.Environments{
		Operations: mockClient,
		Transport:  &MockTransport{},
	}
}

func createRawResource(resourceID string, rcaaRole string) tftypes.Value {
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"id":               tftypes.String,
				"data_access_role": tftypes.String,
				"environment_name": tftypes.String,
				"environment_crn":  tftypes.String,
				"mappings": tftypes.Set{
					ElementType: tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"accessor_crn": tftypes.String,
							"role":         tftypes.String,
						},
					},
				},
				"ranger_audit_role":                   tftypes.String,
				"ranger_cloud_access_authorizer_role": tftypes.String,
				"set_empty_mappings":                  tftypes.Bool,
				"mappings_version":                    tftypes.Number,
			},
		},
		map[string]tftypes.Value{
			"id":               tftypes.NewValue(tftypes.String, resourceID),
			"data_access_role": tftypes.NewValue(tftypes.String, "test-da-role"),
			"environment_name": tftypes.NewValue(tftypes.String, "test-env-name"),
			"environment_crn":  tftypes.NewValue(tftypes.String, "test-env-crn"),
			"mappings": tftypes.NewValue(tftypes.Set{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"accessor_crn": tftypes.String,
						"role":         tftypes.String,
					},
				},
			}, []tftypes.Value{tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"accessor_crn": tftypes.String,
					"role":         tftypes.String,
				},
			}, map[string]tftypes.Value{
				"accessor_crn": tftypes.NewValue(tftypes.String, "test-acrn"),
				"role":         tftypes.NewValue(tftypes.String, "test-role"),
			})}),
			"ranger_audit_role":                   tftypes.NewValue(tftypes.String, "test-ra-role"),
			"ranger_cloud_access_authorizer_role": tftypes.NewValue(tftypes.String, rcaaRole),
			"set_empty_mappings":                  tftypes.NewValue(tftypes.Bool, false),
			"mappings_version":                    tftypes.NewValue(tftypes.Number, 0),
		},
	)
}

func TestCreate(t *testing.T) {
	testCases := map[string]struct {
		expectedResponse                        interface{}
		expectedErrorResponse                   interface{}
		expectedError                           bool
		expectedSummary                         string
		expectedDetail                          string
		expectedID                              string
		expectedRangerCloudAccessAuthorizerRole string
		expectedMappingVersion                  int64
	}{
		"OK": {
			expectedResponse: &operations.SetIDBrokerMappingsOK{
				Payload: &models.SetIDBrokerMappingsResponse{
					RangerCloudAccessAuthorizerRole: "test-rcaa-role",
					MappingsVersion:                 func(i int64) *int64 { return &i }(1),
				},
			},
			expectedErrorResponse:                   nil,
			expectedError:                           false,
			expectedSummary:                         "",
			expectedDetail:                          "",
			expectedID:                              "test-env-crn",
			expectedRangerCloudAccessAuthorizerRole: "test-rcaa-role",
			expectedMappingVersion:                  1,
		},
		"EnvironmentNotFound": {
			expectedResponse: nil,
			expectedErrorResponse: &operations.SetIDBrokerMappingsDefault{
				Payload: &models.Error{
					Code:    "NOT_FOUND",
					Message: "",
				},
			},
			expectedError:                           true,
			expectedSummary:                         "Error applying ID Broker mappings",
			expectedDetail:                          "Environment not found: test-env-crn",
			expectedID:                              "",
			expectedRangerCloudAccessAuthorizerRole: "",
			expectedMappingVersion:                  0,
		},
		"TransportError": {
			expectedResponse:                        nil,
			expectedErrorResponse:                   errors.New("request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)"),
			expectedError:                           true,
			expectedSummary:                         "Create Id Broker Mapping",
			expectedDetail:                          "Failed to create ID Broker mapping, unexpected error: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)",
			expectedID:                              "",
			expectedRangerCloudAccessAuthorizerRole: "",
			expectedMappingVersion:                  0,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()

			mockClient := new(mocks.MockEnvironmentClientService)

			createMatcher := func(params *operations.SetIDBrokerMappingsParams) bool {
				match := *params.Input.DataAccessRole == "test-da-role"
				match = match && *params.Input.EnvironmentName == "test-env-name"
				match = match && *params.Input.EnvironmentName == "test-env-name"
				match = match && params.Input.RangerAuditRole == "test-ra-role"
				match = match && params.Input.RangerCloudAccessAuthorizerRole == ""
				match = match && !*params.Input.SetEmptyMappings
				match = match && len(params.Input.Mappings) == 1 && *params.Input.Mappings[0].AccessorCrn == "test-acrn" && *params.Input.Mappings[0].Role == "test-role"
				return match
			}
			mockClient.On("SetIDBrokerMappings", mock.MatchedBy(createMatcher)).Return(testCase.expectedResponse, testCase.expectedErrorResponse)

			idmResource := &idBrokerMappingsResource{
				client: &cdp.Client{Environments: NewMockEnvironments(mockClient)},
			}

			req := resource.CreateRequest{
				Plan: tfsdk.Plan{
					Raw:    createRawResource("", ""),
					Schema: IDBrokerMappingSchema,
				},
			}

			resp := &resource.CreateResponse{
				State: tfsdk.State{
					Schema: IDBrokerMappingSchema,
				},
			}

			idmResource.Create(ctx, req, resp)

			assert.Equal(t, testCase.expectedError, resp.Diagnostics.HasError())
			if testCase.expectedError {
				assert.Equal(t, testCase.expectedSummary, resp.Diagnostics.Errors()[0].Summary())
				assert.Equal(t, testCase.expectedDetail, resp.Diagnostics.Errors()[0].Detail())
			}

			var state idBrokerMappingsResourceModel
			resp.State.Get(ctx, &state)

			assert.Equal(t, testCase.expectedID, state.ID.ValueString())
			assert.Equal(t, testCase.expectedRangerCloudAccessAuthorizerRole, state.RangerCloudAccessAuthorizerRole.ValueString())
			assert.Equal(t, testCase.expectedMappingVersion, state.MappingsVersion.ValueInt64())

			mockClient.AssertExpectations(t)
		})
	}
}

func TestRead(t *testing.T) {
	testCases := map[string]struct {
		expectedEnvironmentResponse             interface{}
		expectedEnvironmentErrorResponse        interface{}
		expectedResponse                        interface{}
		expectedErrorResponse                   interface{}
		expectedError                           bool
		expectedWarning                         bool
		expectedSummary                         string
		expectedDetail                          string
		expectedID                              string
		expectedRangerCloudAccessAuthorizerRole string
		expectedMappingVersion                  int64
	}{
		"OK": {
			expectedEnvironmentResponse: &operations.DescribeEnvironmentOK{
				Payload: &models.DescribeEnvironmentResponse{
					Environment: &models.Environment{
						Crn: func(s string) *string { return &s }("test-env-crn"),
					},
				},
			},
			expectedResponse: &operations.GetIDBrokerMappingsOK{
				Payload: &models.GetIDBrokerMappingsResponse{
					RangerCloudAccessAuthorizerRole: "test-rcaa-role",
					MappingsVersion:                 func(i int64) *int64 { return &i }(1),
				},
			},
			expectedErrorResponse:                   nil,
			expectedError:                           false,
			expectedSummary:                         "",
			expectedDetail:                          "",
			expectedID:                              "test-env-crn",
			expectedRangerCloudAccessAuthorizerRole: "test-rcaa-role",
			expectedMappingVersion:                  1,
		},
		"EnvironmentNotFound": {
			expectedEnvironmentResponse: nil,
			expectedEnvironmentErrorResponse: &operations.DescribeEnvironmentDefault{
				Payload: &models.Error{
					Code:    "NOT_FOUND",
					Message: "",
				},
			},
			expectedResponse: nil,
			expectedErrorResponse: &operations.GetIDBrokerMappingsDefault{
				Payload: &models.Error{
					Code:    "NOT_FOUND",
					Message: "",
				},
			},
			expectedError:                           false,
			expectedWarning:                         true,
			expectedSummary:                         "Resource not found on provider",
			expectedDetail:                          "Environment not found, removing ID Broker mapping from state.",
			expectedID:                              "",
			expectedRangerCloudAccessAuthorizerRole: "",
			expectedMappingVersion:                  0,
		},
		"TransportError": {
			expectedResponse:                        nil,
			expectedErrorResponse:                   errors.New("request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)"),
			expectedError:                           true,
			expectedSummary:                         "Read Id Broker Mapping",
			expectedDetail:                          "Failed to read ID Broker mapping, unexpected error: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)",
			expectedID:                              "",
			expectedRangerCloudAccessAuthorizerRole: "",
			expectedMappingVersion:                  0,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()

			mockClient := new(mocks.MockEnvironmentClientService)
			mockClient.On("DescribeEnvironment", mock.MatchedBy(func(params *operations.DescribeEnvironmentParams) bool {
				return *params.Input.EnvironmentName == "test-env-name"
			})).Return(testCase.expectedEnvironmentResponse, testCase.expectedEnvironmentErrorResponse)

			if testCase.expectedEnvironmentErrorResponse == nil {
				var readMatcher = func(params *operations.GetIDBrokerMappingsParams) bool {
					return *params.Input.EnvironmentName == "test-env-name"
				}
				mockClient.On("GetIDBrokerMappings", mock.MatchedBy(readMatcher)).Return(testCase.expectedResponse, testCase.expectedErrorResponse)
			}

			idmResource := &idBrokerMappingsResource{
				client: &cdp.Client{Environments: NewMockEnvironments(mockClient)},
			}

			req := resource.ReadRequest{
				State: tfsdk.State{
					Raw:    createRawResource("test-env-crn", ""),
					Schema: IDBrokerMappingSchema,
				},
			}

			resp := &resource.ReadResponse{
				State: tfsdk.State{
					Schema: IDBrokerMappingSchema,
				},
			}

			idmResource.Read(ctx, req, resp)

			assert.Equal(t, testCase.expectedError, resp.Diagnostics.HasError())
			if testCase.expectedError {
				assert.Equal(t, testCase.expectedSummary, resp.Diagnostics.Errors()[0].Summary())
				assert.Equal(t, testCase.expectedDetail, resp.Diagnostics.Errors()[0].Detail())
			}

			if testCase.expectedWarning {
				assert.Equal(t, 1, resp.Diagnostics.WarningsCount())
				assert.Equal(t, testCase.expectedSummary, resp.Diagnostics.Warnings()[0].Summary())
				assert.Equal(t, testCase.expectedDetail, resp.Diagnostics.Warnings()[0].Detail())
				assert.True(t, resp.State.Raw.IsNull())
			}

			var state idBrokerMappingsResourceModel
			resp.State.Get(ctx, &state)

			assert.Equal(t, testCase.expectedID, state.ID.ValueString())
			assert.Equal(t, testCase.expectedRangerCloudAccessAuthorizerRole, state.RangerCloudAccessAuthorizerRole.ValueString())
			assert.Equal(t, testCase.expectedMappingVersion, state.MappingsVersion.ValueInt64())

			mockClient.AssertExpectations(t)
		})
	}
}

func TestUpdate(t *testing.T) {
	testCases := map[string]struct {
		expectedEnvironmentResponse             interface{}
		expectedEnvironmentErrorResponse        interface{}
		expectedResponse                        interface{}
		expectedErrorResponse                   interface{}
		expectedError                           bool
		expectedWarning                         bool
		expectedSummary                         string
		expectedDetail                          string
		expectedID                              string
		expectedRangerCloudAccessAuthorizerRole string
		expectedMappingVersion                  int64
	}{
		"OK": {
			expectedEnvironmentResponse: &operations.DescribeEnvironmentOK{
				Payload: &models.DescribeEnvironmentResponse{
					Environment: &models.Environment{
						Crn: func(s string) *string { return &s }("test-env-crn"),
					},
				},
			},
			expectedResponse: &operations.SetIDBrokerMappingsOK{
				Payload: &models.SetIDBrokerMappingsResponse{
					RangerCloudAccessAuthorizerRole: "test-rcaa-role",
					MappingsVersion:                 func(i int64) *int64 { return &i }(1),
				},
			},
			expectedErrorResponse:                   nil,
			expectedError:                           false,
			expectedSummary:                         "",
			expectedDetail:                          "",
			expectedID:                              "test-env-crn",
			expectedRangerCloudAccessAuthorizerRole: "test-rcaa-role",
			expectedMappingVersion:                  1,
		},
		"EnvironmentNotFound": {
			expectedEnvironmentResponse: nil,
			expectedEnvironmentErrorResponse: &operations.DescribeEnvironmentDefault{
				Payload: &models.Error{
					Code:    "NOT_FOUND",
					Message: "",
				},
			},
			expectedResponse: nil,
			expectedErrorResponse: &operations.GetIDBrokerMappingsDefault{
				Payload: &models.Error{
					Code:    "NOT_FOUND",
					Message: "",
				},
			},
			expectedError:                           false,
			expectedWarning:                         true,
			expectedSummary:                         "Resource not found on provider",
			expectedDetail:                          "Environment not found, removing ID Broker mapping from state.",
			expectedID:                              "",
			expectedRangerCloudAccessAuthorizerRole: "",
			expectedMappingVersion:                  0,
		},
		"TransportError": {
			expectedResponse:                        nil,
			expectedErrorResponse:                   errors.New("request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)"),
			expectedError:                           true,
			expectedSummary:                         "Error setting ID Broker mappings",
			expectedDetail:                          "Got the following error setting ID Broker mappings: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)",
			expectedID:                              "test-env-crn",
			expectedRangerCloudAccessAuthorizerRole: "test-old-rcaa-role",
			expectedMappingVersion:                  0,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()

			mockClient := new(mocks.MockEnvironmentClientService)
			mockClient.On("DescribeEnvironment", mock.MatchedBy(func(params *operations.DescribeEnvironmentParams) bool {
				return *params.Input.EnvironmentName == "test-env-name"
			})).Return(testCase.expectedEnvironmentResponse, testCase.expectedEnvironmentErrorResponse)

			if testCase.expectedEnvironmentErrorResponse == nil {
				updateMatcher := func(params *operations.SetIDBrokerMappingsParams) bool {
					match := *params.Input.DataAccessRole == "test-da-role"
					match = match && *params.Input.EnvironmentName == "test-env-name"
					match = match && *params.Input.EnvironmentName == "test-env-name"
					match = match && params.Input.RangerAuditRole == "test-ra-role"
					match = match && params.Input.RangerCloudAccessAuthorizerRole == "test-rcaa-role"
					match = match && !*params.Input.SetEmptyMappings
					match = match && len(params.Input.Mappings) == 1 && *params.Input.Mappings[0].AccessorCrn == "test-acrn" && *params.Input.Mappings[0].Role == "test-role"
					return match
				}
				mockClient.On("SetIDBrokerMappings", mock.MatchedBy(updateMatcher)).Return(testCase.expectedResponse, testCase.expectedErrorResponse)
			}

			idmResource := &idBrokerMappingsResource{
				client: &cdp.Client{Environments: NewMockEnvironments(mockClient)},
			}

			req := resource.UpdateRequest{
				Plan: tfsdk.Plan{
					Raw:    createRawResource("test-env-crn", "test-rcaa-role"),
					Schema: IDBrokerMappingSchema,
				},
			}

			resp := &resource.UpdateResponse{
				State: tfsdk.State{
					Raw:    createRawResource("test-env-crn", "test-old-rcaa-role"),
					Schema: IDBrokerMappingSchema,
				},
			}

			idmResource.Update(ctx, req, resp)

			assert.Equal(t, testCase.expectedError, resp.Diagnostics.HasError())
			if testCase.expectedError {
				assert.Equal(t, testCase.expectedSummary, resp.Diagnostics.Errors()[0].Summary())
				assert.Equal(t, testCase.expectedDetail, resp.Diagnostics.Errors()[0].Detail())
			}

			if testCase.expectedWarning {
				assert.Equal(t, 1, resp.Diagnostics.WarningsCount())
				assert.Equal(t, testCase.expectedSummary, resp.Diagnostics.Warnings()[0].Summary())
				assert.Equal(t, testCase.expectedDetail, resp.Diagnostics.Warnings()[0].Detail())
				assert.True(t, resp.State.Raw.IsNull())
			}

			var state idBrokerMappingsResourceModel
			resp.State.Get(ctx, &state)

			assert.Equal(t, testCase.expectedID, state.ID.ValueString())
			assert.Equal(t, testCase.expectedRangerCloudAccessAuthorizerRole, state.RangerCloudAccessAuthorizerRole.ValueString())
			assert.Equal(t, testCase.expectedMappingVersion, state.MappingsVersion.ValueInt64())

			mockClient.AssertExpectations(t)
		})
	}
}

func TestDelete(t *testing.T) {
	testCases := map[string]struct {
		expectedEnvironmentResponse             interface{}
		expectedEnvironmentErrorResponse        interface{}
		expectedResponse                        interface{}
		expectedErrorResponse                   interface{}
		expectedError                           bool
		expectedWarning                         bool
		expectedSummary                         string
		expectedDetail                          string
		expectedID                              string
		expectedRangerCloudAccessAuthorizerRole string
		expectedMappingVersion                  int64
	}{
		"OK": {
			expectedEnvironmentResponse: &operations.DescribeEnvironmentOK{
				Payload: &models.DescribeEnvironmentResponse{
					Environment: &models.Environment{
						Crn: func(s string) *string { return &s }("test-env-crn"),
					},
				},
			},
			expectedResponse: &operations.SetIDBrokerMappingsOK{
				Payload: &models.SetIDBrokerMappingsResponse{
					RangerCloudAccessAuthorizerRole: "test-rcaa-role",
					MappingsVersion:                 func(i int64) *int64 { return &i }(1),
				},
			},
			expectedErrorResponse:                   nil,
			expectedError:                           false,
			expectedSummary:                         "",
			expectedDetail:                          "",
			expectedID:                              "",
			expectedRangerCloudAccessAuthorizerRole: "",
			expectedMappingVersion:                  0,
		},
		"EnvironmentNotFound": {
			expectedEnvironmentResponse: nil,
			expectedEnvironmentErrorResponse: &operations.DescribeEnvironmentDefault{
				Payload: &models.Error{
					Code:    "NOT_FOUND",
					Message: "",
				},
			},
			expectedResponse: nil,
			expectedErrorResponse: &operations.GetIDBrokerMappingsDefault{
				Payload: &models.Error{
					Code:    "NOT_FOUND",
					Message: "",
				},
			},
			expectedError:                           false,
			expectedWarning:                         true,
			expectedSummary:                         "Resource not found on provider",
			expectedDetail:                          "Environment not found, removing ID Broker mapping from state.",
			expectedID:                              "",
			expectedRangerCloudAccessAuthorizerRole: "",
			expectedMappingVersion:                  0,
		},
		"TransportError": {
			expectedResponse:                        nil,
			expectedErrorResponse:                   errors.New("request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)"),
			expectedError:                           true,
			expectedSummary:                         "Delete Id Broker Mapping",
			expectedDetail:                          "Failed to delete ID Broker mapping, unexpected error: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)",
			expectedID:                              "",
			expectedRangerCloudAccessAuthorizerRole: "",
			expectedMappingVersion:                  0,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()

			mockClient := new(mocks.MockEnvironmentClientService)
			mockClient.On("DescribeEnvironment", mock.MatchedBy(func(params *operations.DescribeEnvironmentParams) bool {
				return *params.Input.EnvironmentName == "test-env-name"
			})).Return(testCase.expectedEnvironmentResponse, testCase.expectedEnvironmentErrorResponse)

			if testCase.expectedEnvironmentErrorResponse == nil {
				updateMatcher := func(params *operations.SetIDBrokerMappingsParams) bool {
					match := *params.Input.DataAccessRole == "test-da-role"
					match = match && *params.Input.EnvironmentName == "test-env-name"
					match = match && *params.Input.EnvironmentName == "test-env-name"
					match = match && params.Input.RangerAuditRole == "test-ra-role"
					match = match && params.Input.RangerCloudAccessAuthorizerRole == ""
					match = match && *params.Input.SetEmptyMappings
					match = match && len(params.Input.Mappings) == 0
					return match
				}
				mockClient.On("SetIDBrokerMappings", mock.MatchedBy(updateMatcher)).Return(testCase.expectedResponse, testCase.expectedErrorResponse)
			}

			idmResource := &idBrokerMappingsResource{
				client: &cdp.Client{Environments: NewMockEnvironments(mockClient)},
			}

			req := resource.DeleteRequest{
				State: tfsdk.State{
					Raw:    createRawResource("test-env-crn", ""),
					Schema: IDBrokerMappingSchema,
				},
			}

			resp := &resource.DeleteResponse{
				State: tfsdk.State{
					Schema: IDBrokerMappingSchema,
				},
			}

			idmResource.Delete(ctx, req, resp)

			assert.Equal(t, testCase.expectedError, resp.Diagnostics.HasError())
			if testCase.expectedError {
				assert.Equal(t, testCase.expectedSummary, resp.Diagnostics.Errors()[0].Summary())
				assert.Equal(t, testCase.expectedDetail, resp.Diagnostics.Errors()[0].Detail())
			}

			if testCase.expectedWarning {
				assert.Equal(t, 1, resp.Diagnostics.WarningsCount())
				assert.Equal(t, testCase.expectedSummary, resp.Diagnostics.Warnings()[0].Summary())
				assert.Equal(t, testCase.expectedDetail, resp.Diagnostics.Warnings()[0].Detail())
			}

			var state idBrokerMappingsResourceModel
			resp.State.Get(ctx, &state)

			assert.Equal(t, testCase.expectedID, state.ID.ValueString())
			assert.Equal(t, testCase.expectedRangerCloudAccessAuthorizerRole, state.RangerCloudAccessAuthorizerRole.ValueString())
			assert.Equal(t, testCase.expectedMappingVersion, state.MappingsVersion.ValueInt64())
			assert.True(t, resp.State.Raw.IsNull())

			mockClient.AssertExpectations(t)
		})
	}
}
