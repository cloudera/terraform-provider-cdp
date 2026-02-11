// Copyright 2024 Cloudera. All Rights Reserved.
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
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/cloudera/terraform-provider-cdp/utils/test"
)

var azureSchemaElements = []test.ResourceSchemaTestCaseStructure{
	{
		Name:          "id field must exist and be valid",
		Field:         "id",
		Computed:      true,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "polling_options should exist and be valid",
		Field:         "polling_options",
		Computed:      false,
		Required:      false,
		AttributeType: schema.SingleNestedAttribute{},
	},
	{
		Name:          "crn should exist and be valid",
		Field:         "crn",
		Computed:      true,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "create_private_endpoints should exist and be valid",
		Field:         "create_private_endpoints",
		Computed:      false,
		Required:      false,
		AttributeType: schema.BoolAttribute{},
	},
	{
		Name:          "credential_name must exist and be valid",
		Field:         "credential_name",
		Computed:      false,
		Required:      true,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "description should exist and be valid",
		Field:         "description",
		Computed:      true,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "enable_outbound_load_balancer should exist and be valid",
		Field:         "enable_outbound_load_balancer",
		Computed:      false,
		Required:      false,
		AttributeType: schema.BoolAttribute{},
	},
	{
		Name:          "enable_tunnel should exist and be valid",
		Field:         "enable_tunnel",
		Computed:      true,
		Required:      false,
		AttributeType: schema.BoolAttribute{},
	},
	{
		Name:          "endpoint_access_gateway_scheme should exist and be valid",
		Field:         "endpoint_access_gateway_scheme",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "endpoint_access_gateway_subnet_ids should exist and be valid",
		Field:         "endpoint_access_gateway_subnet_ids",
		Computed:      false,
		Required:      false,
		AttributeType: schema.SetAttribute{},
	},
	{
		Name:          "encryption_key_resource_group_name should exist and be valid",
		Field:         "encryption_key_resource_group_name",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "encryption_key_url should exist and be valid",
		Field:         "encryption_key_url",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "encryption_at_host should exist and be valid",
		Field:         "encryption_at_host",
		Computed:      true,
		Required:      false,
		AttributeType: schema.BoolAttribute{},
	},
	{
		Name:          "environment_name should exist and be valid",
		Field:         "environment_name",
		Computed:      false,
		Required:      true,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "existing_network_params should exist and be valid",
		Field:         "existing_network_params",
		Computed:      false,
		Required:      true,
		AttributeType: schema.SingleNestedAttribute{},
	},
	{
		Name:          "freeipa should exist and be valid",
		Field:         "freeipa",
		Computed:      true,
		Required:      false,
		AttributeType: schema.SingleNestedAttribute{},
	},
	{
		Name:          "log_storage should exist and be valid",
		Field:         "log_storage",
		Computed:      false,
		Required:      true,
		AttributeType: schema.SingleNestedAttribute{},
	},
	{
		Name:          "proxy_config_name should exist and be valid",
		Field:         "proxy_config_name",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "public_key should exist and be valid",
		Field:         "public_key",
		Computed:      false,
		Required:      true,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "region should exist and be valid",
		Field:         "region",
		Computed:      false,
		Required:      true,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "report_deployment_logs should exist and be valid",
		Field:         "report_deployment_logs",
		Computed:      true,
		Required:      false,
		AttributeType: schema.BoolAttribute{},
	},
	{
		Name:          "resource_group_name should exist and be valid",
		Field:         "resource_group_name",
		Computed:      true,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "security_access should exist and be valid",
		Field:         "security_access",
		Computed:      false,
		Required:      true,
		AttributeType: schema.SingleNestedAttribute{},
	},
	{
		Name:          "status should exist and be valid",
		Field:         "status",
		Computed:      true,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "status_reason should exist and be valid",
		Field:         "status_reason",
		Computed:      true,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "encryption_user_managed_identity should exist and be valid",
		Field:         "encryption_user_managed_identity",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "tags should exist and be valid",
		Field:         "tags",
		Computed:      true,
		Required:      false,
		AttributeType: schema.MapAttribute{},
	},
	{
		Name:          "use_public_ip should exist and be valid",
		Field:         "use_public_ip",
		Computed:      false,
		Required:      true,
		AttributeType: schema.BoolAttribute{},
	},
	{
		Name:          "workload_analytics should exist and be valid",
		Field:         "workload_analytics",
		Computed:      true,
		Required:      false,
		AttributeType: schema.BoolAttribute{},
	},
	{
		Name:          "compute_cluster should exist and be valid",
		Field:         "compute_cluster",
		Computed:      false,
		Required:      false,
		AttributeType: schema.SingleNestedAttribute{},
	},
	{
		Name:          "delete_options should exist and be valid",
		Field:         "delete_options",
		Computed:      false,
		Required:      false,
		AttributeType: schema.SingleNestedAttribute{},
	},
	{
		Name:          "cascading_delete should exist and be valid",
		Field:         "cascading_delete",
		Computed:      true,
		Required:      false,
		AttributeType: schema.BoolAttribute{},
	},
	{
		Name:          "custom_docker_registry should exist and be valid",
		Field:         "custom_docker_registry",
		Computed:      false,
		Required:      false,
		AttributeType: schema.SingleNestedAttribute{},
	},
	{
		Name:          "security should exist and be valid",
		Field:         "security",
		Computed:      false,
		Required:      false,
		AttributeType: schema.SingleNestedAttribute{},
	},
	{
		Name:          "data_services should exist and be valid",
		Field:         "data_services",
		Computed:      false,
		Required:      false,
		AttributeType: schema.SingleNestedAttribute{},
	},
	{
		Name:          "environment_type should exist and be valid",
		Field:         "environment_type",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "flexible_server_subnet_ids should exist and be valid",
		Field:         "flexible_server_subnet_ids",
		Computed:      false,
		Required:      false,
		AttributeType: schema.SetAttribute{},
	},
	{
		Name:          "availability_zones should exist and be valid",
		Field:         "availability_zones",
		Computed:      false,
		Required:      false,
		AttributeType: schema.SetAttribute{},
	},
}

func TestAzureEnvironmentSchemaSchemaAttributeNumber(t *testing.T) {
	expected := 38
	if len(AzureEnvironmentSchema.Attributes) != expected {
		t.Errorf("The number of fields in the AzureEnvironment schema should be: %d but it is: %d", expected, len(AzureEnvironmentSchema.Attributes))
		t.FailNow()
	}
}

func TestAzureSchemaContainsElements(t *testing.T) {
	test.PerformResourceSchemaValidation(t, AzureEnvironmentSchema.Attributes, azureSchemaElements)
}
