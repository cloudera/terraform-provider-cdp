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
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/mocks"
)

// Tests for updateGcpCredentialIfChanged

func TestUpdateGcpCredentialIfChanged_Changed_CallsAPI(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &gcpEnvironmentResourceModel{
		CredentialName:  types.StringValue(testNewCredentialName),
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		CredentialName: types.StringValue(testOldCredentialName),
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("ChangeEnvironmentCredentialContext", mock.Anything, mock.MatchedBy(func(params *operations.ChangeEnvironmentCredentialParams) bool {
		return params.Input != nil &&
			*params.Input.CredentialName == testNewCredentialName &&
			*params.Input.EnvironmentName == testEnvName
	}), mock.Anything).Return(&operations.ChangeEnvironmentCredentialOK{}, nil)

	result := updateGcpCredentialIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testNewCredentialName), state.CredentialName)
	mockClient.AssertExpectations(t)
}

func TestUpdateGcpCredentialIfChanged_Unchanged_Skips(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &gcpEnvironmentResourceModel{
		CredentialName:  types.StringValue(testSameCredentialName),
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		CredentialName: types.StringValue(testSameCredentialName),
	}
	resp := &resource.UpdateResponse{}

	result := updateGcpCredentialIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "ChangeEnvironmentCredentialContext", mock.Anything, mock.Anything, mock.Anything)
}

// Tests for updateGcpSshKeyIfChanged

func TestUpdateGcpSshKeyIfChanged_Changed_CallsAPI(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &gcpEnvironmentResourceModel{
		PublicKey:       types.StringValue(testNewKey),
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		PublicKey: types.StringValue(testOldKey),
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateSSHKeyContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateSSHKeyParams) bool {
		return params.Input != nil &&
			params.Input.NewPublicKey == testNewKey &&
			*params.Input.Environment == testEnvName
	}), mock.Anything).Return(&operations.UpdateSSHKeyOK{}, nil)

	result := updateGcpSshKeyIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testNewKey), state.PublicKey)
	mockClient.AssertExpectations(t)
}

func TestUpdateGcpSshKeyIfChanged_Unchanged_Skips(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &gcpEnvironmentResourceModel{
		PublicKey:       types.StringValue(testSameKey),
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		PublicKey: types.StringValue(testSameKey),
	}
	resp := &resource.UpdateResponse{}

	result := updateGcpSshKeyIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateSSHKeyContext", mock.Anything, mock.Anything, mock.Anything)
}

// Tests for updateGcpProxyConfigurationIfChanged

func TestUpdateGcpProxyConfigurationIfChanged_Changed_CallsAPI(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &gcpEnvironmentResourceModel{
		ProxyConfigName: types.StringValue(testNewProxyConfigName),
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		ProxyConfigName: types.StringValue(testOldProxyConfigName),
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateProxyConfigContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateProxyConfigParams) bool {
		return params.Input != nil &&
			params.Input.ProxyConfigName == testNewProxyConfigName &&
			*params.Input.Environment == testEnvName &&
			!params.Input.RemoveProxy
	}), mock.Anything).Return(&operations.UpdateProxyConfigOK{}, nil)

	result := updateGcpProxyConfigurationIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testNewProxyConfigName), state.ProxyConfigName)
	mockClient.AssertExpectations(t)
}

func TestUpdateGcpProxyConfigurationIfChanged_Unchanged_Skips(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &gcpEnvironmentResourceModel{
		ProxyConfigName: types.StringValue(testOldProxyConfigName),
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		ProxyConfigName: types.StringValue(testOldProxyConfigName),
	}
	resp := &resource.UpdateResponse{}

	result := updateGcpProxyConfigurationIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateProxyConfigContext", mock.Anything, mock.Anything, mock.Anything)
}

// Tests for updateGcpCustomDockerRegistryIfChanged

func TestUpdateGcpCustomDockerRegistryIfChanged_Changed_CallsAPI(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &gcpEnvironmentResourceModel{
		CustomDockerRegistry: &CustomDockerRegistry{Crn: types.StringValue(testNewDockerRegistryCrn)},
		EnvironmentName:      types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		CustomDockerRegistry: &CustomDockerRegistry{Crn: types.StringValue(testOldDockerRegistryCrn)},
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("UpdateCustomDockerRegistryContext", mock.Anything, mock.MatchedBy(func(params *operations.UpdateCustomDockerRegistryParams) bool {
		return params.Input != nil &&
			*params.Input.CustomDockerRegistry == testNewDockerRegistryCrn &&
			*params.Input.Environment == testEnvName
	}), mock.Anything).Return(&operations.UpdateCustomDockerRegistryOK{}, nil)

	result := updateGcpCustomDockerRegistryIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testNewDockerRegistryCrn), state.CustomDockerRegistry.Crn)
	mockClient.AssertExpectations(t)
}

