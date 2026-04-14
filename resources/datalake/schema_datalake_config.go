// Copyright 2025 Cloudera. All Rights Reserved.
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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

func (e *datalakeConfigDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attr := map[string]schema.Attribute{}
	utils.AppendToDataSourceSchema(attr, generalDatalakeAttributes)
	utils.AppendToDataSourceSchema(attr, map[string]schema.Attribute{
		"aws":   awsSchema,
		"azure": azureSchema,
		"gcp":   gcpSchema,
	})
	resp.Schema = schema.Schema{
		MarkdownDescription: "Collects datalake information for re-creation purposes.",
		Attributes:          attr,
	}
}

var awsSchema = schema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"polling_options": pollingOptionsSchema,
		"creation_date": schema.StringAttribute{
			Computed: true,
		},
		"crn": schema.StringAttribute{
			Computed: true,
		},
		"datalake_name": schema.StringAttribute{
			Computed: true,
		},
		"enable_ranger_raz": schema.BoolAttribute{
			Computed: true,
		},
		"environment_crn": schema.StringAttribute{
			Computed: true,
		},
		"environment_name": schema.StringAttribute{
			Computed: true,
		},
		"image": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"catalog_name": schema.StringAttribute{
					Computed: true,
				},
				"id": schema.StringAttribute{
					Computed: true,
				},
				"os": schema.StringAttribute{
					Computed: true,
				},
			},
		},
		"java_version": schema.StringAttribute{
			Computed: true,
		},
		"recipes": schema.SetNestedAttribute{
			Computed: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"instance_group_name": schema.StringAttribute{
						Computed: true,
					},
					"recipe_names": schema.SetAttribute{
						ElementType: types.StringType,
						Computed:    true,
					},
				},
			},
		},
		"custom_instance_groups": schema.SetNestedAttribute{
			Computed: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Computed: true,
					},
					"instance_type": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
		"runtime": schema.StringAttribute{
			Computed: true,
		},
		"scale": schema.StringAttribute{
			Computed: true,
		},
		"status": schema.StringAttribute{
			Computed: true,
		},
		"status_reason": schema.BoolAttribute{
			Computed: true,
		},
		"multi_az": schema.BoolAttribute{
			Computed: true,
		},
		"tags": schema.MapAttribute{
			ElementType: types.StringType,
			Computed:    true,
		},
		"security": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"se_linux": schema.StringAttribute{
					Computed: true,
				},
			},
		},
		"delete_options": deleteOptionsSchema,
		"certificate_expiration_state": schema.StringAttribute{
			Computed: true,
		},
		"storage_location_base": schema.StringAttribute{
			Computed: true,
		},
		"instance_profile": schema.StringAttribute{
			Computed: true,
		},
		"enable_ranger_rms": schema.BoolAttribute{
			Computed: true,
		},
		"architecture": schema.StringAttribute{
			Computed: true,
		},
	},
}

