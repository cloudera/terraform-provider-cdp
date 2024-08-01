// Copyright 2023 Cloudera. All Rights Reserved.
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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var generalAttributes = map[string]schema.Attribute{
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
			"polling_timeout": schema.Int64Attribute{
				MarkdownDescription: "Timeout value in minutes that specifies for how long should the polling go for resource creation/deletion.",
				Default:             int64default.StaticInt64(60),
				Computed:            true,
				Optional:            true,
			},
			"call_failure_threshold": schema.Int64Attribute{
				MarkdownDescription: "Threshold value that specifies how many times should a single call failure happen before giving up the polling.",
				Default:             int64default.StaticInt64(3),
				Computed:            true,
				Optional:            true,
			},
		},
	},
	"crn": schema.StringAttribute{
		MarkdownDescription: "The CRN of the cluster.",
		Computed:            true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"status": schema.StringAttribute{
		MarkdownDescription: "The last known state of the cluster",
		Description:         "The last known state of the cluster",
		Computed:            true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"name": schema.StringAttribute{
		MarkdownDescription: "The name of the cluster.",
		Required:            true,
	},
	"destroy_options": schema.SingleNestedAttribute{
		Optional:            true,
		Description:         "Cluster deletion options.",
		MarkdownDescription: "Cluster deletion options.",
		Attributes: map[string]schema.Attribute{
			"force_delete_cluster": schema.BoolAttribute{
				MarkdownDescription: "An indicator that will take place once the cluster termination will be performed. " +
					"If it is true, that means if something would go sideways during termination, the operation will proceed, " +
					"however in such a case no notification would come thus it is advisable to check the cloud provider if " +
					"there are no leftover resources once the destroy is finished.",
				Description: "An indicator that will take place once the cluster termination will be performed. " +
					"If it is true, that means if something would go sideways during termination, the operation will proceed, " +
					"however in such a case no notification would come thus it is advisable to check the cloud provider if " +
					"there are no leftover resources once the destroy is finished.",
				Default:  booldefault.StaticBool(false),
				Computed: true,
				Optional: true,
			},
		},
	},
	"custom_configurations_name": schema.StringAttribute{
		MarkdownDescription: "The name of the custom configurations to use for cluster creation.",
		Description:         "The name of the custom configurations to use for cluster creation.",
		Optional:            true,
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
	"request_template": schema.StringAttribute{
		MarkdownDescription: "JSON  template  to  use for cluster creation. This is different from cluster template and would be removed in the future.",
		Description:         "JSON  template  to  use for cluster creation. This is different from cluster template and would be removed in the future.",
		Optional:            true,
	},
	"datahub_database": schema.StringAttribute{
		MarkdownDescription: "Database type for datahub. Currently supported values: NONE, NON_HA, HA",
		Description:         "Database type for datahub. Currently supported values: NONE, NON_HA, HA",
		Optional:            true,
	},
	"cluster_extension": schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"custom_properties": schema.StringAttribute{
				Optional: true,
			},
		},
	},
	"java_version": schema.Int64Attribute{
		MarkdownDescription: "Configure the major version of Java on the cluster.",
		Optional:            true,
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		ElementType: types.StringType,
	},
}

