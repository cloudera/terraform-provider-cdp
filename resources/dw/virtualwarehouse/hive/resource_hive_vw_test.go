// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package hive

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
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
	"github.com/cloudera/terraform-provider-cdp/mocks"
)

var testHiveSchema = schema.Schema{
	MarkdownDescription: "A Hive Virtual Warehouse is service which is able to run big SQL queries.",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"cluster_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The id of the CDW Cluster which the Hive Virtual Warehouse is attached to.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"database_catalog_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The id of the Database Catalog which the Hive Virtual Warehouse is attached to.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The name of the Hive Virtual Warehouse.",
		},
		"image_version": schema.StringAttribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "The version of the Hive Virtual Warehouse image.",
		},
		"group_size": schema.Int64Attribute{
			Required:            true,
			MarkdownDescription: "Nodes per compute group. If specified, forces ‘template’ to be ‘custom’.",
		},
		"platform_jwt_auth": schema.BoolAttribute{
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
			MarkdownDescription: "Value of ‘true’ automatically configures the Virtual Warehouse to support JWTs issued by the CDP JWT token provider. Value of ‘false’ does not enable JWT auth on the Virtual Warehouse. If this field is not specified, it defaults to ‘false’.",
		},
		"ldap_groups": schema.ListAttribute{
			Optional:            true,
			ElementType:         types.StringType,
			MarkdownDescription: "LDAP group names to be enabled to authenticate with.",
		},
		"enable_sso": schema.BoolAttribute{
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
			MarkdownDescription: "Enable SSO for the Virtual Warehouse. If this field is not specified, it defaults to ‘false’.",
		},
		"compactor": schema.BoolAttribute{
			Computed:            true,
			MarkdownDescription: "Boolean value that describes if the Hive Virtual Warehouse is a compactor.",
		},
		"jdbc_url": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "JDBC URL for the Hive Virtual Warehouse.",
		},
		"kerberos_jdbc_url": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Kerberos JDBC URL for the Hive Virtual Warehouse.",
		},
		"hue_url": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Hue URL for the Hive Virtual Warehouse.",
		},
		"jwt_connection_string": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Generic semi-colon delimited list of key-value pairs that contain all necessary information for clients to construct a connection to this Virtual Warehouse using JWTs as the authentication method.",
		},
		"jwt_token_gen_url": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "URL to generate JWT tokens for the Virtual Warehouse by the CDP JWT token provider. Available if platform JWT authentication is enabled.",
		},
		"min_group_count": schema.Int64Attribute{
			Required:            true,
			MarkdownDescription: "Minimum number of available compute groups.",
		},
		"max_group_count": schema.Int64Attribute{
			Required:            true,
			MarkdownDescription: "Maximum number of available compute groups.",
		},
		"disable_auto_suspend": schema.BoolAttribute{
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
			MarkdownDescription: "Boolean value that specifies if auto-suspend should be disabled.",
		},
		"auto_suspend_timeout_seconds": schema.Int64Attribute{
			Optional:            true,
			MarkdownDescription: "The time in seconds after which the compute group should be suspended.",
		},
		"scale_wait_time_seconds": schema.Int64Attribute{
			Optional:            true,
			MarkdownDescription: "Set wait time before a scale event happens.",
		},
		"headroom": schema.Int64Attribute{
			Optional:            true,
			MarkdownDescription: "Set headroom node count. Nodes will be started in case there are no free nodes left to pick up new jobs.",
		},
		"max_concurrent_isolated_queries": schema.Int64Attribute{
			Optional:            true,
			MarkdownDescription: "Maximum number of concurrent isolated queries. If not provided, 0 will be applied. The 0 value means the query isolation functionality will be disabled.",
		},
		"max_nodes_per_isolated_query": schema.Int64Attribute{
			Optional:            true,
			MarkdownDescription: "Maximum number of nodes per isolated query. If not provided, 0 will be applied. The 0 value means the query isolation functionality will be disabled.",
		},
		"aws_options": schema.SingleNestedAttribute{
			MarkdownDescription: "AWS related configuration options that could specify various values that will be used during CDW resource creation.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"availability_zone": schema.StringAttribute{
					Optional:            true,
					MarkdownDescription: "This feature works only for AWS cluster type. An availability zone to host compute instances. If not specified, defaults to a randomly selected availability zone inferred from available subnets.",
				},
				"ebs_llap_spill_gb": schema.Int64Attribute{
					Optional:            true,
					MarkdownDescription: "This feature works only for AWS cluster type. The size of the EBS volume in GB to be used for LLAP spill storage. If not specified, defaults to no extra spill disk.",
				},
				"tags": schema.MapAttribute{
					Optional:            true,
					ElementType:         types.StringType,
					MarkdownDescription: "This feature works only for AWS cluster type. Tags to be applied to the underlying compute nodes.",
				},
			},
		},
		"last_updated": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Timestamp of the last Terraform update of the order.",
		},
		"status": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The status of the Hive Virtual Warehouse.",
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
					Default:             int64default.StaticInt64(40),
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
	},
}

