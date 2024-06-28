// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package aws

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	dwclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
	mocks "github.com/cloudera/terraform-provider-cdp/mocks/github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/client/operations"
)

var testDwClusterSchema = schema.Schema{
	MarkdownDescription: "Creates an AWS Data Warehouse cluster.",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"crn": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The cloudera resource name of the environment that the cluster will read from.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The name of the cluster matches the environment name.",
		},
		"cluster_id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The id of the cluster.",
		},
		"last_updated": schema.StringAttribute{
			Description: "Timestamp of the last Terraform update of the order.",
			Computed:    true,
		},
		"node_role_cdw_managed_policy_arn": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "The managed policy ARN to be attached to the created node instance role.",
		},
		"database_backup_retention_days": schema.Int64Attribute{
			Optional:            true,
			MarkdownDescription: "The number of days to retain database backups.",
		},
		"custom_registry_options": schema.SingleNestedAttribute{
			Optional: true,
			Attributes: map[string]schema.Attribute{
				"registry_type": schema.StringAttribute{
					Required:            true,
					MarkdownDescription: "Registry type, supported values are ECR or ACR.",
				},
				"repository_url": schema.StringAttribute{
					Required:            true,
					MarkdownDescription: "The URL of the registry.",
				},
			},
		},
		"custom_subdomain": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "The custom subdomain to keep compatibility with old URL format.",
		},
		"network_settings": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"worker_subnet_ids": schema.ListAttribute{
					Required:            true,
					ElementType:         types.StringType,
					MarkdownDescription: "The list of subnet IDs for worker nodes.",
				},
				"load_balancer_subnet_ids": schema.ListAttribute{
					Required:            true,
					ElementType:         types.StringType,
					MarkdownDescription: "The list of subnet IDs for the load balancer.",
				},
				"use_overlay_network": schema.BoolAttribute{
					Required:            true,
					MarkdownDescription: "Whether to use overlay network.",
				},
				"whitelist_k8s_cluster_access_ip_cidrs": schema.ListAttribute{
					Optional:            true,
					ElementType:         types.StringType,
					MarkdownDescription: "The list of IP CIDRs to allow access for kubernetes cluster API endpoint.",
				},
				"whitelist_workload_access_ip_cidrs": schema.ListAttribute{
					Optional:            true,
					ElementType:         types.StringType,
					MarkdownDescription: "The list of IP CIDRs to allow access for workload endpoints.",
				},
				"use_private_load_balancer": schema.BoolAttribute{
					Required:            true,
					MarkdownDescription: "Whether to use private IP addresses for the load balancer. Determines workload endpoint access.",
				},
				"use_public_worker_node": schema.BoolAttribute{
					Required:            true,
					MarkdownDescription: "Whether to use public IP addresses for worker nodes.",
				},
			},
		},
		"instance_settings": schema.SingleNestedAttribute{
			Optional: true,
			Attributes: map[string]schema.Attribute{
				"custom_ami_id": schema.StringAttribute{
					Optional:            true,
					MarkdownDescription: "The custom AMI ID to use for worker nodes.",
				},
				"enable_spot_instances": schema.BoolAttribute{
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
					MarkdownDescription: "Whether to use spot instances for worker nodes.",
				},
				"compute_instance_types": schema.ListAttribute{
					Optional:            true,
					ElementType:         types.StringType,
					MarkdownDescription: "The compute instance types that the environment is restricted to use. This affects the creation of virtual warehouses where this restriction will apply. Select an instance type that meets your computing, memory, networking, or storage needs. As of now, only a single instance type can be listed.",
				},
				"additional_instance_types": schema.ListAttribute{
					Optional:            true,
					ElementType:         types.StringType,
					MarkdownDescription: "The additional instance types that the environment is allowed to use, listed in their priority order. They will be used instead of the primary compute instance type in case it is unavailable. You cannot include any instance type that was already indicated in computeInstanceTypes.",
				},
			},
		},
	},
}

type MockTransport struct {
	runtime.ClientTransport
}

