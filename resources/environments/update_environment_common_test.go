// Copyright 2026 Cloudera. All Rights Reserved.
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
	rsschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/mocks"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

const (
	testEnvName            = "test-env"
	testNewKey             = "ssh-rsa NEW_KEY"
	testOldKey             = "ssh-rsa OLD_KEY"
	testSameKey            = "ssh-rsa SAME_KEY"
	testKey                = "ssh-rsa TEST_KEY"
	testServiceUnavailable = "service unavailable"

	testPlanName   = "plan-name"
	testPlanValue  = "plan-value"
	testStateName  = "state-name"
	testStateValue = "state-value"
	testNewName    = "new-name"
	testNewValue   = "new-value"
	testOldName    = "old-name"
	testOldValue   = "old-value"

	testOldProxyConfigName = "proxy1"
	testNewProxyConfigName = "proxy2"

	testOldDockerRegistryCrn = "crn:cdp:docker:us-west-1:old-registry"
	testNewDockerRegistryCrn = "crn:cdp:docker:us-west-1:new-registry"

	testOldCredentialName  = "old-credential"
	testNewCredentialName  = "new-credential"
	testSameCredentialName = "same-credential"

	testGatewaySchemePublic  = "PUBLIC"
	testGatewaySchemePrivate = "PRIVATE"

	testOldCatalogURL  = "https://old-catalog.example.com"
	testNewCatalogURL  = "https://new-catalog.example.com"
	testSameCatalogURL = "https://same-catalog.example.com"

	testClusterInitFailed = "cluster init failed"

	testOldDefaultSG  = "sg-old-default"
	testNewDefaultSG  = "sg-new-default"
	testOldKnoxSG     = "sg-old-knox"
	testNewKnoxSG     = "sg-new-knox"
	testSameDefaultSG = "sg-same-default"
	testSameKnoxSG    = "sg-same-knox"
)

type commonTestFixture struct {
	ctx        context.Context
	mockClient *mocks.MockEnvironmentClientService
	client     *environmentsclient.Environments
	resp       *resource.UpdateResponse
}

func setupCommonTest(t *testing.T) commonTestFixture {
	t.Helper()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	return commonTestFixture{
		ctx:        context.TODO(),
		mockClient: mockClient,
		client:     NewMockEnvironments(mockClient),
		resp:       &resource.UpdateResponse{},
	}
}

// Tests for updateSshKeyIfChanged