type MockTransport struct {
	runtime.ClientTransport
}

func newDwApi(client *mocks.MockDwClientService) *hiveResource {
	return &hiveResource{
		client: &cdp.Client{
			Dw: &dwclient.Dw{
				Operations: client,
				Transport:  MockTransport{},
			}}}
}

func createRawHiveResource() tftypes.Value {
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"id":                  tftypes.String,
				"cluster_id":          tftypes.String,
				"database_catalog_id": tftypes.String,
				"name":                tftypes.String,
				"image_version":       tftypes.String,
				"group_size":          tftypes.Number,
				"platform_jwt_auth":   tftypes.Bool,
				"ldap_groups": tftypes.List{
					ElementType: tftypes.String,
				},
				"enable_sso":                      tftypes.Bool,
				"compactor":                       tftypes.Bool,
				"jdbc_url":                        tftypes.String,
				"kerberos_jdbc_url":               tftypes.String,
				"hue_url":                         tftypes.String,
				"jwt_connection_string":           tftypes.String,
				"jwt_token_gen_url":               tftypes.String,
				"min_group_count":                 tftypes.Number,
				"max_group_count":                 tftypes.Number,
				"disable_auto_suspend":            tftypes.Bool,
				"auto_suspend_timeout_seconds":    tftypes.Number,
				"scale_wait_time_seconds":         tftypes.Number,
				"headroom":                        tftypes.Number,
				"max_concurrent_isolated_queries": tftypes.Number,
				"max_nodes_per_isolated_query":    tftypes.Number,
				"aws_options": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"availability_zone": tftypes.String,
						"ebs_llap_spill_gb": tftypes.Number,
						"tags":              tftypes.Map{ElementType: tftypes.String},
					},
				},
				"last_updated": tftypes.String,
				"status":       tftypes.String,
				"polling_options": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"async":                  tftypes.Bool,
						"polling_timeout":        tftypes.Number,
						"call_failure_threshold": tftypes.Number,
					},
				},
			}},
		map[string]tftypes.Value{
			"id":                              tftypes.NewValue(tftypes.String, ""),
			"cluster_id":                      tftypes.NewValue(tftypes.String, "cluster-id"),
			"database_catalog_id":             tftypes.NewValue(tftypes.String, "database-catalog-id"),
			"name":                            tftypes.NewValue(tftypes.String, ""),
			"image_version":                   tftypes.NewValue(tftypes.String, ""),
			"group_size":                      tftypes.NewValue(tftypes.Number, 10),
			"platform_jwt_auth":               tftypes.NewValue(tftypes.Bool, true),
			"ldap_groups":                     tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{}),
			"enable_sso":                      tftypes.NewValue(tftypes.Bool, true),
			"compactor":                       tftypes.NewValue(tftypes.Bool, false),
			"jdbc_url":                        tftypes.NewValue(tftypes.String, ""),
			"kerberos_jdbc_url":               tftypes.NewValue(tftypes.String, ""),
			"hue_url":                         tftypes.NewValue(tftypes.String, ""),
			"jwt_connection_string":           tftypes.NewValue(tftypes.String, ""),
			"jwt_token_gen_url":               tftypes.NewValue(tftypes.String, ""),
			"min_group_count":                 tftypes.NewValue(tftypes.Number, 1),
			"max_group_count":                 tftypes.NewValue(tftypes.Number, 10),
			"disable_auto_suspend":            tftypes.NewValue(tftypes.Bool, false),
			"auto_suspend_timeout_seconds":    tftypes.NewValue(tftypes.Number, 60),
			"scale_wait_time_seconds":         tftypes.NewValue(tftypes.Number, 60),
			"headroom":                        tftypes.NewValue(tftypes.Number, 10),
			"max_concurrent_isolated_queries": tftypes.NewValue(tftypes.Number, 10),
			"max_nodes_per_isolated_query":    tftypes.NewValue(tftypes.Number, 10),
			"aws_options": tftypes.NewValue(
				tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"availability_zone": tftypes.String,
						"ebs_llap_spill_gb": tftypes.Number,
						"tags":              tftypes.Map{ElementType: tftypes.String},
					}}, map[string]tftypes.Value{
					"availability_zone": tftypes.NewValue(tftypes.String, "us-west-2a"),
					"ebs_llap_spill_gb": tftypes.NewValue(tftypes.Number, 300),
					"tags": tftypes.NewValue(tftypes.Map{ElementType: tftypes.String}, map[string]tftypes.Value{
						"key1":  tftypes.NewValue(tftypes.String, "value1"),
						"owner": tftypes.NewValue(tftypes.String, "dw-terraform@cloudera.com"),
					}),
				}),
			"last_updated": tftypes.NewValue(tftypes.String, ""),
			"status":       tftypes.NewValue(tftypes.String, "Running"),
			"polling_options": tftypes.NewValue(
				tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"async":                  tftypes.Bool,
						"polling_timeout":        tftypes.Number,
						"call_failure_threshold": tftypes.Number,
					}}, map[string]tftypes.Value{
					"async":                  tftypes.NewValue(tftypes.Bool, true),
					"polling_timeout":        tftypes.NewValue(tftypes.Number, 90),
					"call_failure_threshold": tftypes.NewValue(tftypes.Number, 3),
				}),
		})
}