func NewDwApi(client *mocks.MockDwClientService) *dwClusterResource {
	return &dwClusterResource{
		client: &cdp.Client{
			Dw: &dwclient.Dw{
				Operations: client,
				Transport:  MockTransport{},
			}}}
}

func createRawClusterResource() tftypes.Value {
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"id":                               tftypes.String,
				"crn":                              tftypes.String,
				"name":                             tftypes.String,
				"cluster_id":                       tftypes.String,
				"last_updated":                     tftypes.String,
				"node_role_cdw_managed_policy_arn": tftypes.String,
				"database_backup_retention_days":   tftypes.Number,
				"custom_registry_options": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"registry_type":  tftypes.String,
						"repository_url": tftypes.String,
					},
				},
				"custom_subdomain": tftypes.String,
				"network_settings": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"worker_subnet_ids":                     tftypes.List{ElementType: tftypes.String},
						"load_balancer_subnet_ids":              tftypes.List{ElementType: tftypes.String},
						"use_overlay_network":                   tftypes.Bool,
						"whitelist_k8s_cluster_access_ip_cidrs": tftypes.List{ElementType: tftypes.String},
						"whitelist_workload_access_ip_cidrs":    tftypes.List{ElementType: tftypes.String},
						"use_private_load_balancer":             tftypes.Bool,
						"use_public_worker_node":                tftypes.Bool,
					},
				},
				"instance_settings": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"custom_ami_id":             tftypes.String,
						"enable_spot_instances":     tftypes.Bool,
						"compute_instance_types":    tftypes.List{ElementType: tftypes.String},
						"additional_instance_types": tftypes.List{ElementType: tftypes.String},
					},
				},
			}}, map[string]tftypes.Value{
			"id":                               tftypes.NewValue(tftypes.String, ""),
			"crn":                              tftypes.NewValue(tftypes.String, "crn"),
			"name":                             tftypes.NewValue(tftypes.String, ""),
			"cluster_id":                       tftypes.NewValue(tftypes.String, ""),
			"last_updated":                     tftypes.NewValue(tftypes.String, ""),
			"node_role_cdw_managed_policy_arn": tftypes.NewValue(tftypes.String, ""),
			"database_backup_retention_days":   tftypes.NewValue(tftypes.Number, 0),
			"custom_registry_options": tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"registry_type":  tftypes.String,
					"repository_url": tftypes.String,
				},
			}, map[string]tftypes.Value{
				"registry_type":  tftypes.NewValue(tftypes.String, ""),
				"repository_url": tftypes.NewValue(tftypes.String, ""),
			}),
			"custom_subdomain": tftypes.NewValue(tftypes.String, ""),
			"network_settings": tftypes.NewValue(
				tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"worker_subnet_ids":                     tftypes.List{ElementType: tftypes.String},
						"load_balancer_subnet_ids":              tftypes.List{ElementType: tftypes.String},
						"use_overlay_network":                   tftypes.Bool,
						"whitelist_k8s_cluster_access_ip_cidrs": tftypes.List{ElementType: tftypes.String},
						"whitelist_workload_access_ip_cidrs":    tftypes.List{ElementType: tftypes.String},
						"use_private_load_balancer":             tftypes.Bool,
						"use_public_worker_node":                tftypes.Bool,
					},
				}, map[string]tftypes.Value{
					"worker_subnet_ids": tftypes.NewValue(tftypes.List{ElementType: tftypes.String},
						[]tftypes.Value{
							tftypes.NewValue(tftypes.String, "subnet-1"),
							tftypes.NewValue(tftypes.String, "subnet-2"),
							tftypes.NewValue(tftypes.String, "subnet-3"),
						}),
					"load_balancer_subnet_ids": tftypes.NewValue(tftypes.List{ElementType: tftypes.String},
						[]tftypes.Value{
							tftypes.NewValue(tftypes.String, "subnet-4"),
							tftypes.NewValue(tftypes.String, "subnet-5"),
							tftypes.NewValue(tftypes.String, "subnet-6"),
						}),
					"use_overlay_network": tftypes.NewValue(tftypes.Bool, true),
					"whitelist_k8s_cluster_access_ip_cidrs": tftypes.NewValue(tftypes.List{ElementType: tftypes.String},
						[]tftypes.Value{
							tftypes.NewValue(tftypes.String, "cidr-1"),
							tftypes.NewValue(tftypes.String, "cidr-2"),
							tftypes.NewValue(tftypes.String, "cidr-3"),
						}),
					"whitelist_workload_access_ip_cidrs": tftypes.NewValue(tftypes.List{ElementType: tftypes.String},
						[]tftypes.Value{
							tftypes.NewValue(tftypes.String, "cidr-4"),
							tftypes.NewValue(tftypes.String, "cidr-2"),
							tftypes.NewValue(tftypes.String, "cidr-3"),
						}),
					"use_private_load_balancer": tftypes.NewValue(tftypes.Bool, true),
					"use_public_worker_node":    tftypes.NewValue(tftypes.Bool, false),
				},
			),
			"instance_settings": tftypes.NewValue(
				tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"custom_ami_id":             tftypes.String,
						"enable_spot_instances":     tftypes.Bool,
						"compute_instance_types":    tftypes.List{ElementType: tftypes.String},
						"additional_instance_types": tftypes.List{ElementType: tftypes.String},
					}}, map[string]tftypes.Value{
					"custom_ami_id":             tftypes.NewValue(tftypes.String, ""),
					"enable_spot_instances":     tftypes.NewValue(tftypes.Bool, false),
					"compute_instance_types":    tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{}),
					"additional_instance_types": tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{}),
				},
			),
		})
}

