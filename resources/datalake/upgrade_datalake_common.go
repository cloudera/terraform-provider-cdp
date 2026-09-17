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
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	datalakeclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client/operations"
	datalakemodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

func ExecuteDatalakeUpdateOperations[T any](ctx context.Context, plan *T, state *T,
	client *datalakeclient.Datalake, resp *resource.UpdateResponse,
	ops ...func(context.Context, *T, *T, *datalakeclient.Datalake, *resource.UpdateResponse) *resource.UpdateResponse,
) *resource.UpdateResponse {
	for _, op := range ops {
		op(ctx, plan, state, client, resp)
		if resp.Diagnostics.HasError() {
			return resp
		}
	}
	return resp
}

func isDryRun(opts *UpgradeOptions) bool {
	return opts != nil && opts.DryRun.ValueBool()
}

func UpgradeDatalake(ctx context.Context, client *datalakeclient.Datalake,
	datalakeName string, runtime string, imageID string,
	upgradeOpts *UpgradeOptions, pollingOptions *utils.PollingOptions,
	resp *resource.UpdateResponse) {

	dryRun := isDryRun(upgradeOpts)
	if dryRun {
		tflog.Info(ctx, fmt.Sprintf("Performing dry-run upgrade check for datalake '%s'", datalakeName))
	} else {
		tflog.Info(ctx, fmt.Sprintf("Triggering upgrade for datalake '%s'", datalakeName))
	}

	req := &datalakemodels.UpgradeDatalakeRequest{
		DatalakeName: &datalakeName,
		Runtime:      runtime,
	}
	if imageID != "" {
		req.ImageID = imageID
	}
	applyUpgradeOptions(req, upgradeOpts)

	params := operations.NewUpgradeDatalakeParams()
	params.WithInput(req)
	apiResp, err := client.Operations.UpgradeDatalakeContext(ctx, params)
	if err != nil {
		utils.AddDatalakeDiagnosticsError(err, &resp.Diagnostics, "upgrade Datalake")
		return
	}

	if dryRun {
		summarizeDryRunResponse(apiResp.Payload, &resp.Diagnostics)
		return
	}

	if apiResp.Payload == nil || apiResp.Payload.OperationID == "" {
		reason := ""
		if apiResp.Payload != nil {
			reason = apiResp.Payload.Reason
		}
		if reason != "" {
			resp.Diagnostics.AddError("Datalake upgrade was not started", reason)
		} else {
			resp.Diagnostics.AddError("Datalake upgrade was not started",
				"The API returned no operation ID. The upgrade may not be supported for this datalake configuration.")
		}
		return
	}

	if err := waitForDatalakeOperationToComplete(ctx, datalakeName,
		apiResp.Payload.OperationID, client, pollingOptions); err != nil {
		utils.AddDatalakeDiagnosticsError(err, &resp.Diagnostics,
			"wait for Datalake upgrade to complete")
	}
}

func summarizeDryRunResponse(payload *datalakemodels.UpgradeDatalakeResponse, diags *diag.Diagnostics) {
	if payload == nil {
		diags.AddWarning("Datalake upgrade dry run", "No response payload returned.")
		return
	}

	var msg strings.Builder

	if payload.Current != nil {
		msg.WriteString("Current image:\n")
		writeImageInfo(&msg, payload.Current)
	}

	fmt.Fprintf(&msg, "Upgrade candidates: %d", len(payload.UpgradeCandidates))
	for i, candidate := range payload.UpgradeCandidates {
		if candidate == nil {
			continue
		}
		fmt.Fprintf(&msg, "\n\n  Candidate %d:\n", i+1)
		writeImageInfo(&msg, candidate)
	}

	if payload.Reason != "" {
		fmt.Fprintf(&msg, "\n\nReason: %s", payload.Reason)
	}

	diags.AddWarning("Datalake upgrade dry run", msg.String())
}

