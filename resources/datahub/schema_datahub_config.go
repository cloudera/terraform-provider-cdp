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
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var datahubConfigSchema = schema.Schema{
	Description:         "Retrieve configuration of an existing Datahub cluster.",
	MarkdownDescription: "Retrieve configuration of an existing Datahub cluster.",
	Attributes: map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Computed:            true,
			Optional:            true,
			Description:         "Name of the Datahub cluster.",
			MarkdownDescription: "Name of the Datahub cluster.",
		},
		"crn": schema.StringAttribute{
			Computed:            true,
			Optional:            true,
			Description:         "CRN of the Datahub cluster.",
			MarkdownDescription: "CRN of the Datahub cluster.",
		},
		"aws": schema.SingleNestedAttribute{
			Computed:            true,
			Description:         "AWS Datahub cluster configuration.",
			MarkdownDescription: "AWS Datahub cluster configuration.",
			Attributes:          awsDatahubConfigAttributes(),
		},
		"azure": schema.SingleNestedAttribute{
			Computed:            true,
			Description:         "Azure Datahub cluster configuration.",
			MarkdownDescription: "Azure Datahub cluster configuration.",
			Attributes:          azureDatahubConfigAttributes(),
		},
		"gcp": schema.SingleNestedAttribute{
			Computed:            true,
			Description:         "GCP Datahub cluster configuration.",
			MarkdownDescription: "GCP Datahub cluster configuration.",
			Attributes:          gcpDatahubConfigAttributes(),
		},
	},
}

func awsDatahubConfigAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"crn": schema.StringAttribute{
			Computed: true,
		},
		"name": schema.StringAttribute{
			Computed: true,
		},
		"status": schema.StringAttribute{
			Computed: true,
		},
		"environment": schema.StringAttribute{
			Computed: true,
		},
		"cluster_template": schema.StringAttribute{
			Computed: true,
		},
		"multi_az": schema.BoolAttribute{
			Computed: true,
		},
		"security": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"se_linux": schema.StringAttribute{
					Computed: true,
				},
			},
		},
		"instance_group": schema.ListNestedAttribute{
			Computed: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"instance_group_name": schema.StringAttribute{
						Computed: true,
					},
					"availability_zones": schema.ListAttribute{
						Computed:    true,
						ElementType: schema.StringAttribute{}.GetType(),
					},
				},
			},
		},
	}
}

func azureDatahubConfigAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"crn": schema.StringAttribute{
			Computed: true,
		},
		"name": schema.StringAttribute{
			Computed: true,
		},
		"status": schema.StringAttribute{
			Computed: true,
		},
		"environment": schema.StringAttribute{
			Computed: true,
		},
		"cluster_template": schema.StringAttribute{
			Computed: true,
		},
		"multi_az": schema.BoolAttribute{
			Computed: true,
		},
		"security": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"se_linux": schema.StringAttribute{
					Computed: true,
				},
			},
		},
		"instance_group": schema.ListNestedAttribute{
			Computed: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"instance_group_name": schema.StringAttribute{
						Computed: true,
					},
					"availability_zones": schema.ListAttribute{
						Computed:    true,
						ElementType: schema.StringAttribute{}.GetType(),
					},
				},
			},
		},
	}
}

func gcpDatahubConfigAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"crn": schema.StringAttribute{
			Computed: true,
		},
		"name": schema.StringAttribute{
			Computed: true,
		},
		"status": schema.StringAttribute{
			Computed: true,
		},
		"environment_name": schema.StringAttribute{
			Computed: true,
		},
		"cluster_template_name": schema.StringAttribute{
			Computed: true,
		},
		"multi_az": schema.BoolAttribute{
			Computed: true,
		},
		"security": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"se_linux": schema.StringAttribute{
					Computed: true,
				},
			},
		},
		"instance_group": schema.ListNestedAttribute{
			Computed: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"instance_group_name": schema.StringAttribute{
						Computed: true,
					},
					"availability_zones": schema.ListAttribute{
						Computed:    true,
						ElementType: schema.StringAttribute{}.GetType(),
					},
				},
			},
		},
	}
}
