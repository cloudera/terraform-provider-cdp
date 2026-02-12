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
	"context"
	"errors"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/mocks"
)

func createRawAzureEnvironmentResource() tftypes.Value {
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"id":  tftypes.String,
				"crn": tftypes.String,
				"polling_options": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"async":                  tftypes.Bool,
						"polling_timeout":        tftypes.Number,
						"call_failure_threshold": tftypes.Number,
					},
				},
				"status_reason": tftypes.String,
				"tags": tftypes.Map{
					ElementType: tftypes.String,
				},
				"use_public_ip":                      tftypes.Bool,
				"credential_name":                    tftypes.String,
				"description":                        tftypes.String,
				"encryption_key_resource_group_name": tftypes.String,
				"encryption_at_host":                 tftypes.Bool,
				"public_key":                         tftypes.String,
				"status":                             tftypes.String,
				"enable_tunnel":                      tftypes.Bool,
				"log_storage": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"managed_identity":             tftypes.String,
						"storage_location_base":        tftypes.String,
						"backup_storage_location_base": tftypes.String,
					},
				},
				"report_deployment_logs": tftypes.Bool,
				"security_access": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"cidr":                      tftypes.String,
						"default_security_group_id": tftypes.String,
						"default_security_group_ids": tftypes.Set{
							ElementType: tftypes.String,
						},
						"security_group_id_for_knox": tftypes.String,
						"security_group_ids_for_knox": tftypes.Set{
							ElementType: tftypes.String,
						},
					},
				},
				"workload_analytics":       tftypes.Bool,
				"encryption_key_url":       tftypes.String,
				"freeipa":                  FreeIpaDetailsObject,
				"compute_cluster":          AzureComputeClusterObject,
				"region":                   tftypes.String,
				"resource_group_name":      tftypes.String,
				"create_private_endpoints": tftypes.Bool,
				"environment_name":         tftypes.String,
				"cascading_delete":         tftypes.Bool,
				"delete_options": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"cascading": tftypes.Bool,
						"forced":    tftypes.Bool,
					},
				},
				"proxy_config_name":                tftypes.String,
				"endpoint_access_gateway_scheme":   tftypes.String,
				"enable_outbound_load_balancer":    tftypes.Bool,
				"encryption_user_managed_identity": tftypes.String,
				"security": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"se_linux": tftypes.String,
					},
				},
				"custom_docker_registry": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"crn": tftypes.String,
					},
				},
				"environment_type": tftypes.String,
				"data_services": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"shared_managed_identity": tftypes.String,
						"aks_private_dns_zone_id": tftypes.String,
					},
				},
				"flexible_server_subnet_ids": tftypes.Set{
					ElementType: tftypes.String,
				},
				"availability_zones": tftypes.Set{
					ElementType: tftypes.String,
				},
				"existing_network_params": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"aks_private_dns_zone_id":      tftypes.String,
						"database_private_dns_zone_id": tftypes.String,
						"network_id":                   tftypes.String,
						"resource_group_name":          tftypes.String,
						"subnet_ids": tftypes.Set{
							ElementType: tftypes.String,
						},
						"flexible_server_subnet_ids": tftypes.Set{
							ElementType: tftypes.String,
						},
					},
				},
				"endpoint_access_gateway_subnet_ids": tftypes.Set{
					ElementType: tftypes.String,
				},
			},
		},
		map[string]tftypes.Value{
			"id":  tftypes.NewValue(tftypes.String, ""),
			"crn": tftypes.NewValue(tftypes.String, ""),
			"polling_options": tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"async":                  tftypes.Bool,
					"polling_timeout":        tftypes.Number,
					"call_failure_threshold": tftypes.Number,
				},
			}, map[string]tftypes.Value{
				"async":                  tftypes.NewValue(tftypes.Bool, false),
				"polling_timeout":        tftypes.NewValue(tftypes.Number, 100),
				"call_failure_threshold": tftypes.NewValue(tftypes.Number, 3),
			}),
			"status_reason": tftypes.NewValue(tftypes.String, ""),
			"tags": tftypes.NewValue(tftypes.Map{
				ElementType: tftypes.String,
			}, map[string]tftypes.Value{}),
			"use_public_ip":                      tftypes.NewValue(tftypes.Bool, false),
			"credential_name":                    tftypes.NewValue(tftypes.String, ""),
			"description":                        tftypes.NewValue(tftypes.String, ""),
			"encryption_key_resource_group_name": tftypes.NewValue(tftypes.String, ""),
			"encryption_at_host":                 tftypes.NewValue(tftypes.Bool, false),
			"public_key":                         tftypes.NewValue(tftypes.String, ""),
			"status":                             tftypes.NewValue(tftypes.String, ""),
			"enable_tunnel":                      tftypes.NewValue(tftypes.Bool, false),
			"encryption_user_managed_identity":   tftypes.NewValue(tftypes.String, ""),
			"log_storage": tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"managed_identity":             tftypes.String,
					"storage_location_base":        tftypes.String,
					"backup_storage_location_base": tftypes.String,
				},
			}, map[string]tftypes.Value{
				"managed_identity":             tftypes.NewValue(tftypes.String, ""),
				"storage_location_base":        tftypes.NewValue(tftypes.String, ""),
				"backup_storage_location_base": tftypes.NewValue(tftypes.String, ""),
			}),
			"report_deployment_logs": tftypes.NewValue(tftypes.Bool, false),
			"security_access": tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"cidr":                      tftypes.String,
					"default_security_group_id": tftypes.String,
					"default_security_group_ids": tftypes.Set{
						ElementType: tftypes.String,
					},
					"security_group_id_for_knox": tftypes.String,
					"security_group_ids_for_knox": tftypes.Set{
						ElementType: tftypes.String,
					},
				},
			}, map[string]tftypes.Value{
				"cidr":                      tftypes.NewValue(tftypes.String, ""),
				"default_security_group_id": tftypes.NewValue(tftypes.String, ""),
				"default_security_group_ids": tftypes.NewValue(tftypes.Set{
					ElementType: tftypes.String,
				}, []tftypes.Value{}),
				"security_group_id_for_knox": tftypes.NewValue(tftypes.String, ""),
				"security_group_ids_for_knox": tftypes.NewValue(tftypes.Set{
					ElementType: tftypes.String,
				}, []tftypes.Value{}),
			}),
			"workload_analytics": tftypes.NewValue(tftypes.Bool, false),
			"encryption_key_url": tftypes.NewValue(tftypes.String, ""),
			"compute_cluster": tftypes.NewValue(AzureComputeClusterObject, map[string]tftypes.Value{
				"enabled": tftypes.NewValue(tftypes.Bool, false),
				"configuration": tftypes.NewValue(AzureComputeClusterConfigurationObject, map[string]tftypes.Value{
					"private_cluster": tftypes.NewValue(tftypes.Bool, false),
					"kube_api_authorized_ip_ranges": tftypes.NewValue(tftypes.Set{
						ElementType: tftypes.String,
					}, []tftypes.Value{}),
					"worker_node_subnets": tftypes.NewValue(tftypes.Set{
						ElementType: tftypes.String,
					}, []tftypes.Value{}),
					"outbound_type": tftypes.NewValue(tftypes.String, ""),
				}),
			}),
			"freeipa": tftypes.NewValue(FreeIpaDetailsObject, map[string]tftypes.Value{
				"catalog":                 tftypes.NewValue(tftypes.String, ""),
				"architecture":            tftypes.NewValue(tftypes.String, ""),
				"image_id":                tftypes.NewValue(tftypes.String, ""),
				"os":                      tftypes.NewValue(tftypes.String, ""),
				"instance_count_by_group": tftypes.NewValue(tftypes.Number, 0),
				"instance_type":           tftypes.NewValue(tftypes.String, ""),
				"instances": tftypes.NewValue(tftypes.Set{
					ElementType: FreeIpaInstanceObject,
				}, []tftypes.Value{}),
				"multi_az": tftypes.NewValue(tftypes.Bool, false),
				"recipes": tftypes.NewValue(tftypes.Set{
					ElementType: tftypes.String,
				}, []tftypes.Value{}),
			}),
			"region":                   tftypes.NewValue(tftypes.String, ""),
			"resource_group_name":      tftypes.NewValue(tftypes.String, ""),
			"create_private_endpoints": tftypes.NewValue(tftypes.Bool, false),
			"environment_name":         tftypes.NewValue(tftypes.String, ""),
			"cascading_delete":         tftypes.NewValue(tftypes.Bool, true),
			"delete_options": tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"cascading": tftypes.Bool,
					"forced":    tftypes.Bool,
				},
			}, map[string]tftypes.Value{
				"cascading": tftypes.NewValue(tftypes.Bool, true),
				"forced":    tftypes.NewValue(tftypes.Bool, false),
			}),
			"proxy_config_name":              tftypes.NewValue(tftypes.String, ""),
			"endpoint_access_gateway_scheme": tftypes.NewValue(tftypes.String, ""),
			"enable_outbound_load_balancer":  tftypes.NewValue(tftypes.Bool, false),
			"existing_network_params": tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"aks_private_dns_zone_id":      tftypes.String,
					"database_private_dns_zone_id": tftypes.String,
					"network_id":                   tftypes.String,
					"resource_group_name":          tftypes.String,
					"subnet_ids": tftypes.Set{
						ElementType: tftypes.String,
					},
					"flexible_server_subnet_ids": tftypes.Set{
						ElementType: tftypes.String,
					},
				},
			}, map[string]tftypes.Value{
				"aks_private_dns_zone_id":      tftypes.NewValue(tftypes.String, ""),
				"database_private_dns_zone_id": tftypes.NewValue(tftypes.String, ""),
				"network_id":                   tftypes.NewValue(tftypes.String, "some-network"),
				"resource_group_name":          tftypes.NewValue(tftypes.String, "some-resource-group"),
				"subnet_ids": tftypes.NewValue(tftypes.Set{
					ElementType: tftypes.String,
				}, []tftypes.Value{tftypes.NewValue(tftypes.String, "some-subnet-1"), tftypes.NewValue(tftypes.String, "some-subnet-2")}),
				"flexible_server_subnet_ids": tftypes.NewValue(tftypes.Set{
					ElementType: tftypes.String,
				}, []tftypes.Value{}),
			}),
			"security": tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"se_linux": tftypes.String,
				},
			}, map[string]tftypes.Value{
				"se_linux": tftypes.NewValue(tftypes.String, ""),
			}),
			"custom_docker_registry": tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"crn": tftypes.String,
				},
			}, map[string]tftypes.Value{
				"crn": tftypes.NewValue(tftypes.String, ""),
			}),
			"data_services": tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"shared_managed_identity": tftypes.String,
					"aks_private_dns_zone_id": tftypes.String,
				},
			}, map[string]tftypes.Value{
				"shared_managed_identity": tftypes.NewValue(tftypes.String, ""),
				"aks_private_dns_zone_id": tftypes.NewValue(tftypes.String, ""),
			}),
			"flexible_server_subnet_ids": tftypes.NewValue(tftypes.Set{
				ElementType: tftypes.String,
			}, []tftypes.Value{}),
			"environment_type": tftypes.NewValue(tftypes.String, ""),
			"availability_zones": tftypes.NewValue(tftypes.Set{
				ElementType: tftypes.String,
			}, []tftypes.Value{}),
			"endpoint_access_gateway_subnet_ids": tftypes.NewValue(tftypes.Set{
				ElementType: tftypes.String,
			}, []tftypes.Value{}),
		},
	)
}