func writeImageInfo(sb *strings.Builder, img *datalakemodels.ImageInfo) {
	if img.ImageID != "" {
		fmt.Fprintf(sb, "  Image ID: %s\n", img.ImageID)
	}
	if img.ImageName != "" {
		fmt.Fprintf(sb, "  Image name: %s\n", img.ImageName)
	}
	if img.ImageCatalogName != "" {
		fmt.Fprintf(sb, "  Catalog: %s\n", img.ImageCatalogName)
	}
	if img.ComponentVersions != nil {
		cv := img.ComponentVersions
		if cv.Cdp != "" {
			fmt.Fprintf(sb, "  CDP runtime: %s\n", cv.Cdp)
		}
		if cv.Cm != "" {
			fmt.Fprintf(sb, "  CM version: %s\n", cv.Cm)
		}
		if cv.Os != "" {
			os := cv.Os
			if cv.OsPatchLevel != "" {
				os += " (patch: " + cv.OsPatchLevel + ")"
			}
			fmt.Fprintf(sb, "  OS: %s\n", os)
		}
	}
}

func applyUpgradeOptions(req *datalakemodels.UpgradeDatalakeRequest, opts *UpgradeOptions) {
	if opts == nil {
		return
	}
	req.DryRun = opts.DryRun.ValueBool()
	req.LockComponents = opts.LockComponents.ValueBool()
	req.RollingUpgradeEnabled = opts.RollingUpgradeEnabled.ValueBool()
	req.SkipBackup = opts.SkipBackup.ValueBool()
	req.SkipBackupValidation = opts.SkipBackupValidation.ValueBool()
	req.SkipAtlasMetadata = opts.SkipAtlasMetadata.ValueBool()
	req.SkipRangerAudits = opts.SkipRangerAudits.ValueBool()
	req.SkipRangerHmsMetadata = opts.SkipRangerHmsMetadata.ValueBool()
	req.SkipDatahubValidation = opts.SkipDatahubValidation.ValueBool()
}

func waitForDatalakeOperationToComplete(ctx context.Context, datalakeName string,
	operationID string, client *datalakeclient.Datalake,
	pollingOptions *utils.PollingOptions) error {
	timeout, err := utils.CalculateTimeoutOrDefault(ctx, pollingOptions, time.Hour)
	if err != nil {
		return err
	}
	callFailureThresholdVal, failureThresholdError := utils.CalculateCallFailureThresholdOrDefault(ctx, pollingOptions, callFailureThreshold)
	if failureThresholdError != nil {
		return failureThresholdError
	}
	callFailedCount := 0

	stateConf := &retry.StateChangeConf{
		Pending:      []string{"RUNNING", "UNKNOWN"},
		Target:       []string{"FINISHED"},
		Delay:        0,
		Timeout:      *timeout,
		PollInterval: 10 * time.Second,
		Refresh: func() (any, string, error) {
			tflog.Debug(ctx, fmt.Sprintf("Polling operation '%s' for datalake '%s'", operationID, datalakeName))
			params := operations.NewGetOperationParams()
			params.WithInput(&datalakemodels.GetOperationRequest{
				DatalakeName: datalakeName,
				OperationID:  operationID,
			})
			resp, err := client.Operations.GetOperationContext(ctx, params)
			if err != nil {
				callFailedCount++
				if callFailedCount <= callFailureThresholdVal {
					tflog.Warn(ctx, fmt.Sprintf("Error polling operation due to [%s] but threshold limit is not reached yet (%d out of %d).", err.Error(), callFailedCount, callFailureThresholdVal))
					return nil, "RUNNING", nil
				}
				return nil, "", err
			}
			callFailedCount = 0
			if resp.Payload == nil {
				return nil, "UNKNOWN", nil
			}
			status := resp.Payload.OperationStatus
			if status == "" {
				return nil, "UNKNOWN", nil
			}
			tflog.Info(ctx, fmt.Sprintf("Operation '%s' status: %s", operationID, status))
			if status == "FAILED" || status == "CANCELLED" {
				return nil, status, fmt.Errorf("operation '%s' ended with status %s", operationID, status)
			}
			return resp, status, nil
		},
	}
	_, err = stateConf.WaitForStateContext(ctx)
	return err
}
