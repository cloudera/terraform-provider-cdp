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

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	datalakeclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client/operations"
	datalakemodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
	"github.com/cloudera/terraform-provider-cdp/mocks"
)

func TestApplyUpgradeOptions_Nil(t *testing.T) {
	req := &datalakemodels.UpgradeDatalakeRequest{DatalakeName: testDatalakeName_ptr()}
	applyUpgradeOptions(req, nil)
	assert.False(t, req.LockComponents)
	assert.False(t, req.SkipBackup)
}

func TestApplyUpgradeOptions_AllTrue(t *testing.T) {
	req := &datalakemodels.UpgradeDatalakeRequest{DatalakeName: testDatalakeName_ptr()}
	opts := &UpgradeOptions{
		DryRun:                types.BoolValue(true),
		LockComponents:        types.BoolValue(true),
		RollingUpgradeEnabled: types.BoolValue(true),
		SkipBackup:            types.BoolValue(true),
		SkipBackupValidation:  types.BoolValue(true),
		SkipAtlasMetadata:     types.BoolValue(true),
		SkipRangerAudits:      types.BoolValue(true),
		SkipRangerHmsMetadata: types.BoolValue(true),
		SkipDatahubValidation: types.BoolValue(true),
	}
	applyUpgradeOptions(req, opts)
	assert.True(t, req.DryRun)
	assert.True(t, req.LockComponents)
	assert.True(t, req.RollingUpgradeEnabled)
	assert.True(t, req.SkipBackup)
	assert.True(t, req.SkipBackupValidation)
	assert.True(t, req.SkipAtlasMetadata)
	assert.True(t, req.SkipRangerAudits)
	assert.True(t, req.SkipRangerHmsMetadata)
	assert.True(t, req.SkipDatahubValidation)
}

func TestUpgradeDatalake_ApiError(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDatalakeClientService(t)
	client := &datalakeclient.Datalake{Operations: mockClient}

	mockClient.On("UpgradeDatalakeContext", mock.Anything, mock.Anything, mock.Anything).
		Return((*operations.UpgradeDatalakeOK)(nil), errors.New("service unavailable"))

	resp := &resource.UpdateResponse{}
	UpgradeDatalake(ctx, client, testDatalakeName, testNewRuntime, testImageID, nil, nil, resp)

	assert.True(t, resp.Diagnostics.HasError())
}

func TestUpgradeDatalake_OperationFinished(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDatalakeClientService(t)
	client := &datalakeclient.Datalake{Operations: mockClient}

	mockClient.On("UpgradeDatalakeContext", mock.Anything, mock.Anything, mock.Anything).
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
	UpgradeDatalake(ctx, client, testDatalakeName, testNewRuntime, testImageID, nil, nil, resp)

	assert.False(t, resp.Diagnostics.HasError())
}

func TestUpgradeDatalake_OperationFailed(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDatalakeClientService(t)
	client := &datalakeclient.Datalake{Operations: mockClient}

	mockClient.On("UpgradeDatalakeContext", mock.Anything, mock.Anything, mock.Anything).
		Return(&operations.UpgradeDatalakeOK{
			Payload: &datalakemodels.UpgradeDatalakeResponse{OperationID: testOperationID},
		}, nil)

	mockClient.On("GetOperationContext", mock.Anything, mock.Anything, mock.Anything).
		Return(&operations.GetOperationOK{
			Payload: &datalakemodels.GetOperationResponse{OperationStatus: "FAILED"},
		}, nil)

	resp := &resource.UpdateResponse{}
	UpgradeDatalake(ctx, client, testDatalakeName, testNewRuntime, testImageID, nil, nil, resp)

	assert.True(t, resp.Diagnostics.HasError())
}

