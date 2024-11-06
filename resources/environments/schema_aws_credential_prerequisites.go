// Copyright 2024 Cloudera. All Rights Reserved.
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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (d *awsCredentialPrerequisitesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This data source is used to get information required to set up a delegated access role in AWS that can be used to create a CDP credential.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				MarkdownDescription: "The AWS account ID of the identity used by CDP when assuming a delegated access role associated with a CDP credential.",
				Computed:            true,
			},
			"external_id": schema.StringAttribute{
				MarkdownDescription: "The external ID that will be used when assuming a delegated access role associated with a CDP credential.",
				Computed:            true,
			},
			"policy": schema.StringAttribute{
				MarkdownDescription: "The related policy json encoded in base64",
				Computed:            true,
			},
			"policies": schema.MapAttribute{
				Computed:            true,
				MarkdownDescription: "The fine-grained policies related to each service.",
				ElementType:         types.StringType,
			},
		},
	}
}
