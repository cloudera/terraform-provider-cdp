// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package auditcredential

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

func (r *awsGovCloudAuditCredentialResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	attr := map[string]schema.Attribute{}
	utils.AppendToResourceSchema(attr, auditCredentialCommonSchema)
	utils.AppendToResourceSchema(attr, map[string]schema.Attribute{
		"role_arn": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The ARN of the delegated access role.",
		},
	})
	resp.Schema = schema.Schema{
		MarkdownDescription: "The AWS GovCloud audit credential is used for MSA (Monitoring, SPI, and Auditing) access to your AWS GovCloud account.",
		Attributes:          attr,
	}
}