var instanceGroupSchemaAttributes = map[string]schema.Attribute{
	"instance_group": schema.ListNestedAttribute{
		Optional: true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"node_count": schema.Int64Attribute{
					MarkdownDescription: "The cluster node count. Has to be greater or equal than 0 and less than 100,000.",
					Required:            true,
				},
				"instance_group_name": schema.StringAttribute{
					MarkdownDescription: "The name of the instance group.",
					Required:            true,
				},
				"instance_group_type": schema.StringAttribute{
					MarkdownDescription: "The type of the instance group.",
					Required:            true,
				},
				"instance_type": schema.StringAttribute{
					MarkdownDescription: "The cloud provider-side instance type.",
					Required:            true,
				},
				"root_volume_size": schema.Int64Attribute{
					MarkdownDescription: "The size of the root volume in GB",
					Required:            true,
				},
				"recipes": schema.SetAttribute{
					MarkdownDescription: "The set of recipe names that are going to be applied on the given instance group.",
					ElementType:         types.StringType,
					Optional:            true,
				},
				"availability_zones": schema.SetAttribute{
					MarkdownDescription: "The set of availability zones that are going to be used for cluster creation on the given instance group.",
					ElementType:         types.StringType,
					Optional:            true,
				},
				"attached_volume_configuration": schema.ListNestedAttribute{
					Required:            true,
					MarkdownDescription: "Configuration regarding the attached volume to the specific instance group.",
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"volume_size": schema.Int64Attribute{
								MarkdownDescription: "The size of the volume in GB.",
								Required:            true,
							},
							"volume_count": schema.Int64Attribute{
								MarkdownDescription: "The number of volumes to be attached.",
								Required:            true,
							},
							"volume_type": schema.StringAttribute{
								MarkdownDescription: "The - cloud provider - type of the volume.",
								Required:            true,
							},
						},
					},
				},
				"recovery_mode": schema.StringAttribute{
					MarkdownDescription: "The type of the recovery mode.",
					Required:            true,
				},
				"volume_encryption": schema.SingleNestedAttribute{
					MarkdownDescription: "The volume encryption related configuration.",
					Required:            true,
					Attributes: map[string]schema.Attribute{
						"encryption": schema.BoolAttribute{
							Required: true,
						},
					},
				},
			},
		},
	},
}

var azureInstanceGroupSchemaAttributes = map[string]schema.Attribute{
	"instance_group": schema.ListNestedAttribute{
		Optional: true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"node_count": schema.Int64Attribute{
					MarkdownDescription: "The cluster node count. Has to be greater or equal than 0 and less than 100,000.",
					Required:            true,
				},
				"availability_zones": schema.SetAttribute{
					MarkdownDescription: "List of availability zones that this instance group is associated with.",
					ElementType:         types.StringType,
					Optional:            true,
				},
				"instance_group_name": schema.StringAttribute{
					MarkdownDescription: "The name of the instance group.",
					Required:            true,
				},
				"instance_group_type": schema.StringAttribute{
					MarkdownDescription: "The type of the instance group.",
					Required:            true,
				},
				"instance_type": schema.StringAttribute{
					MarkdownDescription: "The cloud provider-side instance type.",
					Required:            true,
				},
				"root_volume_size": schema.Int64Attribute{
					MarkdownDescription: "The size of the root volume in GB",
					Required:            true,
				},
				"recipes": schema.SetAttribute{
					MarkdownDescription: "The set of recipe names that are going to be applied on the given instance group.",
					ElementType:         types.StringType,
					Optional:            true,
				},
				"attached_volume_configuration": schema.ListNestedAttribute{
					Required:            true,
					MarkdownDescription: "Configuration regarding the attached volume to the specific instance group.",
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"volume_size": schema.Int64Attribute{
								MarkdownDescription: "The size of the volume in GB.",
								Required:            true,
							},
							"volume_count": schema.Int64Attribute{
								MarkdownDescription: "The number of volumes to be attached.",
								Required:            true,
							},
							"volume_type": schema.StringAttribute{
								MarkdownDescription: "The - cloud provider - type of the volume.",
								Required:            true,
							},
						},
					},
				},
				"recovery_mode": schema.StringAttribute{
					MarkdownDescription: "The type of the recovery mode.",
					Required:            true,
				},
				"volume_encryption": schema.SingleNestedAttribute{
					MarkdownDescription: "The volume encryption related configuration.",
					Required:            true,
					Attributes: map[string]schema.Attribute{
						"encryption": schema.BoolAttribute{
							Required: true,
						},
					},
				},
			},
		},
	},
}
