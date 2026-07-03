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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/mocks"
)

type gcpUpdateTestFixture struct {
	ctx        context.Context
	mockClient *mocks.MockEnvironmentClientService
	client     *environmentsclient.Environments
	resp       *resource.UpdateResponse
}

func setupGcpUpdateTest(t *testing.T) gcpUpdateTestFixture {
	t.Helper()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	return gcpUpdateTestFixture{
		ctx:        context.TODO(),
		mockClient: mockClient,
		client:     NewMockEnvironments(mockClient),
		resp:       &resource.UpdateResponse{},
	}
}

// Tests for updateGcpCredentialIfChanged

func TestUpdateGcpCredentialIfChanged_Changed_CallsAPI(t *testing.T) {
	f := setupGcpUpdateTest(t)
	plan := &gcpEnvironmentResourceModel{
		CredentialName:  types.StringValue(testNewCredentialName),
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		CredentialName: types.StringValue(testOldCredentialName),
	}

	f.mockClient.On("ChangeEnvironmentCredentialContext", mock.Anything, mock.MatchedBy(func(params *operations.ChangeEnvironmentCredentialParams) bool {
		return params.Input != nil &&
			*params.Input.CredentialName == testNewCredentialName &&
			*params.Input.EnvironmentName == testEnvName
	}), mock.Anything).Return(&operations.ChangeEnvironmentCredentialOK{}, nil)

	result := updateGcpCredentialIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testNewCredentialName), state.CredentialName)
	f.mockClient.AssertExpectations(t)
}

func TestUpdateGcpCredentialIfChanged_Unchanged_Skips(t *testing.T) {
	f := setupGcpUpdateTest(t)
	plan := &gcpEnvironmentResourceModel{
		CredentialName:  types.StringValue(testSameCredentialName),
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		CredentialName: types.StringValue(testSameCredentialName),
	}

	result := updateGcpCredentialIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.False(t, result.Diagnostics.HasError())
	f.mockClient.AssertNotCalled(t, "ChangeEnvironmentCredentialContext", mock.Anything, mock.Anything, mock.Anything)
}

// Tests for updateGcpSshKeyIfChanged

func TestUpdateGcpSshKeyIfChanged_Changed_CallsAPI(t *testing.T) {
	f := setupGcpUpdateTest(t)
	plan := &gcpEnvironmentResourceModel{
		PublicKey:       types.StringValue(testNewKey),
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		PublicKey: types.StringValue(testOldKey),
	}

	f.mockClient.On("UpdateSSHKeyContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateSSHKeyParams) bool {
		return params.Input != nil &&
			params.Input.NewPublicKey == testNewKey &&
			*params.Input.Environment == testEnvName
	}), mock.Anything).Return(&operations.UpdateSSHKeyOK{}, nil)

	result := updateGcpSshKeyIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testNewKey), state.PublicKey)
	f.mockClient.AssertExpectations(t)
}

func TestUpdateGcpSshKeyIfChanged_Unchanged_Skips(t *testing.T) {
	f := setupGcpUpdateTest(t)
	plan := &gcpEnvironmentResourceModel{
		PublicKey:       types.StringValue(testSameKey),
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		PublicKey: types.StringValue(testSameKey),
	}

	result := updateGcpSshKeyIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.False(t, result.Diagnostics.HasError())
	f.mockClient.AssertNotCalled(t, "UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything)
}

// Tests for updateGcpProxyConfigurationIfChanged

func TestUpdateGcpProxyConfigurationIfChanged_Changed_CallsAPI(t *testing.T) {
	f := setupGcpUpdateTest(t)
	plan := &gcpEnvironmentResourceModel{
		ProxyConfigName: types.StringValue(testNewProxyConfigName),
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		ProxyConfigName: types.StringValue(testOldProxyConfigName),
	}

	f.mockClient.On("UpdateProxyConfigContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateProxyConfigParams) bool {
		return params.Input != nil &&
			params.Input.ProxyConfigName == testNewProxyConfigName &&
			*params.Input.Environment == testEnvName &&
			!params.Input.RemoveProxy
	}), mock.Anything).Return(&operations.UpdateProxyConfigOK{}, nil)

	result := updateGcpProxyConfigurationIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testNewProxyConfigName), state.ProxyConfigName)
	f.mockClient.AssertExpectations(t)
}

