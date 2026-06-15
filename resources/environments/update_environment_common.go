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
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

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