type DwClusterTestSuite struct {
	suite.Suite
	expectedCreateResponse *operations.CreateAwsClusterOK
}

func TestDwAwsClusterTestSuite(t *testing.T) {
	suite.Run(t, new(DwClusterTestSuite))
}

func (suite *DwClusterTestSuite) SetupTest() {
	suite.expectedCreateResponse = &operations.CreateAwsClusterOK{
		Payload: &models.CreateAwsClusterResponse{
			ClusterID: "cluster-id"}}
}

func (suite *DwClusterTestSuite) TestDwAwsClusterMetadata() {
	dwApi := NewDwApi(new(mocks.MockDwClientService))
	resp := resource.MetadataResponse{}

	// Function under test
	dwApi.Metadata(
		context.TODO(),
		resource.MetadataRequest{ProviderTypeName: "cdp"},
		&resp,
	)
	suite.Equal("cdp_dw_aws_cluster", resp.TypeName)
}

func (suite *DwClusterTestSuite) TestDwAwsClusterSchema() {
	dwApi := NewDwApi(new(mocks.MockDwClientService))
	resp := resource.SchemaResponse{}

	// Function under test
	dwApi.Schema(
		context.TODO(),
		resource.SchemaRequest{},
		&resp,
	)
	suite.Equal(testDwClusterSchema, resp.Schema)
}

func (suite *DwClusterTestSuite) TestDwAwsClusterCreate_Success() {
	ctx := context.TODO()
	expectedDescribeResponse := &operations.DescribeClusterOK{
		Payload: &models.DescribeClusterResponse{
			Cluster: &models.ClusterSummaryResponse{
				EnvironmentCrn: "crn",
				ID:             "cluster-id",
				Name:           "test-name",
			}}}

	client := new(mocks.MockDwClientService)
	client.On("CreateAwsCluster", mock.Anything).Return(suite.expectedCreateResponse, nil)
	client.On("DescribeCluster", mock.Anything).Return(expectedDescribeResponse, nil)
	dwApi := NewDwApi(client)

	req := resource.CreateRequest{
		Plan: tfsdk.Plan{
			Raw:    createRawClusterResource(),
			Schema: testDwClusterSchema,
		},
	}

	resp := &resource.CreateResponse{
		State: tfsdk.State{
			Schema: testDwClusterSchema,
		},
	}

	// Function under test
	dwApi.Create(ctx, req, resp)
	var result resourceModel
	resp.State.Get(ctx, &result)
	suite.False(resp.Diagnostics.HasError())
	suite.Equal("crn", result.ID.ValueString())
	suite.Equal("crn", result.Crn.ValueString())
	suite.Equal("cluster-id", result.ClusterID.ValueString())
	suite.Equal("test-name", result.Name.ValueString())
}

