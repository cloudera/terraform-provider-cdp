// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package iam

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func (r *samlProvider) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"saml_provider_name": schema.StringAttribute{
				MarkdownDescription: "The name of SAML provider. The name must be unique, must have a maximum of 128 characters, and must contain only alphanumeric characters, \"-\" and \"_\". Names are are not case-sensitive.",
				Required:            true,
			},
			"enable_scim": schema.BoolAttribute{
				MarkdownDescription: "Whether to enable SCIM on this SAML provider. System for Cross-domain Identity Management (SCIM) version 2.0 is a standard for automating the provisioning of user and group identity information from identity provider to CDP. It is not supported for Cloudera for Government.",
				Optional:            true,
			},
			"generate_workload_username_by_email": schema.BoolAttribute{
				MarkdownDescription: "Whether to generate users' workload username by email . The default is to generate workload usernames by identity provider user ID (SAML NameID).",
				Optional:            true,
			},
			"saml_metadata_document": schema.StringAttribute{
				MarkdownDescription: "SAML metadata document XML file. Length of meta data document cannot be more than 200 KB (200,000 bytes). Max Length: 200000",
				Optional:            true,
			},
			"sync_groups_on_login": schema.StringAttribute{
				MarkdownDescription: "Whether to sync group information for users federated with this SAML provider. Group membership can be passed using the https://cdp.cloudera.com/SAML/Attributes/groups SAML assertion. The default is to synchronize group membership.",
				Optional:            true,
			},
		},
	}
}
