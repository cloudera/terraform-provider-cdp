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
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

var gcpInstanceGroupSchemaAttributes = map[string]schema.Attribute{
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
			},
		},
	},
}

func (r *gcpDatahubResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	attr := map[string]schema.Attribute{}
	utils.Append(attr, generalAttributes)
	utils.Append(attr, gcpInstanceGroupSchemaAttributes)
	utils.Append(attr, map[string]schema.Attribute{
		"cluster_template_name": schema.StringAttribute{
			MarkdownDescription: "The name of the cluster template.",
			Optional:            true,
		},
		"cluster_definition_name": schema.StringAttribute{
			MarkdownDescription: "The name of the cluster definition.",
			Optional:            true,
		},
		"environment_name": schema.StringAttribute{
			MarkdownDescription: "The name of the environment where the cluster will belong to.",
			Required:            true,
		},
		"subnet_name": schema.StringAttribute{
			MarkdownDescription: "The subnet name.",
			Optional:            true,
		},
	})
	removeRequiredNonGcpAttributes(attr)
	resp.Schema = schema.Schema{
		MarkdownDescription: "Creates an GCP Data hub cluster.",
		Attributes:          attr,
	}
}

/*
This is a workaround for now since the following keys are uniform over Azure and AWS but as of now for GCP they are
different. Until all these fields will get the same name this deletion-addition shall make less code repetitions
over the schema implementations.
*/
func removeRequiredNonGcpAttributes(attr map[string]schema.Attribute) {
	delete(attr, "environment")
	delete(attr, "cluster_template")
	delete(attr, "cluster_definition")
}
