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
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	datalakeclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client"
)

func upgradeAwsDatalake(ctx context.Context, plan *awsDatalakeResourceModel, state *awsDatalakeResourceModel,
	client *datalakeclient.Datalake, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return ExecuteDatalakeUpdateOperations(ctx, plan, state, client, resp,
		upgradeAwsRuntimeOrImageIfChanged,
	)
}

func upgradeAwsRuntimeOrImageIfChanged(ctx context.Context, plan *awsDatalakeResourceModel,
	state *awsDatalakeResourceModel, client *datalakeclient.Datalake,
	resp *resource.UpdateResponse) *resource.UpdateResponse {

	runtimeChanged := !plan.Runtime.Equal(state.Runtime)
	imageChanged := !reflect.DeepEqual(plan.Image, state.Image)

	if !runtimeChanged && !imageChanged {
		state.UpgradeOptions = plan.UpgradeOptions
		return resp
	}

	var imageID string
	if plan.Image != nil {
		imageID = plan.Image.ID.ValueString()
	}

	UpgradeDatalake(ctx, client,
		plan.DatalakeName.ValueString(),
		plan.Runtime.ValueString(),
		imageID,
		plan.UpgradeOptions,
		plan.PollingOptions,
		resp)

	if resp.Diagnostics.HasError() {
		return resp
	}

	state.UpgradeOptions = plan.UpgradeOptions
	if !isDryRun(plan.UpgradeOptions) {
		state.Runtime = plan.Runtime
		state.Image = plan.Image
	}
	return resp
}
