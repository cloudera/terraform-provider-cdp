// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datahub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client/operations"
	datahubmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

func describeDatahubWithDiagnosticHandle(datahub string, id string, ctx context.Context, client *cdp.Client, diags *diag.Diagnostics, state *tfsdk.State) (*datahubmodels.Cluster, error) {
	tflog.Info(ctx, "About to describe datahub '"+datahub+"'.")
	params := operations.NewDescribeClusterParamsWithContext(ctx)
	params.WithInput(&datahubmodels.DescribeClusterRequest{
		ClusterName: &datahub,
	})
	describeResp, err := client.Datahub.Operations.DescribeCluster(params)
	if err != nil {
		tflog.Warn(ctx, "Something happened during datahub fetch: "+err.Error())
		if isNotFoundError(err) {
			diags.AddWarning("Resource not found on provider", "Datahub not found, removing from state.")
			tflog.Warn(ctx, "Datahub not found, removing from state", map[string]interface{}{
				"id": id,
			})
			state.RemoveResource(ctx)
			return nil, err
		}
		utils.AddDatahubDiagnosticsError(err, diags, "read Datahub")
		return nil, err
	}
	return utils.LogDatahubSilently(ctx, describeResp.GetPayload().Cluster, "Result of describe datahub cluster: "), nil
}