type HiveTestSuite struct {
	suite.Suite
	expectedCreateResponse *operations.CreateVwOK
}

func TestHive(t *testing.T) {
	suite.Run(t, new(HiveTestSuite))
}

func (suite *HiveTestSuite) SetupTest() {
	suite.expectedCreateResponse = &operations.CreateVwOK{Payload: &models.CreateVwResponse{
		VwID: "test-id",
	}}

}

func (suite *HiveTestSuite) TestHiveMetadata() {
	dwApi := newDwApi(new(mocks.MockDwClientService))
	resp := resource.MetadataResponse{}

	// Function under test
	dwApi.Metadata(
		context.TODO(),
		resource.MetadataRequest{ProviderTypeName: "cdp"},
		&resp,
	)
	suite.Equal("cdp_dw_vw_hive", resp.TypeName)
}

func (suite *HiveTestSuite) TestHiveSchema() {
	dwApi := newDwApi(new(mocks.MockDwClientService))
	resp := resource.SchemaResponse{}

	// Function under test
	dwApi.Schema(
		context.TODO(),
		resource.SchemaRequest{},
		&resp,
	)
	suite.Equal(testHiveSchema, resp.Schema)
}

func (suite *HiveTestSuite) TestHiveCreate_Success() {
	ctx := context.TODO()
	expectedDescribeResponse := &operations.DescribeVwOK{
		Payload: &models.DescribeVwResponse{
			Vw: &models.VwSummary{
				ID:        "test-id",
				DbcID:     "database-catalog-id",
				Name:      "test-name",
				VwType:    models.VwTypeHive,
				Compactor: true,
				Endpoints: &models.VwSummaryEndpoints{
					HiveJdbc:            "jdbc://hive",
					HiveKerberosJdbc:    "jdbc://hive",
					Hue:                 "https://hue",
					JwtConnectionString: "connection-string",
					JwtTokenGenURL:      "https://jwt",
				},
			}}}

	client := new(mocks.MockDwClientService)
	client.On("CreateVw", mock.Anything).Return(suite.expectedCreateResponse, nil)
	client.On("DescribeVw", mock.Anything).Return(expectedDescribeResponse, nil)
	dwApi := newDwApi(client)

	req := resource.CreateRequest{
		Plan: tfsdk.Plan{
			Raw:    createRawHiveResource(),
			Schema: testHiveSchema,
		},
	}

	resp := &resource.CreateResponse{
		State: tfsdk.State{
			Schema: testHiveSchema,
		},
	}

	// Function under test
	dwApi.Create(ctx, req, resp)
	var result resourceModel
	resp.State.Get(ctx, &result)
	suite.False(resp.Diagnostics.HasError())
	suite.Equal("test-id", result.ID.ValueString())
	suite.Equal("database-catalog-id", result.DatabaseCatalogID.ValueString())
	suite.Equal("cluster-id", result.ClusterID.ValueString())
	suite.Equal("test-name", result.Name.ValueString())
}

