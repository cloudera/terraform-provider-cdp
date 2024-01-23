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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	mocks "github.com/cloudera/terraform-provider-cdp/mocks/github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createRawAzureImageTermsResource(accepted bool) tftypes.Value {
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"id":       tftypes.String,
				"accepted": tftypes.Bool,
			},
		},
		map[string]tftypes.Value{
			"id":       tftypes.NewValue(tftypes.String, ""),
			"accepted": tftypes.NewValue(tftypes.Bool, accepted),
		},
	)
}

func TestCreateAzureImageTerms(t *testing.T) {
	testCases := map[string]struct {
		expectedResponse      interface{}
		expectedErrorResponse interface{}
		expectedError         bool
		expectedSummary       string
		expectedDetail        string
		expectedAccepted      bool
	}{
		"OK": {
			expectedResponse:      &operations.UpdateAzureImageTermsPolicyOK{},
			expectedErrorResponse: nil,
			expectedError:         false,
			expectedSummary:       "",
			expectedDetail:        "",
			expectedAccepted:      true,
		},
		"TransportError": {
			expectedResponse:      nil,
			expectedErrorResponse: errors.New("request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)"),
			expectedError:         true,
			expectedSummary:       "Create Azure Image Terms Policy",
			expectedDetail:        "Failed to create Azure Image Terms Policy, unexpected error: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)",
			expectedAccepted:      false,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()

			mockClient := new(mocks.MockEnvironmentClientService)

			createMatcher := func(params *operations.UpdateAzureImageTermsPolicyParams) bool {
				return *params.Input.Accepted
			}
			mockClient.On("UpdateAzureImageTermsPolicy", mock.MatchedBy(createMatcher)).Return(testCase.expectedResponse, testCase.expectedErrorResponse)

			aitpResource := &azureImageTermsResource{
				client: &cdp.Client{Environments: NewMockEnvironments(mockClient)},
			}

			req := resource.CreateRequest{
				Plan: tfsdk.Plan{
					Raw:    createRawAzureImageTermsResource(true),
					Schema: AzureImageTermsPolicySchema,
				},
			}

			resp := &resource.CreateResponse{
				State: tfsdk.State{
					Schema: AzureImageTermsPolicySchema,
				},
			}

			aitpResource.Create(ctx, req, resp)

			assert.Equal(t, testCase.expectedError, resp.Diagnostics.HasError())
			if testCase.expectedError {
				assert.Equal(t, testCase.expectedSummary, resp.Diagnostics.Errors()[0].Summary())
				assert.Equal(t, testCase.expectedDetail, resp.Diagnostics.Errors()[0].Detail())
			}

			var state azureImageTermsResourceModel
			resp.State.Get(ctx, &state)

			assert.Equal(t, testCase.expectedAccepted, state.Accepted.ValueBool())

			mockClient.AssertExpectations(t)
		})
	}
}

func TestReadAzureImageTermsPolicy(t *testing.T) {
	testCases := map[string]struct {
		expectedResponse      interface{}
		expectedErrorResponse interface{}
		expectedError         bool
		expectedWarning       bool
		expectedSummary       string
		expectedDetail        string
		expectedAccepted      bool
	}{
		"OK": {
			expectedResponse: &operations.GetAzureImageTermsPolicyOK{
				Payload: &models.GetAzureImageTermsPolicyResponse{
					Accepted: func(in bool) *bool { return &in }(true),
				},
			},
			expectedErrorResponse: nil,
			expectedError:         false,
			expectedSummary:       "",
			expectedDetail:        "",
			expectedAccepted:      true,
		},
		"TransportError": {
			expectedResponse:      nil,
			expectedErrorResponse: errors.New("request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)"),
			expectedError:         true,
			expectedSummary:       "Read Azure Image Terms Policy",
			expectedDetail:        "Failed to read Azure Image Terms Policy, unexpected error: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)",
			expectedAccepted:      false,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()

			mockClient := new(mocks.MockEnvironmentClientService)
			readMatcher := func(params *operations.GetAzureImageTermsPolicyParams) bool {
				return true
			}
			mockClient.On("GetAzureImageTermsPolicy", mock.MatchedBy(readMatcher)).Return(testCase.expectedResponse, testCase.expectedErrorResponse)

			aitpResource := &azureImageTermsResource{
				client: &cdp.Client{Environments: NewMockEnvironments(mockClient)},
			}

			req := resource.ReadRequest{
				State: tfsdk.State{
					Raw:    createRawAzureImageTermsResource(true),
					Schema: AzureImageTermsPolicySchema,
				},
			}

			resp := &resource.ReadResponse{
				State: tfsdk.State{
					Schema: AzureImageTermsPolicySchema,
				},
			}

			aitpResource.Read(ctx, req, resp)

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

			var state azureImageTermsResourceModel
			resp.State.Get(ctx, &state)

			assert.Equal(t, testCase.expectedAccepted, state.Accepted.ValueBool())

			mockClient.AssertExpectations(t)
		})
	}
}