func TestUpdateSshKeyIfChanged_KeyChanged_UpdatesStateAndCallsAPI(t *testing.T) {
	f := setupCommonTest(t)
	planKey := types.StringValue(testNewKey)
	stateKey := types.StringValue(testOldKey)

	f.mockClient.On("UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything).
		Return(&operations.UpdateSSHKeyOK{}, nil)

	result := updateSshKeyIfChanged(f.ctx, f.client, planKey, &stateKey, new(testEnvName), f.resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, testNewKey, stateKey.ValueString())
	f.mockClient.AssertCalled(t, "UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateSshKeyIfChanged_KeyUnchanged_SkipsAPICall(t *testing.T) {
	f := setupCommonTest(t)
	sameKey := types.StringValue(testSameKey)

	result := updateSshKeyIfChanged(f.ctx, f.client, sameKey, new(types.StringValue(testSameKey)), new(testEnvName), f.resp)

	assert.False(t, result.Diagnostics.HasError())
	f.mockClient.AssertNotCalled(t, "UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateSshKeyIfChanged_NullOrUnknownPlan_SkipsAPICall(t *testing.T) {
	tests := []struct {
		name    string
		planKey types.String
	}{
		{"null", types.StringNull()},
		{"unknown", types.StringUnknown()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := setupCommonTest(t)
			stateKey := types.StringValue(testOldKey)

			result := updateSshKeyIfChanged(f.ctx, f.client, tt.planKey, &stateKey, new(testEnvName), f.resp)

			assert.False(t, result.Diagnostics.HasError())
			assert.Equal(t, testOldKey, stateKey.ValueString())
			f.mockClient.AssertNotCalled(t, "UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything)
		})
	}
}

func TestUpdateSshKeyIfChanged_PlanKeyEmpty_AddsValidationError(t *testing.T) {
	f := setupCommonTest(t)
	stateKey := types.StringValue(testOldKey)

	result := updateSshKeyIfChanged(f.ctx, f.client, types.StringValue(""), &stateKey, new(testEnvName), f.resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Equal(t, testOldKey, stateKey.ValueString())
	f.mockClient.AssertNotCalled(t, "UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateSshKeyIfChanged_APIError_AddsDiagnosticError(t *testing.T) {
	f := setupCommonTest(t)
	stateKey := types.StringValue(testOldKey)

	f.mockClient.On("UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpdateSSHKeyOK)(nil), errors.New("API connection failed"))

	result := updateSshKeyIfChanged(f.ctx, f.client, types.StringValue(testNewKey), &stateKey, new(testEnvName), f.resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Equal(t, testOldKey, stateKey.ValueString())
}

// Tests for updateSshKey

func TestUpdateSshKey_Success(t *testing.T) {
	f := setupCommonTest(t)

	f.mockClient.On("UpdateSSHKeyContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateSSHKeyParams) bool {
		return params.Input != nil &&
			params.Input.NewPublicKey == testKey &&
			*params.Input.Environment == testEnvName
	}), mock.Anything).Return(&operations.UpdateSSHKeyOK{}, nil)

	err := updateSshKey(f.ctx, f.client, types.StringValue(testKey), new(testEnvName))

	assert.NoError(t, err)
	f.mockClient.AssertExpectations(t)
}

func TestUpdateSshKey_ReturnsError(t *testing.T) {
	f := setupCommonTest(t)

	f.mockClient.On("UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpdateSSHKeyOK)(nil), errors.New(testServiceUnavailable))

	err := updateSshKey(f.ctx, f.client, types.StringValue(testKey), new(testEnvName))

	assert.EqualError(t, err, testServiceUnavailable)
}

func TestUpdateSshKey_NullOrUnknownOrEmpty_DoesNotCallAPI(t *testing.T) {
	tests := []struct {
		name      string
		publicKey types.String
	}{
		{"null", types.StringNull()},
		{"unknown", types.StringUnknown()},
		{"empty", types.StringValue("")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := setupCommonTest(t)

			err := updateSshKey(f.ctx, f.client, tt.publicKey, new(testEnvName))

			assert.NoError(t, err)
			f.mockClient.AssertNotCalled(t, "UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything)
		})
	}
}

// Test helpers for performEnvironmentUpdate tests

type testUpdateModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

var testUpdateSchema = rsschema.Schema{
	Attributes: map[string]rsschema.Attribute{
		"name":  rsschema.StringAttribute{Required: true},
		"value": rsschema.StringAttribute{Optional: true},
	},
}

func createTestUpdateRaw(name, value string) tftypes.Value {
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"name":  tftypes.String,
				"value": tftypes.String,
			},
		},
		map[string]tftypes.Value{
			"name":  tftypes.NewValue(tftypes.String, name),
			"value": tftypes.NewValue(tftypes.String, value),
		},
	)
}

func createTestUpdateReqResp(planName, planValue, stateName, stateValue string) (resource.UpdateRequest, *resource.UpdateResponse) {
	req := resource.UpdateRequest{
		Plan: tfsdk.Plan{
			Raw:    createTestUpdateRaw(planName, planValue),
			Schema: testUpdateSchema,
		},
		State: tfsdk.State{
			Raw:    createTestUpdateRaw(stateName, stateValue),
			Schema: testUpdateSchema,
		},
	}
	resp := &resource.UpdateResponse{
		State: tfsdk.State{
			Schema: testUpdateSchema,
		},
	}
	return req, resp
}

func TestPerformEnvironmentUpdate_CallsUpdateWithDecodedPlanAndState(t *testing.T) {
	ctx := context.TODO()
	req, resp := createTestUpdateReqResp(testPlanName, testPlanValue, testStateName, testStateValue)

	var capturedPlan, capturedState *testUpdateModel
	updateFn := func(_ context.Context, plan *testUpdateModel, state *testUpdateModel, _ *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
		capturedPlan = plan
		capturedState = state
		return resp
	}

	performEnvironmentUpdate(ctx, req, resp, nil, updateFn)

	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, capturedPlan)
	assert.Equal(t, testPlanName, capturedPlan.Name.ValueString())
	assert.Equal(t, testPlanValue, capturedPlan.Value.ValueString())
	assert.NotNil(t, capturedState)
	assert.Equal(t, testStateName, capturedState.Name.ValueString())
	assert.Equal(t, testStateValue, capturedState.Value.ValueString())
}

func TestPerformEnvironmentUpdate_StateModifiedByUpdateIsPersisted(t *testing.T) {
	ctx := context.TODO()
	req, resp := createTestUpdateReqResp(testNewName, testNewValue, testOldName, testOldValue)

	updateFn := func(_ context.Context, plan *testUpdateModel, state *testUpdateModel, _ *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
		state.Name = plan.Name
		state.Value = plan.Value
		return resp
	}

	performEnvironmentUpdate(ctx, req, resp, nil, updateFn)

	assert.False(t, resp.Diagnostics.HasError())

	var finalState testUpdateModel
	resp.State.Get(ctx, &finalState)
	assert.Equal(t, testNewName, finalState.Name.ValueString())
	assert.Equal(t, testNewValue, finalState.Value.ValueString())
}

func TestPerformEnvironmentUpdate_PlanDecodeError_ReturnsEarlyWithoutCallingUpdate(t *testing.T) {
	ctx := context.TODO()

	invalidRaw := tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"name":  tftypes.String,
				"value": tftypes.Bool,
			},
		},
		map[string]tftypes.Value{
			"name":  tftypes.NewValue(tftypes.String, testStateName),
			"value": tftypes.NewValue(tftypes.Bool, true),
		},
	)

	req := resource.UpdateRequest{
		Plan: tfsdk.Plan{
			Raw:    invalidRaw,
			Schema: testUpdateSchema,
		},
		State: tfsdk.State{
			Raw:    createTestUpdateRaw(testStateName, testStateValue),
			Schema: testUpdateSchema,
		},
	}
	resp := &resource.UpdateResponse{
		State: tfsdk.State{
			Schema: testUpdateSchema,
		},
	}

	updateCalled := false
	updateFn := func(_ context.Context, _ *testUpdateModel, _ *testUpdateModel, _ *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
		updateCalled = true
		return resp
	}

	performEnvironmentUpdate(ctx, req, resp, nil, updateFn)

	assert.True(t, resp.Diagnostics.HasError())
	assert.False(t, updateCalled)
}

