// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package hive

import (
	"time"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

type resourceModel struct {
	ID                types.String          `tfsdk:"id"`
	ClusterID         types.String          `tfsdk:"cluster_id"`
	DatabaseCatalogID types.String          `tfsdk:"database_catalog_id"`
	Name              types.String          `tfsdk:"name"`
	LastUpdated       types.String          `tfsdk:"last_updated"`
	Status            types.String          `tfsdk:"status"`
	PollingOptions    *utils.PollingOptions `tfsdk:"polling_options"`
}

// TODO these are the same everywhere, abstract this
func (p *resourceModel) getPollingTimeout() time.Duration {
	if p.PollingOptions != nil {
		return time.Duration(p.PollingOptions.PollingTimeout.ValueInt64()) * time.Minute
	}
	return 40 * time.Minute
}

func (p *resourceModel) getCallFailureThreshold() int {
	if p.PollingOptions != nil {
		return int(p.PollingOptions.CallFailureThreshold.ValueInt64())
	}
	return 3
}