var azureSchema = schema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"polling_options": pollingOptionsSchema,
		"managed_identity": schema.StringAttribute{
			Computed: true,
		},
		"storage_location_base": schema.StringAttribute{
			Computed: true,
		},
		"certificate_expiration_state": schema.StringAttribute{
			Computed: true,
		},
		"creation_date": schema.StringAttribute{
			Computed: true,
		},
		"crn": schema.StringAttribute{
			Computed: true,
		},
		"datalake_name": schema.StringAttribute{
			Computed: true,
		},
		"enable_ranger_raz": schema.BoolAttribute{
			Computed: true,
		},
		"environment_crn": schema.StringAttribute{
			Computed: true,
		},
		"environment_name": schema.StringAttribute{
			Optional: true,
		},
		"image": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"catalog_name": schema.StringAttribute{
					Computed: true,
				},
				"id": schema.StringAttribute{
					Computed: true,
				},
				"os": schema.StringAttribute{
					Computed: true,
				},
			},
		},
		"java_version": schema.Int32Attribute{
			Computed: true,
		},
		"database_type": schema.StringAttribute{
			Computed: true,
		},
		"recipes": schema.SetNestedAttribute{
			Computed: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"instance_group_name": schema.StringAttribute{
						Computed: true,
					},
					"recipe_names": schema.SetAttribute{
						ElementType: types.StringType,
						Computed:    true,
					},
				},
			},
		},
		"custom_instance_groups": schema.SetNestedAttribute{
			Computed: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Computed: true,
					},
					"instance_type": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
		"runtime": schema.StringAttribute{
			Computed: true,
		},
		"scale": schema.StringAttribute{
			Computed: true,
		},
		"status": schema.StringAttribute{
			Computed: true,
		},
		"status_reason": schema.StringAttribute{
			Computed: true,
		},
		"tags": schema.MapAttribute{
			ElementType: types.StringType,
			Computed:    true,
		},
		"load_balancer_sku": schema.StringAttribute{
			Computed: true,
		},
		"flexible_server_delegated_subnet_id": schema.StringAttribute{
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
		"delete_options": deleteOptionsSchema,
	},
}

var gcpSchema = schema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"polling_options": pollingOptionsSchema,
		"creation_date": schema.StringAttribute{
			Computed: true,
		},
		"crn": schema.StringAttribute{
			Computed: true,
		},
		"datalake_name": schema.StringAttribute{
			Computed: true,
		},
		"enable_ranger_raz": schema.BoolAttribute{
			Computed: true,
		},
		"environment_crn": schema.StringAttribute{
			Computed: true,
		},
		"environment_name": schema.StringAttribute{
			Computed: true,
		},
		"image": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"catalog_name": schema.StringAttribute{
					Computed: true,
				},
				"id": schema.StringAttribute{
					Computed: true,
				},
				"os": schema.StringAttribute{
					Computed: true,
				},
			},
		},
		"java_version": schema.StringAttribute{
			Computed: true,
		},
		"recipes": schema.SetNestedAttribute{
			Computed: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"instance_group_name": schema.StringAttribute{
						Computed: true,
					},
					"recipe_names": schema.SetAttribute{
						ElementType: types.StringType,
						Computed:    true,
					},
				},
			},
		},
		"custom_instance_groups": schema.SetNestedAttribute{
			Computed: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Computed: true,
					},
					"instance_type": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
		"runtime": schema.StringAttribute{
			Computed: true,
		},
		"scale": schema.StringAttribute{
			Computed: true,
		},
		"status": schema.StringAttribute{
			Computed: true,
		},
		"status_reason": schema.BoolAttribute{
			Computed: true,
		},
		"multi_az": schema.BoolAttribute{
			Computed: true,
		},
		"tags": schema.MapAttribute{
			ElementType: types.StringType,
			Computed:    true,
		},
		"security": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"se_linux": schema.StringAttribute{
					Computed: true,
				},
			},
		},
		"delete_options": deleteOptionsSchema,
		"cloud_provider_configuration": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"service_account_email": schema.StringAttribute{
					Computed: true,
				},
				"storage_location": schema.StringAttribute{
					Computed: true,
				},
			},
		},
	},
}

var pollingOptionsSchema = schema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]schema.Attribute{
		"async": schema.BoolAttribute{
			Computed: true,
		},
		"polling_timeout": schema.Int64Attribute{
			Computed: true,
		},
		"call_failure_threshold": schema.Int64Attribute{
			Computed: true,
		},
	},
}

var deleteOptionsSchema = schema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]schema.Attribute{
		"forced": schema.BoolAttribute{
			Computed: true,
		},
	},
}

var generalDatalakeAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Description: "The name of the datalake. Either this or the CRN must be specified - but not both -.",
		Optional:    true,
	},
	"crn": schema.StringAttribute{
		Description: "The CRN of the datalake. Either this or the name must be specified - but not both -.",
		Optional:    true,
	},
}