var AzureComputeClusterObject = tftypes.Object{
	AttributeTypes: map[string]tftypes.Type{
		"enabled":       tftypes.Bool,
		"configuration": AzureComputeClusterConfigurationObject,
	},
}

var AzureComputeClusterConfigurationObject = tftypes.Object{
	AttributeTypes: map[string]tftypes.Type{
		"private_cluster": tftypes.Bool,
		"kube_api_authorized_ip_ranges": tftypes.Set{
			ElementType: tftypes.String,
		},
		"outbound_type": tftypes.String,
		"worker_node_subnets": tftypes.Set{
			ElementType: tftypes.String,
		},
	},
}

func TestCreateAzureEnvironment(t *testing.T) {
	testCases := map[string]struct {
		expectedResponse         interface{}
		expectedDescribeResponse interface{}
		expectedErrorResponse    interface{}
		expectedError            bool
		expectedSummary          string
		expectedDetail           string
		expectedAccepted         bool
	}{
		"OK": {
			expectedResponse: &operations.CreateAzureEnvironmentOK{
				Payload: &models.CreateAzureEnvironmentResponse{
					Environment: &models.Environment{
						Authentication:  &models.Authentication{},
						BackupStorage:   &models.BackupStorage{},
						CloudPlatform:   func(s string) *string { return &s }(""),
						CredentialName:  func(s string) *string { return &s }(""),
						Crn:             func(s string) *string { return &s }(""),
						DataServices:    &models.DataServices{},
						EnvironmentName: func(s string) *string { return &s }(""),
						Freeipa:         &models.FreeipaDetails{},
						LogStorage:      &models.LogStorage{},
						Network: &models.Network{
							SubnetIds: []string{"subnet-1", "subnet-2"},
							Azure: &models.NetworkAzureParams{
								ResourceGroupName: func(s string) *string { return &s }("some-resource-group"),
							},
						},
						ProxyConfig:    &models.ProxyConfig{},
						Region:         func(s string) *string { return &s }(""),
						SecurityAccess: &models.SecurityAccess{},
						Status:         func(s string) *string { return &s }(""),
						Tags:           &models.EnvironmentTags{},
					},
				},
			},
			expectedDescribeResponse: &operations.DescribeEnvironmentOK{
				Payload: &models.DescribeEnvironmentResponse{
					Environment: &models.Environment{
						Authentication:  &models.Authentication{},
						BackupStorage:   &models.BackupStorage{},
						CloudPlatform:   func(s string) *string { return &s }(""),
						CredentialName:  func(s string) *string { return &s }(""),
						Crn:             func(s string) *string { return &s }(""),
						DataServices:    &models.DataServices{},
						EnvironmentName: func(s string) *string { return &s }(""),
						Freeipa:         &models.FreeipaDetails{},
						LogStorage:      &models.LogStorage{},
						Network: &models.Network{
							SubnetIds: []string{"subnet-1", "subnet-2"},
							Azure: &models.NetworkAzureParams{
								ResourceGroupName: func(s string) *string { return &s }("some-resource-group"),
							},
						},
						ProxyConfig:    &models.ProxyConfig{},
						Region:         func(s string) *string { return &s }(""),
						SecurityAccess: &models.SecurityAccess{},
						Status:         func(s string) *string { return &s }("AVAILABLE"),
						Tags:           &models.EnvironmentTags{},
					},
				},
			},
			expectedErrorResponse: nil,
			expectedError:         false,
			expectedSummary:       "",
			expectedDetail:        "",
			expectedAccepted:      true,
		},
		"TransportError": {
			expectedResponse:      nil,
			expectedErrorResponse: errors.New("request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)"),
			expectedError:         true,
			expectedSummary:       "Create Azure Environment",
			expectedDetail:        "Failed to create Azure Environment, unexpected error: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)",
			expectedAccepted:      false,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()

			mockClient := new(mocks.MockEnvironmentClientService)

			createMatcher := func(params *operations.CreateAzureEnvironmentParams) bool {
				return true
			}
			mockClient.On("CreateAzureEnvironment", mock.MatchedBy(createMatcher)).Return(testCase.expectedResponse, testCase.expectedErrorResponse)

			if !testCase.expectedError {
				describeMatcher := func(params *operations.DescribeEnvironmentParams) bool {
					return true
				}
				mockClient.On("DescribeEnvironment", mock.MatchedBy(describeMatcher)).Return(testCase.expectedDescribeResponse, testCase.expectedErrorResponse)
			}
			aitpResource := &azureEnvironmentResource{
				client: &cdp.Client{Environments: NewMockEnvironments(mockClient)},
			}

			req := resource.CreateRequest{
				Plan: tfsdk.Plan{
					Raw:    createRawAzureEnvironmentResource(),
					Schema: AzureEnvironmentSchema,
				},
			}

			resp := &resource.CreateResponse{
				State: tfsdk.State{
					Schema: AzureEnvironmentSchema,
				},
			}

			aitpResource.Create(ctx, req, resp)

			assert.Equal(t, testCase.expectedError, resp.Diagnostics.HasError())
			if testCase.expectedError {
				assert.Equal(t, testCase.expectedSummary, resp.Diagnostics.Errors()[0].Summary())
				assert.Equal(t, testCase.expectedDetail, resp.Diagnostics.Errors()[0].Detail())
			}

			var state azureEnvironmentResourceModel
			resp.State.Get(ctx, &state)

			mockClient.AssertExpectations(t)
		})
	}
}

