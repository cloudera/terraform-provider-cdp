// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package freeipa

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

const (
	testOldCatalogURL  = "https://old-catalog.example.com"
	testNewCatalogURL  = "https://new-catalog.example.com"
	testSameCatalogURL = "https://same-catalog.example.com"
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
	UpdateCatalogIfChanged(ctx, planFreeIpa, &stateFreeIpa, "test-env", newMockEnvClient(mockClient), resp)

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
	UpdateCatalogIfChanged(ctx, planFreeIpa, new(newFreeIpaObject(testSameCatalogURL)), "test-env", newMockEnvClient(mockClient), resp)

	assert.False(t, resp.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "SetCatalog", mock.Anything)
}

func TestSetCatalogIfChanged_PlanCatalogNull_DoesNotCallSetCatalog(t *testing.T) {
	ctx := context.TODO()
	mockClient := new(mocks.MockEnvironmentClientService)

	planFreeIpa := newFreeIpaObjectWithNullCatalog()
	resp := &resource.UpdateResponse{}
	UpdateCatalogIfChanged(ctx, planFreeIpa, new(newFreeIpaObject(testOldCatalogURL)), "test-env", newMockEnvClient(mockClient), resp)

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
	UpdateCatalogIfChanged(ctx, planFreeIpa, &stateFreeIpa, "test-env", newMockEnvClient(mockClient), resp)

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
	UpdateCatalogIfChanged(ctx, planFreeIpa, &stateFreeIpa, "my-env", newMockEnvClient(mockClient), resp)

	assert.False(t, resp.Diagnostics.HasError())
	mockClient.AssertExpectations(t)

	var updatedDetails FreeIpaDetails
	asDiags := stateFreeIpa.As(ctx, &updatedDetails, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
	assert.False(t, asDiags.HasError())
	assert.Equal(t, testNewCatalogURL, updatedDetails.Catalog.ValueString())
}

func newFreeIpaObjectWithRecipes(catalog string, recipes []string) types.Object {
	instances, _ := types.SetValueFrom(context.TODO(), FreeIpaInstanceType, []FreeIpaInstance{})
	var recipesSet types.Set
	if recipes == nil {
		recipesSet = types.SetNull(types.StringType)
	} else {
		recipesSet, _ = types.SetValueFrom(context.TODO(), types.StringType, recipes)
	}
	obj, _ := basetypes.NewObjectValueFrom(context.TODO(), FreeIpaDetailsType.AttrTypes, &FreeIpaDetails{
		Catalog:              types.StringValue(catalog),
		ImageID:              types.StringValue("img-1"),
		Os:                   types.StringValue("centos7"),
		InstanceCountByGroup: types.Int32Value(1),
		InstanceType:         types.StringValue("m5.xlarge"),
		Instances:            instances,
		MultiAz:              types.BoolValue(false),
		Recipes:              recipesSet,
		Architecture:         types.StringValue("X86_64"),
	})
	return obj
}

func TestUpdateRecipesIfChanged_RecipesAdded_CallsAttachOnly(t *testing.T) {
	ctx := context.TODO()
	mockClient := new(mocks.MockEnvironmentClientService)

	planFreeIpa := newFreeIpaObjectWithRecipes(testSameCatalogURL, []string{"recipe-a", "recipe-b", "recipe-c"})
	stateFreeIpa := newFreeIpaObjectWithRecipes(testSameCatalogURL, []string{"recipe-a", "recipe-b"})

	matcher := func(params *operations.AttachFreeIpaRecipesParams) bool {
		return *params.Input.Environment == "test-env" &&
			len(params.Input.Recipes) == 1 &&
			params.Input.Recipes[0] == "recipe-c"
	}
	mockClient.On("AttachFreeIpaRecipesContext", mock.Anything, mock.MatchedBy(matcher)).Return(&operations.AttachFreeIpaRecipesOK{}, nil)

	resp := &resource.UpdateResponse{}
	UpdateRecipesIfChanged(ctx, planFreeIpa, &stateFreeIpa, "test-env", newMockEnvClient(mockClient), resp)

	assert.False(t, resp.Diagnostics.HasError())
	mockClient.AssertExpectations(t)
	mockClient.AssertNotCalled(t, "DetachFreeIpaRecipesContext", mock.Anything, mock.Anything)

	var updatedDetails FreeIpaDetails
	asDiags := stateFreeIpa.As(ctx, &updatedDetails, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
	assert.False(t, asDiags.HasError())
	assert.True(t, updatedDetails.Recipes.Equal(planFreeIpa.Attributes()["recipes"].(types.Set)))
}

func TestUpdateRecipesIfChanged_RecipesRemoved_CallsDetachOnly(t *testing.T) {
	ctx := context.TODO()
	mockClient := new(mocks.MockEnvironmentClientService)

	planFreeIpa := newFreeIpaObjectWithRecipes(testSameCatalogURL, []string{"recipe-a"})
	stateFreeIpa := newFreeIpaObjectWithRecipes(testSameCatalogURL, []string{"recipe-a", "recipe-b"})

	matcher := func(params *operations.DetachFreeIpaRecipesParams) bool {
		return *params.Input.Environment == "test-env" &&
			len(params.Input.Recipes) == 1 &&
			params.Input.Recipes[0] == "recipe-b"
	}
	mockClient.On("DetachFreeIpaRecipesContext", mock.Anything, mock.MatchedBy(matcher)).Return(&operations.DetachFreeIpaRecipesOK{}, nil)

	resp := &resource.UpdateResponse{}
	UpdateRecipesIfChanged(ctx, planFreeIpa, &stateFreeIpa, "test-env", newMockEnvClient(mockClient), resp)

	assert.False(t, resp.Diagnostics.HasError())
	mockClient.AssertExpectations(t)
	mockClient.AssertNotCalled(t, "AttachFreeIpaRecipesContext", mock.Anything, mock.Anything)

	var updatedDetails FreeIpaDetails
	asDiags := stateFreeIpa.As(ctx, &updatedDetails, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
	assert.False(t, asDiags.HasError())
	assert.True(t, updatedDetails.Recipes.Equal(planFreeIpa.Attributes()["recipes"].(types.Set)))
}

func TestUpdateRecipesIfChanged_RecipesSwapped_CallsBoth(t *testing.T) {
	ctx := context.TODO()
	mockClient := new(mocks.MockEnvironmentClientService)

	planFreeIpa := newFreeIpaObjectWithRecipes(testSameCatalogURL, []string{"recipe-a", "recipe-c"})
	stateFreeIpa := newFreeIpaObjectWithRecipes(testSameCatalogURL, []string{"recipe-a", "recipe-b"})

	attachMatcher := func(params *operations.AttachFreeIpaRecipesParams) bool {
		return *params.Input.Environment == "test-env" &&
			len(params.Input.Recipes) == 1 &&
			params.Input.Recipes[0] == "recipe-c"
	}
	detachMatcher := func(params *operations.DetachFreeIpaRecipesParams) bool {
		return *params.Input.Environment == "test-env" &&
			len(params.Input.Recipes) == 1 &&
			params.Input.Recipes[0] == "recipe-b"
	}
	mockClient.On("AttachFreeIpaRecipesContext", mock.Anything, mock.MatchedBy(attachMatcher)).Return(&operations.AttachFreeIpaRecipesOK{}, nil)
	mockClient.On("DetachFreeIpaRecipesContext", mock.Anything, mock.MatchedBy(detachMatcher)).Return(&operations.DetachFreeIpaRecipesOK{}, nil)

	resp := &resource.UpdateResponse{}
	UpdateRecipesIfChanged(ctx, planFreeIpa, &stateFreeIpa, "test-env", newMockEnvClient(mockClient), resp)

	assert.False(t, resp.Diagnostics.HasError())
	mockClient.AssertExpectations(t)

	var updatedDetails FreeIpaDetails
	asDiags := stateFreeIpa.As(ctx, &updatedDetails, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
	assert.False(t, asDiags.HasError())
	assert.True(t, updatedDetails.Recipes.Equal(planFreeIpa.Attributes()["recipes"].(types.Set)))
}

func TestUpdateRecipesIfChanged_RecipesUnchanged_SkipsAPIs(t *testing.T) {
	ctx := context.TODO()
	mockClient := new(mocks.MockEnvironmentClientService)

	planFreeIpa := newFreeIpaObjectWithRecipes(testSameCatalogURL, []string{"recipe-a", "recipe-b"})
	resp := &resource.UpdateResponse{}
	UpdateRecipesIfChanged(ctx, planFreeIpa, new(newFreeIpaObjectWithRecipes(testSameCatalogURL, []string{"recipe-a", "recipe-b"})), "test-env", newMockEnvClient(mockClient), resp)

	assert.False(t, resp.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "AttachFreeIpaRecipesContext", mock.Anything, mock.Anything)
	mockClient.AssertNotCalled(t, "DetachFreeIpaRecipesContext", mock.Anything, mock.Anything)
}

func TestUpdateRecipesIfChanged_SameRecipesDifferentOrder_SkipsAPIs(t *testing.T) {
	ctx := context.TODO()
	mockClient := new(mocks.MockEnvironmentClientService)

	planFreeIpa := newFreeIpaObjectWithRecipes(testSameCatalogURL, []string{"recipe-b", "recipe-a"})
	resp := &resource.UpdateResponse{}
	UpdateRecipesIfChanged(ctx, planFreeIpa, new(newFreeIpaObjectWithRecipes(testSameCatalogURL, []string{"recipe-a", "recipe-b"})), "test-env", newMockEnvClient(mockClient), resp)

	assert.False(t, resp.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "AttachFreeIpaRecipesContext", mock.Anything, mock.Anything)
	mockClient.AssertNotCalled(t, "DetachFreeIpaRecipesContext", mock.Anything, mock.Anything)
}

func TestUpdateRecipesIfChanged_PlanRecipesNull_SkipsAPIs(t *testing.T) {
	ctx := context.TODO()
	mockClient := new(mocks.MockEnvironmentClientService)

	planFreeIpa := newFreeIpaObjectWithRecipes(testSameCatalogURL, nil)
	resp := &resource.UpdateResponse{}
	UpdateRecipesIfChanged(ctx, planFreeIpa, new(newFreeIpaObjectWithRecipes(testSameCatalogURL, []string{"recipe-a"})), "test-env", newMockEnvClient(mockClient), resp)

	assert.False(t, resp.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "AttachFreeIpaRecipesContext", mock.Anything, mock.Anything)
	mockClient.AssertNotCalled(t, "DetachFreeIpaRecipesContext", mock.Anything, mock.Anything)
}

func TestUpdateRecipesIfChanged_AttachFails_AddsDiagnosticsStateUnchanged(t *testing.T) {
	ctx := context.TODO()
	mockClient := new(mocks.MockEnvironmentClientService)

	planFreeIpa := newFreeIpaObjectWithRecipes(testSameCatalogURL, []string{"recipe-a", "recipe-b", "recipe-c"})
	stateFreeIpa := newFreeIpaObjectWithRecipes(testSameCatalogURL, []string{"recipe-a", "recipe-b"})

	mockClient.On("AttachFreeIpaRecipesContext", mock.Anything, mock.Anything).Return((*operations.AttachFreeIpaRecipesOK)(nil), errors.New("API connection failed"))

	resp := &resource.UpdateResponse{}
	UpdateRecipesIfChanged(ctx, planFreeIpa, &stateFreeIpa, "test-env", newMockEnvClient(mockClient), resp)

	assert.True(t, resp.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "DetachFreeIpaRecipesContext", mock.Anything, mock.Anything)

	var stateDetails FreeIpaDetails
	asDiags := stateFreeIpa.As(ctx, &stateDetails, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
	assert.False(t, asDiags.HasError())
	expectedState, _ := types.SetValueFrom(ctx, types.StringType, []string{"recipe-a", "recipe-b"})
	assert.True(t, stateDetails.Recipes.Equal(expectedState))
}

func TestUpdateRecipesIfChanged_DetachFails_AddsDiagnosticsStateUnchanged(t *testing.T) {
	ctx := context.TODO()
	mockClient := new(mocks.MockEnvironmentClientService)

	planFreeIpa := newFreeIpaObjectWithRecipes(testSameCatalogURL, []string{"recipe-a"})
	stateFreeIpa := newFreeIpaObjectWithRecipes(testSameCatalogURL, []string{"recipe-a", "recipe-b"})

	mockClient.On("DetachFreeIpaRecipesContext", mock.Anything, mock.Anything).Return((*operations.DetachFreeIpaRecipesOK)(nil), errors.New("API connection failed"))

	resp := &resource.UpdateResponse{}
	UpdateRecipesIfChanged(ctx, planFreeIpa, &stateFreeIpa, "test-env", newMockEnvClient(mockClient), resp)

	assert.True(t, resp.Diagnostics.HasError())
	mockClient.AssertNotCalled(t, "AttachFreeIpaRecipesContext", mock.Anything, mock.Anything)

	var stateDetails FreeIpaDetails
	asDiags := stateFreeIpa.As(ctx, &stateDetails, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
	assert.False(t, asDiags.HasError())
	expectedState, _ := types.SetValueFrom(ctx, types.StringType, []string{"recipe-a", "recipe-b"})
	assert.True(t, stateDetails.Recipes.Equal(expectedState))
}