func TestUpdateGcpCustomDockerRegistryIfChanged_Unchanged_Skips(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	registry := &CustomDockerRegistry{Crn: types.StringValue(testOldDockerRegistryCrn)}
	plan := &gcpEnvironmentResourceModel{
		CustomDockerRegistry: registry,
		EnvironmentName:      types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		CustomDockerRegistry: registry,
	}
	resp := &resource.UpdateResponse{}

	result := updateGcpCustomDockerRegistryIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "UpdateCustomDockerRegistryContext", mock.Anything, mock.Anything, mock.Anything)
}

// Tests for updateGcpEndpointAccessGatewayIfChanged

func TestUpdateGcpEndpointAccessGatewayIfChanged_Changed_CallsAPI(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &gcpEnvironmentResourceModel{
		EndpointAccessGatewayScheme:    types.StringValue(testGatewaySchemePublic),
		EndpointAccessGatewaySubnetIds: types.SetNull(types.StringType),
		EnvironmentName:                types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		EndpointAccessGatewayScheme:    types.StringValue(testGatewaySchemePrivate),
		EndpointAccessGatewaySubnetIds: types.SetNull(types.StringType),
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("SetEndpointAccessGatewayContext", mock.Anything, mock.MatchedBy(func(params *operations.SetEndpointAccessGatewayParams) bool {
		return params.Input != nil &&
			*params.Input.EndpointAccessGatewayScheme == testGatewaySchemePublic &&
			*params.Input.Environment == testEnvName
	}), mock.Anything).Return(&operations.SetEndpointAccessGatewayOK{
		Payload: &environmentsmodels.SetEndpointAccessGatewayResponse{},
	}, nil)

	result := updateGcpEndpointAccessGatewayIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	assert.Equal(t, types.StringValue(testGatewaySchemePublic), state.EndpointAccessGatewayScheme)
	mockClient.AssertExpectations(t)
}

func TestUpdateGcpEndpointAccessGatewayIfChanged_Unchanged_Skips(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &gcpEnvironmentResourceModel{
		EndpointAccessGatewayScheme:    types.StringValue(testGatewaySchemePublic),
		EndpointAccessGatewaySubnetIds: types.SetNull(types.StringType),
		EnvironmentName:                types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		EndpointAccessGatewayScheme:    types.StringValue(testGatewaySchemePublic),
		EndpointAccessGatewaySubnetIds: types.SetNull(types.StringType),
	}
	resp := &resource.UpdateResponse{}

	result := updateGcpEndpointAccessGatewayIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "SetEndpointAccessGatewayContext", mock.Anything, mock.Anything, mock.Anything)
}

// Tests for updateGcpCatalogIfChanged

func TestUpdateGcpCatalogIfChanged_Changed_CallsAPI(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &gcpEnvironmentResourceModel{
		FreeIpa:         newFreeIpaObject(testNewCatalogURL),
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		FreeIpa: newFreeIpaObject(testOldCatalogURL),
	}
	resp := &resource.UpdateResponse{}

	mockClient.On("SetCatalogContext", mock.Anything, mock.MatchedBy(func(params *operations.SetCatalogParams) bool {
		return params.Input != nil &&
			*params.Input.Catalog == testNewCatalogURL &&
			*params.Input.Environment == testEnvName
	})).Return(&operations.SetCatalogOK{}, nil)

	result := updateGcpCatalogIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertExpectations(t)
}

func TestUpdateGcpCatalogIfChanged_Unchanged_Skips(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockEnvironmentClientService(t)
	client := NewMockEnvironments(mockClient)

	plan := &gcpEnvironmentResourceModel{
		FreeIpa:         newFreeIpaObject(testSameCatalogURL),
		EnvironmentName: types.StringValue(testEnvName),
	}
	state := &gcpEnvironmentResourceModel{
		FreeIpa: newFreeIpaObject(testSameCatalogURL),
	}
	resp := &resource.UpdateResponse{}

	result := updateGcpCatalogIfChanged(ctx, plan, state, client, resp)

	assert.False(t, result.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "SetCatalogContext", mock.Anything, mock.Anything)
}
