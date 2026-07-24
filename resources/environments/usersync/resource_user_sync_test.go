// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package usersync

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

type MockTransport struct {
	runtime.ClientTransport
}

func (m MockTransport) SubmitContext(ctx context.Context, operation *runtime.ClientOperation) (interface{}, error) {
	return nil, nil
}

func NewMockEnvironments(mockClient *mocks.MockEnvironmentClientService) *environmentsclient.Environments {
	return &environmentsclient.Environments{
		Operations: mockClient,
		Transport:  &MockTransport{},
	}
}

func createRawUserSyncResource() tftypes.Value {
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"id": tftypes.String,
				"environment_names": tftypes.Set{
					ElementType: tftypes.String,
				},
				"polling_options": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"async":                  tftypes.Bool,
						"polling_timeout":        tftypes.Number,
						"call_failure_threshold": tftypes.Number,
					},
				},
			},
		},
		map[string]tftypes.Value{
			"id": tftypes.NewValue(tftypes.String, ""),
			"environment_names": tftypes.NewValue(tftypes.Set{
				ElementType: tftypes.String,
			}, []tftypes.Value{tftypes.NewValue(tftypes.String, "test-env")}),
			"polling_options": tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"async":                  tftypes.Bool,
					"polling_timeout":        tftypes.Number,
					"call_failure_threshold": tftypes.Number,
				},
			}, map[string]tftypes.Value{
				"async":                  tftypes.NewValue(tftypes.Bool, true),
				"polling_timeout":        tftypes.NewValue(tftypes.Number, 90),
				"call_failure_threshold": tftypes.NewValue(tftypes.Number, 3),
			}),
		},
	)
}

func createRawUserSyncResourceWithoutEnvNames() tftypes.Value {
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"id": tftypes.String,
				"environment_names": tftypes.Set{
					ElementType: tftypes.String,
				},
				"polling_options": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"async":                  tftypes.Bool,
						"polling_timeout":        tftypes.Number,
						"call_failure_threshold": tftypes.Number,
					},
				},
			},
		},
		map[string]tftypes.Value{
			"id": tftypes.NewValue(tftypes.String, ""),
			"environment_names": tftypes.NewValue(tftypes.Set{
				ElementType: tftypes.String,
			}, []tftypes.Value{}),
			"polling_options": tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"async":                  tftypes.Bool,
					"polling_timeout":        tftypes.Number,
					"call_failure_threshold": tftypes.Number,
				},
			}, map[string]tftypes.Value{
				"async":                  tftypes.NewValue(tftypes.Bool, true),
				"polling_timeout":        tftypes.NewValue(tftypes.Number, 90),
				"call_failure_threshold": tftypes.NewValue(tftypes.Number, 3),
			}),
		},
	)
}

func TestMetadata(t *testing.T) {
	r := NewUserSyncResource()
	req := resource.MetadataRequest{ProviderTypeName: "cdp"}
	resp := &resource.MetadataResponse{}
	r.Metadata(context.TODO(), req, resp)
	assert.Equal(t, "cdp_environments_user_sync", resp.TypeName)
}

func TestSchema(t *testing.T) {
	r := NewUserSyncResource()
	req := resource.SchemaRequest{}
	resp := &resource.SchemaResponse{}
	r.Schema(context.TODO(), req, resp)

	assert.NotNil(t, resp.Schema.Attributes["id"])
	assert.NotNil(t, resp.Schema.Attributes["environment_names"])
	assert.NotNil(t, resp.Schema.Attributes["polling_options"])
}

func TestCreate(t *testing.T) {
	testCases := map[string]struct {
		expectedResponse      interface{}
		expectedErrorResponse interface{}
		expectedError         bool
		expectedSummary       string
		expectedDetail        string
	}{
		"OK": {
			expectedResponse: &operations.SyncAllUsersOK{
				Payload: &models.SyncAllUsersResponse{
					OperationID: strPtr("op-123"),
				},
			},
			expectedErrorResponse: nil,
			expectedError:         false,
		},
		"EnvironmentNotFound": {
			expectedResponse: nil,
			expectedErrorResponse: &operations.SyncAllUsersDefault{
				Payload: &models.Error{
					Code:    "NOT_FOUND",
					Message: "",
				},
			},
			expectedError:   true,
			expectedSummary: "Error in sync all users",
			expectedDetail:  "An environment not found: [\"test-env\"]",
		},
		"TransportError": {
			expectedResponse:      nil,
			expectedErrorResponse: errors.New("request canceled while waiting for connection"),
			expectedError:         true,
			expectedSummary:       "Sync All Users",
			expectedDetail:        "Failed to sync all users, unexpected error: request canceled while waiting for connection",
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()

			mockClient := new(mocks.MockEnvironmentClientService)
			mockClient.On("SyncAllUsersContext", mock.Anything, mock.Anything, mock.Anything).Return(testCase.expectedResponse, testCase.expectedErrorResponse)

			r := &userSyncResource{
				client: &cdp.Client{Environments: NewMockEnvironments(mockClient)},
			}

			req := resource.CreateRequest{
				Plan: tfsdk.Plan{
					Raw:    createRawUserSyncResource(),
					Schema: userSyncSchema,
				},
			}

			resp := &resource.CreateResponse{
				State: tfsdk.State{
					Schema: userSyncSchema,
				},
			}

			r.Create(ctx, req, resp)

			assert.Equal(t, testCase.expectedError, resp.Diagnostics.HasError())
			if testCase.expectedError {
				assert.Equal(t, testCase.expectedSummary, resp.Diagnostics.Errors()[0].Summary())
				assert.Contains(t, resp.Diagnostics.Errors()[0].Detail(), testCase.expectedDetail)
			} else {
				var state userSyncResourceModel
				resp.State.Get(ctx, &state)
				assert.NotEmpty(t, state.ID.ValueString())
			}

			mockClient.AssertExpectations(t)
		})
	}
}