func (suite *HiveTestSuite) TestHiveValidate_Headroom_ScaleWaitTime() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	dwApi := newDwApi(client)

	req := resource.ValidateConfigRequest{
		Config: tfsdk.Config{
			Raw:    createRawHiveResource(),
			Schema: testHiveSchema,
		},
	}

	resp := &resource.ValidateConfigResponse{
		Diagnostics: diag.Diagnostics{},
	}

	// Function under test
	validator := dwApi.ConfigValidators(ctx)
	for _, v := range validator {
		v.ValidateResource(ctx, req, resp)
	}
	suite.Contains(resp.Diagnostics.Errors()[0].Summary(), "Invalid Attribute Combination")
	suite.Contains(resp.Diagnostics.Errors()[0].Detail(), "These attributes cannot be configured together: [scale_wait_time_seconds,headroom]")
}

func (suite *HiveTestSuite) TestHiveCreate_CreationError() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("CreateVw", mock.Anything).Return(&operations.CreateVwOK{}, fmt.Errorf("create failed"))
	client.On("DescribeVw", mock.Anything).Return(&operations.DescribeVwOK{}, nil)
	dwApi := newDwApi(client)

	req := resource.CreateRequest{
		Plan: tfsdk.Plan{
			Raw:    createRawHiveResource(),
			Schema: testHiveSchema,
		},
	}

	resp := &resource.CreateResponse{
		State: tfsdk.State{
			Schema: testHiveSchema,
		},
	}

	// Function under test
	dwApi.Create(ctx, req, resp)
	var result resourceModel
	resp.State.Get(ctx, &result)
	suite.True(resp.Diagnostics.HasError())
	suite.Contains(resp.Diagnostics.Errors()[0].Summary(), "Error creating hive virtual warehouse")
	suite.Contains(resp.Diagnostics.Errors()[0].Detail(), "Could not create hive")
}

func (suite *HiveTestSuite) TestHiveCreate_DescribeError() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("CreateVw", mock.Anything).Return(suite.expectedCreateResponse, nil)
	client.On("DescribeVw", mock.Anything).Return(&operations.DescribeVwOK{}, fmt.Errorf("describe failed"))
	dwApi := newDwApi(client)

	req := resource.CreateRequest{
		Plan: tfsdk.Plan{
			Raw:    createRawHiveResource(),
			Schema: testHiveSchema,
		},
	}

	resp := &resource.CreateResponse{
		State: tfsdk.State{
			Schema: testHiveSchema,
		},
	}

	// Function under test
	dwApi.Create(ctx, req, resp)
	var result resourceModel
	resp.State.Get(ctx, &result)
	suite.True(resp.Diagnostics.HasError())
	suite.Contains(resp.Diagnostics.Errors()[0].Summary(), "Error creating hive virtual warehouse")
	suite.Contains(resp.Diagnostics.Errors()[0].Detail(), "Could not describe hive")
}

