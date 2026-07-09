// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package auditcredential

import (
	"context"
	"errors"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/mocks"
)

func TestCreateAwsAuditCredential(t *testing.T) {
	testCases := map[string]struct {
		expectedResponse      *operations.SetAWSAuditCredentialOK
		expectedErrorResponse error
		expectedError         bool
		expectedSummary       string
		expectedDetail        string
		expectedID            string
	}{
		"OK": {
			expectedResponse: &operations.SetAWSAuditCredentialOK{
				Payload: &models.SetAWSAuditCredentialResponse{
					Credential: testAuditCredential(),
				},
			},
			expectedError: false,
			expectedID:    "test-audit-cred",
		},
		"BadRequest": {
			expectedErrorResponse: &operations.SetAWSAuditCredentialDefault{Payload: &models.Error{Code: "BAD_REQUEST", Message: "Invalid role ARN"}},
			expectedError:         true,
			expectedSummary:       "Create Aws Audit Credential",
			expectedDetail:        "Failed to create AWS Audit Credential, unexpected error: Invalid role ARN",
		},
		"TransportError": {
			expectedErrorResponse: errors.New("connection timeout"),
			expectedError:         true,
			expectedSummary:       "Create Aws Audit Credential",
			expectedDetail:        "Failed to create AWS Audit Credential, unexpected error: connection timeout",
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			mockClient := new(mocks.MockEnvironmentClientService)

			createMatcher := func(params *operations.SetAWSAuditCredentialParams) bool {
				return *params.Input.RoleArn == "arn:aws:iam::123456789012:role/audit-role"
			}
			mockClient.On("SetAWSAuditCredentialContext", mock.Anything, mock.MatchedBy(createMatcher)).Return(testCase.expectedResponse, testCase.expectedErrorResponse)

			r := &awsAuditCredentialResource{client: &cdp.Client{Environments: NewMockEnvironments(mockClient)}}
			schemaResp := &resource.SchemaResponse{}
			r.Schema(ctx, resource.SchemaRequest{}, schemaResp)

			req := resource.CreateRequest{Plan: tfsdk.Plan{Raw: createRawAwsAuditCredentialResource(""), Schema: schemaResp.Schema}}
			resp := &resource.CreateResponse{State: tfsdk.State{Schema: schemaResp.Schema}}

			r.Create(ctx, req, resp)

			assert.Equal(t, testCase.expectedError, resp.Diagnostics.HasError())
			if testCase.expectedError {
				assert.Equal(t, testCase.expectedSummary, resp.Diagnostics.Errors()[0].Summary())
				assert.Equal(t, testCase.expectedDetail, resp.Diagnostics.Errors()[0].Detail())
			} else {
				var state awsAuditCredentialResourceModel
				resp.State.Get(ctx, &state)
				assert.Equal(t, testCase.expectedID, state.ID.ValueString())
			}
			mockClient.AssertExpectations(t)
		})
	}
}

func TestReadAwsAuditCredential(t *testing.T) {
	testCases := map[string]struct {
		expectedResponse      *operations.ListAuditCredentialsOK
		expectedErrorResponse error
		expectedError         bool
		expectedSummary       string
		expectedDetail        string
		expectRemoved         bool
		expectedID            string
	}{
		"OK": {
			expectedResponse: &operations.ListAuditCredentialsOK{
				Payload: &models.ListAuditCredentialsResponse{
					Credentials: []*models.Credential{testAuditCredential()},
				},
			},
			expectedError: false,
			expectedID:    "test-audit-cred",
		},
		"NotFound": {
			expectedResponse: &operations.ListAuditCredentialsOK{
				Payload: &models.ListAuditCredentialsResponse{
					Credentials: []*models.Credential{},
				},
			},
			expectedError: false,
			expectRemoved: true,
		},
		"TransportError": {
			expectedErrorResponse: errors.New("connection timeout"),
			expectedError:         true,
			expectedSummary:       "Read Aws Audit Credential",
			expectedDetail:        "Failed to read AWS Audit Credential, unexpected error: connection timeout",
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			mockClient := new(mocks.MockEnvironmentClientService)
			mockClient.On("ListAuditCredentialsContext", mock.Anything, mock.Anything).Return(testCase.expectedResponse, testCase.expectedErrorResponse)

			r := &awsAuditCredentialResource{client: &cdp.Client{Environments: NewMockEnvironments(mockClient)}}
			schemaResp := &resource.SchemaResponse{}
			r.Schema(ctx, resource.SchemaRequest{}, schemaResp)

			req := resource.ReadRequest{State: tfsdk.State{Raw: createRawAwsAuditCredentialResource("test-audit-cred"), Schema: schemaResp.Schema}}
			resp := &resource.ReadResponse{State: tfsdk.State{Schema: schemaResp.Schema}}

			r.Read(ctx, req, resp)

			assert.Equal(t, testCase.expectedError, resp.Diagnostics.HasError())
			if testCase.expectedError {
				assert.Equal(t, testCase.expectedSummary, resp.Diagnostics.Errors()[0].Summary())
				assert.Equal(t, testCase.expectedDetail, resp.Diagnostics.Errors()[0].Detail())
			} else if testCase.expectRemoved {
				assert.True(t, resp.State.Raw.IsNull())
			} else {
				var state awsAuditCredentialResourceModel
				resp.State.Get(ctx, &state)
				assert.Equal(t, testCase.expectedID, state.ID.ValueString())
			}
			mockClient.AssertExpectations(t)
		})
	}
}

