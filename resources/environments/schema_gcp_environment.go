// Copyright 2023 Cloudera. All Rights Reserved.
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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var GcpEnvironmentSchema = schema.Schema{
	MarkdownDescription: "The environment is a logical entity that represents the association of your user account with multiple compute resources using which you can provision and manage workloads.",
	Attributes: map[string]schema.Attribute{
		"environment_name": schema.StringAttribute{
			MarkdownDescription: "The name of the environment. Must contain only lowercase letters, numbers and hyphens.",
			Required:            true,
		},
		"cascading_delete": schema.BoolAttribute{
			Optional:           true,
			Computed:           true,
			Default:            booldefault.StaticBool(true),
			DeprecationMessage: "Cascading_delete is deprecated. Use delete_options.cascading instead. If latter specified, it will override this value.",
		},
		"delete_options": schema.SingleNestedAttribute{
			MarkdownDescription: "Options for deleting the environment.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"cascading": schema.BoolAttribute{
					MarkdownDescription: "If true, all resources in the environment will be deleted.",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"forced": schema.BoolAttribute{
					MarkdownDescription: "Force delete action removes CDP resources and may leave cloud provider resources running even if the deletion did not succeed.",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
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
		"credential_name": schema.StringAttribute{
			MarkdownDescription: "Name of the credential to use for the environment.",
			Required:            true,
		},
		"region": schema.StringAttribute{
			MarkdownDescription: "The region of the environment.",
			Required:            true,
		},
		"public_key": schema.StringAttribute{
			MarkdownDescription: "Public SSH key string. The associated private key can be used to get root-level access to the Data Lake instance and Data Hub cluster instances.",
			Required:            true,
		},
		"use_public_ip": schema.BoolAttribute{
			MarkdownDescription: "Whether to associate public IPs to the resources within the network or not.",
			Required:            true,
		},
		"existing_network_params": schema.SingleNestedAttribute{
			MarkdownDescription: "Parameters needed to use an existing VPC and Subnets. For now only existing network params is supported.",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"network_name": schema.StringAttribute{
					MarkdownDescription: "The name of the GCP VPC.",
					Required:            true,
				},
				"subnet_names": schema.ListAttribute{
					MarkdownDescription: "One or more subnet names within the VPC. Google VPCs are global, please give subnets from single geographic region only to reduce latency.",
					Required:            true,
					ElementType:         types.StringType,
				},
				"shared_project_id": schema.StringAttribute{
					MarkdownDescription: "The ID of the Google project associated with the VPC.",
					Required:            true,
				},
			},
		},
		"security_access": schema.SingleNestedAttribute{
			MarkdownDescription: "Firewall rules for FreeIPA, Data Lake and Data Hub deployment.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"security_group_id_for_knox": schema.StringAttribute{
					MarkdownDescription: "Firewall rule for Knox hosts.",
					Optional:            true,
					Computed:            true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
				"default_security_group_id": schema.StringAttribute{
					MarkdownDescription: "Firewall rule for other hosts.",
					Optional:            true,
					Computed:            true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
			},
		},
		"log_storage": schema.SingleNestedAttribute{
			MarkdownDescription: "GCP storage configuration for cluster and audit logs.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"storage_location_base": schema.StringAttribute{
					MarkdownDescription: "The Google storage bucket to use. This should be a gs:// url.",
					Required:            true,
				},
				"service_account_email": schema.StringAttribute{
					MarkdownDescription: "Email id of the service account to be associated with the instances. This service account should have \"storage.ObjectCreator\" role on the given storage bucket.",
					Required:            true,
				},
				"backup_storage_location_base": schema.StringAttribute{
					MarkdownDescription: "The Google storage bucket to use. This should be a gs:// url.",
					Optional:            true,
					Computed:            true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
			},
		},
		"description": schema.StringAttribute{
			MarkdownDescription: "A description of the environment.",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"enable_tunnel": schema.BoolAttribute{
			MarkdownDescription: "Whether to enable SSH tunneling for the environment.",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(true),
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"workload_analytics": schema.BoolAttribute{
			MarkdownDescription: "When this is enabled, diagnostic information about job and query execution is sent to Workload Manager for Data Hub clusters created within this environment.",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"report_deployment_logs": schema.BoolAttribute{
			// report_deployment_logs is a deprecated field and should not be used
			MarkdownDescription: " [Deprecated] When true, this will report additional diagnostic information back to Cloudera.",
			Computed:            true,
			Default:             booldefault.StaticBool(false),
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"freeipa": FreeIpaSchema,
		"endpoint_access_gateway_scheme": schema.StringAttribute{
			MarkdownDescription: "The scheme for the endpoint gateway. PUBLIC creates an external endpoint that can be accessed over the Internet. Defaults to PRIVATE which restricts the traffic to be internal to the VPC.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"tags": schema.MapAttribute{
			MarkdownDescription: "Tags that can be attached to GCP resources. Please refer to Google documentation for the rules https://cloud.google.com/compute/docs/labeling-resources#label_format.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers: []planmodifier.Map{
				mapplanmodifier.UseStateForUnknown(),
			},
		},
		"proxy_config_name": schema.StringAttribute{
			MarkdownDescription: "Name of the proxy config to use for the environment.",
			Optional:            true,
		},
		"encryption_key": schema.StringAttribute{
			MarkdownDescription: "Key Resource ID of the customer managed encryption key to encrypt GCP resources.",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString(""),
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"availability_zones": schema.SetAttribute{
			MarkdownDescription: "The zones of the environment in the given region. Multi-zone selection is not supported in GCP yet. It accepts only one zone until support is added.",
			Optional:            true,
			ElementType:         types.StringType,
		},
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"crn": schema.StringAttribute{
			MarkdownDescription: "The CRN of the environment resource.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"status": schema.StringAttribute{
			MarkdownDescription: "The last known status for the environment.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"status_reason": schema.StringAttribute{
			MarkdownDescription: "The last known detailed status reason for the environment.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
	},
}
