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

type gcpSetTestCase struct {
	response      *operations.SetGCPAuditCredentialOK
	errorResponse error
	wantError     bool
	wantID        string
}

var gcpSetTestCases = map[string]gcpSetTestCase{
	"OK": {
		response: &operations.SetGCPAuditCredentialOK{
			Payload: &models.SetGCPAuditCredentialResponse{
				Credential: testGcpAuditCredential(),
			},
		},
		wantID: "test-audit-cred",
	},
	"NilCredential": {
		response: &operations.SetGCPAuditCredentialOK{
			Payload: &models.SetGCPAuditCredentialResponse{
				Credential: nil,
			},
		},
		wantError: true,
	},
	"BadRequest": {
		errorResponse: &operations.SetGCPAuditCredentialDefault{Payload: &models.Error{Code: "BAD_REQUEST", Message: "Invalid key"}},
		wantError:     true,
	},
	"TransportError": {
		errorResponse: errors.New("connection timeout"),
		wantError:     true,
	},
}

func gcpSetMatcher(params *operations.SetGCPAuditCredentialParams) bool {
	return *params.Input.CredentialKey == `{"test":"value"}`
}

func setupGcpResource(t *testing.T, mockClient *mocks.MockEnvironmentClientService) (*gcpAuditCredentialResource, resource.SchemaResponse) {
	t.Helper()
	r := &gcpAuditCredentialResource{client: &cdp.Client{Environments: NewMockEnvironments(mockClient)}}
	schemaResp := &resource.SchemaResponse{}
	r.Schema(context.TODO(), resource.SchemaRequest{}, schemaResp)
	return r, *schemaResp
}

func TestCreateGcpAuditCredential(t *testing.T) {
	for name, tc := range gcpSetTestCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			mockClient := new(mocks.MockEnvironmentClientService)
			mockClient.On("SetGCPAuditCredentialContext", mock.Anything, mock.MatchedBy(gcpSetMatcher)).Return(tc.response, tc.errorResponse)

			r, schemaResp := setupGcpResource(t, mockClient)

			req := resource.CreateRequest{Plan: tfsdk.Plan{Raw: createRawGcpAuditCredentialResource(""), Schema: schemaResp.Schema}}
			resp := &resource.CreateResponse{State: tfsdk.State{Schema: schemaResp.Schema}}

			r.Create(ctx, req, resp)

			assertGcpSetResult(t, ctx, resp.Diagnostics.HasError(), resp.State, tc)
			mockClient.AssertExpectations(t)
		})
	}
}

func TestUpdateGcpAuditCredential(t *testing.T) {
	for name, tc := range gcpSetTestCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			mockClient := new(mocks.MockEnvironmentClientService)
			mockClient.On("SetGCPAuditCredentialContext", mock.Anything, mock.MatchedBy(gcpSetMatcher)).Return(tc.response, tc.errorResponse)

			r, schemaResp := setupGcpResource(t, mockClient)

			req := resource.UpdateRequest{
				Plan:  tfsdk.Plan{Raw: createRawGcpAuditCredentialResource("test-audit-cred"), Schema: schemaResp.Schema},
				State: tfsdk.State{Raw: createRawGcpAuditCredentialResource("test-audit-cred"), Schema: schemaResp.Schema},
			}
			resp := &resource.UpdateResponse{State: tfsdk.State{Schema: schemaResp.Schema}}

			r.Update(ctx, req, resp)

			assertGcpSetResult(t, ctx, resp.Diagnostics.HasError(), resp.State, tc)
			mockClient.AssertExpectations(t)
		})
	}
}

func assertGcpSetResult(t *testing.T, ctx context.Context, hasError bool, state tfsdk.State, tc gcpSetTestCase) {
	t.Helper()
	assert.Equal(t, tc.wantError, hasError)
	if !tc.wantError {
		var s gcpAuditCredentialResourceModel
		state.Get(ctx, &s)
		assert.Equal(t, tc.wantID, s.ID.ValueString())
	}
}

func TestReadGcpAuditCredential(t *testing.T) {
	ctx := context.TODO()
	mockClient := new(mocks.MockEnvironmentClientService)
	mockClient.On("ListAuditCredentialsContext", mock.Anything, mock.Anything).Return(
		&operations.ListAuditCredentialsOK{Payload: &models.ListAuditCredentialsResponse{
			Credentials: []*models.Credential{testGcpAuditCredential()},
		}}, nil)

	r, schemaResp := setupGcpResource(t, mockClient)

	req := resource.ReadRequest{State: tfsdk.State{Raw: createRawGcpAuditCredentialResource("test-audit-cred"), Schema: schemaResp.Schema}}
	resp := &resource.ReadResponse{State: tfsdk.State{Schema: schemaResp.Schema}}

	r.Read(ctx, req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	var state gcpAuditCredentialResourceModel
	resp.State.Get(ctx, &state)
	assert.Equal(t, "test-audit-cred", state.ID.ValueString())
	mockClient.AssertExpectations(t)
}

func TestDeleteGcpAuditCredential(t *testing.T) {
	ctx := context.TODO()
	mockClient := new(mocks.MockEnvironmentClientService)

	deleteMatcher := func(params *operations.DeleteAuditCredentialParams) bool {
		return *params.Input.CredentialName == "test-audit-cred"
	}
	mockClient.On("DeleteAuditCredentialContext", mock.Anything, mock.MatchedBy(deleteMatcher)).Return(
		&operations.DeleteAuditCredentialOK{}, nil)

	r, schemaResp := setupGcpResource(t, mockClient)

	req := resource.DeleteRequest{State: tfsdk.State{Raw: createRawGcpAuditCredentialResource("test-audit-cred"), Schema: schemaResp.Schema}}
	resp := &resource.DeleteResponse{}

	r.Delete(ctx, req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	mockClient.AssertExpectations(t)
}