func TestUpdateAzureImageTermsPolicy(t *testing.T) {
	testCases := map[string]struct {
		expectedResponse      interface{}
		expectedErrorResponse interface{}
		expectedError         bool
		expectedWarning       bool
		expectedSummary       string
		expectedDetail        string
		expectedAccepted      bool
	}{
		"OK": {
			expectedResponse:      &operations.UpdateAzureImageTermsPolicyOK{},
			expectedErrorResponse: nil,
			expectedError:         false,
			expectedSummary:       "",
			expectedDetail:        "",
			expectedAccepted:      true,
		},
		"TransportError": {
			expectedResponse:      nil,
			expectedErrorResponse: errors.New("request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)"),
			expectedError:         true,
			expectedSummary:       "Update Azure Image Terms Policy",
			expectedDetail:        "Failed to update Azure Image Terms Policy, unexpected error: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)",
			expectedAccepted:      false,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()

			mockClient := new(mocks.MockEnvironmentClientService)
			updateMatcher := func(params *operations.UpdateAzureImageTermsPolicyParams) bool {
				return *params.Input.Accepted
			}
			mockClient.On("UpdateAzureImageTermsPolicy", mock.MatchedBy(updateMatcher)).Return(testCase.expectedResponse, testCase.expectedErrorResponse)

			aitpResource := &azureImageTermsResource{
				client: &cdp.Client{Environments: NewMockEnvironments(mockClient)},
			}

			req := resource.UpdateRequest{
				Plan: tfsdk.Plan{
					Raw:    createRawAzureImageTermsResource(true),
					Schema: AzureImageTermsPolicySchema,
				},
			}

			resp := &resource.UpdateResponse{
				State: tfsdk.State{
					Raw:    createRawAzureImageTermsResource(false),
					Schema: AzureImageTermsPolicySchema,
				},
			}

			aitpResource.Update(ctx, req, resp)

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

			var state azureImageTermsResourceModel
			resp.State.Get(ctx, &state)

			assert.Equal(t, testCase.expectedAccepted, state.Accepted.ValueBool())

			mockClient.AssertExpectations(t)
		})
	}
}

func TestDeleteAzureImageTermsPolicy(t *testing.T) {
	testCases := map[string]struct {
		expectedResponse      interface{}
		expectedErrorResponse interface{}
		expectedError         bool
		expectedWarning       bool
		expectedSummary       string
		expectedDetail        string
		expectedAccepted      bool
	}{
		"OK": {
			expectedResponse:      &operations.UpdateAzureImageTermsPolicyOK{},
			expectedErrorResponse: nil,
			expectedError:         false,
			expectedSummary:       "",
			expectedDetail:        "",
			expectedAccepted:      false,
		},
		"TransportError": {
			expectedResponse:      nil,
			expectedErrorResponse: errors.New("request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)"),
			expectedError:         true,
			expectedSummary:       "Delete Azure Image Terms Policy",
			expectedDetail:        "Failed to delete Azure Image Terms Policy, unexpected error: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)",
			expectedAccepted:      false,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()

			mockClient := new(mocks.MockEnvironmentClientService)
			updateMatcher := func(params *operations.UpdateAzureImageTermsPolicyParams) bool {
				return !*params.Input.Accepted
			}
			mockClient.On("UpdateAzureImageTermsPolicy", mock.MatchedBy(updateMatcher)).Return(testCase.expectedResponse, testCase.expectedErrorResponse)

			aitpResource := &azureImageTermsResource{
				client: &cdp.Client{Environments: NewMockEnvironments(mockClient)},
			}

			req := resource.DeleteRequest{
				State: tfsdk.State{
					Raw:    createRawAzureImageTermsResource(true),
					Schema: AzureImageTermsPolicySchema,
				},
			}

			resp := &resource.DeleteResponse{
				State: tfsdk.State{
					Schema: AzureImageTermsPolicySchema,
				},
			}

			aitpResource.Delete(ctx, req, resp)

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

			var state azureImageTermsResourceModel
			resp.State.Get(ctx, &state)

			assert.Equal(t, testCase.expectedAccepted, state.Accepted.ValueBool())
			assert.True(t, resp.State.Raw.IsNull())

			mockClient.AssertExpectations(t)
		})
	}
}