func TestUpgradeDatalake_NoOperationID_ReturnsError(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDatalakeClientService(t)
	client := &datalakeclient.Datalake{Operations: mockClient}

	mockClient.On("UpgradeDatalakeContext", mock.Anything, mock.Anything, mock.Anything).
		Return(&operations.UpgradeDatalakeOK{
			Payload: &datalakemodels.UpgradeDatalakeResponse{},
		}, nil)

	resp := &resource.UpdateResponse{}
	UpgradeDatalake(ctx, client, testDatalakeName, testNewRuntime, testImageID, nil, nil, resp)

	assert.True(t, resp.Diagnostics.HasError())
	assert.Contains(t, resp.Diagnostics.Errors()[0].Detail(), "no operation ID")
	mockClient.AssertNotCalled(t, "GetOperationContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpgradeDatalake_NoOperationID_WithReason_ReturnsError(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDatalakeClientService(t)
	client := &datalakeclient.Datalake{Operations: mockClient}

	mockClient.On("UpgradeDatalakeContext", mock.Anything, mock.Anything, mock.Anything).
		Return(&operations.UpgradeDatalakeOK{
			Payload: &datalakemodels.UpgradeDatalakeResponse{
				Reason: "No compatible images found for the given runtime version",
			},
		}, nil)

	resp := &resource.UpdateResponse{}
	UpgradeDatalake(ctx, client, testDatalakeName, testNewRuntime, testImageID, nil, nil, resp)

	assert.True(t, resp.Diagnostics.HasError())
	assert.Contains(t, resp.Diagnostics.Errors()[0].Detail(), "No compatible images found")
	mockClient.AssertNotCalled(t, "GetOperationContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpgradeDatalake_DryRun_SkipsPolling(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDatalakeClientService(t)
	client := &datalakeclient.Datalake{Operations: mockClient}

	opts := &UpgradeOptions{DryRun: types.BoolValue(true)}

	mockClient.On("UpgradeDatalakeContext", mock.Anything, mock.MatchedBy(func(params *operations.UpgradeDatalakeParams) bool {
		return params.Input != nil && params.Input.DryRun
	}), mock.Anything).
		Return(&operations.UpgradeDatalakeOK{
			Payload: &datalakemodels.UpgradeDatalakeResponse{
				Current: &datalakemodels.ImageInfo{
					ImageID:   "img-current",
					ImageName: "current-image",
					ComponentVersions: &datalakemodels.ImageComponentVersions{
						Cdp: "7.2.17",
						Cm:  "7.11.0",
						Os:  "centos7",
					},
				},
				UpgradeCandidates: []*datalakemodels.ImageInfo{
					{
						ImageID:          "img-1",
						ImageName:        "candidate-1",
						ImageCatalogName: "cdp-default",
						ComponentVersions: &datalakemodels.ImageComponentVersions{
							Cdp:          "7.2.18",
							Cm:           "7.11.1",
							Os:           "centos7",
							OsPatchLevel: "2026-09-01",
						},
					},
					{
						ImageID:   "img-2",
						ImageName: "candidate-2",
						ComponentVersions: &datalakemodels.ImageComponentVersions{
							Cdp: "7.3.0",
							Cm:  "7.12.0",
						},
					},
				},
			},
		}, nil)

	resp := &resource.UpdateResponse{}
	UpgradeDatalake(ctx, client, testDatalakeName, testNewRuntime, testImageID, opts, nil, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.Equal(t, 1, resp.Diagnostics.WarningsCount())

	detail := resp.Diagnostics.Warnings()[0].Detail()
	assert.Contains(t, detail, "Current image:")
	assert.Contains(t, detail, "Image ID: img-current")
	assert.Contains(t, detail, "CDP runtime: 7.2.17")
	assert.Contains(t, detail, "Upgrade candidates: 2")
	assert.Contains(t, detail, "Candidate 1:")
	assert.Contains(t, detail, "Image ID: img-1")
	assert.Contains(t, detail, "CDP runtime: 7.2.18")
	assert.Contains(t, detail, "OS: centos7 (patch: 2026-09-01)")
	assert.Contains(t, detail, "Candidate 2:")
	assert.Contains(t, detail, "Image ID: img-2")
	assert.Contains(t, detail, "CDP runtime: 7.3.0")

	mockClient.AssertNotCalled(t, "GetOperationContext", mock.Anything, mock.Anything, mock.Anything)
}

func TestUpgradeDatalake_DryRun_WithReason(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDatalakeClientService(t)
	client := &datalakeclient.Datalake{Operations: mockClient}

	opts := &UpgradeOptions{DryRun: types.BoolValue(true)}

	mockClient.On("UpgradeDatalakeContext", mock.Anything, mock.Anything, mock.Anything).
		Return(&operations.UpgradeDatalakeOK{
			Payload: &datalakemodels.UpgradeDatalakeResponse{
				Reason: "No compatible images found",
			},
		}, nil)

	resp := &resource.UpdateResponse{}
	UpgradeDatalake(ctx, client, testDatalakeName, testNewRuntime, testImageID, opts, nil, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.Equal(t, 1, resp.Diagnostics.WarningsCount())
	warnings := resp.Diagnostics.Warnings()
	assert.Contains(t, warnings[0].Detail(), "No compatible images found")
}

func TestSummarizeDryRunResponse_NilPayload(t *testing.T) {
	var diags diag.Diagnostics
	summarizeDryRunResponse(nil, &diags)

	assert.False(t, diags.HasError())
	assert.Equal(t, 1, diags.WarningsCount())
	assert.Contains(t, diags.Warnings()[0].Detail(), "No response payload returned.")
}

func TestSummarizeDryRunResponse_NoCandidates(t *testing.T) {
	var diags diag.Diagnostics
	summarizeDryRunResponse(&datalakemodels.UpgradeDatalakeResponse{
		Current: &datalakemodels.ImageInfo{
			ImageID:   "img-current",
			ImageName: "current-image",
			ComponentVersions: &datalakemodels.ImageComponentVersions{
				Cdp: "7.2.17",
			},
		},
		UpgradeCandidates: []*datalakemodels.ImageInfo{},
	}, &diags)

	assert.False(t, diags.HasError())
	assert.Equal(t, 1, diags.WarningsCount())

	detail := diags.Warnings()[0].Detail()
	assert.Contains(t, detail, "Current image:")
	assert.Contains(t, detail, "Image ID: img-current")
	assert.Contains(t, detail, "CDP runtime: 7.2.17")
	assert.Contains(t, detail, "Upgrade candidates: 0")
}

func TestUpgradeDatalake_WithUpgradeOptions(t *testing.T) {
	ctx := context.TODO()
	mockClient := mocks.NewMockDatalakeClientService(t)
	client := &datalakeclient.Datalake{Operations: mockClient}

	opts := &UpgradeOptions{
		RollingUpgradeEnabled: types.BoolValue(true),
		SkipBackup:            types.BoolValue(true),
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
	UpgradeDatalake(ctx, client, testDatalakeName, testNewRuntime, testImageID, opts, nil, resp)

	assert.False(t, resp.Diagnostics.HasError())
}
