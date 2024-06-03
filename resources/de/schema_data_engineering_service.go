// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package de

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var serviceSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"chart_value_overrides": schema.SetNestedAttribute{
			MarkdownDescription: "Chart overrides for enabling a service.",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"chart_name": schema.StringAttribute{
						MarkdownDescription: "Name of the chart that has to be overridden, for eg- \"dex-app\", \"dex-base\".",
						Required:            true,
					},
					"overrides": schema.StringAttribute{
						MarkdownDescription: "Space separated key-value pairs for overriding chart values. The key and the value must be separated using a colon(:) For eg- \"airflow.enabled:true safari.enabled:true\".",
						Required:            true,
					},
				},
			},
		},
		"cpu_requests": schema.StringAttribute{
			MarkdownDescription: "CPU Requests for the entire CDE service quota.",
			Optional:            true,
		},
		"custom_azure_files_configs": schema.SingleNestedAttribute{
			MarkdownDescription: "CDE uses a default public File Shares storage provisioned by AKS. Enable this option to use your own public/private File Shares.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"azure_files_fqdn": schema.StringAttribute{
					MarkdownDescription: "Azure File Share's server address. Defaults to '<storageaccount>.file.core.windows.net'.",
					Optional:            true,
				},
				"resource_group": schema.StringAttribute{
					MarkdownDescription: "Resource Group of the Storage Account.",
					Required:            true,
				},
				"storage_account_name": schema.StringAttribute{
					MarkdownDescription: "Azure Storage Account of the File Share.",
					Required:            true,
				},
			},
		},
		"deploy_previous_version": schema.BoolAttribute{
			MarkdownDescription: "If set to \"true\", the previous version of the CDE service will be deployed.",
			Optional:            true,
		},
		"enable_private_network": schema.BoolAttribute{
			MarkdownDescription: "Create a fully private CDE instance on either Amazon or Azure. This includes services such as Kubernetes, MySQL, etc. For Azure, this will also enable virtual network (VNet) access via private endpoints and private link.",
			Optional:            true,
		},
		"enable_public_endpoint": schema.BoolAttribute{
			MarkdownDescription: "Creates a CDE endpoint (Load Balancer) in a publicly accessible subnet. If set false, the endpoint will be created in a private subnet and you will need to setup access to the endpoint manually in your cloud account.",
			Optional:            true,
		},
		"enable_workload_analytics": schema.BoolAttribute{
			MarkdownDescription: "If set false, diagnostic information about job and query execution is sent to Cloudera Workload Manager. Anonymization can be configured under Environments / Shared Resources / Telemetry. Refer documentation for more info at https://docs.cloudera.com/workload-manager/cloud/index.html.",
			Optional:            true,
		},
		"env": schema.StringAttribute{
			MarkdownDescription: "CDP environment where cde service should be enabled.",
			Required:            true,
		},
		"gpu_requests": schema.StringAttribute{
			MarkdownDescription: "GPU requests for the entire CDE service quota.",
			Optional:            true,
		},
		"initial_instances": schema.Int64Attribute{
			MarkdownDescription: "Initial Instances when the service is enabled.",
			Optional:            true,
		},
		"initial_spot_instances": schema.Int64Attribute{
			MarkdownDescription: "Initial spot Instances when the service is enabled.",
			Optional:            true,
		},
		"instance_type": schema.StringAttribute{
			MarkdownDescription: "Instance type of the cluster for CDE Service.",
			Required:            true,
		},
		"loadbalancer_allowlist": schema.SetAttribute{
			MarkdownDescription: "List of CIDRs that would be allowed to access the load balancer.",
			Optional:            true,
			ElementType:         types.StringType,
		},
		"maximum_instances": schema.Int64Attribute{
			MarkdownDescription: "Maximum Instances for the CDE Service.",
			Required:            true,
		},
		"maximum_spot_instances": schema.Int64Attribute{
			MarkdownDescription: "Maximum Spot Instances for the CDE Service.",
			Optional:            true,
		},
		"memory_requests": schema.StringAttribute{
			MarkdownDescription: "Memory requests for the entire CDE service quota, eg. 100Gi.",
			Optional:            true,
		},
		"minimum_instances": schema.Int64Attribute{
			MarkdownDescription: "Minimum Instances for the CDE Service.",
			Required:            true,
		},
		"minimum_spot_instances": schema.Int64Attribute{
			MarkdownDescription: "Minimum Spot instances for the CDE Service.",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "Name of the CDE Service.",
			Required:            true,
		},
		"network_outbound_type": schema.StringAttribute{
			MarkdownDescription: "Network outbound type. Currently 'udr' is the only supported. Enum: [\"UDR\"]",
			Optional:            true,
		},
		"nfs_storage_class": schema.StringAttribute{
			MarkdownDescription: "NFS Storage class to override the default storage class in private cloud.",
			Optional:            true,
		},
		"resource_pool": schema.StringAttribute{
			MarkdownDescription: "Resource Pool for the CDE service.",
			Optional:            true,
		},
		"root_volume_size": schema.Int64Attribute{
			MarkdownDescription: "EBS volume size in GB.",
			Optional:            true,
		},
		"skip_validation": schema.BoolAttribute{
			MarkdownDescription: "Skip Validation check.",
			Optional:            true,
		},
		"subnets": schema.SetAttribute{
			MarkdownDescription: "List of Subnet IDs of CDP subnets to use for the kubernetes worker node.",
			Optional:            true,
			ElementType:         types.StringType,
		},
		"tags": schema.MapAttribute{
			MarkdownDescription: "User defined labels that tag all provisioned cloud resources.",
			Optional:            true,
			ElementType:         types.StringType,
		},
		"use_ssd": schema.BoolAttribute{
			MarkdownDescription: "Instance local storage (SSD) would be used for the workload filesystem (Example - spark local directory). In case the workload requires more space than what's available in the instance storage, please use an instance type with sufficient instance local storage or choose an instance type without SSD and configure the EBS volume size. Currently supported only for aws services.",
			Optional:            true,
		},
		"whitelist_ips": schema.SetAttribute{
			MarkdownDescription: "List of CIDRs that would be allowed to access kubernetes master API server.",
			Optional:            true,
			ElementType:         types.StringType,
		},
	},
}