func TestReadAzureEnvironment(t *testing.T) {
	testCases := map[string]struct {
		expectedResponse      interface{}
		expectedErrorResponse interface{}
		expectedError         bool
		expectedWarning       bool
		expectedSummary       string
		expectedDetail        string
		expectedAccepted      bool
	}{
		"OK": {
			expectedResponse: &operations.DescribeEnvironmentOK{
				Payload: &models.DescribeEnvironmentResponse{
					Environment: &models.Environment{
						Authentication:  &models.Authentication{},
						BackupStorage:   &models.BackupStorage{},
						CloudPlatform:   func(s string) *string { return &s }(""),
						CredentialName:  func(s string) *string { return &s }(""),
						Crn:             func(s string) *string { return &s }(""),
						DataServices:    &models.DataServices{},
						EnvironmentName: func(s string) *string { return &s }(""),
						Freeipa:         &models.FreeipaDetails{},
						LogStorage:      &models.LogStorage{},
						Network: &models.Network{
							SubnetIds: []string{"subnet-1", "subnet-2"},
							Azure: &models.NetworkAzureParams{
								ResourceGroupName: func(s string) *string { return &s }("some-resource-group"),
							},
						},
						ProxyConfig:    &models.ProxyConfig{},
						Region:         func(s string) *string { return &s }(""),
						SecurityAccess: &models.SecurityAccess{},
						Status:         func(s string) *string { return &s }("AVAILABLE"),
						Tags:           &models.EnvironmentTags{},
					},
				},
			},
			expectedErrorResponse: nil,
			expectedError:         false,
			expectedSummary:       "",
			expectedDetail:        "",
			expectedAccepted:      true,
		},
		"TransportError": {
			expectedResponse:      nil,
			expectedErrorResponse: errors.New("request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)"),
			expectedError:         true,
			expectedSummary:       "Read Environment",
			expectedDetail:        "Failed to read Environment, unexpected error: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)",
			expectedAccepted:      false,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()

			mockClient := new(mocks.MockEnvironmentClientService)
			readMatcher := func(params *operations.DescribeEnvironmentParams) bool {
				return true
			}
			mockClient.On("DescribeEnvironment", mock.MatchedBy(readMatcher)).Return(testCase.expectedResponse, testCase.expectedErrorResponse)

			aitpResource := &azureEnvironmentResource{
				client: &cdp.Client{Environments: NewMockEnvironments(mockClient)},
			}

			req := resource.ReadRequest{
				State: tfsdk.State{
					Raw:    createRawAzureEnvironmentResource(),
					Schema: AzureEnvironmentSchema,
				},
			}

			resp := &resource.ReadResponse{
				State: tfsdk.State{
					Schema: AzureEnvironmentSchema,
				},
			}

			aitpResource.Read(ctx, req, resp)

			assert.Equal(t, testCase.expectedError, resp.Diagnostics.HasError())
			if testCase.expectedError {
				assert.Equal(t, testCase.expectedSummary, resp.Diagnostics.Errors()[0].Summary())
				assert.Equal(t, testCase.expectedDetail, resp.Diagnostics.Errors()[0].Detail())
			}

			if testCase.expectedWarning {
				assert.Equal(t, 1, resp.Diagnostics.WarningsCount())
				assert.Equal(t, testCase.expectedSummary, resp.Diagnostics.Warnings()[0].Summary())
				assert.Equal(t, testCase.expectedDetail, resp.Diagnostics.Warnings()[0].Detail())
				assert.True(t, resp.State.Raw.IsNull())
			}

			var state azureEnvironmentResourceModel
			resp.State.Get(ctx, &state)

			mockClient.AssertExpectations(t)
		})
	}
}

