// Copyright 2025 Cloudera. All Rights Reserved.
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

var awsSchemaElements = []test.ResourceSchemaTestCaseStructure{
	{
		Name:             "id field must exist and be valid",
		Field:            "id",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "crn should exist and be valid",
		Field:            "crn",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "polling_options should exist and be valid",
		Field:            "polling_options",
		Computed:         false,
		ShouldBeRequired: false,
		AttributeType:    schema.SingleNestedAttribute{},
	},
	{
		Name:             "authentication should exist and be valid",
		Field:            "authentication",
		Computed:         false,
		ShouldBeRequired: true,
		AttributeType:    schema.SingleNestedAttribute{},
	},
	{
		Name:             "create_private_subnets should exist and be valid",
		Field:            "create_private_subnets",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.BoolAttribute{},
	},
	{
		Name:             "create_service_endpoints should exist and be valid",
		Field:            "create_service_endpoints",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.BoolAttribute{},
	},
	{
		Name:             "s3_guard_table_name should exist and be valid",
		Field:            "s3_guard_table_name",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "credential_name should exist and be valid",
		Field:            "credential_name",
		Computed:         false,
		ShouldBeRequired: true,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "description should exist and be valid",
		Field:            "description",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "enable_tunnel should exist and be valid",
		Field:            "enable_tunnel",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.BoolAttribute{},
	},
	{
		Name:             "encryption_key_arn should exist and be valid",
		Field:            "encryption_key_arn",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "endpoint_access_gateway_scheme should exist and be valid",
		Field:            "endpoint_access_gateway_scheme",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "endpoint_access_gateway_subnet_ids should exist and be valid",
		Field:            "endpoint_access_gateway_subnet_ids",
		Computed:         false,
		ShouldBeRequired: false,
		AttributeType:    schema.SetAttribute{},
	},
	{
		Name:             "environment_name should exist and be valid",
		Field:            "environment_name",
		Computed:         false,
		ShouldBeRequired: true,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "delete_options should exist and be valid",
		Field:            "delete_options",
		Computed:         false,
		ShouldBeRequired: false,
		AttributeType:    schema.SingleNestedAttribute{},
	},
	{
		Name:             "cascading_delete should exist and be valid",
		Field:            "cascading_delete",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.BoolAttribute{},
	},
	{
		Name:             "freeipa should exist and be valid",
		Field:            "freeipa",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.SingleNestedAttribute{},
	},
	{
		Name:             "log_storage should exist and be valid",
		Field:            "log_storage",
		Computed:         false,
		ShouldBeRequired: true,
		AttributeType:    schema.SingleNestedAttribute{},
	},
	{
		Name:             "region should exist and be valid",
		Field:            "region",
		Computed:         false,
		ShouldBeRequired: true,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "report_deployment_logs should exist and be valid",
		Field:            "report_deployment_logs",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.BoolAttribute{},
	},
	{
		Name:             "proxy_config_name should exist and be valid",
		Field:            "proxy_config_name",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "security_access should exist and be valid",
		Field:            "security_access",
		Computed:         false,
		ShouldBeRequired: true,
		AttributeType:    schema.SingleNestedAttribute{},
	},
	{
		Name:             "status should exist and be valid",
		Field:            "status",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "status_reason should exist and be valid",
		Field:            "status_reason",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "subnet_ids should exist and be valid",
		Field:            "subnet_ids",
		Computed:         false,
		ShouldBeRequired: true,
		AttributeType:    schema.SetAttribute{},
	},
	{
		Name:             "tags should exist and be valid",
		Field:            "tags",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.MapAttribute{},
	},
	{
		Name:             "tunnel_type should exist and be valid",
		Field:            "tunnel_type",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "workload_analytics should exist and be valid",
		Field:            "workload_analytics",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.BoolAttribute{},
	},
	{
		Name:             "vpc_id should exist and be valid",
		Field:            "vpc_id",
		Computed:         false,
		ShouldBeRequired: true,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "compute_cluster should exist and be valid",
		Field:            "compute_cluster",
		Computed:         false,
		ShouldBeRequired: false,
		AttributeType:    schema.SingleNestedAttribute{},
	},
}

func TestAwsEnvironmentSchemaAttributeNumber(t *testing.T) {
	expected := 30
	if len(AwsEnvironmentSchema.Attributes) != expected {
		t.Errorf("The number of fields in the AwsEnvironmentSchema schema should be: %d but it is: %d", expected, len(AwsEnvironmentSchema.Attributes))
		t.FailNow()
	}
}

func TestAwsEnvironmentSchema(t *testing.T) {
	test.PerformResourceSchemaValidation(t, AwsEnvironmentSchema.Attributes, awsSchemaElements)
}