func TestUpdateGcpProxyConfigurationIfChanged_Unchanged_Skips(t *testing.T) {
	f := setupGcpUpdateTest(t)
	plan := &gcpEnvironmentResourceModel{
		ProxyConfigName: types.StringValue(testOldProxyConfigName),
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		ProxyConfigName: types.StringValue(testOldProxyConfigName),
	}

	result := updateGcpProxyConfigurationIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.False(t, result.Diagnostics.HasError())
	f.mockClient.AssertNotCalled(t, "UpdateProxyConfigContext", mock.Anything, mock.Anything, mock.Anything)
}

// Tests for updateGcpCustomDockerRegistryIfChanged

func TestUpdateGcpCustomDockerRegistryIfChanged_Changed_CallsAPI(t *testing.T) {
	f := setupGcpUpdateTest(t)
	plan := &gcpEnvironmentResourceModel{
		CustomDockerRegistry: &CustomDockerRegistry{Crn: types.StringValue(testNewDockerRegistryCrn)},
		EnvironmentName:      types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		CustomDockerRegistry: &CustomDockerRegistry{Crn: types.StringValue(testOldDockerRegistryCrn)},
	}

	f.mockClient.On("UpdateCustomDockerRegistryContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateCustomDockerRegistryParams) bool {
		return params.Input != nil &&
			*params.Input.CustomDockerRegistry == testNewDockerRegistryCrn &&
			*params.Input.Environment == testEnvName
	}), mock.Anything).Return(&operations.UpdateCustomDockerRegistryOK{}, nil)

	result := updateGcpCustomDockerRegistryIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testNewDockerRegistryCrn), state.CustomDockerRegistry.Crn)
	f.mockClient.AssertExpectations(t)
}

func TestUpdateGcpCustomDockerRegistryIfChanged_Unchanged_Skips(t *testing.T) {
	f := setupGcpUpdateTest(t)
	registry := &CustomDockerRegistry{Crn: types.StringValue(testOldDockerRegistryCrn)}
	plan := &gcpEnvironmentResourceModel{
		CustomDockerRegistry: registry,
		EnvironmentName:      types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		CustomDockerRegistry: registry,
	}

	result := updateGcpCustomDockerRegistryIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.False(t, result.Diagnostics.HasError())
	f.mockClient.AssertNotCalled(t, "UpdateCustomDockerRegistryContext", mock.Anything, mock.Anything, mock.Anything)
}

// Tests for updateGcpEndpointAccessGatewayIfChanged

func TestUpdateGcpEndpointAccessGatewayIfChanged_Changed_CallsAPI(t *testing.T) {
	f := setupGcpUpdateTest(t)
	plan := &gcpEnvironmentResourceModel{
		EndpointAccessGatewayScheme:    types.StringValue(testGatewaySchemePublic),
		EndpointAccessGatewaySubnetIds: types.SetNull(types.StringType),
		EnvironmentName:                types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		EndpointAccessGatewayScheme:    types.StringValue(testGatewaySchemePrivate),
		EndpointAccessGatewaySubnetIds: types.SetNull(types.StringType),
	}

	f.mockClient.On("SetEndpointAccessGatewayContext", mock.Anything, mock.MatchedBy(func(params *operations.SetEndpointAccessGatewayParams) bool {
		return params.Input != nil &&
			*params.Input.EndpointAccessGatewayScheme == testGatewaySchemePublic &&
			*params.Input.Environment == testEnvName
	}), mock.Anything).Return(&operations.SetEndpointAccessGatewayOK{
		Payload: &environmentsmodels.SetEndpointAccessGatewayResponse{},
	}, nil)

	result := updateGcpEndpointAccessGatewayIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testGatewaySchemePublic), state.EndpointAccessGatewayScheme)
	f.mockClient.AssertExpectations(t)
}

func TestUpdateGcpEndpointAccessGatewayIfChanged_Unchanged_Skips(t *testing.T) {
	f := setupGcpUpdateTest(t)
	plan := &gcpEnvironmentResourceModel{
		EndpointAccessGatewayScheme:    types.StringValue(testGatewaySchemePublic),
		EndpointAccessGatewaySubnetIds: types.SetNull(types.StringType),
		EnvironmentName:                types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		EndpointAccessGatewayScheme:    types.StringValue(testGatewaySchemePublic),
		EndpointAccessGatewaySubnetIds: types.SetNull(types.StringType),
	}

	result := updateGcpEndpointAccessGatewayIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.False(t, result.Diagnostics.HasError())
	f.mockClient.AssertNotCalled(t, "SetEndpointAccessGatewayContext", mock.Anything, mock.Anything, mock.Anything)
}

