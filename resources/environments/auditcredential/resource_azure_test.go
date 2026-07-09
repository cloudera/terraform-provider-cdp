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

type azureSetTestCase struct {
	response      *operations.SetAzureAuditCredentialOK
	errorResponse error
	wantError     bool
	wantID        string
}

var azureSetTestCases = map[string]azureSetTestCase{
	"OK": {
		response: &operations.SetAzureAuditCredentialOK{
			Payload: &models.SetAzureAuditCredentialResponse{
				Credential: testAzureAuditCredential(),
			},
		},
		wantID: "test-audit-cred",
	},
	"NilCredential": {
		response: &operations.SetAzureAuditCredentialOK{
			Payload: &models.SetAzureAuditCredentialResponse{
				Credential: nil,
			},
		},
		wantError: true,
	},
	"BadRequest": {
		errorResponse: &operations.SetAzureAuditCredentialDefault{Payload: &models.Error{Code: "BAD_REQUEST", Message: "Invalid subscription"}},
		wantError:     true,
	},
	"TransportError": {
		errorResponse: errors.New("connection timeout"),
		wantError:     true,
	},
}

func azureSetMatcher(params *operations.SetAzureAuditCredentialParams) bool {
	return *params.Input.SubscriptionID == "sub-id-123" &&
		*params.Input.TenantID == "tenant-id-456" &&
		*params.Input.AppBased.ApplicationID == "app-id-789" &&
		*params.Input.AppBased.SecretKey == "secret-key-abc"
}

func setupAzureResource(t *testing.T, mockClient *mocks.MockEnvironmentClientService) (*azureAuditCredentialResource, resource.SchemaResponse) {
	t.Helper()
	r := &azureAuditCredentialResource{client: &cdp.Client{Environments: NewMockEnvironments(mockClient)}}
	schemaResp := &resource.SchemaResponse{}
	r.Schema(context.TODO(), resource.SchemaRequest{}, schemaResp)
	return r, *schemaResp
}

func TestCreateAzureAuditCredential(t *testing.T) {
	for name, tc := range azureSetTestCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			mockClient := new(mocks.MockEnvironmentClientService)
			mockClient.On("SetAzureAuditCredentialContext", mock.Anything, mock.MatchedBy(azureSetMatcher)).Return(tc.response, tc.errorResponse)

			r, schemaResp := setupAzureResource(t, mockClient)

			req := resource.CreateRequest{Plan: tfsdk.Plan{Raw: createRawAzureAuditCredentialResource(""), Schema: schemaResp.Schema}}
			resp := &resource.CreateResponse{State: tfsdk.State{Schema: schemaResp.Schema}}

			r.Create(ctx, req, resp)

			assertAzureSetResult(t, ctx, resp.Diagnostics.HasError(), resp.State, tc)
			mockClient.AssertExpectations(t)
		})
	}
}

func TestUpdateAzureAuditCredential(t *testing.T) {
	for name, tc := range azureSetTestCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			mockClient := new(mocks.MockEnvironmentClientService)
			mockClient.On("SetAzureAuditCredentialContext", mock.Anything, mock.MatchedBy(azureSetMatcher)).Return(tc.response, tc.errorResponse)

			r, schemaResp := setupAzureResource(t, mockClient)

			req := resource.UpdateRequest{
				Plan:  tfsdk.Plan{Raw: createRawAzureAuditCredentialResource("test-audit-cred"), Schema: schemaResp.Schema},
				State: tfsdk.State{Raw: createRawAzureAuditCredentialResource("test-audit-cred"), Schema: schemaResp.Schema},
			}
			resp := &resource.UpdateResponse{State: tfsdk.State{Schema: schemaResp.Schema}}

			r.Update(ctx, req, resp)

			assertAzureSetResult(t, ctx, resp.Diagnostics.HasError(), resp.State, tc)
			mockClient.AssertExpectations(t)
		})
	}
}

func assertAzureSetResult(t *testing.T, ctx context.Context, hasError bool, state tfsdk.State, tc azureSetTestCase) {
	t.Helper()
	assert.Equal(t, tc.wantError, hasError)
	if !tc.wantError {
		var s azureAuditCredentialResourceModel
		state.Get(ctx, &s)
		assert.Equal(t, tc.wantID, s.ID.ValueString())
		assert.Equal(t, "sub-id-123", s.SubscriptionID.ValueString())
		assert.Equal(t, "tenant-id-456", s.TenantID.ValueString())
		assert.Equal(t, "app-id-789", s.AppBased.ApplicationID.ValueString())
	}
}

func TestReadAzureAuditCredential(t *testing.T) {
	ctx := context.TODO()
	mockClient := new(mocks.MockEnvironmentClientService)
	mockClient.On("ListAuditCredentialsContext", mock.Anything, mock.Anything).Return(
		&operations.ListAuditCredentialsOK{Payload: &models.ListAuditCredentialsResponse{
			Credentials: []*models.Credential{testAzureAuditCredential()},
		}}, nil)

	r, schemaResp := setupAzureResource(t, mockClient)

	req := resource.ReadRequest{State: tfsdk.State{Raw: createRawAzureAuditCredentialResource("test-audit-cred"), Schema: schemaResp.Schema}}
	resp := &resource.ReadResponse{State: tfsdk.State{Schema: schemaResp.Schema}}

	r.Read(ctx, req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	var state azureAuditCredentialResourceModel
	resp.State.Get(ctx, &state)
	assert.Equal(t, "test-audit-cred", state.ID.ValueString())
	assert.Equal(t, "sub-id-123", state.SubscriptionID.ValueString())
	assert.Equal(t, "tenant-id-456", state.TenantID.ValueString())
	assert.Equal(t, "app-id-789", state.AppBased.ApplicationID.ValueString())
	mockClient.AssertExpectations(t)
}

func TestDeleteAzureAuditCredential(t *testing.T) {
	ctx := context.TODO()
	mockClient := new(mocks.MockEnvironmentClientService)

	deleteMatcher := func(params *operations.DeleteAuditCredentialParams) bool {
		return *params.Input.CredentialName == "test-audit-cred"
	}
	mockClient.On("DeleteAuditCredentialContext", mock.Anything, mock.MatchedBy(deleteMatcher)).Return(
		&operations.DeleteAuditCredentialOK{}, nil)

	r, schemaResp := setupAzureResource(t, mockClient)

	req := resource.DeleteRequest{State: tfsdk.State{Raw: createRawAzureAuditCredentialResource("test-audit-cred"), Schema: schemaResp.Schema}}
	resp := &resource.DeleteResponse{}

	r.Delete(ctx, req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	mockClient.AssertExpectations(t)
}