func TestPerformEnvironmentUpdate_UpdateAddsError_StateSetsOnceAndReturns(t *testing.T) {
	ctx := context.TODO()
	req, resp := createTestUpdateReqResp(testPlanName, testPlanValue, testStateName, testStateValue)

	updateFn := func(_ context.Context, plan *testUpdateModel, state *testUpdateModel, _ *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
		state.Name = plan.Name
		state.Value = plan.Value
		resp.Diagnostics.AddError("update failed", "something went wrong")
		return resp
	}

	performEnvironmentUpdate(ctx, req, resp, nil, updateFn)

	assert.True(t, resp.Diagnostics.HasError())

	var finalState testUpdateModel
	resp.State.Get(ctx, &finalState)
	assert.Equal(t, testPlanName, finalState.Name.ValueString())
	assert.Equal(t, testPlanValue, finalState.Value.ValueString())
}

func TestPerformEnvironmentUpdate_PassesClientToUpdateFunction(t *testing.T) {
	ctx := context.TODO()
	f := setupCommonTest(t)
	req, resp := createTestUpdateReqResp(testPlanName, testPlanValue, testPlanName, testPlanValue)

	var capturedClient *environmentsclient.Environments
	updateFn := func(_ context.Context, _ *testUpdateModel, _ *testUpdateModel, c *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
		capturedClient = c
		return resp
	}

	performEnvironmentUpdate(ctx, req, resp, f.client, updateFn)

	assert.False(t, resp.Diagnostics.HasError())
	assert.Same(t, f.client, capturedClient)
}

// Tests for updateProxyConfigurationIfChanged

func TestUpdateProxyConfigurationIfChanged_NoChange_SkipsAPICall(t *testing.T) {
	f := setupCommonTest(t)

	result := updateProxyConfigurationIfChanged(f.ctx, f.client, new(types.StringValue(testOldProxyConfigName)), new(types.StringValue(testOldProxyConfigName)), new(testEnvName), f.resp)

	assert.False(t, result.Diagnostics.HasError())
	f.mockClient.AssertNotCalled(t, "UpdateProxyConfigContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateProxyConfigurationIfChanged_ProxyRemoved_RemoveProxyTrue(t *testing.T) {
	tests := []struct {
		name    string
		planVal types.String
	}{
		{"to empty string", types.StringValue("")},
		{"to null", types.StringNull()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := setupCommonTest(t)

			f.mockClient.On("UpdateProxyConfigContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateProxyConfigParams) bool {
				return params.Input != nil &&
					params.Input.ProxyConfigName == "" &&
					params.Input.RemoveProxy == true &&
					*params.Input.Environment == testEnvName
			}), mock.Anything).Return(&operations.UpdateProxyConfigOK{}, nil)

			result := updateProxyConfigurationIfChanged(f.ctx, f.client, new(types.StringValue(testOldProxyConfigName)), new(tt.planVal), new(testEnvName), f.resp)

			assert.False(t, result.Diagnostics.HasError())
			f.mockClient.AssertExpectations(t)
		})
	}
}

func TestUpdateProxyConfigurationIfChanged_SkipScenarios(t *testing.T) {
	tests := []struct {
		name  string
		state *types.String
		plan  *types.String
	}{
		{"plan nil", new(types.StringValue(testOldProxyConfigName)), nil},
		{"plan unknown", new(types.StringValue(testOldProxyConfigName)), new(types.StringUnknown())},
		{"state nil", nil, new(types.StringValue(testNewProxyConfigName))},
		{"both null", new(types.StringNull()), new(types.StringNull())},
		{"both empty", new(types.StringValue("")), new(types.StringValue(""))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := setupCommonTest(t)

			result := updateProxyConfigurationIfChanged(f.ctx, f.client, tt.state, tt.plan, new(testEnvName), f.resp)

			assert.False(t, result.Diagnostics.HasError())
			f.mockClient.AssertNotCalled(t, "UpdateProxyConfigContext", mock.Anything, mock.Anything, mock.Anything)
		})
	}
}