// Tests for updateGcpCatalogIfChanged

func TestUpdateGcpCatalogIfChanged_Changed_CallsAPI(t *testing.T) {
	f := setupGcpUpdateTest(t)
	plan := &gcpEnvironmentResourceModel{
		FreeIpa:         newFreeIpaObject(testNewCatalogURL),
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		FreeIpa: newFreeIpaObject(testOldCatalogURL),
	}

	f.mockClient.On("SetCatalogContext", mock.Anything, mock.MatchedBy(func(params *operations.SetCatalogParams) bool {
		return params.Input != nil &&
			*params.Input.Catalog == testNewCatalogURL &&
			*params.Input.Environment == testEnvName
	})).Return(&operations.SetCatalogOK{}, nil)

	result := updateGcpCatalogIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.False(t, result.Diagnostics.HasError())
	f.mockClient.AssertExpectations(t)
}

func TestUpdateGcpCatalogIfChanged_Unchanged_Skips(t *testing.T) {
	f := setupGcpUpdateTest(t)
	plan := &gcpEnvironmentResourceModel{
		FreeIpa:         newFreeIpaObject(testSameCatalogURL),
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		FreeIpa: newFreeIpaObject(testSameCatalogURL),
	}

	result := updateGcpCatalogIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.False(t, result.Diagnostics.HasError())
	f.mockClient.AssertNotCalled(t, "SetCatalogContext", mock.Anything, mock.Anything)
}

// Tests for updateGcpSecurityAccessIfChanged

func TestUpdateGcpSecurityAccessIfChanged_Changed_CallsAPI(t *testing.T) {
	f := setupGcpUpdateTest(t)
	plan := &gcpEnvironmentResourceModel{
		SecurityAccess: &GcpSecurityAccess{
			DefaultSecurityGroupId: types.StringValue(testNewDefaultSG),
			SecurityGroupIdForKnox: types.StringValue(testNewKnoxSG),
		},
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		SecurityAccess: &GcpSecurityAccess{
			DefaultSecurityGroupId: types.StringValue(testOldDefaultSG),
			SecurityGroupIdForKnox: types.StringValue(testOldKnoxSG),
		},
	}

	f.mockClient.On("UpdateSecurityAccessContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateSecurityAccessParams) bool {
		return params.Input != nil &&
			*params.Input.DefaultSecurityGroupID == testNewDefaultSG &&
			*params.Input.GatewayNodeSecurityGroupID == testNewKnoxSG &&
			*params.Input.Environment == testEnvName
	}), mock.Anything).Return(&operations.UpdateSecurityAccessOK{}, nil)

	result := updateGcpSecurityAccessIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testNewDefaultSG), state.SecurityAccess.DefaultSecurityGroupId)
	assert.Equal(t, types.StringValue(testNewKnoxSG), state.SecurityAccess.SecurityGroupIdForKnox)
	f.mockClient.AssertExpectations(t)
}

