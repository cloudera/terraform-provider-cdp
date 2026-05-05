// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package validators

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

type kubeAPIAuthorizedIPRangesPrivateClusterValidator struct{}

func (v kubeAPIAuthorizedIPRangesPrivateClusterValidator) Description(_ context.Context) string {
	return "kube_api_authorized_ip_ranges must not be set when private_cluster is true"
}

func (v kubeAPIAuthorizedIPRangesPrivateClusterValidator) MarkdownDescription(_ context.Context) string {
	return "`kube_api_authorized_ip_ranges` must not be set when `private_cluster` is `true`"
}

func (v kubeAPIAuthorizedIPRangesPrivateClusterValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	privateClusterPath := req.Path.ParentPath().AtName("private_cluster")

	var privateCluster types.Bool
	diags := req.Config.GetAttribute(ctx, privateClusterPath, &privateCluster)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if privateCluster.IsNull() || privateCluster.IsUnknown() {
		return
	}

	ipRangeCount := len(req.ConfigValue.Elements())
	privateClusterGiven := privateCluster.ValueBool()
	kubeAPIAuthorizedIPRangesGiven := ipRangeCount > 0
	kubeAPIAuthorizedIPRangesInvalid := privateClusterGiven && kubeAPIAuthorizedIPRangesGiven

	validationLogJSON := v.prepareValidationLog(req, privateCluster)
	utils.LogSilently(ctx, "Validating kube_api_authorized_ip_ranges configuration: ", validationLogJSON)

	if kubeAPIAuthorizedIPRangesInvalid {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid compute cluster configuration",
			"When `private_cluster` is `true`, `kube_api_authorized_ip_ranges` must not be set.",
		)
	}
}

func (v kubeAPIAuthorizedIPRangesPrivateClusterValidator) prepareValidationLog(req validator.SetRequest, privateCluster types.Bool) []byte {
	validationLog := map[string]interface{}{
		"values": map[string]interface{}{
			"kube_api_authorized_ip_ranges": req.ConfigValue.String(),
			"private_cluster":               privateCluster.String(),
		},
	}
	validationLogJSON, _ := json.Marshal(validationLog)
	return validationLogJSON
}

func KubeAPIAuthorizedIPRangesMustBeEmptyWhenPrivateClusterTrue() validator.Set {
	return kubeAPIAuthorizedIPRangesPrivateClusterValidator{}
}
