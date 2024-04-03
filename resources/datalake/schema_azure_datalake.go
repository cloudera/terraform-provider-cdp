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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var azureDatalakeResourceSchema = schema.Schema{
	MarkdownDescription: "A Data Lake is a service which provides a protective ring around the data stored in a cloud object store, including authentication, authorization, and governance support.",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"polling_options": schema.SingleNestedAttribute{
			MarkdownDescription: "Polling related configuration options that could specify various values that will be used during CDP resource creation.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"async": schema.BoolAttribute{
					MarkdownDescription: "Boolean value that specifies if Terraform should wait for resource creation/deletion.",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
					PlanModifiers: []planmodifier.Bool{
						boolplanmodifier.UseStateForUnknown(),
					},
				},
				"polling_timeout": schema.Int64Attribute{
					MarkdownDescription: "Timeout value in minutes that specifies for how long should the polling go for resource creation/deletion.",
					Default:             int64default.StaticInt64(90),
					Computed:            true,
					Optional:            true,
				},
			},
		},
		"managed_identity": schema.StringAttribute{
			Required: true,
		},
		"storage_location_base": schema.StringAttribute{
			Required: true,
		},
		"certificate_expiration_state": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"cloudera_manager": schema.SingleNestedAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.Object{
				objectplanmodifier.UseStateForUnknown(),
			},
			Attributes: map[string]schema.Attribute{
				"cloudera_manager_repository_url": schema.StringAttribute{
					Computed: true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
				"cloudera_manager_server_url": schema.StringAttribute{
					Computed: true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
				"version": schema.StringAttribute{
					Computed: true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
			},
		},
		"creation_date": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"crn": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"datalake_name": schema.StringAttribute{
			Required: true,
		},
		"enable_ranger_raz": schema.BoolAttribute{
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"environment_crn": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"endpoints": schema.SetNestedAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.Set{
				setplanmodifier.UseStateForUnknown(),
			},
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"display_name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"knox_service": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"mode": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"open": schema.BoolAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"service_name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"service_url": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
		},
		"environment_name": schema.StringAttribute{
			Required: true,
		},
		"image": schema.SingleNestedAttribute{
			Optional: true,
			Attributes: map[string]schema.Attribute{
				"catalog": schema.StringAttribute{
					Optional: true,
				},
				"id": schema.StringAttribute{
					Required: true,
				},
				"os": schema.StringAttribute{
					Optional: true,
				},
			},
		},
		"instance_groups": schema.SetNestedAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.Set{
				setplanmodifier.UseStateForUnknown(),
			},
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"instances": schema.SetNestedAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"discovery_fqdn": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"id": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"instance_group": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"instance_status": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"instance_type_val": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"private_ip": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"public_ip": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"ssh_port": schema.Int64Attribute{
									Computed: true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"state": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"status_reason": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
					},
					"name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
		},
		"java_version": schema.Int64Attribute{
			Optional: true,
		},
		"database_type": schema.StringAttribute{
			MarkdownDescription: "The type of the azure database. FLEXIBLE_SERVER is the next generation managed PostgreSQL service in Azure that provides maximum flexibility over your database, built-in cost-optimizations. SINGLE_SERVER is a fully managed database service with minimal requirements for customizations of the database.",
			Optional:            true,
		},
		"product_versions": schema.SetNestedAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.Set{
				setplanmodifier.UseStateForUnknown(),
			},
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"version": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
		},
		"recipes": schema.SetNestedAttribute{
			Optional: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"instance_group_name": schema.StringAttribute{
						MarkdownDescription: "The name of the designated instance group.",
						Required:            true,
					},
					"recipe_names": schema.SetAttribute{
						MarkdownDescription: "The set of recipe names that are going to be applied on the given instance group.",
						ElementType:         types.StringType,
						Required:            true,
					},
				},
			},
		},
		"runtime": schema.StringAttribute{
			Optional: true,
		},
		"scale": schema.StringAttribute{
			Optional: true,
		},
		"status": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"status_reason": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"tags": schema.MapAttribute{
			Optional:    true,
			ElementType: types.StringType,
		},
		"load_balancer_sku": schema.StringAttribute{
			MarkdownDescription: "Represents the Azure load balancer SKU type. The current default  is BASIC. To disable the load balancer, use type NONE. Possible values: BASIC, STANDARD, NONE",
			Optional:            true,
		},
		"flexible_server_delegated_subnet_id": schema.StringAttribute{
			MarkdownDescription: "This  argument  allows  you  to specify the subnet ID for the subnet within which you want to configure your Azure Flexible Server.",
			Optional:            true,
		},
		"multi_az": schema.BoolAttribute{
			MarkdownDescription: "Creates CDP datalake distributed across multiple availability  zones in an Azure region.",
			Optional:            true,
		},
	},
}