func (suite *HiveTestSuite) TestHiveDeletion_Success() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("DeleteVw", mock.Anything).Return(&operations.DeleteVwOK{}, nil)
	dwApi := newDwApi(client)

	req := resource.DeleteRequest{
		State: tfsdk.State{
			Schema: testHiveSchema,
			Raw:    createRawHiveResource(),
		},
	}

	resp := &resource.DeleteResponse{}

	// Function under test
	dwApi.Delete(ctx, req, resp)
	suite.False(resp.Diagnostics.HasError())
}

func (suite *HiveTestSuite) TestHiveDeletion_ReturnsError() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("DeleteVw", mock.Anything).Return(&operations.DeleteVwOK{}, fmt.Errorf("delete failed"))
	dwApi := newDwApi(client)

	req := resource.DeleteRequest{
		State: tfsdk.State{
			Schema: testHiveSchema,
			Raw:    createRawHiveResource(),
		},
	}

	resp := &resource.DeleteResponse{}

	// Function under test
	dwApi.Delete(ctx, req, resp)
	suite.True(resp.Diagnostics.HasError())
}

func (suite *HiveTestSuite) TestStateRefresh_Success() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("DescribeVw", mock.Anything).Return(
		&operations.DescribeVwOK{
			Payload: &models.DescribeVwResponse{
				Vw: &models.VwSummary{
					ID:     "hive-id",
					Status: "Running",
				},
			},
		},
		nil)
	dwApi := newDwApi(client)

	clusterID := "cluster-id"
	vwID := "hive-id"
	callFailedCount := 0
	callFailureThreshold := 3

	// Function under test
	refresh := dwApi.stateRefresh(ctx, &clusterID, &vwID, &callFailedCount, callFailureThreshold)
	_, status, err := refresh()
	suite.NoError(err)
	suite.Equal("Running", status)
}

func (suite *HiveTestSuite) TestStateRefresh_FailureThresholdReached() {
	ctx := context.TODO()
	client := new(mocks.MockDwClientService)
	client.On("DescribeVw", mock.Anything).Return(
		&operations.DescribeVwOK{}, fmt.Errorf("unknown error"))
	dwApi := newDwApi(client)

	clusterID := "cluster-id"
	vwID := "hive-id"
	callFailedCount := 0
	callFailureThreshold := 3

	// Function under test
	refresh := dwApi.stateRefresh(ctx, &clusterID, &vwID, &callFailedCount, callFailureThreshold)
	var err error
	for i := 0; i <= callFailureThreshold; i++ {
		_, _, err = refresh()
	}
	suite.Error(err, "unknown error")
}

