// Copyright 2023 Cloudera. All Rights Reserved.
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

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

func (r *gcpDatalakeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = gcpDatalakeSchema
}

var gcpDatalakeSchema = getGcpResourceSchema()

func getGcpResourceSchema() schema.Schema {
	attr := map[string]schema.Attribute{}
	utils.AppendToResourceSchema(attr, generalAttributes)
	utils.AppendToResourceSchema(attr, map[string]schema.Attribute{
		"cloud_provider_configuration": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"service_account_email": schema.StringAttribute{
					MarkdownDescription: "Email id of the service account to be associated with the datalake IdBroker instance. This service account should have \"token.creator\" role for one or more storage accounts that has access to storage.",
					Description:         "Email id of the service account to be associated with the datalake IdBroker instance. This service account should have \"token.creator\" role for one or more storage accounts that has access to storage.",
					Required:            true,
				},
				"storage_location": schema.StringAttribute{
					MarkdownDescription: "The location of the GCS bucket to be used as storage. The location has to start with gs:// followed by the bucket name.",
					Description:         "The location of the GCS bucket to be used as storage. The location has to start with gs:// followed by the bucket name.",
					Required:            true,
				},
			},
		},
	})
	return schema.Schema{
		MarkdownDescription: "A Data Lake is a service which provides a protective ring around the data stored in a cloud object store, including authentication, authorization, and governance support.",
		Description:         "A Data Lake is a service which provides a protective ring around the data stored in a cloud object store, including authentication, authorization, and governance support.",
		Attributes:          attr,
	}
}