func TestUpdateAwsAuditCredential(t *testing.T) {
	ctx := context.TODO()
	mockClient := new(mocks.MockEnvironmentClientService)

	createMatcher := func(params *operations.SetAWSAuditCredentialParams) bool {
		return *params.Input.RoleArn == "arn:aws:iam::123456789012:role/audit-role"
	}
	mockClient.On("SetAWSAuditCredentialContext", mock.Anything, mock.MatchedBy(createMatcher)).Return(
		&operations.SetAWSAuditCredentialOK{Payload: &models.SetAWSAuditCredentialResponse{Credential: testAuditCredential()}}, nil)

	r := &awsAuditCredentialResource{client: &cdp.Client{Environments: NewMockEnvironments(mockClient)}}
	schemaResp := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResp)

	req := resource.UpdateRequest{
		Plan:  tfsdk.Plan{Raw: createRawAwsAuditCredentialResource("test-audit-cred"), Schema: schemaResp.Schema},
		State: tfsdk.State{Raw: createRawAwsAuditCredentialResource("test-audit-cred"), Schema: schemaResp.Schema},
	}
	resp := &resource.UpdateResponse{State: tfsdk.State{Schema: schemaResp.Schema}}

	r.Update(ctx, req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	var state awsAuditCredentialResourceModel
	resp.State.Get(ctx, &state)
	assert.Equal(t, "test-audit-cred", state.ID.ValueString())
	mockClient.AssertExpectations(t)
}

func TestDeleteAwsAuditCredential(t *testing.T) {
	testCases := map[string]struct {
		expectedResponse      *operations.DeleteAuditCredentialOK
		expectedErrorResponse error
		expectedError         bool
		expectedSummary       string
		expectedDetail        string
	}{
		"OK": {
			expectedResponse: &operations.DeleteAuditCredentialOK{},
			expectedError:    false,
		},
		"BadRequest": {
			expectedErrorResponse: &operations.DeleteAuditCredentialDefault{Payload: &models.Error{Code: "NOT_FOUND", Message: "Credential not found"}},
			expectedError:         true,
			expectedSummary:       "Delete Aws Audit Credential",
			expectedDetail:        "Failed to delete AWS Audit Credential, unexpected error: Credential not found",
		},
		"TransportError": {
			expectedErrorResponse: errors.New("connection timeout"),
			expectedError:         true,
			expectedSummary:       "Delete Aws Audit Credential",
			expectedDetail:        "Failed to delete AWS Audit Credential, unexpected error: connection timeout",
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			mockClient := new(mocks.MockEnvironmentClientService)

			deleteMatcher := func(params *operations.DeleteAuditCredentialParams) bool {
				return *params.Input.CredentialName == "test-audit-cred"
			}
			mockClient.On("DeleteAuditCredentialContext", mock.Anything, mock.MatchedBy(deleteMatcher)).Return(testCase.expectedResponse, testCase.expectedErrorResponse)

			r := &awsAuditCredentialResource{client: &cdp.Client{Environments: NewMockEnvironments(mockClient)}}
			schemaResp := &resource.SchemaResponse{}
			r.Schema(ctx, resource.SchemaRequest{}, schemaResp)

			req := resource.DeleteRequest{State: tfsdk.State{Raw: createRawAwsAuditCredentialResource("test-audit-cred"), Schema: schemaResp.Schema}}
			resp := &resource.DeleteResponse{}

			r.Delete(ctx, req, resp)

			assert.Equal(t, testCase.expectedError, resp.Diagnostics.HasError())
			if testCase.expectedError {
				assert.Equal(t, testCase.expectedSummary, resp.Diagnostics.Errors()[0].Summary())
				assert.Equal(t, testCase.expectedDetail, resp.Diagnostics.Errors()[0].Detail())
			}
			mockClient.AssertExpectations(t)
		})
	}
}