func TestUpdateGcpSecurityAccessIfChanged_Unchanged_SkipsAPICall(t *testing.T) {
	f := setupGcpUpdateTest(t)
	plan := &gcpEnvironmentResourceModel{
		SecurityAccess: &GcpSecurityAccess{
			DefaultSecurityGroupId: types.StringValue(testSameDefaultSG),
			SecurityGroupIdForKnox: types.StringValue(testSameKnoxSG),
		},
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		SecurityAccess: &GcpSecurityAccess{
			DefaultSecurityGroupId: types.StringValue(testSameDefaultSG),
			SecurityGroupIdForKnox: types.StringValue(testSameKnoxSG),
		},
	}

	result := updateGcpSecurityAccessIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.False(t, result.Diagnostics.HasError())
	f.mockClient.AssertNotCalled(t, "UpdateSecurityAccessContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateGcpSecurityAccessIfChanged_NilPlanSecurityAccess_SkipsAPICall(t *testing.T) {
	f := setupGcpUpdateTest(t)
	plan := &gcpEnvironmentResourceModel{
		SecurityAccess:  nil,
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		SecurityAccess: &GcpSecurityAccess{
			DefaultSecurityGroupId: types.StringValue(testOldDefaultSG),
			SecurityGroupIdForKnox: types.StringValue(testOldKnoxSG),
		},
	}

	result := updateGcpSecurityAccessIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.False(t, result.Diagnostics.HasError())
	f.mockClient.AssertNotCalled(t, "UpdateSecurityAccessContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateGcpSecurityAccessIfChanged_APIError_AddsDiagnostic(t *testing.T) {
	f := setupGcpUpdateTest(t)
	plan := &gcpEnvironmentResourceModel{
		SecurityAccess: &GcpSecurityAccess{
			DefaultSecurityGroupId: types.StringValue(testNewDefaultSG),
			SecurityGroupIdForKnox: types.StringValue(testNewKnoxSG),
		},
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		SecurityAccess: &GcpSecurityAccess{
			DefaultSecurityGroupId: types.StringValue(testOldDefaultSG),
			SecurityGroupIdForKnox: types.StringValue(testOldKnoxSG),
		},
	}

	f.mockClient.On("UpdateSecurityAccessContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpdateSecurityAccessOK)(nil), errors.New(testServiceUnavailable))

	result := updateGcpSecurityAccessIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testOldDefaultSG), state.SecurityAccess.DefaultSecurityGroupId)
	assert.Equal(t, types.StringValue(testOldKnoxSG), state.SecurityAccess.SecurityGroupIdForKnox)
}

// Tests for updateGcpAvailabilityZonesIfChanged

func TestUpdateGcpAvailabilityZonesIfChanged_Changed_CallsAPIAndUpdatesState(t *testing.T) {
	f := setupGcpUpdateTest(t)
	plan := &gcpEnvironmentResourceModel{
		AvailabilityZones: []types.String{types.StringValue("us-central1-a"), types.StringValue("us-central1-b"), types.StringValue("us-central1-c")},
		EnvironmentName:   types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		AvailabilityZones: []types.String{types.StringValue("us-central1-a"), types.StringValue("us-central1-b")},
	}

	f.mockClient.On("UpdateGcpAvailabilityZonesContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateGcpAvailabilityZonesParams) bool {
		return params.Input != nil &&
			*params.Input.Environment == testEnvName &&
			len(params.Input.AvailabilityZones) == 3
	}), mock.Anything).Return(&operations.UpdateGcpAvailabilityZonesOK{}, nil)

	result := updateGcpAvailabilityZonesIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, plan.AvailabilityZones, state.AvailabilityZones)
	f.mockClient.AssertExpectations(t)
}

