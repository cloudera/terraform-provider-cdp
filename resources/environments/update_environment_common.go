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
	"fmt"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

func performEnvironmentUpdate[T any](ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse, client *environmentsclient.Environments, update func(context.Context, *T, *T, *environmentsclient.Environments, *resource.UpdateResponse) *resource.UpdateResponse) {
	var plan T
	var state T
	planDiags := req.Plan.Get(ctx, &plan)
	var stateDiags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(planDiags...)
	resp.Diagnostics.Append(stateDiags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Got Error while trying to set plan")
		return
	}

	update(ctx, &plan, &state, client, resp)

	stateDiags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(stateDiags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func updateSshKeyIfChanged(ctx context.Context, client *environmentsclient.Environments, planKey types.String, stateKey *types.String, envName *string, resp *resource.UpdateResponse) *resource.UpdateResponse {
	if planKey.IsNull() || planKey.IsUnknown() {
		return resp
	}
	if planKey.ValueString() == "" {
		resp.Diagnostics.AddError("Invalid SSH public key", "public_key must be a non-empty, known value.")
		return resp
	}
	if !reflect.DeepEqual(planKey, *stateKey) {
		if err := updateSshKey(ctx, client, planKey, envName); err != nil {
			utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "update SSH key")
			return resp
		}
		*stateKey = planKey
	}
	return resp
}

func updateSshKey(ctx context.Context, client *environmentsclient.Environments, publicKey types.String, env *string) error {
	if publicKey.IsNull() || publicKey.IsUnknown() || len(publicKey.ValueString()) == 0 {
		return nil
	}
	params := operations.NewUpdateSSHKeyParams()
	if !publicKey.IsNull() && len(publicKey.ValueString()) != 0 {
		params.WithInput(&environmentsmodels.UpdateSSHKeyRequest{
			Environment:         env,
			NewPublicKey:        publicKey.ValueString(),
			ExistingPublicKeyID: "",
		})
	}
	tflog.Info(ctx, "Updating SSH key in the environment")
	_, err := client.Operations.UpdateSSHKeyContext(ctx, params)
	return err
}

func updateProxyConfigurationIfChanged(ctx context.Context, client *environmentsclient.Environments, state *types.String, plan *types.String, env *string, resp *resource.UpdateResponse) *resource.UpdateResponse {
	if plan == nil || plan.IsUnknown() {
		return resp
	}
	if state == nil {
		return resp
	}
	if reflect.DeepEqual(*plan, *state) {
		return resp
	}

	removeProxy := plan.IsNull() || plan.ValueString() == ""
	params := operations.NewUpdateProxyConfigParams().WithInput(&environmentsmodels.UpdateProxyConfigRequest{
		Environment:     env,
		ProxyConfigName: plan.ValueString(),
		RemoveProxy:     removeProxy,
	})
	tflog.Info(ctx, "Updating proxy configuration in the environment")
	if _, err := client.Operations.UpdateProxyConfigContext(ctx, params); err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "update proxy configuration")
		return resp
	}

	*state = *plan
	return resp
}

func updateCredentialIfChanged(ctx context.Context, client *environmentsclient.Environments, plan types.String, state *types.String, env *string, resp *resource.UpdateResponse) *resource.UpdateResponse {
	if reflect.DeepEqual(plan, *state) {
		return resp
	}
	params := operations.NewChangeEnvironmentCredentialParams().WithInput(&environmentsmodels.ChangeEnvironmentCredentialRequest{
		CredentialName:  plan.ValueStringPointer(),
		EnvironmentName: env,
	})
	tflog.Info(ctx, "Updating credential for the environment")
	if _, err := client.Operations.ChangeEnvironmentCredentialContext(ctx, params); err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "update credential")
		return resp
	}
	*state = plan
	return resp
}

func SetEndpointAccessGatewayIfChanged(ctx context.Context, planScheme types.String, planSubnetIds types.Set, stateScheme types.String, stateSubnetIds types.Set, environmentName string, client *environmentsclient.Environments, pollingOptions *utils.PollingOptions, diags *diag.Diagnostics) {
	if planScheme.IsNull() || planScheme.IsUnknown() {
		return
	}
	if planSubnetIds.IsUnknown() {
		return
	}

	schemeChanged := !planScheme.Equal(stateScheme)
	subnetIdsChanged := !planSubnetIds.Equal(stateSubnetIds)

	if !schemeChanged && !subnetIdsChanged {
		return
	}

	scheme := planScheme.ValueString()
	tflog.Info(ctx, fmt.Sprintf("Endpoint access gateway change detected for environment '%s', calling SetEndpointAccessGateway.", environmentName))

	params := operations.NewSetEndpointAccessGatewayParams()
	params.WithInput(&environmentsmodels.SetEndpointAccessGatewayRequest{
		EndpointAccessGatewayScheme:    &scheme,
		EndpointAccessGatewaySubnetIds: utils.FromSetValueToStringList(planSubnetIds),
		Environment:                    &environmentName,
	})
	resp, err := client.Operations.SetEndpointAccessGatewayContext(ctx, params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, diags, "set endpoint access gateway")
		return
	}

	if resp.Payload != nil && resp.Payload.OperationID != "" {
		if err := waitForOperationToComplete(ctx, environmentName, resp.Payload.OperationID, client, pollingOptions); err != nil {
			utils.AddEnvironmentDiagnosticsError(err, diags, "wait for set endpoint access gateway operation to complete")
			return
		}
	}
}

func waitForOperationToComplete(ctx context.Context, environmentName string, operationID string, client *environmentsclient.Environments, pollingOptions *utils.PollingOptions) error {
	timeout, err := utils.CalculateTimeoutOrDefault(ctx, pollingOptions, timeoutOneHour)
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
		Refresh: func() (interface{}, string, error) {
			tflog.Debug(ctx, fmt.Sprintf("Polling operation '%s' for environment '%s'", operationID, environmentName))
			params := operations.NewGetOperationParams()
			params.WithInput(&environmentsmodels.GetOperationRequest{
				EnvironmentName: &environmentName,
				OperationID:     operationID,
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
