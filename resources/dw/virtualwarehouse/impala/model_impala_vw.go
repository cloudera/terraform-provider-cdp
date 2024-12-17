// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package impala

import (
	"time"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

const timeZone = time.RFC850

type resourceModel struct {
	ID                types.String          `tfsdk:"id"`
	ClusterID         types.String          `tfsdk:"cluster_id"`
	DatabaseCatalogID types.String          `tfsdk:"database_catalog_id"`
	Name              types.String          `tfsdk:"name"`
	LastUpdated       types.String          `tfsdk:"last_updated"`
	Status            types.String          `tfsdk:"status"`
	ImageVersion      types.String          `tfsdk:"image_version"`
	PollingOptions    *utils.PollingOptions `tfsdk:"polling_options"`
}

func (p *resourceModel) setFromDescribeVwResponse(resp *models.DescribeVwResponse) {
	p.ID = types.StringValue(resp.Vw.DbcID)
	p.DatabaseCatalogID = types.StringValue(resp.Vw.DbcID)
	p.Name = types.StringValue(resp.Vw.Name)
	p.Status = types.StringValue(resp.Vw.Status)
	p.ImageVersion = types.StringValue(resp.Vw.CdhVersion)
	p.LastUpdated = types.StringValue(time.Now().Format(timeZone))
}

func (p *resourceModel) GetPollingOptions() *utils.PollingOptions {
	return p.PollingOptions
}