func TestUpdateGcpAvailabilityZonesIfChanged_Unchanged_SkipsAPICall(t *testing.T) {
	f := setupGcpUpdateTest(t)
	zones := []types.String{types.StringValue("us-central1-a"), types.StringValue("us-central1-b")}
	plan := &gcpEnvironmentResourceModel{
		AvailabilityZones: zones,
		EnvironmentName:   types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		AvailabilityZones: zones,
	}

	result := updateGcpAvailabilityZonesIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.False(t, result.Diagnostics.HasError())
	f.mockClient.AssertNotCalled(t, "UpdateGcpAvailabilityZonesContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateGcpAvailabilityZonesIfChanged_NilPlan_SkipsAPICall(t *testing.T) {
	f := setupGcpUpdateTest(t)
	plan := &gcpEnvironmentResourceModel{
		AvailabilityZones: nil,
		EnvironmentName:   types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		AvailabilityZones: []types.String{types.StringValue("us-central1-a")},
	}

	result := updateGcpAvailabilityZonesIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.False(t, result.Diagnostics.HasError())
	f.mockClient.AssertNotCalled(t, "UpdateGcpAvailabilityZonesContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateGcpAvailabilityZonesIfChanged_EmptyPlan_AddsValidationError(t *testing.T) {
	f := setupGcpUpdateTest(t)
	plan := &gcpEnvironmentResourceModel{
		AvailabilityZones: []types.String{},
		EnvironmentName:   types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		AvailabilityZones: []types.String{types.StringValue("us-central1-a")},
	}

	result := updateGcpAvailabilityZonesIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.True(t, result.Diagnostics.HasError())
	f.mockClient.AssertNotCalled(t, "UpdateGcpAvailabilityZonesContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateGcpAvailabilityZonesIfChanged_APIError_AddsDiagnosticError(t *testing.T) {
	f := setupGcpUpdateTest(t)
	plan := &gcpEnvironmentResourceModel{
		AvailabilityZones: []types.String{types.StringValue("us-central1-a"), types.StringValue("us-central1-b")},
		EnvironmentName:   types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		AvailabilityZones: []types.String{types.StringValue("us-central1-a")},
	}

	f.mockClient.On("UpdateGcpAvailabilityZonesContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpdateGcpAvailabilityZonesOK)(nil), errors.New(testServiceUnavailable))

	result := updateGcpAvailabilityZonesIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Equal(t, []types.String{types.StringValue("us-central1-a")}, state.AvailabilityZones)
}

// Tests for updateGcpTelemetryFeaturesIfChanged

func TestUpdateGcpTelemetryFeaturesIfChanged_Changed_CallsAPIAndUpdatesState(t *testing.T) {
	f := setupGcpUpdateTest(t)
	plan := &gcpEnvironmentResourceModel{
		WorkloadAnalytics: types.BoolValue(true),
		EnvironmentName:   types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		WorkloadAnalytics: types.BoolValue(false),
	}

	f.mockClient.On("SetTelemetryFeaturesContext", mock.Anything, mock.MatchedBy(func(params *operations.SetTelemetryFeaturesParams) bool {
		return params.Input != nil &&
			*params.Input.EnvironmentName == testEnvName &&
			params.Input.WorkloadAnalytics == true
	}), mock.Anything).Return(&operations.SetTelemetryFeaturesOK{}, nil)

	result := updateGcpTelemetryFeaturesIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, types.BoolValue(true), state.WorkloadAnalytics)
	f.mockClient.AssertExpectations(t)
}

func TestUpdateGcpTelemetryFeaturesIfChanged_Unchanged_SkipsAPICall(t *testing.T) {
	f := setupGcpUpdateTest(t)
	plan := &gcpEnvironmentResourceModel{
		WorkloadAnalytics: types.BoolValue(true),
		EnvironmentName:   types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		WorkloadAnalytics: types.BoolValue(true),
	}

	result := updateGcpTelemetryFeaturesIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.False(t, result.Diagnostics.HasError())
	f.mockClient.AssertNotCalled(t, "SetTelemetryFeaturesContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpdateGcpTelemetryFeaturesIfChanged_NullOrUnknownPlan_SkipsAPICall(t *testing.T) {
	tests := []struct {
		name      string
		planValue types.Bool
	}{
		{"null", types.BoolNull()},
		{"unknown", types.BoolUnknown()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := setupGcpUpdateTest(t)
			plan := &gcpEnvironmentResourceModel{
				WorkloadAnalytics: tt.planValue,
				EnvironmentName:   types.StringValue(testEnvName),
			}
			state := &gcpEnvironmentResourceModel{
				WorkloadAnalytics: types.BoolValue(true),
			}

			result := updateGcpTelemetryFeaturesIfChanged(f.ctx, plan, state, f.client, f.resp)

			assert.False(t, result.Diagnostics.HasError())
			f.mockClient.AssertNotCalled(t, "SetTelemetryFeaturesContext", mock.Anything, mock.Anything, mock.Anything)
		})
	}
}

func TestUpdateGcpTelemetryFeaturesIfChanged_APIError_AddsDiagnosticError(t *testing.T) {
	f := setupGcpUpdateTest(t)
	plan := &gcpEnvironmentResourceModel{
		WorkloadAnalytics: types.BoolValue(false),
		EnvironmentName:   types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		WorkloadAnalytics: types.BoolValue(true),
	}

	f.mockClient.On("SetTelemetryFeaturesContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.SetTelemetryFeaturesOK)(nil), errors.New(testServiceUnavailable))

	result := updateGcpTelemetryFeaturesIfChanged(f.ctx, plan, state, f.client, f.resp)

	assert.True(t, result.Diagnostics.HasError())
	assert.Equal(t, types.BoolValue(true), state.WorkloadAnalytics)
}