func (suite *HiveTestSuite) TestConvertToCreateVwRequest_All() {
	ctx := context.TODO()
	plan := resourceModel{
		ClusterID:                    types.StringValue("cluster-id"),
		DatabaseCatalogID:            types.StringValue("database-catalog-id"),
		Name:                         types.StringValue("test-name"),
		ImageVersion:                 types.StringValue("2024.0.19.0-301"),
		GroupSize:                    types.Int64Value(10),
		PlatformJwtAuth:              types.BoolValue(true),
		LdapGroups:                   types.ListValueMust(types.StringType, []attr.Value{types.StringValue("ldap-group")}),
		EnableSSO:                    types.BoolValue(true),
		MinGroupCount:                types.Int64Value(1),
		MaxGroupCount:                types.Int64Value(10),
		DisableAutoSuspend:           types.BoolValue(false),
		AutoSuspendTimeoutSeconds:    types.Int64Value(60),
		Headroom:                     types.Int64Value(2),
		ScaleWaitTimeSeconds:         types.Int64Value(60),
		MaxConcurrentIsolatedQueries: types.Int64Value(5),
		MaxNodesPerIsolatedQuery:     types.Int64Value(2),
		AwsOptions: &awsOptions{
			AvailabilityZone: types.StringValue("us-west-2a"),
			EbsLLAPSpillGb:   types.Int64Value(300),
			Tags:             types.MapValueMust(types.StringType, map[string]attr.Value{"key1": types.StringValue("value1")}),
		},
	}

	req, diags := plan.convertToCreateVwRequest(ctx)
	suite.False(diags.HasError())
	suite.Equal("cluster-id", *req.ClusterID)
	suite.Equal("database-catalog-id", *req.DbcID)
	suite.Equal("test-name", *req.Name)
	suite.Equal("2024.0.19.0-301", req.ImageVersion)
	suite.Equal(int32(10), req.NodeCount) // group_size
	suite.Equal(true, *req.PlatformJwtAuth)
	suite.Equal([]string{"ldap-group"}, req.Config.LdapGroups)
	suite.Equal(true, req.Config.EnableSSO)
	suite.Equal(int32(1), *req.Autoscaling.MinClusters)  // min_group_size
	suite.Equal(int32(10), *req.Autoscaling.MaxClusters) // max_group_size
	suite.Equal(false, req.Autoscaling.DisableAutoSuspend)
	suite.Equal(int32(60), req.Autoscaling.AutoSuspendTimeoutSeconds)
	suite.Equal(int32(60), req.Autoscaling.HiveScaleWaitTimeSeconds) // scale_wait_time_seconds
	suite.Equal(int32(2), req.Autoscaling.HiveDesiredFreeCapacity)   // headroom
	suite.Equal("us-west-2a", req.AvailabilityZone)
	suite.Equal(int32(300), req.EbsLLAPSpillGB)
	suite.Equal("key1", *req.Tags[0].Key)
	suite.Equal("value1", *req.Tags[0].Value)
	suite.Equal(int32(5), req.QueryIsolationOptions.MaxQueries)
	suite.Equal(int32(2), req.QueryIsolationOptions.MaxNodesPerQuery)
}

func (suite *HiveTestSuite) TestConvertToCreateVwRequest_MissingImageVersion() {
	ctx := context.TODO()
	plan := resourceModel{
		ImageVersion: types.StringUnknown(),
		AwsOptions:   &awsOptions{},
	}

	req, diags := plan.convertToCreateVwRequest(ctx)
	suite.False(diags.HasError())
	suite.Equal("", req.ImageVersion)
}

func (suite *HiveTestSuite) TestSetFromDescribeVwResponse() {
	plan := resourceModel{}
	resp := &models.DescribeVwResponse{
		Vw: &models.VwSummary{
			ID:     "test-id",
			DbcID:  "database-catalog-id",
			Name:   "test-name",
			Status: "Running",
			Endpoints: &models.VwSummaryEndpoints{
				HiveJdbc:            "jdbc://hive",
				HiveKerberosJdbc:    "jdbc://hive",
				Hue:                 "https://hue",
				JwtConnectionString: "connection-string",
				JwtTokenGenURL:      "https://jwt",
			},
		},
	}

	plan.setFromDescribeVwResponse(resp)
	suite.Equal("test-id", plan.ID.ValueString())
	suite.Equal("database-catalog-id", plan.DatabaseCatalogID.ValueString())
	suite.Equal("test-name", plan.Name.ValueString())
	suite.Equal("Running", plan.Status.ValueString())
	suite.Equal("jdbc://hive", plan.JdbcUrl.ValueString())
	suite.Equal("jdbc://hive", plan.KerberosJdbcUrl.ValueString())
	suite.Equal("https://hue", plan.HueUrl.ValueString())
	suite.Equal("connection-string", plan.JwtConnectionString.ValueString())
	suite.Equal("https://jwt", plan.JwtTokenGenUrl.ValueString())
}
