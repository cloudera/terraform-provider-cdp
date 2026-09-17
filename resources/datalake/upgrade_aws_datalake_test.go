// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datalake

import (
	"context"
	"errors"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	datalakeclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client/operations"
	datalakemodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
	"github.com/cloudera/terraform-provider-cdp/mocks"
)

const (
	testDatalakeName = "test-datalake"
	testRuntime      = "7.2.17"
	testNewRuntime   = "7.2.18"
	testImageID      = "image-123"
	testNewImageID   = "image-456"
	testOperationID  = "op-abc-123"
)

func newMockDatalake(mockClient *mocks.MockDatalakeClientService) *datalakeclient.Datalake {
	return &datalakeclient.Datalake{Operations: mockClient}
}

func testDatalakeName_ptr() *string {
	s := testDatalakeName
	return &s
}

func TestUpgradeAwsRuntimeOrImageIfChanged_NoChange(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDatalakeClientService(t)
	client := newMockDatalake(mockClient)

	plan := &awsDatalakeResourceModel{
		DatalakeName: types.StringValue(testDatalakeName),
		Runtime:      types.StringValue(testRuntime),
		Image:        &awsDatalakeImage{ID: types.StringValue(testImageID)},
		UpgradeOptions: &UpgradeOptions{
			SkipBackup: types.BoolValue(true),
		},
	}
	state := &awsDatalakeResourceModel{
		DatalakeName: types.StringValue(testDatalakeName),
		Runtime:      types.StringValue(testRuntime),
		Image:        &awsDatalakeImage{ID: types.StringValue(testImageID)},
	}

	resp := &resource.UpdateResponse{}
	upgradeAwsRuntimeOrImageIfChanged(ctx, plan, state, client, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.Equal(t, plan.UpgradeOptions, state.UpgradeOptions)
}

func TestUpgradeAwsRuntimeOrImageIfChanged_RuntimeChanged(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDatalakeClientService(t)
	client := newMockDatalake(mockClient)

	plan := &awsDatalakeResourceModel{
		DatalakeName: types.StringValue(testDatalakeName),
		Runtime:      types.StringValue(testNewRuntime),
		Image:        &awsDatalakeImage{ID: types.StringValue(testImageID)},
	}
	state := &awsDatalakeResourceModel{
		DatalakeName: types.StringValue(testDatalakeName),
		Runtime:      types.StringValue(testRuntime),
		Image:        &awsDatalakeImage{ID: types.StringValue(testImageID)},
	}

	mockClient.On("UpgradeDatalakeContext", mock.Anything, mock.MatchedBy(func(params *operations.UpgradeDatalakeParams) bool {
		return params.Input != nil &&
			*params.Input.DatalakeName == testDatalakeName &&
			params.Input.Runtime == testNewRuntime
	}), mock.Anything).
		Return(&operations.UpgradeDatalakeOK{
			Payload: &datalakemodels.UpgradeDatalakeResponse{OperationID: testOperationID},
		}, nil)

	mockClient.On("GetOperationContext", mock.Anything, mock.MatchedBy(func(params *operations.GetOperationParams) bool {
		return params.Input != nil &&
			params.Input.DatalakeName == testDatalakeName &&
			params.Input.OperationID == testOperationID
	}), mock.Anything).
		Return(&operations.GetOperationOK{
			Payload: &datalakemodels.GetOperationResponse{OperationStatus: "FINISHED"},
		}, nil)

	resp := &resource.UpdateResponse{}
	upgradeAwsRuntimeOrImageIfChanged(ctx, plan, state, client, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.Equal(t, testNewRuntime, state.Runtime.ValueString())
}

func TestUpgradeAwsRuntimeOrImageIfChanged_ImageChanged(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDatalakeClientService(t)
	client := newMockDatalake(mockClient)

	plan := &awsDatalakeResourceModel{
		DatalakeName: types.StringValue(testDatalakeName),
		Runtime:      types.StringValue(testRuntime),
		Image:        &awsDatalakeImage{ID: types.StringValue(testNewImageID)},
	}
	state := &awsDatalakeResourceModel{
		DatalakeName: types.StringValue(testDatalakeName),
		Runtime:      types.StringValue(testRuntime),
		Image:        &awsDatalakeImage{ID: types.StringValue(testImageID)},
	}

	mockClient.On("UpgradeDatalakeContext", mock.Anything, mock.MatchedBy(func(params *operations.UpgradeDatalakeParams) bool {
		return params.Input != nil &&
			params.Input.ImageID == testNewImageID
	}), mock.Anything).
		Return(&operations.UpgradeDatalakeOK{
			Payload: &datalakemodels.UpgradeDatalakeResponse{OperationID: testOperationID},
		}, nil)

	mockClient.On("GetOperationContext", mock.Anything, mock.Anything, mock.Anything).
		Return(&operations.GetOperationOK{
			Payload: &datalakemodels.GetOperationResponse{OperationStatus: "FINISHED"},
		}, nil)

	resp := &resource.UpdateResponse{}
	upgradeAwsRuntimeOrImageIfChanged(ctx, plan, state, client, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.Equal(t, testNewImageID, state.Image.ID.ValueString())
}

func TestUpgradeAwsRuntimeOrImageIfChanged_DryRunDoesNotUpdateState(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDatalakeClientService(t)
	client := newMockDatalake(mockClient)

	plan := &awsDatalakeResourceModel{
		DatalakeName: types.StringValue(testDatalakeName),
		Runtime:      types.StringValue(testNewRuntime),
		Image:        &awsDatalakeImage{ID: types.StringValue(testNewImageID)},
		UpgradeOptions: &UpgradeOptions{
			DryRun: types.BoolValue(true),
		},
	}
	state := &awsDatalakeResourceModel{
		DatalakeName: types.StringValue(testDatalakeName),
		Runtime:      types.StringValue(testRuntime),
		Image:        &awsDatalakeImage{ID: types.StringValue(testImageID)},
	}

	mockClient.On("UpgradeDatalakeContext", mock.Anything, mock.MatchedBy(func(params *operations.UpgradeDatalakeParams) bool {
		return params.Input != nil && params.Input.DryRun
	}), mock.Anything).
		Return(&operations.UpgradeDatalakeOK{
			Payload: &datalakemodels.UpgradeDatalakeResponse{
				UpgradeCandidates: []*datalakemodels.ImageInfo{{ImageID: "img-1"}},
			},
		}, nil)

	resp := &resource.UpdateResponse{}
	upgradeAwsRuntimeOrImageIfChanged(ctx, plan, state, client, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.Equal(t, testRuntime, state.Runtime.ValueString(), "runtime should not change during dry run")
	assert.Equal(t, testImageID, state.Image.ID.ValueString(), "image should not change during dry run")
	assert.Equal(t, plan.UpgradeOptions, state.UpgradeOptions, "upgrade options should still sync")
}

func TestUpgradeAwsRuntimeOrImageIfChanged_ApiError_DoesNotUpdateState(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDatalakeClientService(t)
	client := newMockDatalake(mockClient)

	plan := &awsDatalakeResourceModel{
		DatalakeName: types.StringValue(testDatalakeName),
		Runtime:      types.StringValue(testNewRuntime),
		Image:        &awsDatalakeImage{ID: types.StringValue(testNewImageID)},
		UpgradeOptions: &UpgradeOptions{
			SkipBackup: types.BoolValue(true),
		},
	}
	state := &awsDatalakeResourceModel{
		DatalakeName: types.StringValue(testDatalakeName),
		Runtime:      types.StringValue(testRuntime),
		Image:        &awsDatalakeImage{ID: types.StringValue(testImageID)},
	}

	mockClient.On("UpgradeDatalakeContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpgradeDatalakeOK)(nil), errors.New("service unavailable"))

	resp := &resource.UpdateResponse{}
	upgradeAwsRuntimeOrImageIfChanged(ctx, plan, state, client, resp)

	assert.True(t, resp.Diagnostics.HasError())
	assert.Equal(t, testRuntime, state.Runtime.ValueString(), "runtime should not change on API error")
	assert.Equal(t, testImageID, state.Image.ID.ValueString(), "image should not change on API error")
	assert.Nil(t, state.UpgradeOptions, "upgrade options should not sync on API error")
}

func TestUpgradeAwsRuntimeOrImageIfChanged_NilPlanImage(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDatalakeClientService(t)
	client := newMockDatalake(mockClient)

	plan := &awsDatalakeResourceModel{
		DatalakeName: types.StringValue(testDatalakeName),
		Runtime:      types.StringValue(testRuntime),
		Image:        nil,
	}
	state := &awsDatalakeResourceModel{
		DatalakeName: types.StringValue(testDatalakeName),
		Runtime:      types.StringValue(testRuntime),
		Image:        &awsDatalakeImage{ID: types.StringValue(testImageID)},
	}

	mockClient.On("UpgradeDatalakeContext", mock.Anything, mock.MatchedBy(func(params *operations.UpgradeDatalakeParams) bool {
		return params.Input != nil && params.Input.ImageID == ""
	}), mock.Anything).
		Return(&operations.UpgradeDatalakeOK{
			Payload: &datalakemodels.UpgradeDatalakeResponse{OperationID: testOperationID},
		}, nil)

	mockClient.On("GetOperationContext", mock.Anything, mock.Anything, mock.Anything).
		Return(&operations.GetOperationOK{
			Payload: &datalakemodels.GetOperationResponse{OperationStatus: "FINISHED"},
		}, nil)

	resp := &resource.UpdateResponse{}
	upgradeAwsRuntimeOrImageIfChanged(ctx, plan, state, client, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.Nil(t, state.Image)
}

func TestUpgradeAwsRuntimeOrImageIfChanged_WithUpgradeOptions(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDatalakeClientService(t)
	client := newMockDatalake(mockClient)

	plan := &awsDatalakeResourceModel{
		DatalakeName: types.StringValue(testDatalakeName),
		Runtime:      types.StringValue(testNewRuntime),
		Image:        &awsDatalakeImage{ID: types.StringValue(testImageID)},
		UpgradeOptions: &UpgradeOptions{
			RollingUpgradeEnabled: types.BoolValue(true),
			SkipBackup:            types.BoolValue(true),
		},
	}
	state := &awsDatalakeResourceModel{
		DatalakeName: types.StringValue(testDatalakeName),
		Runtime:      types.StringValue(testRuntime),
		Image:        &awsDatalakeImage{ID: types.StringValue(testImageID)},
	}

	mockClient.On("UpgradeDatalakeContext", mock.Anything, mock.MatchedBy(func(params *operations.UpgradeDatalakeParams) bool {
		return params.Input != nil &&
			params.Input.RollingUpgradeEnabled &&
			params.Input.SkipBackup &&
			!params.Input.LockComponents
	}), mock.Anything).
		Return(&operations.UpgradeDatalakeOK{
			Payload: &datalakemodels.UpgradeDatalakeResponse{OperationID: testOperationID},
		}, nil)

	mockClient.On("GetOperationContext", mock.Anything, mock.Anything, mock.Anything).
		Return(&operations.GetOperationOK{
			Payload: &datalakemodels.GetOperationResponse{OperationStatus: "FINISHED"},
		}, nil)

	resp := &resource.UpdateResponse{}
	upgradeAwsRuntimeOrImageIfChanged(ctx, plan, state, client, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.Equal(t, plan.UpgradeOptions, state.UpgradeOptions)
}