func TestUpdateProxyConfigurationIfChanged_SetsNewProxy(t *testing.T) {
	tests := []struct {
		name     string
		stateVal types.String
	}{
		{"from existing proxy", types.StringValue(testOldProxyConfigName)},
		{"from null", types.StringNull()},
		{"from empty", types.StringValue("")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := setupCommonTest(t)

			f.mockClient.On("UpdateProxyConfigContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateProxyConfigParams) bool {
				return params.Input != nil &&
					params.Input.ProxyConfigName == testNewProxyConfigName &&
					params.Input.RemoveProxy == false &&
					*params.Input.Environment == testEnvName
			}), mock.Anything).Return(&operations.UpdateProxyConfigOK{}, nil)

			result := updateProxyConfigurationIfChanged(f.ctx, f.client, new(tt.stateVal), new(types.StringValue(testNewProxyConfigName)), new(testEnvName), f.resp)

			assert.False(t, result.Diagnostics.HasError())
			f.mockClient.AssertExpectations(t)
		})
	}
}

func TestUpdateProxyConfigurationIfChanged_Success_UpdatesState(t *testing.T) {
	f := setupCommonTest(t)

	f.mockClient.On("UpdateProxyConfigContext", mock.Anything, mock.Anything, mock.Anything).
		Return(&operations.UpdateProxyConfigOK{}, nil)

	state := types.StringValue(testOldProxyConfigName)
	result := updateProxyConfigurationIfChanged(f.ctx, f.client, &state, new(types.StringValue(testNewProxyConfigName)), new(testEnvName), f.resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, testNewProxyConfigName, state.ValueString())
}

func TestUpdateProxyConfigurationIfChanged_APIError_AddsDiagnosticAndPreservesState(t *testing.T) {
	f := setupCommonTest(t)

	f.mockClient.On("UpdateProxyConfigContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpdateProxyConfigOK)(nil), errors.New("API error"))

	state := types.StringValue(testOldProxyConfigName)
	result := updateProxyConfigurationIfChanged(f.ctx, f.client, &state, new(types.StringValue(testNewProxyConfigName)), new(testEnvName), f.resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Equal(t, testOldProxyConfigName, state.ValueString())
}

// Tests for updateCustomDockerRegistryIfChanged

func TestUpdateCustomDockerRegistryIfChanged_NoChange_SkipsAPICall(t *testing.T) {
	f := setupCommonTest(t)
	state := &CustomDockerRegistry{Crn: types.StringValue(testOldDockerRegistryCrn)}
	plan := &CustomDockerRegistry{Crn: types.StringValue(testOldDockerRegistryCrn)}

	result := updateCustomDockerRegistryIfChanged(f.ctx, f.client, state, plan, new(testEnvName), f.resp)

	assert.False(t, result.Diagnostics.HasError())
	f.mockClient.AssertNotCalled(t, "UpdateCustomDockerRegistryContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateCustomDockerRegistryIfChanged_Changed_CallsAPIAndUpdatesState(t *testing.T) {
	f := setupCommonTest(t)
	state := &CustomDockerRegistry{Crn: types.StringValue(testOldDockerRegistryCrn)}
	plan := &CustomDockerRegistry{Crn: types.StringValue(testNewDockerRegistryCrn)}

	f.mockClient.On("UpdateCustomDockerRegistryContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateCustomDockerRegistryParams) bool {
		return params.Input != nil &&
			*params.Input.CustomDockerRegistry == testNewDockerRegistryCrn &&
			*params.Input.Environment == testEnvName
	}), mock.Anything).Return(&operations.UpdateCustomDockerRegistryOK{}, nil)

	result := updateCustomDockerRegistryIfChanged(f.ctx, f.client, state, plan, new(testEnvName), f.resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, testNewDockerRegistryCrn, state.Crn.ValueString())
	f.mockClient.AssertExpectations(t)
}

func TestUpdateCustomDockerRegistryIfChanged_SkipScenarios(t *testing.T) {
	tests := []struct {
		name  string
		state *CustomDockerRegistry
		plan  *CustomDockerRegistry
	}{
		{"plan nil", &CustomDockerRegistry{Crn: types.StringValue(testOldDockerRegistryCrn)}, nil},
		{"state nil", nil, &CustomDockerRegistry{Crn: types.StringValue(testNewDockerRegistryCrn)}},
		{"plan CRN null", &CustomDockerRegistry{Crn: types.StringValue(testOldDockerRegistryCrn)}, &CustomDockerRegistry{Crn: types.StringNull()}},
		{"plan CRN unknown", &CustomDockerRegistry{Crn: types.StringValue(testOldDockerRegistryCrn)}, &CustomDockerRegistry{Crn: types.StringUnknown()}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := setupCommonTest(t)

			result := updateCustomDockerRegistryIfChanged(f.ctx, f.client, tt.state, tt.plan, new(testEnvName), f.resp)

			assert.False(t, result.Diagnostics.HasError())
			f.mockClient.AssertNotCalled(t, "UpdateCustomDockerRegistryContext", mock.Anything, mock.Anything, mock.Anything)
		})
	}
}