func TestDeleteAzureEnvironmentPolicy(t *testing.T) {
	testCases := map[string]struct {
		expectedResponse              interface{}
		expectedErrorResponse         interface{}
		expectedDescribeResponse      interface{}
		expectedDescribeErrorResponse interface{}
		expectedError                 bool
		expectedWarning               bool
		expectedSummary               string
		expectedDetail                string
		expectedAccepted              bool
	}{
		"OK": {
			expectedResponse:         &operations.DeleteEnvironmentOK{},
			expectedDescribeResponse: &operations.DescribeEnvironmentOK{},
			expectedErrorResponse:    nil,
			expectedDescribeErrorResponse: func() *operations.DescribeEnvironmentDefault {
				resp := operations.NewDescribeEnvironmentDefault(404)
				resp.Payload = &models.Error{
					Code:    "NOT_FOUND",
					Message: "",
				}
				return resp
			}(),
			expectedError:    false,
			expectedSummary:  "",
			expectedDetail:   "",
			expectedAccepted: false,
		},
		"TransportError": {
			expectedResponse:      nil,
			expectedErrorResponse: errors.New("request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)"),
			expectedError:         true,
			expectedSummary:       "Delete Environment",
			expectedDetail:        "Failed to delete Environment, unexpected error: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)",
			expectedAccepted:      false,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()

			mockClient := new(mocks.MockEnvironmentClientService)
			delMatcher := func(params *operations.DeleteEnvironmentParams) bool {
				return true
			}
			mockClient.On("DeleteEnvironment", mock.MatchedBy(delMatcher)).Return(testCase.expectedResponse, testCase.expectedErrorResponse)
			if !testCase.expectedError {
				descMatcher := func(params *operations.DescribeEnvironmentParams) bool {
					return true
				}
				mockClient.On("DescribeEnvironment", mock.MatchedBy(descMatcher)).Return(testCase.expectedDescribeResponse, testCase.expectedDescribeErrorResponse)
			}

			aitpResource := &azureEnvironmentResource{
				client: &cdp.Client{Environments: NewMockEnvironments(mockClient)},
			}

			req := resource.DeleteRequest{
				State: tfsdk.State{
					Raw:    createRawAzureEnvironmentResource(),
					Schema: AzureEnvironmentSchema,
				},
			}

			resp := &resource.DeleteResponse{
				State: tfsdk.State{
					Schema: AzureEnvironmentSchema,
				},
			}

			aitpResource.Delete(ctx, req, resp)

			assert.Equal(t, testCase.expectedError, resp.Diagnostics.HasError())
			if testCase.expectedError {
				assert.Equal(t, testCase.expectedSummary, resp.Diagnostics.Errors()[0].Summary())
				assert.Equal(t, testCase.expectedDetail, resp.Diagnostics.Errors()[0].Detail())
			}

			if testCase.expectedWarning {
				assert.Equal(t, 1, resp.Diagnostics.WarningsCount())
				assert.Equal(t, testCase.expectedSummary, resp.Diagnostics.Warnings()[0].Summary())
				assert.Equal(t, testCase.expectedDetail, resp.Diagnostics.Warnings()[0].Detail())
			}

			var state azureEnvironmentResourceModel
			resp.State.Get(ctx, &state)

			assert.True(t, resp.State.Raw.IsNull())

			mockClient.AssertExpectations(t)
		})
	}
}
