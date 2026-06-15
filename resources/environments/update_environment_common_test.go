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
	"github.com/cloudera/terraform-provider-cdp/mocks"
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
)

func TestUpdateSshKeyIfChanged_KeyChanged_UpdatesStateAndCallsAPI(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	planKey := types.StringValue(testNewKey)
	stateKey := types.StringValue(testOldKey)
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything).
		Return(&operations.UpdateSSHKeyOK{}, nil)

	result := updateSshKeyIfChanged(ctx, client, planKey, &stateKey, new(testEnvName), resp)

	if result.Diagnostics.HasError() {
		t.Errorf("expected no errors, got: %v", result.Diagnostics.Errors())
	}
	if stateKey.ValueString() != testNewKey {
		t.Errorf("expected state key to be updated to plan key, got: %s", stateKey.ValueString())
	}
	mockClient.AssertCalled(t, "UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateSshKeyIfChanged_KeyUnchanged_SkipsAPICall(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	sameKey := types.StringValue(testSameKey)
	stateKey := types.StringValue(testSameKey)
	resp := &resource.UpdateResponse{}

	result := updateSshKeyIfChanged(ctx, client, sameKey, &stateKey, new(testEnvName), resp)

	if result.Diagnostics.HasError() {
		t.Errorf("expected no errors, got: %v", result.Diagnostics.Errors())
	}
	mockClient.AssertNotCalled(t, "UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateSshKeyIfChanged_PlanKeyNull_SkipsAPICall(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	planKey := types.StringNull()
	stateKey := types.StringValue(testOldKey)
	resp := &resource.UpdateResponse{}

	result := updateSshKeyIfChanged(ctx, client, planKey, &stateKey, new(testEnvName), resp)

	if result.Diagnostics.HasError() {
		t.Errorf("expected no errors, got: %v", result.Diagnostics.Errors())
	}
	if stateKey.ValueString() != testOldKey {
		t.Errorf("expected state key to remain unchanged, got: %s", stateKey.ValueString())
	}
	mockClient.AssertNotCalled(t, "UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateSshKeyIfChanged_PlanKeyUnknown_SkipsUpdate(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	planKey := types.StringUnknown()
	stateKey := types.StringValue(testOldKey)
	resp := &resource.UpdateResponse{}

	result := updateSshKeyIfChanged(ctx, client, planKey, &stateKey, new(testEnvName), resp)

	if result.Diagnostics.HasError() {
		t.Errorf("expected no errors, got: %v", result.Diagnostics.Errors())
	}
	if stateKey.ValueString() != testOldKey {
		t.Errorf("expected state key to remain unchanged, got: %s", stateKey.ValueString())
	}
	mockClient.AssertNotCalled(t, "UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateSshKeyIfChanged_PlanKeyEmpty_AddsValidationError(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	planKey := types.StringValue("")
	stateKey := types.StringValue(testOldKey)
	resp := &resource.UpdateResponse{}

	result := updateSshKeyIfChanged(ctx, client, planKey, &stateKey, new(testEnvName), resp)

	if !result.Diagnostics.HasError() {
		t.Errorf("expected diagnostics to contain a validation error")
	}
	if stateKey.ValueString() != testOldKey {
		t.Errorf("expected state key to remain unchanged, got: %s", stateKey.ValueString())
	}
	mockClient.AssertNotCalled(t, "UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateSshKeyIfChanged_APIError_AddsDiagnosticError(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	planKey := types.StringValue(testNewKey)
	stateKey := types.StringValue(testOldKey)
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpdateSSHKeyOK)(nil), errors.New("API connection failed"))

	result := updateSshKeyIfChanged(ctx, client, planKey, &stateKey, new(testEnvName), resp)

	if !result.Diagnostics.HasError() {
		t.Errorf("expected diagnostics to contain an error")
	}
	if stateKey.ValueString() != testOldKey {
		t.Errorf("expected state key to remain unchanged on error, got: %s", stateKey.ValueString())
	}
}

func TestUpdateSshKey_Success(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	publicKey := types.StringValue(testKey)

	mockClient.On("UpdateSSHKeyContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateSSHKeyParams) bool {
		return params.Input != nil &&
			params.Input.NewPublicKey == testKey &&
			*params.Input.Environment == testEnvName
	}), mock.Anything).
		Return(&operations.UpdateSSHKeyOK{}, nil)

	err := updateSshKey(ctx, client, publicKey, new(testEnvName))

	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	mockClient.AssertExpectations(t)
}

func TestUpdateSshKey_ReturnsError(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	publicKey := types.StringValue(testKey)

	mockClient.On("UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpdateSSHKeyOK)(nil), errors.New(testServiceUnavailable))

	err := updateSshKey(ctx, client, publicKey, new(testEnvName))

	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	if err.Error() != testServiceUnavailable {
		t.Errorf("expected error message 'service unavailable', got: %s", err.Error())
	}
}

func TestUpdateSshKey_NullPublicKey_DoesNotSetInput(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	publicKey := types.StringNull()

	err := updateSshKey(ctx, client, publicKey, new(testEnvName))

	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	mockClient.AssertNotCalled(t, "UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateSshKey_EmptyPublicKey_DoesNotSetInput(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	publicKey := types.StringValue("")

	err := updateSshKey(ctx, client, publicKey, new(testEnvName))

	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	mockClient.AssertNotCalled(t, "UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateSshKey_UnknownPublicKey_DoesNotCallAPI(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	publicKey := types.StringUnknown()

	err := updateSshKey(ctx, client, publicKey, new(testEnvName))

	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	mockClient.AssertNotCalled(t, "UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything)
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
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)
	req, resp := createTestUpdateReqResp(testPlanName, testPlanValue, testPlanName, testPlanValue)

	var capturedClient *environmentsclient.Environments
	updateFn := func(_ context.Context, _ *testUpdateModel, _ *testUpdateModel, c *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
		capturedClient = c
		return resp
	}

	performEnvironmentUpdate(ctx, req, resp, client, updateFn)

	assert.False(t, resp.Diagnostics.HasError())
	assert.Same(t, client, capturedClient)
}