func TestUpdateCustomDockerRegistryIfChanged_PlanCrnEmpty_CallsAPI(t *testing.T) {
	f := setupCommonTest(t)
	state := &CustomDockerRegistry{Crn: types.StringValue(testOldDockerRegistryCrn)}
	plan := &CustomDockerRegistry{Crn: types.StringValue("")}

	f.mockClient.On("UpdateCustomDockerRegistryContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateCustomDockerRegistryParams) bool {
		return params.Input != nil &&
			*params.Input.CustomDockerRegistry == "" &&
			*params.Input.Environment == testEnvName
	}), mock.Anything).Return(&operations.UpdateCustomDockerRegistryOK{}, nil)

	result := updateCustomDockerRegistryIfChanged(f.ctx, f.client, state, plan, new(testEnvName), f.resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, "", state.Crn.ValueString())
	f.mockClient.AssertExpectations(t)
}

func TestUpdateCustomDockerRegistryIfChanged_APIError_AddsDiagnosticAndPreservesState(t *testing.T) {
	f := setupCommonTest(t)
	state := &CustomDockerRegistry{Crn: types.StringValue(testOldDockerRegistryCrn)}
	plan := &CustomDockerRegistry{Crn: types.StringValue(testNewDockerRegistryCrn)}

	f.mockClient.On("UpdateCustomDockerRegistryContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpdateCustomDockerRegistryOK)(nil), errors.New("API error"))

	result := updateCustomDockerRegistryIfChanged(f.ctx, f.client, state, plan, new(testEnvName), f.resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Equal(t, testOldDockerRegistryCrn, state.Crn.ValueString())
}

// Tests for updateCredential

func TestUpdateCredentialIfChanged_CredentialChanged_UpdatesStateAndCallsAPI(t *testing.T) {
	f := setupCommonTest(t)

	f.mockClient.On("ChangeEnvironmentCredentialContext", mock.Anything, mock.MatchedBy(func(params *operations.ChangeEnvironmentCredentialParams) bool {
		return *params.Input.CredentialName == testNewCredentialName && *params.Input.EnvironmentName == testEnvName
	}), mock.Anything).Return(&operations.ChangeEnvironmentCredentialOK{}, nil)

	state := types.StringValue(testOldCredentialName)
	result := updateCredential(f.ctx, f.client, types.StringValue(testNewCredentialName), &state, new(testEnvName), f.resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, testNewCredentialName, state.ValueString())
	f.mockClient.AssertExpectations(t)
}

func TestUpdateCredentialIfChanged_CredentialUnchanged_SkipsAPICall(t *testing.T) {
	f := setupCommonTest(t)

	state := types.StringValue(testSameCredentialName)
	result := updateCredential(f.ctx, f.client, types.StringValue(testSameCredentialName), &state, new(testEnvName), f.resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, testSameCredentialName, state.ValueString())
	f.mockClient.AssertNotCalled(t, "ChangeEnvironmentCredentialContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateCredentialIfChanged_APIError_AddsDiagnosticAndPreservesState(t *testing.T) {
	f := setupCommonTest(t)

	f.mockClient.On("ChangeEnvironmentCredentialContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.ChangeEnvironmentCredentialOK)(nil), errors.New("API connection failed"))

	state := types.StringValue(testOldCredentialName)
	result := updateCredential(f.ctx, f.client, types.StringValue(testNewCredentialName), &state, new(testEnvName), f.resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Equal(t, testOldCredentialName, state.ValueString())
}

// Tests for updateEndpointAccessGatewayIfChanged

func TestUpdateEndpointAccessGatewayIfChanged_SchemeChanged_CallsApiAndPolls(t *testing.T) {
	mockClient := new(mocks.MockEnvironmentClientService)

	mockClient.On("SetEndpointAccessGatewayContext", mock.Anything, mock.MatchedBy(func(params *operations.SetEndpointAccessGatewayParams) bool {
		return *params.Input.EndpointAccessGatewayScheme == testGatewaySchemePublic &&
			*params.Input.Environment == testEnvName &&
			len(params.Input.EndpointAccessGatewaySubnetIds) == 2
	}), mock.Anything).Return(&operations.SetEndpointAccessGatewayOK{
		Payload: &environmentsmodels.SetEndpointAccessGatewayResponse{OperationID: "op-123"},
	}, nil)

	mockClient.On("GetOperationContext", mock.Anything, mock.MatchedBy(func(params *operations.GetOperationParams) bool {
		return *params.Input.EnvironmentName == testEnvName && params.Input.OperationID == "op-123"
	}), mock.Anything).Return(&operations.GetOperationOK{
		Payload: &environmentsmodels.GetOperationResponse{OperationID: "op-123", OperationStatus: "FINISHED"},
	}, nil)

	resp := &resource.UpdateResponse{}
	updateEndpointAccessGatewayIfChanged(context.TODO(), NewMockEnvironments(mockClient),
		types.StringValue(testGatewaySchemePublic), utils.ToSetValueFromStringList([]string{"subnet-1", "subnet-2"}),
		new(types.StringValue(testGatewaySchemePrivate)), new(utils.ToSetValueFromStringList([]string{"subnet-1", "subnet-2"})),
		testEnvName, nil, resp)

	assert.False(t, resp.Diagnostics.HasError())
	mockClient.AssertExpectations(t)
}

func TestUpdateEndpointAccessGatewayIfChanged_SubnetIdsChanged_CallsApiAndPolls(t *testing.T) {
	mockClient := new(mocks.MockEnvironmentClientService)

	mockClient.On("SetEndpointAccessGatewayContext", mock.Anything, mock.Anything, mock.Anything).Return(&operations.SetEndpointAccessGatewayOK{
		Payload: &environmentsmodels.SetEndpointAccessGatewayResponse{OperationID: "op-456"},
	}, nil)
	mockClient.On("GetOperationContext", mock.Anything, mock.Anything, mock.Anything).Return(&operations.GetOperationOK{
		Payload: &environmentsmodels.GetOperationResponse{OperationID: "op-456", OperationStatus: "FINISHED"},
	}, nil)

	resp := &resource.UpdateResponse{}
	updateEndpointAccessGatewayIfChanged(context.TODO(), NewMockEnvironments(mockClient),
		types.StringValue(testGatewaySchemePublic), utils.ToSetValueFromStringList([]string{"subnet-1", "subnet-2", "subnet-3"}),
		new(types.StringValue(testGatewaySchemePublic)), new(utils.ToSetValueFromStringList([]string{"subnet-1", "subnet-2"})),
		testEnvName, nil, resp)

	assert.False(t, resp.Diagnostics.HasError())
	mockClient.AssertExpectations(t)
}

func TestUpdateEndpointAccessGatewayIfChanged_NoChangeOrNullScheme_DoesNotCallApi(t *testing.T) {
	tests := []struct {
		name        string
		planScheme  types.String
		stateScheme types.String
	}{
		{"nothing changed", types.StringValue(testGatewaySchemePublic), types.StringValue(testGatewaySchemePublic)},
		{"plan scheme null", types.StringNull(), types.StringValue(testGatewaySchemePublic)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := new(mocks.MockEnvironmentClientService)
			resp := &resource.UpdateResponse{}

			updateEndpointAccessGatewayIfChanged(context.TODO(), NewMockEnvironments(mockClient),
				tt.planScheme, utils.ToSetValueFromStringList([]string{"subnet-1"}),
				new(tt.stateScheme), new(utils.ToSetValueFromStringList([]string{"subnet-1"})),
				testEnvName, nil, resp)

			assert.False(t, resp.Diagnostics.HasError())
			mockClient.AssertNotCalled(t, "SetEndpointAccessGatewayContext", mock.Anything, mock.Anything, mock.Anything)
		})
	}
}

func TestUpdateEndpointAccessGatewayIfChanged_ErrorAndEdgeCases(t *testing.T) {
	tests := []struct {
		name            string
		setupMocks      func(*mocks.MockEnvironmentClientService)
		wantError       bool
		assertNoPolling bool
	}{
		{
			"API error",
			func(m *mocks.MockEnvironmentClientService) {
				m.On("SetEndpointAccessGatewayContext", mock.Anything, mock.Anything, mock.Anything).
					Return((*operations.SetEndpointAccessGatewayOK)(nil), errors.New("API connection failed"))
			},
			true, false,
		},
		{
			"operation failed",
			func(m *mocks.MockEnvironmentClientService) {
				m.On("SetEndpointAccessGatewayContext", mock.Anything, mock.Anything, mock.Anything).Return(&operations.SetEndpointAccessGatewayOK{
					Payload: &environmentsmodels.SetEndpointAccessGatewayResponse{OperationID: "op-fail"},
				}, nil)
				m.On("GetOperationContext", mock.Anything, mock.Anything, mock.Anything).Return(&operations.GetOperationOK{
					Payload: &environmentsmodels.GetOperationResponse{OperationID: "op-fail", OperationStatus: "FAILED"},
				}, nil)
			},
			true, false,
		},
		{
			"no operation ID skips polling",
			func(m *mocks.MockEnvironmentClientService) {
				m.On("SetEndpointAccessGatewayContext", mock.Anything, mock.Anything, mock.Anything).Return(&operations.SetEndpointAccessGatewayOK{
					Payload: &environmentsmodels.SetEndpointAccessGatewayResponse{},
				}, nil)
			},
			false, true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := new(mocks.MockEnvironmentClientService)
			tt.setupMocks(mockClient)

			resp := &resource.UpdateResponse{}
			updateEndpointAccessGatewayIfChanged(context.TODO(), NewMockEnvironments(mockClient),
				types.StringValue(testGatewaySchemePublic), utils.ToSetValueFromStringList([]string{"subnet-1"}),
				new(types.StringValue(testGatewaySchemePrivate)), new(utils.ToSetValueFromStringList([]string{"subnet-1"})),
				testEnvName, nil, resp)

			assert.Equal(t, tt.wantError, resp.Diagnostics.HasError())
			if tt.assertNoPolling {
				mockClient.AssertNotCalled(t, "GetOperationContext", mock.Anything, mock.Anything, mock.Anything)
			}
			mockClient.AssertExpectations(t)
		})
	}
}

// Tests for executeUpdateOperations

type updateOpsFixture struct {
	plan  *testUpdateModel
	state *testUpdateModel
	resp  *resource.UpdateResponse
}

func setupUpdateOpsTest() updateOpsFixture {
	return updateOpsFixture{
		plan:  &testUpdateModel{Name: types.StringValue("plan"), Value: types.StringValue("v1")},
		state: &testUpdateModel{Name: types.StringValue("state"), Value: types.StringValue("v0")},
		resp:  &resource.UpdateResponse{},
	}
}

func makeTrackingOp(callOrder *[]int, n int, failAt int) func(context.Context, *testUpdateModel, *testUpdateModel, *environmentsclient.Environments, *resource.UpdateResponse) *resource.UpdateResponse {
	return func(_ context.Context, _ *testUpdateModel, _ *testUpdateModel, _ *environmentsclient.Environments, r *resource.UpdateResponse) *resource.UpdateResponse {
		*callOrder = append(*callOrder, n)
		if n == failAt {
			r.Diagnostics.AddError("test error", "op failed")
		}
		return r
	}
}

func TestExecuteUpdateOperations_AllOpsSucceed_CallsAllInOrder(t *testing.T) {
	f := setupUpdateOpsTest()
	var callOrder []int

	result := executeUpdateOperations(context.TODO(), f.plan, f.state, nil, f.resp,
		makeTrackingOp(&callOrder, 1, -1), makeTrackingOp(&callOrder, 2, -1), makeTrackingOp(&callOrder, 3, -1))

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, []int{1, 2, 3}, callOrder)
}

func TestExecuteUpdateOperations_SecondOpFails_StopsAndReturnsError(t *testing.T) {
	f := setupUpdateOpsTest()
	var callOrder []int

	result := executeUpdateOperations(context.TODO(), f.plan, f.state, nil, f.resp,
		makeTrackingOp(&callOrder, 1, -1), makeTrackingOp(&callOrder, 2, 2), makeTrackingOp(&callOrder, 3, -1))

	assert.True(t, result.Diagnostics.HasError())
	assert.Equal(t, []int{1, 2}, callOrder)
}

func TestExecuteUpdateOperations_EmptyOps_ReturnsWithoutError(t *testing.T) {
	f := setupUpdateOpsTest()

	result := executeUpdateOperations[testUpdateModel](context.TODO(), f.plan, f.state, nil, f.resp)

	assert.False(t, result.Diagnostics.HasError())
}

func TestExecuteUpdateOperations_SingleOp_Success(t *testing.T) {
	f := setupUpdateOpsTest()

	called := false
	op := func(_ context.Context, p *testUpdateModel, s *testUpdateModel, _ *environmentsclient.Environments, r *resource.UpdateResponse) *resource.UpdateResponse {
		called = true
		assert.Equal(t, "plan", p.Name.ValueString())
		assert.Equal(t, "state", s.Name.ValueString())
		return r
	}

	result := executeUpdateOperations(context.TODO(), f.plan, f.state, nil, f.resp, op)

	assert.True(t, called)
	assert.False(t, result.Diagnostics.HasError())
}

// Tests for updateSecurityAccessIfChanged

func TestUpdateSecurityAccessIfChanged_CallsAPIAndHandlesResult(t *testing.T) {
	tests := []struct {
		name          string
		mockReturn    *operations.UpdateSecurityAccessOK
		mockErr       error
		wantError     bool
		wantDefaultSG string
		wantKnoxSG    string
	}{
		{
			"success updates state",
			&operations.UpdateSecurityAccessOK{}, nil,
			false, testNewDefaultSG, testNewKnoxSG,
		},
		{
			"API error preserves state",
			(*operations.UpdateSecurityAccessOK)(nil), errors.New(testServiceUnavailable),
			true, testOldDefaultSG, testOldKnoxSG,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := setupCommonTest(t)
			stateDefaultSG := types.StringValue(testOldDefaultSG)
			stateKnoxSG := types.StringValue(testOldKnoxSG)

			f.mockClient.On("UpdateSecurityAccessContext", mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mockReturn, tt.mockErr)

			result := updateSecurityAccessIfChanged(f.ctx, f.client, types.StringValue(testNewDefaultSG), types.StringValue(testNewKnoxSG), &stateDefaultSG, &stateKnoxSG, new(testEnvName), f.resp)

			assert.Equal(t, tt.wantError, result.Diagnostics.HasError())
			assert.Equal(t, tt.wantDefaultSG, stateDefaultSG.ValueString())
			assert.Equal(t, tt.wantKnoxSG, stateKnoxSG.ValueString())
		})
	}
}

func TestUpdateSecurityAccessIfChanged_Unchanged_SkipsAPICall(t *testing.T) {
	f := setupCommonTest(t)

	result := updateSecurityAccessIfChanged(f.ctx, f.client, types.StringValue(testSameDefaultSG), types.StringValue(testSameKnoxSG), new(types.StringValue(testSameDefaultSG)), new(types.StringValue(testSameKnoxSG)), new(testEnvName), f.resp)

	assert.False(t, result.Diagnostics.HasError())
	f.mockClient.AssertNotCalled(t, "UpdateSecurityAccessContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateSecurityAccessIfChanged_NullOrUnknownPlan_SkipsAPICall(t *testing.T) {
	tests := []struct {
		name        string
		planDefault types.String
		planKnox    types.String
	}{
		{"null", types.StringNull(), types.StringNull()},
		{"unknown", types.StringUnknown(), types.StringUnknown()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := setupCommonTest(t)
			stateDefaultSG := types.StringValue(testOldDefaultSG)
			stateKnoxSG := types.StringValue(testOldKnoxSG)

			result := updateSecurityAccessIfChanged(f.ctx, f.client, tt.planDefault, tt.planKnox, &stateDefaultSG, &stateKnoxSG, new(testEnvName), f.resp)

			assert.False(t, result.Diagnostics.HasError())
			assert.Equal(t, testOldDefaultSG, stateDefaultSG.ValueString())
			assert.Equal(t, testOldKnoxSG, stateKnoxSG.ValueString())
			f.mockClient.AssertNotCalled(t, "UpdateSecurityAccessContext", mock.Anything, mock.Anything, mock.Anything)
		})
	}
}

// Tests for updateSubnetIfChanged

func TestUpdateSubnetIfChanged_Changed_CallsAPIAndUpdatesState(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	planSubnets := utils.ToSetValueFromStringList([]string{testSubnet1, testSubnet2})
	stateSubnets := utils.ToSetValueFromStringList([]string{testSubnet1})
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateSubnetContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateSubnetParams) bool {
		return params.Input != nil &&
			*params.Input.Environment == testEnvName &&
			len(params.Input.SubnetIds) == 2
	}), mock.Anything).Return(&operations.UpdateSubnetOK{}, nil)

	result := updateSubnetIfChanged(ctx, client, planSubnets, &stateSubnets, new(testEnvName), resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, planSubnets, stateSubnets)
	mockClient.AssertExpectations(t)
}

func TestUpdateSubnetIfChanged_Unchanged_SkipsAPICall(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	subnets := utils.ToSetValueFromStringList([]string{testSubnet1, testSubnet2})
	stateSubnets := utils.ToSetValueFromStringList([]string{testSubnet1, testSubnet2})
	resp := &resource.UpdateResponse{}

	result := updateSubnetIfChanged(ctx, client, subnets, &stateSubnets, new(testEnvName), resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateSubnetContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateSubnetIfChanged_APIError_AddsDiagnosticError(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	planSubnets := utils.ToSetValueFromStringList([]string{testSubnet1, testSubnet2})
	stateSubnets := utils.ToSetValueFromStringList([]string{testSubnet1})
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateSubnetContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpdateSubnetOK)(nil), errors.New(testServiceUnavailable))

	result := updateSubnetIfChanged(ctx, client, planSubnets, &stateSubnets, new(testEnvName), resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Equal(t, utils.ToSetValueFromStringList([]string{testSubnet1}), stateSubnets)
}
