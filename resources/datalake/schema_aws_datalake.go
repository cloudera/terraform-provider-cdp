package datalake

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var awsDatalakeResourceSchema schema.Schema = schema.Schema{
	MarkdownDescription: "A Data Lake is a service which provides a protective ring around the data stored in a cloud object store, including authentication, authorization, and governance support.",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"instance_profile": schema.StringAttribute{
			Required: true,
		},
		"storage_bucket_location": schema.StringAttribute{
			Required: true,
		},
		"certificate_expiration_state": schema.StringAttribute{
			Computed: true,
		},
		"cloud_storage_base_location": schema.StringAttribute{
			Computed: true,
		},
		"cloudbreak_version": schema.StringAttribute{
			Computed: true,
		},
		"cloudera_manager": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"cloudera_manager_repository_url": schema.StringAttribute{
					Computed: true,
				},
				"cloudera_manager_server_url": schema.StringAttribute{
					Computed: true,
				},
				"version": schema.StringAttribute{
					Computed: true,
				},
			},
		},
		"creation_date": schema.StringAttribute{
			Computed: true,
		},
		"credential_crn": schema.StringAttribute{
			Computed: true,
		},
		"crn": schema.StringAttribute{
			Computed: true,
		},
		"custom_instance_groups": schema.SetNestedAttribute{
			Optional: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"instance_type": schema.StringAttribute{
						Optional: true,
					},
					"name": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
		"database_availability_type": schema.StringAttribute{
			Optional: true,
			Computed: true,
		},
		"database_engine_version": schema.StringAttribute{
			Optional: true,
			Computed: true,
		},
		"datalake_name": schema.StringAttribute{
			Required: true,
		},
		"datalake_template": schema.StringAttribute{
			Optional: true,
		},
		"enable_ranger_raz": schema.BoolAttribute{
			Optional: true,
			Computed: true,
		},
		"environment_crn": schema.StringAttribute{
			Computed: true,
		},
		"endpoints": schema.SetNestedAttribute{
			Computed: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"display_name": schema.StringAttribute{
						Computed: true,
					},
					"knox_service": schema.StringAttribute{
						Computed: true,
					},
					"mode": schema.StringAttribute{
						Computed: true,
					},
					"open": schema.BoolAttribute{
						Computed: true,
					},
					"service_name": schema.StringAttribute{
						Computed: true,
					},
					"service_url": schema.StringAttribute{
						Computed: true,
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
			},
		},
		"instance_groups": schema.SetNestedAttribute{
			Computed: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"instances": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"ambari_server": schema.BoolAttribute{
									Computed: true,
								},
								"discovery_fqdn": schema.StringAttribute{
									Computed: true,
								},
								"id": schema.StringAttribute{
									Computed: true,
								},
								"instance_group": schema.StringAttribute{
									Computed: true,
								},
								"instance_status": schema.StringAttribute{
									Computed: true,
								},
								"instance_type_val": schema.StringAttribute{
									Computed: true,
								},
								"life_cycle": schema.StringAttribute{
									Computed: true,
								},
								"mounted_volumes": schema.SetNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"device": schema.StringAttribute{
												Computed: true,
											},
											"volume_id": schema.StringAttribute{
												Computed: true,
											},
											"volume_size": schema.StringAttribute{
												Computed: true,
											},
											"volume_type": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
								"private_ip": schema.StringAttribute{
									Computed: true,
								},
								"public_ip": schema.StringAttribute{
									Computed: true,
								},
								"ssh_port": schema.Int64Attribute{
									Computed: true,
								},
								"state": schema.StringAttribute{
									Computed: true,
								},
								"status_reason": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
		"java_version": schema.Int64Attribute{
			Optional: true,
		},
		"load_balancers": schema.SetNestedAttribute{
			Computed: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"resource_id": schema.StringAttribute{
						Computed: true,
					},
					"cloud_dns": schema.StringAttribute{
						Computed: true,
					},
					"fqdn": schema.StringAttribute{
						Computed: true,
					},
					"ip": schema.StringAttribute{
						Computed: true,
					},
					"load_balancer_type": schema.StringAttribute{
						Computed: true,
					},
					"targets": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"listener_id": schema.StringAttribute{
									Computed: true,
								},
								"target_group_id": schema.StringAttribute{
									Computed: true,
								},
								"port": schema.Int64Attribute{
									Computed: true,
								},
								"target_instances": schema.SetAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
							},
						},
					},
				},
			},
		},
		"multi_az": schema.BoolAttribute{
			Optional: true,
			Computed: true,
		},
		"product_versions": schema.SetNestedAttribute{
			Computed: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Computed: true,
					},
					"version": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
		"recipes": schema.SetNestedAttribute{
			Optional: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"instance_group_name": schema.StringAttribute{
						Required: true,
					},
					"recipe_names": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
		"region": schema.StringAttribute{
			Computed: true,
		},
		"runtime": schema.StringAttribute{
			Optional: true,
		},
		"scale": schema.StringAttribute{
			Optional: true,
		},
		"spot_max_price": schema.Float64Attribute{
			Optional: true,
		},
		"spot_percentage": schema.Int64Attribute{
			Optional: true,
		},
		"status": schema.StringAttribute{
			Computed: true,
		},
		"status_reason": schema.StringAttribute{
			Computed: true,
		},
		"tags": schema.MapAttribute{
			Optional:    true,
			Computed:    true,
			ElementType: types.StringType,
		},
	},
}