func (suite *DwClusterTestSuite) TestDwAwsClusterCreate_CreationError() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("CreateAwsCluster", mock.Anything).Return(&operations.CreateAwsClusterOK{}, fmt.Errorf("create failed"))
	client.On("DescribeCluster", mock.Anything).Return(&operations.DescribeClusterOK{}, nil)
	dwApi := NewDwApi(client)

	req := resource.CreateRequest{
		Plan: tfsdk.Plan{
			Raw:    createRawClusterResource(),
			Schema: testDwClusterSchema,
		},
	}

	resp := &resource.CreateResponse{
		State: tfsdk.State{
			Schema: testDwClusterSchema,
		},
	}

	// Function under test
	dwApi.Create(ctx, req, resp)
	var result resourceModel
	resp.State.Get(ctx, &result)
	suite.True(resp.Diagnostics.HasError())
	suite.Contains(resp.Diagnostics.Errors()[0].Summary(), "Error creating data warehouse aws cluster")
	suite.Contains(resp.Diagnostics.Errors()[0].Detail(), "Could not create cluster")
}

func (suite *DwClusterTestSuite) TestDwAwsClusterCreate_DescribeError() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("CreateAwsCluster", mock.Anything).Return(suite.expectedCreateResponse, nil)
	client.On("DescribeCluster", mock.Anything).Return(&operations.DescribeClusterOK{}, fmt.Errorf("describe failed"))
	dwApi := NewDwApi(client)

	req := resource.CreateRequest{
		Plan: tfsdk.Plan{
			Raw:    createRawClusterResource(),
			Schema: testDwClusterSchema,
		},
	}

	resp := &resource.CreateResponse{
		State: tfsdk.State{
			Schema: testDwClusterSchema,
		},
	}

	// Function under test
	dwApi.Create(ctx, req, resp)
	var result resourceModel
	resp.State.Get(ctx, &result)
	suite.True(resp.Diagnostics.HasError())
	suite.Contains(resp.Diagnostics.Errors()[0].Summary(), "Error creating data warehouse aws cluster")
	suite.Contains(resp.Diagnostics.Errors()[0].Detail(), "Could not describe cluster")
}

func (suite *DwClusterTestSuite) TestDwAwsClusterDeletion_Success() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("DeleteCluster", mock.Anything).Return(&operations.DeleteClusterOK{}, nil)
	dwApi := NewDwApi(client)

	req := resource.DeleteRequest{
		State: tfsdk.State{
			Schema: testDwClusterSchema,
			Raw:    createRawClusterResource(),
		},
	}

	resp := &resource.DeleteResponse{}

	// Function under test
	dwApi.Delete(ctx, req, resp)
	suite.False(resp.Diagnostics.HasError())
}

func (suite *DwClusterTestSuite) TestDwAwsClusterDeletion_ReturnsError() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("DeleteCluster", mock.Anything).Return(&operations.DeleteClusterOK{}, fmt.Errorf("delete failed"))
	dwApi := NewDwApi(client)

	req := resource.DeleteRequest{
		State: tfsdk.State{
			Schema: testDwClusterSchema,
			Raw:    createRawClusterResource(),
		},
	}

	resp := &resource.DeleteResponse{}

	// Function under test
	dwApi.Delete(ctx, req, resp)
	suite.True(resp.Diagnostics.HasError())
	suite.Contains(resp.Diagnostics.Errors()[0].Summary(), "Error deleting data warehouse aws cluster")
	suite.Contains(resp.Diagnostics.Errors()[0].Detail(), "Could not delete cluster")
}