func TestCreateWithoutEnvironmentNames(t *testing.T) {
	ctx := context.TODO()

	mockClient := new(mocks.MockEnvironmentClientService)
	mockClient.On("SyncAllUsersContext", mock.Anything, mock.Anything, mock.Anything).Return(
		&operations.SyncAllUsersOK{
			Payload: &models.SyncAllUsersResponse{
				OperationID: strPtr("op-456"),
			},
		}, nil)

	r := &userSyncResource{
		client: &cdp.Client{Environments: NewMockEnvironments(mockClient)},
	}

	req := resource.CreateRequest{
		Plan: tfsdk.Plan{
			Raw:    createRawUserSyncResourceWithoutEnvNames(),
			Schema: userSyncSchema,
		},
	}

	resp := &resource.CreateResponse{
		State: tfsdk.State{
			Schema: userSyncSchema,
		},
	}

	r.Create(ctx, req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	var state userSyncResourceModel
	resp.State.Get(ctx, &state)
	assert.NotEmpty(t, state.ID.ValueString())
	mockClient.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	testCases := map[string]struct {
		expectedResponse      interface{}
		expectedErrorResponse interface{}
		expectedError         bool
		expectedSummary       string
	}{
		"OK": {
			expectedResponse: &operations.SyncAllUsersOK{
				Payload: &models.SyncAllUsersResponse{
					OperationID: strPtr("op-789"),
				},
			},
			expectedErrorResponse: nil,
			expectedError:         false,
		},
		"EnvironmentNotFound": {
			expectedResponse: nil,
			expectedErrorResponse: &operations.SyncAllUsersDefault{
				Payload: &models.Error{
					Code:    "NOT_FOUND",
					Message: "",
				},
			},
			expectedError:   true,
			expectedSummary: "Error in sync all users",
		},
		"TransportError": {
			expectedResponse:      nil,
			expectedErrorResponse: errors.New("connection refused"),
			expectedError:         true,
			expectedSummary:       "Sync All Users",
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()

			mockClient := new(mocks.MockEnvironmentClientService)
			mockClient.On("SyncAllUsersContext", mock.Anything, mock.Anything, mock.Anything).Return(testCase.expectedResponse, testCase.expectedErrorResponse)

			r := &userSyncResource{
				client: &cdp.Client{Environments: NewMockEnvironments(mockClient)},
			}

			req := resource.UpdateRequest{
				Plan: tfsdk.Plan{
					Raw:    createRawUserSyncResource(),
					Schema: userSyncSchema,
				},
			}

			resp := &resource.UpdateResponse{
				State: tfsdk.State{
					Schema: userSyncSchema,
				},
			}

			r.Update(ctx, req, resp)

			assert.Equal(t, testCase.expectedError, resp.Diagnostics.HasError())
			if testCase.expectedError {
				assert.Equal(t, testCase.expectedSummary, resp.Diagnostics.Errors()[0].Summary())
			}

			mockClient.AssertExpectations(t)
		})
	}
}

func TestDelete(t *testing.T) {
	r := NewUserSyncResource()
	req := resource.DeleteRequest{}
	resp := &resource.DeleteResponse{}
	r.(*userSyncResource).Delete(context.TODO(), req, resp)
	assert.False(t, resp.Diagnostics.HasError())
}

func TestRead(t *testing.T) {
	r := NewUserSyncResource()
	req := resource.ReadRequest{}
	resp := &resource.ReadResponse{}
	r.(*userSyncResource).Read(context.TODO(), req, resp)
	assert.False(t, resp.Diagnostics.HasError())
}

func TestCheckUserSyncResponseStatusForError(t *testing.T) {
	testCases := map[string]struct {
		status        models.SyncStatus
		expectedError bool
		expectedState string
	}{
		"COMPLETED": {
			status:        "COMPLETED",
			expectedError: false,
			expectedState: "COMPLETED",
		},
		"RUNNING": {
			status:        "RUNNING",
			expectedError: false,
			expectedState: "RUNNING",
		},
		"FAILED": {
			status:        "FAILED",
			expectedError: true,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			resp := &operations.SyncStatusOK{
				Payload: &models.SyncStatusResponse{
					Status: testCase.status,
				},
			}
			result, state, err := checkUserSyncResponseStatusForError(resp)
			if testCase.expectedError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, testCase.expectedState, state)
			}
		})
	}
}

func TestIsSyncAllUsersNotFoundError(t *testing.T) {
	testCases := map[string]struct {
		err      error
		expected bool
	}{
		"NotFoundError": {
			err: &operations.SyncAllUsersDefault{
				Payload: &models.Error{
					Code:    "NOT_FOUND",
					Message: "",
				},
			},
			expected: true,
		},
		"OtherCdpError": {
			err: &operations.SyncAllUsersDefault{
				Payload: &models.Error{
					Code:    "INVALID_ARGUMENT",
					Message: "",
				},
			},
			expected: false,
		},
		"GenericError": {
			err:      errors.New("generic error"),
			expected: false,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, isSyncAllUsersNotFoundError(testCase.err))
		})
	}
}

func strPtr(s string) *string {
	return &s
}
