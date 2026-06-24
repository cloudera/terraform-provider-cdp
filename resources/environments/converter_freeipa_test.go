// Copyright 2025 Cloudera. All Rights Reserved.
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
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	"github.com/cloudera/terraform-provider-cdp/mocks"
)

func newFreeIpaObject(catalog string) types.Object {
	instances, _ := types.SetValueFrom(context.TODO(), FreeIpaInstanceType, []FreeIpaInstance{})
	recipes, _ := types.SetValueFrom(context.TODO(), types.StringType, []string{})
	obj, _ := basetypes.NewObjectValueFrom(context.TODO(), FreeIpaDetailsType.AttrTypes, &FreeIpaDetails{
		Catalog:              types.StringValue(catalog),
		ImageID:              types.StringValue("img-1"),
		Os:                   types.StringValue("centos7"),
		InstanceCountByGroup: types.Int32Value(1),
		InstanceType:         types.StringValue("m5.xlarge"),
		Instances:            instances,
		MultiAz:              types.BoolValue(false),
		Recipes:              recipes,
		Architecture:         types.StringValue("X86_64"),
	})
	return obj
}

func newFreeIpaObjectWithNullCatalog() types.Object {
	instances, _ := types.SetValueFrom(context.TODO(), FreeIpaInstanceType, []FreeIpaInstance{})
	recipes, _ := types.SetValueFrom(context.TODO(), types.StringType, []string{})
	obj, _ := basetypes.NewObjectValueFrom(context.TODO(), FreeIpaDetailsType.AttrTypes, &FreeIpaDetails{
		Catalog:              types.StringNull(),
		ImageID:              types.StringValue("img-1"),
		Os:                   types.StringValue("centos7"),
		InstanceCountByGroup: types.Int32Value(1),
		InstanceType:         types.StringValue("m5.xlarge"),
		Instances:            instances,
		MultiAz:              types.BoolValue(false),
		Recipes:              recipes,
		Architecture:         types.StringValue("X86_64"),
	})
	return obj
}

func newMockEnvClient(mockOps *mocks.MockEnvironmentClientService) *environmentsclient.Environments {
	return &environmentsclient.Environments{
		Operations: mockOps,
	}
}

func TestSetCatalogIfChanged_CatalogChanged_CallsSetCatalog(t *testing.T) {
	ctx := context.TODO()
	mockClient := new(mocks.MockEnvironmentClientService)

	planFreeIpa := newFreeIpaObject(testNewCatalogURL)
	stateFreeIpa := newFreeIpaObject(testOldCatalogURL)

	matcher := func(params *operations.SetCatalogParams) bool {
		return *params.Input.Catalog == testNewCatalogURL &&
			*params.Input.Environment == "test-env"
	}
	mockClient.On("SetCatalogContext", mock.Anything, mock.MatchedBy(matcher)).Return(&operations.SetCatalogOK{}, nil)

	resp := &resource.UpdateResponse{}
	updateCatalogIfChanged(ctx, planFreeIpa, &stateFreeIpa, "test-env", newMockEnvClient(mockClient), resp)

	assert.False(t, resp.Diagnostics.HasError())
	mockClient.AssertExpectations(t)

	var updatedDetails FreeIpaDetails
	asDiags := stateFreeIpa.As(ctx, &updatedDetails, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
	assert.False(t, asDiags.HasError())
	assert.Equal(t, testNewCatalogURL, updatedDetails.Catalog.ValueString())
}

func TestSetCatalogIfChanged_CatalogUnchanged_DoesNotCallSetCatalog(t *testing.T) {
	ctx := context.TODO()
	mockClient := new(mocks.MockEnvironmentClientService)

	planFreeIpa := newFreeIpaObject(testSameCatalogURL)
	resp := &resource.UpdateResponse{}
	updateCatalogIfChanged(ctx, planFreeIpa, new(newFreeIpaObject(testSameCatalogURL)), "test-env", newMockEnvClient(mockClient), resp)

	assert.False(t, resp.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "SetCatalog", mock.Anything)
}

func TestSetCatalogIfChanged_PlanCatalogNull_DoesNotCallSetCatalog(t *testing.T) {
	ctx := context.TODO()
	mockClient := new(mocks.MockEnvironmentClientService)

	planFreeIpa := newFreeIpaObjectWithNullCatalog()
	resp := &resource.UpdateResponse{}
	updateCatalogIfChanged(ctx, planFreeIpa, new(newFreeIpaObject(testOldCatalogURL)), "test-env", newMockEnvClient(mockClient), resp)

	assert.False(t, resp.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "SetCatalog", mock.Anything)
}

func TestSetCatalogIfChanged_ApiError_AddsDiagnostics(t *testing.T) {
	ctx := context.TODO()
	mockClient := new(mocks.MockEnvironmentClientService)

	planFreeIpa := newFreeIpaObject(testNewCatalogURL)
	stateFreeIpa := newFreeIpaObject(testOldCatalogURL)

	mockClient.On("SetCatalogContext", mock.Anything, mock.Anything).Return((*operations.SetCatalogOK)(nil), errors.New("API connection failed"))

	resp := &resource.UpdateResponse{}
	updateCatalogIfChanged(ctx, planFreeIpa, &stateFreeIpa, "test-env", newMockEnvClient(mockClient), resp)

	assert.True(t, resp.Diagnostics.HasError())
	mockClient.AssertExpectations(t)

	var stateDetails FreeIpaDetails
	asDiags := stateFreeIpa.As(ctx, &stateDetails, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
	assert.False(t, asDiags.HasError())
	assert.Equal(t, testOldCatalogURL, stateDetails.Catalog.ValueString())
}

func TestSetCatalogIfChanged_CatalogChangedFromNull_CallsSetCatalog(t *testing.T) {
	ctx := context.TODO()
	mockClient := new(mocks.MockEnvironmentClientService)

	planFreeIpa := newFreeIpaObject(testNewCatalogURL)
	stateFreeIpa := newFreeIpaObjectWithNullCatalog()

	matcher := func(params *operations.SetCatalogParams) bool {
		return *params.Input.Catalog == testNewCatalogURL &&
			*params.Input.Environment == "my-env"
	}
	mockClient.On("SetCatalogContext", mock.Anything, mock.MatchedBy(matcher)).Return(&operations.SetCatalogOK{}, nil)

	resp := &resource.UpdateResponse{}
	updateCatalogIfChanged(ctx, planFreeIpa, &stateFreeIpa, "my-env", newMockEnvClient(mockClient), resp)

	assert.False(t, resp.Diagnostics.HasError())
	mockClient.AssertExpectations(t)

	var updatedDetails FreeIpaDetails
	asDiags := stateFreeIpa.As(ctx, &updatedDetails, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
	assert.False(t, asDiags.HasError())
	assert.Equal(t, testNewCatalogURL, updatedDetails.Catalog.ValueString())
}
